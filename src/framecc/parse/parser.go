package parse

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"framecc/ast"
)

var frameSizes = map[string]ast.FrameSize{
	"Fixed": ast.SzFixed,
	"Var8":  ast.SzVar8,
	"Var16": ast.SzVar16,
}

var intSizes = map[string]int{
	"int8":  8,
	"int16": 16,
	"int32": 32,
	"int64": 64,
}

var intFlags = map[string]ast.IntegerFlag{
	"LittleEndian": ast.IntLittleEndian,
	"PDPEndian":    ast.IntPDPEndian,
	"RPDPEndian":   ast.IntRPDPEndian,
	"Negate":       ast.IntNegate,
	"Offset128":    ast.IntOffset128,
	"Inverse128":   ast.IntInverse128,
}

type parseContext struct {
	root       *ast.File
	errors     errorList
	itemBuffer []item
	items      chan item
	position   int
	lineMap    []line
	scopeStack []*ast.Scope
}

func Parse(filename string, input string) (*ast.File, errorList) {
	lexer := lex(filename, input)
	context := &parseContext{
		root:       ast.NewFile(filename),
		errors:     make(errorList, 0),
		itemBuffer: make([]item, 0),
		items:      lexer.items,
		lineMap:    mapNewLines(input),
		scopeStack: make([]*ast.Scope, 0),
	}

	context.doParse()

	// recreate the stack incase the parse left us with open scopes..
	context.scopeStack = make([]*ast.Scope, 0)
	context.doResolveDecls(context.root.Scope)

	return context.root, context.errors
}

// mapNewLines creates a mapping of byte indices to line numbers. used for error messages
func mapNewLines(s string) []line {
	linemap := make([]line, 0)
	accum := 0
	for {
		index := strings.Index(s, "\n")
		if index == -1 {
			break
		}
		linemap = append(linemap, line{Pos(accum), Pos(accum + index)})
		s = s[index+1:]
		accum = accum + index
	}
	return linemap
}

// scopeDepth returns the depth of the current scope
func (c *parseContext) scopeDepth() int {
	return len(c.scopeStack)
}

// currentScope returns the scope we're currently parsing within
func (c *parseContext) currentScope() *ast.Scope {
	if c.scopeDepth() == 0 {
		return nil
	}
	return c.scopeStack[len(c.scopeStack)-1]
}

// pushScope enters a new scope by pushing to the scope stack
func (c *parseContext) pushScope(s *ast.Scope) {
	c.scopeStack = append(c.scopeStack, s)
}

// popScope returns from a scope by popping from the scope stack
func (c *parseContext) popScope() *ast.Scope {
	scope := c.currentScope()
	if scope != nil {
		c.scopeStack = c.scopeStack[:len(c.scopeStack)-1]
	}
	return scope
}

// resolveDecl resolves a name to a node in the enclosing scopes
func (c *parseContext) resolveDecl(name string) ast.Node {
	scope := c.popScope()
	defer c.pushScope(scope)

	var node ast.Node
	if scope != nil {
		for _, decl := range scope.S {
			if decl.Identifier() == name {
				node = decl
				break
			}
		}
		if node == nil {
			// Not in this scope.. Go up
			node = c.resolveDecl(name)
		}
	} else {
		// We've recursed up to the top!
	}
	return node
}

// doResolveDecls walks the Ast, attempting to resolve any unresolved type references
func (c *parseContext) doResolveDecls(n ast.Node) {
	switch n := n.(type) {
	case *ast.Scope:
		c.pushScope(n)
		for _, decl := range n.S {
			c.doResolveDecls(decl)
		}
		c.popScope()
	case *ast.DeclReference:
		name := n.DeclName
		if node := c.resolveDecl(name); node != nil {
			n.Object = node
		} else {
			c.errorf(n.Meta.(item), "unresolved reference to type '%v'", n.DeclName)
		}
	case *ast.Struct:
		c.doResolveDecls(n.Scope)
	case *ast.DynamicLength:
		c.doResolveDecls(n.Field)
	case *ast.ArrayType:
		c.doResolveDecls(n.Object)
		c.doResolveDecls(n.Length)
	case *ast.Field:
		c.doResolveDecls(n.Type)
	case *ast.Frame:
		c.doResolveDecls(n.Object)
	case *ast.IntegerType:
	case *ast.StringBaseType:
	case *ast.StaticLength:
	default:
		if n != nil {
			// nil can occur due to lex/parse errors.
			panic(fmt.Sprintf("couldn't do anything with %T\n", n))
		}
	}
}

// peek returns the next item in the buffer without incrementing the position
// peek will fill the buffer from the lex channel if we're out of items
func (c *parseContext) peek() item {
	for c.position >= len(c.itemBuffer) {
		item := <-c.items
		c.itemBuffer = append(c.itemBuffer, item)
	}

	return c.itemBuffer[c.position]
}

// next returns the next item from the buffer and increases the position
func (c *parseContext) next() item {
	item := c.peek()
	c.position = c.position + 1
	return item
}

// prev returns the last item we passed without modifying the position
func (c *parseContext) prev() item {
	return c.itemBuffer[c.position-1]
}

// has determines if the next token matches one of the given itemTypes
// position is not modified
func (c *parseContext) has(types ...itemType) bool {
	item := c.peek()
	for _, typ := range types {
		if item.typ == typ {
			return true
		}
	}
	return false
}

// expect consumes a token if it matches the given type. returns true if the token existed
func (c *parseContext) expect(typ itemType, ignoreSpace bool) bool {
	if ignoreSpace {
		c.consumeWhitespace(typ)
	}

	if c.has(typ) {
		c.next()
		return true
	}
	return false
}

// accept returns the next token if it matches the given type, optionally consuming any preceding whitespace
func (c *parseContext) accept(typ itemType, ignoreSpace bool) (item, error) {
	if ignoreSpace {
		c.consumeWhitespace()
	}

	if c.has(typ) {
		return c.next(), nil
	}
	item := c.peek()
	return item, fmt.Errorf("expected %v, found '%v'", typ, item.val)
}

// error generates a parse error at the given item
func (c *parseContext) error(item item, err error) {
	e := parseError{
		item:  item,
		error: fmt.Errorf("%v:%v: %v", c.root.Name, item.pos.Line(c), err),
	}
	c.errors = append(c.errors, e)
}

// errorf formats and generates a parse error at the given item
func (c *parseContext) errorf(item item, format string, args ...interface{}) {
	c.error(item, fmt.Errorf(format, args...))
}

// consumeWhitespace swallows consecutive whitespace. returns true if any whitespace was found
func (c *parseContext) consumeWhitespace(except ...itemType) bool {
	found := false
	ignore := []itemType{}
Outer:
	for _, typ := range []itemType{itemWhiteSpace, itemEOL, itemComment} {
		for _, exc := range except {
			if typ == exc {
				continue Outer
			}
		}
		ignore = append(ignore, typ)
	}
	for c.has(ignore...) {
		c.next()
		found = true
	}
	return found
}

// doParse is the main parsing routine
func (c *parseContext) doParse() {
	c.pushScope(c.root.Scope)
	defer c.popScope()

	for !c.has(itemEOF) {
		c.consumeWhitespace()
		switch {
		case c.has(itemIdentifier):
			decl := c.parseDecl()
			if decl != nil {
				c.currentScope().Add(decl)
			}
		default:
			item := c.next()
			c.errorf(item, "unexpected %v in global scope", item.typ)
		}
	}
}

// parseDecl parses top level declarations
func (c *parseContext) parseDecl() ast.Node {
	var node ast.Node

	var identifier string
	if item, err := c.accept(itemIdentifier, true); err != nil {
		c.error(item, err)
		return nil
	} else {
		identifier = item.val
	}

	c.consumeWhitespace()

	switch {
	case c.has(itemFrame):
		node = c.parseFrameDecl(identifier)
	case c.has(itemStruct):
		node = c.parseStructDecl(identifier)
	default:
		item := c.next()
		c.errorf(item, "expected frame or struct, found '%v'", item.val)
		return nil
	}

	// Declarations must be followed by EOL
	if !c.expect(itemEOL, true) && !c.expect(itemEOF, true) {
		c.errorf(c.next(), "expected EOL following declaration")
	}

	return node
}

// parseField parses inner-struct field declarations
func (c *parseContext) parseField() ast.Node {
	var identifier string
	if item, err := c.accept(itemIdentifier, true); err != nil {
		c.error(item, err)
	} else {
		identifier = item.val
	}

	typ := c.parseType()

	return &ast.Field{
		Name: identifier,
		Type: typ,
	}
}

// parseType parses the type/structure of a field/decl
func (c *parseContext) parseType() ast.Node {
	c.consumeWhitespace()

	var typ ast.Node
	requireArraySpec := false

	switch {
	case c.has(itemStruct):
		//TODO: should anonymous structs be declared globally?
		anonStructName := "AnonStruct_" + strconv.Itoa(int(c.position))
		typ = c.parseStructDecl(anonStructName)
	case c.has(itemStringType):
		c.next()
		typ = &ast.StringBaseType{}
		requireArraySpec = true
	case c.has(itemIntType):
		typ = c.parseIntType()
	case c.has(itemIdentifier):
		item := c.next()
		typ = &ast.DeclReference{
			DeclName: item.val,
			Meta:     item,
		}
	default:
		item := c.next()
		c.errorf(item, "expected type, found '%v'", item.val)
		return nil
	}

	if c.has(itemLeftSquareBrack) {
		return c.parseArrayType(typ)
	} else if requireArraySpec {
		c.errorf(c.peek(), "expected array expression on type")
	}

	return typ
}

// parseArrayType parses creates arrays over types
func (c *parseContext) parseArrayType(base ast.Node) ast.Node {
	if _, err := c.accept(itemLeftSquareBrack, false); err != nil {
		panic("never reached")
	}

	array := &ast.ArrayType{
		Object: base,
	}
	switch {
	case c.has(itemNumber):
		count := c.parseNumber(c.next())
		array.Length = &ast.StaticLength{count}
	case c.has(itemIdentifier):
		item := c.next()
		declRef := &ast.DeclReference{DeclName: item.val, Meta: item}
		array.Length = &ast.DynamicLength{declRef}
	default:
		c.errorf(c.next(), "unknown array size expression")
	}

	if item, err := c.accept(itemRightSquareBrack, false); err != nil {
		c.errorf(item, "unclosed array expression")
		return nil
	}

	return array
}

// parseIntType parses int{8,16,32,64}<flags>
func (c *parseContext) parseIntType() ast.Node {
	if _, err := c.accept(itemIntType, true); err != nil {
		panic("never reached")
	}

	item := c.prev()
	intType := &ast.IntegerType{
		Signed:    false, // do we need unsigned integers? probably
		Bitsize:   intSizes[item.val],
		Modifiers: ast.IntegerFlag(0),
	}

	if c.expect(itemLeftAngleBrack, false) {
		for {
			if !c.expect(itemFlag, true) {
				c.errorf(item, "expected flag in integer specialization")
			}

			flag := intFlags[c.prev().val]
			intType.Set(flag)

			if !c.expect(itemComma, true) {
				break
			}
		}

		if item, err := c.accept(itemRightAngleBrack, true); err != nil {
			c.errorf(item, "unclosed type specialization on integer type")
			return nil
		}
	}

	return intType
}

// parseStructDecl parses a structure, assuming we're currently at the 'struct' item
func (c *parseContext) parseStructDecl(identifier string) ast.Node {
	if _, err := c.accept(itemStruct, true); err != nil {
		panic("never reached")
	}

	structNode := ast.NewStruct(identifier)
	c.pushScope(structNode.Scope)
	defer c.popScope()

	if item, err := c.accept(itemLeftBrack, true); err != nil {
		c.errorf(item, "expected struct scope")
		return structNode
	}

	for {
		if _, err := c.accept(itemRightBrack, true); err == nil {
			break
		}

		if c.has(itemEOF) {
			c.errorf(c.peek(), "unclosed struct block")
			return structNode
		}

		field := c.parseField()
		c.currentScope().Add(field)
	}

	return structNode
}

// parseFrameDecl parses frame declarations, assuming we're currently at the 'frame' item
func (c *parseContext) parseFrameDecl(identifier string) ast.Node {
	if _, err := c.accept(itemFrame, true); err != nil {
		panic("never reached")
	}

	frame := &ast.Frame{Name: identifier}

	if item, err := c.accept(itemLeftAngleBrack, false); err != nil {
		c.errorf(item, "missing type specialization in frame declaration")
	}

	if item, err := c.accept(itemNumber, true); err != nil {
		c.errorf(item, "missing frame number in declaration")
	} else {
		frame.Number = c.parseNumber(item)
	}

	if item, err := c.accept(itemComma, true); err != nil {
		c.errorf(item, "expected ','")
	}

	if item, err := c.accept(itemLenSpec, true); err != nil {
		c.errorf(item, "missing frame size declaration")
	} else {
		frame.Size = frameSizes[item.val]
	}

	if item, err := c.accept(itemRightAngleBrack, true); err != nil {
		c.errorf(item, "unclosed type specialization in frame declaration")
	}

	frame.Object = c.parseType()

	return frame
}

// parseNumber converts an itemNumber into it's integer representation
func (c *parseContext) parseNumber(item item) int {
	value, err := strconv.Atoi(item.val)
	if err != nil {
		c.error(item, err)
	}
	return value
}

// parseError ties an error message to a given item
type parseError struct {
	item  item
	error error
}

type errorList []parseError

func (errors errorList) String() string {
	var buf bytes.Buffer
	for _, err := range errors {
		buf.WriteString(fmt.Sprintf("%v\n", err.error))
	}
	return buf.String()
}

package parse

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Pos represents a byte position in the original input text from which
// this template was parsed.
type Pos int

func (p Pos) Position() Pos {
	return p
}

// item represents a token or text string returned from the scanner.
type item struct {
	typ itemType // The type of this item.
	val string   // The value of this item.
	pos Pos      // The starting position, in bytes, of this item in the input string.
}

// itemType is a lexed token
type itemType int

//go:generate stringer -type=itemType

const (
	itemError itemType = iota
	itemEOF
	itemEOL
	itemWhiteSpace
	itemLeftBrack
	itemRightBrack
	itemLeftSquareBrack
	itemRightSquareBrack
	itemLeftAngleBrack
	itemRightAngleBrack
	itemComma
	itemIdentifier
	itemNumber
	itemComment
	/* Keywords */
	itemKeyword // Not emitted; used as a delimiter
	itemFrame
	itemStruct
	itemStringType
	itemIntType
	itemLenSpec // Fixed, Var8, or Var16
	itemFlag    // LittleEndian, PDPEndian, RPDPEndian, Negate, Offset128, Inverse128
)

var keywords = map[string]itemType{
	"frame":        itemFrame,
	"struct":       itemStruct,
	"string":       itemStringType,
	"int8":         itemIntType,
	"int16":        itemIntType,
	"int32":        itemIntType,
	"int64":        itemIntType,
	"Fixed":        itemLenSpec,
	"Var8":         itemLenSpec,
	"Var16":        itemLenSpec,
	"LittleEndian": itemFlag,
	"PDPEndian":    itemFlag,
	"RPDPEndian":   itemFlag,
	"Negate":       itemFlag,
	"Offset128":    itemFlag,
	"Inverse128":   itemFlag,
}

const eof = -1

// stateFn represents the state of the scanner as a function that returns the next state.
type stateFn func(*lexer) stateFn

// lexer holds the state of the scanner.
type lexer struct {
	name       string  // the name of the input; used only for error reports
	input      string  // the string being scanned
	state      stateFn // the next lexing function to enter
	stateStack []stateFn
	pos        Pos       // current position in the input
	start      Pos       // start position of this item
	width      Pos       // width of last rune read from input
	lastPos    Pos       // position of most recent item returned by nextItem
	items      chan item // channel of scanned items
	parenDepth int       // nesting depth of { } blocks
}

// pushState pushes the current stateFn to the stack
func (l *lexer) pushState() {
	l.stateStack = append(l.stateStack, l.state)
}

// popState pops the last stateFn from the stack
func (l *lexer) popState() stateFn {
	state := l.stateStack[len(l.stateStack)-1]
	l.stateStack = l.stateStack[:len(l.stateStack)-1]
	return state
}

// detour pushes the current state, then enters a given stateFn
func (l *lexer) detour(fn stateFn) stateFn {
	l.pushState()
	return fn
}

// next returns the next rune in the input.
func (l *lexer) next() rune {
	if int(l.pos) >= len(l.input) {
		l.width = 0
		return eof
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = Pos(w)
	l.pos += l.width
	return r
}

// peek returns but does not consume the next rune in the input.
func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// backup steps back one rune. Can only be called once per call of next.
func (l *lexer) backup() {
	l.pos -= l.width
}

// emit passes an item back to the client.
func (l *lexer) emit(t itemType) {
	l.items <- item{t, l.input[l.start:l.pos], l.start}
	l.start = l.pos
}

// ignore skips over the pending input before this point.
func (l *lexer) ignore() {
	l.start = l.pos
}

// accept consumes the next rune if it's from the valid set.
func (l *lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

// lineNumber reports which line we're on, based on the position of
// the previous item returned by nextItem. Doing it this way
// means we don't have to worry about peek double counting.
func (l *lexer) lineNumber() int {
	return 1 + strings.Count(l.input[:l.lastPos], "\n")
}

// errorf returns an error token and terminates the scan by passing
// back a nil pointer that will be the next state, terminating l.nextItem.
func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- item{itemError, fmt.Sprintf(format, args...), l.start}
	return nil
}

// nextItem returns the next item from the input.
// Called by the parser, not in the lexing goroutine.
func (l *lexer) nextItem() item {
	item := <-l.items
	l.lastPos = item.pos
	return item
}

// drain drains the output so the lexing goroutine will exit.
// Called by the parser, not in the lexing goroutine.
func (l *lexer) drain() {
	for range l.items {
	}
}

// lex creates a new scanner for the input string.
func lex(name, input string) *lexer {
	l := &lexer{
		name:       name,
		input:      input,
		items:      make(chan item),
		stateStack: make([]stateFn, 0),
	}
	go l.run()
	return l
}

// run runs the state machine for the lexer.
func (l *lexer) run() {
	for l.state = lexText; l.state != nil; {
		l.state = l.state(l)
	}
	close(l.items)
}

// state functions

const (
	leftBracket        = '{'
	rightBracket       = '}'
	leftSquareBracket  = '['
	rightSquareBracket = ']'
	leftAngleBracket   = '<'
	rightAngleBracket  = '>'
	leftComment        = "/*"
	rightComment       = "*/"
)

// lexText
func lexText(l *lexer) stateFn {
	switch r := l.peek(); {
	case isWhiteSpace(r):
		return l.detour(lexWhiteSpace)
	case r == '/': // is this enough to detect a comment?
		return l.detour(lexComment)
	case isAlphaNumeric(r):
		return l.detour(lexAlphaNumeric)
	case r == leftAngleBracket:
		return l.detour(lexTypeArgs)
	case r == leftSquareBracket:
		return l.detour(lexArrayExpr)
	case r == leftBracket:
		return l.detour(lexStruct)
	case isEndOfLine(r):
		l.next()
		l.emit(itemEOL)
		return lexText
	}
	l.emit(itemEOF)
	return nil
}

// lexWhiteSpace consumes whitespace, or comments, or EOL
func lexWhiteSpace(l *lexer) stateFn {
	for isWhiteSpace(l.peek()) {
		l.next()
	}
	l.emit(itemWhiteSpace)
	return l.popState()
}

// lexComment scans a comment. The left comment marker is known to be present.
func lexComment(l *lexer) stateFn {
	l.pos += Pos(len(leftComment))
	i := strings.Index(l.input[l.pos:], rightComment)
	if i < 0 {
		return l.errorf("unclosed comment")
	}
	l.pos += Pos(i + len(rightComment))
	l.emit(itemComment)
	return l.popState()
}

// lexAlphaNumeric scans identifiers
func lexAlphaNumeric(l *lexer) stateFn {
	if r := l.peek(); isDigit(r) {
		return lexNumber
	}
Loop:
	for {
		switch r := l.next(); {
		case isAlphaNumeric(r):
			// absorb.
		default:
			l.backup()
			word := l.input[l.start:l.pos]
			if !isWhiteSpace(r) && !isBracket(r) && !isComma(r) && r != eof {
				return l.errorf("bad character %#U", r)
			}
			switch {
			case keywords[word] > itemKeyword:
				l.emit(keywords[word])
			default:
				l.emit(itemIdentifier)
			}
			break Loop
		}
	}
	return l.popState()
}

// lexNumber scans identifiers
func lexNumber(l *lexer) stateFn {
Loop:
	for {
		switch r := l.next(); {
		case isDigit(r):
			// absorb
		default:
			l.backup()
			if !isWhiteSpace(r) && !isBracket(r) && !isComma(r) && r != eof {
				return l.errorf("bad character %#U", r)
			}
			l.emit(itemNumber)
			break Loop
		}
	}
	return l.popState()
}

// lexTypeArgs scans type arguments, eg. type<arg, arg>. The '<' has already been scanned
func lexTypeArgs(l *lexer) stateFn {
	switch r := l.next(); {
	case isWhiteSpace(r):
		l.backup()
		return l.detour(lexWhiteSpace)
	case r == leftAngleBracket:
		l.emit(itemLeftAngleBrack)
		return lexTypeArgs
	case r == rightAngleBracket:
		l.emit(itemRightAngleBrack)
		return l.popState()
	case isComma(r):
		l.emit(itemComma)
		return lexTypeArgs
	default:
		l.backup()
		return l.detour(lexAlphaNumeric)
	}
	panic("never reached")
}

// lexStruct scans the content of a struct declaration
func lexStruct(l *lexer) stateFn {
	switch r := l.next(); {
	case isEndOfLine(r):
		l.emit(itemEOL)
		return lexStruct
	case isWhiteSpace(r):
		return l.detour(lexWhiteSpace)
	case r == leftBracket:
		l.emit(itemLeftBrack)
		l.parenDepth = l.parenDepth + 1
		return lexStruct
	case r == leftSquareBracket:
		l.backup()
		return l.detour(lexArrayExpr)
	case r == leftAngleBracket:
		l.backup()
		return l.detour(lexTypeArgs)
	case r == leftBracket:
		l.backup()
		return l.detour(lexStruct)
	case r == rightBracket:
		l.emit(itemRightBrack)
		l.parenDepth = l.parenDepth - 1
		if l.parenDepth == 0 {
			return l.popState()
		}
		return lexStruct
	default:
		return l.detour(lexAlphaNumeric)
	}
	panic("never reached")
}

// lexArrayExpr scans square bracketed array exprs
func lexArrayExpr(l *lexer) stateFn {
	switch r := l.next(); {
	case isWhiteSpace(r):
		l.backup()
		return l.detour(lexWhiteSpace)
	case r == leftSquareBracket:
		l.emit(itemLeftSquareBrack)
		return lexArrayExpr
	case r == rightSquareBracket:
		l.emit(itemRightSquareBrack)
		return l.popState()
	default:
		l.backup()
		return l.detour(lexAlphaNumeric)
	}
	panic("never reached")
}

// isBracket reports whether r is one of the bracket characters
func isBracket(r rune) bool {
	return r == leftBracket || r == leftAngleBracket || r == leftSquareBracket ||
		r == rightBracket || r == rightAngleBracket || r == rightSquareBracket
}

func isDigit(r rune) bool {
	return unicode.IsDigit(r)
}

// isWhiteSpace reports whether r is a whitespace (ignored) character
func isWhiteSpace(r rune) bool {
	return isSpace(r) || isEndOfLine(r)
}

// isSpace reports whether r is a space character.
func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

func isComma(r rune) bool {
	return r == ','
}

// isEndOfLine reports whether r is an end-of-line character.
func isEndOfLine(r rune) bool {
	return r == '\r' || r == '\n'
}

// isAlphaNumeric reports whether r is an alphabetic, digit, or underscore.
func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}

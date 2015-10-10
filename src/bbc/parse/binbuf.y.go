//line binbuf.y:2
package parse

import __yyfmt__ "fmt"

//line binbuf.y:2
import (
	"bbc/ast"
)

//line binbuf.y:10
type yySymType struct {
	yys     int
	nsl     []ast.Node
	n       ast.Node
	ival    int
	sval    string
	svalarr []string
	length  ast.LengthSpec
	size    ast.FrameSize
}

const tWhitespace = 57346
const tIdentifier = 57347
const tNumber = 57348
const tStruct = 57349
const tType = 57350
const tFrame = 57351
const tFrameFixed = 57352
const tFrameVar8 = 57353
const tFrameVar16 = 57354
const tStringType = 57355
const tByteType = 57356
const tIntegerType = 57357
const tIntegerFlag = 57358
const tEOL = 57359

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"'{'",
	"'}'",
	"'['",
	"']'",
	"'<'",
	"'>'",
	"','",
	"tWhitespace",
	"tIdentifier",
	"tNumber",
	"tStruct",
	"tType",
	"tFrame",
	"tFrameFixed",
	"tFrameVar8",
	"tFrameVar16",
	"tStringType",
	"tByteType",
	"tIntegerType",
	"tIntegerFlag",
	"tEOL",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line binbuf.y:206

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 36
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 57

var yyAct = [...]int{

	2, 33, 55, 7, 42, 5, 18, 9, 30, 31,
	32, 12, 40, 41, 39, 37, 18, 51, 16, 11,
	48, 26, 24, 22, 15, 5, 53, 54, 28, 25,
	43, 17, 46, 19, 52, 45, 3, 21, 4, 1,
	49, 8, 6, 10, 13, 47, 14, 50, 38, 20,
	23, 27, 36, 35, 34, 44, 29,
}
var yyPact = [...]int{

	14, -1000, 14, -1000, 14, -1000, 14, -1000, -1000, 4,
	14, 12, -1000, -1000, -1000, 2, 25, -1000, 33, 10,
	-1000, 14, 19, 16, -1000, -9, -1000, -1000, -8, 21,
	-1000, -1000, -1000, 29, -1000, -1000, -1000, -1000, -1000, 24,
	-1000, -1000, -1000, -8, -1000, 7, -6, 29, 27, 17,
	-1000, -1000, -1000, -1000, -21, -1000,
}
var yyPgo = [...]int{

	0, 56, 55, 1, 54, 53, 52, 51, 50, 49,
	48, 46, 44, 15, 43, 42, 40, 39, 0, 38,
	36,
}
var yyR1 = [...]int{

	0, 17, 15, 15, 14, 14, 11, 1, 1, 1,
	12, 13, 9, 8, 8, 7, 4, 4, 16, 16,
	16, 10, 3, 3, 3, 3, 3, 3, 6, 5,
	2, 19, 20, 20, 18, 18,
}
var yyR2 = [...]int{

	0, 3, 1, 4, 2, 2, 8, 1, 1, 1,
	2, 2, 3, 1, 2, 2, 1, 4, 1, 1,
	3, 1, 1, 1, 1, 1, 1, 2, 1, 1,
	3, 1, 1, 2, 0, 1,
}
var yyChk = [...]int{

	-1000, -17, -18, -20, -19, 11, -15, -18, -20, -18,
	-14, 15, -18, -12, -11, 12, 16, -13, 14, 8,
	-9, 4, 13, -8, -18, 10, 5, -7, 12, -1,
	17, 18, 19, -3, -4, -5, -6, -13, -10, 22,
	20, 21, 12, 9, -2, 6, 8, -3, 13, -16,
	-18, 23, 7, 9, 10, 23,
}
var yyDef = [...]int{

	34, -2, 34, 35, 32, 31, 34, 2, 33, 1,
	34, 0, 3, 4, 5, 0, 0, 10, 0, 0,
	11, 34, 0, 0, 13, 0, 12, 14, 0, 0,
	7, 8, 9, 15, 22, 23, 24, 25, 26, 16,
	29, 28, 21, 0, 27, 0, 34, 6, 0, 0,
	18, 19, 30, 17, 0, 20,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 10, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	8, 3, 9, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 6, 3, 7, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 4, 3, 5,
}
var yyTok2 = [...]int{

	2, 3, 11, 12, 13, 14, 15, 16, 17, 18,
	19, 20, 21, 22, 23, 24,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lookahead func() int
}

func (p *yyParserImpl) Lookahead() int {
	return p.lookahead()
}

func yyNewParser() yyParser {
	p := &yyParserImpl{
		lookahead: func() int { return -1 },
	}
	return p
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yytoken := -1 // yychar translated into internal numbering
	yyrcvr.lookahead = func() int { return yychar }
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yychar = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar, yytoken = yylex1(yylex, &yylval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yychar = -1
		yytoken = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar, yytoken = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yychar = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line binbuf.y:52
		{
			yylex.(*Lexer).Ast().Scope = yyDollar[2].n.(*ast.Scope)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line binbuf.y:56
		{
			yyVAL.n = ast.NewScope()
		}
	case 3:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line binbuf.y:58
		{
			yyDollar[1].n.(*ast.Scope).Add(yyDollar[3].n)
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line binbuf.y:63
		{
			yyVAL.n = yyDollar[2].n
			yylex.(*Lexer).AddDecl(yyVAL.n)
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line binbuf.y:68
		{
			yyVAL.n = yyDollar[2].n
			yylex.(*Lexer).AddDecl(yyVAL.n)
		}
	case 6:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line binbuf.y:76
		{
			yyVAL.n = &ast.Frame{
				Name:   yyDollar[1].sval,
				Number: yyDollar[4].ival,
				Size:   yyDollar[6].size,
				Object: yyDollar[8].n,
			}
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line binbuf.y:88
		{
			yyVAL.size = ast.SzFixed
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line binbuf.y:90
		{
			yyVAL.size = ast.SzVar8
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line binbuf.y:92
		{
			yyVAL.size = ast.SzVar16
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line binbuf.y:97
		{
			yyDollar[2].n.(*ast.Struct).Name = yyDollar[1].sval
			yyVAL.n = yyDollar[2].n
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line binbuf.y:105
		{

			yyVAL.n = &ast.Struct{
				Name:  yylex.(*Lexer).NameAnonStruct(),
				Scope: yyDollar[2].n.(*ast.Scope),
			}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line binbuf.y:116
		{
			yyVAL.n = yyDollar[2].n
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line binbuf.y:120
		{
			yyVAL.n = ast.NewScope()
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line binbuf.y:122
		{
			yyDollar[1].n.(*ast.Scope).Add(yyDollar[2].n.(ast.Node))
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line binbuf.y:127
		{
			yyVAL.n = &ast.Field{
				Name: yyDollar[1].sval,
				Type: yyDollar[2].n.(ast.Node),
			}
		}
	case 17:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line binbuf.y:138
		{
			yyDollar[1].n.(*ast.IntegerType).Modifiers = yyDollar[3].svalarr
			yyVAL.n = yyDollar[1].n
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line binbuf.y:145
		{
			yyVAL.svalarr = make([]string, 0)
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line binbuf.y:147
		{
			yyVAL.svalarr = append(yyVAL.svalarr, yyDollar[1].sval)
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line binbuf.y:149
		{
			yyVAL.svalarr = append(yyVAL.svalarr, yyDollar[3].sval)
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line binbuf.y:154
		{
			yyVAL.n = &ast.DeclReference{
				DeclName: yyDollar[1].sval,
			}
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line binbuf.y:168
		{
			yyVAL.n = &ast.ArrayType{
				Object: yyDollar[1].n,
				Length: yyDollar[2].length,
			}
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line binbuf.y:178
		{
			yyVAL.n = &ast.ByteBaseType{}
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line binbuf.y:183
		{
			yyVAL.n = &ast.StringBaseType{}
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line binbuf.y:188
		{
			yyVAL.length = &ast.StaticLength{
				Length: yyDollar[2].ival,
			}
		}
	}
	goto yystack /* stack new state and value */
}

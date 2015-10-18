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
const tBitStruct = 57352
const tFrameFixed = 57353
const tFrameVar8 = 57354
const tFrameVar16 = 57355
const tStringType = 57356
const tByteType = 57357
const tBitsType = 57358
const tIntegerType = 57359
const tIntegerFlag = 57360
const tEOL = 57361

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
	"tBitStruct",
	"tFrameFixed",
	"tFrameVar8",
	"tFrameVar16",
	"tStringType",
	"tByteType",
	"tBitsType",
	"tIntegerType",
	"tIntegerFlag",
	"tEOL",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line binbuf.y:230

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 40
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 63

var yyAct = [...]int{

	2, 35, 48, 7, 46, 5, 18, 9, 60, 19,
	21, 12, 40, 44, 45, 43, 42, 53, 11, 56,
	32, 33, 34, 26, 18, 28, 16, 19, 17, 24,
	23, 15, 30, 58, 59, 47, 5, 27, 50, 20,
	57, 49, 3, 22, 4, 1, 51, 8, 54, 52,
	6, 55, 10, 13, 14, 41, 25, 29, 37, 39,
	38, 36, 31,
}
var yyPact = [...]int{

	25, -1000, 25, -1000, 25, -1000, 25, -1000, -1000, 3,
	25, 19, -1000, -1000, -1000, 10, 31, -1000, 39, 39,
	16, -1000, 25, -1000, 27, 20, -1000, 2, -1000, -1000,
	-8, 26, -1000, -1000, -1000, 35, -1000, -1000, -1000, -1000,
	-1000, -1000, 30, 35, -1000, -1000, -1000, -8, -1000, 4,
	-6, -1000, 35, 33, 24, -1000, -1000, -1000, -1000, -17,
	-1000,
}
var yyPgo = [...]int{

	0, 62, 2, 1, 61, 60, 59, 58, 57, 56,
	10, 55, 54, 53, 12, 52, 50, 48, 45, 0,
	44, 42,
}
var yyR1 = [...]int{

	0, 18, 16, 16, 15, 15, 12, 1, 1, 1,
	13, 14, 14, 10, 9, 9, 8, 4, 4, 17,
	17, 17, 11, 3, 3, 3, 3, 3, 3, 3,
	6, 5, 7, 7, 2, 20, 21, 21, 19, 19,
}
var yyR2 = [...]int{

	0, 3, 1, 4, 2, 2, 8, 1, 1, 1,
	2, 2, 2, 3, 1, 2, 2, 1, 4, 1,
	1, 3, 1, 1, 1, 1, 1, 1, 1, 2,
	1, 1, 1, 2, 3, 1, 1, 2, 0, 1,
}
var yyChk = [...]int{

	-1000, -18, -19, -21, -20, 11, -16, -19, -21, -19,
	-15, 15, -19, -13, -12, 12, 16, -14, 14, 17,
	8, -10, 4, -10, 13, -9, -19, 10, 5, -8,
	12, -1, 18, 19, 20, -3, -4, -7, -5, -6,
	-14, -11, 24, 23, 21, 22, 12, 9, -2, 6,
	8, -2, -3, 13, -17, -19, 25, 7, 9, 10,
	25,
}
var yyDef = [...]int{

	38, -2, 38, 39, 36, 35, 38, 2, 37, 1,
	38, 0, 3, 4, 5, 0, 0, 10, 0, 0,
	0, 11, 38, 12, 0, 0, 14, 0, 13, 15,
	0, 0, 7, 8, 9, 16, 23, 24, 25, 26,
	27, 28, 17, 32, 31, 30, 22, 0, 29, 0,
	38, 33, 6, 0, 0, 19, 20, 34, 18, 0,
	21,
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
	19, 20, 21, 22, 23, 24, 25, 26,
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
			if _, ok := yyDollar[2].n.(*ast.BitStruct); ok {
				yyDollar[2].n.(*ast.BitStruct).Name = yyDollar[1].sval
			} else {
				yyDollar[2].n.(*ast.Struct).Name = yyDollar[1].sval
			}
			yyVAL.n = yyDollar[2].n
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line binbuf.y:109
		{
			yyVAL.n = &ast.Struct{
				Name:  yylex.(*Lexer).NameAnonStruct(),
				Scope: yyDollar[2].n.(*ast.Scope),
			}
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line binbuf.y:116
		{
			yyVAL.n = &ast.BitStruct{
				Name:  yylex.(*Lexer).NameAnonStruct(),
				Scope: yyDollar[2].n.(*ast.Scope),
			}
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line binbuf.y:126
		{
			yyVAL.n = yyDollar[2].n
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line binbuf.y:130
		{
			yyVAL.n = ast.NewScope()
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line binbuf.y:132
		{
			yyDollar[1].n.(*ast.Scope).Add(yyDollar[2].n.(ast.Node))
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line binbuf.y:137
		{
			yyVAL.n = &ast.Field{
				Name: yyDollar[1].sval,
				Type: yyDollar[2].n.(ast.Node),
			}
		}
	case 18:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line binbuf.y:148
		{
			yyDollar[1].n.(*ast.IntegerType).Modifiers = yyDollar[3].svalarr
			yyVAL.n = yyDollar[1].n
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line binbuf.y:155
		{
			yyVAL.svalarr = make([]string, 0)
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line binbuf.y:157
		{
			yyVAL.svalarr = append(yyVAL.svalarr, yyDollar[1].sval)
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line binbuf.y:159
		{
			yyVAL.svalarr = append(yyVAL.svalarr, yyDollar[3].sval)
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line binbuf.y:164
		{
			yyVAL.n = &ast.DeclReference{
				DeclName: yyDollar[1].sval,
			}
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line binbuf.y:179
		{
			yyVAL.n = &ast.ArrayType{
				Object: yyDollar[1].n,
				Length: yyDollar[2].length,
			}
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line binbuf.y:189
		{
			yyVAL.n = &ast.ByteBaseType{}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line binbuf.y:194
		{
			yyVAL.n = &ast.StringBaseType{}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line binbuf.y:199
		{
			yyVAL.n = &ast.BitsType{Count: 1}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line binbuf.y:201
		{
			if _, ok := yyDollar[2].length.(*ast.StaticLength); !ok {
				yylex.(*Lexer).Error("bit fields require a constant length expression")
			} else {
				yyVAL.n = &ast.BitsType{Count: yyDollar[2].length.(*ast.StaticLength).Length}
			}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line binbuf.y:212
		{
			yyVAL.length = &ast.StaticLength{
				Length: yyDollar[2].ival,
			}
		}
	}
	goto yystack /* stack new state and value */
}

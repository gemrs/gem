//line binbuf.y:2
package parse

import __yyfmt__ "fmt"

//line binbuf.y:2
import (
	"github.com/gemrs/gem/bbc/ast"
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
const tEllipsis = 57349
const tStruct = 57350
const tType = 57351
const tFrame = 57352
const tBitStruct = 57353
const tFrameFixed = 57354
const tFrameVar8 = 57355
const tFrameVar16 = 57356
const tStringType = 57357
const tByteType = 57358
const tBitsType = 57359
const tIntegerType = 57360
const tIntegerFlag = 57361
const tEOL = 57362

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
	"tEllipsis",
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
const yyInitialStackSize = 16

//line binbuf.y:235

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 65

var yyAct = [...]int{

	2, 35, 48, 7, 62, 11, 46, 9, 21, 18,
	5, 12, 19, 40, 5, 24, 44, 45, 43, 42,
	32, 33, 34, 26, 18, 57, 16, 19, 23, 17,
	53, 54, 28, 15, 60, 61, 27, 47, 50, 30,
	20, 59, 58, 49, 22, 3, 51, 4, 1, 52,
	8, 56, 55, 6, 10, 13, 14, 41, 25, 29,
	37, 39, 38, 36, 31,
}
var yyPact = [...]int{

	3, -1000, 3, -1000, 3, -1000, 3, -1000, -1000, -11,
	3, 21, -1000, -1000, -1000, 9, 32, -1000, 40, 40,
	2, -1000, 3, -1000, 26, 27, -1000, 1, -1000, -1000,
	-6, 28, -1000, -1000, -1000, 37, -1000, -1000, -1000, -1000,
	-1000, -1000, 30, 37, -1000, -1000, -1000, -6, -1000, 17,
	-1, -1000, 37, 35, 34, 25, -1000, -1000, -1000, -1000,
	-1000, -22, -1000,
}
var yyPgo = [...]int{

	0, 64, 2, 1, 63, 62, 61, 60, 59, 58,
	8, 57, 56, 55, 13, 54, 53, 52, 48, 0,
	47, 45,
}
var yyR1 = [...]int{

	0, 18, 16, 16, 15, 15, 12, 1, 1, 1,
	13, 14, 14, 10, 9, 9, 8, 4, 4, 17,
	17, 17, 11, 3, 3, 3, 3, 3, 3, 3,
	6, 5, 7, 7, 2, 2, 20, 21, 21, 19,
	19,
}
var yyR2 = [...]int{

	0, 3, 1, 4, 2, 2, 8, 1, 1, 1,
	2, 2, 2, 3, 1, 2, 2, 1, 4, 1,
	1, 3, 1, 1, 1, 1, 1, 1, 1, 2,
	1, 1, 1, 2, 3, 3, 1, 1, 2, 0,
	1,
}
var yyChk = [...]int{

	-1000, -18, -19, -21, -20, 11, -16, -19, -21, -19,
	-15, 16, -19, -13, -12, 12, 17, -14, 15, 18,
	8, -10, 4, -10, 13, -9, -19, 10, 5, -8,
	12, -1, 19, 20, 21, -3, -4, -7, -5, -6,
	-14, -11, 25, 24, 22, 23, 12, 9, -2, 6,
	8, -2, -3, 13, 14, -17, -19, 26, 7, 7,
	9, 10, 26,
}
var yyDef = [...]int{

	39, -2, 39, 40, 37, 36, 39, 2, 38, 1,
	39, 0, 3, 4, 5, 0, 0, 10, 0, 0,
	0, 11, 39, 12, 0, 0, 14, 0, 13, 15,
	0, 0, 7, 8, 9, 16, 23, 24, 25, 26,
	27, 28, 17, 32, 31, 30, 22, 0, 29, 0,
	39, 33, 6, 0, 0, 0, 19, 20, 34, 35,
	18, 0, 21,
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
	19, 20, 21, 22, 23, 24, 25, 26, 27,
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
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
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
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
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
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
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
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
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
			yyrcvr.char = -1
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
		//line binbuf.y:53
		{
			yylex.(*Lexer).Ast().Scope = yyDollar[2].n.(*ast.Scope)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line binbuf.y:57
		{
			yyVAL.n = ast.NewScope()
		}
	case 3:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line binbuf.y:59
		{
			yyDollar[1].n.(*ast.Scope).Add(yyDollar[3].n)
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line binbuf.y:64
		{
			yyVAL.n = yyDollar[2].n
			yylex.(*Lexer).AddDecl(yyVAL.n)
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line binbuf.y:69
		{
			yyVAL.n = yyDollar[2].n
			yylex.(*Lexer).AddDecl(yyVAL.n)
		}
	case 6:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line binbuf.y:77
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
		//line binbuf.y:89
		{
			yyVAL.size = ast.SzFixed
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line binbuf.y:91
		{
			yyVAL.size = ast.SzVar8
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line binbuf.y:93
		{
			yyVAL.size = ast.SzVar16
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line binbuf.y:98
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
		//line binbuf.y:110
		{
			yyVAL.n = &ast.Struct{
				Name:  yylex.(*Lexer).NameAnonStruct(),
				Scope: yyDollar[2].n.(*ast.Scope),
			}
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line binbuf.y:117
		{
			yyVAL.n = &ast.BitStruct{
				Name:  yylex.(*Lexer).NameAnonStruct(),
				Scope: yyDollar[2].n.(*ast.Scope),
			}
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line binbuf.y:127
		{
			yyVAL.n = yyDollar[2].n
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line binbuf.y:131
		{
			yyVAL.n = ast.NewScope()
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line binbuf.y:133
		{
			yyDollar[1].n.(*ast.Scope).Add(yyDollar[2].n.(ast.Node))
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line binbuf.y:138
		{
			yyVAL.n = &ast.Field{
				Name: yyDollar[1].sval,
				Type: yyDollar[2].n.(ast.Node),
			}
		}
	case 18:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line binbuf.y:149
		{
			yyDollar[1].n.(*ast.IntegerType).Modifiers = yyDollar[3].svalarr
			yyVAL.n = yyDollar[1].n
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line binbuf.y:156
		{
			yyVAL.svalarr = make([]string, 0)
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line binbuf.y:158
		{
			yyVAL.svalarr = append(yyVAL.svalarr, yyDollar[1].sval)
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line binbuf.y:160
		{
			yyVAL.svalarr = append(yyVAL.svalarr, yyDollar[3].sval)
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line binbuf.y:165
		{
			yyVAL.n = &ast.DeclReference{
				DeclName: yyDollar[1].sval,
			}
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line binbuf.y:180
		{
			yyVAL.n = &ast.ArrayType{
				Object: yyDollar[1].n,
				Length: yyDollar[2].length,
			}
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line binbuf.y:190
		{
			yyVAL.n = &ast.ByteBaseType{}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line binbuf.y:195
		{
			yyVAL.n = &ast.StringBaseType{}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line binbuf.y:200
		{
			yyVAL.n = &ast.BitsType{Count: 1}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line binbuf.y:202
		{
			if _, ok := yyDollar[2].length.(*ast.StaticLength); !ok {
				yylex.(*Lexer).Error("bit fields require a constant length expression")
			} else {
				yyVAL.n = &ast.BitsType{Count: yyDollar[2].length.(*ast.StaticLength).Length}
			}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line binbuf.y:213
		{
			yyVAL.length = &ast.StaticLength{
				Length: yyDollar[2].ival,
			}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line binbuf.y:219
		{
			yyVAL.length = &ast.RemainingLength{}
		}
	}
	goto yystack /* stack new state and value */
}

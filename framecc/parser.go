package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/sinusoids/gem/framecc/ast"
)

var currentFile *ast.File

func toIfaceSlice(v interface{}) []interface{} {
	if v == nil {
		return nil
	}
	return v.([]interface{})
}

func parseFrameDefinition(inputFile string) (*ast.File, error) {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	currentFile = &ast.File{
		Types:  make(map[string]ast.Type),
		Frames: make(map[string]ast.Frame),
	}

	_, err = ParseReader("", file)
	if err != nil {
		return nil, err
	}

	return currentFile, nil
}

var g = &grammar{
	rules: []*rule{
		{
			name: "Start",
			pos:  position{line: 39, col: 1, offset: 563},
			expr: &actionExpr{
				pos: position{line: 39, col: 10, offset: 572},
				run: (*parser).callonStart1,
				expr: &seqExpr{
					pos: position{line: 39, col: 10, offset: 572},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 39, col: 10, offset: 572},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 39, col: 12, offset: 574},
							label: "decls",
							expr: &oneOrMoreExpr{
								pos: position{line: 39, col: 18, offset: 580},
								expr: &ruleRefExpr{
									pos:  position{line: 39, col: 18, offset: 580},
									name: "Decl",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 39, col: 24, offset: 586},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 39, col: 26, offset: 588},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "Decl",
			pos:  position{line: 43, col: 1, offset: 622},
			expr: &choiceExpr{
				pos: position{line: 43, col: 9, offset: 630},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 43, col: 9, offset: 630},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 43, col: 9, offset: 630},
								name: "_",
							},
							&ruleRefExpr{
								pos:  position{line: 43, col: 11, offset: 632},
								name: "Frame",
							},
							&ruleRefExpr{
								pos:  position{line: 43, col: 17, offset: 638},
								name: "_",
							},
						},
					},
					&seqExpr{
						pos: position{line: 44, col: 9, offset: 648},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 44, col: 9, offset: 648},
								name: "_",
							},
							&ruleRefExpr{
								pos:  position{line: 44, col: 11, offset: 650},
								name: "Struct",
							},
							&ruleRefExpr{
								pos:  position{line: 44, col: 18, offset: 657},
								name: "_",
							},
						},
					},
				},
			},
		},
		{
			name: "Struct",
			pos:  position{line: 46, col: 1, offset: 660},
			expr: &actionExpr{
				pos: position{line: 46, col: 11, offset: 670},
				run: (*parser).callonStruct1,
				expr: &seqExpr{
					pos: position{line: 46, col: 11, offset: 670},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 46, col: 11, offset: 670},
							val:        "struct",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 20, offset: 679},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 46, col: 23, offset: 682},
							label: "identifier",
							expr: &ruleRefExpr{
								pos:  position{line: 46, col: 34, offset: 693},
								name: "Ident",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 40, offset: 699},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 46, col: 42, offset: 701},
							label: "fields",
							expr: &ruleRefExpr{
								pos:  position{line: 46, col: 49, offset: 708},
								name: "StructBlock",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 46, col: 61, offset: 720},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "StructBlock",
			pos:  position{line: 56, col: 1, offset: 868},
			expr: &actionExpr{
				pos: position{line: 56, col: 16, offset: 883},
				run: (*parser).callonStructBlock1,
				expr: &seqExpr{
					pos: position{line: 56, col: 16, offset: 883},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 56, col: 16, offset: 883},
							val:        "{",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 56, col: 20, offset: 887},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 56, col: 22, offset: 889},
							label: "fields",
							expr: &oneOrMoreExpr{
								pos: position{line: 56, col: 29, offset: 896},
								expr: &ruleRefExpr{
									pos:  position{line: 56, col: 29, offset: 896},
									name: "Field",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 56, col: 36, offset: 903},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 56, col: 38, offset: 905},
							val:        "}",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "Field",
			pos:  position{line: 66, col: 1, offset: 1103},
			expr: &actionExpr{
				pos: position{line: 66, col: 10, offset: 1112},
				run: (*parser).callonField1,
				expr: &seqExpr{
					pos: position{line: 66, col: 10, offset: 1112},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 66, col: 10, offset: 1112},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 66, col: 12, offset: 1114},
							label: "identifier",
							expr: &ruleRefExpr{
								pos:  position{line: 66, col: 23, offset: 1125},
								name: "Ident",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 66, col: 29, offset: 1131},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 66, col: 32, offset: 1134},
							label: "typ",
							expr: &ruleRefExpr{
								pos:  position{line: 66, col: 36, offset: 1138},
								name: "Type",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 66, col: 41, offset: 1143},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "Type",
			pos:  position{line: 73, col: 1, offset: 1232},
			expr: &choiceExpr{
				pos: position{line: 73, col: 9, offset: 1240},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 73, col: 9, offset: 1240},
						name: "StringType",
					},
					&ruleRefExpr{
						pos:  position{line: 73, col: 22, offset: 1253},
						name: "IntegerTypeNoFlags",
					},
					&ruleRefExpr{
						pos:  position{line: 73, col: 43, offset: 1274},
						name: "IntegerTypeAndFlags",
					},
					&ruleRefExpr{
						pos:  position{line: 73, col: 65, offset: 1296},
						name: "TypeRef",
					},
				},
			},
		},
		{
			name: "StringType",
			pos:  position{line: 75, col: 1, offset: 1305},
			expr: &choiceExpr{
				pos: position{line: 75, col: 15, offset: 1319},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 75, col: 15, offset: 1319},
						run: (*parser).callonStringType2,
						expr: &seqExpr{
							pos: position{line: 75, col: 15, offset: 1319},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 75, col: 15, offset: 1319},
									val:        "string",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 75, col: 24, offset: 1328},
									val:        "[",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 75, col: 28, offset: 1332},
									label: "size",
									expr: &ruleRefExpr{
										pos:  position{line: 75, col: 33, offset: 1337},
										name: "Number",
									},
								},
								&litMatcher{
									pos:        position{line: 75, col: 40, offset: 1344},
									val:        "]",
									ignoreCase: false,
								},
							},
						},
					},
					&actionExpr{
						pos: position{line: 79, col: 5, offset: 1408},
						run: (*parser).callonStringType9,
						expr: &seqExpr{
							pos: position{line: 79, col: 5, offset: 1408},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 79, col: 5, offset: 1408},
									val:        "string",
									ignoreCase: false,
								},
								&litMatcher{
									pos:        position{line: 79, col: 14, offset: 1417},
									val:        "[",
									ignoreCase: false,
								},
								&labeledExpr{
									pos:   position{line: 79, col: 18, offset: 1421},
									label: "fieldref",
									expr: &ruleRefExpr{
										pos:  position{line: 79, col: 27, offset: 1430},
										name: "Ident",
									},
								},
								&litMatcher{
									pos:        position{line: 79, col: 33, offset: 1436},
									val:        "]",
									ignoreCase: false,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "IntegerTypeNoFlags",
			pos:  position{line: 85, col: 1, offset: 1516},
			expr: &actionExpr{
				pos: position{line: 85, col: 23, offset: 1538},
				run: (*parser).callonIntegerTypeNoFlags1,
				expr: &seqExpr{
					pos: position{line: 85, col: 23, offset: 1538},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 85, col: 23, offset: 1538},
							label: "baseType",
							expr: &ruleRefExpr{
								pos:  position{line: 85, col: 32, offset: 1547},
								name: "IntegerType",
							},
						},
						&notExpr{
							pos: position{line: 85, col: 44, offset: 1559},
							expr: &seqExpr{
								pos: position{line: 85, col: 46, offset: 1561},
								exprs: []interface{}{
									&ruleRefExpr{
										pos:  position{line: 85, col: 46, offset: 1561},
										name: "_",
									},
									&litMatcher{
										pos:        position{line: 85, col: 48, offset: 1563},
										val:        "(",
										ignoreCase: false,
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "IntegerTypeAndFlags",
			pos:  position{line: 89, col: 1, offset: 1595},
			expr: &actionExpr{
				pos: position{line: 89, col: 24, offset: 1618},
				run: (*parser).callonIntegerTypeAndFlags1,
				expr: &seqExpr{
					pos: position{line: 89, col: 24, offset: 1618},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 89, col: 24, offset: 1618},
							label: "baseType",
							expr: &ruleRefExpr{
								pos:  position{line: 89, col: 33, offset: 1627},
								name: "IntegerType",
							},
						},
						&litMatcher{
							pos:        position{line: 89, col: 45, offset: 1639},
							val:        "(",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 89, col: 49, offset: 1643},
							label: "flags",
							expr: &oneOrMoreExpr{
								pos: position{line: 89, col: 55, offset: 1649},
								expr: &seqExpr{
									pos: position{line: 89, col: 56, offset: 1650},
									exprs: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 89, col: 56, offset: 1650},
											name: "IntegerFlag",
										},
										&zeroOrOneExpr{
											pos: position{line: 89, col: 68, offset: 1662},
											expr: &seqExpr{
												pos: position{line: 89, col: 69, offset: 1663},
												exprs: []interface{}{
													&litMatcher{
														pos:        position{line: 89, col: 69, offset: 1663},
														val:        ",",
														ignoreCase: false,
													},
													&ruleRefExpr{
														pos:  position{line: 89, col: 73, offset: 1667},
														name: "_",
													},
												},
											},
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 89, col: 79, offset: 1673},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "IntegerType",
			pos:  position{line: 119, col: 1, offset: 2287},
			expr: &actionExpr{
				pos: position{line: 119, col: 16, offset: 2302},
				run: (*parser).callonIntegerType1,
				expr: &seqExpr{
					pos: position{line: 119, col: 16, offset: 2302},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 119, col: 16, offset: 2302},
							label: "unsigned",
							expr: &zeroOrOneExpr{
								pos: position{line: 119, col: 25, offset: 2311},
								expr: &litMatcher{
									pos:        position{line: 119, col: 25, offset: 2311},
									val:        "u",
									ignoreCase: false,
								},
							},
						},
						&litMatcher{
							pos:        position{line: 119, col: 30, offset: 2316},
							val:        "int",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 119, col: 36, offset: 2322},
							label: "bitsize",
							expr: &ruleRefExpr{
								pos:  position{line: 119, col: 44, offset: 2330},
								name: "Number",
							},
						},
					},
				},
			},
		},
		{
			name: "IntegerFlag",
			pos:  position{line: 130, col: 1, offset: 2480},
			expr: &choiceExpr{
				pos: position{line: 130, col: 16, offset: 2495},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 130, col: 16, offset: 2495},
						val:        "negate",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 130, col: 27, offset: 2506},
						val:        "inv128",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 130, col: 38, offset: 2517},
						val:        "ofs128",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 130, col: 49, offset: 2528},
						val:        "endian(little)",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 130, col: 68, offset: 2547},
						val:        "endian(pdp)",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 130, col: 84, offset: 2563},
						run: (*parser).callonIntegerFlag7,
						expr: &litMatcher{
							pos:        position{line: 130, col: 84, offset: 2563},
							val:        "endian(rpdp)",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TypeRef",
			pos:  position{line: 134, col: 1, offset: 2614},
			expr: &actionExpr{
				pos: position{line: 134, col: 12, offset: 2625},
				run: (*parser).callonTypeRef1,
				expr: &labeledExpr{
					pos:   position{line: 134, col: 12, offset: 2625},
					label: "name",
					expr: &ruleRefExpr{
						pos:  position{line: 134, col: 17, offset: 2630},
						name: "Ident",
					},
				},
			},
		},
		{
			name: "Frame",
			pos:  position{line: 142, col: 1, offset: 2795},
			expr: &actionExpr{
				pos: position{line: 142, col: 10, offset: 2804},
				run: (*parser).callonFrame1,
				expr: &seqExpr{
					pos: position{line: 142, col: 10, offset: 2804},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 142, col: 10, offset: 2804},
							val:        "frame",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 142, col: 18, offset: 2812},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 142, col: 21, offset: 2815},
							label: "identifier",
							expr: &ruleRefExpr{
								pos:  position{line: 142, col: 32, offset: 2826},
								name: "Ident",
							},
						},
						&litMatcher{
							pos:        position{line: 142, col: 38, offset: 2832},
							val:        "(",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 142, col: 42, offset: 2836},
							label: "number",
							expr: &ruleRefExpr{
								pos:  position{line: 142, col: 49, offset: 2843},
								name: "Number",
							},
						},
						&labeledExpr{
							pos:   position{line: 142, col: 56, offset: 2850},
							label: "framesz",
							expr: &zeroOrOneExpr{
								pos: position{line: 142, col: 64, offset: 2858},
								expr: &seqExpr{
									pos: position{line: 142, col: 65, offset: 2859},
									exprs: []interface{}{
										&litMatcher{
											pos:        position{line: 142, col: 65, offset: 2859},
											val:        ",",
											ignoreCase: false,
										},
										&ruleRefExpr{
											pos:  position{line: 142, col: 69, offset: 2863},
											name: "_",
										},
										&ruleRefExpr{
											pos:  position{line: 142, col: 71, offset: 2865},
											name: "FrameSize",
										},
									},
								},
							},
						},
						&litMatcher{
							pos:        position{line: 142, col: 83, offset: 2877},
							val:        ")",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 142, col: 87, offset: 2881},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 142, col: 89, offset: 2883},
							label: "object",
							expr: &ruleRefExpr{
								pos:  position{line: 142, col: 96, offset: 2890},
								name: "TypeRef",
							},
						},
					},
				},
			},
		},
		{
			name: "Ident",
			pos:  position{line: 160, col: 1, offset: 3196},
			expr: &actionExpr{
				pos: position{line: 160, col: 10, offset: 3205},
				run: (*parser).callonIdent1,
				expr: &oneOrMoreExpr{
					pos: position{line: 160, col: 10, offset: 3205},
					expr: &charClassMatcher{
						pos:        position{line: 160, col: 10, offset: 3205},
						val:        "[a-zA-Z0-9_]",
						chars:      []rune{'_'},
						ranges:     []rune{'a', 'z', 'A', 'Z', '0', '9'},
						ignoreCase: false,
						inverted:   false,
					},
				},
			},
		},
		{
			name: "Number",
			pos:  position{line: 164, col: 1, offset: 3255},
			expr: &actionExpr{
				pos: position{line: 164, col: 11, offset: 3265},
				run: (*parser).callonNumber1,
				expr: &oneOrMoreExpr{
					pos: position{line: 164, col: 11, offset: 3265},
					expr: &charClassMatcher{
						pos:        position{line: 164, col: 11, offset: 3265},
						val:        "[0-9]",
						ranges:     []rune{'0', '9'},
						ignoreCase: false,
						inverted:   false,
					},
				},
			},
		},
		{
			name: "Value",
			pos:  position{line: 168, col: 1, offset: 3317},
			expr: &choiceExpr{
				pos: position{line: 168, col: 10, offset: 3326},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 168, col: 10, offset: 3326},
						run: (*parser).callonValue2,
						expr: &oneOrMoreExpr{
							pos: position{line: 168, col: 10, offset: 3326},
							expr: &charClassMatcher{
								pos:        position{line: 168, col: 10, offset: 3326},
								val:        "[a-zA-Z0-9]",
								ranges:     []rune{'a', 'z', 'A', 'Z', '0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 170, col: 5, offset: 3376},
						name: "Ident",
					},
				},
			},
		},
		{
			name: "FrameSize",
			pos:  position{line: 172, col: 1, offset: 3383},
			expr: &choiceExpr{
				pos: position{line: 172, col: 14, offset: 3396},
				alternatives: []interface{}{
					&actionExpr{
						pos: position{line: 172, col: 14, offset: 3396},
						run: (*parser).callonFrameSize2,
						expr: &litMatcher{
							pos:        position{line: 172, col: 14, offset: 3396},
							val:        "var8",
							ignoreCase: false,
						},
					},
					&actionExpr{
						pos: position{line: 174, col: 5, offset: 3436},
						run: (*parser).callonFrameSize4,
						expr: &litMatcher{
							pos:        position{line: 174, col: 5, offset: 3436},
							val:        "var16",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "SourceChar",
			pos:  position{line: 178, col: 1, offset: 3477},
			expr: &anyMatcher{
				line: 178, col: 15, offset: 3491,
			},
		},
		{
			name: "Comment",
			pos:  position{line: 179, col: 1, offset: 3493},
			expr: &choiceExpr{
				pos: position{line: 179, col: 12, offset: 3504},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 179, col: 12, offset: 3504},
						name: "MultiLineComment",
					},
					&ruleRefExpr{
						pos:  position{line: 179, col: 31, offset: 3523},
						name: "SingleLineComment",
					},
				},
			},
		},
		{
			name: "MultiLineComment",
			pos:  position{line: 180, col: 1, offset: 3541},
			expr: &seqExpr{
				pos: position{line: 180, col: 21, offset: 3561},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 180, col: 21, offset: 3561},
						val:        "/*",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 180, col: 26, offset: 3566},
						expr: &seqExpr{
							pos: position{line: 180, col: 28, offset: 3568},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 180, col: 28, offset: 3568},
									expr: &litMatcher{
										pos:        position{line: 180, col: 29, offset: 3569},
										val:        "*/",
										ignoreCase: false,
									},
								},
								&ruleRefExpr{
									pos:  position{line: 180, col: 34, offset: 3574},
									name: "SourceChar",
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 180, col: 48, offset: 3588},
						val:        "*/",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "SingleLineComment",
			pos:  position{line: 181, col: 1, offset: 3593},
			expr: &seqExpr{
				pos: position{line: 181, col: 22, offset: 3614},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 181, col: 22, offset: 3614},
						val:        "//",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 181, col: 27, offset: 3619},
						expr: &seqExpr{
							pos: position{line: 181, col: 29, offset: 3621},
							exprs: []interface{}{
								&notExpr{
									pos: position{line: 181, col: 29, offset: 3621},
									expr: &ruleRefExpr{
										pos:  position{line: 181, col: 30, offset: 3622},
										name: "EOL",
									},
								},
								&ruleRefExpr{
									pos:  position{line: 181, col: 34, offset: 3626},
									name: "SourceChar",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "_",
			pos:  position{line: 183, col: 1, offset: 3641},
			expr: &zeroOrMoreExpr{
				pos: position{line: 183, col: 7, offset: 3647},
				expr: &choiceExpr{
					pos: position{line: 183, col: 9, offset: 3649},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 183, col: 9, offset: 3649},
							name: "Whitespace",
						},
						&ruleRefExpr{
							pos:  position{line: 183, col: 22, offset: 3662},
							name: "EOL",
						},
						&ruleRefExpr{
							pos:  position{line: 183, col: 28, offset: 3668},
							name: "Comment",
						},
					},
				},
			},
		},
		{
			name: "__",
			pos:  position{line: 184, col: 1, offset: 3679},
			expr: &oneOrMoreExpr{
				pos: position{line: 184, col: 7, offset: 3685},
				expr: &choiceExpr{
					pos: position{line: 184, col: 9, offset: 3687},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 184, col: 9, offset: 3687},
							name: "Whitespace",
						},
						&ruleRefExpr{
							pos:  position{line: 184, col: 22, offset: 3700},
							name: "EOL",
						},
						&ruleRefExpr{
							pos:  position{line: 184, col: 28, offset: 3706},
							name: "Comment",
						},
					},
				},
			},
		},
		{
			name: "Whitespace",
			pos:  position{line: 186, col: 1, offset: 3718},
			expr: &charClassMatcher{
				pos:        position{line: 186, col: 15, offset: 3732},
				val:        "[ \\t\\r]",
				chars:      []rune{' ', '\t', '\r'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EOL",
			pos:  position{line: 187, col: 1, offset: 3740},
			expr: &litMatcher{
				pos:        position{line: 187, col: 8, offset: 3747},
				val:        "\n",
				ignoreCase: false,
			},
		},
		{
			name: "EOF",
			pos:  position{line: 188, col: 1, offset: 3752},
			expr: &notExpr{
				pos: position{line: 188, col: 8, offset: 3759},
				expr: &anyMatcher{
					line: 188, col: 9, offset: 3760,
				},
			},
		},
	},
}

func (c *current) onStart1(decls interface{}) (interface{}, error) {
	return currentFile, nil
}

func (p *parser) callonStart1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStart1(stack["decls"])
}

func (c *current) onStruct1(identifier, fields interface{}) (interface{}, error) {
	decl := ast.Struct{
		Name:   identifier.(string),
		Fields: fields.([]ast.Field),
	}

	currentFile.Types[decl.Name] = decl
	return decl, nil
}

func (p *parser) callonStruct1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStruct1(stack["identifier"], stack["fields"])
}

func (c *current) onStructBlock1(fields interface{}) (interface{}, error) {
	fieldsSl := make([]ast.Field, 0)
	fieldsIfSl := toIfaceSlice(fields)
	for _, fieldIfSl := range fieldsIfSl {
		fieldsSl = append(fieldsSl, fieldIfSl.(ast.Field))
	}

	return fieldsSl, nil
}

func (p *parser) callonStructBlock1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStructBlock1(stack["fields"])
}

func (c *current) onField1(identifier, typ interface{}) (interface{}, error) {
	return ast.Field{
		Name: identifier.(string),
		Type: typ.(ast.Type),
	}, nil
}

func (p *parser) callonField1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onField1(stack["identifier"], stack["typ"])
}

func (c *current) onStringType2(size interface{}) (interface{}, error) {
	return ast.StringType{
		Length: size.(int),
	}, nil
}

func (p *parser) callonStringType2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStringType2(stack["size"])
}

func (c *current) onStringType9(fieldref interface{}) (interface{}, error) {
	return ast.VariableStringType{
		FieldRef: fieldref.(string),
	}, nil
}

func (p *parser) callonStringType9() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStringType9(stack["fieldref"])
}

func (c *current) onIntegerTypeNoFlags1(baseType interface{}) (interface{}, error) {
	return baseType, nil
}

func (p *parser) callonIntegerTypeNoFlags1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIntegerTypeNoFlags1(stack["baseType"])
}

func (c *current) onIntegerTypeAndFlags1(baseType, flags interface{}) (interface{}, error) {
	iType := baseType.(ast.IntegerType)
	hasFlag := func(flags interface{}, flag string) bool {
		for _, f_ := range toIfaceSlice(flags) {
			f := toIfaceSlice(f_)
			if string(f[0].([]byte)) == flag {
				return true
			}
		}
		return false
	}

	modMap := map[string]ast.IntegerFlag{
		"negate":         ast.IntNegate,
		"inv128":         ast.IntInv128,
		"ofs128":         ast.IntOfs128,
		"endian(little)": ast.IntLittleEndian,
		"endian(pdp)":    ast.IntPDPEndian,
		"endian(rpdp)":   ast.IntRPDPEndian,
	}

	for k, v := range modMap {
		if hasFlag(flags, k) {
			iType.Modifiers = iType.Modifiers | v
		}
	}

	return iType, nil
}

func (p *parser) callonIntegerTypeAndFlags1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIntegerTypeAndFlags1(stack["baseType"], stack["flags"])
}

func (c *current) onIntegerType1(unsigned, bitsize interface{}) (interface{}, error) {
	signed := true
	if unsigned != nil {
		signed = false
	}
	return ast.IntegerType{
		Signed:  signed,
		Bitsize: bitsize.(int),
	}, nil
}

func (p *parser) callonIntegerType1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIntegerType1(stack["unsigned"], stack["bitsize"])
}

func (c *current) onIntegerFlag7() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonIntegerFlag7() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIntegerFlag7()
}

func (c *current) onTypeRef1(name interface{}) (interface{}, error) {
	if typ, ok := currentFile.Types[name.(string)]; !ok {
		return "", fmt.Errorf("%v: %v", ast.ErrNoSuchType, name.(string))
	} else {
		return typ, nil
	}
}

func (p *parser) callonTypeRef1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTypeRef1(stack["name"])
}

func (c *current) onFrame1(identifier, number, framesz, object interface{}) (interface{}, error) {
	if framesz == nil {
		framesz = ast.SzFixed
	} else {
		framesz = framesz.([]interface{})[2]
	}

	decl := ast.Frame{
		Name:   identifier.(string),
		Number: number.(int),
		Size:   framesz.(ast.FrameSize),
		Object: object.(ast.Type),
	}

	currentFile.Frames[decl.Name] = decl
	return decl, nil
}

func (p *parser) callonFrame1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFrame1(stack["identifier"], stack["number"], stack["framesz"], stack["object"])
}

func (c *current) onIdent1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonIdent1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIdent1()
}

func (c *current) onNumber1() (interface{}, error) {
	return strconv.Atoi(string(c.text))
}

func (p *parser) callonNumber1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onNumber1()
}

func (c *current) onValue2() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonValue2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onValue2()
}

func (c *current) onFrameSize2() (interface{}, error) {
	return ast.SzVar8, nil
}

func (p *parser) callonFrameSize2() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFrameSize2()
}

func (c *current) onFrameSize4() (interface{}, error) {
	return ast.SzVar16, nil
}

func (p *parser) callonFrameSize4() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFrameSize4()
}

var (
	// errNoRule is returned when the grammar to parse has no rule.
	errNoRule = errors.New("grammar has no rule")

	// errInvalidEncoding is returned when the source is not properly
	// utf8-encoded.
	errInvalidEncoding = errors.New("invalid encoding")

	// errNoMatch is returned if no match could be found.
	errNoMatch = errors.New("no match found")
)

// Option is a function that can set an option on the parser. It returns
// the previous setting as an Option.
type Option func(*parser) Option

// Debug creates an Option to set the debug flag to b. When set to true,
// debugging information is printed to stdout while parsing.
//
// The default is false.
func Debug(b bool) Option {
	return func(p *parser) Option {
		old := p.debug
		p.debug = b
		return Debug(old)
	}
}

// Memoize creates an Option to set the memoize flag to b. When set to true,
// the parser will cache all results so each expression is evaluated only
// once. This guarantees linear parsing time even for pathological cases,
// at the expense of more memory and slower times for typical cases.
//
// The default is false.
func Memoize(b bool) Option {
	return func(p *parser) Option {
		old := p.memoize
		p.memoize = b
		return Memoize(old)
	}
}

// Recover creates an Option to set the recover flag to b. When set to
// true, this causes the parser to recover from panics and convert it
// to an error. Setting it to false can be useful while debugging to
// access the full stack trace.
//
// The default is true.
func Recover(b bool) Option {
	return func(p *parser) Option {
		old := p.recover
		p.recover = b
		return Recover(old)
	}
}

// ParseFile parses the file identified by filename.
func ParseFile(filename string, opts ...Option) (interface{}, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ParseReader(filename, f, opts...)
}

// ParseReader parses the data from r using filename as information in the
// error messages.
func ParseReader(filename string, r io.Reader, opts ...Option) (interface{}, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return Parse(filename, b, opts...)
}

// Parse parses the data from b using filename as information in the
// error messages.
func Parse(filename string, b []byte, opts ...Option) (interface{}, error) {
	return newParser(filename, b, opts...).parse(g)
}

// position records a position in the text.
type position struct {
	line, col, offset int
}

func (p position) String() string {
	return fmt.Sprintf("%d:%d [%d]", p.line, p.col, p.offset)
}

// savepoint stores all state required to go back to this point in the
// parser.
type savepoint struct {
	position
	rn rune
	w  int
}

type current struct {
	pos  position // start position of the match
	text []byte   // raw text of the match
}

// the AST types...

type grammar struct {
	pos   position
	rules []*rule
}

type rule struct {
	pos         position
	name        string
	displayName string
	expr        interface{}
}

type choiceExpr struct {
	pos          position
	alternatives []interface{}
}

type actionExpr struct {
	pos  position
	expr interface{}
	run  func(*parser) (interface{}, error)
}

type seqExpr struct {
	pos   position
	exprs []interface{}
}

type labeledExpr struct {
	pos   position
	label string
	expr  interface{}
}

type expr struct {
	pos  position
	expr interface{}
}

type andExpr expr
type notExpr expr
type zeroOrOneExpr expr
type zeroOrMoreExpr expr
type oneOrMoreExpr expr

type ruleRefExpr struct {
	pos  position
	name string
}

type andCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type notCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type litMatcher struct {
	pos        position
	val        string
	ignoreCase bool
}

type charClassMatcher struct {
	pos        position
	val        string
	chars      []rune
	ranges     []rune
	classes    []*unicode.RangeTable
	ignoreCase bool
	inverted   bool
}

type anyMatcher position

// errList cumulates the errors found by the parser.
type errList []error

func (e *errList) add(err error) {
	*e = append(*e, err)
}

func (e errList) err() error {
	if len(e) == 0 {
		return nil
	}
	e.dedupe()
	return e
}

func (e *errList) dedupe() {
	var cleaned []error
	set := make(map[string]bool)
	for _, err := range *e {
		if msg := err.Error(); !set[msg] {
			set[msg] = true
			cleaned = append(cleaned, err)
		}
	}
	*e = cleaned
}

func (e errList) Error() string {
	switch len(e) {
	case 0:
		return ""
	case 1:
		return e[0].Error()
	default:
		var buf bytes.Buffer

		for i, err := range e {
			if i > 0 {
				buf.WriteRune('\n')
			}
			buf.WriteString(err.Error())
		}
		return buf.String()
	}
}

// parserError wraps an error with a prefix indicating the rule in which
// the error occurred. The original error is stored in the Inner field.
type parserError struct {
	Inner  error
	pos    position
	prefix string
}

// Error returns the error message.
func (p *parserError) Error() string {
	return p.prefix + ": " + p.Inner.Error()
}

// newParser creates a parser with the specified input source and options.
func newParser(filename string, b []byte, opts ...Option) *parser {
	p := &parser{
		filename: filename,
		errs:     new(errList),
		data:     b,
		pt:       savepoint{position: position{line: 1}},
		recover:  true,
	}
	p.setOptions(opts)
	return p
}

// setOptions applies the options to the parser.
func (p *parser) setOptions(opts []Option) {
	for _, opt := range opts {
		opt(p)
	}
}

type resultTuple struct {
	v   interface{}
	b   bool
	end savepoint
}

type parser struct {
	filename string
	pt       savepoint
	cur      current

	data []byte
	errs *errList

	recover bool
	debug   bool
	depth   int

	memoize bool
	// memoization table for the packrat algorithm:
	// map[offset in source] map[expression or rule] {value, match}
	memo map[int]map[interface{}]resultTuple

	// rules table, maps the rule identifier to the rule node
	rules map[string]*rule
	// variables stack, map of label to value
	vstack []map[string]interface{}
	// rule stack, allows identification of the current rule in errors
	rstack []*rule

	// stats
	exprCnt int
}

// push a variable set on the vstack.
func (p *parser) pushV() {
	if cap(p.vstack) == len(p.vstack) {
		// create new empty slot in the stack
		p.vstack = append(p.vstack, nil)
	} else {
		// slice to 1 more
		p.vstack = p.vstack[:len(p.vstack)+1]
	}

	// get the last args set
	m := p.vstack[len(p.vstack)-1]
	if m != nil && len(m) == 0 {
		// empty map, all good
		return
	}

	m = make(map[string]interface{})
	p.vstack[len(p.vstack)-1] = m
}

// pop a variable set from the vstack.
func (p *parser) popV() {
	// if the map is not empty, clear it
	m := p.vstack[len(p.vstack)-1]
	if len(m) > 0 {
		// GC that map
		p.vstack[len(p.vstack)-1] = nil
	}
	p.vstack = p.vstack[:len(p.vstack)-1]
}

func (p *parser) print(prefix, s string) string {
	if !p.debug {
		return s
	}

	fmt.Printf("%s %d:%d:%d: %s [%#U]\n",
		prefix, p.pt.line, p.pt.col, p.pt.offset, s, p.pt.rn)
	return s
}

func (p *parser) in(s string) string {
	p.depth++
	return p.print(strings.Repeat(" ", p.depth)+">", s)
}

func (p *parser) out(s string) string {
	p.depth--
	return p.print(strings.Repeat(" ", p.depth)+"<", s)
}

func (p *parser) addErr(err error) {
	p.addErrAt(err, p.pt.position)
}

func (p *parser) addErrAt(err error, pos position) {
	var buf bytes.Buffer
	if p.filename != "" {
		buf.WriteString(p.filename)
	}
	if buf.Len() > 0 {
		buf.WriteString(":")
	}
	buf.WriteString(fmt.Sprintf("%d:%d (%d)", pos.line, pos.col, pos.offset))
	if len(p.rstack) > 0 {
		if buf.Len() > 0 {
			buf.WriteString(": ")
		}
		rule := p.rstack[len(p.rstack)-1]
		if rule.displayName != "" {
			buf.WriteString("rule " + rule.displayName)
		} else {
			buf.WriteString("rule " + rule.name)
		}
	}
	pe := &parserError{Inner: err, prefix: buf.String()}
	p.errs.add(pe)
}

// read advances the parser to the next rune.
func (p *parser) read() {
	p.pt.offset += p.pt.w
	rn, n := utf8.DecodeRune(p.data[p.pt.offset:])
	p.pt.rn = rn
	p.pt.w = n
	p.pt.col++
	if rn == '\n' {
		p.pt.line++
		p.pt.col = 0
	}

	if rn == utf8.RuneError {
		if n > 0 {
			p.addErr(errInvalidEncoding)
		}
	}
}

// restore parser position to the savepoint pt.
func (p *parser) restore(pt savepoint) {
	if p.debug {
		defer p.out(p.in("restore"))
	}
	if pt.offset == p.pt.offset {
		return
	}
	p.pt = pt
}

// get the slice of bytes from the savepoint start to the current position.
func (p *parser) sliceFrom(start savepoint) []byte {
	return p.data[start.position.offset:p.pt.position.offset]
}

func (p *parser) getMemoized(node interface{}) (resultTuple, bool) {
	if len(p.memo) == 0 {
		return resultTuple{}, false
	}
	m := p.memo[p.pt.offset]
	if len(m) == 0 {
		return resultTuple{}, false
	}
	res, ok := m[node]
	return res, ok
}

func (p *parser) setMemoized(pt savepoint, node interface{}, tuple resultTuple) {
	if p.memo == nil {
		p.memo = make(map[int]map[interface{}]resultTuple)
	}
	m := p.memo[pt.offset]
	if m == nil {
		m = make(map[interface{}]resultTuple)
		p.memo[pt.offset] = m
	}
	m[node] = tuple
}

func (p *parser) buildRulesTable(g *grammar) {
	p.rules = make(map[string]*rule, len(g.rules))
	for _, r := range g.rules {
		p.rules[r.name] = r
	}
}

func (p *parser) parse(g *grammar) (val interface{}, err error) {
	if len(g.rules) == 0 {
		p.addErr(errNoRule)
		return nil, p.errs.err()
	}

	// TODO : not super critical but this could be generated
	p.buildRulesTable(g)

	if p.recover {
		// panic can be used in action code to stop parsing immediately
		// and return the panic as an error.
		defer func() {
			if e := recover(); e != nil {
				if p.debug {
					defer p.out(p.in("panic handler"))
				}
				val = nil
				switch e := e.(type) {
				case error:
					p.addErr(e)
				default:
					p.addErr(fmt.Errorf("%v", e))
				}
				err = p.errs.err()
			}
		}()
	}

	// start rule is rule [0]
	p.read() // advance to first rune
	val, ok := p.parseRule(g.rules[0])
	if !ok {
		if len(*p.errs) == 0 {
			// make sure this doesn't go out silently
			p.addErr(errNoMatch)
		}
		return nil, p.errs.err()
	}
	return val, p.errs.err()
}

func (p *parser) parseRule(rule *rule) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRule " + rule.name))
	}

	if p.memoize {
		res, ok := p.getMemoized(rule)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
	}

	start := p.pt
	p.rstack = append(p.rstack, rule)
	p.pushV()
	val, ok := p.parseExpr(rule.expr)
	p.popV()
	p.rstack = p.rstack[:len(p.rstack)-1]
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}

	if p.memoize {
		p.setMemoized(start, rule, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseExpr(expr interface{}) (interface{}, bool) {
	var pt savepoint
	var ok bool

	if p.memoize {
		res, ok := p.getMemoized(expr)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
		pt = p.pt
	}

	p.exprCnt++
	var val interface{}
	switch expr := expr.(type) {
	case *actionExpr:
		val, ok = p.parseActionExpr(expr)
	case *andCodeExpr:
		val, ok = p.parseAndCodeExpr(expr)
	case *andExpr:
		val, ok = p.parseAndExpr(expr)
	case *anyMatcher:
		val, ok = p.parseAnyMatcher(expr)
	case *charClassMatcher:
		val, ok = p.parseCharClassMatcher(expr)
	case *choiceExpr:
		val, ok = p.parseChoiceExpr(expr)
	case *labeledExpr:
		val, ok = p.parseLabeledExpr(expr)
	case *litMatcher:
		val, ok = p.parseLitMatcher(expr)
	case *notCodeExpr:
		val, ok = p.parseNotCodeExpr(expr)
	case *notExpr:
		val, ok = p.parseNotExpr(expr)
	case *oneOrMoreExpr:
		val, ok = p.parseOneOrMoreExpr(expr)
	case *ruleRefExpr:
		val, ok = p.parseRuleRefExpr(expr)
	case *seqExpr:
		val, ok = p.parseSeqExpr(expr)
	case *zeroOrMoreExpr:
		val, ok = p.parseZeroOrMoreExpr(expr)
	case *zeroOrOneExpr:
		val, ok = p.parseZeroOrOneExpr(expr)
	default:
		panic(fmt.Sprintf("unknown expression type %T", expr))
	}
	if p.memoize {
		p.setMemoized(pt, expr, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseActionExpr(act *actionExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseActionExpr"))
	}

	start := p.pt
	val, ok := p.parseExpr(act.expr)
	if ok {
		p.cur.pos = start.position
		p.cur.text = p.sliceFrom(start)
		actVal, err := act.run(p)
		if err != nil {
			p.addErrAt(err, start.position)
		}
		val = actVal
	}
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}
	return val, ok
}

func (p *parser) parseAndCodeExpr(and *andCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndCodeExpr"))
	}

	ok, err := and.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, ok
}

func (p *parser) parseAndExpr(and *andExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(and.expr)
	p.popV()
	p.restore(pt)
	return nil, ok
}

func (p *parser) parseAnyMatcher(any *anyMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAnyMatcher"))
	}

	if p.pt.rn != utf8.RuneError {
		start := p.pt
		p.read()
		return p.sliceFrom(start), true
	}
	return nil, false
}

func (p *parser) parseCharClassMatcher(chr *charClassMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseCharClassMatcher"))
	}

	cur := p.pt.rn
	// can't match EOF
	if cur == utf8.RuneError {
		return nil, false
	}
	start := p.pt
	if chr.ignoreCase {
		cur = unicode.ToLower(cur)
	}

	// try to match in the list of available chars
	for _, rn := range chr.chars {
		if rn == cur {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of ranges
	for i := 0; i < len(chr.ranges); i += 2 {
		if cur >= chr.ranges[i] && cur <= chr.ranges[i+1] {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of Unicode classes
	for _, cl := range chr.classes {
		if unicode.Is(cl, cur) {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	if chr.inverted {
		p.read()
		return p.sliceFrom(start), true
	}
	return nil, false
}

func (p *parser) parseChoiceExpr(ch *choiceExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseChoiceExpr"))
	}

	for _, alt := range ch.alternatives {
		p.pushV()
		val, ok := p.parseExpr(alt)
		p.popV()
		if ok {
			return val, ok
		}
	}
	return nil, false
}

func (p *parser) parseLabeledExpr(lab *labeledExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLabeledExpr"))
	}

	p.pushV()
	val, ok := p.parseExpr(lab.expr)
	p.popV()
	if ok && lab.label != "" {
		m := p.vstack[len(p.vstack)-1]
		m[lab.label] = val
	}
	return val, ok
}

func (p *parser) parseLitMatcher(lit *litMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLitMatcher"))
	}

	start := p.pt
	for _, want := range lit.val {
		cur := p.pt.rn
		if lit.ignoreCase {
			cur = unicode.ToLower(cur)
		}
		if cur != want {
			p.restore(start)
			return nil, false
		}
		p.read()
	}
	return p.sliceFrom(start), true
}

func (p *parser) parseNotCodeExpr(not *notCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotCodeExpr"))
	}

	ok, err := not.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, !ok
}

func (p *parser) parseNotExpr(not *notExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(not.expr)
	p.popV()
	p.restore(pt)
	return nil, !ok
}

func (p *parser) parseOneOrMoreExpr(expr *oneOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseOneOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			if len(vals) == 0 {
				// did not match once, no match
				return nil, false
			}
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseRuleRefExpr(ref *ruleRefExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRuleRefExpr " + ref.name))
	}

	if ref.name == "" {
		panic(fmt.Sprintf("%s: invalid rule: missing name", ref.pos))
	}

	rule := p.rules[ref.name]
	if rule == nil {
		p.addErr(fmt.Errorf("undefined rule: %s", ref.name))
		return nil, false
	}
	return p.parseRule(rule)
}

func (p *parser) parseSeqExpr(seq *seqExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseSeqExpr"))
	}

	var vals []interface{}

	pt := p.pt
	for _, expr := range seq.exprs {
		val, ok := p.parseExpr(expr)
		if !ok {
			p.restore(pt)
			return nil, false
		}
		vals = append(vals, val)
	}
	return vals, true
}

func (p *parser) parseZeroOrMoreExpr(expr *zeroOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseZeroOrOneExpr(expr *zeroOrOneExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrOneExpr"))
	}

	p.pushV()
	val, _ := p.parseExpr(expr.expr)
	p.popV()
	// whether it matched or not, consider it a match
	return val, true
}

func rangeTable(class string) *unicode.RangeTable {
	if rt, ok := unicode.Categories[class]; ok {
		return rt
	}
	if rt, ok := unicode.Properties[class]; ok {
		return rt
	}
	if rt, ok := unicode.Scripts[class]; ok {
		return rt
	}

	// cannot happen
	panic(fmt.Sprintf("invalid Unicode class: %s", class))
}

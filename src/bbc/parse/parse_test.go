package parse

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"bbc/ast"
)

type testCase struct {
	filename string
	source   string
	expected *ast.File
}

/* This set of tests is kind of a mess, since we're deep comparing the entire AST... */
var tests = []testCase{
	{
		filename: "in_file",
		source: `type SomeFrame frame<100, Var8> struct {
	SomeInt int8
	AnotherInt uint24
}`,
		expected: &ast.File{
			Name: "in_file",
			Scope: &ast.Scope{
				S: []ast.Node{
					&ast.Frame{
						Name:   "SomeFrame",
						Number: 100,
						Size:   ast.SzVar8,
						Object: &ast.Struct{
							Name: "AnonStruct_X",
							Scope: &ast.Scope{
								S: []ast.Node{
									&ast.Field{
										Name: "SomeInt",
										Type: &ast.IntegerType{
											Signed:    true,
											Bitsize:   8,
											Modifiers: nil,
										},
									},
									&ast.Field{
										Name: "AnotherInt",
										Type: &ast.IntegerType{
											Signed:    false,
											Bitsize:   24,
											Modifiers: nil,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},

	{
		filename: "in_file",
		source: `type SomeFrame frame<100, Var16> struct {
	SomeInt int8
	AnotherInt uint24
}`,
		expected: &ast.File{
			Name: "in_file",
			Scope: &ast.Scope{
				S: []ast.Node{
					&ast.Frame{
						Name:   "SomeFrame",
						Number: 100,
						Size:   ast.SzVar16,
						Object: &ast.Struct{
							Name: "AnonStruct_X",
							Scope: &ast.Scope{
								S: []ast.Node{
									&ast.Field{
										Name: "SomeInt",
										Type: &ast.IntegerType{
											Signed:    true,
											Bitsize:   8,
											Modifiers: nil,
										},
									},
									&ast.Field{
										Name: "AnotherInt",
										Type: &ast.IntegerType{
											Signed:    false,
											Bitsize:   24,
											Modifiers: nil,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},

	{
		filename: "in_file",
		source: `type SomeStruct struct {
	SomeInt int8
	AnotherInt uint24
}

type SomeFrame frame<100, Fixed> SomeStruct`,
		expected: &ast.File{
			Name: "in_file",
			Scope: &ast.Scope{
				S: []ast.Node{
					&ast.Struct{
						Name: "SomeStruct",
						Scope: &ast.Scope{
							S: []ast.Node{
								&ast.Field{
									Name: "SomeInt",
									Type: &ast.IntegerType{
										Signed:    true,
										Bitsize:   8,
										Modifiers: nil,
									},
								},
								&ast.Field{
									Name: "AnotherInt",
									Type: &ast.IntegerType{
										Signed:    false,
										Bitsize:   24,
										Modifiers: nil,
									},
								},
							},
						},
					},
					&ast.Frame{
						Name:   "SomeFrame",
						Number: 100,
						Size:   ast.SzFixed,
						Object: &ast.DeclReference{
							DeclName: "SomeStruct",
							Object: &ast.Struct{
								Name: "SomeStruct",
								Scope: &ast.Scope{
									S: []ast.Node{
										&ast.Field{
											Name: "SomeInt",
											Type: &ast.IntegerType{
												Signed:    true,
												Bitsize:   8,
												Modifiers: nil,
											},
										},
										&ast.Field{
											Name: "AnotherInt",
											Type: &ast.IntegerType{
												Signed:    false,
												Bitsize:   24,
												Modifiers: nil,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},

	{
		filename: "in_file",
		source: `type SomeStruct struct {
	SomeInt int8
	AnotherInt uint24
}`,
		expected: &ast.File{
			Name: "in_file",
			Scope: &ast.Scope{
				S: []ast.Node{
					&ast.Struct{
						Name: "SomeStruct",
						Scope: &ast.Scope{
							S: []ast.Node{
								&ast.Field{
									Name: "SomeInt",
									Type: &ast.IntegerType{
										Signed:    true,
										Bitsize:   8,
										Modifiers: nil,
									},
								},
								&ast.Field{
									Name: "AnotherInt",
									Type: &ast.IntegerType{
										Signed:    false,
										Bitsize:   24,
										Modifiers: nil,
									},
								},
							},
						},
					},
				},
			},
		},
	},

	{
		filename: "in_file",
		source: `type OuterStruct struct {
	SomeStruct struct {
		SomeInt int8
		AnotherInt uint24
	}
}`,
		expected: &ast.File{
			Name: "in_file",
			Scope: &ast.Scope{
				S: []ast.Node{
					&ast.Struct{
						Name: "OuterStruct",
						Scope: &ast.Scope{
							S: []ast.Node{
								&ast.Field{
									Name: "SomeStruct",
									Type: &ast.Struct{
										Name: "AnonStruct_X",
										Scope: &ast.Scope{
											S: []ast.Node{
												&ast.Field{
													Name: "SomeInt",
													Type: &ast.IntegerType{
														Signed:    true,
														Bitsize:   8,
														Modifiers: nil,
													},
												},
												&ast.Field{
													Name: "AnotherInt",
													Type: &ast.IntegerType{
														Signed:    false,
														Bitsize:   24,
														Modifiers: nil,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},

	{
		filename: "in_file",
		source: `type SomeStruct struct {
	Field uint32
}

type AnotherStruct struct {
	Field uint32
}`,
		expected: &ast.File{
			Name: "in_file",
			Scope: &ast.Scope{
				S: []ast.Node{
					&ast.Struct{
						Name: "SomeStruct",
						Scope: &ast.Scope{
							S: []ast.Node{
								&ast.Field{
									Name: "Field",
									Type: &ast.IntegerType{
										Signed:    false,
										Bitsize:   32,
										Modifiers: nil,
									},
								},
							},
						},
					},
					&ast.Struct{
						Name: "AnotherStruct",
						Scope: &ast.Scope{
							S: []ast.Node{
								&ast.Field{
									Name: "Field",
									Type: &ast.IntegerType{
										Signed:    false,
										Bitsize:   32,
										Modifiers: nil,
									},
								},
							},
						},
					},
				},
			},
		},
	},

	{
		filename: "in_file",
		source: `type SomeStruct struct/**/{
    /* Priority is an integer between 02 * /
         0 Urgent request ie client needs resources for a new area
         1 Preload request pre-login client update
         2 Background request low priority */
	SomeInt int8 /*
  Multi-line comment
*/
	AnotherInt uint24 // single line comment
}`,
		expected: &ast.File{
			Name: "in_file",
			Scope: &ast.Scope{
				S: []ast.Node{
					&ast.Struct{
						Name: "SomeStruct",
						Scope: &ast.Scope{
							S: []ast.Node{
								&ast.Field{
									Name: "SomeInt",
									Type: &ast.IntegerType{
										Signed:    true,
										Bitsize:   8,
										Modifiers: nil,
									},
								},
								&ast.Field{
									Name: "AnotherInt",
									Type: &ast.IntegerType{
										Signed:    false,
										Bitsize:   24,
										Modifiers: nil,
									},
								},
							},
						},
					},
				},
			},
		},
	},

	{
		filename: "in_file",
		source: `type SomeStruct struct {
	SomeInt int8<IntLittleEndian, IntOffset128>
}`,
		expected: &ast.File{
			Name: "in_file",
			Scope: &ast.Scope{
				S: []ast.Node{
					&ast.Struct{
						Name: "SomeStruct",
						Scope: &ast.Scope{
							S: []ast.Node{
								&ast.Field{
									Name: "SomeInt",
									Type: &ast.IntegerType{
										Signed:    true,
										Bitsize:   8,
										Modifiers: []string{"IntLittleEndian", "IntOffset128"},
									},
								},
							},
						},
					},
				},
			},
		},
	},

	{
		filename: "in_file",
		source: `type SomeStruct struct {
	SomeInt int8<IntLittleEndian, IntOffset128>[10]
}`,
		expected: &ast.File{
			Name: "in_file",
			Scope: &ast.Scope{
				S: []ast.Node{
					&ast.Struct{
						Name: "SomeStruct",
						Scope: &ast.Scope{
							S: []ast.Node{
								&ast.Field{
									Name: "SomeInt",
									Type: &ast.ArrayType{
										Object: &ast.IntegerType{
											Signed:    true,
											Bitsize:   8,
											Modifiers: []string{"IntLittleEndian", "IntOffset128"},
										},
										Length: &ast.StaticLength{
											Length: 10,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},

	{
		filename: "in_file",
		source: `type SomeStruct struct {
	SomeString string[256]
}`,
		expected: &ast.File{
			Name: "in_file",
			Scope: &ast.Scope{
				S: []ast.Node{
					&ast.Struct{
						Name: "SomeStruct",
						Scope: &ast.Scope{
							S: []ast.Node{
								&ast.Field{
									Name: "SomeString",
									Type: &ast.ArrayType{
										Object: &ast.StringBaseType{},
										Length: &ast.StaticLength{
											Length: 256,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},

	{
		filename: "in_file",
		source: `type SomeStruct struct {
	SomeBuffer byte[256]
}`,
		expected: &ast.File{
			Name: "in_file",
			Scope: &ast.Scope{
				S: []ast.Node{
					&ast.Struct{
						Name: "SomeStruct",
						Scope: &ast.Scope{
							S: []ast.Node{
								&ast.Field{
									Name: "SomeBuffer",
									Type: &ast.ArrayType{
										Object: &ast.ByteBaseType{},
										Length: &ast.StaticLength{
											Length: 256,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},

	{
		filename: "in_file",
		source: `type SomeStruct struct {
	Field uint32
}

type AnotherStruct struct {
	Field SomeStruct
}`,
		expected: &ast.File{
			Name: "in_file",
			Scope: &ast.Scope{
				S: []ast.Node{
					&ast.Struct{
						Name: "SomeStruct",
						Scope: &ast.Scope{
							S: []ast.Node{
								&ast.Field{
									Name: "Field",
									Type: &ast.IntegerType{
										Signed:    false,
										Bitsize:   32,
										Modifiers: nil,
									},
								},
							},
						},
					},
					&ast.Struct{
						Name: "AnotherStruct",
						Scope: &ast.Scope{
							S: []ast.Node{
								&ast.Field{
									Name: "Field",
									Type: &ast.DeclReference{
										DeclName: "SomeStruct",
										Object: &ast.Struct{
											Name: "SomeStruct",
											Scope: &ast.Scope{
												S: []ast.Node{
													&ast.Field{
														Name: "Field",
														Type: &ast.IntegerType{
															Signed:    false,
															Bitsize:   32,
															Modifiers: nil,
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},
}

func dump(file *ast.File) string {
	astStr, _ := json.Marshal(file)
	return fmt.Sprintf("%v", string(astStr))
}

func TestParser(t *testing.T) {
	yyDebug = 1
	for _, tc := range tests {
		t.Logf("Testing: \n%v", tc.source)
		ast, errors := Parse(tc.filename, tc.source)
		if len(errors) > 0 {
			t.Error(errors)
		}
		if !reflect.DeepEqual(ast, tc.expected) {
			t.Errorf("ast didn't match.\nGot:\t\t%v\nExpected:\t%v", dump(ast), dump(tc.expected))
		}
	}
}

func TestParseErrors(t *testing.T) {
	yyDebug = 1

	// Test that the lexer logs errors correctly
	_, errors := Parse("in_file", "^")
	if len(errors) != 1 {
		t.Errorf("expected 1 error, got %v", len(errors))
	}

	// Test that the parser logs errors correctly
	_, errors = Parse("in_file", "invalid struct }")
	if len(errors) != 1 {
		t.Errorf("expected 1 error, got %v", len(errors))
	}
}

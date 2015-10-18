%{
package parse

import (
	"bbc/ast"
)

%}

%union {
    nsl     []ast.Node
    n       ast.Node
    ival    int
    sval    string
    svalarr []string
    length  ast.LengthSpec
    size    ast.FrameSize
}

%token '{' '}' '[' ']' '<' '>' ','

%token tWhitespace

%token <sval> tIdentifier
%token <ival> tNumber

%token tStruct tType tFrame tBitStruct
%token tFrameFixed tFrameVar8 tFrameVar16
%token <n> tStringType tByteType tBitsType
%token <n> tIntegerType
%token <sval> tIntegerFlag

%token tEOL

%type <size> frame_size
%type <length> array_spec
%type <n> type int_type string_type bytes_type bits_type
%type <n> field
%type <n> field_list
%type <n> scope
%type <n> reference
%type <n> frame struct anon_struct decl
%type <n> decl_list
%type <svalarr> int_modifiers

%start file

%%

file
    : ws✳ decl_list ws✳
      { yylex.(*Lexer).Ast().Scope = $2.(*ast.Scope) }
    ;

decl_list
    : ws✳ { $$ = ast.NewScope() }
    | decl_list ws✳ decl ws✳
      { $1.(*ast.Scope).Add($3) }
    ;

decl
	: tType struct
      {
          $$ = $2
          yylex.(*Lexer).AddDecl($$)
      }
    | tType frame
      {
          $$ = $2
          yylex.(*Lexer).AddDecl($$)
      }
	;

frame
	: tIdentifier tFrame '<' tNumber ',' frame_size '>' type
      {
          $$ = &ast.Frame{
              Name: $1,
              Number: $4,
              Size: $6,
              Object: $8,
          }
      }
	;

frame_size
	: tFrameFixed
      { $$ = ast.SzFixed }
	| tFrameVar8
      { $$ = ast.SzVar8 }
	| tFrameVar16
      { $$ = ast.SzVar16 }
	;

struct
	: tIdentifier anon_struct
      {
          if _, ok := $2.(*ast.BitStruct); ok {
              $2.(*ast.BitStruct).Name = $1
          } else {
              $2.(*ast.Struct).Name = $1
          }
          $$ = $2
      }
    ;

anon_struct
    : tStruct scope
      {
          $$ = &ast.Struct{
	          Name: yylex.(*Lexer).NameAnonStruct(),
              Scope: $2.(*ast.Scope),
          }
      }
    | tBitStruct scope
      {
          $$ = &ast.BitStruct{
	          Name: yylex.(*Lexer).NameAnonStruct(),
              Scope: $2.(*ast.Scope),
          }
      }
    ;

scope
    : '{' field_list '}'
      { $$ = $2 }
    ;

field_list
    : ws✳ { $$ = ast.NewScope() }
    | field_list field
      { $1.(*ast.Scope).Add($2.(ast.Node)) }
    ;

field
    : tIdentifier type
      {
          $$ = &ast.Field{
              Name: $1,
              Type: $2.(ast.Node),
          }
      }
    ;

int_type
    : tIntegerType
    | tIntegerType '<' int_modifiers '>'
      {
          $1.(*ast.IntegerType).Modifiers = $3
          $$ = $1
      }
    ;

int_modifiers
	: ws✳ { $$ = make([]string, 0) }
    | tIntegerFlag
      { $$ = append($$, $1) }
    | int_modifiers ',' tIntegerFlag
      { $$ = append($$, $3) }
    ;

reference
	: tIdentifier
      {
	      $$ = &ast.DeclReference{
			  DeclName: $1,
	      }
      }
    ;

type
    : int_type
    | bits_type
    | string_type
    | bytes_type
    | anon_struct
    | reference
    | type array_spec
      {
          $$ = &ast.ArrayType{
	          Object: $1,
              Length: $2,
          }
      }
    ;

bytes_type
	: tByteType
      { $$ = &ast.ByteBaseType{} }
	;

string_type
	: tStringType
      { $$ = &ast.StringBaseType{} }
	;

bits_type
	: tBitsType
      { $$ = &ast.BitsType{Count: 1} }
	| tBitsType array_spec
      {
          if _, ok := $2.(*ast.StaticLength); !ok {
              yylex.(*Lexer).Error("bit fields require a constant length expression")
          } else {
              $$ = &ast.BitsType{ Count: $2.(*ast.StaticLength).Length }
          }
      }
	;

array_spec
	: '[' tNumber ']'
      {
          $$ = &ast.StaticLength{
	          Length: $2,
          }
      }

ws
	: tWhitespace
	;

ws＋
	: ws
	| ws ws＋

ws✳
	: /* empty */
	| ws＋

%%

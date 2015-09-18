%{
package parse

import (
	"framecc/ast"
)

%}

%union {
    nsl []ast.Node
    n ast.Node
    ival int
    sval string
}

%token '{' '}' '[' ']' '<' '>' ','

%token tWhitespace

%token tIdentifier
%token tNumber

%token tStruct
%token tStringType
%token tIntegerType
%token tIntegerFlag

%token tEOL

%start file

%%

file
    : ws✳ decl_list ws✳
      { yylex.(*Lexer).parseResult.(*ast.File).Scope = $2.n.(*ast.Scope) }
    ;

decl_list
    : ws✳ { $$.n = ast.NewScope() }
    | decl_list ws✳ struct ws✳
      { $1.n.(*ast.Scope).Add($3.n.(*ast.Struct)) }
    ;

struct
	: tIdentifier anon_struct
      {
          $2.n.(*ast.Struct).Name = $1.sval
          $$ = $2
      }
    ;

anon_struct
    : tStruct scope
      {
          $$.n = &ast.Struct{
              Name: "AnonStruct_X",
              Scope: $2.n.(*ast.Scope),
          }
      }
    ;

scope
    : '{' field_list '}'
      { $$.n = $2.n }
    ;

field_list
    : ws✳ { $$.n = ast.NewScope() }
    | field_list ws✳ field ws✳
    { $1.n.(*ast.Scope).Add($3.n.(ast.Node)) }
    ;

field
    : tIdentifier type
      {
          $$.n = &ast.Field{
              Name: $1.sval,
              Type: $2.n.(ast.Node),
          }
      }
    ;

int_type
    : tIntegerType
//  | tIntegerType '<' int_modifiers '>'
    ;

type
    : int_type
    | anon_struct
    ;

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

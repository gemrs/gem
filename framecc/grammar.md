Start <- Declaration+ EOF

Declaration <- Identifier _ Type EOL

Type <- ArrayType / BaseType
ArrayType <- BaseType '[' Expr ']'
BaseType <- IntegerType / StringType / StructType / FrameType / Reference
StructType <- 'struct' '{' Declaration+ '}'
IntegerType <- 'u'? !_ 'int' !_ ('8' / '16' / '32' / '64') Flags?
FrameType <- 'frame' Flags Type
StringType <- 'string' '[' Expr ']'
Reference <- Identifier

//TODO: Expand Expr to do arithmetic?
Expr <- IntegerExpr / Reference
IntegerExpr <- '0x' [0-9a-fA-F]+
            /  [0-9]+

Flags <- '<' (Flag (',' &Flag)?)+ '>'
Flag <- [a-zA-Z0-9_]+

Identifier <- [a-zA-Z][a-zA-Z0-9_]+

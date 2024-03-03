grammar Tiny_Bard;

// Program structure
program : program identifier ':' { ... } .

// Program parts
program_header : 'program' identifier ':' ;
program_body : (declaration ';')* 'begin' (statement ';')* 'end program' ;

// Declarations
declaration : ( 'global' )? (procedure_declaration | variable_declaration) ;

// Procedure declaration
procedure_declaration : procedure_header procedure_body ;

procedure_header : 'procedure' identifier ':' type_mark '(' parameter_list? ')' ;

parameter_list : parameter (',' parameter_list)? ;

parameter : variable_declaration ;

procedure_body : (declaration ';')* 'begin' (statement ';')* 'end procedure' ;

// Variable declaration
variable_declaration : 'variable' identifier ':' type_mark ( '[' bound ']' )? ;

// Type definition
type_mark : 'integer' | 'float' | 'string' | 'bool' ;

// Variable bounds
bound : number ;

// Statements
statement : assignment_statement | if_statement | loop_statement | return_statement ;

// Procedure call
procedure_call : identifier '(' argument_list? ')' ;

// Assignment statement
assignment_statement : destination ':=' expression ;

destination : identifier ( '[' expression ']' )? ;

// If statement
if_statement : 'if' '(' expression ')' 'then' (statement ';')* ('else' (statement ';')*)? 'end if' ;

// Loop statement
loop_statement : 'for' '(' assignment_statement ';' expression ')' (statement ';')* 'end for' ;

// Return statement
return_statement : 'return' expression ;

// Identifiers
identifier : [a-zA-Z][a-zA-Z0-9_]* ;

// Expressions
expression : term
           | expression op1 term
           | 'not' op1 term
           ;

op1 : '+' | '-' ;

// Relational expressions
relation : term
          | relation_op1 term
          | relation_op2 term
          | relation_op3 term
          | relation_op4 term
          | relation_op5 term
          | relation_op6 term
          ;

relation_op1 : '<' ;
relation_op2 : '>' ;
relation_op3 : '<=' ;
relation_op4 : '>=' ;
relation_op5 : '==' ;
relation_op6 : '!=' ;

// Terms
term : factor
      | term op2 factor
      ;

op2 : '*' | '/' ;

// Factors
factor : '(' expression ')'
        | procedure_call
        | '-' factor
        | number
        | string
        | 'true'
        | 'false' ;

// Names (variables or procedure calls)
name : identifier ( '[' expression ']' )? ;

// Argument list
argument_list : expression (',' argument_list)? ;

// Numbers
number : [0-9][0-9_]*[.[0-9_]*] ;

// Strings
string : '"' .* '"' ;

// Whitespace
WS : [ \t\r\n]+ -> skip ;

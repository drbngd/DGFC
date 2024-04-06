grammar example;

set_of_stmts : ( stmt SEMI_COLON )* ;

stmt: if_stmt
    | while_stmt
    | assignment_stmt
    ;

if_stmt: IF_KW cond_expr THEN_KW set_of_stmts (ELSE_KW set_of_stmts)? ;

while_stmt: WHILE_KW cond_expr DO_KW set_of_stmts END_KW WHILE_KW ;

assignment_stmt: ID ( L_BRAKET expr R_BRAKET )? COLON_EQUAL expr ;

cond_expr: expr;

expr: expr ADD_OP term
    | term ;

term: term MULT_OP factor
    | factor ;

factor: L_PAREN expr R_PAREN
    | ID ( L_BRAKET expr R_BRAKET )?
    | NUM
    ;

ADD_OP: '+' | '-' ;

MULT_OP: '*' | '/' ;

IF_KW: 'if' ;
THEN_KW: 'then' ;
ELSE_KW: 'else' ;
WHILE_KW: 'while' ;
DO_KW: 'do' ;
END_KW: 'end' ;
COLON_EQUAL: ':=' ;

L_BRAKET: '[' ;
R_BRAKET: ']' ;

L_PAREN: '(' ;
R_PAREN: ')' ;

SEMI_COLON: ';' ;

ID : [a-z]+ ;             // match lower-case identifiers
NUM: [0-9]+ ;
WS : [ \t\r\n]+ -> skip ; // skip spaces, tabs, newlines
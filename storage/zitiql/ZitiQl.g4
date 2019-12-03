grammar ZitiQl;

// Allow case insensitive token definitions
fragment A : [aA];
fragment B : [bB];
fragment C : [cC];
fragment D : [dD];
fragment E : [eE];
fragment F : [fF];
fragment G : [gG];
fragment H : [hH];
fragment I : [iI];
fragment J : [jJ];
fragment K : [kK];
fragment L : [lL];
fragment M : [mM];
fragment N : [nN];
fragment O : [oO];
fragment P : [pP];
fragment Q : [qQ];
fragment R : [rR];
fragment S : [sS];
fragment T : [tT];
fragment U : [uU];
fragment V : [vV];
fragment W : [wW];
fragment X : [xX];
fragment Y : [yY];
fragment Z : [zZ];

WS: [ \n\t\r];

LPAREN: '(';
RPAREN: ')';
LBRACKET: '[';
RBRACKET: ']';

AND: A N D;
OR: O R;

LT: '<' '='?;
GT: '>' '='?;
EQ: '!'? '=';
CONTAINS: (N O T WS+)? C O N T A I N S;


IN: (N O T WS)? I N;
BETWEEN: (N O T WS+)? B E T W E E N;

BOOL: TRUE | FALSE;

fragment TRUE: T R U E;
fragment  FALSE: F A L S E;

DATETIME: 'datetime(' WS* RFC3339_DATE_TIME WS* ')';

fragment FULL_DATE: YEAR '-' MONTH '-' DAY;
fragment FULL_TIME: HOUR  ':' MINUTE ':' SECOND ('.' SECOND_FRACTION)?  TIME_NUM_OFFSET;
fragment YEAR: INT+;
fragment MONTH: ('0'[1-9]|'1'[012]);
fragment DAY: ('0'[1-9]|[12][0-9]|'3'[01]);
fragment HOUR: ([01][0-9]|'2'[0-3]);
fragment MINUTE: ([0-5][0-9]);
fragment SECOND: ([0-5][0-9]|'60');
fragment SECOND_FRACTION: [0-9]+;
fragment TIME_NUM_OFFSET: (Z|(('+'|'-') TIME_NUM_OFFSET_HOUR ':' TIME_NUM_OFFSET_MINUTE));
fragment TIME_NUM_OFFSET_HOUR: ([01][0-9]|'2'[0-3]);
fragment TIME_NUM_OFFSET_MINUTE: ([0-5][0-9]);

ALL_OF: A L L O F;
NONE_OF: N O N E O F;
ANY_OF: A N Y O F;

STRING: '"' (ESC | SAFECODEPOINT)* '"';
NUMBER: '-'? INT ('.' [0-9] +)? EXP?;
NULL: N U L L;

ASC: A S C;
DESC: D E S C;

SORT: S O R T;
BY: B Y;
SKIP_ROWS: S K I P;
LIMIT_ROWS: L I M I T;
NONE: N O N E;

fragment INT: '0' | [1-9] [0-9]*;
fragment EXP: [Ee] [+\-]? INT;

//Allow escaping of form-feed, new line, line feed, and tab. No support for backspace and other escapables
fragment ESC: '\\' ["\\fnrt];
fragment IDENTSEP: '.';
IDENTIFIER:
      [A-Za-z] [A-Za-z_]* (IDENTSEP [A-Za-z] [A-Za-z_]*)*
    | '\'' [A-Za-z] [A-Za-z_]* (IDENTSEP [A-Za-z] [A-Za-z_]*)* '\'';

// Per RFC
RFC3339_DATE_TIME: FULL_DATE  T  FULL_TIME;

//No control characters
fragment SAFECODEPOINT: ~ ["\\\u0000-\u001F];

// Empty lists not supported as some RDBMs don't allow empty lists
string_array: LBRACKET WS* STRING (WS* ',' WS* STRING)* WS* RBRACKET;
number_array: LBRACKET WS* NUMBER (WS* ',' WS* NUMBER)* WS* RBRACKET;
datetime_array: LBRACKET WS* DATETIME (WS* ',' WS* DATETIME)* WS* RBRACKET;

start: WS* query* WS* EOF #End;

query: expression (WS+ sortBy)? (WS+ skip)? (WS+ limit)? #QueryStmt;

skip: SKIP_ROWS WS+ NUMBER #SkipExpr;

limit: LIMIT_ROWS WS+ (NONE|NUMBER) #LimitExpr;

sortBy: SORT WS+ BY WS+ sortField (WS* ',' WS* sortField)* #SortByExpr;

sortField: IDENTIFIER (WS+ (ASC | DESC))? #SortFieldExpr;

expression:
  operation #OperationOp
  | LPAREN WS* expression WS* RPAREN #Group
  | expression (WS+ AND WS+ expression)+ #AndConjunction
  | expression (WS+ OR WS+ expression)+ #OrConjunction
  | BOOL #BoolConst
  ;


operation:
    binary_lhs WS+ IN WS+ string_array #InStringArrayOp
  | binary_lhs WS+ IN WS+ number_array #InNumberArrayOp
  | binary_lhs WS+ IN WS+ datetime_array #InDatetimeArrayOp
  | binary_lhs WS+ BETWEEN WS+ NUMBER WS+ AND WS+ NUMBER # BetweenNumberOp
  | binary_lhs WS+ BETWEEN WS+ DATETIME WS+ AND WS+ DATETIME # BetweenDateOp
  | binary_lhs WS* LT WS* NUMBER # BinaryLessThanNumberOp
  | binary_lhs WS* LT WS* DATETIME # BinaryLessThanDatetimeOp
  | binary_lhs WS* GT WS* NUMBER # BinaryGreaterThanNumberOp
  | binary_lhs WS* GT WS* DATETIME# BinaryGreaterThanDatetimeOp
  | binary_lhs WS* EQ WS* STRING #BinaryEqualToStringOp
  | binary_lhs WS* EQ WS* NUMBER #BinaryEqualToNumberOp
  | binary_lhs WS* EQ WS* DATETIME #BinaryEqualToDatetimeOp
  | binary_lhs WS* EQ WS* BOOL #BinaryEqualToBoolOp
  | binary_lhs WS* EQ WS* NULL #BinaryEqualToNullOp
  | binary_lhs WS* CONTAINS WS+ (STRING|NUMBER) #BinaryContainsOp
  ;

binary_lhs: IDENTIFIER | set_function;

set_function:
    ALL_OF LPAREN IDENTIFIER RPAREN #SetFunction
  | NONE_OF LPAREN IDENTIFIER RPAREN #SetFunction
  | ANY_OF LPAREN IDENTIFIER RPAREN #SetFunction
  ;
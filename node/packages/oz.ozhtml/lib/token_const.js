var NAMES = {};
var EOF = 1;
NAMES[EOF] = "EOF";
var TAGHEADOPEN = 2;
NAMES[TAGHEADOPEN] = "TAGHEADOPEN";
var TAGHEADCLOSE = 4;
NAMES[TAGHEADCLOSE] = "TAGHEADCLOSE";
var TAGBODY = 5;
NAMES[TAGBODY] = "TAGBODY";
var WORD = 6;
NAMES[WORD] = "WORD";
var WHITESPACE = 7;
NAMES[WHITESPACE] = "WHITESPACE";
var EQUAL = 8;
NAMES[EQUAL] = "EQUAL";
var STRING = 9;
NAMES[STRING] = "STRING";
var MULTIPLICATION = 10;
NAMES[MULTIPLICATION] = "MULTIPLICATION";
var PARENTHESISOPEN = 11;
NAMES[PARENTHESISOPEN] = "PARENTHESISOPEN";
var PARENTHESISCLOSE = 12;
NAMES[PARENTHESISCLOSE] = "PARENTHESISCLOSE";
var BRACKETOPEN = 13;
NAMES[BRACKETOPEN] = "BRACKETOPEN";
var BRACKETCLOSE = 14;
NAMES[BRACKETCLOSE] = "BRACKETCLOSE";
var DOUBLEDOT = 16;
NAMES[DOUBLEDOT] = "DOUBLEDOT";
var DOT = 17;
NAMES[DOT] = "DOT";
var TAGCOMMENT = 18;
NAMES[TAGCOMMENT] = "TAGCOMMENT";
var NUMERIC = 19;
NAMES[NUMERIC] = "NUMERIC";
module.exports = {
  NAMES: NAMES,
  EOF: EOF,
  TAGHEADOPEN: TAGHEADOPEN,
  TAGHEADCLOSE: TAGHEADCLOSE,
  TAGBODY: TAGBODY,
  WORD: WORD,
  WHITESPACE: WHITESPACE,
  EQUAL: EQUAL,
  STRING: STRING,
  MULTIPLICATION: MULTIPLICATION,
  PARENTHESISOPEN: PARENTHESISOPEN,
  PARENTHESISCLOSE: PARENTHESISCLOSE,
  BRACKETOPEN: BRACKETOPEN,
  BRACKETCLOSE: BRACKETCLOSE,
  DOUBLEDOT: DOUBLEDOT,
  DOT: DOT,
  TAGCOMMENT: TAGCOMMENT,
  NUMERIC: NUMERIC
};
//# sourceMappingURL=token_const.js.map
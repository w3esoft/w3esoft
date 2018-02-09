const NAMES={};
const EOF                         = 1;
NAMES[EOF]                  = "EOF";
const TAGHEADOPEN                 = 2;
NAMES[TAGHEADOPEN]          = "TAGHEADOPEN";
const TAGHEADCLOSE                = 4;
NAMES[TAGHEADCLOSE]         = "TAGHEADCLOSE";
const TAGBODY                     = 5;
NAMES[TAGBODY]              = "TAGBODY";
const WORD                        = 6;
NAMES[WORD]                 = "WORD";
const WHITESPACE                  = 7;
NAMES[WHITESPACE]           = "WHITESPACE";
const EQUAL                       = 8;
NAMES[EQUAL]                = "EQUAL";
const STRING                      = 9;
NAMES[STRING]               = "STRING";
const MULTIPLICATION              = 10;
NAMES[MULTIPLICATION]       = "MULTIPLICATION";
const PARENTHESISOPEN             = 11;
NAMES[PARENTHESISOPEN]      = "PARENTHESISOPEN";
const PARENTHESISCLOSE            = 12;
NAMES[PARENTHESISCLOSE]     = "PARENTHESISCLOSE";
const BRACKETOPEN                 = 13;
NAMES[BRACKETOPEN]          = "BRACKETOPEN";
const BRACKETCLOSE                = 14;
NAMES[BRACKETCLOSE]         = "BRACKETCLOSE";
const DOUBLEDOT                   = 16;
NAMES[DOUBLEDOT]            = "DOUBLEDOT";
const DOT                         = 17;
NAMES[DOT]                  = "DOT";
const TAGCOMMENT                  = 18;
NAMES[TAGCOMMENT]           = "TAGCOMMENT";
const NUMERIC                     = 19;
NAMES[NUMERIC]              = "NUMERIC";
module.exports = {
    NAMES,
    EOF,
    TAGHEADOPEN,
    TAGHEADCLOSE,
    TAGBODY,
    WORD,
    WHITESPACE,
    EQUAL,
    STRING,
    MULTIPLICATION,
    PARENTHESISOPEN,
    PARENTHESISCLOSE,
    BRACKETOPEN,
    BRACKETCLOSE,
    DOUBLEDOT,
    DOT,
    TAGCOMMENT,
    NUMERIC
}
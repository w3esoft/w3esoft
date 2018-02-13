"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.Lexer = void 0;

var tokenConst = _interopRequireWildcard(require("./token_const"));

var charConst = _interopRequireWildcard(require("./char_const"));

var _token = require("./token");

function _interopRequireWildcard(obj) { if (obj && obj.__esModule) { return obj; } else { var newObj = {}; if (obj != null) { for (var key in obj) { if (Object.prototype.hasOwnProperty.call(obj, key)) { var desc = Object.defineProperty && Object.getOwnPropertyDescriptor ? Object.getOwnPropertyDescriptor(obj, key) : {}; if (desc.get || desc.set) { Object.defineProperty(newObj, key, desc); } else { newObj[key] = obj[key]; } } } } newObj.default = obj; return newObj; } }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } }

function _createClass(Constructor, protoProps, staticProps) { if (protoProps) _defineProperties(Constructor.prototype, protoProps); if (staticProps) _defineProperties(Constructor, staticProps); return Constructor; }

var Lexer =
/*#__PURE__*/
function () {
  function Lexer(input) {
    _classCallCheck(this, Lexer);

    this.CHAR = -1;
    this.input = input;
    this.tokens = [];
  }

  _createClass(Lexer, [{
    key: "tokenize",
    value: function tokenize() {
      var me = this;
      if (me.tokens.length) return me.tokens.pop();
      var char;

      if (me.CHAR !== -1) {
        char = me.CHAR;
        me.CHAR = -1;
      } else {
        char = me.input.read();
      }

      if (char === charConst.GREATER) {
        char = me.input.read();

        if (char === charConst.BACKSLASH) {
          return new _token.Token(tokenConst.TAGHEADCLOSE);
        } else if (char === charConst.BANG) {
          me.input.read();
          me.input.read();
          var ii = 0;
          var value = "";

          while (true) {
            char = me.input.read();
            value += String.fromCharCode(char);

            if (char === charConst.EOF) {
              break;
            } else if (ii === 0 && char === charConst.MINUS) {
              ii = 1;
            } else if (ii === 1 && char === charConst.MINUS) {
              ii = 2;
            } else if (ii === 2 && char === charConst.MINOR) {
              break;
            } else {
              ii = 0;
            }
          }

          value = value.substring(0, value.length - 3);
          return new _token.Token(tokenConst.TAGCOMMENT, value);
        } else {
          me.CHAR = char;
          return new _token.Token(tokenConst.TAGHEADOPEN);
        }
      } else if (char === charConst.BACKSLASH) {
        char = input.read();

        if (char === charConst.MINOR) {
          console.log(char);
        }

        return new _token.Token(tokenConst.TAGHEADCLOSE);
      } else if (char === charConst.MINOR) {
        var _value = "";
        char = input.read();

        while (char !== charConst.GREATER && char !== charConst.EOF) {
          _value += String.fromCharCode(char);
          char = input.read();
        }

        me.CHAR = char;
        return new _token.Token(tokenConst.TAGBODY, _value);
      } else if (char === charConst.PARENTHESISOPEN) {
        return new _token.Token(tokenConst.PARENTHESISOPEN);
      } else if (char === charConst.PARENTHESISCLOSE) {
        return new _token.Token(tokenConst.PARENTHESISCLOSE);
      } else if (char === charConst.EQUAL) {
        return new _token.Token(tokenConst.EQUAL);
      } else if (char === charConst.MULTIPLICATION) {
        return new _token.Token(tokenConst.MULTIPLICATION);
      } else if (char === charConst.BRACKETOPEN) {
        return new _token.Token(tokenConst.BRACKETOPEN);
      } else if (char === charConst.BRACKETClOSE) {
        return new _token.Token(tokenConst.BRACKETCLOSE);
      } else if (char === charConst.DOT) {
        return new _token.Token(tokenConst.DOT);
      } else if (char === charConst.DOUBLEDOT) {
        return new _token.Token(tokenConst.DOUBLEDOT);
      } else if (char === charConst.DOUBLEQUOTES) {
        var _value2 = "";
        char = input.read();

        while (char !== charConst.DOUBLEQUOTES && char !== charConst.EOF) {
          _value2 += String.fromCharCode(char);
          char = input.read();
        }

        return new _token.Token(tokenConst.STRING, _value2);
      } else if (Lexer.isWhiteSpace(char)) {
        while (Lexer.isWhiteSpace(char)) {
          char = input.read();
        }

        me.CHAR = char;
        return new _token.Token(tokenConst.WHITESPACE);
      } else if (Lexer.isWord(char)) {
        var _value3 = "";

        while (Lexer.isWord(char) || Lexer.isNumeric(char)) {
          _value3 += String.fromCharCode(char);
          char = input.read();
        }

        me.CHAR = char;
        return new _token.Token(tokenConst.WORD, _value3);
      } else if (Lexer.isNumeric(char)) {
        var _value4 = "";

        while (Lexer.isNumeric(char)) {
          _value4 += String.fromCharCode(char);
          char = input.read();
        }

        me.CHAR = char;
        return new _token.Token(tokenConst.NUMERIC, _value4);
      } else if (char === charConst.EOF) {
        return new _token.Token(tokenConst.EOF);
      } else {
        throw "do'nt know charcode " + char;
      }
    }
  }], [{
    key: "isWord",
    value: function isWord(char) {
      return 96 < char && 123 > char || 64 < char && 91 > char || char === charConst.MINUS || char === 95 || char === 35;
    }
  }, {
    key: "isNumeric",
    value: function isNumeric(char) {
      return 47 < char && 58 > char;
    }
  }, {
    key: "isWhiteSpace",
    value: function isWhiteSpace(char) {
      return char === 32 || char === 13 || char === 10;
    }
  }]);

  return Lexer;
}();

exports.Lexer = Lexer;
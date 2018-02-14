function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } }

function _createClass(Constructor, protoProps, staticProps) { if (protoProps) _defineProperties(Constructor.prototype, protoProps); if (staticProps) _defineProperties(Constructor, staticProps); return Constructor; }

var astConst = require("./ast_const");

var tokenConst = require("./token_const");

var _require = require("./ast"),
    Ast = _require.Ast;

var Parser =
/*#__PURE__*/
function () {
  function Parser() {
    _classCallCheck(this, Parser);
  }

  _createClass(Parser, [{
    key: "contructor",
    value: function contructor(lexer) {
      var me = this;
      me.lexer = lexer;
    }
  }, {
    key: "tokenize",
    value: function tokenize() {
      var me = this;

      while (true) {
        var tk = me.lexer.tokenize(); // console.log( (tk.id) +  "  " +tk.toString()  );

        if (tk.isNot(tokenConst.WHITESPACE)) return tk;
      }
    }
  }, {
    key: "pushToken",
    value: function pushToken(tk) {
      var me = this;
      me.lexer.tokens.push(tk);
    }
  }, {
    key: "tokenToAstValue",
    value: function tokenToAstValue(tk) {
      if (tk.is(tokenConst.TAGCOMMENT)) {
        return new Ast(astConst.TAGCOMMENT, tk.value);
      }

      return new Ast(astConst.VALUE, tk.value);
    }
  }, {
    key: "attrValue",
    value: function attrValue() {
      var me = this;
      var tk = tokenize();
      tk.expected([tokenConst.NUMERIC, tokenConst.STRING]);
      return me.tokenToAstValue(tk);
    }
  }, {
    key: "attrKey",
    value: function attrKey() {
      var me = this;
      var tk = tokenize();
      var fields = [];
      var isTemplate = false;
      var isGlobal = false;
      var closeAttr = [];

      if (tk.is(tokenConst.MULTIPLICATION)) {
        isTemplate = true;
        tk = tokenize();
      }

      var operation = [];

      while (true) {
        if (tk.is(tokenConst.PARENTHESISOPEN)) {
          closeAttr.push(tokenConst.PARENTHESISCLOSE);
          operation.push("OUTPUT");
          tk = tokenize();
        } else if (tk.is(tokenConst.BRACKETOPEN)) {
          operation.push("INPUT");
          closeAttr.push(tokenConst.BRACKETCLOSE);
          tk = tokenize();
        } else {
          break;
        }
      }

      var value = tk;
      tk.expected(tokenConst.WORD);
      tk = tokenize();

      if (tk.is(tokenConst.DOUBLEDOT)) {
        isGlobal = true;
        tk = tokenize();
        tk.expected(tokenConst.WORD);
        fields.push(me.tokenToAstValue(tk));
        tk = tokenize();
      }

      while (tk.is(tokenConst.DOT)) {
        tk = tokenize();
        tk.expected(tokenConst.WORD);
        fields.push(me.tokenToAstValue(tk));
        tk = tokenize();
      }

      while (closeAttr.length) {
        var index = closeAttr.pop();
        tk.expected(index);
        tk = tokenize();
      }

      me.pushToken(tk);
      return new Ast(astConst.ATTRKEY, {
        value: me.tokenToAstValue(value),
        fields: fields,
        isGlobal: isGlobal,
        isTemplate: isTemplate,
        operation: operation
      });
    }
  }, {
    key: "tagBody",
    value: function tagBody() {
      var me = this;
      var tk = tokenize();

      if (tk.is(tokenConst.TAGHEADOPEN)) {
        var tkName = tokenize();
        tkName.expected(tokenConst.WORD);
        var attrs = [];
        var body = [];

        while (true) {
          var _tk = tokenize();

          while (_tk.is([tokenConst.WORD, tokenConst.MULTIPLICATION, tokenConst.PARENTHESISOPEN, tokenConst.BRACKETOPEN])) {
            var _attr = {};
            me.pushToken(_tk);
            _attr.key = me.attrKey();
            _tk = tokenize();

            if (_tk.is(tokenConst.EQUAL)) {
              _attr.value = me.attrValue();
              _tk = tokenize();
            }

            attrs.push(new Ast(astConst.ATTR, _attr));
          }

          if (_tk.expected([tokenConst.TAGBODY, tokenConst.TAGHEADOPEN])) {
            me.pushToken(_tk);
            break;
          }
        }

        while (true) {
          tk = tokenize();

          if (tk.is(tokenConst.TAGHEADCLOSE)) {
            var tk1 = tk;
            tk.toString();
            var tk2 = tk = tokenize();
            tk.expected(tokenConst.WORD);

            if (!tk.is(tokenConst.WORD, tkName.value)) {
              me.pushToken(tk2);
              me.pushToken(tk1);
            }

            break;
          }

          me.pushToken(tk);
          var b = me.tagBody();
          body.push(b);
        }

        return new Ast(astConst.TAG, {
          name: me.tokenToAstValue(tkName),
          attrs: attrs,
          body: body
        });
      } else if (tk.is(tokenConst.TAGBODY)) {
        return me.tokenToAstValue(tk);
      } else if (tk.is(tokenConst.TAGCOMMENT)) {
        return me.tokenToAstValue(tk);
      } else {
        tk.unexpected(tk.tokenIndex);
      }
    }
  }, {
    key: "parse",
    value: function parse() {
      var me = this;
      var items = [];

      while (true) {
        var tk = me.tokenize();

        if (tk.is(tokenConst.EOF)) {
          console.log("Token.......................");
          break;
        } else {
          me.pushToken(tk);
          var b = me.tagBody();
          items.push(b);
        }
      }

      return new Ast(astConst.DOCUMENT, items);
    }
  }]);

  return Parser;
}();
//# sourceMappingURL=parser.js.map
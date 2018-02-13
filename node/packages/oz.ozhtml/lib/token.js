var tokenConst = require("./token_const");

var TOKEN_ID = 1;

var Token = function Token(tokenIndex, value) {
  var me = this;
  me.id = 100 + TOKEN_ID++;
  me.tokenIndex = tokenIndex;
  me.value = value;

  me.is = function (tokenIndex, value) {
    if (Array.isArray(tokenIndex)) {
      var b;
      var _iteratorNormalCompletion = true;
      var _didIteratorError = false;
      var _iteratorError = undefined;

      try {
        for (var _iterator = tokenIndex[Symbol.iterator](), _step; !(_iteratorNormalCompletion = (_step = _iterator.next()).done); _iteratorNormalCompletion = true) {
          var _iii = _step.value;
          b = me.is(_iii, value);

          if (b) {
            b = true;
            break;
          }
        }
      } catch (err) {
        _didIteratorError = true;
        _iteratorError = err;
      } finally {
        try {
          if (!_iteratorNormalCompletion && _iterator.return != null) {
            _iterator.return();
          }
        } finally {
          if (_didIteratorError) {
            throw _iteratorError;
          }
        }
      }

      return !!b;
    }

    var b1 = me.tokenIndex === tokenIndex;
    var b2 = me.value === value || value === undefined;
    return b1 && b2;
  };

  me.isNot = function (tokenIndex, value) {
    var me = this;
    return !me.is.apply(me, arguments);
  };

  me.expected = function (tokenIndex, value) {
    var me = this;
    var b = me.is.apply(me, arguments);

    if (!b) {
      var msg = "not expect token " + me.name + "  ";

      if (me.value) {
        msg += " value = " + me.value;
      }

      throw new Error(msg);
    }

    return b;
  };

  me.unexpected = function (tokenIndex, value) {
    var me = this;
    var b = me.isNot.apply(me, arguments);

    if (!b) {
      throw new Error("not expect token " + me.name);
    }

    return b;
  };

  me.name = tokenConst.NAMES[tokenIndex];

  me.toString = function () {
    if (value) {
      return "Token." + me.name + "(" + JSON.stringify(value) + ");";
    } else {
      return "Token." + me.name + "();";
    }
  };
};

module.exports = {
  Token: Token
};
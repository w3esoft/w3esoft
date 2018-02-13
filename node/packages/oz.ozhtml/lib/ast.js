"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.Ast = void 0;

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } }

function _createClass(Constructor, protoProps, staticProps) { if (protoProps) _defineProperties(Constructor.prototype, protoProps); if (staticProps) _defineProperties(Constructor, staticProps); return Constructor; }

var astConst = require("./ast_const");

var Ast =
/*#__PURE__*/
function () {
  function Ast(value, astIndex) {
    _classCallCheck(this, Ast);

    this.astIndex = astIndex;
    this.value = value;
    this.name = astConst.NAMES[astIndex];
  }

  _createClass(Ast, [{
    key: "is",
    value: function is(astIndex, value) {
      var me = this;

      if (Array.isArray(astIndex)) {
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

      return me.astIndex === astIndex;
    }
  }, {
    key: "isNot",
    value: function isNot(astIndex, value) {
      var me = this;
      return !me.is.apply(me, arguments);
    }
  }, {
    key: "expected",
    value: function expected(astIndex, value) {
      var me = this;
      var b = me.is.apply(me, arguments);

      if (!b) {
        throw new Error("unexpected ast " + me.name);
      }

      return b;
    }
  }, {
    key: "unexpected",
    value: function unexpected(astIndex, value) {
      var me = this;
      var b = me.isNot.apply(me, arguments);

      if (!b) {
        throw new Error("unexpected ast " + me.name);
      }

      return b;
    }
  }, {
    key: "toString",
    value: function toString() {
      var me = this;

      if (value) {
        return "Ast." + me.name + "(" + JSON.stringify(value) + ");";
      } else {
        return "Ast." + me.name + "();";
      }
    }
  }, {
    key: "toData",
    value: function toData() {
      var me = this;

      if (me.astIndex === astConst.ATTRKEY) {
        var fields = [];
        var _iteratorNormalCompletion2 = true;
        var _didIteratorError2 = false;
        var _iteratorError2 = undefined;

        try {
          for (var _iterator2 = me.value.fields[Symbol.iterator](), _step2; !(_iteratorNormalCompletion2 = (_step2 = _iterator2.next()).done); _iteratorNormalCompletion2 = true) {
            var _field = _step2.value;
            fields.push(_field.toData());
          }
        } catch (err) {
          _didIteratorError2 = true;
          _iteratorError2 = err;
        } finally {
          try {
            if (!_iteratorNormalCompletion2 && _iterator2.return != null) {
              _iterator2.return();
            }
          } finally {
            if (_didIteratorError2) {
              throw _iteratorError2;
            }
          }
        }

        var _value = me.value.value.toData();

        return {
          type: me.name,
          fields: fields,
          isGlobal: me.value.isGlobal,
          operation: me.value.operation,
          value: _value
        };
      } else if (me.astIndex === astConst.DOCUMENT) {
        var values = [];
        var _iteratorNormalCompletion3 = true;
        var _didIteratorError3 = false;
        var _iteratorError3 = undefined;

        try {
          for (var _iterator3 = me.value[Symbol.iterator](), _step3; !(_iteratorNormalCompletion3 = (_step3 = _iterator3.next()).done); _iteratorNormalCompletion3 = true) {
            var _v = _step3.value;
            values.push(_v.toData());
          }
        } catch (err) {
          _didIteratorError3 = true;
          _iteratorError3 = err;
        } finally {
          try {
            if (!_iteratorNormalCompletion3 && _iterator3.return != null) {
              _iterator3.return();
            }
          } finally {
            if (_didIteratorError3) {
              throw _iteratorError3;
            }
          }
        }

        return {
          type: me.name,
          values: values
        };
      } else if (me.astIndex === astConst.TAG) {
        var attrs = [];
        var _iteratorNormalCompletion4 = true;
        var _didIteratorError4 = false;
        var _iteratorError4 = undefined;

        try {
          for (var _iterator4 = me.value.attrs[Symbol.iterator](), _step4; !(_iteratorNormalCompletion4 = (_step4 = _iterator4.next()).done); _iteratorNormalCompletion4 = true) {
            var _v4 = _step4.value;
            attrs.push(_v4.toData());
          }
        } catch (err) {
          _didIteratorError4 = true;
          _iteratorError4 = err;
        } finally {
          try {
            if (!_iteratorNormalCompletion4 && _iterator4.return != null) {
              _iterator4.return();
            }
          } finally {
            if (_didIteratorError4) {
              throw _iteratorError4;
            }
          }
        }

        var body = [];
        var _iteratorNormalCompletion5 = true;
        var _didIteratorError5 = false;
        var _iteratorError5 = undefined;

        try {
          for (var _iterator5 = me.value.attrs[Symbol.iterator](), _step5; !(_iteratorNormalCompletion5 = (_step5 = _iterator5.next()).done); _iteratorNormalCompletion5 = true) {
            var _v5 = _step5.value;
            body.push(_v5.toData());
          }
        } catch (err) {
          _didIteratorError5 = true;
          _iteratorError5 = err;
        } finally {
          try {
            if (!_iteratorNormalCompletion5 && _iterator5.return != null) {
              _iterator5.return();
            }
          } finally {
            if (_didIteratorError5) {
              throw _iteratorError5;
            }
          }
        }

        return {
          type: me.name,
          name: me.value.name.toData(),
          attrs: attrs,
          body: body
        };
      } else if (me.astIndex === astConst.ATTR) {
        var r = {
          type: me.name,
          attrs: me.value.key.toData()
        };

        if (me.value.value) {
          r.value = me.value.value.toData();
        }

        return r;
      } else if (me.astIndex === astConst.VALUE) {
        return {
          type: me.name,
          value: me.value
        };
      } else if (me.astIndex === astConst.TAGCOMMENT) {
        return {
          type: me.name,
          value: me.value
        };
      }
    }
  }]);

  return Ast;
}();

exports.Ast = Ast;
;
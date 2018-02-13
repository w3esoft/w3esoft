"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.InputString = void 0;

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } }

function _createClass(Constructor, protoProps, staticProps) { if (protoProps) _defineProperties(Constructor.prototype, protoProps); if (staticProps) _defineProperties(Constructor, staticProps); return Constructor; }

var InputString =
/*#__PURE__*/
function () {
  function InputString() {
    _classCallCheck(this, InputString);

    this.pos = 0;
  }

  _createClass(InputString, [{
    key: "read",
    value: function read() {
      var me = this;
      var char = str[me.pos++];
      if (!char) return -1;else return char.charCodeAt(0);
    }
  }]);

  return InputString;
}();

exports.InputString = InputString;
;
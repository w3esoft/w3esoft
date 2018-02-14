const fs = require('fs');

const copyFileSync = function (target, source) {
  mkdirSync(path.dirname(target));
  fs.writeFileSync(target, source);
};

module.exports = {
  copyFileSync
};
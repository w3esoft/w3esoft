const fs = require('fs');

const path = require('path');

const mkdirSync = function (target) {
  if (fs.existsSync(target)) {} else {
    mkdirSync(path.dirname(target));
    fs.mkdirSync(target);
  }
};

module.exports = {
  mkdirSync
};
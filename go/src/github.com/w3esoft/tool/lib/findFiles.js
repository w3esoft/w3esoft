const fs = require('fs');

const path = require('path');

const findFiles = function (o) {
  const createBuild = function (files) {
    const _files = [];

    for (let file of files) {
      const _rules = [];
      let paths = file.split(/[\/\\]/);

      if (!paths[0]) {
        paths.shift();
      }

      let ii;

      for (let p of paths) {
        if (/^(\*\*)\./.test(p)) {
          ii = null;

          _rules.push({
            type: "recursiveExtension",
            value: p.replace(/^(\*\*)\./, "")
          });
        } else if (/^(\*)\./.test(p)) {
          ii = null;

          _rules.push({
            type: "extension",
            value: p.replace(/^(\*)\./, "")
          });
        } else if (/^(\*\*)$/.test(p)) {
          ii = null;

          _rules.push({
            type: "allPath"
          });
        } else if (/^(\*)$/.test(p)) {
          ii = null;

          _rules.push({
            type: "anyPath"
          });
        } else {
          if (ii) {
            ii.value += "/" + p;
          } else {
            ii = {
              type: "path",
              value: p
            };

            _rules.push(ii);
          }
        }
      }

      _files.push(_rules);
    }

    return _files;
  };

  const bFiles = createBuild(o.files);
  const bExcludes = createBuild(o.excludes);
  const showHiddenFolder = !!o.showHiddenFolder;
  const ROOTDIR = o.rootDir;

  const ignore = function (_rootDir, bExcludes) {
    let basename = path.basename(_rootDir);

    if (!showHiddenFolder && basename && basename[0] === ".") {
      return true;
    }

    const _ignore = function (r, _rootDir, bExcludes) {
      let pr = path.relative(r, _rootDir);

      for (let rules of bExcludes) {
        let r_rules = [].concat(rules);
        let rule = r_rules.shift();

        if (rule) {
          const rule = rules[0];

          if (rule.type === "path") {
            if (pr.indexOf(rule.value) === 0) {
              if (r_rules.length == 0) {
                return true;
              } else {
                let v2 = path.resolve(r, rule.value);

                let v4 = _ignore(v2, _rootDir, [r_rules]);

                return v4;
              }
            } else {
              return false;
            }
          } else if (rule.type === "extension") {} else if (rule.type === "recursiveExtension") {} else if (rule.type === "allPath") {
            if (r_rules.length === 0) return true;
            let f1 = pr.split(/[\/\\]/),
                f2 = "",
                f3 = "",
                f4 = [];

            for (let o of f1) {
              f2 += f3 + o;
              f4.push(f2);
              f3 = "/";
            }

            let b = false;

            for (let o of f4) {
              let i1 = path.resolve(r, o);
              b = _ignore(i1, _rootDir, [r_rules]);

              if (b) {
                break;
              }
            }

            return b;
          } else if (rule.type === "anyPath") {}
        }
      }
    };

    return _ignore(ROOTDIR, _rootDir, bExcludes);
  };

  const find = function (rootDir, bFiles, bExcludes) {
    let r_files = [];

    for (let rules of bFiles) {
      let r_rules = [].concat(rules);
      let r_rootDirs = [];
      let rule = r_rules.shift();

      if (rule) {
        const rule = rules[0];

        if (rule.type === "path") {
          if (r_rules.length === 0) {
            let v = findPath(rootDir, rule.value, bExcludes);

            if (v) {
              r_files.push(v);
            }
          } else {
            let v = findPath(rootDir, rule.value, bExcludes);

            if (v) {
              r_rootDirs.push(v);
            }
          }
        } else if (rule.type === "extension") {
          let r = findExtension(rootDir, rule.value, bExcludes);
          r_files = r_files.concat(r);
        } else if (rule.type === "recursiveExtension") {
          let r = findRecursiveExtension(rootDir, rule.value, bExcludes);
          r_files = r_files.concat(r);
        } else if (rule.type === "allPath") {
          if (r_rules.length === 0) {
            let r = findAllPath(rootDir, bExcludes);
            r_files = r_files.concat(r);
          } else {
            let r = findAllPath(rootDir, bExcludes);
            r_rootDirs = r_rootDirs.concat(r);
          }
        } else if (rule.type === "anyPath") {}
      }

      for (let r_rootDir of r_rootDirs) {
        if (r_rules.length) {
          r_files = r_files.concat(find(r_rootDir, [r_rules], bExcludes));
        }
      }
    }

    return r_files.filter(function (item, pos) {
      return r_files.indexOf(item) === pos;
    });
  };

  const findAllPath = function (_rootDir, bExcludes) {
    let r_files = [];

    if (!ignore(_rootDir, bExcludes)) {
      if (fs.existsSync(_rootDir)) {
        r_files.push(_rootDir);
        const stats = fs.lstatSync(_rootDir);

        if (stats.isDirectory()) {
          const r_rootDirs = fs.readdirSync(_rootDir);

          for (let r_rootDir of r_rootDirs) {
            r_files = r_files.concat(findAllPath(path.relative(ROOTDIR, path.resolve(_rootDir, r_rootDir)), bExcludes));
          }
        }
      }
    }

    return r_files;
  };

  const findExtension = function (_rootDir, extension, bExcludes) {
    let r_files = [];

    if (!ignore(_rootDir, bExcludes)) {
      if (fs.existsSync(_rootDir)) {
        const stats = fs.lstatSync(_rootDir);

        if (stats.isDirectory()) {
          const r_rootDirs = fs.readdirSync(_rootDir);

          for (let r_rootDir of r_rootDirs) {
            const stats = fs.lstatSync(r_rootDir);

            if (stats.isFile() && new RegExp(extension + "$", "i").test(_rootDir)) {
              r_files.push(_rootDir);
            }
          }
        }
      }
    }

    return r_files;
  };

  const findRecursiveExtension = function (_rootDir, extension, bExcludes) {
    let r_files = [];

    if (!ignore(_rootDir, bExcludes)) {
      if (fs.existsSync(_rootDir)) {
        const stats = fs.lstatSync(_rootDir);

        if (stats.isDirectory()) {
          const r_rootDirs = fs.readdirSync(_rootDir);

          for (let r_rootDir of r_rootDirs) {
            r_files = r_files.concat(findRecursiveExtension(path.resolve(_rootDir, r_rootDir), extension, bExcludes));
          }
        } else {
          if (new RegExp(extension + "$", "i").test(_rootDir)) {
            r_files.push(_rootDir);
          }
        }
      }
    }

    return r_files;
  };

  const findPath = function (rootDir, _path, bExcludes) {
    let _rootDir = path.resolve(ROOTDIR, _path);

    if (fs.existsSync(_rootDir) && !ignore(_rootDir, bExcludes)) {
      return _rootDir;
    }
  };

  let _files = find(ROOTDIR, bFiles, bExcludes);

  return _files;
};

module.exports = {
  findFiles
};
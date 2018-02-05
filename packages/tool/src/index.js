/* 
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

const fs = require('fs');
const path = require('path');
const child_process = require('child_process');
const cp = require('child_process');
const deleteFolderRecursive = function (path) {
    if (fs.existsSync(path)) {
        fs.readdirSync(path).forEach(function (file, index) {
            var curPath = path + "/" + file;
            if (fs.lstatSync(curPath).isDirectory()) { // recurse
                deleteFolderRecursive(curPath);
            } else { // delete file
                fs.unlinkSync(curPath);
            }
        });
        fs.rmdirSync(path);
    }
};

const mkdirSync = function (target) {
    if (fs.existsSync(target)) {
    } else {
        mkdirSync(path.dirname(target));
        fs.mkdirSync(target);
    }
};
const copyFileSync = function (target, source) {
    mkdirSync(path.dirname(target));
    fs.writeFileSync(target, source);
};



const findFiles = function (o) {
    const createBuild = function (files) {
        const _files = [];
        for (let file of  files) {
            const _rules = [];
            let paths = file.split("/");
            if (!paths[0]) {
                paths.shift();
            }
            let ii;
            for (let p of  paths) {
                if (/^(\*\*)\./.test(p)) {
                    ii =null;
                    _rules.push({
                        type: "recursiveExtension",
                        value: p.replace(/^(\*\*)\./, "")
                    });
                } else if (/^(\*)\./.test(p)) {
                    ii =null;
                    _rules.push({
                        type: "extension",
                        value: p.replace(/^(\*)\./, "")
                    });
                } else if (/^(\*\*)$/.test(p)) {
                    ii =null;
                    _rules.push({
                        type: "allPath"
                    });
                } else if (/^(\*)$/.test(p)) {
                    ii =null;
                    _rules.push({
                        type: "anyPath"
                    });
                } else {
                    if (ii){
                        ii.value +="/"+ p;
                    }else {
                        ii ={
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
    const _ignore = function (_rootDir) {
        let basename = path.basename(_rootDir);
        if (!showHiddenFolder && basename && basename[0] === ".") {
            return true;
        }
        return false;
    };
    const findAllPath = function (_rootDir) {
        let r_files = [];
        if (!_ignore(_rootDir)) {
            if (fs.existsSync(_rootDir)) {
                r_files.push(_rootDir);
                const stats = fs.lstatSync(_rootDir);
                if (stats.isDirectory()) {
                    const r_rootDirs = fs.readdirSync(_rootDir);
                    for (let r_rootDir of r_rootDirs) {
                        r_files = r_files.concat(findAllPath(path.relative(ROOTDIR,path.resolve(_rootDir, r_rootDir))));
                    }
                }
            }
        }
        return r_files;
    };
    const findExtension = function (_rootDir, extension) {
        let r_files = [];
        if (!_ignore(_rootDir)) {            if (fs.existsSync(_rootDir)) {
                const stats = fs.lstatSync(_rootDir);
                if (stats.isDirectory()) {
                    const r_rootDirs = fs.readdirSync(_rootDir);
                    for (let r_rootDir of r_rootDirs) {
                        const stats = fs.lstatSync(r_rootDir);
                        if (stats.isFile()&&new RegExp(extension + "$", "i").test(_rootDir)) {
                            r_files.push(_rootDir);
                        }
                        
                    }
                }
            }
        }
        return r_files;
    };
    
    const findRecursiveExtension = function (_rootDir, extension) {
        let r_files = [];
        if (!_ignore(_rootDir)) {
            if (fs.existsSync(_rootDir)) {
                const stats = fs.lstatSync(_rootDir);
                if (stats.isDirectory()) {
                    const r_rootDirs = fs.readdirSync(_rootDir);
                    for (let r_rootDir of r_rootDirs) {
                        r_files = r_files.concat(findRecursiveExtension(path.resolve(_rootDir, r_rootDir), extension));
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
    const findPath = function (rootDir,_path) {
        let _rootDir= path.resolve(ROOTDIR,_path) ;
        if (fs.existsSync(_rootDir) ) {
            return _rootDir;
        }
    };
    const find = function (rootDir, bFiles) {
        let r_files = [];
        for (let rules of bFiles) {
            let r_rules = [].concat(rules);
            let r_rootDirs = [];
            let rule =r_rules.shift();
            if (rule) {
                const rule = rules[0];
                if (rule.type === "path") {
                    if (r_rules.length === 0) {
                        let v = findPath(rootDir,rule.value);
                        if (v){
                            r_files.push(v);
                        }
                    }else {
                        let v = findPath(rootDir,rule.value);
                        
                        if (v){
                            r_rootDirs.push(v);
                        }
                    }
                } else if (rule.type === "extension") {
                    let r = findExtension(rootDir, rule.value);
                    r_files = r_files.concat(r);
                } else if (rule.type === "recursiveExtension") {
                    let r = findRecursiveExtension(rootDir, rule.value);
                    r_files = r_files.concat(r);
                } else if (rule.type === "allPath") {
                    if (r_rules.length === 0) {          
                        let r = findAllPath(rootDir);
                        r_files = r_files.concat(r);
                    }else {      
                        let r = findAllPath(rootDir);
                        r_rootDirs = r_rootDirs.concat(r); 
                    }
                } else if (rule.type === "anyPath") {
                }
            }
            for (let r_rootDir of   r_rootDirs) {
                if (r_rules.length) {
                    r_files = r_files.concat(find(r_rootDir, [r_rules]));
                }
            }
        }
        return  r_files.filter(function (item, pos) {
            return r_files.indexOf(item) === pos;
        });
    };
    let excludes = find(ROOTDIR, bExcludes);
    let _files = find(ROOTDIR, bFiles);
    let files =[];
    for (let f of _files){
        let b= true;
        for (let e of excludes){
            if (f.indexOf(e)===0){
                b=false;
                break;
            }
        }
        if (b){ 
            files.push(path.relative(ROOTDIR,f));
        }
    }
    return files;
};
const html = require("./html");
module.exports = {
    mkdirSync: mkdirSync,
    copyFileSync:copyFileSync,
    findFiles:findFiles,
    HtmlParser:html.HtmlParser
};
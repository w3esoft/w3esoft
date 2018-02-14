/* 
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

const {copyFileSync} = require("./copyFileSync");
const {deleteFolderRecursive} = require("./deleteFolderRecursive");
const {findFiles} = require("./findFiles");
const {mkdirSync} = require("./mkdirSync");

module.exports = {
    mkdirSync: mkdirSync,
    copyFileSync:copyFileSync,
    findFiles:findFiles,
    deleteFolderRecursive:deleteFolderRecursive
};
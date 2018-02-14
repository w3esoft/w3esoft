const path = require('path');
const fs = require('fs');
const child_process = require('child_process');

module.exports = function(){

	const map = {
			"@w3esoft/craft"		: "./packages/craft"		,
			"@w3esoft/glaive"		: "./packages/glaive"		,
			"@w3esoft/oz"			: "./packages/oz"			,
			"@w3esoft/oz.ozcss"		: "./packages/oz.ozcss"		,
			"@w3esoft/oz.ozhtml"	: "./packages/oz.ozhtml"	,
			"@w3esoft/tool"			: "./packages/tool"			
	};

	let   files = [];
	for (let key  in map) {
		let p = map[key];
		let babelJs=path.resolve("./node_modules/@babel/cli/bin/babel.js");
		if(fs.existsSync(p+"/src")&&fs.existsSync(p+"/.babelrc")){
			files.push(p+"/src/**.js");
			child_process.fork(babelJs,[p+"/src/**.js","--config-file", p +"/.babelrc" ,"-d",p+"/lib","--source-maps"]);
		}
		
	}
	const {findFiles,mkdirSync} = require('@w3esoft/tool');
	let rootDir = path.resolve("../");
	console.log(rootDir);
	console.log(rootDir);
	files =findFiles(
	        {
	            rootDir:rootDir,
	            files: files,
	            excludes: []
	        }
	);
	console.log(files);
	//const {InputString} = require('@w3esoft/glaive');
	//const {Parser,Lexer} = require('@w3esoft/oz.ozhtml');
}
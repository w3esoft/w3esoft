let p = "";

const path = require('path');
const fs = require('fs');
const {findFiles,mkdirSync} = require('@w3esoft/tool');
const {InputString} = require('@w3esoft/glaive');
const {Parser,Lexer} = require('@w3esoft/oz.ozhtml');


const map = {
    "@atomicburst/breadcrumb": p + "atomicburst/packages/breadcrumb",
    "@atomicburst/button": p + "atomicburst/packages/button",
    "@atomicburst/calendar": p + "atomicburst/packages/calendar",
    "@atomicburst/chart": p + "atomicburst/packages/chart",
    "@atomicburst/container": p + "atomicburst/packages/container",
    "@atomicburst/container.layout": p + "atomicburst/packages/container.layout",
    "@atomicburst/container.modal": p + "atomicburst/packages/container.modal",
    "@atomicburst/container.panel": p + "atomicburst/packages/container.panel",
    "@atomicburst/container.viewport": p + "atomicburst/packages/container.viewport",
    "@atomicburst/core": p + "atomicburst/packages/core",
    "@atomicburst/data": p + "atomicburst/packages/data",
    "@atomicburst/datatable": p + "atomicburst/packages/datatable",
    "@atomicburst/form.field": p + "atomicburst/packages/form.field",
    "@atomicburst/form.field.checkbox": p + "atomicburst/packages/form.field.checkbox",
    "@atomicburst/form.field.combo": p + "atomicburst/packages/form.field.combo",
    "@atomicburst/form.field.date": p + "atomicburst/packages/form.field.date",
    "@atomicburst/form.field.file": p + "atomicburst/packages/form.field.file",
    "@atomicburst/form.field.picker.date": p + "atomicburst/packages/form.field.picker.date",
    "@atomicburst/form.field.range": p + "atomicburst/packages/form.field.range",
    "@atomicburst/form.field.text": p + "atomicburst/packages/form.field.text",
    "@atomicburst/form.field.textarea": p + "atomicburst/packages/form.field.textarea",
    "@atomicburst/form.mask": p + "atomicburst/packages/form.mask",
    "@atomicburst/form.validation": p + "atomicburst/packages/form.validation",
    "@atomicburst/imagefield": p + "atomicburst/packages/imagefield",
    "@atomicburst/locale.pt_br": p + "atomicburst/packages/locale.pt_br",
    "@atomicburst/menu": p + "atomicburst/packages/menu",
    "@atomicburst/pagination": p + "atomicburst/packages/pagination",
    "@atomicburst/toolbar": p + "atomicburst/packages/toolbar",
    "@damsistemas/core": p + "damsistemas/packages/core",
    "@damsistemas/crm": p + "damsistemas/packages/crm"
};

let   files = [];
for (let key  in map) {
    files.push(map[key] + "/**.html");
}


let rootDir = path.resolve("F://desenvolvimento/damonline2/angular/packages/");
files =findFiles(
        {
            rootDir:rootDir,
            files: files,
            excludes: ["**/node_modules"]
        }
);
let ii=0;
files.shift();
files.shift();
files.shift();
files.shift();
for (let file of files){
    let d1  =path.resolve(rootDir,file);
    let d2  =path.resolve("test",file);
    mkdirSync(path.dirname(d2));
    let  v=fs.readFileSync(d1);
    let c = v.toString();
    const input =new InputString(c);
    const lexer = new  Lexer(input);
    const parser =  new Parser(lexer);


    console.log("\n\n\n\n");
    console.log(d1)
    console.log((ii++))
    console.log(c);
    console.log("\n\n\n\n");


    let ee=parser.parse().toData();
    console.log(JSON.stringify(ee, null, "\t"))
    fs.writeFileSync(d2,c);
}

const astConst = require("./ast_const");
const HtmlAst = function (astIndex,value){
    let me = this;
    me.astIndex = astIndex;
    me.value = value;
    me.is=function(astIndex,value){
        if (Array.isArray(astIndex)){
            let b ;
            for ( let iii of tokenIndex){
                b =me.is(iii,value);
                if (b){
                    b = true;
                    break;
                }
            }
            return !!b;
        }

        return (me.astIndex===astIndex);
    };
    me.isNot=function(astIndex,value){
        let me = this;
        return !me.is.apply(me,arguments);
    };
    me.expected=function (astIndex,value) {
        let me = this;
        let b = me.is.apply(me,arguments) ;
        if (!b){
            throw new Error("unexpected ast " + me.name);
        }
        return b;
    };
    me.unexpected=function (astIndex,value) {
        let me = this;
        let b = me.isNot.apply(me,arguments);
        if (!b){
            throw new Error("unexpected ast " + me.name);
        }
        return b;
    };
    me.name= astConst.NAMES[astIndex];
    me.toString =function () {
        if (value){
            return "HtmlAst."+me.name+"("+ JSON.stringify(value)+");";
        }else {
            return "HtmlAst."+me.name+"();";
        }
    };
    me.toData =function () {
        if (me.astIndex===astConst.ATTRKEY){
            let fields= [];
            for ( let field of me.value.fields){
                fields.push(field.toData());
            }
            let value = me.value.value.toData();
            return {
                type:me.name,
                fields:fields,
                isGlobal:me.value.isGlobal,
                operation :me.value.operation,
                value:value
            }
        }else if ( me.astIndex===astConst.DOCUMENT ) {
            let values =[];
            for (let v of me.value){
                values.push(v.toData());
            }
            return {
                type:me.name,
                values:values
            }
        }else  if ( me.astIndex===astConst.TAG )  {
            let attrs =[];
            for (let v of me.value.attrs){
                attrs.push(v.toData());
            }
            let body =[];
            for (let v of me.value.attrs){
                body.push(v.toData());
            }
            return {
                type:me.name,
                name:me.value.name.toData(),
                attrs:attrs,
                body:body
            }

        }else if(me.astIndex===astConst.ATTR) {
            let r= {
                type:me.name,
                attrs:me.value.key.toData()
            };
            if (me.value.value){
                r.value = me.value.value.toData();
            }

            return r;
        }else if(me.astIndex===astConst.VALUE) {
            return  {
                type:me.name,
                value:me.value
            };
        } else if(me.astIndex===astConst.TAGCOMMENT)  {
            return  {
                type:me.name,
                value:me.value
            };
        }else {
            console.log(me)
        }


    }
};
module.exports = {HtmlAst};
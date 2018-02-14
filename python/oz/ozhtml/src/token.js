const tokenConst = require("./token_const");
let TOKEN_ID= 1;
const Token = function (tokenIndex,value){
    let me = this;
    me.id=100+(TOKEN_ID++);
    me.tokenIndex = tokenIndex;
    me.value = value;
    me.is=function(tokenIndex,value){
        if (Array.isArray(tokenIndex)){
            let b;
            for ( let iii of tokenIndex){
                b =me.is(iii,value);
                if (b){
                    b = true;
                    break;
                }
            }
            return !!b;
        }
        let b1 = (me.tokenIndex===tokenIndex);

        let b2 = (me.value===value||value===undefined);

        return b1 && b2;
    };
    me.isNot=function(tokenIndex,value){
        let me = this;
        return !me.is.apply(me,arguments);
    };
    me.expected=function (tokenIndex,value) {
        let me = this;
        let b = me.is.apply(me,arguments) ;
        if (!b){
            let msg = "not expect token " + me.name + "  ";
            if (me.value){
                msg+=" value = " + me.value;
            }
            throw new Error(msg);
        }
        return b;
    };
    me.unexpected=function (tokenIndex,value) {
        let me = this;
        let b = me.isNot.apply(me,arguments);
        if (!b){
            throw new Error("not expect token " +me.name);
        }
        return b;
    };
    me.name= tokenConst.NAMES[tokenIndex];
    me.toString =function () {
        if (value){
            return "Token."+me.name+"("+ JSON.stringify(value)+");"
        }else {
            return "Token."+me.name+"();"
        }
    };

};



module.exports = {
    Token
};


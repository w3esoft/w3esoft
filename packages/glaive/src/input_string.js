
const HtmlInputString = function (str){
    let me = this;
    me.pos=0;
    me.read=function(){
        let char = str[me.pos++];
        if (!char)
            return -1;
        else
            return char.charCodeAt(0);

    }
};

module.exports = {
    HtmlInputString
};


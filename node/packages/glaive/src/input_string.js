export  class InputString {
    pos=0;
    read=function(){
        let me = this;
        let char = str[me.pos++];
        if (!char)
            return -1;
        else
            return char.charCodeAt(0);
    }
};
export  class InputString { 
   
    constructor()
	{
    	this.pos=0;
	}
    read(){
        let me = this;
        let char = str[me.pos++];
        if (!char)
            return -1;
        else
            return char.charCodeAt(0);
    }
};
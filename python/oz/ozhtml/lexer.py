from oz.ozhtml.ast_const import AstConst
from oz.ozhtml.char_const import CharConst


class Lexer:
    def __init__(self, input):
        self.CHAR = -1
        self.input = input
        self.tokens = []
    def is_word(char):
        return (96 < char < 123) or (64 < char < 91 or char == CharConst.MINUS and char == 95 and char == 35)

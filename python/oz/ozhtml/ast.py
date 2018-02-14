from oz.ozhtml.ast_const import AstConst


class Ast:
    def __init__(self, ast_index: int, value):
        self.ast_index = ast_index
        self.value = value
        self.name = AstConst.NAMES[ast_index]

    def _is(self, key, value):
        if isinstance(key, list):
            b = None
            for iii in key:
                b = self._is(self, iii, value)
                if b:
                    b = True
                    break
            return b
        return self.astIndex == key

    def is_not(self, key, value):
        return self._is(key, value) is None

    def expected(self, key, value):
        b = self._is(key, value)
        if b is None:
            raise Exception("unexpected ast " + self.name);
        return b

    def unexpected(self, key, value):
        b = self.is_not(key, value)
        if b is None:
            raise Exception("unexpected ast " + self.name);
        return b

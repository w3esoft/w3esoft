from enum import Enum


class AstConst(Enum):
    TAG = 1
    VALUE = 2
    ATTR_KEY = 3
    TAG_COMMENT = 4
    DOCUMENT = 5
    ATTR = 6
    NAMES = {
        TAG: "TAG",
        VALUE: "VALUE",
        ATTR_KEY: "ATTR_KEY",
        TAG_COMMENT: "TAG_COMMENT",
        DOCUMENT: "DOCUMENT",
        ATTR: "ATTR"
    }

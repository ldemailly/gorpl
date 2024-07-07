// Code generated by "stringer -type=Type"; DO NOT EDIT.

package token

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ILLEGAL-0]
	_ = x[EOF-1]
	_ = x[IDENT-2]
	_ = x[INT-3]
	_ = x[ASSIGN-4]
	_ = x[PLUS-5]
	_ = x[MINUS-6]
	_ = x[BANG-7]
	_ = x[ASTERISK-8]
	_ = x[SLASH-9]
	_ = x[PERCENT-10]
	_ = x[LT-11]
	_ = x[GT-12]
	_ = x[EQ-13]
	_ = x[NOTEQ-14]
	_ = x[COMMA-15]
	_ = x[SEMICOLON-16]
	_ = x[LPAREN-17]
	_ = x[RPAREN-18]
	_ = x[LBRACE-19]
	_ = x[RBRACE-20]
	_ = x[FUNCTION-21]
	_ = x[LET-22]
	_ = x[TRUE-23]
	_ = x[FALSE-24]
	_ = x[IF-25]
	_ = x[ELSE-26]
	_ = x[RETURN-27]
	_ = x[STRING-28]
	_ = x[LEN-29]
	_ = x[FIRST-30]
	_ = x[REST-31]
}

const _Type_name = "ILLEGALEOFIDENTINTASSIGNPLUSMINUSBANGASTERISKSLASHPERCENTLTGTEQNOTEQCOMMASEMICOLONLPARENRPARENLBRACERBRACEFUNCTIONLETTRUEFALSEIFELSERETURNSTRINGLENFIRSTREST"

var _Type_index = [...]uint8{0, 7, 10, 15, 18, 24, 28, 33, 37, 45, 50, 57, 59, 61, 63, 68, 73, 82, 88, 94, 100, 106, 114, 117, 121, 126, 128, 132, 138, 144, 147, 152, 156}

func (i Type) String() string {
	if i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}

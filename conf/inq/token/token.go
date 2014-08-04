
package token

import(
	"fmt"
)

type Token struct {
	Type
	Lit []byte
	Pos
}

type Type int

const(
	INVALID Type = iota
	EOF
)

type Pos struct {
	Offset int
	Line int
	Column int
}

func (this Pos) String() string {
	return fmt.Sprintf("Pos(offset=%d, line=%d, column=%d)", this.Offset, this.Line, this.Column)
}

type TokenMap struct {
	typeMap  []string
	idMap map[string]Type
}

func (this TokenMap) Id(tok Type) string {
	if int(tok) < len(this.typeMap) {
		return this.typeMap[tok]
	}
	return "unknown"
}

func (this TokenMap) Type(tok string) Type {
	if typ, exist := this.idMap[tok]; exist {
		return typ
	}
	return INVALID
}

func (this TokenMap) TokenString(tok *Token) string {
	//TODO: refactor to print pos & token string properly
	return fmt.Sprintf("%s(%d,%s)", this.Id(tok.Type), tok.Type, tok.Lit)
}

func (this TokenMap) StringType(typ Type) string {
	return fmt.Sprintf("%s(%d)", this.Id(typ), typ)
}

var TokMap = TokenMap{
	typeMap: []string{
		"INVALID",
		"$",
		"name",
		"host",
		"check",
		"process",
		"service",
		"restart",
		"alert",
		"uint_lit",
		"sized_uint_lit",
		"(",
		")",
		"if",
		"operator",
		"then",
		"for",
		"cycles",
	},

	idMap: map[string]Type {
		"INVALID": 0,
		"$": 1,
		"name": 2,
		"host": 3,
		"check": 4,
		"process": 5,
		"service": 6,
		"restart": 7,
		"alert": 8,
		"uint_lit": 9,
		"sized_uint_lit": 10,
		"(": 11,
		")": 12,
		"if": 13,
		"operator": 14,
		"then": 15,
		"for": 16,
		"cycles": 17,
	},
}


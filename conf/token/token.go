
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
		"check",
		"id",
		"host",
		"process",
		"service",
		"restart",
		"alert",
		"launchctl",
		"upstart",
		"runit",
		"systemd",
		"init.d",
		"uint_lit",
		"sizecode",
		"if",
		"operator",
		"then",
		"for",
		"cycles",
	},

	idMap: map[string]Type {
		"INVALID": 0,
		"$": 1,
		"check": 2,
		"id": 3,
		"host": 4,
		"process": 5,
		"service": 6,
		"restart": 7,
		"alert": 8,
		"launchctl": 9,
		"upstart": 10,
		"runit": 11,
		"systemd": 12,
		"init.d": 13,
		"uint_lit": 14,
		"sizecode": 15,
		"if": 16,
		"operator": 17,
		"then": 18,
		"for": 19,
		"cycles": 20,
	},
}


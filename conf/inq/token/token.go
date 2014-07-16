
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
		"id",
		"host",
		"check",
		"inspect",
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
		"id": 2,
		"host": 3,
		"check": 4,
		"inspect": 5,
		"process": 6,
		"service": 7,
		"restart": 8,
		"alert": 9,
		"launchctl": 10,
		"upstart": 11,
		"runit": 12,
		"systemd": 13,
		"init.d": 14,
		"uint_lit": 15,
		"sizecode": 16,
		"if": 17,
		"operator": 18,
		"then": 19,
		"for": 20,
		"cycles": 21,
	},
}


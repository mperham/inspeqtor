
package lexer

import(
	"fmt"
	"inspeqtor/conf/token"
)

type ActionTable [NumStates] ActionRow

type ActionRow struct {
	Accept token.Type
	Ignore string
}

func (this ActionRow) String() string {
	return fmt.Sprintf("Accept=%d, Ignore=%s", this.Accept, this.Ignore)
}

var ActTab = ActionTable{
 	ActionRow{ // S0
		Accept: 0,
 		Ignore: "",
 	},
 	ActionRow{ // S1
		Accept: -1,
 		Ignore: "!whitespace",
 	},
 	ActionRow{ // S2
		Accept: 0,
 		Ignore: "",
 	},
 	ActionRow{ // S3
		Accept: 14,
 		Ignore: "",
 	},
 	ActionRow{ // S4
		Accept: 17,
 		Ignore: "",
 	},
 	ActionRow{ // S5
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S6
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S7
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S8
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S9
		Accept: 15,
 		Ignore: "",
 	},
 	ActionRow{ // S10
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S11
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S12
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S13
		Accept: 15,
 		Ignore: "",
 	},
 	ActionRow{ // S14
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S15
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S16
		Accept: 15,
 		Ignore: "",
 	},
 	ActionRow{ // S17
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S18
		Accept: -1,
 		Ignore: "!comment",
 	},
 	ActionRow{ // S19
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S20
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S21
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S22
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S23
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S24
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S25
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S26
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S27
		Accept: 16,
 		Ignore: "",
 	},
 	ActionRow{ // S28
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S29
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S30
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S31
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S32
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S33
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S34
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S35
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S36
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S37
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S38
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S39
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S40
		Accept: 19,
 		Ignore: "",
 	},
 	ActionRow{ // S41
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S42
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S43
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S44
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S45
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S46
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S47
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S48
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S49
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S50
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S51
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S52
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S53
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S54
		Accept: 4,
 		Ignore: "",
 	},
 	ActionRow{ // S55
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S56
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S57
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S58
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S59
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S60
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S61
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S62
		Accept: 18,
 		Ignore: "",
 	},
 	ActionRow{ // S63
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S64
		Accept: 8,
 		Ignore: "",
 	},
 	ActionRow{ // S65
		Accept: 2,
 		Ignore: "",
 	},
 	ActionRow{ // S66
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S67
		Accept: 0,
 		Ignore: "",
 	},
 	ActionRow{ // S68
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S69
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S70
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S71
		Accept: 11,
 		Ignore: "",
 	},
 	ActionRow{ // S72
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S73
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S74
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S75
		Accept: 20,
 		Ignore: "",
 	},
 	ActionRow{ // S76
		Accept: 13,
 		Ignore: "",
 	},
 	ActionRow{ // S77
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S78
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S79
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S80
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S81
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S82
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S83
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S84
		Accept: 5,
 		Ignore: "",
 	},
 	ActionRow{ // S85
		Accept: 7,
 		Ignore: "",
 	},
 	ActionRow{ // S86
		Accept: 6,
 		Ignore: "",
 	},
 	ActionRow{ // S87
		Accept: 12,
 		Ignore: "",
 	},
 	ActionRow{ // S88
		Accept: 10,
 		Ignore: "",
 	},
 	ActionRow{ // S89
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S90
		Accept: 9,
 		Ignore: "",
 	},
 		
}

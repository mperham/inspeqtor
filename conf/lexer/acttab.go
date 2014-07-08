
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
		Accept: 6,
 		Ignore: "",
 	},
 	ActionRow{ // S4
		Accept: 5,
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
		Accept: -1,
 		Ignore: "!comment",
 	},
 	ActionRow{ // S10
		Accept: 6,
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
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S14
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S15
		Accept: 4,
 		Ignore: "",
 	},
 	ActionRow{ // S16
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S17
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S18
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S19
		Accept: 3,
 		Ignore: "",
 	},
 	ActionRow{ // S20
		Accept: 7,
 		Ignore: "",
 	},
 	ActionRow{ // S21
		Accept: 2,
 		Ignore: "",
 	},
 		
}

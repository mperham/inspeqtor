
package parser

type(
	actionTable [numStates]actionRow
	actionRow struct {
		canRecover bool
		actions [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* id */
			nil,		/* host */
			shift(3),		/* check */
			shift(4),		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S1
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			accept(true),		/* $ */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S2
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* id */
			shift(6),		/* host */
			nil,		/* check */
			nil,		/* inspect */
			shift(7),		/* process */
			shift(8),		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S3
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* id */
			reduce(3),		/* host, reduce: Preamble */
			nil,		/* check */
			nil,		/* inspect */
			reduce(3),		/* process, reduce: Preamble */
			reduce(3),		/* service, reduce: Preamble */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S4
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* id */
			reduce(4),		/* host, reduce: Preamble */
			nil,		/* check */
			nil,		/* inspect */
			reduce(4),		/* process, reduce: Preamble */
			reduce(4),		/* service, reduce: Preamble */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S5
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(9),		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S6
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			shift(12),		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S7
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(5),		/* id, reduce: Checktype */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S8
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(6),		/* id, reduce: Checktype */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S9
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			shift(12),		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S10
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(2),		/* $, reduce: Check */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			shift(12),		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S11
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(14),		/* $, reduce: RuleList */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			reduce(14),		/* if, reduce: RuleList */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S12
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(15),		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S13
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: Check */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			shift(12),		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S14
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(15),		/* $, reduce: RuleList */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			reduce(15),		/* if, reduce: RuleList */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S15
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			nil,		/* if */
			shift(16),		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S16
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			shift(17),		/* uint_lit */
			nil,		/* sizecode */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S17
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* uint_lit */
			shift(19),		/* sizecode */
			nil,		/* if */
			nil,		/* operator */
			reduce(10),		/* then, reduce: HumanAmount */
			reduce(10),		/* for, reduce: HumanAmount */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S18
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			nil,		/* if */
			nil,		/* operator */
			shift(20),		/* then */
			shift(21),		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S19
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			nil,		/* if */
			nil,		/* operator */
			reduce(11),		/* then, reduce: HumanAmount */
			reduce(11),		/* for, reduce: HumanAmount */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S20
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			shift(23),		/* restart */
			shift(24),		/* alert */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S21
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			shift(26),		/* uint_lit */
			nil,		/* sizecode */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S22
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(12),		/* $, reduce: Rule */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			reduce(12),		/* if, reduce: Rule */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S23
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(7),		/* $, reduce: Action */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			reduce(7),		/* if, reduce: Action */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S24
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(8),		/* $, reduce: Action */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			reduce(8),		/* if, reduce: Action */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S25
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			shift(27),		/* cycles */
			
		},

	},
	actionRow{ // S26
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			reduce(9),		/* cycles, reduce: IntAmount */
			
		},

	},
	actionRow{ // S27
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			nil,		/* if */
			nil,		/* operator */
			shift(28),		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S28
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			shift(23),		/* restart */
			shift(24),		/* alert */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S29
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(13),		/* $, reduce: Rule */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			reduce(13),		/* if, reduce: Rule */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	
}


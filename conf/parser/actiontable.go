
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
			shift(2),		/* check */
			nil,		/* id */
			nil,		/* if */
			nil,		/* > */
			nil,		/* amount */
			nil,		/* then */
			
		},

	},
	actionRow{ // S1
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			accept(true),		/* $ */
			nil,		/* check */
			nil,		/* id */
			nil,		/* if */
			nil,		/* > */
			nil,		/* amount */
			nil,		/* then */
			
		},

	},
	actionRow{ // S2
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			shift(3),		/* id */
			nil,		/* if */
			nil,		/* > */
			nil,		/* amount */
			nil,		/* then */
			
		},

	},
	actionRow{ // S3
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			shift(4),		/* id */
			nil,		/* if */
			nil,		/* > */
			nil,		/* amount */
			nil,		/* then */
			
		},

	},
	actionRow{ // S4
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			shift(5),		/* id */
			shift(8),		/* if */
			nil,		/* > */
			nil,		/* amount */
			nil,		/* then */
			
		},

	},
	actionRow{ // S5
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* id */
			shift(8),		/* if */
			nil,		/* > */
			nil,		/* amount */
			nil,		/* then */
			
		},

	},
	actionRow{ // S6
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(2),		/* $, reduce: Check */
			nil,		/* check */
			nil,		/* id */
			shift(8),		/* if */
			nil,		/* > */
			nil,		/* amount */
			nil,		/* then */
			
		},

	},
	actionRow{ // S7
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(4),		/* $, reduce: RuleList */
			nil,		/* check */
			nil,		/* id */
			reduce(4),		/* if, reduce: RuleList */
			nil,		/* > */
			nil,		/* amount */
			nil,		/* then */
			
		},

	},
	actionRow{ // S8
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			shift(11),		/* id */
			nil,		/* if */
			nil,		/* > */
			nil,		/* amount */
			nil,		/* then */
			
		},

	},
	actionRow{ // S9
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: Check */
			nil,		/* check */
			nil,		/* id */
			shift(8),		/* if */
			nil,		/* > */
			nil,		/* amount */
			nil,		/* then */
			
		},

	},
	actionRow{ // S10
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(5),		/* $, reduce: RuleList */
			nil,		/* check */
			nil,		/* id */
			reduce(5),		/* if, reduce: RuleList */
			nil,		/* > */
			nil,		/* amount */
			nil,		/* then */
			
		},

	},
	actionRow{ // S11
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* id */
			nil,		/* if */
			shift(12),		/* > */
			nil,		/* amount */
			nil,		/* then */
			
		},

	},
	actionRow{ // S12
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* id */
			nil,		/* if */
			nil,		/* > */
			shift(13),		/* amount */
			nil,		/* then */
			
		},

	},
	actionRow{ // S13
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* id */
			nil,		/* if */
			nil,		/* > */
			nil,		/* amount */
			shift(14),		/* then */
			
		},

	},
	actionRow{ // S14
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			shift(15),		/* id */
			nil,		/* if */
			nil,		/* > */
			nil,		/* amount */
			nil,		/* then */
			
		},

	},
	actionRow{ // S15
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(3),		/* $, reduce: Rule */
			nil,		/* check */
			nil,		/* id */
			reduce(3),		/* if, reduce: Rule */
			nil,		/* > */
			nil,		/* amount */
			nil,		/* then */
			
		},

	},
	
}


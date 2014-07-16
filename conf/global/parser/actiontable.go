
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
			nil,		/* [ */
			shift(3),		/* value */
			nil,		/* ] */
			nil,		/* : */
			
		},

	},
	actionRow{ // S1
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			accept(true),		/* $ */
			nil,		/* [ */
			nil,		/* value */
			nil,		/* ] */
			nil,		/* : */
			
		},

	},
	actionRow{ // S2
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: Config */
			shift(7),		/* [ */
			shift(3),		/* value */
			nil,		/* ] */
			nil,		/* : */
			
		},

	},
	actionRow{ // S3
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* [ */
			nil,		/* value */
			nil,		/* ] */
			reduce(8),		/* :, reduce: Key */
			
		},

	},
	actionRow{ // S4
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(5),		/* $, reduce: ConfigList */
			reduce(5),		/* [, reduce: ConfigList */
			reduce(5),		/* value, reduce: ConfigList */
			nil,		/* ] */
			nil,		/* : */
			
		},

	},
	actionRow{ // S5
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* [ */
			nil,		/* value */
			nil,		/* ] */
			shift(9),		/* : */
			
		},

	},
	actionRow{ // S6
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(2),		/* $, reduce: Config */
			shift(10),		/* [ */
			nil,		/* value */
			nil,		/* ] */
			nil,		/* : */
			
		},

	},
	actionRow{ // S7
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* [ */
			shift(11),		/* value */
			nil,		/* ] */
			nil,		/* : */
			
		},

	},
	actionRow{ // S8
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(6),		/* $, reduce: ConfigList */
			reduce(6),		/* [, reduce: ConfigList */
			reduce(6),		/* value, reduce: ConfigList */
			nil,		/* ] */
			nil,		/* : */
			
		},

	},
	actionRow{ // S9
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* [ */
			shift(12),		/* value */
			nil,		/* ] */
			nil,		/* : */
			
		},

	},
	actionRow{ // S10
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* [ */
			shift(14),		/* value */
			nil,		/* ] */
			nil,		/* : */
			
		},

	},
	actionRow{ // S11
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* [ */
			nil,		/* value */
			shift(15),		/* ] */
			nil,		/* : */
			
		},

	},
	actionRow{ // S12
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(9),		/* $, reduce: Value */
			reduce(9),		/* [, reduce: Value */
			reduce(9),		/* value, reduce: Value */
			nil,		/* ] */
			nil,		/* : */
			
		},

	},
	actionRow{ // S13
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(7),		/* $, reduce: KVPair */
			reduce(7),		/* [, reduce: KVPair */
			reduce(7),		/* value, reduce: KVPair */
			nil,		/* ] */
			nil,		/* : */
			
		},

	},
	actionRow{ // S14
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* [ */
			nil,		/* value */
			shift(16),		/* ] */
			nil,		/* : */
			
		},

	},
	actionRow{ // S15
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* [ */
			shift(3),		/* value */
			nil,		/* ] */
			nil,		/* : */
			
		},

	},
	actionRow{ // S16
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* [ */
			shift(3),		/* value */
			nil,		/* ] */
			nil,		/* : */
			
		},

	},
	actionRow{ // S17
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(3),		/* $, reduce: ContextList */
			reduce(3),		/* [, reduce: ContextList */
			shift(3),		/* value */
			nil,		/* ] */
			nil,		/* : */
			
		},

	},
	actionRow{ // S18
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(4),		/* $, reduce: ContextList */
			reduce(4),		/* [, reduce: ContextList */
			shift(3),		/* value */
			nil,		/* ] */
			nil,		/* : */
			
		},

	},
	
}


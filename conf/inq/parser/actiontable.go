
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
			nil,		/* name */
			nil,		/* host */
			shift(3),		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
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
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
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
			nil,		/* name */
			shift(5),		/* host */
			nil,		/* check */
			shift(6),		/* process */
			shift(7),		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
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
			nil,		/* name */
			reduce(4),		/* host, reduce: Preamble */
			nil,		/* check */
			reduce(4),		/* process, reduce: Preamble */
			reduce(4),		/* service, reduce: Preamble */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
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
			shift(8),		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
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
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			shift(11),		/* if */
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
			reduce(5),		/* name, reduce: Checktype */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
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
			reduce(6),		/* name, reduce: Checktype */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
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
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			shift(14),		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			shift(11),		/* if */
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
			reduce(3),		/* $, reduce: Check */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			shift(11),		/* if */
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
			reduce(23),		/* $, reduce: RuleList */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			reduce(23),		/* if, reduce: RuleList */
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
			nil,		/* $ */
			shift(16),		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
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
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			shift(11),		/* if */
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
			reduce(2),		/* $, reduce: Check */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			shift(11),		/* if */
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
			nil,		/* $ */
			shift(19),		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
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
			reduce(24),		/* $, reduce: RuleList */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			reduce(24),		/* if, reduce: RuleList */
			nil,		/* operator */
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
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			shift(21),		/* ( */
			nil,		/* ) */
			nil,		/* if */
			reduce(18),		/* operator, reduce: Metric */
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
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			shift(22),		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S18
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: Check */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			shift(11),		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S19
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(23),		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S20
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			reduce(12),		/* if, reduce: ParameterList */
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
			shift(24),		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			shift(25),		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
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
			nil,		/* $ */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			shift(26),		/* uint_lit */
			shift(28),		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
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
			nil,		/* $ */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			shift(29),		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			reduce(14),		/* if, reduce: Parameters */
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
			nil,		/* $ */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			shift(30),		/* ) */
			nil,		/* if */
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
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			shift(31),		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S26
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			reduce(17),		/* then, reduce: HumanAmount */
			reduce(17),		/* for, reduce: HumanAmount */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S27
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			shift(32),		/* then */
			shift(33),		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S28
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			reduce(16),		/* then, reduce: HumanAmount */
			reduce(16),		/* for, reduce: HumanAmount */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S29
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(19),		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S30
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			reduce(19),		/* operator, reduce: Metric */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S31
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			reduce(20),		/* operator, reduce: Metric */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S32
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			shift(37),		/* restart */
			shift(38),		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S33
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			shift(40),		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S34
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			reduce(13),		/* if, reduce: Parameters */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S35
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(21),		/* $, reduce: Rule */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			reduce(21),		/* if, reduce: Rule */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S36
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(7),		/* $, reduce: ActionList */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			shift(41),		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			reduce(7),		/* if, reduce: ActionList */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S37
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(9),		/* $, reduce: Action */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			reduce(9),		/* ,, reduce: Action */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			reduce(9),		/* if, reduce: Action */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S38
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(10),		/* $, reduce: Action */
			shift(42),		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			reduce(10),		/* ,, reduce: Action */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			reduce(10),		/* if, reduce: Action */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S39
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			shift(43),		/* cycles */
			
		},

	},
	actionRow{ // S40
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			reduce(15),		/* cycles, reduce: IntAmount */
			
		},

	},
	actionRow{ // S41
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			shift(37),		/* restart */
			shift(38),		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S42
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(11),		/* $, reduce: Action */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			reduce(11),		/* ,, reduce: Action */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			reduce(11),		/* if, reduce: Action */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S43
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			shift(45),		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S44
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(8),		/* $, reduce: ActionList */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			reduce(8),		/* if, reduce: ActionList */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S45
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			shift(37),		/* restart */
			shift(38),		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S46
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(22),		/* $, reduce: Rule */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			reduce(22),		/* if, reduce: Rule */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	
}


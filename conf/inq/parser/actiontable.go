
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
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
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
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
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
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
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
			reduce(3),		/* host, reduce: Preamble */
			nil,		/* check */
			reduce(3),		/* process, reduce: Preamble */
			reduce(3),		/* service, reduce: Preamble */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
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
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
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
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
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
			reduce(4),		/* name, reduce: Checktype */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
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
			reduce(5),		/* name, reduce: Checktype */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
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
			nil,		/* restart */
			nil,		/* alert */
			shift(13),		/* with */
			nil,		/* , */
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
	actionRow{ // S9
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(2),		/* $, reduce: Check */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
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
			reduce(19),		/* $, reduce: RuleList */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			reduce(19),		/* if, reduce: RuleList */
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
			shift(15),		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
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
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
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
			nil,		/* $ */
			shift(18),		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
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
	actionRow{ // S14
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(20),		/* $, reduce: RuleList */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			reduce(20),		/* if, reduce: RuleList */
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
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			shift(20),		/* ( */
			nil,		/* ) */
			nil,		/* if */
			reduce(14),		/* operator, reduce: Metric */
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
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			shift(21),		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S17
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: Check */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
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
	actionRow{ // S18
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(22),		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
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
	actionRow{ // S19
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			reduce(8),		/* if, reduce: ParameterList */
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
			shift(23),		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
			nil,		/* uint_lit */
			shift(24),		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
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
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
			shift(25),		/* uint_lit */
			shift(27),		/* sized_uint_lit */
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
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			shift(28),		/* , */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			reduce(10),		/* if, reduce: Parameters */
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
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			shift(29),		/* ) */
			nil,		/* if */
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
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
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
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			reduce(13),		/* then, reduce: HumanAmount */
			reduce(13),		/* for, reduce: HumanAmount */
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
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			shift(31),		/* then */
			shift(32),		/* for */
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
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			reduce(12),		/* then, reduce: HumanAmount */
			reduce(12),		/* for, reduce: HumanAmount */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S28
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(18),		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
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
	actionRow{ // S29
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			reduce(15),		/* operator, reduce: Metric */
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
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			reduce(16),		/* operator, reduce: Metric */
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
			shift(35),		/* restart */
			shift(36),		/* alert */
			nil,		/* with */
			nil,		/* , */
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
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
			shift(38),		/* uint_lit */
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
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			reduce(9),		/* if, reduce: Parameters */
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
			reduce(17),		/* $, reduce: Rule */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			reduce(17),		/* if, reduce: Rule */
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
			reduce(6),		/* $, reduce: Action */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			reduce(6),		/* if, reduce: Action */
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
			reduce(7),		/* $, reduce: Action */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			reduce(7),		/* if, reduce: Action */
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
			nil,		/* $ */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			shift(39),		/* cycles */
			
		},

	},
	actionRow{ // S38
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			reduce(11),		/* cycles, reduce: IntAmount */
			
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
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			shift(40),		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
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
			shift(35),		/* restart */
			shift(36),		/* alert */
			nil,		/* with */
			nil,		/* , */
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
	actionRow{ // S41
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(18),		/* $, reduce: Rule */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* , */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			reduce(18),		/* if, reduce: Rule */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	
}


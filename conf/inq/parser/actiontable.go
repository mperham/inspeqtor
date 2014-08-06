
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
			reduce(5),		/* host, reduce: Preamble */
			nil,		/* check */
			reduce(5),		/* process, reduce: Preamble */
			reduce(5),		/* service, reduce: Preamble */
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
			shift(11),		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			shift(13),		/* if */
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
	actionRow{ // S7
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			reduce(7),		/* name, reduce: Checktype */
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
			shift(11),		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			shift(13),		/* if */
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
			shift(13),		/* if */
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
			reduce(4),		/* $, reduce: Check */
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
			shift(13),		/* if */
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
			shift(18),		/* name */
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
	actionRow{ // S13
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(20),		/* name */
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
	actionRow{ // S14
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
			shift(13),		/* if */
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
			shift(13),		/* if */
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
			shift(13),		/* if */
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
			reduce(25),		/* $, reduce: RuleList */
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
			reduce(25),		/* if, reduce: RuleList */
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
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			reduce(13),		/* if, reduce: ParameterList */
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
			shift(24),		/* ( */
			nil,		/* ) */
			nil,		/* if */
			reduce(19),		/* operator, reduce: Metric */
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
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			shift(25),		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S22
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
			shift(13),		/* if */
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
			shift(26),		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			reduce(15),		/* if, reduce: Parameters */
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
			shift(27),		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			shift(28),		/* uint_lit */
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
			shift(29),		/* uint_lit */
			shift(31),		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
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
			shift(18),		/* name */
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
			shift(33),		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
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
			shift(34),		/* ) */
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
			reduce(18),		/* then, reduce: HumanAmount */
			reduce(18),		/* for, reduce: HumanAmount */
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
			nil,		/* operator */
			shift(35),		/* then */
			shift(36),		/* for */
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
			nil,		/* operator */
			reduce(17),		/* then, reduce: HumanAmount */
			reduce(17),		/* for, reduce: HumanAmount */
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
			nil,		/* if */
			reduce(21),		/* operator, reduce: Metric */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S35
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
			shift(39),		/* restart */
			shift(40),		/* alert */
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
	actionRow{ // S36
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
			shift(42),		/* uint_lit */
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
	actionRow{ // S37
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
	actionRow{ // S38
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(8),		/* $, reduce: ActionList */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			shift(43),		/* , */
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
	actionRow{ // S39
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(10),		/* $, reduce: Action */
			nil,		/* name */
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
	actionRow{ // S40
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(11),		/* $, reduce: Action */
			shift(44),		/* name */
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
			shift(45),		/* cycles */
			
		},

	},
	actionRow{ // S42
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
			reduce(16),		/* cycles, reduce: IntAmount */
			
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
			shift(39),		/* restart */
			shift(40),		/* alert */
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
	actionRow{ // S44
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(12),		/* $, reduce: Action */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* process */
			nil,		/* service */
			reduce(12),		/* ,, reduce: Action */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			reduce(12),		/* if, reduce: Action */
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
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* uint_lit */
			nil,		/* sized_uint_lit */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			shift(47),		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S46
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(9),		/* $, reduce: ActionList */
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
			reduce(9),		/* if, reduce: ActionList */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S47
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
			shift(39),		/* restart */
			shift(40),		/* alert */
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
	actionRow{ // S48
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(23),		/* $, reduce: Rule */
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
			reduce(23),		/* if, reduce: Rule */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	
}


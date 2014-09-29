
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
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
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
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
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
			shift(6),		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
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
			reduce(5),		/* service, reduce: Preamble */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
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
			shift(7),		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
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
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			shift(10),		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			shift(12),		/* if */
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
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
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
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			shift(10),		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			shift(12),		/* if */
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
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			shift(12),		/* if */
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
			reduce(4),		/* $, reduce: Check */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
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
			nil,		/* $ */
			shift(17),		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
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
			reduce(22),		/* $, reduce: RuleList */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			reduce(22),		/* if, reduce: RuleList */
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
			shift(19),		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
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
			nil,		/* $ */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
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
			reduce(2),		/* $, reduce: Check */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			shift(12),		/* if */
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
			reduce(3),		/* $, reduce: Check */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			shift(12),		/* if */
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
			reduce(23),		/* $, reduce: RuleList */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			reduce(23),		/* if, reduce: RuleList */
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
			shift(22),		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
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
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			reduce(12),		/* if, reduce: ParameterList */
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
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			shift(23),		/* : */
			shift(24),		/* ( */
			nil,		/* ) */
			nil,		/* if */
			reduce(17),		/* operator, reduce: Metric */
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
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			shift(25),		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S21
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: Check */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			shift(12),		/* if */
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
			nil,		/* service */
			shift(26),		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			reduce(14),		/* if, reduce: Parameters */
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
			shift(27),		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
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
			shift(28),		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
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
			shift(29),		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
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
			shift(17),		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
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
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			reduce(18),		/* operator, reduce: Metric */
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
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			shift(32),		/* ) */
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
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			reduce(16),		/* then, reduce: HumanAmount */
			reduce(16),		/* for, reduce: HumanAmount */
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
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			shift(33),		/* then */
			shift(34),		/* for */
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
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			reduce(13),		/* if, reduce: Parameters */
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
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			reduce(19),		/* operator, reduce: Metric */
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
			nil,		/* service */
			nil,		/* , */
			shift(37),		/* restart */
			shift(38),		/* alert */
			nil,		/* with */
			nil,		/* : */
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
			shift(39),		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
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
			reduce(20),		/* $, reduce: Rule */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			reduce(20),		/* if, reduce: Rule */
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
			nil,		/* service */
			shift(41),		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
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
			nil,		/* service */
			reduce(9),		/* ,, reduce: Action */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
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
			nil,		/* service */
			reduce(10),		/* ,, reduce: Action */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
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
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			reduce(15),		/* cycles, reduce: IntAmount */
			
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
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			shift(43),		/* cycles */
			
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
			nil,		/* service */
			nil,		/* , */
			shift(37),		/* restart */
			shift(38),		/* alert */
			nil,		/* with */
			nil,		/* : */
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
			nil,		/* service */
			reduce(11),		/* ,, reduce: Action */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
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
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
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
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
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
			nil,		/* service */
			nil,		/* , */
			shift(37),		/* restart */
			shift(38),		/* alert */
			nil,		/* with */
			nil,		/* : */
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
			reduce(21),		/* $, reduce: Rule */
			nil,		/* name */
			nil,		/* host */
			nil,		/* check */
			nil,		/* service */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			reduce(21),		/* if, reduce: Rule */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	
}


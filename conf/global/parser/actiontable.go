package parser

type (
	actionTable [numStates]actionRow
	actionRow   struct {
		canRecover bool
		actions    [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
		canRecover: false,
		actions: [numSymbols]action{
			nil,      /* INVALID */
			nil,      /* $ */
			shift(4), /* send */
			nil,      /* alerts */
			nil,      /* via */
			nil,      /* value */
			nil,      /* with */
			nil,      /* to */
			nil,      /* , */
			shift(5), /* set */

		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: [numSymbols]action{
			nil,          /* INVALID */
			accept(true), /* $ */
			nil,          /* send */
			nil,          /* alerts */
			nil,          /* via */
			nil,          /* value */
			nil,          /* with */
			nil,          /* to */
			nil,          /* , */
			nil,          /* set */

		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(1), /* $, reduce: Config */
			shift(4),  /* send */
			nil,       /* alerts */
			nil,       /* via */
			nil,       /* value */
			nil,       /* with */
			nil,       /* to */
			nil,       /* , */
			shift(5),  /* set */

		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(2), /* $, reduce: Config */
			shift(4),  /* send */
			nil,       /* alerts */
			nil,       /* via */
			nil,       /* value */
			nil,       /* with */
			nil,       /* to */
			nil,       /* , */
			shift(5),  /* set */

		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numSymbols]action{
			nil,      /* INVALID */
			nil,      /* $ */
			nil,      /* send */
			shift(8), /* alerts */
			nil,      /* via */
			nil,      /* value */
			nil,      /* with */
			nil,      /* to */
			nil,      /* , */
			nil,      /* set */

		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numSymbols]action{
			nil,      /* INVALID */
			nil,      /* $ */
			nil,      /* send */
			nil,      /* alerts */
			nil,      /* via */
			shift(9), /* value */
			nil,      /* with */
			nil,      /* to */
			nil,      /* , */
			nil,      /* set */

		},
	},
	actionRow{ // S6
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(3), /* $, reduce: Config */
			nil,       /* send */
			nil,       /* alerts */
			nil,       /* via */
			nil,       /* value */
			nil,       /* with */
			nil,       /* to */
			nil,       /* , */
			nil,       /* set */

		},
	},
	actionRow{ // S7
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(4), /* $, reduce: Config */
			nil,       /* send */
			nil,       /* alerts */
			nil,       /* via */
			nil,       /* value */
			nil,       /* with */
			nil,       /* to */
			nil,       /* , */
			nil,       /* set */

		},
	},
	actionRow{ // S8
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* send */
			nil,       /* alerts */
			shift(10), /* via */
			nil,       /* value */
			nil,       /* with */
			shift(11), /* to */
			nil,       /* , */
			nil,       /* set */

		},
	},
	actionRow{ // S9
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* send */
			nil,       /* alerts */
			nil,       /* via */
			shift(12), /* value */
			nil,       /* with */
			nil,       /* to */
			nil,       /* , */
			nil,       /* set */

		},
	},
	actionRow{ // S10
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* send */
			nil,       /* alerts */
			nil,       /* via */
			shift(13), /* value */
			nil,       /* with */
			nil,       /* to */
			nil,       /* , */
			nil,       /* set */

		},
	},
	actionRow{ // S11
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* send */
			nil,       /* alerts */
			nil,       /* via */
			shift(14), /* value */
			nil,       /* with */
			nil,       /* to */
			nil,       /* , */
			nil,       /* set */

		},
	},
	actionRow{ // S12
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(9), /* $, reduce: SetStatement */
			reduce(9), /* send, reduce: SetStatement */
			nil,       /* alerts */
			nil,       /* via */
			nil,       /* value */
			nil,       /* with */
			nil,       /* to */
			nil,       /* , */
			reduce(9), /* set, reduce: SetStatement */

		},
	},
	actionRow{ // S13
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* send */
			nil,       /* alerts */
			nil,       /* via */
			nil,       /* value */
			shift(15), /* with */
			nil,       /* to */
			nil,       /* , */
			nil,       /* set */

		},
	},
	actionRow{ // S14
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* send */
			nil,       /* alerts */
			shift(16), /* via */
			nil,       /* value */
			nil,       /* with */
			nil,       /* to */
			nil,       /* , */
			nil,       /* set */

		},
	},
	actionRow{ // S15
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* send */
			nil,       /* alerts */
			nil,       /* via */
			shift(17), /* value */
			nil,       /* with */
			nil,       /* to */
			nil,       /* , */
			nil,       /* set */

		},
	},
	actionRow{ // S16
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* send */
			nil,       /* alerts */
			nil,       /* via */
			shift(19), /* value */
			nil,       /* with */
			nil,       /* to */
			nil,       /* , */
			nil,       /* set */

		},
	},
	actionRow{ // S17
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* send */
			nil,       /* alerts */
			nil,       /* via */
			shift(20), /* value */
			nil,       /* with */
			nil,       /* to */
			nil,       /* , */
			nil,       /* set */

		},
	},
	actionRow{ // S18
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(5), /* $, reduce: RouteStatement */
			reduce(5), /* send, reduce: RouteStatement */
			nil,       /* alerts */
			nil,       /* via */
			nil,       /* value */
			nil,       /* with */
			nil,       /* to */
			nil,       /* , */
			reduce(5), /* set, reduce: RouteStatement */

		},
	},
	actionRow{ // S19
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* send */
			nil,       /* alerts */
			nil,       /* via */
			nil,       /* value */
			shift(21), /* with */
			nil,       /* to */
			nil,       /* , */
			nil,       /* set */

		},
	},
	actionRow{ // S20
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(7), /* $, reduce: ChannelParameters */
			reduce(7), /* send, reduce: ChannelParameters */
			nil,       /* alerts */
			nil,       /* via */
			nil,       /* value */
			nil,       /* with */
			nil,       /* to */
			shift(22), /* , */
			reduce(7), /* set, reduce: ChannelParameters */

		},
	},
	actionRow{ // S21
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* send */
			nil,       /* alerts */
			nil,       /* via */
			shift(17), /* value */
			nil,       /* with */
			nil,       /* to */
			nil,       /* , */
			nil,       /* set */

		},
	},
	actionRow{ // S22
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			nil,       /* $ */
			nil,       /* send */
			nil,       /* alerts */
			nil,       /* via */
			shift(17), /* value */
			nil,       /* with */
			nil,       /* to */
			nil,       /* , */
			nil,       /* set */

		},
	},
	actionRow{ // S23
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(6), /* $, reduce: RouteStatement */
			reduce(6), /* send, reduce: RouteStatement */
			nil,       /* alerts */
			nil,       /* via */
			nil,       /* value */
			nil,       /* with */
			nil,       /* to */
			nil,       /* , */
			reduce(6), /* set, reduce: RouteStatement */

		},
	},
	actionRow{ // S24
		canRecover: false,
		actions: [numSymbols]action{
			nil,       /* INVALID */
			reduce(8), /* $, reduce: ChannelParameters */
			reduce(8), /* send, reduce: ChannelParameters */
			nil,       /* alerts */
			nil,       /* via */
			nil,       /* value */
			nil,       /* with */
			nil,       /* to */
			nil,       /* , */
			reduce(8), /* set, reduce: ChannelParameters */

		},
	},
}

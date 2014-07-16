
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
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
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
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
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
			shift(7),		/* host */
			nil,		/* check */
			nil,		/* inspect */
			shift(8),		/* process */
			shift(9),		/* service */
			nil,		/* restart */
			nil,		/* alert */
			shift(10),		/* launchctl */
			shift(11),		/* upstart */
			shift(12),		/* runit */
			shift(13),		/* systemd */
			shift(14),		/* init.d */
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
			reduce(4),		/* host, reduce: Preamble */
			nil,		/* check */
			nil,		/* inspect */
			reduce(4),		/* process, reduce: Preamble */
			reduce(4),		/* service, reduce: Preamble */
			nil,		/* restart */
			nil,		/* alert */
			reduce(4),		/* launchctl, reduce: Preamble */
			reduce(4),		/* upstart, reduce: Preamble */
			reduce(4),		/* runit, reduce: Preamble */
			reduce(4),		/* systemd, reduce: Preamble */
			reduce(4),		/* init.d, reduce: Preamble */
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
			reduce(5),		/* host, reduce: Preamble */
			nil,		/* check */
			nil,		/* inspect */
			reduce(5),		/* process, reduce: Preamble */
			reduce(5),		/* service, reduce: Preamble */
			nil,		/* restart */
			nil,		/* alert */
			reduce(5),		/* launchctl, reduce: Preamble */
			reduce(5),		/* upstart, reduce: Preamble */
			reduce(5),		/* runit, reduce: Preamble */
			reduce(5),		/* systemd, reduce: Preamble */
			reduce(5),		/* init.d, reduce: Preamble */
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
			shift(15),		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
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
			shift(8),		/* process */
			shift(9),		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			nil,		/* uint_lit */
			nil,		/* sizecode */
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
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			shift(19),		/* if */
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
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
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
			reduce(7),		/* id, reduce: Checktype */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			nil,		/* if */
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
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			reduce(10),		/* process, reduce: Inittype */
			reduce(10),		/* service, reduce: Inittype */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			nil,		/* uint_lit */
			nil,		/* sizecode */
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
			nil,		/* $ */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			reduce(11),		/* process, reduce: Inittype */
			reduce(11),		/* service, reduce: Inittype */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			nil,		/* uint_lit */
			nil,		/* sizecode */
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
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			reduce(12),		/* process, reduce: Inittype */
			reduce(12),		/* service, reduce: Inittype */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
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
			nil,		/* $ */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			reduce(13),		/* process, reduce: Inittype */
			reduce(13),		/* service, reduce: Inittype */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			nil,		/* uint_lit */
			nil,		/* sizecode */
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
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			reduce(14),		/* process, reduce: Inittype */
			reduce(14),		/* service, reduce: Inittype */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			nil,		/* uint_lit */
			nil,		/* sizecode */
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
			nil,		/* $ */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			shift(19),		/* if */
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
			shift(21),		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			nil,		/* uint_lit */
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
			reduce(3),		/* $, reduce: Check */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			shift(19),		/* if */
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
			reduce(20),		/* $, reduce: RuleList */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			reduce(20),		/* if, reduce: RuleList */
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
			shift(23),		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			nil,		/* uint_lit */
			nil,		/* sizecode */
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
			reduce(1),		/* $, reduce: Check */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			shift(19),		/* if */
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
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			shift(19),		/* if */
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
			reduce(21),		/* $, reduce: RuleList */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			reduce(21),		/* if, reduce: RuleList */
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
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			nil,		/* if */
			shift(25),		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S24
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
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			shift(19),		/* if */
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
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			shift(26),		/* uint_lit */
			nil,		/* sizecode */
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
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			nil,		/* uint_lit */
			shift(28),		/* sizecode */
			nil,		/* if */
			nil,		/* operator */
			reduce(16),		/* then, reduce: HumanAmount */
			reduce(16),		/* for, reduce: HumanAmount */
			nil,		/* cycles */
			
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
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			nil,		/* if */
			nil,		/* operator */
			shift(29),		/* then */
			shift(30),		/* for */
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
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			nil,		/* if */
			nil,		/* operator */
			reduce(17),		/* then, reduce: HumanAmount */
			reduce(17),		/* for, reduce: HumanAmount */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S29
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
			shift(32),		/* restart */
			shift(33),		/* alert */
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			nil,		/* uint_lit */
			nil,		/* sizecode */
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
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			shift(35),		/* uint_lit */
			nil,		/* sizecode */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S31
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(18),		/* $, reduce: Rule */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			reduce(18),		/* if, reduce: Rule */
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
			reduce(8),		/* $, reduce: Action */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			reduce(8),		/* if, reduce: Action */
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
			reduce(9),		/* $, reduce: Action */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			reduce(9),		/* if, reduce: Action */
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
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			shift(36),		/* cycles */
			
		},

	},
	actionRow{ // S35
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
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			reduce(15),		/* cycles, reduce: IntAmount */
			
		},

	},
	actionRow{ // S36
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
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			nil,		/* if */
			nil,		/* operator */
			shift(37),		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S37
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
			shift(32),		/* restart */
			shift(33),		/* alert */
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			nil,		/* if */
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
			reduce(19),		/* $, reduce: Rule */
			nil,		/* id */
			nil,		/* host */
			nil,		/* check */
			nil,		/* inspect */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			nil,		/* uint_lit */
			nil,		/* sizecode */
			reduce(19),		/* if, reduce: Rule */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	
}


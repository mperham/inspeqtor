
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
			nil,		/* host */
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
			nil,		/* check */
			nil,		/* id */
			nil,		/* host */
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
			nil,		/* check */
			nil,		/* id */
			shift(5),		/* host */
			shift(6),		/* process */
			shift(7),		/* service */
			nil,		/* restart */
			nil,		/* alert */
			shift(8),		/* launchctl */
			shift(9),		/* upstart */
			shift(10),		/* runit */
			shift(11),		/* systemd */
			shift(12),		/* init.d */
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
			nil,		/* check */
			shift(13),		/* id */
			nil,		/* host */
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
	actionRow{ // S4
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* id */
			nil,		/* host */
			shift(6),		/* process */
			shift(7),		/* service */
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
	actionRow{ // S5
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* id */
			nil,		/* host */
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
			shift(17),		/* if */
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
			nil,		/* check */
			reduce(4),		/* id, reduce: Checktype */
			nil,		/* host */
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
	actionRow{ // S7
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			reduce(5),		/* id, reduce: Checktype */
			nil,		/* host */
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
	actionRow{ // S8
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* id */
			nil,		/* host */
			reduce(8),		/* process, reduce: Inittype */
			reduce(8),		/* service, reduce: Inittype */
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
			nil,		/* check */
			nil,		/* id */
			nil,		/* host */
			reduce(9),		/* process, reduce: Inittype */
			reduce(9),		/* service, reduce: Inittype */
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
			nil,		/* check */
			nil,		/* id */
			nil,		/* host */
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
			nil,		/* check */
			nil,		/* id */
			nil,		/* host */
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
			nil,		/* check */
			nil,		/* id */
			nil,		/* host */
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
			nil,		/* check */
			nil,		/* id */
			nil,		/* host */
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
			shift(17),		/* if */
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
			nil,		/* check */
			shift(19),		/* id */
			nil,		/* host */
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
	actionRow{ // S15
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(3),		/* $, reduce: Check */
			nil,		/* check */
			nil,		/* id */
			nil,		/* host */
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
			shift(17),		/* if */
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
			reduce(18),		/* $, reduce: RuleList */
			nil,		/* check */
			nil,		/* id */
			nil,		/* host */
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
			reduce(18),		/* if, reduce: RuleList */
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
			nil,		/* check */
			shift(21),		/* id */
			nil,		/* host */
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
	actionRow{ // S18
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: Check */
			nil,		/* check */
			nil,		/* id */
			nil,		/* host */
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
			shift(17),		/* if */
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
			nil,		/* check */
			nil,		/* id */
			nil,		/* host */
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
			shift(17),		/* if */
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
			reduce(19),		/* $, reduce: RuleList */
			nil,		/* check */
			nil,		/* id */
			nil,		/* host */
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
			reduce(19),		/* if, reduce: RuleList */
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
			nil,		/* check */
			nil,		/* id */
			nil,		/* host */
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
			shift(23),		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S22
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(2),		/* $, reduce: Check */
			nil,		/* check */
			nil,		/* id */
			nil,		/* host */
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
			shift(17),		/* if */
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
			nil,		/* check */
			nil,		/* id */
			nil,		/* host */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			shift(24),		/* uint_lit */
			nil,		/* sizecode */
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
			nil,		/* check */
			nil,		/* id */
			nil,		/* host */
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
			shift(26),		/* sizecode */
			nil,		/* if */
			nil,		/* operator */
			reduce(14),		/* then, reduce: HumanAmount */
			reduce(14),		/* for, reduce: HumanAmount */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S25
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* id */
			nil,		/* host */
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
			shift(27),		/* then */
			shift(28),		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S26
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* id */
			nil,		/* host */
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
			reduce(15),		/* then, reduce: HumanAmount */
			reduce(15),		/* for, reduce: HumanAmount */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S27
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* id */
			nil,		/* host */
			nil,		/* process */
			nil,		/* service */
			shift(30),		/* restart */
			shift(31),		/* alert */
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
	actionRow{ // S28
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* id */
			nil,		/* host */
			nil,		/* process */
			nil,		/* service */
			nil,		/* restart */
			nil,		/* alert */
			nil,		/* launchctl */
			nil,		/* upstart */
			nil,		/* runit */
			nil,		/* systemd */
			nil,		/* init.d */
			shift(33),		/* uint_lit */
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
			reduce(16),		/* $, reduce: Rule */
			nil,		/* check */
			nil,		/* id */
			nil,		/* host */
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
			reduce(16),		/* if, reduce: Rule */
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
			reduce(6),		/* $, reduce: Action */
			nil,		/* check */
			nil,		/* id */
			nil,		/* host */
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
			reduce(6),		/* if, reduce: Action */
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
			reduce(7),		/* $, reduce: Action */
			nil,		/* check */
			nil,		/* id */
			nil,		/* host */
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
			reduce(7),		/* if, reduce: Action */
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
			nil,		/* check */
			nil,		/* id */
			nil,		/* host */
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
			shift(34),		/* cycles */
			
		},

	},
	actionRow{ // S33
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* id */
			nil,		/* host */
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
			reduce(13),		/* cycles, reduce: IntAmount */
			
		},

	},
	actionRow{ // S34
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* id */
			nil,		/* host */
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
			shift(35),		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S35
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* id */
			nil,		/* host */
			nil,		/* process */
			nil,		/* service */
			shift(30),		/* restart */
			shift(31),		/* alert */
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
	actionRow{ // S36
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(17),		/* $, reduce: Rule */
			nil,		/* check */
			nil,		/* id */
			nil,		/* host */
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
			reduce(17),		/* if, reduce: Rule */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	
}


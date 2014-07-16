
/*
*/
package parser

const numNTSymbols = 7
type(
	gotoTable [numStates]gotoRow
	gotoRow	[numNTSymbols] int
)

var gotoTab = gotoTable{
	gotoRow{ // S0
		
		-1, // S'
		1, // Config
		-1, // ContextList
		2, // ConfigList
		4, // KVPair
		5, // Key
		-1, // Value
		

	},
	gotoRow{ // S1
		
		-1, // S'
		-1, // Config
		-1, // ContextList
		-1, // ConfigList
		-1, // KVPair
		-1, // Key
		-1, // Value
		

	},
	gotoRow{ // S2
		
		-1, // S'
		-1, // Config
		6, // ContextList
		-1, // ConfigList
		8, // KVPair
		5, // Key
		-1, // Value
		

	},
	gotoRow{ // S3
		
		-1, // S'
		-1, // Config
		-1, // ContextList
		-1, // ConfigList
		-1, // KVPair
		-1, // Key
		-1, // Value
		

	},
	gotoRow{ // S4
		
		-1, // S'
		-1, // Config
		-1, // ContextList
		-1, // ConfigList
		-1, // KVPair
		-1, // Key
		-1, // Value
		

	},
	gotoRow{ // S5
		
		-1, // S'
		-1, // Config
		-1, // ContextList
		-1, // ConfigList
		-1, // KVPair
		-1, // Key
		-1, // Value
		

	},
	gotoRow{ // S6
		
		-1, // S'
		-1, // Config
		-1, // ContextList
		-1, // ConfigList
		-1, // KVPair
		-1, // Key
		-1, // Value
		

	},
	gotoRow{ // S7
		
		-1, // S'
		-1, // Config
		-1, // ContextList
		-1, // ConfigList
		-1, // KVPair
		-1, // Key
		-1, // Value
		

	},
	gotoRow{ // S8
		
		-1, // S'
		-1, // Config
		-1, // ContextList
		-1, // ConfigList
		-1, // KVPair
		-1, // Key
		-1, // Value
		

	},
	gotoRow{ // S9
		
		-1, // S'
		-1, // Config
		-1, // ContextList
		-1, // ConfigList
		-1, // KVPair
		-1, // Key
		13, // Value
		

	},
	gotoRow{ // S10
		
		-1, // S'
		-1, // Config
		-1, // ContextList
		-1, // ConfigList
		-1, // KVPair
		-1, // Key
		-1, // Value
		

	},
	gotoRow{ // S11
		
		-1, // S'
		-1, // Config
		-1, // ContextList
		-1, // ConfigList
		-1, // KVPair
		-1, // Key
		-1, // Value
		

	},
	gotoRow{ // S12
		
		-1, // S'
		-1, // Config
		-1, // ContextList
		-1, // ConfigList
		-1, // KVPair
		-1, // Key
		-1, // Value
		

	},
	gotoRow{ // S13
		
		-1, // S'
		-1, // Config
		-1, // ContextList
		-1, // ConfigList
		-1, // KVPair
		-1, // Key
		-1, // Value
		

	},
	gotoRow{ // S14
		
		-1, // S'
		-1, // Config
		-1, // ContextList
		-1, // ConfigList
		-1, // KVPair
		-1, // Key
		-1, // Value
		

	},
	gotoRow{ // S15
		
		-1, // S'
		-1, // Config
		-1, // ContextList
		17, // ConfigList
		4, // KVPair
		5, // Key
		-1, // Value
		

	},
	gotoRow{ // S16
		
		-1, // S'
		-1, // Config
		-1, // ContextList
		18, // ConfigList
		4, // KVPair
		5, // Key
		-1, // Value
		

	},
	gotoRow{ // S17
		
		-1, // S'
		-1, // Config
		-1, // ContextList
		-1, // ConfigList
		8, // KVPair
		5, // Key
		-1, // Value
		

	},
	gotoRow{ // S18
		
		-1, // S'
		-1, // Config
		-1, // ContextList
		-1, // ConfigList
		8, // KVPair
		5, // Key
		-1, // Value
		

	},
	
}

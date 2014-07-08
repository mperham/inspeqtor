
/*
*/
package parser

const numNTSymbols = 4
type(
	gotoTable [numStates]gotoRow
	gotoRow	[numNTSymbols] int
)

var gotoTab = gotoTable{
	gotoRow{ // S0
		
		-1, // S'
		1, // Check
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S1
		
		-1, // S'
		-1, // Check
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S2
		
		-1, // S'
		-1, // Check
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S3
		
		-1, // S'
		-1, // Check
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S4
		
		-1, // S'
		-1, // Check
		7, // Rule
		6, // RuleList
		

	},
	gotoRow{ // S5
		
		-1, // S'
		-1, // Check
		7, // Rule
		9, // RuleList
		

	},
	gotoRow{ // S6
		
		-1, // S'
		-1, // Check
		10, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S7
		
		-1, // S'
		-1, // Check
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S8
		
		-1, // S'
		-1, // Check
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S9
		
		-1, // S'
		-1, // Check
		10, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S10
		
		-1, // S'
		-1, // Check
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S11
		
		-1, // S'
		-1, // Check
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S12
		
		-1, // S'
		-1, // Check
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S13
		
		-1, // S'
		-1, // Check
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S14
		
		-1, // S'
		-1, // Check
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S15
		
		-1, // S'
		-1, // Check
		-1, // Rule
		-1, // RuleList
		

	},
	
}

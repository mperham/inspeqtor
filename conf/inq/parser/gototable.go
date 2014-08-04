
/*
*/
package parser

const numNTSymbols = 12
type(
	gotoTable [numStates]gotoRow
	gotoRow	[numNTSymbols] int
)

var gotoTab = gotoTable{
	gotoRow{ // S0
		
		-1, // S'
		1, // Check
		2, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S1
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S2
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		4, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S3
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S4
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S5
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		10, // Rule
		9, // RuleList
		

	},
	gotoRow{ // S6
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S7
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S8
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		12, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S9
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		14, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S10
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S11
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		16, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S12
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		10, // Rule
		17, // RuleList
		

	},
	gotoRow{ // S13
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		19, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S14
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S15
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S16
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S17
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		14, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S18
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S19
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S20
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S21
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		26, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S22
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S23
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S24
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S25
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S26
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S27
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S28
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		33, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S29
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S30
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S31
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		34, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S32
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		37, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S33
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S34
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S35
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S36
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S37
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S38
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S39
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S40
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		41, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	gotoRow{ // S41
		
		-1, // S'
		-1, // Check
		-1, // Preamble
		-1, // Checktype
		-1, // Action
		-1, // ParameterList
		-1, // Parameters
		-1, // IntAmount
		-1, // HumanAmount
		-1, // Metric
		-1, // Rule
		-1, // RuleList
		

	},
	
}

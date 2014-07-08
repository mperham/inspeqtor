
package parser

import "inspeqtor/conf/ast"

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab {
	ProdTabEntry{
		String: `S' : Check	<<  >>`,
		Id: "S'",
		NTType: 0,
		Index: 0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Check : "check" id id id RuleList	<< ast.NewCheck(X[1], X[2], X[3], X[4]), nil >>`,
		Id: "Check",
		NTType: 1,
		Index: 1,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCheck(X[1], X[2], X[3], X[4]), nil
		},
	},
	ProdTabEntry{
		String: `Check : "check" id id RuleList	<< ast.NewCheck(nil, X[1], X[2], X[3]), nil >>`,
		Id: "Check",
		NTType: 1,
		Index: 2,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewCheck(nil, X[1], X[2], X[3]), nil
		},
	},
	ProdTabEntry{
		String: `Rule : "if" id ">" amount "then" id	<< ast.NewRule(X[1], X[2], X[3], X[5]), nil >>`,
		Id: "Rule",
		NTType: 2,
		Index: 3,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewRule(X[1], X[2], X[3], X[5]), nil
		},
	},
	ProdTabEntry{
		String: `RuleList : Rule	<< ast.NewRuleList(X[0]), nil >>`,
		Id: "RuleList",
		NTType: 3,
		Index: 4,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewRuleList(X[0]), nil
		},
	},
	ProdTabEntry{
		String: `RuleList : RuleList Rule	<< ast.AppendRule(X[0], X[1]), nil >>`,
		Id: "RuleList",
		NTType: 3,
		Index: 5,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendRule(X[0], X[1]), nil
		},
	},
	
}

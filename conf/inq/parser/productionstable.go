
package parser

import "inspeqtor/conf/inq/ast"

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
		String: `Check : Preamble Checktype id RuleList	<< ast.NewProcessCheck(X[1], X[2], X[3]), nil >>`,
		Id: "Check",
		NTType: 1,
		Index: 1,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewProcessCheck(X[1], X[2], X[3]), nil
		},
	},
	ProdTabEntry{
		String: `Check : Preamble "host" RuleList	<< ast.NewHostCheck(X[2]), nil >>`,
		Id: "Check",
		NTType: 1,
		Index: 2,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewHostCheck(X[2]), nil
		},
	},
	ProdTabEntry{
		String: `Preamble : "check"	<< X[0], nil >>`,
		Id: "Preamble",
		NTType: 2,
		Index: 3,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Preamble : "inspect"	<< X[0], nil >>`,
		Id: "Preamble",
		NTType: 2,
		Index: 4,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Checktype : "process"	<< X[0], nil >>`,
		Id: "Checktype",
		NTType: 3,
		Index: 5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Checktype : "service"	<< X[0], nil >>`,
		Id: "Checktype",
		NTType: 3,
		Index: 6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Action : "restart"	<< X[0], nil >>`,
		Id: "Action",
		NTType: 4,
		Index: 7,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Action : "alert"	<< X[0], nil >>`,
		Id: "Action",
		NTType: 4,
		Index: 8,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `IntAmount : uint_lit	<< ast.ToUint64(X[0]) >>`,
		Id: "IntAmount",
		NTType: 5,
		Index: 9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.ToUint64(X[0])
		},
	},
	ProdTabEntry{
		String: `HumanAmount : uint_lit	<< ast.HumanAmount(X[0], nil) >>`,
		Id: "HumanAmount",
		NTType: 6,
		Index: 10,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.HumanAmount(X[0], nil)
		},
	},
	ProdTabEntry{
		String: `HumanAmount : uint_lit sizecode	<< ast.HumanAmount(X[0], X[1]) >>`,
		Id: "HumanAmount",
		NTType: 6,
		Index: 11,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.HumanAmount(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Rule : "if" id operator HumanAmount "then" Action	<< ast.NewRule(X[1], X[2], X[3], X[5], uint64(1)), nil >>`,
		Id: "Rule",
		NTType: 7,
		Index: 12,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewRule(X[1], X[2], X[3], X[5], uint64(1)), nil
		},
	},
	ProdTabEntry{
		String: `Rule : "if" id operator HumanAmount "for" IntAmount "cycles" "then" Action	<< ast.NewRule(X[1], X[2], X[3], X[8], X[5]), nil >>`,
		Id: "Rule",
		NTType: 7,
		Index: 13,
		NumSymbols: 9,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewRule(X[1], X[2], X[3], X[8], X[5]), nil
		},
	},
	ProdTabEntry{
		String: `RuleList : Rule	<< ast.NewRuleList(X[0]), nil >>`,
		Id: "RuleList",
		NTType: 8,
		Index: 14,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewRuleList(X[0]), nil
		},
	},
	ProdTabEntry{
		String: `RuleList : RuleList Rule	<< ast.AppendRule(X[0], X[1]), nil >>`,
		Id: "RuleList",
		NTType: 8,
		Index: 15,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendRule(X[0], X[1]), nil
		},
	},
	
}

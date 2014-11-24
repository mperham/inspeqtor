package parser

import "github.com/mperham/inspeqtor/conf/inq/ast"

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Check	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Check : Preamble Checktype name ParameterList RuleList	<< ast.NewProcessCheck(X[1], X[2], X[4], X[3]), nil >>`,
		Id:         "Check",
		NTType:     1,
		Index:      1,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewProcessCheck(X[1], X[2], X[4], X[3]), nil
		},
	},
	ProdTabEntry{
		String: `Check : Preamble Checktype name RuleList	<< ast.NewProcessCheck(X[1], X[2], X[3], map[string]string{}), nil >>`,
		Id:         "Check",
		NTType:     1,
		Index:      2,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewProcessCheck(X[1], X[2], X[3], map[string]string{}), nil
		},
	},
	ProdTabEntry{
		String: `Check : Preamble "host" ParameterList RuleList	<< ast.NewHostCheck(X[3], X[2]), nil >>`,
		Id:         "Check",
		NTType:     1,
		Index:      3,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewHostCheck(X[3], X[2]), nil
		},
	},
	ProdTabEntry{
		String: `Check : Preamble "host" RuleList	<< ast.NewHostCheck(X[2], map[string]string{}), nil >>`,
		Id:         "Check",
		NTType:     1,
		Index:      4,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewHostCheck(X[2], map[string]string{}), nil
		},
	},
	ProdTabEntry{
		String: `Preamble : "check"	<< X[0], nil >>`,
		Id:         "Preamble",
		NTType:     2,
		Index:      5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Checktype : "service"	<< X[0], nil >>`,
		Id:         "Checktype",
		NTType:     3,
		Index:      6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `ActionList : Action	<< []ast.Action{X[0].(ast.Action)}, nil >>`,
		Id:         "ActionList",
		NTType:     4,
		Index:      7,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return []ast.Action{X[0].(ast.Action)}, nil
		},
	},
	ProdTabEntry{
		String: `ActionList : Action "," ActionList	<< ast.AppendAction(X[0], X[2]) >>`,
		Id:         "ActionList",
		NTType:     4,
		Index:      8,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendAction(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Action : "restart"	<< ast.AddAction(X[0], nil) >>`,
		Id:         "Action",
		NTType:     5,
		Index:      9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AddAction(X[0], nil)
		},
	},
	ProdTabEntry{
		String: `Action : "reload"	<< ast.AddAction(X[0], nil) >>`,
		Id:         "Action",
		NTType:     5,
		Index:      10,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AddAction(X[0], nil)
		},
	},
	ProdTabEntry{
		String: `Action : "alert"	<< ast.AddAction(X[0], nil) >>`,
		Id:         "Action",
		NTType:     5,
		Index:      11,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AddAction(X[0], nil)
		},
	},
	ProdTabEntry{
		String: `Action : "alert" name	<< ast.AddAction(X[0], X[1]) >>`,
		Id:         "Action",
		NTType:     5,
		Index:      12,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AddAction(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `ParameterList : "with" Parameters	<< X[1], nil >>`,
		Id:         "ParameterList",
		NTType:     6,
		Index:      13,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[1], nil
		},
	},
	ProdTabEntry{
		String: `Parameters : name name "," Parameters	<< ast.AddParam(X[0], X[1], X[3]) >>`,
		Id:         "Parameters",
		NTType:     7,
		Index:      14,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AddParam(X[0], X[1], X[3])
		},
	},
	ProdTabEntry{
		String: `Parameters : name name	<< ast.AddParam(X[0], X[1], nil) >>`,
		Id:         "Parameters",
		NTType:     7,
		Index:      15,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AddParam(X[0], X[1], nil)
		},
	},
	ProdTabEntry{
		String: `IntAmount : name	<< ast.ToInt64(X[0]) >>`,
		Id:         "IntAmount",
		NTType:     8,
		Index:      16,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.ToInt64(X[0])
		},
	},
	ProdTabEntry{
		String: `HumanAmount : name	<< ast.HumanAmount(X[0]) >>`,
		Id:         "HumanAmount",
		NTType:     9,
		Index:      17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.HumanAmount(X[0])
		},
	},
	ProdTabEntry{
		String: `Metric : name	<< ast.Metric(X[0], nil) >>`,
		Id:         "Metric",
		NTType:     10,
		Index:      18,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Metric(X[0], nil)
		},
	},
	ProdTabEntry{
		String: `Metric : name ":" name	<< ast.Metric(X[0], X[2]) >>`,
		Id:         "Metric",
		NTType:     10,
		Index:      19,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Metric(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Metric : name "(" name ")"	<< ast.Metric(X[0], X[2]) >>`,
		Id:         "Metric",
		NTType:     10,
		Index:      20,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.Metric(X[0], X[2])
		},
	},
	ProdTabEntry{
		String: `Rule : "if" Metric operator HumanAmount "then" ActionList	<< ast.NewRule(X[1], X[2], X[3], X[5], ast.ONE_CYCLE), nil >>`,
		Id:         "Rule",
		NTType:     11,
		Index:      21,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewRule(X[1], X[2], X[3], X[5], ast.ONE_CYCLE), nil
		},
	},
	ProdTabEntry{
		String: `Rule : "if" Metric operator HumanAmount "for" IntAmount "cycles" "then" ActionList	<< ast.NewRule(X[1], X[2], X[3], X[8], X[5]), nil >>`,
		Id:         "Rule",
		NTType:     11,
		Index:      22,
		NumSymbols: 9,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewRule(X[1], X[2], X[3], X[8], X[5]), nil
		},
	},
	ProdTabEntry{
		String: `RuleList : Rule	<< ast.NewRuleList(X[0]), nil >>`,
		Id:         "RuleList",
		NTType:     12,
		Index:      23,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewRuleList(X[0]), nil
		},
	},
	ProdTabEntry{
		String: `RuleList : RuleList Rule	<< ast.AppendRule(X[0], X[1]), nil >>`,
		Id:         "RuleList",
		NTType:     12,
		Index:      24,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendRule(X[0], X[1]), nil
		},
	},
}

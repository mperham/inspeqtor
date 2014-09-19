
package parser

import "github.com/mperham/inspeqtor/conf/global/ast"

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
		String: `S' : Config	<<  >>`,
		Id: "S'",
		NTType: 0,
		Index: 0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Config : SetStatement	<< ast.NewConfig(X[0].(map[string]string), nil) >>`,
		Id: "Config",
		NTType: 1,
		Index: 1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewConfig(X[0].(map[string]string), nil)
		},
	},
	ProdTabEntry{
		String: `Config : RouteStatement	<< ast.NewConfig(map[string]string{}, X[0]) >>`,
		Id: "Config",
		NTType: 1,
		Index: 2,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewConfig(map[string]string{}, X[0])
		},
	},
	ProdTabEntry{
		String: `Config : SetStatement Config	<< ast.AddSet(X[0], X[1]) >>`,
		Id: "Config",
		NTType: 1,
		Index: 3,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AddSet(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `Config : RouteStatement Config	<< ast.AddRoute(X[0], X[1]) >>`,
		Id: "Config",
		NTType: 1,
		Index: 4,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AddRoute(X[0], X[1])
		},
	},
	ProdTabEntry{
		String: `RouteStatement : "send" "alerts" "via" value "with" ChannelParameters	<< ast.NewRoute(nil, X[3], X[5]) >>`,
		Id: "RouteStatement",
		NTType: 2,
		Index: 5,
		NumSymbols: 6,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewRoute(nil, X[3], X[5])
		},
	},
	ProdTabEntry{
		String: `RouteStatement : "send" "alerts" "to" value "via" value "with" ChannelParameters	<< ast.NewRoute(X[3], X[5], X[7]) >>`,
		Id: "RouteStatement",
		NTType: 2,
		Index: 6,
		NumSymbols: 8,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewRoute(X[3], X[5], X[7])
		},
	},
	ProdTabEntry{
		String: `ChannelParameters : value value	<< ast.AppendPair(X[0], X[1], map[string]string{}) >>`,
		Id: "ChannelParameters",
		NTType: 3,
		Index: 7,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendPair(X[0], X[1], map[string]string{})
		},
	},
	ProdTabEntry{
		String: `ChannelParameters : value value "," ChannelParameters	<< ast.AppendPair(X[0], X[1], X[3]) >>`,
		Id: "ChannelParameters",
		NTType: 3,
		Index: 8,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendPair(X[0], X[1], X[3])
		},
	},
	ProdTabEntry{
		String: `SetStatement : "set" value value	<< ast.AppendPair(X[1], X[2], map[string]string{}) >>`,
		Id: "SetStatement",
		NTType: 4,
		Index: 9,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendPair(X[1], X[2], map[string]string{})
		},
	},
	
}

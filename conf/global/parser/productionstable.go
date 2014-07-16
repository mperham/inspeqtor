
package parser

import "inspeqtor/conf/global/ast"

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
		String: `Config : ConfigList	<< ast.GlobalConfig(X[0], nil), nil >>`,
		Id: "Config",
		NTType: 1,
		Index: 1,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.GlobalConfig(X[0], nil), nil
		},
	},
	ProdTabEntry{
		String: `Config : ConfigList ContextList	<< ast.GlobalConfig(X[0], X[1]), nil >>`,
		Id: "Config",
		NTType: 1,
		Index: 2,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.GlobalConfig(X[0], X[1]), nil
		},
	},
	ProdTabEntry{
		String: `ContextList : "[" value "]" ConfigList	<< ast.NewContext(X[1], X[3]), nil >>`,
		Id: "ContextList",
		NTType: 2,
		Index: 3,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewContext(X[1], X[3]), nil
		},
	},
	ProdTabEntry{
		String: `ContextList : ContextList "[" value "]" ConfigList	<< ast.AppendContext(X[0], X[2], X[4]), nil >>`,
		Id: "ContextList",
		NTType: 2,
		Index: 4,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendContext(X[0], X[2], X[4]), nil
		},
	},
	ProdTabEntry{
		String: `ConfigList : KVPair	<< ast.NewPair(X[0]), nil >>`,
		Id: "ConfigList",
		NTType: 3,
		Index: 5,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewPair(X[0]), nil
		},
	},
	ProdTabEntry{
		String: `ConfigList : ConfigList KVPair	<< ast.AppendPair(X[0], X[1]), nil >>`,
		Id: "ConfigList",
		NTType: 3,
		Index: 6,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.AppendPair(X[0], X[1]), nil
		},
	},
	ProdTabEntry{
		String: `KVPair : Key ":" Value	<< ast.NewKVPair(X[0], X[2]), nil >>`,
		Id: "KVPair",
		NTType: 4,
		Index: 7,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return ast.NewKVPair(X[0], X[2]), nil
		},
	},
	ProdTabEntry{
		String: `Key : value	<< X[0], nil >>`,
		Id: "Key",
		NTType: 5,
		Index: 8,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Value : value	<< X[0], nil >>`,
		Id: "Value",
		NTType: 6,
		Index: 9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	
}

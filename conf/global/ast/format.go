package ast

import (
	"inspeqtor/conf/global/token"
	"strings"
)

type KVPair struct {
	Key   string
	Value string
}

func NewKVPair(k interface{}, v interface{}) KVPair {
	key := string(k.(*token.Token).Lit)
	value := strings.Trim(string(v.(*token.Token).Lit), "\"")
	return KVPair{
		key, value,
	}
}

func NewPair(pair interface{}) map[string]string {
	hash := make(map[string]string, 1)
	hash[pair.(KVPair).Key] = pair.(KVPair).Value
	return hash
}

func AppendPair(hash interface{}, pair interface{}) map[string]string {
	h := hash.(map[string]string)
	h[pair.(KVPair).Key] = pair.(KVPair).Value
	return h
}

type Context struct {
	Name   string
	Config map[string]string
}

func NewContext(name interface{}, config interface{}) []Context {
	return []Context{
		Context{string(name.(*token.Token).Lit), config.(map[string]string)},
	}
}

func AppendContext(contexts interface{}, name interface{}, config interface{}) []Context {
	c := Context{string(name.(*token.Token).Lit), config.(map[string]string)}
	return append(contexts.([]Context), c)
}

type ConfigFile struct {
	TopConfig  map[string]string
	Contextual map[string]map[string]string
}

func GlobalConfig(config interface{}, contexts interface{}) ConfigFile {
	g := ConfigFile{
		TopConfig: config.(map[string]string),
	}

	c := contexts.([]Context)
	g.Contextual = make(map[string]map[string]string, len(c))
	for _, ctx := range c {
		g.Contextual[ctx.Name] = ctx.Config
	}
	return g
}

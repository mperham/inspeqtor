package ast

import (
	"inspeqtor/conf/global/token"
)

func AppendPair(k interface{}, v interface{}, hash interface{}) map[string]string {
	key := string(k.(*token.Token).Lit)
	value := string(v.(*token.Token).Lit)

	h := hash.(map[string]string)
	h[key] = value
	return h
}

type Route struct {
	Name    string
	Channel string
	Config  map[string]string
}

type Config struct {
	Variables map[string]string
	Routes    []Route
}

func NewRoute(name interface{}, channel interface{}, config interface{}) (Route, error) {
	nm := ""
	if name != nil {
		nm = string(name.(*token.Token).Lit)
	}
	return Route{nm, string(channel.(*token.Token).Lit), config.(map[string]string)}, nil
}

func AddRoute(route interface{}, config interface{}) (Config, error) {
	c := config.(Config)
	r := route.(Route)
	c.Routes = append(c.Routes, r)
	return c, nil
}

func AddSet(set interface{}, config interface{}) (Config, error) {
	c := config.(Config)
	vars := set.(map[string]string)
	for k, v := range vars {
		c.Variables[k] = v
	}
	return c, nil
}

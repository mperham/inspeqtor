package ast

import (
	"github.com/mperham/inspeqtor/conf/global/token"
	"strconv"
)

func AppendPair(k interface{}, v interface{}, hash interface{}) (map[string]string, error) {
	key := string(k.(*token.Token).Lit)
	value := string(v.(*token.Token).Lit)

	if value[0] == '"' {
		val, err := strconv.Unquote(value)
		if err != nil {
			return nil, err
		}
		value = val
	}

	h := hash.(map[string]string)
	h[key] = value
	return h, nil
}

type Route struct {
	Name    string
	Channel string
	Config  map[string]string
}

type Config struct {
	Variables map[string]string
	Routes    map[string]Route
}

func NewConfig(vars interface{}, route interface{}) (Config, error) {
	routes := map[string]Route{}

	if route != nil {
		r := route.(Route)
		routes[r.Name] = r
	}
	return Config{
		vars.(map[string]string),
		routes,
	}, nil
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
	c.Routes[r.Name] = r
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

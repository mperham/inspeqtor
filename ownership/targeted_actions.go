package ownership

import (
	"fmt"

	"github.com/mperham/inspeqtor"
	"github.com/mperham/inspeqtor/conf/inq/ast"
)

type OwnedAction struct {
	ActionName  string
	ActionOwner string
}

type Owned interface {
	Owner() string
}

func (oa *OwnedAction) Name() string  { return oa.ActionName }
func (oa *OwnedAction) Owner() string { return oa.ActionOwner }

func init() {
	ast.CreateAction = createOwnedAction
	inspeqtor.BuildAction = convertAction
}

func createOwnedAction(props ...string) (ast.Action, error) {
	return &OwnedAction{props[0], props[1]}, nil
}

func convertAction(global *inspeqtor.ConfigFile, check inspeqtor.Eventable, action ast.Action) (inspeqtor.Action, error) {
	switch action.Name() {
	case "alert":
		owner := ""

		own, ok := action.(Owned)
		if ok {
			owner = own.Owner()
		}
		if owner == "" {
			owner = check.Parameter("owner")
		}

		route := global.AlertRoutes[owner]
		if owner == "" && route == nil {
			return nil, fmt.Errorf("Please configure a \"send alerts\" statement in inspeqtor.conf.")
		}
		if route == nil {
			return nil, fmt.Errorf("No such alert route: %s", owner)
		}
		return inspeqtor.Actions["alert"](check, route)
	case "reload":
		return inspeqtor.Actions["reload"](check, nil)
	case "restart":
		return inspeqtor.Actions["restart"](check, nil)
	default:
		return nil, fmt.Errorf("Unknown action: %s", action.Name())
	}
}

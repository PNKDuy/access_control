package request

import (
	"access_control/model"
	"errors"
	"github.com/casbin/casbin/v2"
	casbinpgadapter "github.com/cychiuae/casbin-pg-adapter"
	"strings"
)

type Request struct {
	Role   string `json:"role"`
	Path   string `json:"path"`
	Method string `json:"method"`
	UserId string `json:"user_id"`
}

func (req Request)checkPermissionForUser() bool {
	if strings.Contains(req.Path, "/general/user/") &&!strings.Contains(req.Path, req.UserId) {
		return false
	}
	return true
}

func (req *Request)CheckPermissionService() (string, error) {
	db, err := model.ConnectToPostgresAndReturnEnforcer()
	if err != nil {

	}
	tableName := "casbin"
	adapter, err := casbinpgadapter.NewAdapter(db, tableName)

	e, err := casbin.NewEnforcer("model.conf", adapter)
	if err != nil {
		return "", err
	}
	if req.Role == "user" {
		if !req.checkPermissionForUser() {
			return "", errors.New("cannot access to other accounts")
		}
	}
	ok, err := e.Enforce(req.Role, req.Path, req.Method)
	if err != nil {
		//fmt.Println("Error occurred: ", err.Error())
		return "", nil
	}

	if ok == true {
		// allow
		//fmt.Printf("Allow %s %s %s %s\n", req.userId, req.role, req.path, req.method)
		return "Allow", nil
	} else {
		// deny the request, show an error
		//fmt.Printf("Deny %s %s %s %s\n", req.userId, req.role, req.path, req.method)
		return "Deny", nil
	}

}



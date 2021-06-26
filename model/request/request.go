package request

import (
	"access_control/model"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/casbin/casbin/v2"
	casbinpgadapter "github.com/cychiuae/casbin-pg-adapter"
	kafka "github.com/segmentio/kafka-go"
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

func (req *Request)CheckPermission() (string, error) {
	db, err := model.ConnectToPostgres()
	if err != nil {
		return "", errors.New("cannot connect to database")
	}
	tableName := "casbins"
	adapter, _ := casbinpgadapter.NewAdapter(db, tableName)

	e, err := casbin.NewEnforcer("model.conf", adapter)
	if err != nil {
		return "", err
	}
	if req.Role == "user" {
		if !req.checkPermissionForUser() {
			return "Deny", nil
		}
	}
	ok, err := e.Enforce(req.Role, req.Path, req.Method)
	if err != nil {
		//fmt.Println("Error occurred: ", err.Error())
		return "", nil
	}

	if ok{
		// allow
		//fmt.Printf("Allow %s %s %s %s\n", req.userId, req.role, req.path, req.method)
		msg := "Allow " + req.UserId + " access to " + req.Path + " method " + req.Method
		return msg, nil
	} else {
		// deny the request, show an error
		//fmt.Printf("Deny %s %s %s %s\n", req.userId, req.role, req.path, req.method)
		msg := "Deny " + req.UserId + " access to " + req.Path + " method " + req.Method
		return msg, nil
	}
}



func produceMassge(msg string) error {
	topic := "access-control"
	partition := 0
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(msg)

	conn , err := kafka.DialLeader(context.Background(), "tcp", "45.32.117.131:9092", topic, partition)
	if err != nil {
		return err
	}
	conn.SetWriteDeadline(time.Now().Add(10*time.Second))
	_, err = conn.WriteMessages(
    kafka.Message{Value: []byte(reqBodyBytes.Bytes())},
	)
	if err != nil {
		return err
	}
	if err := conn.Close(); err != nil {
		return err
	}
	return nil
}

func convertFromMessageValueToRequestModel(msgValue []byte) (Request, error) {
	request := Request{}
	err := json.Unmarshal(msgValue, &request)
	if err != nil {
		return request, err
	}
	return request, nil
}


package main

import (
	"bytes"
	"ecs-onboard/model"
	"encoding/json"
	"errors"
	"github.com/gookit/slog"
	"io"
	"net/http"
)

var (
	client = *http.DefaultClient
	url    = "http://0.0.0.0:8081/v1"
)

func main() {

	//tc1_createRoleNsNotFound()
	//tc2_createRoleUserHasAlreadyKeys()
	//tc3_createRoleUserHasOneKey()
	tc4_createRoleNoUser()
	//tc6_rotateRole()
	//tc5_deleteRole()
}

func tc5_deleteRole() {
	u := model.Role{
		Username:  "test3",
		Namespace: "ci12345-native-user",
		SafeId:    "test",
	}
	code, err := httpReq("DELETE", "/user/"+u.RoleName(), nil, nil)
	if err != nil {
		slog.Error(err)
	}
	if code != 200 {
		slog.Error(code)
	}
}

func tc6_rotateRole() {
	u := model.Role{
		Username:  "test3",
		Namespace: "ci12345",
		SafeId:    "test",
	}
	code, err := httpReq("POST", "/rotate-role/"+u.RoleName(), nil, nil)
	if err != nil {
		slog.Error(err)
	}
	if code != 200 {
		slog.Error(code)
	}
}

func tc4_createRoleNoUser() {
	u := model.Role{
		Username:  "test3",
		Namespace: "ci12345",
		SafeId:    "test",
	}
	code, err := httpReq("POST", "/role", u, nil)
	if err != nil {
		slog.Error(err)
	}
	if code != 200 {
		slog.Error(code)
	}
}

func tc3_createRoleUserHasOneKey() {
	u := model.Role{
		Username:  "iamUser1",
		Namespace: "ci45678-iam-user-1key",
		SafeId:    "test",
	}
	code, err := httpReq("POST", "/role", u, nil)
	if err != nil {
		slog.Error(err)
	}
	if code != 200 {
		slog.Error(code)
	}
}

func tc2_createRoleUserHasAlreadyKeys() {
	u := model.Role{
		Username:  "iamUser2",
		Namespace: "ci898640-iam-user-2keys",
		SafeId:    "test",
	}
	code, err := httpReq("POST", "/role", u, nil)
	if err != nil {
		slog.Error(err)
	}
	if code != 200 {
		slog.Error(code)
	}
}

func tc1_createRoleNsNotFound() {
	u := model.Role{
		Username:  "test",
		Namespace: "test",
		SafeId:    "test",
	}
	code, err := httpReq("POST", "/role", u, nil)
	if err != nil {
		slog.Error(err)
	}
	if code != 200 {
		slog.Error(code)
	}
}

func httpReq(method, path string, data any, obj any) (int, error) {
	path = url + path
	var req *http.Request
	if data != nil {
		payload, _ := json.Marshal(data)
		tmp, err := http.NewRequest(method, path, bytes.NewBuffer(payload))
		if err != nil {
			return 500, err
		}
		req = tmp
	} else {
		tmp, err := http.NewRequest(method, path, nil)
		if err != nil {
			return 500, err
		}
		req = tmp
	}
	resp, err := client.Do(req)
	if err != nil {
		return 500, err
	}
	if err != nil {
		return 500, err
	}

	defer resp.Body.Close()
	bodyByte, err := io.ReadAll(resp.Body)
	slog.Info(string(bodyByte))
	if resp.StatusCode > 300 {
		return resp.StatusCode, errors.New(string(bodyByte))
	}

	if len(bodyByte) > 0 && obj != nil {
		if err = json.Unmarshal(bodyByte, &obj); err != nil {
			return 500, err
		}
	}
	return 200, nil

}

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
	//tc1_onboardNsNotFound()
	//tc2_onboardNsWith1NativeUser()
	//tc4_migrate_not_existing()
	//tc3_migrate_existing()
	//tc5_migrate_existing_2keys()

}

func tc5_migrate_existing_2keys() {
	ns := model.MigrateNamespace{
		Namespace: "ci898640-native-user-iam-user-2keys",
	}
	code, err := httpReq("POST", "/namespace/migrate", ns, nil)
	if err != nil {
		slog.Error(err)
	}
	if code != 200 {
		slog.Error(code)
	}
}

func tc4_migrate_not_existing() {
	ns := model.MigrateNamespace{
		Namespace: "ci45m-user-1key",
	}
	code, err := httpReq("POST", "/namespace/migrate", ns, nil)
	if err != nil {
		slog.Error(err)
	}
	if code != 200 {
		slog.Error(code)
	}
}

func tc3_migrate_existing() {
	ns := model.MigrateNamespace{
		Namespace: "ci45678-native-user-iam-user-1key",
	}
	code, err := httpReq("POST", "/namespace/migrate", ns, nil)
	if err != nil {
		slog.Error(err)
	}
	if code != 200 {
		slog.Error(code)
	}
}

func tc2_onboardNsWith1NativeUser() {
	slog.Info("test onboard namespace with 1 native user")
	ns := model.OnboardNamespace{
		Namespace: "ci12345-native-user",
		Username:  "ns-native-user-us1",
	}
	code, err := httpReq("POST", "/namespace/onboard", ns, nil)
	if err != nil {
		slog.Error(err)
	}
	if code != 200 {
		slog.Error(code)
	}
}

func tc1_onboardNsNotFound() {
	slog.Info("test onboard namespace not found")
	ns := model.OnboardNamespace{
		Namespace: "blah-glo",
		Username:  "blah",
	}
	code, err := httpReq("POST", "/namespace/onboard", ns, nil)
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

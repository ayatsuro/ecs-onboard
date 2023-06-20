package main

import (
	"bytes"
	"ecs-onboard/model"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gookit/slog"
	"io"
	"net/http"
	"time"
)

var (
	client = *http.DefaultClient
	url    = "http://0.0.0.0:8081/v1"
)

func main() {
	//tc1_onboardNsNotFound()
	//tc2_onboardNsWith1NativeUser()
	d1, err := time.Parse(time.RFC3339, "2020-04-15T14:49:22Z")
	if err != nil {
		slog.Error(err)
	}
	d2, err := time.Parse(time.RFC3339, "2020-04-15T15:49:22Z")
	if err != nil {
		slog.Error(err)
	}
	fmt.Println(d1.Before(d2))
}

func tc2_onboardNsWith1NativeUser() {
	slog.Info("test onboard namespace with 1 native user")
	ns := model.Namespace{
		Namespace: "ci12345-native-user-slo",
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
	ns := model.Namespace{
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

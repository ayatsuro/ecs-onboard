package service

import (
	"bytes"
	"ecs-onboard/model"
	"encoding/json"
	"errors"
	"github.com/gookit/slog"
	"io"
	"net/http"
)

const objectStore = "/v1/object-store"

var (
	vault_url = "http://127.0.0.1:8200"
	client    = *http.DefaultClient
)

func OnboardNs(ns model.Namespace) error {

	return nil
}

func ReqVault(method, path string, data any, obj any) (int, error) {
	path = vault_url + objectStore + path
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

	req.Header = http.Header{
		"X-VAULT-TOKEN": {"root"},
		"Content-Type":  {"application/json; charset=UTF-8"},
		"Accept":        {"application/json; charset=UTF-8"},
	}
	resp, err := client.Do(req)
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

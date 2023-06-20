package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gookit/slog"
	"io"
	"net/http"
)

var (
	vault_url = "http://127.0.0.1:8200/v1"
	client    = *http.DefaultClient
)

type VaultResponse struct {
	Data map[string]interface{} `json:"data"`
}

func CreatePolicy(policy string) (int, error) {
	path := "object-store/" + policy
	payload := map[string]string{
		"name":   path,
		"policy": fmt.Sprintf(`{"path": { %q: {"capabilities": ["read"] }}}`, path),
	}
	return ReqVault("POST", "/sys/policies/acl/"+path, payload, nil)
}

func ReqVault(method, path string, data any, obj any) (int, error) {
	path = vault_url + path
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
		var vaultResponse VaultResponse
		if err := json.Unmarshal(bodyByte, &vaultResponse); err != nil {
			return 500, err
		}
		jsonByte, err := json.Marshal(vaultResponse.Data)
		if err != nil {
			return 500, err
		}
		if err := json.Unmarshal(jsonByte, &obj); err != nil {
			return 500, err
		}
	}
	return 200, nil

}

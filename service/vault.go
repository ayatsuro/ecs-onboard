package service

import (
	"bytes"
	"ecs-onboard/model"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var (
	vault_url = "http://127.0.0.1:8200"
	client    = *http.DefaultClient
)

func OnboardNs(ns model.Namespace) error {
	return nil
}

func reqVault(method, path string, data any, obj any) error {
	path = vault_url + path
	var req *http.Request
	if data != nil {
		payload, _ := json.Marshal(data)
		req, _ = http.NewRequest(method, path, bytes.NewBuffer(payload))
	} else {
		req, _ = http.NewRequest(method, path, nil)
	}

	req.Header.Add("X-SDS-AUTH-TOKEN", e.token)
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")
	req.Header.Add("Accept", "application/json; charset=UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	// if token has expired, we login again
	if resp.StatusCode == 401 {
		if err := e.login(); err != nil {
			return err
		}
		req.Header.Set("X-SDS-AUTH-TOKEN", e.token)
		resp, err = e.client.Do(req)
		if err != nil {
			return err
		}
	}

	defer resp.Body.Close()
	bodyByte, err := io.ReadAll(resp.Body)
	if resp.StatusCode > 300 {
		return errors.New(resp.Status + " " + string(bodyByte))
	}

	if len(bodyByte) > 0 && obj != nil {
		if err = json.NewDecoder(bytes.NewReader(bodyByte)).Decode(&obj); err != nil {
			return err
		}
	}
	return nil

}

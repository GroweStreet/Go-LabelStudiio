package annotations

import (
	"Lib/config"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func Create(id int, token string, request CreateRequest) (bool, error) {

	r, err := json.Marshal(request)
	if err != nil {
		return false, err
	}

	config.Init(host, port)
	url := fmt.Sprintf("http://%s:%s/api/tasks/%d/annotations/", config.Host(), config.Port(), id)
	payload := bytes.NewReader(r)

	req, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		return false, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Token %s", token))
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}

	defer res.Body.Close()
	return res.StatusCode == http.StatusCreated, nil
}

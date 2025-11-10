package annotations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/GroweStreet/Go-LabelStudiio/config"
	"net/http"
)

func Create(taskId int, request CreateRequest) (bool, error) {

	r, err := json.Marshal(request)
	if err != nil {
		return false, err
	}

	url := fmt.Sprintf("http://%s:%s/api/tasks/%d/annotations/", config.Host(), config.Port(), taskId)
	payload := bytes.NewReader(r)

	req, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		return false, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", config.Token()))
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}

	defer res.Body.Close()
	return res.StatusCode == http.StatusCreated, nil
}

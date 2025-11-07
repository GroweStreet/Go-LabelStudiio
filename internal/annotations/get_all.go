package annotations

import (
	"Lib/config"
	"encoding/json"
	"fmt"
	"net/http"
)

var annotations []Annotation

func FetchAll(id int, token string) ([]Annotation, error) {
	config.Init(host, port)
	url := fmt.Sprintf("http://%s:%s/api/tasks/%d/annotations/", config.Host(), config.Port(), id)
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Token %s", token))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&annotations)
	if err != nil {
		return annotations, err
	}
	return annotations, nil

}

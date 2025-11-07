package annotations

import (
	"Lib/config"
	"fmt"
	"net/http"
)

func Delete(id int, token string) (bool, error) {

	config.Init(host, port)
	url := fmt.Sprintf("http://%s:%s/api/annotations/%d/", config.Host(), config.Port(), id)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return false, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Token %s", token))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	return res.StatusCode == http.StatusNoContent, nil
}

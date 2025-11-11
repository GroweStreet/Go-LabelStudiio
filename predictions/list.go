package predictions

import (
	"fmt"
	config "github.com/GroweStreet/Go-LabelStudiio/config"
	_ "io"
	"net/http"
)

func List() (bool, error) {

	url := fmt.Sprintf("http://%s:%s/api/predictions/", config.Host(), config.Port())

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", config.Token()))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}

	err = res.Body.Close()
	if err != nil {
		return false, err
	}
	return res.StatusCode == http.StatusOK, nil

}

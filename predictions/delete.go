package predictions

import (
	"fmt"
	config "github.com/GroweStreet/Go-LabelStudiio/config"
	"net/http"
)

func Delete(id int) (bool, error) {

	url := fmt.Sprintf("http://%s:%s/api/predictions/%d/", config.Host(), config.Port(), id)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return false, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Token %s", config.Token()))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}

	defer res.Body.Close()

	return res.StatusCode == 200, nil
}

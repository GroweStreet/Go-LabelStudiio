package predictions

import (
	"bytes"
	"encoding/json"
	"fmt"
	config "github.com/GroweStreet/Go-LabelStudiio/config"
	"net/http"
)

func Create(prediction Prediction) (bool, error) {

	url := fmt.Sprintf("http://%s:%s/api/predictions/", config.Host(), config.Port())

	r, err := json.Marshal(prediction)
	if err != nil {
		return false, err
	}

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
	//body, err := io.ReadAll(res.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(res)
	//fmt.Println(string(body))
}

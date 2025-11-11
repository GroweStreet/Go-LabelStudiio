package predictions

import (
	"bytes"
	"encoding/json"
	"fmt"
	config "github.com/GroweStreet/Go-LabelStudiio/config"
	"net/http"
)

func Update(predictionId int, prediction Prediction) (bool, error) {

	r, err := json.Marshal(prediction)
	url := fmt.Sprintf("http://%s:%s/api/predictions/%d/", config.Host(), config.Port(), predictionId)

	payload := bytes.NewReader(r)
	req, err := http.NewRequest(http.MethodPatch, url, payload)
	if err != nil {
		return false, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", config.Token()))
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	return res.StatusCode == http.StatusOK, nil
	//body, _ := io.ReadAll(res.Body)
	//fmt.Println(res)
	//fmt.Println(string(body))
}

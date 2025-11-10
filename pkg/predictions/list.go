package predictions

import (
	"encoding/json"
	"fmt"
	config "github.com/GroweStreet/Go-LabelStudiio/config"
	_ "io"
	"log"
	"net/http"
)

func List() ([]Prediction, error) {

	var prediction []Prediction

	url := fmt.Sprintf("http://%s:%s/api/predictions/", config.Host(), config.Port())

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Authorization", "Token  6a2e95d769a7cdf02097918de4f2574df0804d7c")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	//body, _ := io.ReadAll(res.Body)
	//fmt.Println(res)
	//fmt.Println(string(body))
	err = json.NewDecoder(res.Body).Decode(&prediction)
	if err != nil {
		log.Fatal(err)
	}
	return prediction, nil
}

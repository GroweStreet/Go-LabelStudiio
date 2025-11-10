package predictions

import (
	"Lib/config"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Create(prediction Prediction) {

	url := fmt.Sprintf("http://%s:%s/api/predictions/", config.Host(), config.Port())

	r, err := json.Marshal(prediction)
	if err != nil {
		panic(err)
	}

	payload := bytes.NewReader(r)

	req, _ := http.NewRequest(http.MethodPost, url, payload)
	req.Header.Add("Authorization", "Token  6a2e95d769a7cdf02097918de4f2574df0804d7c")
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	//body, err := io.ReadAll(res.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(res)
	//fmt.Println(string(body))
}

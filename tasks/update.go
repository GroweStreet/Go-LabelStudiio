package tasks

import (
	"fmt"
	"github.com/GroweStreet/Go-LabelStudiio/config"
	"log"
	"net/http"
	"strings"
)

func Update(taskId int) {

	url := fmt.Sprintf("http://%s:%s/api/tasks/%d/", config.Host(), config.Port(), taskId)
	payload := strings.NewReader("{}")
	req, err := http.NewRequest(http.MethodPatch, url, payload)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Token %s", config.Token()))
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
}

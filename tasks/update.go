package tasks

import (
	"encoding/json"
	"fmt"
	"github.com/GroweStreet/Go-LabelStudiio/config"
	"io"
	"log"
	"net/http"
	"strings"
)

func Update(id int) {
	var task Task
	url := fmt.Sprintf("http://%s:%s/api/tasks/%d/", config.Host(), config.Port(), id)
	payload := strings.NewReader("{}")
	req, _ := http.NewRequest(http.MethodPatch, url, payload)

	req.Header.Set("Authorization", fmt.Sprintf("Token %s", config.Token()))
	req.Header.Set("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	err := json.Unmarshal(body, &task)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(res)
	fmt.Println(task)
}

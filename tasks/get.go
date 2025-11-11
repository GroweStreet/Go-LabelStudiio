package tasks

import (
	"encoding/json"
	"fmt"
	"github.com/GroweStreet/Go-LabelStudiio/config"
	"net/http"
)

func GetTask(taskId int) (Task, error) {

	var task Task

	url := fmt.Sprintf("http://%s:%s/api/tasks/%d/", config.Host(), config.Port(), taskId)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return task, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Token %s", config.Token()))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return task, err
	}

	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&task)
	if err != nil {
		return task, err
	}
	return task, nil
}

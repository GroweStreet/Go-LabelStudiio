package tasks

import (
	"github.com/GroweStreet/Go-LabelStudiio/config"
	"encoding/json"
	"fmt"
	"net/http"
)

var task Task

func GetTask(id int) (Task, error) {

	url := fmt.Sprintf("http://%s:%s/api/tasks/%d/", config.Host(), config.Port(), id)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return task, err
	}

	req.Header.Add("Authorization", "Token  6a2e95d769a7cdf02097918de4f2574df0804d7c")
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

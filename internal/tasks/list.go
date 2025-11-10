package tasks

import (
	"encoding/json"
	"fmt"
	"github.com/GroweStreet/Go-LabelStudiio/config"
	"net/http"
)

type Tasks struct {
	TotalAnnotations int    `json:"total_annotations"`
	TotalPredictions int    `json:"total_predictions"`
	Total            int    `json:"total"`
	Tasks            []Task `json:"tasks"`
}

func List(projectId int, token string) (Tasks, error) {

	var tasks Tasks

	url := fmt.Sprintf("http://%s:%s/api/tasks/?project=%d", config.Host(), config.Port(), projectId)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Token %s", token))
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	err := json.NewDecoder(res.Body).Decode(&tasks)
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

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

func List(projectId int, startIx, endIx int) (Tasks, error) {

	var tasks Tasks
	var tempTasks Tasks

	tempTasks.Tasks = make([]Task, 0, 1000)
	url := fmt.Sprintf("http://%s:%s/api/tasks/?project=%d", config.Host(), config.Port(), projectId)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return tasks, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Token %s", config.Token()))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return tempTasks, err
	}
	err = json.NewDecoder(res.Body).Decode(&tempTasks)
	if err != nil {
		return tempTasks, err
	}
	for _, task := range tempTasks.Tasks {
		if task.InnerId >= startIx && task.InnerId <= endIx {
			tasks.Tasks = append(tasks.Tasks, task)
		}
		if task.InnerId == endIx {
			break
		}
	}
	tempTasks.Tasks = nil

	return tasks, nil
}

package tasks

import (
	"encoding/json"
	"fmt"
	"github.com/GroweStreet/Go-LabelStudiio/config"
	"io"
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

	for pageNumber := 0; ; pageNumber++ {

		url := fmt.Sprintf("http://%s:%s/api/tasks/?project=%d&page=%d&page_size=1000", config.Host(), config.Port(), projectId, pageNumber)

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return tasks, err
		}

		req.Header.Set("Authorization", fmt.Sprintf("Token %s", config.Token()))
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return tasks, err
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return tasks, err
		}

		err = json.Unmarshal(body, &tempTasks)
		if err != nil {
			return tempTasks, err
		}

		for _, task := range tempTasks.Tasks {
			if task.InnerId >= startIx && task.InnerId <= endIx {
				tasks.Tasks = append(tasks.Tasks, task)
			}
			if task.InnerId >= endIx {
				return tasks, nil
			}
		}

		tempTasks.Tasks = nil
	}
}

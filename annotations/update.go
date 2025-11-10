package annotations

import (
	"bytes"
	"encoding/json"
	"fmt"
	config "github.com/GroweStreet/Go-LabelStudiio/config"
	"log"
	"net/http"
)

func Update(annotationId int, token string, annotation Annotation) (bool, error) {

	r, err := json.Marshal(annotation)
	if err != nil {
		return false, err
	}

	url := fmt.Sprintf("http://%s:%s/api/annotations/%d/", config.Host(), config.Port(), annotationId)
	payload := bytes.NewReader(r)
	log.Println()

	req, _ := http.NewRequest(http.MethodPatch, url, payload)
	req.Header.Set("Authorization", fmt.Sprintf("Token  %s", token))
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	return res.StatusCode == http.StatusOK, nil
}

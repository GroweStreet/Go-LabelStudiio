package predictions

import (
	"fmt"
	config "github.com/GroweStreet/Go-LabelStudiio/config"
	"io"
	"net/http"
	"strings"
)

func Update(id int) {

	url := fmt.Sprintf("http://%s:%s/api/predictions/%d/", config.Host(), config.Port(), id)

	payload := strings.NewReader("{\n  \"model_version\": \"yolo-v8\",\n  \"result\": [\n    {\n      \"from_name\": \"bboxes\",\n      \"image_rotation\": 0,\n      \"original_height\": 1080,\n      \"original_width\": 1920,\n      \"to_name\": \"image\",\n      \"type\": \"rectanglelabels\",\n      \"value\": {\n        \"height\": 60,\n        \"rotation\": 0,\n        \"values\": {\n          \"rectanglelabels\": [\n            \"Person\"\n          ]\n        },\n        \"width\": 50,\n        \"x\": 20,\n        \"y\": 30\n      }\n    }\n  ],\n  \"score\": 0.95\n}")

	req, _ := http.NewRequest(http.MethodPatch, url, payload)
	req.Header.Add("Authorization", "Token 6a2e95d769a7cdf02097918de4f2574df0804d7c")
	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	fmt.Println(res)
	fmt.Println(string(body))
}

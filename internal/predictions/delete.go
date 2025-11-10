package predictions

import (
	"fmt"
	config "github.com/GroweStreet/Go-LabelStudiio/config"
	"io"
	"net/http"
)

func Delete(id int) {

	url := fmt.Sprintf("http://%s:%s/api/predictions/%d/", config.Host(), config.Port(), id)

	req, _ := http.NewRequest(http.MethodDelete, url, nil)
	req.Header.Add("Authorization", "Token  6a2e95d769a7cdf02097918de4f2574df0804d7c")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	fmt.Println(res)
	fmt.Println(string(body))
}

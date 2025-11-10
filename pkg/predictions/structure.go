package predictions

type Prediction struct {
	Id           int      `json:"id"`
	Result       []Result `json:"result"`
	ModelVersion string   `json:"model_version"`
	CreatedAgo   string   `json:"created_ago"`
	Score        any      `json:"score"`
	Cluster      any      `json:"cluster"`
	Neighbors    any      `json:"neighbors"`
	Mislabeling  float64  `json:"mislabeling"`
	CreatedAt    string   `json:"created_at"`
	UpdatedAt    string   `json:"updated_at"`
	Model        any      `json:"model"`
	ModelRun     any      `json:"model_run"`
	Task         int      `json:"task"`
	Project      int      `json:"project"`
}

type Result struct {
	Id       string `json:"id"`
	Type     string `json:"type"`
	Value    Value  `json:"value"`
	Origin   string `json:"origin"`
	ToName   string `json:"to_name"`
	FromName string `json:"from_name"`
}

type Value struct {
	End     int      `json:"end"`
	Text    string   `json:"text"`
	Start   int      `json:"start"`
	Labels  []string `json:"labels,omitempty"`
	Choices []string `json:"choices,omitempty"`
}

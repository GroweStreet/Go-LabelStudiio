package annotations

type Annotation struct {
	Id     int      `json:"id,omitempty"`
	Result []Result `json:"result"`
}

type Result struct {
	Id       string `json:"id"`
	Type     string `json:"type"`
	Value    Value  `json:"value"`
	Origin   string `json:"origin"`
	ToName   string `json:"to_name"`
	FromName string `json:"from_name"`
	RegionId string `json:"region_id,omitempty"`
}

type Value struct {
	End     int      `json:"end,omitempty"`
	Text    string   `json:"text,omitempty"`
	Start   int      `json:"start,omitempty"`
	Labels  []string `json:"labels,omitempty"`
	Choices []string `json:"choices,omitempty"`
}

type CreateRequest struct {
	Result []Result `json:"result"`
}

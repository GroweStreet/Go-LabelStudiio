package annotations

type Annotation struct {
	Id     int      `json:"id"`
	Result []Result `json:"result"`
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
	End    int      `json:"end"`
	Text   string   `json:"text"`
	Start  int      `json:"start"`
	Labels []string `json:"labels"`
}

type CreateRequest struct {
	Result []Annotation `json:"result"`
}

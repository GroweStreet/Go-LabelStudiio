package tasks

import (
	"encoding/json"
	"github.com/GroweStreet/Go-LabelStudiio/pkg/annotations"
	"time"
)

type Task struct {
	Id                   int                      `json:"id"`
	Data                 json.RawMessage          `json:"data"`
	Predictions          []interface{}            `json:"predictions"`
	Annotations          []annotations.Annotation `json:"annotations"`
	InnerId              int                      `json:"inner_id"`
	CancelledAnnotations int                      `json:"cancelled_annotations"`
	TotalAnnotations     int                      `json:"total_annotations"`
	CompletedAt          time.Time                `json:"completed_at"`
	AnnotionsResults     string                   `json:"annotions_results"`
	PredictionsResults   string                   `json:"predictions_results"`
	DraftExist           bool                     `json:"draft_exist"`
	UpdatedAt            time.Time                `json:"updated_at"`
	IsLabeled            bool                     `json:"is_labeled"`
	fullText             string
}

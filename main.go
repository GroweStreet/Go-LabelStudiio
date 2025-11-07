package main

import (
	ann "Lib/internal/annotations"
	"log"
)

var err error

func main() {

	//pred := predictions.Prediction{
	//	Id: 0,
	//	Result: []predictions.Result{{
	//		Id:   "",
	//		Type: "",
	//		Value: predictions.Value{
	//			End:     0,
	//			Text:    "подписка сбер прайм",
	//			Start:   0,
	//			Labels:  "",
	//			Choices: nil,
	//		},
	//		Origin:   "",
	//		ToName:   "",
	//		FromName: "",
	//	}},
	//	ModelVersion: "",
	//	CreatedAgo:   "",
	//	Score:        nil,
	//	Cluster:      nil,
	//	Neighbors:    nil,
	//	Mislabeling:  0,
	//	CreatedAt:    "",
	//	UpdatedAt:    "",
	//	Model:        nil,
	//	ModelRun:     nil,
	//	Task:         67714,
	//	Project:      1,
	//}

	//predictions.Create(pred)
	//prd, _ := predictions.List()
	//fmt.Println(prd)

	//tasks, err := tasks.GetTask(67721)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(tasks)

	b, err := ann.Delete(135422, "6a2e95d769a7cdf02097918de4f2574df0804d7c")
	if err != nil {
		log.Fatal(err)
	}
	println(b)
	//b := ann.Create(67717)
	//fmt.Println(b)

	//tasks.GetTask(67721)
	//task, _ := tasks.List(1, "6a2e95d769a7cdf02097918de4f2574df0804d7c")
	//fmt.Println(task)
	//annotations, err = ann.GetAll(67721, "6a2e95d769a7cdf02097918de4f2574df0804d7c")
	//fmt.Println(annotations)
}

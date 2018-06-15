package main

import (
	"fmt"

	cl "github.com/mpmlj/clarifai-client-go"
)

const (
	apiKey = "d6d623dad6894b54b6a5feb980edc339"
)

func main() {
	image := "https://c3.staticflickr.com/4/3372/3557249560_fa83b3c878_o.jpg"
	resp := PredictImage(image)
	cl.PP(resp)
}

func PredictImage(image string) *cl.Response {
	var err error
	var sess *cl.Session
	sess = cl.NewApp(apiKey)
	fmt.Println(sess)
	data := cl.InitInputs()
	//	_ = data.AddInput(cl.NewImageFromURL("https://farm3.staticflickr.com/568/21452126474_ab12789b36_o.jpg"), "")
	_ = data.AddInput(cl.NewImageFromURL(image), "")
	data.SetModel(cl.PublicModelTravel)
	resp, err := sess.Predict(data).Do()
	if err != nil {
		panic(err)
	}
	return resp
}

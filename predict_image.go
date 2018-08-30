package main

import (
	"bufio"
	"fmt"
	cl "github.com/mpmlj/clarifai-client-go"
	"log"
	"os"
	"sort"
)

const (
	apiKey     = "d6d623dad6894b54b6a5feb980edc339"
	imagesFile = "images.txt"
	zero       = 0
)

type imageValue struct {
	image string
	value float64
}

func main() {

	mp := make(map[string][]imageValue)
	file, err := os.Open(imagesFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		image := scanner.Text()
		fmt.Println(image)
		//send POST request and unmarshal that data in some struct
		resp := PredictImage(image)

		output := resp.Outputs[zero].Data.Concepts
		for _, element := range output {
			iv := imageValue{image: image, value: element.Value}
			mp[element.Name] = insertSort(mp[element.Name], iv)
		}
	}

	displayMap(mp)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func displayMap(mp map[string][]imageValue) {
	for key, value := range mp {
		fmt.Println(key)
		fmt.Println(value)
		fmt.Println()
	}
}

func insertSort(data []imageValue, el imageValue) []imageValue {
	index := sort.Search(len(data), func(i int) bool { return data[i].value < el.value })
	data = append(data, imageValue{})
	copy(data[index+1:], data[index:])
	data[index] = el
	return data
}

func PredictImage(image string) *cl.Response {
	var err error
	var sess *cl.Session
	sess = cl.NewApp(apiKey)
	data := cl.InitInputs()
	_ = data.AddInput(cl.NewImageFromURL(image), "")
	data.SetModel(cl.PublicModelTravel)
	resp, err := sess.Predict(data).Do()
	if err != nil {
		panic(err)
	}
	return resp
}

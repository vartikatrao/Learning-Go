package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Hello")
	EncodeJson()
	DecodeJson()

}

func EncodeJson() {
	lcoCourses := []course{
		{"reactjs", 299, "facebook", "password123", []string{"web-dev", "js"}},
		{"golang", 199, "golang.org", "password456", []string{"web-dev", "go"}},
		{"python", 299, "python.org", "password789", nil},
	}

	//package encoding/json

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalJson))

}

func DecodeJson() {
	jsonData := []byte(`
	{
		"coursename": "reactjs",
		"Price": 299,
		"website": "facebook",
		"tags": ["web-dev","js"]
	}
`)

	var lcoCourse course
	checkValid := json.Valid(jsonData)
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonData, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	}else{
		fmt.Println("JSON was not valid")
	}

	// some cases where you just want to add data to a key value pair

	var myData map[string]interface{}
	json.Unmarshal(jsonData, &myData)
	fmt.Printf("%#v", myData)
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Label - Request body of a lable
type Label struct {
	Name        string `json:"name"`
	Color       string `json:"color"`
	Description string `json:"description"`
}

// Labels - Array of Labels
type Labels struct {
	Labels []Label `json:"labels"`
}

func main() {
	labelsFile, err := os.Open("labels.json")
	defer labelsFile.Close()
	if err != nil {
		fmt.Println("Coludn't open file")
		os.Exit(1)
	}

	byteValue, _ := ioutil.ReadAll(labelsFile)

	var labels Labels

	json.Unmarshal(byteValue, &labels)

	url := "https://api.github.com/repos/REPO_OWNER/REPO_NAME/labels"

	for i := 0; i < len(labels.Labels); i++ {
		log.Printf("Lable name: " + labels.Labels[i].Name)
		log.Printf("Lable color: " + labels.Labels[i].Color)
		b, _ := json.Marshal(labels.Labels[i])
		var lableToSend = []byte(b)
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(lableToSend))
		req.Header.Set("Authorization", "token YOUR_GITHUB_TOKEN")
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		fmt.Println("Status code: " + string(resp.Status))
		defer resp.Body.Close()
	}
}

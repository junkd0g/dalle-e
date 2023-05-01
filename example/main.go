package main

import (
	"fmt"

	"github.com/junkd0g/dalle-e"
)

const (
	apiKey = "your-api-key"
	domain = "https://api.openai.com"
)

func main() {
	client, _ := dalle.NewClient(apiKey, domain)

	// Generate an image from a prompt.
	prompt := "A painting of a cat sitting on a couch."
	n := 1
	size := "1024x1024"

	response, err := client.GenerateImageV1(prompt, n, size)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(response))
}

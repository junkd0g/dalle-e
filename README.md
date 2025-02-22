[![Go Report Card](https://goreportcard.com/badge/github.com/junkd0g/dalle-e)](https://goreportcard.com/report/github.com/junkd0g/dalle-e)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://pkg.go.dev/badge/github.com/junkd0g/dalle-e.svg)](https://pkg.go.dev/github.com/junkd0g/dalle-e)
# Dalle API Client

## Overview

This dependency provides a simple Go client for interacting with the Dalle API. It allows you to generate images from a prompt by sending HTTP requests to the image generation endpoint. The client is designed to be configurable with options for setting a custom HTTP client and request timeout.

## Features

- Create a new client instance using your API key.
- Optionally specify a custom domain (defaults to `https://api.openai.com`).
- Configure HTTP client settings such as timeout or a custom HTTP client.
- Generate images by providing a prompt, the number of images to generate, and the desired image size.
- Handles request marshalling and HTTP POST operations internally.

## Installation

1. Include the package in your Go project by importing it (for example, `yourmodule/dalle`).
2. Ensure you have a valid API key from the image generation service.

## Usage Example

```go
package main

import (
    "fmt"
    "time"
    "github.com/junkd0g/dalle-e/dalle"
)

func main() {
    // Create a new Dalle client with your API key.
    client, err := dalle.NewClient("your_api_key", "")
    if err != nil {
        fmt.Println("Error creating client:", err)
        return
    }

    // Optionally, set a custom timeout.
    client, err = dalle.NewClient("your_api_key", "", dalle.WithTimeout(30*time.Second))
    if err != nil {
        fmt.Println("Error creating client:", err)
        return
    }

    // Generate an image using a prompt.
    response, err := client.GenerateImageV1("A futuristic city skyline", 1, "1024x1024")
    if err != nil {
        fmt.Println("Error generating image:", err)
        return
    }

    // Process the response (response is a byte slice).
    fmt.Println("Response:", string(response))
}
```

## Configuration Options

- **APIKey (string)**: Your API authentication key. This is required.
- **Domain (string)**: The API domain to use. If left empty, the default value `https://api.openai.com` is used.
- **WithHTTPClient**: An option to provide a custom HTTP client.
- **WithTimeout**: An option to set a specific timeout for API requests.

## Error Handling

- An error is returned if the API key is not provided during client creation.
- Errors during JSON marshalling or the HTTP request are propagated back to the caller.

## Additional Notes

- The `GenerateImageV1` function constructs the request using the provided prompt, number of images (`n`), and size.
- The client’s `post` method constructs the URL using the domain, API version (`/v1`), and endpoint (`/images/generations`).
- Response data is returned as a byte slice. You may need to further process or decode this data according to your application’s requirements.
- The printed output of the API key in the `post` function is for debugging purposes; consider removing or securing it in production.

## 📝 License

This project is licensed under the MIT License. See the LICENSE file for details.

## Authors

* **Iordanis Paschalidis** -[junkd0g](https://github.com/junkd0g)
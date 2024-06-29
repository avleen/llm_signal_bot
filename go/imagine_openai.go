package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

func imagineOpenai(prompt string, requestor string) (string, string, error) {
	client := openai.NewClient(config["OPENAI_API_KEY"])

	err := makeOutputDir(config["IMAGEDIR"])
	if err != nil {
		fmt.Println("Failed to create output directory:", err)
		return "", "", err

	}

	// Generate an image from the text and send it
	jsonResp, err := client.CreateImage(context.Background(),
		openai.ImageRequest{
			Prompt:         prompt,
			Model:          openai.CreateImageModelDallE3,
			ResponseFormat: openai.CreateImageResponseFormatB64JSON,
			N:              1,
			Size:           openai.CreateImageSize1024x1024,
			Quality:        openai.CreateImageQualityStandard,
		},
	)
	if err != nil {
		fmt.Println("Failed to generate image:", err)
		return "", "", err
	}
	revisedPrompt := jsonResp.Data[0].RevisedPrompt
	data := jsonResp.Data[0].B64JSON

	// Decode the base64 image data
	imageData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		fmt.Println("Failed to decode image data:", err)
		return "", "", err
	}

	// Save the imageData to a file in the format:
	// <date>-<time>-<requestor>.png
	filename := fmt.Sprintf("%s-%s-%s.png", time.Now().Format("2006-01-02"), time.Now().Format("15:04:05"), requestor)
	filename = filepath.Join(config["IMAGEDIR"], filename)
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return "", "", err
	}
	defer file.Close()
	// Write the image data to the file
	_, err = file.Write(imageData)
	if err != nil {
		fmt.Println("Failed to write image data to file:", err)
		return "", "", err
	}
	fmt.Printf("Image saved to %s with revised prompt: %s\n", filename, revisedPrompt)

	return filename, revisedPrompt, nil
}

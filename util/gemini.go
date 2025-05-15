package util

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"google.golang.org/genai"
)

var (
	prompFlag string
)

func GeneratePrompt() (string, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GOOGLE_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return "", err
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text("Genera un prompt visual artístico y detallado para una imagen IA."),
		nil,
	)
	if err != nil {
		return "", err
	}

	return result.Text(), nil
}

func GenerateImage(prompt string) ([]byte, error) {
	flag.StringVar(&prompFlag, "promp", prompt, "Use promp to generate images")
	flag.Parse()
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	var imagesPath string = path.Join(currentPath, "images")
	if _, err := os.Stat(imagesPath); os.IsNotExist(err) {
		err := os.Mkdir(imagesPath, os.ModeDir)
		if err != nil {
			log.Fatal(err)
		}
	}
	ctx := context.Background()
	client, _ := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	parts := []*genai.Part{
		genai.NewPartFromText(prompFlag),
	}
	contents := []*genai.Content{
		genai.NewContentFromParts(parts, genai.RoleUser),
	}

	config := &genai.GenerateContentConfig{
		ResponseModalities: []string{"TEXT", "IMAGE"},
	}

	result, _ := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash-exp-image-generation",
		contents,
		config,
	)

	var image []byte
	var noImage error
	for _, part := range result.Candidates[0].Content.Parts {
		if part.Text != "" {
			fmt.Println(part.Text)
			image = nil
			noImage = fmt.Errorf("no image generated")
		} else if part.InlineData != nil {
			imageBytes := part.InlineData.Data
			t := time.Now()
			outputFilename := fmt.Sprintf("%d%02d%02d_%02d%02d%02d.png", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
			_ = os.WriteFile(path.Join(imagesPath, outputFilename), imageBytes, 0644)
			image = imageBytes
			noImage = nil
		}
	}
	return image, noImage
}

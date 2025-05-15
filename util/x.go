package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"

	"github.com/dghubble/oauth1"
)

func uploadMedia(client *http.Client, imageBytes []byte) (string, error) {
	url := "https://upload.x.com/1.1/media/upload.json"

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("media", "image.png")
	if err != nil {
		return "", err
	}
	part.Write(imageBytes)
	writer.Close()

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var res struct {
		MediaIDString string `json:"media_id_string"`
	}
	err = json.Unmarshal(respBody, &res)
	if err != nil {
		return "", err
	}

	return res.MediaIDString, nil
}

func PublishPost(prompt string, img []byte) {
	// Limpia saltos de línea
	re := regexp.MustCompile(`\r?\n`)
	prompt = re.ReplaceAllString(prompt, " ")

	if len(prompt) > 280 {
		prompt = prompt[:280]
	}

	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	tokenSecret := os.Getenv("TOKEN_SECRET")

	if consumerKey == "" || consumerSecret == "" || accessToken == "" || tokenSecret == "" {
		fmt.Println("Por favor define las variables CONSUMER_KEY, CONSUMER_SECRET, ACCESS_TOKEN y TOKEN_SECRET")
		return
	}

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, tokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	url := "https://api.x.com/2/tweets"

	mediaID, err := uploadMedia(httpClient, img)
	if err != nil {
		fmt.Println("Error subiendo media:", err)
		return
	}

	tweet := map[string]interface{}{
		"text": prompt,
		"media": map[string]interface{}{
			"media_ids": []string{mediaID},
		},
	}

	jsonBody, err := json.Marshal(tweet)
	if err != nil {
		fmt.Println("Error codificando JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Error creando request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("Error enviando request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error leyendo respuesta:", err)
		return
	}

	fmt.Printf("Código de estado: %d\n", resp.StatusCode)
	fmt.Println("Respuesta:")
	fmt.Println(string(body))
}

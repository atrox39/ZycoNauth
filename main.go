package main

import (
	"fmt"

	"github.com/atrox39/zyconauth/util"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No se pudo cargar archivo .env")
	}

	prompt, err := util.GeneratePrompt()
	if err != nil {
		fmt.Println("Error generando prompt:", err)
		return
	}

	img, err := util.GenerateImage(prompt)
	if err != nil {
		fmt.Println("Error generando imagen:", err)
		return
	}

	util.PublishPost(prompt, img)
}

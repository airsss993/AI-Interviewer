package ai

import (
	"AI_Interviewer/internal/models"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetAIReply() string {
	postURL := "http://127.0.0.1:8000/chat"

	body := models.RequestAI{Text: "Привет, расскажи о себе."}

	bodyBytes, _ := json.Marshal(&body)

	resp, err := http.Post(postURL, "application/json", bytes.NewReader(bodyBytes))
	if err != nil {
		log.Printf("request error: %s", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Print(err)
		}
	}(resp.Body)

	var respAI models.ResponseAI

	if err := json.NewDecoder(resp.Body).Decode(&respAI); err != nil {
		log.Printf("decode error: %s", err)
		return ""
	}

	log.Printf("AI response: %v", respAI.Response)
	return string(respAI.Response)
}

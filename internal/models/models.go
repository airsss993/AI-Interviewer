package models

type RequestAI struct {
	Text string `json:"text"`
}

type ResponseAI struct {
	Response string `json:"response"`
}

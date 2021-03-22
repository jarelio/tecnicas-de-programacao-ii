package model

type Grade struct {
	ID      string `json:"id"`
	Subject string `json:"subject"`
	Type    string `json:"type"`
	Value   string `json:"value"`
	Student string `json:"student"`
}

type InvalidGrade struct {
	ID      string            `json:"id"`
	Subject map[string]string `json:"subject"`
	Type    string            `json:"type"`
	Value   string            `json:"value"`
	Student string            `json:"student"`
}

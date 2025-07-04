package summarizer

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/deevanshu-k/fealtyx-student-api/src/config"
	"github.com/deevanshu-k/fealtyx-student-api/src/structs"
)

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type OllamaResponse struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Response  string `json:"response"`
}

func SummarizeStudent(student structs.Student) (string, error) {
	ollamaRequest := OllamaRequest{
		Model: config.OLLAMA_MODEL,
		Prompt: fmt.Sprintf("%s Student Information: Name: %s, Age: %d, Email: %s",
			config.SUMMARIZE_STUDENT_PROMPT,
			student.Name, student.Age,
			student.Email,
		),
		Stream: false,
	}

	requestBody, err := json.Marshal(ollamaRequest)
	if err != nil {
		return "", errors.New("Error marshalling ollama request")
	}

	response, err := http.Post(config.OLLAMA_GENERATE_URL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", errors.New("Failed to connect to Ollama")
	}
	defer response.Body.Close()

	responseBodey, err := io.ReadAll(response.Body)
	if err != nil {
		return "", errors.New("Failed to read ollama response body")
	}

	var ollamaResponse OllamaResponse
	err = json.Unmarshal(responseBodey, &ollamaResponse)
	if err != nil {
		return "", errors.New("Failed to unmarshal ollama response")
	}

	return ollamaResponse.Response, nil
}

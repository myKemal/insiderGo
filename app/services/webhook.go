package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/myKemal/insiderGo/app/config"
	"github.com/myKemal/insiderGo/app/model"
	"io/ioutil"
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type WebhookService struct {
	WebhookURL string
	Client     HTTPClient
}

func NewWebhookService() *WebhookService {
	return &WebhookService{
		WebhookURL: config.GetWebHookURL(),
		Client:     &http.Client{},
	}
}

func (s *WebhookService) SendPost(payload model.WebHookPayload) (*model.WebhookResponse, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequest("POST", s.WebhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusAccepted {
		return nil, fmt.Errorf("non-successful status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var webhookResponse model.WebhookResponse
	if err := json.Unmarshal(body, &webhookResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return &webhookResponse, nil
}

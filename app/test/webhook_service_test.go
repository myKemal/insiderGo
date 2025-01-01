package test

import (
	"bytes"
	"github.com/myKemal/insiderGo/app/mock"
	"github.com/myKemal/insiderGo/app/model"
	"github.com/myKemal/insiderGo/app/services"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestWebhookService_SendPost_Success(t *testing.T) {
	mockResponseBody := `{"message": "Accepted", "messageId": "12345"}`
	mockClient := &mock.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(mockResponseBody)),
			}, nil
		},
	}

	service := &services.WebhookService{
		WebhookURL: "http://example.com/webhook",
		Client:     mockClient,
	}

	payload := model.WebHookPayload{
		To:      "+1234567890",
		Content: "Hello, World!",
	}

	response, err := service.SendPost(payload)

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "Accepted", response.Message)
	assert.Equal(t, "12345", response.MessageID)
}

func TestWebhookService_SendPost_Failure(t *testing.T) {
	mockClient := &mock.MockHTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return nil, http.ErrHandlerTimeout
		},
	}

	service := &services.WebhookService{
		WebhookURL: "http://example.com/webhook",
		Client:     mockClient,
	}

	payload := model.WebHookPayload{
		To:      "+1234567890",
		Content: "Hello, World!",
	}

	response, err := service.SendPost(payload)

	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Contains(t, err.Error(), "request failed")
}

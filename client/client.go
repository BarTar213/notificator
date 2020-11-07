package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/BarTar213/notificator/models"
	"github.com/BarTar213/notificator/senders"
)

type Client interface {
	SendInternal(ctx context.Context, templateName string, internal *senders.Internal) (int, *models.Response, error)
	SendEmail(ctx context.Context, templateName string, email *senders.Email) (int, *models.Response, error)
}

type Notificator struct {
	client  *http.Client
	baseURL string
}

func New(address string, timeout time.Duration) Client {
	return &Notificator{
		client: &http.Client{
			Timeout: timeout,
		},
		baseURL: address,
	}
}

func (n *Notificator) SendInternal(ctx context.Context, templateName string, internal *senders.Internal) (int, *models.Response, error) {
	return n.send(ctx, templateName, internal, "internal")
}

func (n *Notificator) SendEmail(ctx context.Context, templateName string, email *senders.Email) (int, *models.Response, error) {
	return n.send(ctx, templateName, email, "email")
}

func (n *Notificator) send(ctx context.Context, templateName string, internal interface{}, senderType string) (int, *models.Response, error) {
	url := fmt.Sprintf("http://%s/%s/send?type=%s", n.baseURL, templateName, senderType)
	resp, err := n.executeRequest(ctx, http.MethodPost, url, internal)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	response := &models.Response{}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return resp.StatusCode, response, nil
}

package client

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

func (n *Notificator) executeRequest(ctx context.Context, method, url string, body interface{}) (*http.Response, error) {
	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(&body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, url, &b)
	if err != nil {
		return nil, err
	}

	return n.client.Do(req)
}

package adapter

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type GatewayAdapter struct {
	url string
	client *http.Client
}

func NewGatewayAdapter(url string, timeout int) GatewayAdapter {

	timeoutDuration := time.Second * time.Duration(timeout)
	client := &http.Client{
		Timeout: timeoutDuration,
	}

	return GatewayAdapter{url: url, client: client}
}

func (g *GatewayAdapter) GetCategoryId(id int) (*http.Response, error) {

	body := map[string]int{
		"id" : id,
	}
	
	jsonBody, _ := json.Marshal(body)
	req, err := http.NewRequest(http.MethodPost, g.url+"/category", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	// header
	req.Header.Set("Content-Type", "application/json")

	return g.client.Do(req)	
}
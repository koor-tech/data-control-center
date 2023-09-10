package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/koor-tech/data-control-center/pkg/config"
)

type Client struct {
	Token     *string
	client    *http.Client
	apiConfig config.API
}

func NewClient(ctx context.Context, apiConfig config.API) *Client {
	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	if apiConfig.InsecureSSL {
		customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	client := &http.Client{Transport: customTransport}
	return &Client{
		apiConfig: apiConfig,
		client:    client,
	}
}

// Auth method is used to authenticate against ceph mgr api
// if the authentication is successful we are going to store the token returned by the api
// in Client.Token, if any error happens this method will return an error
func (c *Client) Auth(ctx context.Context) error {
	payload := AuthRequest{
		Username: c.apiConfig.Username,
		Password: c.apiConfig.Password,
	}

	payloadBytes, _ := json.Marshal(payload)
	resp, err := c.MakeRequest(ctx, http.MethodPost, "/auth", payloadBytes)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("invalid request %w, status code: %d", err, resp.StatusCode)
	}

	var authData ResponseData
	if err := json.NewDecoder(resp.Body).Decode(&authData); err != nil {
		return fmt.Errorf("error decoding response %w", err)
	}

	c.Token = &authData.Token
	return nil
}

func (c *Client) CreateCall(method, apiUrl string, payload *bytes.Buffer) (*http.Request, error) {
	if method == http.MethodGet {
		return http.NewRequest(method, apiUrl, nil)
	}
	return http.NewRequest(method, apiUrl, payload)
}

// MakeRequest makes a GET or POST (for now) requests to interact with the ceph api
func (c *Client) MakeRequest(ctx context.Context, method, url string, payloadBytes []byte) (*http.Response, error) {
	apiUrl := fmt.Sprintf("%s%s", c.apiConfig.Url, url)

	req, err := c.CreateCall(method, apiUrl, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Accept", headers.Accept)
	req.Header.Set("Content-Type", headers.ContentType)

	// the token is used to make request with authentication
	if c.Token != nil {
		req.Header.Set("Authorization", "Bearer "+*c.Token)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	return resp, nil
}

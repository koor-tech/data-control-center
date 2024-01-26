package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/koor-tech/data-control-center/pkg/config"
)

const tokenLifetimeDuration = 15 * time.Minute

type Client struct {
	apiConfig config.API
	client    *http.Client

	authMutex    sync.Mutex
	tokenExpires time.Time
	token        string
}

func NewClient(apiConfig config.API) *Client {
	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	if apiConfig.InsecureSkipVerify {
		customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	client := &http.Client{Transport: customTransport}
	return &Client{
		apiConfig: apiConfig,
		client:    client,

		authMutex:    sync.Mutex{},
		tokenExpires: time.Time{},
	}
}

// auth method is used to authenticate against ceph mgr api
// if the authentication is successful we are going to store the token returned by the api
// in Client.Token, if any error happens this method will return an error
func (c *Client) auth(ctx context.Context) (string, error) {
	payload := AuthRequest{
		Username: c.apiConfig.Username,
		Password: c.apiConfig.Password,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	resp, err := c.MakeRequest(ctx, NewEndpointAuth(payloadBytes), false)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("wrong response code from auth endpoint, status code %d", resp.StatusCode)
	}

	var authData ResponseData
	if err := json.NewDecoder(resp.Body).Decode(&authData); err != nil {
		return "", fmt.Errorf("error decoding response %w", err)
	}

	return authData.Token, nil
}

func (c *Client) GetToken(ctx context.Context) (string, error) {
	c.authMutex.Lock()
	defer c.authMutex.Unlock()

	if time.Since(c.tokenExpires) > tokenLifetimeDuration {
		token, err := c.auth(ctx)
		if err != nil {
			return "", err
		}

		c.token = token
		c.tokenExpires = time.Now().Add(tokenLifetimeDuration)
	}

	return c.token, nil
}

func (c *Client) CreateCall(method, apiUrl string, payload *bytes.Buffer) (*http.Request, error) {
	if method == http.MethodGet || method == http.MethodDelete {
		return http.NewRequest(method, apiUrl, nil)
	}

	return http.NewRequest(method, apiUrl, payload)
}

// MakeRequest makes a GET or POST (for now) requests to interact with the ceph api
// func (c *Client) MakeRequest(ctx context.Context, method, url string, payloadBytes []byte) (*http.Response, error) {
func (c *Client) MakeRequest(ctx context.Context, e *Endpoint, withAuth bool) (*http.Response, error) {
	apiUrl := fmt.Sprintf("%s%s", c.apiConfig.Url, e.Url)
	req, err := c.CreateCall(e.Method, apiUrl, e.Payload)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	for hName, hValue := range e.Headers() {
		req.Header.Set(hName, hValue)
	}

	if withAuth {
		token, err := c.GetToken(ctx)
		if err != nil {
			return nil, fmt.Errorf("error getting token for request: %w", err)
		}

		// the token is used to make request with authentication
		req.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	return resp, nil
}

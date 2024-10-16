package repo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Request struct {
	Name    string `json:"name"`
	Desc    string `json:"description"`
	Private bool   `json:"private"`
	Token   string `json:"token"`
}

func CreateRepo(r Request) (any, error) {
	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// construction of the request, NewRequest doesn't do the request to the remote service
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.github.com/user/repos", bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json") // setting headers // we will send a json to the server
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+r.Token)

	//Do method would fail if remote service is unavailable or ctx is cancelled or timed out
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("failed: %s\n %s status code:", data, resp.Status)
	}

	return fmt.Sprintf("success: %s", data), nil

}

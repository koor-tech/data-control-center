package update

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type release struct {
	ID      int64  `json:"id"`
	TagName string `json:"tag_name"`
	Name    string `json:"name"`
	Draft   bool   `json:"draft"`
}

func GetLatestRelease(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://api.github.com/repos/koor-tech/data-control-center/releases", nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get github latest release, status code: %d", res.StatusCode)
	}

	var dest []release
	if err := json.NewDecoder(res.Body).Decode(&dest); err != nil {
		return "", err
	}

	for _, r := range dest {
		if strings.HasPrefix(r.TagName, "data-control-center-") {
			continue
		}

		return r.Name, nil
	}

	return "UNKNOWN", nil
}

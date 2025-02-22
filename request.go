package dalle

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	authorizationStr    = `Authorization`
	contentTypeStr      = `Content-Type`
	contentTypeValueStr = `application/json`
)

func (c *Client) post(endpoint, version, payload string) ([]byte, error) {
	url := fmt.Sprintf("%s%s%s", c.Domain, version, endpoint)
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Add(authorizationStr, fmt.Sprintf("Bearer %s", c.APIKey))
	req.Header.Add(contentTypeStr, contentTypeValueStr)

	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	return body, err
}

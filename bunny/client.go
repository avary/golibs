package bunny

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/skerkour/golibs/httputils"
)

type Client struct {
	httpClient *http.Client
	accessKey  string
	baseURL    string
}

func NewClient(accessKey string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = httputils.DefaultClient()
	}

	return &Client{
		httpClient: httpClient,
		accessKey:  accessKey,
		baseURL:    "https://api.bunny.net",
	}
}

type requestParams struct {
	Method      string
	URL         string
	Payload     interface{}
	ServerToken *string
}

func (client *Client) request(params requestParams, dst interface{}) error {
	url := client.baseURL + params.URL

	req, err := http.NewRequest(params.Method, url, nil)
	if err != nil {
		return err
	}

	if params.Payload != nil {
		payloadData, err := json.Marshal(params.Payload)
		if err != nil {
			return err
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer(payloadData))
	}

	req.Header.Add(httputils.HeaderAccept, httputils.MediaTypeJson)
	req.Header.Add(httputils.HeaderContentType, httputils.MediaTypeJson)
	req.Header.Add("AccessKey", client.accessKey)

	res, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 {
		var apiErr APIError

		err = json.Unmarshal(body, &apiErr)
		if err != nil {
			return err
		}

		return apiErr
	} else if dst != nil {
		err = json.Unmarshal(body, dst)
	}

	return err
}

type APIError struct {
	Message string
}

func (res APIError) Error() string {
	return res.Message
}

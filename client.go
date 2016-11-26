package ghost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	Username string
	Password string
	APIURL   string
}

type errorObject struct {
	Code   int      `json:"code,omitempty"`
	Mesage string   `json:"message,omitempty"`
	Errors []string `json:"errors,omitempty"`
}

var netClient = &http.Client{
	Timeout: time.Second * 10,
}

func (c *Client) decodeJSON(resp *http.Response, payload interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(payload)
}

func (c *Client) getErrorFromResponse(resp *http.Response) (*errorObject, error) {
	var result map[string]errorObject
	if err := c.decodeJSON(resp, &result); err != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", err)
	}
	s, ok := result["error"]
	if !ok {
		return nil, fmt.Errorf("JSON response does not have error field")
	}
	return &s, nil
}

func (c *Client) checkResponse(resp *http.Response, err error) (*http.Response, error) {
	if err != nil {
		return resp, fmt.Errorf("Error calling the API endpoint: %v", err)
	}
	if 199 >= resp.StatusCode || 300 <= resp.StatusCode {
		//var eo *errorObject
		//var getErr error
		//if eo, getErr = c.getErrorFromResponse(resp); getErr != nil {
		//	return resp, fmt.Errorf("Response did not contain formatted error: %s. HTTP response code: %v. Raw response: %+v", getErr, resp.StatusCode, resp)
		//}
		return resp, fmt.Errorf("Failed call API endpoint. HTTP response code: %v", resp.StatusCode)
	}
	return resp, nil
}

func (c *Client) delete(path string) (*http.Response, error) {
	return c.do("DELETE", path, nil, nil)
}

func (c *Client) put(path string, payload interface{}, headers *map[string]string) (*http.Response, error) {

	if payload != nil {
		data, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		return c.do("PUT", path, bytes.NewBuffer(data), headers)
	}
	return c.do("PUT", path, nil, headers)
}

func (c *Client) post(path string, payload interface{}) (*http.Response, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return c.do("POST", path, bytes.NewBuffer(data), nil)
}

func (c *Client) get(path string) (*http.Response, error) {
	return c.do("GET", path, nil, nil)
}

func (c *Client) do(method, path string, body io.Reader, headers *map[string]string) (*http.Response, error) {
	endpoint := c.APIURL + path
	req, _ := http.NewRequest(method, endpoint, body)
	if headers != nil {
		for k, v := range *headers {
			req.Header.Set(k, v)
		}
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.Username, c.Password)

	resp, err := netClient.Do(req)
	return c.checkResponse(resp, err)
}

// NewClient Return a Ghost client
func NewClient(apiURL string, username string, password string) *Client {
	return &Client{APIURL: apiURL, Username: username, Password: password}
}

// GetApps return all apps
func (c *Client) GetApps() error {
	res, err := c.get("/apps")
	if err != nil {
		return err
	}
	var apps Apps
	json.NewDecoder(res.Body).Decode(&apps)
	if err != nil {
		return err
	}
	fmt.Println(apps)
	return nil
}

func (c *Client) GetApp(id string) error {
	res, err := c.get("/apps/" + id)
	if err != nil {
		return err
	}
	var app App
	err = json.NewDecoder(res.Body).Decode(&app)
	if err != nil {
		return err
	}
	fmt.Println(app)
	return nil
}

func (c *Client) DeleteApp(id string) {

}

func (c *Client) UpdateApp(id string, app *App) {
}

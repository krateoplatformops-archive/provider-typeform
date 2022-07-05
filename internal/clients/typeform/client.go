package typeform

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	baseApiURL = "https://api.typeform.com"
	formsPath  = "/forms"
)

type ClientOpts struct {
	Token      string
	HTTPClient *http.Client
}

type Client interface {
	CreateForm(ctx context.Context, val *Form) (*Form, error)
	GetForm(ctx context.Context, id string) (*Form, error)
	DeleteForm(ctx context.Context, id string) error
}

func NewClient(opts ClientOpts) Client {
	ret := &typeFormClient{
		token:      opts.Token,
		httpClient: opts.HTTPClient,
	}
	if ret.httpClient == nil {
		ret.httpClient = &http.Client{
			Timeout: time.Second * 30,
		}
	}
	return ret
}

type typeFormClient struct {
	token      string
	httpClient *http.Client
}

func (c *typeFormClient) CreateForm(ctx context.Context, val *Form) (*Form, error) {
	dat, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost,
		fmt.Sprintf("%s/forms", baseApiURL), bytes.NewBuffer(dat))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := new(Form)
	if err := c.sendRequest(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *typeFormClient) GetForm(ctx context.Context, id string) (*Form, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/forms/%s", baseApiURL, id), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res := &Form{}
	if err := c.sendRequest(req, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *typeFormClient) DeleteForm(ctx context.Context, id string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/forms/%s", baseApiURL, id), nil)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)

	return c.sendRequest(req, nil)
}

func (c *typeFormClient) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes Errors
		err := json.NewDecoder(res.Body).Decode(&errRes)
		if err == nil {
			return &errRes
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	if v != nil {
		return json.NewDecoder(res.Body).Decode(v)
	}

	return nil
}

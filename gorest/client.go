package gorest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	apiKey   = "9633315885737b9d78e6996a635b1ab7aa21345d8c524f0cb21d24fda0b4cd90"
	endPoint = "https://gorest.co.in/public/v2/users"
)

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout <= 0 {
		return nil, errors.New("timeout can't be <= 0")
	}
	return &Client{
		client: &http.Client{
			Timeout: timeout,
		},
	}, nil
}

func (c *Client) GetUsers() ([]User, error) {
	req, err := http.NewRequest("GET", endPoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	response, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var r usersResponse

	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return r, nil

}

func (c *Client) AddUser(user *User) (*User, error) {
	userStr := fmt.Sprintf(`{"name":"%s", "gender":"%s", "email":"%s", "status":"%s"}`, user.Name, user.Gender, user.Email, user.Status)
	jsonStr := []byte(userStr)

	req, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(jsonStr))

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	response, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	r := new(User)

	err = json.Unmarshal(body, r)

	if err != nil {
		return nil, err
	}

	return r, nil

}

func (c *Client) DeleteUser(id int) error {
	url := fmt.Sprintf("%s/%d", endPoint, id)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	response, err := c.client.Do(req)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	return nil
}

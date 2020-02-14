package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
	endpoint   string
}

func NewClient(endpoint string) (*Client, error) {
	var client = &http.Client{
		Timeout: time.Second * 5,
	}
	return &Client{client, endpoint}, nil
}

func (c *Client) CreateObjectRecord(inputobject Record) (string, error) {
	var object [4]interface{}
	object[0] = inputobject.Name
	object[1] = inputobject.Type
	object[2] = inputobject.Address
	object[3] = inputobject.Zone

	//see signature : NewRequest(method, url string, body io.Reader) (*Request, error)
	//body of type *bytes.Buffer, *bytes.Reader, or *strings.Reader
	objectbyte, _ := json.Marshal(object)
	objectReader := bytes.NewReader(objectbyte)
	//ouchhhhhhhhhhhhh
	request, err := http.NewRequest("POST", c.endpoint+"/v1/dns", objectReader)
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "HBX4FunAndProfit")
	request.Header.Set("X-Terraform-Provider-Dns-Version", "v1")
	request.Header.Set("Authorization", "Basic "+object[0].(string))

	response, err := c.httpClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	//pop last item et renvoyer l'id à terraform
	body, err := ioutil.ReadAll(response.Body)
	var resArray []interface{}
	json.Unmarshal(body, &resArray)
	var resID string
	resID = (resArray[0]).(map[string]interface{})["id"].(string) //cast

	return resID, err

}

func (c *Client) CreateObjectZone(inputobject Zone) (string, error) {
	var object [1]interface{}
	object[0] = inputobject.Name

	//see signature : NewRequest(method, url string, body io.Reader) (*Request, error)
	//body of type *bytes.Buffer, *bytes.Reader, or *strings.Reader
	objectbyte, _ := json.Marshal(object)
	objectReader := bytes.NewReader(objectbyte)
	//ouchhhhhhhhhhhhh
	request, err := http.NewRequest("POST", c.endpoint+"/v1/dns", objectReader)
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "HBX4FunAndProfit")
	request.Header.Set("X-Terraform-Provider-Nectar-Version", "v1")
	request.Header.Set("Authorization", "Basic "+object[0].(string))

	response, err := c.httpClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	//pop last item et renvoyer l'id à terraform
	body, err := ioutil.ReadAll(response.Body)
	var resArray []interface{}
	json.Unmarshal(body, &resArray)
	var resID string
	resID = (resArray[0]).(map[string]interface{})["id"].(string) //cast

	return resID, err

}

func (c *Client) ReadObject(id string) error {
	return nil
}

func (c *Client) DeleteObject(id string) error {

	deleteURL := c.endpoint + "/v1/dns/" + id

	request, err := http.NewRequest("DELETE", deleteURL, nil)
	if err != nil {
		return err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return nil
}

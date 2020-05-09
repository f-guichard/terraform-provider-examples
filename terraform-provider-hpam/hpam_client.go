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
		Timeout: time.Second * 30,
	}
	return &Client{client, endpoint}, nil
}

func (c *Client) CreateObjectSoutien(inputobject Soutien) (string, error) {
	var object [7]interface{}
	object[0] = inputobject.Exploit
	object[1] = inputobject.Support
	object[2] = inputobject.RespExploit
	object[3] = inputobject.RespSupport
	object[4] = inputobject.Etat
	object[5] = inputobject.EdsPeit
    object[6] = inputobject.EdsSat

	objectbyte, _ := json.Marshal(object)
	objectReader := bytes.NewReader(objectbyte)
	request, err := http.NewRequest("POST", c.endpoint+"/v1/26e", objectReader)
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "Tf-AssetMgmt")
	request.Header.Set("X-Terraform-Provider-AssetMgmt-Version", "v1")
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

func (c *Client) CreateObjectBasicat(inputobject Basicat) (string, error) {
	var object [1]interface{}
	object[0] = inputobject.Name

	objectbyte, _ := json.Marshal(object)
	objectReader := bytes.NewReader(objectbyte)
	request, err := http.NewRequest("POST", c.endpoint+"/v1/26e", objectReader)
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "Tf-AssetMgmt")
	request.Header.Set("X-Terraform-Provider-AssetMgmt-Version", "v1")
	request.Header.Set("Authorization", "Basic "+object[0].(string))

	response, err := c.httpClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	var resArray []interface{}
	json.Unmarshal(body, &resArray)
	var resID string
	resID = (resArray[0]).(map[string]interface{})["id"].(string) //cast

	return resID, err

}

func (c *Client) CreateObjectApplication(inputobject Application) (string, error) {
	var object [2]interface{}
	object[0] = inputobject.Modapp
	object[1] = inputobject.Name

	objectbyte, _ := json.Marshal(object)
	objectReader := bytes.NewReader(objectbyte)
	request, err := http.NewRequest("POST", c.endpoint+"/v1/26e", objectReader)
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "Tf-AssetMgmt")
	request.Header.Set("X-Terraform-Provider-AssetMgmt-Version", "v1")
	request.Header.Set("Authorization", "Basic "+object[0].(string))

	response, err := c.httpClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

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

	deleteURL := c.endpoint + "/v1/26e/" + id

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

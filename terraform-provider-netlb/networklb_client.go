package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type NetworklbClient struct {
	httpClient        *http.Client
	endpointNetworklb string
}

func NewClient(endpoint string) (*NetworklbClient, error) {
	var client = &http.Client{
		Timeout: time.Second * 5,
	}
	return &NetworklbClient{client, endpoint}, nil
}

func (nlbc *NetworklbClient) CreateVip(vip Vip) (string, error) {
	var vipobject [3]interface{}
	vipobject[0] = vip.Ipadrress
	vipobject[1] = vip.Port
	vipobject[2] = vip.TLS

	//see signature : NewRequest(method, url string, body io.Reader) (*Request, error)
	//body of type *bytes.Buffer, *bytes.Reader, or *strings.Reader
	vipbyte, _ := json.Marshal(vipobject)
	vipReader := bytes.NewReader(vipbyte)
	//ouchhhhhhhhhhhhh
	request, err := http.NewRequest("POST", nlbc.endpointNetworklb+"/v1/network-lb", vipReader)
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "99-UA")
	request.Header.Set("X-Terraform-Provider-Networklb-Version", "v1")
	request.Header.Set("Authorization", "Basic "+vipobject[0].(string))

	response, err := nlbc.httpClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	//pop last item et renvoyer l'id à terraform
	body, err := ioutil.ReadAll(response.Body)
	var resArray []interface{}
	json.Unmarshal(body, &resArray)
	var vipID string
	vipID = (resArray[0]).(map[string]interface{})["netlbid"].(string) //cast

	return vipID, err

}

func (nlbc *NetworklbClient) ReadFlow(vipID string) error {
	return nil
}

func (nlbc *NetworklbClient) DeleteVip(vipID string) error {

	deleteURL := nlbc.endpointNetworklb + "/v1/network-lb/" + vipID

	request, err := http.NewRequest("DELETE", deleteURL, nil)
	if err != nil {
		return err
	}

	response, err := nlbc.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return nil
}

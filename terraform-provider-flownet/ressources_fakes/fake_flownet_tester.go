package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type FlownetClient struct {
	httpClient      *http.Client
	endpointFlownet string
}

type Flow struct {
	Cidr        string
	Port        string
	Description string
}

func main() {
	fmt.Println("Starting....")
	flow := &Flow{"10.2.256.0/24", "5482", ""}
	//Creation client flowNet
	client, err := NewClient("http://127.0.0.1:30000")
	if err != nil {
		fmt.Printf("Erreur Client : %v", err.Error())
	}
	rez, err := client.CreateFlow(*flow)
	if err != nil {
		fmt.Printf("Erreur CreateFlow : %v", err.Error())
	}
	fmt.Printf("Retour id : %s\n", rez)

	fmt.Printf("Suppression flow %s\n", rez)
	client.DeleteFlow(rez)

	fmt.Printf("Verification statut flow %s\n", rez)
	resp, _ := client.ReadFlow(rez)

	fmt.Printf("Statut flow %s = %v\n", rez, resp)

}

func NewClient(endpoint string) (*FlownetClient, error) {
	var client = &http.Client{
		Timeout: time.Second * 5,
	}
	return &FlownetClient{client, endpoint}, nil
}

func (fnc *FlownetClient) CreateFlow(flow Flow) (string, error) {
	var flowobject [3]string
	flowobject[0] = flow.Cidr
	flowobject[1] = flow.Port

	vector := rand.NewSource(time.Now().UnixNano())
	randomstring := strconv.FormatInt(rand.New(vector).Int63(), 10)
	if len(flow.Description) < 2 {
		flowobject[2] = randomstring
	} else {
		flowobject[2] = flow.Description
	}

	//see signature : NewRequest(method, url string, body io.Reader) (*Request, error)
	//body of type *bytes.Buffer, *bytes.Reader, or *strings.Reader
	flowbyte, _ := json.Marshal(flowobject)
	flowReader := bytes.NewReader(flowbyte)
	//ouchhhhhhhhhhhhh
	request, err := http.NewRequest("POST", fnc.endpointFlownet+"/v1/flows", flowReader)
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "51-UA")
	request.Header.Set("X-Terraform-Provider-Flownet-Version", "v1")
	request.Header.Set("Authorization", "Basic "+flowobject[2])

	response, err := fnc.httpClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	fmt.Printf("Body brut : %s\n", body)
	var resArray []interface{}
	json.Unmarshal(body, &resArray)
	var flowID string
	flowID = (resArray[0]).(map[string]interface{})["flowid"].(string) //cast
	fmt.Printf("Body parse : %v\n", resArray)
	fmt.Printf("FlowID : %v\n", flowID)

	return flowID, err
}

func (fnc *FlownetClient) DeleteFlow(flowID string) error {

	deleteURL := fnc.endpointFlownet + "/v1/flows/" + flowID

	request, err := http.NewRequest("DELETE", deleteURL, nil)
	if err != nil {
		return err
	}

	response, err := fnc.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return nil
}

func (fnc *FlownetClient) ReadFlow(flowID string) (interface{}, error) {

	readURL := fnc.endpointFlownet + "/v1/flows/" + flowID

	request, err := http.NewRequest("GET", readURL, nil)
	if err != nil {
		return nil, err
	}

	response, err := fnc.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	//AWAITED response : [{0:"cidr",1:"port",2:"description"}]
	//pop last item et renvoyer l'id à terraform
	body, err := ioutil.ReadAll(response.Body)
	fmt.Printf("Body brut : %s\n", body)
	var resArray interface{}
	json.Unmarshal(body, &resArray)

	return resArray, err
}

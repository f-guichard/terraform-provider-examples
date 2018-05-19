package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type BoshClient struct {
	httpClient   *http.Client
	endpointBosh string
}

func NewClient(endpoint string) (*BoshClient, error) {
	var client = &http.Client{
		Timeout: time.Second * 5,
	}
	return &BoshClient{client, endpoint}, nil
}

//TODO use reflexion for Stemcell, Release & Deployment instead of duplicate func
func (bdc *BoshClient) CreateStemcell(stemcell Stemcell) (string, error) {
	var stemcellobject [3]interface{}
	stemcellobject[0] = stemcell.Name
	stemcellobject[1] = stemcell.Os
	stemcellobject[2] = stemcell.Version

	//see signature : NewRequest(method, url string, body io.Reader) (*Request, error)
	//body of type *bytes.Buffer, *bytes.Reader, or *strings.Reader
	stemcellbyte, _ := json.Marshal(stemcellobject)
	stemcellReader := bytes.NewReader(stemcellbyte)
	//ouchhhhhhhhhhhhh
	request, err := http.NewRequest("POST", bdc.endpointBosh+"/v1/api-director/stemcells", stemcellReader)
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "99-UA")
	request.Header.Set("X-Terraform-Provider-Bosh-Version", "v1")
	request.Header.Set("Authorization", "Basic "+stemcellobject[0].(string)+"-"+stemcellobject[2].(string))

	response, err := bdc.httpClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	//pop last item et renvoyer l'id à terraform
	body, err := ioutil.ReadAll(response.Body)
	var resArray []interface{}
	json.Unmarshal(body, &resArray)
	var CID string
	CID = (resArray[0]).(map[string]interface{})["cid"].(string) //cast

	return CID, err

}

func (bdc *BoshClient) ReadStemcell(CID string) error {
	return nil
}

func (bdc *BoshClient) DeleteStemcell(CID string) error {

	deleteURL := bdc.endpointBosh + "/v1/api-director/stemcells/" + CID

	request, err := http.NewRequest("DELETE", deleteURL, nil)
	if err != nil {
		return err
	}

	response, err := bdc.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return nil
}

func (bdc *BoshClient) CreateRelease(release Release) (string, error) {
	var releaseobject [4]interface{}
	releaseobject[0] = release.Name
	releaseobject[1] = release.Version
	releaseobject[2] = release.Hash
	releaseobject[3] = release.JobNames

	//see signature : NewRequest(method, url string, body io.Reader) (*Request, error)
	//body of type *bytes.Buffer, *bytes.Reader, or *strings.Reader
	releasebyte, _ := json.Marshal(releaseobject)
	releaseReader := bytes.NewReader(releasebyte)
	//ouchhhhhhhhhhhhh
	request, err := http.NewRequest("POST", bdc.endpointBosh+"/v1/api-director/releases", releaseReader)
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "99-UA")
	request.Header.Set("X-Terraform-Provider-Bosh-Version", "v1")
	request.Header.Set("Authorization", "Basic "+releaseobject[0].(string))

	response, err := bdc.httpClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	//pop last item et renvoyer l'id à terraform
	body, err := ioutil.ReadAll(response.Body)
	var resArray []interface{}
	json.Unmarshal(body, &resArray)
	var CID string
	CID = (resArray[0]).(map[string]interface{})["cid"].(string) //cast

	return CID, err

}

func (bdc *BoshClient) ReadRelease(CID string) error {
	return nil
}

func (bdc *BoshClient) DeleteRelease(CID string) error {

	deleteURL := bdc.endpointBosh + "/v1/api-director/releases/" + CID

	request, err := http.NewRequest("DELETE", deleteURL, nil)
	if err != nil {
		return err
	}

	response, err := bdc.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return nil
}

func (bdc *BoshClient) CreateDeployment(deployment Deployment) (string, error) {
	var deploymentobject [3]interface{}
	deploymentobject[0] = deployment.Name
	deploymentobject[1] = deployment.Releases
	deploymentobject[2] = deployment.Stemcells

	//see signature : NewRequest(method, url string, body io.Reader) (*Request, error)
	//body of type *bytes.Buffer, *bytes.Reader, or *strings.Reader
	deploymentbyte, _ := json.Marshal(deploymentobject)
	deploymentReader := bytes.NewReader(deploymentbyte)
	//ouchhhhhhhhhhhhh
	request, err := http.NewRequest("POST", bdc.endpointBosh+"/v1/api-director/deployments", deploymentReader)
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "99-UA")
	request.Header.Set("X-Terraform-Provider-Bosh-Version", "v1")
	request.Header.Set("Authorization", "Basic "+deploymentobject[0].(string))

	response, err := bdc.httpClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	//pop last item et renvoyer l'id à terraform
	body, err := ioutil.ReadAll(response.Body)
	var resArray []interface{}
	json.Unmarshal(body, &resArray)
	var CID string
	CID = (resArray[0]).(map[string]interface{})["cid"].(string) //cast

	return CID, err

}

func (bdc *BoshClient) ReadDeployment(CID string) error {
	return nil
}

func (bdc *BoshClient) DeleteDeployment(CID string) error {

	deleteURL := bdc.endpointBosh + "/v1/api-director/deployments/" + CID

	request, err := http.NewRequest("DELETE", deleteURL, nil)
	if err != nil {
		return err
	}

	response, err := bdc.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return nil
}

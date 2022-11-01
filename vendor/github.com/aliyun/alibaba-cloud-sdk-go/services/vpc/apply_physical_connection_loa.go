package vpc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ApplyPhysicalConnectionLOA invokes the vpc.ApplyPhysicalConnectionLOA API synchronously
func (client *Client) ApplyPhysicalConnectionLOA(request *ApplyPhysicalConnectionLOARequest) (response *ApplyPhysicalConnectionLOAResponse, err error) {
	response = CreateApplyPhysicalConnectionLOAResponse()
	err = client.DoAction(request, response)
	return
}

// ApplyPhysicalConnectionLOAWithChan invokes the vpc.ApplyPhysicalConnectionLOA API asynchronously
func (client *Client) ApplyPhysicalConnectionLOAWithChan(request *ApplyPhysicalConnectionLOARequest) (<-chan *ApplyPhysicalConnectionLOAResponse, <-chan error) {
	responseChan := make(chan *ApplyPhysicalConnectionLOAResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ApplyPhysicalConnectionLOA(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ApplyPhysicalConnectionLOAWithCallback invokes the vpc.ApplyPhysicalConnectionLOA API asynchronously
func (client *Client) ApplyPhysicalConnectionLOAWithCallback(request *ApplyPhysicalConnectionLOARequest, callback func(response *ApplyPhysicalConnectionLOAResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ApplyPhysicalConnectionLOAResponse
		var err error
		defer close(result)
		response, err = client.ApplyPhysicalConnectionLOA(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ApplyPhysicalConnectionLOARequest is the request struct for api ApplyPhysicalConnectionLOA
type ApplyPhysicalConnectionLOARequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer                    `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string                              `position:"Query" name:"ClientToken"`
	LineType             string                              `position:"Query" name:"LineType"`
	Si                   string                              `position:"Query" name:"Si"`
	PeerLocation         string                              `position:"Query" name:"PeerLocation"`
	ResourceOwnerAccount string                              `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth            requests.Integer                    `position:"Query" name:"Bandwidth"`
	OwnerAccount         string                              `position:"Query" name:"OwnerAccount"`
	ConstructionTime     string                              `position:"Query" name:"ConstructionTime"`
	OwnerId              requests.Integer                    `position:"Query" name:"OwnerId"`
	InstanceId           string                              `position:"Query" name:"InstanceId"`
	CompanyName          string                              `position:"Query" name:"CompanyName"`
	PMInfo               *[]ApplyPhysicalConnectionLOAPMInfo `position:"Query" name:"PMInfo"  type:"Repeated"`
}

// ApplyPhysicalConnectionLOAPMInfo is a repeated param struct in ApplyPhysicalConnectionLOARequest
type ApplyPhysicalConnectionLOAPMInfo struct {
	PMCertificateNo   string `name:"PMCertificateNo"`
	PMName            string `name:"PMName"`
	PMCertificateType string `name:"PMCertificateType"`
	PMGender          string `name:"PMGender"`
	PMContactInfo     string `name:"PMContactInfo"`
}

// ApplyPhysicalConnectionLOAResponse is the response struct for api ApplyPhysicalConnectionLOA
type ApplyPhysicalConnectionLOAResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateApplyPhysicalConnectionLOARequest creates a request to invoke ApplyPhysicalConnectionLOA API
func CreateApplyPhysicalConnectionLOARequest() (request *ApplyPhysicalConnectionLOARequest) {
	request = &ApplyPhysicalConnectionLOARequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ApplyPhysicalConnectionLOA", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateApplyPhysicalConnectionLOAResponse creates a response to parse from ApplyPhysicalConnectionLOA response
func CreateApplyPhysicalConnectionLOAResponse() (response *ApplyPhysicalConnectionLOAResponse) {
	response = &ApplyPhysicalConnectionLOAResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

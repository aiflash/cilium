package ecs

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

// CreateImageComponent invokes the ecs.CreateImageComponent API synchronously
func (client *Client) CreateImageComponent(request *CreateImageComponentRequest) (response *CreateImageComponentResponse, err error) {
	response = CreateCreateImageComponentResponse()
	err = client.DoAction(request, response)
	return
}

// CreateImageComponentWithChan invokes the ecs.CreateImageComponent API asynchronously
func (client *Client) CreateImageComponentWithChan(request *CreateImageComponentRequest) (<-chan *CreateImageComponentResponse, <-chan error) {
	responseChan := make(chan *CreateImageComponentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateImageComponent(request)
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

// CreateImageComponentWithCallback invokes the ecs.CreateImageComponent API asynchronously
func (client *Client) CreateImageComponentWithCallback(request *CreateImageComponentRequest, callback func(response *CreateImageComponentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateImageComponentResponse
		var err error
		defer close(result)
		response, err = client.CreateImageComponent(request)
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

// CreateImageComponentRequest is the request struct for api CreateImageComponent
type CreateImageComponentRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer           `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string                     `position:"Query" name:"ClientToken"`
	Description          string                     `position:"Query" name:"Description"`
	SystemType           string                     `position:"Query" name:"SystemType"`
	Content              string                     `position:"Query" name:"Content"`
	ResourceGroupId      string                     `position:"Query" name:"ResourceGroupId"`
	Tag                  *[]CreateImageComponentTag `position:"Query" name:"Tag"  type:"Repeated"`
	ResourceOwnerAccount string                     `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                     `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer           `position:"Query" name:"OwnerId"`
	ComponentType        string                     `position:"Query" name:"ComponentType"`
	Name                 string                     `position:"Query" name:"Name"`
}

// CreateImageComponentTag is a repeated param struct in CreateImageComponentRequest
type CreateImageComponentTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// CreateImageComponentResponse is the response struct for api CreateImageComponent
type CreateImageComponentResponse struct {
	*responses.BaseResponse
	ImageComponentId string `json:"ImageComponentId" xml:"ImageComponentId"`
	RequestId        string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateImageComponentRequest creates a request to invoke CreateImageComponent API
func CreateCreateImageComponentRequest() (request *CreateImageComponentRequest) {
	request = &CreateImageComponentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CreateImageComponent", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateImageComponentResponse creates a response to parse from CreateImageComponent response
func CreateCreateImageComponentResponse() (response *CreateImageComponentResponse) {
	response = &CreateImageComponentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

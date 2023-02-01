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

// ReleasePublicIpAddress invokes the ecs.ReleasePublicIpAddress API synchronously
func (client *Client) ReleasePublicIpAddress(request *ReleasePublicIpAddressRequest) (response *ReleasePublicIpAddressResponse, err error) {
	response = CreateReleasePublicIpAddressResponse()
	err = client.DoAction(request, response)
	return
}

// ReleasePublicIpAddressWithChan invokes the ecs.ReleasePublicIpAddress API asynchronously
func (client *Client) ReleasePublicIpAddressWithChan(request *ReleasePublicIpAddressRequest) (<-chan *ReleasePublicIpAddressResponse, <-chan error) {
	responseChan := make(chan *ReleasePublicIpAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReleasePublicIpAddress(request)
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

// ReleasePublicIpAddressWithCallback invokes the ecs.ReleasePublicIpAddress API asynchronously
func (client *Client) ReleasePublicIpAddressWithCallback(request *ReleasePublicIpAddressRequest, callback func(response *ReleasePublicIpAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReleasePublicIpAddressResponse
		var err error
		defer close(result)
		response, err = client.ReleasePublicIpAddress(request)
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

// ReleasePublicIpAddressRequest is the request struct for api ReleasePublicIpAddress
type ReleasePublicIpAddressRequest struct {
	*requests.RpcRequest
	DryRun          requests.Boolean `position:"Query" name:"DryRun"`
	PublicIpAddress string           `position:"Query" name:"PublicIpAddress"`
	InstanceId      string           `position:"Query" name:"InstanceId"`
}

// ReleasePublicIpAddressResponse is the response struct for api ReleasePublicIpAddress
type ReleasePublicIpAddressResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	RemainTimes string `json:"RemainTimes" xml:"RemainTimes"`
}

// CreateReleasePublicIpAddressRequest creates a request to invoke ReleasePublicIpAddress API
func CreateReleasePublicIpAddressRequest() (request *ReleasePublicIpAddressRequest) {
	request = &ReleasePublicIpAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ReleasePublicIpAddress", "", "")
	request.Method = requests.POST
	return
}

// CreateReleasePublicIpAddressResponse creates a response to parse from ReleasePublicIpAddress response
func CreateReleasePublicIpAddressResponse() (response *ReleasePublicIpAddressResponse) {
	response = &ReleasePublicIpAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

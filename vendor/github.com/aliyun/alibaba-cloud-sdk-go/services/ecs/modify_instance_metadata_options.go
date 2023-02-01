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

// ModifyInstanceMetadataOptions invokes the ecs.ModifyInstanceMetadataOptions API synchronously
func (client *Client) ModifyInstanceMetadataOptions(request *ModifyInstanceMetadataOptionsRequest) (response *ModifyInstanceMetadataOptionsResponse, err error) {
	response = CreateModifyInstanceMetadataOptionsResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyInstanceMetadataOptionsWithChan invokes the ecs.ModifyInstanceMetadataOptions API asynchronously
func (client *Client) ModifyInstanceMetadataOptionsWithChan(request *ModifyInstanceMetadataOptionsRequest) (<-chan *ModifyInstanceMetadataOptionsResponse, <-chan error) {
	responseChan := make(chan *ModifyInstanceMetadataOptionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyInstanceMetadataOptions(request)
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

// ModifyInstanceMetadataOptionsWithCallback invokes the ecs.ModifyInstanceMetadataOptions API asynchronously
func (client *Client) ModifyInstanceMetadataOptionsWithCallback(request *ModifyInstanceMetadataOptionsRequest, callback func(response *ModifyInstanceMetadataOptionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyInstanceMetadataOptionsResponse
		var err error
		defer close(result)
		response, err = client.ModifyInstanceMetadataOptions(request)
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

// ModifyInstanceMetadataOptionsRequest is the request struct for api ModifyInstanceMetadataOptions
type ModifyInstanceMetadataOptionsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId         requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstanceMetadataTags    string           `position:"Query" name:"InstanceMetadataTags"`
	HttpPutResponseHopLimit requests.Integer `position:"Query" name:"HttpPutResponseHopLimit"`
	HttpEndpoint            string           `position:"Query" name:"HttpEndpoint"`
	ResourceOwnerAccount    string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId                 requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId              string           `position:"Query" name:"InstanceId"`
	HttpTokens              string           `position:"Query" name:"HttpTokens"`
}

// ModifyInstanceMetadataOptionsResponse is the response struct for api ModifyInstanceMetadataOptions
type ModifyInstanceMetadataOptionsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyInstanceMetadataOptionsRequest creates a request to invoke ModifyInstanceMetadataOptions API
func CreateModifyInstanceMetadataOptionsRequest() (request *ModifyInstanceMetadataOptionsRequest) {
	request = &ModifyInstanceMetadataOptionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyInstanceMetadataOptions", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyInstanceMetadataOptionsResponse creates a response to parse from ModifyInstanceMetadataOptions response
func CreateModifyInstanceMetadataOptionsResponse() (response *ModifyInstanceMetadataOptionsResponse) {
	response = &ModifyInstanceMetadataOptionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

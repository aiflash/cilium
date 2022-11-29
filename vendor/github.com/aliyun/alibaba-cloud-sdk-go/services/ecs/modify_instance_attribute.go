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

// ModifyInstanceAttribute invokes the ecs.ModifyInstanceAttribute API synchronously
func (client *Client) ModifyInstanceAttribute(request *ModifyInstanceAttributeRequest) (response *ModifyInstanceAttributeResponse, err error) {
	response = CreateModifyInstanceAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyInstanceAttributeWithChan invokes the ecs.ModifyInstanceAttribute API asynchronously
func (client *Client) ModifyInstanceAttributeWithChan(request *ModifyInstanceAttributeRequest) (<-chan *ModifyInstanceAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyInstanceAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyInstanceAttribute(request)
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

// ModifyInstanceAttributeWithCallback invokes the ecs.ModifyInstanceAttribute API asynchronously
func (client *Client) ModifyInstanceAttributeWithCallback(request *ModifyInstanceAttributeRequest, callback func(response *ModifyInstanceAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyInstanceAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyInstanceAttribute(request)
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

// ModifyInstanceAttributeRequest is the request struct for api ModifyInstanceAttribute
type ModifyInstanceAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId             requests.Integer                               `position:"Query" name:"ResourceOwnerId"`
	Recyclable                  requests.Boolean                               `position:"Query" name:"Recyclable"`
	NetworkInterfaceQueueNumber requests.Integer                               `position:"Query" name:"NetworkInterfaceQueueNumber"`
	Description                 string                                         `position:"Query" name:"Description"`
	DeletionProtection          requests.Boolean                               `position:"Query" name:"DeletionProtection"`
	UserData                    string                                         `position:"Query" name:"UserData"`
	Password                    string                                         `position:"Query" name:"Password"`
	HostName                    string                                         `position:"Query" name:"HostName"`
	ResourceOwnerAccount        string                                         `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                string                                         `position:"Query" name:"OwnerAccount"`
	CreditSpecification         string                                         `position:"Query" name:"CreditSpecification"`
	OwnerId                     requests.Integer                               `position:"Query" name:"OwnerId"`
	SecurityGroupIds            *[]string                                      `position:"Query" name:"SecurityGroupIds"  type:"Repeated"`
	InstanceId                  string                                         `position:"Query" name:"InstanceId"`
	InstanceName                string                                         `position:"Query" name:"InstanceName"`
	RemoteConnectionOptions     ModifyInstanceAttributeRemoteConnectionOptions `position:"Query" name:"RemoteConnectionOptions"  type:"Struct"`
}

// ModifyInstanceAttributeRemoteConnectionOptions is a repeated param struct in ModifyInstanceAttributeRequest
type ModifyInstanceAttributeRemoteConnectionOptions struct {
	Password string `name:"Password"`
	Type     string `name:"Type"`
}

// ModifyInstanceAttributeResponse is the response struct for api ModifyInstanceAttribute
type ModifyInstanceAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyInstanceAttributeRequest creates a request to invoke ModifyInstanceAttribute API
func CreateModifyInstanceAttributeRequest() (request *ModifyInstanceAttributeRequest) {
	request = &ModifyInstanceAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyInstanceAttribute", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyInstanceAttributeResponse creates a response to parse from ModifyInstanceAttribute response
func CreateModifyInstanceAttributeResponse() (response *ModifyInstanceAttributeResponse) {
	response = &ModifyInstanceAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

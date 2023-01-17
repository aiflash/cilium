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

// EipNotifyPaid invokes the ecs.EipNotifyPaid API synchronously
func (client *Client) EipNotifyPaid(request *EipNotifyPaidRequest) (response *EipNotifyPaidResponse, err error) {
	response = CreateEipNotifyPaidResponse()
	err = client.DoAction(request, response)
	return
}

// EipNotifyPaidWithChan invokes the ecs.EipNotifyPaid API asynchronously
func (client *Client) EipNotifyPaidWithChan(request *EipNotifyPaidRequest) (<-chan *EipNotifyPaidResponse, <-chan error) {
	responseChan := make(chan *EipNotifyPaidResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EipNotifyPaid(request)
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

// EipNotifyPaidWithCallback invokes the ecs.EipNotifyPaid API asynchronously
func (client *Client) EipNotifyPaidWithCallback(request *EipNotifyPaidRequest, callback func(response *EipNotifyPaidResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EipNotifyPaidResponse
		var err error
		defer close(result)
		response, err = client.EipNotifyPaid(request)
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

// EipNotifyPaidRequest is the request struct for api EipNotifyPaid
type EipNotifyPaidRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Data                 string           `position:"Query" name:"data"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	UserCidr             string           `position:"Query" name:"UserCidr"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// EipNotifyPaidResponse is the response struct for api EipNotifyPaid
type EipNotifyPaidResponse struct {
	*responses.BaseResponse
	Code      string `json:"code" xml:"code"`
	Success   bool   `json:"success" xml:"success"`
	Message   string `json:"message" xml:"message"`
	Data      string `json:"data" xml:"data"`
	RequestId string `json:"requestId" xml:"requestId"`
}

// CreateEipNotifyPaidRequest creates a request to invoke EipNotifyPaid API
func CreateEipNotifyPaidRequest() (request *EipNotifyPaidRequest) {
	request = &EipNotifyPaidRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "EipNotifyPaid", "", "")
	request.Method = requests.POST
	return
}

// CreateEipNotifyPaidResponse creates a response to parse from EipNotifyPaid response
func CreateEipNotifyPaidResponse() (response *EipNotifyPaidResponse) {
	response = &EipNotifyPaidResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

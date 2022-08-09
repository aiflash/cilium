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

// ModifyBgpGroupAttribute invokes the vpc.ModifyBgpGroupAttribute API synchronously
func (client *Client) ModifyBgpGroupAttribute(request *ModifyBgpGroupAttributeRequest) (response *ModifyBgpGroupAttributeResponse, err error) {
	response = CreateModifyBgpGroupAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyBgpGroupAttributeWithChan invokes the vpc.ModifyBgpGroupAttribute API asynchronously
func (client *Client) ModifyBgpGroupAttributeWithChan(request *ModifyBgpGroupAttributeRequest) (<-chan *ModifyBgpGroupAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyBgpGroupAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyBgpGroupAttribute(request)
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

// ModifyBgpGroupAttributeWithCallback invokes the vpc.ModifyBgpGroupAttribute API asynchronously
func (client *Client) ModifyBgpGroupAttributeWithCallback(request *ModifyBgpGroupAttributeRequest, callback func(response *ModifyBgpGroupAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyBgpGroupAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyBgpGroupAttribute(request)
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

// ModifyBgpGroupAttributeRequest is the request struct for api ModifyBgpGroupAttribute
type ModifyBgpGroupAttributeRequest struct {
	*requests.RpcRequest
	AuthKey                  string           `position:"Query" name:"AuthKey"`
	ResourceOwnerId          requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken              string           `position:"Query" name:"ClientToken"`
	BgpGroupId               string           `position:"Query" name:"BgpGroupId"`
	Description              string           `position:"Query" name:"Description"`
	PeerAsn                  requests.Integer `position:"Query" name:"PeerAsn"`
	IsFakeAsn                requests.Boolean `position:"Query" name:"IsFakeAsn"`
	ResourceOwnerAccount     string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount             string           `position:"Query" name:"OwnerAccount"`
	RouteQuota               requests.Integer `position:"Query" name:"RouteQuota"`
	OwnerId                  requests.Integer `position:"Query" name:"OwnerId"`
	RouteUsageAlarmThreshold requests.Integer `position:"Query" name:"RouteUsageAlarmThreshold"`
	Name                     string           `position:"Query" name:"Name"`
	LocalAsn                 requests.Integer `position:"Query" name:"LocalAsn"`
}

// ModifyBgpGroupAttributeResponse is the response struct for api ModifyBgpGroupAttribute
type ModifyBgpGroupAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyBgpGroupAttributeRequest creates a request to invoke ModifyBgpGroupAttribute API
func CreateModifyBgpGroupAttributeRequest() (request *ModifyBgpGroupAttributeRequest) {
	request = &ModifyBgpGroupAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyBgpGroupAttribute", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyBgpGroupAttributeResponse creates a response to parse from ModifyBgpGroupAttribute response
func CreateModifyBgpGroupAttributeResponse() (response *ModifyBgpGroupAttributeResponse) {
	response = &ModifyBgpGroupAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

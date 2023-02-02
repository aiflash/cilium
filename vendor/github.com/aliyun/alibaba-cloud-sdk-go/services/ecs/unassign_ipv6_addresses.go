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

// UnassignIpv6Addresses invokes the ecs.UnassignIpv6Addresses API synchronously
func (client *Client) UnassignIpv6Addresses(request *UnassignIpv6AddressesRequest) (response *UnassignIpv6AddressesResponse, err error) {
	response = CreateUnassignIpv6AddressesResponse()
	err = client.DoAction(request, response)
	return
}

// UnassignIpv6AddressesWithChan invokes the ecs.UnassignIpv6Addresses API asynchronously
func (client *Client) UnassignIpv6AddressesWithChan(request *UnassignIpv6AddressesRequest) (<-chan *UnassignIpv6AddressesResponse, <-chan error) {
	responseChan := make(chan *UnassignIpv6AddressesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnassignIpv6Addresses(request)
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

// UnassignIpv6AddressesWithCallback invokes the ecs.UnassignIpv6Addresses API asynchronously
func (client *Client) UnassignIpv6AddressesWithCallback(request *UnassignIpv6AddressesRequest, callback func(response *UnassignIpv6AddressesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnassignIpv6AddressesResponse
		var err error
		defer close(result)
		response, err = client.UnassignIpv6Addresses(request)
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

// UnassignIpv6AddressesRequest is the request struct for api UnassignIpv6Addresses
type UnassignIpv6AddressesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Ipv6Prefix           *[]string        `position:"Query" name:"Ipv6Prefix"  type:"Repeated"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	NetworkInterfaceId   string           `position:"Query" name:"NetworkInterfaceId"`
	Ipv6Address          *[]string        `position:"Query" name:"Ipv6Address"  type:"Repeated"`
}

// UnassignIpv6AddressesResponse is the response struct for api UnassignIpv6Addresses
type UnassignIpv6AddressesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUnassignIpv6AddressesRequest creates a request to invoke UnassignIpv6Addresses API
func CreateUnassignIpv6AddressesRequest() (request *UnassignIpv6AddressesRequest) {
	request = &UnassignIpv6AddressesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "UnassignIpv6Addresses", "", "")
	request.Method = requests.POST
	return
}

// CreateUnassignIpv6AddressesResponse creates a response to parse from UnassignIpv6Addresses response
func CreateUnassignIpv6AddressesResponse() (response *UnassignIpv6AddressesResponse) {
	response = &UnassignIpv6AddressesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

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

// DescribeAccountAttributes invokes the ecs.DescribeAccountAttributes API synchronously
func (client *Client) DescribeAccountAttributes(request *DescribeAccountAttributesRequest) (response *DescribeAccountAttributesResponse, err error) {
	response = CreateDescribeAccountAttributesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAccountAttributesWithChan invokes the ecs.DescribeAccountAttributes API asynchronously
func (client *Client) DescribeAccountAttributesWithChan(request *DescribeAccountAttributesRequest) (<-chan *DescribeAccountAttributesResponse, <-chan error) {
	responseChan := make(chan *DescribeAccountAttributesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAccountAttributes(request)
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

// DescribeAccountAttributesWithCallback invokes the ecs.DescribeAccountAttributes API asynchronously
func (client *Client) DescribeAccountAttributesWithCallback(request *DescribeAccountAttributesRequest, callback func(response *DescribeAccountAttributesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAccountAttributesResponse
		var err error
		defer close(result)
		response, err = client.DescribeAccountAttributes(request)
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

// DescribeAccountAttributesRequest is the request struct for api DescribeAccountAttributes
type DescribeAccountAttributesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AttributeName        *[]string        `position:"Query" name:"AttributeName"  type:"Repeated"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	BizAliUid            string           `position:"Query" name:"BizAliUid"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
}

// DescribeAccountAttributesResponse is the response struct for api DescribeAccountAttributes
type DescribeAccountAttributesResponse struct {
	*responses.BaseResponse
	RequestId             string                `json:"RequestId" xml:"RequestId"`
	AccountAttributeItems AccountAttributeItems `json:"AccountAttributeItems" xml:"AccountAttributeItems"`
}

// CreateDescribeAccountAttributesRequest creates a request to invoke DescribeAccountAttributes API
func CreateDescribeAccountAttributesRequest() (request *DescribeAccountAttributesRequest) {
	request = &DescribeAccountAttributesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeAccountAttributes", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeAccountAttributesResponse creates a response to parse from DescribeAccountAttributes response
func CreateDescribeAccountAttributesResponse() (response *DescribeAccountAttributesResponse) {
	response = &DescribeAccountAttributesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

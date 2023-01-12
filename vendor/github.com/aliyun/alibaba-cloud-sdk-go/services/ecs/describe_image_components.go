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

// DescribeImageComponents invokes the ecs.DescribeImageComponents API synchronously
func (client *Client) DescribeImageComponents(request *DescribeImageComponentsRequest) (response *DescribeImageComponentsResponse, err error) {
	response = CreateDescribeImageComponentsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeImageComponentsWithChan invokes the ecs.DescribeImageComponents API asynchronously
func (client *Client) DescribeImageComponentsWithChan(request *DescribeImageComponentsRequest) (<-chan *DescribeImageComponentsResponse, <-chan error) {
	responseChan := make(chan *DescribeImageComponentsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeImageComponents(request)
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

// DescribeImageComponentsWithCallback invokes the ecs.DescribeImageComponents API asynchronously
func (client *Client) DescribeImageComponentsWithCallback(request *DescribeImageComponentsRequest, callback func(response *DescribeImageComponentsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeImageComponentsResponse
		var err error
		defer close(result)
		response, err = client.DescribeImageComponents(request)
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

// DescribeImageComponentsRequest is the request struct for api DescribeImageComponents
type DescribeImageComponentsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer              `position:"Query" name:"ResourceOwnerId"`
	ImageComponentId     *[]string                     `position:"Query" name:"ImageComponentId"  type:"Repeated"`
	ResourceGroupId      string                        `position:"Query" name:"ResourceGroupId"`
	NextToken            string                        `position:"Query" name:"NextToken"`
	Tag                  *[]DescribeImageComponentsTag `position:"Query" name:"Tag"  type:"Repeated"`
	Owner                string                        `position:"Query" name:"Owner"`
	ResourceOwnerAccount string                        `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                        `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer              `position:"Query" name:"OwnerId"`
	Name                 string                        `position:"Query" name:"Name"`
	MaxResults           requests.Integer              `position:"Query" name:"MaxResults"`
}

// DescribeImageComponentsTag is a repeated param struct in DescribeImageComponentsRequest
type DescribeImageComponentsTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// DescribeImageComponentsResponse is the response struct for api DescribeImageComponents
type DescribeImageComponentsResponse struct {
	*responses.BaseResponse
	NextToken      string         `json:"NextToken" xml:"NextToken"`
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	TotalCount     int            `json:"TotalCount" xml:"TotalCount"`
	MaxResults     int            `json:"MaxResults" xml:"MaxResults"`
	ImageComponent ImageComponent `json:"ImageComponent" xml:"ImageComponent"`
}

// CreateDescribeImageComponentsRequest creates a request to invoke DescribeImageComponents API
func CreateDescribeImageComponentsRequest() (request *DescribeImageComponentsRequest) {
	request = &DescribeImageComponentsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeImageComponents", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeImageComponentsResponse creates a response to parse from DescribeImageComponents response
func CreateDescribeImageComponentsResponse() (response *DescribeImageComponentsResponse) {
	response = &DescribeImageComponentsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

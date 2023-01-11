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

// DescribeManagedInstances invokes the ecs.DescribeManagedInstances API synchronously
func (client *Client) DescribeManagedInstances(request *DescribeManagedInstancesRequest) (response *DescribeManagedInstancesResponse, err error) {
	response = CreateDescribeManagedInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeManagedInstancesWithChan invokes the ecs.DescribeManagedInstances API asynchronously
func (client *Client) DescribeManagedInstancesWithChan(request *DescribeManagedInstancesRequest) (<-chan *DescribeManagedInstancesResponse, <-chan error) {
	responseChan := make(chan *DescribeManagedInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeManagedInstances(request)
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

// DescribeManagedInstancesWithCallback invokes the ecs.DescribeManagedInstances API asynchronously
func (client *Client) DescribeManagedInstancesWithCallback(request *DescribeManagedInstancesRequest, callback func(response *DescribeManagedInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeManagedInstancesResponse
		var err error
		defer close(result)
		response, err = client.DescribeManagedInstances(request)
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

// DescribeManagedInstancesRequest is the request struct for api DescribeManagedInstances
type DescribeManagedInstancesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OsType               string           `position:"Query" name:"OsType"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceName         string           `position:"Query" name:"InstanceName"`
	InstanceId           *[]string        `position:"Query" name:"InstanceId"  type:"Repeated"`
	InstanceIp           string           `position:"Query" name:"InstanceIp"`
	ActivationId         string           `position:"Query" name:"ActivationId"`
}

// DescribeManagedInstancesResponse is the response struct for api DescribeManagedInstances
type DescribeManagedInstancesResponse struct {
	*responses.BaseResponse
	PageSize   int64      `json:"PageSize" xml:"PageSize"`
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	PageNumber int64      `json:"PageNumber" xml:"PageNumber"`
	TotalCount int64      `json:"TotalCount" xml:"TotalCount"`
	Instances  []Instance `json:"Instances" xml:"Instances"`
}

// CreateDescribeManagedInstancesRequest creates a request to invoke DescribeManagedInstances API
func CreateDescribeManagedInstancesRequest() (request *DescribeManagedInstancesRequest) {
	request = &DescribeManagedInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeManagedInstances", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeManagedInstancesResponse creates a response to parse from DescribeManagedInstances response
func CreateDescribeManagedInstancesResponse() (response *DescribeManagedInstancesResponse) {
	response = &DescribeManagedInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

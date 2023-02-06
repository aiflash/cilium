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

// CreateDedicatedHostCluster invokes the ecs.CreateDedicatedHostCluster API synchronously
func (client *Client) CreateDedicatedHostCluster(request *CreateDedicatedHostClusterRequest) (response *CreateDedicatedHostClusterResponse, err error) {
	response = CreateCreateDedicatedHostClusterResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDedicatedHostClusterWithChan invokes the ecs.CreateDedicatedHostCluster API asynchronously
func (client *Client) CreateDedicatedHostClusterWithChan(request *CreateDedicatedHostClusterRequest) (<-chan *CreateDedicatedHostClusterResponse, <-chan error) {
	responseChan := make(chan *CreateDedicatedHostClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDedicatedHostCluster(request)
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

// CreateDedicatedHostClusterWithCallback invokes the ecs.CreateDedicatedHostCluster API asynchronously
func (client *Client) CreateDedicatedHostClusterWithCallback(request *CreateDedicatedHostClusterRequest, callback func(response *CreateDedicatedHostClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDedicatedHostClusterResponse
		var err error
		defer close(result)
		response, err = client.CreateDedicatedHostCluster(request)
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

// CreateDedicatedHostClusterRequest is the request struct for api CreateDedicatedHostCluster
type CreateDedicatedHostClusterRequest struct {
	*requests.RpcRequest
	DedicatedHostClusterName string                           `position:"Query" name:"DedicatedHostClusterName"`
	ResourceOwnerId          requests.Integer                 `position:"Query" name:"ResourceOwnerId"`
	Description              string                           `position:"Query" name:"Description"`
	ResourceGroupId          string                           `position:"Query" name:"ResourceGroupId"`
	Tag                      *[]CreateDedicatedHostClusterTag `position:"Query" name:"Tag"  type:"Repeated"`
	DryRun                   requests.Boolean                 `position:"Query" name:"DryRun"`
	ResourceOwnerAccount     string                           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount             string                           `position:"Query" name:"OwnerAccount"`
	OwnerId                  requests.Integer                 `position:"Query" name:"OwnerId"`
	ZoneId                   string                           `position:"Query" name:"ZoneId"`
}

// CreateDedicatedHostClusterTag is a repeated param struct in CreateDedicatedHostClusterRequest
type CreateDedicatedHostClusterTag struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// CreateDedicatedHostClusterResponse is the response struct for api CreateDedicatedHostCluster
type CreateDedicatedHostClusterResponse struct {
	*responses.BaseResponse
	DedicatedHostClusterId string `json:"DedicatedHostClusterId" xml:"DedicatedHostClusterId"`
	RequestId              string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateDedicatedHostClusterRequest creates a request to invoke CreateDedicatedHostCluster API
func CreateCreateDedicatedHostClusterRequest() (request *CreateDedicatedHostClusterRequest) {
	request = &CreateDedicatedHostClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CreateDedicatedHostCluster", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateDedicatedHostClusterResponse creates a response to parse from CreateDedicatedHostCluster response
func CreateCreateDedicatedHostClusterResponse() (response *CreateDedicatedHostClusterResponse) {
	response = &CreateDedicatedHostClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

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

// CreateAutoSnapshotPolicy invokes the ecs.CreateAutoSnapshotPolicy API synchronously
func (client *Client) CreateAutoSnapshotPolicy(request *CreateAutoSnapshotPolicyRequest) (response *CreateAutoSnapshotPolicyResponse, err error) {
	response = CreateCreateAutoSnapshotPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAutoSnapshotPolicyWithChan invokes the ecs.CreateAutoSnapshotPolicy API asynchronously
func (client *Client) CreateAutoSnapshotPolicyWithChan(request *CreateAutoSnapshotPolicyRequest) (<-chan *CreateAutoSnapshotPolicyResponse, <-chan error) {
	responseChan := make(chan *CreateAutoSnapshotPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAutoSnapshotPolicy(request)
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

// CreateAutoSnapshotPolicyWithCallback invokes the ecs.CreateAutoSnapshotPolicy API asynchronously
func (client *Client) CreateAutoSnapshotPolicyWithCallback(request *CreateAutoSnapshotPolicyRequest, callback func(response *CreateAutoSnapshotPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAutoSnapshotPolicyResponse
		var err error
		defer close(result)
		response, err = client.CreateAutoSnapshotPolicy(request)
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

// CreateAutoSnapshotPolicyRequest is the request struct for api CreateAutoSnapshotPolicy
type CreateAutoSnapshotPolicyRequest struct {
	*requests.RpcRequest
	ResourceOwnerId              requests.Integer               `position:"Query" name:"ResourceOwnerId"`
	CopiedSnapshotsRetentionDays requests.Integer               `position:"Query" name:"CopiedSnapshotsRetentionDays"`
	TimePoints                   string                         `position:"Query" name:"timePoints"`
	RepeatWeekdays               string                         `position:"Query" name:"repeatWeekdays"`
	ResourceGroupId              string                         `position:"Query" name:"ResourceGroupId"`
	StorageLocationArn           string                         `position:"Query" name:"StorageLocationArn"`
	Tag                          *[]CreateAutoSnapshotPolicyTag `position:"Query" name:"Tag"  type:"Repeated"`
	EnableCrossRegionCopy        requests.Boolean               `position:"Query" name:"EnableCrossRegionCopy"`
	ResourceOwnerAccount         string                         `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId                      requests.Integer               `position:"Query" name:"OwnerId"`
	AutoSnapshotPolicyName       string                         `position:"Query" name:"autoSnapshotPolicyName"`
	RetentionDays                requests.Integer               `position:"Query" name:"retentionDays"`
	TargetCopyRegions            string                         `position:"Query" name:"TargetCopyRegions"`
}

// CreateAutoSnapshotPolicyTag is a repeated param struct in CreateAutoSnapshotPolicyRequest
type CreateAutoSnapshotPolicyTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateAutoSnapshotPolicyResponse is the response struct for api CreateAutoSnapshotPolicy
type CreateAutoSnapshotPolicyResponse struct {
	*responses.BaseResponse
	AutoSnapshotPolicyId string `json:"AutoSnapshotPolicyId" xml:"AutoSnapshotPolicyId"`
	RequestId            string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateAutoSnapshotPolicyRequest creates a request to invoke CreateAutoSnapshotPolicy API
func CreateCreateAutoSnapshotPolicyRequest() (request *CreateAutoSnapshotPolicyRequest) {
	request = &CreateAutoSnapshotPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CreateAutoSnapshotPolicy", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateAutoSnapshotPolicyResponse creates a response to parse from CreateAutoSnapshotPolicy response
func CreateCreateAutoSnapshotPolicyResponse() (response *CreateAutoSnapshotPolicyResponse) {
	response = &CreateAutoSnapshotPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

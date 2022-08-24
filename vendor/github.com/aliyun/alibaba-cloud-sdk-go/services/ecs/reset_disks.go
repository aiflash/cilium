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

// ResetDisks invokes the ecs.ResetDisks API synchronously
func (client *Client) ResetDisks(request *ResetDisksRequest) (response *ResetDisksResponse, err error) {
	response = CreateResetDisksResponse()
	err = client.DoAction(request, response)
	return
}

// ResetDisksWithChan invokes the ecs.ResetDisks API asynchronously
func (client *Client) ResetDisksWithChan(request *ResetDisksRequest) (<-chan *ResetDisksResponse, <-chan error) {
	responseChan := make(chan *ResetDisksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResetDisks(request)
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

// ResetDisksWithCallback invokes the ecs.ResetDisks API asynchronously
func (client *Client) ResetDisksWithCallback(request *ResetDisksRequest, callback func(response *ResetDisksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResetDisksResponse
		var err error
		defer close(result)
		response, err = client.ResetDisks(request)
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

// ResetDisksRequest is the request struct for api ResetDisks
type ResetDisksRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer  `position:"Query" name:"ResourceOwnerId"`
	DryRun               requests.Boolean  `position:"Query" name:"DryRun"`
	ResourceOwnerAccount string            `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string            `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer  `position:"Query" name:"OwnerId"`
	Disk                 *[]ResetDisksDisk `position:"Query" name:"Disk"  type:"Repeated"`
}

// ResetDisksDisk is a repeated param struct in ResetDisksRequest
type ResetDisksDisk struct {
	SnapshotId string `name:"SnapshotId"`
	DiskId     string `name:"DiskId"`
}

// ResetDisksResponse is the response struct for api ResetDisks
type ResetDisksResponse struct {
	*responses.BaseResponse
	RequestId            string                           `json:"RequestId" xml:"RequestId"`
	OperationProgressSet OperationProgressSetInResetDisks `json:"OperationProgressSet" xml:"OperationProgressSet"`
}

// CreateResetDisksRequest creates a request to invoke ResetDisks API
func CreateResetDisksRequest() (request *ResetDisksRequest) {
	request = &ResetDisksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ResetDisks", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateResetDisksResponse creates a response to parse from ResetDisks response
func CreateResetDisksResponse() (response *ResetDisksResponse) {
	response = &ResetDisksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

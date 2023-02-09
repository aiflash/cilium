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

// ReInitDisk invokes the ecs.ReInitDisk API synchronously
func (client *Client) ReInitDisk(request *ReInitDiskRequest) (response *ReInitDiskResponse, err error) {
	response = CreateReInitDiskResponse()
	err = client.DoAction(request, response)
	return
}

// ReInitDiskWithChan invokes the ecs.ReInitDisk API asynchronously
func (client *Client) ReInitDiskWithChan(request *ReInitDiskRequest) (<-chan *ReInitDiskResponse, <-chan error) {
	responseChan := make(chan *ReInitDiskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReInitDisk(request)
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

// ReInitDiskWithCallback invokes the ecs.ReInitDisk API asynchronously
func (client *Client) ReInitDiskWithCallback(request *ReInitDiskRequest, callback func(response *ReInitDiskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReInitDiskResponse
		var err error
		defer close(result)
		response, err = client.ReInitDisk(request)
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

// ReInitDiskRequest is the request struct for api ReInitDisk
type ReInitDiskRequest struct {
	*requests.RpcRequest
	ResourceOwnerId             requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AutoStartInstance           requests.Boolean `position:"Query" name:"AutoStartInstance"`
	SecurityEnhancementStrategy string           `position:"Query" name:"SecurityEnhancementStrategy"`
	KeyPairName                 string           `position:"Query" name:"KeyPairName"`
	Password                    string           `position:"Query" name:"Password"`
	LoginAsNonRoot              requests.Boolean `position:"Query" name:"LoginAsNonRoot"`
	DiskId                      string           `position:"Query" name:"DiskId"`
	ResourceOwnerAccount        string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                string           `position:"Query" name:"OwnerAccount"`
	OwnerId                     requests.Integer `position:"Query" name:"OwnerId"`
}

// ReInitDiskResponse is the response struct for api ReInitDisk
type ReInitDiskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateReInitDiskRequest creates a request to invoke ReInitDisk API
func CreateReInitDiskRequest() (request *ReInitDiskRequest) {
	request = &ReInitDiskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ReInitDisk", "", "")
	request.Method = requests.POST
	return
}

// CreateReInitDiskResponse creates a response to parse from ReInitDisk response
func CreateReInitDiskResponse() (response *ReInitDiskResponse) {
	response = &ReInitDiskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

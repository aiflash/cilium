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

// CreateImage invokes the ecs.CreateImage API synchronously
func (client *Client) CreateImage(request *CreateImageRequest) (response *CreateImageResponse, err error) {
	response = CreateCreateImageResponse()
	err = client.DoAction(request, response)
	return
}

// CreateImageWithChan invokes the ecs.CreateImage API asynchronously
func (client *Client) CreateImageWithChan(request *CreateImageRequest) (<-chan *CreateImageResponse, <-chan error) {
	responseChan := make(chan *CreateImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateImage(request)
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

// CreateImageWithCallback invokes the ecs.CreateImage API asynchronously
func (client *Client) CreateImageWithCallback(request *CreateImageRequest, callback func(response *CreateImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateImageResponse
		var err error
		defer close(result)
		response, err = client.CreateImage(request)
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

// CreateImageRequest is the request struct for api CreateImage
type CreateImageRequest struct {
	*requests.RpcRequest
	DiskDeviceMapping    *[]CreateImageDiskDeviceMapping `position:"Query" name:"DiskDeviceMapping"  type:"Repeated"`
	ResourceOwnerId      requests.Integer                `position:"Query" name:"ResourceOwnerId"`
	SnapshotId           string                          `position:"Query" name:"SnapshotId"`
	ClientToken          string                          `position:"Query" name:"ClientToken"`
	SystemTag            *[]CreateImageSystemTag         `position:"Query" name:"SystemTag"  type:"Repeated"`
	Description          string                          `position:"Query" name:"Description"`
	Platform             string                          `position:"Query" name:"Platform"`
	ResourceGroupId      string                          `position:"Query" name:"ResourceGroupId"`
	BootMode             string                          `position:"Query" name:"BootMode"`
	ImageName            string                          `position:"Query" name:"ImageName"`
	StorageLocationArn   string                          `position:"Query" name:"StorageLocationArn"`
	Tag                  *[]CreateImageTag               `position:"Query" name:"Tag"  type:"Repeated"`
	Architecture         string                          `position:"Query" name:"Architecture"`
	DetectionStrategy    string                          `position:"Query" name:"DetectionStrategy"`
	ResourceOwnerAccount string                          `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                          `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer                `position:"Query" name:"OwnerId"`
	InstanceId           string                          `position:"Query" name:"InstanceId"`
	ImageFamily          string                          `position:"Query" name:"ImageFamily"`
	ImageVersion         string                          `position:"Query" name:"ImageVersion"`
}

// CreateImageDiskDeviceMapping is a repeated param struct in CreateImageRequest
type CreateImageDiskDeviceMapping struct {
	SnapshotId string `name:"SnapshotId"`
	Size       string `name:"Size"`
	DiskType   string `name:"DiskType"`
	Device     string `name:"Device"`
}

// CreateImageSystemTag is a repeated param struct in CreateImageRequest
type CreateImageSystemTag struct {
	Scope string `name:"Scope"`
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateImageTag is a repeated param struct in CreateImageRequest
type CreateImageTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateImageResponse is the response struct for api CreateImage
type CreateImageResponse struct {
	*responses.BaseResponse
	ImageId   string `json:"ImageId" xml:"ImageId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateImageRequest creates a request to invoke CreateImage API
func CreateCreateImageRequest() (request *CreateImageRequest) {
	request = &CreateImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "CreateImage", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateImageResponse creates a response to parse from CreateImage response
func CreateCreateImageResponse() (response *CreateImageResponse) {
	response = &CreateImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

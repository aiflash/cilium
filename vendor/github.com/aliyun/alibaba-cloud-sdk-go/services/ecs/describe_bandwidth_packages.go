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

// DescribeBandwidthPackages invokes the ecs.DescribeBandwidthPackages API synchronously
func (client *Client) DescribeBandwidthPackages(request *DescribeBandwidthPackagesRequest) (response *DescribeBandwidthPackagesResponse, err error) {
	response = CreateDescribeBandwidthPackagesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeBandwidthPackagesWithChan invokes the ecs.DescribeBandwidthPackages API asynchronously
func (client *Client) DescribeBandwidthPackagesWithChan(request *DescribeBandwidthPackagesRequest) (<-chan *DescribeBandwidthPackagesResponse, <-chan error) {
	responseChan := make(chan *DescribeBandwidthPackagesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBandwidthPackages(request)
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

// DescribeBandwidthPackagesWithCallback invokes the ecs.DescribeBandwidthPackages API asynchronously
func (client *Client) DescribeBandwidthPackagesWithCallback(request *DescribeBandwidthPackagesRequest, callback func(response *DescribeBandwidthPackagesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBandwidthPackagesResponse
		var err error
		defer close(result)
		response, err = client.DescribeBandwidthPackages(request)
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

// DescribeBandwidthPackagesRequest is the request struct for api DescribeBandwidthPackages
type DescribeBandwidthPackagesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	NatGatewayId         string           `position:"Query" name:"NatGatewayId"`
	BandwidthPackageId   string           `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeBandwidthPackagesResponse is the response struct for api DescribeBandwidthPackages
type DescribeBandwidthPackagesResponse struct {
	*responses.BaseResponse
	PageSize          int               `json:"PageSize" xml:"PageSize"`
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	PageNumber        int               `json:"PageNumber" xml:"PageNumber"`
	TotalCount        int               `json:"TotalCount" xml:"TotalCount"`
	BandwidthPackages BandwidthPackages `json:"BandwidthPackages" xml:"BandwidthPackages"`
}

// CreateDescribeBandwidthPackagesRequest creates a request to invoke DescribeBandwidthPackages API
func CreateDescribeBandwidthPackagesRequest() (request *DescribeBandwidthPackagesRequest) {
	request = &DescribeBandwidthPackagesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeBandwidthPackages", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeBandwidthPackagesResponse creates a response to parse from DescribeBandwidthPackages response
func CreateDescribeBandwidthPackagesResponse() (response *DescribeBandwidthPackagesResponse) {
	response = &DescribeBandwidthPackagesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

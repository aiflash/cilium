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

// DescribeSecurityGroups invokes the ecs.DescribeSecurityGroups API synchronously
func (client *Client) DescribeSecurityGroups(request *DescribeSecurityGroupsRequest) (response *DescribeSecurityGroupsResponse, err error) {
	response = CreateDescribeSecurityGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSecurityGroupsWithChan invokes the ecs.DescribeSecurityGroups API asynchronously
func (client *Client) DescribeSecurityGroupsWithChan(request *DescribeSecurityGroupsRequest) (<-chan *DescribeSecurityGroupsResponse, <-chan error) {
	responseChan := make(chan *DescribeSecurityGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSecurityGroups(request)
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

// DescribeSecurityGroupsWithCallback invokes the ecs.DescribeSecurityGroups API asynchronously
func (client *Client) DescribeSecurityGroupsWithCallback(request *DescribeSecurityGroupsRequest, callback func(response *DescribeSecurityGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSecurityGroupsResponse
		var err error
		defer close(result)
		response, err = client.DescribeSecurityGroups(request)
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

// DescribeSecurityGroupsRequest is the request struct for api DescribeSecurityGroups
type DescribeSecurityGroupsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer             `position:"Query" name:"ResourceOwnerId"`
	FuzzyQuery           requests.Boolean             `position:"Query" name:"FuzzyQuery"`
	SecurityGroupId      string                       `position:"Query" name:"SecurityGroupId"`
	IsQueryEcsCount      requests.Boolean             `position:"Query" name:"IsQueryEcsCount"`
	NetworkType          string                       `position:"Query" name:"NetworkType"`
	SecurityGroupName    string                       `position:"Query" name:"SecurityGroupName"`
	PageNumber           requests.Integer             `position:"Query" name:"PageNumber"`
	ResourceGroupId      string                       `position:"Query" name:"ResourceGroupId"`
	NextToken            string                       `position:"Query" name:"NextToken"`
	PageSize             requests.Integer             `position:"Query" name:"PageSize"`
	Tag                  *[]DescribeSecurityGroupsTag `position:"Query" name:"Tag"  type:"Repeated"`
	DryRun               requests.Boolean             `position:"Query" name:"DryRun"`
	ResourceOwnerAccount string                       `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                       `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer             `position:"Query" name:"OwnerId"`
	SecurityGroupIds     string                       `position:"Query" name:"SecurityGroupIds"`
	SecurityGroupType    string                       `position:"Query" name:"SecurityGroupType"`
	VpcId                string                       `position:"Query" name:"VpcId"`
	MaxResults           requests.Integer             `position:"Query" name:"MaxResults"`
}

// DescribeSecurityGroupsTag is a repeated param struct in DescribeSecurityGroupsRequest
type DescribeSecurityGroupsTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// DescribeSecurityGroupsResponse is the response struct for api DescribeSecurityGroups
type DescribeSecurityGroupsResponse struct {
	*responses.BaseResponse
	PageSize       int            `json:"PageSize" xml:"PageSize"`
	PageNumber     int            `json:"PageNumber" xml:"PageNumber"`
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	TotalCount     int            `json:"TotalCount" xml:"TotalCount"`
	RegionId       string         `json:"RegionId" xml:"RegionId"`
	NextToken      string         `json:"NextToken" xml:"NextToken"`
	SecurityGroups SecurityGroups `json:"SecurityGroups" xml:"SecurityGroups"`
}

// CreateDescribeSecurityGroupsRequest creates a request to invoke DescribeSecurityGroups API
func CreateDescribeSecurityGroupsRequest() (request *DescribeSecurityGroupsRequest) {
	request = &DescribeSecurityGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeSecurityGroups", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSecurityGroupsResponse creates a response to parse from DescribeSecurityGroups response
func CreateDescribeSecurityGroupsResponse() (response *DescribeSecurityGroupsResponse) {
	response = &DescribeSecurityGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

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

// DeleteDiagnosticMetricSets invokes the ecs.DeleteDiagnosticMetricSets API synchronously
func (client *Client) DeleteDiagnosticMetricSets(request *DeleteDiagnosticMetricSetsRequest) (response *DeleteDiagnosticMetricSetsResponse, err error) {
	response = CreateDeleteDiagnosticMetricSetsResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDiagnosticMetricSetsWithChan invokes the ecs.DeleteDiagnosticMetricSets API asynchronously
func (client *Client) DeleteDiagnosticMetricSetsWithChan(request *DeleteDiagnosticMetricSetsRequest) (<-chan *DeleteDiagnosticMetricSetsResponse, <-chan error) {
	responseChan := make(chan *DeleteDiagnosticMetricSetsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDiagnosticMetricSets(request)
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

// DeleteDiagnosticMetricSetsWithCallback invokes the ecs.DeleteDiagnosticMetricSets API asynchronously
func (client *Client) DeleteDiagnosticMetricSetsWithCallback(request *DeleteDiagnosticMetricSetsRequest, callback func(response *DeleteDiagnosticMetricSetsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDiagnosticMetricSetsResponse
		var err error
		defer close(result)
		response, err = client.DeleteDiagnosticMetricSets(request)
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

// DeleteDiagnosticMetricSetsRequest is the request struct for api DeleteDiagnosticMetricSets
type DeleteDiagnosticMetricSetsRequest struct {
	*requests.RpcRequest
	MetricSetIds *[]string `position:"Query" name:"MetricSetIds"  type:"Repeated"`
}

// DeleteDiagnosticMetricSetsResponse is the response struct for api DeleteDiagnosticMetricSets
type DeleteDiagnosticMetricSetsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteDiagnosticMetricSetsRequest creates a request to invoke DeleteDiagnosticMetricSets API
func CreateDeleteDiagnosticMetricSetsRequest() (request *DeleteDiagnosticMetricSetsRequest) {
	request = &DeleteDiagnosticMetricSetsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DeleteDiagnosticMetricSets", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteDiagnosticMetricSetsResponse creates a response to parse from DeleteDiagnosticMetricSets response
func CreateDeleteDiagnosticMetricSetsResponse() (response *DeleteDiagnosticMetricSetsResponse) {
	response = &DeleteDiagnosticMetricSetsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

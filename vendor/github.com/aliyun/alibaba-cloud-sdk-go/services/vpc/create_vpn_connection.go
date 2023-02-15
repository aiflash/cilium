package vpc

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

// CreateVpnConnection invokes the vpc.CreateVpnConnection API synchronously
func (client *Client) CreateVpnConnection(request *CreateVpnConnectionRequest) (response *CreateVpnConnectionResponse, err error) {
	response = CreateCreateVpnConnectionResponse()
	err = client.DoAction(request, response)
	return
}

// CreateVpnConnectionWithChan invokes the vpc.CreateVpnConnection API asynchronously
func (client *Client) CreateVpnConnectionWithChan(request *CreateVpnConnectionRequest) (<-chan *CreateVpnConnectionResponse, <-chan error) {
	responseChan := make(chan *CreateVpnConnectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateVpnConnection(request)
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

// CreateVpnConnectionWithCallback invokes the vpc.CreateVpnConnection API asynchronously
func (client *Client) CreateVpnConnectionWithCallback(request *CreateVpnConnectionRequest, callback func(response *CreateVpnConnectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateVpnConnectionResponse
		var err error
		defer close(result)
		response, err = client.CreateVpnConnection(request)
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

// CreateVpnConnectionRequest is the request struct for api CreateVpnConnection
type CreateVpnConnectionRequest struct {
	*requests.RpcRequest
	IkeConfig            string                     `position:"Query" name:"IkeConfig"`
	ResourceOwnerId      requests.Integer           `position:"Query" name:"ResourceOwnerId"`
	AutoConfigRoute      requests.Boolean           `position:"Query" name:"AutoConfigRoute"`
	ClientToken          string                     `position:"Query" name:"ClientToken"`
	IpsecConfig          string                     `position:"Query" name:"IpsecConfig"`
	BgpConfig            string                     `position:"Query" name:"BgpConfig"`
	HealthCheckConfig    string                     `position:"Query" name:"HealthCheckConfig"`
	CustomerGatewayId    string                     `position:"Query" name:"CustomerGatewayId"`
	LocalSubnet          string                     `position:"Query" name:"LocalSubnet"`
	RemoteSubnet         string                     `position:"Query" name:"RemoteSubnet"`
	EffectImmediately    requests.Boolean           `position:"Query" name:"EffectImmediately"`
	ResourceOwnerAccount string                     `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                     `position:"Query" name:"OwnerAccount"`
	VpnGatewayId         string                     `position:"Query" name:"VpnGatewayId"`
	OwnerId              requests.Integer           `position:"Query" name:"OwnerId"`
	EnableDpd            requests.Boolean           `position:"Query" name:"EnableDpd"`
	Tags                 *[]CreateVpnConnectionTags `position:"Query" name:"Tags"  type:"Repeated"`
	RemoteCaCertificate  string                     `position:"Query" name:"RemoteCaCertificate"`
	Name                 string                     `position:"Query" name:"Name"`
	EnableNatTraversal   requests.Boolean           `position:"Query" name:"EnableNatTraversal"`
}

// CreateVpnConnectionTags is a repeated param struct in CreateVpnConnectionRequest
type CreateVpnConnectionTags struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateVpnConnectionResponse is the response struct for api CreateVpnConnection
type CreateVpnConnectionResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	VpnConnectionId string `json:"VpnConnectionId" xml:"VpnConnectionId"`
	Name            string `json:"Name" xml:"Name"`
	CreateTime      int64  `json:"CreateTime" xml:"CreateTime"`
}

// CreateCreateVpnConnectionRequest creates a request to invoke CreateVpnConnection API
func CreateCreateVpnConnectionRequest() (request *CreateVpnConnectionRequest) {
	request = &CreateVpnConnectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateVpnConnection", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateVpnConnectionResponse creates a response to parse from CreateVpnConnection response
func CreateCreateVpnConnectionResponse() (response *CreateVpnConnectionResponse) {
	response = &CreateVpnConnectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}

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

// AssociatedPhysicalConnection is a nested struct in vpc response
type AssociatedPhysicalConnection struct {
	Status                           string `json:"Status" xml:"Status"`
	VlanInterfaceId                  string `json:"VlanInterfaceId" xml:"VlanInterfaceId"`
	CircuitCode                      string `json:"CircuitCode" xml:"CircuitCode"`
	PeerIpv6GatewayIp                string `json:"PeerIpv6GatewayIp" xml:"PeerIpv6GatewayIp"`
	LocalIpv6GatewayIp               string `json:"LocalIpv6GatewayIp" xml:"LocalIpv6GatewayIp"`
	PhysicalConnectionOwnerUid       string `json:"PhysicalConnectionOwnerUid" xml:"PhysicalConnectionOwnerUid"`
	LocalGatewayIp                   string `json:"LocalGatewayIp" xml:"LocalGatewayIp"`
	PhysicalConnectionBusinessStatus string `json:"PhysicalConnectionBusinessStatus" xml:"PhysicalConnectionBusinessStatus"`
	PeeringSubnetMask                string `json:"PeeringSubnetMask" xml:"PeeringSubnetMask"`
	EnableIpv6                       bool   `json:"EnableIpv6" xml:"EnableIpv6"`
	PhysicalConnectionStatus         string `json:"PhysicalConnectionStatus" xml:"PhysicalConnectionStatus"`
	PeerGatewayIp                    string `json:"PeerGatewayIp" xml:"PeerGatewayIp"`
	PeeringIpv6SubnetMask            string `json:"PeeringIpv6SubnetMask" xml:"PeeringIpv6SubnetMask"`
	PhysicalConnectionId             string `json:"PhysicalConnectionId" xml:"PhysicalConnectionId"`
	VlanId                           string `json:"VlanId" xml:"VlanId"`
}

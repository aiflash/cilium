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

// Disk is a nested struct in ecs response
type Disk struct {
	Category                      string                        `json:"Category" xml:"Category"`
	BdfId                         string                        `json:"BdfId" xml:"BdfId"`
	ImageId                       string                        `json:"ImageId" xml:"ImageId"`
	DeleteAutoSnapshot            bool                          `json:"DeleteAutoSnapshot" xml:"DeleteAutoSnapshot"`
	StorageClusterId              string                        `json:"StorageClusterId" xml:"StorageClusterId"`
	AutoSnapshotPolicyId          string                        `json:"AutoSnapshotPolicyId" xml:"AutoSnapshotPolicyId"`
	ProvisionedIops               int64                         `json:"ProvisionedIops" xml:"ProvisionedIops"`
	EnableAutomatedSnapshotPolicy bool                          `json:"EnableAutomatedSnapshotPolicy" xml:"EnableAutomatedSnapshotPolicy"`
	SerialNumber                  string                        `json:"SerialNumber" xml:"SerialNumber"`
	DiskId                        string                        `json:"DiskId" xml:"DiskId"`
	Throughput                    int                           `json:"Throughput" xml:"Throughput"`
	Size                          int                           `json:"Size" xml:"Size"`
	IOPS                          int                           `json:"IOPS" xml:"IOPS"`
	MountInstanceNum              int                           `json:"MountInstanceNum" xml:"MountInstanceNum"`
	RegionId                      string                        `json:"RegionId" xml:"RegionId"`
	StorageSetId                  string                        `json:"StorageSetId" xml:"StorageSetId"`
	ResourceGroupId               string                        `json:"ResourceGroupId" xml:"ResourceGroupId"`
	InstanceId                    string                        `json:"InstanceId" xml:"InstanceId"`
	Description                   string                        `json:"Description" xml:"Description"`
	Type                          string                        `json:"Type" xml:"Type"`
	ExpiredTime                   string                        `json:"ExpiredTime" xml:"ExpiredTime"`
	Device                        string                        `json:"Device" xml:"Device"`
	MultiAttach                   string                        `json:"MultiAttach" xml:"MultiAttach"`
	BurstingEnabled               bool                          `json:"BurstingEnabled" xml:"BurstingEnabled"`
	CreationTime                  string                        `json:"CreationTime" xml:"CreationTime"`
	IOPSRead                      int                           `json:"IOPSRead" xml:"IOPSRead"`
	SourceSnapshotId              string                        `json:"SourceSnapshotId" xml:"SourceSnapshotId"`
	StorageSetPartitionNumber     int                           `json:"StorageSetPartitionNumber" xml:"StorageSetPartitionNumber"`
	Portable                      bool                          `json:"Portable" xml:"Portable"`
	KMSKeyId                      string                        `json:"KMSKeyId" xml:"KMSKeyId"`
	ProductCode                   string                        `json:"ProductCode" xml:"ProductCode"`
	Encrypted                     bool                          `json:"Encrypted" xml:"Encrypted"`
	EnableAutoSnapshot            bool                          `json:"EnableAutoSnapshot" xml:"EnableAutoSnapshot"`
	DetachedTime                  string                        `json:"DetachedTime" xml:"DetachedTime"`
	DeleteWithInstance            bool                          `json:"DeleteWithInstance" xml:"DeleteWithInstance"`
	DiskChargeType                string                        `json:"DiskChargeType" xml:"DiskChargeType"`
	ZoneId                        string                        `json:"ZoneId" xml:"ZoneId"`
	PerformanceLevel              string                        `json:"PerformanceLevel" xml:"PerformanceLevel"`
	Status                        string                        `json:"Status" xml:"Status"`
	DiskName                      string                        `json:"DiskName" xml:"DiskName"`
	IOPSWrite                     int                           `json:"IOPSWrite" xml:"IOPSWrite"`
	AttachedTime                  string                        `json:"AttachedTime" xml:"AttachedTime"`
	Attachments                   Attachments                   `json:"Attachments" xml:"Attachments"`
	Tags                          TagsInDescribeDisks           `json:"Tags" xml:"Tags"`
	MountInstances                MountInstances                `json:"MountInstances" xml:"MountInstances"`
	OperationLocks                OperationLocksInDescribeDisks `json:"OperationLocks" xml:"OperationLocks"`
}

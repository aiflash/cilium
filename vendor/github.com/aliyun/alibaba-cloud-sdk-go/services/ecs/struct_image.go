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

// Image is a nested struct in ecs response
type Image struct {
	ImageId                 string                                      `json:"ImageId" xml:"ImageId"`
	ImageOwnerAlias         string                                      `json:"ImageOwnerAlias" xml:"ImageOwnerAlias"`
	OSName                  string                                      `json:"OSName" xml:"OSName"`
	OSNameEn                string                                      `json:"OSNameEn" xml:"OSNameEn"`
	ImageFamily             string                                      `json:"ImageFamily" xml:"ImageFamily"`
	Architecture            string                                      `json:"Architecture" xml:"Architecture"`
	IsSupportIoOptimized    bool                                        `json:"IsSupportIoOptimized" xml:"IsSupportIoOptimized"`
	Size                    int                                         `json:"Size" xml:"Size"`
	ResourceGroupId         string                                      `json:"ResourceGroupId" xml:"ResourceGroupId"`
	SupplierName            string                                      `json:"SupplierName" xml:"SupplierName"`
	Description             string                                      `json:"Description" xml:"Description"`
	Usage                   string                                      `json:"Usage" xml:"Usage"`
	IsCopied                bool                                        `json:"IsCopied" xml:"IsCopied"`
	LoginAsNonRootSupported bool                                        `json:"LoginAsNonRootSupported" xml:"LoginAsNonRootSupported"`
	ImageVersion            string                                      `json:"ImageVersion" xml:"ImageVersion"`
	OSType                  string                                      `json:"OSType" xml:"OSType"`
	IsSubscribed            bool                                        `json:"IsSubscribed" xml:"IsSubscribed"`
	IsSupportCloudinit      bool                                        `json:"IsSupportCloudinit" xml:"IsSupportCloudinit"`
	CreationTime            string                                      `json:"CreationTime" xml:"CreationTime"`
	ProductCode             string                                      `json:"ProductCode" xml:"ProductCode"`
	Progress                string                                      `json:"Progress" xml:"Progress"`
	Platform                string                                      `json:"Platform" xml:"Platform"`
	IsSelfShared            string                                      `json:"IsSelfShared" xml:"IsSelfShared"`
	ImageName               string                                      `json:"ImageName" xml:"ImageName"`
	Status                  string                                      `json:"Status" xml:"Status"`
	ImageOwnerId            int64                                       `json:"ImageOwnerId" xml:"ImageOwnerId"`
	IsPublic                bool                                        `json:"IsPublic" xml:"IsPublic"`
	Tags                    TagsInDescribeImageFromFamily               `json:"Tags" xml:"Tags"`
	DiskDeviceMappings      DiskDeviceMappingsInDescribeImageFromFamily `json:"DiskDeviceMappings" xml:"DiskDeviceMappings"`
}

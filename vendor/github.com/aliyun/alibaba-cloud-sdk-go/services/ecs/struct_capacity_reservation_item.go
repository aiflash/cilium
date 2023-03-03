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

// CapacityReservationItem is a nested struct in ecs response
type CapacityReservationItem struct {
	PrivatePoolOptionsName          string                                           `json:"PrivatePoolOptionsName" xml:"PrivatePoolOptionsName"`
	EndTimeType                     string                                           `json:"EndTimeType" xml:"EndTimeType"`
	PrivatePoolOptionsMatchCriteria string                                           `json:"PrivatePoolOptionsMatchCriteria" xml:"PrivatePoolOptionsMatchCriteria"`
	TimeSlot                        string                                           `json:"TimeSlot" xml:"TimeSlot"`
	ReservedInstanceId              string                                           `json:"ReservedInstanceId" xml:"ReservedInstanceId"`
	Platform                        string                                           `json:"Platform" xml:"Platform"`
	RegionId                        string                                           `json:"RegionId" xml:"RegionId"`
	StartTime                       string                                           `json:"StartTime" xml:"StartTime"`
	EndTime                         string                                           `json:"EndTime" xml:"EndTime"`
	ResourceGroupId                 string                                           `json:"ResourceGroupId" xml:"ResourceGroupId"`
	SavingPlanId                    string                                           `json:"SavingPlanId" xml:"SavingPlanId"`
	StartTimeType                   string                                           `json:"StartTimeType" xml:"StartTimeType"`
	Status                          string                                           `json:"Status" xml:"Status"`
	InstanceChargeType              string                                           `json:"InstanceChargeType" xml:"InstanceChargeType"`
	Description                     string                                           `json:"Description" xml:"Description"`
	PrivatePoolOptionsId            string                                           `json:"PrivatePoolOptionsId" xml:"PrivatePoolOptionsId"`
	AllocatedResources              AllocatedResourcesInDescribeCapacityReservations `json:"AllocatedResources" xml:"AllocatedResources"`
	Tags                            TagsInDescribeCapacityReservations               `json:"Tags" xml:"Tags"`
}

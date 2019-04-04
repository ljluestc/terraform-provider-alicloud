package ess

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

// ScheduledTask is a nested struct in ess response
type ScheduledTask struct {
	ScheduledTaskId      string `json:"ScheduledTaskId" xml:"ScheduledTaskId"`
	ScheduledTaskName    string `json:"ScheduledTaskName" xml:"ScheduledTaskName"`
	Description          string `json:"Description" xml:"Description"`
	ScheduledAction      string `json:"ScheduledAction" xml:"ScheduledAction"`
	RecurrenceEndTime    string `json:"RecurrenceEndTime" xml:"RecurrenceEndTime"`
	LaunchTime           string `json:"LaunchTime" xml:"LaunchTime"`
	RecurrenceType       string `json:"RecurrenceType" xml:"RecurrenceType"`
	RecurrenceValue      string `json:"RecurrenceValue" xml:"RecurrenceValue"`
	LaunchExpirationTime int    `json:"LaunchExpirationTime" xml:"LaunchExpirationTime"`
	TaskEnabled          bool   `json:"TaskEnabled" xml:"TaskEnabled"`
}

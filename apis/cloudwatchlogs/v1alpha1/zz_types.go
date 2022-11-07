/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
)

// +kubebuilder:skipversion
type Destination struct {
	ARN *string `json:"arn,omitempty"`

	CreationTime *int64 `json:"creationTime,omitempty"`
}

// +kubebuilder:skipversion
type ExportTask struct {
	From *int64 `json:"from,omitempty"`

	LogGroupName *string `json:"logGroupName,omitempty"`

	To *int64 `json:"to,omitempty"`
}

// +kubebuilder:skipversion
type ExportTaskExecutionInfo struct {
	CompletionTime *int64 `json:"completionTime,omitempty"`

	CreationTime *int64 `json:"creationTime,omitempty"`
}

// +kubebuilder:skipversion
type FilteredLogEvent struct {
	IngestionTime *int64 `json:"ingestionTime,omitempty"`

	Timestamp *int64 `json:"timestamp,omitempty"`
}

// +kubebuilder:skipversion
type InputLogEvent struct {
	Timestamp *int64 `json:"timestamp,omitempty"`
}

// +kubebuilder:skipversion
type LogGroup_SDK struct {
	ARN *string `json:"arn,omitempty"`

	CreationTime *int64 `json:"creationTime,omitempty"`

	KMSKeyID *string `json:"kmsKeyID,omitempty"`

	LogGroupName *string `json:"logGroupName,omitempty"`

	MetricFilterCount *int64 `json:"metricFilterCount,omitempty"`
	// The number of days to retain the log events in the specified log group. Possible
	// values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731,
	// 1827, and 3653.
	//
	// To set a log group to never have log events expire, use DeleteRetentionPolicy
	// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DeleteRetentionPolicy.html).
	RetentionInDays *int64 `json:"retentionInDays,omitempty"`

	StoredBytes *int64 `json:"storedBytes,omitempty"`
}

// +kubebuilder:skipversion
type LogStream struct {
	ARN *string `json:"arn,omitempty"`

	CreationTime *int64 `json:"creationTime,omitempty"`

	FirstEventTimestamp *int64 `json:"firstEventTimestamp,omitempty"`

	LastEventTimestamp *int64 `json:"lastEventTimestamp,omitempty"`

	LastIngestionTime *int64 `json:"lastIngestionTime,omitempty"`

	StoredBytes *int64 `json:"storedBytes,omitempty"`
}

// +kubebuilder:skipversion
type MetricFilter struct {
	CreationTime *int64 `json:"creationTime,omitempty"`

	LogGroupName *string `json:"logGroupName,omitempty"`
}

// +kubebuilder:skipversion
type OutputLogEvent struct {
	IngestionTime *int64 `json:"ingestionTime,omitempty"`

	Timestamp *int64 `json:"timestamp,omitempty"`
}

// +kubebuilder:skipversion
type QueryDefinition struct {
	LastModified *int64 `json:"lastModified,omitempty"`
}

// +kubebuilder:skipversion
type QueryInfo struct {
	CreateTime *int64 `json:"createTime,omitempty"`

	LogGroupName *string `json:"logGroupName,omitempty"`
}

// +kubebuilder:skipversion
type ResourcePolicy struct {
	LastUpdatedTime *int64 `json:"lastUpdatedTime,omitempty"`
}

// +kubebuilder:skipversion
type SubscriptionFilter struct {
	CreationTime *int64 `json:"creationTime,omitempty"`

	LogGroupName *string `json:"logGroupName,omitempty"`
}

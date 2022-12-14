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
type AWSECRContainerImageDetails struct {
	Registry *string `json:"registry,omitempty"`

	RepositoryName *string `json:"repositoryName,omitempty"`
}

// +kubebuilder:skipversion
type EnhancedImageScanFinding struct {
	AWSAccountID *string `json:"awsAccountID,omitempty"`
}

// +kubebuilder:skipversion
type ImageReplicationStatus struct {
	RegistryID *string `json:"registryID,omitempty"`
}

// +kubebuilder:skipversion
type PullThroughCacheRule struct {
	RegistryID *string `json:"registryID,omitempty"`
}

// +kubebuilder:skipversion
type RepositoryScanningConfiguration struct {
	RepositoryName *string `json:"repositoryName,omitempty"`
}

// +kubebuilder:skipversion
type RepositoryScanningConfigurationFailure struct {
	RepositoryName *string `json:"repositoryName,omitempty"`
}

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

package secret

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/secretsmanager"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/secretsmanager/v1beta1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateDescribeSecretInput returns input for read
// operation.
func GenerateDescribeSecretInput(cr *svcapitypes.Secret) *svcsdk.DescribeSecretInput {
	res := &svcsdk.DescribeSecretInput{}

	return res
}

// GenerateSecret returns the current state in the form of *svcapitypes.Secret.
func GenerateSecret(resp *svcsdk.DescribeSecretOutput) *svcapitypes.Secret {
	cr := &svcapitypes.Secret{}

	if resp.ARN != nil {
		cr.Status.AtProvider.ARN = resp.ARN
	} else {
		cr.Status.AtProvider.ARN = nil
	}
	if resp.Description != nil {
		cr.Spec.ForProvider.Description = resp.Description
	} else {
		cr.Spec.ForProvider.Description = nil
	}
	if resp.KmsKeyId != nil {
		cr.Spec.ForProvider.KMSKeyID = resp.KmsKeyId
	} else {
		cr.Spec.ForProvider.KMSKeyID = nil
	}
	if resp.ReplicationStatus != nil {
		f10 := []*svcapitypes.ReplicationStatusType{}
		for _, f10iter := range resp.ReplicationStatus {
			f10elem := &svcapitypes.ReplicationStatusType{}
			if f10iter.KmsKeyId != nil {
				f10elem.KMSKeyID = f10iter.KmsKeyId
			}
			if f10iter.LastAccessedDate != nil {
				f10elem.LastAccessedDate = &metav1.Time{*f10iter.LastAccessedDate}
			}
			if f10iter.Region != nil {
				f10elem.Region = f10iter.Region
			}
			if f10iter.Status != nil {
				f10elem.Status = f10iter.Status
			}
			if f10iter.StatusMessage != nil {
				f10elem.StatusMessage = f10iter.StatusMessage
			}
			f10 = append(f10, f10elem)
		}
		cr.Status.AtProvider.ReplicationStatus = f10
	} else {
		cr.Status.AtProvider.ReplicationStatus = nil
	}
	if resp.Tags != nil {
		f14 := []*svcapitypes.Tag{}
		for _, f14iter := range resp.Tags {
			f14elem := &svcapitypes.Tag{}
			if f14iter.Key != nil {
				f14elem.Key = f14iter.Key
			}
			if f14iter.Value != nil {
				f14elem.Value = f14iter.Value
			}
			f14 = append(f14, f14elem)
		}
		cr.Spec.ForProvider.Tags = f14
	} else {
		cr.Spec.ForProvider.Tags = nil
	}

	return cr
}

// GenerateCreateSecretInput returns a create input.
func GenerateCreateSecretInput(cr *svcapitypes.Secret) *svcsdk.CreateSecretInput {
	res := &svcsdk.CreateSecretInput{}

	if cr.Spec.ForProvider.AddReplicaRegions != nil {
		f0 := []*svcsdk.ReplicaRegionType{}
		for _, f0iter := range cr.Spec.ForProvider.AddReplicaRegions {
			f0elem := &svcsdk.ReplicaRegionType{}
			if f0iter.KMSKeyID != nil {
				f0elem.SetKmsKeyId(*f0iter.KMSKeyID)
			}
			if f0iter.Region != nil {
				f0elem.SetRegion(*f0iter.Region)
			}
			f0 = append(f0, f0elem)
		}
		res.SetAddReplicaRegions(f0)
	}
	if cr.Spec.ForProvider.Description != nil {
		res.SetDescription(*cr.Spec.ForProvider.Description)
	}
	if cr.Spec.ForProvider.ForceOverwriteReplicaSecret != nil {
		res.SetForceOverwriteReplicaSecret(*cr.Spec.ForProvider.ForceOverwriteReplicaSecret)
	}
	if cr.Spec.ForProvider.KMSKeyID != nil {
		res.SetKmsKeyId(*cr.Spec.ForProvider.KMSKeyID)
	}
	if cr.Spec.ForProvider.Tags != nil {
		f4 := []*svcsdk.Tag{}
		for _, f4iter := range cr.Spec.ForProvider.Tags {
			f4elem := &svcsdk.Tag{}
			if f4iter.Key != nil {
				f4elem.SetKey(*f4iter.Key)
			}
			if f4iter.Value != nil {
				f4elem.SetValue(*f4iter.Value)
			}
			f4 = append(f4, f4elem)
		}
		res.SetTags(f4)
	}

	return res
}

// GenerateUpdateSecretInput returns an update input.
func GenerateUpdateSecretInput(cr *svcapitypes.Secret) *svcsdk.UpdateSecretInput {
	res := &svcsdk.UpdateSecretInput{}

	if cr.Spec.ForProvider.Description != nil {
		res.SetDescription(*cr.Spec.ForProvider.Description)
	}
	if cr.Spec.ForProvider.KMSKeyID != nil {
		res.SetKmsKeyId(*cr.Spec.ForProvider.KMSKeyID)
	}

	return res
}

// GenerateDeleteSecretInput returns a deletion input.
func GenerateDeleteSecretInput(cr *svcapitypes.Secret) *svcsdk.DeleteSecretInput {
	res := &svcsdk.DeleteSecretInput{}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "ResourceNotFoundException"
}
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

package deployment

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/apigateway"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/apigateway/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateGetDeploymentInput returns input for read
// operation.
func GenerateGetDeploymentInput(cr *svcapitypes.Deployment) *svcsdk.GetDeploymentInput {
	res := &svcsdk.GetDeploymentInput{}

	return res
}

// GenerateDeployment returns the current state in the form of *svcapitypes.Deployment.
func GenerateDeployment(resp *svcsdk.Deployment) *svcapitypes.Deployment {
	cr := &svcapitypes.Deployment{}

	if resp.ApiSummary != nil {
		f0 := map[string]map[string]*svcapitypes.MethodSnapshot{}
		for f0key, f0valiter := range resp.ApiSummary {
			f0val := map[string]*svcapitypes.MethodSnapshot{}
			for f0valkey, f0valvaliter := range f0valiter {
				f0valval := &svcapitypes.MethodSnapshot{}
				if f0valvaliter.ApiKeyRequired != nil {
					f0valval.APIKeyRequired = f0valvaliter.ApiKeyRequired
				}
				if f0valvaliter.AuthorizationType != nil {
					f0valval.AuthorizationType = f0valvaliter.AuthorizationType
				}
				f0val[f0valkey] = f0valval
			}
			f0[f0key] = f0val
		}
		cr.Status.AtProvider.APISummary = f0
	} else {
		cr.Status.AtProvider.APISummary = nil
	}
	if resp.CreatedDate != nil {
		cr.Status.AtProvider.CreatedDate = &metav1.Time{*resp.CreatedDate}
	} else {
		cr.Status.AtProvider.CreatedDate = nil
	}
	if resp.Description != nil {
		cr.Spec.ForProvider.Description = resp.Description
	} else {
		cr.Spec.ForProvider.Description = nil
	}
	if resp.Id != nil {
		cr.Status.AtProvider.ID = resp.Id
	} else {
		cr.Status.AtProvider.ID = nil
	}

	return cr
}

// GenerateCreateDeploymentInput returns a create input.
func GenerateCreateDeploymentInput(cr *svcapitypes.Deployment) *svcsdk.CreateDeploymentInput {
	res := &svcsdk.CreateDeploymentInput{}

	if cr.Spec.ForProvider.CacheClusterEnabled != nil {
		res.SetCacheClusterEnabled(*cr.Spec.ForProvider.CacheClusterEnabled)
	}
	if cr.Spec.ForProvider.CacheClusterSize != nil {
		res.SetCacheClusterSize(*cr.Spec.ForProvider.CacheClusterSize)
	}
	if cr.Spec.ForProvider.CanarySettings != nil {
		f2 := &svcsdk.DeploymentCanarySettings{}
		if cr.Spec.ForProvider.CanarySettings.PercentTraffic != nil {
			f2.SetPercentTraffic(*cr.Spec.ForProvider.CanarySettings.PercentTraffic)
		}
		if cr.Spec.ForProvider.CanarySettings.StageVariableOverrides != nil {
			f2f1 := map[string]*string{}
			for f2f1key, f2f1valiter := range cr.Spec.ForProvider.CanarySettings.StageVariableOverrides {
				var f2f1val string
				f2f1val = *f2f1valiter
				f2f1[f2f1key] = &f2f1val
			}
			f2.SetStageVariableOverrides(f2f1)
		}
		if cr.Spec.ForProvider.CanarySettings.UseStageCache != nil {
			f2.SetUseStageCache(*cr.Spec.ForProvider.CanarySettings.UseStageCache)
		}
		res.SetCanarySettings(f2)
	}
	if cr.Spec.ForProvider.Description != nil {
		res.SetDescription(*cr.Spec.ForProvider.Description)
	}
	if cr.Spec.ForProvider.StageDescription != nil {
		res.SetStageDescription(*cr.Spec.ForProvider.StageDescription)
	}
	if cr.Spec.ForProvider.StageName != nil {
		res.SetStageName(*cr.Spec.ForProvider.StageName)
	}
	if cr.Spec.ForProvider.TracingEnabled != nil {
		res.SetTracingEnabled(*cr.Spec.ForProvider.TracingEnabled)
	}
	if cr.Spec.ForProvider.Variables != nil {
		f7 := map[string]*string{}
		for f7key, f7valiter := range cr.Spec.ForProvider.Variables {
			var f7val string
			f7val = *f7valiter
			f7[f7key] = &f7val
		}
		res.SetVariables(f7)
	}

	return res
}

// GenerateUpdateDeploymentInput returns an update input.
func GenerateUpdateDeploymentInput(cr *svcapitypes.Deployment) *svcsdk.UpdateDeploymentInput {
	res := &svcsdk.UpdateDeploymentInput{}

	return res
}

// GenerateDeleteDeploymentInput returns a deletion input.
func GenerateDeleteDeploymentInput(cr *svcapitypes.Deployment) *svcsdk.DeleteDeploymentInput {
	res := &svcsdk.DeleteDeploymentInput{}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "NotFoundException"
}

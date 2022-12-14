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

// Code generated by angryjet. DO NOT EDIT.

package manualv1alpha1

import (
	"context"
	v1alpha1 "github.com/crossplane-contrib/provider-aws/apis/batch/v1alpha1"
	v1alpha11 "github.com/crossplane-contrib/provider-aws/apis/efs/v1alpha1"
	v1beta1 "github.com/crossplane-contrib/provider-aws/apis/iam/v1beta1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Job.
func (mg *Job) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.DependsOn); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DependsOn[i3].JobID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.DependsOn[i3].JobIDRef,
			Selector:     mg.Spec.ForProvider.DependsOn[i3].JobIDSelector,
			To: reference.To{
				List:    &JobList{},
				Managed: &Job{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.DependsOn[i3].JobID")
		}
		mg.Spec.ForProvider.DependsOn[i3].JobID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.DependsOn[i3].JobIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.ForProvider.JobDefinition,
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.JobDefinitionRef,
		Selector:     mg.Spec.ForProvider.JobDefinitionSelector,
		To: reference.To{
			List:    &JobDefinitionList{},
			Managed: &JobDefinition{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.JobDefinition")
	}
	mg.Spec.ForProvider.JobDefinition = rsp.ResolvedValue
	mg.Spec.ForProvider.JobDefinitionRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: mg.Spec.ForProvider.JobQueue,
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.JobQueueRef,
		Selector:     mg.Spec.ForProvider.JobQueueSelector,
		To: reference.To{
			List:    &v1alpha1.JobQueueList{},
			Managed: &v1alpha1.JobQueue{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.JobQueue")
	}
	mg.Spec.ForProvider.JobQueue = rsp.ResolvedValue
	mg.Spec.ForProvider.JobQueueRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this JobDefinition.
func (mg *JobDefinition) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.ContainerProperties != nil {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ContainerProperties.ExecutionRoleArn),
			Extract:      v1beta1.RoleARN(),
			Reference:    mg.Spec.ForProvider.ContainerProperties.ExecutionRoleARNRef,
			Selector:     mg.Spec.ForProvider.ContainerProperties.ExecutionRoleARNSelector,
			To: reference.To{
				List:    &v1beta1.RoleList{},
				Managed: &v1beta1.Role{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.ContainerProperties.ExecutionRoleArn")
		}
		mg.Spec.ForProvider.ContainerProperties.ExecutionRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.ContainerProperties.ExecutionRoleARNRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.ContainerProperties != nil {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ContainerProperties.JobRoleArn),
			Extract:      v1beta1.RoleARN(),
			Reference:    mg.Spec.ForProvider.ContainerProperties.JobRoleARNRef,
			Selector:     mg.Spec.ForProvider.ContainerProperties.JobRoleARNSelector,
			To: reference.To{
				List:    &v1beta1.RoleList{},
				Managed: &v1beta1.Role{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.ContainerProperties.JobRoleArn")
		}
		mg.Spec.ForProvider.ContainerProperties.JobRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.ContainerProperties.JobRoleARNRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.ContainerProperties != nil {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ContainerProperties.Volumes); i4++ {
			if mg.Spec.ForProvider.ContainerProperties.Volumes[i4].EfsVolumeConfiguration != nil {
				if mg.Spec.ForProvider.ContainerProperties.Volumes[i4].EfsVolumeConfiguration.AuthorizationConfig != nil {
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ContainerProperties.Volumes[i4].EfsVolumeConfiguration.AuthorizationConfig.AccessPointID),
						Extract:      reference.ExternalName(),
						Reference:    mg.Spec.ForProvider.ContainerProperties.Volumes[i4].EfsVolumeConfiguration.AuthorizationConfig.AccessPointIDRef,
						Selector:     mg.Spec.ForProvider.ContainerProperties.Volumes[i4].EfsVolumeConfiguration.AuthorizationConfig.AccessPointIDSelector,
						To: reference.To{
							List:    &v1alpha11.AccessPointList{},
							Managed: &v1alpha11.AccessPoint{},
						},
					})
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.ContainerProperties.Volumes[i4].EfsVolumeConfiguration.AuthorizationConfig.AccessPointID")
					}
					mg.Spec.ForProvider.ContainerProperties.Volumes[i4].EfsVolumeConfiguration.AuthorizationConfig.AccessPointID = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.ContainerProperties.Volumes[i4].EfsVolumeConfiguration.AuthorizationConfig.AccessPointIDRef = rsp.ResolvedReference

				}
			}
		}
	}
	if mg.Spec.ForProvider.ContainerProperties != nil {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.ContainerProperties.Volumes); i4++ {
			if mg.Spec.ForProvider.ContainerProperties.Volumes[i4].EfsVolumeConfiguration != nil {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: mg.Spec.ForProvider.ContainerProperties.Volumes[i4].EfsVolumeConfiguration.FileSystemID,
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.ForProvider.ContainerProperties.Volumes[i4].EfsVolumeConfiguration.FileSystemIDRef,
					Selector:     mg.Spec.ForProvider.ContainerProperties.Volumes[i4].EfsVolumeConfiguration.FileSystemIDSelector,
					To: reference.To{
						List:    &v1alpha11.FileSystemList{},
						Managed: &v1alpha11.FileSystem{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.ContainerProperties.Volumes[i4].EfsVolumeConfiguration.FileSystemID")
				}
				mg.Spec.ForProvider.ContainerProperties.Volumes[i4].EfsVolumeConfiguration.FileSystemID = rsp.ResolvedValue
				mg.Spec.ForProvider.ContainerProperties.Volumes[i4].EfsVolumeConfiguration.FileSystemIDRef = rsp.ResolvedReference

			}
		}
	}
	if mg.Spec.ForProvider.NodeProperties != nil {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.NodeProperties.NodeRangeProperties); i4++ {
			if mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container != nil {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container.ExecutionRoleArn),
					Extract:      v1beta1.RoleARN(),
					Reference:    mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container.ExecutionRoleARNRef,
					Selector:     mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container.ExecutionRoleARNSelector,
					To: reference.To{
						List:    &v1beta1.RoleList{},
						Managed: &v1beta1.Role{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container.ExecutionRoleArn")
				}
				mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container.ExecutionRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container.ExecutionRoleARNRef = rsp.ResolvedReference

			}
		}
	}
	if mg.Spec.ForProvider.NodeProperties != nil {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.NodeProperties.NodeRangeProperties); i4++ {
			if mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container != nil {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container.JobRoleArn),
					Extract:      v1beta1.RoleARN(),
					Reference:    mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container.JobRoleARNRef,
					Selector:     mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container.JobRoleARNSelector,
					To: reference.To{
						List:    &v1beta1.RoleList{},
						Managed: &v1beta1.Role{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container.JobRoleArn")
				}
				mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container.JobRoleArn = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container.JobRoleARNRef = rsp.ResolvedReference

			}
		}
	}
	if mg.Spec.ForProvider.NodeProperties != nil {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.NodeProperties.NodeRangeProperties); i4++ {
			if mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container != nil {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container.Volumes); i6++ {
					if mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container.Volumes[i6].EfsVolumeConfiguration != nil {
						if mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container.Volumes[i6].EfsVolumeConfiguration.AuthorizationConfig != nil {
							rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
								CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container.Volumes[i6].EfsVolumeConfiguration.AuthorizationConfig.AccessPointID),
								Extract:      reference.ExternalName(),
								Reference:    mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container.Volumes[i6].EfsVolumeConfiguration.AuthorizationConfig.AccessPointIDRef,
								Selector:     mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container.Volumes[i6].EfsVolumeConfiguration.AuthorizationConfig.AccessPointIDSelector,
								To: reference.To{
									List:    &v1alpha11.AccessPointList{},
									Managed: &v1alpha11.AccessPoint{},
								},
							})
							if err != nil {
								return errors.Wrap(err, "mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container.Volumes[i6].EfsVolumeConfiguration.AuthorizationConfig.AccessPointID")
							}
							mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container.Volumes[i6].EfsVolumeConfiguration.AuthorizationConfig.AccessPointID = reference.ToPtrValue(rsp.ResolvedValue)
							mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container.Volumes[i6].EfsVolumeConfiguration.AuthorizationConfig.AccessPointIDRef = rsp.ResolvedReference

						}
					}
				}
			}
		}
	}
	if mg.Spec.ForProvider.NodeProperties != nil {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.NodeProperties.NodeRangeProperties); i4++ {
			if mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container != nil {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container.Volumes); i6++ {
					if mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container.Volumes[i6].EfsVolumeConfiguration != nil {
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container.Volumes[i6].EfsVolumeConfiguration.FileSystemID,
							Extract:      reference.ExternalName(),
							Reference:    mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container.Volumes[i6].EfsVolumeConfiguration.FileSystemIDRef,
							Selector:     mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container.Volumes[i6].EfsVolumeConfiguration.FileSystemIDSelector,
							To: reference.To{
								List:    &v1alpha11.FileSystemList{},
								Managed: &v1alpha11.FileSystem{},
							},
						})
						if err != nil {
							return errors.Wrap(err, "mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container.Volumes[i6].EfsVolumeConfiguration.FileSystemID")
						}
						mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container.Volumes[i6].EfsVolumeConfiguration.FileSystemID = rsp.ResolvedValue
						mg.Spec.ForProvider.NodeProperties.NodeRangeProperties[i4].Container.Volumes[i6].EfsVolumeConfiguration.FileSystemIDRef = rsp.ResolvedReference

					}
				}
			}
		}
	}

	return nil
}

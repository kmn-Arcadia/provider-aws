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

package targetgroup

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/elbv2"
	svcsdk "github.com/aws/aws-sdk-go/service/elbv2"
	svcsdkapi "github.com/aws/aws-sdk-go/service/elbv2/elbv2iface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/elbv2/v1alpha1"
	awsclient "github.com/crossplane-contrib/provider-aws/pkg/clients"
)

const (
	errUnexpectedObject = "managed resource is not an TargetGroup resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create TargetGroup in AWS"
	errUpdate        = "cannot update TargetGroup in AWS"
	errDescribe      = "failed to describe TargetGroup"
	errDelete        = "failed to delete TargetGroup"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.TargetGroup)
	if !ok {
		return nil, errors.New(errUnexpectedObject)
	}
	sess, err := awsclient.GetConfigV1(ctx, c.kube, mg, cr.Spec.ForProvider.Region)
	if err != nil {
		return nil, errors.Wrap(err, errCreateSession)
	}
	return newExternal(c.kube, svcapi.New(sess), c.opts), nil
}

func (e *external) Observe(ctx context.Context, mg cpresource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*svcapitypes.TargetGroup)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateDescribeTargetGroupsInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.DescribeTargetGroupsWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	resp = e.filterList(cr, resp)
	if len(resp.TargetGroups) == 0 {
		return managed.ExternalObservation{ResourceExists: false}, nil
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateTargetGroup(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)

	upToDate, err := e.isUpToDate(cr, resp)
	if err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "isUpToDate check failed")
	}
	return e.postObserve(ctx, cr, resp, managed.ExternalObservation{
		ResourceExists:          true,
		ResourceUpToDate:        upToDate,
		ResourceLateInitialized: !cmp.Equal(&cr.Spec.ForProvider, currentSpec),
	}, nil)
}

func (e *external) Create(ctx context.Context, mg cpresource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*svcapitypes.TargetGroup)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateTargetGroupInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateTargetGroupWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, awsclient.Wrap(err, errCreate)
	}

	if resp.TargetGroups != nil {
		f0 := []*svcapitypes.TargetGroup_SDK{}
		for _, f0iter := range resp.TargetGroups {
			f0elem := &svcapitypes.TargetGroup_SDK{}
			if f0iter.HealthCheckEnabled != nil {
				f0elem.HealthCheckEnabled = f0iter.HealthCheckEnabled
			}
			if f0iter.HealthCheckIntervalSeconds != nil {
				f0elem.HealthCheckIntervalSeconds = f0iter.HealthCheckIntervalSeconds
			}
			if f0iter.HealthCheckPath != nil {
				f0elem.HealthCheckPath = f0iter.HealthCheckPath
			}
			if f0iter.HealthCheckPort != nil {
				f0elem.HealthCheckPort = f0iter.HealthCheckPort
			}
			if f0iter.HealthCheckProtocol != nil {
				f0elem.HealthCheckProtocol = f0iter.HealthCheckProtocol
			}
			if f0iter.HealthCheckTimeoutSeconds != nil {
				f0elem.HealthCheckTimeoutSeconds = f0iter.HealthCheckTimeoutSeconds
			}
			if f0iter.HealthyThresholdCount != nil {
				f0elem.HealthyThresholdCount = f0iter.HealthyThresholdCount
			}
			if f0iter.IpAddressType != nil {
				f0elem.IPAddressType = f0iter.IpAddressType
			}
			if f0iter.LoadBalancerArns != nil {
				f0elemf8 := []*string{}
				for _, f0elemf8iter := range f0iter.LoadBalancerArns {
					var f0elemf8elem string
					f0elemf8elem = *f0elemf8iter
					f0elemf8 = append(f0elemf8, &f0elemf8elem)
				}
				f0elem.LoadBalancerARNs = f0elemf8
			}
			if f0iter.Matcher != nil {
				f0elemf9 := &svcapitypes.Matcher{}
				if f0iter.Matcher.GrpcCode != nil {
					f0elemf9.GrpcCode = f0iter.Matcher.GrpcCode
				}
				if f0iter.Matcher.HttpCode != nil {
					f0elemf9.HTTPCode = f0iter.Matcher.HttpCode
				}
				f0elem.Matcher = f0elemf9
			}
			if f0iter.Port != nil {
				f0elem.Port = f0iter.Port
			}
			if f0iter.Protocol != nil {
				f0elem.Protocol = f0iter.Protocol
			}
			if f0iter.ProtocolVersion != nil {
				f0elem.ProtocolVersion = f0iter.ProtocolVersion
			}
			if f0iter.TargetGroupArn != nil {
				f0elem.TargetGroupARN = f0iter.TargetGroupArn
			}
			if f0iter.TargetGroupName != nil {
				f0elem.TargetGroupName = f0iter.TargetGroupName
			}
			if f0iter.TargetType != nil {
				f0elem.TargetType = f0iter.TargetType
			}
			if f0iter.UnhealthyThresholdCount != nil {
				f0elem.UnhealthyThresholdCount = f0iter.UnhealthyThresholdCount
			}
			if f0iter.VpcId != nil {
				f0elem.VPCID = f0iter.VpcId
			}
			f0 = append(f0, f0elem)
		}
		cr.Status.AtProvider.TargetGroups = f0
	} else {
		cr.Status.AtProvider.TargetGroups = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.TargetGroup)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	input := GenerateModifyTargetGroupInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.ModifyTargetGroupWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, awsclient.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.TargetGroup)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteTargetGroupInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return nil
	}
	resp, err := e.client.DeleteTargetGroupWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.ELBV2API, opts []option) *external {
	e := &external{
		kube:           kube,
		client:         client,
		preObserve:     nopPreObserve,
		postObserve:    nopPostObserve,
		lateInitialize: nopLateInitialize,
		isUpToDate:     alwaysUpToDate,
		filterList:     nopFilterList,
		preCreate:      nopPreCreate,
		postCreate:     nopPostCreate,
		preDelete:      nopPreDelete,
		postDelete:     nopPostDelete,
		preUpdate:      nopPreUpdate,
		postUpdate:     nopPostUpdate,
	}
	for _, f := range opts {
		f(e)
	}
	return e
}

type external struct {
	kube           client.Client
	client         svcsdkapi.ELBV2API
	preObserve     func(context.Context, *svcapitypes.TargetGroup, *svcsdk.DescribeTargetGroupsInput) error
	postObserve    func(context.Context, *svcapitypes.TargetGroup, *svcsdk.DescribeTargetGroupsOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	filterList     func(*svcapitypes.TargetGroup, *svcsdk.DescribeTargetGroupsOutput) *svcsdk.DescribeTargetGroupsOutput
	lateInitialize func(*svcapitypes.TargetGroupParameters, *svcsdk.DescribeTargetGroupsOutput) error
	isUpToDate     func(*svcapitypes.TargetGroup, *svcsdk.DescribeTargetGroupsOutput) (bool, error)
	preCreate      func(context.Context, *svcapitypes.TargetGroup, *svcsdk.CreateTargetGroupInput) error
	postCreate     func(context.Context, *svcapitypes.TargetGroup, *svcsdk.CreateTargetGroupOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.TargetGroup, *svcsdk.DeleteTargetGroupInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.TargetGroup, *svcsdk.DeleteTargetGroupOutput, error) error
	preUpdate      func(context.Context, *svcapitypes.TargetGroup, *svcsdk.ModifyTargetGroupInput) error
	postUpdate     func(context.Context, *svcapitypes.TargetGroup, *svcsdk.ModifyTargetGroupOutput, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.TargetGroup, *svcsdk.DescribeTargetGroupsInput) error {
	return nil
}
func nopPostObserve(_ context.Context, _ *svcapitypes.TargetGroup, _ *svcsdk.DescribeTargetGroupsOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopFilterList(_ *svcapitypes.TargetGroup, list *svcsdk.DescribeTargetGroupsOutput) *svcsdk.DescribeTargetGroupsOutput {
	return list
}

func nopLateInitialize(*svcapitypes.TargetGroupParameters, *svcsdk.DescribeTargetGroupsOutput) error {
	return nil
}
func alwaysUpToDate(*svcapitypes.TargetGroup, *svcsdk.DescribeTargetGroupsOutput) (bool, error) {
	return true, nil
}

func nopPreCreate(context.Context, *svcapitypes.TargetGroup, *svcsdk.CreateTargetGroupInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.TargetGroup, _ *svcsdk.CreateTargetGroupOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.TargetGroup, *svcsdk.DeleteTargetGroupInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.TargetGroup, _ *svcsdk.DeleteTargetGroupOutput, err error) error {
	return err
}
func nopPreUpdate(context.Context, *svcapitypes.TargetGroup, *svcsdk.ModifyTargetGroupInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.TargetGroup, _ *svcsdk.ModifyTargetGroupOutput, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}

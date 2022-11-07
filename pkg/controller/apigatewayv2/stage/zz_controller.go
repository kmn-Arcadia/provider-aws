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

package stage

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/apigatewayv2"
	svcsdk "github.com/aws/aws-sdk-go/service/apigatewayv2"
	svcsdkapi "github.com/aws/aws-sdk-go/service/apigatewayv2/apigatewayv2iface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/apigatewayv2/v1alpha1"
	awsclient "github.com/crossplane-contrib/provider-aws/pkg/clients"
)

const (
	errUnexpectedObject = "managed resource is not an Stage resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create Stage in AWS"
	errUpdate        = "cannot update Stage in AWS"
	errDescribe      = "failed to describe Stage"
	errDelete        = "failed to delete Stage"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.Stage)
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
	cr, ok := mg.(*svcapitypes.Stage)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateGetStageInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.GetStageWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateStage(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)

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
	cr, ok := mg.(*svcapitypes.Stage)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateStageInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateStageWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, awsclient.Wrap(err, errCreate)
	}

	if resp.AccessLogSettings != nil {
		f0 := &svcapitypes.AccessLogSettings{}
		if resp.AccessLogSettings.DestinationArn != nil {
			f0.DestinationARN = resp.AccessLogSettings.DestinationArn
		}
		if resp.AccessLogSettings.Format != nil {
			f0.Format = resp.AccessLogSettings.Format
		}
		cr.Spec.ForProvider.AccessLogSettings = f0
	} else {
		cr.Spec.ForProvider.AccessLogSettings = nil
	}
	if resp.ApiGatewayManaged != nil {
		cr.Status.AtProvider.APIGatewayManaged = resp.ApiGatewayManaged
	} else {
		cr.Status.AtProvider.APIGatewayManaged = nil
	}
	if resp.AutoDeploy != nil {
		cr.Spec.ForProvider.AutoDeploy = resp.AutoDeploy
	} else {
		cr.Spec.ForProvider.AutoDeploy = nil
	}
	if resp.ClientCertificateId != nil {
		cr.Spec.ForProvider.ClientCertificateID = resp.ClientCertificateId
	} else {
		cr.Spec.ForProvider.ClientCertificateID = nil
	}
	if resp.CreatedDate != nil {
		cr.Status.AtProvider.CreatedDate = &metav1.Time{*resp.CreatedDate}
	} else {
		cr.Status.AtProvider.CreatedDate = nil
	}
	if resp.DefaultRouteSettings != nil {
		f5 := &svcapitypes.RouteSettings{}
		if resp.DefaultRouteSettings.DataTraceEnabled != nil {
			f5.DataTraceEnabled = resp.DefaultRouteSettings.DataTraceEnabled
		}
		if resp.DefaultRouteSettings.DetailedMetricsEnabled != nil {
			f5.DetailedMetricsEnabled = resp.DefaultRouteSettings.DetailedMetricsEnabled
		}
		if resp.DefaultRouteSettings.LoggingLevel != nil {
			f5.LoggingLevel = resp.DefaultRouteSettings.LoggingLevel
		}
		if resp.DefaultRouteSettings.ThrottlingBurstLimit != nil {
			f5.ThrottlingBurstLimit = resp.DefaultRouteSettings.ThrottlingBurstLimit
		}
		if resp.DefaultRouteSettings.ThrottlingRateLimit != nil {
			f5.ThrottlingRateLimit = resp.DefaultRouteSettings.ThrottlingRateLimit
		}
		cr.Spec.ForProvider.DefaultRouteSettings = f5
	} else {
		cr.Spec.ForProvider.DefaultRouteSettings = nil
	}
	if resp.DeploymentId != nil {
		cr.Spec.ForProvider.DeploymentID = resp.DeploymentId
	} else {
		cr.Spec.ForProvider.DeploymentID = nil
	}
	if resp.Description != nil {
		cr.Spec.ForProvider.Description = resp.Description
	} else {
		cr.Spec.ForProvider.Description = nil
	}
	if resp.LastDeploymentStatusMessage != nil {
		cr.Status.AtProvider.LastDeploymentStatusMessage = resp.LastDeploymentStatusMessage
	} else {
		cr.Status.AtProvider.LastDeploymentStatusMessage = nil
	}
	if resp.LastUpdatedDate != nil {
		cr.Status.AtProvider.LastUpdatedDate = &metav1.Time{*resp.LastUpdatedDate}
	} else {
		cr.Status.AtProvider.LastUpdatedDate = nil
	}
	if resp.RouteSettings != nil {
		f10 := map[string]*svcapitypes.RouteSettings{}
		for f10key, f10valiter := range resp.RouteSettings {
			f10val := &svcapitypes.RouteSettings{}
			if f10valiter.DataTraceEnabled != nil {
				f10val.DataTraceEnabled = f10valiter.DataTraceEnabled
			}
			if f10valiter.DetailedMetricsEnabled != nil {
				f10val.DetailedMetricsEnabled = f10valiter.DetailedMetricsEnabled
			}
			if f10valiter.LoggingLevel != nil {
				f10val.LoggingLevel = f10valiter.LoggingLevel
			}
			if f10valiter.ThrottlingBurstLimit != nil {
				f10val.ThrottlingBurstLimit = f10valiter.ThrottlingBurstLimit
			}
			if f10valiter.ThrottlingRateLimit != nil {
				f10val.ThrottlingRateLimit = f10valiter.ThrottlingRateLimit
			}
			f10[f10key] = f10val
		}
		cr.Spec.ForProvider.RouteSettings = f10
	} else {
		cr.Spec.ForProvider.RouteSettings = nil
	}
	if resp.StageName != nil {
		cr.Status.AtProvider.StageName = resp.StageName
	} else {
		cr.Status.AtProvider.StageName = nil
	}
	if resp.StageVariables != nil {
		f12 := map[string]*string{}
		for f12key, f12valiter := range resp.StageVariables {
			var f12val string
			f12val = *f12valiter
			f12[f12key] = &f12val
		}
		cr.Spec.ForProvider.StageVariables = f12
	} else {
		cr.Spec.ForProvider.StageVariables = nil
	}
	if resp.Tags != nil {
		f13 := map[string]*string{}
		for f13key, f13valiter := range resp.Tags {
			var f13val string
			f13val = *f13valiter
			f13[f13key] = &f13val
		}
		cr.Spec.ForProvider.Tags = f13
	} else {
		cr.Spec.ForProvider.Tags = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.Stage)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	input := GenerateUpdateStageInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.UpdateStageWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, awsclient.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.Stage)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteStageInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return nil
	}
	resp, err := e.client.DeleteStageWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.ApiGatewayV2API, opts []option) *external {
	e := &external{
		kube:           kube,
		client:         client,
		preObserve:     nopPreObserve,
		postObserve:    nopPostObserve,
		lateInitialize: nopLateInitialize,
		isUpToDate:     alwaysUpToDate,
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
	client         svcsdkapi.ApiGatewayV2API
	preObserve     func(context.Context, *svcapitypes.Stage, *svcsdk.GetStageInput) error
	postObserve    func(context.Context, *svcapitypes.Stage, *svcsdk.GetStageOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	lateInitialize func(*svcapitypes.StageParameters, *svcsdk.GetStageOutput) error
	isUpToDate     func(*svcapitypes.Stage, *svcsdk.GetStageOutput) (bool, error)
	preCreate      func(context.Context, *svcapitypes.Stage, *svcsdk.CreateStageInput) error
	postCreate     func(context.Context, *svcapitypes.Stage, *svcsdk.CreateStageOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.Stage, *svcsdk.DeleteStageInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.Stage, *svcsdk.DeleteStageOutput, error) error
	preUpdate      func(context.Context, *svcapitypes.Stage, *svcsdk.UpdateStageInput) error
	postUpdate     func(context.Context, *svcapitypes.Stage, *svcsdk.UpdateStageOutput, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.Stage, *svcsdk.GetStageInput) error {
	return nil
}

func nopPostObserve(_ context.Context, _ *svcapitypes.Stage, _ *svcsdk.GetStageOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopLateInitialize(*svcapitypes.StageParameters, *svcsdk.GetStageOutput) error {
	return nil
}
func alwaysUpToDate(*svcapitypes.Stage, *svcsdk.GetStageOutput) (bool, error) {
	return true, nil
}

func nopPreCreate(context.Context, *svcapitypes.Stage, *svcsdk.CreateStageInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.Stage, _ *svcsdk.CreateStageOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.Stage, *svcsdk.DeleteStageInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.Stage, _ *svcsdk.DeleteStageOutput, err error) error {
	return err
}
func nopPreUpdate(context.Context, *svcapitypes.Stage, *svcsdk.UpdateStageInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.Stage, _ *svcsdk.UpdateStageOutput, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}

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

package domain

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/cloudsearch"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/cloudsearch/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateDescribeDomainsInput returns input for read
// operation.
func GenerateDescribeDomainsInput(cr *svcapitypes.Domain) *svcsdk.DescribeDomainsInput {
	res := &svcsdk.DescribeDomainsInput{}

	if cr.Spec.ForProvider.DomainName != nil {
		f0 := []*string{}
		f0 = append(f0, cr.Spec.ForProvider.DomainName)
		res.SetDomainNames(f0)
	}

	return res
}

// GenerateDomain returns the current state in the form of *svcapitypes.Domain.
func GenerateDomain(resp *svcsdk.DescribeDomainsOutput) *svcapitypes.Domain {
	cr := &svcapitypes.Domain{}

	found := false
	for _, elem := range resp.DomainStatusList {
		if elem.ARN != nil {
			cr.Status.AtProvider.ARN = elem.ARN
		} else {
			cr.Status.AtProvider.ARN = nil
		}
		if elem.Created != nil {
			cr.Status.AtProvider.Created = elem.Created
		} else {
			cr.Status.AtProvider.Created = nil
		}
		if elem.Deleted != nil {
			cr.Status.AtProvider.Deleted = elem.Deleted
		} else {
			cr.Status.AtProvider.Deleted = nil
		}
		if elem.DocService != nil {
			f3 := &svcapitypes.ServiceEndpoint{}
			if elem.DocService.Endpoint != nil {
				f3.Endpoint = elem.DocService.Endpoint
			}
			cr.Status.AtProvider.DocService = f3
		} else {
			cr.Status.AtProvider.DocService = nil
		}
		if elem.DomainId != nil {
			cr.Status.AtProvider.DomainID = elem.DomainId
		} else {
			cr.Status.AtProvider.DomainID = nil
		}
		if elem.DomainName != nil {
			cr.Spec.ForProvider.DomainName = elem.DomainName
		} else {
			cr.Spec.ForProvider.DomainName = nil
		}
		if elem.Limits != nil {
			f6 := &svcapitypes.Limits{}
			if elem.Limits.MaximumPartitionCount != nil {
				f6.MaximumPartitionCount = elem.Limits.MaximumPartitionCount
			}
			if elem.Limits.MaximumReplicationCount != nil {
				f6.MaximumReplicationCount = elem.Limits.MaximumReplicationCount
			}
			cr.Status.AtProvider.Limits = f6
		} else {
			cr.Status.AtProvider.Limits = nil
		}
		if elem.Processing != nil {
			cr.Status.AtProvider.Processing = elem.Processing
		} else {
			cr.Status.AtProvider.Processing = nil
		}
		if elem.RequiresIndexDocuments != nil {
			cr.Status.AtProvider.RequiresIndexDocuments = elem.RequiresIndexDocuments
		} else {
			cr.Status.AtProvider.RequiresIndexDocuments = nil
		}
		if elem.SearchInstanceCount != nil {
			cr.Status.AtProvider.SearchInstanceCount = elem.SearchInstanceCount
		} else {
			cr.Status.AtProvider.SearchInstanceCount = nil
		}
		if elem.SearchInstanceType != nil {
			cr.Status.AtProvider.SearchInstanceType = elem.SearchInstanceType
		} else {
			cr.Status.AtProvider.SearchInstanceType = nil
		}
		if elem.SearchPartitionCount != nil {
			cr.Status.AtProvider.SearchPartitionCount = elem.SearchPartitionCount
		} else {
			cr.Status.AtProvider.SearchPartitionCount = nil
		}
		if elem.SearchService != nil {
			f12 := &svcapitypes.ServiceEndpoint{}
			if elem.SearchService.Endpoint != nil {
				f12.Endpoint = elem.SearchService.Endpoint
			}
			cr.Status.AtProvider.SearchService = f12
		} else {
			cr.Status.AtProvider.SearchService = nil
		}
		found = true
		break
	}
	if !found {
		return cr
	}

	return cr
}

// GenerateCreateDomainInput returns a create input.
func GenerateCreateDomainInput(cr *svcapitypes.Domain) *svcsdk.CreateDomainInput {
	res := &svcsdk.CreateDomainInput{}

	if cr.Spec.ForProvider.DomainName != nil {
		res.SetDomainName(*cr.Spec.ForProvider.DomainName)
	}

	return res
}

// GenerateDeleteDomainInput returns a deletion input.
func GenerateDeleteDomainInput(cr *svcapitypes.Domain) *svcsdk.DeleteDomainInput {
	res := &svcsdk.DeleteDomainInput{}

	if cr.Spec.ForProvider.DomainName != nil {
		res.SetDomainName(*cr.Spec.ForProvider.DomainName)
	}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "UNKNOWN"
}
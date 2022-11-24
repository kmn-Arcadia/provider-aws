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

package broker

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/mq"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/mq/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateDescribeBrokerInput returns input for read
// operation.
func GenerateDescribeBrokerInput(cr *svcapitypes.Broker) *svcsdk.DescribeBrokerInput {
	res := &svcsdk.DescribeBrokerInput{}

	if cr.Status.AtProvider.BrokerID != nil {
		res.SetBrokerId(*cr.Status.AtProvider.BrokerID)
	}

	return res
}

// GenerateBroker returns the current state in the form of *svcapitypes.Broker.
func GenerateBroker(resp *svcsdk.DescribeBrokerResponse) *svcapitypes.Broker {
	cr := &svcapitypes.Broker{}

	if resp.AuthenticationStrategy != nil {
		cr.Spec.ForProvider.AuthenticationStrategy = resp.AuthenticationStrategy
	} else {
		cr.Spec.ForProvider.AuthenticationStrategy = nil
	}
	if resp.AutoMinorVersionUpgrade != nil {
		cr.Spec.ForProvider.AutoMinorVersionUpgrade = resp.AutoMinorVersionUpgrade
	} else {
		cr.Spec.ForProvider.AutoMinorVersionUpgrade = nil
	}
	if resp.BrokerArn != nil {
		cr.Status.AtProvider.BrokerARN = resp.BrokerArn
	} else {
		cr.Status.AtProvider.BrokerARN = nil
	}
	if resp.BrokerId != nil {
		cr.Status.AtProvider.BrokerID = resp.BrokerId
	} else {
		cr.Status.AtProvider.BrokerID = nil
	}
	if resp.BrokerInstances != nil {
		f5 := []*svcapitypes.BrokerInstance{}
		for _, f5iter := range resp.BrokerInstances {
			f5elem := &svcapitypes.BrokerInstance{}
			if f5iter.ConsoleURL != nil {
				f5elem.ConsoleURL = f5iter.ConsoleURL
			}
			if f5iter.Endpoints != nil {
				f5elemf1 := []*string{}
				for _, f5elemf1iter := range f5iter.Endpoints {
					var f5elemf1elem string
					f5elemf1elem = *f5elemf1iter
					f5elemf1 = append(f5elemf1, &f5elemf1elem)
				}
				f5elem.Endpoints = f5elemf1
			}
			if f5iter.IpAddress != nil {
				f5elem.IPAddress = f5iter.IpAddress
			}
			f5 = append(f5, f5elem)
		}
		cr.Status.AtProvider.BrokerInstances = f5
	} else {
		cr.Status.AtProvider.BrokerInstances = nil
	}
	if resp.BrokerState != nil {
		cr.Status.AtProvider.BrokerState = resp.BrokerState
	} else {
		cr.Status.AtProvider.BrokerState = nil
	}
	if resp.Configurations != nil {
		f8 := &svcapitypes.Configurations{}
		if resp.Configurations.Current != nil {
			f8f0 := &svcapitypes.ConfigurationID{}
			if resp.Configurations.Current.Id != nil {
				f8f0.ID = resp.Configurations.Current.Id
			}
			if resp.Configurations.Current.Revision != nil {
				f8f0.Revision = resp.Configurations.Current.Revision
			}
			f8.Current = f8f0
		}
		if resp.Configurations.History != nil {
			f8f1 := []*svcapitypes.ConfigurationID{}
			for _, f8f1iter := range resp.Configurations.History {
				f8f1elem := &svcapitypes.ConfigurationID{}
				if f8f1iter.Id != nil {
					f8f1elem.ID = f8f1iter.Id
				}
				if f8f1iter.Revision != nil {
					f8f1elem.Revision = f8f1iter.Revision
				}
				f8f1 = append(f8f1, f8f1elem)
			}
			f8.History = f8f1
		}
		if resp.Configurations.Pending != nil {
			f8f2 := &svcapitypes.ConfigurationID{}
			if resp.Configurations.Pending.Id != nil {
				f8f2.ID = resp.Configurations.Pending.Id
			}
			if resp.Configurations.Pending.Revision != nil {
				f8f2.Revision = resp.Configurations.Pending.Revision
			}
			f8.Pending = f8f2
		}
		cr.Status.AtProvider.Configurations = f8
	} else {
		cr.Status.AtProvider.Configurations = nil
	}
	if resp.Created != nil {
		cr.Status.AtProvider.Created = &metav1.Time{*resp.Created}
	} else {
		cr.Status.AtProvider.Created = nil
	}
	if resp.DeploymentMode != nil {
		cr.Spec.ForProvider.DeploymentMode = resp.DeploymentMode
	} else {
		cr.Spec.ForProvider.DeploymentMode = nil
	}
	if resp.EncryptionOptions != nil {
		f11 := &svcapitypes.EncryptionOptions{}
		if resp.EncryptionOptions.KmsKeyId != nil {
			f11.KMSKeyID = resp.EncryptionOptions.KmsKeyId
		}
		if resp.EncryptionOptions.UseAwsOwnedKey != nil {
			f11.UseAWSOwnedKey = resp.EncryptionOptions.UseAwsOwnedKey
		}
		cr.Spec.ForProvider.EncryptionOptions = f11
	} else {
		cr.Spec.ForProvider.EncryptionOptions = nil
	}
	if resp.EngineType != nil {
		cr.Spec.ForProvider.EngineType = resp.EngineType
	} else {
		cr.Spec.ForProvider.EngineType = nil
	}
	if resp.EngineVersion != nil {
		cr.Spec.ForProvider.EngineVersion = resp.EngineVersion
	} else {
		cr.Spec.ForProvider.EngineVersion = nil
	}
	if resp.HostInstanceType != nil {
		cr.Spec.ForProvider.HostInstanceType = resp.HostInstanceType
	} else {
		cr.Spec.ForProvider.HostInstanceType = nil
	}
	if resp.LdapServerMetadata != nil {
		f15 := &svcapitypes.LDAPServerMetadataInput{}
		if resp.LdapServerMetadata.Hosts != nil {
			f15f0 := []*string{}
			for _, f15f0iter := range resp.LdapServerMetadata.Hosts {
				var f15f0elem string
				f15f0elem = *f15f0iter
				f15f0 = append(f15f0, &f15f0elem)
			}
			f15.Hosts = f15f0
		}
		if resp.LdapServerMetadata.RoleBase != nil {
			f15.RoleBase = resp.LdapServerMetadata.RoleBase
		}
		if resp.LdapServerMetadata.RoleName != nil {
			f15.RoleName = resp.LdapServerMetadata.RoleName
		}
		if resp.LdapServerMetadata.RoleSearchMatching != nil {
			f15.RoleSearchMatching = resp.LdapServerMetadata.RoleSearchMatching
		}
		if resp.LdapServerMetadata.RoleSearchSubtree != nil {
			f15.RoleSearchSubtree = resp.LdapServerMetadata.RoleSearchSubtree
		}
		if resp.LdapServerMetadata.ServiceAccountUsername != nil {
			f15.ServiceAccountUsername = resp.LdapServerMetadata.ServiceAccountUsername
		}
		if resp.LdapServerMetadata.UserBase != nil {
			f15.UserBase = resp.LdapServerMetadata.UserBase
		}
		if resp.LdapServerMetadata.UserRoleName != nil {
			f15.UserRoleName = resp.LdapServerMetadata.UserRoleName
		}
		if resp.LdapServerMetadata.UserSearchMatching != nil {
			f15.UserSearchMatching = resp.LdapServerMetadata.UserSearchMatching
		}
		if resp.LdapServerMetadata.UserSearchSubtree != nil {
			f15.UserSearchSubtree = resp.LdapServerMetadata.UserSearchSubtree
		}
		cr.Spec.ForProvider.LDAPServerMetadata = f15
	} else {
		cr.Spec.ForProvider.LDAPServerMetadata = nil
	}
	if resp.Logs != nil {
		f16 := &svcapitypes.Logs{}
		if resp.Logs.Audit != nil {
			f16.Audit = resp.Logs.Audit
		}
		if resp.Logs.General != nil {
			f16.General = resp.Logs.General
		}
		cr.Spec.ForProvider.Logs = f16
	} else {
		cr.Spec.ForProvider.Logs = nil
	}
	if resp.MaintenanceWindowStartTime != nil {
		f17 := &svcapitypes.WeeklyStartTime{}
		if resp.MaintenanceWindowStartTime.DayOfWeek != nil {
			f17.DayOfWeek = resp.MaintenanceWindowStartTime.DayOfWeek
		}
		if resp.MaintenanceWindowStartTime.TimeOfDay != nil {
			f17.TimeOfDay = resp.MaintenanceWindowStartTime.TimeOfDay
		}
		if resp.MaintenanceWindowStartTime.TimeZone != nil {
			f17.TimeZone = resp.MaintenanceWindowStartTime.TimeZone
		}
		cr.Spec.ForProvider.MaintenanceWindowStartTime = f17
	} else {
		cr.Spec.ForProvider.MaintenanceWindowStartTime = nil
	}
	if resp.PendingAuthenticationStrategy != nil {
		cr.Status.AtProvider.PendingAuthenticationStrategy = resp.PendingAuthenticationStrategy
	} else {
		cr.Status.AtProvider.PendingAuthenticationStrategy = nil
	}
	if resp.PendingEngineVersion != nil {
		cr.Status.AtProvider.PendingEngineVersion = resp.PendingEngineVersion
	} else {
		cr.Status.AtProvider.PendingEngineVersion = nil
	}
	if resp.PendingHostInstanceType != nil {
		cr.Status.AtProvider.PendingHostInstanceType = resp.PendingHostInstanceType
	} else {
		cr.Status.AtProvider.PendingHostInstanceType = nil
	}
	if resp.PendingLdapServerMetadata != nil {
		f21 := &svcapitypes.LDAPServerMetadataOutput{}
		if resp.PendingLdapServerMetadata.Hosts != nil {
			f21f0 := []*string{}
			for _, f21f0iter := range resp.PendingLdapServerMetadata.Hosts {
				var f21f0elem string
				f21f0elem = *f21f0iter
				f21f0 = append(f21f0, &f21f0elem)
			}
			f21.Hosts = f21f0
		}
		if resp.PendingLdapServerMetadata.RoleBase != nil {
			f21.RoleBase = resp.PendingLdapServerMetadata.RoleBase
		}
		if resp.PendingLdapServerMetadata.RoleName != nil {
			f21.RoleName = resp.PendingLdapServerMetadata.RoleName
		}
		if resp.PendingLdapServerMetadata.RoleSearchMatching != nil {
			f21.RoleSearchMatching = resp.PendingLdapServerMetadata.RoleSearchMatching
		}
		if resp.PendingLdapServerMetadata.RoleSearchSubtree != nil {
			f21.RoleSearchSubtree = resp.PendingLdapServerMetadata.RoleSearchSubtree
		}
		if resp.PendingLdapServerMetadata.ServiceAccountUsername != nil {
			f21.ServiceAccountUsername = resp.PendingLdapServerMetadata.ServiceAccountUsername
		}
		if resp.PendingLdapServerMetadata.UserBase != nil {
			f21.UserBase = resp.PendingLdapServerMetadata.UserBase
		}
		if resp.PendingLdapServerMetadata.UserRoleName != nil {
			f21.UserRoleName = resp.PendingLdapServerMetadata.UserRoleName
		}
		if resp.PendingLdapServerMetadata.UserSearchMatching != nil {
			f21.UserSearchMatching = resp.PendingLdapServerMetadata.UserSearchMatching
		}
		if resp.PendingLdapServerMetadata.UserSearchSubtree != nil {
			f21.UserSearchSubtree = resp.PendingLdapServerMetadata.UserSearchSubtree
		}
		cr.Status.AtProvider.PendingLDAPServerMetadata = f21
	} else {
		cr.Status.AtProvider.PendingLDAPServerMetadata = nil
	}
	if resp.PendingSecurityGroups != nil {
		f22 := []*string{}
		for _, f22iter := range resp.PendingSecurityGroups {
			var f22elem string
			f22elem = *f22iter
			f22 = append(f22, &f22elem)
		}
		cr.Status.AtProvider.PendingSecurityGroups = f22
	} else {
		cr.Status.AtProvider.PendingSecurityGroups = nil
	}
	if resp.PubliclyAccessible != nil {
		cr.Spec.ForProvider.PubliclyAccessible = resp.PubliclyAccessible
	} else {
		cr.Spec.ForProvider.PubliclyAccessible = nil
	}
	if resp.StorageType != nil {
		cr.Spec.ForProvider.StorageType = resp.StorageType
	} else {
		cr.Spec.ForProvider.StorageType = nil
	}
	if resp.Tags != nil {
		f27 := map[string]*string{}
		for f27key, f27valiter := range resp.Tags {
			var f27val string
			f27val = *f27valiter
			f27[f27key] = &f27val
		}
		cr.Spec.ForProvider.Tags = f27
	} else {
		cr.Spec.ForProvider.Tags = nil
	}
	if resp.Users != nil {
		f28 := []*svcapitypes.UserSummary{}
		for _, f28iter := range resp.Users {
			f28elem := &svcapitypes.UserSummary{}
			if f28iter.PendingChange != nil {
				f28elem.PendingChange = f28iter.PendingChange
			}
			if f28iter.Username != nil {
				f28elem.Username = f28iter.Username
			}
			f28 = append(f28, f28elem)
		}
		cr.Status.AtProvider.Users = f28
	} else {
		cr.Status.AtProvider.Users = nil
	}

	return cr
}

// GenerateCreateBrokerRequest returns a create input.
func GenerateCreateBrokerRequest(cr *svcapitypes.Broker) *svcsdk.CreateBrokerRequest {
	res := &svcsdk.CreateBrokerRequest{}

	if cr.Spec.ForProvider.AuthenticationStrategy != nil {
		res.SetAuthenticationStrategy(*cr.Spec.ForProvider.AuthenticationStrategy)
	}
	if cr.Spec.ForProvider.AutoMinorVersionUpgrade != nil {
		res.SetAutoMinorVersionUpgrade(*cr.Spec.ForProvider.AutoMinorVersionUpgrade)
	}
	if cr.Spec.ForProvider.Configuration != nil {
		f2 := &svcsdk.ConfigurationId{}
		if cr.Spec.ForProvider.Configuration.ID != nil {
			f2.SetId(*cr.Spec.ForProvider.Configuration.ID)
		}
		if cr.Spec.ForProvider.Configuration.Revision != nil {
			f2.SetRevision(*cr.Spec.ForProvider.Configuration.Revision)
		}
		res.SetConfiguration(f2)
	}
	if cr.Spec.ForProvider.CreatorRequestID != nil {
		res.SetCreatorRequestId(*cr.Spec.ForProvider.CreatorRequestID)
	}
	if cr.Spec.ForProvider.DeploymentMode != nil {
		res.SetDeploymentMode(*cr.Spec.ForProvider.DeploymentMode)
	}
	if cr.Spec.ForProvider.EncryptionOptions != nil {
		f5 := &svcsdk.EncryptionOptions{}
		if cr.Spec.ForProvider.EncryptionOptions.KMSKeyID != nil {
			f5.SetKmsKeyId(*cr.Spec.ForProvider.EncryptionOptions.KMSKeyID)
		}
		if cr.Spec.ForProvider.EncryptionOptions.UseAWSOwnedKey != nil {
			f5.SetUseAwsOwnedKey(*cr.Spec.ForProvider.EncryptionOptions.UseAWSOwnedKey)
		}
		res.SetEncryptionOptions(f5)
	}
	if cr.Spec.ForProvider.EngineType != nil {
		res.SetEngineType(*cr.Spec.ForProvider.EngineType)
	}
	if cr.Spec.ForProvider.EngineVersion != nil {
		res.SetEngineVersion(*cr.Spec.ForProvider.EngineVersion)
	}
	if cr.Spec.ForProvider.HostInstanceType != nil {
		res.SetHostInstanceType(*cr.Spec.ForProvider.HostInstanceType)
	}
	if cr.Spec.ForProvider.LDAPServerMetadata != nil {
		f9 := &svcsdk.LdapServerMetadataInput{}
		if cr.Spec.ForProvider.LDAPServerMetadata.Hosts != nil {
			f9f0 := []*string{}
			for _, f9f0iter := range cr.Spec.ForProvider.LDAPServerMetadata.Hosts {
				var f9f0elem string
				f9f0elem = *f9f0iter
				f9f0 = append(f9f0, &f9f0elem)
			}
			f9.SetHosts(f9f0)
		}
		if cr.Spec.ForProvider.LDAPServerMetadata.RoleBase != nil {
			f9.SetRoleBase(*cr.Spec.ForProvider.LDAPServerMetadata.RoleBase)
		}
		if cr.Spec.ForProvider.LDAPServerMetadata.RoleName != nil {
			f9.SetRoleName(*cr.Spec.ForProvider.LDAPServerMetadata.RoleName)
		}
		if cr.Spec.ForProvider.LDAPServerMetadata.RoleSearchMatching != nil {
			f9.SetRoleSearchMatching(*cr.Spec.ForProvider.LDAPServerMetadata.RoleSearchMatching)
		}
		if cr.Spec.ForProvider.LDAPServerMetadata.RoleSearchSubtree != nil {
			f9.SetRoleSearchSubtree(*cr.Spec.ForProvider.LDAPServerMetadata.RoleSearchSubtree)
		}
		if cr.Spec.ForProvider.LDAPServerMetadata.ServiceAccountPassword != nil {
			f9.SetServiceAccountPassword(*cr.Spec.ForProvider.LDAPServerMetadata.ServiceAccountPassword)
		}
		if cr.Spec.ForProvider.LDAPServerMetadata.ServiceAccountUsername != nil {
			f9.SetServiceAccountUsername(*cr.Spec.ForProvider.LDAPServerMetadata.ServiceAccountUsername)
		}
		if cr.Spec.ForProvider.LDAPServerMetadata.UserBase != nil {
			f9.SetUserBase(*cr.Spec.ForProvider.LDAPServerMetadata.UserBase)
		}
		if cr.Spec.ForProvider.LDAPServerMetadata.UserRoleName != nil {
			f9.SetUserRoleName(*cr.Spec.ForProvider.LDAPServerMetadata.UserRoleName)
		}
		if cr.Spec.ForProvider.LDAPServerMetadata.UserSearchMatching != nil {
			f9.SetUserSearchMatching(*cr.Spec.ForProvider.LDAPServerMetadata.UserSearchMatching)
		}
		if cr.Spec.ForProvider.LDAPServerMetadata.UserSearchSubtree != nil {
			f9.SetUserSearchSubtree(*cr.Spec.ForProvider.LDAPServerMetadata.UserSearchSubtree)
		}
		res.SetLdapServerMetadata(f9)
	}
	if cr.Spec.ForProvider.Logs != nil {
		f10 := &svcsdk.Logs{}
		if cr.Spec.ForProvider.Logs.Audit != nil {
			f10.SetAudit(*cr.Spec.ForProvider.Logs.Audit)
		}
		if cr.Spec.ForProvider.Logs.General != nil {
			f10.SetGeneral(*cr.Spec.ForProvider.Logs.General)
		}
		res.SetLogs(f10)
	}
	if cr.Spec.ForProvider.MaintenanceWindowStartTime != nil {
		f11 := &svcsdk.WeeklyStartTime{}
		if cr.Spec.ForProvider.MaintenanceWindowStartTime.DayOfWeek != nil {
			f11.SetDayOfWeek(*cr.Spec.ForProvider.MaintenanceWindowStartTime.DayOfWeek)
		}
		if cr.Spec.ForProvider.MaintenanceWindowStartTime.TimeOfDay != nil {
			f11.SetTimeOfDay(*cr.Spec.ForProvider.MaintenanceWindowStartTime.TimeOfDay)
		}
		if cr.Spec.ForProvider.MaintenanceWindowStartTime.TimeZone != nil {
			f11.SetTimeZone(*cr.Spec.ForProvider.MaintenanceWindowStartTime.TimeZone)
		}
		res.SetMaintenanceWindowStartTime(f11)
	}
	if cr.Spec.ForProvider.PubliclyAccessible != nil {
		res.SetPubliclyAccessible(*cr.Spec.ForProvider.PubliclyAccessible)
	}
	if cr.Spec.ForProvider.StorageType != nil {
		res.SetStorageType(*cr.Spec.ForProvider.StorageType)
	}
	if cr.Spec.ForProvider.Tags != nil {
		f14 := map[string]*string{}
		for f14key, f14valiter := range cr.Spec.ForProvider.Tags {
			var f14val string
			f14val = *f14valiter
			f14[f14key] = &f14val
		}
		res.SetTags(f14)
	}

	return res
}

// GenerateUpdateBrokerRequest returns an update input.
func GenerateUpdateBrokerRequest(cr *svcapitypes.Broker) *svcsdk.UpdateBrokerRequest {
	res := &svcsdk.UpdateBrokerRequest{}

	if cr.Spec.ForProvider.AuthenticationStrategy != nil {
		res.SetAuthenticationStrategy(*cr.Spec.ForProvider.AuthenticationStrategy)
	}
	if cr.Spec.ForProvider.AutoMinorVersionUpgrade != nil {
		res.SetAutoMinorVersionUpgrade(*cr.Spec.ForProvider.AutoMinorVersionUpgrade)
	}
	if cr.Status.AtProvider.BrokerID != nil {
		res.SetBrokerId(*cr.Status.AtProvider.BrokerID)
	}
	if cr.Spec.ForProvider.Configuration != nil {
		f3 := &svcsdk.ConfigurationId{}
		if cr.Spec.ForProvider.Configuration.ID != nil {
			f3.SetId(*cr.Spec.ForProvider.Configuration.ID)
		}
		if cr.Spec.ForProvider.Configuration.Revision != nil {
			f3.SetRevision(*cr.Spec.ForProvider.Configuration.Revision)
		}
		res.SetConfiguration(f3)
	}
	if cr.Spec.ForProvider.EngineVersion != nil {
		res.SetEngineVersion(*cr.Spec.ForProvider.EngineVersion)
	}
	if cr.Spec.ForProvider.HostInstanceType != nil {
		res.SetHostInstanceType(*cr.Spec.ForProvider.HostInstanceType)
	}
	if cr.Spec.ForProvider.LDAPServerMetadata != nil {
		f6 := &svcsdk.LdapServerMetadataInput{}
		if cr.Spec.ForProvider.LDAPServerMetadata.Hosts != nil {
			f6f0 := []*string{}
			for _, f6f0iter := range cr.Spec.ForProvider.LDAPServerMetadata.Hosts {
				var f6f0elem string
				f6f0elem = *f6f0iter
				f6f0 = append(f6f0, &f6f0elem)
			}
			f6.SetHosts(f6f0)
		}
		if cr.Spec.ForProvider.LDAPServerMetadata.RoleBase != nil {
			f6.SetRoleBase(*cr.Spec.ForProvider.LDAPServerMetadata.RoleBase)
		}
		if cr.Spec.ForProvider.LDAPServerMetadata.RoleName != nil {
			f6.SetRoleName(*cr.Spec.ForProvider.LDAPServerMetadata.RoleName)
		}
		if cr.Spec.ForProvider.LDAPServerMetadata.RoleSearchMatching != nil {
			f6.SetRoleSearchMatching(*cr.Spec.ForProvider.LDAPServerMetadata.RoleSearchMatching)
		}
		if cr.Spec.ForProvider.LDAPServerMetadata.RoleSearchSubtree != nil {
			f6.SetRoleSearchSubtree(*cr.Spec.ForProvider.LDAPServerMetadata.RoleSearchSubtree)
		}
		if cr.Spec.ForProvider.LDAPServerMetadata.ServiceAccountPassword != nil {
			f6.SetServiceAccountPassword(*cr.Spec.ForProvider.LDAPServerMetadata.ServiceAccountPassword)
		}
		if cr.Spec.ForProvider.LDAPServerMetadata.ServiceAccountUsername != nil {
			f6.SetServiceAccountUsername(*cr.Spec.ForProvider.LDAPServerMetadata.ServiceAccountUsername)
		}
		if cr.Spec.ForProvider.LDAPServerMetadata.UserBase != nil {
			f6.SetUserBase(*cr.Spec.ForProvider.LDAPServerMetadata.UserBase)
		}
		if cr.Spec.ForProvider.LDAPServerMetadata.UserRoleName != nil {
			f6.SetUserRoleName(*cr.Spec.ForProvider.LDAPServerMetadata.UserRoleName)
		}
		if cr.Spec.ForProvider.LDAPServerMetadata.UserSearchMatching != nil {
			f6.SetUserSearchMatching(*cr.Spec.ForProvider.LDAPServerMetadata.UserSearchMatching)
		}
		if cr.Spec.ForProvider.LDAPServerMetadata.UserSearchSubtree != nil {
			f6.SetUserSearchSubtree(*cr.Spec.ForProvider.LDAPServerMetadata.UserSearchSubtree)
		}
		res.SetLdapServerMetadata(f6)
	}
	if cr.Spec.ForProvider.Logs != nil {
		f7 := &svcsdk.Logs{}
		if cr.Spec.ForProvider.Logs.Audit != nil {
			f7.SetAudit(*cr.Spec.ForProvider.Logs.Audit)
		}
		if cr.Spec.ForProvider.Logs.General != nil {
			f7.SetGeneral(*cr.Spec.ForProvider.Logs.General)
		}
		res.SetLogs(f7)
	}
	if cr.Spec.ForProvider.MaintenanceWindowStartTime != nil {
		f8 := &svcsdk.WeeklyStartTime{}
		if cr.Spec.ForProvider.MaintenanceWindowStartTime.DayOfWeek != nil {
			f8.SetDayOfWeek(*cr.Spec.ForProvider.MaintenanceWindowStartTime.DayOfWeek)
		}
		if cr.Spec.ForProvider.MaintenanceWindowStartTime.TimeOfDay != nil {
			f8.SetTimeOfDay(*cr.Spec.ForProvider.MaintenanceWindowStartTime.TimeOfDay)
		}
		if cr.Spec.ForProvider.MaintenanceWindowStartTime.TimeZone != nil {
			f8.SetTimeZone(*cr.Spec.ForProvider.MaintenanceWindowStartTime.TimeZone)
		}
		res.SetMaintenanceWindowStartTime(f8)
	}

	return res
}

// GenerateDeleteBrokerInput returns a deletion input.
func GenerateDeleteBrokerInput(cr *svcapitypes.Broker) *svcsdk.DeleteBrokerInput {
	res := &svcsdk.DeleteBrokerInput{}

	if cr.Status.AtProvider.BrokerID != nil {
		res.SetBrokerId(*cr.Status.AtProvider.BrokerID)
	}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "NotFoundException"
}
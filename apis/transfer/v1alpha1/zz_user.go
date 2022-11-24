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
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// UserParameters defines the desired state of User
type UserParameters struct {
	// Region is which region the User will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// The landing directory (folder) for a user when they log in to the server
	// using the client.
	//
	// A HomeDirectory example is /bucket_name/home/mydirectory.
	HomeDirectory *string `json:"homeDirectory,omitempty"`
	// Logical directory mappings that specify what Amazon S3 or Amazon EFS paths
	// and keys should be visible to your user and how you want to make them visible.
	// You must specify the Entry and Target pair, where Entry shows how the path
	// is made visible and Target is the actual Amazon S3 or Amazon EFS path. If
	// you only specify a target, it is displayed as is. You also must ensure that
	// your Amazon Web Services Identity and Access Management (IAM) role provides
	// access to paths in Target. This value can only be set when HomeDirectoryType
	// is set to LOGICAL.
	//
	// The following is an Entry and Target pair example.
	//
	// [ { "Entry": "/directory1", "Target": "/bucket_name/home/mydirectory" } ]
	//
	// In most cases, you can use this value instead of the session policy to lock
	// your user down to the designated home directory ("chroot"). To do this, you
	// can set Entry to / and set Target to the HomeDirectory parameter value.
	//
	// The following is an Entry and Target pair example for chroot.
	//
	// [ { "Entry": "/", "Target": "/bucket_name/home/mydirectory" } ]
	HomeDirectoryMappings []*HomeDirectoryMapEntry `json:"homeDirectoryMappings,omitempty"`
	// The type of landing directory (folder) you want your users' home directory
	// to be when they log into the server. If you set it to PATH, the user will
	// see the absolute Amazon S3 bucket or EFS paths as is in their file transfer
	// protocol clients. If you set it LOGICAL, you need to provide mappings in
	// the HomeDirectoryMappings for how you want to make Amazon S3 or EFS paths
	// visible to your users.
	HomeDirectoryType *string `json:"homeDirectoryType,omitempty"`
	// A session policy for your user so that you can use the same IAM role across
	// multiple users. This policy scopes down user access to portions of their
	// Amazon S3 bucket. Variables that you can use inside this policy include ${Transfer:UserName},
	// ${Transfer:HomeDirectory}, and ${Transfer:HomeBucket}.
	//
	// This only applies when the domain of ServerId is S3. EFS does not use session
	// policies.
	//
	// For session policies, Amazon Web Services Transfer Family stores the policy
	// as a JSON blob, instead of the Amazon Resource Name (ARN) of the policy.
	// You save the policy as a JSON blob and pass it in the Policy argument.
	//
	// For an example of a session policy, see Example session policy (https://docs.aws.amazon.com/transfer/latest/userguide/session-policy.html).
	//
	// For more information, see AssumeRole (https://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRole.html)
	// in the Amazon Web Services Security Token Service API Reference.
	Policy *string `json:"policy,omitempty"`
	// Specifies the full POSIX identity, including user ID (Uid), group ID (Gid),
	// and any secondary groups IDs (SecondaryGids), that controls your users' access
	// to your Amazon EFS file systems. The POSIX permissions that are set on files
	// and directories in Amazon EFS determine the level of access your users get
	// when transferring files into and out of your Amazon EFS file systems.
	PosixProfile *PosixProfile `json:"posixProfile,omitempty"`
	// The public portion of the Secure Shell (SSH) key used to authenticate the
	// user to the server.
	//
	// Currently, Transfer Family does not accept elliptical curve keys (keys beginning
	// with ecdsa).
	SshPublicKeyBody *string `json:"sshPublicKeyBody,omitempty"`
	// Key-value pairs that can be used to group and search for users. Tags are
	// metadata attached to users for any purpose.
	Tags                 []*Tag `json:"tags,omitempty"`
	CustomUserParameters `json:",inline"`
}

// UserSpec defines the desired state of User
type UserSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       UserParameters `json:"forProvider"`
}

// UserObservation defines the observed state of User
type UserObservation struct {
	// The ID of the server that the user is attached to.
	ServerID *string `json:"serverID,omitempty"`
	// A unique string that identifies a user account associated with a server.
	UserName *string `json:"userName,omitempty"`
}

// UserStatus defines the observed state of User.
type UserStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          UserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// User is the Schema for the Users API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserSpec   `json:"spec"`
	Status            UserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserList contains a list of Users
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User `json:"items"`
}

// Repository type metadata.
var (
	UserKind             = "User"
	UserGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserKind}.String()
	UserKindAPIVersion   = UserKind + "." + GroupVersion.String()
	UserGroupVersionKind = GroupVersion.WithKind(UserKind)
)

func init() {
	SchemeBuilder.Register(&User{}, &UserList{})
}
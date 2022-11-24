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

// FileSystemParameters defines the desired state of FileSystem
type FileSystemParameters struct {
	// Region is which region the FileSystem will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// Used to create a file system that uses One Zone storage classes. It specifies
	// the Amazon Web Services Availability Zone in which to create the file system.
	// Use the format us-east-1a to specify the Availability Zone. For more information
	// about One Zone storage classes, see Using EFS storage classes (https://docs.aws.amazon.com/efs/latest/ug/storage-classes.html)
	// in the Amazon EFS User Guide.
	//
	// One Zone storage classes are not available in all Availability Zones in Amazon
	// Web Services Regions where Amazon EFS is available.
	AvailabilityZoneName *string `json:"availabilityZoneName,omitempty"`
	// Specifies whether automatic backups are enabled on the file system that you
	// are creating. Set the value to true to enable automatic backups. If you are
	// creating a file system that uses One Zone storage classes, automatic backups
	// are enabled by default. For more information, see Automatic backups (https://docs.aws.amazon.com/efs/latest/ug/awsbackup.html#automatic-backups)
	// in the Amazon EFS User Guide.
	//
	// Default is false. However, if you specify an AvailabilityZoneName, the default
	// is true.
	//
	// Backup is not available in all Amazon Web Services Regions where Amazon EFS
	// is available.
	Backup *bool `json:"backup,omitempty"`
	// A Boolean value that, if true, creates an encrypted file system. When creating
	// an encrypted file system, you have the option of specifying an existing Key
	// Management Service key (KMS key). If you don't specify a KMS key, then the
	// default KMS key for Amazon EFS, /aws/elasticfilesystem, is used to protect
	// the encrypted file system.
	Encrypted *bool `json:"encrypted,omitempty"`
	// The ID of the KMS key that you want to use to protect the encrypted file
	// system. This parameter is required only if you want to use a non-default
	// KMS key. If this parameter is not specified, the default KMS key for Amazon
	// EFS is used. You can specify a KMS key ID using the following formats:
	//
	//    * Key ID - A unique identifier of the key, for example 1234abcd-12ab-34cd-56ef-1234567890ab.
	//
	//    * ARN - An Amazon Resource Name (ARN) for the key, for example arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab.
	//
	//    * Key alias - A previously created display name for a key, for example
	//    alias/projectKey1.
	//
	//    * Key alias ARN - An ARN for a key alias, for example arn:aws:kms:us-west-2:444455556666:alias/projectKey1.
	//
	// If you use KmsKeyId, you must set the CreateFileSystemRequest$Encrypted parameter
	// to true.
	//
	// EFS accepts only symmetric KMS keys. You cannot use asymmetric KMS keys with
	// Amazon EFS file systems.
	KMSKeyID *string `json:"kmsKeyID,omitempty"`
	// The performance mode of the file system. We recommend generalPurpose performance
	// mode for most file systems. File systems using the maxIO performance mode
	// can scale to higher levels of aggregate throughput and operations per second
	// with a tradeoff of slightly higher latencies for most file operations. The
	// performance mode can't be changed after the file system has been created.
	//
	// The maxIO mode is not supported on file systems using One Zone storage classes.
	PerformanceMode *string `json:"performanceMode,omitempty"`
	// Use to create one or more tags associated with the file system. Each tag
	// is a user-defined key-value pair. Name your file system on creation by including
	// a "Key":"Name","Value":"{value}" key-value pair. Each key must be unique.
	// For more information, see Tagging Amazon Web Services resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// in the Amazon Web Services General Reference Guide.
	Tags []*Tag `json:"tags,omitempty"`
	// Specifies the throughput mode for the file system, either bursting or provisioned.
	// If you set ThroughputMode to provisioned, you must also set a value for ProvisionedThroughputInMibps.
	// After you create the file system, you can decrease your file system's throughput
	// in Provisioned Throughput mode or change between the throughput modes, as
	// long as it’s been more than 24 hours since the last decrease or throughput
	// mode change. For more information, see Specifying throughput with provisioned
	// mode (https://docs.aws.amazon.com/efs/latest/ug/performance.html#provisioned-throughput)
	// in the Amazon EFS User Guide.
	//
	// Default is bursting.
	ThroughputMode             *string `json:"throughputMode,omitempty"`
	CustomFileSystemParameters `json:",inline"`
}

// FileSystemSpec defines the desired state of FileSystem
type FileSystemSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       FileSystemParameters `json:"forProvider"`
}

// FileSystemObservation defines the observed state of FileSystem
type FileSystemObservation struct {
	// The unique and consistent identifier of the Availability Zone in which the
	// file system's One Zone storage classes exist. For example, use1-az1 is an
	// Availability Zone ID for the us-east-1 Amazon Web Services Region, and it
	// has the same location in every Amazon Web Services account.
	AvailabilityZoneID *string `json:"availabilityZoneID,omitempty"`
	// The time that the file system was created, in seconds (since 1970-01-01T00:00:00Z).
	CreationTime *metav1.Time `json:"creationTime,omitempty"`
	// The opaque string specified in the request.
	CreationToken *string `json:"creationToken,omitempty"`
	// The Amazon Resource Name (ARN) for the EFS file system, in the format arn:aws:elasticfilesystem:region:account-id:file-system/file-system-id
	// . Example with sample data: arn:aws:elasticfilesystem:us-west-2:1111333322228888:file-system/fs-01234567
	FileSystemARN *string `json:"fileSystemARN,omitempty"`
	// The ID of the file system, assigned by Amazon EFS.
	FileSystemID *string `json:"fileSystemID,omitempty"`
	// The lifecycle phase of the file system.
	LifeCycleState *string `json:"lifeCycleState,omitempty"`
	// You can add tags to a file system, including a Name tag. For more information,
	// see CreateFileSystem. If the file system has a Name tag, Amazon EFS returns
	// the value in this field.
	Name *string `json:"name,omitempty"`
	// The current number of mount targets that the file system has. For more information,
	// see CreateMountTarget.
	NumberOfMountTargets *int64 `json:"numberOfMountTargets,omitempty"`
	// The Amazon Web Services account that created the file system. If the file
	// system was created by an IAM user, the parent account to which the user belongs
	// is the owner.
	OwnerID *string `json:"ownerID,omitempty"`
	// The latest known metered size (in bytes) of data stored in the file system,
	// in its Value field, and the time at which that size was determined in its
	// Timestamp field. The Timestamp value is the integer number of seconds since
	// 1970-01-01T00:00:00Z. The SizeInBytes value doesn't represent the size of
	// a consistent snapshot of the file system, but it is eventually consistent
	// when there are no writes to the file system. That is, SizeInBytes represents
	// actual size only if the file system is not modified for a period longer than
	// a couple of hours. Otherwise, the value is not the exact size that the file
	// system was at any point in time.
	SizeInBytes *FileSystemSize `json:"sizeInBytes,omitempty"`
}

// FileSystemStatus defines the observed state of FileSystem.
type FileSystemStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          FileSystemObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FileSystem is the Schema for the FileSystems API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type FileSystem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FileSystemSpec   `json:"spec"`
	Status            FileSystemStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FileSystemList contains a list of FileSystems
type FileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FileSystem `json:"items"`
}

// Repository type metadata.
var (
	FileSystemKind             = "FileSystem"
	FileSystemGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FileSystemKind}.String()
	FileSystemKindAPIVersion   = FileSystemKind + "." + GroupVersion.String()
	FileSystemGroupVersionKind = GroupVersion.WithKind(FileSystemKind)
)

func init() {
	SchemeBuilder.Register(&FileSystem{}, &FileSystemList{})
}

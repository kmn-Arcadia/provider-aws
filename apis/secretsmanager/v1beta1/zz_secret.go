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

package v1beta1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// SecretParameters defines the desired state of Secret
type SecretParameters struct {
	// Region is which region the Secret will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// A list of Regions and KMS keys to replicate secrets.
	 AddReplicaRegions []*ReplicaRegionType `json:"addReplicaRegions,omitempty"` 
	// The description of the secret.
	 Description *string `json:"description,omitempty"` 
	// Specifies whether to overwrite a secret with the same name in the destination
// Region.
	 ForceOverwriteReplicaSecret *bool `json:"forceOverwriteReplicaSecret,omitempty"` 
	// The ARN, key ID, or alias of the KMS key that Secrets Manager uses to encrypt
// the secret value in the secret.
// 
// To use a KMS key in a different account, use the key ARN or the alias ARN.
// 
// If you don't specify this value, then Secrets Manager uses the key aws/secretsmanager.
// If that key doesn't yet exist, then Secrets Manager creates it for you automatically
// the first time it encrypts the secret value.
// 
// If the secret is in a different Amazon Web Services account from the credentials
// calling the API, then you can't use aws/secretsmanager to encrypt the secret,
// and you must create and use a customer managed KMS key.
	 KMSKeyID *string `json:"kmsKeyID,omitempty"` 
	// A list of tags to attach to the secret. Each tag is a key and value pair
// of strings in a JSON text string, for example:
// 
// [{"Key":"CostCenter","Value":"12345"},{"Key":"environment","Value":"production"}]
// 
// Secrets Manager tag key names are case sensitive. A tag with the key "ABC"
// is a different tag from one with key "abc".
// 
// If you check tags in permissions policies as part of your security strategy,
// then adding or removing a tag can change permissions. If the completion of
// this operation would result in you losing your permissions for this secret,
// then Secrets Manager blocks the operation and returns an Access Denied error.
// For more information, see Control access to secrets using tags (https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html#tag-secrets-abac)
// and Limit access to identities with tags that match secrets' tags (https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_examples.html#auth-and-access_tags2).
// 
// For information about how to format a JSON parameter for the various command
// line tool environments, see Using JSON for Parameters (https://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#cli-using-param-json).
// If your command-line tool or SDK requires quotation marks around the parameter,
// you should use single quotes to avoid confusion with the double quotes required
// in the JSON text.
// 
// The following restrictions apply to tags:
// 
//    * Maximum number of tags per secret: 50
// 
//    * Maximum key length: 127 Unicode characters in UTF-8
// 
//    * Maximum value length: 255 Unicode characters in UTF-8
// 
//    * Tag keys and values are case sensitive.
// 
//    * Do not use the aws: prefix in your tag names or values because Amazon
//    Web Services reserves it for Amazon Web Services use. You can't edit or
//    delete tag names or values with this prefix. Tags with this prefix do
//    not count against your tags per secret limit.
// 
//    * If you use your tagging schema across multiple services and resources,
//    other services might have restrictions on allowed characters. Generally
//    allowed characters: letters, spaces, and numbers representable in UTF-8,
//    plus the following special characters: + - = . _ : / @.
	 Tags []*Tag `json:"tags,omitempty"` 
	CustomSecretParameters `json:",inline"`
}

// SecretSpec defines the desired state of Secret
type SecretSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider SecretParameters `json:"forProvider"`
}

// SecretObservation defines the observed state of Secret
type SecretObservation struct {
	// The ARN of the new secret. The ARN includes the name of the secret followed
// by six random characters. This ensures that if you create a new secret with
// the same name as a deleted secret, then users with access to the old secret
// don't get access to the new secret because the ARNs are different.
	ARN *string `json:"arn,omitempty"`
	// A list of the replicas of this secret and their status:
// 
//    * Failed, which indicates that the replica was not created.
// 
//    * InProgress, which indicates that Secrets Manager is in the process of
//    creating the replica.
// 
//    * InSync, which indicates that the replica was created.
	ReplicationStatus []*ReplicationStatusType `json:"replicationStatus,omitempty"`
}

// SecretStatus defines the observed state of Secret.
type SecretStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider SecretObservation `json:"atProvider,omitempty"`
}


// +kubebuilder:object:root=true

// Secret is the Schema for the Secrets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Secret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec   SecretSpec   `json:"spec"`
	Status SecretStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretList contains a list of Secrets
type SecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items []Secret `json:"items"`
}

// Repository type metadata.
var (
	SecretKind             = "Secret"
	SecretGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretKind}.String()
	SecretKindAPIVersion   = SecretKind + "." + GroupVersion.String()
	SecretGroupVersionKind = GroupVersion.WithKind(SecretKind)
)

func init() {
	SchemeBuilder.Register(&Secret{}, &SecretList{})
}

//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChildShard) DeepCopyInto(out *ChildShard) {
	*out = *in
	if in.HashKeyRange != nil {
		in, out := &in.HashKeyRange, &out.HashKeyRange
		*out = new(HashKeyRange)
		(*in).DeepCopyInto(*out)
	}
	if in.ShardID != nil {
		in, out := &in.ShardID, &out.ShardID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChildShard.
func (in *ChildShard) DeepCopy() *ChildShard {
	if in == nil {
		return nil
	}
	out := new(ChildShard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Consumer) DeepCopyInto(out *Consumer) {
	*out = *in
	if in.ConsumerCreationTimestamp != nil {
		in, out := &in.ConsumerCreationTimestamp, &out.ConsumerCreationTimestamp
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Consumer.
func (in *Consumer) DeepCopy() *Consumer {
	if in == nil {
		return nil
	}
	out := new(Consumer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsumerDescription) DeepCopyInto(out *ConsumerDescription) {
	*out = *in
	if in.ConsumerCreationTimestamp != nil {
		in, out := &in.ConsumerCreationTimestamp, &out.ConsumerCreationTimestamp
		*out = (*in).DeepCopy()
	}
	if in.StreamARN != nil {
		in, out := &in.StreamARN, &out.StreamARN
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsumerDescription.
func (in *ConsumerDescription) DeepCopy() *ConsumerDescription {
	if in == nil {
		return nil
	}
	out := new(ConsumerDescription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomStreamParameters) DeepCopyInto(out *CustomStreamParameters) {
	*out = *in
	if in.RetentionPeriodHours != nil {
		in, out := &in.RetentionPeriodHours, &out.RetentionPeriodHours
		*out = new(int64)
		**out = **in
	}
	if in.KMSKeyARN != nil {
		in, out := &in.KMSKeyARN, &out.KMSKeyARN
		*out = new(string)
		**out = **in
	}
	if in.KMSKeyARNRef != nil {
		in, out := &in.KMSKeyARNRef, &out.KMSKeyARNRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.KMSKeyARNSelector != nil {
		in, out := &in.KMSKeyARNSelector, &out.KMSKeyARNSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.EnhancedMetrics != nil {
		in, out := &in.EnhancedMetrics, &out.EnhancedMetrics
		*out = make([]*EnhancedMetrics, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(EnhancedMetrics)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]CustomTag, len(*in))
		copy(*out, *in)
	}
	if in.EnforceConsumerDeletion != nil {
		in, out := &in.EnforceConsumerDeletion, &out.EnforceConsumerDeletion
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomStreamParameters.
func (in *CustomStreamParameters) DeepCopy() *CustomStreamParameters {
	if in == nil {
		return nil
	}
	out := new(CustomStreamParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomTag) DeepCopyInto(out *CustomTag) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomTag.
func (in *CustomTag) DeepCopy() *CustomTag {
	if in == nil {
		return nil
	}
	out := new(CustomTag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnhancedMetrics) DeepCopyInto(out *EnhancedMetrics) {
	*out = *in
	if in.ShardLevelMetrics != nil {
		in, out := &in.ShardLevelMetrics, &out.ShardLevelMetrics
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnhancedMetrics.
func (in *EnhancedMetrics) DeepCopy() *EnhancedMetrics {
	if in == nil {
		return nil
	}
	out := new(EnhancedMetrics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HashKeyRange) DeepCopyInto(out *HashKeyRange) {
	*out = *in
	if in.EndingHashKey != nil {
		in, out := &in.EndingHashKey, &out.EndingHashKey
		*out = new(string)
		**out = **in
	}
	if in.StartingHashKey != nil {
		in, out := &in.StartingHashKey, &out.StartingHashKey
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HashKeyRange.
func (in *HashKeyRange) DeepCopy() *HashKeyRange {
	if in == nil {
		return nil
	}
	out := new(HashKeyRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PutRecordsRequestEntry) DeepCopyInto(out *PutRecordsRequestEntry) {
	*out = *in
	if in.ExplicitHashKey != nil {
		in, out := &in.ExplicitHashKey, &out.ExplicitHashKey
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PutRecordsRequestEntry.
func (in *PutRecordsRequestEntry) DeepCopy() *PutRecordsRequestEntry {
	if in == nil {
		return nil
	}
	out := new(PutRecordsRequestEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PutRecordsResultEntry) DeepCopyInto(out *PutRecordsResultEntry) {
	*out = *in
	if in.SequenceNumber != nil {
		in, out := &in.SequenceNumber, &out.SequenceNumber
		*out = new(string)
		**out = **in
	}
	if in.ShardID != nil {
		in, out := &in.ShardID, &out.ShardID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PutRecordsResultEntry.
func (in *PutRecordsResultEntry) DeepCopy() *PutRecordsResultEntry {
	if in == nil {
		return nil
	}
	out := new(PutRecordsResultEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Record) DeepCopyInto(out *Record) {
	*out = *in
	if in.ApproximateArrivalTimestamp != nil {
		in, out := &in.ApproximateArrivalTimestamp, &out.ApproximateArrivalTimestamp
		*out = (*in).DeepCopy()
	}
	if in.EncryptionType != nil {
		in, out := &in.EncryptionType, &out.EncryptionType
		*out = new(string)
		**out = **in
	}
	if in.SequenceNumber != nil {
		in, out := &in.SequenceNumber, &out.SequenceNumber
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Record.
func (in *Record) DeepCopy() *Record {
	if in == nil {
		return nil
	}
	out := new(Record)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SequenceNumberRange) DeepCopyInto(out *SequenceNumberRange) {
	*out = *in
	if in.EndingSequenceNumber != nil {
		in, out := &in.EndingSequenceNumber, &out.EndingSequenceNumber
		*out = new(string)
		**out = **in
	}
	if in.StartingSequenceNumber != nil {
		in, out := &in.StartingSequenceNumber, &out.StartingSequenceNumber
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SequenceNumberRange.
func (in *SequenceNumberRange) DeepCopy() *SequenceNumberRange {
	if in == nil {
		return nil
	}
	out := new(SequenceNumberRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Shard) DeepCopyInto(out *Shard) {
	*out = *in
	if in.AdjacentParentShardID != nil {
		in, out := &in.AdjacentParentShardID, &out.AdjacentParentShardID
		*out = new(string)
		**out = **in
	}
	if in.HashKeyRange != nil {
		in, out := &in.HashKeyRange, &out.HashKeyRange
		*out = new(HashKeyRange)
		(*in).DeepCopyInto(*out)
	}
	if in.ParentShardID != nil {
		in, out := &in.ParentShardID, &out.ParentShardID
		*out = new(string)
		**out = **in
	}
	if in.SequenceNumberRange != nil {
		in, out := &in.SequenceNumberRange, &out.SequenceNumberRange
		*out = new(SequenceNumberRange)
		(*in).DeepCopyInto(*out)
	}
	if in.ShardID != nil {
		in, out := &in.ShardID, &out.ShardID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Shard.
func (in *Shard) DeepCopy() *Shard {
	if in == nil {
		return nil
	}
	out := new(Shard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShardFilter) DeepCopyInto(out *ShardFilter) {
	*out = *in
	if in.ShardID != nil {
		in, out := &in.ShardID, &out.ShardID
		*out = new(string)
		**out = **in
	}
	if in.Timestamp != nil {
		in, out := &in.Timestamp, &out.Timestamp
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShardFilter.
func (in *ShardFilter) DeepCopy() *ShardFilter {
	if in == nil {
		return nil
	}
	out := new(ShardFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StartingPosition) DeepCopyInto(out *StartingPosition) {
	*out = *in
	if in.SequenceNumber != nil {
		in, out := &in.SequenceNumber, &out.SequenceNumber
		*out = new(string)
		**out = **in
	}
	if in.Timestamp != nil {
		in, out := &in.Timestamp, &out.Timestamp
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StartingPosition.
func (in *StartingPosition) DeepCopy() *StartingPosition {
	if in == nil {
		return nil
	}
	out := new(StartingPosition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Stream) DeepCopyInto(out *Stream) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Stream.
func (in *Stream) DeepCopy() *Stream {
	if in == nil {
		return nil
	}
	out := new(Stream)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Stream) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamDescription) DeepCopyInto(out *StreamDescription) {
	*out = *in
	if in.EncryptionType != nil {
		in, out := &in.EncryptionType, &out.EncryptionType
		*out = new(string)
		**out = **in
	}
	if in.EnhancedMonitoring != nil {
		in, out := &in.EnhancedMonitoring, &out.EnhancedMonitoring
		*out = make([]*EnhancedMetrics, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(EnhancedMetrics)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.HasMoreShards != nil {
		in, out := &in.HasMoreShards, &out.HasMoreShards
		*out = new(bool)
		**out = **in
	}
	if in.KeyID != nil {
		in, out := &in.KeyID, &out.KeyID
		*out = new(string)
		**out = **in
	}
	if in.RetentionPeriodHours != nil {
		in, out := &in.RetentionPeriodHours, &out.RetentionPeriodHours
		*out = new(int64)
		**out = **in
	}
	if in.Shards != nil {
		in, out := &in.Shards, &out.Shards
		*out = make([]*Shard, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Shard)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.StreamARN != nil {
		in, out := &in.StreamARN, &out.StreamARN
		*out = new(string)
		**out = **in
	}
	if in.StreamCreationTimestamp != nil {
		in, out := &in.StreamCreationTimestamp, &out.StreamCreationTimestamp
		*out = (*in).DeepCopy()
	}
	if in.StreamModeDetails != nil {
		in, out := &in.StreamModeDetails, &out.StreamModeDetails
		*out = new(StreamModeDetails)
		(*in).DeepCopyInto(*out)
	}
	if in.StreamName != nil {
		in, out := &in.StreamName, &out.StreamName
		*out = new(string)
		**out = **in
	}
	if in.StreamStatus != nil {
		in, out := &in.StreamStatus, &out.StreamStatus
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamDescription.
func (in *StreamDescription) DeepCopy() *StreamDescription {
	if in == nil {
		return nil
	}
	out := new(StreamDescription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamDescriptionSummary) DeepCopyInto(out *StreamDescriptionSummary) {
	*out = *in
	if in.EncryptionType != nil {
		in, out := &in.EncryptionType, &out.EncryptionType
		*out = new(string)
		**out = **in
	}
	if in.EnhancedMonitoring != nil {
		in, out := &in.EnhancedMonitoring, &out.EnhancedMonitoring
		*out = make([]*EnhancedMetrics, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(EnhancedMetrics)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.KeyID != nil {
		in, out := &in.KeyID, &out.KeyID
		*out = new(string)
		**out = **in
	}
	if in.RetentionPeriodHours != nil {
		in, out := &in.RetentionPeriodHours, &out.RetentionPeriodHours
		*out = new(int64)
		**out = **in
	}
	if in.StreamARN != nil {
		in, out := &in.StreamARN, &out.StreamARN
		*out = new(string)
		**out = **in
	}
	if in.StreamCreationTimestamp != nil {
		in, out := &in.StreamCreationTimestamp, &out.StreamCreationTimestamp
		*out = (*in).DeepCopy()
	}
	if in.StreamModeDetails != nil {
		in, out := &in.StreamModeDetails, &out.StreamModeDetails
		*out = new(StreamModeDetails)
		(*in).DeepCopyInto(*out)
	}
	if in.StreamName != nil {
		in, out := &in.StreamName, &out.StreamName
		*out = new(string)
		**out = **in
	}
	if in.StreamStatus != nil {
		in, out := &in.StreamStatus, &out.StreamStatus
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamDescriptionSummary.
func (in *StreamDescriptionSummary) DeepCopy() *StreamDescriptionSummary {
	if in == nil {
		return nil
	}
	out := new(StreamDescriptionSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamList) DeepCopyInto(out *StreamList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Stream, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamList.
func (in *StreamList) DeepCopy() *StreamList {
	if in == nil {
		return nil
	}
	out := new(StreamList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StreamList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamModeDetails) DeepCopyInto(out *StreamModeDetails) {
	*out = *in
	if in.StreamMode != nil {
		in, out := &in.StreamMode, &out.StreamMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamModeDetails.
func (in *StreamModeDetails) DeepCopy() *StreamModeDetails {
	if in == nil {
		return nil
	}
	out := new(StreamModeDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamObservation) DeepCopyInto(out *StreamObservation) {
	*out = *in
	if in.EncryptionType != nil {
		in, out := &in.EncryptionType, &out.EncryptionType
		*out = new(string)
		**out = **in
	}
	if in.EnhancedMonitoring != nil {
		in, out := &in.EnhancedMonitoring, &out.EnhancedMonitoring
		*out = make([]*EnhancedMetrics, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(EnhancedMetrics)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.HasMoreShards != nil {
		in, out := &in.HasMoreShards, &out.HasMoreShards
		*out = new(bool)
		**out = **in
	}
	if in.KeyID != nil {
		in, out := &in.KeyID, &out.KeyID
		*out = new(string)
		**out = **in
	}
	if in.RetentionPeriodHours != nil {
		in, out := &in.RetentionPeriodHours, &out.RetentionPeriodHours
		*out = new(int64)
		**out = **in
	}
	if in.Shards != nil {
		in, out := &in.Shards, &out.Shards
		*out = make([]*Shard, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Shard)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.StreamARN != nil {
		in, out := &in.StreamARN, &out.StreamARN
		*out = new(string)
		**out = **in
	}
	if in.StreamStatus != nil {
		in, out := &in.StreamStatus, &out.StreamStatus
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamObservation.
func (in *StreamObservation) DeepCopy() *StreamObservation {
	if in == nil {
		return nil
	}
	out := new(StreamObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamParameters) DeepCopyInto(out *StreamParameters) {
	*out = *in
	if in.ShardCount != nil {
		in, out := &in.ShardCount, &out.ShardCount
		*out = new(int64)
		**out = **in
	}
	if in.StreamModeDetails != nil {
		in, out := &in.StreamModeDetails, &out.StreamModeDetails
		*out = new(StreamModeDetails)
		(*in).DeepCopyInto(*out)
	}
	in.CustomStreamParameters.DeepCopyInto(&out.CustomStreamParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamParameters.
func (in *StreamParameters) DeepCopy() *StreamParameters {
	if in == nil {
		return nil
	}
	out := new(StreamParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamSpec) DeepCopyInto(out *StreamSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamSpec.
func (in *StreamSpec) DeepCopy() *StreamSpec {
	if in == nil {
		return nil
	}
	out := new(StreamSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamStatus) DeepCopyInto(out *StreamStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamStatus.
func (in *StreamStatus) DeepCopy() *StreamStatus {
	if in == nil {
		return nil
	}
	out := new(StreamStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscribeToShardEvent) DeepCopyInto(out *SubscribeToShardEvent) {
	*out = *in
	if in.ContinuationSequenceNumber != nil {
		in, out := &in.ContinuationSequenceNumber, &out.ContinuationSequenceNumber
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscribeToShardEvent.
func (in *SubscribeToShardEvent) DeepCopy() *SubscribeToShardEvent {
	if in == nil {
		return nil
	}
	out := new(SubscribeToShardEvent)
	in.DeepCopyInto(out)
	return out
}
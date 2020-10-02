//
// Copyright 2020 The Feast Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: feast/core/IngestionJob.proto

package core

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Status of a Feast Ingestion Job
type IngestionJobStatus int32

const (
	// Job status is not known.
	IngestionJobStatus_UNKNOWN IngestionJobStatus = 0
	// Import job is submitted to runner and currently pending for executing
	IngestionJobStatus_PENDING IngestionJobStatus = 1
	// Import job is currently running in the runner
	IngestionJobStatus_RUNNING IngestionJobStatus = 2
	// Runner's reported the import job has completed (applicable to batch job)
	IngestionJobStatus_COMPLETED IngestionJobStatus = 3
	// When user sent abort command, but it's still running
	IngestionJobStatus_ABORTING IngestionJobStatus = 4
	// User initiated abort job
	IngestionJobStatus_ABORTED IngestionJobStatus = 5
	// Runner's reported that the import job failed to run or there is a failure during job
	IngestionJobStatus_ERROR IngestionJobStatus = 6
	// job has been suspended and waiting for cleanup
	IngestionJobStatus_SUSPENDING IngestionJobStatus = 7
	// job has been suspended
	IngestionJobStatus_SUSPENDED IngestionJobStatus = 8
)

// Enum value maps for IngestionJobStatus.
var (
	IngestionJobStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "PENDING",
		2: "RUNNING",
		3: "COMPLETED",
		4: "ABORTING",
		5: "ABORTED",
		6: "ERROR",
		7: "SUSPENDING",
		8: "SUSPENDED",
	}
	IngestionJobStatus_value = map[string]int32{
		"UNKNOWN":    0,
		"PENDING":    1,
		"RUNNING":    2,
		"COMPLETED":  3,
		"ABORTING":   4,
		"ABORTED":    5,
		"ERROR":      6,
		"SUSPENDING": 7,
		"SUSPENDED":  8,
	}
)

func (x IngestionJobStatus) Enum() *IngestionJobStatus {
	p := new(IngestionJobStatus)
	*p = x
	return p
}

func (x IngestionJobStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IngestionJobStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_feast_core_IngestionJob_proto_enumTypes[0].Descriptor()
}

func (IngestionJobStatus) Type() protoreflect.EnumType {
	return &file_feast_core_IngestionJob_proto_enumTypes[0]
}

func (x IngestionJobStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IngestionJobStatus.Descriptor instead.
func (IngestionJobStatus) EnumDescriptor() ([]byte, []int) {
	return file_feast_core_IngestionJob_proto_rawDescGZIP(), []int{0}
}

// Represents Feast Injestion Job
type IngestionJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Job ID assigned by Feast
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// External job ID specific to the runner.
	// For DirectRunner jobs, this is identical to id. For DataflowRunner jobs, this refers to the Dataflow job ID.
	ExternalId string             `protobuf:"bytes,2,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	Status     IngestionJobStatus `protobuf:"varint,3,opt,name=status,proto3,enum=feast.core.IngestionJobStatus" json:"status,omitempty"`
	// Source this job is reading from.
	Source *Source `protobuf:"bytes,5,opt,name=source,proto3" json:"source,omitempty"`
	// Store this job is writing to.
	Stores []*Store `protobuf:"bytes,6,rep,name=stores,proto3" json:"stores,omitempty"`
	// List of Feature Set References
	FeatureSetReferences []*FeatureSetReference `protobuf:"bytes,7,rep,name=feature_set_references,json=featureSetReferences,proto3" json:"feature_set_references,omitempty"`
}

func (x *IngestionJob) Reset() {
	*x = IngestionJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feast_core_IngestionJob_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IngestionJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IngestionJob) ProtoMessage() {}

func (x *IngestionJob) ProtoReflect() protoreflect.Message {
	mi := &file_feast_core_IngestionJob_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IngestionJob.ProtoReflect.Descriptor instead.
func (*IngestionJob) Descriptor() ([]byte, []int) {
	return file_feast_core_IngestionJob_proto_rawDescGZIP(), []int{0}
}

func (x *IngestionJob) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *IngestionJob) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *IngestionJob) GetStatus() IngestionJobStatus {
	if x != nil {
		return x.Status
	}
	return IngestionJobStatus_UNKNOWN
}

func (x *IngestionJob) GetSource() *Source {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *IngestionJob) GetStores() []*Store {
	if x != nil {
		return x.Stores
	}
	return nil
}

func (x *IngestionJob) GetFeatureSetReferences() []*FeatureSetReference {
	if x != nil {
		return x.FeatureSetReferences
	}
	return nil
}

// Config for bi-directional communication channel between Core Service and Ingestion Job
type SpecsStreamingUpdateConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// out-channel for publishing new FeatureSetSpecs (by Core).
	// IngestionJob use it as source of existing FeatureSetSpecs and new real-time updates
	Source *KafkaSourceConfig `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	// ack-channel for sending acknowledgments when new FeatureSetSpecs is installed in Job
	Ack *KafkaSourceConfig `protobuf:"bytes,2,opt,name=ack,proto3" json:"ack,omitempty"`
}

func (x *SpecsStreamingUpdateConfig) Reset() {
	*x = SpecsStreamingUpdateConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feast_core_IngestionJob_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpecsStreamingUpdateConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpecsStreamingUpdateConfig) ProtoMessage() {}

func (x *SpecsStreamingUpdateConfig) ProtoReflect() protoreflect.Message {
	mi := &file_feast_core_IngestionJob_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpecsStreamingUpdateConfig.ProtoReflect.Descriptor instead.
func (*SpecsStreamingUpdateConfig) Descriptor() ([]byte, []int) {
	return file_feast_core_IngestionJob_proto_rawDescGZIP(), []int{1}
}

func (x *SpecsStreamingUpdateConfig) GetSource() *KafkaSourceConfig {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *SpecsStreamingUpdateConfig) GetAck() *KafkaSourceConfig {
	if x != nil {
		return x.Ack
	}
	return nil
}

type FeatureSetSpecAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FeatureSetReference string `protobuf:"bytes,1,opt,name=feature_set_reference,json=featureSetReference,proto3" json:"feature_set_reference,omitempty"`
	FeatureSetVersion   int32  `protobuf:"varint,2,opt,name=feature_set_version,json=featureSetVersion,proto3" json:"feature_set_version,omitempty"`
	JobName             string `protobuf:"bytes,3,opt,name=job_name,json=jobName,proto3" json:"job_name,omitempty"`
}

func (x *FeatureSetSpecAck) Reset() {
	*x = FeatureSetSpecAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feast_core_IngestionJob_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureSetSpecAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureSetSpecAck) ProtoMessage() {}

func (x *FeatureSetSpecAck) ProtoReflect() protoreflect.Message {
	mi := &file_feast_core_IngestionJob_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureSetSpecAck.ProtoReflect.Descriptor instead.
func (*FeatureSetSpecAck) Descriptor() ([]byte, []int) {
	return file_feast_core_IngestionJob_proto_rawDescGZIP(), []int{2}
}

func (x *FeatureSetSpecAck) GetFeatureSetReference() string {
	if x != nil {
		return x.FeatureSetReference
	}
	return ""
}

func (x *FeatureSetSpecAck) GetFeatureSetVersion() int32 {
	if x != nil {
		return x.FeatureSetVersion
	}
	return 0
}

func (x *FeatureSetSpecAck) GetJobName() string {
	if x != nil {
		return x.JobName
	}
	return ""
}

var File_feast_core_IngestionJob_proto protoreflect.FileDescriptor

var file_feast_core_IngestionJob_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x49, 0x6e, 0x67,
	0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x24, 0x66, 0x65, 0x61,
	0x73, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53,
	0x65, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x16, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x66, 0x65, 0x61, 0x73, 0x74,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xab, 0x02, 0x0a, 0x0c, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x4a, 0x6f, 0x62, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x66,
	0x65, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x06, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x73, 0x12, 0x55, 0x0a, 0x16, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x73,
	0x65, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x52, 0x14, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x74,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05,
	0x22, 0x84, 0x01, 0x0a, 0x1a, 0x53, 0x70, 0x65, 0x63, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x35, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4b, 0x61, 0x66,
	0x6b, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x03, 0x61, 0x63, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x03, 0x61, 0x63, 0x6b, 0x22, 0x92, 0x01, 0x0a, 0x11, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x53, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x41, 0x63, 0x6b, 0x12, 0x32, 0x0a,
	0x15, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x12, 0x2e, 0x0a, 0x13, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x74,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x19, 0x0a, 0x08, 0x6a, 0x6f, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x6f, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x2a, 0x8f, 0x01, 0x0a,
	0x12, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f,
	0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x42, 0x4f,
	0x52, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x42, 0x4f, 0x52, 0x54,
	0x45, 0x44, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x06, 0x12,
	0x0e, 0x0a, 0x0a, 0x53, 0x55, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x07, 0x12,
	0x0d, 0x0a, 0x09, 0x53, 0x55, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x08, 0x42, 0x5a,
	0x0a, 0x10, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x42, 0x11, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6f, 0x62,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x66, 0x65, 0x61, 0x73,
	0x74, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x66, 0x65, 0x61, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_feast_core_IngestionJob_proto_rawDescOnce sync.Once
	file_feast_core_IngestionJob_proto_rawDescData = file_feast_core_IngestionJob_proto_rawDesc
)

func file_feast_core_IngestionJob_proto_rawDescGZIP() []byte {
	file_feast_core_IngestionJob_proto_rawDescOnce.Do(func() {
		file_feast_core_IngestionJob_proto_rawDescData = protoimpl.X.CompressGZIP(file_feast_core_IngestionJob_proto_rawDescData)
	})
	return file_feast_core_IngestionJob_proto_rawDescData
}

var file_feast_core_IngestionJob_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_feast_core_IngestionJob_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_feast_core_IngestionJob_proto_goTypes = []interface{}{
	(IngestionJobStatus)(0),            // 0: feast.core.IngestionJobStatus
	(*IngestionJob)(nil),               // 1: feast.core.IngestionJob
	(*SpecsStreamingUpdateConfig)(nil), // 2: feast.core.SpecsStreamingUpdateConfig
	(*FeatureSetSpecAck)(nil),          // 3: feast.core.FeatureSetSpecAck
	(*Source)(nil),                     // 4: feast.core.Source
	(*Store)(nil),                      // 5: feast.core.Store
	(*FeatureSetReference)(nil),        // 6: feast.core.FeatureSetReference
	(*KafkaSourceConfig)(nil),          // 7: feast.core.KafkaSourceConfig
}
var file_feast_core_IngestionJob_proto_depIdxs = []int32{
	0, // 0: feast.core.IngestionJob.status:type_name -> feast.core.IngestionJobStatus
	4, // 1: feast.core.IngestionJob.source:type_name -> feast.core.Source
	5, // 2: feast.core.IngestionJob.stores:type_name -> feast.core.Store
	6, // 3: feast.core.IngestionJob.feature_set_references:type_name -> feast.core.FeatureSetReference
	7, // 4: feast.core.SpecsStreamingUpdateConfig.source:type_name -> feast.core.KafkaSourceConfig
	7, // 5: feast.core.SpecsStreamingUpdateConfig.ack:type_name -> feast.core.KafkaSourceConfig
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_feast_core_IngestionJob_proto_init() }
func file_feast_core_IngestionJob_proto_init() {
	if File_feast_core_IngestionJob_proto != nil {
		return
	}
	file_feast_core_FeatureSetReference_proto_init()
	file_feast_core_Store_proto_init()
	file_feast_core_Source_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_feast_core_IngestionJob_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IngestionJob); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_feast_core_IngestionJob_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpecsStreamingUpdateConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_feast_core_IngestionJob_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureSetSpecAck); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_feast_core_IngestionJob_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_feast_core_IngestionJob_proto_goTypes,
		DependencyIndexes: file_feast_core_IngestionJob_proto_depIdxs,
		EnumInfos:         file_feast_core_IngestionJob_proto_enumTypes,
		MessageInfos:      file_feast_core_IngestionJob_proto_msgTypes,
	}.Build()
	File_feast_core_IngestionJob_proto = out.File
	file_feast_core_IngestionJob_proto_rawDesc = nil
	file_feast_core_IngestionJob_proto_goTypes = nil
	file_feast_core_IngestionJob_proto_depIdxs = nil
}
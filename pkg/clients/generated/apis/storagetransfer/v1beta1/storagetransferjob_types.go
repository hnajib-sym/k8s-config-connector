// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type JobAccessKeyId struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the field's value. Cannot be used if 'value' is specified. */
	// +optional
	ValueFrom *JobValueFrom `json:"valueFrom,omitempty"`
}

type JobAwsAccessKey struct {
	/* AWS Key ID. */
	AccessKeyId JobAccessKeyId `json:"accessKeyId"`

	/* AWS Secret Access Key. */
	SecretAccessKey JobSecretAccessKey `json:"secretAccessKey"`
}

type JobAwsS3DataSource struct {
	/* AWS credentials block. */
	// +optional
	AwsAccessKey *JobAwsAccessKey `json:"awsAccessKey,omitempty"`

	/* S3 Bucket name. */
	BucketName string `json:"bucketName"`

	/* The Amazon Resource Name (ARN) of the role to support temporary credentials via 'AssumeRoleWithWebIdentity'. For more information about ARNs, see [IAM ARNs](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-arns). When a role ARN is provided, Transfer Service fetches temporary credentials for the session using a 'AssumeRoleWithWebIdentity' call for the provided role using the [GoogleServiceAccount][] for this project. */
	// +optional
	RoleArn *string `json:"roleArn,omitempty"`
}

type JobAzureBlobStorageDataSource struct {
	/* Credentials used to authenticate API requests to Azure. */
	AzureCredentials JobAzureCredentials `json:"azureCredentials"`

	/* The container to transfer from the Azure Storage account. */
	Container string `json:"container"`

	/* Root path to transfer objects. Must be an empty string or full path name that ends with a '/'. This field is treated as an object prefix. As such, it should generally not begin with a '/'. */
	// +optional
	Path *string `json:"path,omitempty"`

	/* The name of the Azure Storage account. */
	StorageAccount string `json:"storageAccount"`
}

type JobAzureCredentials struct {
	/* Azure shared access signature. */
	SasToken JobSasToken `json:"sasToken"`
}

type JobGcsDataSink struct {
	BucketRef v1alpha1.ResourceRef `json:"bucketRef"`

	/* Google Cloud Storage path in bucket to transfer. */
	// +optional
	Path *string `json:"path,omitempty"`
}

type JobGcsDataSource struct {
	BucketRef v1alpha1.ResourceRef `json:"bucketRef"`

	/* Google Cloud Storage path in bucket to transfer. */
	// +optional
	Path *string `json:"path,omitempty"`
}

type JobHttpDataSource struct {
	/* The URL that points to the file that stores the object list entries. This file must allow public access. Currently, only URLs with HTTP and HTTPS schemes are supported. */
	ListUrl string `json:"listUrl"`
}

type JobNotificationConfig struct {
	/* Event types for which a notification is desired. If empty, send notifications for all event types. The valid types are "TRANSFER_OPERATION_SUCCESS", "TRANSFER_OPERATION_FAILED", "TRANSFER_OPERATION_ABORTED". */
	// +optional
	EventTypes []string `json:"eventTypes,omitempty"`

	/* The desired format of the notification message payloads. One of "NONE" or "JSON". */
	PayloadFormat string `json:"payloadFormat"`

	/* The PubSubTopic to which to publish notifications. */
	TopicRef v1alpha1.ResourceRef `json:"topicRef"`
}

type JobObjectConditions struct {
	/* exclude_prefixes must follow the requirements described for include_prefixes. */
	// +optional
	ExcludePrefixes []string `json:"excludePrefixes,omitempty"`

	/* If include_refixes is specified, objects that satisfy the object conditions must have names that start with one of the include_prefixes and that do not start with any of the exclude_prefixes. If include_prefixes is not specified, all objects except those that have names starting with one of the exclude_prefixes must satisfy the object conditions. */
	// +optional
	IncludePrefixes []string `json:"includePrefixes,omitempty"`

	/* A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s". */
	// +optional
	MaxTimeElapsedSinceLastModification *string `json:"maxTimeElapsedSinceLastModification,omitempty"`

	/* A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s". */
	// +optional
	MinTimeElapsedSinceLastModification *string `json:"minTimeElapsedSinceLastModification,omitempty"`
}

type JobPosixDataSink struct {
	/* Root directory path to the filesystem. */
	RootDirectory string `json:"rootDirectory"`
}

type JobPosixDataSource struct {
	/* Root directory path to the filesystem. */
	RootDirectory string `json:"rootDirectory"`
}

type JobSasToken struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the field's value. Cannot be used if 'value' is specified. */
	// +optional
	ValueFrom *JobValueFrom `json:"valueFrom,omitempty"`
}

type JobSchedule struct {
	/* Interval between the start of each scheduled transfer. If unspecified, the default value is 24 hours. This value may not be less than 1 hour. A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s". */
	// +optional
	RepeatInterval *string `json:"repeatInterval,omitempty"`

	/* The last day the recurring transfer will be run. If schedule_end_date is the same as schedule_start_date, the transfer will be executed only once. */
	// +optional
	ScheduleEndDate *JobScheduleEndDate `json:"scheduleEndDate,omitempty"`

	/* The first day the recurring transfer is scheduled to run. If schedule_start_date is in the past, the transfer will run for the first time on the following day. */
	ScheduleStartDate JobScheduleStartDate `json:"scheduleStartDate"`

	/* The time in UTC at which the transfer will be scheduled to start in a day. Transfers may start later than this time. If not specified, recurring and one-time transfers that are scheduled to run today will run immediately; recurring transfers that are scheduled to run on a future date will start at approximately midnight UTC on that date. Note that when configuring a transfer with the Cloud Platform Console, the transfer's start time in a day is specified in your local timezone. */
	// +optional
	StartTimeOfDay *JobStartTimeOfDay `json:"startTimeOfDay,omitempty"`
}

type JobScheduleEndDate struct {
	/* Day of month. Must be from 1 to 31 and valid for the year and month. */
	Day int `json:"day"`

	/* Month of year. Must be from 1 to 12. */
	Month int `json:"month"`

	/* Year of date. Must be from 1 to 9999. */
	Year int `json:"year"`
}

type JobScheduleStartDate struct {
	/* Day of month. Must be from 1 to 31 and valid for the year and month. */
	Day int `json:"day"`

	/* Month of year. Must be from 1 to 12. */
	Month int `json:"month"`

	/* Year of date. Must be from 1 to 9999. */
	Year int `json:"year"`
}

type JobSecretAccessKey struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the field's value. Cannot be used if 'value' is specified. */
	// +optional
	ValueFrom *JobValueFrom `json:"valueFrom,omitempty"`
}

type JobStartTimeOfDay struct {
	/* Hours of day in 24 hour format. Should be from 0 to 23. */
	Hours int `json:"hours"`

	/* Minutes of hour of day. Must be from 0 to 59. */
	Minutes int `json:"minutes"`

	/* Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999. */
	Nanos int `json:"nanos"`

	/* Seconds of minutes of the time. Must normally be from 0 to 59. */
	Seconds int `json:"seconds"`
}

type JobTransferOptions struct {
	/* Whether objects should be deleted from the source after they are transferred to the sink. Note that this option and delete_objects_unique_in_sink are mutually exclusive. */
	// +optional
	DeleteObjectsFromSourceAfterTransfer *bool `json:"deleteObjectsFromSourceAfterTransfer,omitempty"`

	/* Whether objects that exist only in the sink should be deleted. Note that this option and delete_objects_from_source_after_transfer are mutually exclusive. */
	// +optional
	DeleteObjectsUniqueInSink *bool `json:"deleteObjectsUniqueInSink,omitempty"`

	/* Whether overwriting objects that already exist in the sink is allowed. */
	// +optional
	OverwriteObjectsAlreadyExistingInSink *bool `json:"overwriteObjectsAlreadyExistingInSink,omitempty"`

	/* When to overwrite objects that already exist in the sink. If not set, overwrite behavior is determined by overwriteObjectsAlreadyExistingInSink. */
	// +optional
	OverwriteWhen *string `json:"overwriteWhen,omitempty"`
}

type JobTransferSpec struct {
	/* An AWS S3 data source. */
	// +optional
	AwsS3DataSource *JobAwsS3DataSource `json:"awsS3DataSource,omitempty"`

	/* An Azure Blob Storage data source. */
	// +optional
	AzureBlobStorageDataSource *JobAzureBlobStorageDataSource `json:"azureBlobStorageDataSource,omitempty"`

	/* A Google Cloud Storage data sink. */
	// +optional
	GcsDataSink *JobGcsDataSink `json:"gcsDataSink,omitempty"`

	/* A Google Cloud Storage data source. */
	// +optional
	GcsDataSource *JobGcsDataSource `json:"gcsDataSource,omitempty"`

	/* A HTTP URL data source. */
	// +optional
	HttpDataSource *JobHttpDataSource `json:"httpDataSource,omitempty"`

	/* Only objects that satisfy these object conditions are included in the set of data source and data sink objects. Object conditions based on objects' last_modification_time do not exclude objects in a data sink. */
	// +optional
	ObjectConditions *JobObjectConditions `json:"objectConditions,omitempty"`

	/* A POSIX filesystem data sink. */
	// +optional
	PosixDataSink *JobPosixDataSink `json:"posixDataSink,omitempty"`

	/* A POSIX filesystem data source. */
	// +optional
	PosixDataSource *JobPosixDataSource `json:"posixDataSource,omitempty"`

	/* Characteristics of how to treat files from datasource and sink during job. If the option delete_objects_unique_in_sink is true, object conditions based on objects' last_modification_time are ignored and do not exclude objects in a data source or a data sink. */
	// +optional
	TransferOptions *JobTransferOptions `json:"transferOptions,omitempty"`
}

type JobValueFrom struct {
	/* Reference to a value with the given key in the given Secret in the resource's namespace. */
	// +optional
	SecretKeyRef *v1alpha1.ResourceRef `json:"secretKeyRef,omitempty"`
}

type StorageTransferJobSpec struct {
	/* Unique description to identify the Transfer Job. */
	Description string `json:"description"`

	/* Notification configuration. */
	// +optional
	NotificationConfig *JobNotificationConfig `json:"notificationConfig,omitempty"`

	/* Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run. */
	// +optional
	Schedule *JobSchedule `json:"schedule,omitempty"`

	/* Status of the job. Default: ENABLED. NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation. */
	// +optional
	Status *string `json:"status,omitempty"`

	/* Transfer specification. */
	TransferSpec JobTransferSpec `json:"transferSpec"`
}

type StorageTransferJobStatus struct {
	/* Conditions represent the latest available observations of the
	   StorageTransferJob's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* When the Transfer Job was created. */
	// +optional
	CreationTime *string `json:"creationTime,omitempty"`

	/* When the Transfer Job was deleted. */
	// +optional
	DeletionTime *string `json:"deletionTime,omitempty"`

	/* When the Transfer Job was last modified. */
	// +optional
	LastModificationTime *string `json:"lastModificationTime,omitempty"`

	/* The name of the Transfer Job. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StorageTransferJob is the Schema for the storagetransfer API
// +k8s:openapi-gen=true
type StorageTransferJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StorageTransferJobSpec   `json:"spec,omitempty"`
	Status StorageTransferJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StorageTransferJobList contains a list of StorageTransferJob
type StorageTransferJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageTransferJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StorageTransferJob{}, &StorageTransferJobList{})
}

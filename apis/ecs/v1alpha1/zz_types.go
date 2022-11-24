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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
)

// +kubebuilder:skipversion
type AWSVPCConfiguration struct {
	AssignPublicIP *string `json:"assignPublicIP,omitempty"`

	SecurityGroups []*string `json:"securityGroups,omitempty"`

	Subnets []*string `json:"subnets,omitempty"`
}

// +kubebuilder:skipversion
type Attachment struct {
	Details []*KeyValuePair `json:"details,omitempty"`

	ID *string `json:"id,omitempty"`

	Status *string `json:"status,omitempty"`

	Type *string `json:"type_,omitempty"`
}

// +kubebuilder:skipversion
type AttachmentStateChange struct {
	AttachmentARN *string `json:"attachmentARN,omitempty"`

	Status *string `json:"status,omitempty"`
}

// +kubebuilder:skipversion
type Attribute struct {
	Name *string `json:"name,omitempty"`

	TargetID *string `json:"targetID,omitempty"`

	TargetType *string `json:"targetType,omitempty"`

	Value *string `json:"value,omitempty"`
}

// +kubebuilder:skipversion
type AutoScalingGroupProvider struct {
	AutoScalingGroupARN *string `json:"autoScalingGroupARN,omitempty"`
}

// +kubebuilder:skipversion
type CapacityProvider struct {
	CapacityProviderARN *string `json:"capacityProviderARN,omitempty"`

	Name *string `json:"name,omitempty"`

	Tags []*Tag `json:"tags,omitempty"`

	UpdateStatusReason *string `json:"updateStatusReason,omitempty"`
}

// +kubebuilder:skipversion
type CapacityProviderStrategyItem struct {
	Base *int64 `json:"base,omitempty"`

	CapacityProvider *string `json:"capacityProvider,omitempty"`

	Weight *int64 `json:"weight,omitempty"`
}

// +kubebuilder:skipversion
type ClusterConfiguration struct {
	// The details of the execute command configuration.
	ExecuteCommandConfiguration *ExecuteCommandConfiguration `json:"executeCommandConfiguration,omitempty"`
}

// +kubebuilder:skipversion
type ClusterSetting struct {
	Name *string `json:"name,omitempty"`

	Value *string `json:"value,omitempty"`
}

// +kubebuilder:skipversion
type Cluster_SDK struct {
	ActiveServicesCount *int64 `json:"activeServicesCount,omitempty"`

	Attachments []*Attachment `json:"attachments,omitempty"`

	AttachmentsStatus *string `json:"attachmentsStatus,omitempty"`

	CapacityProviders []*string `json:"capacityProviders,omitempty"`

	ClusterARN *string `json:"clusterARN,omitempty"`

	ClusterName *string `json:"clusterName,omitempty"`
	// The execute command configuration for the cluster.
	Configuration *ClusterConfiguration `json:"configuration,omitempty"`

	DefaultCapacityProviderStrategy []*CapacityProviderStrategyItem `json:"defaultCapacityProviderStrategy,omitempty"`

	PendingTasksCount *int64 `json:"pendingTasksCount,omitempty"`

	RegisteredContainerInstancesCount *int64 `json:"registeredContainerInstancesCount,omitempty"`

	RunningTasksCount *int64 `json:"runningTasksCount,omitempty"`

	Settings []*ClusterSetting `json:"settings,omitempty"`

	Statistics []*KeyValuePair `json:"statistics,omitempty"`

	Status *string `json:"status,omitempty"`

	Tags []*Tag `json:"tags,omitempty"`
}

// +kubebuilder:skipversion
type Container struct {
	ContainerARN *string `json:"containerARN,omitempty"`

	CPU *string `json:"cpu,omitempty"`

	ExitCode *int64 `json:"exitCode,omitempty"`

	Image *string `json:"image,omitempty"`

	ImageDigest *string `json:"imageDigest,omitempty"`

	LastStatus *string `json:"lastStatus,omitempty"`

	Memory *string `json:"memory,omitempty"`

	MemoryReservation *string `json:"memoryReservation,omitempty"`

	Name *string `json:"name,omitempty"`

	Reason *string `json:"reason,omitempty"`

	RuntimeID *string `json:"runtimeID,omitempty"`

	TaskARN *string `json:"taskARN,omitempty"`
}

// +kubebuilder:skipversion
type ContainerDefinition struct {
	Command []*string `json:"command,omitempty"`

	CPU *int64 `json:"cpu,omitempty"`

	DependsOn []*ContainerDependency `json:"dependsOn,omitempty"`

	DisableNetworking *bool `json:"disableNetworking,omitempty"`

	DNSSearchDomains []*string `json:"dnsSearchDomains,omitempty"`

	DNSServers []*string `json:"dnsServers,omitempty"`

	DockerLabels map[string]*string `json:"dockerLabels,omitempty"`

	DockerSecurityOptions []*string `json:"dockerSecurityOptions,omitempty"`

	EntryPoint []*string `json:"entryPoint,omitempty"`

	Environment []*KeyValuePair `json:"environment,omitempty"`

	EnvironmentFiles []*EnvironmentFile `json:"environmentFiles,omitempty"`

	Essential *bool `json:"essential,omitempty"`

	ExtraHosts []*HostEntry `json:"extraHosts,omitempty"`
	// The FireLens configuration for the container. This is used to specify and
	// configure a log router for container logs. For more information, see Custom
	// Log Routing (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html)
	// in the Amazon Elastic Container Service Developer Guide.
	FirelensConfiguration *FirelensConfiguration `json:"firelensConfiguration,omitempty"`
	// An object representing a container health check. Health check parameters
	// that are specified in a container definition override any Docker health checks
	// that exist in the container image (such as those specified in a parent image
	// or from the image's Dockerfile).
	//
	// The Amazon ECS container agent only monitors and reports on the health checks
	// specified in the task definition. Amazon ECS does not monitor Docker health
	// checks that are embedded in a container image and not specified in the container
	// definition. Health check parameters that are specified in a container definition
	// override any Docker health checks that exist in the container image.
	//
	// You can view the health status of both individual containers and a task with
	// the DescribeTasks API operation or when viewing the task details in the console.
	//
	// The following describes the possible healthStatus values for a container:
	//
	//    * HEALTHY-The container health check has passed successfully.
	//
	//    * UNHEALTHY-The container health check has failed.
	//
	//    * UNKNOWN-The container health check is being evaluated or there's no
	//    container health check defined.
	//
	// The following describes the possible healthStatus values for a task. The
	// container health check status of nonessential containers do not have an effect
	// on the health status of a task.
	//
	//    * HEALTHY-All essential containers within the task have passed their health
	//    checks.
	//
	//    * UNHEALTHY-One or more essential containers have failed their health
	//    check.
	//
	//    * UNKNOWN-The essential containers within the task are still having their
	//    health checks evaluated or there are no container health checks defined.
	//
	// If a task is run manually, and not as part of a service, the task will continue
	// its lifecycle regardless of its health status. For tasks that are part of
	// a service, if the task reports as unhealthy then the task will be stopped
	// and the service scheduler will replace it.
	//
	// The following are notes about container health check support:
	//
	//    * Container health checks require version 1.17.0 or greater of the Amazon
	//    ECS container agent. For more information, see Updating the Amazon ECS
	//    Container Agent (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-update.html).
	//
	//    * Container health checks are supported for Fargate tasks if you're using
	//    platform version 1.1.0 or greater. For more information, see Fargate Platform
	//    Versions (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html).
	//
	//    * Container health checks aren't supported for tasks that are part of
	//    a service that's configured to use a Classic Load Balancer.
	HealthCheck *HealthCheck `json:"healthCheck,omitempty"`

	Hostname *string `json:"hostname,omitempty"`

	Image *string `json:"image,omitempty"`

	Interactive *bool `json:"interactive,omitempty"`

	Links []*string `json:"links,omitempty"`
	// Linux-specific options that are applied to the container, such as Linux KernelCapabilities.
	LinuxParameters *LinuxParameters `json:"linuxParameters,omitempty"`
	// The log configuration for the container. This parameter maps to LogConfig
	// in the Create a container (https://docs.docker.com/engine/api/v1.35/#operation/ContainerCreate)
	// section of the Docker Remote API (https://docs.docker.com/engine/api/v1.35/)
	// and the --log-driver option to docker run (https://docs.docker.com/engine/reference/commandline/run/).
	//
	// By default, containers use the same logging driver that the Docker daemon
	// uses. However, the container might use a different logging driver than the
	// Docker daemon by specifying a log driver configuration in the container definition.
	// For more information about the options for different supported log drivers,
	// see Configure logging drivers (https://docs.docker.com/engine/admin/logging/overview/)
	// in the Docker documentation.
	//
	// Understand the following when specifying a log configuration for your containers.
	//
	//    * Amazon ECS currently supports a subset of the logging drivers available
	//    to the Docker daemon (shown in the valid values below). Additional log
	//    drivers may be available in future releases of the Amazon ECS container
	//    agent.
	//
	//    * This parameter requires version 1.18 of the Docker Remote API or greater
	//    on your container instance.
	//
	//    * For tasks that are hosted on Amazon EC2 instances, the Amazon ECS container
	//    agent must register the available logging drivers with the ECS_AVAILABLE_LOGGING_DRIVERS
	//    environment variable before containers placed on that instance can use
	//    these log configuration options. For more information, see Amazon ECS
	//    container agent configuration (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-agent-config.html)
	//    in the Amazon Elastic Container Service Developer Guide.
	//
	//    * For tasks that are on Fargate, because you don't have access to the
	//    underlying infrastructure your tasks are hosted on, any additional software
	//    needed must be installed outside of the task. For example, the Fluentd
	//    output aggregators or a remote host running Logstash to send Gelf logs
	//    to.
	LogConfiguration *LogConfiguration `json:"logConfiguration,omitempty"`

	Memory *int64 `json:"memory,omitempty"`

	MemoryReservation *int64 `json:"memoryReservation,omitempty"`

	MountPoints []*MountPoint `json:"mountPoints,omitempty"`

	Name *string `json:"name,omitempty"`

	PortMappings []*PortMapping `json:"portMappings,omitempty"`

	Privileged *bool `json:"privileged,omitempty"`

	PseudoTerminal *bool `json:"pseudoTerminal,omitempty"`

	ReadonlyRootFilesystem *bool `json:"readonlyRootFilesystem,omitempty"`
	// The repository credentials for private registry authentication.
	RepositoryCredentials *RepositoryCredentials `json:"repositoryCredentials,omitempty"`

	ResourceRequirements []*ResourceRequirement `json:"resourceRequirements,omitempty"`

	Secrets []*Secret `json:"secrets,omitempty"`

	StartTimeout *int64 `json:"startTimeout,omitempty"`

	StopTimeout *int64 `json:"stopTimeout,omitempty"`

	SystemControls []*SystemControl `json:"systemControls,omitempty"`

	Ulimits []*Ulimit `json:"ulimits,omitempty"`

	User *string `json:"user,omitempty"`

	VolumesFrom []*VolumeFrom `json:"volumesFrom,omitempty"`

	WorkingDirectory *string `json:"workingDirectory,omitempty"`
}

// +kubebuilder:skipversion
type ContainerDependency struct {
	Condition *string `json:"condition,omitempty"`

	ContainerName *string `json:"containerName,omitempty"`
}

// +kubebuilder:skipversion
type ContainerInstance struct {
	AgentConnected *bool `json:"agentConnected,omitempty"`

	Attachments []*Attachment `json:"attachments,omitempty"`

	CapacityProviderName *string `json:"capacityProviderName,omitempty"`

	ContainerInstanceARN *string `json:"containerInstanceARN,omitempty"`

	EC2InstanceID *string `json:"ec2InstanceID,omitempty"`

	PendingTasksCount *int64 `json:"pendingTasksCount,omitempty"`

	RegisteredAt *metav1.Time `json:"registeredAt,omitempty"`

	RunningTasksCount *int64 `json:"runningTasksCount,omitempty"`

	Status *string `json:"status,omitempty"`

	StatusReason *string `json:"statusReason,omitempty"`

	Tags []*Tag `json:"tags,omitempty"`
}

// +kubebuilder:skipversion
type ContainerOverride struct {
	Command []*string `json:"command,omitempty"`

	CPU *int64 `json:"cpu,omitempty"`

	Environment []*KeyValuePair `json:"environment,omitempty"`

	EnvironmentFiles []*EnvironmentFile `json:"environmentFiles,omitempty"`

	Memory *int64 `json:"memory,omitempty"`

	MemoryReservation *int64 `json:"memoryReservation,omitempty"`

	Name *string `json:"name,omitempty"`

	ResourceRequirements []*ResourceRequirement `json:"resourceRequirements,omitempty"`
}

// +kubebuilder:skipversion
type ContainerStateChange struct {
	ContainerName *string `json:"containerName,omitempty"`

	ExitCode *int64 `json:"exitCode,omitempty"`

	ImageDigest *string `json:"imageDigest,omitempty"`

	Reason *string `json:"reason,omitempty"`

	RuntimeID *string `json:"runtimeID,omitempty"`

	Status *string `json:"status,omitempty"`
}

// +kubebuilder:skipversion
type Deployment struct {
	CapacityProviderStrategy []*CapacityProviderStrategyItem `json:"capacityProviderStrategy,omitempty"`

	CreatedAt *metav1.Time `json:"createdAt,omitempty"`

	DesiredCount *int64 `json:"desiredCount,omitempty"`

	FailedTasks *int64 `json:"failedTasks,omitempty"`

	ID *string `json:"id,omitempty"`

	LaunchType *string `json:"launchType,omitempty"`
	// An object representing the network configuration for a task or service.
	NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty"`

	PendingCount *int64 `json:"pendingCount,omitempty"`

	PlatformFamily *string `json:"platformFamily,omitempty"`

	PlatformVersion *string `json:"platformVersion,omitempty"`

	RolloutState *string `json:"rolloutState,omitempty"`

	RolloutStateReason *string `json:"rolloutStateReason,omitempty"`

	RunningCount *int64 `json:"runningCount,omitempty"`

	Status *string `json:"status,omitempty"`

	TaskDefinition *string `json:"taskDefinition,omitempty"`

	UpdatedAt *metav1.Time `json:"updatedAt,omitempty"`
}

// +kubebuilder:skipversion
type DeploymentCircuitBreaker struct {
	Enable *bool `json:"enable,omitempty"`

	Rollback *bool `json:"rollback,omitempty"`
}

// +kubebuilder:skipversion
type DeploymentConfiguration struct {
	//
	// The deployment circuit breaker can only be used for services using the rolling
	// update (ECS) deployment type that aren't behind a Classic Load Balancer.
	//
	// The deployment circuit breaker determines whether a service deployment will
	// fail if the service can't reach a steady state. If enabled, a service deployment
	// will transition to a failed state and stop launching new tasks. You can also
	// configure Amazon ECS to roll back your service to the last completed deployment
	// after a failure. For more information, see Rolling update (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-type-ecs.html)
	// in the Amazon Elastic Container Service Developer Guide.
	DeploymentCircuitBreaker *DeploymentCircuitBreaker `json:"deploymentCircuitBreaker,omitempty"`

	MaximumPercent *int64 `json:"maximumPercent,omitempty"`

	MinimumHealthyPercent *int64 `json:"minimumHealthyPercent,omitempty"`
}

// +kubebuilder:skipversion
type DeploymentController struct {
	Type *string `json:"type_,omitempty"`
}

// +kubebuilder:skipversion
type Device struct {
	ContainerPath *string `json:"containerPath,omitempty"`

	HostPath *string `json:"hostPath,omitempty"`

	Permissions []*string `json:"permissions,omitempty"`
}

// +kubebuilder:skipversion
type DockerVolumeConfiguration struct {
	Autoprovision *bool `json:"autoprovision,omitempty"`

	Driver *string `json:"driver,omitempty"`

	DriverOpts map[string]*string `json:"driverOpts,omitempty"`

	Labels map[string]*string `json:"labels,omitempty"`

	Scope *string `json:"scope,omitempty"`
}

// +kubebuilder:skipversion
type EFSAuthorizationConfig struct {
	AccessPointID *string `json:"accessPointID,omitempty"`

	IAM *string `json:"iam,omitempty"`
}

// +kubebuilder:skipversion
type EFSVolumeConfiguration struct {
	// The authorization configuration details for the Amazon EFS file system.
	AuthorizationConfig *EFSAuthorizationConfig `json:"authorizationConfig,omitempty"`

	FileSystemID *string `json:"fileSystemID,omitempty"`

	RootDirectory *string `json:"rootDirectory,omitempty"`

	TransitEncryption *string `json:"transitEncryption,omitempty"`

	TransitEncryptionPort *int64 `json:"transitEncryptionPort,omitempty"`
}

// +kubebuilder:skipversion
type EnvironmentFile struct {
	Type *string `json:"type_,omitempty"`

	Value *string `json:"value,omitempty"`
}

// +kubebuilder:skipversion
type EphemeralStorage struct {
	SizeInGiB *int64 `json:"sizeInGiB,omitempty"`
}

// +kubebuilder:skipversion
type ExecuteCommandConfiguration struct {
	KMSKeyID *string `json:"kmsKeyID,omitempty"`
	// The log configuration for the results of the execute command actions. The
	// logs can be sent to CloudWatch Logs or an Amazon S3 bucket.
	LogConfiguration *ExecuteCommandLogConfiguration `json:"logConfiguration,omitempty"`

	Logging *string `json:"logging,omitempty"`
}

// +kubebuilder:skipversion
type ExecuteCommandLogConfiguration struct {
	CloudWatchEncryptionEnabled *bool `json:"cloudWatchEncryptionEnabled,omitempty"`

	CloudWatchLogGroupName *string `json:"cloudWatchLogGroupName,omitempty"`

	S3BucketName *string `json:"s3BucketName,omitempty"`

	S3EncryptionEnabled *bool `json:"s3EncryptionEnabled,omitempty"`

	S3KeyPrefix *string `json:"s3KeyPrefix,omitempty"`
}

// +kubebuilder:skipversion
type FSxWindowsFileServerAuthorizationConfig struct {
	CredentialsParameter *string `json:"credentialsParameter,omitempty"`

	Domain *string `json:"domain,omitempty"`
}

// +kubebuilder:skipversion
type FSxWindowsFileServerVolumeConfiguration struct {
	// The authorization configuration details for Amazon FSx for Windows File Server
	// file system. See FSxWindowsFileServerVolumeConfiguration (https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_FSxWindowsFileServerVolumeConfiguration.html)
	// in the Amazon Elastic Container Service API Reference.
	//
	// For more information and the input format, see Amazon FSx for Windows File
	// Server Volumes (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/wfsx-volumes.html)
	// in the Amazon Elastic Container Service Developer Guide.
	AuthorizationConfig *FSxWindowsFileServerAuthorizationConfig `json:"authorizationConfig,omitempty"`

	FileSystemID *string `json:"fileSystemID,omitempty"`

	RootDirectory *string `json:"rootDirectory,omitempty"`
}

// +kubebuilder:skipversion
type Failure struct {
	ARN *string `json:"arn,omitempty"`

	Detail *string `json:"detail,omitempty"`

	Reason *string `json:"reason,omitempty"`
}

// +kubebuilder:skipversion
type FirelensConfiguration struct {
	Options map[string]*string `json:"options,omitempty"`

	Type *string `json:"type_,omitempty"`
}

// +kubebuilder:skipversion
type HealthCheck struct {
	Command []*string `json:"command,omitempty"`

	Interval *int64 `json:"interval,omitempty"`

	Retries *int64 `json:"retries,omitempty"`

	StartPeriod *int64 `json:"startPeriod,omitempty"`

	Timeout *int64 `json:"timeout,omitempty"`
}

// +kubebuilder:skipversion
type HostEntry struct {
	Hostname *string `json:"hostname,omitempty"`

	IPAddress *string `json:"ipAddress,omitempty"`
}

// +kubebuilder:skipversion
type HostVolumeProperties struct {
	SourcePath *string `json:"sourcePath,omitempty"`
}

// +kubebuilder:skipversion
type InferenceAccelerator struct {
	DeviceName *string `json:"deviceName,omitempty"`

	DeviceType *string `json:"deviceType,omitempty"`
}

// +kubebuilder:skipversion
type InferenceAcceleratorOverride struct {
	DeviceName *string `json:"deviceName,omitempty"`

	DeviceType *string `json:"deviceType,omitempty"`
}

// +kubebuilder:skipversion
type InstanceHealthCheckResult struct {
	LastStatusChange *metav1.Time `json:"lastStatusChange,omitempty"`

	LastUpdated *metav1.Time `json:"lastUpdated,omitempty"`
}

// +kubebuilder:skipversion
type KernelCapabilities struct {
	Add []*string `json:"add,omitempty"`

	Drop []*string `json:"drop,omitempty"`
}

// +kubebuilder:skipversion
type KeyValuePair struct {
	Name *string `json:"name,omitempty"`

	Value *string `json:"value,omitempty"`
}

// +kubebuilder:skipversion
type LinuxParameters struct {
	// The Linux capabilities for the container that are added to or dropped from
	// the default configuration provided by Docker. For more information about
	// the default capabilities and the non-default available capabilities, see
	// Runtime privilege and Linux capabilities (https://docs.docker.com/engine/reference/run/#runtime-privilege-and-linux-capabilities)
	// in the Docker run reference. For more detailed information about these Linux
	// capabilities, see the capabilities(7) (http://man7.org/linux/man-pages/man7/capabilities.7.html)
	// Linux manual page.
	Capabilities *KernelCapabilities `json:"capabilities,omitempty"`

	Devices []*Device `json:"devices,omitempty"`

	InitProcessEnabled *bool `json:"initProcessEnabled,omitempty"`

	MaxSwap *int64 `json:"maxSwap,omitempty"`

	SharedMemorySize *int64 `json:"sharedMemorySize,omitempty"`

	Swappiness *int64 `json:"swappiness,omitempty"`

	Tmpfs []*Tmpfs `json:"tmpfs,omitempty"`
}

// +kubebuilder:skipversion
type LoadBalancer struct {
	ContainerName *string `json:"containerName,omitempty"`

	ContainerPort *int64 `json:"containerPort,omitempty"`

	LoadBalancerName *string `json:"loadBalancerName,omitempty"`

	TargetGroupARN *string `json:"targetGroupARN,omitempty"`
}

// +kubebuilder:skipversion
type LogConfiguration struct {
	LogDriver *string `json:"logDriver,omitempty"`

	Options map[string]*string `json:"options,omitempty"`

	SecretOptions []*Secret `json:"secretOptions,omitempty"`
}

// +kubebuilder:skipversion
type ManagedAgent struct {
	LastStartedAt *metav1.Time `json:"lastStartedAt,omitempty"`

	LastStatus *string `json:"lastStatus,omitempty"`

	Reason *string `json:"reason,omitempty"`
}

// +kubebuilder:skipversion
type ManagedAgentStateChange struct {
	ContainerName *string `json:"containerName,omitempty"`

	Reason *string `json:"reason,omitempty"`

	Status *string `json:"status,omitempty"`
}

// +kubebuilder:skipversion
type MountPoint struct {
	ContainerPath *string `json:"containerPath,omitempty"`

	ReadOnly *bool `json:"readOnly,omitempty"`

	SourceVolume *string `json:"sourceVolume,omitempty"`
}

// +kubebuilder:skipversion
type NetworkBinding struct {
	BindIP *string `json:"bindIP,omitempty"`

	ContainerPort *int64 `json:"containerPort,omitempty"`

	HostPort *int64 `json:"hostPort,omitempty"`

	Protocol *string `json:"protocol,omitempty"`
}

// +kubebuilder:skipversion
type NetworkConfiguration struct {
	// An object representing the networking details for a task or service.
	AWSvpcConfiguration *AWSVPCConfiguration `json:"awsvpcConfiguration,omitempty"`
}

// +kubebuilder:skipversion
type NetworkInterface struct {
	AttachmentID *string `json:"attachmentID,omitempty"`

	IPv6Address *string `json:"ipv6Address,omitempty"`

	PrivateIPv4Address *string `json:"privateIPv4Address,omitempty"`
}

// +kubebuilder:skipversion
type PlacementConstraint struct {
	Expression *string `json:"expression,omitempty"`

	Type *string `json:"type_,omitempty"`
}

// +kubebuilder:skipversion
type PlacementStrategy struct {
	Field *string `json:"field,omitempty"`

	Type *string `json:"type_,omitempty"`
}

// +kubebuilder:skipversion
type PlatformDevice struct {
	ID *string `json:"id,omitempty"`
}

// +kubebuilder:skipversion
type PortMapping struct {
	ContainerPort *int64 `json:"containerPort,omitempty"`

	HostPort *int64 `json:"hostPort,omitempty"`

	Protocol *string `json:"protocol,omitempty"`
}

// +kubebuilder:skipversion
type ProxyConfiguration struct {
	ContainerName *string `json:"containerName,omitempty"`

	Properties []*KeyValuePair `json:"properties,omitempty"`

	Type *string `json:"type_,omitempty"`
}

// +kubebuilder:skipversion
type RepositoryCredentials struct {
	CredentialsParameter *string `json:"credentialsParameter,omitempty"`
}

// +kubebuilder:skipversion
type Resource struct {
	DoubleValue *float64 `json:"doubleValue,omitempty"`

	IntegerValue *int64 `json:"integerValue,omitempty"`

	Name *string `json:"name,omitempty"`

	StringSetValue []*string `json:"stringSetValue,omitempty"`

	Type *string `json:"type_,omitempty"`
}

// +kubebuilder:skipversion
type ResourceRequirement struct {
	Type *string `json:"type_,omitempty"`

	Value *string `json:"value,omitempty"`
}

// +kubebuilder:skipversion
type RuntimePlatform struct {
	CPUArchitecture *string `json:"cpuArchitecture,omitempty"`

	OperatingSystemFamily *string `json:"operatingSystemFamily,omitempty"`
}

// +kubebuilder:skipversion
type Scale struct {
	Unit *string `json:"unit,omitempty"`

	Value *float64 `json:"value,omitempty"`
}

// +kubebuilder:skipversion
type Secret struct {
	Name *string `json:"name,omitempty"`

	ValueFrom *string `json:"valueFrom,omitempty"`
}

// +kubebuilder:skipversion
type ServiceEvent struct {
	CreatedAt *metav1.Time `json:"createdAt,omitempty"`

	ID *string `json:"id,omitempty"`

	Message *string `json:"message,omitempty"`
}

// +kubebuilder:skipversion
type ServiceRegistry struct {
	ContainerName *string `json:"containerName,omitempty"`

	ContainerPort *int64 `json:"containerPort,omitempty"`

	Port *int64 `json:"port,omitempty"`

	RegistryARN *string `json:"registryARN,omitempty"`
}

// +kubebuilder:skipversion
type Service_SDK struct {
	CapacityProviderStrategy []*CapacityProviderStrategyItem `json:"capacityProviderStrategy,omitempty"`

	ClusterARN *string `json:"clusterARN,omitempty"`

	CreatedAt *metav1.Time `json:"createdAt,omitempty"`

	CreatedBy *string `json:"createdBy,omitempty"`
	// Optional deployment parameters that control how many tasks run during a deployment
	// and the ordering of stopping and starting tasks.
	DeploymentConfiguration *DeploymentConfiguration `json:"deploymentConfiguration,omitempty"`
	// The deployment controller to use for the service. For more information, see
	// Amazon ECS Deployment Types (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
	// in the Amazon Elastic Container Service Developer Guide.
	DeploymentController *DeploymentController `json:"deploymentController,omitempty"`

	Deployments []*Deployment `json:"deployments,omitempty"`

	DesiredCount *int64 `json:"desiredCount,omitempty"`

	EnableECSManagedTags *bool `json:"enableECSManagedTags,omitempty"`

	EnableExecuteCommand *bool `json:"enableExecuteCommand,omitempty"`

	Events []*ServiceEvent `json:"events,omitempty"`

	HealthCheckGracePeriodSeconds *int64 `json:"healthCheckGracePeriodSeconds,omitempty"`

	LaunchType *string `json:"launchType,omitempty"`

	LoadBalancers []*LoadBalancer `json:"loadBalancers,omitempty"`
	// An object representing the network configuration for a task or service.
	NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty"`

	PendingCount *int64 `json:"pendingCount,omitempty"`

	PlacementConstraints []*PlacementConstraint `json:"placementConstraints,omitempty"`

	PlacementStrategy []*PlacementStrategy `json:"placementStrategy,omitempty"`

	PlatformFamily *string `json:"platformFamily,omitempty"`

	PlatformVersion *string `json:"platformVersion,omitempty"`

	PropagateTags *string `json:"propagateTags,omitempty"`

	RoleARN *string `json:"roleARN,omitempty"`

	RunningCount *int64 `json:"runningCount,omitempty"`

	SchedulingStrategy *string `json:"schedulingStrategy,omitempty"`

	ServiceARN *string `json:"serviceARN,omitempty"`

	ServiceName *string `json:"serviceName,omitempty"`

	ServiceRegistries []*ServiceRegistry `json:"serviceRegistries,omitempty"`

	Status *string `json:"status,omitempty"`

	Tags []*Tag `json:"tags,omitempty"`

	TaskDefinition *string `json:"taskDefinition,omitempty"`

	TaskSets []*TaskSet `json:"taskSets,omitempty"`
}

// +kubebuilder:skipversion
type Session struct {
	SessionID *string `json:"sessionID,omitempty"`

	StreamURL *string `json:"streamURL,omitempty"`
}

// +kubebuilder:skipversion
type Setting struct {
	PrincipalARN *string `json:"principalARN,omitempty"`

	Value *string `json:"value,omitempty"`
}

// +kubebuilder:skipversion
type SystemControl struct {
	Namespace *string `json:"namespace,omitempty"`

	Value *string `json:"value,omitempty"`
}

// +kubebuilder:skipversion
type Tag struct {
	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`
}

// +kubebuilder:skipversion
type Task struct {
	Attachments []*Attachment `json:"attachments,omitempty"`

	AvailabilityZone *string `json:"availabilityZone,omitempty"`

	CapacityProviderName *string `json:"capacityProviderName,omitempty"`

	ClusterARN *string `json:"clusterARN,omitempty"`

	ConnectivityAt *metav1.Time `json:"connectivityAt,omitempty"`

	ContainerInstanceARN *string `json:"containerInstanceARN,omitempty"`

	CPU *string `json:"cpu,omitempty"`

	CreatedAt *metav1.Time `json:"createdAt,omitempty"`

	DesiredStatus *string `json:"desiredStatus,omitempty"`

	EnableExecuteCommand *bool `json:"enableExecuteCommand,omitempty"`
	// The amount of ephemeral storage to allocate for the task. This parameter
	// is used to expand the total amount of ephemeral storage available, beyond
	// the default amount, for tasks hosted on Fargate. For more information, see
	// Fargate task storage (https://docs.aws.amazon.com/AmazonECS/latest/userguide/using_data_volumes.html)
	// in the Amazon ECS User Guide for Fargate.
	//
	// This parameter is only supported for tasks hosted on Fargate using Linux
	// platform version 1.4.0 or later. This parameter is not supported for Windows
	// containers on Fargate.
	EphemeralStorage *EphemeralStorage `json:"ephemeralStorage,omitempty"`

	ExecutionStoppedAt *metav1.Time `json:"executionStoppedAt,omitempty"`

	Group *string `json:"group,omitempty"`

	InferenceAccelerators []*InferenceAccelerator `json:"inferenceAccelerators,omitempty"`

	LastStatus *string `json:"lastStatus,omitempty"`

	LaunchType *string `json:"launchType,omitempty"`

	Memory *string `json:"memory,omitempty"`

	PlatformFamily *string `json:"platformFamily,omitempty"`

	PlatformVersion *string `json:"platformVersion,omitempty"`

	PullStartedAt *metav1.Time `json:"pullStartedAt,omitempty"`

	PullStoppedAt *metav1.Time `json:"pullStoppedAt,omitempty"`

	StartedAt *metav1.Time `json:"startedAt,omitempty"`

	StartedBy *string `json:"startedBy,omitempty"`

	StoppedAt *metav1.Time `json:"stoppedAt,omitempty"`

	StoppedReason *string `json:"stoppedReason,omitempty"`

	StoppingAt *metav1.Time `json:"stoppingAt,omitempty"`

	Tags []*Tag `json:"tags,omitempty"`

	TaskARN *string `json:"taskARN,omitempty"`

	TaskDefinitionARN *string `json:"taskDefinitionARN,omitempty"`
}

// +kubebuilder:skipversion
type TaskDefinitionPlacementConstraint struct {
	Expression *string `json:"expression,omitempty"`

	Type *string `json:"type_,omitempty"`
}

// +kubebuilder:skipversion
type TaskDefinition_SDK struct {
	Compatibilities []*string `json:"compatibilities,omitempty"`

	ContainerDefinitions []*ContainerDefinition `json:"containerDefinitions,omitempty"`

	CPU *string `json:"cpu,omitempty"`

	DeregisteredAt *metav1.Time `json:"deregisteredAt,omitempty"`
	// The amount of ephemeral storage to allocate for the task. This parameter
	// is used to expand the total amount of ephemeral storage available, beyond
	// the default amount, for tasks hosted on Fargate. For more information, see
	// Fargate task storage (https://docs.aws.amazon.com/AmazonECS/latest/userguide/using_data_volumes.html)
	// in the Amazon ECS User Guide for Fargate.
	//
	// This parameter is only supported for tasks hosted on Fargate using Linux
	// platform version 1.4.0 or later. This parameter is not supported for Windows
	// containers on Fargate.
	EphemeralStorage *EphemeralStorage `json:"ephemeralStorage,omitempty"`

	ExecutionRoleARN *string `json:"executionRoleARN,omitempty"`

	Family *string `json:"family,omitempty"`

	InferenceAccelerators []*InferenceAccelerator `json:"inferenceAccelerators,omitempty"`

	IPcMode *string `json:"ipcMode,omitempty"`

	Memory *string `json:"memory,omitempty"`

	NetworkMode *string `json:"networkMode,omitempty"`

	PidMode *string `json:"pidMode,omitempty"`

	PlacementConstraints []*TaskDefinitionPlacementConstraint `json:"placementConstraints,omitempty"`
	// The configuration details for the App Mesh proxy.
	//
	// For tasks that use the EC2 launch type, the container instances require at
	// least version 1.26.0 of the container agent and at least version 1.26.0-1
	// of the ecs-init package to use a proxy configuration. If your container instances
	// are launched from the Amazon ECS optimized AMI version 20190301 or later,
	// then they contain the required versions of the container agent and ecs-init.
	// For more information, see Amazon ECS-optimized Linux AMI (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-optimized_AMI.html)
	ProxyConfiguration *ProxyConfiguration `json:"proxyConfiguration,omitempty"`

	RegisteredAt *metav1.Time `json:"registeredAt,omitempty"`

	RegisteredBy *string `json:"registeredBy,omitempty"`

	RequiresAttributes []*Attribute `json:"requiresAttributes,omitempty"`

	RequiresCompatibilities []*string `json:"requiresCompatibilities,omitempty"`

	Revision *int64 `json:"revision,omitempty"`
	// Information about the platform for the Amazon ECS service or task.
	//
	// For more informataion about RuntimePlatform, see RuntimePlatform (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#runtime-platform)
	// in the Amazon Elastic Container Service Developer Guide.
	RuntimePlatform *RuntimePlatform `json:"runtimePlatform,omitempty"`

	Status *string `json:"status,omitempty"`

	TaskDefinitionARN *string `json:"taskDefinitionARN,omitempty"`

	TaskRoleARN *string `json:"taskRoleARN,omitempty"`

	Volumes []*Volume `json:"volumes,omitempty"`
}

// +kubebuilder:skipversion
type TaskOverride struct {
	CPU *string `json:"cpu,omitempty"`
	// The amount of ephemeral storage to allocate for the task. This parameter
	// is used to expand the total amount of ephemeral storage available, beyond
	// the default amount, for tasks hosted on Fargate. For more information, see
	// Fargate task storage (https://docs.aws.amazon.com/AmazonECS/latest/userguide/using_data_volumes.html)
	// in the Amazon ECS User Guide for Fargate.
	//
	// This parameter is only supported for tasks hosted on Fargate using Linux
	// platform version 1.4.0 or later. This parameter is not supported for Windows
	// containers on Fargate.
	EphemeralStorage *EphemeralStorage `json:"ephemeralStorage,omitempty"`

	ExecutionRoleARN *string `json:"executionRoleARN,omitempty"`

	Memory *string `json:"memory,omitempty"`

	TaskRoleARN *string `json:"taskRoleARN,omitempty"`
}

// +kubebuilder:skipversion
type TaskSet struct {
	CapacityProviderStrategy []*CapacityProviderStrategyItem `json:"capacityProviderStrategy,omitempty"`

	ClusterARN *string `json:"clusterARN,omitempty"`

	ComputedDesiredCount *int64 `json:"computedDesiredCount,omitempty"`

	CreatedAt *metav1.Time `json:"createdAt,omitempty"`

	ExternalID *string `json:"externalID,omitempty"`

	ID *string `json:"id,omitempty"`

	LaunchType *string `json:"launchType,omitempty"`

	LoadBalancers []*LoadBalancer `json:"loadBalancers,omitempty"`
	// An object representing the network configuration for a task or service.
	NetworkConfiguration *NetworkConfiguration `json:"networkConfiguration,omitempty"`

	PendingCount *int64 `json:"pendingCount,omitempty"`

	PlatformFamily *string `json:"platformFamily,omitempty"`

	PlatformVersion *string `json:"platformVersion,omitempty"`

	RunningCount *int64 `json:"runningCount,omitempty"`
	// A floating-point percentage of the desired number of tasks to place and keep
	// running in the task set.
	Scale *Scale `json:"scale,omitempty"`

	ServiceARN *string `json:"serviceARN,omitempty"`

	ServiceRegistries []*ServiceRegistry `json:"serviceRegistries,omitempty"`

	StabilityStatus *string `json:"stabilityStatus,omitempty"`

	StabilityStatusAt *metav1.Time `json:"stabilityStatusAt,omitempty"`

	StartedBy *string `json:"startedBy,omitempty"`

	Status *string `json:"status,omitempty"`

	Tags []*Tag `json:"tags,omitempty"`

	TaskDefinition *string `json:"taskDefinition,omitempty"`

	TaskSetARN *string `json:"taskSetARN,omitempty"`

	UpdatedAt *metav1.Time `json:"updatedAt,omitempty"`
}

// +kubebuilder:skipversion
type Tmpfs struct {
	ContainerPath *string `json:"containerPath,omitempty"`

	MountOptions []*string `json:"mountOptions,omitempty"`

	Size *int64 `json:"size,omitempty"`
}

// +kubebuilder:skipversion
type Ulimit struct {
	HardLimit *int64 `json:"hardLimit,omitempty"`

	Name *string `json:"name,omitempty"`

	SoftLimit *int64 `json:"softLimit,omitempty"`
}

// +kubebuilder:skipversion
type VersionInfo struct {
	AgentHash *string `json:"agentHash,omitempty"`

	AgentVersion *string `json:"agentVersion,omitempty"`

	DockerVersion *string `json:"dockerVersion,omitempty"`
}

// +kubebuilder:skipversion
type Volume struct {
	// This parameter is specified when you're using Docker volumes. Docker volumes
	// are only supported when you're using the EC2 launch type. Windows containers
	// only support the use of the local driver. To use bind mounts, specify a host
	// instead.
	DockerVolumeConfiguration *DockerVolumeConfiguration `json:"dockerVolumeConfiguration,omitempty"`
	// This parameter is specified when you're using an Amazon Elastic File System
	// file system for task storage. For more information, see Amazon EFS Volumes
	// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/efs-volumes.html)
	// in the Amazon Elastic Container Service Developer Guide.
	EFSVolumeConfiguration *EFSVolumeConfiguration `json:"efsVolumeConfiguration,omitempty"`
	// This parameter is specified when you're using Amazon FSx for Windows File
	// Server (https://docs.aws.amazon.com/fsx/latest/WindowsGuide/what-is.html)
	// file system for task storage.
	//
	// For more information and the input format, see Amazon FSx for Windows File
	// Server Volumes (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/wfsx-volumes.html)
	// in the Amazon Elastic Container Service Developer Guide.
	FsxWindowsFileServerVolumeConfiguration *FSxWindowsFileServerVolumeConfiguration `json:"fsxWindowsFileServerVolumeConfiguration,omitempty"`
	// Details on a container instance bind mount host volume.
	Host *HostVolumeProperties `json:"host,omitempty"`

	Name *string `json:"name,omitempty"`
}

// +kubebuilder:skipversion
type VolumeFrom struct {
	ReadOnly *bool `json:"readOnly,omitempty"`

	SourceContainer *string `json:"sourceContainer,omitempty"`
}

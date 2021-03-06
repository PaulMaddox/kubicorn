// +build go1.9

// Copyright 2017 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder
// commit ID: dab57ee609fffdc578f48519d5cdc980efd8cc00

package streamanalytics

import original "github.com/Azure/azure-sdk-for-go/services/streamanalytics/mgmt/2016-03-01/streamanalytics"

type InputsClient = original.InputsClient
type BindingType = original.BindingType

const (
	BindingTypeMicrosoftMachineLearningWebService    BindingType = original.BindingTypeMicrosoftMachineLearningWebService
	BindingTypeMicrosoftStreamAnalyticsJavascriptUdf BindingType = original.BindingTypeMicrosoftStreamAnalyticsJavascriptUdf
)

type CompatibilityLevel = original.CompatibilityLevel

const (
	OneFullStopZero CompatibilityLevel = original.OneFullStopZero
)

type Encoding = original.Encoding

const (
	UTF8 Encoding = original.UTF8
)

type EventsOutOfOrderPolicy = original.EventsOutOfOrderPolicy

const (
	Adjust EventsOutOfOrderPolicy = original.Adjust
	Drop   EventsOutOfOrderPolicy = original.Drop
)

type JSONOutputSerializationFormat = original.JSONOutputSerializationFormat

const (
	Array         JSONOutputSerializationFormat = original.Array
	LineSeparated JSONOutputSerializationFormat = original.LineSeparated
)

type OutputErrorPolicy = original.OutputErrorPolicy

const (
	OutputErrorPolicyDrop OutputErrorPolicy = original.OutputErrorPolicyDrop
	OutputErrorPolicyStop OutputErrorPolicy = original.OutputErrorPolicyStop
)

type OutputStartMode = original.OutputStartMode

const (
	CustomTime          OutputStartMode = original.CustomTime
	JobStartTime        OutputStartMode = original.JobStartTime
	LastOutputEventTime OutputStartMode = original.LastOutputEventTime
)

type SkuName = original.SkuName

const (
	Standard SkuName = original.Standard
)

type Type = original.Type

const (
	TypeAvro Type = original.TypeAvro
	TypeCsv  Type = original.TypeCsv
	TypeJSON Type = original.TypeJSON
)

type TypeFunctionBinding = original.TypeFunctionBinding

const (
	TypeMicrosoftMachineLearningWebService    TypeFunctionBinding = original.TypeMicrosoftMachineLearningWebService
	TypeMicrosoftStreamAnalyticsJavascriptUdf TypeFunctionBinding = original.TypeMicrosoftStreamAnalyticsJavascriptUdf
)

type TypeFunctionProperties = original.TypeFunctionProperties

const (
	TypeScalar TypeFunctionProperties = original.TypeScalar
)

type TypeInputProperties = original.TypeInputProperties

const (
	TypeReference TypeInputProperties = original.TypeReference
	TypeStream    TypeInputProperties = original.TypeStream
)

type TypeOutputDataSource = original.TypeOutputDataSource

const (
	TypeMicrosoftDataLakeAccounts   TypeOutputDataSource = original.TypeMicrosoftDataLakeAccounts
	TypeMicrosoftServiceBusEventHub TypeOutputDataSource = original.TypeMicrosoftServiceBusEventHub
	TypeMicrosoftServiceBusQueue    TypeOutputDataSource = original.TypeMicrosoftServiceBusQueue
	TypeMicrosoftServiceBusTopic    TypeOutputDataSource = original.TypeMicrosoftServiceBusTopic
	TypeMicrosoftSQLServerDatabase  TypeOutputDataSource = original.TypeMicrosoftSQLServerDatabase
	TypeMicrosoftStorageBlob        TypeOutputDataSource = original.TypeMicrosoftStorageBlob
	TypeMicrosoftStorageDocumentDB  TypeOutputDataSource = original.TypeMicrosoftStorageDocumentDB
	TypeMicrosoftStorageTable       TypeOutputDataSource = original.TypeMicrosoftStorageTable
	TypePowerBI                     TypeOutputDataSource = original.TypePowerBI
)

type TypeReferenceInputDataSource = original.TypeReferenceInputDataSource

const (
	TypeReferenceInputDataSourceTypeMicrosoftStorageBlob TypeReferenceInputDataSource = original.TypeReferenceInputDataSourceTypeMicrosoftStorageBlob
)

type TypeStreamInputDataSource = original.TypeStreamInputDataSource

const (
	TypeStreamInputDataSourceTypeMicrosoftDevicesIotHubs     TypeStreamInputDataSource = original.TypeStreamInputDataSourceTypeMicrosoftDevicesIotHubs
	TypeStreamInputDataSourceTypeMicrosoftServiceBusEventHub TypeStreamInputDataSource = original.TypeStreamInputDataSourceTypeMicrosoftServiceBusEventHub
	TypeStreamInputDataSourceTypeMicrosoftStorageBlob        TypeStreamInputDataSource = original.TypeStreamInputDataSourceTypeMicrosoftStorageBlob
)

type UdfType = original.UdfType

const (
	Scalar UdfType = original.Scalar
)

type AvroSerialization = original.AvroSerialization
type AzureDataLakeStoreOutputDataSource = original.AzureDataLakeStoreOutputDataSource
type AzureDataLakeStoreOutputDataSourceProperties = original.AzureDataLakeStoreOutputDataSourceProperties
type AzureMachineLearningWebServiceFunctionBinding = original.AzureMachineLearningWebServiceFunctionBinding
type AzureMachineLearningWebServiceFunctionBindingProperties = original.AzureMachineLearningWebServiceFunctionBindingProperties
type AzureMachineLearningWebServiceFunctionBindingRetrievalProperties = original.AzureMachineLearningWebServiceFunctionBindingRetrievalProperties
type AzureMachineLearningWebServiceFunctionRetrieveDefaultDefinitionParameters = original.AzureMachineLearningWebServiceFunctionRetrieveDefaultDefinitionParameters
type AzureMachineLearningWebServiceInputColumn = original.AzureMachineLearningWebServiceInputColumn
type AzureMachineLearningWebServiceInputs = original.AzureMachineLearningWebServiceInputs
type AzureMachineLearningWebServiceOutputColumn = original.AzureMachineLearningWebServiceOutputColumn
type AzureSQLDatabaseDataSourceProperties = original.AzureSQLDatabaseDataSourceProperties
type AzureSQLDatabaseOutputDataSource = original.AzureSQLDatabaseOutputDataSource
type AzureSQLDatabaseOutputDataSourceProperties = original.AzureSQLDatabaseOutputDataSourceProperties
type AzureTableOutputDataSource = original.AzureTableOutputDataSource
type AzureTableOutputDataSourceProperties = original.AzureTableOutputDataSourceProperties
type BlobDataSourceProperties = original.BlobDataSourceProperties
type BlobOutputDataSource = original.BlobOutputDataSource
type BlobOutputDataSourceProperties = original.BlobOutputDataSourceProperties
type BlobReferenceInputDataSource = original.BlobReferenceInputDataSource
type BlobReferenceInputDataSourceProperties = original.BlobReferenceInputDataSourceProperties
type BlobStreamInputDataSource = original.BlobStreamInputDataSource
type BlobStreamInputDataSourceProperties = original.BlobStreamInputDataSourceProperties
type CsvSerialization = original.CsvSerialization
type CsvSerializationProperties = original.CsvSerializationProperties
type DiagnosticCondition = original.DiagnosticCondition
type Diagnostics = original.Diagnostics
type DocumentDbOutputDataSource = original.DocumentDbOutputDataSource
type DocumentDbOutputDataSourceProperties = original.DocumentDbOutputDataSourceProperties
type ErrorResponse = original.ErrorResponse
type EventHubDataSourceProperties = original.EventHubDataSourceProperties
type EventHubOutputDataSource = original.EventHubOutputDataSource
type EventHubOutputDataSourceProperties = original.EventHubOutputDataSourceProperties
type EventHubStreamInputDataSource = original.EventHubStreamInputDataSource
type EventHubStreamInputDataSourceProperties = original.EventHubStreamInputDataSourceProperties
type Function = original.Function
type FunctionBinding = original.FunctionBinding
type FunctionInput = original.FunctionInput
type FunctionListResult = original.FunctionListResult
type FunctionOutput = original.FunctionOutput
type FunctionProperties = original.FunctionProperties
type FunctionRetrieveDefaultDefinitionParameters = original.FunctionRetrieveDefaultDefinitionParameters
type Input = original.Input
type InputListResult = original.InputListResult
type InputProperties = original.InputProperties
type IoTHubStreamInputDataSource = original.IoTHubStreamInputDataSource
type IoTHubStreamInputDataSourceProperties = original.IoTHubStreamInputDataSourceProperties
type JavaScriptFunctionBinding = original.JavaScriptFunctionBinding
type JavaScriptFunctionBindingProperties = original.JavaScriptFunctionBindingProperties
type JavaScriptFunctionBindingRetrievalProperties = original.JavaScriptFunctionBindingRetrievalProperties
type JavaScriptFunctionRetrieveDefaultDefinitionParameters = original.JavaScriptFunctionRetrieveDefaultDefinitionParameters
type JSONSerialization = original.JSONSerialization
type JSONSerializationProperties = original.JSONSerializationProperties
type OAuthBasedDataSourceProperties = original.OAuthBasedDataSourceProperties
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type Output = original.Output
type OutputDataSource = original.OutputDataSource
type OutputListResult = original.OutputListResult
type OutputProperties = original.OutputProperties
type PowerBIOutputDataSource = original.PowerBIOutputDataSource
type PowerBIOutputDataSourceProperties = original.PowerBIOutputDataSourceProperties
type ReferenceInputDataSource = original.ReferenceInputDataSource
type ReferenceInputProperties = original.ReferenceInputProperties
type Resource = original.Resource
type ResourceTestStatus = original.ResourceTestStatus
type ScalarFunctionConfiguration = original.ScalarFunctionConfiguration
type ScalarFunctionProperties = original.ScalarFunctionProperties
type Serialization = original.Serialization
type ServiceBusDataSourceProperties = original.ServiceBusDataSourceProperties
type ServiceBusQueueOutputDataSource = original.ServiceBusQueueOutputDataSource
type ServiceBusQueueOutputDataSourceProperties = original.ServiceBusQueueOutputDataSourceProperties
type ServiceBusTopicOutputDataSource = original.ServiceBusTopicOutputDataSource
type ServiceBusTopicOutputDataSourceProperties = original.ServiceBusTopicOutputDataSourceProperties
type Sku = original.Sku
type StartStreamingJobParameters = original.StartStreamingJobParameters
type StorageAccount = original.StorageAccount
type StreamingJob = original.StreamingJob
type StreamingJobListResult = original.StreamingJobListResult
type StreamingJobProperties = original.StreamingJobProperties
type StreamInputDataSource = original.StreamInputDataSource
type StreamInputProperties = original.StreamInputProperties
type SubResource = original.SubResource
type SubscriptionQuota = original.SubscriptionQuota
type SubscriptionQuotaProperties = original.SubscriptionQuotaProperties
type SubscriptionQuotasListResult = original.SubscriptionQuotasListResult
type Transformation = original.Transformation
type TransformationProperties = original.TransformationProperties
type OperationsClient = original.OperationsClient
type SubscriptionsClient = original.SubscriptionsClient
type FunctionsClient = original.FunctionsClient
type OutputsClient = original.OutputsClient
type StreamingJobsClient = original.StreamingJobsClient
type TransformationsClient = original.TransformationsClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ManagementClient = original.ManagementClient

func NewFunctionsClient(subscriptionID string) FunctionsClient {
	return original.NewFunctionsClient(subscriptionID)
}
func NewFunctionsClientWithBaseURI(baseURI string, subscriptionID string) FunctionsClient {
	return original.NewFunctionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewInputsClient(subscriptionID string) InputsClient {
	return original.NewInputsClient(subscriptionID)
}
func NewInputsClientWithBaseURI(baseURI string, subscriptionID string) InputsClient {
	return original.NewInputsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSubscriptionsClient(subscriptionID string) SubscriptionsClient {
	return original.NewSubscriptionsClient(subscriptionID)
}
func NewSubscriptionsClientWithBaseURI(baseURI string, subscriptionID string) SubscriptionsClient {
	return original.NewSubscriptionsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
func New(subscriptionID string) ManagementClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) ManagementClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewOutputsClient(subscriptionID string) OutputsClient {
	return original.NewOutputsClient(subscriptionID)
}
func NewOutputsClientWithBaseURI(baseURI string, subscriptionID string) OutputsClient {
	return original.NewOutputsClientWithBaseURI(baseURI, subscriptionID)
}
func NewStreamingJobsClient(subscriptionID string) StreamingJobsClient {
	return original.NewStreamingJobsClient(subscriptionID)
}
func NewStreamingJobsClientWithBaseURI(baseURI string, subscriptionID string) StreamingJobsClient {
	return original.NewStreamingJobsClientWithBaseURI(baseURI, subscriptionID)
}
func NewTransformationsClient(subscriptionID string) TransformationsClient {
	return original.NewTransformationsClient(subscriptionID)
}
func NewTransformationsClientWithBaseURI(baseURI string, subscriptionID string) TransformationsClient {
	return original.NewTransformationsClientWithBaseURI(baseURI, subscriptionID)
}

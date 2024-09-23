# Execution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildOutputId** | Pointer to **string** |  | [optional] 
**BuildOutputIdObject** | Pointer to [**Output**](Output.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**DeployOutputId** | Pointer to **string** |  | [optional] 
**DeployOutputIdObject** | Pointer to [**Output**](Output.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**JobId** | Pointer to **string** |  | [optional] 
**JobIdObject** | Pointer to [**Job**](Job.md) |  | [optional] 
**PublishOutputId** | Pointer to **string** |  | [optional] 
**PublishOutputIdObject** | Pointer to [**Output**](Output.md) |  | [optional] 
**ReferencedByTriggerHasExecutionExecutionIdObjects** | Pointer to [**[]TriggerHasExecution**](TriggerHasExecution.md) |  | [optional] 
**TestOutputId** | Pointer to **string** |  | [optional] 
**TestOutputIdObject** | Pointer to [**Output**](Output.md) |  | [optional] 
**TriggerId** | Pointer to **string** |  | [optional] 
**TriggerIdObject** | Pointer to [**Trigger**](Trigger.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**ValidateOutputId** | Pointer to **string** |  | [optional] 
**ValidateOutputIdObject** | Pointer to [**Output**](Output.md) |  | [optional] 

## Methods

### NewExecution

`func NewExecution() *Execution`

NewExecution instantiates a new Execution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionWithDefaults

`func NewExecutionWithDefaults() *Execution`

NewExecutionWithDefaults instantiates a new Execution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildOutputId

`func (o *Execution) GetBuildOutputId() string`

GetBuildOutputId returns the BuildOutputId field if non-nil, zero value otherwise.

### GetBuildOutputIdOk

`func (o *Execution) GetBuildOutputIdOk() (*string, bool)`

GetBuildOutputIdOk returns a tuple with the BuildOutputId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildOutputId

`func (o *Execution) SetBuildOutputId(v string)`

SetBuildOutputId sets BuildOutputId field to given value.

### HasBuildOutputId

`func (o *Execution) HasBuildOutputId() bool`

HasBuildOutputId returns a boolean if a field has been set.

### GetBuildOutputIdObject

`func (o *Execution) GetBuildOutputIdObject() Output`

GetBuildOutputIdObject returns the BuildOutputIdObject field if non-nil, zero value otherwise.

### GetBuildOutputIdObjectOk

`func (o *Execution) GetBuildOutputIdObjectOk() (*Output, bool)`

GetBuildOutputIdObjectOk returns a tuple with the BuildOutputIdObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildOutputIdObject

`func (o *Execution) SetBuildOutputIdObject(v Output)`

SetBuildOutputIdObject sets BuildOutputIdObject field to given value.

### HasBuildOutputIdObject

`func (o *Execution) HasBuildOutputIdObject() bool`

HasBuildOutputIdObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Execution) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Execution) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Execution) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Execution) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Execution) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Execution) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Execution) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Execution) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeployOutputId

`func (o *Execution) GetDeployOutputId() string`

GetDeployOutputId returns the DeployOutputId field if non-nil, zero value otherwise.

### GetDeployOutputIdOk

`func (o *Execution) GetDeployOutputIdOk() (*string, bool)`

GetDeployOutputIdOk returns a tuple with the DeployOutputId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployOutputId

`func (o *Execution) SetDeployOutputId(v string)`

SetDeployOutputId sets DeployOutputId field to given value.

### HasDeployOutputId

`func (o *Execution) HasDeployOutputId() bool`

HasDeployOutputId returns a boolean if a field has been set.

### GetDeployOutputIdObject

`func (o *Execution) GetDeployOutputIdObject() Output`

GetDeployOutputIdObject returns the DeployOutputIdObject field if non-nil, zero value otherwise.

### GetDeployOutputIdObjectOk

`func (o *Execution) GetDeployOutputIdObjectOk() (*Output, bool)`

GetDeployOutputIdObjectOk returns a tuple with the DeployOutputIdObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployOutputIdObject

`func (o *Execution) SetDeployOutputIdObject(v Output)`

SetDeployOutputIdObject sets DeployOutputIdObject field to given value.

### HasDeployOutputIdObject

`func (o *Execution) HasDeployOutputIdObject() bool`

HasDeployOutputIdObject returns a boolean if a field has been set.

### GetId

`func (o *Execution) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Execution) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Execution) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Execution) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobId

`func (o *Execution) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *Execution) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *Execution) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *Execution) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetJobIdObject

`func (o *Execution) GetJobIdObject() Job`

GetJobIdObject returns the JobIdObject field if non-nil, zero value otherwise.

### GetJobIdObjectOk

`func (o *Execution) GetJobIdObjectOk() (*Job, bool)`

GetJobIdObjectOk returns a tuple with the JobIdObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobIdObject

`func (o *Execution) SetJobIdObject(v Job)`

SetJobIdObject sets JobIdObject field to given value.

### HasJobIdObject

`func (o *Execution) HasJobIdObject() bool`

HasJobIdObject returns a boolean if a field has been set.

### GetPublishOutputId

`func (o *Execution) GetPublishOutputId() string`

GetPublishOutputId returns the PublishOutputId field if non-nil, zero value otherwise.

### GetPublishOutputIdOk

`func (o *Execution) GetPublishOutputIdOk() (*string, bool)`

GetPublishOutputIdOk returns a tuple with the PublishOutputId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishOutputId

`func (o *Execution) SetPublishOutputId(v string)`

SetPublishOutputId sets PublishOutputId field to given value.

### HasPublishOutputId

`func (o *Execution) HasPublishOutputId() bool`

HasPublishOutputId returns a boolean if a field has been set.

### GetPublishOutputIdObject

`func (o *Execution) GetPublishOutputIdObject() Output`

GetPublishOutputIdObject returns the PublishOutputIdObject field if non-nil, zero value otherwise.

### GetPublishOutputIdObjectOk

`func (o *Execution) GetPublishOutputIdObjectOk() (*Output, bool)`

GetPublishOutputIdObjectOk returns a tuple with the PublishOutputIdObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishOutputIdObject

`func (o *Execution) SetPublishOutputIdObject(v Output)`

SetPublishOutputIdObject sets PublishOutputIdObject field to given value.

### HasPublishOutputIdObject

`func (o *Execution) HasPublishOutputIdObject() bool`

HasPublishOutputIdObject returns a boolean if a field has been set.

### GetReferencedByTriggerHasExecutionExecutionIdObjects

`func (o *Execution) GetReferencedByTriggerHasExecutionExecutionIdObjects() []TriggerHasExecution`

GetReferencedByTriggerHasExecutionExecutionIdObjects returns the ReferencedByTriggerHasExecutionExecutionIdObjects field if non-nil, zero value otherwise.

### GetReferencedByTriggerHasExecutionExecutionIdObjectsOk

`func (o *Execution) GetReferencedByTriggerHasExecutionExecutionIdObjectsOk() (*[]TriggerHasExecution, bool)`

GetReferencedByTriggerHasExecutionExecutionIdObjectsOk returns a tuple with the ReferencedByTriggerHasExecutionExecutionIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByTriggerHasExecutionExecutionIdObjects

`func (o *Execution) SetReferencedByTriggerHasExecutionExecutionIdObjects(v []TriggerHasExecution)`

SetReferencedByTriggerHasExecutionExecutionIdObjects sets ReferencedByTriggerHasExecutionExecutionIdObjects field to given value.

### HasReferencedByTriggerHasExecutionExecutionIdObjects

`func (o *Execution) HasReferencedByTriggerHasExecutionExecutionIdObjects() bool`

HasReferencedByTriggerHasExecutionExecutionIdObjects returns a boolean if a field has been set.

### SetReferencedByTriggerHasExecutionExecutionIdObjectsNil

`func (o *Execution) SetReferencedByTriggerHasExecutionExecutionIdObjectsNil(b bool)`

 SetReferencedByTriggerHasExecutionExecutionIdObjectsNil sets the value for ReferencedByTriggerHasExecutionExecutionIdObjects to be an explicit nil

### UnsetReferencedByTriggerHasExecutionExecutionIdObjects
`func (o *Execution) UnsetReferencedByTriggerHasExecutionExecutionIdObjects()`

UnsetReferencedByTriggerHasExecutionExecutionIdObjects ensures that no value is present for ReferencedByTriggerHasExecutionExecutionIdObjects, not even an explicit nil
### GetTestOutputId

`func (o *Execution) GetTestOutputId() string`

GetTestOutputId returns the TestOutputId field if non-nil, zero value otherwise.

### GetTestOutputIdOk

`func (o *Execution) GetTestOutputIdOk() (*string, bool)`

GetTestOutputIdOk returns a tuple with the TestOutputId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestOutputId

`func (o *Execution) SetTestOutputId(v string)`

SetTestOutputId sets TestOutputId field to given value.

### HasTestOutputId

`func (o *Execution) HasTestOutputId() bool`

HasTestOutputId returns a boolean if a field has been set.

### GetTestOutputIdObject

`func (o *Execution) GetTestOutputIdObject() Output`

GetTestOutputIdObject returns the TestOutputIdObject field if non-nil, zero value otherwise.

### GetTestOutputIdObjectOk

`func (o *Execution) GetTestOutputIdObjectOk() (*Output, bool)`

GetTestOutputIdObjectOk returns a tuple with the TestOutputIdObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestOutputIdObject

`func (o *Execution) SetTestOutputIdObject(v Output)`

SetTestOutputIdObject sets TestOutputIdObject field to given value.

### HasTestOutputIdObject

`func (o *Execution) HasTestOutputIdObject() bool`

HasTestOutputIdObject returns a boolean if a field has been set.

### GetTriggerId

`func (o *Execution) GetTriggerId() string`

GetTriggerId returns the TriggerId field if non-nil, zero value otherwise.

### GetTriggerIdOk

`func (o *Execution) GetTriggerIdOk() (*string, bool)`

GetTriggerIdOk returns a tuple with the TriggerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerId

`func (o *Execution) SetTriggerId(v string)`

SetTriggerId sets TriggerId field to given value.

### HasTriggerId

`func (o *Execution) HasTriggerId() bool`

HasTriggerId returns a boolean if a field has been set.

### GetTriggerIdObject

`func (o *Execution) GetTriggerIdObject() Trigger`

GetTriggerIdObject returns the TriggerIdObject field if non-nil, zero value otherwise.

### GetTriggerIdObjectOk

`func (o *Execution) GetTriggerIdObjectOk() (*Trigger, bool)`

GetTriggerIdObjectOk returns a tuple with the TriggerIdObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerIdObject

`func (o *Execution) SetTriggerIdObject(v Trigger)`

SetTriggerIdObject sets TriggerIdObject field to given value.

### HasTriggerIdObject

`func (o *Execution) HasTriggerIdObject() bool`

HasTriggerIdObject returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Execution) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Execution) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Execution) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Execution) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetValidateOutputId

`func (o *Execution) GetValidateOutputId() string`

GetValidateOutputId returns the ValidateOutputId field if non-nil, zero value otherwise.

### GetValidateOutputIdOk

`func (o *Execution) GetValidateOutputIdOk() (*string, bool)`

GetValidateOutputIdOk returns a tuple with the ValidateOutputId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateOutputId

`func (o *Execution) SetValidateOutputId(v string)`

SetValidateOutputId sets ValidateOutputId field to given value.

### HasValidateOutputId

`func (o *Execution) HasValidateOutputId() bool`

HasValidateOutputId returns a boolean if a field has been set.

### GetValidateOutputIdObject

`func (o *Execution) GetValidateOutputIdObject() Output`

GetValidateOutputIdObject returns the ValidateOutputIdObject field if non-nil, zero value otherwise.

### GetValidateOutputIdObjectOk

`func (o *Execution) GetValidateOutputIdObjectOk() (*Output, bool)`

GetValidateOutputIdObjectOk returns a tuple with the ValidateOutputIdObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateOutputIdObject

`func (o *Execution) SetValidateOutputIdObject(v Output)`

SetValidateOutputIdObject sets ValidateOutputIdObject field to given value.

### HasValidateOutputIdObject

`func (o *Execution) HasValidateOutputIdObject() bool`

HasValidateOutputIdObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



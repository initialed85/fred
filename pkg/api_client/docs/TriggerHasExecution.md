# TriggerHasExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**ExecutionId** | Pointer to **string** |  | [optional] 
**ExecutionIdObject** | Pointer to [**Execution**](Execution.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**TriggerId** | Pointer to **string** |  | [optional] 
**TriggerIdObject** | Pointer to [**Trigger**](Trigger.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTriggerHasExecution

`func NewTriggerHasExecution() *TriggerHasExecution`

NewTriggerHasExecution instantiates a new TriggerHasExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerHasExecutionWithDefaults

`func NewTriggerHasExecutionWithDefaults() *TriggerHasExecution`

NewTriggerHasExecutionWithDefaults instantiates a new TriggerHasExecution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *TriggerHasExecution) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TriggerHasExecution) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TriggerHasExecution) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TriggerHasExecution) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *TriggerHasExecution) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *TriggerHasExecution) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *TriggerHasExecution) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *TriggerHasExecution) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetExecutionId

`func (o *TriggerHasExecution) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *TriggerHasExecution) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *TriggerHasExecution) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *TriggerHasExecution) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetExecutionIdObject

`func (o *TriggerHasExecution) GetExecutionIdObject() Execution`

GetExecutionIdObject returns the ExecutionIdObject field if non-nil, zero value otherwise.

### GetExecutionIdObjectOk

`func (o *TriggerHasExecution) GetExecutionIdObjectOk() (*Execution, bool)`

GetExecutionIdObjectOk returns a tuple with the ExecutionIdObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionIdObject

`func (o *TriggerHasExecution) SetExecutionIdObject(v Execution)`

SetExecutionIdObject sets ExecutionIdObject field to given value.

### HasExecutionIdObject

`func (o *TriggerHasExecution) HasExecutionIdObject() bool`

HasExecutionIdObject returns a boolean if a field has been set.

### GetId

`func (o *TriggerHasExecution) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TriggerHasExecution) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TriggerHasExecution) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TriggerHasExecution) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTriggerId

`func (o *TriggerHasExecution) GetTriggerId() string`

GetTriggerId returns the TriggerId field if non-nil, zero value otherwise.

### GetTriggerIdOk

`func (o *TriggerHasExecution) GetTriggerIdOk() (*string, bool)`

GetTriggerIdOk returns a tuple with the TriggerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerId

`func (o *TriggerHasExecution) SetTriggerId(v string)`

SetTriggerId sets TriggerId field to given value.

### HasTriggerId

`func (o *TriggerHasExecution) HasTriggerId() bool`

HasTriggerId returns a boolean if a field has been set.

### GetTriggerIdObject

`func (o *TriggerHasExecution) GetTriggerIdObject() Trigger`

GetTriggerIdObject returns the TriggerIdObject field if non-nil, zero value otherwise.

### GetTriggerIdObjectOk

`func (o *TriggerHasExecution) GetTriggerIdObjectOk() (*Trigger, bool)`

GetTriggerIdObjectOk returns a tuple with the TriggerIdObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerIdObject

`func (o *TriggerHasExecution) SetTriggerIdObject(v Trigger)`

SetTriggerIdObject sets TriggerIdObject field to given value.

### HasTriggerIdObject

`func (o *TriggerHasExecution) HasTriggerIdObject() bool`

HasTriggerIdObject returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TriggerHasExecution) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TriggerHasExecution) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TriggerHasExecution) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TriggerHasExecution) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



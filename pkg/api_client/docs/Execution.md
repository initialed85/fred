# Execution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeId** | Pointer to **string** |  | [optional] 
**ChangeIdObject** | Pointer to [**Change**](Change.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**EndedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**JobExecutorClaimedUntil** | Pointer to **time.Time** |  | [optional] 
**JobId** | Pointer to **string** |  | [optional] 
**JobIdObject** | Pointer to [**Job**](Job.md) |  | [optional] 
**ReferencedByOutputExecutionIdObjects** | Pointer to [**[]Output**](Output.md) |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TriggerId** | Pointer to **string** |  | [optional] 
**TriggerIdObject** | Pointer to [**Trigger**](Trigger.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

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

### GetChangeId

`func (o *Execution) GetChangeId() string`

GetChangeId returns the ChangeId field if non-nil, zero value otherwise.

### GetChangeIdOk

`func (o *Execution) GetChangeIdOk() (*string, bool)`

GetChangeIdOk returns a tuple with the ChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeId

`func (o *Execution) SetChangeId(v string)`

SetChangeId sets ChangeId field to given value.

### HasChangeId

`func (o *Execution) HasChangeId() bool`

HasChangeId returns a boolean if a field has been set.

### GetChangeIdObject

`func (o *Execution) GetChangeIdObject() Change`

GetChangeIdObject returns the ChangeIdObject field if non-nil, zero value otherwise.

### GetChangeIdObjectOk

`func (o *Execution) GetChangeIdObjectOk() (*Change, bool)`

GetChangeIdObjectOk returns a tuple with the ChangeIdObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeIdObject

`func (o *Execution) SetChangeIdObject(v Change)`

SetChangeIdObject sets ChangeIdObject field to given value.

### HasChangeIdObject

`func (o *Execution) HasChangeIdObject() bool`

HasChangeIdObject returns a boolean if a field has been set.

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

### GetEndedAt

`func (o *Execution) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *Execution) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *Execution) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *Execution) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

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

### GetJobExecutorClaimedUntil

`func (o *Execution) GetJobExecutorClaimedUntil() time.Time`

GetJobExecutorClaimedUntil returns the JobExecutorClaimedUntil field if non-nil, zero value otherwise.

### GetJobExecutorClaimedUntilOk

`func (o *Execution) GetJobExecutorClaimedUntilOk() (*time.Time, bool)`

GetJobExecutorClaimedUntilOk returns a tuple with the JobExecutorClaimedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobExecutorClaimedUntil

`func (o *Execution) SetJobExecutorClaimedUntil(v time.Time)`

SetJobExecutorClaimedUntil sets JobExecutorClaimedUntil field to given value.

### HasJobExecutorClaimedUntil

`func (o *Execution) HasJobExecutorClaimedUntil() bool`

HasJobExecutorClaimedUntil returns a boolean if a field has been set.

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

### GetReferencedByOutputExecutionIdObjects

`func (o *Execution) GetReferencedByOutputExecutionIdObjects() []Output`

GetReferencedByOutputExecutionIdObjects returns the ReferencedByOutputExecutionIdObjects field if non-nil, zero value otherwise.

### GetReferencedByOutputExecutionIdObjectsOk

`func (o *Execution) GetReferencedByOutputExecutionIdObjectsOk() (*[]Output, bool)`

GetReferencedByOutputExecutionIdObjectsOk returns a tuple with the ReferencedByOutputExecutionIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByOutputExecutionIdObjects

`func (o *Execution) SetReferencedByOutputExecutionIdObjects(v []Output)`

SetReferencedByOutputExecutionIdObjects sets ReferencedByOutputExecutionIdObjects field to given value.

### HasReferencedByOutputExecutionIdObjects

`func (o *Execution) HasReferencedByOutputExecutionIdObjects() bool`

HasReferencedByOutputExecutionIdObjects returns a boolean if a field has been set.

### GetStartedAt

`func (o *Execution) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Execution) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Execution) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *Execution) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *Execution) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Execution) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Execution) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Execution) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



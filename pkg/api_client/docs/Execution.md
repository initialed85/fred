# Execution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**EndedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**M2mRuleTriggerJobId** | Pointer to **string** |  | [optional] 
**M2mRuleTriggerJobIdObject** | Pointer to [**M2MRuleTriggerJob**](M2MRuleTriggerJob.md) |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TaskId** | Pointer to **string** |  | [optional] 
**TaskIdObject** | Pointer to [**Task**](Task.md) |  | [optional] 
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

### GetM2mRuleTriggerJobId

`func (o *Execution) GetM2mRuleTriggerJobId() string`

GetM2mRuleTriggerJobId returns the M2mRuleTriggerJobId field if non-nil, zero value otherwise.

### GetM2mRuleTriggerJobIdOk

`func (o *Execution) GetM2mRuleTriggerJobIdOk() (*string, bool)`

GetM2mRuleTriggerJobIdOk returns a tuple with the M2mRuleTriggerJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM2mRuleTriggerJobId

`func (o *Execution) SetM2mRuleTriggerJobId(v string)`

SetM2mRuleTriggerJobId sets M2mRuleTriggerJobId field to given value.

### HasM2mRuleTriggerJobId

`func (o *Execution) HasM2mRuleTriggerJobId() bool`

HasM2mRuleTriggerJobId returns a boolean if a field has been set.

### GetM2mRuleTriggerJobIdObject

`func (o *Execution) GetM2mRuleTriggerJobIdObject() M2MRuleTriggerJob`

GetM2mRuleTriggerJobIdObject returns the M2mRuleTriggerJobIdObject field if non-nil, zero value otherwise.

### GetM2mRuleTriggerJobIdObjectOk

`func (o *Execution) GetM2mRuleTriggerJobIdObjectOk() (*M2MRuleTriggerJob, bool)`

GetM2mRuleTriggerJobIdObjectOk returns a tuple with the M2mRuleTriggerJobIdObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM2mRuleTriggerJobIdObject

`func (o *Execution) SetM2mRuleTriggerJobIdObject(v M2MRuleTriggerJob)`

SetM2mRuleTriggerJobIdObject sets M2mRuleTriggerJobIdObject field to given value.

### HasM2mRuleTriggerJobIdObject

`func (o *Execution) HasM2mRuleTriggerJobIdObject() bool`

HasM2mRuleTriggerJobIdObject returns a boolean if a field has been set.

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

### GetTaskId

`func (o *Execution) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *Execution) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *Execution) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *Execution) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetTaskIdObject

`func (o *Execution) GetTaskIdObject() Task`

GetTaskIdObject returns the TaskIdObject field if non-nil, zero value otherwise.

### GetTaskIdObjectOk

`func (o *Execution) GetTaskIdObjectOk() (*Task, bool)`

GetTaskIdObjectOk returns a tuple with the TaskIdObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskIdObject

`func (o *Execution) SetTaskIdObject(v Task)`

SetTaskIdObject sets TaskIdObject field to given value.

### HasTaskIdObject

`func (o *Execution) HasTaskIdObject() bool`

HasTaskIdObject returns a boolean if a field has been set.

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



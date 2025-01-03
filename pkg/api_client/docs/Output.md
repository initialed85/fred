# Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**EndedAt** | Pointer to **time.Time** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**ExitStatus** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Logid** | Pointer to **string** |  | [optional] 
**LogidObject** | Pointer to [**Log**](Log.md) |  | [optional] 
**ReferencedByLogOutputIdObjects** | Pointer to [**[]Log**](Log.md) |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TaskId** | Pointer to **string** |  | [optional] 
**TaskIdObject** | Pointer to [**Task**](Task.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewOutput

`func NewOutput() *Output`

NewOutput instantiates a new Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputWithDefaults

`func NewOutputWithDefaults() *Output`

NewOutputWithDefaults instantiates a new Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Output) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Output) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Output) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Output) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Output) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Output) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Output) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Output) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetEndedAt

`func (o *Output) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *Output) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *Output) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *Output) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetError

`func (o *Output) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Output) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Output) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Output) HasError() bool`

HasError returns a boolean if a field has been set.

### GetExitStatus

`func (o *Output) GetExitStatus() int64`

GetExitStatus returns the ExitStatus field if non-nil, zero value otherwise.

### GetExitStatusOk

`func (o *Output) GetExitStatusOk() (*int64, bool)`

GetExitStatusOk returns a tuple with the ExitStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitStatus

`func (o *Output) SetExitStatus(v int64)`

SetExitStatus sets ExitStatus field to given value.

### HasExitStatus

`func (o *Output) HasExitStatus() bool`

HasExitStatus returns a boolean if a field has been set.

### GetId

`func (o *Output) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Output) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Output) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Output) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogid

`func (o *Output) GetLogid() string`

GetLogid returns the Logid field if non-nil, zero value otherwise.

### GetLogidOk

`func (o *Output) GetLogidOk() (*string, bool)`

GetLogidOk returns a tuple with the Logid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogid

`func (o *Output) SetLogid(v string)`

SetLogid sets Logid field to given value.

### HasLogid

`func (o *Output) HasLogid() bool`

HasLogid returns a boolean if a field has been set.

### GetLogidObject

`func (o *Output) GetLogidObject() Log`

GetLogidObject returns the LogidObject field if non-nil, zero value otherwise.

### GetLogidObjectOk

`func (o *Output) GetLogidObjectOk() (*Log, bool)`

GetLogidObjectOk returns a tuple with the LogidObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogidObject

`func (o *Output) SetLogidObject(v Log)`

SetLogidObject sets LogidObject field to given value.

### HasLogidObject

`func (o *Output) HasLogidObject() bool`

HasLogidObject returns a boolean if a field has been set.

### GetReferencedByLogOutputIdObjects

`func (o *Output) GetReferencedByLogOutputIdObjects() []Log`

GetReferencedByLogOutputIdObjects returns the ReferencedByLogOutputIdObjects field if non-nil, zero value otherwise.

### GetReferencedByLogOutputIdObjectsOk

`func (o *Output) GetReferencedByLogOutputIdObjectsOk() (*[]Log, bool)`

GetReferencedByLogOutputIdObjectsOk returns a tuple with the ReferencedByLogOutputIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByLogOutputIdObjects

`func (o *Output) SetReferencedByLogOutputIdObjects(v []Log)`

SetReferencedByLogOutputIdObjects sets ReferencedByLogOutputIdObjects field to given value.

### HasReferencedByLogOutputIdObjects

`func (o *Output) HasReferencedByLogOutputIdObjects() bool`

HasReferencedByLogOutputIdObjects returns a boolean if a field has been set.

### SetReferencedByLogOutputIdObjectsNil

`func (o *Output) SetReferencedByLogOutputIdObjectsNil(b bool)`

 SetReferencedByLogOutputIdObjectsNil sets the value for ReferencedByLogOutputIdObjects to be an explicit nil

### UnsetReferencedByLogOutputIdObjects
`func (o *Output) UnsetReferencedByLogOutputIdObjects()`

UnsetReferencedByLogOutputIdObjects ensures that no value is present for ReferencedByLogOutputIdObjects, not even an explicit nil
### GetStartedAt

`func (o *Output) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Output) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Output) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *Output) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *Output) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Output) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Output) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Output) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaskId

`func (o *Output) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *Output) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *Output) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *Output) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetTaskIdObject

`func (o *Output) GetTaskIdObject() Task`

GetTaskIdObject returns the TaskIdObject field if non-nil, zero value otherwise.

### GetTaskIdObjectOk

`func (o *Output) GetTaskIdObjectOk() (*Task, bool)`

GetTaskIdObjectOk returns a tuple with the TaskIdObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskIdObject

`func (o *Output) SetTaskIdObject(v Task)`

SetTaskIdObject sets TaskIdObject field to given value.

### HasTaskIdObject

`func (o *Output) HasTaskIdObject() bool`

HasTaskIdObject returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Output) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Output) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Output) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Output) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



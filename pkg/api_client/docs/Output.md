# Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ReferencedByExecutionBuildOutputIdObjects** | Pointer to [**[]Execution**](Execution.md) |  | [optional] 
**ReferencedByExecutionDeployOutputIdObjects** | Pointer to [**[]Execution**](Execution.md) |  | [optional] 
**ReferencedByExecutionPublishOutputIdObjects** | Pointer to [**[]Execution**](Execution.md) |  | [optional] 
**ReferencedByExecutionTestOutputIdObjects** | Pointer to [**[]Execution**](Execution.md) |  | [optional] 
**ReferencedByExecutionValidateOutputIdObjects** | Pointer to [**[]Execution**](Execution.md) |  | [optional] 
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

### GetReferencedByExecutionBuildOutputIdObjects

`func (o *Output) GetReferencedByExecutionBuildOutputIdObjects() []Execution`

GetReferencedByExecutionBuildOutputIdObjects returns the ReferencedByExecutionBuildOutputIdObjects field if non-nil, zero value otherwise.

### GetReferencedByExecutionBuildOutputIdObjectsOk

`func (o *Output) GetReferencedByExecutionBuildOutputIdObjectsOk() (*[]Execution, bool)`

GetReferencedByExecutionBuildOutputIdObjectsOk returns a tuple with the ReferencedByExecutionBuildOutputIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByExecutionBuildOutputIdObjects

`func (o *Output) SetReferencedByExecutionBuildOutputIdObjects(v []Execution)`

SetReferencedByExecutionBuildOutputIdObjects sets ReferencedByExecutionBuildOutputIdObjects field to given value.

### HasReferencedByExecutionBuildOutputIdObjects

`func (o *Output) HasReferencedByExecutionBuildOutputIdObjects() bool`

HasReferencedByExecutionBuildOutputIdObjects returns a boolean if a field has been set.

### SetReferencedByExecutionBuildOutputIdObjectsNil

`func (o *Output) SetReferencedByExecutionBuildOutputIdObjectsNil(b bool)`

 SetReferencedByExecutionBuildOutputIdObjectsNil sets the value for ReferencedByExecutionBuildOutputIdObjects to be an explicit nil

### UnsetReferencedByExecutionBuildOutputIdObjects
`func (o *Output) UnsetReferencedByExecutionBuildOutputIdObjects()`

UnsetReferencedByExecutionBuildOutputIdObjects ensures that no value is present for ReferencedByExecutionBuildOutputIdObjects, not even an explicit nil
### GetReferencedByExecutionDeployOutputIdObjects

`func (o *Output) GetReferencedByExecutionDeployOutputIdObjects() []Execution`

GetReferencedByExecutionDeployOutputIdObjects returns the ReferencedByExecutionDeployOutputIdObjects field if non-nil, zero value otherwise.

### GetReferencedByExecutionDeployOutputIdObjectsOk

`func (o *Output) GetReferencedByExecutionDeployOutputIdObjectsOk() (*[]Execution, bool)`

GetReferencedByExecutionDeployOutputIdObjectsOk returns a tuple with the ReferencedByExecutionDeployOutputIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByExecutionDeployOutputIdObjects

`func (o *Output) SetReferencedByExecutionDeployOutputIdObjects(v []Execution)`

SetReferencedByExecutionDeployOutputIdObjects sets ReferencedByExecutionDeployOutputIdObjects field to given value.

### HasReferencedByExecutionDeployOutputIdObjects

`func (o *Output) HasReferencedByExecutionDeployOutputIdObjects() bool`

HasReferencedByExecutionDeployOutputIdObjects returns a boolean if a field has been set.

### SetReferencedByExecutionDeployOutputIdObjectsNil

`func (o *Output) SetReferencedByExecutionDeployOutputIdObjectsNil(b bool)`

 SetReferencedByExecutionDeployOutputIdObjectsNil sets the value for ReferencedByExecutionDeployOutputIdObjects to be an explicit nil

### UnsetReferencedByExecutionDeployOutputIdObjects
`func (o *Output) UnsetReferencedByExecutionDeployOutputIdObjects()`

UnsetReferencedByExecutionDeployOutputIdObjects ensures that no value is present for ReferencedByExecutionDeployOutputIdObjects, not even an explicit nil
### GetReferencedByExecutionPublishOutputIdObjects

`func (o *Output) GetReferencedByExecutionPublishOutputIdObjects() []Execution`

GetReferencedByExecutionPublishOutputIdObjects returns the ReferencedByExecutionPublishOutputIdObjects field if non-nil, zero value otherwise.

### GetReferencedByExecutionPublishOutputIdObjectsOk

`func (o *Output) GetReferencedByExecutionPublishOutputIdObjectsOk() (*[]Execution, bool)`

GetReferencedByExecutionPublishOutputIdObjectsOk returns a tuple with the ReferencedByExecutionPublishOutputIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByExecutionPublishOutputIdObjects

`func (o *Output) SetReferencedByExecutionPublishOutputIdObjects(v []Execution)`

SetReferencedByExecutionPublishOutputIdObjects sets ReferencedByExecutionPublishOutputIdObjects field to given value.

### HasReferencedByExecutionPublishOutputIdObjects

`func (o *Output) HasReferencedByExecutionPublishOutputIdObjects() bool`

HasReferencedByExecutionPublishOutputIdObjects returns a boolean if a field has been set.

### SetReferencedByExecutionPublishOutputIdObjectsNil

`func (o *Output) SetReferencedByExecutionPublishOutputIdObjectsNil(b bool)`

 SetReferencedByExecutionPublishOutputIdObjectsNil sets the value for ReferencedByExecutionPublishOutputIdObjects to be an explicit nil

### UnsetReferencedByExecutionPublishOutputIdObjects
`func (o *Output) UnsetReferencedByExecutionPublishOutputIdObjects()`

UnsetReferencedByExecutionPublishOutputIdObjects ensures that no value is present for ReferencedByExecutionPublishOutputIdObjects, not even an explicit nil
### GetReferencedByExecutionTestOutputIdObjects

`func (o *Output) GetReferencedByExecutionTestOutputIdObjects() []Execution`

GetReferencedByExecutionTestOutputIdObjects returns the ReferencedByExecutionTestOutputIdObjects field if non-nil, zero value otherwise.

### GetReferencedByExecutionTestOutputIdObjectsOk

`func (o *Output) GetReferencedByExecutionTestOutputIdObjectsOk() (*[]Execution, bool)`

GetReferencedByExecutionTestOutputIdObjectsOk returns a tuple with the ReferencedByExecutionTestOutputIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByExecutionTestOutputIdObjects

`func (o *Output) SetReferencedByExecutionTestOutputIdObjects(v []Execution)`

SetReferencedByExecutionTestOutputIdObjects sets ReferencedByExecutionTestOutputIdObjects field to given value.

### HasReferencedByExecutionTestOutputIdObjects

`func (o *Output) HasReferencedByExecutionTestOutputIdObjects() bool`

HasReferencedByExecutionTestOutputIdObjects returns a boolean if a field has been set.

### SetReferencedByExecutionTestOutputIdObjectsNil

`func (o *Output) SetReferencedByExecutionTestOutputIdObjectsNil(b bool)`

 SetReferencedByExecutionTestOutputIdObjectsNil sets the value for ReferencedByExecutionTestOutputIdObjects to be an explicit nil

### UnsetReferencedByExecutionTestOutputIdObjects
`func (o *Output) UnsetReferencedByExecutionTestOutputIdObjects()`

UnsetReferencedByExecutionTestOutputIdObjects ensures that no value is present for ReferencedByExecutionTestOutputIdObjects, not even an explicit nil
### GetReferencedByExecutionValidateOutputIdObjects

`func (o *Output) GetReferencedByExecutionValidateOutputIdObjects() []Execution`

GetReferencedByExecutionValidateOutputIdObjects returns the ReferencedByExecutionValidateOutputIdObjects field if non-nil, zero value otherwise.

### GetReferencedByExecutionValidateOutputIdObjectsOk

`func (o *Output) GetReferencedByExecutionValidateOutputIdObjectsOk() (*[]Execution, bool)`

GetReferencedByExecutionValidateOutputIdObjectsOk returns a tuple with the ReferencedByExecutionValidateOutputIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByExecutionValidateOutputIdObjects

`func (o *Output) SetReferencedByExecutionValidateOutputIdObjects(v []Execution)`

SetReferencedByExecutionValidateOutputIdObjects sets ReferencedByExecutionValidateOutputIdObjects field to given value.

### HasReferencedByExecutionValidateOutputIdObjects

`func (o *Output) HasReferencedByExecutionValidateOutputIdObjects() bool`

HasReferencedByExecutionValidateOutputIdObjects returns a boolean if a field has been set.

### SetReferencedByExecutionValidateOutputIdObjectsNil

`func (o *Output) SetReferencedByExecutionValidateOutputIdObjectsNil(b bool)`

 SetReferencedByExecutionValidateOutputIdObjectsNil sets the value for ReferencedByExecutionValidateOutputIdObjects to be an explicit nil

### UnsetReferencedByExecutionValidateOutputIdObjects
`func (o *Output) UnsetReferencedByExecutionValidateOutputIdObjects()`

UnsetReferencedByExecutionValidateOutputIdObjects ensures that no value is present for ReferencedByExecutionValidateOutputIdObjects, not even an explicit nil
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



# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**ReferencedByJobBuildTaskIdObjects** | Pointer to [**[]Job**](Job.md) |  | [optional] 
**ReferencedByJobDeployTaskIdObjects** | Pointer to [**[]Job**](Job.md) |  | [optional] 
**ReferencedByJobPublishTaskIdObjects** | Pointer to [**[]Job**](Job.md) |  | [optional] 
**ReferencedByJobTestTaskIdObjects** | Pointer to [**[]Job**](Job.md) |  | [optional] 
**ReferencedByJobValidateTaskIdObjects** | Pointer to [**[]Job**](Job.md) |  | [optional] 
**ReferencedByOutputTaskIdObjects** | Pointer to [**[]Output**](Output.md) |  | [optional] 
**Script** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTask

`func NewTask() *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Task) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Task) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Task) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Task) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Task) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Task) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Task) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Task) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetId

`func (o *Task) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Task) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Task) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Task) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *Task) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Task) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Task) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *Task) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetName

`func (o *Task) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Task) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Task) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Task) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatform

`func (o *Task) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *Task) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *Task) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *Task) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetReferencedByJobBuildTaskIdObjects

`func (o *Task) GetReferencedByJobBuildTaskIdObjects() []Job`

GetReferencedByJobBuildTaskIdObjects returns the ReferencedByJobBuildTaskIdObjects field if non-nil, zero value otherwise.

### GetReferencedByJobBuildTaskIdObjectsOk

`func (o *Task) GetReferencedByJobBuildTaskIdObjectsOk() (*[]Job, bool)`

GetReferencedByJobBuildTaskIdObjectsOk returns a tuple with the ReferencedByJobBuildTaskIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByJobBuildTaskIdObjects

`func (o *Task) SetReferencedByJobBuildTaskIdObjects(v []Job)`

SetReferencedByJobBuildTaskIdObjects sets ReferencedByJobBuildTaskIdObjects field to given value.

### HasReferencedByJobBuildTaskIdObjects

`func (o *Task) HasReferencedByJobBuildTaskIdObjects() bool`

HasReferencedByJobBuildTaskIdObjects returns a boolean if a field has been set.

### SetReferencedByJobBuildTaskIdObjectsNil

`func (o *Task) SetReferencedByJobBuildTaskIdObjectsNil(b bool)`

 SetReferencedByJobBuildTaskIdObjectsNil sets the value for ReferencedByJobBuildTaskIdObjects to be an explicit nil

### UnsetReferencedByJobBuildTaskIdObjects
`func (o *Task) UnsetReferencedByJobBuildTaskIdObjects()`

UnsetReferencedByJobBuildTaskIdObjects ensures that no value is present for ReferencedByJobBuildTaskIdObjects, not even an explicit nil
### GetReferencedByJobDeployTaskIdObjects

`func (o *Task) GetReferencedByJobDeployTaskIdObjects() []Job`

GetReferencedByJobDeployTaskIdObjects returns the ReferencedByJobDeployTaskIdObjects field if non-nil, zero value otherwise.

### GetReferencedByJobDeployTaskIdObjectsOk

`func (o *Task) GetReferencedByJobDeployTaskIdObjectsOk() (*[]Job, bool)`

GetReferencedByJobDeployTaskIdObjectsOk returns a tuple with the ReferencedByJobDeployTaskIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByJobDeployTaskIdObjects

`func (o *Task) SetReferencedByJobDeployTaskIdObjects(v []Job)`

SetReferencedByJobDeployTaskIdObjects sets ReferencedByJobDeployTaskIdObjects field to given value.

### HasReferencedByJobDeployTaskIdObjects

`func (o *Task) HasReferencedByJobDeployTaskIdObjects() bool`

HasReferencedByJobDeployTaskIdObjects returns a boolean if a field has been set.

### SetReferencedByJobDeployTaskIdObjectsNil

`func (o *Task) SetReferencedByJobDeployTaskIdObjectsNil(b bool)`

 SetReferencedByJobDeployTaskIdObjectsNil sets the value for ReferencedByJobDeployTaskIdObjects to be an explicit nil

### UnsetReferencedByJobDeployTaskIdObjects
`func (o *Task) UnsetReferencedByJobDeployTaskIdObjects()`

UnsetReferencedByJobDeployTaskIdObjects ensures that no value is present for ReferencedByJobDeployTaskIdObjects, not even an explicit nil
### GetReferencedByJobPublishTaskIdObjects

`func (o *Task) GetReferencedByJobPublishTaskIdObjects() []Job`

GetReferencedByJobPublishTaskIdObjects returns the ReferencedByJobPublishTaskIdObjects field if non-nil, zero value otherwise.

### GetReferencedByJobPublishTaskIdObjectsOk

`func (o *Task) GetReferencedByJobPublishTaskIdObjectsOk() (*[]Job, bool)`

GetReferencedByJobPublishTaskIdObjectsOk returns a tuple with the ReferencedByJobPublishTaskIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByJobPublishTaskIdObjects

`func (o *Task) SetReferencedByJobPublishTaskIdObjects(v []Job)`

SetReferencedByJobPublishTaskIdObjects sets ReferencedByJobPublishTaskIdObjects field to given value.

### HasReferencedByJobPublishTaskIdObjects

`func (o *Task) HasReferencedByJobPublishTaskIdObjects() bool`

HasReferencedByJobPublishTaskIdObjects returns a boolean if a field has been set.

### SetReferencedByJobPublishTaskIdObjectsNil

`func (o *Task) SetReferencedByJobPublishTaskIdObjectsNil(b bool)`

 SetReferencedByJobPublishTaskIdObjectsNil sets the value for ReferencedByJobPublishTaskIdObjects to be an explicit nil

### UnsetReferencedByJobPublishTaskIdObjects
`func (o *Task) UnsetReferencedByJobPublishTaskIdObjects()`

UnsetReferencedByJobPublishTaskIdObjects ensures that no value is present for ReferencedByJobPublishTaskIdObjects, not even an explicit nil
### GetReferencedByJobTestTaskIdObjects

`func (o *Task) GetReferencedByJobTestTaskIdObjects() []Job`

GetReferencedByJobTestTaskIdObjects returns the ReferencedByJobTestTaskIdObjects field if non-nil, zero value otherwise.

### GetReferencedByJobTestTaskIdObjectsOk

`func (o *Task) GetReferencedByJobTestTaskIdObjectsOk() (*[]Job, bool)`

GetReferencedByJobTestTaskIdObjectsOk returns a tuple with the ReferencedByJobTestTaskIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByJobTestTaskIdObjects

`func (o *Task) SetReferencedByJobTestTaskIdObjects(v []Job)`

SetReferencedByJobTestTaskIdObjects sets ReferencedByJobTestTaskIdObjects field to given value.

### HasReferencedByJobTestTaskIdObjects

`func (o *Task) HasReferencedByJobTestTaskIdObjects() bool`

HasReferencedByJobTestTaskIdObjects returns a boolean if a field has been set.

### SetReferencedByJobTestTaskIdObjectsNil

`func (o *Task) SetReferencedByJobTestTaskIdObjectsNil(b bool)`

 SetReferencedByJobTestTaskIdObjectsNil sets the value for ReferencedByJobTestTaskIdObjects to be an explicit nil

### UnsetReferencedByJobTestTaskIdObjects
`func (o *Task) UnsetReferencedByJobTestTaskIdObjects()`

UnsetReferencedByJobTestTaskIdObjects ensures that no value is present for ReferencedByJobTestTaskIdObjects, not even an explicit nil
### GetReferencedByJobValidateTaskIdObjects

`func (o *Task) GetReferencedByJobValidateTaskIdObjects() []Job`

GetReferencedByJobValidateTaskIdObjects returns the ReferencedByJobValidateTaskIdObjects field if non-nil, zero value otherwise.

### GetReferencedByJobValidateTaskIdObjectsOk

`func (o *Task) GetReferencedByJobValidateTaskIdObjectsOk() (*[]Job, bool)`

GetReferencedByJobValidateTaskIdObjectsOk returns a tuple with the ReferencedByJobValidateTaskIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByJobValidateTaskIdObjects

`func (o *Task) SetReferencedByJobValidateTaskIdObjects(v []Job)`

SetReferencedByJobValidateTaskIdObjects sets ReferencedByJobValidateTaskIdObjects field to given value.

### HasReferencedByJobValidateTaskIdObjects

`func (o *Task) HasReferencedByJobValidateTaskIdObjects() bool`

HasReferencedByJobValidateTaskIdObjects returns a boolean if a field has been set.

### SetReferencedByJobValidateTaskIdObjectsNil

`func (o *Task) SetReferencedByJobValidateTaskIdObjectsNil(b bool)`

 SetReferencedByJobValidateTaskIdObjectsNil sets the value for ReferencedByJobValidateTaskIdObjects to be an explicit nil

### UnsetReferencedByJobValidateTaskIdObjects
`func (o *Task) UnsetReferencedByJobValidateTaskIdObjects()`

UnsetReferencedByJobValidateTaskIdObjects ensures that no value is present for ReferencedByJobValidateTaskIdObjects, not even an explicit nil
### GetReferencedByOutputTaskIdObjects

`func (o *Task) GetReferencedByOutputTaskIdObjects() []Output`

GetReferencedByOutputTaskIdObjects returns the ReferencedByOutputTaskIdObjects field if non-nil, zero value otherwise.

### GetReferencedByOutputTaskIdObjectsOk

`func (o *Task) GetReferencedByOutputTaskIdObjectsOk() (*[]Output, bool)`

GetReferencedByOutputTaskIdObjectsOk returns a tuple with the ReferencedByOutputTaskIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByOutputTaskIdObjects

`func (o *Task) SetReferencedByOutputTaskIdObjects(v []Output)`

SetReferencedByOutputTaskIdObjects sets ReferencedByOutputTaskIdObjects field to given value.

### HasReferencedByOutputTaskIdObjects

`func (o *Task) HasReferencedByOutputTaskIdObjects() bool`

HasReferencedByOutputTaskIdObjects returns a boolean if a field has been set.

### SetReferencedByOutputTaskIdObjectsNil

`func (o *Task) SetReferencedByOutputTaskIdObjectsNil(b bool)`

 SetReferencedByOutputTaskIdObjectsNil sets the value for ReferencedByOutputTaskIdObjects to be an explicit nil

### UnsetReferencedByOutputTaskIdObjects
`func (o *Task) UnsetReferencedByOutputTaskIdObjects()`

UnsetReferencedByOutputTaskIdObjects ensures that no value is present for ReferencedByOutputTaskIdObjects, not even an explicit nil
### GetScript

`func (o *Task) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *Task) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *Task) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *Task) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Task) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Task) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Task) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Task) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



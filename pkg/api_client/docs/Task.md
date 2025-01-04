# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Index** | Pointer to **int64** |  | [optional] 
**JobId** | Pointer to **string** |  | [optional] 
**JobIdObject** | Pointer to [**Job**](Job.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
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

### GetIndex

`func (o *Task) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Task) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Task) SetIndex(v int64)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *Task) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetJobId

`func (o *Task) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *Task) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *Task) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *Task) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetJobIdObject

`func (o *Task) GetJobIdObject() Job`

GetJobIdObject returns the JobIdObject field if non-nil, zero value otherwise.

### GetJobIdObjectOk

`func (o *Task) GetJobIdObjectOk() (*Job, bool)`

GetJobIdObjectOk returns a tuple with the JobIdObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobIdObject

`func (o *Task) SetJobIdObject(v Job)`

SetJobIdObject sets JobIdObject field to given value.

### HasJobIdObject

`func (o *Task) HasJobIdObject() bool`

HasJobIdObject returns a boolean if a field has been set.

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



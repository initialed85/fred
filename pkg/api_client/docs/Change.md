# Change

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ReferencedByTriggerChangeIdObjects** | Pointer to [**[]Trigger**](Trigger.md) |  | [optional] 
**RepositoryId** | Pointer to **string** |  | [optional] 
**RepositoryIdObject** | Pointer to [**Repository**](Repository.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewChange

`func NewChange() *Change`

NewChange instantiates a new Change object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeWithDefaults

`func NewChangeWithDefaults() *Change`

NewChangeWithDefaults instantiates a new Change object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Change) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Change) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Change) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Change) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Change) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Change) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Change) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Change) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetId

`func (o *Change) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Change) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Change) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Change) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReferencedByTriggerChangeIdObjects

`func (o *Change) GetReferencedByTriggerChangeIdObjects() []Trigger`

GetReferencedByTriggerChangeIdObjects returns the ReferencedByTriggerChangeIdObjects field if non-nil, zero value otherwise.

### GetReferencedByTriggerChangeIdObjectsOk

`func (o *Change) GetReferencedByTriggerChangeIdObjectsOk() (*[]Trigger, bool)`

GetReferencedByTriggerChangeIdObjectsOk returns a tuple with the ReferencedByTriggerChangeIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByTriggerChangeIdObjects

`func (o *Change) SetReferencedByTriggerChangeIdObjects(v []Trigger)`

SetReferencedByTriggerChangeIdObjects sets ReferencedByTriggerChangeIdObjects field to given value.

### HasReferencedByTriggerChangeIdObjects

`func (o *Change) HasReferencedByTriggerChangeIdObjects() bool`

HasReferencedByTriggerChangeIdObjects returns a boolean if a field has been set.

### SetReferencedByTriggerChangeIdObjectsNil

`func (o *Change) SetReferencedByTriggerChangeIdObjectsNil(b bool)`

 SetReferencedByTriggerChangeIdObjectsNil sets the value for ReferencedByTriggerChangeIdObjects to be an explicit nil

### UnsetReferencedByTriggerChangeIdObjects
`func (o *Change) UnsetReferencedByTriggerChangeIdObjects()`

UnsetReferencedByTriggerChangeIdObjects ensures that no value is present for ReferencedByTriggerChangeIdObjects, not even an explicit nil
### GetRepositoryId

`func (o *Change) GetRepositoryId() string`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *Change) GetRepositoryIdOk() (*string, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *Change) SetRepositoryId(v string)`

SetRepositoryId sets RepositoryId field to given value.

### HasRepositoryId

`func (o *Change) HasRepositoryId() bool`

HasRepositoryId returns a boolean if a field has been set.

### GetRepositoryIdObject

`func (o *Change) GetRepositoryIdObject() Repository`

GetRepositoryIdObject returns the RepositoryIdObject field if non-nil, zero value otherwise.

### GetRepositoryIdObjectOk

`func (o *Change) GetRepositoryIdObjectOk() (*Repository, bool)`

GetRepositoryIdObjectOk returns a tuple with the RepositoryIdObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryIdObject

`func (o *Change) SetRepositoryIdObject(v Repository)`

SetRepositoryIdObject sets RepositoryIdObject field to given value.

### HasRepositoryIdObject

`func (o *Change) HasRepositoryIdObject() bool`

HasRepositoryIdObject returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Change) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Change) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Change) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Change) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



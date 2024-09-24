# Change

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthoredAt** | Pointer to **time.Time** |  | [optional] 
**AuthoredBy** | Pointer to **string** |  | [optional] 
**BranchName** | Pointer to **string** |  | [optional] 
**CommitHash** | Pointer to **string** |  | [optional] 
**CommittedAt** | Pointer to **time.Time** |  | [optional] 
**CommittedBy** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**ReferencedByTriggerChangeIdObjects** | Pointer to [**[]Trigger**](Trigger.md) |  | [optional] 
**RepositoryId** | Pointer to **string** |  | [optional] 
**RepositoryIdObject** | Pointer to [**Repository**](Repository.md) |  | [optional] 
**TriggerProducedAt** | Pointer to **time.Time** |  | [optional] 
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

### GetAuthoredAt

`func (o *Change) GetAuthoredAt() time.Time`

GetAuthoredAt returns the AuthoredAt field if non-nil, zero value otherwise.

### GetAuthoredAtOk

`func (o *Change) GetAuthoredAtOk() (*time.Time, bool)`

GetAuthoredAtOk returns a tuple with the AuthoredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthoredAt

`func (o *Change) SetAuthoredAt(v time.Time)`

SetAuthoredAt sets AuthoredAt field to given value.

### HasAuthoredAt

`func (o *Change) HasAuthoredAt() bool`

HasAuthoredAt returns a boolean if a field has been set.

### GetAuthoredBy

`func (o *Change) GetAuthoredBy() string`

GetAuthoredBy returns the AuthoredBy field if non-nil, zero value otherwise.

### GetAuthoredByOk

`func (o *Change) GetAuthoredByOk() (*string, bool)`

GetAuthoredByOk returns a tuple with the AuthoredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthoredBy

`func (o *Change) SetAuthoredBy(v string)`

SetAuthoredBy sets AuthoredBy field to given value.

### HasAuthoredBy

`func (o *Change) HasAuthoredBy() bool`

HasAuthoredBy returns a boolean if a field has been set.

### GetBranchName

`func (o *Change) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *Change) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *Change) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.

### HasBranchName

`func (o *Change) HasBranchName() bool`

HasBranchName returns a boolean if a field has been set.

### GetCommitHash

`func (o *Change) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *Change) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *Change) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *Change) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### GetCommittedAt

`func (o *Change) GetCommittedAt() time.Time`

GetCommittedAt returns the CommittedAt field if non-nil, zero value otherwise.

### GetCommittedAtOk

`func (o *Change) GetCommittedAtOk() (*time.Time, bool)`

GetCommittedAtOk returns a tuple with the CommittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommittedAt

`func (o *Change) SetCommittedAt(v time.Time)`

SetCommittedAt sets CommittedAt field to given value.

### HasCommittedAt

`func (o *Change) HasCommittedAt() bool`

HasCommittedAt returns a boolean if a field has been set.

### GetCommittedBy

`func (o *Change) GetCommittedBy() string`

GetCommittedBy returns the CommittedBy field if non-nil, zero value otherwise.

### GetCommittedByOk

`func (o *Change) GetCommittedByOk() (*string, bool)`

GetCommittedByOk returns a tuple with the CommittedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommittedBy

`func (o *Change) SetCommittedBy(v string)`

SetCommittedBy sets CommittedBy field to given value.

### HasCommittedBy

`func (o *Change) HasCommittedBy() bool`

HasCommittedBy returns a boolean if a field has been set.

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

### GetMessage

`func (o *Change) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Change) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Change) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Change) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

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

### GetTriggerProducedAt

`func (o *Change) GetTriggerProducedAt() time.Time`

GetTriggerProducedAt returns the TriggerProducedAt field if non-nil, zero value otherwise.

### GetTriggerProducedAtOk

`func (o *Change) GetTriggerProducedAtOk() (*time.Time, bool)`

GetTriggerProducedAtOk returns a tuple with the TriggerProducedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerProducedAt

`func (o *Change) SetTriggerProducedAt(v time.Time)`

SetTriggerProducedAt sets TriggerProducedAt field to given value.

### HasTriggerProducedAt

`func (o *Change) HasTriggerProducedAt() bool`

HasTriggerProducedAt returns a boolean if a field has been set.

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



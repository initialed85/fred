# Repository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**HandledAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ReferencedByChangeRepositoryIdObjects** | Pointer to [**[]Change**](Change.md) |  | [optional] 
**ReferencedByJobRepositoryIdObjects** | Pointer to [**[]Job**](Job.md) |  | [optional] 
**RepositorySyncerClaimedUntil** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewRepository

`func NewRepository() *Repository`

NewRepository instantiates a new Repository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryWithDefaults

`func NewRepositoryWithDefaults() *Repository`

NewRepositoryWithDefaults instantiates a new Repository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Repository) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Repository) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Repository) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Repository) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Repository) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Repository) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Repository) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Repository) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetHandledAt

`func (o *Repository) GetHandledAt() time.Time`

GetHandledAt returns the HandledAt field if non-nil, zero value otherwise.

### GetHandledAtOk

`func (o *Repository) GetHandledAtOk() (*time.Time, bool)`

GetHandledAtOk returns a tuple with the HandledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandledAt

`func (o *Repository) SetHandledAt(v time.Time)`

SetHandledAt sets HandledAt field to given value.

### HasHandledAt

`func (o *Repository) HasHandledAt() bool`

HasHandledAt returns a boolean if a field has been set.

### GetId

`func (o *Repository) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Repository) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Repository) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Repository) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Repository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Repository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Repository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Repository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReferencedByChangeRepositoryIdObjects

`func (o *Repository) GetReferencedByChangeRepositoryIdObjects() []Change`

GetReferencedByChangeRepositoryIdObjects returns the ReferencedByChangeRepositoryIdObjects field if non-nil, zero value otherwise.

### GetReferencedByChangeRepositoryIdObjectsOk

`func (o *Repository) GetReferencedByChangeRepositoryIdObjectsOk() (*[]Change, bool)`

GetReferencedByChangeRepositoryIdObjectsOk returns a tuple with the ReferencedByChangeRepositoryIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByChangeRepositoryIdObjects

`func (o *Repository) SetReferencedByChangeRepositoryIdObjects(v []Change)`

SetReferencedByChangeRepositoryIdObjects sets ReferencedByChangeRepositoryIdObjects field to given value.

### HasReferencedByChangeRepositoryIdObjects

`func (o *Repository) HasReferencedByChangeRepositoryIdObjects() bool`

HasReferencedByChangeRepositoryIdObjects returns a boolean if a field has been set.

### GetReferencedByJobRepositoryIdObjects

`func (o *Repository) GetReferencedByJobRepositoryIdObjects() []Job`

GetReferencedByJobRepositoryIdObjects returns the ReferencedByJobRepositoryIdObjects field if non-nil, zero value otherwise.

### GetReferencedByJobRepositoryIdObjectsOk

`func (o *Repository) GetReferencedByJobRepositoryIdObjectsOk() (*[]Job, bool)`

GetReferencedByJobRepositoryIdObjectsOk returns a tuple with the ReferencedByJobRepositoryIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByJobRepositoryIdObjects

`func (o *Repository) SetReferencedByJobRepositoryIdObjects(v []Job)`

SetReferencedByJobRepositoryIdObjects sets ReferencedByJobRepositoryIdObjects field to given value.

### HasReferencedByJobRepositoryIdObjects

`func (o *Repository) HasReferencedByJobRepositoryIdObjects() bool`

HasReferencedByJobRepositoryIdObjects returns a boolean if a field has been set.

### GetRepositorySyncerClaimedUntil

`func (o *Repository) GetRepositorySyncerClaimedUntil() time.Time`

GetRepositorySyncerClaimedUntil returns the RepositorySyncerClaimedUntil field if non-nil, zero value otherwise.

### GetRepositorySyncerClaimedUntilOk

`func (o *Repository) GetRepositorySyncerClaimedUntilOk() (*time.Time, bool)`

GetRepositorySyncerClaimedUntilOk returns a tuple with the RepositorySyncerClaimedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositorySyncerClaimedUntil

`func (o *Repository) SetRepositorySyncerClaimedUntil(v time.Time)`

SetRepositorySyncerClaimedUntil sets RepositorySyncerClaimedUntil field to given value.

### HasRepositorySyncerClaimedUntil

`func (o *Repository) HasRepositorySyncerClaimedUntil() bool`

HasRepositorySyncerClaimedUntil returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Repository) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Repository) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Repository) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Repository) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *Repository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Repository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Repository) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Repository) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



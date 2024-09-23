# Repository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastSynced** | Pointer to **time.Time** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**ReferencedByChangeRepositoryIdObjects** | Pointer to [**[]Change**](Change.md) |  | [optional] 
**ReferencedByRuleRepositoryIdObjects** | Pointer to [**[]Rule**](Rule.md) |  | [optional] 
**SshKey** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

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

### GetLastSynced

`func (o *Repository) GetLastSynced() time.Time`

GetLastSynced returns the LastSynced field if non-nil, zero value otherwise.

### GetLastSyncedOk

`func (o *Repository) GetLastSyncedOk() (*time.Time, bool)`

GetLastSyncedOk returns a tuple with the LastSynced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSynced

`func (o *Repository) SetLastSynced(v time.Time)`

SetLastSynced sets LastSynced field to given value.

### HasLastSynced

`func (o *Repository) HasLastSynced() bool`

HasLastSynced returns a boolean if a field has been set.

### GetPassword

`func (o *Repository) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Repository) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Repository) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Repository) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

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

### SetReferencedByChangeRepositoryIdObjectsNil

`func (o *Repository) SetReferencedByChangeRepositoryIdObjectsNil(b bool)`

 SetReferencedByChangeRepositoryIdObjectsNil sets the value for ReferencedByChangeRepositoryIdObjects to be an explicit nil

### UnsetReferencedByChangeRepositoryIdObjects
`func (o *Repository) UnsetReferencedByChangeRepositoryIdObjects()`

UnsetReferencedByChangeRepositoryIdObjects ensures that no value is present for ReferencedByChangeRepositoryIdObjects, not even an explicit nil
### GetReferencedByRuleRepositoryIdObjects

`func (o *Repository) GetReferencedByRuleRepositoryIdObjects() []Rule`

GetReferencedByRuleRepositoryIdObjects returns the ReferencedByRuleRepositoryIdObjects field if non-nil, zero value otherwise.

### GetReferencedByRuleRepositoryIdObjectsOk

`func (o *Repository) GetReferencedByRuleRepositoryIdObjectsOk() (*[]Rule, bool)`

GetReferencedByRuleRepositoryIdObjectsOk returns a tuple with the ReferencedByRuleRepositoryIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByRuleRepositoryIdObjects

`func (o *Repository) SetReferencedByRuleRepositoryIdObjects(v []Rule)`

SetReferencedByRuleRepositoryIdObjects sets ReferencedByRuleRepositoryIdObjects field to given value.

### HasReferencedByRuleRepositoryIdObjects

`func (o *Repository) HasReferencedByRuleRepositoryIdObjects() bool`

HasReferencedByRuleRepositoryIdObjects returns a boolean if a field has been set.

### SetReferencedByRuleRepositoryIdObjectsNil

`func (o *Repository) SetReferencedByRuleRepositoryIdObjectsNil(b bool)`

 SetReferencedByRuleRepositoryIdObjectsNil sets the value for ReferencedByRuleRepositoryIdObjects to be an explicit nil

### UnsetReferencedByRuleRepositoryIdObjects
`func (o *Repository) UnsetReferencedByRuleRepositoryIdObjects()`

UnsetReferencedByRuleRepositoryIdObjects ensures that no value is present for ReferencedByRuleRepositoryIdObjects, not even an explicit nil
### GetSshKey

`func (o *Repository) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *Repository) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *Repository) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *Repository) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

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

### GetUsername

`func (o *Repository) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Repository) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Repository) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Repository) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



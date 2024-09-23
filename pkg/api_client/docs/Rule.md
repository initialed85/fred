# Rule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ReferencedByJobRuleIdObjects** | Pointer to [**[]Job**](Job.md) |  | [optional] 
**ReferencedByRuleRequiresJobRuleIdObjects** | Pointer to [**[]RuleRequiresJob**](RuleRequiresJob.md) |  | [optional] 
**ReferencedByTriggerRuleIdObjects** | Pointer to [**[]Trigger**](Trigger.md) |  | [optional] 
**RepositoryId** | Pointer to **string** |  | [optional] 
**RepositoryIdObject** | Pointer to [**Repository**](Repository.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRule

`func NewRule() *Rule`

NewRule instantiates a new Rule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleWithDefaults

`func NewRuleWithDefaults() *Rule`

NewRuleWithDefaults instantiates a new Rule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Rule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Rule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Rule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Rule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Rule) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Rule) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Rule) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Rule) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetId

`func (o *Rule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Rule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Rule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Rule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReferencedByJobRuleIdObjects

`func (o *Rule) GetReferencedByJobRuleIdObjects() []Job`

GetReferencedByJobRuleIdObjects returns the ReferencedByJobRuleIdObjects field if non-nil, zero value otherwise.

### GetReferencedByJobRuleIdObjectsOk

`func (o *Rule) GetReferencedByJobRuleIdObjectsOk() (*[]Job, bool)`

GetReferencedByJobRuleIdObjectsOk returns a tuple with the ReferencedByJobRuleIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByJobRuleIdObjects

`func (o *Rule) SetReferencedByJobRuleIdObjects(v []Job)`

SetReferencedByJobRuleIdObjects sets ReferencedByJobRuleIdObjects field to given value.

### HasReferencedByJobRuleIdObjects

`func (o *Rule) HasReferencedByJobRuleIdObjects() bool`

HasReferencedByJobRuleIdObjects returns a boolean if a field has been set.

### SetReferencedByJobRuleIdObjectsNil

`func (o *Rule) SetReferencedByJobRuleIdObjectsNil(b bool)`

 SetReferencedByJobRuleIdObjectsNil sets the value for ReferencedByJobRuleIdObjects to be an explicit nil

### UnsetReferencedByJobRuleIdObjects
`func (o *Rule) UnsetReferencedByJobRuleIdObjects()`

UnsetReferencedByJobRuleIdObjects ensures that no value is present for ReferencedByJobRuleIdObjects, not even an explicit nil
### GetReferencedByRuleRequiresJobRuleIdObjects

`func (o *Rule) GetReferencedByRuleRequiresJobRuleIdObjects() []RuleRequiresJob`

GetReferencedByRuleRequiresJobRuleIdObjects returns the ReferencedByRuleRequiresJobRuleIdObjects field if non-nil, zero value otherwise.

### GetReferencedByRuleRequiresJobRuleIdObjectsOk

`func (o *Rule) GetReferencedByRuleRequiresJobRuleIdObjectsOk() (*[]RuleRequiresJob, bool)`

GetReferencedByRuleRequiresJobRuleIdObjectsOk returns a tuple with the ReferencedByRuleRequiresJobRuleIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByRuleRequiresJobRuleIdObjects

`func (o *Rule) SetReferencedByRuleRequiresJobRuleIdObjects(v []RuleRequiresJob)`

SetReferencedByRuleRequiresJobRuleIdObjects sets ReferencedByRuleRequiresJobRuleIdObjects field to given value.

### HasReferencedByRuleRequiresJobRuleIdObjects

`func (o *Rule) HasReferencedByRuleRequiresJobRuleIdObjects() bool`

HasReferencedByRuleRequiresJobRuleIdObjects returns a boolean if a field has been set.

### SetReferencedByRuleRequiresJobRuleIdObjectsNil

`func (o *Rule) SetReferencedByRuleRequiresJobRuleIdObjectsNil(b bool)`

 SetReferencedByRuleRequiresJobRuleIdObjectsNil sets the value for ReferencedByRuleRequiresJobRuleIdObjects to be an explicit nil

### UnsetReferencedByRuleRequiresJobRuleIdObjects
`func (o *Rule) UnsetReferencedByRuleRequiresJobRuleIdObjects()`

UnsetReferencedByRuleRequiresJobRuleIdObjects ensures that no value is present for ReferencedByRuleRequiresJobRuleIdObjects, not even an explicit nil
### GetReferencedByTriggerRuleIdObjects

`func (o *Rule) GetReferencedByTriggerRuleIdObjects() []Trigger`

GetReferencedByTriggerRuleIdObjects returns the ReferencedByTriggerRuleIdObjects field if non-nil, zero value otherwise.

### GetReferencedByTriggerRuleIdObjectsOk

`func (o *Rule) GetReferencedByTriggerRuleIdObjectsOk() (*[]Trigger, bool)`

GetReferencedByTriggerRuleIdObjectsOk returns a tuple with the ReferencedByTriggerRuleIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByTriggerRuleIdObjects

`func (o *Rule) SetReferencedByTriggerRuleIdObjects(v []Trigger)`

SetReferencedByTriggerRuleIdObjects sets ReferencedByTriggerRuleIdObjects field to given value.

### HasReferencedByTriggerRuleIdObjects

`func (o *Rule) HasReferencedByTriggerRuleIdObjects() bool`

HasReferencedByTriggerRuleIdObjects returns a boolean if a field has been set.

### SetReferencedByTriggerRuleIdObjectsNil

`func (o *Rule) SetReferencedByTriggerRuleIdObjectsNil(b bool)`

 SetReferencedByTriggerRuleIdObjectsNil sets the value for ReferencedByTriggerRuleIdObjects to be an explicit nil

### UnsetReferencedByTriggerRuleIdObjects
`func (o *Rule) UnsetReferencedByTriggerRuleIdObjects()`

UnsetReferencedByTriggerRuleIdObjects ensures that no value is present for ReferencedByTriggerRuleIdObjects, not even an explicit nil
### GetRepositoryId

`func (o *Rule) GetRepositoryId() string`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *Rule) GetRepositoryIdOk() (*string, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *Rule) SetRepositoryId(v string)`

SetRepositoryId sets RepositoryId field to given value.

### HasRepositoryId

`func (o *Rule) HasRepositoryId() bool`

HasRepositoryId returns a boolean if a field has been set.

### GetRepositoryIdObject

`func (o *Rule) GetRepositoryIdObject() Repository`

GetRepositoryIdObject returns the RepositoryIdObject field if non-nil, zero value otherwise.

### GetRepositoryIdObjectOk

`func (o *Rule) GetRepositoryIdObjectOk() (*Repository, bool)`

GetRepositoryIdObjectOk returns a tuple with the RepositoryIdObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryIdObject

`func (o *Rule) SetRepositoryIdObject(v Repository)`

SetRepositoryIdObject sets RepositoryIdObject field to given value.

### HasRepositoryIdObject

`func (o *Rule) HasRepositoryIdObject() bool`

HasRepositoryIdObject returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Rule) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Rule) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Rule) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Rule) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



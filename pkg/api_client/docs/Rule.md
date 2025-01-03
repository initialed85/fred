# Rule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchName** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**JobTriggerJobId** | Pointer to **string** |  | [optional] 
**JobTriggerJobIdObject** | Pointer to [**Job**](Job.md) |  | [optional] 
**ReferencedByJobRuleTriggerRuleIdObjects** | Pointer to [**[]Rule**](Rule.md) |  | [optional] 
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

### GetBranchName

`func (o *Rule) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *Rule) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *Rule) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.

### HasBranchName

`func (o *Rule) HasBranchName() bool`

HasBranchName returns a boolean if a field has been set.

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

### GetJobTriggerJobId

`func (o *Rule) GetJobTriggerJobId() string`

GetJobTriggerJobId returns the JobTriggerJobId field if non-nil, zero value otherwise.

### GetJobTriggerJobIdOk

`func (o *Rule) GetJobTriggerJobIdOk() (*string, bool)`

GetJobTriggerJobIdOk returns a tuple with the JobTriggerJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTriggerJobId

`func (o *Rule) SetJobTriggerJobId(v string)`

SetJobTriggerJobId sets JobTriggerJobId field to given value.

### HasJobTriggerJobId

`func (o *Rule) HasJobTriggerJobId() bool`

HasJobTriggerJobId returns a boolean if a field has been set.

### GetJobTriggerJobIdObject

`func (o *Rule) GetJobTriggerJobIdObject() Job`

GetJobTriggerJobIdObject returns the JobTriggerJobIdObject field if non-nil, zero value otherwise.

### GetJobTriggerJobIdObjectOk

`func (o *Rule) GetJobTriggerJobIdObjectOk() (*Job, bool)`

GetJobTriggerJobIdObjectOk returns a tuple with the JobTriggerJobIdObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTriggerJobIdObject

`func (o *Rule) SetJobTriggerJobIdObject(v Job)`

SetJobTriggerJobIdObject sets JobTriggerJobIdObject field to given value.

### HasJobTriggerJobIdObject

`func (o *Rule) HasJobTriggerJobIdObject() bool`

HasJobTriggerJobIdObject returns a boolean if a field has been set.

### GetReferencedByJobRuleTriggerRuleIdObjects

`func (o *Rule) GetReferencedByJobRuleTriggerRuleIdObjects() []Rule`

GetReferencedByJobRuleTriggerRuleIdObjects returns the ReferencedByJobRuleTriggerRuleIdObjects field if non-nil, zero value otherwise.

### GetReferencedByJobRuleTriggerRuleIdObjectsOk

`func (o *Rule) GetReferencedByJobRuleTriggerRuleIdObjectsOk() (*[]Rule, bool)`

GetReferencedByJobRuleTriggerRuleIdObjectsOk returns a tuple with the ReferencedByJobRuleTriggerRuleIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByJobRuleTriggerRuleIdObjects

`func (o *Rule) SetReferencedByJobRuleTriggerRuleIdObjects(v []Rule)`

SetReferencedByJobRuleTriggerRuleIdObjects sets ReferencedByJobRuleTriggerRuleIdObjects field to given value.

### HasReferencedByJobRuleTriggerRuleIdObjects

`func (o *Rule) HasReferencedByJobRuleTriggerRuleIdObjects() bool`

HasReferencedByJobRuleTriggerRuleIdObjects returns a boolean if a field has been set.

### SetReferencedByJobRuleTriggerRuleIdObjectsNil

`func (o *Rule) SetReferencedByJobRuleTriggerRuleIdObjectsNil(b bool)`

 SetReferencedByJobRuleTriggerRuleIdObjectsNil sets the value for ReferencedByJobRuleTriggerRuleIdObjects to be an explicit nil

### UnsetReferencedByJobRuleTriggerRuleIdObjects
`func (o *Rule) UnsetReferencedByJobRuleTriggerRuleIdObjects()`

UnsetReferencedByJobRuleTriggerRuleIdObjects ensures that no value is present for ReferencedByJobRuleTriggerRuleIdObjects, not even an explicit nil
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



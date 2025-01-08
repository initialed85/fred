# Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branches** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ReferencedByDependsOnSinkJobIdObjects** | Pointer to [**[]DependsOn**](DependsOn.md) |  | [optional] 
**ReferencedByDependsOnSourceJobIdObjects** | Pointer to [**[]DependsOn**](DependsOn.md) |  | [optional] 
**ReferencedByExecutionJobIdObjects** | Pointer to [**[]Execution**](Execution.md) |  | [optional] 
**ReferencedByTaskJobIdObjects** | Pointer to [**[]Task**](Task.md) |  | [optional] 
**RepositoryId** | Pointer to **string** |  | [optional] 
**RepositoryIdObject** | Pointer to [**Repository**](Repository.md) |  | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewJob

`func NewJob() *Job`

NewJob instantiates a new Job object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobWithDefaults

`func NewJobWithDefaults() *Job`

NewJobWithDefaults instantiates a new Job object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranches

`func (o *Job) GetBranches() string`

GetBranches returns the Branches field if non-nil, zero value otherwise.

### GetBranchesOk

`func (o *Job) GetBranchesOk() (*string, bool)`

GetBranchesOk returns a tuple with the Branches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranches

`func (o *Job) SetBranches(v string)`

SetBranches sets Branches field to given value.

### HasBranches

`func (o *Job) HasBranches() bool`

HasBranches returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Job) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Job) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Job) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Job) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Job) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Job) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Job) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Job) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetId

`func (o *Job) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Job) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Job) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Job) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Job) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Job) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Job) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Job) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReferencedByDependsOnSinkJobIdObjects

`func (o *Job) GetReferencedByDependsOnSinkJobIdObjects() []DependsOn`

GetReferencedByDependsOnSinkJobIdObjects returns the ReferencedByDependsOnSinkJobIdObjects field if non-nil, zero value otherwise.

### GetReferencedByDependsOnSinkJobIdObjectsOk

`func (o *Job) GetReferencedByDependsOnSinkJobIdObjectsOk() (*[]DependsOn, bool)`

GetReferencedByDependsOnSinkJobIdObjectsOk returns a tuple with the ReferencedByDependsOnSinkJobIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByDependsOnSinkJobIdObjects

`func (o *Job) SetReferencedByDependsOnSinkJobIdObjects(v []DependsOn)`

SetReferencedByDependsOnSinkJobIdObjects sets ReferencedByDependsOnSinkJobIdObjects field to given value.

### HasReferencedByDependsOnSinkJobIdObjects

`func (o *Job) HasReferencedByDependsOnSinkJobIdObjects() bool`

HasReferencedByDependsOnSinkJobIdObjects returns a boolean if a field has been set.

### GetReferencedByDependsOnSourceJobIdObjects

`func (o *Job) GetReferencedByDependsOnSourceJobIdObjects() []DependsOn`

GetReferencedByDependsOnSourceJobIdObjects returns the ReferencedByDependsOnSourceJobIdObjects field if non-nil, zero value otherwise.

### GetReferencedByDependsOnSourceJobIdObjectsOk

`func (o *Job) GetReferencedByDependsOnSourceJobIdObjectsOk() (*[]DependsOn, bool)`

GetReferencedByDependsOnSourceJobIdObjectsOk returns a tuple with the ReferencedByDependsOnSourceJobIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByDependsOnSourceJobIdObjects

`func (o *Job) SetReferencedByDependsOnSourceJobIdObjects(v []DependsOn)`

SetReferencedByDependsOnSourceJobIdObjects sets ReferencedByDependsOnSourceJobIdObjects field to given value.

### HasReferencedByDependsOnSourceJobIdObjects

`func (o *Job) HasReferencedByDependsOnSourceJobIdObjects() bool`

HasReferencedByDependsOnSourceJobIdObjects returns a boolean if a field has been set.

### GetReferencedByExecutionJobIdObjects

`func (o *Job) GetReferencedByExecutionJobIdObjects() []Execution`

GetReferencedByExecutionJobIdObjects returns the ReferencedByExecutionJobIdObjects field if non-nil, zero value otherwise.

### GetReferencedByExecutionJobIdObjectsOk

`func (o *Job) GetReferencedByExecutionJobIdObjectsOk() (*[]Execution, bool)`

GetReferencedByExecutionJobIdObjectsOk returns a tuple with the ReferencedByExecutionJobIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByExecutionJobIdObjects

`func (o *Job) SetReferencedByExecutionJobIdObjects(v []Execution)`

SetReferencedByExecutionJobIdObjects sets ReferencedByExecutionJobIdObjects field to given value.

### HasReferencedByExecutionJobIdObjects

`func (o *Job) HasReferencedByExecutionJobIdObjects() bool`

HasReferencedByExecutionJobIdObjects returns a boolean if a field has been set.

### GetReferencedByTaskJobIdObjects

`func (o *Job) GetReferencedByTaskJobIdObjects() []Task`

GetReferencedByTaskJobIdObjects returns the ReferencedByTaskJobIdObjects field if non-nil, zero value otherwise.

### GetReferencedByTaskJobIdObjectsOk

`func (o *Job) GetReferencedByTaskJobIdObjectsOk() (*[]Task, bool)`

GetReferencedByTaskJobIdObjectsOk returns a tuple with the ReferencedByTaskJobIdObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByTaskJobIdObjects

`func (o *Job) SetReferencedByTaskJobIdObjects(v []Task)`

SetReferencedByTaskJobIdObjects sets ReferencedByTaskJobIdObjects field to given value.

### HasReferencedByTaskJobIdObjects

`func (o *Job) HasReferencedByTaskJobIdObjects() bool`

HasReferencedByTaskJobIdObjects returns a boolean if a field has been set.

### GetRepositoryId

`func (o *Job) GetRepositoryId() string`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *Job) GetRepositoryIdOk() (*string, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *Job) SetRepositoryId(v string)`

SetRepositoryId sets RepositoryId field to given value.

### HasRepositoryId

`func (o *Job) HasRepositoryId() bool`

HasRepositoryId returns a boolean if a field has been set.

### GetRepositoryIdObject

`func (o *Job) GetRepositoryIdObject() Repository`

GetRepositoryIdObject returns the RepositoryIdObject field if non-nil, zero value otherwise.

### GetRepositoryIdObjectOk

`func (o *Job) GetRepositoryIdObjectOk() (*Repository, bool)`

GetRepositoryIdObjectOk returns a tuple with the RepositoryIdObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryIdObject

`func (o *Job) SetRepositoryIdObject(v Repository)`

SetRepositoryIdObject sets RepositoryIdObject field to given value.

### HasRepositoryIdObject

`func (o *Job) HasRepositoryIdObject() bool`

HasRepositoryIdObject returns a boolean if a field has been set.

### GetTags

`func (o *Job) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Job) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Job) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Job) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Job) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Job) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Job) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Job) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



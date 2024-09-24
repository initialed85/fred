/*
Djangolang

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_client

import (
	"encoding/json"
	"time"
)

// checks if the Change type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Change{}

// Change struct for Change
type Change struct {
	AuthoredAt                         *time.Time  `json:"authored_at,omitempty"`
	AuthoredBy                         *string     `json:"authored_by,omitempty"`
	BranchName                         *string     `json:"branch_name,omitempty"`
	CommitHash                         *string     `json:"commit_hash,omitempty"`
	CommittedAt                        *time.Time  `json:"committed_at,omitempty"`
	CommittedBy                        *string     `json:"committed_by,omitempty"`
	CreatedAt                          *time.Time  `json:"created_at,omitempty"`
	DeletedAt                          *time.Time  `json:"deleted_at,omitempty"`
	Id                                 *string     `json:"id,omitempty"`
	Message                            *string     `json:"message,omitempty"`
	ReferencedByTriggerChangeIdObjects []Trigger   `json:"referenced_by_trigger_change_id_objects,omitempty"`
	RepositoryId                       *string     `json:"repository_id,omitempty"`
	RepositoryIdObject                 *Repository `json:"repository_id_object,omitempty"`
	TriggerProducedAt                  *time.Time  `json:"trigger_produced_at,omitempty"`
	UpdatedAt                          *time.Time  `json:"updated_at,omitempty"`
}

// NewChange instantiates a new Change object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChange() *Change {
	this := Change{}
	return &this
}

// NewChangeWithDefaults instantiates a new Change object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeWithDefaults() *Change {
	this := Change{}
	return &this
}

// GetAuthoredAt returns the AuthoredAt field value if set, zero value otherwise.
func (o *Change) GetAuthoredAt() time.Time {
	if o == nil || IsNil(o.AuthoredAt) {
		var ret time.Time
		return ret
	}
	return *o.AuthoredAt
}

// GetAuthoredAtOk returns a tuple with the AuthoredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetAuthoredAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AuthoredAt) {
		return nil, false
	}
	return o.AuthoredAt, true
}

// HasAuthoredAt returns a boolean if a field has been set.
func (o *Change) HasAuthoredAt() bool {
	if o != nil && !IsNil(o.AuthoredAt) {
		return true
	}

	return false
}

// SetAuthoredAt gets a reference to the given time.Time and assigns it to the AuthoredAt field.
func (o *Change) SetAuthoredAt(v time.Time) {
	o.AuthoredAt = &v
}

// GetAuthoredBy returns the AuthoredBy field value if set, zero value otherwise.
func (o *Change) GetAuthoredBy() string {
	if o == nil || IsNil(o.AuthoredBy) {
		var ret string
		return ret
	}
	return *o.AuthoredBy
}

// GetAuthoredByOk returns a tuple with the AuthoredBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetAuthoredByOk() (*string, bool) {
	if o == nil || IsNil(o.AuthoredBy) {
		return nil, false
	}
	return o.AuthoredBy, true
}

// HasAuthoredBy returns a boolean if a field has been set.
func (o *Change) HasAuthoredBy() bool {
	if o != nil && !IsNil(o.AuthoredBy) {
		return true
	}

	return false
}

// SetAuthoredBy gets a reference to the given string and assigns it to the AuthoredBy field.
func (o *Change) SetAuthoredBy(v string) {
	o.AuthoredBy = &v
}

// GetBranchName returns the BranchName field value if set, zero value otherwise.
func (o *Change) GetBranchName() string {
	if o == nil || IsNil(o.BranchName) {
		var ret string
		return ret
	}
	return *o.BranchName
}

// GetBranchNameOk returns a tuple with the BranchName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetBranchNameOk() (*string, bool) {
	if o == nil || IsNil(o.BranchName) {
		return nil, false
	}
	return o.BranchName, true
}

// HasBranchName returns a boolean if a field has been set.
func (o *Change) HasBranchName() bool {
	if o != nil && !IsNil(o.BranchName) {
		return true
	}

	return false
}

// SetBranchName gets a reference to the given string and assigns it to the BranchName field.
func (o *Change) SetBranchName(v string) {
	o.BranchName = &v
}

// GetCommitHash returns the CommitHash field value if set, zero value otherwise.
func (o *Change) GetCommitHash() string {
	if o == nil || IsNil(o.CommitHash) {
		var ret string
		return ret
	}
	return *o.CommitHash
}

// GetCommitHashOk returns a tuple with the CommitHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetCommitHashOk() (*string, bool) {
	if o == nil || IsNil(o.CommitHash) {
		return nil, false
	}
	return o.CommitHash, true
}

// HasCommitHash returns a boolean if a field has been set.
func (o *Change) HasCommitHash() bool {
	if o != nil && !IsNil(o.CommitHash) {
		return true
	}

	return false
}

// SetCommitHash gets a reference to the given string and assigns it to the CommitHash field.
func (o *Change) SetCommitHash(v string) {
	o.CommitHash = &v
}

// GetCommittedAt returns the CommittedAt field value if set, zero value otherwise.
func (o *Change) GetCommittedAt() time.Time {
	if o == nil || IsNil(o.CommittedAt) {
		var ret time.Time
		return ret
	}
	return *o.CommittedAt
}

// GetCommittedAtOk returns a tuple with the CommittedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetCommittedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CommittedAt) {
		return nil, false
	}
	return o.CommittedAt, true
}

// HasCommittedAt returns a boolean if a field has been set.
func (o *Change) HasCommittedAt() bool {
	if o != nil && !IsNil(o.CommittedAt) {
		return true
	}

	return false
}

// SetCommittedAt gets a reference to the given time.Time and assigns it to the CommittedAt field.
func (o *Change) SetCommittedAt(v time.Time) {
	o.CommittedAt = &v
}

// GetCommittedBy returns the CommittedBy field value if set, zero value otherwise.
func (o *Change) GetCommittedBy() string {
	if o == nil || IsNil(o.CommittedBy) {
		var ret string
		return ret
	}
	return *o.CommittedBy
}

// GetCommittedByOk returns a tuple with the CommittedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetCommittedByOk() (*string, bool) {
	if o == nil || IsNil(o.CommittedBy) {
		return nil, false
	}
	return o.CommittedBy, true
}

// HasCommittedBy returns a boolean if a field has been set.
func (o *Change) HasCommittedBy() bool {
	if o != nil && !IsNil(o.CommittedBy) {
		return true
	}

	return false
}

// SetCommittedBy gets a reference to the given string and assigns it to the CommittedBy field.
func (o *Change) SetCommittedBy(v string) {
	o.CommittedBy = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Change) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Change) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Change) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *Change) GetDeletedAt() time.Time {
	if o == nil || IsNil(o.DeletedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *Change) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *Change) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Change) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Change) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Change) SetId(v string) {
	o.Id = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Change) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Change) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Change) SetMessage(v string) {
	o.Message = &v
}

// GetReferencedByTriggerChangeIdObjects returns the ReferencedByTriggerChangeIdObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Change) GetReferencedByTriggerChangeIdObjects() []Trigger {
	if o == nil {
		var ret []Trigger
		return ret
	}
	return o.ReferencedByTriggerChangeIdObjects
}

// GetReferencedByTriggerChangeIdObjectsOk returns a tuple with the ReferencedByTriggerChangeIdObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Change) GetReferencedByTriggerChangeIdObjectsOk() ([]Trigger, bool) {
	if o == nil || IsNil(o.ReferencedByTriggerChangeIdObjects) {
		return nil, false
	}
	return o.ReferencedByTriggerChangeIdObjects, true
}

// HasReferencedByTriggerChangeIdObjects returns a boolean if a field has been set.
func (o *Change) HasReferencedByTriggerChangeIdObjects() bool {
	if o != nil && !IsNil(o.ReferencedByTriggerChangeIdObjects) {
		return true
	}

	return false
}

// SetReferencedByTriggerChangeIdObjects gets a reference to the given []Trigger and assigns it to the ReferencedByTriggerChangeIdObjects field.
func (o *Change) SetReferencedByTriggerChangeIdObjects(v []Trigger) {
	o.ReferencedByTriggerChangeIdObjects = v
}

// GetRepositoryId returns the RepositoryId field value if set, zero value otherwise.
func (o *Change) GetRepositoryId() string {
	if o == nil || IsNil(o.RepositoryId) {
		var ret string
		return ret
	}
	return *o.RepositoryId
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetRepositoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.RepositoryId) {
		return nil, false
	}
	return o.RepositoryId, true
}

// HasRepositoryId returns a boolean if a field has been set.
func (o *Change) HasRepositoryId() bool {
	if o != nil && !IsNil(o.RepositoryId) {
		return true
	}

	return false
}

// SetRepositoryId gets a reference to the given string and assigns it to the RepositoryId field.
func (o *Change) SetRepositoryId(v string) {
	o.RepositoryId = &v
}

// GetRepositoryIdObject returns the RepositoryIdObject field value if set, zero value otherwise.
func (o *Change) GetRepositoryIdObject() Repository {
	if o == nil || IsNil(o.RepositoryIdObject) {
		var ret Repository
		return ret
	}
	return *o.RepositoryIdObject
}

// GetRepositoryIdObjectOk returns a tuple with the RepositoryIdObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetRepositoryIdObjectOk() (*Repository, bool) {
	if o == nil || IsNil(o.RepositoryIdObject) {
		return nil, false
	}
	return o.RepositoryIdObject, true
}

// HasRepositoryIdObject returns a boolean if a field has been set.
func (o *Change) HasRepositoryIdObject() bool {
	if o != nil && !IsNil(o.RepositoryIdObject) {
		return true
	}

	return false
}

// SetRepositoryIdObject gets a reference to the given Repository and assigns it to the RepositoryIdObject field.
func (o *Change) SetRepositoryIdObject(v Repository) {
	o.RepositoryIdObject = &v
}

// GetTriggerProducedAt returns the TriggerProducedAt field value if set, zero value otherwise.
func (o *Change) GetTriggerProducedAt() time.Time {
	if o == nil || IsNil(o.TriggerProducedAt) {
		var ret time.Time
		return ret
	}
	return *o.TriggerProducedAt
}

// GetTriggerProducedAtOk returns a tuple with the TriggerProducedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetTriggerProducedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TriggerProducedAt) {
		return nil, false
	}
	return o.TriggerProducedAt, true
}

// HasTriggerProducedAt returns a boolean if a field has been set.
func (o *Change) HasTriggerProducedAt() bool {
	if o != nil && !IsNil(o.TriggerProducedAt) {
		return true
	}

	return false
}

// SetTriggerProducedAt gets a reference to the given time.Time and assigns it to the TriggerProducedAt field.
func (o *Change) SetTriggerProducedAt(v time.Time) {
	o.TriggerProducedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Change) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Change) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Change) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o Change) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Change) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthoredAt) {
		toSerialize["authored_at"] = o.AuthoredAt
	}
	if !IsNil(o.AuthoredBy) {
		toSerialize["authored_by"] = o.AuthoredBy
	}
	if !IsNil(o.BranchName) {
		toSerialize["branch_name"] = o.BranchName
	}
	if !IsNil(o.CommitHash) {
		toSerialize["commit_hash"] = o.CommitHash
	}
	if !IsNil(o.CommittedAt) {
		toSerialize["committed_at"] = o.CommittedAt
	}
	if !IsNil(o.CommittedBy) {
		toSerialize["committed_by"] = o.CommittedBy
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if o.ReferencedByTriggerChangeIdObjects != nil {
		toSerialize["referenced_by_trigger_change_id_objects"] = o.ReferencedByTriggerChangeIdObjects
	}
	if !IsNil(o.RepositoryId) {
		toSerialize["repository_id"] = o.RepositoryId
	}
	if !IsNil(o.RepositoryIdObject) {
		toSerialize["repository_id_object"] = o.RepositoryIdObject
	}
	if !IsNil(o.TriggerProducedAt) {
		toSerialize["trigger_produced_at"] = o.TriggerProducedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableChange struct {
	value *Change
	isSet bool
}

func (v NullableChange) Get() *Change {
	return v.value
}

func (v *NullableChange) Set(val *Change) {
	v.value = val
	v.isSet = true
}

func (v NullableChange) IsSet() bool {
	return v.isSet
}

func (v *NullableChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChange(val *Change) *NullableChange {
	return &NullableChange{value: val, isSet: true}
}

func (v NullableChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

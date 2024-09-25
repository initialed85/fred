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

// checks if the TriggerHasExecution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TriggerHasExecution{}

// TriggerHasExecution struct for TriggerHasExecution
type TriggerHasExecution struct {
	CreatedAt         *time.Time `json:"created_at,omitempty"`
	DeletedAt         *time.Time `json:"deleted_at,omitempty"`
	ExecutionId       *string    `json:"execution_id,omitempty"`
	ExecutionIdObject *Execution `json:"execution_id_object,omitempty"`
	Id                *string    `json:"id,omitempty"`
	TriggerId         *string    `json:"trigger_id,omitempty"`
	TriggerIdObject   *Trigger   `json:"trigger_id_object,omitempty"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
}

// NewTriggerHasExecution instantiates a new TriggerHasExecution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerHasExecution() *TriggerHasExecution {
	this := TriggerHasExecution{}
	return &this
}

// NewTriggerHasExecutionWithDefaults instantiates a new TriggerHasExecution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerHasExecutionWithDefaults() *TriggerHasExecution {
	this := TriggerHasExecution{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TriggerHasExecution) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerHasExecution) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TriggerHasExecution) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *TriggerHasExecution) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *TriggerHasExecution) GetDeletedAt() time.Time {
	if o == nil || IsNil(o.DeletedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerHasExecution) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *TriggerHasExecution) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *TriggerHasExecution) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise.
func (o *TriggerHasExecution) GetExecutionId() string {
	if o == nil || IsNil(o.ExecutionId) {
		var ret string
		return ret
	}
	return *o.ExecutionId
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerHasExecution) GetExecutionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutionId) {
		return nil, false
	}
	return o.ExecutionId, true
}

// HasExecutionId returns a boolean if a field has been set.
func (o *TriggerHasExecution) HasExecutionId() bool {
	if o != nil && !IsNil(o.ExecutionId) {
		return true
	}

	return false
}

// SetExecutionId gets a reference to the given string and assigns it to the ExecutionId field.
func (o *TriggerHasExecution) SetExecutionId(v string) {
	o.ExecutionId = &v
}

// GetExecutionIdObject returns the ExecutionIdObject field value if set, zero value otherwise.
func (o *TriggerHasExecution) GetExecutionIdObject() Execution {
	if o == nil || IsNil(o.ExecutionIdObject) {
		var ret Execution
		return ret
	}
	return *o.ExecutionIdObject
}

// GetExecutionIdObjectOk returns a tuple with the ExecutionIdObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerHasExecution) GetExecutionIdObjectOk() (*Execution, bool) {
	if o == nil || IsNil(o.ExecutionIdObject) {
		return nil, false
	}
	return o.ExecutionIdObject, true
}

// HasExecutionIdObject returns a boolean if a field has been set.
func (o *TriggerHasExecution) HasExecutionIdObject() bool {
	if o != nil && !IsNil(o.ExecutionIdObject) {
		return true
	}

	return false
}

// SetExecutionIdObject gets a reference to the given Execution and assigns it to the ExecutionIdObject field.
func (o *TriggerHasExecution) SetExecutionIdObject(v Execution) {
	o.ExecutionIdObject = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TriggerHasExecution) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerHasExecution) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TriggerHasExecution) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TriggerHasExecution) SetId(v string) {
	o.Id = &v
}

// GetTriggerId returns the TriggerId field value if set, zero value otherwise.
func (o *TriggerHasExecution) GetTriggerId() string {
	if o == nil || IsNil(o.TriggerId) {
		var ret string
		return ret
	}
	return *o.TriggerId
}

// GetTriggerIdOk returns a tuple with the TriggerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerHasExecution) GetTriggerIdOk() (*string, bool) {
	if o == nil || IsNil(o.TriggerId) {
		return nil, false
	}
	return o.TriggerId, true
}

// HasTriggerId returns a boolean if a field has been set.
func (o *TriggerHasExecution) HasTriggerId() bool {
	if o != nil && !IsNil(o.TriggerId) {
		return true
	}

	return false
}

// SetTriggerId gets a reference to the given string and assigns it to the TriggerId field.
func (o *TriggerHasExecution) SetTriggerId(v string) {
	o.TriggerId = &v
}

// GetTriggerIdObject returns the TriggerIdObject field value if set, zero value otherwise.
func (o *TriggerHasExecution) GetTriggerIdObject() Trigger {
	if o == nil || IsNil(o.TriggerIdObject) {
		var ret Trigger
		return ret
	}
	return *o.TriggerIdObject
}

// GetTriggerIdObjectOk returns a tuple with the TriggerIdObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerHasExecution) GetTriggerIdObjectOk() (*Trigger, bool) {
	if o == nil || IsNil(o.TriggerIdObject) {
		return nil, false
	}
	return o.TriggerIdObject, true
}

// HasTriggerIdObject returns a boolean if a field has been set.
func (o *TriggerHasExecution) HasTriggerIdObject() bool {
	if o != nil && !IsNil(o.TriggerIdObject) {
		return true
	}

	return false
}

// SetTriggerIdObject gets a reference to the given Trigger and assigns it to the TriggerIdObject field.
func (o *TriggerHasExecution) SetTriggerIdObject(v Trigger) {
	o.TriggerIdObject = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TriggerHasExecution) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerHasExecution) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TriggerHasExecution) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *TriggerHasExecution) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o TriggerHasExecution) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TriggerHasExecution) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if !IsNil(o.ExecutionId) {
		toSerialize["execution_id"] = o.ExecutionId
	}
	if !IsNil(o.ExecutionIdObject) {
		toSerialize["execution_id_object"] = o.ExecutionIdObject
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.TriggerId) {
		toSerialize["trigger_id"] = o.TriggerId
	}
	if !IsNil(o.TriggerIdObject) {
		toSerialize["trigger_id_object"] = o.TriggerIdObject
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableTriggerHasExecution struct {
	value *TriggerHasExecution
	isSet bool
}

func (v NullableTriggerHasExecution) Get() *TriggerHasExecution {
	return v.value
}

func (v *NullableTriggerHasExecution) Set(val *TriggerHasExecution) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerHasExecution) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerHasExecution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerHasExecution(val *TriggerHasExecution) *NullableTriggerHasExecution {
	return &NullableTriggerHasExecution{value: val, isSet: true}
}

func (v NullableTriggerHasExecution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerHasExecution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

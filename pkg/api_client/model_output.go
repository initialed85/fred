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

// checks if the Output type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Output{}

// Output struct for Output
type Output struct {
	CreatedAt                                    *time.Time  `json:"created_at,omitempty"`
	DeletedAt                                    *time.Time  `json:"deleted_at,omitempty"`
	Id                                           *string     `json:"id,omitempty"`
	ReferencedByExecutionBuildOutputIdObjects    []Execution `json:"referenced_by_execution_build_output_id_objects,omitempty"`
	ReferencedByExecutionDeployOutputIdObjects   []Execution `json:"referenced_by_execution_deploy_output_id_objects,omitempty"`
	ReferencedByExecutionPublishOutputIdObjects  []Execution `json:"referenced_by_execution_publish_output_id_objects,omitempty"`
	ReferencedByExecutionTestOutputIdObjects     []Execution `json:"referenced_by_execution_test_output_id_objects,omitempty"`
	ReferencedByExecutionValidateOutputIdObjects []Execution `json:"referenced_by_execution_validate_output_id_objects,omitempty"`
	TaskId                                       *string     `json:"task_id,omitempty"`
	TaskIdObject                                 *Task       `json:"task_id_object,omitempty"`
	UpdatedAt                                    *time.Time  `json:"updated_at,omitempty"`
}

// NewOutput instantiates a new Output object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutput() *Output {
	this := Output{}
	return &this
}

// NewOutputWithDefaults instantiates a new Output object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputWithDefaults() *Output {
	this := Output{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Output) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Output) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Output) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Output) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *Output) GetDeletedAt() time.Time {
	if o == nil || IsNil(o.DeletedAt) {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Output) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeletedAt) {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *Output) HasDeletedAt() bool {
	if o != nil && !IsNil(o.DeletedAt) {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *Output) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Output) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Output) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Output) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Output) SetId(v string) {
	o.Id = &v
}

// GetReferencedByExecutionBuildOutputIdObjects returns the ReferencedByExecutionBuildOutputIdObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Output) GetReferencedByExecutionBuildOutputIdObjects() []Execution {
	if o == nil {
		var ret []Execution
		return ret
	}
	return o.ReferencedByExecutionBuildOutputIdObjects
}

// GetReferencedByExecutionBuildOutputIdObjectsOk returns a tuple with the ReferencedByExecutionBuildOutputIdObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Output) GetReferencedByExecutionBuildOutputIdObjectsOk() ([]Execution, bool) {
	if o == nil || IsNil(o.ReferencedByExecutionBuildOutputIdObjects) {
		return nil, false
	}
	return o.ReferencedByExecutionBuildOutputIdObjects, true
}

// HasReferencedByExecutionBuildOutputIdObjects returns a boolean if a field has been set.
func (o *Output) HasReferencedByExecutionBuildOutputIdObjects() bool {
	if o != nil && !IsNil(o.ReferencedByExecutionBuildOutputIdObjects) {
		return true
	}

	return false
}

// SetReferencedByExecutionBuildOutputIdObjects gets a reference to the given []Execution and assigns it to the ReferencedByExecutionBuildOutputIdObjects field.
func (o *Output) SetReferencedByExecutionBuildOutputIdObjects(v []Execution) {
	o.ReferencedByExecutionBuildOutputIdObjects = v
}

// GetReferencedByExecutionDeployOutputIdObjects returns the ReferencedByExecutionDeployOutputIdObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Output) GetReferencedByExecutionDeployOutputIdObjects() []Execution {
	if o == nil {
		var ret []Execution
		return ret
	}
	return o.ReferencedByExecutionDeployOutputIdObjects
}

// GetReferencedByExecutionDeployOutputIdObjectsOk returns a tuple with the ReferencedByExecutionDeployOutputIdObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Output) GetReferencedByExecutionDeployOutputIdObjectsOk() ([]Execution, bool) {
	if o == nil || IsNil(o.ReferencedByExecutionDeployOutputIdObjects) {
		return nil, false
	}
	return o.ReferencedByExecutionDeployOutputIdObjects, true
}

// HasReferencedByExecutionDeployOutputIdObjects returns a boolean if a field has been set.
func (o *Output) HasReferencedByExecutionDeployOutputIdObjects() bool {
	if o != nil && !IsNil(o.ReferencedByExecutionDeployOutputIdObjects) {
		return true
	}

	return false
}

// SetReferencedByExecutionDeployOutputIdObjects gets a reference to the given []Execution and assigns it to the ReferencedByExecutionDeployOutputIdObjects field.
func (o *Output) SetReferencedByExecutionDeployOutputIdObjects(v []Execution) {
	o.ReferencedByExecutionDeployOutputIdObjects = v
}

// GetReferencedByExecutionPublishOutputIdObjects returns the ReferencedByExecutionPublishOutputIdObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Output) GetReferencedByExecutionPublishOutputIdObjects() []Execution {
	if o == nil {
		var ret []Execution
		return ret
	}
	return o.ReferencedByExecutionPublishOutputIdObjects
}

// GetReferencedByExecutionPublishOutputIdObjectsOk returns a tuple with the ReferencedByExecutionPublishOutputIdObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Output) GetReferencedByExecutionPublishOutputIdObjectsOk() ([]Execution, bool) {
	if o == nil || IsNil(o.ReferencedByExecutionPublishOutputIdObjects) {
		return nil, false
	}
	return o.ReferencedByExecutionPublishOutputIdObjects, true
}

// HasReferencedByExecutionPublishOutputIdObjects returns a boolean if a field has been set.
func (o *Output) HasReferencedByExecutionPublishOutputIdObjects() bool {
	if o != nil && !IsNil(o.ReferencedByExecutionPublishOutputIdObjects) {
		return true
	}

	return false
}

// SetReferencedByExecutionPublishOutputIdObjects gets a reference to the given []Execution and assigns it to the ReferencedByExecutionPublishOutputIdObjects field.
func (o *Output) SetReferencedByExecutionPublishOutputIdObjects(v []Execution) {
	o.ReferencedByExecutionPublishOutputIdObjects = v
}

// GetReferencedByExecutionTestOutputIdObjects returns the ReferencedByExecutionTestOutputIdObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Output) GetReferencedByExecutionTestOutputIdObjects() []Execution {
	if o == nil {
		var ret []Execution
		return ret
	}
	return o.ReferencedByExecutionTestOutputIdObjects
}

// GetReferencedByExecutionTestOutputIdObjectsOk returns a tuple with the ReferencedByExecutionTestOutputIdObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Output) GetReferencedByExecutionTestOutputIdObjectsOk() ([]Execution, bool) {
	if o == nil || IsNil(o.ReferencedByExecutionTestOutputIdObjects) {
		return nil, false
	}
	return o.ReferencedByExecutionTestOutputIdObjects, true
}

// HasReferencedByExecutionTestOutputIdObjects returns a boolean if a field has been set.
func (o *Output) HasReferencedByExecutionTestOutputIdObjects() bool {
	if o != nil && !IsNil(o.ReferencedByExecutionTestOutputIdObjects) {
		return true
	}

	return false
}

// SetReferencedByExecutionTestOutputIdObjects gets a reference to the given []Execution and assigns it to the ReferencedByExecutionTestOutputIdObjects field.
func (o *Output) SetReferencedByExecutionTestOutputIdObjects(v []Execution) {
	o.ReferencedByExecutionTestOutputIdObjects = v
}

// GetReferencedByExecutionValidateOutputIdObjects returns the ReferencedByExecutionValidateOutputIdObjects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Output) GetReferencedByExecutionValidateOutputIdObjects() []Execution {
	if o == nil {
		var ret []Execution
		return ret
	}
	return o.ReferencedByExecutionValidateOutputIdObjects
}

// GetReferencedByExecutionValidateOutputIdObjectsOk returns a tuple with the ReferencedByExecutionValidateOutputIdObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Output) GetReferencedByExecutionValidateOutputIdObjectsOk() ([]Execution, bool) {
	if o == nil || IsNil(o.ReferencedByExecutionValidateOutputIdObjects) {
		return nil, false
	}
	return o.ReferencedByExecutionValidateOutputIdObjects, true
}

// HasReferencedByExecutionValidateOutputIdObjects returns a boolean if a field has been set.
func (o *Output) HasReferencedByExecutionValidateOutputIdObjects() bool {
	if o != nil && !IsNil(o.ReferencedByExecutionValidateOutputIdObjects) {
		return true
	}

	return false
}

// SetReferencedByExecutionValidateOutputIdObjects gets a reference to the given []Execution and assigns it to the ReferencedByExecutionValidateOutputIdObjects field.
func (o *Output) SetReferencedByExecutionValidateOutputIdObjects(v []Execution) {
	o.ReferencedByExecutionValidateOutputIdObjects = v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *Output) GetTaskId() string {
	if o == nil || IsNil(o.TaskId) {
		var ret string
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Output) GetTaskIdOk() (*string, bool) {
	if o == nil || IsNil(o.TaskId) {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *Output) HasTaskId() bool {
	if o != nil && !IsNil(o.TaskId) {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given string and assigns it to the TaskId field.
func (o *Output) SetTaskId(v string) {
	o.TaskId = &v
}

// GetTaskIdObject returns the TaskIdObject field value if set, zero value otherwise.
func (o *Output) GetTaskIdObject() Task {
	if o == nil || IsNil(o.TaskIdObject) {
		var ret Task
		return ret
	}
	return *o.TaskIdObject
}

// GetTaskIdObjectOk returns a tuple with the TaskIdObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Output) GetTaskIdObjectOk() (*Task, bool) {
	if o == nil || IsNil(o.TaskIdObject) {
		return nil, false
	}
	return o.TaskIdObject, true
}

// HasTaskIdObject returns a boolean if a field has been set.
func (o *Output) HasTaskIdObject() bool {
	if o != nil && !IsNil(o.TaskIdObject) {
		return true
	}

	return false
}

// SetTaskIdObject gets a reference to the given Task and assigns it to the TaskIdObject field.
func (o *Output) SetTaskIdObject(v Task) {
	o.TaskIdObject = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Output) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Output) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Output) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Output) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o Output) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Output) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DeletedAt) {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.ReferencedByExecutionBuildOutputIdObjects != nil {
		toSerialize["referenced_by_execution_build_output_id_objects"] = o.ReferencedByExecutionBuildOutputIdObjects
	}
	if o.ReferencedByExecutionDeployOutputIdObjects != nil {
		toSerialize["referenced_by_execution_deploy_output_id_objects"] = o.ReferencedByExecutionDeployOutputIdObjects
	}
	if o.ReferencedByExecutionPublishOutputIdObjects != nil {
		toSerialize["referenced_by_execution_publish_output_id_objects"] = o.ReferencedByExecutionPublishOutputIdObjects
	}
	if o.ReferencedByExecutionTestOutputIdObjects != nil {
		toSerialize["referenced_by_execution_test_output_id_objects"] = o.ReferencedByExecutionTestOutputIdObjects
	}
	if o.ReferencedByExecutionValidateOutputIdObjects != nil {
		toSerialize["referenced_by_execution_validate_output_id_objects"] = o.ReferencedByExecutionValidateOutputIdObjects
	}
	if !IsNil(o.TaskId) {
		toSerialize["task_id"] = o.TaskId
	}
	if !IsNil(o.TaskIdObject) {
		toSerialize["task_id_object"] = o.TaskIdObject
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableOutput struct {
	value *Output
	isSet bool
}

func (v NullableOutput) Get() *Output {
	return v.value
}

func (v *NullableOutput) Set(val *Output) {
	v.value = val
	v.isSet = true
}

func (v NullableOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutput(val *Output) *NullableOutput {
	return &NullableOutput{value: val, isSet: true}
}

func (v NullableOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

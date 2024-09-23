/*
Djangolang

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api_client

import (
	"encoding/json"
)

// checks if the ResponseWithGenericOfTrigger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseWithGenericOfTrigger{}

// ResponseWithGenericOfTrigger struct for ResponseWithGenericOfTrigger
type ResponseWithGenericOfTrigger struct {
	Count *int64 `json:"count,omitempty"`
	Error []string `json:"error,omitempty"`
	Limit *int64 `json:"limit,omitempty"`
	Objects []Trigger `json:"objects,omitempty"`
	Offset *int64 `json:"offset,omitempty"`
	Status *int64 `json:"status,omitempty"`
	Success *bool `json:"success,omitempty"`
	TotalCount *int64 `json:"total_count,omitempty"`
}

// NewResponseWithGenericOfTrigger instantiates a new ResponseWithGenericOfTrigger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseWithGenericOfTrigger() *ResponseWithGenericOfTrigger {
	this := ResponseWithGenericOfTrigger{}
	return &this
}

// NewResponseWithGenericOfTriggerWithDefaults instantiates a new ResponseWithGenericOfTrigger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseWithGenericOfTriggerWithDefaults() *ResponseWithGenericOfTrigger {
	this := ResponseWithGenericOfTrigger{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ResponseWithGenericOfTrigger) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseWithGenericOfTrigger) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ResponseWithGenericOfTrigger) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *ResponseWithGenericOfTrigger) SetCount(v int64) {
	o.Count = &v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponseWithGenericOfTrigger) GetError() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseWithGenericOfTrigger) GetErrorOk() ([]string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ResponseWithGenericOfTrigger) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given []string and assigns it to the Error field.
func (o *ResponseWithGenericOfTrigger) SetError(v []string) {
	o.Error = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ResponseWithGenericOfTrigger) GetLimit() int64 {
	if o == nil || IsNil(o.Limit) {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseWithGenericOfTrigger) GetLimitOk() (*int64, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ResponseWithGenericOfTrigger) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *ResponseWithGenericOfTrigger) SetLimit(v int64) {
	o.Limit = &v
}

// GetObjects returns the Objects field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponseWithGenericOfTrigger) GetObjects() []Trigger {
	if o == nil {
		var ret []Trigger
		return ret
	}
	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseWithGenericOfTrigger) GetObjectsOk() ([]Trigger, bool) {
	if o == nil || IsNil(o.Objects) {
		return nil, false
	}
	return o.Objects, true
}

// HasObjects returns a boolean if a field has been set.
func (o *ResponseWithGenericOfTrigger) HasObjects() bool {
	if o != nil && !IsNil(o.Objects) {
		return true
	}

	return false
}

// SetObjects gets a reference to the given []Trigger and assigns it to the Objects field.
func (o *ResponseWithGenericOfTrigger) SetObjects(v []Trigger) {
	o.Objects = v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ResponseWithGenericOfTrigger) GetOffset() int64 {
	if o == nil || IsNil(o.Offset) {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseWithGenericOfTrigger) GetOffsetOk() (*int64, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ResponseWithGenericOfTrigger) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *ResponseWithGenericOfTrigger) SetOffset(v int64) {
	o.Offset = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ResponseWithGenericOfTrigger) GetStatus() int64 {
	if o == nil || IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseWithGenericOfTrigger) GetStatusOk() (*int64, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ResponseWithGenericOfTrigger) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *ResponseWithGenericOfTrigger) SetStatus(v int64) {
	o.Status = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ResponseWithGenericOfTrigger) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseWithGenericOfTrigger) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ResponseWithGenericOfTrigger) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ResponseWithGenericOfTrigger) SetSuccess(v bool) {
	o.Success = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ResponseWithGenericOfTrigger) GetTotalCount() int64 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseWithGenericOfTrigger) GetTotalCountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ResponseWithGenericOfTrigger) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *ResponseWithGenericOfTrigger) SetTotalCount(v int64) {
	o.TotalCount = &v
}

func (o ResponseWithGenericOfTrigger) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseWithGenericOfTrigger) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if o.Objects != nil {
		toSerialize["objects"] = o.Objects
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.TotalCount) {
		toSerialize["total_count"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableResponseWithGenericOfTrigger struct {
	value *ResponseWithGenericOfTrigger
	isSet bool
}

func (v NullableResponseWithGenericOfTrigger) Get() *ResponseWithGenericOfTrigger {
	return v.value
}

func (v *NullableResponseWithGenericOfTrigger) Set(val *ResponseWithGenericOfTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseWithGenericOfTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseWithGenericOfTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseWithGenericOfTrigger(val *ResponseWithGenericOfTrigger) *NullableResponseWithGenericOfTrigger {
	return &NullableResponseWithGenericOfTrigger{value: val, isSet: true}
}

func (v NullableResponseWithGenericOfTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseWithGenericOfTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



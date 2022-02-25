/*
Landscape omnikeeper REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package okclient

import (
	"encoding/json"
)

// ChangeDataCell struct for ChangeDataCell
type ChangeDataCell struct {
	Name NullableString `json:"name,omitempty"`
	Value *AttributeValueDTO `json:"value,omitempty"`
	Changeable *bool `json:"changeable,omitempty"`
}

// NewChangeDataCell instantiates a new ChangeDataCell object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeDataCell() *ChangeDataCell {
	this := ChangeDataCell{}
	return &this
}

// NewChangeDataCellWithDefaults instantiates a new ChangeDataCell object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeDataCellWithDefaults() *ChangeDataCell {
	this := ChangeDataCell{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChangeDataCell) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChangeDataCell) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ChangeDataCell) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ChangeDataCell) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ChangeDataCell) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ChangeDataCell) UnsetName() {
	o.Name.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ChangeDataCell) GetValue() AttributeValueDTO {
	if o == nil || o.Value == nil {
		var ret AttributeValueDTO
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeDataCell) GetValueOk() (*AttributeValueDTO, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ChangeDataCell) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given AttributeValueDTO and assigns it to the Value field.
func (o *ChangeDataCell) SetValue(v AttributeValueDTO) {
	o.Value = &v
}

// GetChangeable returns the Changeable field value if set, zero value otherwise.
func (o *ChangeDataCell) GetChangeable() bool {
	if o == nil || o.Changeable == nil {
		var ret bool
		return ret
	}
	return *o.Changeable
}

// GetChangeableOk returns a tuple with the Changeable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeDataCell) GetChangeableOk() (*bool, bool) {
	if o == nil || o.Changeable == nil {
		return nil, false
	}
	return o.Changeable, true
}

// HasChangeable returns a boolean if a field has been set.
func (o *ChangeDataCell) HasChangeable() bool {
	if o != nil && o.Changeable != nil {
		return true
	}

	return false
}

// SetChangeable gets a reference to the given bool and assigns it to the Changeable field.
func (o *ChangeDataCell) SetChangeable(v bool) {
	o.Changeable = &v
}

func (o ChangeDataCell) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Changeable != nil {
		toSerialize["changeable"] = o.Changeable
	}
	return json.Marshal(toSerialize)
}

type NullableChangeDataCell struct {
	value *ChangeDataCell
	isSet bool
}

func (v NullableChangeDataCell) Get() *ChangeDataCell {
	return v.value
}

func (v *NullableChangeDataCell) Set(val *ChangeDataCell) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeDataCell) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeDataCell) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeDataCell(val *ChangeDataCell) *NullableChangeDataCell {
	return &NullableChangeDataCell{value: val, isSet: true}
}

func (v NullableChangeDataCell) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeDataCell) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



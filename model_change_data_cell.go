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

// checks if the ChangeDataCell type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeDataCell{}

// ChangeDataCell struct for ChangeDataCell
type ChangeDataCell struct {
	Id *string `json:"id,omitempty"`
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

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChangeDataCell) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeDataCell) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChangeDataCell) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ChangeDataCell) SetId(v string) {
	o.Id = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ChangeDataCell) GetValue() AttributeValueDTO {
	if o == nil || isNil(o.Value) {
		var ret AttributeValueDTO
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeDataCell) GetValueOk() (*AttributeValueDTO, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ChangeDataCell) HasValue() bool {
	if o != nil && !isNil(o.Value) {
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
	if o == nil || isNil(o.Changeable) {
		var ret bool
		return ret
	}
	return *o.Changeable
}

// GetChangeableOk returns a tuple with the Changeable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeDataCell) GetChangeableOk() (*bool, bool) {
	if o == nil || isNil(o.Changeable) {
		return nil, false
	}
	return o.Changeable, true
}

// HasChangeable returns a boolean if a field has been set.
func (o *ChangeDataCell) HasChangeable() bool {
	if o != nil && !isNil(o.Changeable) {
		return true
	}

	return false
}

// SetChangeable gets a reference to the given bool and assigns it to the Changeable field.
func (o *ChangeDataCell) SetChangeable(v bool) {
	o.Changeable = &v
}

func (o ChangeDataCell) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeDataCell) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !isNil(o.Changeable) {
		toSerialize["changeable"] = o.Changeable
	}
	return toSerialize, nil
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



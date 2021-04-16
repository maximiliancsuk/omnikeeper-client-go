/*
 * Landscape omnikeeper REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package okclient

import (
	"encoding/json"
)

// GridViewConfiguration struct for GridViewConfiguration
type GridViewConfiguration struct {
	ShowCIIDColumn *bool `json:"showCIIDColumn,omitempty"`
	WriteLayer *int64 `json:"writeLayer,omitempty"`
	ReadLayerset []int64 `json:"readLayerset,omitempty"`
	Columns []GridViewColumn `json:"columns,omitempty"`
	Trait NullableString `json:"trait,omitempty"`
}

// NewGridViewConfiguration instantiates a new GridViewConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGridViewConfiguration() *GridViewConfiguration {
	this := GridViewConfiguration{}
	return &this
}

// NewGridViewConfigurationWithDefaults instantiates a new GridViewConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGridViewConfigurationWithDefaults() *GridViewConfiguration {
	this := GridViewConfiguration{}
	return &this
}

// GetShowCIIDColumn returns the ShowCIIDColumn field value if set, zero value otherwise.
func (o *GridViewConfiguration) GetShowCIIDColumn() bool {
	if o == nil || o.ShowCIIDColumn == nil {
		var ret bool
		return ret
	}
	return *o.ShowCIIDColumn
}

// GetShowCIIDColumnOk returns a tuple with the ShowCIIDColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridViewConfiguration) GetShowCIIDColumnOk() (*bool, bool) {
	if o == nil || o.ShowCIIDColumn == nil {
		return nil, false
	}
	return o.ShowCIIDColumn, true
}

// HasShowCIIDColumn returns a boolean if a field has been set.
func (o *GridViewConfiguration) HasShowCIIDColumn() bool {
	if o != nil && o.ShowCIIDColumn != nil {
		return true
	}

	return false
}

// SetShowCIIDColumn gets a reference to the given bool and assigns it to the ShowCIIDColumn field.
func (o *GridViewConfiguration) SetShowCIIDColumn(v bool) {
	o.ShowCIIDColumn = &v
}

// GetWriteLayer returns the WriteLayer field value if set, zero value otherwise.
func (o *GridViewConfiguration) GetWriteLayer() int64 {
	if o == nil || o.WriteLayer == nil {
		var ret int64
		return ret
	}
	return *o.WriteLayer
}

// GetWriteLayerOk returns a tuple with the WriteLayer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridViewConfiguration) GetWriteLayerOk() (*int64, bool) {
	if o == nil || o.WriteLayer == nil {
		return nil, false
	}
	return o.WriteLayer, true
}

// HasWriteLayer returns a boolean if a field has been set.
func (o *GridViewConfiguration) HasWriteLayer() bool {
	if o != nil && o.WriteLayer != nil {
		return true
	}

	return false
}

// SetWriteLayer gets a reference to the given int64 and assigns it to the WriteLayer field.
func (o *GridViewConfiguration) SetWriteLayer(v int64) {
	o.WriteLayer = &v
}

// GetReadLayerset returns the ReadLayerset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GridViewConfiguration) GetReadLayerset() []int64 {
	if o == nil  {
		var ret []int64
		return ret
	}
	return o.ReadLayerset
}

// GetReadLayersetOk returns a tuple with the ReadLayerset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GridViewConfiguration) GetReadLayersetOk() (*[]int64, bool) {
	if o == nil || o.ReadLayerset == nil {
		return nil, false
	}
	return &o.ReadLayerset, true
}

// HasReadLayerset returns a boolean if a field has been set.
func (o *GridViewConfiguration) HasReadLayerset() bool {
	if o != nil && o.ReadLayerset != nil {
		return true
	}

	return false
}

// SetReadLayerset gets a reference to the given []int64 and assigns it to the ReadLayerset field.
func (o *GridViewConfiguration) SetReadLayerset(v []int64) {
	o.ReadLayerset = v
}

// GetColumns returns the Columns field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GridViewConfiguration) GetColumns() []GridViewColumn {
	if o == nil  {
		var ret []GridViewColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GridViewConfiguration) GetColumnsOk() (*[]GridViewColumn, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *GridViewConfiguration) HasColumns() bool {
	if o != nil && o.Columns != nil {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []GridViewColumn and assigns it to the Columns field.
func (o *GridViewConfiguration) SetColumns(v []GridViewColumn) {
	o.Columns = v
}

// GetTrait returns the Trait field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GridViewConfiguration) GetTrait() string {
	if o == nil || o.Trait.Get() == nil {
		var ret string
		return ret
	}
	return *o.Trait.Get()
}

// GetTraitOk returns a tuple with the Trait field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GridViewConfiguration) GetTraitOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Trait.Get(), o.Trait.IsSet()
}

// HasTrait returns a boolean if a field has been set.
func (o *GridViewConfiguration) HasTrait() bool {
	if o != nil && o.Trait.IsSet() {
		return true
	}

	return false
}

// SetTrait gets a reference to the given NullableString and assigns it to the Trait field.
func (o *GridViewConfiguration) SetTrait(v string) {
	o.Trait.Set(&v)
}
// SetTraitNil sets the value for Trait to be an explicit nil
func (o *GridViewConfiguration) SetTraitNil() {
	o.Trait.Set(nil)
}

// UnsetTrait ensures that no value is present for Trait, not even an explicit nil
func (o *GridViewConfiguration) UnsetTrait() {
	o.Trait.Unset()
}

func (o GridViewConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ShowCIIDColumn != nil {
		toSerialize["showCIIDColumn"] = o.ShowCIIDColumn
	}
	if o.WriteLayer != nil {
		toSerialize["writeLayer"] = o.WriteLayer
	}
	if o.ReadLayerset != nil {
		toSerialize["readLayerset"] = o.ReadLayerset
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.Trait.IsSet() {
		toSerialize["trait"] = o.Trait.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableGridViewConfiguration struct {
	value *GridViewConfiguration
	isSet bool
}

func (v NullableGridViewConfiguration) Get() *GridViewConfiguration {
	return v.value
}

func (v *NullableGridViewConfiguration) Set(val *GridViewConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableGridViewConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableGridViewConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGridViewConfiguration(val *GridViewConfiguration) *NullableGridViewConfiguration {
	return &NullableGridViewConfiguration{value: val, isSet: true}
}

func (v NullableGridViewConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGridViewConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


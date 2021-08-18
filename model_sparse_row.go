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

// SparseRow struct for SparseRow
type SparseRow struct {
	Ciid *string `json:"ciid,omitempty"`
	Cells []ChangeDataCell `json:"cells,omitempty"`
}

// NewSparseRow instantiates a new SparseRow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSparseRow() *SparseRow {
	this := SparseRow{}
	return &this
}

// NewSparseRowWithDefaults instantiates a new SparseRow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSparseRowWithDefaults() *SparseRow {
	this := SparseRow{}
	return &this
}

// GetCiid returns the Ciid field value if set, zero value otherwise.
func (o *SparseRow) GetCiid() string {
	if o == nil || o.Ciid == nil {
		var ret string
		return ret
	}
	return *o.Ciid
}

// GetCiidOk returns a tuple with the Ciid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SparseRow) GetCiidOk() (*string, bool) {
	if o == nil || o.Ciid == nil {
		return nil, false
	}
	return o.Ciid, true
}

// HasCiid returns a boolean if a field has been set.
func (o *SparseRow) HasCiid() bool {
	if o != nil && o.Ciid != nil {
		return true
	}

	return false
}

// SetCiid gets a reference to the given string and assigns it to the Ciid field.
func (o *SparseRow) SetCiid(v string) {
	o.Ciid = &v
}

// GetCells returns the Cells field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SparseRow) GetCells() []ChangeDataCell {
	if o == nil  {
		var ret []ChangeDataCell
		return ret
	}
	return o.Cells
}

// GetCellsOk returns a tuple with the Cells field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SparseRow) GetCellsOk() (*[]ChangeDataCell, bool) {
	if o == nil || o.Cells == nil {
		return nil, false
	}
	return &o.Cells, true
}

// HasCells returns a boolean if a field has been set.
func (o *SparseRow) HasCells() bool {
	if o != nil && o.Cells != nil {
		return true
	}

	return false
}

// SetCells gets a reference to the given []ChangeDataCell and assigns it to the Cells field.
func (o *SparseRow) SetCells(v []ChangeDataCell) {
	o.Cells = v
}

func (o SparseRow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ciid != nil {
		toSerialize["ciid"] = o.Ciid
	}
	if o.Cells != nil {
		toSerialize["cells"] = o.Cells
	}
	return json.Marshal(toSerialize)
}

type NullableSparseRow struct {
	value *SparseRow
	isSet bool
}

func (v NullableSparseRow) Get() *SparseRow {
	return v.value
}

func (v *NullableSparseRow) Set(val *SparseRow) {
	v.value = val
	v.isSet = true
}

func (v NullableSparseRow) IsSet() bool {
	return v.isSet
}

func (v *NullableSparseRow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSparseRow(val *SparseRow) *NullableSparseRow {
	return &NullableSparseRow{value: val, isSet: true}
}

func (v NullableSparseRow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSparseRow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// ODataEntitySetInfo struct for ODataEntitySetInfo
type ODataEntitySetInfo struct {
	TypeAnnotation *ODataTypeAnnotation `json:"typeAnnotation,omitempty"`
	Url NullableString `json:"url,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Title NullableString `json:"title,omitempty"`
}

// NewODataEntitySetInfo instantiates a new ODataEntitySetInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewODataEntitySetInfo() *ODataEntitySetInfo {
	this := ODataEntitySetInfo{}
	return &this
}

// NewODataEntitySetInfoWithDefaults instantiates a new ODataEntitySetInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewODataEntitySetInfoWithDefaults() *ODataEntitySetInfo {
	this := ODataEntitySetInfo{}
	return &this
}

// GetTypeAnnotation returns the TypeAnnotation field value if set, zero value otherwise.
func (o *ODataEntitySetInfo) GetTypeAnnotation() ODataTypeAnnotation {
	if o == nil || o.TypeAnnotation == nil {
		var ret ODataTypeAnnotation
		return ret
	}
	return *o.TypeAnnotation
}

// GetTypeAnnotationOk returns a tuple with the TypeAnnotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ODataEntitySetInfo) GetTypeAnnotationOk() (*ODataTypeAnnotation, bool) {
	if o == nil || o.TypeAnnotation == nil {
		return nil, false
	}
	return o.TypeAnnotation, true
}

// HasTypeAnnotation returns a boolean if a field has been set.
func (o *ODataEntitySetInfo) HasTypeAnnotation() bool {
	if o != nil && o.TypeAnnotation != nil {
		return true
	}

	return false
}

// SetTypeAnnotation gets a reference to the given ODataTypeAnnotation and assigns it to the TypeAnnotation field.
func (o *ODataEntitySetInfo) SetTypeAnnotation(v ODataTypeAnnotation) {
	o.TypeAnnotation = &v
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ODataEntitySetInfo) GetUrl() string {
	if o == nil || o.Url.Get() == nil {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ODataEntitySetInfo) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *ODataEntitySetInfo) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *ODataEntitySetInfo) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *ODataEntitySetInfo) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *ODataEntitySetInfo) UnsetUrl() {
	o.Url.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ODataEntitySetInfo) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ODataEntitySetInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ODataEntitySetInfo) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ODataEntitySetInfo) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ODataEntitySetInfo) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ODataEntitySetInfo) UnsetName() {
	o.Name.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ODataEntitySetInfo) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ODataEntitySetInfo) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *ODataEntitySetInfo) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *ODataEntitySetInfo) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *ODataEntitySetInfo) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *ODataEntitySetInfo) UnsetTitle() {
	o.Title.Unset()
}

func (o ODataEntitySetInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TypeAnnotation != nil {
		toSerialize["typeAnnotation"] = o.TypeAnnotation
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableODataEntitySetInfo struct {
	value *ODataEntitySetInfo
	isSet bool
}

func (v NullableODataEntitySetInfo) Get() *ODataEntitySetInfo {
	return v.value
}

func (v *NullableODataEntitySetInfo) Set(val *ODataEntitySetInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableODataEntitySetInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableODataEntitySetInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableODataEntitySetInfo(val *ODataEntitySetInfo) *NullableODataEntitySetInfo {
	return &NullableODataEntitySetInfo{value: val, isSet: true}
}

func (v NullableODataEntitySetInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableODataEntitySetInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



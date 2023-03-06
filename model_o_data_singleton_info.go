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

// checks if the ODataSingletonInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ODataSingletonInfo{}

// ODataSingletonInfo struct for ODataSingletonInfo
type ODataSingletonInfo struct {
	TypeAnnotation *ODataTypeAnnotation `json:"typeAnnotation,omitempty"`
	Url NullableString `json:"url,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Title NullableString `json:"title,omitempty"`
}

// NewODataSingletonInfo instantiates a new ODataSingletonInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewODataSingletonInfo() *ODataSingletonInfo {
	this := ODataSingletonInfo{}
	return &this
}

// NewODataSingletonInfoWithDefaults instantiates a new ODataSingletonInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewODataSingletonInfoWithDefaults() *ODataSingletonInfo {
	this := ODataSingletonInfo{}
	return &this
}

// GetTypeAnnotation returns the TypeAnnotation field value if set, zero value otherwise.
func (o *ODataSingletonInfo) GetTypeAnnotation() ODataTypeAnnotation {
	if o == nil || IsNil(o.TypeAnnotation) {
		var ret ODataTypeAnnotation
		return ret
	}
	return *o.TypeAnnotation
}

// GetTypeAnnotationOk returns a tuple with the TypeAnnotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ODataSingletonInfo) GetTypeAnnotationOk() (*ODataTypeAnnotation, bool) {
	if o == nil || IsNil(o.TypeAnnotation) {
		return nil, false
	}
	return o.TypeAnnotation, true
}

// HasTypeAnnotation returns a boolean if a field has been set.
func (o *ODataSingletonInfo) HasTypeAnnotation() bool {
	if o != nil && !IsNil(o.TypeAnnotation) {
		return true
	}

	return false
}

// SetTypeAnnotation gets a reference to the given ODataTypeAnnotation and assigns it to the TypeAnnotation field.
func (o *ODataSingletonInfo) SetTypeAnnotation(v ODataTypeAnnotation) {
	o.TypeAnnotation = &v
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ODataSingletonInfo) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ODataSingletonInfo) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *ODataSingletonInfo) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *ODataSingletonInfo) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *ODataSingletonInfo) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *ODataSingletonInfo) UnsetUrl() {
	o.Url.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ODataSingletonInfo) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ODataSingletonInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ODataSingletonInfo) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ODataSingletonInfo) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ODataSingletonInfo) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ODataSingletonInfo) UnsetName() {
	o.Name.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ODataSingletonInfo) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ODataSingletonInfo) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *ODataSingletonInfo) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *ODataSingletonInfo) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *ODataSingletonInfo) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *ODataSingletonInfo) UnsetTitle() {
	o.Title.Unset()
}

func (o ODataSingletonInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ODataSingletonInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TypeAnnotation) {
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
	return toSerialize, nil
}

type NullableODataSingletonInfo struct {
	value *ODataSingletonInfo
	isSet bool
}

func (v NullableODataSingletonInfo) Get() *ODataSingletonInfo {
	return v.value
}

func (v *NullableODataSingletonInfo) Set(val *ODataSingletonInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableODataSingletonInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableODataSingletonInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableODataSingletonInfo(val *ODataSingletonInfo) *NullableODataSingletonInfo {
	return &NullableODataSingletonInfo{value: val, isSet: true}
}

func (v NullableODataSingletonInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableODataSingletonInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
Landscape omnikeeper REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package okclient

import (
	"encoding/json"
	"fmt"
)

// AttributeValueType the model 'AttributeValueType'
type AttributeValueType string

// List of AttributeValueType
const (
	ATTRIBUTEVALUETYPE_TEXT AttributeValueType = "Text"
	ATTRIBUTEVALUETYPE_MULTILINE_TEXT AttributeValueType = "MultilineText"
	ATTRIBUTEVALUETYPE_INTEGER AttributeValueType = "Integer"
	ATTRIBUTEVALUETYPE_JSON AttributeValueType = "JSON"
	ATTRIBUTEVALUETYPE_YAML AttributeValueType = "YAML"
	ATTRIBUTEVALUETYPE_IMAGE AttributeValueType = "Image"
	ATTRIBUTEVALUETYPE_MASK AttributeValueType = "Mask"
	ATTRIBUTEVALUETYPE_DOUBLE AttributeValueType = "Double"
	ATTRIBUTEVALUETYPE_BOOLEAN AttributeValueType = "Boolean"
	ATTRIBUTEVALUETYPE_DATE_TIME_WITH_OFFSET AttributeValueType = "DateTimeWithOffset"
)

// All allowed values of AttributeValueType enum
var AllowedAttributeValueTypeEnumValues = []AttributeValueType{
	"Text",
	"MultilineText",
	"Integer",
	"JSON",
	"YAML",
	"Image",
	"Mask",
	"Double",
	"Boolean",
	"DateTimeWithOffset",
}

func (v *AttributeValueType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AttributeValueType(value)
	for _, existing := range AllowedAttributeValueTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AttributeValueType", value)
}

// NewAttributeValueTypeFromValue returns a pointer to a valid AttributeValueType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAttributeValueTypeFromValue(v string) (*AttributeValueType, error) {
	ev := AttributeValueType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AttributeValueType: valid values are %v", v, AllowedAttributeValueTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AttributeValueType) IsValid() bool {
	for _, existing := range AllowedAttributeValueTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AttributeValueType value
func (v AttributeValueType) Ptr() *AttributeValueType {
	return &v
}

type NullableAttributeValueType struct {
	value *AttributeValueType
	isSet bool
}

func (v NullableAttributeValueType) Get() *AttributeValueType {
	return v.value
}

func (v *NullableAttributeValueType) Set(val *AttributeValueType) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributeValueType) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributeValueType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributeValueType(val *AttributeValueType) *NullableAttributeValueType {
	return &NullableAttributeValueType{value: val, isSet: true}
}

func (v NullableAttributeValueType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributeValueType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


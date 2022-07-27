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

// SameTempIDHandling the model 'SameTempIDHandling'
type SameTempIDHandling string

// List of SameTempIDHandling
const (
	SAMETEMPIDHANDLING_DROP_AND_WARN SameTempIDHandling = "DropAndWarn"
	SAMETEMPIDHANDLING_DROP SameTempIDHandling = "Drop"
)

// All allowed values of SameTempIDHandling enum
var AllowedSameTempIDHandlingEnumValues = []SameTempIDHandling{
	"DropAndWarn",
	"Drop",
}

func (v *SameTempIDHandling) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SameTempIDHandling(value)
	for _, existing := range AllowedSameTempIDHandlingEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SameTempIDHandling", value)
}

// NewSameTempIDHandlingFromValue returns a pointer to a valid SameTempIDHandling
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSameTempIDHandlingFromValue(v string) (*SameTempIDHandling, error) {
	ev := SameTempIDHandling(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SameTempIDHandling: valid values are %v", v, AllowedSameTempIDHandlingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SameTempIDHandling) IsValid() bool {
	for _, existing := range AllowedSameTempIDHandlingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SameTempIDHandling value
func (v SameTempIDHandling) Ptr() *SameTempIDHandling {
	return &v
}

type NullableSameTempIDHandling struct {
	value *SameTempIDHandling
	isSet bool
}

func (v NullableSameTempIDHandling) Get() *SameTempIDHandling {
	return v.value
}

func (v *NullableSameTempIDHandling) Set(val *SameTempIDHandling) {
	v.value = val
	v.isSet = true
}

func (v NullableSameTempIDHandling) IsSet() bool {
	return v.isSet
}

func (v *NullableSameTempIDHandling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSameTempIDHandling(val *SameTempIDHandling) *NullableSameTempIDHandling {
	return &NullableSameTempIDHandling{value: val, isSet: true}
}

func (v NullableSameTempIDHandling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSameTempIDHandling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

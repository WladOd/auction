/*
Immudb Cloud Service

Specification of API to interact with Immudb Cloud Service.

API version: 1.0.0
Contact: contact@codenotary.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storage

import (
	"encoding/json"
	"fmt"
)

// FieldType the model 'FieldType'
type FieldType string

// List of FieldType
const (
	STRING FieldType = "STRING"
	BOOLEAN FieldType = "BOOLEAN"
	INTEGER FieldType = "INTEGER"
	DOUBLE FieldType = "DOUBLE"
)

// All allowed values of FieldType enum
var AllowedFieldTypeEnumValues = []FieldType{
	"STRING",
	"BOOLEAN",
	"INTEGER",
	"DOUBLE",
}

func (v *FieldType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldType(value)
	for _, existing := range AllowedFieldTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldType", value)
}

// NewFieldTypeFromValue returns a pointer to a valid FieldType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldTypeFromValue(v string) (*FieldType, error) {
	ev := FieldType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldType: valid values are %v", v, AllowedFieldTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldType) IsValid() bool {
	for _, existing := range AllowedFieldTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FieldType value
func (v FieldType) Ptr() *FieldType {
	return &v
}

type NullableFieldType struct {
	value *FieldType
	isSet bool
}

func (v NullableFieldType) Get() *FieldType {
	return v.value
}

func (v *NullableFieldType) Set(val *FieldType) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldType) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldType(val *FieldType) *NullableFieldType {
	return &NullableFieldType{value: val, isSet: true}
}

func (v NullableFieldType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

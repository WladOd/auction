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
)

// checks if the Index type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Index{}

// Index struct for Index
type Index struct {
	Fields []string `json:"fields"`
	IsUnique bool `json:"isUnique"`
}

// NewIndex instantiates a new Index object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndex(fields []string, isUnique bool) *Index {
	this := Index{}
	this.Fields = fields
	this.IsUnique = isUnique
	return &this
}

// NewIndexWithDefaults instantiates a new Index object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexWithDefaults() *Index {
	this := Index{}
	return &this
}

// GetFields returns the Fields field value
func (o *Index) GetFields() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *Index) GetFieldsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *Index) SetFields(v []string) {
	o.Fields = v
}

// GetIsUnique returns the IsUnique field value
func (o *Index) GetIsUnique() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsUnique
}

// GetIsUniqueOk returns a tuple with the IsUnique field value
// and a boolean to check if the value has been set.
func (o *Index) GetIsUniqueOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsUnique, true
}

// SetIsUnique sets field value
func (o *Index) SetIsUnique(v bool) {
	o.IsUnique = v
}

func (o Index) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Index) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fields"] = o.Fields
	toSerialize["isUnique"] = o.IsUnique
	return toSerialize, nil
}

type NullableIndex struct {
	value *Index
	isSet bool
}

func (v NullableIndex) Get() *Index {
	return v.value
}

func (v *NullableIndex) Set(val *Index) {
	v.value = val
	v.isSet = true
}

func (v NullableIndex) IsSet() bool {
	return v.isSet
}

func (v *NullableIndex) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndex(val *Index) *NullableIndex {
	return &NullableIndex{value: val, isSet: true}
}

func (v NullableIndex) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndex) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


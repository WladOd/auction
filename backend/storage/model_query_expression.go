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

// checks if the QueryExpression type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryExpression{}

// QueryExpression struct for QueryExpression
type QueryExpression struct {
	FieldComparisons []FieldComparison `json:"fieldComparisons,omitempty"`
}

// NewQueryExpression instantiates a new QueryExpression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryExpression() *QueryExpression {
	this := QueryExpression{}
	return &this
}

// NewQueryExpressionWithDefaults instantiates a new QueryExpression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryExpressionWithDefaults() *QueryExpression {
	this := QueryExpression{}
	return &this
}

// GetFieldComparisons returns the FieldComparisons field value if set, zero value otherwise.
func (o *QueryExpression) GetFieldComparisons() []FieldComparison {
	if o == nil || IsNil(o.FieldComparisons) {
		var ret []FieldComparison
		return ret
	}
	return o.FieldComparisons
}

// GetFieldComparisonsOk returns a tuple with the FieldComparisons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryExpression) GetFieldComparisonsOk() ([]FieldComparison, bool) {
	if o == nil || IsNil(o.FieldComparisons) {
		return nil, false
	}
	return o.FieldComparisons, true
}

// HasFieldComparisons returns a boolean if a field has been set.
func (o *QueryExpression) HasFieldComparisons() bool {
	if o != nil && !IsNil(o.FieldComparisons) {
		return true
	}

	return false
}

// SetFieldComparisons gets a reference to the given []FieldComparison and assigns it to the FieldComparisons field.
func (o *QueryExpression) SetFieldComparisons(v []FieldComparison) {
	o.FieldComparisons = v
}

func (o QueryExpression) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryExpression) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FieldComparisons) {
		toSerialize["fieldComparisons"] = o.FieldComparisons
	}
	return toSerialize, nil
}

type NullableQueryExpression struct {
	value *QueryExpression
	isSet bool
}

func (v NullableQueryExpression) Get() *QueryExpression {
	return v.value
}

func (v *NullableQueryExpression) Set(val *QueryExpression) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryExpression(val *QueryExpression) *NullableQueryExpression {
	return &NullableQueryExpression{value: val, isSet: true}
}

func (v NullableQueryExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



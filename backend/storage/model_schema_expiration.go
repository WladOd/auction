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

// checks if the SchemaExpiration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemaExpiration{}

// SchemaExpiration struct for SchemaExpiration
type SchemaExpiration struct {
	// Entry expiration time (unix timestamp in seconds)
	ExpiresAt *string `json:"expiresAt,omitempty"`
}

// NewSchemaExpiration instantiates a new SchemaExpiration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaExpiration() *SchemaExpiration {
	this := SchemaExpiration{}
	return &this
}

// NewSchemaExpirationWithDefaults instantiates a new SchemaExpiration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaExpirationWithDefaults() *SchemaExpiration {
	this := SchemaExpiration{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *SchemaExpiration) GetExpiresAt() string {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaExpiration) GetExpiresAtOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *SchemaExpiration) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *SchemaExpiration) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

func (o SchemaExpiration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemaExpiration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	return toSerialize, nil
}

type NullableSchemaExpiration struct {
	value *SchemaExpiration
	isSet bool
}

func (v NullableSchemaExpiration) Get() *SchemaExpiration {
	return v.value
}

func (v *NullableSchemaExpiration) Set(val *SchemaExpiration) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaExpiration) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaExpiration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaExpiration(val *SchemaExpiration) *NullableSchemaExpiration {
	return &NullableSchemaExpiration{value: val, isSet: true}
}

func (v NullableSchemaExpiration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaExpiration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


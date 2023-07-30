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

// checks if the SchemaSignature type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemaSignature{}

// SchemaSignature Signature for the new state value
type SchemaSignature struct {
	PublicKey *string `json:"publicKey,omitempty"`
	Signature *string `json:"signature,omitempty"`
}

// NewSchemaSignature instantiates a new SchemaSignature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaSignature() *SchemaSignature {
	this := SchemaSignature{}
	return &this
}

// NewSchemaSignatureWithDefaults instantiates a new SchemaSignature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaSignatureWithDefaults() *SchemaSignature {
	this := SchemaSignature{}
	return &this
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *SchemaSignature) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSignature) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *SchemaSignature) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *SchemaSignature) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *SchemaSignature) GetSignature() string {
	if o == nil || IsNil(o.Signature) {
		var ret string
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaSignature) GetSignatureOk() (*string, bool) {
	if o == nil || IsNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *SchemaSignature) HasSignature() bool {
	if o != nil && !IsNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given string and assigns it to the Signature field.
func (o *SchemaSignature) SetSignature(v string) {
	o.Signature = &v
}

func (o SchemaSignature) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemaSignature) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PublicKey) {
		toSerialize["publicKey"] = o.PublicKey
	}
	if !IsNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}
	return toSerialize, nil
}

type NullableSchemaSignature struct {
	value *SchemaSignature
	isSet bool
}

func (v NullableSchemaSignature) Get() *SchemaSignature {
	return v.value
}

func (v *NullableSchemaSignature) Set(val *SchemaSignature) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaSignature) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaSignature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaSignature(val *SchemaSignature) *NullableSchemaSignature {
	return &NullableSchemaSignature{value: val, isSet: true}
}

func (v NullableSchemaSignature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaSignature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



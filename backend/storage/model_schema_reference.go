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

// checks if the SchemaReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemaReference{}

// SchemaReference struct for SchemaReference
type SchemaReference struct {
	// Transaction if when the reference key was set
	Tx *string `json:"tx,omitempty"`
	// Reference key
	Key *string `json:"key,omitempty"`
	// At which transaction the key is bound, 0 if reference is not bound and should read the most recent reference
	AtTx *string `json:"atTx,omitempty"`
	Metadata *SchemaKVMetadata `json:"metadata,omitempty"`
	// Revision of the reference entry
	Revision *string `json:"revision,omitempty"`
}

// NewSchemaReference instantiates a new SchemaReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaReference() *SchemaReference {
	this := SchemaReference{}
	return &this
}

// NewSchemaReferenceWithDefaults instantiates a new SchemaReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaReferenceWithDefaults() *SchemaReference {
	this := SchemaReference{}
	return &this
}

// GetTx returns the Tx field value if set, zero value otherwise.
func (o *SchemaReference) GetTx() string {
	if o == nil || IsNil(o.Tx) {
		var ret string
		return ret
	}
	return *o.Tx
}

// GetTxOk returns a tuple with the Tx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaReference) GetTxOk() (*string, bool) {
	if o == nil || IsNil(o.Tx) {
		return nil, false
	}
	return o.Tx, true
}

// HasTx returns a boolean if a field has been set.
func (o *SchemaReference) HasTx() bool {
	if o != nil && !IsNil(o.Tx) {
		return true
	}

	return false
}

// SetTx gets a reference to the given string and assigns it to the Tx field.
func (o *SchemaReference) SetTx(v string) {
	o.Tx = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SchemaReference) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaReference) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SchemaReference) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *SchemaReference) SetKey(v string) {
	o.Key = &v
}

// GetAtTx returns the AtTx field value if set, zero value otherwise.
func (o *SchemaReference) GetAtTx() string {
	if o == nil || IsNil(o.AtTx) {
		var ret string
		return ret
	}
	return *o.AtTx
}

// GetAtTxOk returns a tuple with the AtTx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaReference) GetAtTxOk() (*string, bool) {
	if o == nil || IsNil(o.AtTx) {
		return nil, false
	}
	return o.AtTx, true
}

// HasAtTx returns a boolean if a field has been set.
func (o *SchemaReference) HasAtTx() bool {
	if o != nil && !IsNil(o.AtTx) {
		return true
	}

	return false
}

// SetAtTx gets a reference to the given string and assigns it to the AtTx field.
func (o *SchemaReference) SetAtTx(v string) {
	o.AtTx = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SchemaReference) GetMetadata() SchemaKVMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret SchemaKVMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaReference) GetMetadataOk() (*SchemaKVMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SchemaReference) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given SchemaKVMetadata and assigns it to the Metadata field.
func (o *SchemaReference) SetMetadata(v SchemaKVMetadata) {
	o.Metadata = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *SchemaReference) GetRevision() string {
	if o == nil || IsNil(o.Revision) {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaReference) GetRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *SchemaReference) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *SchemaReference) SetRevision(v string) {
	o.Revision = &v
}

func (o SchemaReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemaReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tx) {
		toSerialize["tx"] = o.Tx
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.AtTx) {
		toSerialize["atTx"] = o.AtTx
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	return toSerialize, nil
}

type NullableSchemaReference struct {
	value *SchemaReference
	isSet bool
}

func (v NullableSchemaReference) Get() *SchemaReference {
	return v.value
}

func (v *NullableSchemaReference) Set(val *SchemaReference) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaReference) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaReference(val *SchemaReference) *NullableSchemaReference {
	return &NullableSchemaReference{value: val, isSet: true}
}

func (v NullableSchemaReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// checks if the CollectionCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionCreateRequest{}

// CollectionCreateRequest struct for CollectionCreateRequest
type CollectionCreateRequest struct {
	IdFieldName *string `json:"idFieldName,omitempty"`
	Fields []Field `json:"fields,omitempty"`
	Indexes []Index `json:"indexes,omitempty"`
}

// NewCollectionCreateRequest instantiates a new CollectionCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionCreateRequest() *CollectionCreateRequest {
	this := CollectionCreateRequest{}
	return &this
}

// NewCollectionCreateRequestWithDefaults instantiates a new CollectionCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionCreateRequestWithDefaults() *CollectionCreateRequest {
	this := CollectionCreateRequest{}
	return &this
}

// GetIdFieldName returns the IdFieldName field value if set, zero value otherwise.
func (o *CollectionCreateRequest) GetIdFieldName() string {
	if o == nil || IsNil(o.IdFieldName) {
		var ret string
		return ret
	}
	return *o.IdFieldName
}

// GetIdFieldNameOk returns a tuple with the IdFieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionCreateRequest) GetIdFieldNameOk() (*string, bool) {
	if o == nil || IsNil(o.IdFieldName) {
		return nil, false
	}
	return o.IdFieldName, true
}

// HasIdFieldName returns a boolean if a field has been set.
func (o *CollectionCreateRequest) HasIdFieldName() bool {
	if o != nil && !IsNil(o.IdFieldName) {
		return true
	}

	return false
}

// SetIdFieldName gets a reference to the given string and assigns it to the IdFieldName field.
func (o *CollectionCreateRequest) SetIdFieldName(v string) {
	o.IdFieldName = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *CollectionCreateRequest) GetFields() []Field {
	if o == nil || IsNil(o.Fields) {
		var ret []Field
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionCreateRequest) GetFieldsOk() ([]Field, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *CollectionCreateRequest) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []Field and assigns it to the Fields field.
func (o *CollectionCreateRequest) SetFields(v []Field) {
	o.Fields = v
}

// GetIndexes returns the Indexes field value if set, zero value otherwise.
func (o *CollectionCreateRequest) GetIndexes() []Index {
	if o == nil || IsNil(o.Indexes) {
		var ret []Index
		return ret
	}
	return o.Indexes
}

// GetIndexesOk returns a tuple with the Indexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionCreateRequest) GetIndexesOk() ([]Index, bool) {
	if o == nil || IsNil(o.Indexes) {
		return nil, false
	}
	return o.Indexes, true
}

// HasIndexes returns a boolean if a field has been set.
func (o *CollectionCreateRequest) HasIndexes() bool {
	if o != nil && !IsNil(o.Indexes) {
		return true
	}

	return false
}

// SetIndexes gets a reference to the given []Index and assigns it to the Indexes field.
func (o *CollectionCreateRequest) SetIndexes(v []Index) {
	o.Indexes = v
}

func (o CollectionCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IdFieldName) {
		toSerialize["idFieldName"] = o.IdFieldName
	}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !IsNil(o.Indexes) {
		toSerialize["indexes"] = o.Indexes
	}
	return toSerialize, nil
}

type NullableCollectionCreateRequest struct {
	value *CollectionCreateRequest
	isSet bool
}

func (v NullableCollectionCreateRequest) Get() *CollectionCreateRequest {
	return v.value
}

func (v *NullableCollectionCreateRequest) Set(val *CollectionCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionCreateRequest(val *CollectionCreateRequest) *NullableCollectionCreateRequest {
	return &NullableCollectionCreateRequest{value: val, isSet: true}
}

func (v NullableCollectionCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



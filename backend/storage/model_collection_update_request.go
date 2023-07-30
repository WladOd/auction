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

// checks if the CollectionUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionUpdateRequest{}

// CollectionUpdateRequest struct for CollectionUpdateRequest
type CollectionUpdateRequest struct {
	IdFieldName string `json:"idFieldName"`
}

// NewCollectionUpdateRequest instantiates a new CollectionUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionUpdateRequest(idFieldName string) *CollectionUpdateRequest {
	this := CollectionUpdateRequest{}
	this.IdFieldName = idFieldName
	return &this
}

// NewCollectionUpdateRequestWithDefaults instantiates a new CollectionUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionUpdateRequestWithDefaults() *CollectionUpdateRequest {
	this := CollectionUpdateRequest{}
	return &this
}

// GetIdFieldName returns the IdFieldName field value
func (o *CollectionUpdateRequest) GetIdFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdFieldName
}

// GetIdFieldNameOk returns a tuple with the IdFieldName field value
// and a boolean to check if the value has been set.
func (o *CollectionUpdateRequest) GetIdFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdFieldName, true
}

// SetIdFieldName sets field value
func (o *CollectionUpdateRequest) SetIdFieldName(v string) {
	o.IdFieldName = v
}

func (o CollectionUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["idFieldName"] = o.IdFieldName
	return toSerialize, nil
}

type NullableCollectionUpdateRequest struct {
	value *CollectionUpdateRequest
	isSet bool
}

func (v NullableCollectionUpdateRequest) Get() *CollectionUpdateRequest {
	return v.value
}

func (v *NullableCollectionUpdateRequest) Set(val *CollectionUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionUpdateRequest(val *CollectionUpdateRequest) *NullableCollectionUpdateRequest {
	return &NullableCollectionUpdateRequest{value: val, isSet: true}
}

func (v NullableCollectionUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


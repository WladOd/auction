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

// checks if the DocumentInsertManyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentInsertManyResponse{}

// DocumentInsertManyResponse struct for DocumentInsertManyResponse
type DocumentInsertManyResponse struct {
	TransactionId *string `json:"transactionId,omitempty"`
	DocumentIds []string `json:"documentIds"`
}

// NewDocumentInsertManyResponse instantiates a new DocumentInsertManyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentInsertManyResponse(documentIds []string) *DocumentInsertManyResponse {
	this := DocumentInsertManyResponse{}
	this.DocumentIds = documentIds
	return &this
}

// NewDocumentInsertManyResponseWithDefaults instantiates a new DocumentInsertManyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentInsertManyResponseWithDefaults() *DocumentInsertManyResponse {
	this := DocumentInsertManyResponse{}
	return &this
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *DocumentInsertManyResponse) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentInsertManyResponse) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *DocumentInsertManyResponse) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *DocumentInsertManyResponse) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetDocumentIds returns the DocumentIds field value
func (o *DocumentInsertManyResponse) GetDocumentIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DocumentIds
}

// GetDocumentIdsOk returns a tuple with the DocumentIds field value
// and a boolean to check if the value has been set.
func (o *DocumentInsertManyResponse) GetDocumentIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocumentIds, true
}

// SetDocumentIds sets field value
func (o *DocumentInsertManyResponse) SetDocumentIds(v []string) {
	o.DocumentIds = v
}

func (o DocumentInsertManyResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentInsertManyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TransactionId) {
		toSerialize["transactionId"] = o.TransactionId
	}
	toSerialize["documentIds"] = o.DocumentIds
	return toSerialize, nil
}

type NullableDocumentInsertManyResponse struct {
	value *DocumentInsertManyResponse
	isSet bool
}

func (v NullableDocumentInsertManyResponse) Get() *DocumentInsertManyResponse {
	return v.value
}

func (v *NullableDocumentInsertManyResponse) Set(val *DocumentInsertManyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentInsertManyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentInsertManyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentInsertManyResponse(val *DocumentInsertManyResponse) *NullableDocumentInsertManyResponse {
	return &NullableDocumentInsertManyResponse{value: val, isSet: true}
}

func (v NullableDocumentInsertManyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentInsertManyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



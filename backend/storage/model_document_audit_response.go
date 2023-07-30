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

// checks if the DocumentAuditResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentAuditResponse{}

// DocumentAuditResponse struct for DocumentAuditResponse
type DocumentAuditResponse struct {
	Revisions []DocumentAtRevision `json:"revisions"`
}

// NewDocumentAuditResponse instantiates a new DocumentAuditResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentAuditResponse(revisions []DocumentAtRevision) *DocumentAuditResponse {
	this := DocumentAuditResponse{}
	this.Revisions = revisions
	return &this
}

// NewDocumentAuditResponseWithDefaults instantiates a new DocumentAuditResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentAuditResponseWithDefaults() *DocumentAuditResponse {
	this := DocumentAuditResponse{}
	return &this
}

// GetRevisions returns the Revisions field value
func (o *DocumentAuditResponse) GetRevisions() []DocumentAtRevision {
	if o == nil {
		var ret []DocumentAtRevision
		return ret
	}

	return o.Revisions
}

// GetRevisionsOk returns a tuple with the Revisions field value
// and a boolean to check if the value has been set.
func (o *DocumentAuditResponse) GetRevisionsOk() ([]DocumentAtRevision, bool) {
	if o == nil {
		return nil, false
	}
	return o.Revisions, true
}

// SetRevisions sets field value
func (o *DocumentAuditResponse) SetRevisions(v []DocumentAtRevision) {
	o.Revisions = v
}

func (o DocumentAuditResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentAuditResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["revisions"] = o.Revisions
	return toSerialize, nil
}

type NullableDocumentAuditResponse struct {
	value *DocumentAuditResponse
	isSet bool
}

func (v NullableDocumentAuditResponse) Get() *DocumentAuditResponse {
	return v.value
}

func (v *NullableDocumentAuditResponse) Set(val *DocumentAuditResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentAuditResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentAuditResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentAuditResponse(val *DocumentAuditResponse) *NullableDocumentAuditResponse {
	return &NullableDocumentAuditResponse{value: val, isSet: true}
}

func (v NullableDocumentAuditResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentAuditResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



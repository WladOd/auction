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

// checks if the DocumentSearchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentSearchResponse{}

// DocumentSearchResponse struct for DocumentSearchResponse
type DocumentSearchResponse struct {
	Page int32 `json:"page"`
	PerPage int32 `json:"perPage"`
	SearchId string `json:"searchId"`
	Revisions []DocumentAtRevision `json:"revisions"`
}

// NewDocumentSearchResponse instantiates a new DocumentSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentSearchResponse(page int32, perPage int32, searchId string, revisions []DocumentAtRevision) *DocumentSearchResponse {
	this := DocumentSearchResponse{}
	this.Page = page
	this.PerPage = perPage
	this.SearchId = searchId
	this.Revisions = revisions
	return &this
}

// NewDocumentSearchResponseWithDefaults instantiates a new DocumentSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentSearchResponseWithDefaults() *DocumentSearchResponse {
	this := DocumentSearchResponse{}
	return &this
}

// GetPage returns the Page field value
func (o *DocumentSearchResponse) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *DocumentSearchResponse) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *DocumentSearchResponse) SetPage(v int32) {
	o.Page = v
}

// GetPerPage returns the PerPage field value
func (o *DocumentSearchResponse) GetPerPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PerPage
}

// GetPerPageOk returns a tuple with the PerPage field value
// and a boolean to check if the value has been set.
func (o *DocumentSearchResponse) GetPerPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PerPage, true
}

// SetPerPage sets field value
func (o *DocumentSearchResponse) SetPerPage(v int32) {
	o.PerPage = v
}

// GetSearchId returns the SearchId field value
func (o *DocumentSearchResponse) GetSearchId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SearchId
}

// GetSearchIdOk returns a tuple with the SearchId field value
// and a boolean to check if the value has been set.
func (o *DocumentSearchResponse) GetSearchIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SearchId, true
}

// SetSearchId sets field value
func (o *DocumentSearchResponse) SetSearchId(v string) {
	o.SearchId = v
}

// GetRevisions returns the Revisions field value
func (o *DocumentSearchResponse) GetRevisions() []DocumentAtRevision {
	if o == nil {
		var ret []DocumentAtRevision
		return ret
	}

	return o.Revisions
}

// GetRevisionsOk returns a tuple with the Revisions field value
// and a boolean to check if the value has been set.
func (o *DocumentSearchResponse) GetRevisionsOk() ([]DocumentAtRevision, bool) {
	if o == nil {
		return nil, false
	}
	return o.Revisions, true
}

// SetRevisions sets field value
func (o *DocumentSearchResponse) SetRevisions(v []DocumentAtRevision) {
	o.Revisions = v
}

func (o DocumentSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentSearchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["page"] = o.Page
	toSerialize["perPage"] = o.PerPage
	toSerialize["searchId"] = o.SearchId
	toSerialize["revisions"] = o.Revisions
	return toSerialize, nil
}

type NullableDocumentSearchResponse struct {
	value *DocumentSearchResponse
	isSet bool
}

func (v NullableDocumentSearchResponse) Get() *DocumentSearchResponse {
	return v.value
}

func (v *NullableDocumentSearchResponse) Set(val *DocumentSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentSearchResponse(val *DocumentSearchResponse) *NullableDocumentSearchResponse {
	return &NullableDocumentSearchResponse{value: val, isSet: true}
}

func (v NullableDocumentSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



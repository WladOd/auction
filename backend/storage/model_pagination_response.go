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

// checks if the PaginationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginationResponse{}

// PaginationResponse struct for PaginationResponse
type PaginationResponse struct {
	Page int32 `json:"page"`
	PerPage int32 `json:"perPage"`
}

// NewPaginationResponse instantiates a new PaginationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationResponse(page int32, perPage int32) *PaginationResponse {
	this := PaginationResponse{}
	this.Page = page
	this.PerPage = perPage
	return &this
}

// NewPaginationResponseWithDefaults instantiates a new PaginationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationResponseWithDefaults() *PaginationResponse {
	this := PaginationResponse{}
	return &this
}

// GetPage returns the Page field value
func (o *PaginationResponse) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *PaginationResponse) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *PaginationResponse) SetPage(v int32) {
	o.Page = v
}

// GetPerPage returns the PerPage field value
func (o *PaginationResponse) GetPerPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PerPage
}

// GetPerPageOk returns a tuple with the PerPage field value
// and a boolean to check if the value has been set.
func (o *PaginationResponse) GetPerPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PerPage, true
}

// SetPerPage sets field value
func (o *PaginationResponse) SetPerPage(v int32) {
	o.PerPage = v
}

func (o PaginationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["page"] = o.Page
	toSerialize["perPage"] = o.PerPage
	return toSerialize, nil
}

type NullablePaginationResponse struct {
	value *PaginationResponse
	isSet bool
}

func (v NullablePaginationResponse) Get() *PaginationResponse {
	return v.value
}

func (v *NullablePaginationResponse) Set(val *PaginationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationResponse(val *PaginationResponse) *NullablePaginationResponse {
	return &NullablePaginationResponse{value: val, isSet: true}
}

func (v NullablePaginationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



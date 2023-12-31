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
	"time"
)

// checks if the ExportInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportInformation{}

// ExportInformation struct for ExportInformation
type ExportInformation struct {
	Id string `json:"id"`
	Message string `json:"message"`
	Date time.Time `json:"date"`
}

// NewExportInformation instantiates a new ExportInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportInformation(id string, message string, date time.Time) *ExportInformation {
	this := ExportInformation{}
	this.Id = id
	this.Message = message
	this.Date = date
	return &this
}

// NewExportInformationWithDefaults instantiates a new ExportInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportInformationWithDefaults() *ExportInformation {
	this := ExportInformation{}
	return &this
}

// GetId returns the Id field value
func (o *ExportInformation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExportInformation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExportInformation) SetId(v string) {
	o.Id = v
}

// GetMessage returns the Message field value
func (o *ExportInformation) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ExportInformation) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ExportInformation) SetMessage(v string) {
	o.Message = v
}

// GetDate returns the Date field value
func (o *ExportInformation) GetDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *ExportInformation) GetDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *ExportInformation) SetDate(v time.Time) {
	o.Date = v
}

func (o ExportInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExportInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["message"] = o.Message
	toSerialize["date"] = o.Date
	return toSerialize, nil
}

type NullableExportInformation struct {
	value *ExportInformation
	isSet bool
}

func (v NullableExportInformation) Get() *ExportInformation {
	return v.value
}

func (v *NullableExportInformation) Set(val *ExportInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableExportInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableExportInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportInformation(val *ExportInformation) *NullableExportInformation {
	return &NullableExportInformation{value: val, isSet: true}
}

func (v NullableExportInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



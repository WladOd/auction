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

// checks if the CreateCollectionErrReply type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCollectionErrReply{}

// CreateCollectionErrReply struct for CreateCollectionErrReply
type CreateCollectionErrReply struct {
	// application-specific error code
	Code int32 `json:"code"`
	// application-level error message, for debugging
	Error string `json:"error"`
	// user-level status message
	Status string `json:"status"`
	// the maximum number of collections a tenant can create with the current license
	MaxNumberOfCollections int32 `json:"max_number_of_collections"`
	// ledger DB max size exceed
	MaxLedgerDbSize float64 `json:"max_ledger_db_size"`
	// the type of the license of the tenant
	LicenseType string `json:"license_type"`
}

// NewCreateCollectionErrReply instantiates a new CreateCollectionErrReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCollectionErrReply(code int32, error_ string, status string, maxNumberOfCollections int32, maxLedgerDbSize float64, licenseType string) *CreateCollectionErrReply {
	this := CreateCollectionErrReply{}
	this.Code = code
	this.Error = error_
	this.Status = status
	this.MaxNumberOfCollections = maxNumberOfCollections
	this.MaxLedgerDbSize = maxLedgerDbSize
	this.LicenseType = licenseType
	return &this
}

// NewCreateCollectionErrReplyWithDefaults instantiates a new CreateCollectionErrReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCollectionErrReplyWithDefaults() *CreateCollectionErrReply {
	this := CreateCollectionErrReply{}
	return &this
}

// GetCode returns the Code field value
func (o *CreateCollectionErrReply) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *CreateCollectionErrReply) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *CreateCollectionErrReply) SetCode(v int32) {
	o.Code = v
}

// GetError returns the Error field value
func (o *CreateCollectionErrReply) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *CreateCollectionErrReply) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *CreateCollectionErrReply) SetError(v string) {
	o.Error = v
}

// GetStatus returns the Status field value
func (o *CreateCollectionErrReply) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CreateCollectionErrReply) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CreateCollectionErrReply) SetStatus(v string) {
	o.Status = v
}

// GetMaxNumberOfCollections returns the MaxNumberOfCollections field value
func (o *CreateCollectionErrReply) GetMaxNumberOfCollections() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxNumberOfCollections
}

// GetMaxNumberOfCollectionsOk returns a tuple with the MaxNumberOfCollections field value
// and a boolean to check if the value has been set.
func (o *CreateCollectionErrReply) GetMaxNumberOfCollectionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxNumberOfCollections, true
}

// SetMaxNumberOfCollections sets field value
func (o *CreateCollectionErrReply) SetMaxNumberOfCollections(v int32) {
	o.MaxNumberOfCollections = v
}

// GetMaxLedgerDbSize returns the MaxLedgerDbSize field value
func (o *CreateCollectionErrReply) GetMaxLedgerDbSize() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.MaxLedgerDbSize
}

// GetMaxLedgerDbSizeOk returns a tuple with the MaxLedgerDbSize field value
// and a boolean to check if the value has been set.
func (o *CreateCollectionErrReply) GetMaxLedgerDbSizeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxLedgerDbSize, true
}

// SetMaxLedgerDbSize sets field value
func (o *CreateCollectionErrReply) SetMaxLedgerDbSize(v float64) {
	o.MaxLedgerDbSize = v
}

// GetLicenseType returns the LicenseType field value
func (o *CreateCollectionErrReply) GetLicenseType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LicenseType
}

// GetLicenseTypeOk returns a tuple with the LicenseType field value
// and a boolean to check if the value has been set.
func (o *CreateCollectionErrReply) GetLicenseTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LicenseType, true
}

// SetLicenseType sets field value
func (o *CreateCollectionErrReply) SetLicenseType(v string) {
	o.LicenseType = v
}

func (o CreateCollectionErrReply) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCollectionErrReply) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["error"] = o.Error
	toSerialize["status"] = o.Status
	toSerialize["max_number_of_collections"] = o.MaxNumberOfCollections
	toSerialize["max_ledger_db_size"] = o.MaxLedgerDbSize
	toSerialize["license_type"] = o.LicenseType
	return toSerialize, nil
}

type NullableCreateCollectionErrReply struct {
	value *CreateCollectionErrReply
	isSet bool
}

func (v NullableCreateCollectionErrReply) Get() *CreateCollectionErrReply {
	return v.value
}

func (v *NullableCreateCollectionErrReply) Set(val *CreateCollectionErrReply) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCollectionErrReply) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCollectionErrReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCollectionErrReply(val *CreateCollectionErrReply) *NullableCreateCollectionErrReply {
	return &NullableCreateCollectionErrReply{value: val, isSet: true}
}

func (v NullableCreateCollectionErrReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCollectionErrReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



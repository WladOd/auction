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

// checks if the SchemaTx type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemaTx{}

// SchemaTx Transaction to verify
type SchemaTx struct {
	Header *SchemaTxHeader `json:"header,omitempty"`
	// Raw entry values
	Entries []SchemaTxEntry `json:"entries,omitempty"`
	// KV entries in the transaction (parsed)
	KvEntries []SchemaEntry `json:"kvEntries,omitempty"`
	// Sorted Set entries in the transaction (parsed)
	ZEntries []SchemaZEntry `json:"zEntries,omitempty"`
}

// NewSchemaTx instantiates a new SchemaTx object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaTx() *SchemaTx {
	this := SchemaTx{}
	return &this
}

// NewSchemaTxWithDefaults instantiates a new SchemaTx object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaTxWithDefaults() *SchemaTx {
	this := SchemaTx{}
	return &this
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *SchemaTx) GetHeader() SchemaTxHeader {
	if o == nil || IsNil(o.Header) {
		var ret SchemaTxHeader
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaTx) GetHeaderOk() (*SchemaTxHeader, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *SchemaTx) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given SchemaTxHeader and assigns it to the Header field.
func (o *SchemaTx) SetHeader(v SchemaTxHeader) {
	o.Header = &v
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *SchemaTx) GetEntries() []SchemaTxEntry {
	if o == nil || IsNil(o.Entries) {
		var ret []SchemaTxEntry
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaTx) GetEntriesOk() ([]SchemaTxEntry, bool) {
	if o == nil || IsNil(o.Entries) {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *SchemaTx) HasEntries() bool {
	if o != nil && !IsNil(o.Entries) {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []SchemaTxEntry and assigns it to the Entries field.
func (o *SchemaTx) SetEntries(v []SchemaTxEntry) {
	o.Entries = v
}

// GetKvEntries returns the KvEntries field value if set, zero value otherwise.
func (o *SchemaTx) GetKvEntries() []SchemaEntry {
	if o == nil || IsNil(o.KvEntries) {
		var ret []SchemaEntry
		return ret
	}
	return o.KvEntries
}

// GetKvEntriesOk returns a tuple with the KvEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaTx) GetKvEntriesOk() ([]SchemaEntry, bool) {
	if o == nil || IsNil(o.KvEntries) {
		return nil, false
	}
	return o.KvEntries, true
}

// HasKvEntries returns a boolean if a field has been set.
func (o *SchemaTx) HasKvEntries() bool {
	if o != nil && !IsNil(o.KvEntries) {
		return true
	}

	return false
}

// SetKvEntries gets a reference to the given []SchemaEntry and assigns it to the KvEntries field.
func (o *SchemaTx) SetKvEntries(v []SchemaEntry) {
	o.KvEntries = v
}

// GetZEntries returns the ZEntries field value if set, zero value otherwise.
func (o *SchemaTx) GetZEntries() []SchemaZEntry {
	if o == nil || IsNil(o.ZEntries) {
		var ret []SchemaZEntry
		return ret
	}
	return o.ZEntries
}

// GetZEntriesOk returns a tuple with the ZEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaTx) GetZEntriesOk() ([]SchemaZEntry, bool) {
	if o == nil || IsNil(o.ZEntries) {
		return nil, false
	}
	return o.ZEntries, true
}

// HasZEntries returns a boolean if a field has been set.
func (o *SchemaTx) HasZEntries() bool {
	if o != nil && !IsNil(o.ZEntries) {
		return true
	}

	return false
}

// SetZEntries gets a reference to the given []SchemaZEntry and assigns it to the ZEntries field.
func (o *SchemaTx) SetZEntries(v []SchemaZEntry) {
	o.ZEntries = v
}

func (o SchemaTx) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemaTx) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	if !IsNil(o.Entries) {
		toSerialize["entries"] = o.Entries
	}
	if !IsNil(o.KvEntries) {
		toSerialize["kvEntries"] = o.KvEntries
	}
	if !IsNil(o.ZEntries) {
		toSerialize["zEntries"] = o.ZEntries
	}
	return toSerialize, nil
}

type NullableSchemaTx struct {
	value *SchemaTx
	isSet bool
}

func (v NullableSchemaTx) Get() *SchemaTx {
	return v.value
}

func (v *NullableSchemaTx) Set(val *SchemaTx) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaTx) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaTx) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaTx(val *SchemaTx) *NullableSchemaTx {
	return &NullableSchemaTx{value: val, isSet: true}
}

func (v NullableSchemaTx) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaTx) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



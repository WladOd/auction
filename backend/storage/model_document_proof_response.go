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

// checks if the DocumentProofResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentProofResponse{}

// DocumentProofResponse struct for DocumentProofResponse
type DocumentProofResponse struct {
	Database string `json:"database"`
	CollectionId int64 `json:"collectionId"`
	IdFieldName string `json:"idFieldName"`
	EncodedDocument string `json:"encodedDocument"`
	VerifiableTx SchemaVerifiableTxV2 `json:"verifiableTx"`
}

// NewDocumentProofResponse instantiates a new DocumentProofResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentProofResponse(database string, collectionId int64, idFieldName string, encodedDocument string, verifiableTx SchemaVerifiableTxV2) *DocumentProofResponse {
	this := DocumentProofResponse{}
	this.Database = database
	this.CollectionId = collectionId
	this.IdFieldName = idFieldName
	this.EncodedDocument = encodedDocument
	this.VerifiableTx = verifiableTx
	return &this
}

// NewDocumentProofResponseWithDefaults instantiates a new DocumentProofResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentProofResponseWithDefaults() *DocumentProofResponse {
	this := DocumentProofResponse{}
	return &this
}

// GetDatabase returns the Database field value
func (o *DocumentProofResponse) GetDatabase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value
// and a boolean to check if the value has been set.
func (o *DocumentProofResponse) GetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Database, true
}

// SetDatabase sets field value
func (o *DocumentProofResponse) SetDatabase(v string) {
	o.Database = v
}

// GetCollectionId returns the CollectionId field value
func (o *DocumentProofResponse) GetCollectionId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CollectionId
}

// GetCollectionIdOk returns a tuple with the CollectionId field value
// and a boolean to check if the value has been set.
func (o *DocumentProofResponse) GetCollectionIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionId, true
}

// SetCollectionId sets field value
func (o *DocumentProofResponse) SetCollectionId(v int64) {
	o.CollectionId = v
}

// GetIdFieldName returns the IdFieldName field value
func (o *DocumentProofResponse) GetIdFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdFieldName
}

// GetIdFieldNameOk returns a tuple with the IdFieldName field value
// and a boolean to check if the value has been set.
func (o *DocumentProofResponse) GetIdFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdFieldName, true
}

// SetIdFieldName sets field value
func (o *DocumentProofResponse) SetIdFieldName(v string) {
	o.IdFieldName = v
}

// GetEncodedDocument returns the EncodedDocument field value
func (o *DocumentProofResponse) GetEncodedDocument() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EncodedDocument
}

// GetEncodedDocumentOk returns a tuple with the EncodedDocument field value
// and a boolean to check if the value has been set.
func (o *DocumentProofResponse) GetEncodedDocumentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EncodedDocument, true
}

// SetEncodedDocument sets field value
func (o *DocumentProofResponse) SetEncodedDocument(v string) {
	o.EncodedDocument = v
}

// GetVerifiableTx returns the VerifiableTx field value
func (o *DocumentProofResponse) GetVerifiableTx() SchemaVerifiableTxV2 {
	if o == nil {
		var ret SchemaVerifiableTxV2
		return ret
	}

	return o.VerifiableTx
}

// GetVerifiableTxOk returns a tuple with the VerifiableTx field value
// and a boolean to check if the value has been set.
func (o *DocumentProofResponse) GetVerifiableTxOk() (*SchemaVerifiableTxV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerifiableTx, true
}

// SetVerifiableTx sets field value
func (o *DocumentProofResponse) SetVerifiableTx(v SchemaVerifiableTxV2) {
	o.VerifiableTx = v
}

func (o DocumentProofResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentProofResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["database"] = o.Database
	toSerialize["collectionId"] = o.CollectionId
	toSerialize["idFieldName"] = o.IdFieldName
	toSerialize["encodedDocument"] = o.EncodedDocument
	toSerialize["verifiableTx"] = o.VerifiableTx
	return toSerialize, nil
}

type NullableDocumentProofResponse struct {
	value *DocumentProofResponse
	isSet bool
}

func (v NullableDocumentProofResponse) Get() *DocumentProofResponse {
	return v.value
}

func (v *NullableDocumentProofResponse) Set(val *DocumentProofResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentProofResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentProofResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentProofResponse(val *DocumentProofResponse) *NullableDocumentProofResponse {
	return &NullableDocumentProofResponse{value: val, isSet: true}
}

func (v NullableDocumentProofResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentProofResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



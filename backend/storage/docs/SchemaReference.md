# SchemaReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tx** | Pointer to **string** | Transaction if when the reference key was set | [optional] 
**Key** | Pointer to **string** | Reference key | [optional] 
**AtTx** | Pointer to **string** | At which transaction the key is bound, 0 if reference is not bound and should read the most recent reference | [optional] 
**Metadata** | Pointer to [**SchemaKVMetadata**](SchemaKVMetadata.md) |  | [optional] 
**Revision** | Pointer to **string** | Revision of the reference entry | [optional] 

## Methods

### NewSchemaReference

`func NewSchemaReference() *SchemaReference`

NewSchemaReference instantiates a new SchemaReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaReferenceWithDefaults

`func NewSchemaReferenceWithDefaults() *SchemaReference`

NewSchemaReferenceWithDefaults instantiates a new SchemaReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTx

`func (o *SchemaReference) GetTx() string`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *SchemaReference) GetTxOk() (*string, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *SchemaReference) SetTx(v string)`

SetTx sets Tx field to given value.

### HasTx

`func (o *SchemaReference) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetKey

`func (o *SchemaReference) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SchemaReference) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SchemaReference) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SchemaReference) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetAtTx

`func (o *SchemaReference) GetAtTx() string`

GetAtTx returns the AtTx field if non-nil, zero value otherwise.

### GetAtTxOk

`func (o *SchemaReference) GetAtTxOk() (*string, bool)`

GetAtTxOk returns a tuple with the AtTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtTx

`func (o *SchemaReference) SetAtTx(v string)`

SetAtTx sets AtTx field to given value.

### HasAtTx

`func (o *SchemaReference) HasAtTx() bool`

HasAtTx returns a boolean if a field has been set.

### GetMetadata

`func (o *SchemaReference) GetMetadata() SchemaKVMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SchemaReference) GetMetadataOk() (*SchemaKVMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SchemaReference) SetMetadata(v SchemaKVMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SchemaReference) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRevision

`func (o *SchemaReference) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *SchemaReference) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *SchemaReference) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *SchemaReference) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



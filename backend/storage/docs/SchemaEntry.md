# SchemaEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tx** | Pointer to **string** | Transaction id at which the target value was set (i.e. not the reference transaction id) | [optional] 
**Key** | Pointer to **string** | Key of the target value (i.e. not the reference entry) | [optional] 
**Value** | Pointer to **string** | Value | [optional] 
**ReferencedBy** | Pointer to [**SchemaReference**](SchemaReference.md) |  | [optional] 
**Metadata** | Pointer to [**SchemaKVMetadata**](SchemaKVMetadata.md) |  | [optional] 
**Expired** | Pointer to **bool** | If set to true, this entry has expired and the value is not retrieved | [optional] 
**Revision** | Pointer to **string** | Key&#39;s revision, in case of GetAt it will be 0 | [optional] 

## Methods

### NewSchemaEntry

`func NewSchemaEntry() *SchemaEntry`

NewSchemaEntry instantiates a new SchemaEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaEntryWithDefaults

`func NewSchemaEntryWithDefaults() *SchemaEntry`

NewSchemaEntryWithDefaults instantiates a new SchemaEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTx

`func (o *SchemaEntry) GetTx() string`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *SchemaEntry) GetTxOk() (*string, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *SchemaEntry) SetTx(v string)`

SetTx sets Tx field to given value.

### HasTx

`func (o *SchemaEntry) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetKey

`func (o *SchemaEntry) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SchemaEntry) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SchemaEntry) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SchemaEntry) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *SchemaEntry) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SchemaEntry) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SchemaEntry) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SchemaEntry) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetReferencedBy

`func (o *SchemaEntry) GetReferencedBy() SchemaReference`

GetReferencedBy returns the ReferencedBy field if non-nil, zero value otherwise.

### GetReferencedByOk

`func (o *SchemaEntry) GetReferencedByOk() (*SchemaReference, bool)`

GetReferencedByOk returns a tuple with the ReferencedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedBy

`func (o *SchemaEntry) SetReferencedBy(v SchemaReference)`

SetReferencedBy sets ReferencedBy field to given value.

### HasReferencedBy

`func (o *SchemaEntry) HasReferencedBy() bool`

HasReferencedBy returns a boolean if a field has been set.

### GetMetadata

`func (o *SchemaEntry) GetMetadata() SchemaKVMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SchemaEntry) GetMetadataOk() (*SchemaKVMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SchemaEntry) SetMetadata(v SchemaKVMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SchemaEntry) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetExpired

`func (o *SchemaEntry) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *SchemaEntry) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *SchemaEntry) SetExpired(v bool)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *SchemaEntry) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetRevision

`func (o *SchemaEntry) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *SchemaEntry) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *SchemaEntry) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *SchemaEntry) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



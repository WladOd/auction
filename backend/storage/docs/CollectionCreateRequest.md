# CollectionCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdFieldName** | Pointer to **string** |  | [optional] 
**Fields** | Pointer to [**[]Field**](Field.md) |  | [optional] 
**Indexes** | Pointer to [**[]Index**](Index.md) |  | [optional] 

## Methods

### NewCollectionCreateRequest

`func NewCollectionCreateRequest() *CollectionCreateRequest`

NewCollectionCreateRequest instantiates a new CollectionCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionCreateRequestWithDefaults

`func NewCollectionCreateRequestWithDefaults() *CollectionCreateRequest`

NewCollectionCreateRequestWithDefaults instantiates a new CollectionCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdFieldName

`func (o *CollectionCreateRequest) GetIdFieldName() string`

GetIdFieldName returns the IdFieldName field if non-nil, zero value otherwise.

### GetIdFieldNameOk

`func (o *CollectionCreateRequest) GetIdFieldNameOk() (*string, bool)`

GetIdFieldNameOk returns a tuple with the IdFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdFieldName

`func (o *CollectionCreateRequest) SetIdFieldName(v string)`

SetIdFieldName sets IdFieldName field to given value.

### HasIdFieldName

`func (o *CollectionCreateRequest) HasIdFieldName() bool`

HasIdFieldName returns a boolean if a field has been set.

### GetFields

`func (o *CollectionCreateRequest) GetFields() []Field`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *CollectionCreateRequest) GetFieldsOk() (*[]Field, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *CollectionCreateRequest) SetFields(v []Field)`

SetFields sets Fields field to given value.

### HasFields

`func (o *CollectionCreateRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetIndexes

`func (o *CollectionCreateRequest) GetIndexes() []Index`

GetIndexes returns the Indexes field if non-nil, zero value otherwise.

### GetIndexesOk

`func (o *CollectionCreateRequest) GetIndexesOk() (*[]Index, bool)`

GetIndexesOk returns a tuple with the Indexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexes

`func (o *CollectionCreateRequest) SetIndexes(v []Index)`

SetIndexes sets Indexes field to given value.

### HasIndexes

`func (o *CollectionCreateRequest) HasIndexes() bool`

HasIndexes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



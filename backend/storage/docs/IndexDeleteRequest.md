# IndexDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | **string** |  | 
**Fields** | **[]string** |  | 

## Methods

### NewIndexDeleteRequest

`func NewIndexDeleteRequest(collection string, fields []string, ) *IndexDeleteRequest`

NewIndexDeleteRequest instantiates a new IndexDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexDeleteRequestWithDefaults

`func NewIndexDeleteRequestWithDefaults() *IndexDeleteRequest`

NewIndexDeleteRequestWithDefaults instantiates a new IndexDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *IndexDeleteRequest) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *IndexDeleteRequest) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *IndexDeleteRequest) SetCollection(v string)`

SetCollection sets Collection field to given value.


### GetFields

`func (o *IndexDeleteRequest) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *IndexDeleteRequest) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *IndexDeleteRequest) SetFields(v []string)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



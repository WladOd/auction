# DocumentsCountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | **string** |  | 
**Count** | **int32** |  | 

## Methods

### NewDocumentsCountResponse

`func NewDocumentsCountResponse(collection string, count int32, ) *DocumentsCountResponse`

NewDocumentsCountResponse instantiates a new DocumentsCountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentsCountResponseWithDefaults

`func NewDocumentsCountResponseWithDefaults() *DocumentsCountResponse`

NewDocumentsCountResponseWithDefaults instantiates a new DocumentsCountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *DocumentsCountResponse) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *DocumentsCountResponse) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *DocumentsCountResponse) SetCollection(v string)`

SetCollection sets Collection field to given value.


### GetCount

`func (o *DocumentsCountResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DocumentsCountResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DocumentsCountResponse) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



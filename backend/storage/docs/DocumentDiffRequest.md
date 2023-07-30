# DocumentDiffRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **int32** |  | 
**PerPage** | **int32** |  | 
**Desc** | Pointer to **bool** |  | [optional] 

## Methods

### NewDocumentDiffRequest

`func NewDocumentDiffRequest(page int32, perPage int32, ) *DocumentDiffRequest`

NewDocumentDiffRequest instantiates a new DocumentDiffRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentDiffRequestWithDefaults

`func NewDocumentDiffRequestWithDefaults() *DocumentDiffRequest`

NewDocumentDiffRequestWithDefaults instantiates a new DocumentDiffRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *DocumentDiffRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *DocumentDiffRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *DocumentDiffRequest) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPerPage

`func (o *DocumentDiffRequest) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *DocumentDiffRequest) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *DocumentDiffRequest) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.


### GetDesc

`func (o *DocumentDiffRequest) GetDesc() bool`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *DocumentDiffRequest) GetDescOk() (*bool, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *DocumentDiffRequest) SetDesc(v bool)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *DocumentDiffRequest) HasDesc() bool`

HasDesc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



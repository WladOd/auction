# DocumentSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **int32** |  | 
**PerPage** | **int32** |  | 
**SearchId** | Pointer to **string** |  | [optional] 
**KeepOpen** | Pointer to **bool** |  | [optional] 
**Query** | Pointer to [**Query**](Query.md) |  | [optional] 

## Methods

### NewDocumentSearchRequest

`func NewDocumentSearchRequest(page int32, perPage int32, ) *DocumentSearchRequest`

NewDocumentSearchRequest instantiates a new DocumentSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentSearchRequestWithDefaults

`func NewDocumentSearchRequestWithDefaults() *DocumentSearchRequest`

NewDocumentSearchRequestWithDefaults instantiates a new DocumentSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *DocumentSearchRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *DocumentSearchRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *DocumentSearchRequest) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPerPage

`func (o *DocumentSearchRequest) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *DocumentSearchRequest) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *DocumentSearchRequest) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.


### GetSearchId

`func (o *DocumentSearchRequest) GetSearchId() string`

GetSearchId returns the SearchId field if non-nil, zero value otherwise.

### GetSearchIdOk

`func (o *DocumentSearchRequest) GetSearchIdOk() (*string, bool)`

GetSearchIdOk returns a tuple with the SearchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchId

`func (o *DocumentSearchRequest) SetSearchId(v string)`

SetSearchId sets SearchId field to given value.

### HasSearchId

`func (o *DocumentSearchRequest) HasSearchId() bool`

HasSearchId returns a boolean if a field has been set.

### GetKeepOpen

`func (o *DocumentSearchRequest) GetKeepOpen() bool`

GetKeepOpen returns the KeepOpen field if non-nil, zero value otherwise.

### GetKeepOpenOk

`func (o *DocumentSearchRequest) GetKeepOpenOk() (*bool, bool)`

GetKeepOpenOk returns a tuple with the KeepOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepOpen

`func (o *DocumentSearchRequest) SetKeepOpen(v bool)`

SetKeepOpen sets KeepOpen field to given value.

### HasKeepOpen

`func (o *DocumentSearchRequest) HasKeepOpen() bool`

HasKeepOpen returns a boolean if a field has been set.

### GetQuery

`func (o *DocumentSearchRequest) GetQuery() Query`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *DocumentSearchRequest) GetQueryOk() (*Query, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *DocumentSearchRequest) SetQuery(v Query)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *DocumentSearchRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



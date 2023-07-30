# DocumentSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **int32** |  | 
**PerPage** | **int32** |  | 
**SearchId** | **string** |  | 
**Revisions** | [**[]DocumentAtRevision**](DocumentAtRevision.md) |  | 

## Methods

### NewDocumentSearchResponse

`func NewDocumentSearchResponse(page int32, perPage int32, searchId string, revisions []DocumentAtRevision, ) *DocumentSearchResponse`

NewDocumentSearchResponse instantiates a new DocumentSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentSearchResponseWithDefaults

`func NewDocumentSearchResponseWithDefaults() *DocumentSearchResponse`

NewDocumentSearchResponseWithDefaults instantiates a new DocumentSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *DocumentSearchResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *DocumentSearchResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *DocumentSearchResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPerPage

`func (o *DocumentSearchResponse) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *DocumentSearchResponse) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *DocumentSearchResponse) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.


### GetSearchId

`func (o *DocumentSearchResponse) GetSearchId() string`

GetSearchId returns the SearchId field if non-nil, zero value otherwise.

### GetSearchIdOk

`func (o *DocumentSearchResponse) GetSearchIdOk() (*string, bool)`

GetSearchIdOk returns a tuple with the SearchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchId

`func (o *DocumentSearchResponse) SetSearchId(v string)`

SetSearchId sets SearchId field to given value.


### GetRevisions

`func (o *DocumentSearchResponse) GetRevisions() []DocumentAtRevision`

GetRevisions returns the Revisions field if non-nil, zero value otherwise.

### GetRevisionsOk

`func (o *DocumentSearchResponse) GetRevisionsOk() (*[]DocumentAtRevision, bool)`

GetRevisionsOk returns a tuple with the Revisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisions

`func (o *DocumentSearchResponse) SetRevisions(v []DocumentAtRevision)`

SetRevisions sets Revisions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



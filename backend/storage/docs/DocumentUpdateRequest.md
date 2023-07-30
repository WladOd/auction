# DocumentUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Document** | **map[string]interface{}** |  | 
**Query** | [**Query**](Query.md) |  | 

## Methods

### NewDocumentUpdateRequest

`func NewDocumentUpdateRequest(document map[string]interface{}, query Query, ) *DocumentUpdateRequest`

NewDocumentUpdateRequest instantiates a new DocumentUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentUpdateRequestWithDefaults

`func NewDocumentUpdateRequestWithDefaults() *DocumentUpdateRequest`

NewDocumentUpdateRequestWithDefaults instantiates a new DocumentUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocument

`func (o *DocumentUpdateRequest) GetDocument() map[string]interface{}`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *DocumentUpdateRequest) GetDocumentOk() (*map[string]interface{}, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *DocumentUpdateRequest) SetDocument(v map[string]interface{})`

SetDocument sets Document field to given value.


### GetQuery

`func (o *DocumentUpdateRequest) GetQuery() Query`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *DocumentUpdateRequest) GetQueryOk() (*Query, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *DocumentUpdateRequest) SetQuery(v Query)`

SetQuery sets Query field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



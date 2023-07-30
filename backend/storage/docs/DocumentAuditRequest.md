# DocumentAuditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **int32** |  | 
**PerPage** | **int32** |  | 
**Desc** | **bool** |  | 

## Methods

### NewDocumentAuditRequest

`func NewDocumentAuditRequest(page int32, perPage int32, desc bool, ) *DocumentAuditRequest`

NewDocumentAuditRequest instantiates a new DocumentAuditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentAuditRequestWithDefaults

`func NewDocumentAuditRequestWithDefaults() *DocumentAuditRequest`

NewDocumentAuditRequestWithDefaults instantiates a new DocumentAuditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *DocumentAuditRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *DocumentAuditRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *DocumentAuditRequest) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPerPage

`func (o *DocumentAuditRequest) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *DocumentAuditRequest) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *DocumentAuditRequest) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.


### GetDesc

`func (o *DocumentAuditRequest) GetDesc() bool`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *DocumentAuditRequest) GetDescOk() (*bool, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *DocumentAuditRequest) SetDesc(v bool)`

SetDesc sets Desc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



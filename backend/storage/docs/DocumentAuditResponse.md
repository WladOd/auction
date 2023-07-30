# DocumentAuditResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revisions** | [**[]DocumentAtRevision**](DocumentAtRevision.md) |  | 

## Methods

### NewDocumentAuditResponse

`func NewDocumentAuditResponse(revisions []DocumentAtRevision, ) *DocumentAuditResponse`

NewDocumentAuditResponse instantiates a new DocumentAuditResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentAuditResponseWithDefaults

`func NewDocumentAuditResponseWithDefaults() *DocumentAuditResponse`

NewDocumentAuditResponseWithDefaults instantiates a new DocumentAuditResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisions

`func (o *DocumentAuditResponse) GetRevisions() []DocumentAtRevision`

GetRevisions returns the Revisions field if non-nil, zero value otherwise.

### GetRevisionsOk

`func (o *DocumentAuditResponse) GetRevisionsOk() (*[]DocumentAtRevision, bool)`

GetRevisionsOk returns a tuple with the Revisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisions

`func (o *DocumentAuditResponse) SetRevisions(v []DocumentAtRevision)`

SetRevisions sets Revisions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



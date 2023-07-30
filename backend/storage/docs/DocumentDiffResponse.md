# DocumentDiffResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revisions** | [**[]DocumentAtRevision**](DocumentAtRevision.md) |  | 
**Diffs** | [**[]DocumentDiff**](DocumentDiff.md) |  | 

## Methods

### NewDocumentDiffResponse

`func NewDocumentDiffResponse(revisions []DocumentAtRevision, diffs []DocumentDiff, ) *DocumentDiffResponse`

NewDocumentDiffResponse instantiates a new DocumentDiffResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentDiffResponseWithDefaults

`func NewDocumentDiffResponseWithDefaults() *DocumentDiffResponse`

NewDocumentDiffResponseWithDefaults instantiates a new DocumentDiffResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisions

`func (o *DocumentDiffResponse) GetRevisions() []DocumentAtRevision`

GetRevisions returns the Revisions field if non-nil, zero value otherwise.

### GetRevisionsOk

`func (o *DocumentDiffResponse) GetRevisionsOk() (*[]DocumentAtRevision, bool)`

GetRevisionsOk returns a tuple with the Revisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisions

`func (o *DocumentDiffResponse) SetRevisions(v []DocumentAtRevision)`

SetRevisions sets Revisions field to given value.


### GetDiffs

`func (o *DocumentDiffResponse) GetDiffs() []DocumentDiff`

GetDiffs returns the Diffs field if non-nil, zero value otherwise.

### GetDiffsOk

`func (o *DocumentDiffResponse) GetDiffsOk() (*[]DocumentDiff, bool)`

GetDiffsOk returns a tuple with the Diffs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffs

`func (o *DocumentDiffResponse) SetDiffs(v []DocumentDiff)`

SetDiffs sets Diffs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



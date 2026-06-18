# ActorStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalBuilds** | Pointer to **int32** |  | [optional] 
**TotalRuns** | Pointer to **int32** |  | [optional] 
**TotalUsers** | Pointer to **int32** |  | [optional] 
**TotalUsers7Days** | Pointer to **int32** |  | [optional] 
**TotalUsers30Days** | Pointer to **int32** |  | [optional] 
**TotalUsers90Days** | Pointer to **int32** |  | [optional] 
**TotalMetamorphs** | Pointer to **int32** |  | [optional] 
**LastRunStartedAt** | Pointer to **time.Time** |  | [optional] 
**ActorReviewCount** | Pointer to **int32** |  | [optional] 
**ActorReviewRating** | Pointer to **float32** |  | [optional] 
**BookmarkCount** | Pointer to **int32** |  | [optional] 
**PublicActorRunStats30Days** | Pointer to [**ActorStatsPublicActorRunStats30Days**](ActorStatsPublicActorRunStats30Days.md) |  | [optional] 

## Methods

### NewActorStats

`func NewActorStats() *ActorStats`

NewActorStats instantiates a new ActorStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActorStatsWithDefaults

`func NewActorStatsWithDefaults() *ActorStats`

NewActorStatsWithDefaults instantiates a new ActorStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalBuilds

`func (o *ActorStats) GetTotalBuilds() int32`

GetTotalBuilds returns the TotalBuilds field if non-nil, zero value otherwise.

### GetTotalBuildsOk

`func (o *ActorStats) GetTotalBuildsOk() (*int32, bool)`

GetTotalBuildsOk returns a tuple with the TotalBuilds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBuilds

`func (o *ActorStats) SetTotalBuilds(v int32)`

SetTotalBuilds sets TotalBuilds field to given value.

### HasTotalBuilds

`func (o *ActorStats) HasTotalBuilds() bool`

HasTotalBuilds returns a boolean if a field has been set.

### GetTotalRuns

`func (o *ActorStats) GetTotalRuns() int32`

GetTotalRuns returns the TotalRuns field if non-nil, zero value otherwise.

### GetTotalRunsOk

`func (o *ActorStats) GetTotalRunsOk() (*int32, bool)`

GetTotalRunsOk returns a tuple with the TotalRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRuns

`func (o *ActorStats) SetTotalRuns(v int32)`

SetTotalRuns sets TotalRuns field to given value.

### HasTotalRuns

`func (o *ActorStats) HasTotalRuns() bool`

HasTotalRuns returns a boolean if a field has been set.

### GetTotalUsers

`func (o *ActorStats) GetTotalUsers() int32`

GetTotalUsers returns the TotalUsers field if non-nil, zero value otherwise.

### GetTotalUsersOk

`func (o *ActorStats) GetTotalUsersOk() (*int32, bool)`

GetTotalUsersOk returns a tuple with the TotalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsers

`func (o *ActorStats) SetTotalUsers(v int32)`

SetTotalUsers sets TotalUsers field to given value.

### HasTotalUsers

`func (o *ActorStats) HasTotalUsers() bool`

HasTotalUsers returns a boolean if a field has been set.

### GetTotalUsers7Days

`func (o *ActorStats) GetTotalUsers7Days() int32`

GetTotalUsers7Days returns the TotalUsers7Days field if non-nil, zero value otherwise.

### GetTotalUsers7DaysOk

`func (o *ActorStats) GetTotalUsers7DaysOk() (*int32, bool)`

GetTotalUsers7DaysOk returns a tuple with the TotalUsers7Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsers7Days

`func (o *ActorStats) SetTotalUsers7Days(v int32)`

SetTotalUsers7Days sets TotalUsers7Days field to given value.

### HasTotalUsers7Days

`func (o *ActorStats) HasTotalUsers7Days() bool`

HasTotalUsers7Days returns a boolean if a field has been set.

### GetTotalUsers30Days

`func (o *ActorStats) GetTotalUsers30Days() int32`

GetTotalUsers30Days returns the TotalUsers30Days field if non-nil, zero value otherwise.

### GetTotalUsers30DaysOk

`func (o *ActorStats) GetTotalUsers30DaysOk() (*int32, bool)`

GetTotalUsers30DaysOk returns a tuple with the TotalUsers30Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsers30Days

`func (o *ActorStats) SetTotalUsers30Days(v int32)`

SetTotalUsers30Days sets TotalUsers30Days field to given value.

### HasTotalUsers30Days

`func (o *ActorStats) HasTotalUsers30Days() bool`

HasTotalUsers30Days returns a boolean if a field has been set.

### GetTotalUsers90Days

`func (o *ActorStats) GetTotalUsers90Days() int32`

GetTotalUsers90Days returns the TotalUsers90Days field if non-nil, zero value otherwise.

### GetTotalUsers90DaysOk

`func (o *ActorStats) GetTotalUsers90DaysOk() (*int32, bool)`

GetTotalUsers90DaysOk returns a tuple with the TotalUsers90Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsers90Days

`func (o *ActorStats) SetTotalUsers90Days(v int32)`

SetTotalUsers90Days sets TotalUsers90Days field to given value.

### HasTotalUsers90Days

`func (o *ActorStats) HasTotalUsers90Days() bool`

HasTotalUsers90Days returns a boolean if a field has been set.

### GetTotalMetamorphs

`func (o *ActorStats) GetTotalMetamorphs() int32`

GetTotalMetamorphs returns the TotalMetamorphs field if non-nil, zero value otherwise.

### GetTotalMetamorphsOk

`func (o *ActorStats) GetTotalMetamorphsOk() (*int32, bool)`

GetTotalMetamorphsOk returns a tuple with the TotalMetamorphs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMetamorphs

`func (o *ActorStats) SetTotalMetamorphs(v int32)`

SetTotalMetamorphs sets TotalMetamorphs field to given value.

### HasTotalMetamorphs

`func (o *ActorStats) HasTotalMetamorphs() bool`

HasTotalMetamorphs returns a boolean if a field has been set.

### GetLastRunStartedAt

`func (o *ActorStats) GetLastRunStartedAt() time.Time`

GetLastRunStartedAt returns the LastRunStartedAt field if non-nil, zero value otherwise.

### GetLastRunStartedAtOk

`func (o *ActorStats) GetLastRunStartedAtOk() (*time.Time, bool)`

GetLastRunStartedAtOk returns a tuple with the LastRunStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunStartedAt

`func (o *ActorStats) SetLastRunStartedAt(v time.Time)`

SetLastRunStartedAt sets LastRunStartedAt field to given value.

### HasLastRunStartedAt

`func (o *ActorStats) HasLastRunStartedAt() bool`

HasLastRunStartedAt returns a boolean if a field has been set.

### GetActorReviewCount

`func (o *ActorStats) GetActorReviewCount() int32`

GetActorReviewCount returns the ActorReviewCount field if non-nil, zero value otherwise.

### GetActorReviewCountOk

`func (o *ActorStats) GetActorReviewCountOk() (*int32, bool)`

GetActorReviewCountOk returns a tuple with the ActorReviewCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorReviewCount

`func (o *ActorStats) SetActorReviewCount(v int32)`

SetActorReviewCount sets ActorReviewCount field to given value.

### HasActorReviewCount

`func (o *ActorStats) HasActorReviewCount() bool`

HasActorReviewCount returns a boolean if a field has been set.

### GetActorReviewRating

`func (o *ActorStats) GetActorReviewRating() float32`

GetActorReviewRating returns the ActorReviewRating field if non-nil, zero value otherwise.

### GetActorReviewRatingOk

`func (o *ActorStats) GetActorReviewRatingOk() (*float32, bool)`

GetActorReviewRatingOk returns a tuple with the ActorReviewRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorReviewRating

`func (o *ActorStats) SetActorReviewRating(v float32)`

SetActorReviewRating sets ActorReviewRating field to given value.

### HasActorReviewRating

`func (o *ActorStats) HasActorReviewRating() bool`

HasActorReviewRating returns a boolean if a field has been set.

### GetBookmarkCount

`func (o *ActorStats) GetBookmarkCount() int32`

GetBookmarkCount returns the BookmarkCount field if non-nil, zero value otherwise.

### GetBookmarkCountOk

`func (o *ActorStats) GetBookmarkCountOk() (*int32, bool)`

GetBookmarkCountOk returns a tuple with the BookmarkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookmarkCount

`func (o *ActorStats) SetBookmarkCount(v int32)`

SetBookmarkCount sets BookmarkCount field to given value.

### HasBookmarkCount

`func (o *ActorStats) HasBookmarkCount() bool`

HasBookmarkCount returns a boolean if a field has been set.

### GetPublicActorRunStats30Days

`func (o *ActorStats) GetPublicActorRunStats30Days() ActorStatsPublicActorRunStats30Days`

GetPublicActorRunStats30Days returns the PublicActorRunStats30Days field if non-nil, zero value otherwise.

### GetPublicActorRunStats30DaysOk

`func (o *ActorStats) GetPublicActorRunStats30DaysOk() (*ActorStatsPublicActorRunStats30Days, bool)`

GetPublicActorRunStats30DaysOk returns a tuple with the PublicActorRunStats30Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicActorRunStats30Days

`func (o *ActorStats) SetPublicActorRunStats30Days(v ActorStatsPublicActorRunStats30Days)`

SetPublicActorRunStats30Days sets PublicActorRunStats30Days field to given value.

### HasPublicActorRunStats30Days

`func (o *ActorStats) HasPublicActorRunStats30Days() bool`

HasPublicActorRunStats30Days returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



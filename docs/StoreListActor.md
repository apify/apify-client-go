# StoreListActor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Title** | **string** |  | 
**Name** | **string** |  | 
**Username** | **string** |  | 
**UserFullName** | **string** |  | 
**Description** | **string** |  | 
**Categories** | Pointer to **[]string** |  | [optional] 
**Notice** | Pointer to **string** |  | [optional] 
**PictureUrl** | Pointer to **NullableString** |  | [optional] 
**UserPictureUrl** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Stats** | [**ActorStats**](ActorStats.md) |  | 
**CurrentPricingInfo** | [**CurrentPricingInfo**](CurrentPricingInfo.md) |  | 
**IsWhiteListedForAgenticPayments** | Pointer to **NullableBool** | Whether the Actor is whitelisted for agentic payment processing. | [optional] 
**ActorReviewCount** | Pointer to **int32** |  | [optional] 
**ActorReviewRating** | Pointer to **float32** |  | [optional] 
**BookmarkCount** | Pointer to **int32** |  | [optional] 
**Badge** | Pointer to **NullableString** |  | [optional] 
**ReadmeSummary** | Pointer to **string** | A brief, LLM-generated readme summary | [optional] 

## Methods

### NewStoreListActor

`func NewStoreListActor(id string, title string, name string, username string, userFullName string, description string, stats ActorStats, currentPricingInfo CurrentPricingInfo, ) *StoreListActor`

NewStoreListActor instantiates a new StoreListActor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreListActorWithDefaults

`func NewStoreListActorWithDefaults() *StoreListActor`

NewStoreListActorWithDefaults instantiates a new StoreListActor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StoreListActor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StoreListActor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StoreListActor) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *StoreListActor) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StoreListActor) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StoreListActor) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetName

`func (o *StoreListActor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StoreListActor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StoreListActor) SetName(v string)`

SetName sets Name field to given value.


### GetUsername

`func (o *StoreListActor) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *StoreListActor) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *StoreListActor) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetUserFullName

`func (o *StoreListActor) GetUserFullName() string`

GetUserFullName returns the UserFullName field if non-nil, zero value otherwise.

### GetUserFullNameOk

`func (o *StoreListActor) GetUserFullNameOk() (*string, bool)`

GetUserFullNameOk returns a tuple with the UserFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFullName

`func (o *StoreListActor) SetUserFullName(v string)`

SetUserFullName sets UserFullName field to given value.


### GetDescription

`func (o *StoreListActor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StoreListActor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StoreListActor) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCategories

`func (o *StoreListActor) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *StoreListActor) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *StoreListActor) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *StoreListActor) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetNotice

`func (o *StoreListActor) GetNotice() string`

GetNotice returns the Notice field if non-nil, zero value otherwise.

### GetNoticeOk

`func (o *StoreListActor) GetNoticeOk() (*string, bool)`

GetNoticeOk returns a tuple with the Notice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotice

`func (o *StoreListActor) SetNotice(v string)`

SetNotice sets Notice field to given value.

### HasNotice

`func (o *StoreListActor) HasNotice() bool`

HasNotice returns a boolean if a field has been set.

### GetPictureUrl

`func (o *StoreListActor) GetPictureUrl() string`

GetPictureUrl returns the PictureUrl field if non-nil, zero value otherwise.

### GetPictureUrlOk

`func (o *StoreListActor) GetPictureUrlOk() (*string, bool)`

GetPictureUrlOk returns a tuple with the PictureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPictureUrl

`func (o *StoreListActor) SetPictureUrl(v string)`

SetPictureUrl sets PictureUrl field to given value.

### HasPictureUrl

`func (o *StoreListActor) HasPictureUrl() bool`

HasPictureUrl returns a boolean if a field has been set.

### SetPictureUrlNil

`func (o *StoreListActor) SetPictureUrlNil(b bool)`

 SetPictureUrlNil sets the value for PictureUrl to be an explicit nil

### UnsetPictureUrl
`func (o *StoreListActor) UnsetPictureUrl()`

UnsetPictureUrl ensures that no value is present for PictureUrl, not even an explicit nil
### GetUserPictureUrl

`func (o *StoreListActor) GetUserPictureUrl() string`

GetUserPictureUrl returns the UserPictureUrl field if non-nil, zero value otherwise.

### GetUserPictureUrlOk

`func (o *StoreListActor) GetUserPictureUrlOk() (*string, bool)`

GetUserPictureUrlOk returns a tuple with the UserPictureUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPictureUrl

`func (o *StoreListActor) SetUserPictureUrl(v string)`

SetUserPictureUrl sets UserPictureUrl field to given value.

### HasUserPictureUrl

`func (o *StoreListActor) HasUserPictureUrl() bool`

HasUserPictureUrl returns a boolean if a field has been set.

### SetUserPictureUrlNil

`func (o *StoreListActor) SetUserPictureUrlNil(b bool)`

 SetUserPictureUrlNil sets the value for UserPictureUrl to be an explicit nil

### UnsetUserPictureUrl
`func (o *StoreListActor) UnsetUserPictureUrl()`

UnsetUserPictureUrl ensures that no value is present for UserPictureUrl, not even an explicit nil
### GetUrl

`func (o *StoreListActor) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StoreListActor) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StoreListActor) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *StoreListActor) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *StoreListActor) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *StoreListActor) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetStats

`func (o *StoreListActor) GetStats() ActorStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *StoreListActor) GetStatsOk() (*ActorStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *StoreListActor) SetStats(v ActorStats)`

SetStats sets Stats field to given value.


### GetCurrentPricingInfo

`func (o *StoreListActor) GetCurrentPricingInfo() CurrentPricingInfo`

GetCurrentPricingInfo returns the CurrentPricingInfo field if non-nil, zero value otherwise.

### GetCurrentPricingInfoOk

`func (o *StoreListActor) GetCurrentPricingInfoOk() (*CurrentPricingInfo, bool)`

GetCurrentPricingInfoOk returns a tuple with the CurrentPricingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPricingInfo

`func (o *StoreListActor) SetCurrentPricingInfo(v CurrentPricingInfo)`

SetCurrentPricingInfo sets CurrentPricingInfo field to given value.


### GetIsWhiteListedForAgenticPayments

`func (o *StoreListActor) GetIsWhiteListedForAgenticPayments() bool`

GetIsWhiteListedForAgenticPayments returns the IsWhiteListedForAgenticPayments field if non-nil, zero value otherwise.

### GetIsWhiteListedForAgenticPaymentsOk

`func (o *StoreListActor) GetIsWhiteListedForAgenticPaymentsOk() (*bool, bool)`

GetIsWhiteListedForAgenticPaymentsOk returns a tuple with the IsWhiteListedForAgenticPayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWhiteListedForAgenticPayments

`func (o *StoreListActor) SetIsWhiteListedForAgenticPayments(v bool)`

SetIsWhiteListedForAgenticPayments sets IsWhiteListedForAgenticPayments field to given value.

### HasIsWhiteListedForAgenticPayments

`func (o *StoreListActor) HasIsWhiteListedForAgenticPayments() bool`

HasIsWhiteListedForAgenticPayments returns a boolean if a field has been set.

### SetIsWhiteListedForAgenticPaymentsNil

`func (o *StoreListActor) SetIsWhiteListedForAgenticPaymentsNil(b bool)`

 SetIsWhiteListedForAgenticPaymentsNil sets the value for IsWhiteListedForAgenticPayments to be an explicit nil

### UnsetIsWhiteListedForAgenticPayments
`func (o *StoreListActor) UnsetIsWhiteListedForAgenticPayments()`

UnsetIsWhiteListedForAgenticPayments ensures that no value is present for IsWhiteListedForAgenticPayments, not even an explicit nil
### GetActorReviewCount

`func (o *StoreListActor) GetActorReviewCount() int32`

GetActorReviewCount returns the ActorReviewCount field if non-nil, zero value otherwise.

### GetActorReviewCountOk

`func (o *StoreListActor) GetActorReviewCountOk() (*int32, bool)`

GetActorReviewCountOk returns a tuple with the ActorReviewCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorReviewCount

`func (o *StoreListActor) SetActorReviewCount(v int32)`

SetActorReviewCount sets ActorReviewCount field to given value.

### HasActorReviewCount

`func (o *StoreListActor) HasActorReviewCount() bool`

HasActorReviewCount returns a boolean if a field has been set.

### GetActorReviewRating

`func (o *StoreListActor) GetActorReviewRating() float32`

GetActorReviewRating returns the ActorReviewRating field if non-nil, zero value otherwise.

### GetActorReviewRatingOk

`func (o *StoreListActor) GetActorReviewRatingOk() (*float32, bool)`

GetActorReviewRatingOk returns a tuple with the ActorReviewRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorReviewRating

`func (o *StoreListActor) SetActorReviewRating(v float32)`

SetActorReviewRating sets ActorReviewRating field to given value.

### HasActorReviewRating

`func (o *StoreListActor) HasActorReviewRating() bool`

HasActorReviewRating returns a boolean if a field has been set.

### GetBookmarkCount

`func (o *StoreListActor) GetBookmarkCount() int32`

GetBookmarkCount returns the BookmarkCount field if non-nil, zero value otherwise.

### GetBookmarkCountOk

`func (o *StoreListActor) GetBookmarkCountOk() (*int32, bool)`

GetBookmarkCountOk returns a tuple with the BookmarkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookmarkCount

`func (o *StoreListActor) SetBookmarkCount(v int32)`

SetBookmarkCount sets BookmarkCount field to given value.

### HasBookmarkCount

`func (o *StoreListActor) HasBookmarkCount() bool`

HasBookmarkCount returns a boolean if a field has been set.

### GetBadge

`func (o *StoreListActor) GetBadge() string`

GetBadge returns the Badge field if non-nil, zero value otherwise.

### GetBadgeOk

`func (o *StoreListActor) GetBadgeOk() (*string, bool)`

GetBadgeOk returns a tuple with the Badge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadge

`func (o *StoreListActor) SetBadge(v string)`

SetBadge sets Badge field to given value.

### HasBadge

`func (o *StoreListActor) HasBadge() bool`

HasBadge returns a boolean if a field has been set.

### SetBadgeNil

`func (o *StoreListActor) SetBadgeNil(b bool)`

 SetBadgeNil sets the value for Badge to be an explicit nil

### UnsetBadge
`func (o *StoreListActor) UnsetBadge()`

UnsetBadge ensures that no value is present for Badge, not even an explicit nil
### GetReadmeSummary

`func (o *StoreListActor) GetReadmeSummary() string`

GetReadmeSummary returns the ReadmeSummary field if non-nil, zero value otherwise.

### GetReadmeSummaryOk

`func (o *StoreListActor) GetReadmeSummaryOk() (*string, bool)`

GetReadmeSummaryOk returns a tuple with the ReadmeSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadmeSummary

`func (o *StoreListActor) SetReadmeSummary(v string)`

SetReadmeSummary sets ReadmeSummary field to given value.

### HasReadmeSummary

`func (o *StoreListActor) HasReadmeSummary() bool`

HasReadmeSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



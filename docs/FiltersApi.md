# \FiltersApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListFilterValues**](FiltersApi.md#ListFilterValues) | **Get** /data/v1/filters/{FILTER_ID} | Lists values for a specific filter
[**ListFilters**](FiltersApi.md#ListFilters) | **Get** /data/v1/filters | List Filters


# **ListFilterValues**
> ListFilterValuesResponse ListFilterValues(ctx, fILTERID, optional)
Lists values for a specific filter

Deprecated: The API has been replaced by the list-dimension-values API call.  Lists the values for a filter along with a total count of related views. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fILTERID** | **string**| ID of the Filter | 
 **optional** | ***ListFilterValuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListFilterValuesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Number of items to include in the response | [default to 25]
 **page** | **optional.Int32**| Offset by this many pages, of the size of &#x60;limit&#x60; | [default to 1]
 **filters** | [**optional.Interface of []string**](string.md)| Filter key:value pairs. Must be provided as an array query string parameter (e.g. filters[]&#x3D;operating_system:windows&amp;filters[]&#x3D;country:US). Possible filter names are the same as returned by the List Filters endpoint.  | 
 **timeframe** | [**optional.Interface of []string**](string.md)| Timeframe window to limit results by. Must be provided as an array query string parameter (e.g. timeframe[]&#x3D;). Accepted formats are...   * array of epoch timestamps e.g. timeframe[]&#x3D;1498867200&amp;timeframe[]&#x3D;1498953600   * duration string e.g. timeframe[]&#x3D;24:hours or timeframe[]&#x3D;7:days.  | 

### Return type

[**ListFilterValuesResponse**](ListFilterValuesResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFilters**
> ListFiltersResponse ListFilters(ctx, )
List Filters

Deprecated: The API has been replaced by the list-dimensions API call.  Lists all the filters broken out into basic and advanced. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListFiltersResponse**](ListFiltersResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


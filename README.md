# Go API client for muxgo

Mux is how developers build online video. This API encompasses both
Mux Video and Mux Data functionality to help you build your video-related
projects better and faster than ever before.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 0.13.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./muxgo"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.mux.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AssetsApi* | [**CreateAsset**](docs/AssetsApi.md#createasset) | **Post** /video/v1/assets | Create an asset
*AssetsApi* | [**CreateAssetPlaybackId**](docs/AssetsApi.md#createassetplaybackid) | **Post** /video/v1/assets/{ASSET_ID}/playback-ids | Create a playback ID
*AssetsApi* | [**CreateAssetTrack**](docs/AssetsApi.md#createassettrack) | **Post** /video/v1/assets/{ASSET_ID}/tracks | Create an asset track
*AssetsApi* | [**DeleteAsset**](docs/AssetsApi.md#deleteasset) | **Delete** /video/v1/assets/{ASSET_ID} | Delete an asset
*AssetsApi* | [**DeleteAssetPlaybackId**](docs/AssetsApi.md#deleteassetplaybackid) | **Delete** /video/v1/assets/{ASSET_ID}/playback-ids/{PLAYBACK_ID} | Delete a playback ID
*AssetsApi* | [**DeleteAssetTrack**](docs/AssetsApi.md#deleteassettrack) | **Delete** /video/v1/assets/{ASSET_ID}/tracks/{TRACK_ID} | Delete an asset track
*AssetsApi* | [**GetAsset**](docs/AssetsApi.md#getasset) | **Get** /video/v1/assets/{ASSET_ID} | Retrieve an asset
*AssetsApi* | [**GetAssetInputInfo**](docs/AssetsApi.md#getassetinputinfo) | **Get** /video/v1/assets/{ASSET_ID}/input-info | Retrieve asset input info
*AssetsApi* | [**GetAssetPlaybackId**](docs/AssetsApi.md#getassetplaybackid) | **Get** /video/v1/assets/{ASSET_ID}/playback-ids/{PLAYBACK_ID} | Retrieve a playback ID
*AssetsApi* | [**ListAssets**](docs/AssetsApi.md#listassets) | **Get** /video/v1/assets | List assets
*AssetsApi* | [**UpdateAssetMasterAccess**](docs/AssetsApi.md#updateassetmasteraccess) | **Put** /video/v1/assets/{ASSET_ID}/master-access | Update master access
*AssetsApi* | [**UpdateAssetMp4Support**](docs/AssetsApi.md#updateassetmp4support) | **Put** /video/v1/assets/{ASSET_ID}/mp4-support | Update MP4 support
*DeliveryUsageApi* | [**ListDeliveryUsage**](docs/DeliveryUsageApi.md#listdeliveryusage) | **Get** /video/v1/delivery-usage | List Usage
*DimensionsApi* | [**ListDimensionValues**](docs/DimensionsApi.md#listdimensionvalues) | **Get** /data/v1/dimensions/{DIMENSION_ID} | Lists the values for a specific dimension
*DimensionsApi* | [**ListDimensions**](docs/DimensionsApi.md#listdimensions) | **Get** /data/v1/dimensions | List Dimensions
*DirectUploadsApi* | [**CancelDirectUpload**](docs/DirectUploadsApi.md#canceldirectupload) | **Put** /video/v1/uploads/{UPLOAD_ID}/cancel | Cancel a direct upload
*DirectUploadsApi* | [**CreateDirectUpload**](docs/DirectUploadsApi.md#createdirectupload) | **Post** /video/v1/uploads | Create a new direct upload URL
*DirectUploadsApi* | [**GetDirectUpload**](docs/DirectUploadsApi.md#getdirectupload) | **Get** /video/v1/uploads/{UPLOAD_ID} | Retrieve a single direct upload&#39;s info
*DirectUploadsApi* | [**ListDirectUploads**](docs/DirectUploadsApi.md#listdirectuploads) | **Get** /video/v1/uploads | List direct uploads
*ErrorsApi* | [**ListErrors**](docs/ErrorsApi.md#listerrors) | **Get** /data/v1/errors | List Errors
*ExportsApi* | [**ListExports**](docs/ExportsApi.md#listexports) | **Get** /data/v1/exports | List property video view export links
*FiltersApi* | [**ListFilterValues**](docs/FiltersApi.md#listfiltervalues) | **Get** /data/v1/filters/{FILTER_ID} | Lists values for a specific filter
*FiltersApi* | [**ListFilters**](docs/FiltersApi.md#listfilters) | **Get** /data/v1/filters | List Filters
*IncidentsApi* | [**GetIncident**](docs/IncidentsApi.md#getincident) | **Get** /data/v1/incidents/{INCIDENT_ID} | Get an Incident
*IncidentsApi* | [**ListIncidents**](docs/IncidentsApi.md#listincidents) | **Get** /data/v1/incidents | List Incidents
*IncidentsApi* | [**ListRelatedIncidents**](docs/IncidentsApi.md#listrelatedincidents) | **Get** /data/v1/incidents/{INCIDENT_ID}/related | List Related Incidents
*LiveStreamsApi* | [**CreateLiveStream**](docs/LiveStreamsApi.md#createlivestream) | **Post** /video/v1/live-streams | Create a live stream
*LiveStreamsApi* | [**CreateLiveStreamPlaybackId**](docs/LiveStreamsApi.md#createlivestreamplaybackid) | **Post** /video/v1/live-streams/{LIVE_STREAM_ID}/playback-ids | Create a live stream playback ID
*LiveStreamsApi* | [**CreateLiveStreamSimulcastTarget**](docs/LiveStreamsApi.md#createlivestreamsimulcasttarget) | **Post** /video/v1/live-streams/{LIVE_STREAM_ID}/simulcast-targets | Create a live stream simulcast target
*LiveStreamsApi* | [**DeleteLiveStream**](docs/LiveStreamsApi.md#deletelivestream) | **Delete** /video/v1/live-streams/{LIVE_STREAM_ID} | Delete a live stream
*LiveStreamsApi* | [**DeleteLiveStreamPlaybackId**](docs/LiveStreamsApi.md#deletelivestreamplaybackid) | **Delete** /video/v1/live-streams/{LIVE_STREAM_ID}/playback-ids/{PLAYBACK_ID} | Delete a live stream playback ID
*LiveStreamsApi* | [**DeleteLiveStreamSimulcastTarget**](docs/LiveStreamsApi.md#deletelivestreamsimulcasttarget) | **Delete** /video/v1/live-streams/{LIVE_STREAM_ID}/simulcast-targets/{SIMULCAST_TARGET_ID} | Delete a Live Stream Simulcast Target
*LiveStreamsApi* | [**DisableLiveStream**](docs/LiveStreamsApi.md#disablelivestream) | **Put** /video/v1/live-streams/{LIVE_STREAM_ID}/disable | Disable a live stream
*LiveStreamsApi* | [**EnableLiveStream**](docs/LiveStreamsApi.md#enablelivestream) | **Put** /video/v1/live-streams/{LIVE_STREAM_ID}/enable | Enable a live stream
*LiveStreamsApi* | [**GetLiveStream**](docs/LiveStreamsApi.md#getlivestream) | **Get** /video/v1/live-streams/{LIVE_STREAM_ID} | Retrieve a live stream
*LiveStreamsApi* | [**GetLiveStreamSimulcastTarget**](docs/LiveStreamsApi.md#getlivestreamsimulcasttarget) | **Get** /video/v1/live-streams/{LIVE_STREAM_ID}/simulcast-targets/{SIMULCAST_TARGET_ID} | Retrieve a Live Stream Simulcast Target
*LiveStreamsApi* | [**ListLiveStreams**](docs/LiveStreamsApi.md#listlivestreams) | **Get** /video/v1/live-streams | List live streams
*LiveStreamsApi* | [**ResetStreamKey**](docs/LiveStreamsApi.md#resetstreamkey) | **Post** /video/v1/live-streams/{LIVE_STREAM_ID}/reset-stream-key | Reset a live stream’s stream key
*LiveStreamsApi* | [**SignalLiveStreamComplete**](docs/LiveStreamsApi.md#signallivestreamcomplete) | **Put** /video/v1/live-streams/{LIVE_STREAM_ID}/complete | Signal a live stream is finished
*MetricsApi* | [**GetMetricTimeseriesData**](docs/MetricsApi.md#getmetrictimeseriesdata) | **Get** /data/v1/metrics/{METRIC_ID}/timeseries | Get metric timeseries data
*MetricsApi* | [**GetOverallValues**](docs/MetricsApi.md#getoverallvalues) | **Get** /data/v1/metrics/{METRIC_ID}/overall | Get Overall values
*MetricsApi* | [**ListAllMetricValues**](docs/MetricsApi.md#listallmetricvalues) | **Get** /data/v1/metrics/comparison | List all metric values
*MetricsApi* | [**ListBreakdownValues**](docs/MetricsApi.md#listbreakdownvalues) | **Get** /data/v1/metrics/{METRIC_ID}/breakdown | List breakdown values
*MetricsApi* | [**ListInsights**](docs/MetricsApi.md#listinsights) | **Get** /data/v1/metrics/{METRIC_ID}/insights | List Insights
*PlaybackIDApi* | [**GetAssetOrLivestreamId**](docs/PlaybackIDApi.md#getassetorlivestreamid) | **Get** /video/v1/playback-ids/{PLAYBACK_ID} | Retrieve an Asset or Live Stream ID
*RealTimeApi* | [**GetRealtimeBreakdown**](docs/RealTimeApi.md#getrealtimebreakdown) | **Get** /data/v1/realtime/metrics/{REALTIME_METRIC_ID}/breakdown | Get Real-Time Breakdown
*RealTimeApi* | [**GetRealtimeHistogramTimeseries**](docs/RealTimeApi.md#getrealtimehistogramtimeseries) | **Get** /data/v1/realtime/metrics/{REALTIME_METRIC_ID}/histogram-timeseries | Get Real-Time Histogram Timeseries
*RealTimeApi* | [**GetRealtimeTimeseries**](docs/RealTimeApi.md#getrealtimetimeseries) | **Get** /data/v1/realtime/metrics/{REALTIME_METRIC_ID}/timeseries | Get Real-Time Timeseries
*RealTimeApi* | [**ListRealtimeDimensions**](docs/RealTimeApi.md#listrealtimedimensions) | **Get** /data/v1/realtime/dimensions | List Real-Time Dimensions
*RealTimeApi* | [**ListRealtimeMetrics**](docs/RealTimeApi.md#listrealtimemetrics) | **Get** /data/v1/realtime/metrics | List Real-Time Metrics
*URLSigningKeysApi* | [**CreateUrlSigningKey**](docs/URLSigningKeysApi.md#createurlsigningkey) | **Post** /video/v1/signing-keys | Create a URL signing key
*URLSigningKeysApi* | [**DeleteUrlSigningKey**](docs/URLSigningKeysApi.md#deleteurlsigningkey) | **Delete** /video/v1/signing-keys/{SIGNING_KEY_ID} | Delete a URL signing key
*URLSigningKeysApi* | [**GetUrlSigningKey**](docs/URLSigningKeysApi.md#geturlsigningkey) | **Get** /video/v1/signing-keys/{SIGNING_KEY_ID} | Retrieve a URL signing key
*URLSigningKeysApi* | [**ListUrlSigningKeys**](docs/URLSigningKeysApi.md#listurlsigningkeys) | **Get** /video/v1/signing-keys | List URL signing keys
*VideoViewsApi* | [**GetVideoView**](docs/VideoViewsApi.md#getvideoview) | **Get** /data/v1/video-views/{VIDEO_VIEW_ID} | Get a Video View
*VideoViewsApi* | [**ListVideoViews**](docs/VideoViewsApi.md#listvideoviews) | **Get** /data/v1/video-views | List Video Views


## Documentation For Models

 - [AbridgedVideoView](docs/AbridgedVideoView.md)
 - [Asset](docs/Asset.md)
 - [AssetErrors](docs/AssetErrors.md)
 - [AssetMaster](docs/AssetMaster.md)
 - [AssetNonStandardInputReasons](docs/AssetNonStandardInputReasons.md)
 - [AssetRecordingTimes](docs/AssetRecordingTimes.md)
 - [AssetResponse](docs/AssetResponse.md)
 - [AssetStaticRenditions](docs/AssetStaticRenditions.md)
 - [AssetStaticRenditionsFiles](docs/AssetStaticRenditionsFiles.md)
 - [BreakdownValue](docs/BreakdownValue.md)
 - [CreateAssetRequest](docs/CreateAssetRequest.md)
 - [CreateLiveStreamRequest](docs/CreateLiveStreamRequest.md)
 - [CreatePlaybackIDRequest](docs/CreatePlaybackIDRequest.md)
 - [CreatePlaybackIDResponse](docs/CreatePlaybackIDResponse.md)
 - [CreateSimulcastTargetRequest](docs/CreateSimulcastTargetRequest.md)
 - [CreateTrackRequest](docs/CreateTrackRequest.md)
 - [CreateTrackResponse](docs/CreateTrackResponse.md)
 - [CreateUploadRequest](docs/CreateUploadRequest.md)
 - [DeliveryReport](docs/DeliveryReport.md)
 - [DimensionValue](docs/DimensionValue.md)
 - [DisableLiveStreamResponse](docs/DisableLiveStreamResponse.md)
 - [EnableLiveStreamResponse](docs/EnableLiveStreamResponse.md)
 - [Error](docs/Error.md)
 - [FilterValue](docs/FilterValue.md)
 - [GetAssetInputInfoResponse](docs/GetAssetInputInfoResponse.md)
 - [GetAssetOrLiveStreamIdResponse](docs/GetAssetOrLiveStreamIdResponse.md)
 - [GetAssetOrLiveStreamIdResponseData](docs/GetAssetOrLiveStreamIdResponseData.md)
 - [GetAssetOrLiveStreamIdResponseDataObject](docs/GetAssetOrLiveStreamIdResponseDataObject.md)
 - [GetAssetPlaybackIDResponse](docs/GetAssetPlaybackIDResponse.md)
 - [GetMetricTimeseriesDataResponse](docs/GetMetricTimeseriesDataResponse.md)
 - [GetOverallValuesResponse](docs/GetOverallValuesResponse.md)
 - [GetRealTimeBreakdownResponse](docs/GetRealTimeBreakdownResponse.md)
 - [GetRealTimeHistogramTimeseriesResponse](docs/GetRealTimeHistogramTimeseriesResponse.md)
 - [GetRealTimeHistogramTimeseriesResponseMeta](docs/GetRealTimeHistogramTimeseriesResponseMeta.md)
 - [GetRealTimeTimeseriesResponse](docs/GetRealTimeTimeseriesResponse.md)
 - [Incident](docs/Incident.md)
 - [IncidentBreakdown](docs/IncidentBreakdown.md)
 - [IncidentNotification](docs/IncidentNotification.md)
 - [IncidentNotificationRule](docs/IncidentNotificationRule.md)
 - [IncidentResponse](docs/IncidentResponse.md)
 - [InputFile](docs/InputFile.md)
 - [InputInfo](docs/InputInfo.md)
 - [InputSettings](docs/InputSettings.md)
 - [InputSettingsOverlaySettings](docs/InputSettingsOverlaySettings.md)
 - [InputTrack](docs/InputTrack.md)
 - [Insight](docs/Insight.md)
 - [ListAllMetricValuesResponse](docs/ListAllMetricValuesResponse.md)
 - [ListAssetsResponse](docs/ListAssetsResponse.md)
 - [ListBreakdownValuesResponse](docs/ListBreakdownValuesResponse.md)
 - [ListDeliveryUsageResponse](docs/ListDeliveryUsageResponse.md)
 - [ListDimensionValuesResponse](docs/ListDimensionValuesResponse.md)
 - [ListDimensionsResponse](docs/ListDimensionsResponse.md)
 - [ListErrorsResponse](docs/ListErrorsResponse.md)
 - [ListExportsResponse](docs/ListExportsResponse.md)
 - [ListFilterValuesResponse](docs/ListFilterValuesResponse.md)
 - [ListFiltersResponse](docs/ListFiltersResponse.md)
 - [ListFiltersResponseData](docs/ListFiltersResponseData.md)
 - [ListIncidentsResponse](docs/ListIncidentsResponse.md)
 - [ListInsightsResponse](docs/ListInsightsResponse.md)
 - [ListLiveStreamsResponse](docs/ListLiveStreamsResponse.md)
 - [ListRealTimeDimensionsResponse](docs/ListRealTimeDimensionsResponse.md)
 - [ListRealTimeDimensionsResponseData](docs/ListRealTimeDimensionsResponseData.md)
 - [ListRealTimeMetricsResponse](docs/ListRealTimeMetricsResponse.md)
 - [ListRelatedIncidentsResponse](docs/ListRelatedIncidentsResponse.md)
 - [ListSigningKeysResponse](docs/ListSigningKeysResponse.md)
 - [ListUploadsResponse](docs/ListUploadsResponse.md)
 - [ListVideoViewsResponse](docs/ListVideoViewsResponse.md)
 - [LiveStream](docs/LiveStream.md)
 - [LiveStreamResponse](docs/LiveStreamResponse.md)
 - [Metric](docs/Metric.md)
 - [NotificationRule](docs/NotificationRule.md)
 - [OverallValues](docs/OverallValues.md)
 - [PlaybackID](docs/PlaybackID.md)
 - [PlaybackPolicy](docs/PlaybackPolicy.md)
 - [RealTimeBreakdownValue](docs/RealTimeBreakdownValue.md)
 - [RealTimeHistogramTimeseriesBucket](docs/RealTimeHistogramTimeseriesBucket.md)
 - [RealTimeHistogramTimeseriesBucketValues](docs/RealTimeHistogramTimeseriesBucketValues.md)
 - [RealTimeHistogramTimeseriesDatapoint](docs/RealTimeHistogramTimeseriesDatapoint.md)
 - [RealTimeTimeseriesDatapoint](docs/RealTimeTimeseriesDatapoint.md)
 - [Score](docs/Score.md)
 - [SignalLiveStreamCompleteResponse](docs/SignalLiveStreamCompleteResponse.md)
 - [SigningKey](docs/SigningKey.md)
 - [SigningKeyResponse](docs/SigningKeyResponse.md)
 - [SimulcastTarget](docs/SimulcastTarget.md)
 - [SimulcastTargetResponse](docs/SimulcastTargetResponse.md)
 - [Track](docs/Track.md)
 - [UpdateAssetMP4SupportRequest](docs/UpdateAssetMP4SupportRequest.md)
 - [UpdateAssetMasterAccessRequest](docs/UpdateAssetMasterAccessRequest.md)
 - [Upload](docs/Upload.md)
 - [UploadError](docs/UploadError.md)
 - [UploadResponse](docs/UploadResponse.md)
 - [VideoView](docs/VideoView.md)
 - [VideoViewEvent](docs/VideoViewEvent.md)
 - [VideoViewResponse](docs/VideoViewResponse.md)


## Documentation For Authorization



### accessToken

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author




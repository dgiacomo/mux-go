/*
 * Mux API
 *
 * Mux is how developers build online video. This API encompasses both Mux Video and Mux Data functionality to help you build your video-related projects better and faster than ever before. 
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package muxgo

import (
	"encoding/json"
)

// LiveStream struct for LiveStream
type LiveStream struct {
	// Unique identifier for the Live Stream. Max 255 characters.
	Id *string `json:"id,omitempty"`
	// Time the Live Stream was created, defined as a Unix timestamp (seconds since epoch).
	CreatedAt *string `json:"created_at,omitempty"`
	// Unique key used for streaming to a Mux RTMP endpoint. This should be considered as sensitive as credentials, anyone with this stream key can begin streaming.
	StreamKey *string `json:"stream_key,omitempty"`
	// The Asset that is currently being created if there is an active broadcast.
	ActiveAssetId *string `json:"active_asset_id,omitempty"`
	// An array of strings with the most recent Assets that were created from this live stream.
	RecentAssetIds *[]string `json:"recent_asset_ids,omitempty"`
	// `idle` indicates that there is no active broadcast. `active` indicates that there is an active broadcast and `disabled` status indicates that no future RTMP streams can be published.
	Status *string `json:"status,omitempty"`
	// An array of Playback ID objects. Use these to create HLS playback URLs. See [Play your videos](https://docs.mux.com/guides/video/play-your-videos) for more details.
	PlaybackIds *[]PlaybackID `json:"playback_ids,omitempty"`
	NewAssetSettings *CreateAssetRequest `json:"new_asset_settings,omitempty"`
	// Arbitrary metadata set for the asset. Max 255 characters.
	Passthrough *string `json:"passthrough,omitempty"`
	// When live streaming software disconnects from Mux, either intentionally or due to a drop in the network, the Reconnect Window is the time in seconds that Mux should wait for the streaming software to reconnect before considering the live stream finished and completing the recorded asset. **Min**: 0.1s. **Max**: 300s (5 minutes).
	ReconnectWindow *float32 `json:"reconnect_window,omitempty"`
	// Latency is the time from when the streamer does something in real life to when you see it happen in the player. Set this if you want lower latency for your live stream. **Note**: Reconnect windows are incompatible with Reduced Latency and will always be set to zero (0) seconds. See the [Reduce live stream latency guide](https://docs.mux.com/guides/video/reduce-live-stream-latency) to understand the tradeoffs.
	ReducedLatency *bool `json:"reduced_latency,omitempty"`
	// Each Simulcast Target contains configuration details to broadcast (or \"restream\") a live stream to a third-party streaming service. [See the Stream live to 3rd party platforms guide](https://docs.mux.com/guides/video/stream-live-to-3rd-party-platforms).
	SimulcastTargets *[]SimulcastTarget `json:"simulcast_targets,omitempty"`
	// True means this live stream is a test live stream. Test live streams can be used to help evaluate the Mux Video APIs for free. There is no limit on the number of test live streams, but they are watermarked with the Mux logo, limited to 5 minutes, and deleted after 24 hours.
	Test *bool `json:"test,omitempty"`
}

// NewLiveStream instantiates a new LiveStream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveStream() *LiveStream {
	this := LiveStream{}
	var reconnectWindow float32 = 60
	this.ReconnectWindow = &reconnectWindow
	return &this
}

// NewLiveStreamWithDefaults instantiates a new LiveStream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveStreamWithDefaults() *LiveStream {
	this := LiveStream{}
	var reconnectWindow float32 = 60
	this.ReconnectWindow = &reconnectWindow
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LiveStream) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStream) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LiveStream) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LiveStream) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *LiveStream) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStream) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LiveStream) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *LiveStream) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetStreamKey returns the StreamKey field value if set, zero value otherwise.
func (o *LiveStream) GetStreamKey() string {
	if o == nil || o.StreamKey == nil {
		var ret string
		return ret
	}
	return *o.StreamKey
}

// GetStreamKeyOk returns a tuple with the StreamKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStream) GetStreamKeyOk() (*string, bool) {
	if o == nil || o.StreamKey == nil {
		return nil, false
	}
	return o.StreamKey, true
}

// HasStreamKey returns a boolean if a field has been set.
func (o *LiveStream) HasStreamKey() bool {
	if o != nil && o.StreamKey != nil {
		return true
	}

	return false
}

// SetStreamKey gets a reference to the given string and assigns it to the StreamKey field.
func (o *LiveStream) SetStreamKey(v string) {
	o.StreamKey = &v
}

// GetActiveAssetId returns the ActiveAssetId field value if set, zero value otherwise.
func (o *LiveStream) GetActiveAssetId() string {
	if o == nil || o.ActiveAssetId == nil {
		var ret string
		return ret
	}
	return *o.ActiveAssetId
}

// GetActiveAssetIdOk returns a tuple with the ActiveAssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStream) GetActiveAssetIdOk() (*string, bool) {
	if o == nil || o.ActiveAssetId == nil {
		return nil, false
	}
	return o.ActiveAssetId, true
}

// HasActiveAssetId returns a boolean if a field has been set.
func (o *LiveStream) HasActiveAssetId() bool {
	if o != nil && o.ActiveAssetId != nil {
		return true
	}

	return false
}

// SetActiveAssetId gets a reference to the given string and assigns it to the ActiveAssetId field.
func (o *LiveStream) SetActiveAssetId(v string) {
	o.ActiveAssetId = &v
}

// GetRecentAssetIds returns the RecentAssetIds field value if set, zero value otherwise.
func (o *LiveStream) GetRecentAssetIds() []string {
	if o == nil || o.RecentAssetIds == nil {
		var ret []string
		return ret
	}
	return *o.RecentAssetIds
}

// GetRecentAssetIdsOk returns a tuple with the RecentAssetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStream) GetRecentAssetIdsOk() (*[]string, bool) {
	if o == nil || o.RecentAssetIds == nil {
		return nil, false
	}
	return o.RecentAssetIds, true
}

// HasRecentAssetIds returns a boolean if a field has been set.
func (o *LiveStream) HasRecentAssetIds() bool {
	if o != nil && o.RecentAssetIds != nil {
		return true
	}

	return false
}

// SetRecentAssetIds gets a reference to the given []string and assigns it to the RecentAssetIds field.
func (o *LiveStream) SetRecentAssetIds(v []string) {
	o.RecentAssetIds = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LiveStream) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStream) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LiveStream) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LiveStream) SetStatus(v string) {
	o.Status = &v
}

// GetPlaybackIds returns the PlaybackIds field value if set, zero value otherwise.
func (o *LiveStream) GetPlaybackIds() []PlaybackID {
	if o == nil || o.PlaybackIds == nil {
		var ret []PlaybackID
		return ret
	}
	return *o.PlaybackIds
}

// GetPlaybackIdsOk returns a tuple with the PlaybackIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStream) GetPlaybackIdsOk() (*[]PlaybackID, bool) {
	if o == nil || o.PlaybackIds == nil {
		return nil, false
	}
	return o.PlaybackIds, true
}

// HasPlaybackIds returns a boolean if a field has been set.
func (o *LiveStream) HasPlaybackIds() bool {
	if o != nil && o.PlaybackIds != nil {
		return true
	}

	return false
}

// SetPlaybackIds gets a reference to the given []PlaybackID and assigns it to the PlaybackIds field.
func (o *LiveStream) SetPlaybackIds(v []PlaybackID) {
	o.PlaybackIds = &v
}

// GetNewAssetSettings returns the NewAssetSettings field value if set, zero value otherwise.
func (o *LiveStream) GetNewAssetSettings() CreateAssetRequest {
	if o == nil || o.NewAssetSettings == nil {
		var ret CreateAssetRequest
		return ret
	}
	return *o.NewAssetSettings
}

// GetNewAssetSettingsOk returns a tuple with the NewAssetSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStream) GetNewAssetSettingsOk() (*CreateAssetRequest, bool) {
	if o == nil || o.NewAssetSettings == nil {
		return nil, false
	}
	return o.NewAssetSettings, true
}

// HasNewAssetSettings returns a boolean if a field has been set.
func (o *LiveStream) HasNewAssetSettings() bool {
	if o != nil && o.NewAssetSettings != nil {
		return true
	}

	return false
}

// SetNewAssetSettings gets a reference to the given CreateAssetRequest and assigns it to the NewAssetSettings field.
func (o *LiveStream) SetNewAssetSettings(v CreateAssetRequest) {
	o.NewAssetSettings = &v
}

// GetPassthrough returns the Passthrough field value if set, zero value otherwise.
func (o *LiveStream) GetPassthrough() string {
	if o == nil || o.Passthrough == nil {
		var ret string
		return ret
	}
	return *o.Passthrough
}

// GetPassthroughOk returns a tuple with the Passthrough field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStream) GetPassthroughOk() (*string, bool) {
	if o == nil || o.Passthrough == nil {
		return nil, false
	}
	return o.Passthrough, true
}

// HasPassthrough returns a boolean if a field has been set.
func (o *LiveStream) HasPassthrough() bool {
	if o != nil && o.Passthrough != nil {
		return true
	}

	return false
}

// SetPassthrough gets a reference to the given string and assigns it to the Passthrough field.
func (o *LiveStream) SetPassthrough(v string) {
	o.Passthrough = &v
}

// GetReconnectWindow returns the ReconnectWindow field value if set, zero value otherwise.
func (o *LiveStream) GetReconnectWindow() float32 {
	if o == nil || o.ReconnectWindow == nil {
		var ret float32
		return ret
	}
	return *o.ReconnectWindow
}

// GetReconnectWindowOk returns a tuple with the ReconnectWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStream) GetReconnectWindowOk() (*float32, bool) {
	if o == nil || o.ReconnectWindow == nil {
		return nil, false
	}
	return o.ReconnectWindow, true
}

// HasReconnectWindow returns a boolean if a field has been set.
func (o *LiveStream) HasReconnectWindow() bool {
	if o != nil && o.ReconnectWindow != nil {
		return true
	}

	return false
}

// SetReconnectWindow gets a reference to the given float32 and assigns it to the ReconnectWindow field.
func (o *LiveStream) SetReconnectWindow(v float32) {
	o.ReconnectWindow = &v
}

// GetReducedLatency returns the ReducedLatency field value if set, zero value otherwise.
func (o *LiveStream) GetReducedLatency() bool {
	if o == nil || o.ReducedLatency == nil {
		var ret bool
		return ret
	}
	return *o.ReducedLatency
}

// GetReducedLatencyOk returns a tuple with the ReducedLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStream) GetReducedLatencyOk() (*bool, bool) {
	if o == nil || o.ReducedLatency == nil {
		return nil, false
	}
	return o.ReducedLatency, true
}

// HasReducedLatency returns a boolean if a field has been set.
func (o *LiveStream) HasReducedLatency() bool {
	if o != nil && o.ReducedLatency != nil {
		return true
	}

	return false
}

// SetReducedLatency gets a reference to the given bool and assigns it to the ReducedLatency field.
func (o *LiveStream) SetReducedLatency(v bool) {
	o.ReducedLatency = &v
}

// GetSimulcastTargets returns the SimulcastTargets field value if set, zero value otherwise.
func (o *LiveStream) GetSimulcastTargets() []SimulcastTarget {
	if o == nil || o.SimulcastTargets == nil {
		var ret []SimulcastTarget
		return ret
	}
	return *o.SimulcastTargets
}

// GetSimulcastTargetsOk returns a tuple with the SimulcastTargets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStream) GetSimulcastTargetsOk() (*[]SimulcastTarget, bool) {
	if o == nil || o.SimulcastTargets == nil {
		return nil, false
	}
	return o.SimulcastTargets, true
}

// HasSimulcastTargets returns a boolean if a field has been set.
func (o *LiveStream) HasSimulcastTargets() bool {
	if o != nil && o.SimulcastTargets != nil {
		return true
	}

	return false
}

// SetSimulcastTargets gets a reference to the given []SimulcastTarget and assigns it to the SimulcastTargets field.
func (o *LiveStream) SetSimulcastTargets(v []SimulcastTarget) {
	o.SimulcastTargets = &v
}

// GetTest returns the Test field value if set, zero value otherwise.
func (o *LiveStream) GetTest() bool {
	if o == nil || o.Test == nil {
		var ret bool
		return ret
	}
	return *o.Test
}

// GetTestOk returns a tuple with the Test field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStream) GetTestOk() (*bool, bool) {
	if o == nil || o.Test == nil {
		return nil, false
	}
	return o.Test, true
}

// HasTest returns a boolean if a field has been set.
func (o *LiveStream) HasTest() bool {
	if o != nil && o.Test != nil {
		return true
	}

	return false
}

// SetTest gets a reference to the given bool and assigns it to the Test field.
func (o *LiveStream) SetTest(v bool) {
	o.Test = &v
}

func (o LiveStream) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.StreamKey != nil {
		toSerialize["stream_key"] = o.StreamKey
	}
	if o.ActiveAssetId != nil {
		toSerialize["active_asset_id"] = o.ActiveAssetId
	}
	if o.RecentAssetIds != nil {
		toSerialize["recent_asset_ids"] = o.RecentAssetIds
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.PlaybackIds != nil {
		toSerialize["playback_ids"] = o.PlaybackIds
	}
	if o.NewAssetSettings != nil {
		toSerialize["new_asset_settings"] = o.NewAssetSettings
	}
	if o.Passthrough != nil {
		toSerialize["passthrough"] = o.Passthrough
	}
	if o.ReconnectWindow != nil {
		toSerialize["reconnect_window"] = o.ReconnectWindow
	}
	if o.ReducedLatency != nil {
		toSerialize["reduced_latency"] = o.ReducedLatency
	}
	if o.SimulcastTargets != nil {
		toSerialize["simulcast_targets"] = o.SimulcastTargets
	}
	if o.Test != nil {
		toSerialize["test"] = o.Test
	}
	return json.Marshal(toSerialize)
}

type NullableLiveStream struct {
	value *LiveStream
	isSet bool
}

func (v NullableLiveStream) Get() *LiveStream {
	return v.value
}

func (v *NullableLiveStream) Set(val *LiveStream) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveStream) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveStream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveStream(val *LiveStream) *NullableLiveStream {
	return &NullableLiveStream{value: val, isSet: true}
}

func (v NullableLiveStream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveStream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



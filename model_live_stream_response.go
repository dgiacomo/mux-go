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

// LiveStreamResponse struct for LiveStreamResponse
type LiveStreamResponse struct {
	Data *LiveStream `json:"data,omitempty"`
}

// NewLiveStreamResponse instantiates a new LiveStreamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveStreamResponse() *LiveStreamResponse {
	this := LiveStreamResponse{}
	return &this
}

// NewLiveStreamResponseWithDefaults instantiates a new LiveStreamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveStreamResponseWithDefaults() *LiveStreamResponse {
	this := LiveStreamResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LiveStreamResponse) GetData() LiveStream {
	if o == nil || o.Data == nil {
		var ret LiveStream
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStreamResponse) GetDataOk() (*LiveStream, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LiveStreamResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given LiveStream and assigns it to the Data field.
func (o *LiveStreamResponse) SetData(v LiveStream) {
	o.Data = &v
}

func (o LiveStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableLiveStreamResponse struct {
	value *LiveStreamResponse
	isSet bool
}

func (v NullableLiveStreamResponse) Get() *LiveStreamResponse {
	return v.value
}

func (v *NullableLiveStreamResponse) Set(val *LiveStreamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveStreamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveStreamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveStreamResponse(val *LiveStreamResponse) *NullableLiveStreamResponse {
	return &NullableLiveStreamResponse{value: val, isSet: true}
}

func (v NullableLiveStreamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveStreamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



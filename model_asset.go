// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

type Asset struct {
	// Unique identifier for the Asset. Max 255 characters.
	Id string `json:"id,omitempty"`
	// Time the Asset was created, defined as a Unix timestamp (seconds since epoch).
	CreatedAt string `json:"created_at,omitempty"`
	// The status of the asset.
	Status string `json:"status,omitempty"`
	// The duration of the asset in seconds (max duration for a single asset is 12 hours).
	Duration float64 `json:"duration,omitempty"`
	// The maximum resolution that has been stored for the asset. The asset may be delivered at lower resolutions depending on the device and bandwidth, however it cannot be delivered at a higher value than is stored.
	MaxStoredResolution string `json:"max_stored_resolution,omitempty"`
	// The maximum frame rate that has been stored for the asset. The asset may be delivered at lower frame rates depending on the device and bandwidth, however it cannot be delivered at a higher value than is stored. This field may return -1 if the frame rate of the input cannot be reliably determined.
	MaxStoredFrameRate float64 `json:"max_stored_frame_rate,omitempty"`
	// The aspect ratio of the asset in the form of `width:height`, for example `16:9`.
	AspectRatio string `json:"aspect_ratio,omitempty"`
	// An array of Playback ID objects. Use these to create HLS playback URLs. See [Play your videos](https://docs.mux.com/guides/video/play-your-videos) for more details.
	PlaybackIds []PlaybackId `json:"playback_ids,omitempty"`
	// The individual media tracks that make up an asset.
	Tracks         []Track     `json:"tracks,omitempty"`
	Errors         AssetErrors `json:"errors,omitempty"`
	PerTitleEncode bool        `json:"per_title_encode,omitempty"`
	// Unique identifier for the Direct Upload. This is an optional parameter added when the asset is created from a direct upload.
	UploadId string `json:"upload_id,omitempty"`
	// Whether the asset is created from a live stream and the live stream is currently `active` and not in `idle` state.
	IsLive bool `json:"is_live,omitempty"`
	// Arbitrary metadata set for the asset. Max 255 characters.
	Passthrough string `json:"passthrough,omitempty"`
	// Unique identifier for the live stream. This is an optional parameter added when the asset is created from a live stream.
	LiveStreamId string      `json:"live_stream_id,omitempty"`
	Master       AssetMaster `json:"master,omitempty"`
	MasterAccess string      `json:"master_access,omitempty"`
	Mp4Support   string      `json:"mp4_support,omitempty"`
	// Asset Identifier of the video used as the source for creating the clip.
	SourceAssetId string `json:"source_asset_id,omitempty"`
	// Normalize the audio track loudness level. This parameter is only applicable to on-demand (not live) assets.
	NormalizeAudio   bool                  `json:"normalize_audio,omitempty"`
	StaticRenditions AssetStaticRenditions `json:"static_renditions,omitempty"`
	// An array of individual live stream recording sessions. A recording session is created on each encoder connection during the live stream
	RecordingTimes          []AssetRecordingTimes        `json:"recording_times,omitempty"`
	NonStandardInputReasons AssetNonStandardInputReasons `json:"non_standard_input_reasons,omitempty"`
	// True means this live stream is a test asset. A test asset can help evaluate the Mux Video APIs without incurring any cost. There is no limit on number of test assets created. Test assets are watermarked with the Mux logo, limited to 10 seconds, and deleted after 24 hrs.
	Test bool `json:"test,omitempty"`
}

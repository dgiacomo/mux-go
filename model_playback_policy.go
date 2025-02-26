// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

// PlaybackPolicy : * `public` playback IDs are accessible by constructing an HLS url like `https://stream.mux.com/${PLAYBACK_ID}`  * `signed` playback IDS should be used with tokens `https://stream.mux.com/${PLAYBACK_ID}?token={TOKEN}`. See [Secure video playback](https://docs.mux.com/guides/video/secure-video-playback) for details about creating tokens.
type PlaybackPolicy string

// List of PlaybackPolicy
const (
	PUBLIC PlaybackPolicy = "public"
	SIGNED PlaybackPolicy = "signed"
)

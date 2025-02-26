// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

type CreateTrackRequest struct {
	Url      string `json:"url"`
	Type     string `json:"type"`
	TextType string `json:"text_type"`
	// The language code value must be a valid BCP 47 specification compliant value. For example, en for English or en-US for the US version of English.
	LanguageCode string `json:"language_code"`
	// The name of the track containing a human-readable description. This value must be unique across all the text type and subtitles text type tracks. HLS manifest will associate subtitle text track with this value. For example, set the value to \"English\" for subtitles text track with language_code as en-US. If this parameter is not included, Mux will auto-populate based on the language_code value.
	Name string `json:"name,omitempty"`
	// Indicates the track provides Subtitles for the Deaf or Hard-of-hearing (SDH).
	ClosedCaptions bool `json:"closed_captions,omitempty"`
	// Arbitrary metadata set for the track either when creating the asset or track.
	Passthrough string `json:"passthrough,omitempty"`
}

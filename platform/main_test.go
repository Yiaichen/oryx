// Copyright (c) 2026 Winlin
//
// SPDX-License-Identifier: MIT
package main

import "testing"

func TestStreamLimitDefaults(t *testing.T) {
	for _, tc := range []struct {
		name                               string
		forward, vlive, camera             string
		wantForward, wantVLive, wantCamera string
	}{
		{name: "defaults", wantForward: "10", wantVLive: "20", wantCamera: "10"},
		{name: "custom limits", forward: "6", vlive: "30", camera: "8", wantForward: "6", wantVLive: "30", wantCamera: "8"},
		{name: "explicit legacy limit", vlive: "10", wantForward: "10", wantVLive: "10", wantCamera: "10"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("SRS_FORWARD_LIMIT", tc.forward)
			t.Setenv("SRS_VLIVE_LIMIT", tc.vlive)
			t.Setenv("SRS_CAMERA_LIMIT", tc.camera)

			setStreamLimitDefaults()
			// Applying defaults again must preserve the initialized values.
			setStreamLimitDefaults()

			if got := envForwardLimit(); got != tc.wantForward {
				t.Errorf("forward limit = %q, want %q", got, tc.wantForward)
			}
			if got := envVLiveLimit(); got != tc.wantVLive {
				t.Errorf("virtual live limit = %q, want %q", got, tc.wantVLive)
			}
			if got := envCameraLimit(); got != tc.wantCamera {
				t.Errorf("camera limit = %q, want %q", got, tc.wantCamera)
			}
			t.Logf("forward=%s virtualLive=%s camera=%s", envForwardLimit(), envVLiveLimit(), envCameraLimit())
		})
	}
}

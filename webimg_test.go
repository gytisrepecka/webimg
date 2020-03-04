package webimg

import "testing"

func TestWatermark(t *testing.T) {
    // want := nil
    var want error = nil
    if got := Watermark("cmd/webimg/smplayer_preferences.jpg", "cmd/webimg/watermark_inretio-logo.png", "cmd/webimg/test_result_img.jpg", 30, 30, 70); got != want {
        t.Errorf("Watermark = %q, want %q", got, want)
    }
}

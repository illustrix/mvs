package auto

import "testing"

func TestGrabScreen(t *testing.T) {
	auto := NewAuto()
	img, err := auto.GrabScreen()
	if err != nil {
		t.Fatalf("Failed to grab screen: %v", err)
	}
	auto.AnalyticsColor(img)
}

package input

import (
	"testing"
	"time"
)

func TestLatencyRecorderComputesPercentile(t *testing.T) {
	var recorder LatencyRecorder
	recorder.Observe(10 * time.Millisecond)
	recorder.Observe(20 * time.Millisecond)
	recorder.Observe(30 * time.Millisecond)
	if got := recorder.Percentile(95); got != 30*time.Millisecond {
		t.Fatalf("p95 = %s", got)
	}
}

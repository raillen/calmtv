package input

import (
	"math"
	"sort"
	"sync"
	"time"
)

type LatencyRecorder struct {
	mu      sync.Mutex
	samples []time.Duration
}

func (r *LatencyRecorder) Observe(duration time.Duration) {
	if duration < 0 {
		return
	}
	r.mu.Lock()
	r.samples = append(r.samples, duration)
	r.mu.Unlock()
}

func (r *LatencyRecorder) Percentile(percent float64) time.Duration {
	r.mu.Lock()
	defer r.mu.Unlock()
	if len(r.samples) == 0 {
		return 0
	}
	samples := append([]time.Duration(nil), r.samples...)
	sort.Slice(samples, func(i, j int) bool { return samples[i] < samples[j] })
	if percent <= 0 {
		return samples[0]
	}
	if percent >= 100 {
		return samples[len(samples)-1]
	}
	index := int(math.Ceil(percent/100*float64(len(samples)))) - 1
	if index >= len(samples) {
		index = len(samples) - 1
	}
	return samples[index]
}

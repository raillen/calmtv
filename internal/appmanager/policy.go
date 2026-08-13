package appmanager

import "fmt"

type ResourceClass string

const (
	Light          ResourceClass = "light"
	Medium         ResourceClass = "medium"
	Heavy          ResourceClass = "heavy"
	ExclusiveHeavy ResourceClass = "exclusive-heavy"
)

type ResourcePolicy struct {
	MemoryHigh string
	MemoryMax  string
	CPUWeight  int
	IOWeight   int
}

func DefaultPolicy(class ResourceClass) ResourcePolicy {
	switch class {
	case Light:
		return ResourcePolicy{MemoryHigh: "128M", MemoryMax: "256M", CPUWeight: 100, IOWeight: 100}
	case Medium:
		return ResourcePolicy{MemoryHigh: "384M", MemoryMax: "512M", CPUWeight: 100, IOWeight: 100}
	case Heavy:
		return ResourcePolicy{MemoryHigh: "768M", MemoryMax: "1G", CPUWeight: 100, IOWeight: 100}
	case ExclusiveHeavy:
		return ResourcePolicy{MemoryHigh: "1G", MemoryMax: "1536M", CPUWeight: 100, IOWeight: 100}
	default:
		return ResourcePolicy{}
	}
}

func (p ResourcePolicy) Validate() error {
	if p.MemoryHigh == "" || p.MemoryMax == "" || p.CPUWeight < 1 || p.CPUWeight > 10000 || p.IOWeight < 1 || p.IOWeight > 10000 {
		return fmt.Errorf("invalid resource policy")
	}
	return nil
}

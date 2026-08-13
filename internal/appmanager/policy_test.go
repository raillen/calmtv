package appmanager

import "testing"

func TestDefaultPolicyKeepsExclusiveAppsBounded(t *testing.T) {
	policy := DefaultPolicy(ExclusiveHeavy)
	if err := policy.Validate(); err != nil {
		t.Fatal(err)
	}
	if policy.MemoryMax != "1536M" || policy.CPUWeight != 100 {
		t.Fatalf("unexpected policy: %#v", policy)
	}
}

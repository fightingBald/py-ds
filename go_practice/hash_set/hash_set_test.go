package hashset

import "testing"

func TestSetBasicOperations(t *testing.T) {
	var s Set

	s.Add("python")
	s.Add("go")
	s.Add("python")

	if got := s.Len(); got != 2 {
		t.Fatalf("expected len 2, got %d", got)
	}

	if !s.Contains("go") {
		t.Fatalf("expected set to contain 'go'")
	}
	if s.Contains("java") {
		t.Fatalf("did not expect set to contain 'java'")
	}

	items := s.Items()
	if len(items) != 2 {
		t.Fatalf("items should have len 2, got %d", len(items))
	}

	seen := make(map[string]struct{})
	for _, v := range items {
		seen[v] = struct{}{}
	}
	if _, ok := seen["python"]; !ok {
		t.Fatalf("items missing 'python'")
	}
	if _, ok := seen["go"]; !ok {
		t.Fatalf("items missing 'go'")
	}

	s.Remove("python")
	if s.Contains("python") {
		t.Fatalf("expected 'python' to be removed")
	}
	if got := s.Len(); got != 1 {
		t.Fatalf("expected len 1 after removal, got %d", got)
	}
}

func TestSetRemoveMissingPanics(t *testing.T) {
	var s Set
	mustPanic(t, func() { s.Remove("missing") })
}

func mustPanic(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if recover() == nil {
			t.Fatalf("expected panic, got none")
		}
	}()
	fn()
}

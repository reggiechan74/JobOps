package model

import "testing"

func TestStageStateGlyph(t *testing.T) {
	cases := map[StageState]string{Missing: "○", Draft: "◐", Final: "✓"}
	for st, want := range cases {
		if got := st.Glyph(); got != want {
			t.Errorf("StageState(%d).Glyph() = %q, want %q", st, got, want)
		}
	}
}

func TestLifecycleValidAndDefault(t *testing.T) {
	if !Interviewing.Valid() {
		t.Errorf("Interviewing should be valid")
	}
	if Lifecycle("bogus").Valid() {
		t.Errorf("bogus should be invalid")
	}
	if DefaultLifecycle() != Interested {
		t.Errorf("default = %q, want %q", DefaultLifecycle(), Interested)
	}
}

func TestLifecycleLabel(t *testing.T) {
	if Interviewing.Label() != "Interviewing" {
		t.Errorf("Label() = %q, want Interviewing", Interviewing.Label())
	}
}

package commons

import "testing"

func TestClassifySkill_Languages(t *testing.T) {
	t.Parallel()
	for _, tag := range []string{"go", "Go", "python", "RUST", "c++", "JavaScript"} {
		if got := ClassifySkill(tag); got != SkillLanguage {
			t.Errorf("ClassifySkill(%q) = %v, want Language", tag, got)
		}
	}
}

func TestClassifySkill_Domains(t *testing.T) {
	t.Parallel()
	for _, tag := range []string{"backend", "frontend", "devops", "machine-learning", "database", "security"} {
		if got := ClassifySkill(tag); got != SkillDomain {
			t.Errorf("ClassifySkill(%q) = %v, want Domain", tag, got)
		}
	}
}

func TestClassifySkill_DomainPrefix(t *testing.T) {
	t.Parallel()
	if got := ClassifySkill("machine-learning-systems"); got != SkillDomain {
		t.Errorf("ClassifySkill(%q) = %v, want Domain", "machine-learning-systems", got)
	}
}

func TestClassifySkill_Capabilities(t *testing.T) {
	t.Parallel()
	for _, tag := range []string{"testing", "cleanup", "lifecycle", "refactoring", "polecat"} {
		if got := ClassifySkill(tag); got != SkillCapability {
			t.Errorf("ClassifySkill(%q) = %v, want Capability", tag, got)
		}
	}
}

func TestSkillCategory_String(t *testing.T) {
	t.Parallel()
	tests := []struct {
		cat  SkillCategory
		want string
	}{
		{SkillLanguage, "Languages"},
		{SkillDomain, "Domains"},
		{SkillCapability, "Capabilities"},
	}
	for _, tt := range tests {
		if got := tt.cat.String(); got != tt.want {
			t.Errorf("SkillCategory(%d).String() = %q, want %q", tt.cat, got, tt.want)
		}
	}
}

package commons

import "strings"

// SkillCategory represents the taxonomy category for a skill tag.
type SkillCategory int

const (
	SkillLanguage   SkillCategory = iota // Programming language
	SkillDomain                          // Technical domain
	SkillCapability                      // General capability (default)
)

// String returns the human-readable category name.
func (c SkillCategory) String() string {
	switch c {
	case SkillLanguage:
		return "Languages"
	case SkillDomain:
		return "Domains"
	case SkillCapability:
		return "Capabilities"
	default:
		return "Capabilities"
	}
}

// Languages is the set of recognized programming language tags.
var Languages = map[string]bool{
	"c": true, "c++": true, "go": true, "rust": true, "python": true,
	"javascript": true, "typescript": true, "java": true, "ruby": true,
	"shell": true, "assembly": true, "makefile": true, "openscad": true,
	"kotlin": true, "swift": true, "scala": true, "haskell": true,
	"perl": true, "php": true, "lua": true, "r": true, "dart": true,
	"elixir": true, "erlang": true, "clojure": true, "zig": true,
	"nim": true, "ocaml": true, "f#": true, "c#": true, "html": true,
	"css": true, "sql": true, "matlab": true, "julia": true,
}

// domainPrefixes is the set of recognized technical domain tag prefixes.
var domainPrefixes = []string{
	"operating-systems", "systems-programming", "audio-processing",
	"text-processing", "hardware-design", "web-development",
	"machine-learning", "data-engineering", "mobile-development",
	"devops", "security", "database", "networking", "cloud",
	"frontend", "backend", "fullstack", "game-development",
	"embedded", "blockchain", "ai", "ml", "infrastructure",
}

// ClassifySkill returns the taxonomy category for a skill tag.
func ClassifySkill(tag string) SkillCategory {
	lower := strings.ToLower(tag)
	if Languages[lower] {
		return SkillLanguage
	}
	for _, prefix := range domainPrefixes {
		if lower == prefix || strings.HasPrefix(lower, prefix) {
			return SkillDomain
		}
	}
	return SkillCapability
}

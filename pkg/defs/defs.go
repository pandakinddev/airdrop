package defs

// Stage ...
type Stage int

const (
	// StageUnknown ...
	StageUnknown Stage = 0 + iota
	// StageRegistration ...
	StageRegistration
	// StageRewards ...
	StageRewards
	// StageEnded ...
	StageEnded
)

// String ...
func (s Stage) String() string {
	switch s {

	case StageRegistration:
		return "registration"
	case StageRewards:
		return "claiming"
	case StageEnded:
		return "ended"
	default:
		return "nearby"
	}
}

// SocialSource ...
type SocialSource byte

const (
	// UnknownSource ...
	UnknownSource SocialSource = 0 + iota
	// FacebookSource ...
	FacebookSource
)

package wasendcore


// Group participant role.
type GroupParticipantRole string

const (
	// Regular participant.
	GroupParticipantRole_PARTICIPANT GroupParticipantRole = "PARTICIPANT"
	// Group admin.
	GroupParticipantRole_ADMIN GroupParticipantRole = "ADMIN"
	// Group owner.
	GroupParticipantRole_OWNER GroupParticipantRole = "OWNER"
)


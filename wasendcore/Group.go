package wasendcore


// Group information.
type Group struct {
	// Addressing mode.
	AddressingMode *string `field:"required" json:"addressingMode" yaml:"addressingMode"`
	// Announcement version ID.
	AnnounceVersionId *string `field:"required" json:"announceVersionId" yaml:"announceVersionId"`
	// Creator's country code.
	CreatorCountryCode *string `field:"required" json:"creatorCountryCode" yaml:"creatorCountryCode"`
	// Default membership approval mode.
	DefaultMembershipApprovalMode *string `field:"required" json:"defaultMembershipApprovalMode" yaml:"defaultMembershipApprovalMode"`
	// Disappearing messages timer in seconds.
	DisappearingTimer *float64 `field:"required" json:"disappearingTimer" yaml:"disappearingTimer"`
	// When the group was created.
	GroupCreated *string `field:"required" json:"groupCreated" yaml:"groupCreated"`
	// Whether the group is an announcement group.
	IsAnnounce *bool `field:"required" json:"isAnnounce" yaml:"isAnnounce"`
	// Whether the group is a default subgroup.
	IsDefaultSubGroup *bool `field:"required" json:"isDefaultSubGroup" yaml:"isDefaultSubGroup"`
	// Whether the group is ephemeral.
	IsEphemeral *bool `field:"required" json:"isEphemeral" yaml:"isEphemeral"`
	// Whether the group is incognito.
	IsIncognito *bool `field:"required" json:"isIncognito" yaml:"isIncognito"`
	// Whether join approval is required.
	IsJoinApprovalRequired *bool `field:"required" json:"isJoinApprovalRequired" yaml:"isJoinApprovalRequired"`
	// Whether the group is locked.
	IsLocked *bool `field:"required" json:"isLocked" yaml:"isLocked"`
	// Whether the group is a parent group.
	IsParent *bool `field:"required" json:"isParent" yaml:"isParent"`
	// Group JID.
	Jid *string `field:"required" json:"jid" yaml:"jid"`
	// Linked parent group JID.
	LinkedParentJid *string `field:"required" json:"linkedParentJid" yaml:"linkedParentJid"`
	// Member add mode.
	MemberAddMode *string `field:"required" json:"memberAddMode" yaml:"memberAddMode"`
	// Group name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// When the name was set.
	NameSetAt *string `field:"required" json:"nameSetAt" yaml:"nameSetAt"`
	// Who set the name.
	NameSetBy *string `field:"required" json:"nameSetBy" yaml:"nameSetBy"`
	// Phone number of who set the name.
	NameSetByPn *string `field:"required" json:"nameSetByPn" yaml:"nameSetByPn"`
	// Owner JID.
	OwnerJid *string `field:"required" json:"ownerJid" yaml:"ownerJid"`
	// Owner phone number.
	OwnerPn *string `field:"required" json:"ownerPn" yaml:"ownerPn"`
	// Group participants.
	Participants *[]*GroupParticipant `field:"required" json:"participants" yaml:"participants"`
	// Participant version ID.
	ParticipantVersionId *string `field:"required" json:"participantVersionId" yaml:"participantVersionId"`
	// Group topic/description.
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// Whether the topic was deleted.
	TopicDeleted *bool `field:"required" json:"topicDeleted" yaml:"topicDeleted"`
	// Topic ID.
	TopicId *string `field:"required" json:"topicId" yaml:"topicId"`
	// When the topic was set.
	TopicSetAt *string `field:"required" json:"topicSetAt" yaml:"topicSetAt"`
	// Who set the topic.
	TopicSetBy *string `field:"required" json:"topicSetBy" yaml:"topicSetBy"`
	// Phone number of who set the topic.
	TopicSetByPn *string `field:"required" json:"topicSetByPn" yaml:"topicSetByPn"`
}


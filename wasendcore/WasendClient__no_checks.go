//go:build no_runtime_type_checking

package wasendcore

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WasendClient) validateAddGroupParticipantsParameters(sessionId *string, groupId *string, request *ParticipantsRequest) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateCheckContactExistsParameters(sessionId *string, params *CheckContactExistsQueryParams) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateCreateGroupParameters(sessionId *string, request *CreateGroupRequest) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateCreateSessionParameters(request *SessionCreateRequest) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateDeleteGroupPictureParameters(sessionId *string, groupId *string) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateDeleteSessionParameters(sessionId *string) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateDemoteGroupParticipantsParameters(sessionId *string, groupId *string, request *ParticipantsRequest) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateGetContactParameters(sessionId *string, params *GetContactQueryParams) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateGetContactProfilePictureParameters(sessionId *string, params *GetProfilePictureQueryParams) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateGetContactsParameters(sessionId *string, params *GetContactsQueryParams) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateGetGroupParameters(sessionId *string, groupId *string) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateGetGroupInfoAdminOnlyParameters(sessionId *string, groupId *string) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateGetGroupInviteCodeParameters(sessionId *string, groupId *string) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateGetGroupJoinInfoParameters(sessionId *string, code *string) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateGetGroupMessagesAdminOnlyParameters(sessionId *string, groupId *string) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateGetGroupParticipantsParameters(sessionId *string, groupId *string) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateGetGroupPictureParameters(sessionId *string, groupId *string) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateGetGroupsParameters(sessionId *string, params *GetGroupsQueryParams) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateGetGroupsCountParameters(sessionId *string) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateGetMessageParameters(messageId *string) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateJoinGroupParameters(sessionId *string, request *JoinGroupRequest) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateLeaveGroupParameters(sessionId *string, groupId *string) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateProcessMessageParameters(request *SendRequest) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validatePromoteGroupParticipantsParameters(sessionId *string, groupId *string, request *ParticipantsRequest) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateRemoveGroupParticipantsParameters(sessionId *string, groupId *string, request *ParticipantsRequest) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateRestartSessionParameters(sessionId *string) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateRetrieveQRCodeParameters(sessionId *string) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateRetrieveSessionInfoParameters(sessionId *string) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateRevokeGroupInviteCodeParameters(sessionId *string, groupId *string) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateSendParameters(request *SendRequest) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateSendFileParameters(request *MessageFileRequest) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateSendImageParameters(request *MessageImageRequest) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateSendLinkCustomPreviewParameters(request *MessageLinkCustomPreviewRequest) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateSendMessageParameters(request *MessageRequest) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateSendSeenParameters(request *SendSeenRequest) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateSendTextParameters(request *MessageTextRequest) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateSendVideoParameters(request *MessageVideoRequest) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateSendVoiceParameters(request *MessageVoiceRequest) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateSetGroupDescriptionParameters(sessionId *string, groupId *string, request *DescriptionRequest) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateSetGroupInfoAdminOnlyParameters(sessionId *string, groupId *string, request *SettingsSecurityChangeInfo) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateSetGroupMessagesAdminOnlyParameters(sessionId *string, groupId *string, request *SettingsSecurityChangeInfo) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateSetGroupPictureParameters(sessionId *string, groupId *string, request *ProfilePictureRequest) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateSetGroupSubjectParameters(sessionId *string, groupId *string, request *SubjectRequest) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateStartSessionParameters(sessionId *string) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateStartTypingParameters(request *ChatRequest) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateStopSessionParameters(sessionId *string) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateStopTypingParameters(request *ChatRequest) error {
	return nil
}

func validateNewWasendClientParameters(config *WasendConfig) error {
	return nil
}


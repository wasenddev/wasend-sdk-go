//go:build !no_runtime_type_checking

package wasendcore

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (w *jsiiProxy_WasendClient) validateAddGroupParticipantsParameters(sessionId *string, groupId *string, request *ParticipantsRequest) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	if groupId == nil {
		return fmt.Errorf("parameter groupId is required, but nil was provided")
	}

	if request == nil {
		return fmt.Errorf("parameter request is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(request, func() string { return "parameter request" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateCheckContactExistsParameters(sessionId *string, params *CheckContactExistsQueryParams) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	if params == nil {
		return fmt.Errorf("parameter params is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(params, func() string { return "parameter params" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateCreateGroupParameters(sessionId *string, request *CreateGroupRequest) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	if request == nil {
		return fmt.Errorf("parameter request is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(request, func() string { return "parameter request" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateCreateSessionParameters(request *SessionCreateRequest) error {
	if request == nil {
		return fmt.Errorf("parameter request is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(request, func() string { return "parameter request" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateDeleteGroupPictureParameters(sessionId *string, groupId *string) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	if groupId == nil {
		return fmt.Errorf("parameter groupId is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateDeleteSessionParameters(sessionId *string) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateDemoteGroupParticipantsParameters(sessionId *string, groupId *string, request *ParticipantsRequest) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	if groupId == nil {
		return fmt.Errorf("parameter groupId is required, but nil was provided")
	}

	if request == nil {
		return fmt.Errorf("parameter request is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(request, func() string { return "parameter request" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateGetAllChatsParameters(session *string, options *GetChatsOptions) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateGetChatPictureParameters(session *string, chatId *string, options *GetChatPictureOptions) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	if chatId == nil {
		return fmt.Errorf("parameter chatId is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateGetChatsOverviewParameters(session *string, options *GetChatsOverviewOptions) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateGetContactParameters(sessionId *string, params *GetContactQueryParams) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	if params == nil {
		return fmt.Errorf("parameter params is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(params, func() string { return "parameter params" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateGetContactProfilePictureParameters(sessionId *string, params *GetProfilePictureQueryParams) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	if params == nil {
		return fmt.Errorf("parameter params is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(params, func() string { return "parameter params" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateGetContactsParameters(sessionId *string, params *GetContactsQueryParams) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(params, func() string { return "parameter params" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateGetGroupParameters(sessionId *string, groupId *string) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	if groupId == nil {
		return fmt.Errorf("parameter groupId is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateGetGroupInfoAdminOnlyParameters(sessionId *string, groupId *string) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	if groupId == nil {
		return fmt.Errorf("parameter groupId is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateGetGroupInviteCodeParameters(sessionId *string, groupId *string) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	if groupId == nil {
		return fmt.Errorf("parameter groupId is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateGetGroupJoinInfoParameters(sessionId *string, code *string) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	if code == nil {
		return fmt.Errorf("parameter code is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateGetGroupMessagesAdminOnlyParameters(sessionId *string, groupId *string) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	if groupId == nil {
		return fmt.Errorf("parameter groupId is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateGetGroupParticipantsParameters(sessionId *string, groupId *string) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	if groupId == nil {
		return fmt.Errorf("parameter groupId is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateGetGroupPictureParameters(sessionId *string, groupId *string) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	if groupId == nil {
		return fmt.Errorf("parameter groupId is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateGetGroupsParameters(sessionId *string, params *GetGroupsQueryParams) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(params, func() string { return "parameter params" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateGetGroupsCountParameters(sessionId *string) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateGetMessageParameters(messageId *string) error {
	if messageId == nil {
		return fmt.Errorf("parameter messageId is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateGetMessageByIdParameters(session *string, chatId *string, messageId *string, options *GetMessageByIdOptions) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	if chatId == nil {
		return fmt.Errorf("parameter chatId is required, but nil was provided")
	}

	if messageId == nil {
		return fmt.Errorf("parameter messageId is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateGetMessagesParameters(session *string, chatId *string, options *GetMessagesOptions) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	if chatId == nil {
		return fmt.Errorf("parameter chatId is required, but nil was provided")
	}

	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateJoinGroupParameters(sessionId *string, request *JoinGroupRequest) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	if request == nil {
		return fmt.Errorf("parameter request is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(request, func() string { return "parameter request" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateLeaveGroupParameters(sessionId *string, groupId *string) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	if groupId == nil {
		return fmt.Errorf("parameter groupId is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateProcessMessageParameters(request *SendRequest) error {
	if request == nil {
		return fmt.Errorf("parameter request is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(request, func() string { return "parameter request" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validatePromoteGroupParticipantsParameters(sessionId *string, groupId *string, request *ParticipantsRequest) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	if groupId == nil {
		return fmt.Errorf("parameter groupId is required, but nil was provided")
	}

	if request == nil {
		return fmt.Errorf("parameter request is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(request, func() string { return "parameter request" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateReadMessagesParameters(session *string, chatId *string, options *ReadMessagesOptions) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	if chatId == nil {
		return fmt.Errorf("parameter chatId is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateRemoveGroupParticipantsParameters(sessionId *string, groupId *string, request *ParticipantsRequest) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	if groupId == nil {
		return fmt.Errorf("parameter groupId is required, but nil was provided")
	}

	if request == nil {
		return fmt.Errorf("parameter request is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(request, func() string { return "parameter request" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateRestartSessionParameters(sessionId *string) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateRetrieveQRCodeParameters(sessionId *string) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateRetrieveSessionInfoParameters(sessionId *string) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateRevokeGroupInviteCodeParameters(sessionId *string, groupId *string) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	if groupId == nil {
		return fmt.Errorf("parameter groupId is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateSendParameters(request *SendRequest) error {
	if request == nil {
		return fmt.Errorf("parameter request is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(request, func() string { return "parameter request" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateSendFileParameters(request *MessageFileRequest) error {
	if request == nil {
		return fmt.Errorf("parameter request is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(request, func() string { return "parameter request" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateSendImageParameters(request *MessageImageRequest) error {
	if request == nil {
		return fmt.Errorf("parameter request is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(request, func() string { return "parameter request" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateSendLinkCustomPreviewParameters(request *MessageLinkCustomPreviewRequest) error {
	if request == nil {
		return fmt.Errorf("parameter request is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(request, func() string { return "parameter request" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateSendMessageParameters(request *MessageRequest) error {
	if request == nil {
		return fmt.Errorf("parameter request is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(request, func() string { return "parameter request" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateSendReactionParameters(request *MessageReactionRequest) error {
	if request == nil {
		return fmt.Errorf("parameter request is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(request, func() string { return "parameter request" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateSendSeenParameters(request *SendSeenRequest) error {
	if request == nil {
		return fmt.Errorf("parameter request is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(request, func() string { return "parameter request" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateSendTextParameters(request *MessageTextRequest) error {
	if request == nil {
		return fmt.Errorf("parameter request is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(request, func() string { return "parameter request" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateSendVideoParameters(request *MessageVideoRequest) error {
	if request == nil {
		return fmt.Errorf("parameter request is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(request, func() string { return "parameter request" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateSendVoiceParameters(request *MessageVoiceRequest) error {
	if request == nil {
		return fmt.Errorf("parameter request is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(request, func() string { return "parameter request" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateSetGroupDescriptionParameters(sessionId *string, groupId *string, request *DescriptionRequest) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	if groupId == nil {
		return fmt.Errorf("parameter groupId is required, but nil was provided")
	}

	if request == nil {
		return fmt.Errorf("parameter request is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(request, func() string { return "parameter request" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateSetGroupInfoAdminOnlyParameters(sessionId *string, groupId *string, request *SettingsSecurityChangeInfo) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	if groupId == nil {
		return fmt.Errorf("parameter groupId is required, but nil was provided")
	}

	if request == nil {
		return fmt.Errorf("parameter request is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(request, func() string { return "parameter request" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateSetGroupMessagesAdminOnlyParameters(sessionId *string, groupId *string, request *SettingsSecurityChangeInfo) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	if groupId == nil {
		return fmt.Errorf("parameter groupId is required, but nil was provided")
	}

	if request == nil {
		return fmt.Errorf("parameter request is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(request, func() string { return "parameter request" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateSetGroupPictureParameters(sessionId *string, groupId *string, request *ProfilePictureRequest) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	if groupId == nil {
		return fmt.Errorf("parameter groupId is required, but nil was provided")
	}

	if request == nil {
		return fmt.Errorf("parameter request is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(request, func() string { return "parameter request" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateSetGroupSubjectParameters(sessionId *string, groupId *string, request *SubjectRequest) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	if groupId == nil {
		return fmt.Errorf("parameter groupId is required, but nil was provided")
	}

	if request == nil {
		return fmt.Errorf("parameter request is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(request, func() string { return "parameter request" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateStartSessionParameters(sessionId *string) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateStartTypingParameters(request *ChatRequest) error {
	if request == nil {
		return fmt.Errorf("parameter request is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(request, func() string { return "parameter request" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateStopSessionParameters(sessionId *string) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateStopTypingParameters(request *ChatRequest) error {
	if request == nil {
		return fmt.Errorf("parameter request is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(request, func() string { return "parameter request" }); err != nil {
		return err
	}

	return nil
}

func validateNewWasendClientParameters(config *WasendConfig) error {
	if config == nil {
		return fmt.Errorf("parameter config is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(config, func() string { return "parameter config" }); err != nil {
		return err
	}

	return nil
}


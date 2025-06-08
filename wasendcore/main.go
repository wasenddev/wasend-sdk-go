// Wasend SDK for multiple programming languages
package wasendcore

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@wasend/core.AccountInfo",
		reflect.TypeOf((*AccountInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.ApiResponse",
		reflect.TypeOf((*ApiResponse)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.ChatRequest",
		reflect.TypeOf((*ChatRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.CheckContactExistsQueryParams",
		reflect.TypeOf((*CheckContactExistsQueryParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.Contact",
		reflect.TypeOf((*Contact)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.CountResponse",
		reflect.TypeOf((*CountResponse)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.CreateGroupParticipant",
		reflect.TypeOf((*CreateGroupParticipant)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.CreateGroupRequest",
		reflect.TypeOf((*CreateGroupRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.DescriptionRequest",
		reflect.TypeOf((*DescriptionRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.DownstreamInfo",
		reflect.TypeOf((*DownstreamInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.EngineStatus",
		reflect.TypeOf((*EngineStatus)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.GetAllSessionsResponse",
		reflect.TypeOf((*GetAllSessionsResponse)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.GetContactQueryParams",
		reflect.TypeOf((*GetContactQueryParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.GetContactsQueryParams",
		reflect.TypeOf((*GetContactsQueryParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.GetGroupsQueryParams",
		reflect.TypeOf((*GetGroupsQueryParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.GetProfilePictureQueryParams",
		reflect.TypeOf((*GetProfilePictureQueryParams)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.GowsStatus",
		reflect.TypeOf((*GowsStatus)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.Group",
		reflect.TypeOf((*Group)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.GroupParticipant",
		reflect.TypeOf((*GroupParticipant)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@wasend/core.GroupParticipantRole",
		reflect.TypeOf((*GroupParticipantRole)(nil)).Elem(),
		map[string]interface{}{
			"PARTICIPANT": GroupParticipantRole_PARTICIPANT,
			"ADMIN": GroupParticipantRole_ADMIN,
			"OWNER": GroupParticipantRole_OWNER,
		},
	)
	_jsii_.RegisterStruct(
		"@wasend/core.GroupPictureResponse",
		reflect.TypeOf((*GroupPictureResponse)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.GroupSettings",
		reflect.TypeOf((*GroupSettings)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@wasend/core.GroupStatus",
		reflect.TypeOf((*GroupStatus)(nil)).Elem(),
		map[string]interface{}{
			"ACTIVE": GroupStatus_ACTIVE,
			"ARCHIVED": GroupStatus_ARCHIVED,
			"DELETED": GroupStatus_DELETED,
		},
	)
	_jsii_.RegisterStruct(
		"@wasend/core.GrpcStatus",
		reflect.TypeOf((*GrpcStatus)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.JoinGroupRequest",
		reflect.TypeOf((*JoinGroupRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.JoinGroupResponse",
		reflect.TypeOf((*JoinGroupResponse)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.LinkPreviewRequest",
		reflect.TypeOf((*LinkPreviewRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.Message",
		reflect.TypeOf((*Message)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.MessageFileRequest",
		reflect.TypeOf((*MessageFileRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.MessageImageRequest",
		reflect.TypeOf((*MessageImageRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.MessageLinkCustomPreviewRequest",
		reflect.TypeOf((*MessageLinkCustomPreviewRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.MessageRequest",
		reflect.TypeOf((*MessageRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.MessageTextRequest",
		reflect.TypeOf((*MessageTextRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.MessageVideoRequest",
		reflect.TypeOf((*MessageVideoRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.MessageVoiceRequest",
		reflect.TypeOf((*MessageVoiceRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.NowebConfig",
		reflect.TypeOf((*NowebConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.NowebStoreConfig",
		reflect.TypeOf((*NowebStoreConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.ParticipantsRequest",
		reflect.TypeOf((*ParticipantsRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.ProfilePictureRequest",
		reflect.TypeOf((*ProfilePictureRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.ProfilePictureResponse",
		reflect.TypeOf((*ProfilePictureResponse)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.QRCodeResponse",
		reflect.TypeOf((*QRCodeResponse)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.SendRequest",
		reflect.TypeOf((*SendRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.SendSeenRequest",
		reflect.TypeOf((*SendSeenRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.Session",
		reflect.TypeOf((*Session)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.SessionConfig",
		reflect.TypeOf((*SessionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.SessionCreateRequest",
		reflect.TypeOf((*SessionCreateRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.SessionDetails",
		reflect.TypeOf((*SessionDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.SessionListItem",
		reflect.TypeOf((*SessionListItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.SessionMetadata",
		reflect.TypeOf((*SessionMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.SettingsSecurityChangeInfo",
		reflect.TypeOf((*SettingsSecurityChangeInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.SubjectRequest",
		reflect.TypeOf((*SubjectRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.WAMessage",
		reflect.TypeOf((*WAMessage)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.WANumberExistResult",
		reflect.TypeOf((*WANumberExistResult)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@wasend/core.WasendClient",
		reflect.TypeOf((*WasendClient)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addGroupParticipants", GoMethod: "AddGroupParticipants"},
			_jsii_.MemberMethod{JsiiMethod: "checkContactExists", GoMethod: "CheckContactExists"},
			_jsii_.MemberMethod{JsiiMethod: "createGroup", GoMethod: "CreateGroup"},
			_jsii_.MemberMethod{JsiiMethod: "createSession", GoMethod: "CreateSession"},
			_jsii_.MemberMethod{JsiiMethod: "deleteGroupPicture", GoMethod: "DeleteGroupPicture"},
			_jsii_.MemberMethod{JsiiMethod: "deleteSession", GoMethod: "DeleteSession"},
			_jsii_.MemberMethod{JsiiMethod: "demoteGroupParticipants", GoMethod: "DemoteGroupParticipants"},
			_jsii_.MemberMethod{JsiiMethod: "getContact", GoMethod: "GetContact"},
			_jsii_.MemberMethod{JsiiMethod: "getContactProfilePicture", GoMethod: "GetContactProfilePicture"},
			_jsii_.MemberMethod{JsiiMethod: "getContacts", GoMethod: "GetContacts"},
			_jsii_.MemberMethod{JsiiMethod: "getGroup", GoMethod: "GetGroup"},
			_jsii_.MemberMethod{JsiiMethod: "getGroupInfoAdminOnly", GoMethod: "GetGroupInfoAdminOnly"},
			_jsii_.MemberMethod{JsiiMethod: "getGroupInviteCode", GoMethod: "GetGroupInviteCode"},
			_jsii_.MemberMethod{JsiiMethod: "getGroupJoinInfo", GoMethod: "GetGroupJoinInfo"},
			_jsii_.MemberMethod{JsiiMethod: "getGroupMessagesAdminOnly", GoMethod: "GetGroupMessagesAdminOnly"},
			_jsii_.MemberMethod{JsiiMethod: "getGroupParticipants", GoMethod: "GetGroupParticipants"},
			_jsii_.MemberMethod{JsiiMethod: "getGroupPicture", GoMethod: "GetGroupPicture"},
			_jsii_.MemberMethod{JsiiMethod: "getGroups", GoMethod: "GetGroups"},
			_jsii_.MemberMethod{JsiiMethod: "getGroupsCount", GoMethod: "GetGroupsCount"},
			_jsii_.MemberMethod{JsiiMethod: "getMessage", GoMethod: "GetMessage"},
			_jsii_.MemberMethod{JsiiMethod: "joinGroup", GoMethod: "JoinGroup"},
			_jsii_.MemberMethod{JsiiMethod: "leaveGroup", GoMethod: "LeaveGroup"},
			_jsii_.MemberMethod{JsiiMethod: "processMessage", GoMethod: "ProcessMessage"},
			_jsii_.MemberMethod{JsiiMethod: "promoteGroupParticipants", GoMethod: "PromoteGroupParticipants"},
			_jsii_.MemberMethod{JsiiMethod: "removeGroupParticipants", GoMethod: "RemoveGroupParticipants"},
			_jsii_.MemberMethod{JsiiMethod: "restartSession", GoMethod: "RestartSession"},
			_jsii_.MemberMethod{JsiiMethod: "retrieveAccount", GoMethod: "RetrieveAccount"},
			_jsii_.MemberMethod{JsiiMethod: "retrieveAllSessions", GoMethod: "RetrieveAllSessions"},
			_jsii_.MemberMethod{JsiiMethod: "retrieveConfig", GoMethod: "RetrieveConfig"},
			_jsii_.MemberMethod{JsiiMethod: "retrieveQRCode", GoMethod: "RetrieveQRCode"},
			_jsii_.MemberMethod{JsiiMethod: "retrieveSessionInfo", GoMethod: "RetrieveSessionInfo"},
			_jsii_.MemberMethod{JsiiMethod: "revokeGroupInviteCode", GoMethod: "RevokeGroupInviteCode"},
			_jsii_.MemberMethod{JsiiMethod: "send", GoMethod: "Send"},
			_jsii_.MemberMethod{JsiiMethod: "sendFile", GoMethod: "SendFile"},
			_jsii_.MemberMethod{JsiiMethod: "sendImage", GoMethod: "SendImage"},
			_jsii_.MemberMethod{JsiiMethod: "sendLinkCustomPreview", GoMethod: "SendLinkCustomPreview"},
			_jsii_.MemberMethod{JsiiMethod: "sendMessage", GoMethod: "SendMessage"},
			_jsii_.MemberMethod{JsiiMethod: "sendSeen", GoMethod: "SendSeen"},
			_jsii_.MemberMethod{JsiiMethod: "sendText", GoMethod: "SendText"},
			_jsii_.MemberMethod{JsiiMethod: "sendVideo", GoMethod: "SendVideo"},
			_jsii_.MemberMethod{JsiiMethod: "sendVoice", GoMethod: "SendVoice"},
			_jsii_.MemberMethod{JsiiMethod: "setGroupDescription", GoMethod: "SetGroupDescription"},
			_jsii_.MemberMethod{JsiiMethod: "setGroupInfoAdminOnly", GoMethod: "SetGroupInfoAdminOnly"},
			_jsii_.MemberMethod{JsiiMethod: "setGroupMessagesAdminOnly", GoMethod: "SetGroupMessagesAdminOnly"},
			_jsii_.MemberMethod{JsiiMethod: "setGroupPicture", GoMethod: "SetGroupPicture"},
			_jsii_.MemberMethod{JsiiMethod: "setGroupSubject", GoMethod: "SetGroupSubject"},
			_jsii_.MemberMethod{JsiiMethod: "startSession", GoMethod: "StartSession"},
			_jsii_.MemberMethod{JsiiMethod: "startTyping", GoMethod: "StartTyping"},
			_jsii_.MemberMethod{JsiiMethod: "stopSession", GoMethod: "StopSession"},
			_jsii_.MemberMethod{JsiiMethod: "stopTyping", GoMethod: "StopTyping"},
		},
		func() interface{} {
			return &jsiiProxy_WasendClient{}
		},
	)
	_jsii_.RegisterStruct(
		"@wasend/core.WasendConfig",
		reflect.TypeOf((*WasendConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.WasendConfigInfo",
		reflect.TypeOf((*WasendConfigInfo)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@wasend/core.WasendUtils",
		reflect.TypeOf((*WasendUtils)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_WasendUtils{}
		},
	)
}

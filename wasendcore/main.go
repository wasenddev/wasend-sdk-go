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
		"@wasend/core.GowsStatus",
		reflect.TypeOf((*GowsStatus)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.GrpcStatus",
		reflect.TypeOf((*GrpcStatus)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.Message",
		reflect.TypeOf((*Message)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@wasend/core.MessageRequest",
		reflect.TypeOf((*MessageRequest)(nil)).Elem(),
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
		"@wasend/core.QRCodeResponse",
		reflect.TypeOf((*QRCodeResponse)(nil)).Elem(),
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
	_jsii_.RegisterClass(
		"@wasend/core.WasendClient",
		reflect.TypeOf((*WasendClient)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "createSession", GoMethod: "CreateSession"},
			_jsii_.MemberMethod{JsiiMethod: "deleteSession", GoMethod: "DeleteSession"},
			_jsii_.MemberMethod{JsiiMethod: "getMessage", GoMethod: "GetMessage"},
			_jsii_.MemberMethod{JsiiMethod: "restartSession", GoMethod: "RestartSession"},
			_jsii_.MemberMethod{JsiiMethod: "retrieveAccount", GoMethod: "RetrieveAccount"},
			_jsii_.MemberMethod{JsiiMethod: "retrieveAllSessions", GoMethod: "RetrieveAllSessions"},
			_jsii_.MemberMethod{JsiiMethod: "retrieveConfig", GoMethod: "RetrieveConfig"},
			_jsii_.MemberMethod{JsiiMethod: "retrieveQRCode", GoMethod: "RetrieveQRCode"},
			_jsii_.MemberMethod{JsiiMethod: "retrieveSessionInfo", GoMethod: "RetrieveSessionInfo"},
			_jsii_.MemberMethod{JsiiMethod: "sendMessage", GoMethod: "SendMessage"},
			_jsii_.MemberMethod{JsiiMethod: "startSession", GoMethod: "StartSession"},
			_jsii_.MemberMethod{JsiiMethod: "stopSession", GoMethod: "StopSession"},
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

package wasendcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/wasenddev/wasend-sdk-go/wasendcore/jsii"
)

// Main Wasend SDK Client.
//
// This class provides access to the Wasend API for sending messages,
// managing contacts, and other messaging operations.
//
// Example:
//   const client = new WasendClient({
//     apiKey: 'your-api-key'
//   });
//
//   const result = await client.sendMessage({
//     to: '+1234567890',
//     content: 'Hello, World!'
//   });
//
type WasendClient interface {
	CreateSession(request *SessionCreateRequest) *Session
	DeleteSession(sessionId *string)
	// Get message by ID.
	//
	// Returns: Promise resolving to the API response.
	GetMessage(messageId *string) *ApiResponse
	RestartSession(sessionId *string) *Session
	// Get account information.
	//
	// Returns: Promise resolving to the API response.
	RetrieveAccount() *ApiResponse
	RetrieveAllSessions() *GetAllSessionsResponse
	// Get the current configuration.
	//
	// Returns: The configuration object (without sensitive data).
	RetrieveConfig() *WasendConfigInfo
	RetrieveQRCode(sessionId *string) *QRCodeResponse
	RetrieveSessionInfo(sessionId *string) *Session
	// Send a message to a recipient.
	//
	// Returns: Promise resolving to the API response.
	SendMessage(request *MessageRequest) *ApiResponse
	StartSession(sessionId *string) *Session
	StopSession(sessionId *string) *Session
}

// The jsii proxy struct for WasendClient
type jsiiProxy_WasendClient struct {
	_ byte // padding
}

// Creates a new Wasend client instance.
func NewWasendClient(config *WasendConfig) WasendClient {
	_init_.Initialize()

	if err := validateNewWasendClientParameters(config); err != nil {
		panic(err)
	}
	j := jsiiProxy_WasendClient{}

	_jsii_.Create(
		"@wasend/core.WasendClient",
		[]interface{}{config},
		&j,
	)

	return &j
}

// Creates a new Wasend client instance.
func NewWasendClient_Override(w WasendClient, config *WasendConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@wasend/core.WasendClient",
		[]interface{}{config},
		w,
	)
}

func (w *jsiiProxy_WasendClient) CreateSession(request *SessionCreateRequest) *Session {
	if err := w.validateCreateSessionParameters(request); err != nil {
		panic(err)
	}
	var returns *Session

	_jsii_.Invoke(
		w,
		"createSession",
		[]interface{}{request},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) DeleteSession(sessionId *string) {
	if err := w.validateDeleteSessionParameters(sessionId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"deleteSession",
		[]interface{}{sessionId},
	)
}

func (w *jsiiProxy_WasendClient) GetMessage(messageId *string) *ApiResponse {
	if err := w.validateGetMessageParameters(messageId); err != nil {
		panic(err)
	}
	var returns *ApiResponse

	_jsii_.Invoke(
		w,
		"getMessage",
		[]interface{}{messageId},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) RestartSession(sessionId *string) *Session {
	if err := w.validateRestartSessionParameters(sessionId); err != nil {
		panic(err)
	}
	var returns *Session

	_jsii_.Invoke(
		w,
		"restartSession",
		[]interface{}{sessionId},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) RetrieveAccount() *ApiResponse {
	var returns *ApiResponse

	_jsii_.Invoke(
		w,
		"retrieveAccount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) RetrieveAllSessions() *GetAllSessionsResponse {
	var returns *GetAllSessionsResponse

	_jsii_.Invoke(
		w,
		"retrieveAllSessions",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) RetrieveConfig() *WasendConfigInfo {
	var returns *WasendConfigInfo

	_jsii_.Invoke(
		w,
		"retrieveConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) RetrieveQRCode(sessionId *string) *QRCodeResponse {
	if err := w.validateRetrieveQRCodeParameters(sessionId); err != nil {
		panic(err)
	}
	var returns *QRCodeResponse

	_jsii_.Invoke(
		w,
		"retrieveQRCode",
		[]interface{}{sessionId},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) RetrieveSessionInfo(sessionId *string) *Session {
	if err := w.validateRetrieveSessionInfoParameters(sessionId); err != nil {
		panic(err)
	}
	var returns *Session

	_jsii_.Invoke(
		w,
		"retrieveSessionInfo",
		[]interface{}{sessionId},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) SendMessage(request *MessageRequest) *ApiResponse {
	if err := w.validateSendMessageParameters(request); err != nil {
		panic(err)
	}
	var returns *ApiResponse

	_jsii_.Invoke(
		w,
		"sendMessage",
		[]interface{}{request},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) StartSession(sessionId *string) *Session {
	if err := w.validateStartSessionParameters(sessionId); err != nil {
		panic(err)
	}
	var returns *Session

	_jsii_.Invoke(
		w,
		"startSession",
		[]interface{}{sessionId},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WasendClient) StopSession(sessionId *string) *Session {
	if err := w.validateStopSessionParameters(sessionId); err != nil {
		panic(err)
	}
	var returns *Session

	_jsii_.Invoke(
		w,
		"stopSession",
		[]interface{}{sessionId},
		&returns,
	)

	return returns
}


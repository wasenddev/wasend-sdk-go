//go:build no_runtime_type_checking

package wasendcore

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WasendClient) validateCreateSessionParameters(request *SessionCreateRequest) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateDeleteSessionParameters(sessionId *string) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateGetMessageParameters(messageId *string) error {
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

func (w *jsiiProxy_WasendClient) validateSendMessageParameters(request *MessageRequest) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateStartSessionParameters(sessionId *string) error {
	return nil
}

func (w *jsiiProxy_WasendClient) validateStopSessionParameters(sessionId *string) error {
	return nil
}

func validateNewWasendClientParameters(config *WasendConfig) error {
	return nil
}


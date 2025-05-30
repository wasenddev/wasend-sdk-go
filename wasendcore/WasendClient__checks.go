//go:build !no_runtime_type_checking

package wasendcore

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (w *jsiiProxy_WasendClient) validateCreateSessionParameters(request *SessionCreateRequest) error {
	if request == nil {
		return fmt.Errorf("parameter request is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(request, func() string { return "parameter request" }); err != nil {
		return err
	}

	return nil
}

func (w *jsiiProxy_WasendClient) validateDeleteSessionParameters(sessionId *string) error {
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

func (w *jsiiProxy_WasendClient) validateSendMessageParameters(request *MessageRequest) error {
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

func (w *jsiiProxy_WasendClient) validateStopSessionParameters(sessionId *string) error {
	if sessionId == nil {
		return fmt.Errorf("parameter sessionId is required, but nil was provided")
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


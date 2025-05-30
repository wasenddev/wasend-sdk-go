//go:build !no_runtime_type_checking

package wasendcore

import (
	"fmt"
)

func validateWasendUtils_FormatPhoneNumberParameters(phoneNumber *string) error {
	if phoneNumber == nil {
		return fmt.Errorf("parameter phoneNumber is required, but nil was provided")
	}

	return nil
}

func validateWasendUtils_IsValidPhoneNumberParameters(phoneNumber *string) error {
	if phoneNumber == nil {
		return fmt.Errorf("parameter phoneNumber is required, but nil was provided")
	}

	return nil
}


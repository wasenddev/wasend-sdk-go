//go:build no_runtime_type_checking

package wasendcore

// Building without runtime type checking enabled, so all the below just return nil

func validateWasendUtils_FormatPhoneNumberParameters(phoneNumber *string) error {
	return nil
}

func validateWasendUtils_IsValidPhoneNumberParameters(phoneNumber *string) error {
	return nil
}


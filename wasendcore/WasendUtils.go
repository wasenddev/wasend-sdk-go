package wasendcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/wasenddev/wasend-sdk-go/wasendcore/jsii"
)

// Utility functions for the Wasend SDK.
type WasendUtils interface {
}

// The jsii proxy struct for WasendUtils
type jsiiProxy_WasendUtils struct {
	_ byte // padding
}

func NewWasendUtils() WasendUtils {
	_init_.Initialize()

	j := jsiiProxy_WasendUtils{}

	_jsii_.Create(
		"@wasend/core.WasendUtils",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewWasendUtils_Override(w WasendUtils) {
	_init_.Initialize()

	_jsii_.Create(
		"@wasend/core.WasendUtils",
		nil, // no parameters
		w,
	)
}

// Format a phone number to international format.
//
// Returns: Formatted phone number.
func WasendUtils_FormatPhoneNumber(phoneNumber *string, countryCode *string) *string {
	_init_.Initialize()

	if err := validateWasendUtils_FormatPhoneNumberParameters(phoneNumber); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"@wasend/core.WasendUtils",
		"formatPhoneNumber",
		[]interface{}{phoneNumber, countryCode},
		&returns,
	)

	return returns
}

// Check if a phone number is valid.
//
// Returns: Whether the phone number is valid.
func WasendUtils_IsValidPhoneNumber(phoneNumber *string) *bool {
	_init_.Initialize()

	if err := validateWasendUtils_IsValidPhoneNumberParameters(phoneNumber); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@wasend/core.WasendUtils",
		"isValidPhoneNumber",
		[]interface{}{phoneNumber},
		&returns,
	)

	return returns
}


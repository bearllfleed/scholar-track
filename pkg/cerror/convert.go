package cerror

import "errors"

func CheckIsDefinedUserError(err error) bool {
	if errors.Is(err, ErrUserHasExists) || errors.Is(err, ErrUserNotFound) || errors.Is(err, ErrEmailHasExists) ||
		errors.Is(err, ErrPhoneHasExists) || errors.Is(err, ErrPasswordWrong) || errors.Is(err, ErrUserDisabled) ||
		errors.Is(err, ErrNotLogin) || errors.Is(err, ErrPermissionDenied) || errors.Is(err, ErrTokenInvalid) ||
		errors.Is(err, ErrTokenExpired) || errors.Is(err, ErrTokenNotMatch) {
		return true
	}
	return false
}

func CheckIsDefinedRoleError(err error) bool {
	if errors.Is(err, ErrRoleHasExists) || errors.Is(err, ErrRoleNotExists) || errors.Is(err, ErrParentRoleNotExists) ||
		errors.Is(err, ErrInvalidRole) {
		return true
	}
	return false
}

func CheckIsDefinedApiError(err error) bool {
	if errors.Is(err, ErrApiAlreadyExists) || errors.Is(err, ErrApiNotFound) || errors.Is(err, ErrHasSameApi) ||
		errors.Is(err, ErrCasbinUpdateFailed) {
		return true
	}
	return false
}

func CheckIsDefinedCategoryError(err error) bool {
	if errors.Is(err, ErrCategoryNameAlreadyExists) || errors.Is(err, ErrCategoryNotFound) {
		return true
	}
	return false
}

func CheckIsDefinedMenuError(err error) bool {
	if errors.Is(err, ErrMenuNotFound) || errors.Is(err, ErrMenuHasExists) || errors.Is(err, ErrMenuParentNotFound) ||
		errors.Is(err, ErrMenuInvalid) {
		return true
	}
	return false
}

func CheckIsDefinedFileError(err error) bool {
	if errors.Is(err, ErrFileNotFound) {
		return true
	}
	return false
}

func CheckIsDictionaryError(err error) bool {
	if errors.Is(err, ErrDictionaryNotFound) || errors.Is(err, ErrDictionaryHasExists) ||
		errors.Is(err, ErrDictionaryDetailNotFound) || errors.Is(err, ErrDictionaryDetailHasExists) {
		return true
	}
	return false
}

func CheckIsAchieveError(err error) bool {
	if errors.Is(err, ErrAchieveHasExists) || errors.Is(err, ErrAchieveHasExists) {
		return true
	}
	return false
}

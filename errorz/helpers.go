package errorz

import (
	"fmt"
	"strings"
)

func NewNotFound() *ApiError {
	return &ApiError{
		AppCode: NotFoundCode,
		Message: NotFoundMessage,
		Status:  NotFoundStatus,
	}
}

func NewUnhandled(cause error) *ApiError {
	return &ApiError{
		AppCode: UnhandledCode,
		Message: UnhandledMessage,
		Status:  UnhandledStatus,
		Cause:   cause,
	}
}

func NewEntityCanNotBeDeleted() *ApiError {
	return &ApiError{
		AppCode: EntityCanNotBeDeletedCode,
		Message: EntityCanNotBeDeletedMessage,
		Status:  EntityCanNotBeDeletedStatus,
	}
}

func NewEntityCanNotBeDeletedFrom(err error) *ApiError {
	return &ApiError{
		AppCode:     EntityCanNotBeDeletedCode,
		Message:     EntityCanNotBeDeletedMessage,
		Status:      EntityCanNotBeDeletedStatus,
		Cause:       err,
		AppendCause: true,
	}
}

func NewEntityCanNotBeUpdatedFrom(err error) *ApiError {
	return &ApiError{
		AppCode:     EntityCanNotBeUpdatedCode,
		Message:     EntityCanNotBeUpdatedMessage,
		Status:      EntityCanNotBeUpdatedStatus,
		Cause:       err,
		AppendCause: true,
	}
}

func NewFieldApiError(fieldError *FieldError) *ApiError {
	return &ApiError{
		AppCode:     InvalidFieldCode,
		Message:     InvalidFieldMessage,
		Status:      InvalidFieldStatus,
		Cause:       fieldError,
		AppendCause: true,
	}
}

func NewCouldNotValidate(err error) *ApiError {
	return &ApiError{
		AppCode: CouldNotValidateCode,
		Message: CouldNotValidateMessage,
		Status:  CouldNotValidateStatus,
		Cause:   err,
	}
}

// NewUnauthorized represents a generic unauthorized request that conveys no additional token status
func NewUnauthorized() *ApiError {
	return &ApiError{
		AppCode: UnauthorizedCode,
		Message: UnauthorizedMessage,
		Status:  UnauthorizedStatus,
	}
}

// NewUnauthorizedTokensMissing represents an unauthorized request due to a lack of any supported security token being provided
func NewUnauthorizedTokensMissing() *ApiError {
	return &ApiError{
		AppCode: UnauthorizedCode,
		Message: UnauthorizedMessage,
		Status:  UnauthorizedStatus,
		Headers: map[string][]string{
			//single string value separate by commas due to OpenAPI 2.0 header limitations (1 value for 1 header)
			//CSV values for www-authenticate allowed per RFCs.
			"WWW-Authenticate": {
				`zt-session realm="zt-session" error="missing" error_description="no token was provided"` + "," +
					`Bearer realm="openziti-oidc" error="missing" error_description="no token was provided"`,
			},
		},
	}
}

// NewUnauthorizedOidcExpired represents an unauthorized request that the provided OIDC token has expired
func NewUnauthorizedOidcExpired() *ApiError {
	return &ApiError{
		AppCode: UnauthorizedCode,
		Message: UnauthorizedMessage,
		Status:  UnauthorizedStatus,
		Headers: map[string][]string{
			"WWW-Authenticate": {
				`Bearer realm="openziti-oidc" error="expired" error_description="token expired"`,
			},
		},
	}
}

// NewUnauthorizedOidcInvalid represents an unauthorized request that the provided OIDC token is invalid
func NewUnauthorizedOidcInvalid() *ApiError {
	return &ApiError{
		AppCode: UnauthorizedCode,
		Message: UnauthorizedMessage,
		Status:  UnauthorizedStatus,
		Headers: map[string][]string{
			"WWW-Authenticate": {
				`Bearer realm="openziti-oidc" error="invalid" error_description="token is invalid"`,
			},
		},
	}
}

// NewUnauthorizedZtSessionInvalid represents an unauthorized request that the provided legacy (zt-session) token is invalid
func NewUnauthorizedZtSessionInvalid() *ApiError {
	return &ApiError{
		AppCode: UnauthorizedCode,
		Message: UnauthorizedMessage,
		Status:  UnauthorizedStatus,
		Headers: map[string][]string{
			"WWW-Authenticate": {
				`zt-session realm="zt-session" error="invalid" error_description="token is invalid"`,
			},
		},
	}
}

// NewUnauthorizedSecondaryTokenMissing represents an unauthorized request that the required additional JWT token, ext-jwt-signers configuration, is missing
func NewUnauthorizedSecondaryTokenMissing(extJwtIds, issuers []string) *ApiError {

	extJwtIdsCsv := strings.Join(extJwtIds, "|")
	issuersCsv := strings.Join(issuers, "|")

	return &ApiError{
		AppCode: UnauthorizedCode,
		Message: UnauthorizedMessage,
		Status:  UnauthorizedStatus,
		Headers: map[string][]string{
			"WWW-Authenticate": {
				fmt.Sprintf(`Bearer realm="openziti-secondary-ext-jwt" error="missing" error_description="this request requires an additional bearer token" id="%s" issuer="%s"`, extJwtIdsCsv, issuersCsv),
			},
		},
	}
}

// NewUnauthorizedSecondaryTokenExpired represents an unauthorized request that the required additional JWT token, ext-jwt-signers configuration, is expired
func NewUnauthorizedSecondaryTokenExpired(extJwtIds, issuers []string) *ApiError {
	extJwtIdsCsv := strings.Join(extJwtIds, "|")
	issuersCsv := strings.Join(issuers, "|")

	return &ApiError{
		AppCode: UnauthorizedCode,
		Message: UnauthorizedMessage,
		Status:  UnauthorizedStatus,
		Headers: map[string][]string{
			"WWW-Authenticate": {
				fmt.Sprintf(`Bearer realm="openziti-secondary-ext-jwt" error="expired" error_description="this request requires an additional bearer token that has expired" id="%s" issuer="%s"`, extJwtIdsCsv, issuersCsv),
			},
		},
	}
}

// NewUnauthorizedSecondaryTokenInvalid represents an unauthorized request that the required additional JWT token, ext-jwt-signers configuration, is invalid
func NewUnauthorizedSecondaryTokenInvalid(extJwtIds, issuers []string) *ApiError {
	extJwtIdsCsv := strings.Join(extJwtIds, "|")
	issuersCsv := strings.Join(issuers, "|")

	return &ApiError{
		AppCode: UnauthorizedCode,
		Message: UnauthorizedMessage,
		Status:  UnauthorizedStatus,
		Headers: map[string][]string{
			"WWW-Authenticate": {
				fmt.Sprintf(`Bearer realm="openziti-secondary-ext-jwt" error="invalid" error_description="this request requires an additional bearer token that is invalid" id="%s" issuer="%s"`, extJwtIdsCsv, issuersCsv),
			},
		},
	}
}

func NewInvalidFilter(cause error) *ApiError {
	return &ApiError{
		AppCode:     InvalidFilterCode,
		Message:     InvalidFilterMessage,
		Status:      InvalidFilterStatus,
		Cause:       cause,
		AppendCause: true,
	}
}

func NewInvalidPagination(err error) *ApiError {
	return &ApiError{
		AppCode:     InvalidPaginationCode,
		Message:     InvalidPaginationMessage,
		Status:      InvalidPaginationStatus,
		Cause:       err,
		AppendCause: true,
	}
}

func NewInvalidSort(err error) *ApiError {
	return &ApiError{
		AppCode:     InvalidSortCode,
		Message:     InvalidSortMessage,
		Status:      InvalidSortStatus,
		Cause:       err,
		AppendCause: true,
	}
}

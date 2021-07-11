package error_code

import "github.com/goools/tools/errorx"

func createErrorHandler(code int) func(error) error {
	return func(err error) error {
		return errorx.NewError(code, err)
	}
}

var (
	NewParameterInvalid = errorx.CreateErrorFuncHandler(CustomForbiddenParameterInvalid)
)

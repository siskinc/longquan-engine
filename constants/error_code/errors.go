package error_code

import (
	"fmt"

	"github.com/goools/tools/errorx"
)

var (
	NewParameterInvalid = errorx.CreateErrorFuncHandler(CustomForbiddenParameterInvalid)
)

var (
	ParameterInvalidIDError = NewParameterInvalid(fmt.Errorf("request id is invalid"))
)

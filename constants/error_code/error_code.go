package error_code

// 通用错误
const (
	CustomForbidden                 = 400001
	CustomForbiddenParameterInvalid = 400002
)

// 未知错误
const (
	CustomForbiddenNotFoundNamespace   = 401001 // 未查找到命名空间
	CustomForbiddenNotFoundPropertySet = 401002 // 未查找到属性集
)

// 冲突错误
const (
	CustomForbiddenConflictPropertySet = 402001
	CustomForbiddenConflictProperty    = 402002
)

package reason

var (
	InternalServerError = "internal server error"
	RequestFormError    = "request format is not valid"
)

var (
	UserAlreadyExist    = "user already exist."
	UserNotFound        = "user not exist."
	FailedLogin         = "failed to login, check you email or password."
	RegisterFailed      = "failed to register user."
	Unauthorized        = "unauthorized, invalid token."
	InvalidRefreshToken = "invalid refresh token."
)

var (
	CategoryNotFound        = "category not found"
	CategoryCannotCreate    = "cannot create category"
	CategoryCannotBrowse    = "cannot browse category"
	CategoryCannotUpdate    = "cannot update category"
	CategoryCannotDelete    = "cannot delete category"
	CategoryCannotGetDetail = "cannot get detail"
)

var (
	ProductNotFound        = "product not found"
	ProductCannotCreate    = "cannot create product"
	ProductCannotBrowse    = "cannot browse product"
	ProductCannotUpdate    = "cannot update product"
	ProductCannotDelete    = "cannot delete product"
	ProductCannotGetDetail = "cannot get detail"
)

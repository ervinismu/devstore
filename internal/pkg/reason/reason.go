package reason

var (
	InternalServerError = "internal server error"
	RequestFormError    = "request format is not valid"
)

// session
var (
	UserAlreadyExist   = "user already exist."
	RegisterFailed     = "failed to register user."
	Unauthorized       = "unauthorized request."
	InvalidAccess      = "you don't have access to this resource."
	FailedLogin        = "failed login, please check your email or password."
	FailedLogout       = "failed logout."
	FailedRefreshToken = "failed to refresh token, please check your token." //nolint
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

var (
	InsufficientStock = "insufficient stock"
	FailedAddCart     = "failed to add cart"
	FailedCheckout    = "failed to checkout order"
)

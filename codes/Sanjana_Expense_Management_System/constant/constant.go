package constant

const (
	AllowOrigins = "*"

	POST = "POST"
	GET  = "GET"

	Origin        = "Origin"
	ContextType   = "Context-Type"
	Authorization = "Authorization"
)

const (
	TableUsers    = "users"
	FieldName     = "name"
	FieldPassword = "password"

	TableExpenses = "expenses"
	FieldCategory = "category"
)

const (
	PasswordRegexp = `/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$/`
)

const (
	Token  = "token"
	Expiry = "exp"
)

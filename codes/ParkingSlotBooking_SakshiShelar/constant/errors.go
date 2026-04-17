package constant

//db migration errors
const (
	DbMigrationFailed = "Failed to migrate to db: %s"
)

//validation errors
const (
	FieldRequired = "%s field is required"
)

//login errors
const (
	ErrorFetchingFromDb     = "Error fetching data from db: %s"
	InvalidUsernamePassword = "Invalid username or password: %s"
	PasswordMismatch        = "Passwords do not match"
	TokenGenerationFailed   = "Failed to generate tpken"
)

//request error
const (
	UnexpectedValue     = "Unexpected values for request"
	InvalidInputPayload = "Invalid input payload"

	UserNotFound         = "User not found"
	AuthenticationFailed = "Authentication failed"

	InternalServer = "Internal Server Error"

	UserUnauthorized = "Unauthorized"
	HeaderMissing    = "Header Not Found"

	TokenInvalidated = "token is invalidated"
)

//logout errors
const (
	RedisInitFailed      = "Failed to connect to redis: %s"
	FailedToSetInRedis   = "failed to set in redis"
	FailedToDelFromRedis = "Failed to delete from redis"

	UserAlreadyLoggedOut = "the token has expired. please generate a new token."
)

//json marshal unmarshal erros
const (
	JsonMarshalFailed   = "Failed to marshal data"
	JsonUnmarshalFailed = "Failed to unmarshal data"
)

//available slots error
const (
	NoFreeSlots    = "No free slots available"
	StatusNotFound = "Status Not Found"
)

//user Bookings error
const (
	NoBookingsFound = "No bookings found for user"
)

//cancel booking error
const (
	BookingNotFound    = "Given booking not found for user"
	AlreadyCancelled   = "Booking is already cancelled"
	FailedToCancel     = "Failed to cancel booking"
	UserIsUnauthorized = "User is unauthorized to access this booking"
	NoUpdate           = "no update"

)

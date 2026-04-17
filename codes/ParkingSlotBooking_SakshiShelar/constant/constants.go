package constant

import "time"

// routes
const (
	AuthPrefix     = "/auth"
	Login          = "/login"
	Logout         = "/logout"
	AvailableSlots = "/slots/available"

	BookingsPrefix = "/bookings"
	CreateBookings = "/create"
	UserBookings   = "/my-bookings"
	CancelBookings = "/cancel/:id"
)

// port
const (
	RunningOn        = "Running Server on port : %v"
	PortDefaultValue = 8080
)

// common fields to set
const (
	Token    = "token"
	RedisKey = "redisKey"
	Email    = "email"
	Expiry   = "exp"
)
const (
	DSNString = "host=%s port=%s dbname=%s user=%s password=%s TimeZone=%s"
)

// redis contants
const (
	RedisAddress      = "localhost:6379"
	RedisPassword     = ""
	RedisDb           = 0
	RedisPoolSize     = 10
	RedisMinIdleConns = 5
	RedisDialTimeout  = 5 * time.Second
	RedisReadTimeout  = 3 * time.Second
	RedisWriteTimeout = 3 * time.Second
)

// db migration success
const (
	DbMigrationSuccess = "DB Migration performed successfully"
)

// jwt secret key and constants
const SecretKey = "Apr/meTe4sxpBwxb36ISTRNnHc4y+Y34KjQ/ntwB1Kw="

const (
	Authorization = "Authorization"
	Bearer        = "Bearer "
	Redis         = "redis"
	Header        = "Header"
)

// table name constanst
const (
	UsersTable    = "users"
	SlotsTable    = "slots"
	BookingsTable = "bookings"
)

// login constants
const (
	EmailField       = "email= ?"
	UserLoginSuccess = "user logged in successfully"
)

// logout constants
const UserLogoutSuccess = "user logged out successfully"

// available slots constants
const (
	StatusField         = "status= ?"
	SlotFree            = "FREE"
	SLotsFetchedSuccess = "slots fetched successfully"
	AvailableSlotsKey   = "available_slots"
)

// bookings constants
const (
	BookingsFetchedSuccess = "bookings fetched successfully"
)

// cancel booking constants
const (
	BookingsIdField = "id= ?"
	SlotsCancelled = "CANCELLED"
	BookingCancelledSuccess = "Booking cancelled successfully"
)

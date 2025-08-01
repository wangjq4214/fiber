package basicauth

import (
	"crypto/subtle"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/utils/v2"
)

// Config defines the config for middleware.
type Config struct {
	// Next defines a function to skip this middleware when returned true.
	//
	// Optional. Default: nil
	Next func(c fiber.Ctx) bool

	// Users defines the allowed credentials
	//
	// Required. Default: map[string]string{}
	Users map[string]string

	// Authorizer defines a function you can pass
	// to check the credentials however you want.
	// It will be called with a username and password
	// and is expected to return true or false to indicate
	// that the credentials were approved or not.
	//
	// Optional. Default: nil.
	Authorizer func(string, string) bool

	// Unauthorized defines the response body for unauthorized responses.
	// By default it will return with a 401 Unauthorized and the correct WWW-Auth header
	//
	// Optional. Default: nil
	Unauthorized fiber.Handler

	// Realm is a string to define realm attribute of BasicAuth.
	// the realm identifies the system to authenticate against
	// and can be used by clients to save credentials
	//
	// Optional. Default: "Restricted".
	Realm string

	// Charset defines the value for the charset parameter in the
	// WWW-Authenticate header. According to RFC 7617 clients can use
	// this value to interpret credentials correctly.
	//
	// Optional. Default: "UTF-8".
	Charset string

	// StorePassword determines if the plaintext password should be stored
	// in the context for later retrieval via PasswordFromContext.
	//
	// Optional. Default: false.
	StorePassword bool
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	Next:          nil,
	Users:         map[string]string{},
	Realm:         "Restricted",
	Charset:       "UTF-8",
	StorePassword: false,
	Authorizer:    nil,
	Unauthorized:  nil,
}

// Helper function to set default values
func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	// Set default values
	if cfg.Next == nil {
		cfg.Next = ConfigDefault.Next
	}
	if cfg.Users == nil {
		cfg.Users = ConfigDefault.Users
	}
	if cfg.Realm == "" {
		cfg.Realm = ConfigDefault.Realm
	}
	if cfg.Charset == "" {
		cfg.Charset = ConfigDefault.Charset
	}
	if cfg.Authorizer == nil {
		cfg.Authorizer = func(user, pass string) bool {
			userPwd, exist := cfg.Users[user]
			return exist && subtle.ConstantTimeCompare(utils.UnsafeBytes(userPwd), utils.UnsafeBytes(pass)) == 1
		}
	}
	if cfg.Unauthorized == nil {
		cfg.Unauthorized = func(c fiber.Ctx) error {
			header := "Basic realm=" + strconv.Quote(cfg.Realm)
			if cfg.Charset != "" {
				header += ", charset=" + strconv.Quote(cfg.Charset)
			}
			c.Set(fiber.HeaderWWWAuthenticate, header)
			c.Set(fiber.HeaderCacheControl, "no-store")
			c.Set(fiber.HeaderVary, fiber.HeaderAuthorization)
			return c.SendStatus(fiber.StatusUnauthorized)
		}
	}
	return cfg
}

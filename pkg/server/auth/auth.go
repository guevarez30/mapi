package auth

import (
	"fmt"
	"os"
	"strings"

	"github.com/clerkinc/clerk-sdk-go/clerk"

	"github.com/gofiber/fiber/v2"
)

type Auth struct {
	Client clerk.Client
}

func Initiatlize() Auth {
	CLERK_KEY := os.Getenv("CLERK_KEY")
	client, err := clerk.NewClient(CLERK_KEY)
	if err != nil {
		panic("Falied to connect to auth service")
	}
	auth := Auth{
		Client: client,
	}
	return auth
}

func (auth Auth) Protect(c *fiber.Ctx) error {

	//Skip auth for dev purposes
	GO_ENV := os.Getenv("GO_ENV")
	if GO_ENV == "development" {
		return c.Next()
	}

	headers := c.GetReqHeaders()
	sessionToken := headers["Authorization"]
	sessionToken = strings.TrimPrefix(sessionToken, "Bearer ")

	// Validate token
	session, err := auth.Client.VerifyToken(sessionToken)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized)
	}

	// Get user information
	user, err := auth.Client.Users().Read(session.Claims.Subject)
	if err != nil {
		return fiber.NewError(fiber.StatusForbidden)
	}
	fmt.Println(user.EmailAddresses)

	return c.Next()
}

package main

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hamza1312/chax/schema"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		panic("Skill issue")
	}

	e := db.AutoMigrate(&schema.User{})
	if e != nil {
		panic(e)
	}
	app := fiber.New()
	// Gets profile with authentication header
	app.Post("/accounts/profile", func(c *fiber.Ctx) error {
		// Get the token from the header
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.JSON(&fiber.Map{
				"success": false,
				"message": "No token provided",
			})
		}
		// Verify the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			return c.JSON(&fiber.Map{
				"success": false,
				"message": "Invalid token",
			})
		}
		// Get the user id from the token
		claims := token.Claims.(jwt.MapClaims)
		var user schema.User

		if err := db.Where("id = ?", claims["id"]).First(&user).Error; err != nil {
			return c.JSON(&fiber.Map{
				"success": false,
				"message": "Invalid token",
			})
		}
		return c.JSON(&fiber.Map{
			"success": true,
			"user":    user,
		})
	})
	app.Post("/accounts/login", func(c *fiber.Ctx) error {
		var Body struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := c.BodyParser(&Body); err != nil {
			return err
		}
		var user schema.User
		if err := db.Where("username = ?", Body.Username).First(&user).Error; err != nil {
			return c.JSON(&fiber.Map{
				"success": false,
				"message": "Invalid username or password",
			})
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Body.Password)); err != nil {
			return c.JSON(&fiber.Map{
				"success": false,
				"message": "Invalid username or password",
			})
		}
		// Generate a token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id": user.ID,
		})
		tokenString, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		return c.JSON(&fiber.Map{
			"success": true,
			"token":   tokenString,
		})
	})

	app.Post("/accounts/new", func(c *fiber.Ctx) error {
		var Body struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := c.BodyParser(&Body); err != nil {
			return err
		}
		// Verify that the username exists
		if len(Body.Username) < 3 {
			return c.JSON(&fiber.Map{
				"success": false,
				"message": "Username must be at least 3 characters long",
			})
		}
		// Verify that the password exists
		if len(Body.Password) < 6 {
			return c.JSON(&fiber.Map{
				"success": false,
				"message": "Password must be at least 6 characters long",
			})
		}
		if strings.Contains(strings.TrimSpace(Body.Username), " ") {
			return c.JSON(&fiber.Map{
				"success": false,
				"message": "Username cannot contain spaces",
			})
		}
		// Encrypt the password
		pass, err := bcrypt.GenerateFromPassword([]byte(Body.Password), 10)
		if err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
			})
		}
		user := schema.User{
			Username: Body.Username,
			Password: string(pass),
		}

		if err := db.Create(&user).Error; err != nil {
			return err
		}
		// Generate a token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id": user.ID,
		})
		tokenString, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		return c.JSON(&fiber.Map{
			"success": true,
			"token":   tokenString,
		})
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"success": true,
		})
	})

	log.Fatal(app.Listen(":3000"))
}

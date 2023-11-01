package handlers

import (
	"time"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/lmfundora/netflixProject/config"

	"github.com/gofiber/fiber/v2"
	"github.com/lmfundora/netflixProject/database"
	"github.com/lmfundora/netflixProject/models"
)

func LogInUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	//fmt.Println(loginRequest.Name)
	result := database.DB.Db.Where("email = ? AND password = ?", user.Email, user.Password).First(&user)
	if result.Error != nil {
		return NotFound(c)
	}

	day := time.Hour * 24

	// Create the JWT claims, which includes the user ID and expiry time
	claims := jtoken.MapClaims{
		"ID":	user.ID,
		"name": user.Name,
		"phone": user.Phone,
		"email": user.Email,
		"exp":   time.Now().Add(day * 1).Unix(),
	}

	// Create token
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Return the token
	return c.JSON(models.LoginResponse{
		Token: t,
	})
}

func ListUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.DB.Db.Find((&users))

	return c.Status(200).JSON(users)
}

func CreateUsers(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	database.DB.Db.Create(&user)

	return c.Status(200).JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	user := models.User{}
	id := c.Params("id")

	result := database.DB.Db.Where("id = ?", id).Delete(&user)
	if result.Error != nil {
		return NotFound(c)
	}

	return ListUsers(c)
}

func UpdateUser(c *fiber.Ctx) error {
	user := new(models.User)
	id := c.Params("id")

	// Parsing the request body
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).SendString(err.Error())
	}

	// Write updated values to the database
	result := database.DB.Db.Model(&user).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return EditUser(c)
	}

	return ShowUser(c)
}

func EditUser(c *fiber.Ctx) error {
	user := models.User{}
	id := c.Params("id")

	result := database.DB.Db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return NotFound(c)
	}

	return c.JSON(fiber.Map{
		"Title":    "Edit User",
		"Subtitle": "Edit your user",
		"User":     user,
	})
}

func ShowUser(c *fiber.Ctx) error {
	user := models.User{}
	id := c.Params("id")

	database.DB.Db.Where("id = ?", id).First(&user)

	return c.JSON(fiber.Map{
		"Title": "Single User",
		"User":  user,
	})
}

func NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).SendFile("./public/404.html")
}

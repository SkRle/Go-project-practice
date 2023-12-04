package controllers

import (
	"go-project/database"
	m "go-project/models"
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func FindIDNameLastName(db *gorm.DB) *gorm.DB {
	return db.Where("employee_id = ? OR name = ? OR last_name = ?", 100)
}

func GetProfileUser(c *fiber.Ctx) error {

	db := database.DBConn
	var users []m.UserProfile

	db.Find(&users)
	return c.Status(200).JSON(users)
}

func AddUserProfile(c *fiber.Ctx) error {
	db := database.DBConn
	var user m.UserProfile

	emailFormat := `\S+@\S+\.\S+`
	namePattern := "^[A-Za-z]+$"

	regExpEmail := regexp.MustCompile(emailFormat)
	regExpName := regexp.MustCompile(namePattern)
	regEXPLastName := regexp.MustCompile(namePattern) //

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if !regExpEmail.MatchString(user.Email) || !regExpName.MatchString(user.Name) || !regEXPLastName.MatchString(user.LastName) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "data is not valid",
		})
	}

	validate := validator.New()
	errors := validate.Struct(user)
	if errors != nil {
		log.Print("Error")
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}
	//if have no any problem then send json to front
	db.Create(&user)
	return c.Status(201).JSON(user)
}

func UpdateUserProfile(c *fiber.Ctx) error {
	db := database.DBConn
	var users m.UserProfile
	id := c.Params("id")

	if err := c.BodyParser(&users); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&users)
	return c.Status(200).JSON(users)
}

func RemoveUserFile(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var users m.UserProfile

	result := db.Delete(&users, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func GetUserProfileGroup(c *fiber.Ctx) error {
	db := database.DBConn
	var user []m.UserProfile
	var sum_GenZ int
	var sum_GenY int
	var sum_GenX int
	var sum_BabyBoomer int
	var sum_GI_Generation int

	db.Find(&user)

	var dataResults []m.UsersRes
	for _, v := range user { //1 inet 112 //2 inet1 113
		Generation := ""
		if v.Age > 75 {
			Generation = "G.I. Generation"
			sum_GI_Generation++
		} else if v.Age >= 57 && v.Age <= 75 {
			Generation = "Baby Boomer"
			sum_BabyBoomer++
		} else if v.Age >= 42 && v.Age <= 56 {
			Generation = "GenX"
			sum_GenX++
		} else if v.Age >= 24 && v.Age <= 41 {
			Generation = "GenY"
			sum_GenY++
		} else if v.Age < 24 {
			Generation = "GenZ"
			sum_GenZ++
		}

		d := m.UsersRes{
			Employee_id: v.Employee_id,
			Name:        v.Name,
			LastName:    v.LastName,
			Birthday:    v.Birthday,
			Age:         v.Age,
			Email:       v.Email,
			Tel:         v.Tel,
			Generation:  Generation,
		}
		dataResults = append(dataResults, d)
		// sumAmount += v.Amount
	}

	r := m.ResultData{
		All_users:    len(user),
		Data:         dataResults,
		GenZ:         sum_GenZ,
		GenY:         sum_GenY,
		GenX:         sum_GenX,
		BabyBoomer:   sum_BabyBoomer,
		GIGeneration: sum_GI_Generation,
	}
	return c.Status(200).JSON(r)
}

func SearchValue(c *fiber.Ctx) error {
	db := database.DBConn
	var users []m.UserProfile

	searchValue := c.Params("value")

	db.Where("employee_id = ? OR name = ? OR last_name = ?", searchValue, searchValue, searchValue).Find(&users)

	if len(users) == 0 {
		return c.Status(404).SendString("Not Found Any Data")
	}
	return c.Status(200).JSON(users)
}

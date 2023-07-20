package lead

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/karimosman181/simpleGoCRM/database"
)

type Lead struct {
	gorm.Model
	Name    string
	Company string
	Email   string
	Phone   string
}

/**
 *
 * get all leads record
 **/
func GetLeads(c *fiber.Ctx) {
	//get db connection
	db := database.DBConn

	var leads []Lead

	//get records from DB
	db.Find(&leads)

	//return response
	c.JSON(leads)
}

/**
 *
 * get lead by id
 **/
func GetLead(c *fiber.Ctx) {

	//get id from params
	id := c.Params("id")

	//get db connection
	db := database.DBConn

	var lead Lead

	//search table for lead
	db.Find(&lead, id)

	//return response
	c.JSON(lead)
}

/**
 *
 * create new lead
 **/
func NewLeads(c *fiber.Ctx) {}

/**
 *
 * delete lead by id
 **/
func DeleteLeads(c *fiber.Ctx) {}

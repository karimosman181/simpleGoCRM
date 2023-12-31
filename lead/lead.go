package lead

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/karimosman181/simpleGoCRM/database"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
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
func NewLeads(c *fiber.Ctx) {
	//get db connection
	db := database.DBConn

	//define lead
	lead := new(Lead)

	//check for error after parsing body
	if err := c.BodyParser(lead); err != nil {

		//return error
		c.Status(503).Send(err)
		return
	}

	//save new lead
	db.Create(&lead)
	c.JSON(lead)
}

/**
 *
 * delete lead by id
 **/
func DeleteLeads(c *fiber.Ctx) {
	//get id from params
	id := c.Params("id")

	//get db connection
	db := database.DBConn

	var lead Lead

	//search for lead if exists
	db.First(&lead, id)

	//check if exits
	if lead.Name == "" {
		//return error
		c.Status(404).Send("Lead not Found")
		return
	}

	//delete lead
	db.Delete(&lead)

	//return success
	c.Send("Lead deleted successfully")
}

/**
 *
 * update lead
 **/
func UpdateLeads(c *fiber.Ctx) {

	//get id from params
	id := c.Params("id")

	//get db connection
	db := database.DBConn

	var lead Lead

	//search for lead if exists
	db.First(&lead, id)

	//check if exits
	if lead.Name == "" {
		//return error
		c.Status(404).Send("Lead not Found")
		return
	}

	//define lead
	uplead := new(Lead)

	//check for error after parsing body
	if err := c.BodyParser(uplead); err != nil {

		//return error
		c.Status(503).Send(err)
		return
	}

	if uplead.Name != "" {
		lead.Name = uplead.Name
	}
	if uplead.Company != "" {
		lead.Company = uplead.Company
	}
	if uplead.Phone != "" {
		lead.Phone = uplead.Phone
	}
	if uplead.Email != "" {
		lead.Email = uplead.Email
	}

	//update record
	db.Save(&lead)

	c.JSON(lead)

}

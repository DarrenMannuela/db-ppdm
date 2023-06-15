package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRestContact(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rest_contact")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rest_contact

	for rows.Next() {
		var rest_contact dto.Rest_contact
		if err := rows.Scan(&rest_contact.Restriction_id, &rest_contact.Restriction_version, &rest_contact.Business_associate_id, &rest_contact.Contact_obs_no, &rest_contact.Active_ind, &rest_contact.Address_obs_no, &rest_contact.Address_source, &rest_contact.Effective_date, &rest_contact.Expiry_date, &rest_contact.Phone_num, &rest_contact.Phone_num_id, &rest_contact.Ppdm_guid, &rest_contact.Primary_contact_ind, &rest_contact.Remark, &rest_contact.Source, &rest_contact.Row_changed_by, &rest_contact.Row_changed_date, &rest_contact.Row_created_by, &rest_contact.Row_created_date, &rest_contact.Row_effective_date, &rest_contact.Row_expiry_date, &rest_contact.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rest_contact to result
		result = append(result, rest_contact)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRestContact(c *fiber.Ctx) error {
	var rest_contact dto.Rest_contact

	setDefaults(&rest_contact)

	if err := json.Unmarshal(c.Body(), &rest_contact); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rest_contact values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	rest_contact.Row_created_date = formatDate(rest_contact.Row_created_date)
	rest_contact.Row_changed_date = formatDate(rest_contact.Row_changed_date)
	rest_contact.Effective_date = formatDateString(rest_contact.Effective_date)
	rest_contact.Expiry_date = formatDateString(rest_contact.Expiry_date)
	rest_contact.Row_effective_date = formatDateString(rest_contact.Row_effective_date)
	rest_contact.Row_expiry_date = formatDateString(rest_contact.Row_expiry_date)






	rows, err := stmt.Exec(rest_contact.Restriction_id, rest_contact.Restriction_version, rest_contact.Business_associate_id, rest_contact.Contact_obs_no, rest_contact.Active_ind, rest_contact.Address_obs_no, rest_contact.Address_source, rest_contact.Effective_date, rest_contact.Expiry_date, rest_contact.Phone_num, rest_contact.Phone_num_id, rest_contact.Ppdm_guid, rest_contact.Primary_contact_ind, rest_contact.Remark, rest_contact.Source, rest_contact.Row_changed_by, rest_contact.Row_changed_date, rest_contact.Row_created_by, rest_contact.Row_created_date, rest_contact.Row_effective_date, rest_contact.Row_expiry_date, rest_contact.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRestContact(c *fiber.Ctx) error {
	var rest_contact dto.Rest_contact

	setDefaults(&rest_contact)

	if err := json.Unmarshal(c.Body(), &rest_contact); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rest_contact.Restriction_id = id

    if rest_contact.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rest_contact set restriction_version = :1, business_associate_id = :2, contact_obs_no = :3, active_ind = :4, address_obs_no = :5, address_source = :6, effective_date = :7, expiry_date = :8, phone_num = :9, phone_num_id = :10, ppdm_guid = :11, primary_contact_ind = :12, remark = :13, source = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where restriction_id = :22")
	if err != nil {
		return err
	}

	rest_contact.Row_changed_date = formatDate(rest_contact.Row_changed_date)
	rest_contact.Effective_date = formatDateString(rest_contact.Effective_date)
	rest_contact.Expiry_date = formatDateString(rest_contact.Expiry_date)
	rest_contact.Row_effective_date = formatDateString(rest_contact.Row_effective_date)
	rest_contact.Row_expiry_date = formatDateString(rest_contact.Row_expiry_date)






	rows, err := stmt.Exec(rest_contact.Restriction_version, rest_contact.Business_associate_id, rest_contact.Contact_obs_no, rest_contact.Active_ind, rest_contact.Address_obs_no, rest_contact.Address_source, rest_contact.Effective_date, rest_contact.Expiry_date, rest_contact.Phone_num, rest_contact.Phone_num_id, rest_contact.Ppdm_guid, rest_contact.Primary_contact_ind, rest_contact.Remark, rest_contact.Source, rest_contact.Row_changed_by, rest_contact.Row_changed_date, rest_contact.Row_created_by, rest_contact.Row_effective_date, rest_contact.Row_expiry_date, rest_contact.Row_quality, rest_contact.Restriction_id)
	if err != nil {
		return err
	}

	if count, err := rows.RowsAffected(); err == nil {
		if count == 0 {
			return c.Status(404).SendString("No matching record found")
		}
	} else {
		return err
	}

	return c.Status(201).JSON(rows)
}

func PatchRestContact(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rest_contact set "
	values := []interface{}{}
	i := 1
	for key, value := range updatedFields {
		query += key + " = :" + strconv.Itoa(i)
		i++
		if i <= len(updatedFields) {
			query += ", "
		}
		if key == "row_changed_date" {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDate(&date)
				value = formattedDate
			}
		} else if key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where restriction_id = :id"

	stmt, err := db.Prepare(query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	values = append(values, id)
	res, err := stmt.Exec(values...)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if count, err := res.RowsAffected(); err == nil {
		if count == 0 {
			return c.Status(404).SendString("No matching record found")
		}
	} else {
		return err
	}

	return c.JSON(res)
}

func DeleteRestContact(c *fiber.Ctx) error {
	id := c.Params("id")
	var rest_contact dto.Rest_contact
	rest_contact.Restriction_id = id

	stmt, err := db.Prepare("delete from rest_contact where restriction_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rest_contact.Restriction_id)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return c.Status(404).SendString("No matching record found")
	}

	return c.JSON(rows)
}



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetApplicBa(c *fiber.Ctx) error {
	rows, err := db.Query("select * from applic_ba")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Applic_ba

	for rows.Next() {
		var applic_ba dto.Applic_ba
		if err := rows.Scan(&applic_ba.Application_id, &applic_ba.Application_ba_id, &applic_ba.Application_ba_obs_no, &applic_ba.Active_ind, &applic_ba.Application_ba_role, &applic_ba.Contact_ba_id, &applic_ba.Description, &applic_ba.Effective_date, &applic_ba.Expiry_date, &applic_ba.Ppdm_guid, &applic_ba.Remark, &applic_ba.Source, &applic_ba.Row_changed_by, &applic_ba.Row_changed_date, &applic_ba.Row_created_by, &applic_ba.Row_created_date, &applic_ba.Row_effective_date, &applic_ba.Row_expiry_date, &applic_ba.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append applic_ba to result
		result = append(result, applic_ba)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetApplicBa(c *fiber.Ctx) error {
	var applic_ba dto.Applic_ba

	setDefaults(&applic_ba)

	if err := json.Unmarshal(c.Body(), &applic_ba); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into applic_ba values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	applic_ba.Row_created_date = formatDate(applic_ba.Row_created_date)
	applic_ba.Row_changed_date = formatDate(applic_ba.Row_changed_date)
	applic_ba.Effective_date = formatDateString(applic_ba.Effective_date)
	applic_ba.Expiry_date = formatDateString(applic_ba.Expiry_date)
	applic_ba.Row_effective_date = formatDateString(applic_ba.Row_effective_date)
	applic_ba.Row_expiry_date = formatDateString(applic_ba.Row_expiry_date)






	rows, err := stmt.Exec(applic_ba.Application_id, applic_ba.Application_ba_id, applic_ba.Application_ba_obs_no, applic_ba.Active_ind, applic_ba.Application_ba_role, applic_ba.Contact_ba_id, applic_ba.Description, applic_ba.Effective_date, applic_ba.Expiry_date, applic_ba.Ppdm_guid, applic_ba.Remark, applic_ba.Source, applic_ba.Row_changed_by, applic_ba.Row_changed_date, applic_ba.Row_created_by, applic_ba.Row_created_date, applic_ba.Row_effective_date, applic_ba.Row_expiry_date, applic_ba.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateApplicBa(c *fiber.Ctx) error {
	var applic_ba dto.Applic_ba

	setDefaults(&applic_ba)

	if err := json.Unmarshal(c.Body(), &applic_ba); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	applic_ba.Application_id = id

    if applic_ba.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update applic_ba set application_ba_id = :1, application_ba_obs_no = :2, active_ind = :3, application_ba_role = :4, contact_ba_id = :5, description = :6, effective_date = :7, expiry_date = :8, ppdm_guid = :9, remark = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where application_id = :19")
	if err != nil {
		return err
	}

	applic_ba.Row_changed_date = formatDate(applic_ba.Row_changed_date)
	applic_ba.Effective_date = formatDateString(applic_ba.Effective_date)
	applic_ba.Expiry_date = formatDateString(applic_ba.Expiry_date)
	applic_ba.Row_effective_date = formatDateString(applic_ba.Row_effective_date)
	applic_ba.Row_expiry_date = formatDateString(applic_ba.Row_expiry_date)






	rows, err := stmt.Exec(applic_ba.Application_ba_id, applic_ba.Application_ba_obs_no, applic_ba.Active_ind, applic_ba.Application_ba_role, applic_ba.Contact_ba_id, applic_ba.Description, applic_ba.Effective_date, applic_ba.Expiry_date, applic_ba.Ppdm_guid, applic_ba.Remark, applic_ba.Source, applic_ba.Row_changed_by, applic_ba.Row_changed_date, applic_ba.Row_created_by, applic_ba.Row_effective_date, applic_ba.Row_expiry_date, applic_ba.Row_quality, applic_ba.Application_id)
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

func PatchApplicBa(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update applic_ba set "
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
	query += " where application_id = :id"

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

func DeleteApplicBa(c *fiber.Ctx) error {
	id := c.Params("id")
	var applic_ba dto.Applic_ba
	applic_ba.Application_id = id

	stmt, err := db.Prepare("delete from applic_ba where application_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(applic_ba.Application_id)
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



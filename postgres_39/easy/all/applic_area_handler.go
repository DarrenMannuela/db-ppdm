package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetApplicArea(c *fiber.Ctx) error {
	rows, err := db.Query("select * from applic_area")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Applic_area

	for rows.Next() {
		var applic_area dto.Applic_area
		if err := rows.Scan(&applic_area.Application_id, &applic_area.Area_id, &applic_area.Area_type, &applic_area.Active_ind, &applic_area.Description, &applic_area.Effective_date, &applic_area.Expiry_date, &applic_area.Ppdm_guid, &applic_area.Remark, &applic_area.Source, &applic_area.Row_changed_by, &applic_area.Row_changed_date, &applic_area.Row_created_by, &applic_area.Row_created_date, &applic_area.Row_effective_date, &applic_area.Row_expiry_date, &applic_area.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append applic_area to result
		result = append(result, applic_area)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetApplicArea(c *fiber.Ctx) error {
	var applic_area dto.Applic_area

	setDefaults(&applic_area)

	if err := json.Unmarshal(c.Body(), &applic_area); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into applic_area values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	applic_area.Row_created_date = formatDate(applic_area.Row_created_date)
	applic_area.Row_changed_date = formatDate(applic_area.Row_changed_date)
	applic_area.Effective_date = formatDateString(applic_area.Effective_date)
	applic_area.Expiry_date = formatDateString(applic_area.Expiry_date)
	applic_area.Row_effective_date = formatDateString(applic_area.Row_effective_date)
	applic_area.Row_expiry_date = formatDateString(applic_area.Row_expiry_date)






	rows, err := stmt.Exec(applic_area.Application_id, applic_area.Area_id, applic_area.Area_type, applic_area.Active_ind, applic_area.Description, applic_area.Effective_date, applic_area.Expiry_date, applic_area.Ppdm_guid, applic_area.Remark, applic_area.Source, applic_area.Row_changed_by, applic_area.Row_changed_date, applic_area.Row_created_by, applic_area.Row_created_date, applic_area.Row_effective_date, applic_area.Row_expiry_date, applic_area.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateApplicArea(c *fiber.Ctx) error {
	var applic_area dto.Applic_area

	setDefaults(&applic_area)

	if err := json.Unmarshal(c.Body(), &applic_area); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	applic_area.Application_id = id

    if applic_area.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update applic_area set area_id = :1, area_type = :2, active_ind = :3, description = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where application_id = :17")
	if err != nil {
		return err
	}

	applic_area.Row_changed_date = formatDate(applic_area.Row_changed_date)
	applic_area.Effective_date = formatDateString(applic_area.Effective_date)
	applic_area.Expiry_date = formatDateString(applic_area.Expiry_date)
	applic_area.Row_effective_date = formatDateString(applic_area.Row_effective_date)
	applic_area.Row_expiry_date = formatDateString(applic_area.Row_expiry_date)






	rows, err := stmt.Exec(applic_area.Area_id, applic_area.Area_type, applic_area.Active_ind, applic_area.Description, applic_area.Effective_date, applic_area.Expiry_date, applic_area.Ppdm_guid, applic_area.Remark, applic_area.Source, applic_area.Row_changed_by, applic_area.Row_changed_date, applic_area.Row_created_by, applic_area.Row_effective_date, applic_area.Row_expiry_date, applic_area.Row_quality, applic_area.Application_id)
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

func PatchApplicArea(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update applic_area set "
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

func DeleteApplicArea(c *fiber.Ctx) error {
	id := c.Params("id")
	var applic_area dto.Applic_area
	applic_area.Application_id = id

	stmt, err := db.Prepare("delete from applic_area where application_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(applic_area.Application_id)
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



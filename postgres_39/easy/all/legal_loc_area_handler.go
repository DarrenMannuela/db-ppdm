package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLegalLocArea(c *fiber.Ctx) error {
	rows, err := db.Query("select * from legal_loc_area")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Legal_loc_area

	for rows.Next() {
		var legal_loc_area dto.Legal_loc_area
		if err := rows.Scan(&legal_loc_area.Legal_loc_id, &legal_loc_area.Location_type, &legal_loc_area.Area_id, &legal_loc_area.Area_type, &legal_loc_area.Active_ind, &legal_loc_area.Effective_date, &legal_loc_area.Expiry_date, &legal_loc_area.Ppdm_guid, &legal_loc_area.Remark, &legal_loc_area.Source, &legal_loc_area.Row_changed_by, &legal_loc_area.Row_changed_date, &legal_loc_area.Row_created_by, &legal_loc_area.Row_created_date, &legal_loc_area.Row_effective_date, &legal_loc_area.Row_expiry_date, &legal_loc_area.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append legal_loc_area to result
		result = append(result, legal_loc_area)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLegalLocArea(c *fiber.Ctx) error {
	var legal_loc_area dto.Legal_loc_area

	setDefaults(&legal_loc_area)

	if err := json.Unmarshal(c.Body(), &legal_loc_area); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into legal_loc_area values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	legal_loc_area.Row_created_date = formatDate(legal_loc_area.Row_created_date)
	legal_loc_area.Row_changed_date = formatDate(legal_loc_area.Row_changed_date)
	legal_loc_area.Effective_date = formatDateString(legal_loc_area.Effective_date)
	legal_loc_area.Expiry_date = formatDateString(legal_loc_area.Expiry_date)
	legal_loc_area.Row_effective_date = formatDateString(legal_loc_area.Row_effective_date)
	legal_loc_area.Row_expiry_date = formatDateString(legal_loc_area.Row_expiry_date)






	rows, err := stmt.Exec(legal_loc_area.Legal_loc_id, legal_loc_area.Location_type, legal_loc_area.Area_id, legal_loc_area.Area_type, legal_loc_area.Active_ind, legal_loc_area.Effective_date, legal_loc_area.Expiry_date, legal_loc_area.Ppdm_guid, legal_loc_area.Remark, legal_loc_area.Source, legal_loc_area.Row_changed_by, legal_loc_area.Row_changed_date, legal_loc_area.Row_created_by, legal_loc_area.Row_created_date, legal_loc_area.Row_effective_date, legal_loc_area.Row_expiry_date, legal_loc_area.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLegalLocArea(c *fiber.Ctx) error {
	var legal_loc_area dto.Legal_loc_area

	setDefaults(&legal_loc_area)

	if err := json.Unmarshal(c.Body(), &legal_loc_area); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	legal_loc_area.Legal_loc_id = id

    if legal_loc_area.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update legal_loc_area set location_type = :1, area_id = :2, area_type = :3, active_ind = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where legal_loc_id = :17")
	if err != nil {
		return err
	}

	legal_loc_area.Row_changed_date = formatDate(legal_loc_area.Row_changed_date)
	legal_loc_area.Effective_date = formatDateString(legal_loc_area.Effective_date)
	legal_loc_area.Expiry_date = formatDateString(legal_loc_area.Expiry_date)
	legal_loc_area.Row_effective_date = formatDateString(legal_loc_area.Row_effective_date)
	legal_loc_area.Row_expiry_date = formatDateString(legal_loc_area.Row_expiry_date)






	rows, err := stmt.Exec(legal_loc_area.Location_type, legal_loc_area.Area_id, legal_loc_area.Area_type, legal_loc_area.Active_ind, legal_loc_area.Effective_date, legal_loc_area.Expiry_date, legal_loc_area.Ppdm_guid, legal_loc_area.Remark, legal_loc_area.Source, legal_loc_area.Row_changed_by, legal_loc_area.Row_changed_date, legal_loc_area.Row_created_by, legal_loc_area.Row_effective_date, legal_loc_area.Row_expiry_date, legal_loc_area.Row_quality, legal_loc_area.Legal_loc_id)
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

func PatchLegalLocArea(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update legal_loc_area set "
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
	query += " where legal_loc_id = :id"

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

func DeleteLegalLocArea(c *fiber.Ctx) error {
	id := c.Params("id")
	var legal_loc_area dto.Legal_loc_area
	legal_loc_area.Legal_loc_id = id

	stmt, err := db.Prepare("delete from legal_loc_area where legal_loc_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(legal_loc_area.Legal_loc_id)
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



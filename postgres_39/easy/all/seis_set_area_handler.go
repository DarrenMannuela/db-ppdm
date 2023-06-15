package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisSetArea(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_set_area")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_set_area

	for rows.Next() {
		var seis_set_area dto.Seis_set_area
		if err := rows.Scan(&seis_set_area.Seis_set_subtype, &seis_set_area.Seis_set_id, &seis_set_area.Area_id, &seis_set_area.Area_type, &seis_set_area.Active_ind, &seis_set_area.Effective_date, &seis_set_area.Expiry_date, &seis_set_area.Ppdm_guid, &seis_set_area.Remark, &seis_set_area.Source, &seis_set_area.Row_changed_by, &seis_set_area.Row_changed_date, &seis_set_area.Row_created_by, &seis_set_area.Row_created_date, &seis_set_area.Row_effective_date, &seis_set_area.Row_expiry_date, &seis_set_area.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_set_area to result
		result = append(result, seis_set_area)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisSetArea(c *fiber.Ctx) error {
	var seis_set_area dto.Seis_set_area

	setDefaults(&seis_set_area)

	if err := json.Unmarshal(c.Body(), &seis_set_area); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_set_area values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	seis_set_area.Row_created_date = formatDate(seis_set_area.Row_created_date)
	seis_set_area.Row_changed_date = formatDate(seis_set_area.Row_changed_date)
	seis_set_area.Effective_date = formatDateString(seis_set_area.Effective_date)
	seis_set_area.Expiry_date = formatDateString(seis_set_area.Expiry_date)
	seis_set_area.Row_effective_date = formatDateString(seis_set_area.Row_effective_date)
	seis_set_area.Row_expiry_date = formatDateString(seis_set_area.Row_expiry_date)






	rows, err := stmt.Exec(seis_set_area.Seis_set_subtype, seis_set_area.Seis_set_id, seis_set_area.Area_id, seis_set_area.Area_type, seis_set_area.Active_ind, seis_set_area.Effective_date, seis_set_area.Expiry_date, seis_set_area.Ppdm_guid, seis_set_area.Remark, seis_set_area.Source, seis_set_area.Row_changed_by, seis_set_area.Row_changed_date, seis_set_area.Row_created_by, seis_set_area.Row_created_date, seis_set_area.Row_effective_date, seis_set_area.Row_expiry_date, seis_set_area.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisSetArea(c *fiber.Ctx) error {
	var seis_set_area dto.Seis_set_area

	setDefaults(&seis_set_area)

	if err := json.Unmarshal(c.Body(), &seis_set_area); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_set_area.Seis_set_subtype = id

    if seis_set_area.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_set_area set seis_set_id = :1, area_id = :2, area_type = :3, active_ind = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where seis_set_subtype = :17")
	if err != nil {
		return err
	}

	seis_set_area.Row_changed_date = formatDate(seis_set_area.Row_changed_date)
	seis_set_area.Effective_date = formatDateString(seis_set_area.Effective_date)
	seis_set_area.Expiry_date = formatDateString(seis_set_area.Expiry_date)
	seis_set_area.Row_effective_date = formatDateString(seis_set_area.Row_effective_date)
	seis_set_area.Row_expiry_date = formatDateString(seis_set_area.Row_expiry_date)






	rows, err := stmt.Exec(seis_set_area.Seis_set_id, seis_set_area.Area_id, seis_set_area.Area_type, seis_set_area.Active_ind, seis_set_area.Effective_date, seis_set_area.Expiry_date, seis_set_area.Ppdm_guid, seis_set_area.Remark, seis_set_area.Source, seis_set_area.Row_changed_by, seis_set_area.Row_changed_date, seis_set_area.Row_created_by, seis_set_area.Row_effective_date, seis_set_area.Row_expiry_date, seis_set_area.Row_quality, seis_set_area.Seis_set_subtype)
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

func PatchSeisSetArea(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_set_area set "
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
	query += " where seis_set_subtype = :id"

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

func DeleteSeisSetArea(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_set_area dto.Seis_set_area
	seis_set_area.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_set_area where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_set_area.Seis_set_subtype)
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



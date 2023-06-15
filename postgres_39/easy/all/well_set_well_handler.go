package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellSetWell(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_set_well")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_set_well

	for rows.Next() {
		var well_set_well dto.Well_set_well
		if err := rows.Scan(&well_set_well.Well_set_id, &well_set_well.Uwi, &well_set_well.Well_set_well_obs_no, &well_set_well.Active_ind, &well_set_well.Effective_date, &well_set_well.Expiry_date, &well_set_well.Ppdm_guid, &well_set_well.Remark, &well_set_well.Source, &well_set_well.Well_set_role, &well_set_well.Row_changed_by, &well_set_well.Row_changed_date, &well_set_well.Row_created_by, &well_set_well.Row_created_date, &well_set_well.Row_effective_date, &well_set_well.Row_expiry_date, &well_set_well.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_set_well to result
		result = append(result, well_set_well)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellSetWell(c *fiber.Ctx) error {
	var well_set_well dto.Well_set_well

	setDefaults(&well_set_well)

	if err := json.Unmarshal(c.Body(), &well_set_well); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_set_well values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	well_set_well.Row_created_date = formatDate(well_set_well.Row_created_date)
	well_set_well.Row_changed_date = formatDate(well_set_well.Row_changed_date)
	well_set_well.Effective_date = formatDateString(well_set_well.Effective_date)
	well_set_well.Expiry_date = formatDateString(well_set_well.Expiry_date)
	well_set_well.Row_effective_date = formatDateString(well_set_well.Row_effective_date)
	well_set_well.Row_expiry_date = formatDateString(well_set_well.Row_expiry_date)






	rows, err := stmt.Exec(well_set_well.Well_set_id, well_set_well.Uwi, well_set_well.Well_set_well_obs_no, well_set_well.Active_ind, well_set_well.Effective_date, well_set_well.Expiry_date, well_set_well.Ppdm_guid, well_set_well.Remark, well_set_well.Source, well_set_well.Well_set_role, well_set_well.Row_changed_by, well_set_well.Row_changed_date, well_set_well.Row_created_by, well_set_well.Row_created_date, well_set_well.Row_effective_date, well_set_well.Row_expiry_date, well_set_well.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellSetWell(c *fiber.Ctx) error {
	var well_set_well dto.Well_set_well

	setDefaults(&well_set_well)

	if err := json.Unmarshal(c.Body(), &well_set_well); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_set_well.Well_set_id = id

    if well_set_well.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_set_well set uwi = :1, well_set_well_obs_no = :2, active_ind = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, source = :8, well_set_role = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where well_set_id = :17")
	if err != nil {
		return err
	}

	well_set_well.Row_changed_date = formatDate(well_set_well.Row_changed_date)
	well_set_well.Effective_date = formatDateString(well_set_well.Effective_date)
	well_set_well.Expiry_date = formatDateString(well_set_well.Expiry_date)
	well_set_well.Row_effective_date = formatDateString(well_set_well.Row_effective_date)
	well_set_well.Row_expiry_date = formatDateString(well_set_well.Row_expiry_date)






	rows, err := stmt.Exec(well_set_well.Uwi, well_set_well.Well_set_well_obs_no, well_set_well.Active_ind, well_set_well.Effective_date, well_set_well.Expiry_date, well_set_well.Ppdm_guid, well_set_well.Remark, well_set_well.Source, well_set_well.Well_set_role, well_set_well.Row_changed_by, well_set_well.Row_changed_date, well_set_well.Row_created_by, well_set_well.Row_effective_date, well_set_well.Row_expiry_date, well_set_well.Row_quality, well_set_well.Well_set_id)
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

func PatchWellSetWell(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_set_well set "
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
	query += " where well_set_id = :id"

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

func DeleteWellSetWell(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_set_well dto.Well_set_well
	well_set_well.Well_set_id = id

	stmt, err := db.Prepare("delete from well_set_well where well_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_set_well.Well_set_id)
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



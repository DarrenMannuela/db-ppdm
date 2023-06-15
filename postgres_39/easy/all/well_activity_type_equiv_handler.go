package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellActivityTypeEquiv(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_activity_type_equiv")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_activity_type_equiv

	for rows.Next() {
		var well_activity_type_equiv dto.Well_activity_type_equiv
		if err := rows.Scan(&well_activity_type_equiv.Activity_set_id, &well_activity_type_equiv.Activity_type_id, &well_activity_type_equiv.Activity_set_id2, &well_activity_type_equiv.Activity_type_id2, &well_activity_type_equiv.Equiv_obs_no, &well_activity_type_equiv.Active_ind, &well_activity_type_equiv.Effective_date, &well_activity_type_equiv.Equiv_owner_ba_id, &well_activity_type_equiv.Equiv_type, &well_activity_type_equiv.Expiry_date, &well_activity_type_equiv.Full_name, &well_activity_type_equiv.Ppdm_guid, &well_activity_type_equiv.Remark, &well_activity_type_equiv.Source, &well_activity_type_equiv.Row_changed_by, &well_activity_type_equiv.Row_changed_date, &well_activity_type_equiv.Row_created_by, &well_activity_type_equiv.Row_created_date, &well_activity_type_equiv.Row_effective_date, &well_activity_type_equiv.Row_expiry_date, &well_activity_type_equiv.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_activity_type_equiv to result
		result = append(result, well_activity_type_equiv)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellActivityTypeEquiv(c *fiber.Ctx) error {
	var well_activity_type_equiv dto.Well_activity_type_equiv

	setDefaults(&well_activity_type_equiv)

	if err := json.Unmarshal(c.Body(), &well_activity_type_equiv); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_activity_type_equiv values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	well_activity_type_equiv.Row_created_date = formatDate(well_activity_type_equiv.Row_created_date)
	well_activity_type_equiv.Row_changed_date = formatDate(well_activity_type_equiv.Row_changed_date)
	well_activity_type_equiv.Effective_date = formatDateString(well_activity_type_equiv.Effective_date)
	well_activity_type_equiv.Expiry_date = formatDateString(well_activity_type_equiv.Expiry_date)
	well_activity_type_equiv.Row_effective_date = formatDateString(well_activity_type_equiv.Row_effective_date)
	well_activity_type_equiv.Row_expiry_date = formatDateString(well_activity_type_equiv.Row_expiry_date)






	rows, err := stmt.Exec(well_activity_type_equiv.Activity_set_id, well_activity_type_equiv.Activity_type_id, well_activity_type_equiv.Activity_set_id2, well_activity_type_equiv.Activity_type_id2, well_activity_type_equiv.Equiv_obs_no, well_activity_type_equiv.Active_ind, well_activity_type_equiv.Effective_date, well_activity_type_equiv.Equiv_owner_ba_id, well_activity_type_equiv.Equiv_type, well_activity_type_equiv.Expiry_date, well_activity_type_equiv.Full_name, well_activity_type_equiv.Ppdm_guid, well_activity_type_equiv.Remark, well_activity_type_equiv.Source, well_activity_type_equiv.Row_changed_by, well_activity_type_equiv.Row_changed_date, well_activity_type_equiv.Row_created_by, well_activity_type_equiv.Row_created_date, well_activity_type_equiv.Row_effective_date, well_activity_type_equiv.Row_expiry_date, well_activity_type_equiv.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellActivityTypeEquiv(c *fiber.Ctx) error {
	var well_activity_type_equiv dto.Well_activity_type_equiv

	setDefaults(&well_activity_type_equiv)

	if err := json.Unmarshal(c.Body(), &well_activity_type_equiv); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_activity_type_equiv.Activity_set_id = id

    if well_activity_type_equiv.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_activity_type_equiv set activity_type_id = :1, activity_set_id2 = :2, activity_type_id2 = :3, equiv_obs_no = :4, active_ind = :5, effective_date = :6, equiv_owner_ba_id = :7, equiv_type = :8, expiry_date = :9, full_name = :10, ppdm_guid = :11, remark = :12, source = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where activity_set_id = :21")
	if err != nil {
		return err
	}

	well_activity_type_equiv.Row_changed_date = formatDate(well_activity_type_equiv.Row_changed_date)
	well_activity_type_equiv.Effective_date = formatDateString(well_activity_type_equiv.Effective_date)
	well_activity_type_equiv.Expiry_date = formatDateString(well_activity_type_equiv.Expiry_date)
	well_activity_type_equiv.Row_effective_date = formatDateString(well_activity_type_equiv.Row_effective_date)
	well_activity_type_equiv.Row_expiry_date = formatDateString(well_activity_type_equiv.Row_expiry_date)






	rows, err := stmt.Exec(well_activity_type_equiv.Activity_type_id, well_activity_type_equiv.Activity_set_id2, well_activity_type_equiv.Activity_type_id2, well_activity_type_equiv.Equiv_obs_no, well_activity_type_equiv.Active_ind, well_activity_type_equiv.Effective_date, well_activity_type_equiv.Equiv_owner_ba_id, well_activity_type_equiv.Equiv_type, well_activity_type_equiv.Expiry_date, well_activity_type_equiv.Full_name, well_activity_type_equiv.Ppdm_guid, well_activity_type_equiv.Remark, well_activity_type_equiv.Source, well_activity_type_equiv.Row_changed_by, well_activity_type_equiv.Row_changed_date, well_activity_type_equiv.Row_created_by, well_activity_type_equiv.Row_effective_date, well_activity_type_equiv.Row_expiry_date, well_activity_type_equiv.Row_quality, well_activity_type_equiv.Activity_set_id)
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

func PatchWellActivityTypeEquiv(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_activity_type_equiv set "
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
	query += " where activity_set_id = :id"

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

func DeleteWellActivityTypeEquiv(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_activity_type_equiv dto.Well_activity_type_equiv
	well_activity_type_equiv.Activity_set_id = id

	stmt, err := db.Prepare("delete from well_activity_type_equiv where activity_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_activity_type_equiv.Activity_set_id)
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



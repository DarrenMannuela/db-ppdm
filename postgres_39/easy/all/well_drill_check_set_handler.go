package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellDrillCheckSet(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_drill_check_set")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_drill_check_set

	for rows.Next() {
		var well_drill_check_set dto.Well_drill_check_set
		if err := rows.Scan(&well_drill_check_set.Check_set_id, &well_drill_check_set.Active_ind, &well_drill_check_set.Effective_date, &well_drill_check_set.Expiry_date, &well_drill_check_set.Full_name, &well_drill_check_set.Owner_ba_id, &well_drill_check_set.Ppdm_guid, &well_drill_check_set.Reference_num, &well_drill_check_set.Remark, &well_drill_check_set.Source, &well_drill_check_set.Row_changed_by, &well_drill_check_set.Row_changed_date, &well_drill_check_set.Row_created_by, &well_drill_check_set.Row_created_date, &well_drill_check_set.Row_effective_date, &well_drill_check_set.Row_expiry_date, &well_drill_check_set.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_drill_check_set to result
		result = append(result, well_drill_check_set)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellDrillCheckSet(c *fiber.Ctx) error {
	var well_drill_check_set dto.Well_drill_check_set

	setDefaults(&well_drill_check_set)

	if err := json.Unmarshal(c.Body(), &well_drill_check_set); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_drill_check_set values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	well_drill_check_set.Row_created_date = formatDate(well_drill_check_set.Row_created_date)
	well_drill_check_set.Row_changed_date = formatDate(well_drill_check_set.Row_changed_date)
	well_drill_check_set.Effective_date = formatDateString(well_drill_check_set.Effective_date)
	well_drill_check_set.Expiry_date = formatDateString(well_drill_check_set.Expiry_date)
	well_drill_check_set.Row_effective_date = formatDateString(well_drill_check_set.Row_effective_date)
	well_drill_check_set.Row_expiry_date = formatDateString(well_drill_check_set.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_check_set.Check_set_id, well_drill_check_set.Active_ind, well_drill_check_set.Effective_date, well_drill_check_set.Expiry_date, well_drill_check_set.Full_name, well_drill_check_set.Owner_ba_id, well_drill_check_set.Ppdm_guid, well_drill_check_set.Reference_num, well_drill_check_set.Remark, well_drill_check_set.Source, well_drill_check_set.Row_changed_by, well_drill_check_set.Row_changed_date, well_drill_check_set.Row_created_by, well_drill_check_set.Row_created_date, well_drill_check_set.Row_effective_date, well_drill_check_set.Row_expiry_date, well_drill_check_set.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellDrillCheckSet(c *fiber.Ctx) error {
	var well_drill_check_set dto.Well_drill_check_set

	setDefaults(&well_drill_check_set)

	if err := json.Unmarshal(c.Body(), &well_drill_check_set); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_drill_check_set.Check_set_id = id

    if well_drill_check_set.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_drill_check_set set active_ind = :1, effective_date = :2, expiry_date = :3, full_name = :4, owner_ba_id = :5, ppdm_guid = :6, reference_num = :7, remark = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where check_set_id = :17")
	if err != nil {
		return err
	}

	well_drill_check_set.Row_changed_date = formatDate(well_drill_check_set.Row_changed_date)
	well_drill_check_set.Effective_date = formatDateString(well_drill_check_set.Effective_date)
	well_drill_check_set.Expiry_date = formatDateString(well_drill_check_set.Expiry_date)
	well_drill_check_set.Row_effective_date = formatDateString(well_drill_check_set.Row_effective_date)
	well_drill_check_set.Row_expiry_date = formatDateString(well_drill_check_set.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_check_set.Active_ind, well_drill_check_set.Effective_date, well_drill_check_set.Expiry_date, well_drill_check_set.Full_name, well_drill_check_set.Owner_ba_id, well_drill_check_set.Ppdm_guid, well_drill_check_set.Reference_num, well_drill_check_set.Remark, well_drill_check_set.Source, well_drill_check_set.Row_changed_by, well_drill_check_set.Row_changed_date, well_drill_check_set.Row_created_by, well_drill_check_set.Row_effective_date, well_drill_check_set.Row_expiry_date, well_drill_check_set.Row_quality, well_drill_check_set.Check_set_id)
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

func PatchWellDrillCheckSet(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_drill_check_set set "
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
	query += " where check_set_id = :id"

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

func DeleteWellDrillCheckSet(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_drill_check_set dto.Well_drill_check_set
	well_drill_check_set.Check_set_id = id

	stmt, err := db.Prepare("delete from well_drill_check_set where check_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_drill_check_set.Check_set_id)
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



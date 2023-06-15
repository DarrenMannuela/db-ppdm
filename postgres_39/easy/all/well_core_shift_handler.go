package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellCoreShift(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_core_shift")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_core_shift

	for rows.Next() {
		var well_core_shift dto.Well_core_shift
		if err := rows.Scan(&well_core_shift.Uwi, &well_core_shift.Source, &well_core_shift.Core_id, &well_core_shift.Shift_obs_no, &well_core_shift.Active_ind, &well_core_shift.Base_depth, &well_core_shift.Base_depth_ouom, &well_core_shift.Base_shift_depth, &well_core_shift.Base_shift_depth_ouom, &well_core_shift.Core_shift_company, &well_core_shift.Core_shift_ind, &well_core_shift.Core_shift_method, &well_core_shift.Effective_date, &well_core_shift.Expiry_date, &well_core_shift.Ppdm_guid, &well_core_shift.Remark, &well_core_shift.Top_depth, &well_core_shift.Top_depth_ouom, &well_core_shift.Top_shift_depth, &well_core_shift.Top_shift_depth_ouom, &well_core_shift.Row_changed_by, &well_core_shift.Row_changed_date, &well_core_shift.Row_created_by, &well_core_shift.Row_created_date, &well_core_shift.Row_effective_date, &well_core_shift.Row_expiry_date, &well_core_shift.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_core_shift to result
		result = append(result, well_core_shift)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellCoreShift(c *fiber.Ctx) error {
	var well_core_shift dto.Well_core_shift

	setDefaults(&well_core_shift)

	if err := json.Unmarshal(c.Body(), &well_core_shift); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_core_shift values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27)")
	if err != nil {
		return err
	}
	well_core_shift.Row_created_date = formatDate(well_core_shift.Row_created_date)
	well_core_shift.Row_changed_date = formatDate(well_core_shift.Row_changed_date)
	well_core_shift.Effective_date = formatDateString(well_core_shift.Effective_date)
	well_core_shift.Expiry_date = formatDateString(well_core_shift.Expiry_date)
	well_core_shift.Row_effective_date = formatDateString(well_core_shift.Row_effective_date)
	well_core_shift.Row_expiry_date = formatDateString(well_core_shift.Row_expiry_date)






	rows, err := stmt.Exec(well_core_shift.Uwi, well_core_shift.Source, well_core_shift.Core_id, well_core_shift.Shift_obs_no, well_core_shift.Active_ind, well_core_shift.Base_depth, well_core_shift.Base_depth_ouom, well_core_shift.Base_shift_depth, well_core_shift.Base_shift_depth_ouom, well_core_shift.Core_shift_company, well_core_shift.Core_shift_ind, well_core_shift.Core_shift_method, well_core_shift.Effective_date, well_core_shift.Expiry_date, well_core_shift.Ppdm_guid, well_core_shift.Remark, well_core_shift.Top_depth, well_core_shift.Top_depth_ouom, well_core_shift.Top_shift_depth, well_core_shift.Top_shift_depth_ouom, well_core_shift.Row_changed_by, well_core_shift.Row_changed_date, well_core_shift.Row_created_by, well_core_shift.Row_created_date, well_core_shift.Row_effective_date, well_core_shift.Row_expiry_date, well_core_shift.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellCoreShift(c *fiber.Ctx) error {
	var well_core_shift dto.Well_core_shift

	setDefaults(&well_core_shift)

	if err := json.Unmarshal(c.Body(), &well_core_shift); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_core_shift.Uwi = id

    if well_core_shift.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_core_shift set source = :1, core_id = :2, shift_obs_no = :3, active_ind = :4, base_depth = :5, base_depth_ouom = :6, base_shift_depth = :7, base_shift_depth_ouom = :8, core_shift_company = :9, core_shift_ind = :10, core_shift_method = :11, effective_date = :12, expiry_date = :13, ppdm_guid = :14, remark = :15, top_depth = :16, top_depth_ouom = :17, top_shift_depth = :18, top_shift_depth_ouom = :19, row_changed_by = :20, row_changed_date = :21, row_created_by = :22, row_effective_date = :23, row_expiry_date = :24, row_quality = :25 where uwi = :27")
	if err != nil {
		return err
	}

	well_core_shift.Row_changed_date = formatDate(well_core_shift.Row_changed_date)
	well_core_shift.Effective_date = formatDateString(well_core_shift.Effective_date)
	well_core_shift.Expiry_date = formatDateString(well_core_shift.Expiry_date)
	well_core_shift.Row_effective_date = formatDateString(well_core_shift.Row_effective_date)
	well_core_shift.Row_expiry_date = formatDateString(well_core_shift.Row_expiry_date)






	rows, err := stmt.Exec(well_core_shift.Source, well_core_shift.Core_id, well_core_shift.Shift_obs_no, well_core_shift.Active_ind, well_core_shift.Base_depth, well_core_shift.Base_depth_ouom, well_core_shift.Base_shift_depth, well_core_shift.Base_shift_depth_ouom, well_core_shift.Core_shift_company, well_core_shift.Core_shift_ind, well_core_shift.Core_shift_method, well_core_shift.Effective_date, well_core_shift.Expiry_date, well_core_shift.Ppdm_guid, well_core_shift.Remark, well_core_shift.Top_depth, well_core_shift.Top_depth_ouom, well_core_shift.Top_shift_depth, well_core_shift.Top_shift_depth_ouom, well_core_shift.Row_changed_by, well_core_shift.Row_changed_date, well_core_shift.Row_created_by, well_core_shift.Row_effective_date, well_core_shift.Row_expiry_date, well_core_shift.Row_quality, well_core_shift.Uwi)
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

func PatchWellCoreShift(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_core_shift set "
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
	query += " where uwi = :id"

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

func DeleteWellCoreShift(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_core_shift dto.Well_core_shift
	well_core_shift.Uwi = id

	stmt, err := db.Prepare("delete from well_core_shift where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_core_shift.Uwi)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellCompletion(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_completion")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_completion

	for rows.Next() {
		var well_completion dto.Well_completion
		if err := rows.Scan(&well_completion.Uwi, &well_completion.Source, &well_completion.Completion_obs_no, &well_completion.Active_ind, &well_completion.Base_depth, &well_completion.Base_depth_ouom, &well_completion.Base_strat_unit_id, &well_completion.Completion_date, &well_completion.Completion_method, &well_completion.Completion_strat_unit_id, &well_completion.Completion_type, &well_completion.Effective_date, &well_completion.Expiry_date, &well_completion.Ppdm_guid, &well_completion.Remark, &well_completion.Strat_name_set_id, &well_completion.Top_depth, &well_completion.Top_depth_ouom, &well_completion.Top_strat_unit_id, &well_completion.Row_changed_by, &well_completion.Row_changed_date, &well_completion.Row_created_by, &well_completion.Row_created_date, &well_completion.Row_effective_date, &well_completion.Row_expiry_date, &well_completion.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_completion to result
		result = append(result, well_completion)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellCompletion(c *fiber.Ctx) error {
	var well_completion dto.Well_completion

	setDefaults(&well_completion)

	if err := json.Unmarshal(c.Body(), &well_completion); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_completion values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26)")
	if err != nil {
		return err
	}
	well_completion.Row_created_date = formatDate(well_completion.Row_created_date)
	well_completion.Row_changed_date = formatDate(well_completion.Row_changed_date)
	well_completion.Completion_date = formatDateString(well_completion.Completion_date)
	well_completion.Effective_date = formatDateString(well_completion.Effective_date)
	well_completion.Expiry_date = formatDateString(well_completion.Expiry_date)
	well_completion.Row_effective_date = formatDateString(well_completion.Row_effective_date)
	well_completion.Row_expiry_date = formatDateString(well_completion.Row_expiry_date)






	rows, err := stmt.Exec(well_completion.Uwi, well_completion.Source, well_completion.Completion_obs_no, well_completion.Active_ind, well_completion.Base_depth, well_completion.Base_depth_ouom, well_completion.Base_strat_unit_id, well_completion.Completion_date, well_completion.Completion_method, well_completion.Completion_strat_unit_id, well_completion.Completion_type, well_completion.Effective_date, well_completion.Expiry_date, well_completion.Ppdm_guid, well_completion.Remark, well_completion.Strat_name_set_id, well_completion.Top_depth, well_completion.Top_depth_ouom, well_completion.Top_strat_unit_id, well_completion.Row_changed_by, well_completion.Row_changed_date, well_completion.Row_created_by, well_completion.Row_created_date, well_completion.Row_effective_date, well_completion.Row_expiry_date, well_completion.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellCompletion(c *fiber.Ctx) error {
	var well_completion dto.Well_completion

	setDefaults(&well_completion)

	if err := json.Unmarshal(c.Body(), &well_completion); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_completion.Uwi = id

    if well_completion.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_completion set source = :1, completion_obs_no = :2, active_ind = :3, base_depth = :4, base_depth_ouom = :5, base_strat_unit_id = :6, completion_date = :7, completion_method = :8, completion_strat_unit_id = :9, completion_type = :10, effective_date = :11, expiry_date = :12, ppdm_guid = :13, remark = :14, strat_name_set_id = :15, top_depth = :16, top_depth_ouom = :17, top_strat_unit_id = :18, row_changed_by = :19, row_changed_date = :20, row_created_by = :21, row_effective_date = :22, row_expiry_date = :23, row_quality = :24 where uwi = :26")
	if err != nil {
		return err
	}

	well_completion.Row_changed_date = formatDate(well_completion.Row_changed_date)
	well_completion.Completion_date = formatDateString(well_completion.Completion_date)
	well_completion.Effective_date = formatDateString(well_completion.Effective_date)
	well_completion.Expiry_date = formatDateString(well_completion.Expiry_date)
	well_completion.Row_effective_date = formatDateString(well_completion.Row_effective_date)
	well_completion.Row_expiry_date = formatDateString(well_completion.Row_expiry_date)






	rows, err := stmt.Exec(well_completion.Source, well_completion.Completion_obs_no, well_completion.Active_ind, well_completion.Base_depth, well_completion.Base_depth_ouom, well_completion.Base_strat_unit_id, well_completion.Completion_date, well_completion.Completion_method, well_completion.Completion_strat_unit_id, well_completion.Completion_type, well_completion.Effective_date, well_completion.Expiry_date, well_completion.Ppdm_guid, well_completion.Remark, well_completion.Strat_name_set_id, well_completion.Top_depth, well_completion.Top_depth_ouom, well_completion.Top_strat_unit_id, well_completion.Row_changed_by, well_completion.Row_changed_date, well_completion.Row_created_by, well_completion.Row_effective_date, well_completion.Row_expiry_date, well_completion.Row_quality, well_completion.Uwi)
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

func PatchWellCompletion(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_completion set "
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
		} else if key == "completion_date" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellCompletion(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_completion dto.Well_completion
	well_completion.Uwi = id

	stmt, err := db.Prepare("delete from well_completion where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_completion.Uwi)
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



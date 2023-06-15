package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellPerforation(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_perforation")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_perforation

	for rows.Next() {
		var well_perforation dto.Well_perforation
		if err := rows.Scan(&well_perforation.Uwi, &well_perforation.Source, &well_perforation.Perforation_obs_no, &well_perforation.Active_ind, &well_perforation.Base_depth, &well_perforation.Base_depth_ouom, &well_perforation.Base_strat_unit_id, &well_perforation.Completion_obs_no, &well_perforation.Completion_source, &well_perforation.Completion_status, &well_perforation.Completion_status_type, &well_perforation.Current_status_date, &well_perforation.Effective_date, &well_perforation.Expiry_date, &well_perforation.Perforation_angle, &well_perforation.Perforation_ba_id, &well_perforation.Perforation_count, &well_perforation.Perforation_date, &well_perforation.Perforation_density, &well_perforation.Perforation_diameter, &well_perforation.Perforation_diameter_ouom, &well_perforation.Perforation_method, &well_perforation.Perforation_per_uom, &well_perforation.Perforation_phase, &well_perforation.Perforation_type, &well_perforation.Ppdm_guid, &well_perforation.Remark, &well_perforation.Strat_name_set_id, &well_perforation.Top_depth, &well_perforation.Top_depth_ouom, &well_perforation.Top_strat_unit_id, &well_perforation.Row_changed_by, &well_perforation.Row_changed_date, &well_perforation.Row_created_by, &well_perforation.Row_created_date, &well_perforation.Row_effective_date, &well_perforation.Row_expiry_date, &well_perforation.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_perforation to result
		result = append(result, well_perforation)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellPerforation(c *fiber.Ctx) error {
	var well_perforation dto.Well_perforation

	setDefaults(&well_perforation)

	if err := json.Unmarshal(c.Body(), &well_perforation); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_perforation values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38)")
	if err != nil {
		return err
	}
	well_perforation.Row_created_date = formatDate(well_perforation.Row_created_date)
	well_perforation.Row_changed_date = formatDate(well_perforation.Row_changed_date)
	well_perforation.Current_status_date = formatDateString(well_perforation.Current_status_date)
	well_perforation.Effective_date = formatDateString(well_perforation.Effective_date)
	well_perforation.Expiry_date = formatDateString(well_perforation.Expiry_date)
	well_perforation.Perforation_date = formatDateString(well_perforation.Perforation_date)
	well_perforation.Row_effective_date = formatDateString(well_perforation.Row_effective_date)
	well_perforation.Row_expiry_date = formatDateString(well_perforation.Row_expiry_date)






	rows, err := stmt.Exec(well_perforation.Uwi, well_perforation.Source, well_perforation.Perforation_obs_no, well_perforation.Active_ind, well_perforation.Base_depth, well_perforation.Base_depth_ouom, well_perforation.Base_strat_unit_id, well_perforation.Completion_obs_no, well_perforation.Completion_source, well_perforation.Completion_status, well_perforation.Completion_status_type, well_perforation.Current_status_date, well_perforation.Effective_date, well_perforation.Expiry_date, well_perforation.Perforation_angle, well_perforation.Perforation_ba_id, well_perforation.Perforation_count, well_perforation.Perforation_date, well_perforation.Perforation_density, well_perforation.Perforation_diameter, well_perforation.Perforation_diameter_ouom, well_perforation.Perforation_method, well_perforation.Perforation_per_uom, well_perforation.Perforation_phase, well_perforation.Perforation_type, well_perforation.Ppdm_guid, well_perforation.Remark, well_perforation.Strat_name_set_id, well_perforation.Top_depth, well_perforation.Top_depth_ouom, well_perforation.Top_strat_unit_id, well_perforation.Row_changed_by, well_perforation.Row_changed_date, well_perforation.Row_created_by, well_perforation.Row_created_date, well_perforation.Row_effective_date, well_perforation.Row_expiry_date, well_perforation.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellPerforation(c *fiber.Ctx) error {
	var well_perforation dto.Well_perforation

	setDefaults(&well_perforation)

	if err := json.Unmarshal(c.Body(), &well_perforation); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_perforation.Uwi = id

    if well_perforation.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_perforation set source = :1, perforation_obs_no = :2, active_ind = :3, base_depth = :4, base_depth_ouom = :5, base_strat_unit_id = :6, completion_obs_no = :7, completion_source = :8, completion_status = :9, completion_status_type = :10, current_status_date = :11, effective_date = :12, expiry_date = :13, perforation_angle = :14, perforation_ba_id = :15, perforation_count = :16, perforation_date = :17, perforation_density = :18, perforation_diameter = :19, perforation_diameter_ouom = :20, perforation_method = :21, perforation_per_uom = :22, perforation_phase = :23, perforation_type = :24, ppdm_guid = :25, remark = :26, strat_name_set_id = :27, top_depth = :28, top_depth_ouom = :29, top_strat_unit_id = :30, row_changed_by = :31, row_changed_date = :32, row_created_by = :33, row_effective_date = :34, row_expiry_date = :35, row_quality = :36 where uwi = :38")
	if err != nil {
		return err
	}

	well_perforation.Row_changed_date = formatDate(well_perforation.Row_changed_date)
	well_perforation.Current_status_date = formatDateString(well_perforation.Current_status_date)
	well_perforation.Effective_date = formatDateString(well_perforation.Effective_date)
	well_perforation.Expiry_date = formatDateString(well_perforation.Expiry_date)
	well_perforation.Perforation_date = formatDateString(well_perforation.Perforation_date)
	well_perforation.Row_effective_date = formatDateString(well_perforation.Row_effective_date)
	well_perforation.Row_expiry_date = formatDateString(well_perforation.Row_expiry_date)






	rows, err := stmt.Exec(well_perforation.Source, well_perforation.Perforation_obs_no, well_perforation.Active_ind, well_perforation.Base_depth, well_perforation.Base_depth_ouom, well_perforation.Base_strat_unit_id, well_perforation.Completion_obs_no, well_perforation.Completion_source, well_perforation.Completion_status, well_perforation.Completion_status_type, well_perforation.Current_status_date, well_perforation.Effective_date, well_perforation.Expiry_date, well_perforation.Perforation_angle, well_perforation.Perforation_ba_id, well_perforation.Perforation_count, well_perforation.Perforation_date, well_perforation.Perforation_density, well_perforation.Perforation_diameter, well_perforation.Perforation_diameter_ouom, well_perforation.Perforation_method, well_perforation.Perforation_per_uom, well_perforation.Perforation_phase, well_perforation.Perforation_type, well_perforation.Ppdm_guid, well_perforation.Remark, well_perforation.Strat_name_set_id, well_perforation.Top_depth, well_perforation.Top_depth_ouom, well_perforation.Top_strat_unit_id, well_perforation.Row_changed_by, well_perforation.Row_changed_date, well_perforation.Row_created_by, well_perforation.Row_effective_date, well_perforation.Row_expiry_date, well_perforation.Row_quality, well_perforation.Uwi)
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

func PatchWellPerforation(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_perforation set "
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
		} else if key == "current_status_date" || key == "effective_date" || key == "expiry_date" || key == "perforation_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellPerforation(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_perforation dto.Well_perforation
	well_perforation.Uwi = id

	stmt, err := db.Prepare("delete from well_perforation where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_perforation.Uwi)
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



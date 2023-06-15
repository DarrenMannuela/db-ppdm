package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellDrillRemark(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_drill_remark")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_drill_remark

	for rows.Next() {
		var well_drill_remark dto.Well_drill_remark
		if err := rows.Scan(&well_drill_remark.Uwi, &well_drill_remark.Remark_seq_no, &well_drill_remark.Active_ind, &well_drill_remark.Base_depth, &well_drill_remark.Base_depth_ouom, &well_drill_remark.Effective_date, &well_drill_remark.End_period_obs_no, &well_drill_remark.Expiry_date, &well_drill_remark.Ppdm_guid, &well_drill_remark.Remark, &well_drill_remark.Remark_by_ba_id, &well_drill_remark.Remark_date, &well_drill_remark.Remark_type, &well_drill_remark.Source, &well_drill_remark.Start_period_obs_no, &well_drill_remark.Strat_name_set_id, &well_drill_remark.Strat_unit_id, &well_drill_remark.Top_depth, &well_drill_remark.Top_depth_ouom, &well_drill_remark.Row_changed_by, &well_drill_remark.Row_changed_date, &well_drill_remark.Row_created_by, &well_drill_remark.Row_created_date, &well_drill_remark.Row_effective_date, &well_drill_remark.Row_expiry_date, &well_drill_remark.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_drill_remark to result
		result = append(result, well_drill_remark)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellDrillRemark(c *fiber.Ctx) error {
	var well_drill_remark dto.Well_drill_remark

	setDefaults(&well_drill_remark)

	if err := json.Unmarshal(c.Body(), &well_drill_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_drill_remark values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26)")
	if err != nil {
		return err
	}
	well_drill_remark.Row_created_date = formatDate(well_drill_remark.Row_created_date)
	well_drill_remark.Row_changed_date = formatDate(well_drill_remark.Row_changed_date)
	well_drill_remark.Effective_date = formatDateString(well_drill_remark.Effective_date)
	well_drill_remark.Expiry_date = formatDateString(well_drill_remark.Expiry_date)
	well_drill_remark.Remark_date = formatDateString(well_drill_remark.Remark_date)
	well_drill_remark.Row_effective_date = formatDateString(well_drill_remark.Row_effective_date)
	well_drill_remark.Row_expiry_date = formatDateString(well_drill_remark.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_remark.Uwi, well_drill_remark.Remark_seq_no, well_drill_remark.Active_ind, well_drill_remark.Base_depth, well_drill_remark.Base_depth_ouom, well_drill_remark.Effective_date, well_drill_remark.End_period_obs_no, well_drill_remark.Expiry_date, well_drill_remark.Ppdm_guid, well_drill_remark.Remark, well_drill_remark.Remark_by_ba_id, well_drill_remark.Remark_date, well_drill_remark.Remark_type, well_drill_remark.Source, well_drill_remark.Start_period_obs_no, well_drill_remark.Strat_name_set_id, well_drill_remark.Strat_unit_id, well_drill_remark.Top_depth, well_drill_remark.Top_depth_ouom, well_drill_remark.Row_changed_by, well_drill_remark.Row_changed_date, well_drill_remark.Row_created_by, well_drill_remark.Row_created_date, well_drill_remark.Row_effective_date, well_drill_remark.Row_expiry_date, well_drill_remark.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellDrillRemark(c *fiber.Ctx) error {
	var well_drill_remark dto.Well_drill_remark

	setDefaults(&well_drill_remark)

	if err := json.Unmarshal(c.Body(), &well_drill_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_drill_remark.Uwi = id

    if well_drill_remark.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_drill_remark set remark_seq_no = :1, active_ind = :2, base_depth = :3, base_depth_ouom = :4, effective_date = :5, end_period_obs_no = :6, expiry_date = :7, ppdm_guid = :8, remark = :9, remark_by_ba_id = :10, remark_date = :11, remark_type = :12, source = :13, start_period_obs_no = :14, strat_name_set_id = :15, strat_unit_id = :16, top_depth = :17, top_depth_ouom = :18, row_changed_by = :19, row_changed_date = :20, row_created_by = :21, row_effective_date = :22, row_expiry_date = :23, row_quality = :24 where uwi = :26")
	if err != nil {
		return err
	}

	well_drill_remark.Row_changed_date = formatDate(well_drill_remark.Row_changed_date)
	well_drill_remark.Effective_date = formatDateString(well_drill_remark.Effective_date)
	well_drill_remark.Expiry_date = formatDateString(well_drill_remark.Expiry_date)
	well_drill_remark.Remark_date = formatDateString(well_drill_remark.Remark_date)
	well_drill_remark.Row_effective_date = formatDateString(well_drill_remark.Row_effective_date)
	well_drill_remark.Row_expiry_date = formatDateString(well_drill_remark.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_remark.Remark_seq_no, well_drill_remark.Active_ind, well_drill_remark.Base_depth, well_drill_remark.Base_depth_ouom, well_drill_remark.Effective_date, well_drill_remark.End_period_obs_no, well_drill_remark.Expiry_date, well_drill_remark.Ppdm_guid, well_drill_remark.Remark, well_drill_remark.Remark_by_ba_id, well_drill_remark.Remark_date, well_drill_remark.Remark_type, well_drill_remark.Source, well_drill_remark.Start_period_obs_no, well_drill_remark.Strat_name_set_id, well_drill_remark.Strat_unit_id, well_drill_remark.Top_depth, well_drill_remark.Top_depth_ouom, well_drill_remark.Row_changed_by, well_drill_remark.Row_changed_date, well_drill_remark.Row_created_by, well_drill_remark.Row_effective_date, well_drill_remark.Row_expiry_date, well_drill_remark.Row_quality, well_drill_remark.Uwi)
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

func PatchWellDrillRemark(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_drill_remark set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "remark_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellDrillRemark(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_drill_remark dto.Well_drill_remark
	well_drill_remark.Uwi = id

	stmt, err := db.Prepare("delete from well_drill_remark where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_drill_remark.Uwi)
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



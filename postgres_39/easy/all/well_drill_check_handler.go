package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellDrillCheck(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_drill_check")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_drill_check

	for rows.Next() {
		var well_drill_check dto.Well_drill_check
		if err := rows.Scan(&well_drill_check.Uwi, &well_drill_check.Period_obs_no, &well_drill_check.Drill_check_obs_no, &well_drill_check.Active_ind, &well_drill_check.Check_set_id, &well_drill_check.Check_type, &well_drill_check.Contractor_name, &well_drill_check.Contractor_rep_ba_id, &well_drill_check.Deficient_ind, &well_drill_check.Deficient_period, &well_drill_check.Deficient_period_ouom, &well_drill_check.Deficient_period_uom, &well_drill_check.Effective_date, &well_drill_check.Expiry_date, &well_drill_check.Operator_name, &well_drill_check.Operator_rep_ba_id, &well_drill_check.Passed_ind, &well_drill_check.Ppdm_guid, &well_drill_check.Recorded_time, &well_drill_check.Recorded_timezone, &well_drill_check.Remark, &well_drill_check.Source, &well_drill_check.Row_changed_by, &well_drill_check.Row_changed_date, &well_drill_check.Row_created_by, &well_drill_check.Row_created_date, &well_drill_check.Row_effective_date, &well_drill_check.Row_expiry_date, &well_drill_check.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_drill_check to result
		result = append(result, well_drill_check)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellDrillCheck(c *fiber.Ctx) error {
	var well_drill_check dto.Well_drill_check

	setDefaults(&well_drill_check)

	if err := json.Unmarshal(c.Body(), &well_drill_check); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_drill_check values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29)")
	if err != nil {
		return err
	}
	well_drill_check.Row_created_date = formatDate(well_drill_check.Row_created_date)
	well_drill_check.Row_changed_date = formatDate(well_drill_check.Row_changed_date)
	well_drill_check.Effective_date = formatDateString(well_drill_check.Effective_date)
	well_drill_check.Expiry_date = formatDateString(well_drill_check.Expiry_date)
	well_drill_check.Recorded_time = formatDateString(well_drill_check.Recorded_time)
	well_drill_check.Row_effective_date = formatDateString(well_drill_check.Row_effective_date)
	well_drill_check.Row_expiry_date = formatDateString(well_drill_check.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_check.Uwi, well_drill_check.Period_obs_no, well_drill_check.Drill_check_obs_no, well_drill_check.Active_ind, well_drill_check.Check_set_id, well_drill_check.Check_type, well_drill_check.Contractor_name, well_drill_check.Contractor_rep_ba_id, well_drill_check.Deficient_ind, well_drill_check.Deficient_period, well_drill_check.Deficient_period_ouom, well_drill_check.Deficient_period_uom, well_drill_check.Effective_date, well_drill_check.Expiry_date, well_drill_check.Operator_name, well_drill_check.Operator_rep_ba_id, well_drill_check.Passed_ind, well_drill_check.Ppdm_guid, well_drill_check.Recorded_time, well_drill_check.Recorded_timezone, well_drill_check.Remark, well_drill_check.Source, well_drill_check.Row_changed_by, well_drill_check.Row_changed_date, well_drill_check.Row_created_by, well_drill_check.Row_created_date, well_drill_check.Row_effective_date, well_drill_check.Row_expiry_date, well_drill_check.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellDrillCheck(c *fiber.Ctx) error {
	var well_drill_check dto.Well_drill_check

	setDefaults(&well_drill_check)

	if err := json.Unmarshal(c.Body(), &well_drill_check); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_drill_check.Uwi = id

    if well_drill_check.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_drill_check set period_obs_no = :1, drill_check_obs_no = :2, active_ind = :3, check_set_id = :4, check_type = :5, contractor_name = :6, contractor_rep_ba_id = :7, deficient_ind = :8, deficient_period = :9, deficient_period_ouom = :10, deficient_period_uom = :11, effective_date = :12, expiry_date = :13, operator_name = :14, operator_rep_ba_id = :15, passed_ind = :16, ppdm_guid = :17, recorded_time = :18, recorded_timezone = :19, remark = :20, source = :21, row_changed_by = :22, row_changed_date = :23, row_created_by = :24, row_effective_date = :25, row_expiry_date = :26, row_quality = :27 where uwi = :29")
	if err != nil {
		return err
	}

	well_drill_check.Row_changed_date = formatDate(well_drill_check.Row_changed_date)
	well_drill_check.Effective_date = formatDateString(well_drill_check.Effective_date)
	well_drill_check.Expiry_date = formatDateString(well_drill_check.Expiry_date)
	well_drill_check.Recorded_time = formatDateString(well_drill_check.Recorded_time)
	well_drill_check.Row_effective_date = formatDateString(well_drill_check.Row_effective_date)
	well_drill_check.Row_expiry_date = formatDateString(well_drill_check.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_check.Period_obs_no, well_drill_check.Drill_check_obs_no, well_drill_check.Active_ind, well_drill_check.Check_set_id, well_drill_check.Check_type, well_drill_check.Contractor_name, well_drill_check.Contractor_rep_ba_id, well_drill_check.Deficient_ind, well_drill_check.Deficient_period, well_drill_check.Deficient_period_ouom, well_drill_check.Deficient_period_uom, well_drill_check.Effective_date, well_drill_check.Expiry_date, well_drill_check.Operator_name, well_drill_check.Operator_rep_ba_id, well_drill_check.Passed_ind, well_drill_check.Ppdm_guid, well_drill_check.Recorded_time, well_drill_check.Recorded_timezone, well_drill_check.Remark, well_drill_check.Source, well_drill_check.Row_changed_by, well_drill_check.Row_changed_date, well_drill_check.Row_created_by, well_drill_check.Row_effective_date, well_drill_check.Row_expiry_date, well_drill_check.Row_quality, well_drill_check.Uwi)
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

func PatchWellDrillCheck(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_drill_check set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "recorded_time" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellDrillCheck(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_drill_check dto.Well_drill_check
	well_drill_check.Uwi = id

	stmt, err := db.Prepare("delete from well_drill_check where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_drill_check.Uwi)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellTestRecovery(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_test_recovery")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_test_recovery

	for rows.Next() {
		var well_test_recovery dto.Well_test_recovery
		if err := rows.Scan(&well_test_recovery.Uwi, &well_test_recovery.Source, &well_test_recovery.Test_type, &well_test_recovery.Run_num, &well_test_recovery.Test_num, &well_test_recovery.Recovery_obs_no, &well_test_recovery.Active_ind, &well_test_recovery.Effective_date, &well_test_recovery.Expiry_date, &well_test_recovery.Multiple_test_ind, &well_test_recovery.Period_obs_no, &well_test_recovery.Period_type, &well_test_recovery.Ppdm_guid, &well_test_recovery.Recovery_amount, &well_test_recovery.Recovery_amount_ouom, &well_test_recovery.Recovery_amount_percent, &well_test_recovery.Recovery_amount_uom, &well_test_recovery.Recovery_desc, &well_test_recovery.Recovery_method, &well_test_recovery.Recovery_show_type, &well_test_recovery.Recovery_type, &well_test_recovery.Remark, &well_test_recovery.Reverse_circulation_ind, &well_test_recovery.Row_changed_by, &well_test_recovery.Row_changed_date, &well_test_recovery.Row_created_by, &well_test_recovery.Row_created_date, &well_test_recovery.Row_effective_date, &well_test_recovery.Row_expiry_date, &well_test_recovery.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_test_recovery to result
		result = append(result, well_test_recovery)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellTestRecovery(c *fiber.Ctx) error {
	var well_test_recovery dto.Well_test_recovery

	setDefaults(&well_test_recovery)

	if err := json.Unmarshal(c.Body(), &well_test_recovery); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_test_recovery values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30)")
	if err != nil {
		return err
	}
	well_test_recovery.Row_created_date = formatDate(well_test_recovery.Row_created_date)
	well_test_recovery.Row_changed_date = formatDate(well_test_recovery.Row_changed_date)
	well_test_recovery.Effective_date = formatDateString(well_test_recovery.Effective_date)
	well_test_recovery.Expiry_date = formatDateString(well_test_recovery.Expiry_date)
	well_test_recovery.Row_effective_date = formatDateString(well_test_recovery.Row_effective_date)
	well_test_recovery.Row_expiry_date = formatDateString(well_test_recovery.Row_expiry_date)






	rows, err := stmt.Exec(well_test_recovery.Uwi, well_test_recovery.Source, well_test_recovery.Test_type, well_test_recovery.Run_num, well_test_recovery.Test_num, well_test_recovery.Recovery_obs_no, well_test_recovery.Active_ind, well_test_recovery.Effective_date, well_test_recovery.Expiry_date, well_test_recovery.Multiple_test_ind, well_test_recovery.Period_obs_no, well_test_recovery.Period_type, well_test_recovery.Ppdm_guid, well_test_recovery.Recovery_amount, well_test_recovery.Recovery_amount_ouom, well_test_recovery.Recovery_amount_percent, well_test_recovery.Recovery_amount_uom, well_test_recovery.Recovery_desc, well_test_recovery.Recovery_method, well_test_recovery.Recovery_show_type, well_test_recovery.Recovery_type, well_test_recovery.Remark, well_test_recovery.Reverse_circulation_ind, well_test_recovery.Row_changed_by, well_test_recovery.Row_changed_date, well_test_recovery.Row_created_by, well_test_recovery.Row_created_date, well_test_recovery.Row_effective_date, well_test_recovery.Row_expiry_date, well_test_recovery.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellTestRecovery(c *fiber.Ctx) error {
	var well_test_recovery dto.Well_test_recovery

	setDefaults(&well_test_recovery)

	if err := json.Unmarshal(c.Body(), &well_test_recovery); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_test_recovery.Uwi = id

    if well_test_recovery.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_test_recovery set source = :1, test_type = :2, run_num = :3, test_num = :4, recovery_obs_no = :5, active_ind = :6, effective_date = :7, expiry_date = :8, multiple_test_ind = :9, period_obs_no = :10, period_type = :11, ppdm_guid = :12, recovery_amount = :13, recovery_amount_ouom = :14, recovery_amount_percent = :15, recovery_amount_uom = :16, recovery_desc = :17, recovery_method = :18, recovery_show_type = :19, recovery_type = :20, remark = :21, reverse_circulation_ind = :22, row_changed_by = :23, row_changed_date = :24, row_created_by = :25, row_effective_date = :26, row_expiry_date = :27, row_quality = :28 where uwi = :30")
	if err != nil {
		return err
	}

	well_test_recovery.Row_changed_date = formatDate(well_test_recovery.Row_changed_date)
	well_test_recovery.Effective_date = formatDateString(well_test_recovery.Effective_date)
	well_test_recovery.Expiry_date = formatDateString(well_test_recovery.Expiry_date)
	well_test_recovery.Row_effective_date = formatDateString(well_test_recovery.Row_effective_date)
	well_test_recovery.Row_expiry_date = formatDateString(well_test_recovery.Row_expiry_date)






	rows, err := stmt.Exec(well_test_recovery.Source, well_test_recovery.Test_type, well_test_recovery.Run_num, well_test_recovery.Test_num, well_test_recovery.Recovery_obs_no, well_test_recovery.Active_ind, well_test_recovery.Effective_date, well_test_recovery.Expiry_date, well_test_recovery.Multiple_test_ind, well_test_recovery.Period_obs_no, well_test_recovery.Period_type, well_test_recovery.Ppdm_guid, well_test_recovery.Recovery_amount, well_test_recovery.Recovery_amount_ouom, well_test_recovery.Recovery_amount_percent, well_test_recovery.Recovery_amount_uom, well_test_recovery.Recovery_desc, well_test_recovery.Recovery_method, well_test_recovery.Recovery_show_type, well_test_recovery.Recovery_type, well_test_recovery.Remark, well_test_recovery.Reverse_circulation_ind, well_test_recovery.Row_changed_by, well_test_recovery.Row_changed_date, well_test_recovery.Row_created_by, well_test_recovery.Row_effective_date, well_test_recovery.Row_expiry_date, well_test_recovery.Row_quality, well_test_recovery.Uwi)
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

func PatchWellTestRecovery(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_test_recovery set "
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

func DeleteWellTestRecovery(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_test_recovery dto.Well_test_recovery
	well_test_recovery.Uwi = id

	stmt, err := db.Prepare("delete from well_test_recovery where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_test_recovery.Uwi)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellCore(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_core")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_core

	for rows.Next() {
		var well_core dto.Well_core
		if err := rows.Scan(&well_core.Uwi, &well_core.Source, &well_core.Core_id, &well_core.Active_ind, &well_core.Base_depth, &well_core.Base_depth_ouom, &well_core.Contractor, &well_core.Core_barrel_size, &well_core.Core_barrel_size_ouom, &well_core.Core_diameter, &well_core.Core_diameter_ouom, &well_core.Core_handling_type, &well_core.Core_oriented_ind, &well_core.Core_show_type, &well_core.Core_type, &well_core.Coring_fluid, &well_core.Digit_avail_ind, &well_core.Effective_date, &well_core.Expiry_date, &well_core.Gamma_correlation_ind, &well_core.Operation_seq_no, &well_core.Ppdm_guid, &well_core.Primary_core_strat_unit_id, &well_core.Recovered_amount, &well_core.Recovered_amount_ouom, &well_core.Recovered_amount_uom, &well_core.Recovery_date, &well_core.Remark, &well_core.Reported_core_num, &well_core.Run_num, &well_core.Shot_recovered_count, &well_core.Sidewall_ind, &well_core.Strat_name_set_id, &well_core.Top_depth, &well_core.Top_depth_ouom, &well_core.Total_shot_count, &well_core.Row_changed_by, &well_core.Row_changed_date, &well_core.Row_created_by, &well_core.Row_created_date, &well_core.Row_effective_date, &well_core.Row_expiry_date, &well_core.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_core to result
		result = append(result, well_core)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellCore(c *fiber.Ctx) error {
	var well_core dto.Well_core

	setDefaults(&well_core)

	if err := json.Unmarshal(c.Body(), &well_core); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_core values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43)")
	if err != nil {
		return err
	}
	well_core.Row_created_date = formatDate(well_core.Row_created_date)
	well_core.Row_changed_date = formatDate(well_core.Row_changed_date)
	well_core.Effective_date = formatDateString(well_core.Effective_date)
	well_core.Expiry_date = formatDateString(well_core.Expiry_date)
	well_core.Recovery_date = formatDateString(well_core.Recovery_date)
	well_core.Row_effective_date = formatDateString(well_core.Row_effective_date)
	well_core.Row_expiry_date = formatDateString(well_core.Row_expiry_date)






	rows, err := stmt.Exec(well_core.Uwi, well_core.Source, well_core.Core_id, well_core.Active_ind, well_core.Base_depth, well_core.Base_depth_ouom, well_core.Contractor, well_core.Core_barrel_size, well_core.Core_barrel_size_ouom, well_core.Core_diameter, well_core.Core_diameter_ouom, well_core.Core_handling_type, well_core.Core_oriented_ind, well_core.Core_show_type, well_core.Core_type, well_core.Coring_fluid, well_core.Digit_avail_ind, well_core.Effective_date, well_core.Expiry_date, well_core.Gamma_correlation_ind, well_core.Operation_seq_no, well_core.Ppdm_guid, well_core.Primary_core_strat_unit_id, well_core.Recovered_amount, well_core.Recovered_amount_ouom, well_core.Recovered_amount_uom, well_core.Recovery_date, well_core.Remark, well_core.Reported_core_num, well_core.Run_num, well_core.Shot_recovered_count, well_core.Sidewall_ind, well_core.Strat_name_set_id, well_core.Top_depth, well_core.Top_depth_ouom, well_core.Total_shot_count, well_core.Row_changed_by, well_core.Row_changed_date, well_core.Row_created_by, well_core.Row_created_date, well_core.Row_effective_date, well_core.Row_expiry_date, well_core.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellCore(c *fiber.Ctx) error {
	var well_core dto.Well_core

	setDefaults(&well_core)

	if err := json.Unmarshal(c.Body(), &well_core); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_core.Uwi = id

    if well_core.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_core set source = :1, core_id = :2, active_ind = :3, base_depth = :4, base_depth_ouom = :5, contractor = :6, core_barrel_size = :7, core_barrel_size_ouom = :8, core_diameter = :9, core_diameter_ouom = :10, core_handling_type = :11, core_oriented_ind = :12, core_show_type = :13, core_type = :14, coring_fluid = :15, digit_avail_ind = :16, effective_date = :17, expiry_date = :18, gamma_correlation_ind = :19, operation_seq_no = :20, ppdm_guid = :21, primary_core_strat_unit_id = :22, recovered_amount = :23, recovered_amount_ouom = :24, recovered_amount_uom = :25, recovery_date = :26, remark = :27, reported_core_num = :28, run_num = :29, shot_recovered_count = :30, sidewall_ind = :31, strat_name_set_id = :32, top_depth = :33, top_depth_ouom = :34, total_shot_count = :35, row_changed_by = :36, row_changed_date = :37, row_created_by = :38, row_effective_date = :39, row_expiry_date = :40, row_quality = :41 where uwi = :43")
	if err != nil {
		return err
	}

	well_core.Row_changed_date = formatDate(well_core.Row_changed_date)
	well_core.Effective_date = formatDateString(well_core.Effective_date)
	well_core.Expiry_date = formatDateString(well_core.Expiry_date)
	well_core.Recovery_date = formatDateString(well_core.Recovery_date)
	well_core.Row_effective_date = formatDateString(well_core.Row_effective_date)
	well_core.Row_expiry_date = formatDateString(well_core.Row_expiry_date)






	rows, err := stmt.Exec(well_core.Source, well_core.Core_id, well_core.Active_ind, well_core.Base_depth, well_core.Base_depth_ouom, well_core.Contractor, well_core.Core_barrel_size, well_core.Core_barrel_size_ouom, well_core.Core_diameter, well_core.Core_diameter_ouom, well_core.Core_handling_type, well_core.Core_oriented_ind, well_core.Core_show_type, well_core.Core_type, well_core.Coring_fluid, well_core.Digit_avail_ind, well_core.Effective_date, well_core.Expiry_date, well_core.Gamma_correlation_ind, well_core.Operation_seq_no, well_core.Ppdm_guid, well_core.Primary_core_strat_unit_id, well_core.Recovered_amount, well_core.Recovered_amount_ouom, well_core.Recovered_amount_uom, well_core.Recovery_date, well_core.Remark, well_core.Reported_core_num, well_core.Run_num, well_core.Shot_recovered_count, well_core.Sidewall_ind, well_core.Strat_name_set_id, well_core.Top_depth, well_core.Top_depth_ouom, well_core.Total_shot_count, well_core.Row_changed_by, well_core.Row_changed_date, well_core.Row_created_by, well_core.Row_effective_date, well_core.Row_expiry_date, well_core.Row_quality, well_core.Uwi)
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

func PatchWellCore(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_core set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "recovery_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellCore(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_core dto.Well_core
	well_core.Uwi = id

	stmt, err := db.Prepare("delete from well_core where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_core.Uwi)
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



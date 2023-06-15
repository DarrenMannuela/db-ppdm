package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellPressure(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_pressure")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_pressure

	for rows.Next() {
		var well_pressure dto.Well_pressure
		if err := rows.Scan(&well_pressure.Uwi, &well_pressure.Source, &well_pressure.Pressure_obs_no, &well_pressure.Active_ind, &well_pressure.Assigned_field, &well_pressure.Base_depth, &well_pressure.Base_depth_ouom, &well_pressure.Base_strat_unit_id, &well_pressure.Effective_date, &well_pressure.Event_obs_no, &well_pressure.Event_source, &well_pressure.Expiry_date, &well_pressure.Flow_casing_pressure, &well_pressure.Flow_casing_pressure_ouom, &well_pressure.Flow_tubing_pressure, &well_pressure.Flow_tubing_pressure_ouom, &well_pressure.Init_reservoir_pressure, &well_pressure.Init_reservoir_press_ouom, &well_pressure.Pool_datum, &well_pressure.Pool_datum_depth, &well_pressure.Pool_datum_depth_ouom, &well_pressure.Pool_id, &well_pressure.Ppdm_guid, &well_pressure.Prod_string_id, &well_pressure.Prod_string_source, &well_pressure.Pr_str_form_obs_no, &well_pressure.Remark, &well_pressure.Shutin_casing_pressure, &well_pressure.Shutin_casing_pressure_ouom, &well_pressure.Shutin_tubing_pressure, &well_pressure.Shutin_tubing_pressure_ouom, &well_pressure.Strat_name_set_id, &well_pressure.Top_depth, &well_pressure.Top_depth_ouom, &well_pressure.Top_strat_unit_id, &well_pressure.Well_datum_depth, &well_pressure.Well_datum_ouom, &well_pressure.Well_test_num, &well_pressure.Well_test_run_num, &well_pressure.Well_test_source, &well_pressure.Well_test_type, &well_pressure.Row_changed_by, &well_pressure.Row_changed_date, &well_pressure.Row_created_by, &well_pressure.Row_created_date, &well_pressure.Row_effective_date, &well_pressure.Row_expiry_date, &well_pressure.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_pressure to result
		result = append(result, well_pressure)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellPressure(c *fiber.Ctx) error {
	var well_pressure dto.Well_pressure

	setDefaults(&well_pressure)

	if err := json.Unmarshal(c.Body(), &well_pressure); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_pressure values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48)")
	if err != nil {
		return err
	}
	well_pressure.Row_created_date = formatDate(well_pressure.Row_created_date)
	well_pressure.Row_changed_date = formatDate(well_pressure.Row_changed_date)
	well_pressure.Effective_date = formatDateString(well_pressure.Effective_date)
	well_pressure.Expiry_date = formatDateString(well_pressure.Expiry_date)
	well_pressure.Row_effective_date = formatDateString(well_pressure.Row_effective_date)
	well_pressure.Row_expiry_date = formatDateString(well_pressure.Row_expiry_date)






	rows, err := stmt.Exec(well_pressure.Uwi, well_pressure.Source, well_pressure.Pressure_obs_no, well_pressure.Active_ind, well_pressure.Assigned_field, well_pressure.Base_depth, well_pressure.Base_depth_ouom, well_pressure.Base_strat_unit_id, well_pressure.Effective_date, well_pressure.Event_obs_no, well_pressure.Event_source, well_pressure.Expiry_date, well_pressure.Flow_casing_pressure, well_pressure.Flow_casing_pressure_ouom, well_pressure.Flow_tubing_pressure, well_pressure.Flow_tubing_pressure_ouom, well_pressure.Init_reservoir_pressure, well_pressure.Init_reservoir_press_ouom, well_pressure.Pool_datum, well_pressure.Pool_datum_depth, well_pressure.Pool_datum_depth_ouom, well_pressure.Pool_id, well_pressure.Ppdm_guid, well_pressure.Prod_string_id, well_pressure.Prod_string_source, well_pressure.Pr_str_form_obs_no, well_pressure.Remark, well_pressure.Shutin_casing_pressure, well_pressure.Shutin_casing_pressure_ouom, well_pressure.Shutin_tubing_pressure, well_pressure.Shutin_tubing_pressure_ouom, well_pressure.Strat_name_set_id, well_pressure.Top_depth, well_pressure.Top_depth_ouom, well_pressure.Top_strat_unit_id, well_pressure.Well_datum_depth, well_pressure.Well_datum_ouom, well_pressure.Well_test_num, well_pressure.Well_test_run_num, well_pressure.Well_test_source, well_pressure.Well_test_type, well_pressure.Row_changed_by, well_pressure.Row_changed_date, well_pressure.Row_created_by, well_pressure.Row_created_date, well_pressure.Row_effective_date, well_pressure.Row_expiry_date, well_pressure.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellPressure(c *fiber.Ctx) error {
	var well_pressure dto.Well_pressure

	setDefaults(&well_pressure)

	if err := json.Unmarshal(c.Body(), &well_pressure); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_pressure.Uwi = id

    if well_pressure.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_pressure set source = :1, pressure_obs_no = :2, active_ind = :3, assigned_field = :4, base_depth = :5, base_depth_ouom = :6, base_strat_unit_id = :7, effective_date = :8, event_obs_no = :9, event_source = :10, expiry_date = :11, flow_casing_pressure = :12, flow_casing_pressure_ouom = :13, flow_tubing_pressure = :14, flow_tubing_pressure_ouom = :15, init_reservoir_pressure = :16, init_reservoir_press_ouom = :17, pool_datum = :18, pool_datum_depth = :19, pool_datum_depth_ouom = :20, pool_id = :21, ppdm_guid = :22, prod_string_id = :23, prod_string_source = :24, pr_str_form_obs_no = :25, remark = :26, shutin_casing_pressure = :27, shutin_casing_pressure_ouom = :28, shutin_tubing_pressure = :29, shutin_tubing_pressure_ouom = :30, strat_name_set_id = :31, top_depth = :32, top_depth_ouom = :33, top_strat_unit_id = :34, well_datum_depth = :35, well_datum_ouom = :36, well_test_num = :37, well_test_run_num = :38, well_test_source = :39, well_test_type = :40, row_changed_by = :41, row_changed_date = :42, row_created_by = :43, row_effective_date = :44, row_expiry_date = :45, row_quality = :46 where uwi = :48")
	if err != nil {
		return err
	}

	well_pressure.Row_changed_date = formatDate(well_pressure.Row_changed_date)
	well_pressure.Effective_date = formatDateString(well_pressure.Effective_date)
	well_pressure.Expiry_date = formatDateString(well_pressure.Expiry_date)
	well_pressure.Row_effective_date = formatDateString(well_pressure.Row_effective_date)
	well_pressure.Row_expiry_date = formatDateString(well_pressure.Row_expiry_date)






	rows, err := stmt.Exec(well_pressure.Source, well_pressure.Pressure_obs_no, well_pressure.Active_ind, well_pressure.Assigned_field, well_pressure.Base_depth, well_pressure.Base_depth_ouom, well_pressure.Base_strat_unit_id, well_pressure.Effective_date, well_pressure.Event_obs_no, well_pressure.Event_source, well_pressure.Expiry_date, well_pressure.Flow_casing_pressure, well_pressure.Flow_casing_pressure_ouom, well_pressure.Flow_tubing_pressure, well_pressure.Flow_tubing_pressure_ouom, well_pressure.Init_reservoir_pressure, well_pressure.Init_reservoir_press_ouom, well_pressure.Pool_datum, well_pressure.Pool_datum_depth, well_pressure.Pool_datum_depth_ouom, well_pressure.Pool_id, well_pressure.Ppdm_guid, well_pressure.Prod_string_id, well_pressure.Prod_string_source, well_pressure.Pr_str_form_obs_no, well_pressure.Remark, well_pressure.Shutin_casing_pressure, well_pressure.Shutin_casing_pressure_ouom, well_pressure.Shutin_tubing_pressure, well_pressure.Shutin_tubing_pressure_ouom, well_pressure.Strat_name_set_id, well_pressure.Top_depth, well_pressure.Top_depth_ouom, well_pressure.Top_strat_unit_id, well_pressure.Well_datum_depth, well_pressure.Well_datum_ouom, well_pressure.Well_test_num, well_pressure.Well_test_run_num, well_pressure.Well_test_source, well_pressure.Well_test_type, well_pressure.Row_changed_by, well_pressure.Row_changed_date, well_pressure.Row_created_by, well_pressure.Row_effective_date, well_pressure.Row_expiry_date, well_pressure.Row_quality, well_pressure.Uwi)
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

func PatchWellPressure(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_pressure set "
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

func DeleteWellPressure(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_pressure dto.Well_pressure
	well_pressure.Uwi = id

	stmt, err := db.Prepare("delete from well_pressure where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_pressure.Uwi)
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



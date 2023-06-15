package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellTreatment(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_treatment")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_treatment

	for rows.Next() {
		var well_treatment dto.Well_treatment
		if err := rows.Scan(&well_treatment.Uwi, &well_treatment.Source, &well_treatment.Treatment_type, &well_treatment.Treatment_obs_no, &well_treatment.Active_ind, &well_treatment.Additive_type, &well_treatment.Base_depth, &well_treatment.Base_depth_ouom, &well_treatment.Base_strat_unit_id, &well_treatment.Completion_obs_no, &well_treatment.Completion_source, &well_treatment.Effective_date, &well_treatment.Expiry_date, &well_treatment.Form_break_pressure, &well_treatment.Form_break_pressure_ouom, &well_treatment.Injection_rate, &well_treatment.Injection_rate_ouom, &well_treatment.Instant_si_pressure, &well_treatment.Instant_si_pressure_ouom, &well_treatment.Ppdm_guid, &well_treatment.Proppant_agent_amount, &well_treatment.Proppant_agent_amount_ouom, &well_treatment.Proppant_agent_amount_uom, &well_treatment.Proppant_agent_mesh_size, &well_treatment.Proppant_agent_type, &well_treatment.Remark, &well_treatment.Run_num, &well_treatment.Stage_no, &well_treatment.Strat_name_set_id, &well_treatment.Test_num, &well_treatment.Test_type, &well_treatment.Top_depth, &well_treatment.Top_depth_ouom, &well_treatment.Top_strat_unit_id, &well_treatment.Treatment_amount, &well_treatment.Treatment_amount_ouom, &well_treatment.Treatment_amount_uom, &well_treatment.Treatment_ba_id, &well_treatment.Treatment_date, &well_treatment.Treatment_fluid_type, &well_treatment.Treatment_pressure, &well_treatment.Treatment_pressure_ouom, &well_treatment.Treatment_status, &well_treatment.Treatment_status_type, &well_treatment.Well_test_source, &well_treatment.Row_changed_by, &well_treatment.Row_changed_date, &well_treatment.Row_created_by, &well_treatment.Row_created_date, &well_treatment.Row_effective_date, &well_treatment.Row_expiry_date, &well_treatment.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_treatment to result
		result = append(result, well_treatment)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellTreatment(c *fiber.Ctx) error {
	var well_treatment dto.Well_treatment

	setDefaults(&well_treatment)

	if err := json.Unmarshal(c.Body(), &well_treatment); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_treatment values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52)")
	if err != nil {
		return err
	}
	well_treatment.Row_created_date = formatDate(well_treatment.Row_created_date)
	well_treatment.Row_changed_date = formatDate(well_treatment.Row_changed_date)
	well_treatment.Effective_date = formatDateString(well_treatment.Effective_date)
	well_treatment.Expiry_date = formatDateString(well_treatment.Expiry_date)
	well_treatment.Treatment_date = formatDateString(well_treatment.Treatment_date)
	well_treatment.Row_effective_date = formatDateString(well_treatment.Row_effective_date)
	well_treatment.Row_expiry_date = formatDateString(well_treatment.Row_expiry_date)






	rows, err := stmt.Exec(well_treatment.Uwi, well_treatment.Source, well_treatment.Treatment_type, well_treatment.Treatment_obs_no, well_treatment.Active_ind, well_treatment.Additive_type, well_treatment.Base_depth, well_treatment.Base_depth_ouom, well_treatment.Base_strat_unit_id, well_treatment.Completion_obs_no, well_treatment.Completion_source, well_treatment.Effective_date, well_treatment.Expiry_date, well_treatment.Form_break_pressure, well_treatment.Form_break_pressure_ouom, well_treatment.Injection_rate, well_treatment.Injection_rate_ouom, well_treatment.Instant_si_pressure, well_treatment.Instant_si_pressure_ouom, well_treatment.Ppdm_guid, well_treatment.Proppant_agent_amount, well_treatment.Proppant_agent_amount_ouom, well_treatment.Proppant_agent_amount_uom, well_treatment.Proppant_agent_mesh_size, well_treatment.Proppant_agent_type, well_treatment.Remark, well_treatment.Run_num, well_treatment.Stage_no, well_treatment.Strat_name_set_id, well_treatment.Test_num, well_treatment.Test_type, well_treatment.Top_depth, well_treatment.Top_depth_ouom, well_treatment.Top_strat_unit_id, well_treatment.Treatment_amount, well_treatment.Treatment_amount_ouom, well_treatment.Treatment_amount_uom, well_treatment.Treatment_ba_id, well_treatment.Treatment_date, well_treatment.Treatment_fluid_type, well_treatment.Treatment_pressure, well_treatment.Treatment_pressure_ouom, well_treatment.Treatment_status, well_treatment.Treatment_status_type, well_treatment.Well_test_source, well_treatment.Row_changed_by, well_treatment.Row_changed_date, well_treatment.Row_created_by, well_treatment.Row_created_date, well_treatment.Row_effective_date, well_treatment.Row_expiry_date, well_treatment.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellTreatment(c *fiber.Ctx) error {
	var well_treatment dto.Well_treatment

	setDefaults(&well_treatment)

	if err := json.Unmarshal(c.Body(), &well_treatment); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_treatment.Uwi = id

    if well_treatment.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_treatment set source = :1, treatment_type = :2, treatment_obs_no = :3, active_ind = :4, additive_type = :5, base_depth = :6, base_depth_ouom = :7, base_strat_unit_id = :8, completion_obs_no = :9, completion_source = :10, effective_date = :11, expiry_date = :12, form_break_pressure = :13, form_break_pressure_ouom = :14, injection_rate = :15, injection_rate_ouom = :16, instant_si_pressure = :17, instant_si_pressure_ouom = :18, ppdm_guid = :19, proppant_agent_amount = :20, proppant_agent_amount_ouom = :21, proppant_agent_amount_uom = :22, proppant_agent_mesh_size = :23, proppant_agent_type = :24, remark = :25, run_num = :26, stage_no = :27, strat_name_set_id = :28, test_num = :29, test_type = :30, top_depth = :31, top_depth_ouom = :32, top_strat_unit_id = :33, treatment_amount = :34, treatment_amount_ouom = :35, treatment_amount_uom = :36, treatment_ba_id = :37, treatment_date = :38, treatment_fluid_type = :39, treatment_pressure = :40, treatment_pressure_ouom = :41, treatment_status = :42, treatment_status_type = :43, well_test_source = :44, row_changed_by = :45, row_changed_date = :46, row_created_by = :47, row_effective_date = :48, row_expiry_date = :49, row_quality = :50 where uwi = :52")
	if err != nil {
		return err
	}

	well_treatment.Row_changed_date = formatDate(well_treatment.Row_changed_date)
	well_treatment.Effective_date = formatDateString(well_treatment.Effective_date)
	well_treatment.Expiry_date = formatDateString(well_treatment.Expiry_date)
	well_treatment.Treatment_date = formatDateString(well_treatment.Treatment_date)
	well_treatment.Row_effective_date = formatDateString(well_treatment.Row_effective_date)
	well_treatment.Row_expiry_date = formatDateString(well_treatment.Row_expiry_date)






	rows, err := stmt.Exec(well_treatment.Source, well_treatment.Treatment_type, well_treatment.Treatment_obs_no, well_treatment.Active_ind, well_treatment.Additive_type, well_treatment.Base_depth, well_treatment.Base_depth_ouom, well_treatment.Base_strat_unit_id, well_treatment.Completion_obs_no, well_treatment.Completion_source, well_treatment.Effective_date, well_treatment.Expiry_date, well_treatment.Form_break_pressure, well_treatment.Form_break_pressure_ouom, well_treatment.Injection_rate, well_treatment.Injection_rate_ouom, well_treatment.Instant_si_pressure, well_treatment.Instant_si_pressure_ouom, well_treatment.Ppdm_guid, well_treatment.Proppant_agent_amount, well_treatment.Proppant_agent_amount_ouom, well_treatment.Proppant_agent_amount_uom, well_treatment.Proppant_agent_mesh_size, well_treatment.Proppant_agent_type, well_treatment.Remark, well_treatment.Run_num, well_treatment.Stage_no, well_treatment.Strat_name_set_id, well_treatment.Test_num, well_treatment.Test_type, well_treatment.Top_depth, well_treatment.Top_depth_ouom, well_treatment.Top_strat_unit_id, well_treatment.Treatment_amount, well_treatment.Treatment_amount_ouom, well_treatment.Treatment_amount_uom, well_treatment.Treatment_ba_id, well_treatment.Treatment_date, well_treatment.Treatment_fluid_type, well_treatment.Treatment_pressure, well_treatment.Treatment_pressure_ouom, well_treatment.Treatment_status, well_treatment.Treatment_status_type, well_treatment.Well_test_source, well_treatment.Row_changed_by, well_treatment.Row_changed_date, well_treatment.Row_created_by, well_treatment.Row_effective_date, well_treatment.Row_expiry_date, well_treatment.Row_quality, well_treatment.Uwi)
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

func PatchWellTreatment(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_treatment set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "treatment_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellTreatment(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_treatment dto.Well_treatment
	well_treatment.Uwi = id

	stmt, err := db.Prepare("delete from well_treatment where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_treatment.Uwi)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellTestComputAnal(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_test_comput_anal")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_test_comput_anal

	for rows.Next() {
		var well_test_comput_anal dto.Well_test_comput_anal
		if err := rows.Scan(&well_test_comput_anal.Uwi, &well_test_comput_anal.Source, &well_test_comput_anal.Test_type, &well_test_comput_anal.Run_num, &well_test_comput_anal.Test_num, &well_test_comput_anal.Report_no, &well_test_comput_anal.Active_ind, &well_test_comput_anal.Analysis_company, &well_test_comput_anal.Computed_permeability, &well_test_comput_anal.Computed_permeability_ouom, &well_test_comput_anal.Confidence_limit, &well_test_comput_anal.Confidence_limit_ouom, &well_test_comput_anal.Effective_date, &well_test_comput_anal.Est_damage_ratio, &well_test_comput_anal.Expiry_date, &well_test_comput_anal.Extrap_pressure_percent, &well_test_comput_anal.Final_reservoir_pressure, &well_test_comput_anal.Final_reservoir_press_ouom, &well_test_comput_anal.Gauge_depth, &well_test_comput_anal.Gauge_depth_ouom, &well_test_comput_anal.Investigation_radius, &well_test_comput_anal.Investigation_radius_ouom, &well_test_comput_anal.Max_reservoir_pressure, &well_test_comput_anal.Max_reservoir_pressure_ouom, &well_test_comput_anal.Potmtrc_surf_height, &well_test_comput_anal.Potmtrc_surf_height_ouom, &well_test_comput_anal.Ppdm_guid, &well_test_comput_anal.Production_index_rate, &well_test_comput_anal.Production_index_rate_ouom, &well_test_comput_anal.Recorder_id, &well_test_comput_anal.Remark, &well_test_comput_anal.Row_changed_by, &well_test_comput_anal.Row_changed_date, &well_test_comput_anal.Row_created_by, &well_test_comput_anal.Row_created_date, &well_test_comput_anal.Row_effective_date, &well_test_comput_anal.Row_expiry_date, &well_test_comput_anal.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_test_comput_anal to result
		result = append(result, well_test_comput_anal)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellTestComputAnal(c *fiber.Ctx) error {
	var well_test_comput_anal dto.Well_test_comput_anal

	setDefaults(&well_test_comput_anal)

	if err := json.Unmarshal(c.Body(), &well_test_comput_anal); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_test_comput_anal values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38)")
	if err != nil {
		return err
	}
	well_test_comput_anal.Row_created_date = formatDate(well_test_comput_anal.Row_created_date)
	well_test_comput_anal.Row_changed_date = formatDate(well_test_comput_anal.Row_changed_date)
	well_test_comput_anal.Effective_date = formatDateString(well_test_comput_anal.Effective_date)
	well_test_comput_anal.Expiry_date = formatDateString(well_test_comput_anal.Expiry_date)
	well_test_comput_anal.Row_effective_date = formatDateString(well_test_comput_anal.Row_effective_date)
	well_test_comput_anal.Row_expiry_date = formatDateString(well_test_comput_anal.Row_expiry_date)






	rows, err := stmt.Exec(well_test_comput_anal.Uwi, well_test_comput_anal.Source, well_test_comput_anal.Test_type, well_test_comput_anal.Run_num, well_test_comput_anal.Test_num, well_test_comput_anal.Report_no, well_test_comput_anal.Active_ind, well_test_comput_anal.Analysis_company, well_test_comput_anal.Computed_permeability, well_test_comput_anal.Computed_permeability_ouom, well_test_comput_anal.Confidence_limit, well_test_comput_anal.Confidence_limit_ouom, well_test_comput_anal.Effective_date, well_test_comput_anal.Est_damage_ratio, well_test_comput_anal.Expiry_date, well_test_comput_anal.Extrap_pressure_percent, well_test_comput_anal.Final_reservoir_pressure, well_test_comput_anal.Final_reservoir_press_ouom, well_test_comput_anal.Gauge_depth, well_test_comput_anal.Gauge_depth_ouom, well_test_comput_anal.Investigation_radius, well_test_comput_anal.Investigation_radius_ouom, well_test_comput_anal.Max_reservoir_pressure, well_test_comput_anal.Max_reservoir_pressure_ouom, well_test_comput_anal.Potmtrc_surf_height, well_test_comput_anal.Potmtrc_surf_height_ouom, well_test_comput_anal.Ppdm_guid, well_test_comput_anal.Production_index_rate, well_test_comput_anal.Production_index_rate_ouom, well_test_comput_anal.Recorder_id, well_test_comput_anal.Remark, well_test_comput_anal.Row_changed_by, well_test_comput_anal.Row_changed_date, well_test_comput_anal.Row_created_by, well_test_comput_anal.Row_created_date, well_test_comput_anal.Row_effective_date, well_test_comput_anal.Row_expiry_date, well_test_comput_anal.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellTestComputAnal(c *fiber.Ctx) error {
	var well_test_comput_anal dto.Well_test_comput_anal

	setDefaults(&well_test_comput_anal)

	if err := json.Unmarshal(c.Body(), &well_test_comput_anal); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_test_comput_anal.Uwi = id

    if well_test_comput_anal.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_test_comput_anal set source = :1, test_type = :2, run_num = :3, test_num = :4, report_no = :5, active_ind = :6, analysis_company = :7, computed_permeability = :8, computed_permeability_ouom = :9, confidence_limit = :10, confidence_limit_ouom = :11, effective_date = :12, est_damage_ratio = :13, expiry_date = :14, extrap_pressure_percent = :15, final_reservoir_pressure = :16, final_reservoir_press_ouom = :17, gauge_depth = :18, gauge_depth_ouom = :19, investigation_radius = :20, investigation_radius_ouom = :21, max_reservoir_pressure = :22, max_reservoir_pressure_ouom = :23, potmtrc_surf_height = :24, potmtrc_surf_height_ouom = :25, ppdm_guid = :26, production_index_rate = :27, production_index_rate_ouom = :28, recorder_id = :29, remark = :30, row_changed_by = :31, row_changed_date = :32, row_created_by = :33, row_effective_date = :34, row_expiry_date = :35, row_quality = :36 where uwi = :38")
	if err != nil {
		return err
	}

	well_test_comput_anal.Row_changed_date = formatDate(well_test_comput_anal.Row_changed_date)
	well_test_comput_anal.Effective_date = formatDateString(well_test_comput_anal.Effective_date)
	well_test_comput_anal.Expiry_date = formatDateString(well_test_comput_anal.Expiry_date)
	well_test_comput_anal.Row_effective_date = formatDateString(well_test_comput_anal.Row_effective_date)
	well_test_comput_anal.Row_expiry_date = formatDateString(well_test_comput_anal.Row_expiry_date)






	rows, err := stmt.Exec(well_test_comput_anal.Source, well_test_comput_anal.Test_type, well_test_comput_anal.Run_num, well_test_comput_anal.Test_num, well_test_comput_anal.Report_no, well_test_comput_anal.Active_ind, well_test_comput_anal.Analysis_company, well_test_comput_anal.Computed_permeability, well_test_comput_anal.Computed_permeability_ouom, well_test_comput_anal.Confidence_limit, well_test_comput_anal.Confidence_limit_ouom, well_test_comput_anal.Effective_date, well_test_comput_anal.Est_damage_ratio, well_test_comput_anal.Expiry_date, well_test_comput_anal.Extrap_pressure_percent, well_test_comput_anal.Final_reservoir_pressure, well_test_comput_anal.Final_reservoir_press_ouom, well_test_comput_anal.Gauge_depth, well_test_comput_anal.Gauge_depth_ouom, well_test_comput_anal.Investigation_radius, well_test_comput_anal.Investigation_radius_ouom, well_test_comput_anal.Max_reservoir_pressure, well_test_comput_anal.Max_reservoir_pressure_ouom, well_test_comput_anal.Potmtrc_surf_height, well_test_comput_anal.Potmtrc_surf_height_ouom, well_test_comput_anal.Ppdm_guid, well_test_comput_anal.Production_index_rate, well_test_comput_anal.Production_index_rate_ouom, well_test_comput_anal.Recorder_id, well_test_comput_anal.Remark, well_test_comput_anal.Row_changed_by, well_test_comput_anal.Row_changed_date, well_test_comput_anal.Row_created_by, well_test_comput_anal.Row_effective_date, well_test_comput_anal.Row_expiry_date, well_test_comput_anal.Row_quality, well_test_comput_anal.Uwi)
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

func PatchWellTestComputAnal(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_test_comput_anal set "
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

func DeleteWellTestComputAnal(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_test_comput_anal dto.Well_test_comput_anal
	well_test_comput_anal.Uwi = id

	stmt, err := db.Prepare("delete from well_test_comput_anal where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_test_comput_anal.Uwi)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellTestAnalysis(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_test_analysis")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_test_analysis

	for rows.Next() {
		var well_test_analysis dto.Well_test_analysis
		if err := rows.Scan(&well_test_analysis.Uwi, &well_test_analysis.Source, &well_test_analysis.Test_type, &well_test_analysis.Run_num, &well_test_analysis.Test_num, &well_test_analysis.Period_type, &well_test_analysis.Period_obs_no, &well_test_analysis.Analysis_obs_no, &well_test_analysis.Active_ind, &well_test_analysis.Bsw, &well_test_analysis.Completion_obs_no, &well_test_analysis.Condensate_gravity, &well_test_analysis.Condensate_ratio, &well_test_analysis.Condensate_ratio_ouom, &well_test_analysis.Condensate_temperature, &well_test_analysis.Condensate_temperature_ouom, &well_test_analysis.Effective_date, &well_test_analysis.Expiry_date, &well_test_analysis.Fluid_type, &well_test_analysis.Gas_content, &well_test_analysis.Gas_gravity, &well_test_analysis.Gor, &well_test_analysis.Gor_ouom, &well_test_analysis.Gwr, &well_test_analysis.Gwr_ouom, &well_test_analysis.H2s_percent, &well_test_analysis.Lgr, &well_test_analysis.Lgr_ouom, &well_test_analysis.Oil_density, &well_test_analysis.Oil_density_ouom, &well_test_analysis.Oil_gravity, &well_test_analysis.Oil_temperature, &well_test_analysis.Oil_temperature_ouom, &well_test_analysis.Ppdm_guid, &well_test_analysis.Remark, &well_test_analysis.Salinity_type, &well_test_analysis.Sulphur_percent, &well_test_analysis.Water_cut, &well_test_analysis.Water_resistivity, &well_test_analysis.Water_resistivity_ouom, &well_test_analysis.Water_salinity, &well_test_analysis.Water_salinity_ouom, &well_test_analysis.Water_temperature, &well_test_analysis.Water_temperature_ouom, &well_test_analysis.Wor, &well_test_analysis.Row_changed_by, &well_test_analysis.Row_changed_date, &well_test_analysis.Row_created_by, &well_test_analysis.Row_created_date, &well_test_analysis.Row_effective_date, &well_test_analysis.Row_expiry_date, &well_test_analysis.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_test_analysis to result
		result = append(result, well_test_analysis)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellTestAnalysis(c *fiber.Ctx) error {
	var well_test_analysis dto.Well_test_analysis

	setDefaults(&well_test_analysis)

	if err := json.Unmarshal(c.Body(), &well_test_analysis); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_test_analysis values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52)")
	if err != nil {
		return err
	}
	well_test_analysis.Row_created_date = formatDate(well_test_analysis.Row_created_date)
	well_test_analysis.Row_changed_date = formatDate(well_test_analysis.Row_changed_date)
	well_test_analysis.Effective_date = formatDateString(well_test_analysis.Effective_date)
	well_test_analysis.Expiry_date = formatDateString(well_test_analysis.Expiry_date)
	well_test_analysis.Row_effective_date = formatDateString(well_test_analysis.Row_effective_date)
	well_test_analysis.Row_expiry_date = formatDateString(well_test_analysis.Row_expiry_date)






	rows, err := stmt.Exec(well_test_analysis.Uwi, well_test_analysis.Source, well_test_analysis.Test_type, well_test_analysis.Run_num, well_test_analysis.Test_num, well_test_analysis.Period_type, well_test_analysis.Period_obs_no, well_test_analysis.Analysis_obs_no, well_test_analysis.Active_ind, well_test_analysis.Bsw, well_test_analysis.Completion_obs_no, well_test_analysis.Condensate_gravity, well_test_analysis.Condensate_ratio, well_test_analysis.Condensate_ratio_ouom, well_test_analysis.Condensate_temperature, well_test_analysis.Condensate_temperature_ouom, well_test_analysis.Effective_date, well_test_analysis.Expiry_date, well_test_analysis.Fluid_type, well_test_analysis.Gas_content, well_test_analysis.Gas_gravity, well_test_analysis.Gor, well_test_analysis.Gor_ouom, well_test_analysis.Gwr, well_test_analysis.Gwr_ouom, well_test_analysis.H2s_percent, well_test_analysis.Lgr, well_test_analysis.Lgr_ouom, well_test_analysis.Oil_density, well_test_analysis.Oil_density_ouom, well_test_analysis.Oil_gravity, well_test_analysis.Oil_temperature, well_test_analysis.Oil_temperature_ouom, well_test_analysis.Ppdm_guid, well_test_analysis.Remark, well_test_analysis.Salinity_type, well_test_analysis.Sulphur_percent, well_test_analysis.Water_cut, well_test_analysis.Water_resistivity, well_test_analysis.Water_resistivity_ouom, well_test_analysis.Water_salinity, well_test_analysis.Water_salinity_ouom, well_test_analysis.Water_temperature, well_test_analysis.Water_temperature_ouom, well_test_analysis.Wor, well_test_analysis.Row_changed_by, well_test_analysis.Row_changed_date, well_test_analysis.Row_created_by, well_test_analysis.Row_created_date, well_test_analysis.Row_effective_date, well_test_analysis.Row_expiry_date, well_test_analysis.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellTestAnalysis(c *fiber.Ctx) error {
	var well_test_analysis dto.Well_test_analysis

	setDefaults(&well_test_analysis)

	if err := json.Unmarshal(c.Body(), &well_test_analysis); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_test_analysis.Uwi = id

    if well_test_analysis.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_test_analysis set source = :1, test_type = :2, run_num = :3, test_num = :4, period_type = :5, period_obs_no = :6, analysis_obs_no = :7, active_ind = :8, bsw = :9, completion_obs_no = :10, condensate_gravity = :11, condensate_ratio = :12, condensate_ratio_ouom = :13, condensate_temperature = :14, condensate_temperature_ouom = :15, effective_date = :16, expiry_date = :17, fluid_type = :18, gas_content = :19, gas_gravity = :20, gor = :21, gor_ouom = :22, gwr = :23, gwr_ouom = :24, h2s_percent = :25, lgr = :26, lgr_ouom = :27, oil_density = :28, oil_density_ouom = :29, oil_gravity = :30, oil_temperature = :31, oil_temperature_ouom = :32, ppdm_guid = :33, remark = :34, salinity_type = :35, sulphur_percent = :36, water_cut = :37, water_resistivity = :38, water_resistivity_ouom = :39, water_salinity = :40, water_salinity_ouom = :41, water_temperature = :42, water_temperature_ouom = :43, wor = :44, row_changed_by = :45, row_changed_date = :46, row_created_by = :47, row_effective_date = :48, row_expiry_date = :49, row_quality = :50 where uwi = :52")
	if err != nil {
		return err
	}

	well_test_analysis.Row_changed_date = formatDate(well_test_analysis.Row_changed_date)
	well_test_analysis.Effective_date = formatDateString(well_test_analysis.Effective_date)
	well_test_analysis.Expiry_date = formatDateString(well_test_analysis.Expiry_date)
	well_test_analysis.Row_effective_date = formatDateString(well_test_analysis.Row_effective_date)
	well_test_analysis.Row_expiry_date = formatDateString(well_test_analysis.Row_expiry_date)






	rows, err := stmt.Exec(well_test_analysis.Source, well_test_analysis.Test_type, well_test_analysis.Run_num, well_test_analysis.Test_num, well_test_analysis.Period_type, well_test_analysis.Period_obs_no, well_test_analysis.Analysis_obs_no, well_test_analysis.Active_ind, well_test_analysis.Bsw, well_test_analysis.Completion_obs_no, well_test_analysis.Condensate_gravity, well_test_analysis.Condensate_ratio, well_test_analysis.Condensate_ratio_ouom, well_test_analysis.Condensate_temperature, well_test_analysis.Condensate_temperature_ouom, well_test_analysis.Effective_date, well_test_analysis.Expiry_date, well_test_analysis.Fluid_type, well_test_analysis.Gas_content, well_test_analysis.Gas_gravity, well_test_analysis.Gor, well_test_analysis.Gor_ouom, well_test_analysis.Gwr, well_test_analysis.Gwr_ouom, well_test_analysis.H2s_percent, well_test_analysis.Lgr, well_test_analysis.Lgr_ouom, well_test_analysis.Oil_density, well_test_analysis.Oil_density_ouom, well_test_analysis.Oil_gravity, well_test_analysis.Oil_temperature, well_test_analysis.Oil_temperature_ouom, well_test_analysis.Ppdm_guid, well_test_analysis.Remark, well_test_analysis.Salinity_type, well_test_analysis.Sulphur_percent, well_test_analysis.Water_cut, well_test_analysis.Water_resistivity, well_test_analysis.Water_resistivity_ouom, well_test_analysis.Water_salinity, well_test_analysis.Water_salinity_ouom, well_test_analysis.Water_temperature, well_test_analysis.Water_temperature_ouom, well_test_analysis.Wor, well_test_analysis.Row_changed_by, well_test_analysis.Row_changed_date, well_test_analysis.Row_created_by, well_test_analysis.Row_effective_date, well_test_analysis.Row_expiry_date, well_test_analysis.Row_quality, well_test_analysis.Uwi)
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

func PatchWellTestAnalysis(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_test_analysis set "
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

func DeleteWellTestAnalysis(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_test_analysis dto.Well_test_analysis
	well_test_analysis.Uwi = id

	stmt, err := db.Prepare("delete from well_test_analysis where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_test_analysis.Uwi)
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



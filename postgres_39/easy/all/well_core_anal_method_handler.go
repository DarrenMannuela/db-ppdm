package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellCoreAnalMethod(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_core_anal_method")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_core_anal_method

	for rows.Next() {
		var well_core_anal_method dto.Well_core_anal_method
		if err := rows.Scan(&well_core_anal_method.Uwi, &well_core_anal_method.Source, &well_core_anal_method.Core_id, &well_core_anal_method.Analysis_obs_no, &well_core_anal_method.Active_ind, &well_core_anal_method.Core_solvent, &well_core_anal_method.Cutting_fluid, &well_core_anal_method.Density_analysis, &well_core_anal_method.Drying_method, &well_core_anal_method.Drying_temperature, &well_core_anal_method.Drying_temperature_ouom, &well_core_anal_method.Drying_time_elapsed, &well_core_anal_method.Drying_time_elapsed_ouom, &well_core_anal_method.Effective_date, &well_core_anal_method.Expiry_date, &well_core_anal_method.Extract_method, &well_core_anal_method.Extract_time_elapsed, &well_core_anal_method.Extract_time_elapsed_ouom, &well_core_anal_method.Fluid_sat_analysis, &well_core_anal_method.Permeability_analysis, &well_core_anal_method.Porosity_analysis, &well_core_anal_method.Ppdm_guid, &well_core_anal_method.Remark, &well_core_anal_method.Row_changed_by, &well_core_anal_method.Row_changed_date, &well_core_anal_method.Row_created_by, &well_core_anal_method.Row_created_date, &well_core_anal_method.Row_effective_date, &well_core_anal_method.Row_expiry_date, &well_core_anal_method.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_core_anal_method to result
		result = append(result, well_core_anal_method)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellCoreAnalMethod(c *fiber.Ctx) error {
	var well_core_anal_method dto.Well_core_anal_method

	setDefaults(&well_core_anal_method)

	if err := json.Unmarshal(c.Body(), &well_core_anal_method); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_core_anal_method values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30)")
	if err != nil {
		return err
	}
	well_core_anal_method.Row_created_date = formatDate(well_core_anal_method.Row_created_date)
	well_core_anal_method.Row_changed_date = formatDate(well_core_anal_method.Row_changed_date)
	well_core_anal_method.Effective_date = formatDateString(well_core_anal_method.Effective_date)
	well_core_anal_method.Expiry_date = formatDateString(well_core_anal_method.Expiry_date)
	well_core_anal_method.Row_effective_date = formatDateString(well_core_anal_method.Row_effective_date)
	well_core_anal_method.Row_expiry_date = formatDateString(well_core_anal_method.Row_expiry_date)






	rows, err := stmt.Exec(well_core_anal_method.Uwi, well_core_anal_method.Source, well_core_anal_method.Core_id, well_core_anal_method.Analysis_obs_no, well_core_anal_method.Active_ind, well_core_anal_method.Core_solvent, well_core_anal_method.Cutting_fluid, well_core_anal_method.Density_analysis, well_core_anal_method.Drying_method, well_core_anal_method.Drying_temperature, well_core_anal_method.Drying_temperature_ouom, well_core_anal_method.Drying_time_elapsed, well_core_anal_method.Drying_time_elapsed_ouom, well_core_anal_method.Effective_date, well_core_anal_method.Expiry_date, well_core_anal_method.Extract_method, well_core_anal_method.Extract_time_elapsed, well_core_anal_method.Extract_time_elapsed_ouom, well_core_anal_method.Fluid_sat_analysis, well_core_anal_method.Permeability_analysis, well_core_anal_method.Porosity_analysis, well_core_anal_method.Ppdm_guid, well_core_anal_method.Remark, well_core_anal_method.Row_changed_by, well_core_anal_method.Row_changed_date, well_core_anal_method.Row_created_by, well_core_anal_method.Row_created_date, well_core_anal_method.Row_effective_date, well_core_anal_method.Row_expiry_date, well_core_anal_method.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellCoreAnalMethod(c *fiber.Ctx) error {
	var well_core_anal_method dto.Well_core_anal_method

	setDefaults(&well_core_anal_method)

	if err := json.Unmarshal(c.Body(), &well_core_anal_method); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_core_anal_method.Uwi = id

    if well_core_anal_method.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_core_anal_method set source = :1, core_id = :2, analysis_obs_no = :3, active_ind = :4, core_solvent = :5, cutting_fluid = :6, density_analysis = :7, drying_method = :8, drying_temperature = :9, drying_temperature_ouom = :10, drying_time_elapsed = :11, drying_time_elapsed_ouom = :12, effective_date = :13, expiry_date = :14, extract_method = :15, extract_time_elapsed = :16, extract_time_elapsed_ouom = :17, fluid_sat_analysis = :18, permeability_analysis = :19, porosity_analysis = :20, ppdm_guid = :21, remark = :22, row_changed_by = :23, row_changed_date = :24, row_created_by = :25, row_effective_date = :26, row_expiry_date = :27, row_quality = :28 where uwi = :30")
	if err != nil {
		return err
	}

	well_core_anal_method.Row_changed_date = formatDate(well_core_anal_method.Row_changed_date)
	well_core_anal_method.Effective_date = formatDateString(well_core_anal_method.Effective_date)
	well_core_anal_method.Expiry_date = formatDateString(well_core_anal_method.Expiry_date)
	well_core_anal_method.Row_effective_date = formatDateString(well_core_anal_method.Row_effective_date)
	well_core_anal_method.Row_expiry_date = formatDateString(well_core_anal_method.Row_expiry_date)






	rows, err := stmt.Exec(well_core_anal_method.Source, well_core_anal_method.Core_id, well_core_anal_method.Analysis_obs_no, well_core_anal_method.Active_ind, well_core_anal_method.Core_solvent, well_core_anal_method.Cutting_fluid, well_core_anal_method.Density_analysis, well_core_anal_method.Drying_method, well_core_anal_method.Drying_temperature, well_core_anal_method.Drying_temperature_ouom, well_core_anal_method.Drying_time_elapsed, well_core_anal_method.Drying_time_elapsed_ouom, well_core_anal_method.Effective_date, well_core_anal_method.Expiry_date, well_core_anal_method.Extract_method, well_core_anal_method.Extract_time_elapsed, well_core_anal_method.Extract_time_elapsed_ouom, well_core_anal_method.Fluid_sat_analysis, well_core_anal_method.Permeability_analysis, well_core_anal_method.Porosity_analysis, well_core_anal_method.Ppdm_guid, well_core_anal_method.Remark, well_core_anal_method.Row_changed_by, well_core_anal_method.Row_changed_date, well_core_anal_method.Row_created_by, well_core_anal_method.Row_effective_date, well_core_anal_method.Row_expiry_date, well_core_anal_method.Row_quality, well_core_anal_method.Uwi)
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

func PatchWellCoreAnalMethod(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_core_anal_method set "
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

func DeleteWellCoreAnalMethod(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_core_anal_method dto.Well_core_anal_method
	well_core_anal_method.Uwi = id

	stmt, err := db.Prepare("delete from well_core_anal_method where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_core_anal_method.Uwi)
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



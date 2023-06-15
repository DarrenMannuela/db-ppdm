package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellCoreSampleAnal(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_core_sample_anal")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_core_sample_anal

	for rows.Next() {
		var well_core_sample_anal dto.Well_core_sample_anal
		if err := rows.Scan(&well_core_sample_anal.Uwi, &well_core_sample_anal.Source, &well_core_sample_anal.Core_id, &well_core_sample_anal.Analysis_obs_no, &well_core_sample_anal.Sample_num, &well_core_sample_anal.Sample_analysis_obs_no, &well_core_sample_anal.Active_ind, &well_core_sample_anal.Analysis_source, &well_core_sample_anal.Bulk_density, &well_core_sample_anal.Bulk_density_ouom, &well_core_sample_anal.Bulk_mass_oil_sat, &well_core_sample_anal.Bulk_mass_oil_sat_ouom, &well_core_sample_anal.Bulk_mass_sand_sat, &well_core_sample_anal.Bulk_mass_sand_sat_ouom, &well_core_sample_anal.Bulk_mass_water_sat, &well_core_sample_anal.Bulk_mass_water_sat_ouom, &well_core_sample_anal.Bulk_volume_oil_sat, &well_core_sample_anal.Bulk_volume_water_sat, &well_core_sample_anal.Confine_perm_pressure, &well_core_sample_anal.Confine_perm_pressure_ouom, &well_core_sample_anal.Confine_por_pressure, &well_core_sample_anal.Confine_por_pressure_ouom, &well_core_sample_anal.Confine_sat_pressure, &well_core_sample_anal.Confine_sat_pressure_ouom, &well_core_sample_anal.Effective_date, &well_core_sample_anal.Effective_porosity, &well_core_sample_anal.Expiry_date, &well_core_sample_anal.Gas_sat_volume, &well_core_sample_anal.Grain_density, &well_core_sample_anal.Grain_density_ouom, &well_core_sample_anal.Grain_mass_oil_sat, &well_core_sample_anal.Grain_mass_oil_sat_ouom, &well_core_sample_anal.Grain_mass_water_sat, &well_core_sample_anal.Grain_mass_water_sat_ouom, &well_core_sample_anal.Interval_depth, &well_core_sample_anal.Interval_depth_ouom, &well_core_sample_anal.Interval_length, &well_core_sample_anal.Interval_length_ouom, &well_core_sample_anal.K90, &well_core_sample_anal.K90_ouom, &well_core_sample_anal.K90_qualifier, &well_core_sample_anal.Kmax, &well_core_sample_anal.Kmax_ouom, &well_core_sample_anal.Kmax_qualifier, &well_core_sample_anal.Kvert, &well_core_sample_anal.Kvert_ouom, &well_core_sample_anal.Kvert_qualifier, &well_core_sample_anal.Oil_sat, &well_core_sample_anal.Pore_volume_gas_sat, &well_core_sample_anal.Pore_volume_oil_sat, &well_core_sample_anal.Pore_volume_water_sat, &well_core_sample_anal.Porosity, &well_core_sample_anal.Ppdm_guid, &well_core_sample_anal.Remark, &well_core_sample_anal.Top_depth, &well_core_sample_anal.Top_depth_ouom, &well_core_sample_anal.Water_sat, &well_core_sample_anal.Row_changed_by, &well_core_sample_anal.Row_changed_date, &well_core_sample_anal.Row_created_by, &well_core_sample_anal.Row_created_date, &well_core_sample_anal.Row_effective_date, &well_core_sample_anal.Row_expiry_date, &well_core_sample_anal.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_core_sample_anal to result
		result = append(result, well_core_sample_anal)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellCoreSampleAnal(c *fiber.Ctx) error {
	var well_core_sample_anal dto.Well_core_sample_anal

	setDefaults(&well_core_sample_anal)

	if err := json.Unmarshal(c.Body(), &well_core_sample_anal); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_core_sample_anal values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64)")
	if err != nil {
		return err
	}
	well_core_sample_anal.Row_created_date = formatDate(well_core_sample_anal.Row_created_date)
	well_core_sample_anal.Row_changed_date = formatDate(well_core_sample_anal.Row_changed_date)
	well_core_sample_anal.Effective_date = formatDateString(well_core_sample_anal.Effective_date)
	well_core_sample_anal.Expiry_date = formatDateString(well_core_sample_anal.Expiry_date)
	well_core_sample_anal.Row_effective_date = formatDateString(well_core_sample_anal.Row_effective_date)
	well_core_sample_anal.Row_expiry_date = formatDateString(well_core_sample_anal.Row_expiry_date)






	rows, err := stmt.Exec(well_core_sample_anal.Uwi, well_core_sample_anal.Source, well_core_sample_anal.Core_id, well_core_sample_anal.Analysis_obs_no, well_core_sample_anal.Sample_num, well_core_sample_anal.Sample_analysis_obs_no, well_core_sample_anal.Active_ind, well_core_sample_anal.Analysis_source, well_core_sample_anal.Bulk_density, well_core_sample_anal.Bulk_density_ouom, well_core_sample_anal.Bulk_mass_oil_sat, well_core_sample_anal.Bulk_mass_oil_sat_ouom, well_core_sample_anal.Bulk_mass_sand_sat, well_core_sample_anal.Bulk_mass_sand_sat_ouom, well_core_sample_anal.Bulk_mass_water_sat, well_core_sample_anal.Bulk_mass_water_sat_ouom, well_core_sample_anal.Bulk_volume_oil_sat, well_core_sample_anal.Bulk_volume_water_sat, well_core_sample_anal.Confine_perm_pressure, well_core_sample_anal.Confine_perm_pressure_ouom, well_core_sample_anal.Confine_por_pressure, well_core_sample_anal.Confine_por_pressure_ouom, well_core_sample_anal.Confine_sat_pressure, well_core_sample_anal.Confine_sat_pressure_ouom, well_core_sample_anal.Effective_date, well_core_sample_anal.Effective_porosity, well_core_sample_anal.Expiry_date, well_core_sample_anal.Gas_sat_volume, well_core_sample_anal.Grain_density, well_core_sample_anal.Grain_density_ouom, well_core_sample_anal.Grain_mass_oil_sat, well_core_sample_anal.Grain_mass_oil_sat_ouom, well_core_sample_anal.Grain_mass_water_sat, well_core_sample_anal.Grain_mass_water_sat_ouom, well_core_sample_anal.Interval_depth, well_core_sample_anal.Interval_depth_ouom, well_core_sample_anal.Interval_length, well_core_sample_anal.Interval_length_ouom, well_core_sample_anal.K90, well_core_sample_anal.K90_ouom, well_core_sample_anal.K90_qualifier, well_core_sample_anal.Kmax, well_core_sample_anal.Kmax_ouom, well_core_sample_anal.Kmax_qualifier, well_core_sample_anal.Kvert, well_core_sample_anal.Kvert_ouom, well_core_sample_anal.Kvert_qualifier, well_core_sample_anal.Oil_sat, well_core_sample_anal.Pore_volume_gas_sat, well_core_sample_anal.Pore_volume_oil_sat, well_core_sample_anal.Pore_volume_water_sat, well_core_sample_anal.Porosity, well_core_sample_anal.Ppdm_guid, well_core_sample_anal.Remark, well_core_sample_anal.Top_depth, well_core_sample_anal.Top_depth_ouom, well_core_sample_anal.Water_sat, well_core_sample_anal.Row_changed_by, well_core_sample_anal.Row_changed_date, well_core_sample_anal.Row_created_by, well_core_sample_anal.Row_created_date, well_core_sample_anal.Row_effective_date, well_core_sample_anal.Row_expiry_date, well_core_sample_anal.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellCoreSampleAnal(c *fiber.Ctx) error {
	var well_core_sample_anal dto.Well_core_sample_anal

	setDefaults(&well_core_sample_anal)

	if err := json.Unmarshal(c.Body(), &well_core_sample_anal); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_core_sample_anal.Uwi = id

    if well_core_sample_anal.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_core_sample_anal set source = :1, core_id = :2, analysis_obs_no = :3, sample_num = :4, sample_analysis_obs_no = :5, active_ind = :6, analysis_source = :7, bulk_density = :8, bulk_density_ouom = :9, bulk_mass_oil_sat = :10, bulk_mass_oil_sat_ouom = :11, bulk_mass_sand_sat = :12, bulk_mass_sand_sat_ouom = :13, bulk_mass_water_sat = :14, bulk_mass_water_sat_ouom = :15, bulk_volume_oil_sat = :16, bulk_volume_water_sat = :17, confine_perm_pressure = :18, confine_perm_pressure_ouom = :19, confine_por_pressure = :20, confine_por_pressure_ouom = :21, confine_sat_pressure = :22, confine_sat_pressure_ouom = :23, effective_date = :24, effective_porosity = :25, expiry_date = :26, gas_sat_volume = :27, grain_density = :28, grain_density_ouom = :29, grain_mass_oil_sat = :30, grain_mass_oil_sat_ouom = :31, grain_mass_water_sat = :32, grain_mass_water_sat_ouom = :33, interval_depth = :34, interval_depth_ouom = :35, interval_length = :36, interval_length_ouom = :37, k90 = :38, k90_ouom = :39, k90_qualifier = :40, kmax = :41, kmax_ouom = :42, kmax_qualifier = :43, kvert = :44, kvert_ouom = :45, kvert_qualifier = :46, oil_sat = :47, pore_volume_gas_sat = :48, pore_volume_oil_sat = :49, pore_volume_water_sat = :50, porosity = :51, ppdm_guid = :52, remark = :53, top_depth = :54, top_depth_ouom = :55, water_sat = :56, row_changed_by = :57, row_changed_date = :58, row_created_by = :59, row_effective_date = :60, row_expiry_date = :61, row_quality = :62 where uwi = :64")
	if err != nil {
		return err
	}

	well_core_sample_anal.Row_changed_date = formatDate(well_core_sample_anal.Row_changed_date)
	well_core_sample_anal.Effective_date = formatDateString(well_core_sample_anal.Effective_date)
	well_core_sample_anal.Expiry_date = formatDateString(well_core_sample_anal.Expiry_date)
	well_core_sample_anal.Row_effective_date = formatDateString(well_core_sample_anal.Row_effective_date)
	well_core_sample_anal.Row_expiry_date = formatDateString(well_core_sample_anal.Row_expiry_date)






	rows, err := stmt.Exec(well_core_sample_anal.Source, well_core_sample_anal.Core_id, well_core_sample_anal.Analysis_obs_no, well_core_sample_anal.Sample_num, well_core_sample_anal.Sample_analysis_obs_no, well_core_sample_anal.Active_ind, well_core_sample_anal.Analysis_source, well_core_sample_anal.Bulk_density, well_core_sample_anal.Bulk_density_ouom, well_core_sample_anal.Bulk_mass_oil_sat, well_core_sample_anal.Bulk_mass_oil_sat_ouom, well_core_sample_anal.Bulk_mass_sand_sat, well_core_sample_anal.Bulk_mass_sand_sat_ouom, well_core_sample_anal.Bulk_mass_water_sat, well_core_sample_anal.Bulk_mass_water_sat_ouom, well_core_sample_anal.Bulk_volume_oil_sat, well_core_sample_anal.Bulk_volume_water_sat, well_core_sample_anal.Confine_perm_pressure, well_core_sample_anal.Confine_perm_pressure_ouom, well_core_sample_anal.Confine_por_pressure, well_core_sample_anal.Confine_por_pressure_ouom, well_core_sample_anal.Confine_sat_pressure, well_core_sample_anal.Confine_sat_pressure_ouom, well_core_sample_anal.Effective_date, well_core_sample_anal.Effective_porosity, well_core_sample_anal.Expiry_date, well_core_sample_anal.Gas_sat_volume, well_core_sample_anal.Grain_density, well_core_sample_anal.Grain_density_ouom, well_core_sample_anal.Grain_mass_oil_sat, well_core_sample_anal.Grain_mass_oil_sat_ouom, well_core_sample_anal.Grain_mass_water_sat, well_core_sample_anal.Grain_mass_water_sat_ouom, well_core_sample_anal.Interval_depth, well_core_sample_anal.Interval_depth_ouom, well_core_sample_anal.Interval_length, well_core_sample_anal.Interval_length_ouom, well_core_sample_anal.K90, well_core_sample_anal.K90_ouom, well_core_sample_anal.K90_qualifier, well_core_sample_anal.Kmax, well_core_sample_anal.Kmax_ouom, well_core_sample_anal.Kmax_qualifier, well_core_sample_anal.Kvert, well_core_sample_anal.Kvert_ouom, well_core_sample_anal.Kvert_qualifier, well_core_sample_anal.Oil_sat, well_core_sample_anal.Pore_volume_gas_sat, well_core_sample_anal.Pore_volume_oil_sat, well_core_sample_anal.Pore_volume_water_sat, well_core_sample_anal.Porosity, well_core_sample_anal.Ppdm_guid, well_core_sample_anal.Remark, well_core_sample_anal.Top_depth, well_core_sample_anal.Top_depth_ouom, well_core_sample_anal.Water_sat, well_core_sample_anal.Row_changed_by, well_core_sample_anal.Row_changed_date, well_core_sample_anal.Row_created_by, well_core_sample_anal.Row_effective_date, well_core_sample_anal.Row_expiry_date, well_core_sample_anal.Row_quality, well_core_sample_anal.Uwi)
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

func PatchWellCoreSampleAnal(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_core_sample_anal set "
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

func DeleteWellCoreSampleAnal(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_core_sample_anal dto.Well_core_sample_anal
	well_core_sample_anal.Uwi = id

	stmt, err := db.Prepare("delete from well_core_sample_anal where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_core_sample_anal.Uwi)
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



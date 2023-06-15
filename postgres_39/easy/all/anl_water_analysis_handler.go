package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlWaterAnalysis(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_water_analysis")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_water_analysis

	for rows.Next() {
		var anl_water_analysis dto.Anl_water_analysis
		if err := rows.Scan(&anl_water_analysis.Analysis_id, &anl_water_analysis.Anl_source, &anl_water_analysis.Water_analysis_obs_no, &anl_water_analysis.Active_ind, &anl_water_analysis.Barium, &anl_water_analysis.Barium_ouom, &anl_water_analysis.Bicarbonate, &anl_water_analysis.Bicarbonate_ouom, &anl_water_analysis.Boron, &anl_water_analysis.Boron_ouom, &anl_water_analysis.Bromine, &anl_water_analysis.Bromine_ouom, &anl_water_analysis.Calcium, &anl_water_analysis.Calcium_ouom, &anl_water_analysis.Carbonate, &anl_water_analysis.Carbonate_ouom, &anl_water_analysis.Chlorine, &anl_water_analysis.Chlorine_ouom, &anl_water_analysis.Conductivity, &anl_water_analysis.Conductivity_ouom, &anl_water_analysis.Effective_date, &anl_water_analysis.Expiry_date, &anl_water_analysis.Hydrogen_sulphide, &anl_water_analysis.Hydrogen_sulphide_ouom, &anl_water_analysis.Hydroxide, &anl_water_analysis.Hydroxide_ouom, &anl_water_analysis.Iodine, &anl_water_analysis.Iodine_ouom, &anl_water_analysis.Ion_ouom, &anl_water_analysis.Ion_uom, &anl_water_analysis.Iron, &anl_water_analysis.Iron_ouom, &anl_water_analysis.Magnesium, &anl_water_analysis.Magnesium_ouom, &anl_water_analysis.Organics_desc, &anl_water_analysis.Ph_observed, &anl_water_analysis.Ph_temp, &anl_water_analysis.Ph_temp_ouom, &anl_water_analysis.Potassium, &anl_water_analysis.Potassium_ouom, &anl_water_analysis.Ppdm_guid, &anl_water_analysis.Preferred_ind, &anl_water_analysis.Problem_ind, &anl_water_analysis.Recovery_desc, &anl_water_analysis.Refractive_index, &anl_water_analysis.Refractive_index_temp, &anl_water_analysis.Refractive_index_temp_ouom, &anl_water_analysis.Remark, &anl_water_analysis.Reservoir_temperature, &anl_water_analysis.Reservoir_temperature_ouom, &anl_water_analysis.Rw_temperature, &anl_water_analysis.Rw_temperature_ouom, &anl_water_analysis.Sample_water_type, &anl_water_analysis.Sodium, &anl_water_analysis.Sodium_ouom, &anl_water_analysis.Sodium_potassium, &anl_water_analysis.Sodium_potassium_ouom, &anl_water_analysis.Source, &anl_water_analysis.Step_seq_no, &anl_water_analysis.Strontium, &anl_water_analysis.Strontium_ouom, &anl_water_analysis.Sulphate, &anl_water_analysis.Sulphate_ouom, &anl_water_analysis.Water_density, &anl_water_analysis.Water_density_ouom, &anl_water_analysis.Water_ph, &anl_water_analysis.Water_resistivity, &anl_water_analysis.Water_resistivity_ouom, &anl_water_analysis.Water_type, &anl_water_analysis.Row_changed_by, &anl_water_analysis.Row_changed_date, &anl_water_analysis.Row_created_by, &anl_water_analysis.Row_created_date, &anl_water_analysis.Row_effective_date, &anl_water_analysis.Row_expiry_date, &anl_water_analysis.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_water_analysis to result
		result = append(result, anl_water_analysis)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlWaterAnalysis(c *fiber.Ctx) error {
	var anl_water_analysis dto.Anl_water_analysis

	setDefaults(&anl_water_analysis)

	if err := json.Unmarshal(c.Body(), &anl_water_analysis); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_water_analysis values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76)")
	if err != nil {
		return err
	}
	anl_water_analysis.Row_created_date = formatDate(anl_water_analysis.Row_created_date)
	anl_water_analysis.Row_changed_date = formatDate(anl_water_analysis.Row_changed_date)
	anl_water_analysis.Effective_date = formatDateString(anl_water_analysis.Effective_date)
	anl_water_analysis.Expiry_date = formatDateString(anl_water_analysis.Expiry_date)
	anl_water_analysis.Row_effective_date = formatDateString(anl_water_analysis.Row_effective_date)
	anl_water_analysis.Row_expiry_date = formatDateString(anl_water_analysis.Row_expiry_date)






	rows, err := stmt.Exec(anl_water_analysis.Analysis_id, anl_water_analysis.Anl_source, anl_water_analysis.Water_analysis_obs_no, anl_water_analysis.Active_ind, anl_water_analysis.Barium, anl_water_analysis.Barium_ouom, anl_water_analysis.Bicarbonate, anl_water_analysis.Bicarbonate_ouom, anl_water_analysis.Boron, anl_water_analysis.Boron_ouom, anl_water_analysis.Bromine, anl_water_analysis.Bromine_ouom, anl_water_analysis.Calcium, anl_water_analysis.Calcium_ouom, anl_water_analysis.Carbonate, anl_water_analysis.Carbonate_ouom, anl_water_analysis.Chlorine, anl_water_analysis.Chlorine_ouom, anl_water_analysis.Conductivity, anl_water_analysis.Conductivity_ouom, anl_water_analysis.Effective_date, anl_water_analysis.Expiry_date, anl_water_analysis.Hydrogen_sulphide, anl_water_analysis.Hydrogen_sulphide_ouom, anl_water_analysis.Hydroxide, anl_water_analysis.Hydroxide_ouom, anl_water_analysis.Iodine, anl_water_analysis.Iodine_ouom, anl_water_analysis.Ion_ouom, anl_water_analysis.Ion_uom, anl_water_analysis.Iron, anl_water_analysis.Iron_ouom, anl_water_analysis.Magnesium, anl_water_analysis.Magnesium_ouom, anl_water_analysis.Organics_desc, anl_water_analysis.Ph_observed, anl_water_analysis.Ph_temp, anl_water_analysis.Ph_temp_ouom, anl_water_analysis.Potassium, anl_water_analysis.Potassium_ouom, anl_water_analysis.Ppdm_guid, anl_water_analysis.Preferred_ind, anl_water_analysis.Problem_ind, anl_water_analysis.Recovery_desc, anl_water_analysis.Refractive_index, anl_water_analysis.Refractive_index_temp, anl_water_analysis.Refractive_index_temp_ouom, anl_water_analysis.Remark, anl_water_analysis.Reservoir_temperature, anl_water_analysis.Reservoir_temperature_ouom, anl_water_analysis.Rw_temperature, anl_water_analysis.Rw_temperature_ouom, anl_water_analysis.Sample_water_type, anl_water_analysis.Sodium, anl_water_analysis.Sodium_ouom, anl_water_analysis.Sodium_potassium, anl_water_analysis.Sodium_potassium_ouom, anl_water_analysis.Source, anl_water_analysis.Step_seq_no, anl_water_analysis.Strontium, anl_water_analysis.Strontium_ouom, anl_water_analysis.Sulphate, anl_water_analysis.Sulphate_ouom, anl_water_analysis.Water_density, anl_water_analysis.Water_density_ouom, anl_water_analysis.Water_ph, anl_water_analysis.Water_resistivity, anl_water_analysis.Water_resistivity_ouom, anl_water_analysis.Water_type, anl_water_analysis.Row_changed_by, anl_water_analysis.Row_changed_date, anl_water_analysis.Row_created_by, anl_water_analysis.Row_created_date, anl_water_analysis.Row_effective_date, anl_water_analysis.Row_expiry_date, anl_water_analysis.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlWaterAnalysis(c *fiber.Ctx) error {
	var anl_water_analysis dto.Anl_water_analysis

	setDefaults(&anl_water_analysis)

	if err := json.Unmarshal(c.Body(), &anl_water_analysis); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_water_analysis.Analysis_id = id

    if anl_water_analysis.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_water_analysis set anl_source = :1, water_analysis_obs_no = :2, active_ind = :3, barium = :4, barium_ouom = :5, bicarbonate = :6, bicarbonate_ouom = :7, boron = :8, boron_ouom = :9, bromine = :10, bromine_ouom = :11, calcium = :12, calcium_ouom = :13, carbonate = :14, carbonate_ouom = :15, chlorine = :16, chlorine_ouom = :17, conductivity = :18, conductivity_ouom = :19, effective_date = :20, expiry_date = :21, hydrogen_sulphide = :22, hydrogen_sulphide_ouom = :23, hydroxide = :24, hydroxide_ouom = :25, iodine = :26, iodine_ouom = :27, ion_ouom = :28, ion_uom = :29, iron = :30, iron_ouom = :31, magnesium = :32, magnesium_ouom = :33, organics_desc = :34, ph_observed = :35, ph_temp = :36, ph_temp_ouom = :37, potassium = :38, potassium_ouom = :39, ppdm_guid = :40, preferred_ind = :41, problem_ind = :42, recovery_desc = :43, refractive_index = :44, refractive_index_temp = :45, refractive_index_temp_ouom = :46, remark = :47, reservoir_temperature = :48, reservoir_temperature_ouom = :49, rw_temperature = :50, rw_temperature_ouom = :51, sample_water_type = :52, sodium = :53, sodium_ouom = :54, sodium_potassium = :55, sodium_potassium_ouom = :56, source = :57, step_seq_no = :58, strontium = :59, strontium_ouom = :60, sulphate = :61, sulphate_ouom = :62, water_density = :63, water_density_ouom = :64, water_ph = :65, water_resistivity = :66, water_resistivity_ouom = :67, water_type = :68, row_changed_by = :69, row_changed_date = :70, row_created_by = :71, row_effective_date = :72, row_expiry_date = :73, row_quality = :74 where analysis_id = :76")
	if err != nil {
		return err
	}

	anl_water_analysis.Row_changed_date = formatDate(anl_water_analysis.Row_changed_date)
	anl_water_analysis.Effective_date = formatDateString(anl_water_analysis.Effective_date)
	anl_water_analysis.Expiry_date = formatDateString(anl_water_analysis.Expiry_date)
	anl_water_analysis.Row_effective_date = formatDateString(anl_water_analysis.Row_effective_date)
	anl_water_analysis.Row_expiry_date = formatDateString(anl_water_analysis.Row_expiry_date)






	rows, err := stmt.Exec(anl_water_analysis.Anl_source, anl_water_analysis.Water_analysis_obs_no, anl_water_analysis.Active_ind, anl_water_analysis.Barium, anl_water_analysis.Barium_ouom, anl_water_analysis.Bicarbonate, anl_water_analysis.Bicarbonate_ouom, anl_water_analysis.Boron, anl_water_analysis.Boron_ouom, anl_water_analysis.Bromine, anl_water_analysis.Bromine_ouom, anl_water_analysis.Calcium, anl_water_analysis.Calcium_ouom, anl_water_analysis.Carbonate, anl_water_analysis.Carbonate_ouom, anl_water_analysis.Chlorine, anl_water_analysis.Chlorine_ouom, anl_water_analysis.Conductivity, anl_water_analysis.Conductivity_ouom, anl_water_analysis.Effective_date, anl_water_analysis.Expiry_date, anl_water_analysis.Hydrogen_sulphide, anl_water_analysis.Hydrogen_sulphide_ouom, anl_water_analysis.Hydroxide, anl_water_analysis.Hydroxide_ouom, anl_water_analysis.Iodine, anl_water_analysis.Iodine_ouom, anl_water_analysis.Ion_ouom, anl_water_analysis.Ion_uom, anl_water_analysis.Iron, anl_water_analysis.Iron_ouom, anl_water_analysis.Magnesium, anl_water_analysis.Magnesium_ouom, anl_water_analysis.Organics_desc, anl_water_analysis.Ph_observed, anl_water_analysis.Ph_temp, anl_water_analysis.Ph_temp_ouom, anl_water_analysis.Potassium, anl_water_analysis.Potassium_ouom, anl_water_analysis.Ppdm_guid, anl_water_analysis.Preferred_ind, anl_water_analysis.Problem_ind, anl_water_analysis.Recovery_desc, anl_water_analysis.Refractive_index, anl_water_analysis.Refractive_index_temp, anl_water_analysis.Refractive_index_temp_ouom, anl_water_analysis.Remark, anl_water_analysis.Reservoir_temperature, anl_water_analysis.Reservoir_temperature_ouom, anl_water_analysis.Rw_temperature, anl_water_analysis.Rw_temperature_ouom, anl_water_analysis.Sample_water_type, anl_water_analysis.Sodium, anl_water_analysis.Sodium_ouom, anl_water_analysis.Sodium_potassium, anl_water_analysis.Sodium_potassium_ouom, anl_water_analysis.Source, anl_water_analysis.Step_seq_no, anl_water_analysis.Strontium, anl_water_analysis.Strontium_ouom, anl_water_analysis.Sulphate, anl_water_analysis.Sulphate_ouom, anl_water_analysis.Water_density, anl_water_analysis.Water_density_ouom, anl_water_analysis.Water_ph, anl_water_analysis.Water_resistivity, anl_water_analysis.Water_resistivity_ouom, anl_water_analysis.Water_type, anl_water_analysis.Row_changed_by, anl_water_analysis.Row_changed_date, anl_water_analysis.Row_created_by, anl_water_analysis.Row_effective_date, anl_water_analysis.Row_expiry_date, anl_water_analysis.Row_quality, anl_water_analysis.Analysis_id)
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

func PatchAnlWaterAnalysis(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_water_analysis set "
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
	query += " where analysis_id = :id"

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

func DeleteAnlWaterAnalysis(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_water_analysis dto.Anl_water_analysis
	anl_water_analysis.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_water_analysis where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_water_analysis.Analysis_id)
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



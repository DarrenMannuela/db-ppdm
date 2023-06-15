package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlOilAnalysis(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_oil_analysis")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_oil_analysis

	for rows.Next() {
		var anl_oil_analysis dto.Anl_oil_analysis
		if err := rows.Scan(&anl_oil_analysis.Analysis_id, &anl_oil_analysis.Anl_source, &anl_oil_analysis.Oil_analysis_obs_no, &anl_oil_analysis.Absolute_gravity, &anl_oil_analysis.Absolute_gravity_ouom, &anl_oil_analysis.Active_ind, &anl_oil_analysis.Asphaltene_content, &anl_oil_analysis.Asphaltene_content_ouom, &anl_oil_analysis.Asph_content_preferred_ind, &anl_oil_analysis.Asph_content_problem_ind, &anl_oil_analysis.Asph_content_step_seq_no, &anl_oil_analysis.Characterize_factor, &anl_oil_analysis.Characterize_factor_ouom, &anl_oil_analysis.Cloud_point_preferred_ind, &anl_oil_analysis.Cloud_point_problem_ind, &anl_oil_analysis.Cloud_point_step_seq_no, &anl_oil_analysis.Cloud_point_temp, &anl_oil_analysis.Cloud_point_temp_ouom, &anl_oil_analysis.Dew_point_preferred_ind, &anl_oil_analysis.Dew_point_problem_ind, &anl_oil_analysis.Dew_point_step_seq_no, &anl_oil_analysis.Dew_point_temp, &anl_oil_analysis.Dew_point_temp_ouom, &anl_oil_analysis.Distillation_base_type, &anl_oil_analysis.Distillation_preferred_ind, &anl_oil_analysis.Distillation_problem_ind, &anl_oil_analysis.Distillation_step_seq_no, &anl_oil_analysis.Effective_date, &anl_oil_analysis.Expiry_date, &anl_oil_analysis.Flash_point_preferred_ind, &anl_oil_analysis.Flash_point_problem_ind, &anl_oil_analysis.Flash_point_step_seq_no, &anl_oil_analysis.Flash_point_temp, &anl_oil_analysis.Flash_point_temp_ouom, &anl_oil_analysis.Gor, &anl_oil_analysis.Gor_ouom, &anl_oil_analysis.Measured_pressure, &anl_oil_analysis.Measured_pressure_ouom, &anl_oil_analysis.Measured_volume, &anl_oil_analysis.Measured_volume_ouom, &anl_oil_analysis.Method_type, &anl_oil_analysis.Oil_analysis_temp, &anl_oil_analysis.Oil_analysis_temp_ouom, &anl_oil_analysis.Oil_color, &anl_oil_analysis.Oil_color_preferred_ind, &anl_oil_analysis.Oil_color_problem_ind, &anl_oil_analysis.Oil_color_step_seq_no, &anl_oil_analysis.Oil_density, &anl_oil_analysis.Oil_density_ouom, &anl_oil_analysis.Oil_density_preferred_ind, &anl_oil_analysis.Oil_density_problem_ind, &anl_oil_analysis.Oil_density_step_seq_no, &anl_oil_analysis.Oil_gravity, &anl_oil_analysis.Oil_gravity_ouom, &anl_oil_analysis.Oil_gravity_preferred_ind, &anl_oil_analysis.Oil_gravity_step_seq_no, &anl_oil_analysis.Oil_gravity_temp, &anl_oil_analysis.Oil_gravity_temp_ouom, &anl_oil_analysis.Oil_gravity_temp_prefer_ind, &anl_oil_analysis.Oil_gravity_temp_prob_ind, &anl_oil_analysis.Oil_gravity_temp_step_seq, &anl_oil_analysis.Oil_type, &anl_oil_analysis.Ppdm_guid, &anl_oil_analysis.Relative_density, &anl_oil_analysis.Relative_density_ouom, &anl_oil_analysis.Relative_std_density, &anl_oil_analysis.Relative_std_density_ouom, &anl_oil_analysis.Remark, &anl_oil_analysis.Reservoir_press, &anl_oil_analysis.Reservoir_press_ouom, &anl_oil_analysis.Reservoir_temp, &anl_oil_analysis.Reservoir_temp_ouom, &anl_oil_analysis.Sediment_content, &anl_oil_analysis.Sediment_content_ouom, &anl_oil_analysis.Sediment_preferred_ind, &anl_oil_analysis.Sediment_problem_ind, &anl_oil_analysis.Sediment_step_seq_no, &anl_oil_analysis.Shrinkage_factor, &anl_oil_analysis.Shrink_factor_preferred_ind, &anl_oil_analysis.Shrink_factor_problem_ind, &anl_oil_analysis.Shrink_factor_step_seq_no, &anl_oil_analysis.Source, &anl_oil_analysis.Sulphur_content, &anl_oil_analysis.Vapour_press, &anl_oil_analysis.Vapour_press_ouom, &anl_oil_analysis.Vapour_press_temp, &anl_oil_analysis.Vapour_press_temp_ouom, &anl_oil_analysis.Water_content, &anl_oil_analysis.Water_content_ouom, &anl_oil_analysis.Wax_content, &anl_oil_analysis.Wax_content_ouom, &anl_oil_analysis.Wax_content_preferred_ind, &anl_oil_analysis.Wax_content_problem_ind, &anl_oil_analysis.Wax_content_step_seq_no, &anl_oil_analysis.Row_changed_by, &anl_oil_analysis.Row_changed_date, &anl_oil_analysis.Row_created_by, &anl_oil_analysis.Row_created_date, &anl_oil_analysis.Row_effective_date, &anl_oil_analysis.Row_expiry_date, &anl_oil_analysis.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_oil_analysis to result
		result = append(result, anl_oil_analysis)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlOilAnalysis(c *fiber.Ctx) error {
	var anl_oil_analysis dto.Anl_oil_analysis

	setDefaults(&anl_oil_analysis)

	if err := json.Unmarshal(c.Body(), &anl_oil_analysis); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_oil_analysis values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96, :97, :98, :99, :100, :101)")
	if err != nil {
		return err
	}
	anl_oil_analysis.Row_created_date = formatDate(anl_oil_analysis.Row_created_date)
	anl_oil_analysis.Row_changed_date = formatDate(anl_oil_analysis.Row_changed_date)
	anl_oil_analysis.Effective_date = formatDateString(anl_oil_analysis.Effective_date)
	anl_oil_analysis.Expiry_date = formatDateString(anl_oil_analysis.Expiry_date)
	anl_oil_analysis.Row_effective_date = formatDateString(anl_oil_analysis.Row_effective_date)
	anl_oil_analysis.Row_expiry_date = formatDateString(anl_oil_analysis.Row_expiry_date)






	rows, err := stmt.Exec(anl_oil_analysis.Analysis_id, anl_oil_analysis.Anl_source, anl_oil_analysis.Oil_analysis_obs_no, anl_oil_analysis.Absolute_gravity, anl_oil_analysis.Absolute_gravity_ouom, anl_oil_analysis.Active_ind, anl_oil_analysis.Asphaltene_content, anl_oil_analysis.Asphaltene_content_ouom, anl_oil_analysis.Asph_content_preferred_ind, anl_oil_analysis.Asph_content_problem_ind, anl_oil_analysis.Asph_content_step_seq_no, anl_oil_analysis.Characterize_factor, anl_oil_analysis.Characterize_factor_ouom, anl_oil_analysis.Cloud_point_preferred_ind, anl_oil_analysis.Cloud_point_problem_ind, anl_oil_analysis.Cloud_point_step_seq_no, anl_oil_analysis.Cloud_point_temp, anl_oil_analysis.Cloud_point_temp_ouom, anl_oil_analysis.Dew_point_preferred_ind, anl_oil_analysis.Dew_point_problem_ind, anl_oil_analysis.Dew_point_step_seq_no, anl_oil_analysis.Dew_point_temp, anl_oil_analysis.Dew_point_temp_ouom, anl_oil_analysis.Distillation_base_type, anl_oil_analysis.Distillation_preferred_ind, anl_oil_analysis.Distillation_problem_ind, anl_oil_analysis.Distillation_step_seq_no, anl_oil_analysis.Effective_date, anl_oil_analysis.Expiry_date, anl_oil_analysis.Flash_point_preferred_ind, anl_oil_analysis.Flash_point_problem_ind, anl_oil_analysis.Flash_point_step_seq_no, anl_oil_analysis.Flash_point_temp, anl_oil_analysis.Flash_point_temp_ouom, anl_oil_analysis.Gor, anl_oil_analysis.Gor_ouom, anl_oil_analysis.Measured_pressure, anl_oil_analysis.Measured_pressure_ouom, anl_oil_analysis.Measured_volume, anl_oil_analysis.Measured_volume_ouom, anl_oil_analysis.Method_type, anl_oil_analysis.Oil_analysis_temp, anl_oil_analysis.Oil_analysis_temp_ouom, anl_oil_analysis.Oil_color, anl_oil_analysis.Oil_color_preferred_ind, anl_oil_analysis.Oil_color_problem_ind, anl_oil_analysis.Oil_color_step_seq_no, anl_oil_analysis.Oil_density, anl_oil_analysis.Oil_density_ouom, anl_oil_analysis.Oil_density_preferred_ind, anl_oil_analysis.Oil_density_problem_ind, anl_oil_analysis.Oil_density_step_seq_no, anl_oil_analysis.Oil_gravity, anl_oil_analysis.Oil_gravity_ouom, anl_oil_analysis.Oil_gravity_preferred_ind, anl_oil_analysis.Oil_gravity_step_seq_no, anl_oil_analysis.Oil_gravity_temp, anl_oil_analysis.Oil_gravity_temp_ouom, anl_oil_analysis.Oil_gravity_temp_prefer_ind, anl_oil_analysis.Oil_gravity_temp_prob_ind, anl_oil_analysis.Oil_gravity_temp_step_seq, anl_oil_analysis.Oil_type, anl_oil_analysis.Ppdm_guid, anl_oil_analysis.Relative_density, anl_oil_analysis.Relative_density_ouom, anl_oil_analysis.Relative_std_density, anl_oil_analysis.Relative_std_density_ouom, anl_oil_analysis.Remark, anl_oil_analysis.Reservoir_press, anl_oil_analysis.Reservoir_press_ouom, anl_oil_analysis.Reservoir_temp, anl_oil_analysis.Reservoir_temp_ouom, anl_oil_analysis.Sediment_content, anl_oil_analysis.Sediment_content_ouom, anl_oil_analysis.Sediment_preferred_ind, anl_oil_analysis.Sediment_problem_ind, anl_oil_analysis.Sediment_step_seq_no, anl_oil_analysis.Shrinkage_factor, anl_oil_analysis.Shrink_factor_preferred_ind, anl_oil_analysis.Shrink_factor_problem_ind, anl_oil_analysis.Shrink_factor_step_seq_no, anl_oil_analysis.Source, anl_oil_analysis.Sulphur_content, anl_oil_analysis.Vapour_press, anl_oil_analysis.Vapour_press_ouom, anl_oil_analysis.Vapour_press_temp, anl_oil_analysis.Vapour_press_temp_ouom, anl_oil_analysis.Water_content, anl_oil_analysis.Water_content_ouom, anl_oil_analysis.Wax_content, anl_oil_analysis.Wax_content_ouom, anl_oil_analysis.Wax_content_preferred_ind, anl_oil_analysis.Wax_content_problem_ind, anl_oil_analysis.Wax_content_step_seq_no, anl_oil_analysis.Row_changed_by, anl_oil_analysis.Row_changed_date, anl_oil_analysis.Row_created_by, anl_oil_analysis.Row_created_date, anl_oil_analysis.Row_effective_date, anl_oil_analysis.Row_expiry_date, anl_oil_analysis.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlOilAnalysis(c *fiber.Ctx) error {
	var anl_oil_analysis dto.Anl_oil_analysis

	setDefaults(&anl_oil_analysis)

	if err := json.Unmarshal(c.Body(), &anl_oil_analysis); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_oil_analysis.Analysis_id = id

    if anl_oil_analysis.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_oil_analysis set anl_source = :1, oil_analysis_obs_no = :2, absolute_gravity = :3, absolute_gravity_ouom = :4, active_ind = :5, asphaltene_content = :6, asphaltene_content_ouom = :7, asph_content_preferred_ind = :8, asph_content_problem_ind = :9, asph_content_step_seq_no = :10, characterize_factor = :11, characterize_factor_ouom = :12, cloud_point_preferred_ind = :13, cloud_point_problem_ind = :14, cloud_point_step_seq_no = :15, cloud_point_temp = :16, cloud_point_temp_ouom = :17, dew_point_preferred_ind = :18, dew_point_problem_ind = :19, dew_point_step_seq_no = :20, dew_point_temp = :21, dew_point_temp_ouom = :22, distillation_base_type = :23, distillation_preferred_ind = :24, distillation_problem_ind = :25, distillation_step_seq_no = :26, effective_date = :27, expiry_date = :28, flash_point_preferred_ind = :29, flash_point_problem_ind = :30, flash_point_step_seq_no = :31, flash_point_temp = :32, flash_point_temp_ouom = :33, gor = :34, gor_ouom = :35, measured_pressure = :36, measured_pressure_ouom = :37, measured_volume = :38, measured_volume_ouom = :39, method_type = :40, oil_analysis_temp = :41, oil_analysis_temp_ouom = :42, oil_color = :43, oil_color_preferred_ind = :44, oil_color_problem_ind = :45, oil_color_step_seq_no = :46, oil_density = :47, oil_density_ouom = :48, oil_density_preferred_ind = :49, oil_density_problem_ind = :50, oil_density_step_seq_no = :51, oil_gravity = :52, oil_gravity_ouom = :53, oil_gravity_preferred_ind = :54, oil_gravity_step_seq_no = :55, oil_gravity_temp = :56, oil_gravity_temp_ouom = :57, oil_gravity_temp_prefer_ind = :58, oil_gravity_temp_prob_ind = :59, oil_gravity_temp_step_seq = :60, oil_type = :61, ppdm_guid = :62, relative_density = :63, relative_density_ouom = :64, relative_std_density = :65, relative_std_density_ouom = :66, remark = :67, reservoir_press = :68, reservoir_press_ouom = :69, reservoir_temp = :70, reservoir_temp_ouom = :71, sediment_content = :72, sediment_content_ouom = :73, sediment_preferred_ind = :74, sediment_problem_ind = :75, sediment_step_seq_no = :76, shrinkage_factor = :77, shrink_factor_preferred_ind = :78, shrink_factor_problem_ind = :79, shrink_factor_step_seq_no = :80, source = :81, sulphur_content = :82, vapour_press = :83, vapour_press_ouom = :84, vapour_press_temp = :85, vapour_press_temp_ouom = :86, water_content = :87, water_content_ouom = :88, wax_content = :89, wax_content_ouom = :90, wax_content_preferred_ind = :91, wax_content_problem_ind = :92, wax_content_step_seq_no = :93, row_changed_by = :94, row_changed_date = :95, row_created_by = :96, row_effective_date = :97, row_expiry_date = :98, row_quality = :99 where analysis_id = :101")
	if err != nil {
		return err
	}

	anl_oil_analysis.Row_changed_date = formatDate(anl_oil_analysis.Row_changed_date)
	anl_oil_analysis.Effective_date = formatDateString(anl_oil_analysis.Effective_date)
	anl_oil_analysis.Expiry_date = formatDateString(anl_oil_analysis.Expiry_date)
	anl_oil_analysis.Row_effective_date = formatDateString(anl_oil_analysis.Row_effective_date)
	anl_oil_analysis.Row_expiry_date = formatDateString(anl_oil_analysis.Row_expiry_date)






	rows, err := stmt.Exec(anl_oil_analysis.Anl_source, anl_oil_analysis.Oil_analysis_obs_no, anl_oil_analysis.Absolute_gravity, anl_oil_analysis.Absolute_gravity_ouom, anl_oil_analysis.Active_ind, anl_oil_analysis.Asphaltene_content, anl_oil_analysis.Asphaltene_content_ouom, anl_oil_analysis.Asph_content_preferred_ind, anl_oil_analysis.Asph_content_problem_ind, anl_oil_analysis.Asph_content_step_seq_no, anl_oil_analysis.Characterize_factor, anl_oil_analysis.Characterize_factor_ouom, anl_oil_analysis.Cloud_point_preferred_ind, anl_oil_analysis.Cloud_point_problem_ind, anl_oil_analysis.Cloud_point_step_seq_no, anl_oil_analysis.Cloud_point_temp, anl_oil_analysis.Cloud_point_temp_ouom, anl_oil_analysis.Dew_point_preferred_ind, anl_oil_analysis.Dew_point_problem_ind, anl_oil_analysis.Dew_point_step_seq_no, anl_oil_analysis.Dew_point_temp, anl_oil_analysis.Dew_point_temp_ouom, anl_oil_analysis.Distillation_base_type, anl_oil_analysis.Distillation_preferred_ind, anl_oil_analysis.Distillation_problem_ind, anl_oil_analysis.Distillation_step_seq_no, anl_oil_analysis.Effective_date, anl_oil_analysis.Expiry_date, anl_oil_analysis.Flash_point_preferred_ind, anl_oil_analysis.Flash_point_problem_ind, anl_oil_analysis.Flash_point_step_seq_no, anl_oil_analysis.Flash_point_temp, anl_oil_analysis.Flash_point_temp_ouom, anl_oil_analysis.Gor, anl_oil_analysis.Gor_ouom, anl_oil_analysis.Measured_pressure, anl_oil_analysis.Measured_pressure_ouom, anl_oil_analysis.Measured_volume, anl_oil_analysis.Measured_volume_ouom, anl_oil_analysis.Method_type, anl_oil_analysis.Oil_analysis_temp, anl_oil_analysis.Oil_analysis_temp_ouom, anl_oil_analysis.Oil_color, anl_oil_analysis.Oil_color_preferred_ind, anl_oil_analysis.Oil_color_problem_ind, anl_oil_analysis.Oil_color_step_seq_no, anl_oil_analysis.Oil_density, anl_oil_analysis.Oil_density_ouom, anl_oil_analysis.Oil_density_preferred_ind, anl_oil_analysis.Oil_density_problem_ind, anl_oil_analysis.Oil_density_step_seq_no, anl_oil_analysis.Oil_gravity, anl_oil_analysis.Oil_gravity_ouom, anl_oil_analysis.Oil_gravity_preferred_ind, anl_oil_analysis.Oil_gravity_step_seq_no, anl_oil_analysis.Oil_gravity_temp, anl_oil_analysis.Oil_gravity_temp_ouom, anl_oil_analysis.Oil_gravity_temp_prefer_ind, anl_oil_analysis.Oil_gravity_temp_prob_ind, anl_oil_analysis.Oil_gravity_temp_step_seq, anl_oil_analysis.Oil_type, anl_oil_analysis.Ppdm_guid, anl_oil_analysis.Relative_density, anl_oil_analysis.Relative_density_ouom, anl_oil_analysis.Relative_std_density, anl_oil_analysis.Relative_std_density_ouom, anl_oil_analysis.Remark, anl_oil_analysis.Reservoir_press, anl_oil_analysis.Reservoir_press_ouom, anl_oil_analysis.Reservoir_temp, anl_oil_analysis.Reservoir_temp_ouom, anl_oil_analysis.Sediment_content, anl_oil_analysis.Sediment_content_ouom, anl_oil_analysis.Sediment_preferred_ind, anl_oil_analysis.Sediment_problem_ind, anl_oil_analysis.Sediment_step_seq_no, anl_oil_analysis.Shrinkage_factor, anl_oil_analysis.Shrink_factor_preferred_ind, anl_oil_analysis.Shrink_factor_problem_ind, anl_oil_analysis.Shrink_factor_step_seq_no, anl_oil_analysis.Source, anl_oil_analysis.Sulphur_content, anl_oil_analysis.Vapour_press, anl_oil_analysis.Vapour_press_ouom, anl_oil_analysis.Vapour_press_temp, anl_oil_analysis.Vapour_press_temp_ouom, anl_oil_analysis.Water_content, anl_oil_analysis.Water_content_ouom, anl_oil_analysis.Wax_content, anl_oil_analysis.Wax_content_ouom, anl_oil_analysis.Wax_content_preferred_ind, anl_oil_analysis.Wax_content_problem_ind, anl_oil_analysis.Wax_content_step_seq_no, anl_oil_analysis.Row_changed_by, anl_oil_analysis.Row_changed_date, anl_oil_analysis.Row_created_by, anl_oil_analysis.Row_effective_date, anl_oil_analysis.Row_expiry_date, anl_oil_analysis.Row_quality, anl_oil_analysis.Analysis_id)
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

func PatchAnlOilAnalysis(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_oil_analysis set "
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

func DeleteAnlOilAnalysis(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_oil_analysis dto.Anl_oil_analysis
	anl_oil_analysis.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_oil_analysis where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_oil_analysis.Analysis_id)
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



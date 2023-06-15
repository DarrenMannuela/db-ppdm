package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSampleOrigin(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sample_origin")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sample_origin

	for rows.Next() {
		var sample_origin dto.Sample_origin
		if err := rows.Scan(&sample_origin.Sample_id, &sample_origin.Collection_obs_no, &sample_origin.Active_ind, &sample_origin.Analysis_id, &sample_origin.Analysis_source, &sample_origin.Anl_step_seq_no, &sample_origin.Area_id, &sample_origin.Area_type, &sample_origin.Collected_by_ba_id, &sample_origin.Collected_for_ba_id, &sample_origin.Collect_method, &sample_origin.Core_type, &sample_origin.Effective_date, &sample_origin.Expiry_date, &sample_origin.Field_station_id, &sample_origin.Land_right_id, &sample_origin.Land_right_subtype, &sample_origin.Meas_section_id, &sample_origin.Meas_section_source, &sample_origin.Portion_diameter, &sample_origin.Portion_diameter_ouom, &sample_origin.Portion_length, &sample_origin.Portion_length_ouom, &sample_origin.Portion_volume, &sample_origin.Portion_volume_ouom, &sample_origin.Portion_weight, &sample_origin.Portion_weight_ouom, &sample_origin.Ppdm_guid, &sample_origin.Prod_string_id, &sample_origin.Prod_string_source, &sample_origin.Project_id, &sample_origin.Remark, &sample_origin.Sample_part_id, &sample_origin.Source, &sample_origin.Well_core_id, &sample_origin.Well_source, &sample_origin.Well_survey_id, &sample_origin.Well_survey_source, &sample_origin.Well_uwi, &sample_origin.Row_changed_by, &sample_origin.Row_changed_date, &sample_origin.Row_created_by, &sample_origin.Row_created_date, &sample_origin.Row_effective_date, &sample_origin.Row_expiry_date, &sample_origin.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sample_origin to result
		result = append(result, sample_origin)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSampleOrigin(c *fiber.Ctx) error {
	var sample_origin dto.Sample_origin

	setDefaults(&sample_origin)

	if err := json.Unmarshal(c.Body(), &sample_origin); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sample_origin values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46)")
	if err != nil {
		return err
	}
	sample_origin.Row_created_date = formatDate(sample_origin.Row_created_date)
	sample_origin.Row_changed_date = formatDate(sample_origin.Row_changed_date)
	sample_origin.Effective_date = formatDateString(sample_origin.Effective_date)
	sample_origin.Expiry_date = formatDateString(sample_origin.Expiry_date)
	sample_origin.Row_effective_date = formatDateString(sample_origin.Row_effective_date)
	sample_origin.Row_expiry_date = formatDateString(sample_origin.Row_expiry_date)






	rows, err := stmt.Exec(sample_origin.Sample_id, sample_origin.Collection_obs_no, sample_origin.Active_ind, sample_origin.Analysis_id, sample_origin.Analysis_source, sample_origin.Anl_step_seq_no, sample_origin.Area_id, sample_origin.Area_type, sample_origin.Collected_by_ba_id, sample_origin.Collected_for_ba_id, sample_origin.Collect_method, sample_origin.Core_type, sample_origin.Effective_date, sample_origin.Expiry_date, sample_origin.Field_station_id, sample_origin.Land_right_id, sample_origin.Land_right_subtype, sample_origin.Meas_section_id, sample_origin.Meas_section_source, sample_origin.Portion_diameter, sample_origin.Portion_diameter_ouom, sample_origin.Portion_length, sample_origin.Portion_length_ouom, sample_origin.Portion_volume, sample_origin.Portion_volume_ouom, sample_origin.Portion_weight, sample_origin.Portion_weight_ouom, sample_origin.Ppdm_guid, sample_origin.Prod_string_id, sample_origin.Prod_string_source, sample_origin.Project_id, sample_origin.Remark, sample_origin.Sample_part_id, sample_origin.Source, sample_origin.Well_core_id, sample_origin.Well_source, sample_origin.Well_survey_id, sample_origin.Well_survey_source, sample_origin.Well_uwi, sample_origin.Row_changed_by, sample_origin.Row_changed_date, sample_origin.Row_created_by, sample_origin.Row_created_date, sample_origin.Row_effective_date, sample_origin.Row_expiry_date, sample_origin.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSampleOrigin(c *fiber.Ctx) error {
	var sample_origin dto.Sample_origin

	setDefaults(&sample_origin)

	if err := json.Unmarshal(c.Body(), &sample_origin); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sample_origin.Sample_id = id

    if sample_origin.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sample_origin set collection_obs_no = :1, active_ind = :2, analysis_id = :3, analysis_source = :4, anl_step_seq_no = :5, area_id = :6, area_type = :7, collected_by_ba_id = :8, collected_for_ba_id = :9, collect_method = :10, core_type = :11, effective_date = :12, expiry_date = :13, field_station_id = :14, land_right_id = :15, land_right_subtype = :16, meas_section_id = :17, meas_section_source = :18, portion_diameter = :19, portion_diameter_ouom = :20, portion_length = :21, portion_length_ouom = :22, portion_volume = :23, portion_volume_ouom = :24, portion_weight = :25, portion_weight_ouom = :26, ppdm_guid = :27, prod_string_id = :28, prod_string_source = :29, project_id = :30, remark = :31, sample_part_id = :32, source = :33, well_core_id = :34, well_source = :35, well_survey_id = :36, well_survey_source = :37, well_uwi = :38, row_changed_by = :39, row_changed_date = :40, row_created_by = :41, row_effective_date = :42, row_expiry_date = :43, row_quality = :44 where sample_id = :46")
	if err != nil {
		return err
	}

	sample_origin.Row_changed_date = formatDate(sample_origin.Row_changed_date)
	sample_origin.Effective_date = formatDateString(sample_origin.Effective_date)
	sample_origin.Expiry_date = formatDateString(sample_origin.Expiry_date)
	sample_origin.Row_effective_date = formatDateString(sample_origin.Row_effective_date)
	sample_origin.Row_expiry_date = formatDateString(sample_origin.Row_expiry_date)






	rows, err := stmt.Exec(sample_origin.Collection_obs_no, sample_origin.Active_ind, sample_origin.Analysis_id, sample_origin.Analysis_source, sample_origin.Anl_step_seq_no, sample_origin.Area_id, sample_origin.Area_type, sample_origin.Collected_by_ba_id, sample_origin.Collected_for_ba_id, sample_origin.Collect_method, sample_origin.Core_type, sample_origin.Effective_date, sample_origin.Expiry_date, sample_origin.Field_station_id, sample_origin.Land_right_id, sample_origin.Land_right_subtype, sample_origin.Meas_section_id, sample_origin.Meas_section_source, sample_origin.Portion_diameter, sample_origin.Portion_diameter_ouom, sample_origin.Portion_length, sample_origin.Portion_length_ouom, sample_origin.Portion_volume, sample_origin.Portion_volume_ouom, sample_origin.Portion_weight, sample_origin.Portion_weight_ouom, sample_origin.Ppdm_guid, sample_origin.Prod_string_id, sample_origin.Prod_string_source, sample_origin.Project_id, sample_origin.Remark, sample_origin.Sample_part_id, sample_origin.Source, sample_origin.Well_core_id, sample_origin.Well_source, sample_origin.Well_survey_id, sample_origin.Well_survey_source, sample_origin.Well_uwi, sample_origin.Row_changed_by, sample_origin.Row_changed_date, sample_origin.Row_created_by, sample_origin.Row_effective_date, sample_origin.Row_expiry_date, sample_origin.Row_quality, sample_origin.Sample_id)
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

func PatchSampleOrigin(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sample_origin set "
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
	query += " where sample_id = :id"

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

func DeleteSampleOrigin(c *fiber.Ctx) error {
	id := c.Params("id")
	var sample_origin dto.Sample_origin
	sample_origin.Sample_id = id

	stmt, err := db.Prepare("delete from sample_origin where sample_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sample_origin.Sample_id)
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



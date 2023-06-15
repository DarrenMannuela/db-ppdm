package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisSpSurvey(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_sp_survey")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_sp_survey

	for rows.Next() {
		var seis_sp_survey dto.Seis_sp_survey
		if err := rows.Scan(&seis_sp_survey.Seis_set_subtype, &seis_sp_survey.Seis_set_id, &seis_sp_survey.Seis_point_id, &seis_sp_survey.Seis_sp_survey_obs_no, &seis_sp_survey.Active_ind, &seis_sp_survey.Azimuth, &seis_sp_survey.Azimuth_ouom, &seis_sp_survey.Effective_date, &seis_sp_survey.Ew_direction, &seis_sp_survey.Ew_distance, &seis_sp_survey.Ew_distance_ouom, &seis_sp_survey.Ew_start_line, &seis_sp_survey.Expiry_date, &seis_sp_survey.Location_type, &seis_sp_survey.Monument_id, &seis_sp_survey.Monument_sf_subtype, &seis_sp_survey.North_type, &seis_sp_survey.Ns_direction, &seis_sp_survey.Ns_distance, &seis_sp_survey.Ns_distance_ouom, &seis_sp_survey.Ns_start_line, &seis_sp_survey.Orientation, &seis_sp_survey.Point_offset, &seis_sp_survey.Point_offset_ouom, &seis_sp_survey.Ppdm_guid, &seis_sp_survey.Reference_loc, &seis_sp_survey.Remark, &seis_sp_survey.Source, &seis_sp_survey.Surface_loc, &seis_sp_survey.Well_node_id, &seis_sp_survey.Row_changed_by, &seis_sp_survey.Row_changed_date, &seis_sp_survey.Row_created_by, &seis_sp_survey.Row_created_date, &seis_sp_survey.Row_effective_date, &seis_sp_survey.Row_expiry_date, &seis_sp_survey.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_sp_survey to result
		result = append(result, seis_sp_survey)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisSpSurvey(c *fiber.Ctx) error {
	var seis_sp_survey dto.Seis_sp_survey

	setDefaults(&seis_sp_survey)

	if err := json.Unmarshal(c.Body(), &seis_sp_survey); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_sp_survey values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37)")
	if err != nil {
		return err
	}
	seis_sp_survey.Row_created_date = formatDate(seis_sp_survey.Row_created_date)
	seis_sp_survey.Row_changed_date = formatDate(seis_sp_survey.Row_changed_date)
	seis_sp_survey.Effective_date = formatDateString(seis_sp_survey.Effective_date)
	seis_sp_survey.Expiry_date = formatDateString(seis_sp_survey.Expiry_date)
	seis_sp_survey.Row_effective_date = formatDateString(seis_sp_survey.Row_effective_date)
	seis_sp_survey.Row_expiry_date = formatDateString(seis_sp_survey.Row_expiry_date)






	rows, err := stmt.Exec(seis_sp_survey.Seis_set_subtype, seis_sp_survey.Seis_set_id, seis_sp_survey.Seis_point_id, seis_sp_survey.Seis_sp_survey_obs_no, seis_sp_survey.Active_ind, seis_sp_survey.Azimuth, seis_sp_survey.Azimuth_ouom, seis_sp_survey.Effective_date, seis_sp_survey.Ew_direction, seis_sp_survey.Ew_distance, seis_sp_survey.Ew_distance_ouom, seis_sp_survey.Ew_start_line, seis_sp_survey.Expiry_date, seis_sp_survey.Location_type, seis_sp_survey.Monument_id, seis_sp_survey.Monument_sf_subtype, seis_sp_survey.North_type, seis_sp_survey.Ns_direction, seis_sp_survey.Ns_distance, seis_sp_survey.Ns_distance_ouom, seis_sp_survey.Ns_start_line, seis_sp_survey.Orientation, seis_sp_survey.Point_offset, seis_sp_survey.Point_offset_ouom, seis_sp_survey.Ppdm_guid, seis_sp_survey.Reference_loc, seis_sp_survey.Remark, seis_sp_survey.Source, seis_sp_survey.Surface_loc, seis_sp_survey.Well_node_id, seis_sp_survey.Row_changed_by, seis_sp_survey.Row_changed_date, seis_sp_survey.Row_created_by, seis_sp_survey.Row_created_date, seis_sp_survey.Row_effective_date, seis_sp_survey.Row_expiry_date, seis_sp_survey.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisSpSurvey(c *fiber.Ctx) error {
	var seis_sp_survey dto.Seis_sp_survey

	setDefaults(&seis_sp_survey)

	if err := json.Unmarshal(c.Body(), &seis_sp_survey); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_sp_survey.Seis_set_subtype = id

    if seis_sp_survey.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_sp_survey set seis_set_id = :1, seis_point_id = :2, seis_sp_survey_obs_no = :3, active_ind = :4, azimuth = :5, azimuth_ouom = :6, effective_date = :7, ew_direction = :8, ew_distance = :9, ew_distance_ouom = :10, ew_start_line = :11, expiry_date = :12, location_type = :13, monument_id = :14, monument_sf_subtype = :15, north_type = :16, ns_direction = :17, ns_distance = :18, ns_distance_ouom = :19, ns_start_line = :20, orientation = :21, point_offset = :22, point_offset_ouom = :23, ppdm_guid = :24, reference_loc = :25, remark = :26, source = :27, surface_loc = :28, well_node_id = :29, row_changed_by = :30, row_changed_date = :31, row_created_by = :32, row_effective_date = :33, row_expiry_date = :34, row_quality = :35 where seis_set_subtype = :37")
	if err != nil {
		return err
	}

	seis_sp_survey.Row_changed_date = formatDate(seis_sp_survey.Row_changed_date)
	seis_sp_survey.Effective_date = formatDateString(seis_sp_survey.Effective_date)
	seis_sp_survey.Expiry_date = formatDateString(seis_sp_survey.Expiry_date)
	seis_sp_survey.Row_effective_date = formatDateString(seis_sp_survey.Row_effective_date)
	seis_sp_survey.Row_expiry_date = formatDateString(seis_sp_survey.Row_expiry_date)






	rows, err := stmt.Exec(seis_sp_survey.Seis_set_id, seis_sp_survey.Seis_point_id, seis_sp_survey.Seis_sp_survey_obs_no, seis_sp_survey.Active_ind, seis_sp_survey.Azimuth, seis_sp_survey.Azimuth_ouom, seis_sp_survey.Effective_date, seis_sp_survey.Ew_direction, seis_sp_survey.Ew_distance, seis_sp_survey.Ew_distance_ouom, seis_sp_survey.Ew_start_line, seis_sp_survey.Expiry_date, seis_sp_survey.Location_type, seis_sp_survey.Monument_id, seis_sp_survey.Monument_sf_subtype, seis_sp_survey.North_type, seis_sp_survey.Ns_direction, seis_sp_survey.Ns_distance, seis_sp_survey.Ns_distance_ouom, seis_sp_survey.Ns_start_line, seis_sp_survey.Orientation, seis_sp_survey.Point_offset, seis_sp_survey.Point_offset_ouom, seis_sp_survey.Ppdm_guid, seis_sp_survey.Reference_loc, seis_sp_survey.Remark, seis_sp_survey.Source, seis_sp_survey.Surface_loc, seis_sp_survey.Well_node_id, seis_sp_survey.Row_changed_by, seis_sp_survey.Row_changed_date, seis_sp_survey.Row_created_by, seis_sp_survey.Row_effective_date, seis_sp_survey.Row_expiry_date, seis_sp_survey.Row_quality, seis_sp_survey.Seis_set_subtype)
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

func PatchSeisSpSurvey(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_sp_survey set "
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
	query += " where seis_set_subtype = :id"

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

func DeleteSeisSpSurvey(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_sp_survey dto.Seis_sp_survey
	seis_sp_survey.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_sp_survey where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_sp_survey.Seis_set_subtype)
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



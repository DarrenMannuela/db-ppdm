package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisAcqtnSurvey(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_acqtn_survey")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_acqtn_survey

	for rows.Next() {
		var seis_acqtn_survey dto.Seis_acqtn_survey
		if err := rows.Scan(&seis_acqtn_survey.Seis_set_subtype, &seis_acqtn_survey.Seis_acqtn_survey_id, &seis_acqtn_survey.Acqtn_survey_name, &seis_acqtn_survey.Active_ind, &seis_acqtn_survey.Area_id, &seis_acqtn_survey.Area_type, &seis_acqtn_survey.Completed_date, &seis_acqtn_survey.Completed_date_desc, &seis_acqtn_survey.Effective_date, &seis_acqtn_survey.Expiry_date, &seis_acqtn_survey.Ppdm_guid, &seis_acqtn_survey.Remark, &seis_acqtn_survey.Seis_dimension, &seis_acqtn_survey.Shot_for, &seis_acqtn_survey.Source, &seis_acqtn_survey.Start_date, &seis_acqtn_survey.Start_date_desc, &seis_acqtn_survey.Row_changed_by, &seis_acqtn_survey.Row_changed_date, &seis_acqtn_survey.Row_created_by, &seis_acqtn_survey.Row_created_date, &seis_acqtn_survey.Row_effective_date, &seis_acqtn_survey.Row_expiry_date, &seis_acqtn_survey.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_acqtn_survey to result
		result = append(result, seis_acqtn_survey)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisAcqtnSurvey(c *fiber.Ctx) error {
	var seis_acqtn_survey dto.Seis_acqtn_survey

	setDefaults(&seis_acqtn_survey)

	if err := json.Unmarshal(c.Body(), &seis_acqtn_survey); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_acqtn_survey values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24)")
	if err != nil {
		return err
	}
	seis_acqtn_survey.Row_created_date = formatDate(seis_acqtn_survey.Row_created_date)
	seis_acqtn_survey.Row_changed_date = formatDate(seis_acqtn_survey.Row_changed_date)
	seis_acqtn_survey.Completed_date = formatDateString(seis_acqtn_survey.Completed_date)
	seis_acqtn_survey.Effective_date = formatDateString(seis_acqtn_survey.Effective_date)
	seis_acqtn_survey.Expiry_date = formatDateString(seis_acqtn_survey.Expiry_date)
	seis_acqtn_survey.Start_date = formatDateString(seis_acqtn_survey.Start_date)
	seis_acqtn_survey.Row_effective_date = formatDateString(seis_acqtn_survey.Row_effective_date)
	seis_acqtn_survey.Row_expiry_date = formatDateString(seis_acqtn_survey.Row_expiry_date)






	rows, err := stmt.Exec(seis_acqtn_survey.Seis_set_subtype, seis_acqtn_survey.Seis_acqtn_survey_id, seis_acqtn_survey.Acqtn_survey_name, seis_acqtn_survey.Active_ind, seis_acqtn_survey.Area_id, seis_acqtn_survey.Area_type, seis_acqtn_survey.Completed_date, seis_acqtn_survey.Completed_date_desc, seis_acqtn_survey.Effective_date, seis_acqtn_survey.Expiry_date, seis_acqtn_survey.Ppdm_guid, seis_acqtn_survey.Remark, seis_acqtn_survey.Seis_dimension, seis_acqtn_survey.Shot_for, seis_acqtn_survey.Source, seis_acqtn_survey.Start_date, seis_acqtn_survey.Start_date_desc, seis_acqtn_survey.Row_changed_by, seis_acqtn_survey.Row_changed_date, seis_acqtn_survey.Row_created_by, seis_acqtn_survey.Row_created_date, seis_acqtn_survey.Row_effective_date, seis_acqtn_survey.Row_expiry_date, seis_acqtn_survey.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisAcqtnSurvey(c *fiber.Ctx) error {
	var seis_acqtn_survey dto.Seis_acqtn_survey

	setDefaults(&seis_acqtn_survey)

	if err := json.Unmarshal(c.Body(), &seis_acqtn_survey); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_acqtn_survey.Seis_set_subtype = id

    if seis_acqtn_survey.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_acqtn_survey set seis_acqtn_survey_id = :1, acqtn_survey_name = :2, active_ind = :3, area_id = :4, area_type = :5, completed_date = :6, completed_date_desc = :7, effective_date = :8, expiry_date = :9, ppdm_guid = :10, remark = :11, seis_dimension = :12, shot_for = :13, source = :14, start_date = :15, start_date_desc = :16, row_changed_by = :17, row_changed_date = :18, row_created_by = :19, row_effective_date = :20, row_expiry_date = :21, row_quality = :22 where seis_set_subtype = :24")
	if err != nil {
		return err
	}

	seis_acqtn_survey.Row_changed_date = formatDate(seis_acqtn_survey.Row_changed_date)
	seis_acqtn_survey.Completed_date = formatDateString(seis_acqtn_survey.Completed_date)
	seis_acqtn_survey.Effective_date = formatDateString(seis_acqtn_survey.Effective_date)
	seis_acqtn_survey.Expiry_date = formatDateString(seis_acqtn_survey.Expiry_date)
	seis_acqtn_survey.Start_date = formatDateString(seis_acqtn_survey.Start_date)
	seis_acqtn_survey.Row_effective_date = formatDateString(seis_acqtn_survey.Row_effective_date)
	seis_acqtn_survey.Row_expiry_date = formatDateString(seis_acqtn_survey.Row_expiry_date)






	rows, err := stmt.Exec(seis_acqtn_survey.Seis_acqtn_survey_id, seis_acqtn_survey.Acqtn_survey_name, seis_acqtn_survey.Active_ind, seis_acqtn_survey.Area_id, seis_acqtn_survey.Area_type, seis_acqtn_survey.Completed_date, seis_acqtn_survey.Completed_date_desc, seis_acqtn_survey.Effective_date, seis_acqtn_survey.Expiry_date, seis_acqtn_survey.Ppdm_guid, seis_acqtn_survey.Remark, seis_acqtn_survey.Seis_dimension, seis_acqtn_survey.Shot_for, seis_acqtn_survey.Source, seis_acqtn_survey.Start_date, seis_acqtn_survey.Start_date_desc, seis_acqtn_survey.Row_changed_by, seis_acqtn_survey.Row_changed_date, seis_acqtn_survey.Row_created_by, seis_acqtn_survey.Row_effective_date, seis_acqtn_survey.Row_expiry_date, seis_acqtn_survey.Row_quality, seis_acqtn_survey.Seis_set_subtype)
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

func PatchSeisAcqtnSurvey(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_acqtn_survey set "
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
		} else if key == "completed_date" || key == "effective_date" || key == "expiry_date" || key == "start_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteSeisAcqtnSurvey(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_acqtn_survey dto.Seis_acqtn_survey
	seis_acqtn_survey.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_acqtn_survey where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_acqtn_survey.Seis_set_subtype)
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



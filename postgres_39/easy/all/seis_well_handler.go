package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisWell(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_well")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_well

	for rows.Next() {
		var seis_well dto.Seis_well
		if err := rows.Scan(&seis_well.Seis_set_subtype, &seis_well.Seis_set_id, &seis_well.Active_ind, &seis_well.Checkshot_survey_type, &seis_well.Checkshot_velocity, &seis_well.Checkshot_velocity_ouom, &seis_well.Depth_datum, &seis_well.Depth_datum_elev, &seis_well.Depth_datum_elev_ouom, &seis_well.Dir_survey_id, &seis_well.Dir_survey_source, &seis_well.Dir_survey_uwi, &seis_well.Effective_date, &seis_well.Expiry_date, &seis_well.Ppdm_guid, &seis_well.Receiver_uwi, &seis_well.Remark, &seis_well.Seismic_path, &seis_well.Source, &seis_well.Source_uwi, &seis_well.Survey_id, &seis_well.Survey_run_num, &seis_well.Vsp_type, &seis_well.Row_changed_by, &seis_well.Row_changed_date, &seis_well.Row_created_by, &seis_well.Row_created_date, &seis_well.Row_effective_date, &seis_well.Row_expiry_date, &seis_well.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_well to result
		result = append(result, seis_well)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisWell(c *fiber.Ctx) error {
	var seis_well dto.Seis_well

	setDefaults(&seis_well)

	if err := json.Unmarshal(c.Body(), &seis_well); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_well values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30)")
	if err != nil {
		return err
	}
	seis_well.Row_created_date = formatDate(seis_well.Row_created_date)
	seis_well.Row_changed_date = formatDate(seis_well.Row_changed_date)
	seis_well.Effective_date = formatDateString(seis_well.Effective_date)
	seis_well.Expiry_date = formatDateString(seis_well.Expiry_date)
	seis_well.Row_effective_date = formatDateString(seis_well.Row_effective_date)
	seis_well.Row_expiry_date = formatDateString(seis_well.Row_expiry_date)






	rows, err := stmt.Exec(seis_well.Seis_set_subtype, seis_well.Seis_set_id, seis_well.Active_ind, seis_well.Checkshot_survey_type, seis_well.Checkshot_velocity, seis_well.Checkshot_velocity_ouom, seis_well.Depth_datum, seis_well.Depth_datum_elev, seis_well.Depth_datum_elev_ouom, seis_well.Dir_survey_id, seis_well.Dir_survey_source, seis_well.Dir_survey_uwi, seis_well.Effective_date, seis_well.Expiry_date, seis_well.Ppdm_guid, seis_well.Receiver_uwi, seis_well.Remark, seis_well.Seismic_path, seis_well.Source, seis_well.Source_uwi, seis_well.Survey_id, seis_well.Survey_run_num, seis_well.Vsp_type, seis_well.Row_changed_by, seis_well.Row_changed_date, seis_well.Row_created_by, seis_well.Row_created_date, seis_well.Row_effective_date, seis_well.Row_expiry_date, seis_well.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisWell(c *fiber.Ctx) error {
	var seis_well dto.Seis_well

	setDefaults(&seis_well)

	if err := json.Unmarshal(c.Body(), &seis_well); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_well.Seis_set_subtype = id

    if seis_well.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_well set seis_set_id = :1, active_ind = :2, checkshot_survey_type = :3, checkshot_velocity = :4, checkshot_velocity_ouom = :5, depth_datum = :6, depth_datum_elev = :7, depth_datum_elev_ouom = :8, dir_survey_id = :9, dir_survey_source = :10, dir_survey_uwi = :11, effective_date = :12, expiry_date = :13, ppdm_guid = :14, receiver_uwi = :15, remark = :16, seismic_path = :17, source = :18, source_uwi = :19, survey_id = :20, survey_run_num = :21, vsp_type = :22, row_changed_by = :23, row_changed_date = :24, row_created_by = :25, row_effective_date = :26, row_expiry_date = :27, row_quality = :28 where seis_set_subtype = :30")
	if err != nil {
		return err
	}

	seis_well.Row_changed_date = formatDate(seis_well.Row_changed_date)
	seis_well.Effective_date = formatDateString(seis_well.Effective_date)
	seis_well.Expiry_date = formatDateString(seis_well.Expiry_date)
	seis_well.Row_effective_date = formatDateString(seis_well.Row_effective_date)
	seis_well.Row_expiry_date = formatDateString(seis_well.Row_expiry_date)






	rows, err := stmt.Exec(seis_well.Seis_set_id, seis_well.Active_ind, seis_well.Checkshot_survey_type, seis_well.Checkshot_velocity, seis_well.Checkshot_velocity_ouom, seis_well.Depth_datum, seis_well.Depth_datum_elev, seis_well.Depth_datum_elev_ouom, seis_well.Dir_survey_id, seis_well.Dir_survey_source, seis_well.Dir_survey_uwi, seis_well.Effective_date, seis_well.Expiry_date, seis_well.Ppdm_guid, seis_well.Receiver_uwi, seis_well.Remark, seis_well.Seismic_path, seis_well.Source, seis_well.Source_uwi, seis_well.Survey_id, seis_well.Survey_run_num, seis_well.Vsp_type, seis_well.Row_changed_by, seis_well.Row_changed_date, seis_well.Row_created_by, seis_well.Row_effective_date, seis_well.Row_expiry_date, seis_well.Row_quality, seis_well.Seis_set_subtype)
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

func PatchSeisWell(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_well set "
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

func DeleteSeisWell(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_well dto.Seis_well
	seis_well.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_well where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_well.Seis_set_subtype)
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



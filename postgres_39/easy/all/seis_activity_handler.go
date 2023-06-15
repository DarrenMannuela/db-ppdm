package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisActivity(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_activity")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_activity

	for rows.Next() {
		var seis_activity dto.Seis_activity
		if err := rows.Scan(&seis_activity.Seis_set_subtype, &seis_activity.Seis_set_id, &seis_activity.Activity_obs_no, &seis_activity.Active_ind, &seis_activity.Activity_duration, &seis_activity.Activity_duration_ouom, &seis_activity.Area_size, &seis_activity.Area_size_ouom, &seis_activity.Crew_company_ba_id, &seis_activity.Crew_id, &seis_activity.Effective_date, &seis_activity.End_date, &seis_activity.End_time, &seis_activity.End_timezone, &seis_activity.Expiry_date, &seis_activity.First_nline_no, &seis_activity.First_seis_point_id, &seis_activity.First_xline_no, &seis_activity.Last_nline_no, &seis_activity.Last_seis_point_id, &seis_activity.Last_xline_no, &seis_activity.Line_length, &seis_activity.Line_length_ouom, &seis_activity.Owner_ba_id, &seis_activity.Ppdm_guid, &seis_activity.Remark, &seis_activity.Seis_activity_class, &seis_activity.Seis_activity_type, &seis_activity.Source, &seis_activity.Start_date, &seis_activity.Start_time, &seis_activity.Start_timezone, &seis_activity.Total_time_elapsed, &seis_activity.Total_time_elapsed_ouom, &seis_activity.Row_changed_by, &seis_activity.Row_changed_date, &seis_activity.Row_created_by, &seis_activity.Row_created_date, &seis_activity.Row_effective_date, &seis_activity.Row_expiry_date, &seis_activity.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_activity to result
		result = append(result, seis_activity)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisActivity(c *fiber.Ctx) error {
	var seis_activity dto.Seis_activity

	setDefaults(&seis_activity)

	if err := json.Unmarshal(c.Body(), &seis_activity); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_activity values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41)")
	if err != nil {
		return err
	}
	seis_activity.Row_created_date = formatDate(seis_activity.Row_created_date)
	seis_activity.Row_changed_date = formatDate(seis_activity.Row_changed_date)
	seis_activity.Effective_date = formatDateString(seis_activity.Effective_date)
	seis_activity.End_date = formatDateString(seis_activity.End_date)
	seis_activity.End_time = formatDateString(seis_activity.End_time)
	seis_activity.Expiry_date = formatDateString(seis_activity.Expiry_date)
	seis_activity.Start_date = formatDateString(seis_activity.Start_date)
	seis_activity.Start_time = formatDateString(seis_activity.Start_time)
	seis_activity.Row_effective_date = formatDateString(seis_activity.Row_effective_date)
	seis_activity.Row_expiry_date = formatDateString(seis_activity.Row_expiry_date)






	rows, err := stmt.Exec(seis_activity.Seis_set_subtype, seis_activity.Seis_set_id, seis_activity.Activity_obs_no, seis_activity.Active_ind, seis_activity.Activity_duration, seis_activity.Activity_duration_ouom, seis_activity.Area_size, seis_activity.Area_size_ouom, seis_activity.Crew_company_ba_id, seis_activity.Crew_id, seis_activity.Effective_date, seis_activity.End_date, seis_activity.End_time, seis_activity.End_timezone, seis_activity.Expiry_date, seis_activity.First_nline_no, seis_activity.First_seis_point_id, seis_activity.First_xline_no, seis_activity.Last_nline_no, seis_activity.Last_seis_point_id, seis_activity.Last_xline_no, seis_activity.Line_length, seis_activity.Line_length_ouom, seis_activity.Owner_ba_id, seis_activity.Ppdm_guid, seis_activity.Remark, seis_activity.Seis_activity_class, seis_activity.Seis_activity_type, seis_activity.Source, seis_activity.Start_date, seis_activity.Start_time, seis_activity.Start_timezone, seis_activity.Total_time_elapsed, seis_activity.Total_time_elapsed_ouom, seis_activity.Row_changed_by, seis_activity.Row_changed_date, seis_activity.Row_created_by, seis_activity.Row_created_date, seis_activity.Row_effective_date, seis_activity.Row_expiry_date, seis_activity.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisActivity(c *fiber.Ctx) error {
	var seis_activity dto.Seis_activity

	setDefaults(&seis_activity)

	if err := json.Unmarshal(c.Body(), &seis_activity); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_activity.Seis_set_subtype = id

    if seis_activity.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_activity set seis_set_id = :1, activity_obs_no = :2, active_ind = :3, activity_duration = :4, activity_duration_ouom = :5, area_size = :6, area_size_ouom = :7, crew_company_ba_id = :8, crew_id = :9, effective_date = :10, end_date = :11, end_time = :12, end_timezone = :13, expiry_date = :14, first_nline_no = :15, first_seis_point_id = :16, first_xline_no = :17, last_nline_no = :18, last_seis_point_id = :19, last_xline_no = :20, line_length = :21, line_length_ouom = :22, owner_ba_id = :23, ppdm_guid = :24, remark = :25, seis_activity_class = :26, seis_activity_type = :27, source = :28, start_date = :29, start_time = :30, start_timezone = :31, total_time_elapsed = :32, total_time_elapsed_ouom = :33, row_changed_by = :34, row_changed_date = :35, row_created_by = :36, row_effective_date = :37, row_expiry_date = :38, row_quality = :39 where seis_set_subtype = :41")
	if err != nil {
		return err
	}

	seis_activity.Row_changed_date = formatDate(seis_activity.Row_changed_date)
	seis_activity.Effective_date = formatDateString(seis_activity.Effective_date)
	seis_activity.End_date = formatDateString(seis_activity.End_date)
	seis_activity.End_time = formatDateString(seis_activity.End_time)
	seis_activity.Expiry_date = formatDateString(seis_activity.Expiry_date)
	seis_activity.Start_date = formatDateString(seis_activity.Start_date)
	seis_activity.Start_time = formatDateString(seis_activity.Start_time)
	seis_activity.Row_effective_date = formatDateString(seis_activity.Row_effective_date)
	seis_activity.Row_expiry_date = formatDateString(seis_activity.Row_expiry_date)






	rows, err := stmt.Exec(seis_activity.Seis_set_id, seis_activity.Activity_obs_no, seis_activity.Active_ind, seis_activity.Activity_duration, seis_activity.Activity_duration_ouom, seis_activity.Area_size, seis_activity.Area_size_ouom, seis_activity.Crew_company_ba_id, seis_activity.Crew_id, seis_activity.Effective_date, seis_activity.End_date, seis_activity.End_time, seis_activity.End_timezone, seis_activity.Expiry_date, seis_activity.First_nline_no, seis_activity.First_seis_point_id, seis_activity.First_xline_no, seis_activity.Last_nline_no, seis_activity.Last_seis_point_id, seis_activity.Last_xline_no, seis_activity.Line_length, seis_activity.Line_length_ouom, seis_activity.Owner_ba_id, seis_activity.Ppdm_guid, seis_activity.Remark, seis_activity.Seis_activity_class, seis_activity.Seis_activity_type, seis_activity.Source, seis_activity.Start_date, seis_activity.Start_time, seis_activity.Start_timezone, seis_activity.Total_time_elapsed, seis_activity.Total_time_elapsed_ouom, seis_activity.Row_changed_by, seis_activity.Row_changed_date, seis_activity.Row_created_by, seis_activity.Row_effective_date, seis_activity.Row_expiry_date, seis_activity.Row_quality, seis_activity.Seis_set_subtype)
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

func PatchSeisActivity(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_activity set "
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
		} else if key == "effective_date" || key == "end_date" || key == "end_time" || key == "expiry_date" || key == "start_date" || key == "start_time" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteSeisActivity(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_activity dto.Seis_activity
	seis_activity.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_activity where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_activity.Seis_set_subtype)
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



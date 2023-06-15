package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisPointSummary(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_point_summary")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_point_summary

	for rows.Next() {
		var seis_point_summary dto.Seis_point_summary
		if err := rows.Scan(&seis_point_summary.Seis_set_subtype, &seis_point_summary.Seis_set_id, &seis_point_summary.Seis_summary_obs_no, &seis_point_summary.Active_ind, &seis_point_summary.Area_size, &seis_point_summary.Area_size_ouom, &seis_point_summary.Cdp_coverage, &seis_point_summary.Effective_date, &seis_point_summary.Expiry_date, &seis_point_summary.First_nline_no, &seis_point_summary.First_seis_point_id, &seis_point_summary.First_xline_no, &seis_point_summary.Last_nline_no, &seis_point_summary.Last_seis_point_id, &seis_point_summary.Last_xline_no, &seis_point_summary.Line_length, &seis_point_summary.Line_length_ouom, &seis_point_summary.Ppdm_guid, &seis_point_summary.Remark, &seis_point_summary.Seis_station_type, &seis_point_summary.Source, &seis_point_summary.Summary_type, &seis_point_summary.Row_changed_by, &seis_point_summary.Row_changed_date, &seis_point_summary.Row_created_by, &seis_point_summary.Row_created_date, &seis_point_summary.Row_effective_date, &seis_point_summary.Row_expiry_date, &seis_point_summary.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_point_summary to result
		result = append(result, seis_point_summary)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisPointSummary(c *fiber.Ctx) error {
	var seis_point_summary dto.Seis_point_summary

	setDefaults(&seis_point_summary)

	if err := json.Unmarshal(c.Body(), &seis_point_summary); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_point_summary values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29)")
	if err != nil {
		return err
	}
	seis_point_summary.Row_created_date = formatDate(seis_point_summary.Row_created_date)
	seis_point_summary.Row_changed_date = formatDate(seis_point_summary.Row_changed_date)
	seis_point_summary.Effective_date = formatDateString(seis_point_summary.Effective_date)
	seis_point_summary.Expiry_date = formatDateString(seis_point_summary.Expiry_date)
	seis_point_summary.Row_effective_date = formatDateString(seis_point_summary.Row_effective_date)
	seis_point_summary.Row_expiry_date = formatDateString(seis_point_summary.Row_expiry_date)






	rows, err := stmt.Exec(seis_point_summary.Seis_set_subtype, seis_point_summary.Seis_set_id, seis_point_summary.Seis_summary_obs_no, seis_point_summary.Active_ind, seis_point_summary.Area_size, seis_point_summary.Area_size_ouom, seis_point_summary.Cdp_coverage, seis_point_summary.Effective_date, seis_point_summary.Expiry_date, seis_point_summary.First_nline_no, seis_point_summary.First_seis_point_id, seis_point_summary.First_xline_no, seis_point_summary.Last_nline_no, seis_point_summary.Last_seis_point_id, seis_point_summary.Last_xline_no, seis_point_summary.Line_length, seis_point_summary.Line_length_ouom, seis_point_summary.Ppdm_guid, seis_point_summary.Remark, seis_point_summary.Seis_station_type, seis_point_summary.Source, seis_point_summary.Summary_type, seis_point_summary.Row_changed_by, seis_point_summary.Row_changed_date, seis_point_summary.Row_created_by, seis_point_summary.Row_created_date, seis_point_summary.Row_effective_date, seis_point_summary.Row_expiry_date, seis_point_summary.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisPointSummary(c *fiber.Ctx) error {
	var seis_point_summary dto.Seis_point_summary

	setDefaults(&seis_point_summary)

	if err := json.Unmarshal(c.Body(), &seis_point_summary); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_point_summary.Seis_set_subtype = id

    if seis_point_summary.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_point_summary set seis_set_id = :1, seis_summary_obs_no = :2, active_ind = :3, area_size = :4, area_size_ouom = :5, cdp_coverage = :6, effective_date = :7, expiry_date = :8, first_nline_no = :9, first_seis_point_id = :10, first_xline_no = :11, last_nline_no = :12, last_seis_point_id = :13, last_xline_no = :14, line_length = :15, line_length_ouom = :16, ppdm_guid = :17, remark = :18, seis_station_type = :19, source = :20, summary_type = :21, row_changed_by = :22, row_changed_date = :23, row_created_by = :24, row_effective_date = :25, row_expiry_date = :26, row_quality = :27 where seis_set_subtype = :29")
	if err != nil {
		return err
	}

	seis_point_summary.Row_changed_date = formatDate(seis_point_summary.Row_changed_date)
	seis_point_summary.Effective_date = formatDateString(seis_point_summary.Effective_date)
	seis_point_summary.Expiry_date = formatDateString(seis_point_summary.Expiry_date)
	seis_point_summary.Row_effective_date = formatDateString(seis_point_summary.Row_effective_date)
	seis_point_summary.Row_expiry_date = formatDateString(seis_point_summary.Row_expiry_date)






	rows, err := stmt.Exec(seis_point_summary.Seis_set_id, seis_point_summary.Seis_summary_obs_no, seis_point_summary.Active_ind, seis_point_summary.Area_size, seis_point_summary.Area_size_ouom, seis_point_summary.Cdp_coverage, seis_point_summary.Effective_date, seis_point_summary.Expiry_date, seis_point_summary.First_nline_no, seis_point_summary.First_seis_point_id, seis_point_summary.First_xline_no, seis_point_summary.Last_nline_no, seis_point_summary.Last_seis_point_id, seis_point_summary.Last_xline_no, seis_point_summary.Line_length, seis_point_summary.Line_length_ouom, seis_point_summary.Ppdm_guid, seis_point_summary.Remark, seis_point_summary.Seis_station_type, seis_point_summary.Source, seis_point_summary.Summary_type, seis_point_summary.Row_changed_by, seis_point_summary.Row_changed_date, seis_point_summary.Row_created_by, seis_point_summary.Row_effective_date, seis_point_summary.Row_expiry_date, seis_point_summary.Row_quality, seis_point_summary.Seis_set_subtype)
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

func PatchSeisPointSummary(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_point_summary set "
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

func DeleteSeisPointSummary(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_point_summary dto.Seis_point_summary
	seis_point_summary.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_point_summary where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_point_summary.Seis_set_subtype)
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



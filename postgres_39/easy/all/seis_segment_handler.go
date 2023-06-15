package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisSegment(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_segment")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_segment

	for rows.Next() {
		var seis_segment dto.Seis_segment
		if err := rows.Scan(&seis_segment.Seis_set_subtype, &seis_segment.Segment_id, &seis_segment.Acqtn_design_id, &seis_segment.Active_ind, &seis_segment.Area_id, &seis_segment.Area_type, &seis_segment.Business_associate_id, &seis_segment.Effective_date, &seis_segment.Expiry_date, &seis_segment.Ppdm_guid, &seis_segment.Refraction_reflection, &seis_segment.Remark, &seis_segment.Row_audit_by, &seis_segment.Row_audit_date, &seis_segment.Seis_dimension, &seis_segment.Seis_line_id, &seis_segment.Seis_line_set_subtype, &seis_segment.Seis_segment_reason, &seis_segment.Seis_spectrum_type, &seis_segment.Seis_station_type, &seis_segment.Source, &seis_segment.Test_experimental, &seis_segment.Row_changed_by, &seis_segment.Row_changed_date, &seis_segment.Row_created_by, &seis_segment.Row_created_date, &seis_segment.Row_effective_date, &seis_segment.Row_expiry_date, &seis_segment.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_segment to result
		result = append(result, seis_segment)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisSegment(c *fiber.Ctx) error {
	var seis_segment dto.Seis_segment

	setDefaults(&seis_segment)

	if err := json.Unmarshal(c.Body(), &seis_segment); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_segment values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29)")
	if err != nil {
		return err
	}
	seis_segment.Row_created_date = formatDate(seis_segment.Row_created_date)
	seis_segment.Row_changed_date = formatDate(seis_segment.Row_changed_date)
	seis_segment.Effective_date = formatDateString(seis_segment.Effective_date)
	seis_segment.Expiry_date = formatDateString(seis_segment.Expiry_date)
	seis_segment.Row_audit_date = formatDateString(seis_segment.Row_audit_date)
	seis_segment.Row_effective_date = formatDateString(seis_segment.Row_effective_date)
	seis_segment.Row_expiry_date = formatDateString(seis_segment.Row_expiry_date)






	rows, err := stmt.Exec(seis_segment.Seis_set_subtype, seis_segment.Segment_id, seis_segment.Acqtn_design_id, seis_segment.Active_ind, seis_segment.Area_id, seis_segment.Area_type, seis_segment.Business_associate_id, seis_segment.Effective_date, seis_segment.Expiry_date, seis_segment.Ppdm_guid, seis_segment.Refraction_reflection, seis_segment.Remark, seis_segment.Row_audit_by, seis_segment.Row_audit_date, seis_segment.Seis_dimension, seis_segment.Seis_line_id, seis_segment.Seis_line_set_subtype, seis_segment.Seis_segment_reason, seis_segment.Seis_spectrum_type, seis_segment.Seis_station_type, seis_segment.Source, seis_segment.Test_experimental, seis_segment.Row_changed_by, seis_segment.Row_changed_date, seis_segment.Row_created_by, seis_segment.Row_created_date, seis_segment.Row_effective_date, seis_segment.Row_expiry_date, seis_segment.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisSegment(c *fiber.Ctx) error {
	var seis_segment dto.Seis_segment

	setDefaults(&seis_segment)

	if err := json.Unmarshal(c.Body(), &seis_segment); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_segment.Seis_set_subtype = id

    if seis_segment.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_segment set segment_id = :1, acqtn_design_id = :2, active_ind = :3, area_id = :4, area_type = :5, business_associate_id = :6, effective_date = :7, expiry_date = :8, ppdm_guid = :9, refraction_reflection = :10, remark = :11, row_audit_by = :12, row_audit_date = :13, seis_dimension = :14, seis_line_id = :15, seis_line_set_subtype = :16, seis_segment_reason = :17, seis_spectrum_type = :18, seis_station_type = :19, source = :20, test_experimental = :21, row_changed_by = :22, row_changed_date = :23, row_created_by = :24, row_effective_date = :25, row_expiry_date = :26, row_quality = :27 where seis_set_subtype = :29")
	if err != nil {
		return err
	}

	seis_segment.Row_changed_date = formatDate(seis_segment.Row_changed_date)
	seis_segment.Effective_date = formatDateString(seis_segment.Effective_date)
	seis_segment.Expiry_date = formatDateString(seis_segment.Expiry_date)
	seis_segment.Row_audit_date = formatDateString(seis_segment.Row_audit_date)
	seis_segment.Row_effective_date = formatDateString(seis_segment.Row_effective_date)
	seis_segment.Row_expiry_date = formatDateString(seis_segment.Row_expiry_date)






	rows, err := stmt.Exec(seis_segment.Segment_id, seis_segment.Acqtn_design_id, seis_segment.Active_ind, seis_segment.Area_id, seis_segment.Area_type, seis_segment.Business_associate_id, seis_segment.Effective_date, seis_segment.Expiry_date, seis_segment.Ppdm_guid, seis_segment.Refraction_reflection, seis_segment.Remark, seis_segment.Row_audit_by, seis_segment.Row_audit_date, seis_segment.Seis_dimension, seis_segment.Seis_line_id, seis_segment.Seis_line_set_subtype, seis_segment.Seis_segment_reason, seis_segment.Seis_spectrum_type, seis_segment.Seis_station_type, seis_segment.Source, seis_segment.Test_experimental, seis_segment.Row_changed_by, seis_segment.Row_changed_date, seis_segment.Row_created_by, seis_segment.Row_effective_date, seis_segment.Row_expiry_date, seis_segment.Row_quality, seis_segment.Seis_set_subtype)
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

func PatchSeisSegment(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_segment set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "row_audit_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteSeisSegment(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_segment dto.Seis_segment
	seis_segment.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_segment where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_segment.Seis_set_subtype)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisInspection(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_inspection")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_inspection

	for rows.Next() {
		var seis_inspection dto.Seis_inspection
		if err := rows.Scan(&seis_inspection.Inspection_id, &seis_inspection.Active_ind, &seis_inspection.Effective_date, &seis_inspection.Expiry_date, &seis_inspection.Inspected_length, &seis_inspection.Inspected_length_ouom, &seis_inspection.Inspecting_ba_id, &seis_inspection.Inspection_date, &seis_inspection.Inspection_status, &seis_inspection.Insp_loc_address_obs_no, &seis_inspection.Insp_loc_ba_id, &seis_inspection.Insp_loc_source, &seis_inspection.Min_length, &seis_inspection.Min_length_ouom, &seis_inspection.Ppdm_guid, &seis_inspection.Reason_type, &seis_inspection.Remark, &seis_inspection.Scheduled_date, &seis_inspection.Source, &seis_inspection.Row_changed_by, &seis_inspection.Row_changed_date, &seis_inspection.Row_created_by, &seis_inspection.Row_created_date, &seis_inspection.Row_effective_date, &seis_inspection.Row_expiry_date, &seis_inspection.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_inspection to result
		result = append(result, seis_inspection)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisInspection(c *fiber.Ctx) error {
	var seis_inspection dto.Seis_inspection

	setDefaults(&seis_inspection)

	if err := json.Unmarshal(c.Body(), &seis_inspection); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_inspection values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26)")
	if err != nil {
		return err
	}
	seis_inspection.Row_created_date = formatDate(seis_inspection.Row_created_date)
	seis_inspection.Row_changed_date = formatDate(seis_inspection.Row_changed_date)
	seis_inspection.Effective_date = formatDateString(seis_inspection.Effective_date)
	seis_inspection.Expiry_date = formatDateString(seis_inspection.Expiry_date)
	seis_inspection.Inspection_date = formatDateString(seis_inspection.Inspection_date)
	seis_inspection.Scheduled_date = formatDateString(seis_inspection.Scheduled_date)
	seis_inspection.Row_effective_date = formatDateString(seis_inspection.Row_effective_date)
	seis_inspection.Row_expiry_date = formatDateString(seis_inspection.Row_expiry_date)






	rows, err := stmt.Exec(seis_inspection.Inspection_id, seis_inspection.Active_ind, seis_inspection.Effective_date, seis_inspection.Expiry_date, seis_inspection.Inspected_length, seis_inspection.Inspected_length_ouom, seis_inspection.Inspecting_ba_id, seis_inspection.Inspection_date, seis_inspection.Inspection_status, seis_inspection.Insp_loc_address_obs_no, seis_inspection.Insp_loc_ba_id, seis_inspection.Insp_loc_source, seis_inspection.Min_length, seis_inspection.Min_length_ouom, seis_inspection.Ppdm_guid, seis_inspection.Reason_type, seis_inspection.Remark, seis_inspection.Scheduled_date, seis_inspection.Source, seis_inspection.Row_changed_by, seis_inspection.Row_changed_date, seis_inspection.Row_created_by, seis_inspection.Row_created_date, seis_inspection.Row_effective_date, seis_inspection.Row_expiry_date, seis_inspection.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisInspection(c *fiber.Ctx) error {
	var seis_inspection dto.Seis_inspection

	setDefaults(&seis_inspection)

	if err := json.Unmarshal(c.Body(), &seis_inspection); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_inspection.Inspection_id = id

    if seis_inspection.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_inspection set active_ind = :1, effective_date = :2, expiry_date = :3, inspected_length = :4, inspected_length_ouom = :5, inspecting_ba_id = :6, inspection_date = :7, inspection_status = :8, insp_loc_address_obs_no = :9, insp_loc_ba_id = :10, insp_loc_source = :11, min_length = :12, min_length_ouom = :13, ppdm_guid = :14, reason_type = :15, remark = :16, scheduled_date = :17, source = :18, row_changed_by = :19, row_changed_date = :20, row_created_by = :21, row_effective_date = :22, row_expiry_date = :23, row_quality = :24 where inspection_id = :26")
	if err != nil {
		return err
	}

	seis_inspection.Row_changed_date = formatDate(seis_inspection.Row_changed_date)
	seis_inspection.Effective_date = formatDateString(seis_inspection.Effective_date)
	seis_inspection.Expiry_date = formatDateString(seis_inspection.Expiry_date)
	seis_inspection.Inspection_date = formatDateString(seis_inspection.Inspection_date)
	seis_inspection.Scheduled_date = formatDateString(seis_inspection.Scheduled_date)
	seis_inspection.Row_effective_date = formatDateString(seis_inspection.Row_effective_date)
	seis_inspection.Row_expiry_date = formatDateString(seis_inspection.Row_expiry_date)






	rows, err := stmt.Exec(seis_inspection.Active_ind, seis_inspection.Effective_date, seis_inspection.Expiry_date, seis_inspection.Inspected_length, seis_inspection.Inspected_length_ouom, seis_inspection.Inspecting_ba_id, seis_inspection.Inspection_date, seis_inspection.Inspection_status, seis_inspection.Insp_loc_address_obs_no, seis_inspection.Insp_loc_ba_id, seis_inspection.Insp_loc_source, seis_inspection.Min_length, seis_inspection.Min_length_ouom, seis_inspection.Ppdm_guid, seis_inspection.Reason_type, seis_inspection.Remark, seis_inspection.Scheduled_date, seis_inspection.Source, seis_inspection.Row_changed_by, seis_inspection.Row_changed_date, seis_inspection.Row_created_by, seis_inspection.Row_effective_date, seis_inspection.Row_expiry_date, seis_inspection.Row_quality, seis_inspection.Inspection_id)
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

func PatchSeisInspection(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_inspection set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "inspection_date" || key == "scheduled_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where inspection_id = :id"

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

func DeleteSeisInspection(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_inspection dto.Seis_inspection
	seis_inspection.Inspection_id = id

	stmt, err := db.Prepare("delete from seis_inspection where inspection_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_inspection.Inspection_id)
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



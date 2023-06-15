package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisRecord(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_record")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_record

	for rows.Next() {
		var seis_record dto.Seis_record
		if err := rows.Scan(&seis_record.Seis_set_subtype, &seis_record.Seis_set_id, &seis_record.Record_id, &seis_record.Active_ind, &seis_record.Actual_acqtn_design_id, &seis_record.Effective_date, &seis_record.Expiry_date, &seis_record.Field_file_number, &seis_record.Information_item_id, &seis_record.Info_item_subtype, &seis_record.Logical_record_number, &seis_record.Patch_id, &seis_record.Patch_used_ind, &seis_record.Ppdm_guid, &seis_record.Rcrd_channel_count, &seis_record.Recording_remark, &seis_record.Record_number, &seis_record.Record_quality, &seis_record.Record_type, &seis_record.Remark, &seis_record.Seis_shot_point_id, &seis_record.Source, &seis_record.Tape_number, &seis_record.Time_delay, &seis_record.Time_delay_ouom, &seis_record.Uphole_time, &seis_record.Uphole_time_ouom, &seis_record.Vessel_config_obs_no, &seis_record.Vessel_id, &seis_record.Vessel_sf_subtype, &seis_record.X_offset, &seis_record.Y_offset, &seis_record.Z_offset, &seis_record.Row_changed_by, &seis_record.Row_changed_date, &seis_record.Row_created_by, &seis_record.Row_created_date, &seis_record.Row_effective_date, &seis_record.Row_expiry_date, &seis_record.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_record to result
		result = append(result, seis_record)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisRecord(c *fiber.Ctx) error {
	var seis_record dto.Seis_record

	setDefaults(&seis_record)

	if err := json.Unmarshal(c.Body(), &seis_record); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_record values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40)")
	if err != nil {
		return err
	}
	seis_record.Row_created_date = formatDate(seis_record.Row_created_date)
	seis_record.Row_changed_date = formatDate(seis_record.Row_changed_date)
	seis_record.Effective_date = formatDateString(seis_record.Effective_date)
	seis_record.Expiry_date = formatDateString(seis_record.Expiry_date)
	seis_record.Row_effective_date = formatDateString(seis_record.Row_effective_date)
	seis_record.Row_expiry_date = formatDateString(seis_record.Row_expiry_date)






	rows, err := stmt.Exec(seis_record.Seis_set_subtype, seis_record.Seis_set_id, seis_record.Record_id, seis_record.Active_ind, seis_record.Actual_acqtn_design_id, seis_record.Effective_date, seis_record.Expiry_date, seis_record.Field_file_number, seis_record.Information_item_id, seis_record.Info_item_subtype, seis_record.Logical_record_number, seis_record.Patch_id, seis_record.Patch_used_ind, seis_record.Ppdm_guid, seis_record.Rcrd_channel_count, seis_record.Recording_remark, seis_record.Record_number, seis_record.Record_quality, seis_record.Record_type, seis_record.Remark, seis_record.Seis_shot_point_id, seis_record.Source, seis_record.Tape_number, seis_record.Time_delay, seis_record.Time_delay_ouom, seis_record.Uphole_time, seis_record.Uphole_time_ouom, seis_record.Vessel_config_obs_no, seis_record.Vessel_id, seis_record.Vessel_sf_subtype, seis_record.X_offset, seis_record.Y_offset, seis_record.Z_offset, seis_record.Row_changed_by, seis_record.Row_changed_date, seis_record.Row_created_by, seis_record.Row_created_date, seis_record.Row_effective_date, seis_record.Row_expiry_date, seis_record.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisRecord(c *fiber.Ctx) error {
	var seis_record dto.Seis_record

	setDefaults(&seis_record)

	if err := json.Unmarshal(c.Body(), &seis_record); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_record.Seis_set_subtype = id

    if seis_record.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_record set seis_set_id = :1, record_id = :2, active_ind = :3, actual_acqtn_design_id = :4, effective_date = :5, expiry_date = :6, field_file_number = :7, information_item_id = :8, info_item_subtype = :9, logical_record_number = :10, patch_id = :11, patch_used_ind = :12, ppdm_guid = :13, rcrd_channel_count = :14, recording_remark = :15, record_number = :16, record_quality = :17, record_type = :18, remark = :19, seis_shot_point_id = :20, source = :21, tape_number = :22, time_delay = :23, time_delay_ouom = :24, uphole_time = :25, uphole_time_ouom = :26, vessel_config_obs_no = :27, vessel_id = :28, vessel_sf_subtype = :29, x_offset = :30, y_offset = :31, z_offset = :32, row_changed_by = :33, row_changed_date = :34, row_created_by = :35, row_effective_date = :36, row_expiry_date = :37, row_quality = :38 where seis_set_subtype = :40")
	if err != nil {
		return err
	}

	seis_record.Row_changed_date = formatDate(seis_record.Row_changed_date)
	seis_record.Effective_date = formatDateString(seis_record.Effective_date)
	seis_record.Expiry_date = formatDateString(seis_record.Expiry_date)
	seis_record.Row_effective_date = formatDateString(seis_record.Row_effective_date)
	seis_record.Row_expiry_date = formatDateString(seis_record.Row_expiry_date)






	rows, err := stmt.Exec(seis_record.Seis_set_id, seis_record.Record_id, seis_record.Active_ind, seis_record.Actual_acqtn_design_id, seis_record.Effective_date, seis_record.Expiry_date, seis_record.Field_file_number, seis_record.Information_item_id, seis_record.Info_item_subtype, seis_record.Logical_record_number, seis_record.Patch_id, seis_record.Patch_used_ind, seis_record.Ppdm_guid, seis_record.Rcrd_channel_count, seis_record.Recording_remark, seis_record.Record_number, seis_record.Record_quality, seis_record.Record_type, seis_record.Remark, seis_record.Seis_shot_point_id, seis_record.Source, seis_record.Tape_number, seis_record.Time_delay, seis_record.Time_delay_ouom, seis_record.Uphole_time, seis_record.Uphole_time_ouom, seis_record.Vessel_config_obs_no, seis_record.Vessel_id, seis_record.Vessel_sf_subtype, seis_record.X_offset, seis_record.Y_offset, seis_record.Z_offset, seis_record.Row_changed_by, seis_record.Row_changed_date, seis_record.Row_created_by, seis_record.Row_effective_date, seis_record.Row_expiry_date, seis_record.Row_quality, seis_record.Seis_set_subtype)
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

func PatchSeisRecord(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_record set "
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

func DeleteSeisRecord(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_record dto.Seis_record
	seis_record.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_record where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_record.Seis_set_subtype)
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



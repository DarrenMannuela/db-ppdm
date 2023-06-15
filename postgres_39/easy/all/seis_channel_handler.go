package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisChannel(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_channel")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_channel

	for rows.Next() {
		var seis_channel dto.Seis_channel
		if err := rows.Scan(&seis_channel.Seis_set_subtype, &seis_channel.Seis_set_id, &seis_channel.Record_id, &seis_channel.Channel_id, &seis_channel.Active_ind, &seis_channel.Channel_num, &seis_channel.Channel_type, &seis_channel.Effective_date, &seis_channel.Expiry_date, &seis_channel.File_num, &seis_channel.Ppdm_guid, &seis_channel.Remark, &seis_channel.Seis_receiver_point_id, &seis_channel.Source, &seis_channel.Streamer_id, &seis_channel.Vessel_config_obs_no, &seis_channel.Vessel_id, &seis_channel.Vessel_sf_subtype, &seis_channel.X_offset, &seis_channel.X_offset_ouom, &seis_channel.Y_offset, &seis_channel.Y_offset_ouom, &seis_channel.Z_offset, &seis_channel.Z_offset_ouom, &seis_channel.Row_changed_by, &seis_channel.Row_changed_date, &seis_channel.Row_created_by, &seis_channel.Row_created_date, &seis_channel.Row_effective_date, &seis_channel.Row_expiry_date, &seis_channel.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_channel to result
		result = append(result, seis_channel)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisChannel(c *fiber.Ctx) error {
	var seis_channel dto.Seis_channel

	setDefaults(&seis_channel)

	if err := json.Unmarshal(c.Body(), &seis_channel); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_channel values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31)")
	if err != nil {
		return err
	}
	seis_channel.Row_created_date = formatDate(seis_channel.Row_created_date)
	seis_channel.Row_changed_date = formatDate(seis_channel.Row_changed_date)
	seis_channel.Effective_date = formatDateString(seis_channel.Effective_date)
	seis_channel.Expiry_date = formatDateString(seis_channel.Expiry_date)
	seis_channel.Row_effective_date = formatDateString(seis_channel.Row_effective_date)
	seis_channel.Row_expiry_date = formatDateString(seis_channel.Row_expiry_date)






	rows, err := stmt.Exec(seis_channel.Seis_set_subtype, seis_channel.Seis_set_id, seis_channel.Record_id, seis_channel.Channel_id, seis_channel.Active_ind, seis_channel.Channel_num, seis_channel.Channel_type, seis_channel.Effective_date, seis_channel.Expiry_date, seis_channel.File_num, seis_channel.Ppdm_guid, seis_channel.Remark, seis_channel.Seis_receiver_point_id, seis_channel.Source, seis_channel.Streamer_id, seis_channel.Vessel_config_obs_no, seis_channel.Vessel_id, seis_channel.Vessel_sf_subtype, seis_channel.X_offset, seis_channel.X_offset_ouom, seis_channel.Y_offset, seis_channel.Y_offset_ouom, seis_channel.Z_offset, seis_channel.Z_offset_ouom, seis_channel.Row_changed_by, seis_channel.Row_changed_date, seis_channel.Row_created_by, seis_channel.Row_created_date, seis_channel.Row_effective_date, seis_channel.Row_expiry_date, seis_channel.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisChannel(c *fiber.Ctx) error {
	var seis_channel dto.Seis_channel

	setDefaults(&seis_channel)

	if err := json.Unmarshal(c.Body(), &seis_channel); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_channel.Seis_set_subtype = id

    if seis_channel.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_channel set seis_set_id = :1, record_id = :2, channel_id = :3, active_ind = :4, channel_num = :5, channel_type = :6, effective_date = :7, expiry_date = :8, file_num = :9, ppdm_guid = :10, remark = :11, seis_receiver_point_id = :12, source = :13, streamer_id = :14, vessel_config_obs_no = :15, vessel_id = :16, vessel_sf_subtype = :17, x_offset = :18, x_offset_ouom = :19, y_offset = :20, y_offset_ouom = :21, z_offset = :22, z_offset_ouom = :23, row_changed_by = :24, row_changed_date = :25, row_created_by = :26, row_effective_date = :27, row_expiry_date = :28, row_quality = :29 where seis_set_subtype = :31")
	if err != nil {
		return err
	}

	seis_channel.Row_changed_date = formatDate(seis_channel.Row_changed_date)
	seis_channel.Effective_date = formatDateString(seis_channel.Effective_date)
	seis_channel.Expiry_date = formatDateString(seis_channel.Expiry_date)
	seis_channel.Row_effective_date = formatDateString(seis_channel.Row_effective_date)
	seis_channel.Row_expiry_date = formatDateString(seis_channel.Row_expiry_date)






	rows, err := stmt.Exec(seis_channel.Seis_set_id, seis_channel.Record_id, seis_channel.Channel_id, seis_channel.Active_ind, seis_channel.Channel_num, seis_channel.Channel_type, seis_channel.Effective_date, seis_channel.Expiry_date, seis_channel.File_num, seis_channel.Ppdm_guid, seis_channel.Remark, seis_channel.Seis_receiver_point_id, seis_channel.Source, seis_channel.Streamer_id, seis_channel.Vessel_config_obs_no, seis_channel.Vessel_id, seis_channel.Vessel_sf_subtype, seis_channel.X_offset, seis_channel.X_offset_ouom, seis_channel.Y_offset, seis_channel.Y_offset_ouom, seis_channel.Z_offset, seis_channel.Z_offset_ouom, seis_channel.Row_changed_by, seis_channel.Row_changed_date, seis_channel.Row_created_by, seis_channel.Row_effective_date, seis_channel.Row_expiry_date, seis_channel.Row_quality, seis_channel.Seis_set_subtype)
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

func PatchSeisChannel(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_channel set "
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

func DeleteSeisChannel(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_channel dto.Seis_channel
	seis_channel.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_channel where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_channel.Seis_set_subtype)
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



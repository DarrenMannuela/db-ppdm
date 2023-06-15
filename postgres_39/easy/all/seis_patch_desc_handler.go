package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisPatchDesc(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_patch_desc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_patch_desc

	for rows.Next() {
		var seis_patch_desc dto.Seis_patch_desc
		if err := rows.Scan(&seis_patch_desc.Patch_id, &seis_patch_desc.Patch_desc_id, &seis_patch_desc.Active_ind, &seis_patch_desc.Effective_date, &seis_patch_desc.End_channel, &seis_patch_desc.End_station_id, &seis_patch_desc.End_x_offset, &seis_patch_desc.End_y_offset, &seis_patch_desc.End_z_offset, &seis_patch_desc.Expiry_date, &seis_patch_desc.Offset_ouom, &seis_patch_desc.Ppdm_guid, &seis_patch_desc.Recorded_line, &seis_patch_desc.Remark, &seis_patch_desc.Source, &seis_patch_desc.Start_channel, &seis_patch_desc.Start_station_id, &seis_patch_desc.Start_x_offset, &seis_patch_desc.Start_y_offset, &seis_patch_desc.Start_z_offset, &seis_patch_desc.Row_changed_by, &seis_patch_desc.Row_changed_date, &seis_patch_desc.Row_created_by, &seis_patch_desc.Row_created_date, &seis_patch_desc.Row_effective_date, &seis_patch_desc.Row_expiry_date, &seis_patch_desc.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_patch_desc to result
		result = append(result, seis_patch_desc)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisPatchDesc(c *fiber.Ctx) error {
	var seis_patch_desc dto.Seis_patch_desc

	setDefaults(&seis_patch_desc)

	if err := json.Unmarshal(c.Body(), &seis_patch_desc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_patch_desc values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27)")
	if err != nil {
		return err
	}
	seis_patch_desc.Row_created_date = formatDate(seis_patch_desc.Row_created_date)
	seis_patch_desc.Row_changed_date = formatDate(seis_patch_desc.Row_changed_date)
	seis_patch_desc.Effective_date = formatDateString(seis_patch_desc.Effective_date)
	seis_patch_desc.Expiry_date = formatDateString(seis_patch_desc.Expiry_date)
	seis_patch_desc.Row_effective_date = formatDateString(seis_patch_desc.Row_effective_date)
	seis_patch_desc.Row_expiry_date = formatDateString(seis_patch_desc.Row_expiry_date)






	rows, err := stmt.Exec(seis_patch_desc.Patch_id, seis_patch_desc.Patch_desc_id, seis_patch_desc.Active_ind, seis_patch_desc.Effective_date, seis_patch_desc.End_channel, seis_patch_desc.End_station_id, seis_patch_desc.End_x_offset, seis_patch_desc.End_y_offset, seis_patch_desc.End_z_offset, seis_patch_desc.Expiry_date, seis_patch_desc.Offset_ouom, seis_patch_desc.Ppdm_guid, seis_patch_desc.Recorded_line, seis_patch_desc.Remark, seis_patch_desc.Source, seis_patch_desc.Start_channel, seis_patch_desc.Start_station_id, seis_patch_desc.Start_x_offset, seis_patch_desc.Start_y_offset, seis_patch_desc.Start_z_offset, seis_patch_desc.Row_changed_by, seis_patch_desc.Row_changed_date, seis_patch_desc.Row_created_by, seis_patch_desc.Row_created_date, seis_patch_desc.Row_effective_date, seis_patch_desc.Row_expiry_date, seis_patch_desc.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisPatchDesc(c *fiber.Ctx) error {
	var seis_patch_desc dto.Seis_patch_desc

	setDefaults(&seis_patch_desc)

	if err := json.Unmarshal(c.Body(), &seis_patch_desc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_patch_desc.Patch_id = id

    if seis_patch_desc.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_patch_desc set patch_desc_id = :1, active_ind = :2, effective_date = :3, end_channel = :4, end_station_id = :5, end_x_offset = :6, end_y_offset = :7, end_z_offset = :8, expiry_date = :9, offset_ouom = :10, ppdm_guid = :11, recorded_line = :12, remark = :13, source = :14, start_channel = :15, start_station_id = :16, start_x_offset = :17, start_y_offset = :18, start_z_offset = :19, row_changed_by = :20, row_changed_date = :21, row_created_by = :22, row_effective_date = :23, row_expiry_date = :24, row_quality = :25 where patch_id = :27")
	if err != nil {
		return err
	}

	seis_patch_desc.Row_changed_date = formatDate(seis_patch_desc.Row_changed_date)
	seis_patch_desc.Effective_date = formatDateString(seis_patch_desc.Effective_date)
	seis_patch_desc.Expiry_date = formatDateString(seis_patch_desc.Expiry_date)
	seis_patch_desc.Row_effective_date = formatDateString(seis_patch_desc.Row_effective_date)
	seis_patch_desc.Row_expiry_date = formatDateString(seis_patch_desc.Row_expiry_date)






	rows, err := stmt.Exec(seis_patch_desc.Patch_desc_id, seis_patch_desc.Active_ind, seis_patch_desc.Effective_date, seis_patch_desc.End_channel, seis_patch_desc.End_station_id, seis_patch_desc.End_x_offset, seis_patch_desc.End_y_offset, seis_patch_desc.End_z_offset, seis_patch_desc.Expiry_date, seis_patch_desc.Offset_ouom, seis_patch_desc.Ppdm_guid, seis_patch_desc.Recorded_line, seis_patch_desc.Remark, seis_patch_desc.Source, seis_patch_desc.Start_channel, seis_patch_desc.Start_station_id, seis_patch_desc.Start_x_offset, seis_patch_desc.Start_y_offset, seis_patch_desc.Start_z_offset, seis_patch_desc.Row_changed_by, seis_patch_desc.Row_changed_date, seis_patch_desc.Row_created_by, seis_patch_desc.Row_effective_date, seis_patch_desc.Row_expiry_date, seis_patch_desc.Row_quality, seis_patch_desc.Patch_id)
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

func PatchSeisPatchDesc(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_patch_desc set "
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
	query += " where patch_id = :id"

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

func DeleteSeisPatchDesc(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_patch_desc dto.Seis_patch_desc
	seis_patch_desc.Patch_id = id

	stmt, err := db.Prepare("delete from seis_patch_desc where patch_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_patch_desc.Patch_id)
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



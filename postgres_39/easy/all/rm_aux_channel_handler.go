package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmAuxChannel(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_aux_channel")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_aux_channel

	for rows.Next() {
		var rm_aux_channel dto.Rm_aux_channel
		if err := rows.Scan(&rm_aux_channel.Info_item_subtype, &rm_aux_channel.Information_item_id, &rm_aux_channel.Channel_id, &rm_aux_channel.Active_ind, &rm_aux_channel.Channel_num, &rm_aux_channel.Channel_type, &rm_aux_channel.Description, &rm_aux_channel.Effective_date, &rm_aux_channel.Expiry_date, &rm_aux_channel.Ppdm_guid, &rm_aux_channel.Remark, &rm_aux_channel.Source, &rm_aux_channel.Row_changed_by, &rm_aux_channel.Row_changed_date, &rm_aux_channel.Row_created_by, &rm_aux_channel.Row_created_date, &rm_aux_channel.Row_effective_date, &rm_aux_channel.Row_expiry_date, &rm_aux_channel.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_aux_channel to result
		result = append(result, rm_aux_channel)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmAuxChannel(c *fiber.Ctx) error {
	var rm_aux_channel dto.Rm_aux_channel

	setDefaults(&rm_aux_channel)

	if err := json.Unmarshal(c.Body(), &rm_aux_channel); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_aux_channel values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	rm_aux_channel.Row_created_date = formatDate(rm_aux_channel.Row_created_date)
	rm_aux_channel.Row_changed_date = formatDate(rm_aux_channel.Row_changed_date)
	rm_aux_channel.Effective_date = formatDateString(rm_aux_channel.Effective_date)
	rm_aux_channel.Expiry_date = formatDateString(rm_aux_channel.Expiry_date)
	rm_aux_channel.Row_effective_date = formatDateString(rm_aux_channel.Row_effective_date)
	rm_aux_channel.Row_expiry_date = formatDateString(rm_aux_channel.Row_expiry_date)






	rows, err := stmt.Exec(rm_aux_channel.Info_item_subtype, rm_aux_channel.Information_item_id, rm_aux_channel.Channel_id, rm_aux_channel.Active_ind, rm_aux_channel.Channel_num, rm_aux_channel.Channel_type, rm_aux_channel.Description, rm_aux_channel.Effective_date, rm_aux_channel.Expiry_date, rm_aux_channel.Ppdm_guid, rm_aux_channel.Remark, rm_aux_channel.Source, rm_aux_channel.Row_changed_by, rm_aux_channel.Row_changed_date, rm_aux_channel.Row_created_by, rm_aux_channel.Row_created_date, rm_aux_channel.Row_effective_date, rm_aux_channel.Row_expiry_date, rm_aux_channel.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmAuxChannel(c *fiber.Ctx) error {
	var rm_aux_channel dto.Rm_aux_channel

	setDefaults(&rm_aux_channel)

	if err := json.Unmarshal(c.Body(), &rm_aux_channel); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_aux_channel.Info_item_subtype = id

    if rm_aux_channel.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_aux_channel set information_item_id = :1, channel_id = :2, active_ind = :3, channel_num = :4, channel_type = :5, description = :6, effective_date = :7, expiry_date = :8, ppdm_guid = :9, remark = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where info_item_subtype = :19")
	if err != nil {
		return err
	}

	rm_aux_channel.Row_changed_date = formatDate(rm_aux_channel.Row_changed_date)
	rm_aux_channel.Effective_date = formatDateString(rm_aux_channel.Effective_date)
	rm_aux_channel.Expiry_date = formatDateString(rm_aux_channel.Expiry_date)
	rm_aux_channel.Row_effective_date = formatDateString(rm_aux_channel.Row_effective_date)
	rm_aux_channel.Row_expiry_date = formatDateString(rm_aux_channel.Row_expiry_date)






	rows, err := stmt.Exec(rm_aux_channel.Information_item_id, rm_aux_channel.Channel_id, rm_aux_channel.Active_ind, rm_aux_channel.Channel_num, rm_aux_channel.Channel_type, rm_aux_channel.Description, rm_aux_channel.Effective_date, rm_aux_channel.Expiry_date, rm_aux_channel.Ppdm_guid, rm_aux_channel.Remark, rm_aux_channel.Source, rm_aux_channel.Row_changed_by, rm_aux_channel.Row_changed_date, rm_aux_channel.Row_created_by, rm_aux_channel.Row_effective_date, rm_aux_channel.Row_expiry_date, rm_aux_channel.Row_quality, rm_aux_channel.Info_item_subtype)
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

func PatchRmAuxChannel(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_aux_channel set "
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
	query += " where info_item_subtype = :id"

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

func DeleteRmAuxChannel(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_aux_channel dto.Rm_aux_channel
	rm_aux_channel.Info_item_subtype = id

	stmt, err := db.Prepare("delete from rm_aux_channel where info_item_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_aux_channel.Info_item_subtype)
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



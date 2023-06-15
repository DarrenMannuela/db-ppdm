package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRaMudLogColor(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ra_mud_log_color")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ra_mud_log_color

	for rows.Next() {
		var ra_mud_log_color dto.Ra_mud_log_color
		if err := rows.Scan(&ra_mud_log_color.Mud_log_color, &ra_mud_log_color.Alias_id, &ra_mud_log_color.Abbreviation, &ra_mud_log_color.Active_ind, &ra_mud_log_color.Alias_long_name, &ra_mud_log_color.Alias_reason_type, &ra_mud_log_color.Alias_short_name, &ra_mud_log_color.Alias_type, &ra_mud_log_color.Amended_date, &ra_mud_log_color.Created_date, &ra_mud_log_color.Effective_date, &ra_mud_log_color.Expiry_date, &ra_mud_log_color.Mud_log_color, &ra_mud_log_color.Original_ind, &ra_mud_log_color.Owner_ba_id, &ra_mud_log_color.Ppdm_guid, &ra_mud_log_color.Preferred_ind, &ra_mud_log_color.Reason_desc, &ra_mud_log_color.Remark, &ra_mud_log_color.Source, &ra_mud_log_color.Source_document, &ra_mud_log_color.Struckoff_date, &ra_mud_log_color.Sw_application_id, &ra_mud_log_color.Use_rule, &ra_mud_log_color.Row_changed_by, &ra_mud_log_color.Row_changed_date, &ra_mud_log_color.Row_created_by, &ra_mud_log_color.Row_created_date, &ra_mud_log_color.Row_effective_date, &ra_mud_log_color.Row_expiry_date, &ra_mud_log_color.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ra_mud_log_color to result
		result = append(result, ra_mud_log_color)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRaMudLogColor(c *fiber.Ctx) error {
	var ra_mud_log_color dto.Ra_mud_log_color

	setDefaults(&ra_mud_log_color)

	if err := json.Unmarshal(c.Body(), &ra_mud_log_color); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ra_mud_log_color values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31)")
	if err != nil {
		return err
	}
	ra_mud_log_color.Row_created_date = formatDate(ra_mud_log_color.Row_created_date)
	ra_mud_log_color.Row_changed_date = formatDate(ra_mud_log_color.Row_changed_date)
	ra_mud_log_color.Amended_date = formatDateString(ra_mud_log_color.Amended_date)
	ra_mud_log_color.Created_date = formatDateString(ra_mud_log_color.Created_date)
	ra_mud_log_color.Effective_date = formatDateString(ra_mud_log_color.Effective_date)
	ra_mud_log_color.Expiry_date = formatDateString(ra_mud_log_color.Expiry_date)
	ra_mud_log_color.Struckoff_date = formatDateString(ra_mud_log_color.Struckoff_date)
	ra_mud_log_color.Row_effective_date = formatDateString(ra_mud_log_color.Row_effective_date)
	ra_mud_log_color.Row_expiry_date = formatDateString(ra_mud_log_color.Row_expiry_date)






	rows, err := stmt.Exec(ra_mud_log_color.Mud_log_color, ra_mud_log_color.Alias_id, ra_mud_log_color.Abbreviation, ra_mud_log_color.Active_ind, ra_mud_log_color.Alias_long_name, ra_mud_log_color.Alias_reason_type, ra_mud_log_color.Alias_short_name, ra_mud_log_color.Alias_type, ra_mud_log_color.Amended_date, ra_mud_log_color.Created_date, ra_mud_log_color.Effective_date, ra_mud_log_color.Expiry_date, ra_mud_log_color.Mud_log_color, ra_mud_log_color.Original_ind, ra_mud_log_color.Owner_ba_id, ra_mud_log_color.Ppdm_guid, ra_mud_log_color.Preferred_ind, ra_mud_log_color.Reason_desc, ra_mud_log_color.Remark, ra_mud_log_color.Source, ra_mud_log_color.Source_document, ra_mud_log_color.Struckoff_date, ra_mud_log_color.Sw_application_id, ra_mud_log_color.Use_rule, ra_mud_log_color.Row_changed_by, ra_mud_log_color.Row_changed_date, ra_mud_log_color.Row_created_by, ra_mud_log_color.Row_created_date, ra_mud_log_color.Row_effective_date, ra_mud_log_color.Row_expiry_date, ra_mud_log_color.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRaMudLogColor(c *fiber.Ctx) error {
	var ra_mud_log_color dto.Ra_mud_log_color

	setDefaults(&ra_mud_log_color)

	if err := json.Unmarshal(c.Body(), &ra_mud_log_color); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ra_mud_log_color.Mud_log_color = id

    if ra_mud_log_color.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ra_mud_log_color set alias_id = :1, abbreviation = :2, active_ind = :3, alias_long_name = :4, alias_reason_type = :5, alias_short_name = :6, alias_type = :7, amended_date = :8, created_date = :9, effective_date = :10, expiry_date = :11, mud_log_color = :12, original_ind = :13, owner_ba_id = :14, ppdm_guid = :15, preferred_ind = :16, reason_desc = :17, remark = :18, source = :19, source_document = :20, struckoff_date = :21, sw_application_id = :22, use_rule = :23, row_changed_by = :24, row_changed_date = :25, row_created_by = :26, row_effective_date = :27, row_expiry_date = :28, row_quality = :29 where mud_log_color = :31")
	if err != nil {
		return err
	}

	ra_mud_log_color.Row_changed_date = formatDate(ra_mud_log_color.Row_changed_date)
	ra_mud_log_color.Amended_date = formatDateString(ra_mud_log_color.Amended_date)
	ra_mud_log_color.Created_date = formatDateString(ra_mud_log_color.Created_date)
	ra_mud_log_color.Effective_date = formatDateString(ra_mud_log_color.Effective_date)
	ra_mud_log_color.Expiry_date = formatDateString(ra_mud_log_color.Expiry_date)
	ra_mud_log_color.Struckoff_date = formatDateString(ra_mud_log_color.Struckoff_date)
	ra_mud_log_color.Row_effective_date = formatDateString(ra_mud_log_color.Row_effective_date)
	ra_mud_log_color.Row_expiry_date = formatDateString(ra_mud_log_color.Row_expiry_date)






	rows, err := stmt.Exec(ra_mud_log_color.Alias_id, ra_mud_log_color.Abbreviation, ra_mud_log_color.Active_ind, ra_mud_log_color.Alias_long_name, ra_mud_log_color.Alias_reason_type, ra_mud_log_color.Alias_short_name, ra_mud_log_color.Alias_type, ra_mud_log_color.Amended_date, ra_mud_log_color.Created_date, ra_mud_log_color.Effective_date, ra_mud_log_color.Expiry_date, ra_mud_log_color.Mud_log_color, ra_mud_log_color.Original_ind, ra_mud_log_color.Owner_ba_id, ra_mud_log_color.Ppdm_guid, ra_mud_log_color.Preferred_ind, ra_mud_log_color.Reason_desc, ra_mud_log_color.Remark, ra_mud_log_color.Source, ra_mud_log_color.Source_document, ra_mud_log_color.Struckoff_date, ra_mud_log_color.Sw_application_id, ra_mud_log_color.Use_rule, ra_mud_log_color.Row_changed_by, ra_mud_log_color.Row_changed_date, ra_mud_log_color.Row_created_by, ra_mud_log_color.Row_effective_date, ra_mud_log_color.Row_expiry_date, ra_mud_log_color.Row_quality, ra_mud_log_color.Mud_log_color)
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

func PatchRaMudLogColor(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ra_mud_log_color set "
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
		} else if key == "amended_date" || key == "created_date" || key == "effective_date" || key == "expiry_date" || key == "struckoff_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where mud_log_color = :id"

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

func DeleteRaMudLogColor(c *fiber.Ctx) error {
	id := c.Params("id")
	var ra_mud_log_color dto.Ra_mud_log_color
	ra_mud_log_color.Mud_log_color = id

	stmt, err := db.Prepare("delete from ra_mud_log_color where mud_log_color = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ra_mud_log_color.Mud_log_color)
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



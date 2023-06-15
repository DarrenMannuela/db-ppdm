package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmFileContent(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_file_content")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_file_content

	for rows.Next() {
		var rm_file_content dto.Rm_file_content
		if err := rows.Scan(&rm_file_content.File_id, &rm_file_content.Active_ind, &rm_file_content.Digital_format, &rm_file_content.Effective_date, &rm_file_content.Expiry_date, &rm_file_content.File_content, &rm_file_content.File_size, &rm_file_content.File_size_uom, &rm_file_content.Information_item_id, &rm_file_content.Info_item_subtype, &rm_file_content.Physical_item_id, &rm_file_content.Ppdm_guid, &rm_file_content.Remark, &rm_file_content.Source, &rm_file_content.Sw_application_id, &rm_file_content.Row_changed_by, &rm_file_content.Row_changed_date, &rm_file_content.Row_created_by, &rm_file_content.Row_created_date, &rm_file_content.Row_effective_date, &rm_file_content.Row_expiry_date, &rm_file_content.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_file_content to result
		result = append(result, rm_file_content)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmFileContent(c *fiber.Ctx) error {
	var rm_file_content dto.Rm_file_content

	setDefaults(&rm_file_content)

	if err := json.Unmarshal(c.Body(), &rm_file_content); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_file_content values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	rm_file_content.Row_created_date = formatDate(rm_file_content.Row_created_date)
	rm_file_content.Row_changed_date = formatDate(rm_file_content.Row_changed_date)
	rm_file_content.Effective_date = formatDateString(rm_file_content.Effective_date)
	rm_file_content.Expiry_date = formatDateString(rm_file_content.Expiry_date)
	rm_file_content.Row_effective_date = formatDateString(rm_file_content.Row_effective_date)
	rm_file_content.Row_expiry_date = formatDateString(rm_file_content.Row_expiry_date)






	rows, err := stmt.Exec(rm_file_content.File_id, rm_file_content.Active_ind, rm_file_content.Digital_format, rm_file_content.Effective_date, rm_file_content.Expiry_date, rm_file_content.File_content, rm_file_content.File_size, rm_file_content.File_size_uom, rm_file_content.Information_item_id, rm_file_content.Info_item_subtype, rm_file_content.Physical_item_id, rm_file_content.Ppdm_guid, rm_file_content.Remark, rm_file_content.Source, rm_file_content.Sw_application_id, rm_file_content.Row_changed_by, rm_file_content.Row_changed_date, rm_file_content.Row_created_by, rm_file_content.Row_created_date, rm_file_content.Row_effective_date, rm_file_content.Row_expiry_date, rm_file_content.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmFileContent(c *fiber.Ctx) error {
	var rm_file_content dto.Rm_file_content

	setDefaults(&rm_file_content)

	if err := json.Unmarshal(c.Body(), &rm_file_content); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_file_content.File_id = id

    if rm_file_content.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_file_content set active_ind = :1, digital_format = :2, effective_date = :3, expiry_date = :4, file_content = :5, file_size = :6, file_size_uom = :7, information_item_id = :8, info_item_subtype = :9, physical_item_id = :10, ppdm_guid = :11, remark = :12, source = :13, sw_application_id = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where file_id = :22")
	if err != nil {
		return err
	}

	rm_file_content.Row_changed_date = formatDate(rm_file_content.Row_changed_date)
	rm_file_content.Effective_date = formatDateString(rm_file_content.Effective_date)
	rm_file_content.Expiry_date = formatDateString(rm_file_content.Expiry_date)
	rm_file_content.Row_effective_date = formatDateString(rm_file_content.Row_effective_date)
	rm_file_content.Row_expiry_date = formatDateString(rm_file_content.Row_expiry_date)






	rows, err := stmt.Exec(rm_file_content.Active_ind, rm_file_content.Digital_format, rm_file_content.Effective_date, rm_file_content.Expiry_date, rm_file_content.File_content, rm_file_content.File_size, rm_file_content.File_size_uom, rm_file_content.Information_item_id, rm_file_content.Info_item_subtype, rm_file_content.Physical_item_id, rm_file_content.Ppdm_guid, rm_file_content.Remark, rm_file_content.Source, rm_file_content.Sw_application_id, rm_file_content.Row_changed_by, rm_file_content.Row_changed_date, rm_file_content.Row_created_by, rm_file_content.Row_effective_date, rm_file_content.Row_expiry_date, rm_file_content.Row_quality, rm_file_content.File_id)
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

func PatchRmFileContent(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_file_content set "
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
	query += " where file_id = :id"

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

func DeleteRmFileContent(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_file_content dto.Rm_file_content
	rm_file_content.File_id = id

	stmt, err := db.Prepare("delete from rm_file_content where file_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_file_content.File_id)
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



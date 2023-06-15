package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmDocument(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_document")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_document

	for rows.Next() {
		var rm_document dto.Rm_document
		if err := rows.Scan(&rm_document.Info_item_subtype, &rm_document.Information_item_id, &rm_document.Active_ind, &rm_document.Document_status, &rm_document.Document_type, &rm_document.Effective_date, &rm_document.Expiry_date, &rm_document.Ppdm_guid, &rm_document.Remark, &rm_document.Source, &rm_document.Source_document_id, &rm_document.Row_changed_by, &rm_document.Row_changed_date, &rm_document.Row_created_by, &rm_document.Row_created_date, &rm_document.Row_effective_date, &rm_document.Row_expiry_date, &rm_document.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_document to result
		result = append(result, rm_document)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmDocument(c *fiber.Ctx) error {
	var rm_document dto.Rm_document

	setDefaults(&rm_document)

	if err := json.Unmarshal(c.Body(), &rm_document); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_document values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	rm_document.Row_created_date = formatDate(rm_document.Row_created_date)
	rm_document.Row_changed_date = formatDate(rm_document.Row_changed_date)
	rm_document.Effective_date = formatDateString(rm_document.Effective_date)
	rm_document.Expiry_date = formatDateString(rm_document.Expiry_date)
	rm_document.Row_effective_date = formatDateString(rm_document.Row_effective_date)
	rm_document.Row_expiry_date = formatDateString(rm_document.Row_expiry_date)






	rows, err := stmt.Exec(rm_document.Info_item_subtype, rm_document.Information_item_id, rm_document.Active_ind, rm_document.Document_status, rm_document.Document_type, rm_document.Effective_date, rm_document.Expiry_date, rm_document.Ppdm_guid, rm_document.Remark, rm_document.Source, rm_document.Source_document_id, rm_document.Row_changed_by, rm_document.Row_changed_date, rm_document.Row_created_by, rm_document.Row_created_date, rm_document.Row_effective_date, rm_document.Row_expiry_date, rm_document.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmDocument(c *fiber.Ctx) error {
	var rm_document dto.Rm_document

	setDefaults(&rm_document)

	if err := json.Unmarshal(c.Body(), &rm_document); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_document.Info_item_subtype = id

    if rm_document.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_document set information_item_id = :1, active_ind = :2, document_status = :3, document_type = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, source = :9, source_document_id = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where info_item_subtype = :18")
	if err != nil {
		return err
	}

	rm_document.Row_changed_date = formatDate(rm_document.Row_changed_date)
	rm_document.Effective_date = formatDateString(rm_document.Effective_date)
	rm_document.Expiry_date = formatDateString(rm_document.Expiry_date)
	rm_document.Row_effective_date = formatDateString(rm_document.Row_effective_date)
	rm_document.Row_expiry_date = formatDateString(rm_document.Row_expiry_date)






	rows, err := stmt.Exec(rm_document.Information_item_id, rm_document.Active_ind, rm_document.Document_status, rm_document.Document_type, rm_document.Effective_date, rm_document.Expiry_date, rm_document.Ppdm_guid, rm_document.Remark, rm_document.Source, rm_document.Source_document_id, rm_document.Row_changed_by, rm_document.Row_changed_date, rm_document.Row_created_by, rm_document.Row_effective_date, rm_document.Row_expiry_date, rm_document.Row_quality, rm_document.Info_item_subtype)
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

func PatchRmDocument(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_document set "
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

func DeleteRmDocument(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_document dto.Rm_document
	rm_document.Info_item_subtype = id

	stmt, err := db.Prepare("delete from rm_document where info_item_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_document.Info_item_subtype)
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



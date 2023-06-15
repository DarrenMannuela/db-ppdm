package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmCopyRecord(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_copy_record")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_copy_record

	for rows.Next() {
		var rm_copy_record dto.Rm_copy_record
		if err := rows.Scan(&rm_copy_record.Old_physical_item_id, &rm_copy_record.New_physical_item_id, &rm_copy_record.Copy_seq_no, &rm_copy_record.Active_ind, &rm_copy_record.Effective_date, &rm_copy_record.Expiry_date, &rm_copy_record.New_record_no, &rm_copy_record.Old_record_no, &rm_copy_record.Ppdm_guid, &rm_copy_record.Remark, &rm_copy_record.Source, &rm_copy_record.Row_changed_by, &rm_copy_record.Row_changed_date, &rm_copy_record.Row_created_by, &rm_copy_record.Row_created_date, &rm_copy_record.Row_effective_date, &rm_copy_record.Row_expiry_date, &rm_copy_record.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_copy_record to result
		result = append(result, rm_copy_record)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmCopyRecord(c *fiber.Ctx) error {
	var rm_copy_record dto.Rm_copy_record

	setDefaults(&rm_copy_record)

	if err := json.Unmarshal(c.Body(), &rm_copy_record); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_copy_record values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	rm_copy_record.Row_created_date = formatDate(rm_copy_record.Row_created_date)
	rm_copy_record.Row_changed_date = formatDate(rm_copy_record.Row_changed_date)
	rm_copy_record.Effective_date = formatDateString(rm_copy_record.Effective_date)
	rm_copy_record.Expiry_date = formatDateString(rm_copy_record.Expiry_date)
	rm_copy_record.Row_effective_date = formatDateString(rm_copy_record.Row_effective_date)
	rm_copy_record.Row_expiry_date = formatDateString(rm_copy_record.Row_expiry_date)






	rows, err := stmt.Exec(rm_copy_record.Old_physical_item_id, rm_copy_record.New_physical_item_id, rm_copy_record.Copy_seq_no, rm_copy_record.Active_ind, rm_copy_record.Effective_date, rm_copy_record.Expiry_date, rm_copy_record.New_record_no, rm_copy_record.Old_record_no, rm_copy_record.Ppdm_guid, rm_copy_record.Remark, rm_copy_record.Source, rm_copy_record.Row_changed_by, rm_copy_record.Row_changed_date, rm_copy_record.Row_created_by, rm_copy_record.Row_created_date, rm_copy_record.Row_effective_date, rm_copy_record.Row_expiry_date, rm_copy_record.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmCopyRecord(c *fiber.Ctx) error {
	var rm_copy_record dto.Rm_copy_record

	setDefaults(&rm_copy_record)

	if err := json.Unmarshal(c.Body(), &rm_copy_record); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_copy_record.Old_physical_item_id = id

    if rm_copy_record.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_copy_record set new_physical_item_id = :1, copy_seq_no = :2, active_ind = :3, effective_date = :4, expiry_date = :5, new_record_no = :6, old_record_no = :7, ppdm_guid = :8, remark = :9, source = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where old_physical_item_id = :18")
	if err != nil {
		return err
	}

	rm_copy_record.Row_changed_date = formatDate(rm_copy_record.Row_changed_date)
	rm_copy_record.Effective_date = formatDateString(rm_copy_record.Effective_date)
	rm_copy_record.Expiry_date = formatDateString(rm_copy_record.Expiry_date)
	rm_copy_record.Row_effective_date = formatDateString(rm_copy_record.Row_effective_date)
	rm_copy_record.Row_expiry_date = formatDateString(rm_copy_record.Row_expiry_date)






	rows, err := stmt.Exec(rm_copy_record.New_physical_item_id, rm_copy_record.Copy_seq_no, rm_copy_record.Active_ind, rm_copy_record.Effective_date, rm_copy_record.Expiry_date, rm_copy_record.New_record_no, rm_copy_record.Old_record_no, rm_copy_record.Ppdm_guid, rm_copy_record.Remark, rm_copy_record.Source, rm_copy_record.Row_changed_by, rm_copy_record.Row_changed_date, rm_copy_record.Row_created_by, rm_copy_record.Row_effective_date, rm_copy_record.Row_expiry_date, rm_copy_record.Row_quality, rm_copy_record.Old_physical_item_id)
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

func PatchRmCopyRecord(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_copy_record set "
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
	query += " where old_physical_item_id = :id"

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

func DeleteRmCopyRecord(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_copy_record dto.Rm_copy_record
	rm_copy_record.Old_physical_item_id = id

	stmt, err := db.Prepare("delete from rm_copy_record where old_physical_item_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_copy_record.Old_physical_item_id)
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



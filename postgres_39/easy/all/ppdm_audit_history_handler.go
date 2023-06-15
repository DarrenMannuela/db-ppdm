package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmAuditHistory(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_audit_history")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_audit_history

	for rows.Next() {
		var ppdm_audit_history dto.Ppdm_audit_history
		if err := rows.Scan(&ppdm_audit_history.System_id, &ppdm_audit_history.Table_name, &ppdm_audit_history.Column_name, &ppdm_audit_history.Audit_row_guid, &ppdm_audit_history.Audit_seq_no, &ppdm_audit_history.Active_ind, &ppdm_audit_history.Audit_authorized_by_ba_id, &ppdm_audit_history.Audit_created_by_ba_id, &ppdm_audit_history.Audit_datetime, &ppdm_audit_history.Audit_desc, &ppdm_audit_history.Audit_reason, &ppdm_audit_history.Audit_source, &ppdm_audit_history.Audit_type, &ppdm_audit_history.Audit_verified_by_ba_id, &ppdm_audit_history.Data_type, &ppdm_audit_history.Effective_date, &ppdm_audit_history.Expiry_date, &ppdm_audit_history.New_date_value, &ppdm_audit_history.New_numeric_value, &ppdm_audit_history.New_numeric_value_ouom, &ppdm_audit_history.New_numeric_value_uom, &ppdm_audit_history.New_text_value, &ppdm_audit_history.Null_description, &ppdm_audit_history.Original_date_value, &ppdm_audit_history.Original_numeric_value, &ppdm_audit_history.Original_numeric_value_ouom, &ppdm_audit_history.Original_numeric_value_uom, &ppdm_audit_history.Original_text_value, &ppdm_audit_history.Ppdm_guid, &ppdm_audit_history.Remark, &ppdm_audit_history.Retention_period, &ppdm_audit_history.Source, &ppdm_audit_history.Row_changed_by, &ppdm_audit_history.Row_changed_date, &ppdm_audit_history.Row_created_by, &ppdm_audit_history.Row_created_date, &ppdm_audit_history.Row_effective_date, &ppdm_audit_history.Row_expiry_date, &ppdm_audit_history.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_audit_history to result
		result = append(result, ppdm_audit_history)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmAuditHistory(c *fiber.Ctx) error {
	var ppdm_audit_history dto.Ppdm_audit_history

	setDefaults(&ppdm_audit_history)

	if err := json.Unmarshal(c.Body(), &ppdm_audit_history); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_audit_history values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39)")
	if err != nil {
		return err
	}
	ppdm_audit_history.Row_created_date = formatDate(ppdm_audit_history.Row_created_date)
	ppdm_audit_history.Row_changed_date = formatDate(ppdm_audit_history.Row_changed_date)
	ppdm_audit_history.Audit_datetime = formatDateString(ppdm_audit_history.Audit_datetime)
	ppdm_audit_history.Effective_date = formatDateString(ppdm_audit_history.Effective_date)
	ppdm_audit_history.Expiry_date = formatDateString(ppdm_audit_history.Expiry_date)
	ppdm_audit_history.New_date_value = formatDateString(ppdm_audit_history.New_date_value)
	ppdm_audit_history.Original_date_value = formatDateString(ppdm_audit_history.Original_date_value)
	ppdm_audit_history.Row_effective_date = formatDateString(ppdm_audit_history.Row_effective_date)
	ppdm_audit_history.Row_expiry_date = formatDateString(ppdm_audit_history.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_audit_history.System_id, ppdm_audit_history.Table_name, ppdm_audit_history.Column_name, ppdm_audit_history.Audit_row_guid, ppdm_audit_history.Audit_seq_no, ppdm_audit_history.Active_ind, ppdm_audit_history.Audit_authorized_by_ba_id, ppdm_audit_history.Audit_created_by_ba_id, ppdm_audit_history.Audit_datetime, ppdm_audit_history.Audit_desc, ppdm_audit_history.Audit_reason, ppdm_audit_history.Audit_source, ppdm_audit_history.Audit_type, ppdm_audit_history.Audit_verified_by_ba_id, ppdm_audit_history.Data_type, ppdm_audit_history.Effective_date, ppdm_audit_history.Expiry_date, ppdm_audit_history.New_date_value, ppdm_audit_history.New_numeric_value, ppdm_audit_history.New_numeric_value_ouom, ppdm_audit_history.New_numeric_value_uom, ppdm_audit_history.New_text_value, ppdm_audit_history.Null_description, ppdm_audit_history.Original_date_value, ppdm_audit_history.Original_numeric_value, ppdm_audit_history.Original_numeric_value_ouom, ppdm_audit_history.Original_numeric_value_uom, ppdm_audit_history.Original_text_value, ppdm_audit_history.Ppdm_guid, ppdm_audit_history.Remark, ppdm_audit_history.Retention_period, ppdm_audit_history.Source, ppdm_audit_history.Row_changed_by, ppdm_audit_history.Row_changed_date, ppdm_audit_history.Row_created_by, ppdm_audit_history.Row_created_date, ppdm_audit_history.Row_effective_date, ppdm_audit_history.Row_expiry_date, ppdm_audit_history.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmAuditHistory(c *fiber.Ctx) error {
	var ppdm_audit_history dto.Ppdm_audit_history

	setDefaults(&ppdm_audit_history)

	if err := json.Unmarshal(c.Body(), &ppdm_audit_history); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_audit_history.System_id = id

    if ppdm_audit_history.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_audit_history set table_name = :1, column_name = :2, audit_row_guid = :3, audit_seq_no = :4, active_ind = :5, audit_authorized_by_ba_id = :6, audit_created_by_ba_id = :7, audit_datetime = :8, audit_desc = :9, audit_reason = :10, audit_source = :11, audit_type = :12, audit_verified_by_ba_id = :13, data_type = :14, effective_date = :15, expiry_date = :16, new_date_value = :17, new_numeric_value = :18, new_numeric_value_ouom = :19, new_numeric_value_uom = :20, new_text_value = :21, null_description = :22, original_date_value = :23, original_numeric_value = :24, original_numeric_value_ouom = :25, original_numeric_value_uom = :26, original_text_value = :27, ppdm_guid = :28, remark = :29, retention_period = :30, source = :31, row_changed_by = :32, row_changed_date = :33, row_created_by = :34, row_effective_date = :35, row_expiry_date = :36, row_quality = :37 where system_id = :39")
	if err != nil {
		return err
	}

	ppdm_audit_history.Row_changed_date = formatDate(ppdm_audit_history.Row_changed_date)
	ppdm_audit_history.Audit_datetime = formatDateString(ppdm_audit_history.Audit_datetime)
	ppdm_audit_history.Effective_date = formatDateString(ppdm_audit_history.Effective_date)
	ppdm_audit_history.Expiry_date = formatDateString(ppdm_audit_history.Expiry_date)
	ppdm_audit_history.New_date_value = formatDateString(ppdm_audit_history.New_date_value)
	ppdm_audit_history.Original_date_value = formatDateString(ppdm_audit_history.Original_date_value)
	ppdm_audit_history.Row_effective_date = formatDateString(ppdm_audit_history.Row_effective_date)
	ppdm_audit_history.Row_expiry_date = formatDateString(ppdm_audit_history.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_audit_history.Table_name, ppdm_audit_history.Column_name, ppdm_audit_history.Audit_row_guid, ppdm_audit_history.Audit_seq_no, ppdm_audit_history.Active_ind, ppdm_audit_history.Audit_authorized_by_ba_id, ppdm_audit_history.Audit_created_by_ba_id, ppdm_audit_history.Audit_datetime, ppdm_audit_history.Audit_desc, ppdm_audit_history.Audit_reason, ppdm_audit_history.Audit_source, ppdm_audit_history.Audit_type, ppdm_audit_history.Audit_verified_by_ba_id, ppdm_audit_history.Data_type, ppdm_audit_history.Effective_date, ppdm_audit_history.Expiry_date, ppdm_audit_history.New_date_value, ppdm_audit_history.New_numeric_value, ppdm_audit_history.New_numeric_value_ouom, ppdm_audit_history.New_numeric_value_uom, ppdm_audit_history.New_text_value, ppdm_audit_history.Null_description, ppdm_audit_history.Original_date_value, ppdm_audit_history.Original_numeric_value, ppdm_audit_history.Original_numeric_value_ouom, ppdm_audit_history.Original_numeric_value_uom, ppdm_audit_history.Original_text_value, ppdm_audit_history.Ppdm_guid, ppdm_audit_history.Remark, ppdm_audit_history.Retention_period, ppdm_audit_history.Source, ppdm_audit_history.Row_changed_by, ppdm_audit_history.Row_changed_date, ppdm_audit_history.Row_created_by, ppdm_audit_history.Row_effective_date, ppdm_audit_history.Row_expiry_date, ppdm_audit_history.Row_quality, ppdm_audit_history.System_id)
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

func PatchPpdmAuditHistory(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_audit_history set "
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
		} else if key == "audit_datetime" || key == "effective_date" || key == "expiry_date" || key == "new_date_value" || key == "original_date_value" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where system_id = :id"

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

func DeletePpdmAuditHistory(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_audit_history dto.Ppdm_audit_history
	ppdm_audit_history.System_id = id

	stmt, err := db.Prepare("delete from ppdm_audit_history where system_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_audit_history.System_id)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmAuditHistoryRem(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_audit_history_rem")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_audit_history_rem

	for rows.Next() {
		var ppdm_audit_history_rem dto.Ppdm_audit_history_rem
		if err := rows.Scan(&ppdm_audit_history_rem.System_id, &ppdm_audit_history_rem.Table_name, &ppdm_audit_history_rem.Column_name, &ppdm_audit_history_rem.Audit_row_guid, &ppdm_audit_history_rem.Audit_seq_no, &ppdm_audit_history_rem.Remark_seq_no, &ppdm_audit_history_rem.Active_ind, &ppdm_audit_history_rem.Audit_datetime, &ppdm_audit_history_rem.Effective_date, &ppdm_audit_history_rem.Expiry_date, &ppdm_audit_history_rem.Ppdm_guid, &ppdm_audit_history_rem.Remark, &ppdm_audit_history_rem.Remark_about_ba_id, &ppdm_audit_history_rem.Remark_by_ba_id, &ppdm_audit_history_rem.Remark_for_ba_id, &ppdm_audit_history_rem.Remark_type, &ppdm_audit_history_rem.Remark_use_type, &ppdm_audit_history_rem.Retention_period, &ppdm_audit_history_rem.Source, &ppdm_audit_history_rem.Row_changed_by, &ppdm_audit_history_rem.Row_changed_date, &ppdm_audit_history_rem.Row_created_by, &ppdm_audit_history_rem.Row_created_date, &ppdm_audit_history_rem.Row_effective_date, &ppdm_audit_history_rem.Row_expiry_date, &ppdm_audit_history_rem.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_audit_history_rem to result
		result = append(result, ppdm_audit_history_rem)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmAuditHistoryRem(c *fiber.Ctx) error {
	var ppdm_audit_history_rem dto.Ppdm_audit_history_rem

	setDefaults(&ppdm_audit_history_rem)

	if err := json.Unmarshal(c.Body(), &ppdm_audit_history_rem); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_audit_history_rem values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26)")
	if err != nil {
		return err
	}
	ppdm_audit_history_rem.Row_created_date = formatDate(ppdm_audit_history_rem.Row_created_date)
	ppdm_audit_history_rem.Row_changed_date = formatDate(ppdm_audit_history_rem.Row_changed_date)
	ppdm_audit_history_rem.Audit_datetime = formatDateString(ppdm_audit_history_rem.Audit_datetime)
	ppdm_audit_history_rem.Effective_date = formatDateString(ppdm_audit_history_rem.Effective_date)
	ppdm_audit_history_rem.Expiry_date = formatDateString(ppdm_audit_history_rem.Expiry_date)
	ppdm_audit_history_rem.Row_effective_date = formatDateString(ppdm_audit_history_rem.Row_effective_date)
	ppdm_audit_history_rem.Row_expiry_date = formatDateString(ppdm_audit_history_rem.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_audit_history_rem.System_id, ppdm_audit_history_rem.Table_name, ppdm_audit_history_rem.Column_name, ppdm_audit_history_rem.Audit_row_guid, ppdm_audit_history_rem.Audit_seq_no, ppdm_audit_history_rem.Remark_seq_no, ppdm_audit_history_rem.Active_ind, ppdm_audit_history_rem.Audit_datetime, ppdm_audit_history_rem.Effective_date, ppdm_audit_history_rem.Expiry_date, ppdm_audit_history_rem.Ppdm_guid, ppdm_audit_history_rem.Remark, ppdm_audit_history_rem.Remark_about_ba_id, ppdm_audit_history_rem.Remark_by_ba_id, ppdm_audit_history_rem.Remark_for_ba_id, ppdm_audit_history_rem.Remark_type, ppdm_audit_history_rem.Remark_use_type, ppdm_audit_history_rem.Retention_period, ppdm_audit_history_rem.Source, ppdm_audit_history_rem.Row_changed_by, ppdm_audit_history_rem.Row_changed_date, ppdm_audit_history_rem.Row_created_by, ppdm_audit_history_rem.Row_created_date, ppdm_audit_history_rem.Row_effective_date, ppdm_audit_history_rem.Row_expiry_date, ppdm_audit_history_rem.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmAuditHistoryRem(c *fiber.Ctx) error {
	var ppdm_audit_history_rem dto.Ppdm_audit_history_rem

	setDefaults(&ppdm_audit_history_rem)

	if err := json.Unmarshal(c.Body(), &ppdm_audit_history_rem); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_audit_history_rem.System_id = id

    if ppdm_audit_history_rem.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_audit_history_rem set table_name = :1, column_name = :2, audit_row_guid = :3, audit_seq_no = :4, remark_seq_no = :5, active_ind = :6, audit_datetime = :7, effective_date = :8, expiry_date = :9, ppdm_guid = :10, remark = :11, remark_about_ba_id = :12, remark_by_ba_id = :13, remark_for_ba_id = :14, remark_type = :15, remark_use_type = :16, retention_period = :17, source = :18, row_changed_by = :19, row_changed_date = :20, row_created_by = :21, row_effective_date = :22, row_expiry_date = :23, row_quality = :24 where system_id = :26")
	if err != nil {
		return err
	}

	ppdm_audit_history_rem.Row_changed_date = formatDate(ppdm_audit_history_rem.Row_changed_date)
	ppdm_audit_history_rem.Audit_datetime = formatDateString(ppdm_audit_history_rem.Audit_datetime)
	ppdm_audit_history_rem.Effective_date = formatDateString(ppdm_audit_history_rem.Effective_date)
	ppdm_audit_history_rem.Expiry_date = formatDateString(ppdm_audit_history_rem.Expiry_date)
	ppdm_audit_history_rem.Row_effective_date = formatDateString(ppdm_audit_history_rem.Row_effective_date)
	ppdm_audit_history_rem.Row_expiry_date = formatDateString(ppdm_audit_history_rem.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_audit_history_rem.Table_name, ppdm_audit_history_rem.Column_name, ppdm_audit_history_rem.Audit_row_guid, ppdm_audit_history_rem.Audit_seq_no, ppdm_audit_history_rem.Remark_seq_no, ppdm_audit_history_rem.Active_ind, ppdm_audit_history_rem.Audit_datetime, ppdm_audit_history_rem.Effective_date, ppdm_audit_history_rem.Expiry_date, ppdm_audit_history_rem.Ppdm_guid, ppdm_audit_history_rem.Remark, ppdm_audit_history_rem.Remark_about_ba_id, ppdm_audit_history_rem.Remark_by_ba_id, ppdm_audit_history_rem.Remark_for_ba_id, ppdm_audit_history_rem.Remark_type, ppdm_audit_history_rem.Remark_use_type, ppdm_audit_history_rem.Retention_period, ppdm_audit_history_rem.Source, ppdm_audit_history_rem.Row_changed_by, ppdm_audit_history_rem.Row_changed_date, ppdm_audit_history_rem.Row_created_by, ppdm_audit_history_rem.Row_effective_date, ppdm_audit_history_rem.Row_expiry_date, ppdm_audit_history_rem.Row_quality, ppdm_audit_history_rem.System_id)
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

func PatchPpdmAuditHistoryRem(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_audit_history_rem set "
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
		} else if key == "audit_datetime" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeletePpdmAuditHistoryRem(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_audit_history_rem dto.Ppdm_audit_history_rem
	ppdm_audit_history_rem.System_id = id

	stmt, err := db.Prepare("delete from ppdm_audit_history_rem where system_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_audit_history_rem.System_id)
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



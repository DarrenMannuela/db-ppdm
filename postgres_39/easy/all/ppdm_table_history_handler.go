package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmTableHistory(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_table_history")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_table_history

	for rows.Next() {
		var ppdm_table_history dto.Ppdm_table_history
		if err := rows.Scan(&ppdm_table_history.System_id, &ppdm_table_history.Table_name, &ppdm_table_history.History_seq_no, &ppdm_table_history.Active_ind, &ppdm_table_history.Audit_reason, &ppdm_table_history.Authorized_by_ba_id, &ppdm_table_history.Delete_record, &ppdm_table_history.Effective_date, &ppdm_table_history.Expiry_date, &ppdm_table_history.Ppdm_guid, &ppdm_table_history.Remark, &ppdm_table_history.Retention_period, &ppdm_table_history.Source, &ppdm_table_history.Row_changed_by, &ppdm_table_history.Row_changed_date, &ppdm_table_history.Row_created_by, &ppdm_table_history.Row_created_date, &ppdm_table_history.Row_effective_date, &ppdm_table_history.Row_expiry_date, &ppdm_table_history.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_table_history to result
		result = append(result, ppdm_table_history)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmTableHistory(c *fiber.Ctx) error {
	var ppdm_table_history dto.Ppdm_table_history

	setDefaults(&ppdm_table_history)

	if err := json.Unmarshal(c.Body(), &ppdm_table_history); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_table_history values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	ppdm_table_history.Row_created_date = formatDate(ppdm_table_history.Row_created_date)
	ppdm_table_history.Row_changed_date = formatDate(ppdm_table_history.Row_changed_date)
	ppdm_table_history.Effective_date = formatDateString(ppdm_table_history.Effective_date)
	ppdm_table_history.Expiry_date = formatDateString(ppdm_table_history.Expiry_date)
	ppdm_table_history.Row_effective_date = formatDateString(ppdm_table_history.Row_effective_date)
	ppdm_table_history.Row_expiry_date = formatDateString(ppdm_table_history.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_table_history.System_id, ppdm_table_history.Table_name, ppdm_table_history.History_seq_no, ppdm_table_history.Active_ind, ppdm_table_history.Audit_reason, ppdm_table_history.Authorized_by_ba_id, ppdm_table_history.Delete_record, ppdm_table_history.Effective_date, ppdm_table_history.Expiry_date, ppdm_table_history.Ppdm_guid, ppdm_table_history.Remark, ppdm_table_history.Retention_period, ppdm_table_history.Source, ppdm_table_history.Row_changed_by, ppdm_table_history.Row_changed_date, ppdm_table_history.Row_created_by, ppdm_table_history.Row_created_date, ppdm_table_history.Row_effective_date, ppdm_table_history.Row_expiry_date, ppdm_table_history.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmTableHistory(c *fiber.Ctx) error {
	var ppdm_table_history dto.Ppdm_table_history

	setDefaults(&ppdm_table_history)

	if err := json.Unmarshal(c.Body(), &ppdm_table_history); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_table_history.System_id = id

    if ppdm_table_history.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_table_history set table_name = :1, history_seq_no = :2, active_ind = :3, audit_reason = :4, authorized_by_ba_id = :5, delete_record = :6, effective_date = :7, expiry_date = :8, ppdm_guid = :9, remark = :10, retention_period = :11, source = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where system_id = :20")
	if err != nil {
		return err
	}

	ppdm_table_history.Row_changed_date = formatDate(ppdm_table_history.Row_changed_date)
	ppdm_table_history.Effective_date = formatDateString(ppdm_table_history.Effective_date)
	ppdm_table_history.Expiry_date = formatDateString(ppdm_table_history.Expiry_date)
	ppdm_table_history.Row_effective_date = formatDateString(ppdm_table_history.Row_effective_date)
	ppdm_table_history.Row_expiry_date = formatDateString(ppdm_table_history.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_table_history.Table_name, ppdm_table_history.History_seq_no, ppdm_table_history.Active_ind, ppdm_table_history.Audit_reason, ppdm_table_history.Authorized_by_ba_id, ppdm_table_history.Delete_record, ppdm_table_history.Effective_date, ppdm_table_history.Expiry_date, ppdm_table_history.Ppdm_guid, ppdm_table_history.Remark, ppdm_table_history.Retention_period, ppdm_table_history.Source, ppdm_table_history.Row_changed_by, ppdm_table_history.Row_changed_date, ppdm_table_history.Row_created_by, ppdm_table_history.Row_effective_date, ppdm_table_history.Row_expiry_date, ppdm_table_history.Row_quality, ppdm_table_history.System_id)
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

func PatchPpdmTableHistory(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_table_history set "
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

func DeletePpdmTableHistory(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_table_history dto.Ppdm_table_history
	ppdm_table_history.System_id = id

	stmt, err := db.Prepare("delete from ppdm_table_history where system_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_table_history.System_id)
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



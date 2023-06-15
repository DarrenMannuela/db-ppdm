package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_print_well_report/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmQualityControl(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_quality_control")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_quality_control

	for rows.Next() {
		var ppdm_quality_control dto.Ppdm_quality_control
		if err := rows.Scan(&ppdm_quality_control.System_id, &ppdm_quality_control.Table_name, &ppdm_quality_control.Qc_seq_no, &ppdm_quality_control.Active_ind, &ppdm_quality_control.Checked_by_ba_id, &ppdm_quality_control.Column_name, &ppdm_quality_control.Current_date_value, &ppdm_quality_control.Current_numeric_value, &ppdm_quality_control.Current_numeric_value_ouom, &ppdm_quality_control.Current_numeric_value_uom, &ppdm_quality_control.Current_row_guid, &ppdm_quality_control.Current_text_value, &ppdm_quality_control.Data_type, &ppdm_quality_control.Done_by_ba_id, &ppdm_quality_control.Effective_date, &ppdm_quality_control.Expiry_date, &ppdm_quality_control.Null_description, &ppdm_quality_control.Ppdm_guid, &ppdm_quality_control.Qc_datetime, &ppdm_quality_control.Qc_desc, &ppdm_quality_control.Qc_status, &ppdm_quality_control.Qc_type, &ppdm_quality_control.Quality_type, &ppdm_quality_control.Remark, &ppdm_quality_control.Retention_period, &ppdm_quality_control.Source, &ppdm_quality_control.Row_changed_by, &ppdm_quality_control.Row_changed_date, &ppdm_quality_control.Row_created_by, &ppdm_quality_control.Row_created_date, &ppdm_quality_control.Row_effective_date, &ppdm_quality_control.Row_expiry_date, &ppdm_quality_control.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_quality_control to result
		result = append(result, ppdm_quality_control)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmQualityControl(c *fiber.Ctx) error {
	var ppdm_quality_control dto.Ppdm_quality_control

	setDefaults(&ppdm_quality_control)

	if err := json.Unmarshal(c.Body(), &ppdm_quality_control); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_quality_control values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33)")
	if err != nil {
		return err
	}
	ppdm_quality_control.Row_created_date = formatDate(ppdm_quality_control.Row_created_date)
	ppdm_quality_control.Row_changed_date = formatDate(ppdm_quality_control.Row_changed_date)
	ppdm_quality_control.Current_date_value = formatDateString(ppdm_quality_control.Current_date_value)
	ppdm_quality_control.Effective_date = formatDateString(ppdm_quality_control.Effective_date)
	ppdm_quality_control.Expiry_date = formatDateString(ppdm_quality_control.Expiry_date)
	ppdm_quality_control.Qc_datetime = formatDateString(ppdm_quality_control.Qc_datetime)
	ppdm_quality_control.Row_effective_date = formatDateString(ppdm_quality_control.Row_effective_date)
	ppdm_quality_control.Row_expiry_date = formatDateString(ppdm_quality_control.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_quality_control.System_id, ppdm_quality_control.Table_name, ppdm_quality_control.Qc_seq_no, ppdm_quality_control.Active_ind, ppdm_quality_control.Checked_by_ba_id, ppdm_quality_control.Column_name, ppdm_quality_control.Current_date_value, ppdm_quality_control.Current_numeric_value, ppdm_quality_control.Current_numeric_value_ouom, ppdm_quality_control.Current_numeric_value_uom, ppdm_quality_control.Current_row_guid, ppdm_quality_control.Current_text_value, ppdm_quality_control.Data_type, ppdm_quality_control.Done_by_ba_id, ppdm_quality_control.Effective_date, ppdm_quality_control.Expiry_date, ppdm_quality_control.Null_description, ppdm_quality_control.Ppdm_guid, ppdm_quality_control.Qc_datetime, ppdm_quality_control.Qc_desc, ppdm_quality_control.Qc_status, ppdm_quality_control.Qc_type, ppdm_quality_control.Quality_type, ppdm_quality_control.Remark, ppdm_quality_control.Retention_period, ppdm_quality_control.Source, ppdm_quality_control.Row_changed_by, ppdm_quality_control.Row_changed_date, ppdm_quality_control.Row_created_by, ppdm_quality_control.Row_created_date, ppdm_quality_control.Row_effective_date, ppdm_quality_control.Row_expiry_date, ppdm_quality_control.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmQualityControl(c *fiber.Ctx) error {
	var ppdm_quality_control dto.Ppdm_quality_control

	setDefaults(&ppdm_quality_control)

	if err := json.Unmarshal(c.Body(), &ppdm_quality_control); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_quality_control.System_id = id

    if ppdm_quality_control.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_quality_control set table_name = :1, qc_seq_no = :2, active_ind = :3, checked_by_ba_id = :4, column_name = :5, current_date_value = :6, current_numeric_value = :7, current_numeric_value_ouom = :8, current_numeric_value_uom = :9, current_row_guid = :10, current_text_value = :11, data_type = :12, done_by_ba_id = :13, effective_date = :14, expiry_date = :15, null_description = :16, ppdm_guid = :17, qc_datetime = :18, qc_desc = :19, qc_status = :20, qc_type = :21, quality_type = :22, remark = :23, retention_period = :24, source = :25, row_changed_by = :26, row_changed_date = :27, row_created_by = :28, row_effective_date = :29, row_expiry_date = :30, row_quality = :31 where system_id = :33")
	if err != nil {
		return err
	}

	ppdm_quality_control.Row_changed_date = formatDate(ppdm_quality_control.Row_changed_date)
	ppdm_quality_control.Current_date_value = formatDateString(ppdm_quality_control.Current_date_value)
	ppdm_quality_control.Effective_date = formatDateString(ppdm_quality_control.Effective_date)
	ppdm_quality_control.Expiry_date = formatDateString(ppdm_quality_control.Expiry_date)
	ppdm_quality_control.Qc_datetime = formatDateString(ppdm_quality_control.Qc_datetime)
	ppdm_quality_control.Row_effective_date = formatDateString(ppdm_quality_control.Row_effective_date)
	ppdm_quality_control.Row_expiry_date = formatDateString(ppdm_quality_control.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_quality_control.Table_name, ppdm_quality_control.Qc_seq_no, ppdm_quality_control.Active_ind, ppdm_quality_control.Checked_by_ba_id, ppdm_quality_control.Column_name, ppdm_quality_control.Current_date_value, ppdm_quality_control.Current_numeric_value, ppdm_quality_control.Current_numeric_value_ouom, ppdm_quality_control.Current_numeric_value_uom, ppdm_quality_control.Current_row_guid, ppdm_quality_control.Current_text_value, ppdm_quality_control.Data_type, ppdm_quality_control.Done_by_ba_id, ppdm_quality_control.Effective_date, ppdm_quality_control.Expiry_date, ppdm_quality_control.Null_description, ppdm_quality_control.Ppdm_guid, ppdm_quality_control.Qc_datetime, ppdm_quality_control.Qc_desc, ppdm_quality_control.Qc_status, ppdm_quality_control.Qc_type, ppdm_quality_control.Quality_type, ppdm_quality_control.Remark, ppdm_quality_control.Retention_period, ppdm_quality_control.Source, ppdm_quality_control.Row_changed_by, ppdm_quality_control.Row_changed_date, ppdm_quality_control.Row_created_by, ppdm_quality_control.Row_effective_date, ppdm_quality_control.Row_expiry_date, ppdm_quality_control.Row_quality, ppdm_quality_control.System_id)
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

func PatchPpdmQualityControl(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_quality_control set "
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
		} else if key == "current_date_value" || key == "effective_date" || key == "expiry_date" || key == "qc_datetime" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDate(&date)
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

func DeletePpdmQualityControl(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_quality_control dto.Ppdm_quality_control
	ppdm_quality_control.System_id = id

	stmt, err := db.Prepare("delete from ppdm_quality_control where system_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_quality_control.System_id)
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



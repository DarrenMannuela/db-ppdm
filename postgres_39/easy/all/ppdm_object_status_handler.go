package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmObjectStatus(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_object_status")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_object_status

	for rows.Next() {
		var ppdm_object_status dto.Ppdm_object_status
		if err := rows.Scan(&ppdm_object_status.Status_id, &ppdm_object_status.Status_obs_no, &ppdm_object_status.Active_ind, &ppdm_object_status.Code_version_obs_no, &ppdm_object_status.Code_version_source, &ppdm_object_status.Column_name, &ppdm_object_status.Constraint_name, &ppdm_object_status.Effective_date, &ppdm_object_status.Expiry_date, &ppdm_object_status.Index_id, &ppdm_object_status.Object_name, &ppdm_object_status.Object_status, &ppdm_object_status.Object_type, &ppdm_object_status.Ppdm_guid, &ppdm_object_status.Procedure_id, &ppdm_object_status.Property_set_id, &ppdm_object_status.Remark, &ppdm_object_status.Rule_id, &ppdm_object_status.Source, &ppdm_object_status.System_id, &ppdm_object_status.Table_name, &ppdm_object_status.Row_changed_by, &ppdm_object_status.Row_changed_date, &ppdm_object_status.Row_created_by, &ppdm_object_status.Row_created_date, &ppdm_object_status.Row_effective_date, &ppdm_object_status.Row_expiry_date, &ppdm_object_status.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_object_status to result
		result = append(result, ppdm_object_status)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmObjectStatus(c *fiber.Ctx) error {
	var ppdm_object_status dto.Ppdm_object_status

	setDefaults(&ppdm_object_status)

	if err := json.Unmarshal(c.Body(), &ppdm_object_status); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_object_status values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28)")
	if err != nil {
		return err
	}
	ppdm_object_status.Row_created_date = formatDate(ppdm_object_status.Row_created_date)
	ppdm_object_status.Row_changed_date = formatDate(ppdm_object_status.Row_changed_date)
	ppdm_object_status.Effective_date = formatDateString(ppdm_object_status.Effective_date)
	ppdm_object_status.Expiry_date = formatDateString(ppdm_object_status.Expiry_date)
	ppdm_object_status.Row_effective_date = formatDateString(ppdm_object_status.Row_effective_date)
	ppdm_object_status.Row_expiry_date = formatDateString(ppdm_object_status.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_object_status.Status_id, ppdm_object_status.Status_obs_no, ppdm_object_status.Active_ind, ppdm_object_status.Code_version_obs_no, ppdm_object_status.Code_version_source, ppdm_object_status.Column_name, ppdm_object_status.Constraint_name, ppdm_object_status.Effective_date, ppdm_object_status.Expiry_date, ppdm_object_status.Index_id, ppdm_object_status.Object_name, ppdm_object_status.Object_status, ppdm_object_status.Object_type, ppdm_object_status.Ppdm_guid, ppdm_object_status.Procedure_id, ppdm_object_status.Property_set_id, ppdm_object_status.Remark, ppdm_object_status.Rule_id, ppdm_object_status.Source, ppdm_object_status.System_id, ppdm_object_status.Table_name, ppdm_object_status.Row_changed_by, ppdm_object_status.Row_changed_date, ppdm_object_status.Row_created_by, ppdm_object_status.Row_created_date, ppdm_object_status.Row_effective_date, ppdm_object_status.Row_expiry_date, ppdm_object_status.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmObjectStatus(c *fiber.Ctx) error {
	var ppdm_object_status dto.Ppdm_object_status

	setDefaults(&ppdm_object_status)

	if err := json.Unmarshal(c.Body(), &ppdm_object_status); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_object_status.Status_id = id

    if ppdm_object_status.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_object_status set status_obs_no = :1, active_ind = :2, code_version_obs_no = :3, code_version_source = :4, column_name = :5, constraint_name = :6, effective_date = :7, expiry_date = :8, index_id = :9, object_name = :10, object_status = :11, object_type = :12, ppdm_guid = :13, procedure_id = :14, property_set_id = :15, remark = :16, rule_id = :17, source = :18, system_id = :19, table_name = :20, row_changed_by = :21, row_changed_date = :22, row_created_by = :23, row_effective_date = :24, row_expiry_date = :25, row_quality = :26 where status_id = :28")
	if err != nil {
		return err
	}

	ppdm_object_status.Row_changed_date = formatDate(ppdm_object_status.Row_changed_date)
	ppdm_object_status.Effective_date = formatDateString(ppdm_object_status.Effective_date)
	ppdm_object_status.Expiry_date = formatDateString(ppdm_object_status.Expiry_date)
	ppdm_object_status.Row_effective_date = formatDateString(ppdm_object_status.Row_effective_date)
	ppdm_object_status.Row_expiry_date = formatDateString(ppdm_object_status.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_object_status.Status_obs_no, ppdm_object_status.Active_ind, ppdm_object_status.Code_version_obs_no, ppdm_object_status.Code_version_source, ppdm_object_status.Column_name, ppdm_object_status.Constraint_name, ppdm_object_status.Effective_date, ppdm_object_status.Expiry_date, ppdm_object_status.Index_id, ppdm_object_status.Object_name, ppdm_object_status.Object_status, ppdm_object_status.Object_type, ppdm_object_status.Ppdm_guid, ppdm_object_status.Procedure_id, ppdm_object_status.Property_set_id, ppdm_object_status.Remark, ppdm_object_status.Rule_id, ppdm_object_status.Source, ppdm_object_status.System_id, ppdm_object_status.Table_name, ppdm_object_status.Row_changed_by, ppdm_object_status.Row_changed_date, ppdm_object_status.Row_created_by, ppdm_object_status.Row_effective_date, ppdm_object_status.Row_expiry_date, ppdm_object_status.Row_quality, ppdm_object_status.Status_id)
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

func PatchPpdmObjectStatus(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_object_status set "
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
	query += " where status_id = :id"

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

func DeletePpdmObjectStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_object_status dto.Ppdm_object_status
	ppdm_object_status.Status_id = id

	stmt, err := db.Prepare("delete from ppdm_object_status where status_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_object_status.Status_id)
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



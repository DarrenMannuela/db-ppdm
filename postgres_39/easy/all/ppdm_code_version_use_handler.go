package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmCodeVersionUse(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_code_version_use")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_code_version_use

	for rows.Next() {
		var ppdm_code_version_use dto.Ppdm_code_version_use
		if err := rows.Scan(&ppdm_code_version_use.System_id, &ppdm_code_version_use.Table_name, &ppdm_code_version_use.Source, &ppdm_code_version_use.Version_obs_no, &ppdm_code_version_use.Use_obs_no, &ppdm_code_version_use.Active_ind, &ppdm_code_version_use.Effective_date, &ppdm_code_version_use.Expiry_date, &ppdm_code_version_use.Group_id, &ppdm_code_version_use.Ppdm_guid, &ppdm_code_version_use.Preferred_ind, &ppdm_code_version_use.Procedure_id, &ppdm_code_version_use.Remark, &ppdm_code_version_use.Source_document_id, &ppdm_code_version_use.Sw_application_id, &ppdm_code_version_use.Use_owner_ba_id, &ppdm_code_version_use.Use_rule_desc, &ppdm_code_version_use.Row_changed_by, &ppdm_code_version_use.Row_changed_date, &ppdm_code_version_use.Row_created_by, &ppdm_code_version_use.Row_created_date, &ppdm_code_version_use.Row_effective_date, &ppdm_code_version_use.Row_expiry_date, &ppdm_code_version_use.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_code_version_use to result
		result = append(result, ppdm_code_version_use)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmCodeVersionUse(c *fiber.Ctx) error {
	var ppdm_code_version_use dto.Ppdm_code_version_use

	setDefaults(&ppdm_code_version_use)

	if err := json.Unmarshal(c.Body(), &ppdm_code_version_use); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_code_version_use values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24)")
	if err != nil {
		return err
	}
	ppdm_code_version_use.Row_created_date = formatDate(ppdm_code_version_use.Row_created_date)
	ppdm_code_version_use.Row_changed_date = formatDate(ppdm_code_version_use.Row_changed_date)
	ppdm_code_version_use.Effective_date = formatDateString(ppdm_code_version_use.Effective_date)
	ppdm_code_version_use.Expiry_date = formatDateString(ppdm_code_version_use.Expiry_date)
	ppdm_code_version_use.Row_effective_date = formatDateString(ppdm_code_version_use.Row_effective_date)
	ppdm_code_version_use.Row_expiry_date = formatDateString(ppdm_code_version_use.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_code_version_use.System_id, ppdm_code_version_use.Table_name, ppdm_code_version_use.Source, ppdm_code_version_use.Version_obs_no, ppdm_code_version_use.Use_obs_no, ppdm_code_version_use.Active_ind, ppdm_code_version_use.Effective_date, ppdm_code_version_use.Expiry_date, ppdm_code_version_use.Group_id, ppdm_code_version_use.Ppdm_guid, ppdm_code_version_use.Preferred_ind, ppdm_code_version_use.Procedure_id, ppdm_code_version_use.Remark, ppdm_code_version_use.Source_document_id, ppdm_code_version_use.Sw_application_id, ppdm_code_version_use.Use_owner_ba_id, ppdm_code_version_use.Use_rule_desc, ppdm_code_version_use.Row_changed_by, ppdm_code_version_use.Row_changed_date, ppdm_code_version_use.Row_created_by, ppdm_code_version_use.Row_created_date, ppdm_code_version_use.Row_effective_date, ppdm_code_version_use.Row_expiry_date, ppdm_code_version_use.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmCodeVersionUse(c *fiber.Ctx) error {
	var ppdm_code_version_use dto.Ppdm_code_version_use

	setDefaults(&ppdm_code_version_use)

	if err := json.Unmarshal(c.Body(), &ppdm_code_version_use); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_code_version_use.System_id = id

    if ppdm_code_version_use.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_code_version_use set table_name = :1, source = :2, version_obs_no = :3, use_obs_no = :4, active_ind = :5, effective_date = :6, expiry_date = :7, group_id = :8, ppdm_guid = :9, preferred_ind = :10, procedure_id = :11, remark = :12, source_document_id = :13, sw_application_id = :14, use_owner_ba_id = :15, use_rule_desc = :16, row_changed_by = :17, row_changed_date = :18, row_created_by = :19, row_effective_date = :20, row_expiry_date = :21, row_quality = :22 where system_id = :24")
	if err != nil {
		return err
	}

	ppdm_code_version_use.Row_changed_date = formatDate(ppdm_code_version_use.Row_changed_date)
	ppdm_code_version_use.Effective_date = formatDateString(ppdm_code_version_use.Effective_date)
	ppdm_code_version_use.Expiry_date = formatDateString(ppdm_code_version_use.Expiry_date)
	ppdm_code_version_use.Row_effective_date = formatDateString(ppdm_code_version_use.Row_effective_date)
	ppdm_code_version_use.Row_expiry_date = formatDateString(ppdm_code_version_use.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_code_version_use.Table_name, ppdm_code_version_use.Source, ppdm_code_version_use.Version_obs_no, ppdm_code_version_use.Use_obs_no, ppdm_code_version_use.Active_ind, ppdm_code_version_use.Effective_date, ppdm_code_version_use.Expiry_date, ppdm_code_version_use.Group_id, ppdm_code_version_use.Ppdm_guid, ppdm_code_version_use.Preferred_ind, ppdm_code_version_use.Procedure_id, ppdm_code_version_use.Remark, ppdm_code_version_use.Source_document_id, ppdm_code_version_use.Sw_application_id, ppdm_code_version_use.Use_owner_ba_id, ppdm_code_version_use.Use_rule_desc, ppdm_code_version_use.Row_changed_by, ppdm_code_version_use.Row_changed_date, ppdm_code_version_use.Row_created_by, ppdm_code_version_use.Row_effective_date, ppdm_code_version_use.Row_expiry_date, ppdm_code_version_use.Row_quality, ppdm_code_version_use.System_id)
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

func PatchPpdmCodeVersionUse(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_code_version_use set "
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

func DeletePpdmCodeVersionUse(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_code_version_use dto.Ppdm_code_version_use
	ppdm_code_version_use.System_id = id

	stmt, err := db.Prepare("delete from ppdm_code_version_use where system_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_code_version_use.System_id)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetContProvision(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cont_provision")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cont_provision

	for rows.Next() {
		var cont_provision dto.Cont_provision
		if err := rows.Scan(&cont_provision.Contract_id, &cont_provision.Provision_id, &cont_provision.Active_ind, &cont_provision.Clause_heading, &cont_provision.Clause_number, &cont_provision.Cont_provision_type, &cont_provision.Effective_date, &cont_provision.Expiry_date, &cont_provision.Modified_ind, &cont_provision.Ppdm_guid, &cont_provision.Provision_desc, &cont_provision.Remark, &cont_provision.Source, &cont_provision.Source_document_id, &cont_provision.Row_changed_by, &cont_provision.Row_changed_date, &cont_provision.Row_created_by, &cont_provision.Row_created_date, &cont_provision.Row_effective_date, &cont_provision.Row_expiry_date, &cont_provision.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cont_provision to result
		result = append(result, cont_provision)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetContProvision(c *fiber.Ctx) error {
	var cont_provision dto.Cont_provision

	setDefaults(&cont_provision)

	if err := json.Unmarshal(c.Body(), &cont_provision); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cont_provision values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	cont_provision.Row_created_date = formatDate(cont_provision.Row_created_date)
	cont_provision.Row_changed_date = formatDate(cont_provision.Row_changed_date)
	cont_provision.Effective_date = formatDateString(cont_provision.Effective_date)
	cont_provision.Expiry_date = formatDateString(cont_provision.Expiry_date)
	cont_provision.Row_effective_date = formatDateString(cont_provision.Row_effective_date)
	cont_provision.Row_expiry_date = formatDateString(cont_provision.Row_expiry_date)






	rows, err := stmt.Exec(cont_provision.Contract_id, cont_provision.Provision_id, cont_provision.Active_ind, cont_provision.Clause_heading, cont_provision.Clause_number, cont_provision.Cont_provision_type, cont_provision.Effective_date, cont_provision.Expiry_date, cont_provision.Modified_ind, cont_provision.Ppdm_guid, cont_provision.Provision_desc, cont_provision.Remark, cont_provision.Source, cont_provision.Source_document_id, cont_provision.Row_changed_by, cont_provision.Row_changed_date, cont_provision.Row_created_by, cont_provision.Row_created_date, cont_provision.Row_effective_date, cont_provision.Row_expiry_date, cont_provision.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateContProvision(c *fiber.Ctx) error {
	var cont_provision dto.Cont_provision

	setDefaults(&cont_provision)

	if err := json.Unmarshal(c.Body(), &cont_provision); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cont_provision.Contract_id = id

    if cont_provision.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cont_provision set provision_id = :1, active_ind = :2, clause_heading = :3, clause_number = :4, cont_provision_type = :5, effective_date = :6, expiry_date = :7, modified_ind = :8, ppdm_guid = :9, provision_desc = :10, remark = :11, source = :12, source_document_id = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where contract_id = :21")
	if err != nil {
		return err
	}

	cont_provision.Row_changed_date = formatDate(cont_provision.Row_changed_date)
	cont_provision.Effective_date = formatDateString(cont_provision.Effective_date)
	cont_provision.Expiry_date = formatDateString(cont_provision.Expiry_date)
	cont_provision.Row_effective_date = formatDateString(cont_provision.Row_effective_date)
	cont_provision.Row_expiry_date = formatDateString(cont_provision.Row_expiry_date)






	rows, err := stmt.Exec(cont_provision.Provision_id, cont_provision.Active_ind, cont_provision.Clause_heading, cont_provision.Clause_number, cont_provision.Cont_provision_type, cont_provision.Effective_date, cont_provision.Expiry_date, cont_provision.Modified_ind, cont_provision.Ppdm_guid, cont_provision.Provision_desc, cont_provision.Remark, cont_provision.Source, cont_provision.Source_document_id, cont_provision.Row_changed_by, cont_provision.Row_changed_date, cont_provision.Row_created_by, cont_provision.Row_effective_date, cont_provision.Row_expiry_date, cont_provision.Row_quality, cont_provision.Contract_id)
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

func PatchContProvision(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cont_provision set "
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
	query += " where contract_id = :id"

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

func DeleteContProvision(c *fiber.Ctx) error {
	id := c.Params("id")
	var cont_provision dto.Cont_provision
	cont_provision.Contract_id = id

	stmt, err := db.Prepare("delete from cont_provision where contract_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cont_provision.Contract_id)
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



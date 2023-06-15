package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetContJurisdiction(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cont_jurisdiction")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cont_jurisdiction

	for rows.Next() {
		var cont_jurisdiction dto.Cont_jurisdiction
		if err := rows.Scan(&cont_jurisdiction.Contract_id, &cont_jurisdiction.Jurisdiction, &cont_jurisdiction.Active_ind, &cont_jurisdiction.Effective_date, &cont_jurisdiction.Expiry_date, &cont_jurisdiction.Ppdm_guid, &cont_jurisdiction.Remark, &cont_jurisdiction.Source, &cont_jurisdiction.Row_changed_by, &cont_jurisdiction.Row_changed_date, &cont_jurisdiction.Row_created_by, &cont_jurisdiction.Row_created_date, &cont_jurisdiction.Row_effective_date, &cont_jurisdiction.Row_expiry_date, &cont_jurisdiction.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cont_jurisdiction to result
		result = append(result, cont_jurisdiction)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetContJurisdiction(c *fiber.Ctx) error {
	var cont_jurisdiction dto.Cont_jurisdiction

	setDefaults(&cont_jurisdiction)

	if err := json.Unmarshal(c.Body(), &cont_jurisdiction); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cont_jurisdiction values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15)")
	if err != nil {
		return err
	}
	cont_jurisdiction.Row_created_date = formatDate(cont_jurisdiction.Row_created_date)
	cont_jurisdiction.Row_changed_date = formatDate(cont_jurisdiction.Row_changed_date)
	cont_jurisdiction.Effective_date = formatDateString(cont_jurisdiction.Effective_date)
	cont_jurisdiction.Expiry_date = formatDateString(cont_jurisdiction.Expiry_date)
	cont_jurisdiction.Row_effective_date = formatDateString(cont_jurisdiction.Row_effective_date)
	cont_jurisdiction.Row_expiry_date = formatDateString(cont_jurisdiction.Row_expiry_date)






	rows, err := stmt.Exec(cont_jurisdiction.Contract_id, cont_jurisdiction.Jurisdiction, cont_jurisdiction.Active_ind, cont_jurisdiction.Effective_date, cont_jurisdiction.Expiry_date, cont_jurisdiction.Ppdm_guid, cont_jurisdiction.Remark, cont_jurisdiction.Source, cont_jurisdiction.Row_changed_by, cont_jurisdiction.Row_changed_date, cont_jurisdiction.Row_created_by, cont_jurisdiction.Row_created_date, cont_jurisdiction.Row_effective_date, cont_jurisdiction.Row_expiry_date, cont_jurisdiction.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateContJurisdiction(c *fiber.Ctx) error {
	var cont_jurisdiction dto.Cont_jurisdiction

	setDefaults(&cont_jurisdiction)

	if err := json.Unmarshal(c.Body(), &cont_jurisdiction); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cont_jurisdiction.Contract_id = id

    if cont_jurisdiction.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cont_jurisdiction set jurisdiction = :1, active_ind = :2, effective_date = :3, expiry_date = :4, ppdm_guid = :5, remark = :6, source = :7, row_changed_by = :8, row_changed_date = :9, row_created_by = :10, row_effective_date = :11, row_expiry_date = :12, row_quality = :13 where contract_id = :15")
	if err != nil {
		return err
	}

	cont_jurisdiction.Row_changed_date = formatDate(cont_jurisdiction.Row_changed_date)
	cont_jurisdiction.Effective_date = formatDateString(cont_jurisdiction.Effective_date)
	cont_jurisdiction.Expiry_date = formatDateString(cont_jurisdiction.Expiry_date)
	cont_jurisdiction.Row_effective_date = formatDateString(cont_jurisdiction.Row_effective_date)
	cont_jurisdiction.Row_expiry_date = formatDateString(cont_jurisdiction.Row_expiry_date)






	rows, err := stmt.Exec(cont_jurisdiction.Jurisdiction, cont_jurisdiction.Active_ind, cont_jurisdiction.Effective_date, cont_jurisdiction.Expiry_date, cont_jurisdiction.Ppdm_guid, cont_jurisdiction.Remark, cont_jurisdiction.Source, cont_jurisdiction.Row_changed_by, cont_jurisdiction.Row_changed_date, cont_jurisdiction.Row_created_by, cont_jurisdiction.Row_effective_date, cont_jurisdiction.Row_expiry_date, cont_jurisdiction.Row_quality, cont_jurisdiction.Contract_id)
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

func PatchContJurisdiction(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cont_jurisdiction set "
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

func DeleteContJurisdiction(c *fiber.Ctx) error {
	id := c.Params("id")
	var cont_jurisdiction dto.Cont_jurisdiction
	cont_jurisdiction.Contract_id = id

	stmt, err := db.Prepare("delete from cont_jurisdiction where contract_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cont_jurisdiction.Contract_id)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetContStatus(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cont_status")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cont_status

	for rows.Next() {
		var cont_status dto.Cont_status
		if err := rows.Scan(&cont_status.Contract_id, &cont_status.Status_id, &cont_status.Active_ind, &cont_status.Contract_status, &cont_status.Contract_status_type, &cont_status.Effective_date, &cont_status.Expiry_date, &cont_status.Ppdm_guid, &cont_status.Remark, &cont_status.Source, &cont_status.Row_changed_by, &cont_status.Row_changed_date, &cont_status.Row_created_by, &cont_status.Row_created_date, &cont_status.Row_effective_date, &cont_status.Row_expiry_date, &cont_status.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cont_status to result
		result = append(result, cont_status)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetContStatus(c *fiber.Ctx) error {
	var cont_status dto.Cont_status

	setDefaults(&cont_status)

	if err := json.Unmarshal(c.Body(), &cont_status); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cont_status values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	cont_status.Row_created_date = formatDate(cont_status.Row_created_date)
	cont_status.Row_changed_date = formatDate(cont_status.Row_changed_date)
	cont_status.Effective_date = formatDateString(cont_status.Effective_date)
	cont_status.Expiry_date = formatDateString(cont_status.Expiry_date)
	cont_status.Row_effective_date = formatDateString(cont_status.Row_effective_date)
	cont_status.Row_expiry_date = formatDateString(cont_status.Row_expiry_date)






	rows, err := stmt.Exec(cont_status.Contract_id, cont_status.Status_id, cont_status.Active_ind, cont_status.Contract_status, cont_status.Contract_status_type, cont_status.Effective_date, cont_status.Expiry_date, cont_status.Ppdm_guid, cont_status.Remark, cont_status.Source, cont_status.Row_changed_by, cont_status.Row_changed_date, cont_status.Row_created_by, cont_status.Row_created_date, cont_status.Row_effective_date, cont_status.Row_expiry_date, cont_status.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateContStatus(c *fiber.Ctx) error {
	var cont_status dto.Cont_status

	setDefaults(&cont_status)

	if err := json.Unmarshal(c.Body(), &cont_status); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cont_status.Contract_id = id

    if cont_status.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cont_status set status_id = :1, active_ind = :2, contract_status = :3, contract_status_type = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where contract_id = :17")
	if err != nil {
		return err
	}

	cont_status.Row_changed_date = formatDate(cont_status.Row_changed_date)
	cont_status.Effective_date = formatDateString(cont_status.Effective_date)
	cont_status.Expiry_date = formatDateString(cont_status.Expiry_date)
	cont_status.Row_effective_date = formatDateString(cont_status.Row_effective_date)
	cont_status.Row_expiry_date = formatDateString(cont_status.Row_expiry_date)






	rows, err := stmt.Exec(cont_status.Status_id, cont_status.Active_ind, cont_status.Contract_status, cont_status.Contract_status_type, cont_status.Effective_date, cont_status.Expiry_date, cont_status.Ppdm_guid, cont_status.Remark, cont_status.Source, cont_status.Row_changed_by, cont_status.Row_changed_date, cont_status.Row_created_by, cont_status.Row_effective_date, cont_status.Row_expiry_date, cont_status.Row_quality, cont_status.Contract_id)
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

func PatchContStatus(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cont_status set "
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

func DeleteContStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var cont_status dto.Cont_status
	cont_status.Contract_id = id

	stmt, err := db.Prepare("delete from cont_status where contract_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cont_status.Contract_id)
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



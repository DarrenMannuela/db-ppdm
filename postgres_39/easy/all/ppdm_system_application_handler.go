package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmSystemApplication(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_system_application")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_system_application

	for rows.Next() {
		var ppdm_system_application dto.Ppdm_system_application
		if err := rows.Scan(&ppdm_system_application.System_id, &ppdm_system_application.Sw_application_id, &ppdm_system_application.Application_seq_no, &ppdm_system_application.Active_ind, &ppdm_system_application.Application_function_desc, &ppdm_system_application.Effective_date, &ppdm_system_application.Expiry_date, &ppdm_system_application.System_id, &ppdm_system_application.Ppdm_guid, &ppdm_system_application.Remark, &ppdm_system_application.Source, &ppdm_system_application.Row_changed_by, &ppdm_system_application.Row_changed_date, &ppdm_system_application.Row_created_by, &ppdm_system_application.Row_created_date, &ppdm_system_application.Row_effective_date, &ppdm_system_application.Row_expiry_date, &ppdm_system_application.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_system_application to result
		result = append(result, ppdm_system_application)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmSystemApplication(c *fiber.Ctx) error {
	var ppdm_system_application dto.Ppdm_system_application

	setDefaults(&ppdm_system_application)

	if err := json.Unmarshal(c.Body(), &ppdm_system_application); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_system_application values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	ppdm_system_application.Row_created_date = formatDate(ppdm_system_application.Row_created_date)
	ppdm_system_application.Row_changed_date = formatDate(ppdm_system_application.Row_changed_date)
	ppdm_system_application.Effective_date = formatDateString(ppdm_system_application.Effective_date)
	ppdm_system_application.Expiry_date = formatDateString(ppdm_system_application.Expiry_date)
	ppdm_system_application.Row_effective_date = formatDateString(ppdm_system_application.Row_effective_date)
	ppdm_system_application.Row_expiry_date = formatDateString(ppdm_system_application.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_system_application.System_id, ppdm_system_application.Sw_application_id, ppdm_system_application.Application_seq_no, ppdm_system_application.Active_ind, ppdm_system_application.Application_function_desc, ppdm_system_application.Effective_date, ppdm_system_application.Expiry_date, ppdm_system_application.System_id, ppdm_system_application.Ppdm_guid, ppdm_system_application.Remark, ppdm_system_application.Source, ppdm_system_application.Row_changed_by, ppdm_system_application.Row_changed_date, ppdm_system_application.Row_created_by, ppdm_system_application.Row_created_date, ppdm_system_application.Row_effective_date, ppdm_system_application.Row_expiry_date, ppdm_system_application.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmSystemApplication(c *fiber.Ctx) error {
	var ppdm_system_application dto.Ppdm_system_application

	setDefaults(&ppdm_system_application)

	if err := json.Unmarshal(c.Body(), &ppdm_system_application); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_system_application.System_id = id

    if ppdm_system_application.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_system_application set sw_application_id = :1, application_seq_no = :2, active_ind = :3, application_function_desc = :4, effective_date = :5, expiry_date = :6, system_id = :7, ppdm_guid = :8, remark = :9, source = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where system_id = :18")
	if err != nil {
		return err
	}

	ppdm_system_application.Row_changed_date = formatDate(ppdm_system_application.Row_changed_date)
	ppdm_system_application.Effective_date = formatDateString(ppdm_system_application.Effective_date)
	ppdm_system_application.Expiry_date = formatDateString(ppdm_system_application.Expiry_date)
	ppdm_system_application.Row_effective_date = formatDateString(ppdm_system_application.Row_effective_date)
	ppdm_system_application.Row_expiry_date = formatDateString(ppdm_system_application.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_system_application.Sw_application_id, ppdm_system_application.Application_seq_no, ppdm_system_application.Active_ind, ppdm_system_application.Application_function_desc, ppdm_system_application.Effective_date, ppdm_system_application.Expiry_date, ppdm_system_application.System_id, ppdm_system_application.Ppdm_guid, ppdm_system_application.Remark, ppdm_system_application.Source, ppdm_system_application.Row_changed_by, ppdm_system_application.Row_changed_date, ppdm_system_application.Row_created_by, ppdm_system_application.Row_effective_date, ppdm_system_application.Row_expiry_date, ppdm_system_application.Row_quality, ppdm_system_application.System_id)
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

func PatchPpdmSystemApplication(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_system_application set "
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

func DeletePpdmSystemApplication(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_system_application dto.Ppdm_system_application
	ppdm_system_application.System_id = id

	stmt, err := db.Prepare("delete from ppdm_system_application where system_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_system_application.System_id)
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



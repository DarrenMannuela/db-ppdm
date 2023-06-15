package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmCircProcess(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_circ_process")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_circ_process

	for rows.Next() {
		var rm_circ_process dto.Rm_circ_process
		if err := rows.Scan(&rm_circ_process.Circ_id, &rm_circ_process.Data_circ_process_id, &rm_circ_process.Active_ind, &rm_circ_process.Data_circ_process, &rm_circ_process.Data_circ_process_date, &rm_circ_process.Effective_date, &rm_circ_process.Expiry_date, &rm_circ_process.Ppdm_guid, &rm_circ_process.Remark, &rm_circ_process.Source, &rm_circ_process.Row_changed_by, &rm_circ_process.Row_changed_date, &rm_circ_process.Row_created_by, &rm_circ_process.Row_created_date, &rm_circ_process.Row_effective_date, &rm_circ_process.Row_expiry_date, &rm_circ_process.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_circ_process to result
		result = append(result, rm_circ_process)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmCircProcess(c *fiber.Ctx) error {
	var rm_circ_process dto.Rm_circ_process

	setDefaults(&rm_circ_process)

	if err := json.Unmarshal(c.Body(), &rm_circ_process); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_circ_process values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	rm_circ_process.Row_created_date = formatDate(rm_circ_process.Row_created_date)
	rm_circ_process.Row_changed_date = formatDate(rm_circ_process.Row_changed_date)
	rm_circ_process.Data_circ_process_date = formatDateString(rm_circ_process.Data_circ_process_date)
	rm_circ_process.Effective_date = formatDateString(rm_circ_process.Effective_date)
	rm_circ_process.Expiry_date = formatDateString(rm_circ_process.Expiry_date)
	rm_circ_process.Row_effective_date = formatDateString(rm_circ_process.Row_effective_date)
	rm_circ_process.Row_expiry_date = formatDateString(rm_circ_process.Row_expiry_date)






	rows, err := stmt.Exec(rm_circ_process.Circ_id, rm_circ_process.Data_circ_process_id, rm_circ_process.Active_ind, rm_circ_process.Data_circ_process, rm_circ_process.Data_circ_process_date, rm_circ_process.Effective_date, rm_circ_process.Expiry_date, rm_circ_process.Ppdm_guid, rm_circ_process.Remark, rm_circ_process.Source, rm_circ_process.Row_changed_by, rm_circ_process.Row_changed_date, rm_circ_process.Row_created_by, rm_circ_process.Row_created_date, rm_circ_process.Row_effective_date, rm_circ_process.Row_expiry_date, rm_circ_process.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmCircProcess(c *fiber.Ctx) error {
	var rm_circ_process dto.Rm_circ_process

	setDefaults(&rm_circ_process)

	if err := json.Unmarshal(c.Body(), &rm_circ_process); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_circ_process.Circ_id = id

    if rm_circ_process.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_circ_process set data_circ_process_id = :1, active_ind = :2, data_circ_process = :3, data_circ_process_date = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where circ_id = :17")
	if err != nil {
		return err
	}

	rm_circ_process.Row_changed_date = formatDate(rm_circ_process.Row_changed_date)
	rm_circ_process.Data_circ_process_date = formatDateString(rm_circ_process.Data_circ_process_date)
	rm_circ_process.Effective_date = formatDateString(rm_circ_process.Effective_date)
	rm_circ_process.Expiry_date = formatDateString(rm_circ_process.Expiry_date)
	rm_circ_process.Row_effective_date = formatDateString(rm_circ_process.Row_effective_date)
	rm_circ_process.Row_expiry_date = formatDateString(rm_circ_process.Row_expiry_date)






	rows, err := stmt.Exec(rm_circ_process.Data_circ_process_id, rm_circ_process.Active_ind, rm_circ_process.Data_circ_process, rm_circ_process.Data_circ_process_date, rm_circ_process.Effective_date, rm_circ_process.Expiry_date, rm_circ_process.Ppdm_guid, rm_circ_process.Remark, rm_circ_process.Source, rm_circ_process.Row_changed_by, rm_circ_process.Row_changed_date, rm_circ_process.Row_created_by, rm_circ_process.Row_effective_date, rm_circ_process.Row_expiry_date, rm_circ_process.Row_quality, rm_circ_process.Circ_id)
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

func PatchRmCircProcess(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_circ_process set "
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
		} else if key == "data_circ_process_date" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where circ_id = :id"

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

func DeleteRmCircProcess(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_circ_process dto.Rm_circ_process
	rm_circ_process.Circ_id = id

	stmt, err := db.Prepare("delete from rm_circ_process where circ_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_circ_process.Circ_id)
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



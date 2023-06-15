package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRInformationProcess(c *fiber.Ctx) error {
	rows, err := db.Query("select * from r_information_process")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.R_information_process

	for rows.Next() {
		var r_information_process dto.R_information_process
		if err := rows.Scan(&r_information_process.Information_process, &r_information_process.Abbreviation, &r_information_process.Active_ind, &r_information_process.Effective_date, &r_information_process.Expiry_date, &r_information_process.Long_name, &r_information_process.Ppdm_guid, &r_information_process.Remark, &r_information_process.Short_name, &r_information_process.Source, &r_information_process.Row_changed_by, &r_information_process.Row_changed_date, &r_information_process.Row_created_by, &r_information_process.Row_created_date, &r_information_process.Row_effective_date, &r_information_process.Row_expiry_date, &r_information_process.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append r_information_process to result
		result = append(result, r_information_process)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRInformationProcess(c *fiber.Ctx) error {
	var r_information_process dto.R_information_process

	setDefaults(&r_information_process)

	if err := json.Unmarshal(c.Body(), &r_information_process); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into r_information_process values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	r_information_process.Row_created_date = formatDate(r_information_process.Row_created_date)
	r_information_process.Row_changed_date = formatDate(r_information_process.Row_changed_date)
	r_information_process.Effective_date = formatDateString(r_information_process.Effective_date)
	r_information_process.Expiry_date = formatDateString(r_information_process.Expiry_date)
	r_information_process.Row_effective_date = formatDateString(r_information_process.Row_effective_date)
	r_information_process.Row_expiry_date = formatDateString(r_information_process.Row_expiry_date)






	rows, err := stmt.Exec(r_information_process.Information_process, r_information_process.Abbreviation, r_information_process.Active_ind, r_information_process.Effective_date, r_information_process.Expiry_date, r_information_process.Long_name, r_information_process.Ppdm_guid, r_information_process.Remark, r_information_process.Short_name, r_information_process.Source, r_information_process.Row_changed_by, r_information_process.Row_changed_date, r_information_process.Row_created_by, r_information_process.Row_created_date, r_information_process.Row_effective_date, r_information_process.Row_expiry_date, r_information_process.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRInformationProcess(c *fiber.Ctx) error {
	var r_information_process dto.R_information_process

	setDefaults(&r_information_process)

	if err := json.Unmarshal(c.Body(), &r_information_process); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	r_information_process.Information_process = id

    if r_information_process.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update r_information_process set abbreviation = :1, active_ind = :2, effective_date = :3, expiry_date = :4, long_name = :5, ppdm_guid = :6, remark = :7, short_name = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where information_process = :17")
	if err != nil {
		return err
	}

	r_information_process.Row_changed_date = formatDate(r_information_process.Row_changed_date)
	r_information_process.Effective_date = formatDateString(r_information_process.Effective_date)
	r_information_process.Expiry_date = formatDateString(r_information_process.Expiry_date)
	r_information_process.Row_effective_date = formatDateString(r_information_process.Row_effective_date)
	r_information_process.Row_expiry_date = formatDateString(r_information_process.Row_expiry_date)






	rows, err := stmt.Exec(r_information_process.Abbreviation, r_information_process.Active_ind, r_information_process.Effective_date, r_information_process.Expiry_date, r_information_process.Long_name, r_information_process.Ppdm_guid, r_information_process.Remark, r_information_process.Short_name, r_information_process.Source, r_information_process.Row_changed_by, r_information_process.Row_changed_date, r_information_process.Row_created_by, r_information_process.Row_effective_date, r_information_process.Row_expiry_date, r_information_process.Row_quality, r_information_process.Information_process)
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

func PatchRInformationProcess(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update r_information_process set "
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
	query += " where information_process = :id"

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

func DeleteRInformationProcess(c *fiber.Ctx) error {
	id := c.Params("id")
	var r_information_process dto.R_information_process
	r_information_process.Information_process = id

	stmt, err := db.Prepare("delete from r_information_process where information_process = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(r_information_process.Information_process)
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



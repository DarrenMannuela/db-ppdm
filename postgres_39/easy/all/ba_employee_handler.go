package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetBaEmployee(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ba_employee")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ba_employee

	for rows.Next() {
		var ba_employee dto.Ba_employee
		if err := rows.Scan(&ba_employee.Employer_ba_id, &ba_employee.Employee_ba_id, &ba_employee.Employee_obs_no, &ba_employee.Active_ind, &ba_employee.Effective_date, &ba_employee.Employee_position, &ba_employee.Expiry_date, &ba_employee.Ppdm_guid, &ba_employee.Remark, &ba_employee.Source, &ba_employee.Status, &ba_employee.Row_changed_by, &ba_employee.Row_changed_date, &ba_employee.Row_created_by, &ba_employee.Row_created_date, &ba_employee.Row_effective_date, &ba_employee.Row_expiry_date, &ba_employee.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ba_employee to result
		result = append(result, ba_employee)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetBaEmployee(c *fiber.Ctx) error {
	var ba_employee dto.Ba_employee

	setDefaults(&ba_employee)

	if err := json.Unmarshal(c.Body(), &ba_employee); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ba_employee values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	ba_employee.Row_created_date = formatDate(ba_employee.Row_created_date)
	ba_employee.Row_changed_date = formatDate(ba_employee.Row_changed_date)
	ba_employee.Effective_date = formatDateString(ba_employee.Effective_date)
	ba_employee.Expiry_date = formatDateString(ba_employee.Expiry_date)
	ba_employee.Row_effective_date = formatDateString(ba_employee.Row_effective_date)
	ba_employee.Row_expiry_date = formatDateString(ba_employee.Row_expiry_date)






	rows, err := stmt.Exec(ba_employee.Employer_ba_id, ba_employee.Employee_ba_id, ba_employee.Employee_obs_no, ba_employee.Active_ind, ba_employee.Effective_date, ba_employee.Employee_position, ba_employee.Expiry_date, ba_employee.Ppdm_guid, ba_employee.Remark, ba_employee.Source, ba_employee.Status, ba_employee.Row_changed_by, ba_employee.Row_changed_date, ba_employee.Row_created_by, ba_employee.Row_created_date, ba_employee.Row_effective_date, ba_employee.Row_expiry_date, ba_employee.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateBaEmployee(c *fiber.Ctx) error {
	var ba_employee dto.Ba_employee

	setDefaults(&ba_employee)

	if err := json.Unmarshal(c.Body(), &ba_employee); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ba_employee.Employer_ba_id = id

    if ba_employee.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ba_employee set employee_ba_id = :1, employee_obs_no = :2, active_ind = :3, effective_date = :4, employee_position = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, source = :9, status = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where employer_ba_id = :18")
	if err != nil {
		return err
	}

	ba_employee.Row_changed_date = formatDate(ba_employee.Row_changed_date)
	ba_employee.Effective_date = formatDateString(ba_employee.Effective_date)
	ba_employee.Expiry_date = formatDateString(ba_employee.Expiry_date)
	ba_employee.Row_effective_date = formatDateString(ba_employee.Row_effective_date)
	ba_employee.Row_expiry_date = formatDateString(ba_employee.Row_expiry_date)






	rows, err := stmt.Exec(ba_employee.Employee_ba_id, ba_employee.Employee_obs_no, ba_employee.Active_ind, ba_employee.Effective_date, ba_employee.Employee_position, ba_employee.Expiry_date, ba_employee.Ppdm_guid, ba_employee.Remark, ba_employee.Source, ba_employee.Status, ba_employee.Row_changed_by, ba_employee.Row_changed_date, ba_employee.Row_created_by, ba_employee.Row_effective_date, ba_employee.Row_expiry_date, ba_employee.Row_quality, ba_employee.Employer_ba_id)
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

func PatchBaEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ba_employee set "
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
	query += " where employer_ba_id = :id"

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

func DeleteBaEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	var ba_employee dto.Ba_employee
	ba_employee.Employer_ba_id = id

	stmt, err := db.Prepare("delete from ba_employee where employer_ba_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ba_employee.Employer_ba_id)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetInterestSet(c *fiber.Ctx) error {
	rows, err := db.Query("select * from interest_set")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Interest_set

	for rows.Next() {
		var interest_set dto.Interest_set
		if err := rows.Scan(&interest_set.Interest_set_id, &interest_set.Interest_set_seq_no, &interest_set.Active_ind, &interest_set.Effective_date, &interest_set.Expiry_date, &interest_set.Interest_set_type, &interest_set.Ppdm_guid, &interest_set.Remark, &interest_set.Source, &interest_set.Row_changed_by, &interest_set.Row_changed_date, &interest_set.Row_created_by, &interest_set.Row_created_date, &interest_set.Row_effective_date, &interest_set.Row_expiry_date, &interest_set.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append interest_set to result
		result = append(result, interest_set)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetInterestSet(c *fiber.Ctx) error {
	var interest_set dto.Interest_set

	setDefaults(&interest_set)

	if err := json.Unmarshal(c.Body(), &interest_set); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into interest_set values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16)")
	if err != nil {
		return err
	}
	interest_set.Row_created_date = formatDate(interest_set.Row_created_date)
	interest_set.Row_changed_date = formatDate(interest_set.Row_changed_date)
	interest_set.Effective_date = formatDateString(interest_set.Effective_date)
	interest_set.Expiry_date = formatDateString(interest_set.Expiry_date)
	interest_set.Row_effective_date = formatDateString(interest_set.Row_effective_date)
	interest_set.Row_expiry_date = formatDateString(interest_set.Row_expiry_date)






	rows, err := stmt.Exec(interest_set.Interest_set_id, interest_set.Interest_set_seq_no, interest_set.Active_ind, interest_set.Effective_date, interest_set.Expiry_date, interest_set.Interest_set_type, interest_set.Ppdm_guid, interest_set.Remark, interest_set.Source, interest_set.Row_changed_by, interest_set.Row_changed_date, interest_set.Row_created_by, interest_set.Row_created_date, interest_set.Row_effective_date, interest_set.Row_expiry_date, interest_set.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateInterestSet(c *fiber.Ctx) error {
	var interest_set dto.Interest_set

	setDefaults(&interest_set)

	if err := json.Unmarshal(c.Body(), &interest_set); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	interest_set.Interest_set_id = id

    if interest_set.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update interest_set set interest_set_seq_no = :1, active_ind = :2, effective_date = :3, expiry_date = :4, interest_set_type = :5, ppdm_guid = :6, remark = :7, source = :8, row_changed_by = :9, row_changed_date = :10, row_created_by = :11, row_effective_date = :12, row_expiry_date = :13, row_quality = :14 where interest_set_id = :16")
	if err != nil {
		return err
	}

	interest_set.Row_changed_date = formatDate(interest_set.Row_changed_date)
	interest_set.Effective_date = formatDateString(interest_set.Effective_date)
	interest_set.Expiry_date = formatDateString(interest_set.Expiry_date)
	interest_set.Row_effective_date = formatDateString(interest_set.Row_effective_date)
	interest_set.Row_expiry_date = formatDateString(interest_set.Row_expiry_date)






	rows, err := stmt.Exec(interest_set.Interest_set_seq_no, interest_set.Active_ind, interest_set.Effective_date, interest_set.Expiry_date, interest_set.Interest_set_type, interest_set.Ppdm_guid, interest_set.Remark, interest_set.Source, interest_set.Row_changed_by, interest_set.Row_changed_date, interest_set.Row_created_by, interest_set.Row_effective_date, interest_set.Row_expiry_date, interest_set.Row_quality, interest_set.Interest_set_id)
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

func PatchInterestSet(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update interest_set set "
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
	query += " where interest_set_id = :id"

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

func DeleteInterestSet(c *fiber.Ctx) error {
	id := c.Params("id")
	var interest_set dto.Interest_set
	interest_set.Interest_set_id = id

	stmt, err := db.Prepare("delete from interest_set where interest_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(interest_set.Interest_set_id)
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



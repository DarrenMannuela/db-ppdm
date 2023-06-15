package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRMaceralAmountType(c *fiber.Ctx) error {
	rows, err := db.Query("select * from r_maceral_amount_type")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.R_maceral_amount_type

	for rows.Next() {
		var r_maceral_amount_type dto.R_maceral_amount_type
		if err := rows.Scan(&r_maceral_amount_type.Maceral_amount_type, &r_maceral_amount_type.Active_ind, &r_maceral_amount_type.Effective_date, &r_maceral_amount_type.Expiry_date, &r_maceral_amount_type.Ppdm_guid, &r_maceral_amount_type.Remark, &r_maceral_amount_type.Source, &r_maceral_amount_type.Row_changed_by, &r_maceral_amount_type.Row_changed_date, &r_maceral_amount_type.Row_created_by, &r_maceral_amount_type.Row_created_date, &r_maceral_amount_type.Row_effective_date, &r_maceral_amount_type.Row_expiry_date, &r_maceral_amount_type.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append r_maceral_amount_type to result
		result = append(result, r_maceral_amount_type)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRMaceralAmountType(c *fiber.Ctx) error {
	var r_maceral_amount_type dto.R_maceral_amount_type

	setDefaults(&r_maceral_amount_type)

	if err := json.Unmarshal(c.Body(), &r_maceral_amount_type); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into r_maceral_amount_type values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14)")
	if err != nil {
		return err
	}
	r_maceral_amount_type.Row_created_date = formatDate(r_maceral_amount_type.Row_created_date)
	r_maceral_amount_type.Row_changed_date = formatDate(r_maceral_amount_type.Row_changed_date)
	r_maceral_amount_type.Effective_date = formatDateString(r_maceral_amount_type.Effective_date)
	r_maceral_amount_type.Expiry_date = formatDateString(r_maceral_amount_type.Expiry_date)
	r_maceral_amount_type.Row_effective_date = formatDateString(r_maceral_amount_type.Row_effective_date)
	r_maceral_amount_type.Row_expiry_date = formatDateString(r_maceral_amount_type.Row_expiry_date)






	rows, err := stmt.Exec(r_maceral_amount_type.Maceral_amount_type, r_maceral_amount_type.Active_ind, r_maceral_amount_type.Effective_date, r_maceral_amount_type.Expiry_date, r_maceral_amount_type.Ppdm_guid, r_maceral_amount_type.Remark, r_maceral_amount_type.Source, r_maceral_amount_type.Row_changed_by, r_maceral_amount_type.Row_changed_date, r_maceral_amount_type.Row_created_by, r_maceral_amount_type.Row_created_date, r_maceral_amount_type.Row_effective_date, r_maceral_amount_type.Row_expiry_date, r_maceral_amount_type.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRMaceralAmountType(c *fiber.Ctx) error {
	var r_maceral_amount_type dto.R_maceral_amount_type

	setDefaults(&r_maceral_amount_type)

	if err := json.Unmarshal(c.Body(), &r_maceral_amount_type); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	r_maceral_amount_type.Maceral_amount_type = id

    if r_maceral_amount_type.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update r_maceral_amount_type set active_ind = :1, effective_date = :2, expiry_date = :3, ppdm_guid = :4, remark = :5, source = :6, row_changed_by = :7, row_changed_date = :8, row_created_by = :9, row_effective_date = :10, row_expiry_date = :11, row_quality = :12 where maceral_amount_type = :14")
	if err != nil {
		return err
	}

	r_maceral_amount_type.Row_changed_date = formatDate(r_maceral_amount_type.Row_changed_date)
	r_maceral_amount_type.Effective_date = formatDateString(r_maceral_amount_type.Effective_date)
	r_maceral_amount_type.Expiry_date = formatDateString(r_maceral_amount_type.Expiry_date)
	r_maceral_amount_type.Row_effective_date = formatDateString(r_maceral_amount_type.Row_effective_date)
	r_maceral_amount_type.Row_expiry_date = formatDateString(r_maceral_amount_type.Row_expiry_date)






	rows, err := stmt.Exec(r_maceral_amount_type.Active_ind, r_maceral_amount_type.Effective_date, r_maceral_amount_type.Expiry_date, r_maceral_amount_type.Ppdm_guid, r_maceral_amount_type.Remark, r_maceral_amount_type.Source, r_maceral_amount_type.Row_changed_by, r_maceral_amount_type.Row_changed_date, r_maceral_amount_type.Row_created_by, r_maceral_amount_type.Row_effective_date, r_maceral_amount_type.Row_expiry_date, r_maceral_amount_type.Row_quality, r_maceral_amount_type.Maceral_amount_type)
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

func PatchRMaceralAmountType(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update r_maceral_amount_type set "
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
	query += " where maceral_amount_type = :id"

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

func DeleteRMaceralAmountType(c *fiber.Ctx) error {
	id := c.Params("id")
	var r_maceral_amount_type dto.R_maceral_amount_type
	r_maceral_amount_type.Maceral_amount_type = id

	stmt, err := db.Prepare("delete from r_maceral_amount_type where maceral_amount_type = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(r_maceral_amount_type.Maceral_amount_type)
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



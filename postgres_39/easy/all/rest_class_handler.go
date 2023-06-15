package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRestClass(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rest_class")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rest_class

	for rows.Next() {
		var rest_class dto.Rest_class
		if err := rows.Scan(&rest_class.Restriction_class_id, &rest_class.Active_ind, &rest_class.Effective_date, &rest_class.Expiry_date, &rest_class.Ppdm_guid, &rest_class.Remark, &rest_class.Rest_class_name, &rest_class.Source, &rest_class.Row_changed_by, &rest_class.Row_changed_date, &rest_class.Row_created_by, &rest_class.Row_created_date, &rest_class.Row_effective_date, &rest_class.Row_expiry_date, &rest_class.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rest_class to result
		result = append(result, rest_class)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRestClass(c *fiber.Ctx) error {
	var rest_class dto.Rest_class

	setDefaults(&rest_class)

	if err := json.Unmarshal(c.Body(), &rest_class); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rest_class values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15)")
	if err != nil {
		return err
	}
	rest_class.Row_created_date = formatDate(rest_class.Row_created_date)
	rest_class.Row_changed_date = formatDate(rest_class.Row_changed_date)
	rest_class.Effective_date = formatDateString(rest_class.Effective_date)
	rest_class.Expiry_date = formatDateString(rest_class.Expiry_date)
	rest_class.Row_effective_date = formatDateString(rest_class.Row_effective_date)
	rest_class.Row_expiry_date = formatDateString(rest_class.Row_expiry_date)






	rows, err := stmt.Exec(rest_class.Restriction_class_id, rest_class.Active_ind, rest_class.Effective_date, rest_class.Expiry_date, rest_class.Ppdm_guid, rest_class.Remark, rest_class.Rest_class_name, rest_class.Source, rest_class.Row_changed_by, rest_class.Row_changed_date, rest_class.Row_created_by, rest_class.Row_created_date, rest_class.Row_effective_date, rest_class.Row_expiry_date, rest_class.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRestClass(c *fiber.Ctx) error {
	var rest_class dto.Rest_class

	setDefaults(&rest_class)

	if err := json.Unmarshal(c.Body(), &rest_class); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rest_class.Restriction_class_id = id

    if rest_class.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rest_class set active_ind = :1, effective_date = :2, expiry_date = :3, ppdm_guid = :4, remark = :5, rest_class_name = :6, source = :7, row_changed_by = :8, row_changed_date = :9, row_created_by = :10, row_effective_date = :11, row_expiry_date = :12, row_quality = :13 where restriction_class_id = :15")
	if err != nil {
		return err
	}

	rest_class.Row_changed_date = formatDate(rest_class.Row_changed_date)
	rest_class.Effective_date = formatDateString(rest_class.Effective_date)
	rest_class.Expiry_date = formatDateString(rest_class.Expiry_date)
	rest_class.Row_effective_date = formatDateString(rest_class.Row_effective_date)
	rest_class.Row_expiry_date = formatDateString(rest_class.Row_expiry_date)






	rows, err := stmt.Exec(rest_class.Active_ind, rest_class.Effective_date, rest_class.Expiry_date, rest_class.Ppdm_guid, rest_class.Remark, rest_class.Rest_class_name, rest_class.Source, rest_class.Row_changed_by, rest_class.Row_changed_date, rest_class.Row_created_by, rest_class.Row_effective_date, rest_class.Row_expiry_date, rest_class.Row_quality, rest_class.Restriction_class_id)
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

func PatchRestClass(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rest_class set "
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
	query += " where restriction_class_id = :id"

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

func DeleteRestClass(c *fiber.Ctx) error {
	id := c.Params("id")
	var rest_class dto.Rest_class
	rest_class.Restriction_class_id = id

	stmt, err := db.Prepare("delete from rest_class where restriction_class_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rest_class.Restriction_class_id)
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



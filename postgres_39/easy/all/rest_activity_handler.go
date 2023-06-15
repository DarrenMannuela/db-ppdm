package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRestActivity(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rest_activity")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rest_activity

	for rows.Next() {
		var rest_activity dto.Rest_activity
		if err := rows.Scan(&rest_activity.Restriction_id, &rest_activity.Restriction_version, &rest_activity.Restricted_activity, &rest_activity.Active_ind, &rest_activity.Effective_date, &rest_activity.Expiry_date, &rest_activity.Ppdm_guid, &rest_activity.Remark, &rest_activity.Source, &rest_activity.Row_changed_by, &rest_activity.Row_changed_date, &rest_activity.Row_created_by, &rest_activity.Row_created_date, &rest_activity.Row_effective_date, &rest_activity.Row_expiry_date, &rest_activity.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rest_activity to result
		result = append(result, rest_activity)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRestActivity(c *fiber.Ctx) error {
	var rest_activity dto.Rest_activity

	setDefaults(&rest_activity)

	if err := json.Unmarshal(c.Body(), &rest_activity); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rest_activity values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16)")
	if err != nil {
		return err
	}
	rest_activity.Row_created_date = formatDate(rest_activity.Row_created_date)
	rest_activity.Row_changed_date = formatDate(rest_activity.Row_changed_date)
	rest_activity.Effective_date = formatDateString(rest_activity.Effective_date)
	rest_activity.Expiry_date = formatDateString(rest_activity.Expiry_date)
	rest_activity.Row_effective_date = formatDateString(rest_activity.Row_effective_date)
	rest_activity.Row_expiry_date = formatDateString(rest_activity.Row_expiry_date)






	rows, err := stmt.Exec(rest_activity.Restriction_id, rest_activity.Restriction_version, rest_activity.Restricted_activity, rest_activity.Active_ind, rest_activity.Effective_date, rest_activity.Expiry_date, rest_activity.Ppdm_guid, rest_activity.Remark, rest_activity.Source, rest_activity.Row_changed_by, rest_activity.Row_changed_date, rest_activity.Row_created_by, rest_activity.Row_created_date, rest_activity.Row_effective_date, rest_activity.Row_expiry_date, rest_activity.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRestActivity(c *fiber.Ctx) error {
	var rest_activity dto.Rest_activity

	setDefaults(&rest_activity)

	if err := json.Unmarshal(c.Body(), &rest_activity); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rest_activity.Restriction_id = id

    if rest_activity.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rest_activity set restriction_version = :1, restricted_activity = :2, active_ind = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, source = :8, row_changed_by = :9, row_changed_date = :10, row_created_by = :11, row_effective_date = :12, row_expiry_date = :13, row_quality = :14 where restriction_id = :16")
	if err != nil {
		return err
	}

	rest_activity.Row_changed_date = formatDate(rest_activity.Row_changed_date)
	rest_activity.Effective_date = formatDateString(rest_activity.Effective_date)
	rest_activity.Expiry_date = formatDateString(rest_activity.Expiry_date)
	rest_activity.Row_effective_date = formatDateString(rest_activity.Row_effective_date)
	rest_activity.Row_expiry_date = formatDateString(rest_activity.Row_expiry_date)






	rows, err := stmt.Exec(rest_activity.Restriction_version, rest_activity.Restricted_activity, rest_activity.Active_ind, rest_activity.Effective_date, rest_activity.Expiry_date, rest_activity.Ppdm_guid, rest_activity.Remark, rest_activity.Source, rest_activity.Row_changed_by, rest_activity.Row_changed_date, rest_activity.Row_created_by, rest_activity.Row_effective_date, rest_activity.Row_expiry_date, rest_activity.Row_quality, rest_activity.Restriction_id)
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

func PatchRestActivity(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rest_activity set "
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
	query += " where restriction_id = :id"

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

func DeleteRestActivity(c *fiber.Ctx) error {
	id := c.Params("id")
	var rest_activity dto.Rest_activity
	rest_activity.Restriction_id = id

	stmt, err := db.Prepare("delete from rest_activity where restriction_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rest_activity.Restriction_id)
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



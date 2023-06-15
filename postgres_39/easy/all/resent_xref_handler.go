package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetResentXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from resent_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Resent_xref

	for rows.Next() {
		var resent_xref dto.Resent_xref
		if err := rows.Scan(&resent_xref.Resent_id, &resent_xref.Resent_id2, &resent_xref.Xref_obs_no, &resent_xref.Active_ind, &resent_xref.Effective_date, &resent_xref.Expiry_date, &resent_xref.Ppdm_guid, &resent_xref.Remark, &resent_xref.Source, &resent_xref.Xref_type, &resent_xref.Row_changed_by, &resent_xref.Row_changed_date, &resent_xref.Row_created_by, &resent_xref.Row_created_date, &resent_xref.Row_effective_date, &resent_xref.Row_expiry_date, &resent_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append resent_xref to result
		result = append(result, resent_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetResentXref(c *fiber.Ctx) error {
	var resent_xref dto.Resent_xref

	setDefaults(&resent_xref)

	if err := json.Unmarshal(c.Body(), &resent_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into resent_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	resent_xref.Row_created_date = formatDate(resent_xref.Row_created_date)
	resent_xref.Row_changed_date = formatDate(resent_xref.Row_changed_date)
	resent_xref.Effective_date = formatDateString(resent_xref.Effective_date)
	resent_xref.Expiry_date = formatDateString(resent_xref.Expiry_date)
	resent_xref.Row_effective_date = formatDateString(resent_xref.Row_effective_date)
	resent_xref.Row_expiry_date = formatDateString(resent_xref.Row_expiry_date)






	rows, err := stmt.Exec(resent_xref.Resent_id, resent_xref.Resent_id2, resent_xref.Xref_obs_no, resent_xref.Active_ind, resent_xref.Effective_date, resent_xref.Expiry_date, resent_xref.Ppdm_guid, resent_xref.Remark, resent_xref.Source, resent_xref.Xref_type, resent_xref.Row_changed_by, resent_xref.Row_changed_date, resent_xref.Row_created_by, resent_xref.Row_created_date, resent_xref.Row_effective_date, resent_xref.Row_expiry_date, resent_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateResentXref(c *fiber.Ctx) error {
	var resent_xref dto.Resent_xref

	setDefaults(&resent_xref)

	if err := json.Unmarshal(c.Body(), &resent_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	resent_xref.Resent_id = id

    if resent_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update resent_xref set resent_id2 = :1, xref_obs_no = :2, active_ind = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, source = :8, xref_type = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where resent_id = :17")
	if err != nil {
		return err
	}

	resent_xref.Row_changed_date = formatDate(resent_xref.Row_changed_date)
	resent_xref.Effective_date = formatDateString(resent_xref.Effective_date)
	resent_xref.Expiry_date = formatDateString(resent_xref.Expiry_date)
	resent_xref.Row_effective_date = formatDateString(resent_xref.Row_effective_date)
	resent_xref.Row_expiry_date = formatDateString(resent_xref.Row_expiry_date)






	rows, err := stmt.Exec(resent_xref.Resent_id2, resent_xref.Xref_obs_no, resent_xref.Active_ind, resent_xref.Effective_date, resent_xref.Expiry_date, resent_xref.Ppdm_guid, resent_xref.Remark, resent_xref.Source, resent_xref.Xref_type, resent_xref.Row_changed_by, resent_xref.Row_changed_date, resent_xref.Row_created_by, resent_xref.Row_effective_date, resent_xref.Row_expiry_date, resent_xref.Row_quality, resent_xref.Resent_id)
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

func PatchResentXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update resent_xref set "
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
	query += " where resent_id = :id"

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

func DeleteResentXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var resent_xref dto.Resent_xref
	resent_xref.Resent_id = id

	stmt, err := db.Prepare("delete from resent_xref where resent_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(resent_xref.Resent_id)
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



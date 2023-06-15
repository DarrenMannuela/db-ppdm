package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlValidBa(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_valid_ba")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_valid_ba

	for rows.Next() {
		var anl_valid_ba dto.Anl_valid_ba
		if err := rows.Scan(&anl_valid_ba.Method_set_id, &anl_valid_ba.Method_id, &anl_valid_ba.Valid_ba_id, &anl_valid_ba.Valid_ba_obs_no, &anl_valid_ba.Active_ind, &anl_valid_ba.Ba_role_type, &anl_valid_ba.Confidence_type, &anl_valid_ba.Effective_date, &anl_valid_ba.Expiry_date, &anl_valid_ba.Ppdm_guid, &anl_valid_ba.Remark, &anl_valid_ba.Source, &anl_valid_ba.Row_changed_by, &anl_valid_ba.Row_changed_date, &anl_valid_ba.Row_created_by, &anl_valid_ba.Row_created_date, &anl_valid_ba.Row_effective_date, &anl_valid_ba.Row_expiry_date, &anl_valid_ba.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_valid_ba to result
		result = append(result, anl_valid_ba)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlValidBa(c *fiber.Ctx) error {
	var anl_valid_ba dto.Anl_valid_ba

	setDefaults(&anl_valid_ba)

	if err := json.Unmarshal(c.Body(), &anl_valid_ba); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_valid_ba values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	anl_valid_ba.Row_created_date = formatDate(anl_valid_ba.Row_created_date)
	anl_valid_ba.Row_changed_date = formatDate(anl_valid_ba.Row_changed_date)
	anl_valid_ba.Effective_date = formatDateString(anl_valid_ba.Effective_date)
	anl_valid_ba.Expiry_date = formatDateString(anl_valid_ba.Expiry_date)
	anl_valid_ba.Row_effective_date = formatDateString(anl_valid_ba.Row_effective_date)
	anl_valid_ba.Row_expiry_date = formatDateString(anl_valid_ba.Row_expiry_date)






	rows, err := stmt.Exec(anl_valid_ba.Method_set_id, anl_valid_ba.Method_id, anl_valid_ba.Valid_ba_id, anl_valid_ba.Valid_ba_obs_no, anl_valid_ba.Active_ind, anl_valid_ba.Ba_role_type, anl_valid_ba.Confidence_type, anl_valid_ba.Effective_date, anl_valid_ba.Expiry_date, anl_valid_ba.Ppdm_guid, anl_valid_ba.Remark, anl_valid_ba.Source, anl_valid_ba.Row_changed_by, anl_valid_ba.Row_changed_date, anl_valid_ba.Row_created_by, anl_valid_ba.Row_created_date, anl_valid_ba.Row_effective_date, anl_valid_ba.Row_expiry_date, anl_valid_ba.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlValidBa(c *fiber.Ctx) error {
	var anl_valid_ba dto.Anl_valid_ba

	setDefaults(&anl_valid_ba)

	if err := json.Unmarshal(c.Body(), &anl_valid_ba); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_valid_ba.Method_set_id = id

    if anl_valid_ba.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_valid_ba set method_id = :1, valid_ba_id = :2, valid_ba_obs_no = :3, active_ind = :4, ba_role_type = :5, confidence_type = :6, effective_date = :7, expiry_date = :8, ppdm_guid = :9, remark = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where method_set_id = :19")
	if err != nil {
		return err
	}

	anl_valid_ba.Row_changed_date = formatDate(anl_valid_ba.Row_changed_date)
	anl_valid_ba.Effective_date = formatDateString(anl_valid_ba.Effective_date)
	anl_valid_ba.Expiry_date = formatDateString(anl_valid_ba.Expiry_date)
	anl_valid_ba.Row_effective_date = formatDateString(anl_valid_ba.Row_effective_date)
	anl_valid_ba.Row_expiry_date = formatDateString(anl_valid_ba.Row_expiry_date)






	rows, err := stmt.Exec(anl_valid_ba.Method_id, anl_valid_ba.Valid_ba_id, anl_valid_ba.Valid_ba_obs_no, anl_valid_ba.Active_ind, anl_valid_ba.Ba_role_type, anl_valid_ba.Confidence_type, anl_valid_ba.Effective_date, anl_valid_ba.Expiry_date, anl_valid_ba.Ppdm_guid, anl_valid_ba.Remark, anl_valid_ba.Source, anl_valid_ba.Row_changed_by, anl_valid_ba.Row_changed_date, anl_valid_ba.Row_created_by, anl_valid_ba.Row_effective_date, anl_valid_ba.Row_expiry_date, anl_valid_ba.Row_quality, anl_valid_ba.Method_set_id)
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

func PatchAnlValidBa(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_valid_ba set "
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
	query += " where method_set_id = :id"

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

func DeleteAnlValidBa(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_valid_ba dto.Anl_valid_ba
	anl_valid_ba.Method_set_id = id

	stmt, err := db.Prepare("delete from anl_valid_ba where method_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_valid_ba.Method_set_id)
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



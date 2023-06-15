package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRVolumeMethod(c *fiber.Ctx) error {
	rows, err := db.Query("select * from r_volume_method")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.R_volume_method

	for rows.Next() {
		var r_volume_method dto.R_volume_method
		if err := rows.Scan(&r_volume_method.Volume_method, &r_volume_method.Abbreviation, &r_volume_method.Active_ind, &r_volume_method.Effective_date, &r_volume_method.Expiry_date, &r_volume_method.Long_name, &r_volume_method.Ppdm_guid, &r_volume_method.Remark, &r_volume_method.Short_name, &r_volume_method.Source, &r_volume_method.Row_changed_by, &r_volume_method.Row_changed_date, &r_volume_method.Row_created_by, &r_volume_method.Row_created_date, &r_volume_method.Row_effective_date, &r_volume_method.Row_expiry_date, &r_volume_method.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append r_volume_method to result
		result = append(result, r_volume_method)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRVolumeMethod(c *fiber.Ctx) error {
	var r_volume_method dto.R_volume_method

	setDefaults(&r_volume_method)

	if err := json.Unmarshal(c.Body(), &r_volume_method); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into r_volume_method values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	r_volume_method.Row_created_date = formatDate(r_volume_method.Row_created_date)
	r_volume_method.Row_changed_date = formatDate(r_volume_method.Row_changed_date)
	r_volume_method.Effective_date = formatDateString(r_volume_method.Effective_date)
	r_volume_method.Expiry_date = formatDateString(r_volume_method.Expiry_date)
	r_volume_method.Row_effective_date = formatDateString(r_volume_method.Row_effective_date)
	r_volume_method.Row_expiry_date = formatDateString(r_volume_method.Row_expiry_date)






	rows, err := stmt.Exec(r_volume_method.Volume_method, r_volume_method.Abbreviation, r_volume_method.Active_ind, r_volume_method.Effective_date, r_volume_method.Expiry_date, r_volume_method.Long_name, r_volume_method.Ppdm_guid, r_volume_method.Remark, r_volume_method.Short_name, r_volume_method.Source, r_volume_method.Row_changed_by, r_volume_method.Row_changed_date, r_volume_method.Row_created_by, r_volume_method.Row_created_date, r_volume_method.Row_effective_date, r_volume_method.Row_expiry_date, r_volume_method.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRVolumeMethod(c *fiber.Ctx) error {
	var r_volume_method dto.R_volume_method

	setDefaults(&r_volume_method)

	if err := json.Unmarshal(c.Body(), &r_volume_method); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	r_volume_method.Volume_method = id

    if r_volume_method.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update r_volume_method set abbreviation = :1, active_ind = :2, effective_date = :3, expiry_date = :4, long_name = :5, ppdm_guid = :6, remark = :7, short_name = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where volume_method = :17")
	if err != nil {
		return err
	}

	r_volume_method.Row_changed_date = formatDate(r_volume_method.Row_changed_date)
	r_volume_method.Effective_date = formatDateString(r_volume_method.Effective_date)
	r_volume_method.Expiry_date = formatDateString(r_volume_method.Expiry_date)
	r_volume_method.Row_effective_date = formatDateString(r_volume_method.Row_effective_date)
	r_volume_method.Row_expiry_date = formatDateString(r_volume_method.Row_expiry_date)






	rows, err := stmt.Exec(r_volume_method.Abbreviation, r_volume_method.Active_ind, r_volume_method.Effective_date, r_volume_method.Expiry_date, r_volume_method.Long_name, r_volume_method.Ppdm_guid, r_volume_method.Remark, r_volume_method.Short_name, r_volume_method.Source, r_volume_method.Row_changed_by, r_volume_method.Row_changed_date, r_volume_method.Row_created_by, r_volume_method.Row_effective_date, r_volume_method.Row_expiry_date, r_volume_method.Row_quality, r_volume_method.Volume_method)
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

func PatchRVolumeMethod(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update r_volume_method set "
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
	query += " where volume_method = :id"

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

func DeleteRVolumeMethod(c *fiber.Ctx) error {
	id := c.Params("id")
	var r_volume_method dto.R_volume_method
	r_volume_method.Volume_method = id

	stmt, err := db.Prepare("delete from r_volume_method where volume_method = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(r_volume_method.Volume_method)
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



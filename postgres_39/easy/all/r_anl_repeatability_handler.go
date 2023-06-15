package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRAnlRepeatability(c *fiber.Ctx) error {
	rows, err := db.Query("select * from r_anl_repeatability")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.R_anl_repeatability

	for rows.Next() {
		var r_anl_repeatability dto.R_anl_repeatability
		if err := rows.Scan(&r_anl_repeatability.Repeatability_type, &r_anl_repeatability.Abbreviation, &r_anl_repeatability.Active_ind, &r_anl_repeatability.Effective_date, &r_anl_repeatability.Expiry_date, &r_anl_repeatability.Long_name, &r_anl_repeatability.Ppdm_guid, &r_anl_repeatability.Remark, &r_anl_repeatability.Short_name, &r_anl_repeatability.Source, &r_anl_repeatability.Row_changed_by, &r_anl_repeatability.Row_changed_date, &r_anl_repeatability.Row_created_by, &r_anl_repeatability.Row_created_date, &r_anl_repeatability.Row_effective_date, &r_anl_repeatability.Row_expiry_date, &r_anl_repeatability.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append r_anl_repeatability to result
		result = append(result, r_anl_repeatability)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRAnlRepeatability(c *fiber.Ctx) error {
	var r_anl_repeatability dto.R_anl_repeatability

	setDefaults(&r_anl_repeatability)

	if err := json.Unmarshal(c.Body(), &r_anl_repeatability); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into r_anl_repeatability values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	r_anl_repeatability.Row_created_date = formatDate(r_anl_repeatability.Row_created_date)
	r_anl_repeatability.Row_changed_date = formatDate(r_anl_repeatability.Row_changed_date)
	r_anl_repeatability.Effective_date = formatDateString(r_anl_repeatability.Effective_date)
	r_anl_repeatability.Expiry_date = formatDateString(r_anl_repeatability.Expiry_date)
	r_anl_repeatability.Row_effective_date = formatDateString(r_anl_repeatability.Row_effective_date)
	r_anl_repeatability.Row_expiry_date = formatDateString(r_anl_repeatability.Row_expiry_date)






	rows, err := stmt.Exec(r_anl_repeatability.Repeatability_type, r_anl_repeatability.Abbreviation, r_anl_repeatability.Active_ind, r_anl_repeatability.Effective_date, r_anl_repeatability.Expiry_date, r_anl_repeatability.Long_name, r_anl_repeatability.Ppdm_guid, r_anl_repeatability.Remark, r_anl_repeatability.Short_name, r_anl_repeatability.Source, r_anl_repeatability.Row_changed_by, r_anl_repeatability.Row_changed_date, r_anl_repeatability.Row_created_by, r_anl_repeatability.Row_created_date, r_anl_repeatability.Row_effective_date, r_anl_repeatability.Row_expiry_date, r_anl_repeatability.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRAnlRepeatability(c *fiber.Ctx) error {
	var r_anl_repeatability dto.R_anl_repeatability

	setDefaults(&r_anl_repeatability)

	if err := json.Unmarshal(c.Body(), &r_anl_repeatability); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	r_anl_repeatability.Repeatability_type = id

    if r_anl_repeatability.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update r_anl_repeatability set abbreviation = :1, active_ind = :2, effective_date = :3, expiry_date = :4, long_name = :5, ppdm_guid = :6, remark = :7, short_name = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where repeatability_type = :17")
	if err != nil {
		return err
	}

	r_anl_repeatability.Row_changed_date = formatDate(r_anl_repeatability.Row_changed_date)
	r_anl_repeatability.Effective_date = formatDateString(r_anl_repeatability.Effective_date)
	r_anl_repeatability.Expiry_date = formatDateString(r_anl_repeatability.Expiry_date)
	r_anl_repeatability.Row_effective_date = formatDateString(r_anl_repeatability.Row_effective_date)
	r_anl_repeatability.Row_expiry_date = formatDateString(r_anl_repeatability.Row_expiry_date)






	rows, err := stmt.Exec(r_anl_repeatability.Abbreviation, r_anl_repeatability.Active_ind, r_anl_repeatability.Effective_date, r_anl_repeatability.Expiry_date, r_anl_repeatability.Long_name, r_anl_repeatability.Ppdm_guid, r_anl_repeatability.Remark, r_anl_repeatability.Short_name, r_anl_repeatability.Source, r_anl_repeatability.Row_changed_by, r_anl_repeatability.Row_changed_date, r_anl_repeatability.Row_created_by, r_anl_repeatability.Row_effective_date, r_anl_repeatability.Row_expiry_date, r_anl_repeatability.Row_quality, r_anl_repeatability.Repeatability_type)
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

func PatchRAnlRepeatability(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update r_anl_repeatability set "
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
	query += " where repeatability_type = :id"

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

func DeleteRAnlRepeatability(c *fiber.Ctx) error {
	id := c.Params("id")
	var r_anl_repeatability dto.R_anl_repeatability
	r_anl_repeatability.Repeatability_type = id

	stmt, err := db.Prepare("delete from r_anl_repeatability where repeatability_type = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(r_anl_repeatability.Repeatability_type)
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



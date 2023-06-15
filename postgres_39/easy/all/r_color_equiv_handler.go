package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRColorEquiv(c *fiber.Ctx) error {
	rows, err := db.Query("select * from r_color_equiv")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.R_color_equiv

	for rows.Next() {
		var r_color_equiv dto.R_color_equiv
		if err := rows.Scan(&r_color_equiv.Color_1, &r_color_equiv.Color_2, &r_color_equiv.Abbreviation, &r_color_equiv.Active_ind, &r_color_equiv.Effective_date, &r_color_equiv.Equivalent_type, &r_color_equiv.Expiry_date, &r_color_equiv.Long_name, &r_color_equiv.Ppdm_guid, &r_color_equiv.Remark, &r_color_equiv.Short_name, &r_color_equiv.Source, &r_color_equiv.Row_changed_by, &r_color_equiv.Row_changed_date, &r_color_equiv.Row_created_by, &r_color_equiv.Row_created_date, &r_color_equiv.Row_effective_date, &r_color_equiv.Row_expiry_date, &r_color_equiv.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append r_color_equiv to result
		result = append(result, r_color_equiv)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRColorEquiv(c *fiber.Ctx) error {
	var r_color_equiv dto.R_color_equiv

	setDefaults(&r_color_equiv)

	if err := json.Unmarshal(c.Body(), &r_color_equiv); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into r_color_equiv values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	r_color_equiv.Row_created_date = formatDate(r_color_equiv.Row_created_date)
	r_color_equiv.Row_changed_date = formatDate(r_color_equiv.Row_changed_date)
	r_color_equiv.Effective_date = formatDateString(r_color_equiv.Effective_date)
	r_color_equiv.Expiry_date = formatDateString(r_color_equiv.Expiry_date)
	r_color_equiv.Row_effective_date = formatDateString(r_color_equiv.Row_effective_date)
	r_color_equiv.Row_expiry_date = formatDateString(r_color_equiv.Row_expiry_date)






	rows, err := stmt.Exec(r_color_equiv.Color_1, r_color_equiv.Color_2, r_color_equiv.Abbreviation, r_color_equiv.Active_ind, r_color_equiv.Effective_date, r_color_equiv.Equivalent_type, r_color_equiv.Expiry_date, r_color_equiv.Long_name, r_color_equiv.Ppdm_guid, r_color_equiv.Remark, r_color_equiv.Short_name, r_color_equiv.Source, r_color_equiv.Row_changed_by, r_color_equiv.Row_changed_date, r_color_equiv.Row_created_by, r_color_equiv.Row_created_date, r_color_equiv.Row_effective_date, r_color_equiv.Row_expiry_date, r_color_equiv.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRColorEquiv(c *fiber.Ctx) error {
	var r_color_equiv dto.R_color_equiv

	setDefaults(&r_color_equiv)

	if err := json.Unmarshal(c.Body(), &r_color_equiv); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	r_color_equiv.Color_1 = id

    if r_color_equiv.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update r_color_equiv set color_2 = :1, abbreviation = :2, active_ind = :3, effective_date = :4, equivalent_type = :5, expiry_date = :6, long_name = :7, ppdm_guid = :8, remark = :9, short_name = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where color_1 = :19")
	if err != nil {
		return err
	}

	r_color_equiv.Row_changed_date = formatDate(r_color_equiv.Row_changed_date)
	r_color_equiv.Effective_date = formatDateString(r_color_equiv.Effective_date)
	r_color_equiv.Expiry_date = formatDateString(r_color_equiv.Expiry_date)
	r_color_equiv.Row_effective_date = formatDateString(r_color_equiv.Row_effective_date)
	r_color_equiv.Row_expiry_date = formatDateString(r_color_equiv.Row_expiry_date)






	rows, err := stmt.Exec(r_color_equiv.Color_2, r_color_equiv.Abbreviation, r_color_equiv.Active_ind, r_color_equiv.Effective_date, r_color_equiv.Equivalent_type, r_color_equiv.Expiry_date, r_color_equiv.Long_name, r_color_equiv.Ppdm_guid, r_color_equiv.Remark, r_color_equiv.Short_name, r_color_equiv.Source, r_color_equiv.Row_changed_by, r_color_equiv.Row_changed_date, r_color_equiv.Row_created_by, r_color_equiv.Row_effective_date, r_color_equiv.Row_expiry_date, r_color_equiv.Row_quality, r_color_equiv.Color_1)
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

func PatchRColorEquiv(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update r_color_equiv set "
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
	query += " where color_1 = :id"

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

func DeleteRColorEquiv(c *fiber.Ctx) error {
	id := c.Params("id")
	var r_color_equiv dto.R_color_equiv
	r_color_equiv.Color_1 = id

	stmt, err := db.Prepare("delete from r_color_equiv where color_1 = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(r_color_equiv.Color_1)
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



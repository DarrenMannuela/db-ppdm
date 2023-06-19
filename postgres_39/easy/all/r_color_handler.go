package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRColor(c *fiber.Ctx) error {
	rows, err := db.Query("select * from r_color")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.R_color

	for rows.Next() {
		var r_color dto.R_color
		if err := rows.Scan(&r_color.Color, &r_color.Abbreviation, &r_color.Active_ind, &r_color.Color_interpretation, &r_color.Color_palette, &r_color.Effective_date, &r_color.Expiry_date, &r_color.Long_name, &r_color.Ppdm_guid, &r_color.Remark, &r_color.Short_name, &r_color.Source, &r_color.Vitrinite_equivalence, &r_color.Row_changed_by, &r_color.Row_changed_date, &r_color.Row_created_by, &r_color.Row_created_date, &r_color.Row_effective_date, &r_color.Row_expiry_date, &r_color.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append r_color to result
		result = append(result, r_color)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRColor(c *fiber.Ctx) error {
	var r_color dto.R_color

	setDefaults(&r_color)

	if err := json.Unmarshal(c.Body(), &r_color); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into r_color values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	r_color.Row_created_date = formatDate(r_color.Row_created_date)
	r_color.Row_changed_date = formatDate(r_color.Row_changed_date)
	r_color.Effective_date = formatDateString(r_color.Effective_date)
	r_color.Expiry_date = formatDateString(r_color.Expiry_date)
	r_color.Row_effective_date = formatDateString(r_color.Row_effective_date)
	r_color.Row_expiry_date = formatDateString(r_color.Row_expiry_date)






	rows, err := stmt.Exec(r_color.Color, r_color.Abbreviation, r_color.Active_ind, r_color.Color_interpretation, r_color.Color_palette, r_color.Effective_date, r_color.Expiry_date, r_color.Long_name, r_color.Ppdm_guid, r_color.Remark, r_color.Short_name, r_color.Source, r_color.Vitrinite_equivalence, r_color.Row_changed_by, r_color.Row_changed_date, r_color.Row_created_by, r_color.Row_created_date, r_color.Row_effective_date, r_color.Row_expiry_date, r_color.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRColor(c *fiber.Ctx) error {
	var r_color dto.R_color

	setDefaults(&r_color)

	if err := json.Unmarshal(c.Body(), &r_color); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	r_color.Color = id

    if r_color.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update r_color set abbreviation = :1, active_ind = :2, color_interpretation = :3, color_palette = :4, effective_date = :5, expiry_date = :6, long_name = :7, ppdm_guid = :8, remark = :9, short_name = :10, source = :11, vitrinite_equivalence = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where color = :20")
	if err != nil {
		return err
	}

	r_color.Row_changed_date = formatDate(r_color.Row_changed_date)
	r_color.Effective_date = formatDateString(r_color.Effective_date)
	r_color.Expiry_date = formatDateString(r_color.Expiry_date)
	r_color.Row_effective_date = formatDateString(r_color.Row_effective_date)
	r_color.Row_expiry_date = formatDateString(r_color.Row_expiry_date)






	rows, err := stmt.Exec(r_color.Abbreviation, r_color.Active_ind, r_color.Color_interpretation, r_color.Color_palette, r_color.Effective_date, r_color.Expiry_date, r_color.Long_name, r_color.Ppdm_guid, r_color.Remark, r_color.Short_name, r_color.Source, r_color.Vitrinite_equivalence, r_color.Row_changed_by, r_color.Row_changed_date, r_color.Row_created_by, r_color.Row_effective_date, r_color.Row_expiry_date, r_color.Row_quality, r_color.Color)
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

func PatchRColor(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update r_color set "
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
	query += " where color = :id"

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

func DeleteRColor(c *fiber.Ctx) error {
	id := c.Params("id")
	var r_color dto.R_color
	r_color.Color = id

	stmt, err := db.Prepare("delete from r_color where color = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(r_color.Color)
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


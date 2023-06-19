package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRColorPalette(c *fiber.Ctx) error {
	rows, err := db.Query("select * from r_color_palette")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.R_color_palette

	for rows.Next() {
		var r_color_palette dto.R_color_palette
		if err := rows.Scan(&r_color_palette.Color_palette, &r_color_palette.Abbreviation, &r_color_palette.Active_ind, &r_color_palette.Effective_date, &r_color_palette.Expiry_date, &r_color_palette.Long_name, &r_color_palette.Palette_purpose, &r_color_palette.Palette_type, &r_color_palette.Physical_item_id, &r_color_palette.Ppdm_guid, &r_color_palette.Remark, &r_color_palette.Short_name, &r_color_palette.Source, &r_color_palette.Row_changed_by, &r_color_palette.Row_changed_date, &r_color_palette.Row_created_by, &r_color_palette.Row_created_date, &r_color_palette.Row_effective_date, &r_color_palette.Row_expiry_date, &r_color_palette.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append r_color_palette to result
		result = append(result, r_color_palette)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRColorPalette(c *fiber.Ctx) error {
	var r_color_palette dto.R_color_palette

	setDefaults(&r_color_palette)

	if err := json.Unmarshal(c.Body(), &r_color_palette); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into r_color_palette values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	r_color_palette.Row_created_date = formatDate(r_color_palette.Row_created_date)
	r_color_palette.Row_changed_date = formatDate(r_color_palette.Row_changed_date)
	r_color_palette.Effective_date = formatDateString(r_color_palette.Effective_date)
	r_color_palette.Expiry_date = formatDateString(r_color_palette.Expiry_date)
	r_color_palette.Row_effective_date = formatDateString(r_color_palette.Row_effective_date)
	r_color_palette.Row_expiry_date = formatDateString(r_color_palette.Row_expiry_date)






	rows, err := stmt.Exec(r_color_palette.Color_palette, r_color_palette.Abbreviation, r_color_palette.Active_ind, r_color_palette.Effective_date, r_color_palette.Expiry_date, r_color_palette.Long_name, r_color_palette.Palette_purpose, r_color_palette.Palette_type, r_color_palette.Physical_item_id, r_color_palette.Ppdm_guid, r_color_palette.Remark, r_color_palette.Short_name, r_color_palette.Source, r_color_palette.Row_changed_by, r_color_palette.Row_changed_date, r_color_palette.Row_created_by, r_color_palette.Row_created_date, r_color_palette.Row_effective_date, r_color_palette.Row_expiry_date, r_color_palette.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRColorPalette(c *fiber.Ctx) error {
	var r_color_palette dto.R_color_palette

	setDefaults(&r_color_palette)

	if err := json.Unmarshal(c.Body(), &r_color_palette); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	r_color_palette.Color_palette = id

    if r_color_palette.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update r_color_palette set abbreviation = :1, active_ind = :2, effective_date = :3, expiry_date = :4, long_name = :5, palette_purpose = :6, palette_type = :7, physical_item_id = :8, ppdm_guid = :9, remark = :10, short_name = :11, source = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where color_palette = :20")
	if err != nil {
		return err
	}

	r_color_palette.Row_changed_date = formatDate(r_color_palette.Row_changed_date)
	r_color_palette.Effective_date = formatDateString(r_color_palette.Effective_date)
	r_color_palette.Expiry_date = formatDateString(r_color_palette.Expiry_date)
	r_color_palette.Row_effective_date = formatDateString(r_color_palette.Row_effective_date)
	r_color_palette.Row_expiry_date = formatDateString(r_color_palette.Row_expiry_date)






	rows, err := stmt.Exec(r_color_palette.Abbreviation, r_color_palette.Active_ind, r_color_palette.Effective_date, r_color_palette.Expiry_date, r_color_palette.Long_name, r_color_palette.Palette_purpose, r_color_palette.Palette_type, r_color_palette.Physical_item_id, r_color_palette.Ppdm_guid, r_color_palette.Remark, r_color_palette.Short_name, r_color_palette.Source, r_color_palette.Row_changed_by, r_color_palette.Row_changed_date, r_color_palette.Row_created_by, r_color_palette.Row_effective_date, r_color_palette.Row_expiry_date, r_color_palette.Row_quality, r_color_palette.Color_palette)
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

func PatchRColorPalette(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update r_color_palette set "
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
	query += " where color_palette = :id"

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

func DeleteRColorPalette(c *fiber.Ctx) error {
	id := c.Params("id")
	var r_color_palette dto.R_color_palette
	r_color_palette.Color_palette = id

	stmt, err := db.Prepare("delete from r_color_palette where color_palette = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(r_color_palette.Color_palette)
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


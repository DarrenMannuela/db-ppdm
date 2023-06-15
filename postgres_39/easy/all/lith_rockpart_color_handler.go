package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLithRockpartColor(c *fiber.Ctx) error {
	rows, err := db.Query("select * from lith_rockpart_color")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Lith_rockpart_color

	for rows.Next() {
		var lith_rockpart_color dto.Lith_rockpart_color
		if err := rows.Scan(&lith_rockpart_color.Lithology_log_id, &lith_rockpart_color.Lith_log_source, &lith_rockpart_color.Depth_obs_no, &lith_rockpart_color.Rock_type, &lith_rockpart_color.Rock_type_obs_no, &lith_rockpart_color.Rockpart_name, &lith_rockpart_color.Basic_color, &lith_rockpart_color.Active_ind, &lith_rockpart_color.Color_adjective, &lith_rockpart_color.Color_distribution, &lith_rockpart_color.Dominant_ind, &lith_rockpart_color.Effective_date, &lith_rockpart_color.Expiry_date, &lith_rockpart_color.Intensity, &lith_rockpart_color.Ppdm_guid, &lith_rockpart_color.Remark, &lith_rockpart_color.Weathered_ind, &lith_rockpart_color.Row_changed_by, &lith_rockpart_color.Row_changed_date, &lith_rockpart_color.Row_created_by, &lith_rockpart_color.Row_created_date, &lith_rockpart_color.Row_effective_date, &lith_rockpart_color.Row_expiry_date, &lith_rockpart_color.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append lith_rockpart_color to result
		result = append(result, lith_rockpart_color)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLithRockpartColor(c *fiber.Ctx) error {
	var lith_rockpart_color dto.Lith_rockpart_color

	setDefaults(&lith_rockpart_color)

	if err := json.Unmarshal(c.Body(), &lith_rockpart_color); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into lith_rockpart_color values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24)")
	if err != nil {
		return err
	}
	lith_rockpart_color.Row_created_date = formatDate(lith_rockpart_color.Row_created_date)
	lith_rockpart_color.Row_changed_date = formatDate(lith_rockpart_color.Row_changed_date)
	lith_rockpart_color.Effective_date = formatDateString(lith_rockpart_color.Effective_date)
	lith_rockpart_color.Expiry_date = formatDateString(lith_rockpart_color.Expiry_date)
	lith_rockpart_color.Row_effective_date = formatDateString(lith_rockpart_color.Row_effective_date)
	lith_rockpart_color.Row_expiry_date = formatDateString(lith_rockpart_color.Row_expiry_date)






	rows, err := stmt.Exec(lith_rockpart_color.Lithology_log_id, lith_rockpart_color.Lith_log_source, lith_rockpart_color.Depth_obs_no, lith_rockpart_color.Rock_type, lith_rockpart_color.Rock_type_obs_no, lith_rockpart_color.Rockpart_name, lith_rockpart_color.Basic_color, lith_rockpart_color.Active_ind, lith_rockpart_color.Color_adjective, lith_rockpart_color.Color_distribution, lith_rockpart_color.Dominant_ind, lith_rockpart_color.Effective_date, lith_rockpart_color.Expiry_date, lith_rockpart_color.Intensity, lith_rockpart_color.Ppdm_guid, lith_rockpart_color.Remark, lith_rockpart_color.Weathered_ind, lith_rockpart_color.Row_changed_by, lith_rockpart_color.Row_changed_date, lith_rockpart_color.Row_created_by, lith_rockpart_color.Row_created_date, lith_rockpart_color.Row_effective_date, lith_rockpart_color.Row_expiry_date, lith_rockpart_color.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLithRockpartColor(c *fiber.Ctx) error {
	var lith_rockpart_color dto.Lith_rockpart_color

	setDefaults(&lith_rockpart_color)

	if err := json.Unmarshal(c.Body(), &lith_rockpart_color); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	lith_rockpart_color.Lithology_log_id = id

    if lith_rockpart_color.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update lith_rockpart_color set lith_log_source = :1, depth_obs_no = :2, rock_type = :3, rock_type_obs_no = :4, rockpart_name = :5, basic_color = :6, active_ind = :7, color_adjective = :8, color_distribution = :9, dominant_ind = :10, effective_date = :11, expiry_date = :12, intensity = :13, ppdm_guid = :14, remark = :15, weathered_ind = :16, row_changed_by = :17, row_changed_date = :18, row_created_by = :19, row_effective_date = :20, row_expiry_date = :21, row_quality = :22 where lithology_log_id = :24")
	if err != nil {
		return err
	}

	lith_rockpart_color.Row_changed_date = formatDate(lith_rockpart_color.Row_changed_date)
	lith_rockpart_color.Effective_date = formatDateString(lith_rockpart_color.Effective_date)
	lith_rockpart_color.Expiry_date = formatDateString(lith_rockpart_color.Expiry_date)
	lith_rockpart_color.Row_effective_date = formatDateString(lith_rockpart_color.Row_effective_date)
	lith_rockpart_color.Row_expiry_date = formatDateString(lith_rockpart_color.Row_expiry_date)






	rows, err := stmt.Exec(lith_rockpart_color.Lith_log_source, lith_rockpart_color.Depth_obs_no, lith_rockpart_color.Rock_type, lith_rockpart_color.Rock_type_obs_no, lith_rockpart_color.Rockpart_name, lith_rockpart_color.Basic_color, lith_rockpart_color.Active_ind, lith_rockpart_color.Color_adjective, lith_rockpart_color.Color_distribution, lith_rockpart_color.Dominant_ind, lith_rockpart_color.Effective_date, lith_rockpart_color.Expiry_date, lith_rockpart_color.Intensity, lith_rockpart_color.Ppdm_guid, lith_rockpart_color.Remark, lith_rockpart_color.Weathered_ind, lith_rockpart_color.Row_changed_by, lith_rockpart_color.Row_changed_date, lith_rockpart_color.Row_created_by, lith_rockpart_color.Row_effective_date, lith_rockpart_color.Row_expiry_date, lith_rockpart_color.Row_quality, lith_rockpart_color.Lithology_log_id)
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

func PatchLithRockpartColor(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update lith_rockpart_color set "
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
	query += " where lithology_log_id = :id"

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

func DeleteLithRockpartColor(c *fiber.Ctx) error {
	id := c.Params("id")
	var lith_rockpart_color dto.Lith_rockpart_color
	lith_rockpart_color.Lithology_log_id = id

	stmt, err := db.Prepare("delete from lith_rockpart_color where lithology_log_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(lith_rockpart_color.Lithology_log_id)
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



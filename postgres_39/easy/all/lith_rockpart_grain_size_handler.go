package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLithRockpartGrainSize(c *fiber.Ctx) error {
	rows, err := db.Query("select * from lith_rockpart_grain_size")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Lith_rockpart_grain_size

	for rows.Next() {
		var lith_rockpart_grain_size dto.Lith_rockpart_grain_size
		if err := rows.Scan(&lith_rockpart_grain_size.Lithology_log_id, &lith_rockpart_grain_size.Lith_log_source, &lith_rockpart_grain_size.Depth_obs_no, &lith_rockpart_grain_size.Scale_scheme, &lith_rockpart_grain_size.Rock_type, &lith_rockpart_grain_size.Rock_type_obs_no, &lith_rockpart_grain_size.Rockpart_name, &lith_rockpart_grain_size.Grain_size, &lith_rockpart_grain_size.Active_ind, &lith_rockpart_grain_size.Distribution_ind, &lith_rockpart_grain_size.Effective_date, &lith_rockpart_grain_size.Expiry_date, &lith_rockpart_grain_size.Ppdm_guid, &lith_rockpart_grain_size.Remark, &lith_rockpart_grain_size.Size_type_ind, &lith_rockpart_grain_size.Row_changed_by, &lith_rockpart_grain_size.Row_changed_date, &lith_rockpart_grain_size.Row_created_by, &lith_rockpart_grain_size.Row_created_date, &lith_rockpart_grain_size.Row_effective_date, &lith_rockpart_grain_size.Row_expiry_date, &lith_rockpart_grain_size.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append lith_rockpart_grain_size to result
		result = append(result, lith_rockpart_grain_size)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLithRockpartGrainSize(c *fiber.Ctx) error {
	var lith_rockpart_grain_size dto.Lith_rockpart_grain_size

	setDefaults(&lith_rockpart_grain_size)

	if err := json.Unmarshal(c.Body(), &lith_rockpart_grain_size); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into lith_rockpart_grain_size values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	lith_rockpart_grain_size.Row_created_date = formatDate(lith_rockpart_grain_size.Row_created_date)
	lith_rockpart_grain_size.Row_changed_date = formatDate(lith_rockpart_grain_size.Row_changed_date)
	lith_rockpart_grain_size.Effective_date = formatDateString(lith_rockpart_grain_size.Effective_date)
	lith_rockpart_grain_size.Expiry_date = formatDateString(lith_rockpart_grain_size.Expiry_date)
	lith_rockpart_grain_size.Row_effective_date = formatDateString(lith_rockpart_grain_size.Row_effective_date)
	lith_rockpart_grain_size.Row_expiry_date = formatDateString(lith_rockpart_grain_size.Row_expiry_date)






	rows, err := stmt.Exec(lith_rockpart_grain_size.Lithology_log_id, lith_rockpart_grain_size.Lith_log_source, lith_rockpart_grain_size.Depth_obs_no, lith_rockpart_grain_size.Scale_scheme, lith_rockpart_grain_size.Rock_type, lith_rockpart_grain_size.Rock_type_obs_no, lith_rockpart_grain_size.Rockpart_name, lith_rockpart_grain_size.Grain_size, lith_rockpart_grain_size.Active_ind, lith_rockpart_grain_size.Distribution_ind, lith_rockpart_grain_size.Effective_date, lith_rockpart_grain_size.Expiry_date, lith_rockpart_grain_size.Ppdm_guid, lith_rockpart_grain_size.Remark, lith_rockpart_grain_size.Size_type_ind, lith_rockpart_grain_size.Row_changed_by, lith_rockpart_grain_size.Row_changed_date, lith_rockpart_grain_size.Row_created_by, lith_rockpart_grain_size.Row_created_date, lith_rockpart_grain_size.Row_effective_date, lith_rockpart_grain_size.Row_expiry_date, lith_rockpart_grain_size.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLithRockpartGrainSize(c *fiber.Ctx) error {
	var lith_rockpart_grain_size dto.Lith_rockpart_grain_size

	setDefaults(&lith_rockpart_grain_size)

	if err := json.Unmarshal(c.Body(), &lith_rockpart_grain_size); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	lith_rockpart_grain_size.Lithology_log_id = id

    if lith_rockpart_grain_size.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update lith_rockpart_grain_size set lith_log_source = :1, depth_obs_no = :2, scale_scheme = :3, rock_type = :4, rock_type_obs_no = :5, rockpart_name = :6, grain_size = :7, active_ind = :8, distribution_ind = :9, effective_date = :10, expiry_date = :11, ppdm_guid = :12, remark = :13, size_type_ind = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where lithology_log_id = :22")
	if err != nil {
		return err
	}

	lith_rockpart_grain_size.Row_changed_date = formatDate(lith_rockpart_grain_size.Row_changed_date)
	lith_rockpart_grain_size.Effective_date = formatDateString(lith_rockpart_grain_size.Effective_date)
	lith_rockpart_grain_size.Expiry_date = formatDateString(lith_rockpart_grain_size.Expiry_date)
	lith_rockpart_grain_size.Row_effective_date = formatDateString(lith_rockpart_grain_size.Row_effective_date)
	lith_rockpart_grain_size.Row_expiry_date = formatDateString(lith_rockpart_grain_size.Row_expiry_date)






	rows, err := stmt.Exec(lith_rockpart_grain_size.Lith_log_source, lith_rockpart_grain_size.Depth_obs_no, lith_rockpart_grain_size.Scale_scheme, lith_rockpart_grain_size.Rock_type, lith_rockpart_grain_size.Rock_type_obs_no, lith_rockpart_grain_size.Rockpart_name, lith_rockpart_grain_size.Grain_size, lith_rockpart_grain_size.Active_ind, lith_rockpart_grain_size.Distribution_ind, lith_rockpart_grain_size.Effective_date, lith_rockpart_grain_size.Expiry_date, lith_rockpart_grain_size.Ppdm_guid, lith_rockpart_grain_size.Remark, lith_rockpart_grain_size.Size_type_ind, lith_rockpart_grain_size.Row_changed_by, lith_rockpart_grain_size.Row_changed_date, lith_rockpart_grain_size.Row_created_by, lith_rockpart_grain_size.Row_effective_date, lith_rockpart_grain_size.Row_expiry_date, lith_rockpart_grain_size.Row_quality, lith_rockpart_grain_size.Lithology_log_id)
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

func PatchLithRockpartGrainSize(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update lith_rockpart_grain_size set "
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

func DeleteLithRockpartGrainSize(c *fiber.Ctx) error {
	id := c.Params("id")
	var lith_rockpart_grain_size dto.Lith_rockpart_grain_size
	lith_rockpart_grain_size.Lithology_log_id = id

	stmt, err := db.Prepare("delete from lith_rockpart_grain_size where lithology_log_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(lith_rockpart_grain_size.Lithology_log_id)
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



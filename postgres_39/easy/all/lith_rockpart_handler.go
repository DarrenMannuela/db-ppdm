package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLithRockpart(c *fiber.Ctx) error {
	rows, err := db.Query("select * from lith_rockpart")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Lith_rockpart

	for rows.Next() {
		var lith_rockpart dto.Lith_rockpart
		if err := rows.Scan(&lith_rockpart.Lithology_log_id, &lith_rockpart.Lith_log_source, &lith_rockpart.Depth_obs_no, &lith_rockpart.Rock_type, &lith_rockpart.Rock_type_obs_no, &lith_rockpart.Rockpart_name, &lith_rockpart.Active_ind, &lith_rockpart.Effective_date, &lith_rockpart.Expiry_date, &lith_rockpart.Lith_pattern_code, &lith_rockpart.Ppdm_guid, &lith_rockpart.Remark, &lith_rockpart.Rockpart_percent, &lith_rockpart.Rockpart_rel_abundance, &lith_rockpart.Rockpart_texture, &lith_rockpart.Rockpart_type, &lith_rockpart.Row_changed_by, &lith_rockpart.Row_changed_date, &lith_rockpart.Row_created_by, &lith_rockpart.Row_created_date, &lith_rockpart.Row_effective_date, &lith_rockpart.Row_expiry_date, &lith_rockpart.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append lith_rockpart to result
		result = append(result, lith_rockpart)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLithRockpart(c *fiber.Ctx) error {
	var lith_rockpart dto.Lith_rockpart

	setDefaults(&lith_rockpart)

	if err := json.Unmarshal(c.Body(), &lith_rockpart); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into lith_rockpart values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	lith_rockpart.Row_created_date = formatDate(lith_rockpart.Row_created_date)
	lith_rockpart.Row_changed_date = formatDate(lith_rockpart.Row_changed_date)
	lith_rockpart.Effective_date = formatDateString(lith_rockpart.Effective_date)
	lith_rockpart.Expiry_date = formatDateString(lith_rockpart.Expiry_date)
	lith_rockpart.Row_effective_date = formatDateString(lith_rockpart.Row_effective_date)
	lith_rockpart.Row_expiry_date = formatDateString(lith_rockpart.Row_expiry_date)






	rows, err := stmt.Exec(lith_rockpart.Lithology_log_id, lith_rockpart.Lith_log_source, lith_rockpart.Depth_obs_no, lith_rockpart.Rock_type, lith_rockpart.Rock_type_obs_no, lith_rockpart.Rockpart_name, lith_rockpart.Active_ind, lith_rockpart.Effective_date, lith_rockpart.Expiry_date, lith_rockpart.Lith_pattern_code, lith_rockpart.Ppdm_guid, lith_rockpart.Remark, lith_rockpart.Rockpart_percent, lith_rockpart.Rockpart_rel_abundance, lith_rockpart.Rockpart_texture, lith_rockpart.Rockpart_type, lith_rockpart.Row_changed_by, lith_rockpart.Row_changed_date, lith_rockpart.Row_created_by, lith_rockpart.Row_created_date, lith_rockpart.Row_effective_date, lith_rockpart.Row_expiry_date, lith_rockpart.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLithRockpart(c *fiber.Ctx) error {
	var lith_rockpart dto.Lith_rockpart

	setDefaults(&lith_rockpart)

	if err := json.Unmarshal(c.Body(), &lith_rockpart); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	lith_rockpart.Lithology_log_id = id

    if lith_rockpart.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update lith_rockpart set lith_log_source = :1, depth_obs_no = :2, rock_type = :3, rock_type_obs_no = :4, rockpart_name = :5, active_ind = :6, effective_date = :7, expiry_date = :8, lith_pattern_code = :9, ppdm_guid = :10, remark = :11, rockpart_percent = :12, rockpart_rel_abundance = :13, rockpart_texture = :14, rockpart_type = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where lithology_log_id = :23")
	if err != nil {
		return err
	}

	lith_rockpart.Row_changed_date = formatDate(lith_rockpart.Row_changed_date)
	lith_rockpart.Effective_date = formatDateString(lith_rockpart.Effective_date)
	lith_rockpart.Expiry_date = formatDateString(lith_rockpart.Expiry_date)
	lith_rockpart.Row_effective_date = formatDateString(lith_rockpart.Row_effective_date)
	lith_rockpart.Row_expiry_date = formatDateString(lith_rockpart.Row_expiry_date)






	rows, err := stmt.Exec(lith_rockpart.Lith_log_source, lith_rockpart.Depth_obs_no, lith_rockpart.Rock_type, lith_rockpart.Rock_type_obs_no, lith_rockpart.Rockpart_name, lith_rockpart.Active_ind, lith_rockpart.Effective_date, lith_rockpart.Expiry_date, lith_rockpart.Lith_pattern_code, lith_rockpart.Ppdm_guid, lith_rockpart.Remark, lith_rockpart.Rockpart_percent, lith_rockpart.Rockpart_rel_abundance, lith_rockpart.Rockpart_texture, lith_rockpart.Rockpart_type, lith_rockpart.Row_changed_by, lith_rockpart.Row_changed_date, lith_rockpart.Row_created_by, lith_rockpart.Row_effective_date, lith_rockpart.Row_expiry_date, lith_rockpart.Row_quality, lith_rockpart.Lithology_log_id)
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

func PatchLithRockpart(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update lith_rockpart set "
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

func DeleteLithRockpart(c *fiber.Ctx) error {
	id := c.Params("id")
	var lith_rockpart dto.Lith_rockpart
	lith_rockpart.Lithology_log_id = id

	stmt, err := db.Prepare("delete from lith_rockpart where lithology_log_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(lith_rockpart.Lithology_log_id)
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



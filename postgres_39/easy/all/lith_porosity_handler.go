package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLithPorosity(c *fiber.Ctx) error {
	rows, err := db.Query("select * from lith_porosity")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Lith_porosity

	for rows.Next() {
		var lith_porosity dto.Lith_porosity
		if err := rows.Scan(&lith_porosity.Lithology_log_id, &lith_porosity.Lith_log_source, &lith_porosity.Depth_obs_no, &lith_porosity.Rock_type, &lith_porosity.Rock_type_obs_no, &lith_porosity.Porosity_grade, &lith_porosity.Active_ind, &lith_porosity.Effective_date, &lith_porosity.Expiry_date, &lith_porosity.Matrix_class, &lith_porosity.Percent_of_sample, &lith_porosity.Porosity_scale, &lith_porosity.Porosity_type, &lith_porosity.Ppdm_guid, &lith_porosity.Remark, &lith_porosity.Sample_rel_abundance, &lith_porosity.Row_changed_by, &lith_porosity.Row_changed_date, &lith_porosity.Row_created_by, &lith_porosity.Row_created_date, &lith_porosity.Row_effective_date, &lith_porosity.Row_expiry_date, &lith_porosity.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append lith_porosity to result
		result = append(result, lith_porosity)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLithPorosity(c *fiber.Ctx) error {
	var lith_porosity dto.Lith_porosity

	setDefaults(&lith_porosity)

	if err := json.Unmarshal(c.Body(), &lith_porosity); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into lith_porosity values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	lith_porosity.Row_created_date = formatDate(lith_porosity.Row_created_date)
	lith_porosity.Row_changed_date = formatDate(lith_porosity.Row_changed_date)
	lith_porosity.Effective_date = formatDateString(lith_porosity.Effective_date)
	lith_porosity.Expiry_date = formatDateString(lith_porosity.Expiry_date)
	lith_porosity.Row_effective_date = formatDateString(lith_porosity.Row_effective_date)
	lith_porosity.Row_expiry_date = formatDateString(lith_porosity.Row_expiry_date)






	rows, err := stmt.Exec(lith_porosity.Lithology_log_id, lith_porosity.Lith_log_source, lith_porosity.Depth_obs_no, lith_porosity.Rock_type, lith_porosity.Rock_type_obs_no, lith_porosity.Porosity_grade, lith_porosity.Active_ind, lith_porosity.Effective_date, lith_porosity.Expiry_date, lith_porosity.Matrix_class, lith_porosity.Percent_of_sample, lith_porosity.Porosity_scale, lith_porosity.Porosity_type, lith_porosity.Ppdm_guid, lith_porosity.Remark, lith_porosity.Sample_rel_abundance, lith_porosity.Row_changed_by, lith_porosity.Row_changed_date, lith_porosity.Row_created_by, lith_porosity.Row_created_date, lith_porosity.Row_effective_date, lith_porosity.Row_expiry_date, lith_porosity.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLithPorosity(c *fiber.Ctx) error {
	var lith_porosity dto.Lith_porosity

	setDefaults(&lith_porosity)

	if err := json.Unmarshal(c.Body(), &lith_porosity); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	lith_porosity.Lithology_log_id = id

    if lith_porosity.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update lith_porosity set lith_log_source = :1, depth_obs_no = :2, rock_type = :3, rock_type_obs_no = :4, porosity_grade = :5, active_ind = :6, effective_date = :7, expiry_date = :8, matrix_class = :9, percent_of_sample = :10, porosity_scale = :11, porosity_type = :12, ppdm_guid = :13, remark = :14, sample_rel_abundance = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where lithology_log_id = :23")
	if err != nil {
		return err
	}

	lith_porosity.Row_changed_date = formatDate(lith_porosity.Row_changed_date)
	lith_porosity.Effective_date = formatDateString(lith_porosity.Effective_date)
	lith_porosity.Expiry_date = formatDateString(lith_porosity.Expiry_date)
	lith_porosity.Row_effective_date = formatDateString(lith_porosity.Row_effective_date)
	lith_porosity.Row_expiry_date = formatDateString(lith_porosity.Row_expiry_date)






	rows, err := stmt.Exec(lith_porosity.Lith_log_source, lith_porosity.Depth_obs_no, lith_porosity.Rock_type, lith_porosity.Rock_type_obs_no, lith_porosity.Porosity_grade, lith_porosity.Active_ind, lith_porosity.Effective_date, lith_porosity.Expiry_date, lith_porosity.Matrix_class, lith_porosity.Percent_of_sample, lith_porosity.Porosity_scale, lith_porosity.Porosity_type, lith_porosity.Ppdm_guid, lith_porosity.Remark, lith_porosity.Sample_rel_abundance, lith_porosity.Row_changed_by, lith_porosity.Row_changed_date, lith_porosity.Row_created_by, lith_porosity.Row_effective_date, lith_porosity.Row_expiry_date, lith_porosity.Row_quality, lith_porosity.Lithology_log_id)
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

func PatchLithPorosity(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update lith_porosity set "
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

func DeleteLithPorosity(c *fiber.Ctx) error {
	id := c.Params("id")
	var lith_porosity dto.Lith_porosity
	lith_porosity.Lithology_log_id = id

	stmt, err := db.Prepare("delete from lith_porosity where lithology_log_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(lith_porosity.Lithology_log_id)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPaleoSumSample(c *fiber.Ctx) error {
	rows, err := db.Query("select * from paleo_sum_sample")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Paleo_sum_sample

	for rows.Next() {
		var paleo_sum_sample dto.Paleo_sum_sample
		if err := rows.Scan(&paleo_sum_sample.Paleo_summary_id, &paleo_sum_sample.Lith_sample_id, &paleo_sum_sample.Active_ind, &paleo_sum_sample.Diversity_count, &paleo_sum_sample.Effective_date, &paleo_sum_sample.Expiry_date, &paleo_sum_sample.First_sample_ind, &paleo_sum_sample.Last_sample_ind, &paleo_sum_sample.Ppdm_guid, &paleo_sum_sample.Remark, &paleo_sum_sample.Source, &paleo_sum_sample.Row_changed_by, &paleo_sum_sample.Row_changed_date, &paleo_sum_sample.Row_created_by, &paleo_sum_sample.Row_created_date, &paleo_sum_sample.Row_effective_date, &paleo_sum_sample.Row_expiry_date, &paleo_sum_sample.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append paleo_sum_sample to result
		result = append(result, paleo_sum_sample)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPaleoSumSample(c *fiber.Ctx) error {
	var paleo_sum_sample dto.Paleo_sum_sample

	setDefaults(&paleo_sum_sample)

	if err := json.Unmarshal(c.Body(), &paleo_sum_sample); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into paleo_sum_sample values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	paleo_sum_sample.Row_created_date = formatDate(paleo_sum_sample.Row_created_date)
	paleo_sum_sample.Row_changed_date = formatDate(paleo_sum_sample.Row_changed_date)
	paleo_sum_sample.Effective_date = formatDateString(paleo_sum_sample.Effective_date)
	paleo_sum_sample.Expiry_date = formatDateString(paleo_sum_sample.Expiry_date)
	paleo_sum_sample.Row_effective_date = formatDateString(paleo_sum_sample.Row_effective_date)
	paleo_sum_sample.Row_expiry_date = formatDateString(paleo_sum_sample.Row_expiry_date)






	rows, err := stmt.Exec(paleo_sum_sample.Paleo_summary_id, paleo_sum_sample.Lith_sample_id, paleo_sum_sample.Active_ind, paleo_sum_sample.Diversity_count, paleo_sum_sample.Effective_date, paleo_sum_sample.Expiry_date, paleo_sum_sample.First_sample_ind, paleo_sum_sample.Last_sample_ind, paleo_sum_sample.Ppdm_guid, paleo_sum_sample.Remark, paleo_sum_sample.Source, paleo_sum_sample.Row_changed_by, paleo_sum_sample.Row_changed_date, paleo_sum_sample.Row_created_by, paleo_sum_sample.Row_created_date, paleo_sum_sample.Row_effective_date, paleo_sum_sample.Row_expiry_date, paleo_sum_sample.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePaleoSumSample(c *fiber.Ctx) error {
	var paleo_sum_sample dto.Paleo_sum_sample

	setDefaults(&paleo_sum_sample)

	if err := json.Unmarshal(c.Body(), &paleo_sum_sample); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	paleo_sum_sample.Paleo_summary_id = id

    if paleo_sum_sample.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update paleo_sum_sample set lith_sample_id = :1, active_ind = :2, diversity_count = :3, effective_date = :4, expiry_date = :5, first_sample_ind = :6, last_sample_ind = :7, ppdm_guid = :8, remark = :9, source = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where paleo_summary_id = :18")
	if err != nil {
		return err
	}

	paleo_sum_sample.Row_changed_date = formatDate(paleo_sum_sample.Row_changed_date)
	paleo_sum_sample.Effective_date = formatDateString(paleo_sum_sample.Effective_date)
	paleo_sum_sample.Expiry_date = formatDateString(paleo_sum_sample.Expiry_date)
	paleo_sum_sample.Row_effective_date = formatDateString(paleo_sum_sample.Row_effective_date)
	paleo_sum_sample.Row_expiry_date = formatDateString(paleo_sum_sample.Row_expiry_date)






	rows, err := stmt.Exec(paleo_sum_sample.Lith_sample_id, paleo_sum_sample.Active_ind, paleo_sum_sample.Diversity_count, paleo_sum_sample.Effective_date, paleo_sum_sample.Expiry_date, paleo_sum_sample.First_sample_ind, paleo_sum_sample.Last_sample_ind, paleo_sum_sample.Ppdm_guid, paleo_sum_sample.Remark, paleo_sum_sample.Source, paleo_sum_sample.Row_changed_by, paleo_sum_sample.Row_changed_date, paleo_sum_sample.Row_created_by, paleo_sum_sample.Row_effective_date, paleo_sum_sample.Row_expiry_date, paleo_sum_sample.Row_quality, paleo_sum_sample.Paleo_summary_id)
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

func PatchPaleoSumSample(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update paleo_sum_sample set "
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
	query += " where paleo_summary_id = :id"

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

func DeletePaleoSumSample(c *fiber.Ctx) error {
	id := c.Params("id")
	var paleo_sum_sample dto.Paleo_sum_sample
	paleo_sum_sample.Paleo_summary_id = id

	stmt, err := db.Prepare("delete from paleo_sum_sample where paleo_summary_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(paleo_sum_sample.Paleo_summary_id)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmLithSample(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_lith_sample")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_lith_sample

	for rows.Next() {
		var rm_lith_sample dto.Rm_lith_sample
		if err := rows.Scan(&rm_lith_sample.Info_item_subtype, &rm_lith_sample.Information_item_id, &rm_lith_sample.Active_ind, &rm_lith_sample.Effective_date, &rm_lith_sample.Expiry_date, &rm_lith_sample.Ppdm_guid, &rm_lith_sample.Remark, &rm_lith_sample.Source, &rm_lith_sample.Row_changed_by, &rm_lith_sample.Row_changed_date, &rm_lith_sample.Row_created_by, &rm_lith_sample.Row_created_date, &rm_lith_sample.Row_effective_date, &rm_lith_sample.Row_expiry_date, &rm_lith_sample.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_lith_sample to result
		result = append(result, rm_lith_sample)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmLithSample(c *fiber.Ctx) error {
	var rm_lith_sample dto.Rm_lith_sample

	setDefaults(&rm_lith_sample)

	if err := json.Unmarshal(c.Body(), &rm_lith_sample); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_lith_sample values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15)")
	if err != nil {
		return err
	}
	rm_lith_sample.Row_created_date = formatDate(rm_lith_sample.Row_created_date)
	rm_lith_sample.Row_changed_date = formatDate(rm_lith_sample.Row_changed_date)
	rm_lith_sample.Effective_date = formatDateString(rm_lith_sample.Effective_date)
	rm_lith_sample.Expiry_date = formatDateString(rm_lith_sample.Expiry_date)
	rm_lith_sample.Row_effective_date = formatDateString(rm_lith_sample.Row_effective_date)
	rm_lith_sample.Row_expiry_date = formatDateString(rm_lith_sample.Row_expiry_date)






	rows, err := stmt.Exec(rm_lith_sample.Info_item_subtype, rm_lith_sample.Information_item_id, rm_lith_sample.Active_ind, rm_lith_sample.Effective_date, rm_lith_sample.Expiry_date, rm_lith_sample.Ppdm_guid, rm_lith_sample.Remark, rm_lith_sample.Source, rm_lith_sample.Row_changed_by, rm_lith_sample.Row_changed_date, rm_lith_sample.Row_created_by, rm_lith_sample.Row_created_date, rm_lith_sample.Row_effective_date, rm_lith_sample.Row_expiry_date, rm_lith_sample.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmLithSample(c *fiber.Ctx) error {
	var rm_lith_sample dto.Rm_lith_sample

	setDefaults(&rm_lith_sample)

	if err := json.Unmarshal(c.Body(), &rm_lith_sample); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_lith_sample.Info_item_subtype = id

    if rm_lith_sample.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_lith_sample set information_item_id = :1, active_ind = :2, effective_date = :3, expiry_date = :4, ppdm_guid = :5, remark = :6, source = :7, row_changed_by = :8, row_changed_date = :9, row_created_by = :10, row_effective_date = :11, row_expiry_date = :12, row_quality = :13 where info_item_subtype = :15")
	if err != nil {
		return err
	}

	rm_lith_sample.Row_changed_date = formatDate(rm_lith_sample.Row_changed_date)
	rm_lith_sample.Effective_date = formatDateString(rm_lith_sample.Effective_date)
	rm_lith_sample.Expiry_date = formatDateString(rm_lith_sample.Expiry_date)
	rm_lith_sample.Row_effective_date = formatDateString(rm_lith_sample.Row_effective_date)
	rm_lith_sample.Row_expiry_date = formatDateString(rm_lith_sample.Row_expiry_date)






	rows, err := stmt.Exec(rm_lith_sample.Information_item_id, rm_lith_sample.Active_ind, rm_lith_sample.Effective_date, rm_lith_sample.Expiry_date, rm_lith_sample.Ppdm_guid, rm_lith_sample.Remark, rm_lith_sample.Source, rm_lith_sample.Row_changed_by, rm_lith_sample.Row_changed_date, rm_lith_sample.Row_created_by, rm_lith_sample.Row_effective_date, rm_lith_sample.Row_expiry_date, rm_lith_sample.Row_quality, rm_lith_sample.Info_item_subtype)
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

func PatchRmLithSample(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_lith_sample set "
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
	query += " where info_item_subtype = :id"

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

func DeleteRmLithSample(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_lith_sample dto.Rm_lith_sample
	rm_lith_sample.Info_item_subtype = id

	stmt, err := db.Prepare("delete from rm_lith_sample where info_item_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_lith_sample.Info_item_subtype)
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



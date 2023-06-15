package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellLogCurveSplice(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_log_curve_splice")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_log_curve_splice

	for rows.Next() {
		var well_log_curve_splice dto.Well_log_curve_splice
		if err := rows.Scan(&well_log_curve_splice.Uwi, &well_log_curve_splice.Curve_id, &well_log_curve_splice.Splice_obs_no, &well_log_curve_splice.Active_ind, &well_log_curve_splice.Created_by_ba_id, &well_log_curve_splice.Effective_date, &well_log_curve_splice.Expiry_date, &well_log_curve_splice.Information_item_id, &well_log_curve_splice.Info_item_subtype, &well_log_curve_splice.Max_value_index, &well_log_curve_splice.Min_value_index, &well_log_curve_splice.Ppdm_guid, &well_log_curve_splice.Remark, &well_log_curve_splice.Source, &well_log_curve_splice.Source_curve_id, &well_log_curve_splice.Row_changed_by, &well_log_curve_splice.Row_changed_date, &well_log_curve_splice.Row_created_by, &well_log_curve_splice.Row_created_date, &well_log_curve_splice.Row_effective_date, &well_log_curve_splice.Row_expiry_date, &well_log_curve_splice.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_log_curve_splice to result
		result = append(result, well_log_curve_splice)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellLogCurveSplice(c *fiber.Ctx) error {
	var well_log_curve_splice dto.Well_log_curve_splice

	setDefaults(&well_log_curve_splice)

	if err := json.Unmarshal(c.Body(), &well_log_curve_splice); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_log_curve_splice values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	well_log_curve_splice.Row_created_date = formatDate(well_log_curve_splice.Row_created_date)
	well_log_curve_splice.Row_changed_date = formatDate(well_log_curve_splice.Row_changed_date)
	well_log_curve_splice.Effective_date = formatDateString(well_log_curve_splice.Effective_date)
	well_log_curve_splice.Expiry_date = formatDateString(well_log_curve_splice.Expiry_date)
	well_log_curve_splice.Row_effective_date = formatDateString(well_log_curve_splice.Row_effective_date)
	well_log_curve_splice.Row_expiry_date = formatDateString(well_log_curve_splice.Row_expiry_date)






	rows, err := stmt.Exec(well_log_curve_splice.Uwi, well_log_curve_splice.Curve_id, well_log_curve_splice.Splice_obs_no, well_log_curve_splice.Active_ind, well_log_curve_splice.Created_by_ba_id, well_log_curve_splice.Effective_date, well_log_curve_splice.Expiry_date, well_log_curve_splice.Information_item_id, well_log_curve_splice.Info_item_subtype, well_log_curve_splice.Max_value_index, well_log_curve_splice.Min_value_index, well_log_curve_splice.Ppdm_guid, well_log_curve_splice.Remark, well_log_curve_splice.Source, well_log_curve_splice.Source_curve_id, well_log_curve_splice.Row_changed_by, well_log_curve_splice.Row_changed_date, well_log_curve_splice.Row_created_by, well_log_curve_splice.Row_created_date, well_log_curve_splice.Row_effective_date, well_log_curve_splice.Row_expiry_date, well_log_curve_splice.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellLogCurveSplice(c *fiber.Ctx) error {
	var well_log_curve_splice dto.Well_log_curve_splice

	setDefaults(&well_log_curve_splice)

	if err := json.Unmarshal(c.Body(), &well_log_curve_splice); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_log_curve_splice.Uwi = id

    if well_log_curve_splice.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_log_curve_splice set curve_id = :1, splice_obs_no = :2, active_ind = :3, created_by_ba_id = :4, effective_date = :5, expiry_date = :6, information_item_id = :7, info_item_subtype = :8, max_value_index = :9, min_value_index = :10, ppdm_guid = :11, remark = :12, source = :13, source_curve_id = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where uwi = :22")
	if err != nil {
		return err
	}

	well_log_curve_splice.Row_changed_date = formatDate(well_log_curve_splice.Row_changed_date)
	well_log_curve_splice.Effective_date = formatDateString(well_log_curve_splice.Effective_date)
	well_log_curve_splice.Expiry_date = formatDateString(well_log_curve_splice.Expiry_date)
	well_log_curve_splice.Row_effective_date = formatDateString(well_log_curve_splice.Row_effective_date)
	well_log_curve_splice.Row_expiry_date = formatDateString(well_log_curve_splice.Row_expiry_date)






	rows, err := stmt.Exec(well_log_curve_splice.Curve_id, well_log_curve_splice.Splice_obs_no, well_log_curve_splice.Active_ind, well_log_curve_splice.Created_by_ba_id, well_log_curve_splice.Effective_date, well_log_curve_splice.Expiry_date, well_log_curve_splice.Information_item_id, well_log_curve_splice.Info_item_subtype, well_log_curve_splice.Max_value_index, well_log_curve_splice.Min_value_index, well_log_curve_splice.Ppdm_guid, well_log_curve_splice.Remark, well_log_curve_splice.Source, well_log_curve_splice.Source_curve_id, well_log_curve_splice.Row_changed_by, well_log_curve_splice.Row_changed_date, well_log_curve_splice.Row_created_by, well_log_curve_splice.Row_effective_date, well_log_curve_splice.Row_expiry_date, well_log_curve_splice.Row_quality, well_log_curve_splice.Uwi)
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

func PatchWellLogCurveSplice(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_log_curve_splice set "
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
	query += " where uwi = :id"

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

func DeleteWellLogCurveSplice(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_log_curve_splice dto.Well_log_curve_splice
	well_log_curve_splice.Uwi = id

	stmt, err := db.Prepare("delete from well_log_curve_splice where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_log_curve_splice.Uwi)
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



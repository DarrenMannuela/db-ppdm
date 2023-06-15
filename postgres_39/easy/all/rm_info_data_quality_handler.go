package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmInfoDataQuality(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_info_data_quality")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_info_data_quality

	for rows.Next() {
		var rm_info_data_quality dto.Rm_info_data_quality
		if err := rows.Scan(&rm_info_data_quality.Info_item_subtype, &rm_info_data_quality.Information_item_id, &rm_info_data_quality.Quality_id, &rm_info_data_quality.Active_ind, &rm_info_data_quality.Attribute_accuracy, &rm_info_data_quality.Attribute_accuracy_desc, &rm_info_data_quality.Completeness_desc, &rm_info_data_quality.Corrected_date, &rm_info_data_quality.Deficiency_ind, &rm_info_data_quality.Deficiency_type, &rm_info_data_quality.Described_by_ba_id, &rm_info_data_quality.Description, &rm_info_data_quality.Effective_date, &rm_info_data_quality.Expiry_date, &rm_info_data_quality.Information_quality, &rm_info_data_quality.Information_quality_desc, &rm_info_data_quality.Ppdm_guid, &rm_info_data_quality.Quality_date, &rm_info_data_quality.Remark, &rm_info_data_quality.Source, &rm_info_data_quality.Row_changed_by, &rm_info_data_quality.Row_changed_date, &rm_info_data_quality.Row_created_by, &rm_info_data_quality.Row_created_date, &rm_info_data_quality.Row_effective_date, &rm_info_data_quality.Row_expiry_date, &rm_info_data_quality.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_info_data_quality to result
		result = append(result, rm_info_data_quality)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmInfoDataQuality(c *fiber.Ctx) error {
	var rm_info_data_quality dto.Rm_info_data_quality

	setDefaults(&rm_info_data_quality)

	if err := json.Unmarshal(c.Body(), &rm_info_data_quality); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_info_data_quality values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27)")
	if err != nil {
		return err
	}
	rm_info_data_quality.Row_created_date = formatDate(rm_info_data_quality.Row_created_date)
	rm_info_data_quality.Row_changed_date = formatDate(rm_info_data_quality.Row_changed_date)
	rm_info_data_quality.Corrected_date = formatDateString(rm_info_data_quality.Corrected_date)
	rm_info_data_quality.Effective_date = formatDateString(rm_info_data_quality.Effective_date)
	rm_info_data_quality.Expiry_date = formatDateString(rm_info_data_quality.Expiry_date)
	rm_info_data_quality.Quality_date = formatDateString(rm_info_data_quality.Quality_date)
	rm_info_data_quality.Row_effective_date = formatDateString(rm_info_data_quality.Row_effective_date)
	rm_info_data_quality.Row_expiry_date = formatDateString(rm_info_data_quality.Row_expiry_date)






	rows, err := stmt.Exec(rm_info_data_quality.Info_item_subtype, rm_info_data_quality.Information_item_id, rm_info_data_quality.Quality_id, rm_info_data_quality.Active_ind, rm_info_data_quality.Attribute_accuracy, rm_info_data_quality.Attribute_accuracy_desc, rm_info_data_quality.Completeness_desc, rm_info_data_quality.Corrected_date, rm_info_data_quality.Deficiency_ind, rm_info_data_quality.Deficiency_type, rm_info_data_quality.Described_by_ba_id, rm_info_data_quality.Description, rm_info_data_quality.Effective_date, rm_info_data_quality.Expiry_date, rm_info_data_quality.Information_quality, rm_info_data_quality.Information_quality_desc, rm_info_data_quality.Ppdm_guid, rm_info_data_quality.Quality_date, rm_info_data_quality.Remark, rm_info_data_quality.Source, rm_info_data_quality.Row_changed_by, rm_info_data_quality.Row_changed_date, rm_info_data_quality.Row_created_by, rm_info_data_quality.Row_created_date, rm_info_data_quality.Row_effective_date, rm_info_data_quality.Row_expiry_date, rm_info_data_quality.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmInfoDataQuality(c *fiber.Ctx) error {
	var rm_info_data_quality dto.Rm_info_data_quality

	setDefaults(&rm_info_data_quality)

	if err := json.Unmarshal(c.Body(), &rm_info_data_quality); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_info_data_quality.Info_item_subtype = id

    if rm_info_data_quality.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_info_data_quality set information_item_id = :1, quality_id = :2, active_ind = :3, attribute_accuracy = :4, attribute_accuracy_desc = :5, completeness_desc = :6, corrected_date = :7, deficiency_ind = :8, deficiency_type = :9, described_by_ba_id = :10, description = :11, effective_date = :12, expiry_date = :13, information_quality = :14, information_quality_desc = :15, ppdm_guid = :16, quality_date = :17, remark = :18, source = :19, row_changed_by = :20, row_changed_date = :21, row_created_by = :22, row_effective_date = :23, row_expiry_date = :24, row_quality = :25 where info_item_subtype = :27")
	if err != nil {
		return err
	}

	rm_info_data_quality.Row_changed_date = formatDate(rm_info_data_quality.Row_changed_date)
	rm_info_data_quality.Corrected_date = formatDateString(rm_info_data_quality.Corrected_date)
	rm_info_data_quality.Effective_date = formatDateString(rm_info_data_quality.Effective_date)
	rm_info_data_quality.Expiry_date = formatDateString(rm_info_data_quality.Expiry_date)
	rm_info_data_quality.Quality_date = formatDateString(rm_info_data_quality.Quality_date)
	rm_info_data_quality.Row_effective_date = formatDateString(rm_info_data_quality.Row_effective_date)
	rm_info_data_quality.Row_expiry_date = formatDateString(rm_info_data_quality.Row_expiry_date)






	rows, err := stmt.Exec(rm_info_data_quality.Information_item_id, rm_info_data_quality.Quality_id, rm_info_data_quality.Active_ind, rm_info_data_quality.Attribute_accuracy, rm_info_data_quality.Attribute_accuracy_desc, rm_info_data_quality.Completeness_desc, rm_info_data_quality.Corrected_date, rm_info_data_quality.Deficiency_ind, rm_info_data_quality.Deficiency_type, rm_info_data_quality.Described_by_ba_id, rm_info_data_quality.Description, rm_info_data_quality.Effective_date, rm_info_data_quality.Expiry_date, rm_info_data_quality.Information_quality, rm_info_data_quality.Information_quality_desc, rm_info_data_quality.Ppdm_guid, rm_info_data_quality.Quality_date, rm_info_data_quality.Remark, rm_info_data_quality.Source, rm_info_data_quality.Row_changed_by, rm_info_data_quality.Row_changed_date, rm_info_data_quality.Row_created_by, rm_info_data_quality.Row_effective_date, rm_info_data_quality.Row_expiry_date, rm_info_data_quality.Row_quality, rm_info_data_quality.Info_item_subtype)
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

func PatchRmInfoDataQuality(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_info_data_quality set "
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
		} else if key == "corrected_date" || key == "effective_date" || key == "expiry_date" || key == "quality_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteRmInfoDataQuality(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_info_data_quality dto.Rm_info_data_quality
	rm_info_data_quality.Info_item_subtype = id

	stmt, err := db.Prepare("delete from rm_info_data_quality where info_item_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_info_data_quality.Info_item_subtype)
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



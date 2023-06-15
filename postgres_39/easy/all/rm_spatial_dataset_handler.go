package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmSpatialDataset(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_spatial_dataset")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_spatial_dataset

	for rows.Next() {
		var rm_spatial_dataset dto.Rm_spatial_dataset
		if err := rows.Scan(&rm_spatial_dataset.Info_item_subtype, &rm_spatial_dataset.Information_item_id, &rm_spatial_dataset.Active_ind, &rm_spatial_dataset.Effective_date, &rm_spatial_dataset.Expiry_date, &rm_spatial_dataset.Ppdm_guid, &rm_spatial_dataset.Remark, &rm_spatial_dataset.Source, &rm_spatial_dataset.Row_changed_by, &rm_spatial_dataset.Row_changed_date, &rm_spatial_dataset.Row_created_by, &rm_spatial_dataset.Row_created_date, &rm_spatial_dataset.Row_effective_date, &rm_spatial_dataset.Row_expiry_date, &rm_spatial_dataset.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_spatial_dataset to result
		result = append(result, rm_spatial_dataset)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmSpatialDataset(c *fiber.Ctx) error {
	var rm_spatial_dataset dto.Rm_spatial_dataset

	setDefaults(&rm_spatial_dataset)

	if err := json.Unmarshal(c.Body(), &rm_spatial_dataset); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_spatial_dataset values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15)")
	if err != nil {
		return err
	}
	rm_spatial_dataset.Row_created_date = formatDate(rm_spatial_dataset.Row_created_date)
	rm_spatial_dataset.Row_changed_date = formatDate(rm_spatial_dataset.Row_changed_date)
	rm_spatial_dataset.Effective_date = formatDateString(rm_spatial_dataset.Effective_date)
	rm_spatial_dataset.Expiry_date = formatDateString(rm_spatial_dataset.Expiry_date)
	rm_spatial_dataset.Row_effective_date = formatDateString(rm_spatial_dataset.Row_effective_date)
	rm_spatial_dataset.Row_expiry_date = formatDateString(rm_spatial_dataset.Row_expiry_date)






	rows, err := stmt.Exec(rm_spatial_dataset.Info_item_subtype, rm_spatial_dataset.Information_item_id, rm_spatial_dataset.Active_ind, rm_spatial_dataset.Effective_date, rm_spatial_dataset.Expiry_date, rm_spatial_dataset.Ppdm_guid, rm_spatial_dataset.Remark, rm_spatial_dataset.Source, rm_spatial_dataset.Row_changed_by, rm_spatial_dataset.Row_changed_date, rm_spatial_dataset.Row_created_by, rm_spatial_dataset.Row_created_date, rm_spatial_dataset.Row_effective_date, rm_spatial_dataset.Row_expiry_date, rm_spatial_dataset.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmSpatialDataset(c *fiber.Ctx) error {
	var rm_spatial_dataset dto.Rm_spatial_dataset

	setDefaults(&rm_spatial_dataset)

	if err := json.Unmarshal(c.Body(), &rm_spatial_dataset); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_spatial_dataset.Info_item_subtype = id

    if rm_spatial_dataset.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_spatial_dataset set information_item_id = :1, active_ind = :2, effective_date = :3, expiry_date = :4, ppdm_guid = :5, remark = :6, source = :7, row_changed_by = :8, row_changed_date = :9, row_created_by = :10, row_effective_date = :11, row_expiry_date = :12, row_quality = :13 where info_item_subtype = :15")
	if err != nil {
		return err
	}

	rm_spatial_dataset.Row_changed_date = formatDate(rm_spatial_dataset.Row_changed_date)
	rm_spatial_dataset.Effective_date = formatDateString(rm_spatial_dataset.Effective_date)
	rm_spatial_dataset.Expiry_date = formatDateString(rm_spatial_dataset.Expiry_date)
	rm_spatial_dataset.Row_effective_date = formatDateString(rm_spatial_dataset.Row_effective_date)
	rm_spatial_dataset.Row_expiry_date = formatDateString(rm_spatial_dataset.Row_expiry_date)






	rows, err := stmt.Exec(rm_spatial_dataset.Information_item_id, rm_spatial_dataset.Active_ind, rm_spatial_dataset.Effective_date, rm_spatial_dataset.Expiry_date, rm_spatial_dataset.Ppdm_guid, rm_spatial_dataset.Remark, rm_spatial_dataset.Source, rm_spatial_dataset.Row_changed_by, rm_spatial_dataset.Row_changed_date, rm_spatial_dataset.Row_created_by, rm_spatial_dataset.Row_effective_date, rm_spatial_dataset.Row_expiry_date, rm_spatial_dataset.Row_quality, rm_spatial_dataset.Info_item_subtype)
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

func PatchRmSpatialDataset(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_spatial_dataset set "
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

func DeleteRmSpatialDataset(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_spatial_dataset dto.Rm_spatial_dataset
	rm_spatial_dataset.Info_item_subtype = id

	stmt, err := db.Prepare("delete from rm_spatial_dataset where info_item_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_spatial_dataset.Info_item_subtype)
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



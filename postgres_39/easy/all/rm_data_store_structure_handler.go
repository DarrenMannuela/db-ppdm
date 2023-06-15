package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmDataStoreStructure(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_data_store_structure")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_data_store_structure

	for rows.Next() {
		var rm_data_store_structure dto.Rm_data_store_structure
		if err := rows.Scan(&rm_data_store_structure.Store_id, &rm_data_store_structure.Structure_obs_no, &rm_data_store_structure.Active_ind, &rm_data_store_structure.Dim_height, &rm_data_store_structure.Dim_height_ouom, &rm_data_store_structure.Dim_height_uom, &rm_data_store_structure.Dim_length, &rm_data_store_structure.Dim_length_ouom, &rm_data_store_structure.Dim_length_uom, &rm_data_store_structure.Dim_weight, &rm_data_store_structure.Dim_weight_ouom, &rm_data_store_structure.Dim_weight_uom, &rm_data_store_structure.Dim_width, &rm_data_store_structure.Dim_width_ouom, &rm_data_store_structure.Dim_width_uom, &rm_data_store_structure.Effective_date, &rm_data_store_structure.Expiry_date, &rm_data_store_structure.Lower_contained_store_id, &rm_data_store_structure.Ppdm_guid, &rm_data_store_structure.Remark, &rm_data_store_structure.Source, &rm_data_store_structure.Total_capacity, &rm_data_store_structure.Upper_contained_store_id, &rm_data_store_structure.Used_capacity, &rm_data_store_structure.Used_capacity_date, &rm_data_store_structure.Row_changed_by, &rm_data_store_structure.Row_changed_date, &rm_data_store_structure.Row_created_by, &rm_data_store_structure.Row_created_date, &rm_data_store_structure.Row_effective_date, &rm_data_store_structure.Row_expiry_date, &rm_data_store_structure.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_data_store_structure to result
		result = append(result, rm_data_store_structure)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmDataStoreStructure(c *fiber.Ctx) error {
	var rm_data_store_structure dto.Rm_data_store_structure

	setDefaults(&rm_data_store_structure)

	if err := json.Unmarshal(c.Body(), &rm_data_store_structure); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_data_store_structure values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32)")
	if err != nil {
		return err
	}
	rm_data_store_structure.Row_created_date = formatDate(rm_data_store_structure.Row_created_date)
	rm_data_store_structure.Row_changed_date = formatDate(rm_data_store_structure.Row_changed_date)
	rm_data_store_structure.Effective_date = formatDateString(rm_data_store_structure.Effective_date)
	rm_data_store_structure.Expiry_date = formatDateString(rm_data_store_structure.Expiry_date)
	rm_data_store_structure.Used_capacity_date = formatDateString(rm_data_store_structure.Used_capacity_date)
	rm_data_store_structure.Row_effective_date = formatDateString(rm_data_store_structure.Row_effective_date)
	rm_data_store_structure.Row_expiry_date = formatDateString(rm_data_store_structure.Row_expiry_date)






	rows, err := stmt.Exec(rm_data_store_structure.Store_id, rm_data_store_structure.Structure_obs_no, rm_data_store_structure.Active_ind, rm_data_store_structure.Dim_height, rm_data_store_structure.Dim_height_ouom, rm_data_store_structure.Dim_height_uom, rm_data_store_structure.Dim_length, rm_data_store_structure.Dim_length_ouom, rm_data_store_structure.Dim_length_uom, rm_data_store_structure.Dim_weight, rm_data_store_structure.Dim_weight_ouom, rm_data_store_structure.Dim_weight_uom, rm_data_store_structure.Dim_width, rm_data_store_structure.Dim_width_ouom, rm_data_store_structure.Dim_width_uom, rm_data_store_structure.Effective_date, rm_data_store_structure.Expiry_date, rm_data_store_structure.Lower_contained_store_id, rm_data_store_structure.Ppdm_guid, rm_data_store_structure.Remark, rm_data_store_structure.Source, rm_data_store_structure.Total_capacity, rm_data_store_structure.Upper_contained_store_id, rm_data_store_structure.Used_capacity, rm_data_store_structure.Used_capacity_date, rm_data_store_structure.Row_changed_by, rm_data_store_structure.Row_changed_date, rm_data_store_structure.Row_created_by, rm_data_store_structure.Row_created_date, rm_data_store_structure.Row_effective_date, rm_data_store_structure.Row_expiry_date, rm_data_store_structure.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmDataStoreStructure(c *fiber.Ctx) error {
	var rm_data_store_structure dto.Rm_data_store_structure

	setDefaults(&rm_data_store_structure)

	if err := json.Unmarshal(c.Body(), &rm_data_store_structure); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_data_store_structure.Store_id = id

    if rm_data_store_structure.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_data_store_structure set structure_obs_no = :1, active_ind = :2, dim_height = :3, dim_height_ouom = :4, dim_height_uom = :5, dim_length = :6, dim_length_ouom = :7, dim_length_uom = :8, dim_weight = :9, dim_weight_ouom = :10, dim_weight_uom = :11, dim_width = :12, dim_width_ouom = :13, dim_width_uom = :14, effective_date = :15, expiry_date = :16, lower_contained_store_id = :17, ppdm_guid = :18, remark = :19, source = :20, total_capacity = :21, upper_contained_store_id = :22, used_capacity = :23, used_capacity_date = :24, row_changed_by = :25, row_changed_date = :26, row_created_by = :27, row_effective_date = :28, row_expiry_date = :29, row_quality = :30 where store_id = :32")
	if err != nil {
		return err
	}

	rm_data_store_structure.Row_changed_date = formatDate(rm_data_store_structure.Row_changed_date)
	rm_data_store_structure.Effective_date = formatDateString(rm_data_store_structure.Effective_date)
	rm_data_store_structure.Expiry_date = formatDateString(rm_data_store_structure.Expiry_date)
	rm_data_store_structure.Used_capacity_date = formatDateString(rm_data_store_structure.Used_capacity_date)
	rm_data_store_structure.Row_effective_date = formatDateString(rm_data_store_structure.Row_effective_date)
	rm_data_store_structure.Row_expiry_date = formatDateString(rm_data_store_structure.Row_expiry_date)






	rows, err := stmt.Exec(rm_data_store_structure.Structure_obs_no, rm_data_store_structure.Active_ind, rm_data_store_structure.Dim_height, rm_data_store_structure.Dim_height_ouom, rm_data_store_structure.Dim_height_uom, rm_data_store_structure.Dim_length, rm_data_store_structure.Dim_length_ouom, rm_data_store_structure.Dim_length_uom, rm_data_store_structure.Dim_weight, rm_data_store_structure.Dim_weight_ouom, rm_data_store_structure.Dim_weight_uom, rm_data_store_structure.Dim_width, rm_data_store_structure.Dim_width_ouom, rm_data_store_structure.Dim_width_uom, rm_data_store_structure.Effective_date, rm_data_store_structure.Expiry_date, rm_data_store_structure.Lower_contained_store_id, rm_data_store_structure.Ppdm_guid, rm_data_store_structure.Remark, rm_data_store_structure.Source, rm_data_store_structure.Total_capacity, rm_data_store_structure.Upper_contained_store_id, rm_data_store_structure.Used_capacity, rm_data_store_structure.Used_capacity_date, rm_data_store_structure.Row_changed_by, rm_data_store_structure.Row_changed_date, rm_data_store_structure.Row_created_by, rm_data_store_structure.Row_effective_date, rm_data_store_structure.Row_expiry_date, rm_data_store_structure.Row_quality, rm_data_store_structure.Store_id)
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

func PatchRmDataStoreStructure(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_data_store_structure set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "used_capacity_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where store_id = :id"

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

func DeleteRmDataStoreStructure(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_data_store_structure dto.Rm_data_store_structure
	rm_data_store_structure.Store_id = id

	stmt, err := db.Prepare("delete from rm_data_store_structure where store_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_data_store_structure.Store_id)
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



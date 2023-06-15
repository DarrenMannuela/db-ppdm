package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmDataStoreHier(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_data_store_hier")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_data_store_hier

	for rows.Next() {
		var rm_data_store_hier dto.Rm_data_store_hier
		if err := rows.Scan(&rm_data_store_hier.Data_store_hier_id, &rm_data_store_hier.Active_ind, &rm_data_store_hier.Effective_date, &rm_data_store_hier.Expiry_date, &rm_data_store_hier.Hierarchy_name, &rm_data_store_hier.Ppdm_guid, &rm_data_store_hier.Remark, &rm_data_store_hier.Source, &rm_data_store_hier.Row_changed_by, &rm_data_store_hier.Row_changed_date, &rm_data_store_hier.Row_created_by, &rm_data_store_hier.Row_created_date, &rm_data_store_hier.Row_effective_date, &rm_data_store_hier.Row_expiry_date, &rm_data_store_hier.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_data_store_hier to result
		result = append(result, rm_data_store_hier)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmDataStoreHier(c *fiber.Ctx) error {
	var rm_data_store_hier dto.Rm_data_store_hier

	setDefaults(&rm_data_store_hier)

	if err := json.Unmarshal(c.Body(), &rm_data_store_hier); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_data_store_hier values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15)")
	if err != nil {
		return err
	}
	rm_data_store_hier.Row_created_date = formatDate(rm_data_store_hier.Row_created_date)
	rm_data_store_hier.Row_changed_date = formatDate(rm_data_store_hier.Row_changed_date)
	rm_data_store_hier.Effective_date = formatDateString(rm_data_store_hier.Effective_date)
	rm_data_store_hier.Expiry_date = formatDateString(rm_data_store_hier.Expiry_date)
	rm_data_store_hier.Row_effective_date = formatDateString(rm_data_store_hier.Row_effective_date)
	rm_data_store_hier.Row_expiry_date = formatDateString(rm_data_store_hier.Row_expiry_date)






	rows, err := stmt.Exec(rm_data_store_hier.Data_store_hier_id, rm_data_store_hier.Active_ind, rm_data_store_hier.Effective_date, rm_data_store_hier.Expiry_date, rm_data_store_hier.Hierarchy_name, rm_data_store_hier.Ppdm_guid, rm_data_store_hier.Remark, rm_data_store_hier.Source, rm_data_store_hier.Row_changed_by, rm_data_store_hier.Row_changed_date, rm_data_store_hier.Row_created_by, rm_data_store_hier.Row_created_date, rm_data_store_hier.Row_effective_date, rm_data_store_hier.Row_expiry_date, rm_data_store_hier.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmDataStoreHier(c *fiber.Ctx) error {
	var rm_data_store_hier dto.Rm_data_store_hier

	setDefaults(&rm_data_store_hier)

	if err := json.Unmarshal(c.Body(), &rm_data_store_hier); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_data_store_hier.Data_store_hier_id = id

    if rm_data_store_hier.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_data_store_hier set active_ind = :1, effective_date = :2, expiry_date = :3, hierarchy_name = :4, ppdm_guid = :5, remark = :6, source = :7, row_changed_by = :8, row_changed_date = :9, row_created_by = :10, row_effective_date = :11, row_expiry_date = :12, row_quality = :13 where data_store_hier_id = :15")
	if err != nil {
		return err
	}

	rm_data_store_hier.Row_changed_date = formatDate(rm_data_store_hier.Row_changed_date)
	rm_data_store_hier.Effective_date = formatDateString(rm_data_store_hier.Effective_date)
	rm_data_store_hier.Expiry_date = formatDateString(rm_data_store_hier.Expiry_date)
	rm_data_store_hier.Row_effective_date = formatDateString(rm_data_store_hier.Row_effective_date)
	rm_data_store_hier.Row_expiry_date = formatDateString(rm_data_store_hier.Row_expiry_date)






	rows, err := stmt.Exec(rm_data_store_hier.Active_ind, rm_data_store_hier.Effective_date, rm_data_store_hier.Expiry_date, rm_data_store_hier.Hierarchy_name, rm_data_store_hier.Ppdm_guid, rm_data_store_hier.Remark, rm_data_store_hier.Source, rm_data_store_hier.Row_changed_by, rm_data_store_hier.Row_changed_date, rm_data_store_hier.Row_created_by, rm_data_store_hier.Row_effective_date, rm_data_store_hier.Row_expiry_date, rm_data_store_hier.Row_quality, rm_data_store_hier.Data_store_hier_id)
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

func PatchRmDataStoreHier(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_data_store_hier set "
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
	query += " where data_store_hier_id = :id"

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

func DeleteRmDataStoreHier(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_data_store_hier dto.Rm_data_store_hier
	rm_data_store_hier.Data_store_hier_id = id

	stmt, err := db.Prepare("delete from rm_data_store_hier where data_store_hier_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_data_store_hier.Data_store_hier_id)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmMap(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_map")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_map

	for rows.Next() {
		var rm_map dto.Rm_map
		if err := rows.Scan(&rm_map.Info_item_subtype, &rm_map.Information_item_id, &rm_map.Active_ind, &rm_map.Effective_date, &rm_map.Expiry_date, &rm_map.Map_scale, &rm_map.Ppdm_guid, &rm_map.Remark, &rm_map.Source, &rm_map.Row_changed_by, &rm_map.Row_changed_date, &rm_map.Row_created_by, &rm_map.Row_created_date, &rm_map.Row_effective_date, &rm_map.Row_expiry_date, &rm_map.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_map to result
		result = append(result, rm_map)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmMap(c *fiber.Ctx) error {
	var rm_map dto.Rm_map

	setDefaults(&rm_map)

	if err := json.Unmarshal(c.Body(), &rm_map); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_map values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16)")
	if err != nil {
		return err
	}
	rm_map.Row_created_date = formatDate(rm_map.Row_created_date)
	rm_map.Row_changed_date = formatDate(rm_map.Row_changed_date)
	rm_map.Effective_date = formatDateString(rm_map.Effective_date)
	rm_map.Expiry_date = formatDateString(rm_map.Expiry_date)
	rm_map.Row_effective_date = formatDateString(rm_map.Row_effective_date)
	rm_map.Row_expiry_date = formatDateString(rm_map.Row_expiry_date)






	rows, err := stmt.Exec(rm_map.Info_item_subtype, rm_map.Information_item_id, rm_map.Active_ind, rm_map.Effective_date, rm_map.Expiry_date, rm_map.Map_scale, rm_map.Ppdm_guid, rm_map.Remark, rm_map.Source, rm_map.Row_changed_by, rm_map.Row_changed_date, rm_map.Row_created_by, rm_map.Row_created_date, rm_map.Row_effective_date, rm_map.Row_expiry_date, rm_map.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmMap(c *fiber.Ctx) error {
	var rm_map dto.Rm_map

	setDefaults(&rm_map)

	if err := json.Unmarshal(c.Body(), &rm_map); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_map.Info_item_subtype = id

    if rm_map.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_map set information_item_id = :1, active_ind = :2, effective_date = :3, expiry_date = :4, map_scale = :5, ppdm_guid = :6, remark = :7, source = :8, row_changed_by = :9, row_changed_date = :10, row_created_by = :11, row_effective_date = :12, row_expiry_date = :13, row_quality = :14 where info_item_subtype = :16")
	if err != nil {
		return err
	}

	rm_map.Row_changed_date = formatDate(rm_map.Row_changed_date)
	rm_map.Effective_date = formatDateString(rm_map.Effective_date)
	rm_map.Expiry_date = formatDateString(rm_map.Expiry_date)
	rm_map.Row_effective_date = formatDateString(rm_map.Row_effective_date)
	rm_map.Row_expiry_date = formatDateString(rm_map.Row_expiry_date)






	rows, err := stmt.Exec(rm_map.Information_item_id, rm_map.Active_ind, rm_map.Effective_date, rm_map.Expiry_date, rm_map.Map_scale, rm_map.Ppdm_guid, rm_map.Remark, rm_map.Source, rm_map.Row_changed_by, rm_map.Row_changed_date, rm_map.Row_created_by, rm_map.Row_effective_date, rm_map.Row_expiry_date, rm_map.Row_quality, rm_map.Info_item_subtype)
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

func PatchRmMap(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_map set "
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

func DeleteRmMap(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_map dto.Rm_map
	rm_map.Info_item_subtype = id

	stmt, err := db.Prepare("delete from rm_map where info_item_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_map.Info_item_subtype)
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



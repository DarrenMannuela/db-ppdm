package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmFossil(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_fossil")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_fossil

	for rows.Next() {
		var rm_fossil dto.Rm_fossil
		if err := rows.Scan(&rm_fossil.Info_item_subtype, &rm_fossil.Information_item_id, &rm_fossil.Active_ind, &rm_fossil.Effective_date, &rm_fossil.Expiry_date, &rm_fossil.Ppdm_guid, &rm_fossil.Remark, &rm_fossil.Source, &rm_fossil.Type_fossil_type, &rm_fossil.Row_changed_by, &rm_fossil.Row_changed_date, &rm_fossil.Row_created_by, &rm_fossil.Row_created_date, &rm_fossil.Row_effective_date, &rm_fossil.Row_expiry_date, &rm_fossil.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_fossil to result
		result = append(result, rm_fossil)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmFossil(c *fiber.Ctx) error {
	var rm_fossil dto.Rm_fossil

	setDefaults(&rm_fossil)

	if err := json.Unmarshal(c.Body(), &rm_fossil); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_fossil values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16)")
	if err != nil {
		return err
	}
	rm_fossil.Row_created_date = formatDate(rm_fossil.Row_created_date)
	rm_fossil.Row_changed_date = formatDate(rm_fossil.Row_changed_date)
	rm_fossil.Effective_date = formatDateString(rm_fossil.Effective_date)
	rm_fossil.Expiry_date = formatDateString(rm_fossil.Expiry_date)
	rm_fossil.Row_effective_date = formatDateString(rm_fossil.Row_effective_date)
	rm_fossil.Row_expiry_date = formatDateString(rm_fossil.Row_expiry_date)






	rows, err := stmt.Exec(rm_fossil.Info_item_subtype, rm_fossil.Information_item_id, rm_fossil.Active_ind, rm_fossil.Effective_date, rm_fossil.Expiry_date, rm_fossil.Ppdm_guid, rm_fossil.Remark, rm_fossil.Source, rm_fossil.Type_fossil_type, rm_fossil.Row_changed_by, rm_fossil.Row_changed_date, rm_fossil.Row_created_by, rm_fossil.Row_created_date, rm_fossil.Row_effective_date, rm_fossil.Row_expiry_date, rm_fossil.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmFossil(c *fiber.Ctx) error {
	var rm_fossil dto.Rm_fossil

	setDefaults(&rm_fossil)

	if err := json.Unmarshal(c.Body(), &rm_fossil); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_fossil.Info_item_subtype = id

    if rm_fossil.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_fossil set information_item_id = :1, active_ind = :2, effective_date = :3, expiry_date = :4, ppdm_guid = :5, remark = :6, source = :7, type_fossil_type = :8, row_changed_by = :9, row_changed_date = :10, row_created_by = :11, row_effective_date = :12, row_expiry_date = :13, row_quality = :14 where info_item_subtype = :16")
	if err != nil {
		return err
	}

	rm_fossil.Row_changed_date = formatDate(rm_fossil.Row_changed_date)
	rm_fossil.Effective_date = formatDateString(rm_fossil.Effective_date)
	rm_fossil.Expiry_date = formatDateString(rm_fossil.Expiry_date)
	rm_fossil.Row_effective_date = formatDateString(rm_fossil.Row_effective_date)
	rm_fossil.Row_expiry_date = formatDateString(rm_fossil.Row_expiry_date)






	rows, err := stmt.Exec(rm_fossil.Information_item_id, rm_fossil.Active_ind, rm_fossil.Effective_date, rm_fossil.Expiry_date, rm_fossil.Ppdm_guid, rm_fossil.Remark, rm_fossil.Source, rm_fossil.Type_fossil_type, rm_fossil.Row_changed_by, rm_fossil.Row_changed_date, rm_fossil.Row_created_by, rm_fossil.Row_effective_date, rm_fossil.Row_expiry_date, rm_fossil.Row_quality, rm_fossil.Info_item_subtype)
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

func PatchRmFossil(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_fossil set "
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

func DeleteRmFossil(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_fossil dto.Rm_fossil
	rm_fossil.Info_item_subtype = id

	stmt, err := db.Prepare("delete from rm_fossil where info_item_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_fossil.Info_item_subtype)
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



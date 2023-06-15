package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmComposite(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_composite")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_composite

	for rows.Next() {
		var rm_composite dto.Rm_composite
		if err := rows.Scan(&rm_composite.Info_item_subtype, &rm_composite.Information_item_id, &rm_composite.Active_ind, &rm_composite.Effective_date, &rm_composite.Expiry_date, &rm_composite.Ppdm_guid, &rm_composite.Remark, &rm_composite.Source, &rm_composite.Row_changed_by, &rm_composite.Row_changed_date, &rm_composite.Row_created_by, &rm_composite.Row_created_date, &rm_composite.Row_effective_date, &rm_composite.Row_expiry_date, &rm_composite.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_composite to result
		result = append(result, rm_composite)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmComposite(c *fiber.Ctx) error {
	var rm_composite dto.Rm_composite

	setDefaults(&rm_composite)

	if err := json.Unmarshal(c.Body(), &rm_composite); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_composite values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15)")
	if err != nil {
		return err
	}
	rm_composite.Row_created_date = formatDate(rm_composite.Row_created_date)
	rm_composite.Row_changed_date = formatDate(rm_composite.Row_changed_date)
	rm_composite.Effective_date = formatDateString(rm_composite.Effective_date)
	rm_composite.Expiry_date = formatDateString(rm_composite.Expiry_date)
	rm_composite.Row_effective_date = formatDateString(rm_composite.Row_effective_date)
	rm_composite.Row_expiry_date = formatDateString(rm_composite.Row_expiry_date)






	rows, err := stmt.Exec(rm_composite.Info_item_subtype, rm_composite.Information_item_id, rm_composite.Active_ind, rm_composite.Effective_date, rm_composite.Expiry_date, rm_composite.Ppdm_guid, rm_composite.Remark, rm_composite.Source, rm_composite.Row_changed_by, rm_composite.Row_changed_date, rm_composite.Row_created_by, rm_composite.Row_created_date, rm_composite.Row_effective_date, rm_composite.Row_expiry_date, rm_composite.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmComposite(c *fiber.Ctx) error {
	var rm_composite dto.Rm_composite

	setDefaults(&rm_composite)

	if err := json.Unmarshal(c.Body(), &rm_composite); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_composite.Info_item_subtype = id

    if rm_composite.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_composite set information_item_id = :1, active_ind = :2, effective_date = :3, expiry_date = :4, ppdm_guid = :5, remark = :6, source = :7, row_changed_by = :8, row_changed_date = :9, row_created_by = :10, row_effective_date = :11, row_expiry_date = :12, row_quality = :13 where info_item_subtype = :15")
	if err != nil {
		return err
	}

	rm_composite.Row_changed_date = formatDate(rm_composite.Row_changed_date)
	rm_composite.Effective_date = formatDateString(rm_composite.Effective_date)
	rm_composite.Expiry_date = formatDateString(rm_composite.Expiry_date)
	rm_composite.Row_effective_date = formatDateString(rm_composite.Row_effective_date)
	rm_composite.Row_expiry_date = formatDateString(rm_composite.Row_expiry_date)






	rows, err := stmt.Exec(rm_composite.Information_item_id, rm_composite.Active_ind, rm_composite.Effective_date, rm_composite.Expiry_date, rm_composite.Ppdm_guid, rm_composite.Remark, rm_composite.Source, rm_composite.Row_changed_by, rm_composite.Row_changed_date, rm_composite.Row_created_by, rm_composite.Row_effective_date, rm_composite.Row_expiry_date, rm_composite.Row_quality, rm_composite.Info_item_subtype)
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

func PatchRmComposite(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_composite set "
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

func DeleteRmComposite(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_composite dto.Rm_composite
	rm_composite.Info_item_subtype = id

	stmt, err := db.Prepare("delete from rm_composite where info_item_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_composite.Info_item_subtype)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetCatAdditiveGroup(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cat_additive_group")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cat_additive_group

	for rows.Next() {
		var cat_additive_group dto.Cat_additive_group
		if err := rows.Scan(&cat_additive_group.Additive_group_id, &cat_additive_group.Active_ind, &cat_additive_group.Effective_date, &cat_additive_group.Expiry_date, &cat_additive_group.Long_name, &cat_additive_group.Owner_ba_id, &cat_additive_group.Ppdm_guid, &cat_additive_group.Remark, &cat_additive_group.Short_name, &cat_additive_group.Source, &cat_additive_group.Source_document_id, &cat_additive_group.Substance_id, &cat_additive_group.Row_changed_by, &cat_additive_group.Row_changed_date, &cat_additive_group.Row_created_by, &cat_additive_group.Row_created_date, &cat_additive_group.Row_effective_date, &cat_additive_group.Row_expiry_date, &cat_additive_group.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cat_additive_group to result
		result = append(result, cat_additive_group)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetCatAdditiveGroup(c *fiber.Ctx) error {
	var cat_additive_group dto.Cat_additive_group

	setDefaults(&cat_additive_group)

	if err := json.Unmarshal(c.Body(), &cat_additive_group); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cat_additive_group values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	cat_additive_group.Row_created_date = formatDate(cat_additive_group.Row_created_date)
	cat_additive_group.Row_changed_date = formatDate(cat_additive_group.Row_changed_date)
	cat_additive_group.Effective_date = formatDateString(cat_additive_group.Effective_date)
	cat_additive_group.Expiry_date = formatDateString(cat_additive_group.Expiry_date)
	cat_additive_group.Row_effective_date = formatDateString(cat_additive_group.Row_effective_date)
	cat_additive_group.Row_expiry_date = formatDateString(cat_additive_group.Row_expiry_date)






	rows, err := stmt.Exec(cat_additive_group.Additive_group_id, cat_additive_group.Active_ind, cat_additive_group.Effective_date, cat_additive_group.Expiry_date, cat_additive_group.Long_name, cat_additive_group.Owner_ba_id, cat_additive_group.Ppdm_guid, cat_additive_group.Remark, cat_additive_group.Short_name, cat_additive_group.Source, cat_additive_group.Source_document_id, cat_additive_group.Substance_id, cat_additive_group.Row_changed_by, cat_additive_group.Row_changed_date, cat_additive_group.Row_created_by, cat_additive_group.Row_created_date, cat_additive_group.Row_effective_date, cat_additive_group.Row_expiry_date, cat_additive_group.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateCatAdditiveGroup(c *fiber.Ctx) error {
	var cat_additive_group dto.Cat_additive_group

	setDefaults(&cat_additive_group)

	if err := json.Unmarshal(c.Body(), &cat_additive_group); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cat_additive_group.Additive_group_id = id

    if cat_additive_group.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cat_additive_group set active_ind = :1, effective_date = :2, expiry_date = :3, long_name = :4, owner_ba_id = :5, ppdm_guid = :6, remark = :7, short_name = :8, source = :9, source_document_id = :10, substance_id = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where additive_group_id = :19")
	if err != nil {
		return err
	}

	cat_additive_group.Row_changed_date = formatDate(cat_additive_group.Row_changed_date)
	cat_additive_group.Effective_date = formatDateString(cat_additive_group.Effective_date)
	cat_additive_group.Expiry_date = formatDateString(cat_additive_group.Expiry_date)
	cat_additive_group.Row_effective_date = formatDateString(cat_additive_group.Row_effective_date)
	cat_additive_group.Row_expiry_date = formatDateString(cat_additive_group.Row_expiry_date)






	rows, err := stmt.Exec(cat_additive_group.Active_ind, cat_additive_group.Effective_date, cat_additive_group.Expiry_date, cat_additive_group.Long_name, cat_additive_group.Owner_ba_id, cat_additive_group.Ppdm_guid, cat_additive_group.Remark, cat_additive_group.Short_name, cat_additive_group.Source, cat_additive_group.Source_document_id, cat_additive_group.Substance_id, cat_additive_group.Row_changed_by, cat_additive_group.Row_changed_date, cat_additive_group.Row_created_by, cat_additive_group.Row_effective_date, cat_additive_group.Row_expiry_date, cat_additive_group.Row_quality, cat_additive_group.Additive_group_id)
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

func PatchCatAdditiveGroup(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cat_additive_group set "
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
	query += " where additive_group_id = :id"

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

func DeleteCatAdditiveGroup(c *fiber.Ctx) error {
	id := c.Params("id")
	var cat_additive_group dto.Cat_additive_group
	cat_additive_group.Additive_group_id = id

	stmt, err := db.Prepare("delete from cat_additive_group where additive_group_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cat_additive_group.Additive_group_id)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetCatAdditiveType(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cat_additive_type")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cat_additive_type

	for rows.Next() {
		var cat_additive_type dto.Cat_additive_type
		if err := rows.Scan(&cat_additive_type.Additive_type, &cat_additive_type.Abbreviation, &cat_additive_type.Active_ind, &cat_additive_type.Additive_group, &cat_additive_type.Effective_date, &cat_additive_type.Expiry_date, &cat_additive_type.Long_name, &cat_additive_type.Ppdm_guid, &cat_additive_type.Remark, &cat_additive_type.Short_name, &cat_additive_type.Source, &cat_additive_type.Row_changed_by, &cat_additive_type.Row_changed_date, &cat_additive_type.Row_created_by, &cat_additive_type.Row_created_date, &cat_additive_type.Row_effective_date, &cat_additive_type.Row_expiry_date, &cat_additive_type.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cat_additive_type to result
		result = append(result, cat_additive_type)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetCatAdditiveType(c *fiber.Ctx) error {
	var cat_additive_type dto.Cat_additive_type

	setDefaults(&cat_additive_type)

	if err := json.Unmarshal(c.Body(), &cat_additive_type); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cat_additive_type values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	cat_additive_type.Row_created_date = formatDate(cat_additive_type.Row_created_date)
	cat_additive_type.Row_changed_date = formatDate(cat_additive_type.Row_changed_date)
	cat_additive_type.Effective_date = formatDateString(cat_additive_type.Effective_date)
	cat_additive_type.Expiry_date = formatDateString(cat_additive_type.Expiry_date)
	cat_additive_type.Row_effective_date = formatDateString(cat_additive_type.Row_effective_date)
	cat_additive_type.Row_expiry_date = formatDateString(cat_additive_type.Row_expiry_date)






	rows, err := stmt.Exec(cat_additive_type.Additive_type, cat_additive_type.Abbreviation, cat_additive_type.Active_ind, cat_additive_type.Additive_group, cat_additive_type.Effective_date, cat_additive_type.Expiry_date, cat_additive_type.Long_name, cat_additive_type.Ppdm_guid, cat_additive_type.Remark, cat_additive_type.Short_name, cat_additive_type.Source, cat_additive_type.Row_changed_by, cat_additive_type.Row_changed_date, cat_additive_type.Row_created_by, cat_additive_type.Row_created_date, cat_additive_type.Row_effective_date, cat_additive_type.Row_expiry_date, cat_additive_type.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateCatAdditiveType(c *fiber.Ctx) error {
	var cat_additive_type dto.Cat_additive_type

	setDefaults(&cat_additive_type)

	if err := json.Unmarshal(c.Body(), &cat_additive_type); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cat_additive_type.Additive_type = id

    if cat_additive_type.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cat_additive_type set abbreviation = :1, active_ind = :2, additive_group = :3, effective_date = :4, expiry_date = :5, long_name = :6, ppdm_guid = :7, remark = :8, short_name = :9, source = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where additive_type = :18")
	if err != nil {
		return err
	}

	cat_additive_type.Row_changed_date = formatDate(cat_additive_type.Row_changed_date)
	cat_additive_type.Effective_date = formatDateString(cat_additive_type.Effective_date)
	cat_additive_type.Expiry_date = formatDateString(cat_additive_type.Expiry_date)
	cat_additive_type.Row_effective_date = formatDateString(cat_additive_type.Row_effective_date)
	cat_additive_type.Row_expiry_date = formatDateString(cat_additive_type.Row_expiry_date)






	rows, err := stmt.Exec(cat_additive_type.Abbreviation, cat_additive_type.Active_ind, cat_additive_type.Additive_group, cat_additive_type.Effective_date, cat_additive_type.Expiry_date, cat_additive_type.Long_name, cat_additive_type.Ppdm_guid, cat_additive_type.Remark, cat_additive_type.Short_name, cat_additive_type.Source, cat_additive_type.Row_changed_by, cat_additive_type.Row_changed_date, cat_additive_type.Row_created_by, cat_additive_type.Row_effective_date, cat_additive_type.Row_expiry_date, cat_additive_type.Row_quality, cat_additive_type.Additive_type)
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

func PatchCatAdditiveType(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cat_additive_type set "
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
	query += " where additive_type = :id"

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

func DeleteCatAdditiveType(c *fiber.Ctx) error {
	id := c.Params("id")
	var cat_additive_type dto.Cat_additive_type
	cat_additive_type.Additive_type = id

	stmt, err := db.Prepare("delete from cat_additive_type where additive_type = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cat_additive_type.Additive_type)
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



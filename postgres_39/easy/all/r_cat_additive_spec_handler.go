package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRCatAdditiveSpec(c *fiber.Ctx) error {
	rows, err := db.Query("select * from r_cat_additive_spec")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.R_cat_additive_spec

	for rows.Next() {
		var r_cat_additive_spec dto.R_cat_additive_spec
		if err := rows.Scan(&r_cat_additive_spec.Additive_spec_type, &r_cat_additive_spec.Abbreviation, &r_cat_additive_spec.Active_ind, &r_cat_additive_spec.Effective_date, &r_cat_additive_spec.Expiry_date, &r_cat_additive_spec.Long_name, &r_cat_additive_spec.Ppdm_guid, &r_cat_additive_spec.Property_set_id, &r_cat_additive_spec.Remark, &r_cat_additive_spec.Short_name, &r_cat_additive_spec.Source, &r_cat_additive_spec.Row_changed_by, &r_cat_additive_spec.Row_changed_date, &r_cat_additive_spec.Row_created_by, &r_cat_additive_spec.Row_created_date, &r_cat_additive_spec.Row_effective_date, &r_cat_additive_spec.Row_expiry_date, &r_cat_additive_spec.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append r_cat_additive_spec to result
		result = append(result, r_cat_additive_spec)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRCatAdditiveSpec(c *fiber.Ctx) error {
	var r_cat_additive_spec dto.R_cat_additive_spec

	setDefaults(&r_cat_additive_spec)

	if err := json.Unmarshal(c.Body(), &r_cat_additive_spec); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into r_cat_additive_spec values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	r_cat_additive_spec.Row_created_date = formatDate(r_cat_additive_spec.Row_created_date)
	r_cat_additive_spec.Row_changed_date = formatDate(r_cat_additive_spec.Row_changed_date)
	r_cat_additive_spec.Effective_date = formatDateString(r_cat_additive_spec.Effective_date)
	r_cat_additive_spec.Expiry_date = formatDateString(r_cat_additive_spec.Expiry_date)
	r_cat_additive_spec.Row_effective_date = formatDateString(r_cat_additive_spec.Row_effective_date)
	r_cat_additive_spec.Row_expiry_date = formatDateString(r_cat_additive_spec.Row_expiry_date)






	rows, err := stmt.Exec(r_cat_additive_spec.Additive_spec_type, r_cat_additive_spec.Abbreviation, r_cat_additive_spec.Active_ind, r_cat_additive_spec.Effective_date, r_cat_additive_spec.Expiry_date, r_cat_additive_spec.Long_name, r_cat_additive_spec.Ppdm_guid, r_cat_additive_spec.Property_set_id, r_cat_additive_spec.Remark, r_cat_additive_spec.Short_name, r_cat_additive_spec.Source, r_cat_additive_spec.Row_changed_by, r_cat_additive_spec.Row_changed_date, r_cat_additive_spec.Row_created_by, r_cat_additive_spec.Row_created_date, r_cat_additive_spec.Row_effective_date, r_cat_additive_spec.Row_expiry_date, r_cat_additive_spec.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRCatAdditiveSpec(c *fiber.Ctx) error {
	var r_cat_additive_spec dto.R_cat_additive_spec

	setDefaults(&r_cat_additive_spec)

	if err := json.Unmarshal(c.Body(), &r_cat_additive_spec); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	r_cat_additive_spec.Additive_spec_type = id

    if r_cat_additive_spec.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update r_cat_additive_spec set abbreviation = :1, active_ind = :2, effective_date = :3, expiry_date = :4, long_name = :5, ppdm_guid = :6, property_set_id = :7, remark = :8, short_name = :9, source = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where additive_spec_type = :18")
	if err != nil {
		return err
	}

	r_cat_additive_spec.Row_changed_date = formatDate(r_cat_additive_spec.Row_changed_date)
	r_cat_additive_spec.Effective_date = formatDateString(r_cat_additive_spec.Effective_date)
	r_cat_additive_spec.Expiry_date = formatDateString(r_cat_additive_spec.Expiry_date)
	r_cat_additive_spec.Row_effective_date = formatDateString(r_cat_additive_spec.Row_effective_date)
	r_cat_additive_spec.Row_expiry_date = formatDateString(r_cat_additive_spec.Row_expiry_date)






	rows, err := stmt.Exec(r_cat_additive_spec.Abbreviation, r_cat_additive_spec.Active_ind, r_cat_additive_spec.Effective_date, r_cat_additive_spec.Expiry_date, r_cat_additive_spec.Long_name, r_cat_additive_spec.Ppdm_guid, r_cat_additive_spec.Property_set_id, r_cat_additive_spec.Remark, r_cat_additive_spec.Short_name, r_cat_additive_spec.Source, r_cat_additive_spec.Row_changed_by, r_cat_additive_spec.Row_changed_date, r_cat_additive_spec.Row_created_by, r_cat_additive_spec.Row_effective_date, r_cat_additive_spec.Row_expiry_date, r_cat_additive_spec.Row_quality, r_cat_additive_spec.Additive_spec_type)
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

func PatchRCatAdditiveSpec(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update r_cat_additive_spec set "
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
	query += " where additive_spec_type = :id"

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

func DeleteRCatAdditiveSpec(c *fiber.Ctx) error {
	id := c.Params("id")
	var r_cat_additive_spec dto.R_cat_additive_spec
	r_cat_additive_spec.Additive_spec_type = id

	stmt, err := db.Prepare("delete from r_cat_additive_spec where additive_spec_type = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(r_cat_additive_spec.Additive_spec_type)
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



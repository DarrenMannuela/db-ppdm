package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetCatAdditiveXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cat_additive_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cat_additive_xref

	for rows.Next() {
		var cat_additive_xref dto.Cat_additive_xref
		if err := rows.Scan(&cat_additive_xref.Catalogue_additive_id, &cat_additive_xref.Catalogue_additive_id2, &cat_additive_xref.Xref_obs_no, &cat_additive_xref.Active_ind, &cat_additive_xref.Additive_xref_type, &cat_additive_xref.Effective_date, &cat_additive_xref.Expiry_date, &cat_additive_xref.Ppdm_guid, &cat_additive_xref.Remark, &cat_additive_xref.Source, &cat_additive_xref.Row_changed_by, &cat_additive_xref.Row_changed_date, &cat_additive_xref.Row_created_by, &cat_additive_xref.Row_created_date, &cat_additive_xref.Row_effective_date, &cat_additive_xref.Row_expiry_date, &cat_additive_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cat_additive_xref to result
		result = append(result, cat_additive_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetCatAdditiveXref(c *fiber.Ctx) error {
	var cat_additive_xref dto.Cat_additive_xref

	setDefaults(&cat_additive_xref)

	if err := json.Unmarshal(c.Body(), &cat_additive_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cat_additive_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	cat_additive_xref.Row_created_date = formatDate(cat_additive_xref.Row_created_date)
	cat_additive_xref.Row_changed_date = formatDate(cat_additive_xref.Row_changed_date)
	cat_additive_xref.Effective_date = formatDateString(cat_additive_xref.Effective_date)
	cat_additive_xref.Expiry_date = formatDateString(cat_additive_xref.Expiry_date)
	cat_additive_xref.Row_effective_date = formatDateString(cat_additive_xref.Row_effective_date)
	cat_additive_xref.Row_expiry_date = formatDateString(cat_additive_xref.Row_expiry_date)






	rows, err := stmt.Exec(cat_additive_xref.Catalogue_additive_id, cat_additive_xref.Catalogue_additive_id2, cat_additive_xref.Xref_obs_no, cat_additive_xref.Active_ind, cat_additive_xref.Additive_xref_type, cat_additive_xref.Effective_date, cat_additive_xref.Expiry_date, cat_additive_xref.Ppdm_guid, cat_additive_xref.Remark, cat_additive_xref.Source, cat_additive_xref.Row_changed_by, cat_additive_xref.Row_changed_date, cat_additive_xref.Row_created_by, cat_additive_xref.Row_created_date, cat_additive_xref.Row_effective_date, cat_additive_xref.Row_expiry_date, cat_additive_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateCatAdditiveXref(c *fiber.Ctx) error {
	var cat_additive_xref dto.Cat_additive_xref

	setDefaults(&cat_additive_xref)

	if err := json.Unmarshal(c.Body(), &cat_additive_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cat_additive_xref.Catalogue_additive_id = id

    if cat_additive_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cat_additive_xref set catalogue_additive_id2 = :1, xref_obs_no = :2, active_ind = :3, additive_xref_type = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where catalogue_additive_id = :17")
	if err != nil {
		return err
	}

	cat_additive_xref.Row_changed_date = formatDate(cat_additive_xref.Row_changed_date)
	cat_additive_xref.Effective_date = formatDateString(cat_additive_xref.Effective_date)
	cat_additive_xref.Expiry_date = formatDateString(cat_additive_xref.Expiry_date)
	cat_additive_xref.Row_effective_date = formatDateString(cat_additive_xref.Row_effective_date)
	cat_additive_xref.Row_expiry_date = formatDateString(cat_additive_xref.Row_expiry_date)






	rows, err := stmt.Exec(cat_additive_xref.Catalogue_additive_id2, cat_additive_xref.Xref_obs_no, cat_additive_xref.Active_ind, cat_additive_xref.Additive_xref_type, cat_additive_xref.Effective_date, cat_additive_xref.Expiry_date, cat_additive_xref.Ppdm_guid, cat_additive_xref.Remark, cat_additive_xref.Source, cat_additive_xref.Row_changed_by, cat_additive_xref.Row_changed_date, cat_additive_xref.Row_created_by, cat_additive_xref.Row_effective_date, cat_additive_xref.Row_expiry_date, cat_additive_xref.Row_quality, cat_additive_xref.Catalogue_additive_id)
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

func PatchCatAdditiveXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cat_additive_xref set "
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
	query += " where catalogue_additive_id = :id"

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

func DeleteCatAdditiveXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var cat_additive_xref dto.Cat_additive_xref
	cat_additive_xref.Catalogue_additive_id = id

	stmt, err := db.Prepare("delete from cat_additive_xref where catalogue_additive_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cat_additive_xref.Catalogue_additive_id)
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



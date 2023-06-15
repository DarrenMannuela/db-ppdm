package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetCatAdditiveAllowance(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cat_additive_allowance")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cat_additive_allowance

	for rows.Next() {
		var cat_additive_allowance dto.Cat_additive_allowance
		if err := rows.Scan(&cat_additive_allowance.Allowance_id, &cat_additive_allowance.Active_ind, &cat_additive_allowance.Additive_group_id, &cat_additive_allowance.Allowed_ind, &cat_additive_allowance.Business_associate_id, &cat_additive_allowance.Catalogue_additive_id, &cat_additive_allowance.Disallowed_ind, &cat_additive_allowance.Effective_date, &cat_additive_allowance.Expiry_date, &cat_additive_allowance.Ppdm_guid, &cat_additive_allowance.Remark, &cat_additive_allowance.Source, &cat_additive_allowance.Source_document_id, &cat_additive_allowance.Substance_id, &cat_additive_allowance.Use_condition, &cat_additive_allowance.Row_changed_by, &cat_additive_allowance.Row_changed_date, &cat_additive_allowance.Row_created_by, &cat_additive_allowance.Row_created_date, &cat_additive_allowance.Row_effective_date, &cat_additive_allowance.Row_expiry_date, &cat_additive_allowance.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cat_additive_allowance to result
		result = append(result, cat_additive_allowance)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetCatAdditiveAllowance(c *fiber.Ctx) error {
	var cat_additive_allowance dto.Cat_additive_allowance

	setDefaults(&cat_additive_allowance)

	if err := json.Unmarshal(c.Body(), &cat_additive_allowance); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cat_additive_allowance values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	cat_additive_allowance.Row_created_date = formatDate(cat_additive_allowance.Row_created_date)
	cat_additive_allowance.Row_changed_date = formatDate(cat_additive_allowance.Row_changed_date)
	cat_additive_allowance.Effective_date = formatDateString(cat_additive_allowance.Effective_date)
	cat_additive_allowance.Expiry_date = formatDateString(cat_additive_allowance.Expiry_date)
	cat_additive_allowance.Row_effective_date = formatDateString(cat_additive_allowance.Row_effective_date)
	cat_additive_allowance.Row_expiry_date = formatDateString(cat_additive_allowance.Row_expiry_date)






	rows, err := stmt.Exec(cat_additive_allowance.Allowance_id, cat_additive_allowance.Active_ind, cat_additive_allowance.Additive_group_id, cat_additive_allowance.Allowed_ind, cat_additive_allowance.Business_associate_id, cat_additive_allowance.Catalogue_additive_id, cat_additive_allowance.Disallowed_ind, cat_additive_allowance.Effective_date, cat_additive_allowance.Expiry_date, cat_additive_allowance.Ppdm_guid, cat_additive_allowance.Remark, cat_additive_allowance.Source, cat_additive_allowance.Source_document_id, cat_additive_allowance.Substance_id, cat_additive_allowance.Use_condition, cat_additive_allowance.Row_changed_by, cat_additive_allowance.Row_changed_date, cat_additive_allowance.Row_created_by, cat_additive_allowance.Row_created_date, cat_additive_allowance.Row_effective_date, cat_additive_allowance.Row_expiry_date, cat_additive_allowance.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateCatAdditiveAllowance(c *fiber.Ctx) error {
	var cat_additive_allowance dto.Cat_additive_allowance

	setDefaults(&cat_additive_allowance)

	if err := json.Unmarshal(c.Body(), &cat_additive_allowance); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cat_additive_allowance.Allowance_id = id

    if cat_additive_allowance.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cat_additive_allowance set active_ind = :1, additive_group_id = :2, allowed_ind = :3, business_associate_id = :4, catalogue_additive_id = :5, disallowed_ind = :6, effective_date = :7, expiry_date = :8, ppdm_guid = :9, remark = :10, source = :11, source_document_id = :12, substance_id = :13, use_condition = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where allowance_id = :22")
	if err != nil {
		return err
	}

	cat_additive_allowance.Row_changed_date = formatDate(cat_additive_allowance.Row_changed_date)
	cat_additive_allowance.Effective_date = formatDateString(cat_additive_allowance.Effective_date)
	cat_additive_allowance.Expiry_date = formatDateString(cat_additive_allowance.Expiry_date)
	cat_additive_allowance.Row_effective_date = formatDateString(cat_additive_allowance.Row_effective_date)
	cat_additive_allowance.Row_expiry_date = formatDateString(cat_additive_allowance.Row_expiry_date)






	rows, err := stmt.Exec(cat_additive_allowance.Active_ind, cat_additive_allowance.Additive_group_id, cat_additive_allowance.Allowed_ind, cat_additive_allowance.Business_associate_id, cat_additive_allowance.Catalogue_additive_id, cat_additive_allowance.Disallowed_ind, cat_additive_allowance.Effective_date, cat_additive_allowance.Expiry_date, cat_additive_allowance.Ppdm_guid, cat_additive_allowance.Remark, cat_additive_allowance.Source, cat_additive_allowance.Source_document_id, cat_additive_allowance.Substance_id, cat_additive_allowance.Use_condition, cat_additive_allowance.Row_changed_by, cat_additive_allowance.Row_changed_date, cat_additive_allowance.Row_created_by, cat_additive_allowance.Row_effective_date, cat_additive_allowance.Row_expiry_date, cat_additive_allowance.Row_quality, cat_additive_allowance.Allowance_id)
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

func PatchCatAdditiveAllowance(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cat_additive_allowance set "
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
	query += " where allowance_id = :id"

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

func DeleteCatAdditiveAllowance(c *fiber.Ctx) error {
	id := c.Params("id")
	var cat_additive_allowance dto.Cat_additive_allowance
	cat_additive_allowance.Allowance_id = id

	stmt, err := db.Prepare("delete from cat_additive_allowance where allowance_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cat_additive_allowance.Allowance_id)
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



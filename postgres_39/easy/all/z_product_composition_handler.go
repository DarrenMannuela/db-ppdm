package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetZProductComposition(c *fiber.Ctx) error {
	rows, err := db.Query("select * from z_product_composition")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Z_product_composition

	for rows.Next() {
		var z_product_composition dto.Z_product_composition
		if err := rows.Scan(&z_product_composition.Conversion_id, &z_product_composition.Active_ind, &z_product_composition.Defined_by_ba_id, &z_product_composition.Effective_date, &z_product_composition.Expiry_date, &z_product_composition.Formula, &z_product_composition.Ppdm_guid, &z_product_composition.Product_component_type, &z_product_composition.Product_type, &z_product_composition.Remark, &z_product_composition.Source, &z_product_composition.Source_document_id, &z_product_composition.Subproduct_type, &z_product_composition.Row_changed_by, &z_product_composition.Row_changed_date, &z_product_composition.Row_created_by, &z_product_composition.Row_created_date, &z_product_composition.Row_effective_date, &z_product_composition.Row_expiry_date, &z_product_composition.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append z_product_composition to result
		result = append(result, z_product_composition)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetZProductComposition(c *fiber.Ctx) error {
	var z_product_composition dto.Z_product_composition

	setDefaults(&z_product_composition)

	if err := json.Unmarshal(c.Body(), &z_product_composition); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into z_product_composition values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	z_product_composition.Row_created_date = formatDate(z_product_composition.Row_created_date)
	z_product_composition.Row_changed_date = formatDate(z_product_composition.Row_changed_date)
	z_product_composition.Effective_date = formatDateString(z_product_composition.Effective_date)
	z_product_composition.Expiry_date = formatDateString(z_product_composition.Expiry_date)
	z_product_composition.Row_effective_date = formatDateString(z_product_composition.Row_effective_date)
	z_product_composition.Row_expiry_date = formatDateString(z_product_composition.Row_expiry_date)






	rows, err := stmt.Exec(z_product_composition.Conversion_id, z_product_composition.Active_ind, z_product_composition.Defined_by_ba_id, z_product_composition.Effective_date, z_product_composition.Expiry_date, z_product_composition.Formula, z_product_composition.Ppdm_guid, z_product_composition.Product_component_type, z_product_composition.Product_type, z_product_composition.Remark, z_product_composition.Source, z_product_composition.Source_document_id, z_product_composition.Subproduct_type, z_product_composition.Row_changed_by, z_product_composition.Row_changed_date, z_product_composition.Row_created_by, z_product_composition.Row_created_date, z_product_composition.Row_effective_date, z_product_composition.Row_expiry_date, z_product_composition.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateZProductComposition(c *fiber.Ctx) error {
	var z_product_composition dto.Z_product_composition

	setDefaults(&z_product_composition)

	if err := json.Unmarshal(c.Body(), &z_product_composition); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	z_product_composition.Conversion_id = id

    if z_product_composition.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update z_product_composition set active_ind = :1, defined_by_ba_id = :2, effective_date = :3, expiry_date = :4, formula = :5, ppdm_guid = :6, product_component_type = :7, product_type = :8, remark = :9, source = :10, source_document_id = :11, subproduct_type = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where conversion_id = :20")
	if err != nil {
		return err
	}

	z_product_composition.Row_changed_date = formatDate(z_product_composition.Row_changed_date)
	z_product_composition.Effective_date = formatDateString(z_product_composition.Effective_date)
	z_product_composition.Expiry_date = formatDateString(z_product_composition.Expiry_date)
	z_product_composition.Row_effective_date = formatDateString(z_product_composition.Row_effective_date)
	z_product_composition.Row_expiry_date = formatDateString(z_product_composition.Row_expiry_date)






	rows, err := stmt.Exec(z_product_composition.Active_ind, z_product_composition.Defined_by_ba_id, z_product_composition.Effective_date, z_product_composition.Expiry_date, z_product_composition.Formula, z_product_composition.Ppdm_guid, z_product_composition.Product_component_type, z_product_composition.Product_type, z_product_composition.Remark, z_product_composition.Source, z_product_composition.Source_document_id, z_product_composition.Subproduct_type, z_product_composition.Row_changed_by, z_product_composition.Row_changed_date, z_product_composition.Row_created_by, z_product_composition.Row_effective_date, z_product_composition.Row_expiry_date, z_product_composition.Row_quality, z_product_composition.Conversion_id)
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

func PatchZProductComposition(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update z_product_composition set "
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
	query += " where conversion_id = :id"

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

func DeleteZProductComposition(c *fiber.Ctx) error {
	id := c.Params("id")
	var z_product_composition dto.Z_product_composition
	z_product_composition.Conversion_id = id

	stmt, err := db.Prepare("delete from z_product_composition where conversion_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(z_product_composition.Conversion_id)
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



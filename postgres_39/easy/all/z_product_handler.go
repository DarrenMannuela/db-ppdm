package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetZProduct(c *fiber.Ctx) error {
	rows, err := db.Query("select * from z_product")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Z_product

	for rows.Next() {
		var z_product dto.Z_product
		if err := rows.Scan(&z_product.Product_type, &z_product.Abbreviation, &z_product.Active_ind, &z_product.Conversion_quantity, &z_product.Effective_date, &z_product.Expiry_date, &z_product.Long_name, &z_product.Ppdm_guid, &z_product.Preferred_uom, &z_product.Property_set_id, &z_product.Remark, &z_product.Short_name, &z_product.Source, &z_product.Row_changed_by, &z_product.Row_changed_date, &z_product.Row_created_by, &z_product.Row_created_date, &z_product.Row_effective_date, &z_product.Row_expiry_date, &z_product.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append z_product to result
		result = append(result, z_product)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetZProduct(c *fiber.Ctx) error {
	var z_product dto.Z_product

	setDefaults(&z_product)

	if err := json.Unmarshal(c.Body(), &z_product); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into z_product values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	z_product.Row_created_date = formatDate(z_product.Row_created_date)
	z_product.Row_changed_date = formatDate(z_product.Row_changed_date)
	z_product.Effective_date = formatDateString(z_product.Effective_date)
	z_product.Expiry_date = formatDateString(z_product.Expiry_date)
	z_product.Row_effective_date = formatDateString(z_product.Row_effective_date)
	z_product.Row_expiry_date = formatDateString(z_product.Row_expiry_date)






	rows, err := stmt.Exec(z_product.Product_type, z_product.Abbreviation, z_product.Active_ind, z_product.Conversion_quantity, z_product.Effective_date, z_product.Expiry_date, z_product.Long_name, z_product.Ppdm_guid, z_product.Preferred_uom, z_product.Property_set_id, z_product.Remark, z_product.Short_name, z_product.Source, z_product.Row_changed_by, z_product.Row_changed_date, z_product.Row_created_by, z_product.Row_created_date, z_product.Row_effective_date, z_product.Row_expiry_date, z_product.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateZProduct(c *fiber.Ctx) error {
	var z_product dto.Z_product

	setDefaults(&z_product)

	if err := json.Unmarshal(c.Body(), &z_product); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	z_product.Product_type = id

    if z_product.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update z_product set abbreviation = :1, active_ind = :2, conversion_quantity = :3, effective_date = :4, expiry_date = :5, long_name = :6, ppdm_guid = :7, preferred_uom = :8, property_set_id = :9, remark = :10, short_name = :11, source = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where product_type = :20")
	if err != nil {
		return err
	}

	z_product.Row_changed_date = formatDate(z_product.Row_changed_date)
	z_product.Effective_date = formatDateString(z_product.Effective_date)
	z_product.Expiry_date = formatDateString(z_product.Expiry_date)
	z_product.Row_effective_date = formatDateString(z_product.Row_effective_date)
	z_product.Row_expiry_date = formatDateString(z_product.Row_expiry_date)






	rows, err := stmt.Exec(z_product.Abbreviation, z_product.Active_ind, z_product.Conversion_quantity, z_product.Effective_date, z_product.Expiry_date, z_product.Long_name, z_product.Ppdm_guid, z_product.Preferred_uom, z_product.Property_set_id, z_product.Remark, z_product.Short_name, z_product.Source, z_product.Row_changed_by, z_product.Row_changed_date, z_product.Row_created_by, z_product.Row_effective_date, z_product.Row_expiry_date, z_product.Row_quality, z_product.Product_type)
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

func PatchZProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update z_product set "
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
	query += " where product_type = :id"

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

func DeleteZProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var z_product dto.Z_product
	z_product.Product_type = id

	stmt, err := db.Prepare("delete from z_product where product_type = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(z_product.Product_type)
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



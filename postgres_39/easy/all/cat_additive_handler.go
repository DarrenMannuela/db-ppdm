package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetCatAdditive(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cat_additive")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cat_additive

	for rows.Next() {
		var cat_additive dto.Cat_additive
		if err := rows.Scan(&cat_additive.Catalogue_additive_id, &cat_additive.Active_ind, &cat_additive.Additive_group, &cat_additive.Additive_name, &cat_additive.Additive_type, &cat_additive.Effective_date, &cat_additive.Expiry_date, &cat_additive.Manufacturer, &cat_additive.Ppdm_guid, &cat_additive.Purchase_quantity, &cat_additive.Purchase_quantity_type, &cat_additive.Purchase_quantity_uom, &cat_additive.Remark, &cat_additive.Source, &cat_additive.Supplier, &cat_additive.Upc_code, &cat_additive.Row_changed_by, &cat_additive.Row_changed_date, &cat_additive.Row_created_by, &cat_additive.Row_created_date, &cat_additive.Row_effective_date, &cat_additive.Row_expiry_date, &cat_additive.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cat_additive to result
		result = append(result, cat_additive)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetCatAdditive(c *fiber.Ctx) error {
	var cat_additive dto.Cat_additive

	setDefaults(&cat_additive)

	if err := json.Unmarshal(c.Body(), &cat_additive); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cat_additive values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	cat_additive.Row_created_date = formatDate(cat_additive.Row_created_date)
	cat_additive.Row_changed_date = formatDate(cat_additive.Row_changed_date)
	cat_additive.Effective_date = formatDateString(cat_additive.Effective_date)
	cat_additive.Expiry_date = formatDateString(cat_additive.Expiry_date)
	cat_additive.Row_effective_date = formatDateString(cat_additive.Row_effective_date)
	cat_additive.Row_expiry_date = formatDateString(cat_additive.Row_expiry_date)






	rows, err := stmt.Exec(cat_additive.Catalogue_additive_id, cat_additive.Active_ind, cat_additive.Additive_group, cat_additive.Additive_name, cat_additive.Additive_type, cat_additive.Effective_date, cat_additive.Expiry_date, cat_additive.Manufacturer, cat_additive.Ppdm_guid, cat_additive.Purchase_quantity, cat_additive.Purchase_quantity_type, cat_additive.Purchase_quantity_uom, cat_additive.Remark, cat_additive.Source, cat_additive.Supplier, cat_additive.Upc_code, cat_additive.Row_changed_by, cat_additive.Row_changed_date, cat_additive.Row_created_by, cat_additive.Row_created_date, cat_additive.Row_effective_date, cat_additive.Row_expiry_date, cat_additive.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateCatAdditive(c *fiber.Ctx) error {
	var cat_additive dto.Cat_additive

	setDefaults(&cat_additive)

	if err := json.Unmarshal(c.Body(), &cat_additive); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cat_additive.Catalogue_additive_id = id

    if cat_additive.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cat_additive set active_ind = :1, additive_group = :2, additive_name = :3, additive_type = :4, effective_date = :5, expiry_date = :6, manufacturer = :7, ppdm_guid = :8, purchase_quantity = :9, purchase_quantity_type = :10, purchase_quantity_uom = :11, remark = :12, source = :13, supplier = :14, upc_code = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where catalogue_additive_id = :23")
	if err != nil {
		return err
	}

	cat_additive.Row_changed_date = formatDate(cat_additive.Row_changed_date)
	cat_additive.Effective_date = formatDateString(cat_additive.Effective_date)
	cat_additive.Expiry_date = formatDateString(cat_additive.Expiry_date)
	cat_additive.Row_effective_date = formatDateString(cat_additive.Row_effective_date)
	cat_additive.Row_expiry_date = formatDateString(cat_additive.Row_expiry_date)






	rows, err := stmt.Exec(cat_additive.Active_ind, cat_additive.Additive_group, cat_additive.Additive_name, cat_additive.Additive_type, cat_additive.Effective_date, cat_additive.Expiry_date, cat_additive.Manufacturer, cat_additive.Ppdm_guid, cat_additive.Purchase_quantity, cat_additive.Purchase_quantity_type, cat_additive.Purchase_quantity_uom, cat_additive.Remark, cat_additive.Source, cat_additive.Supplier, cat_additive.Upc_code, cat_additive.Row_changed_by, cat_additive.Row_changed_date, cat_additive.Row_created_by, cat_additive.Row_effective_date, cat_additive.Row_expiry_date, cat_additive.Row_quality, cat_additive.Catalogue_additive_id)
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

func PatchCatAdditive(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cat_additive set "
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

func DeleteCatAdditive(c *fiber.Ctx) error {
	id := c.Params("id")
	var cat_additive dto.Cat_additive
	cat_additive.Catalogue_additive_id = id

	stmt, err := db.Prepare("delete from cat_additive where catalogue_additive_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cat_additive.Catalogue_additive_id)
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



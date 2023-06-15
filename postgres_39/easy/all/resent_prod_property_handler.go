package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetResentProdProperty(c *fiber.Ctx) error {
	rows, err := db.Query("select * from resent_prod_property")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Resent_prod_property

	for rows.Next() {
		var resent_prod_property dto.Resent_prod_property
		if err := rows.Scan(&resent_prod_property.Resent_id, &resent_prod_property.Reserve_class_id, &resent_prod_property.Property_seq_no, &resent_prod_property.Active_ind, &resent_prod_property.Effective_date, &resent_prod_property.Expiry_date, &resent_prod_property.Heat_content, &resent_prod_property.Heat_content_ouom, &resent_prod_property.Loss_factor, &resent_prod_property.Oil_density, &resent_prod_property.Oil_density_ouom, &resent_prod_property.Ppdm_guid, &resent_prod_property.Product_type, &resent_prod_property.Remark, &resent_prod_property.Source, &resent_prod_property.Sulphur_content, &resent_prod_property.Sulphur_content_ouom, &resent_prod_property.Row_changed_by, &resent_prod_property.Row_changed_date, &resent_prod_property.Row_created_by, &resent_prod_property.Row_created_date, &resent_prod_property.Row_effective_date, &resent_prod_property.Row_expiry_date, &resent_prod_property.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append resent_prod_property to result
		result = append(result, resent_prod_property)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetResentProdProperty(c *fiber.Ctx) error {
	var resent_prod_property dto.Resent_prod_property

	setDefaults(&resent_prod_property)

	if err := json.Unmarshal(c.Body(), &resent_prod_property); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into resent_prod_property values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24)")
	if err != nil {
		return err
	}
	resent_prod_property.Row_created_date = formatDate(resent_prod_property.Row_created_date)
	resent_prod_property.Row_changed_date = formatDate(resent_prod_property.Row_changed_date)
	resent_prod_property.Effective_date = formatDateString(resent_prod_property.Effective_date)
	resent_prod_property.Expiry_date = formatDateString(resent_prod_property.Expiry_date)
	resent_prod_property.Row_effective_date = formatDateString(resent_prod_property.Row_effective_date)
	resent_prod_property.Row_expiry_date = formatDateString(resent_prod_property.Row_expiry_date)






	rows, err := stmt.Exec(resent_prod_property.Resent_id, resent_prod_property.Reserve_class_id, resent_prod_property.Property_seq_no, resent_prod_property.Active_ind, resent_prod_property.Effective_date, resent_prod_property.Expiry_date, resent_prod_property.Heat_content, resent_prod_property.Heat_content_ouom, resent_prod_property.Loss_factor, resent_prod_property.Oil_density, resent_prod_property.Oil_density_ouom, resent_prod_property.Ppdm_guid, resent_prod_property.Product_type, resent_prod_property.Remark, resent_prod_property.Source, resent_prod_property.Sulphur_content, resent_prod_property.Sulphur_content_ouom, resent_prod_property.Row_changed_by, resent_prod_property.Row_changed_date, resent_prod_property.Row_created_by, resent_prod_property.Row_created_date, resent_prod_property.Row_effective_date, resent_prod_property.Row_expiry_date, resent_prod_property.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateResentProdProperty(c *fiber.Ctx) error {
	var resent_prod_property dto.Resent_prod_property

	setDefaults(&resent_prod_property)

	if err := json.Unmarshal(c.Body(), &resent_prod_property); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	resent_prod_property.Resent_id = id

    if resent_prod_property.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update resent_prod_property set reserve_class_id = :1, property_seq_no = :2, active_ind = :3, effective_date = :4, expiry_date = :5, heat_content = :6, heat_content_ouom = :7, loss_factor = :8, oil_density = :9, oil_density_ouom = :10, ppdm_guid = :11, product_type = :12, remark = :13, source = :14, sulphur_content = :15, sulphur_content_ouom = :16, row_changed_by = :17, row_changed_date = :18, row_created_by = :19, row_effective_date = :20, row_expiry_date = :21, row_quality = :22 where resent_id = :24")
	if err != nil {
		return err
	}

	resent_prod_property.Row_changed_date = formatDate(resent_prod_property.Row_changed_date)
	resent_prod_property.Effective_date = formatDateString(resent_prod_property.Effective_date)
	resent_prod_property.Expiry_date = formatDateString(resent_prod_property.Expiry_date)
	resent_prod_property.Row_effective_date = formatDateString(resent_prod_property.Row_effective_date)
	resent_prod_property.Row_expiry_date = formatDateString(resent_prod_property.Row_expiry_date)






	rows, err := stmt.Exec(resent_prod_property.Reserve_class_id, resent_prod_property.Property_seq_no, resent_prod_property.Active_ind, resent_prod_property.Effective_date, resent_prod_property.Expiry_date, resent_prod_property.Heat_content, resent_prod_property.Heat_content_ouom, resent_prod_property.Loss_factor, resent_prod_property.Oil_density, resent_prod_property.Oil_density_ouom, resent_prod_property.Ppdm_guid, resent_prod_property.Product_type, resent_prod_property.Remark, resent_prod_property.Source, resent_prod_property.Sulphur_content, resent_prod_property.Sulphur_content_ouom, resent_prod_property.Row_changed_by, resent_prod_property.Row_changed_date, resent_prod_property.Row_created_by, resent_prod_property.Row_effective_date, resent_prod_property.Row_expiry_date, resent_prod_property.Row_quality, resent_prod_property.Resent_id)
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

func PatchResentProdProperty(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update resent_prod_property set "
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
	query += " where resent_id = :id"

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

func DeleteResentProdProperty(c *fiber.Ctx) error {
	id := c.Params("id")
	var resent_prod_property dto.Resent_prod_property
	resent_prod_property.Resent_id = id

	stmt, err := db.Prepare("delete from resent_prod_property where resent_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(resent_prod_property.Resent_id)
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



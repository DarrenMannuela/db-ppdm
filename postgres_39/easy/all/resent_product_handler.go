package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetResentProduct(c *fiber.Ctx) error {
	rows, err := db.Query("select * from resent_product")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Resent_product

	for rows.Next() {
		var resent_product dto.Resent_product
		if err := rows.Scan(&resent_product.Resent_id, &resent_product.Reserve_class_id, &resent_product.Product_type, &resent_product.Active_ind, &resent_product.Effective_date, &resent_product.Expiry_date, &resent_product.Ppdm_guid, &resent_product.Remark, &resent_product.Report_ind, &resent_product.Source, &resent_product.Row_changed_by, &resent_product.Row_changed_date, &resent_product.Row_created_by, &resent_product.Row_created_date, &resent_product.Row_effective_date, &resent_product.Row_expiry_date, &resent_product.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append resent_product to result
		result = append(result, resent_product)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetResentProduct(c *fiber.Ctx) error {
	var resent_product dto.Resent_product

	setDefaults(&resent_product)

	if err := json.Unmarshal(c.Body(), &resent_product); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into resent_product values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	resent_product.Row_created_date = formatDate(resent_product.Row_created_date)
	resent_product.Row_changed_date = formatDate(resent_product.Row_changed_date)
	resent_product.Effective_date = formatDateString(resent_product.Effective_date)
	resent_product.Expiry_date = formatDateString(resent_product.Expiry_date)
	resent_product.Row_effective_date = formatDateString(resent_product.Row_effective_date)
	resent_product.Row_expiry_date = formatDateString(resent_product.Row_expiry_date)






	rows, err := stmt.Exec(resent_product.Resent_id, resent_product.Reserve_class_id, resent_product.Product_type, resent_product.Active_ind, resent_product.Effective_date, resent_product.Expiry_date, resent_product.Ppdm_guid, resent_product.Remark, resent_product.Report_ind, resent_product.Source, resent_product.Row_changed_by, resent_product.Row_changed_date, resent_product.Row_created_by, resent_product.Row_created_date, resent_product.Row_effective_date, resent_product.Row_expiry_date, resent_product.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateResentProduct(c *fiber.Ctx) error {
	var resent_product dto.Resent_product

	setDefaults(&resent_product)

	if err := json.Unmarshal(c.Body(), &resent_product); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	resent_product.Resent_id = id

    if resent_product.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update resent_product set reserve_class_id = :1, product_type = :2, active_ind = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, report_ind = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where resent_id = :17")
	if err != nil {
		return err
	}

	resent_product.Row_changed_date = formatDate(resent_product.Row_changed_date)
	resent_product.Effective_date = formatDateString(resent_product.Effective_date)
	resent_product.Expiry_date = formatDateString(resent_product.Expiry_date)
	resent_product.Row_effective_date = formatDateString(resent_product.Row_effective_date)
	resent_product.Row_expiry_date = formatDateString(resent_product.Row_expiry_date)






	rows, err := stmt.Exec(resent_product.Reserve_class_id, resent_product.Product_type, resent_product.Active_ind, resent_product.Effective_date, resent_product.Expiry_date, resent_product.Ppdm_guid, resent_product.Remark, resent_product.Report_ind, resent_product.Source, resent_product.Row_changed_by, resent_product.Row_changed_date, resent_product.Row_created_by, resent_product.Row_effective_date, resent_product.Row_expiry_date, resent_product.Row_quality, resent_product.Resent_id)
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

func PatchResentProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update resent_product set "
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

func DeleteResentProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var resent_product dto.Resent_product
	resent_product.Resent_id = id

	stmt, err := db.Prepare("delete from resent_product where resent_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(resent_product.Resent_id)
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



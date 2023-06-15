package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLandSale(c *fiber.Ctx) error {
	rows, err := db.Query("select * from land_sale")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Land_sale

	for rows.Next() {
		var land_sale dto.Land_sale
		if err := rows.Scan(&land_sale.Jurisdiction, &land_sale.Land_sale_number, &land_sale.Active_ind, &land_sale.Close_date, &land_sale.Description, &land_sale.Effective_date, &land_sale.Expiry_date, &land_sale.Land_sale_name, &land_sale.Open_date, &land_sale.Ppdm_guid, &land_sale.Publish_date, &land_sale.Remark, &land_sale.Sale_date, &land_sale.Sold_size, &land_sale.Sold_size_ouom, &land_sale.Source, &land_sale.Total_bonus, &land_sale.Total_bonus_ouom, &land_sale.Total_size, &land_sale.Total_size_ouom, &land_sale.Row_changed_by, &land_sale.Row_changed_date, &land_sale.Row_created_by, &land_sale.Row_created_date, &land_sale.Row_effective_date, &land_sale.Row_expiry_date, &land_sale.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append land_sale to result
		result = append(result, land_sale)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLandSale(c *fiber.Ctx) error {
	var land_sale dto.Land_sale

	setDefaults(&land_sale)

	if err := json.Unmarshal(c.Body(), &land_sale); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into land_sale values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27)")
	if err != nil {
		return err
	}
	land_sale.Row_created_date = formatDate(land_sale.Row_created_date)
	land_sale.Row_changed_date = formatDate(land_sale.Row_changed_date)
	land_sale.Close_date = formatDateString(land_sale.Close_date)
	land_sale.Effective_date = formatDateString(land_sale.Effective_date)
	land_sale.Expiry_date = formatDateString(land_sale.Expiry_date)
	land_sale.Open_date = formatDateString(land_sale.Open_date)
	land_sale.Publish_date = formatDateString(land_sale.Publish_date)
	land_sale.Sale_date = formatDateString(land_sale.Sale_date)
	land_sale.Row_effective_date = formatDateString(land_sale.Row_effective_date)
	land_sale.Row_expiry_date = formatDateString(land_sale.Row_expiry_date)






	rows, err := stmt.Exec(land_sale.Jurisdiction, land_sale.Land_sale_number, land_sale.Active_ind, land_sale.Close_date, land_sale.Description, land_sale.Effective_date, land_sale.Expiry_date, land_sale.Land_sale_name, land_sale.Open_date, land_sale.Ppdm_guid, land_sale.Publish_date, land_sale.Remark, land_sale.Sale_date, land_sale.Sold_size, land_sale.Sold_size_ouom, land_sale.Source, land_sale.Total_bonus, land_sale.Total_bonus_ouom, land_sale.Total_size, land_sale.Total_size_ouom, land_sale.Row_changed_by, land_sale.Row_changed_date, land_sale.Row_created_by, land_sale.Row_created_date, land_sale.Row_effective_date, land_sale.Row_expiry_date, land_sale.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLandSale(c *fiber.Ctx) error {
	var land_sale dto.Land_sale

	setDefaults(&land_sale)

	if err := json.Unmarshal(c.Body(), &land_sale); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	land_sale.Jurisdiction = id

    if land_sale.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update land_sale set land_sale_number = :1, active_ind = :2, close_date = :3, description = :4, effective_date = :5, expiry_date = :6, land_sale_name = :7, open_date = :8, ppdm_guid = :9, publish_date = :10, remark = :11, sale_date = :12, sold_size = :13, sold_size_ouom = :14, source = :15, total_bonus = :16, total_bonus_ouom = :17, total_size = :18, total_size_ouom = :19, row_changed_by = :20, row_changed_date = :21, row_created_by = :22, row_effective_date = :23, row_expiry_date = :24, row_quality = :25 where jurisdiction = :27")
	if err != nil {
		return err
	}

	land_sale.Row_changed_date = formatDate(land_sale.Row_changed_date)
	land_sale.Close_date = formatDateString(land_sale.Close_date)
	land_sale.Effective_date = formatDateString(land_sale.Effective_date)
	land_sale.Expiry_date = formatDateString(land_sale.Expiry_date)
	land_sale.Open_date = formatDateString(land_sale.Open_date)
	land_sale.Publish_date = formatDateString(land_sale.Publish_date)
	land_sale.Sale_date = formatDateString(land_sale.Sale_date)
	land_sale.Row_effective_date = formatDateString(land_sale.Row_effective_date)
	land_sale.Row_expiry_date = formatDateString(land_sale.Row_expiry_date)






	rows, err := stmt.Exec(land_sale.Land_sale_number, land_sale.Active_ind, land_sale.Close_date, land_sale.Description, land_sale.Effective_date, land_sale.Expiry_date, land_sale.Land_sale_name, land_sale.Open_date, land_sale.Ppdm_guid, land_sale.Publish_date, land_sale.Remark, land_sale.Sale_date, land_sale.Sold_size, land_sale.Sold_size_ouom, land_sale.Source, land_sale.Total_bonus, land_sale.Total_bonus_ouom, land_sale.Total_size, land_sale.Total_size_ouom, land_sale.Row_changed_by, land_sale.Row_changed_date, land_sale.Row_created_by, land_sale.Row_effective_date, land_sale.Row_expiry_date, land_sale.Row_quality, land_sale.Jurisdiction)
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

func PatchLandSale(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update land_sale set "
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
		} else if key == "close_date" || key == "effective_date" || key == "expiry_date" || key == "open_date" || key == "publish_date" || key == "sale_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where jurisdiction = :id"

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

func DeleteLandSale(c *fiber.Ctx) error {
	id := c.Params("id")
	var land_sale dto.Land_sale
	land_sale.Jurisdiction = id

	stmt, err := db.Prepare("delete from land_sale where jurisdiction = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(land_sale.Jurisdiction)
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



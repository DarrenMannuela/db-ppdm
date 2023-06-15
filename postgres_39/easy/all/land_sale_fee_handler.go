package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLandSaleFee(c *fiber.Ctx) error {
	rows, err := db.Query("select * from land_sale_fee")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Land_sale_fee

	for rows.Next() {
		var land_sale_fee dto.Land_sale_fee
		if err := rows.Scan(&land_sale_fee.Jurisdiction, &land_sale_fee.Land_sale_number, &land_sale_fee.Rate_schedule_id, &land_sale_fee.Active_ind, &land_sale_fee.Effective_date, &land_sale_fee.Expiry_date, &land_sale_fee.Ppdm_guid, &land_sale_fee.Remark, &land_sale_fee.Source, &land_sale_fee.Row_changed_by, &land_sale_fee.Row_changed_date, &land_sale_fee.Row_created_by, &land_sale_fee.Row_created_date, &land_sale_fee.Row_effective_date, &land_sale_fee.Row_expiry_date, &land_sale_fee.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append land_sale_fee to result
		result = append(result, land_sale_fee)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLandSaleFee(c *fiber.Ctx) error {
	var land_sale_fee dto.Land_sale_fee

	setDefaults(&land_sale_fee)

	if err := json.Unmarshal(c.Body(), &land_sale_fee); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into land_sale_fee values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16)")
	if err != nil {
		return err
	}
	land_sale_fee.Row_created_date = formatDate(land_sale_fee.Row_created_date)
	land_sale_fee.Row_changed_date = formatDate(land_sale_fee.Row_changed_date)
	land_sale_fee.Effective_date = formatDateString(land_sale_fee.Effective_date)
	land_sale_fee.Expiry_date = formatDateString(land_sale_fee.Expiry_date)
	land_sale_fee.Row_effective_date = formatDateString(land_sale_fee.Row_effective_date)
	land_sale_fee.Row_expiry_date = formatDateString(land_sale_fee.Row_expiry_date)






	rows, err := stmt.Exec(land_sale_fee.Jurisdiction, land_sale_fee.Land_sale_number, land_sale_fee.Rate_schedule_id, land_sale_fee.Active_ind, land_sale_fee.Effective_date, land_sale_fee.Expiry_date, land_sale_fee.Ppdm_guid, land_sale_fee.Remark, land_sale_fee.Source, land_sale_fee.Row_changed_by, land_sale_fee.Row_changed_date, land_sale_fee.Row_created_by, land_sale_fee.Row_created_date, land_sale_fee.Row_effective_date, land_sale_fee.Row_expiry_date, land_sale_fee.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLandSaleFee(c *fiber.Ctx) error {
	var land_sale_fee dto.Land_sale_fee

	setDefaults(&land_sale_fee)

	if err := json.Unmarshal(c.Body(), &land_sale_fee); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	land_sale_fee.Jurisdiction = id

    if land_sale_fee.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update land_sale_fee set land_sale_number = :1, rate_schedule_id = :2, active_ind = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, source = :8, row_changed_by = :9, row_changed_date = :10, row_created_by = :11, row_effective_date = :12, row_expiry_date = :13, row_quality = :14 where jurisdiction = :16")
	if err != nil {
		return err
	}

	land_sale_fee.Row_changed_date = formatDate(land_sale_fee.Row_changed_date)
	land_sale_fee.Effective_date = formatDateString(land_sale_fee.Effective_date)
	land_sale_fee.Expiry_date = formatDateString(land_sale_fee.Expiry_date)
	land_sale_fee.Row_effective_date = formatDateString(land_sale_fee.Row_effective_date)
	land_sale_fee.Row_expiry_date = formatDateString(land_sale_fee.Row_expiry_date)






	rows, err := stmt.Exec(land_sale_fee.Land_sale_number, land_sale_fee.Rate_schedule_id, land_sale_fee.Active_ind, land_sale_fee.Effective_date, land_sale_fee.Expiry_date, land_sale_fee.Ppdm_guid, land_sale_fee.Remark, land_sale_fee.Source, land_sale_fee.Row_changed_by, land_sale_fee.Row_changed_date, land_sale_fee.Row_created_by, land_sale_fee.Row_effective_date, land_sale_fee.Row_expiry_date, land_sale_fee.Row_quality, land_sale_fee.Jurisdiction)
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

func PatchLandSaleFee(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update land_sale_fee set "
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

func DeleteLandSaleFee(c *fiber.Ctx) error {
	id := c.Params("id")
	var land_sale_fee dto.Land_sale_fee
	land_sale_fee.Jurisdiction = id

	stmt, err := db.Prepare("delete from land_sale_fee where jurisdiction = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(land_sale_fee.Jurisdiction)
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



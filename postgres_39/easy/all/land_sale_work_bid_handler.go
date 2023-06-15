package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLandSaleWorkBid(c *fiber.Ctx) error {
	rows, err := db.Query("select * from land_sale_work_bid")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Land_sale_work_bid

	for rows.Next() {
		var land_sale_work_bid dto.Land_sale_work_bid
		if err := rows.Scan(&land_sale_work_bid.Jurisdiction, &land_sale_work_bid.Land_sale_number, &land_sale_work_bid.Land_sale_offering_id, &land_sale_work_bid.Land_offering_bid_id, &land_sale_work_bid.Work_bid_id, &land_sale_work_bid.Active_ind, &land_sale_work_bid.Critical_date, &land_sale_work_bid.Currency_conversion, &land_sale_work_bid.Currency_ouom, &land_sale_work_bid.Description, &land_sale_work_bid.Effective_date, &land_sale_work_bid.Expiry_date, &land_sale_work_bid.Ppdm_guid, &land_sale_work_bid.Remark, &land_sale_work_bid.Source, &land_sale_work_bid.Work_bid_count, &land_sale_work_bid.Work_bid_count_uom, &land_sale_work_bid.Work_bid_type, &land_sale_work_bid.Work_bid_value, &land_sale_work_bid.Row_changed_by, &land_sale_work_bid.Row_changed_date, &land_sale_work_bid.Row_created_by, &land_sale_work_bid.Row_created_date, &land_sale_work_bid.Row_effective_date, &land_sale_work_bid.Row_expiry_date, &land_sale_work_bid.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append land_sale_work_bid to result
		result = append(result, land_sale_work_bid)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLandSaleWorkBid(c *fiber.Ctx) error {
	var land_sale_work_bid dto.Land_sale_work_bid

	setDefaults(&land_sale_work_bid)

	if err := json.Unmarshal(c.Body(), &land_sale_work_bid); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into land_sale_work_bid values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26)")
	if err != nil {
		return err
	}
	land_sale_work_bid.Row_created_date = formatDate(land_sale_work_bid.Row_created_date)
	land_sale_work_bid.Row_changed_date = formatDate(land_sale_work_bid.Row_changed_date)
	land_sale_work_bid.Critical_date = formatDateString(land_sale_work_bid.Critical_date)
	land_sale_work_bid.Effective_date = formatDateString(land_sale_work_bid.Effective_date)
	land_sale_work_bid.Expiry_date = formatDateString(land_sale_work_bid.Expiry_date)
	land_sale_work_bid.Row_effective_date = formatDateString(land_sale_work_bid.Row_effective_date)
	land_sale_work_bid.Row_expiry_date = formatDateString(land_sale_work_bid.Row_expiry_date)






	rows, err := stmt.Exec(land_sale_work_bid.Jurisdiction, land_sale_work_bid.Land_sale_number, land_sale_work_bid.Land_sale_offering_id, land_sale_work_bid.Land_offering_bid_id, land_sale_work_bid.Work_bid_id, land_sale_work_bid.Active_ind, land_sale_work_bid.Critical_date, land_sale_work_bid.Currency_conversion, land_sale_work_bid.Currency_ouom, land_sale_work_bid.Description, land_sale_work_bid.Effective_date, land_sale_work_bid.Expiry_date, land_sale_work_bid.Ppdm_guid, land_sale_work_bid.Remark, land_sale_work_bid.Source, land_sale_work_bid.Work_bid_count, land_sale_work_bid.Work_bid_count_uom, land_sale_work_bid.Work_bid_type, land_sale_work_bid.Work_bid_value, land_sale_work_bid.Row_changed_by, land_sale_work_bid.Row_changed_date, land_sale_work_bid.Row_created_by, land_sale_work_bid.Row_created_date, land_sale_work_bid.Row_effective_date, land_sale_work_bid.Row_expiry_date, land_sale_work_bid.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLandSaleWorkBid(c *fiber.Ctx) error {
	var land_sale_work_bid dto.Land_sale_work_bid

	setDefaults(&land_sale_work_bid)

	if err := json.Unmarshal(c.Body(), &land_sale_work_bid); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	land_sale_work_bid.Jurisdiction = id

    if land_sale_work_bid.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update land_sale_work_bid set land_sale_number = :1, land_sale_offering_id = :2, land_offering_bid_id = :3, work_bid_id = :4, active_ind = :5, critical_date = :6, currency_conversion = :7, currency_ouom = :8, description = :9, effective_date = :10, expiry_date = :11, ppdm_guid = :12, remark = :13, source = :14, work_bid_count = :15, work_bid_count_uom = :16, work_bid_type = :17, work_bid_value = :18, row_changed_by = :19, row_changed_date = :20, row_created_by = :21, row_effective_date = :22, row_expiry_date = :23, row_quality = :24 where jurisdiction = :26")
	if err != nil {
		return err
	}

	land_sale_work_bid.Row_changed_date = formatDate(land_sale_work_bid.Row_changed_date)
	land_sale_work_bid.Critical_date = formatDateString(land_sale_work_bid.Critical_date)
	land_sale_work_bid.Effective_date = formatDateString(land_sale_work_bid.Effective_date)
	land_sale_work_bid.Expiry_date = formatDateString(land_sale_work_bid.Expiry_date)
	land_sale_work_bid.Row_effective_date = formatDateString(land_sale_work_bid.Row_effective_date)
	land_sale_work_bid.Row_expiry_date = formatDateString(land_sale_work_bid.Row_expiry_date)






	rows, err := stmt.Exec(land_sale_work_bid.Land_sale_number, land_sale_work_bid.Land_sale_offering_id, land_sale_work_bid.Land_offering_bid_id, land_sale_work_bid.Work_bid_id, land_sale_work_bid.Active_ind, land_sale_work_bid.Critical_date, land_sale_work_bid.Currency_conversion, land_sale_work_bid.Currency_ouom, land_sale_work_bid.Description, land_sale_work_bid.Effective_date, land_sale_work_bid.Expiry_date, land_sale_work_bid.Ppdm_guid, land_sale_work_bid.Remark, land_sale_work_bid.Source, land_sale_work_bid.Work_bid_count, land_sale_work_bid.Work_bid_count_uom, land_sale_work_bid.Work_bid_type, land_sale_work_bid.Work_bid_value, land_sale_work_bid.Row_changed_by, land_sale_work_bid.Row_changed_date, land_sale_work_bid.Row_created_by, land_sale_work_bid.Row_effective_date, land_sale_work_bid.Row_expiry_date, land_sale_work_bid.Row_quality, land_sale_work_bid.Jurisdiction)
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

func PatchLandSaleWorkBid(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update land_sale_work_bid set "
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
		} else if key == "critical_date" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteLandSaleWorkBid(c *fiber.Ctx) error {
	id := c.Params("id")
	var land_sale_work_bid dto.Land_sale_work_bid
	land_sale_work_bid.Jurisdiction = id

	stmt, err := db.Prepare("delete from land_sale_work_bid where jurisdiction = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(land_sale_work_bid.Jurisdiction)
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



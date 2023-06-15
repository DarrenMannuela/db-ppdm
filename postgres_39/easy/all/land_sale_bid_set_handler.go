package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLandSaleBidSet(c *fiber.Ctx) error {
	rows, err := db.Query("select * from land_sale_bid_set")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Land_sale_bid_set

	for rows.Next() {
		var land_sale_bid_set dto.Land_sale_bid_set
		if err := rows.Scan(&land_sale_bid_set.Land_sale_bid_set_id, &land_sale_bid_set.Active_ind, &land_sale_bid_set.Bid_status, &land_sale_bid_set.Confidential_ind, &land_sale_bid_set.Contingency_desc, &land_sale_bid_set.Contingency_ind, &land_sale_bid_set.Effective_date, &land_sale_bid_set.Expiry_date, &land_sale_bid_set.Owner_ba_id, &land_sale_bid_set.Ppdm_guid, &land_sale_bid_set.Remark, &land_sale_bid_set.Set_status_date, &land_sale_bid_set.Source, &land_sale_bid_set.Successful_ind, &land_sale_bid_set.Row_changed_by, &land_sale_bid_set.Row_changed_date, &land_sale_bid_set.Row_created_by, &land_sale_bid_set.Row_created_date, &land_sale_bid_set.Row_effective_date, &land_sale_bid_set.Row_expiry_date, &land_sale_bid_set.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append land_sale_bid_set to result
		result = append(result, land_sale_bid_set)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLandSaleBidSet(c *fiber.Ctx) error {
	var land_sale_bid_set dto.Land_sale_bid_set

	setDefaults(&land_sale_bid_set)

	if err := json.Unmarshal(c.Body(), &land_sale_bid_set); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into land_sale_bid_set values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	land_sale_bid_set.Row_created_date = formatDate(land_sale_bid_set.Row_created_date)
	land_sale_bid_set.Row_changed_date = formatDate(land_sale_bid_set.Row_changed_date)
	land_sale_bid_set.Effective_date = formatDateString(land_sale_bid_set.Effective_date)
	land_sale_bid_set.Expiry_date = formatDateString(land_sale_bid_set.Expiry_date)
	land_sale_bid_set.Set_status_date = formatDateString(land_sale_bid_set.Set_status_date)
	land_sale_bid_set.Row_effective_date = formatDateString(land_sale_bid_set.Row_effective_date)
	land_sale_bid_set.Row_expiry_date = formatDateString(land_sale_bid_set.Row_expiry_date)






	rows, err := stmt.Exec(land_sale_bid_set.Land_sale_bid_set_id, land_sale_bid_set.Active_ind, land_sale_bid_set.Bid_status, land_sale_bid_set.Confidential_ind, land_sale_bid_set.Contingency_desc, land_sale_bid_set.Contingency_ind, land_sale_bid_set.Effective_date, land_sale_bid_set.Expiry_date, land_sale_bid_set.Owner_ba_id, land_sale_bid_set.Ppdm_guid, land_sale_bid_set.Remark, land_sale_bid_set.Set_status_date, land_sale_bid_set.Source, land_sale_bid_set.Successful_ind, land_sale_bid_set.Row_changed_by, land_sale_bid_set.Row_changed_date, land_sale_bid_set.Row_created_by, land_sale_bid_set.Row_created_date, land_sale_bid_set.Row_effective_date, land_sale_bid_set.Row_expiry_date, land_sale_bid_set.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLandSaleBidSet(c *fiber.Ctx) error {
	var land_sale_bid_set dto.Land_sale_bid_set

	setDefaults(&land_sale_bid_set)

	if err := json.Unmarshal(c.Body(), &land_sale_bid_set); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	land_sale_bid_set.Land_sale_bid_set_id = id

    if land_sale_bid_set.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update land_sale_bid_set set active_ind = :1, bid_status = :2, confidential_ind = :3, contingency_desc = :4, contingency_ind = :5, effective_date = :6, expiry_date = :7, owner_ba_id = :8, ppdm_guid = :9, remark = :10, set_status_date = :11, source = :12, successful_ind = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where land_sale_bid_set_id = :21")
	if err != nil {
		return err
	}

	land_sale_bid_set.Row_changed_date = formatDate(land_sale_bid_set.Row_changed_date)
	land_sale_bid_set.Effective_date = formatDateString(land_sale_bid_set.Effective_date)
	land_sale_bid_set.Expiry_date = formatDateString(land_sale_bid_set.Expiry_date)
	land_sale_bid_set.Set_status_date = formatDateString(land_sale_bid_set.Set_status_date)
	land_sale_bid_set.Row_effective_date = formatDateString(land_sale_bid_set.Row_effective_date)
	land_sale_bid_set.Row_expiry_date = formatDateString(land_sale_bid_set.Row_expiry_date)






	rows, err := stmt.Exec(land_sale_bid_set.Active_ind, land_sale_bid_set.Bid_status, land_sale_bid_set.Confidential_ind, land_sale_bid_set.Contingency_desc, land_sale_bid_set.Contingency_ind, land_sale_bid_set.Effective_date, land_sale_bid_set.Expiry_date, land_sale_bid_set.Owner_ba_id, land_sale_bid_set.Ppdm_guid, land_sale_bid_set.Remark, land_sale_bid_set.Set_status_date, land_sale_bid_set.Source, land_sale_bid_set.Successful_ind, land_sale_bid_set.Row_changed_by, land_sale_bid_set.Row_changed_date, land_sale_bid_set.Row_created_by, land_sale_bid_set.Row_effective_date, land_sale_bid_set.Row_expiry_date, land_sale_bid_set.Row_quality, land_sale_bid_set.Land_sale_bid_set_id)
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

func PatchLandSaleBidSet(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update land_sale_bid_set set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "set_status_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where land_sale_bid_set_id = :id"

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

func DeleteLandSaleBidSet(c *fiber.Ctx) error {
	id := c.Params("id")
	var land_sale_bid_set dto.Land_sale_bid_set
	land_sale_bid_set.Land_sale_bid_set_id = id

	stmt, err := db.Prepare("delete from land_sale_bid_set where land_sale_bid_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(land_sale_bid_set.Land_sale_bid_set_id)
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



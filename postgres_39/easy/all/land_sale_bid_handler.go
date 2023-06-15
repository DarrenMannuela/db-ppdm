package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLandSaleBid(c *fiber.Ctx) error {
	rows, err := db.Query("select * from land_sale_bid")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Land_sale_bid

	for rows.Next() {
		var land_sale_bid dto.Land_sale_bid
		if err := rows.Scan(&land_sale_bid.Jurisdiction, &land_sale_bid.Land_sale_number, &land_sale_bid.Land_sale_offering_id, &land_sale_bid.Land_offering_bid_id, &land_sale_bid.Active_ind, &land_sale_bid.Address_for_service, &land_sale_bid.Address_obs_no, &land_sale_bid.Address_source, &land_sale_bid.Bidder, &land_sale_bid.Bidder_type, &land_sale_bid.Bid_status, &land_sale_bid.Bid_status_date, &land_sale_bid.Bid_submit_date, &land_sale_bid.Bid_type, &land_sale_bid.Cash_bid_type, &land_sale_bid.Confidential_ind, &land_sale_bid.Contingency_desc, &land_sale_bid.Contingency_ind, &land_sale_bid.Effective_date, &land_sale_bid.Expiry_date, &land_sale_bid.Payment_instruction_id, &land_sale_bid.Ppdm_guid, &land_sale_bid.Priority_order, &land_sale_bid.Remark, &land_sale_bid.Response, &land_sale_bid.Response_date, &land_sale_bid.Source, &land_sale_bid.Successful_ind, &land_sale_bid.Work_bid_ind, &land_sale_bid.Work_bid_remark, &land_sale_bid.Row_changed_by, &land_sale_bid.Row_changed_date, &land_sale_bid.Row_created_by, &land_sale_bid.Row_created_date, &land_sale_bid.Row_effective_date, &land_sale_bid.Row_expiry_date, &land_sale_bid.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append land_sale_bid to result
		result = append(result, land_sale_bid)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLandSaleBid(c *fiber.Ctx) error {
	var land_sale_bid dto.Land_sale_bid

	setDefaults(&land_sale_bid)

	if err := json.Unmarshal(c.Body(), &land_sale_bid); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into land_sale_bid values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37)")
	if err != nil {
		return err
	}
	land_sale_bid.Row_created_date = formatDate(land_sale_bid.Row_created_date)
	land_sale_bid.Row_changed_date = formatDate(land_sale_bid.Row_changed_date)
	land_sale_bid.Bid_status_date = formatDateString(land_sale_bid.Bid_status_date)
	land_sale_bid.Bid_submit_date = formatDateString(land_sale_bid.Bid_submit_date)
	land_sale_bid.Effective_date = formatDateString(land_sale_bid.Effective_date)
	land_sale_bid.Expiry_date = formatDateString(land_sale_bid.Expiry_date)
	land_sale_bid.Response_date = formatDateString(land_sale_bid.Response_date)
	land_sale_bid.Row_effective_date = formatDateString(land_sale_bid.Row_effective_date)
	land_sale_bid.Row_expiry_date = formatDateString(land_sale_bid.Row_expiry_date)






	rows, err := stmt.Exec(land_sale_bid.Jurisdiction, land_sale_bid.Land_sale_number, land_sale_bid.Land_sale_offering_id, land_sale_bid.Land_offering_bid_id, land_sale_bid.Active_ind, land_sale_bid.Address_for_service, land_sale_bid.Address_obs_no, land_sale_bid.Address_source, land_sale_bid.Bidder, land_sale_bid.Bidder_type, land_sale_bid.Bid_status, land_sale_bid.Bid_status_date, land_sale_bid.Bid_submit_date, land_sale_bid.Bid_type, land_sale_bid.Cash_bid_type, land_sale_bid.Confidential_ind, land_sale_bid.Contingency_desc, land_sale_bid.Contingency_ind, land_sale_bid.Effective_date, land_sale_bid.Expiry_date, land_sale_bid.Payment_instruction_id, land_sale_bid.Ppdm_guid, land_sale_bid.Priority_order, land_sale_bid.Remark, land_sale_bid.Response, land_sale_bid.Response_date, land_sale_bid.Source, land_sale_bid.Successful_ind, land_sale_bid.Work_bid_ind, land_sale_bid.Work_bid_remark, land_sale_bid.Row_changed_by, land_sale_bid.Row_changed_date, land_sale_bid.Row_created_by, land_sale_bid.Row_created_date, land_sale_bid.Row_effective_date, land_sale_bid.Row_expiry_date, land_sale_bid.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLandSaleBid(c *fiber.Ctx) error {
	var land_sale_bid dto.Land_sale_bid

	setDefaults(&land_sale_bid)

	if err := json.Unmarshal(c.Body(), &land_sale_bid); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	land_sale_bid.Jurisdiction = id

    if land_sale_bid.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update land_sale_bid set land_sale_number = :1, land_sale_offering_id = :2, land_offering_bid_id = :3, active_ind = :4, address_for_service = :5, address_obs_no = :6, address_source = :7, bidder = :8, bidder_type = :9, bid_status = :10, bid_status_date = :11, bid_submit_date = :12, bid_type = :13, cash_bid_type = :14, confidential_ind = :15, contingency_desc = :16, contingency_ind = :17, effective_date = :18, expiry_date = :19, payment_instruction_id = :20, ppdm_guid = :21, priority_order = :22, remark = :23, response = :24, response_date = :25, source = :26, successful_ind = :27, work_bid_ind = :28, work_bid_remark = :29, row_changed_by = :30, row_changed_date = :31, row_created_by = :32, row_effective_date = :33, row_expiry_date = :34, row_quality = :35 where jurisdiction = :37")
	if err != nil {
		return err
	}

	land_sale_bid.Row_changed_date = formatDate(land_sale_bid.Row_changed_date)
	land_sale_bid.Bid_status_date = formatDateString(land_sale_bid.Bid_status_date)
	land_sale_bid.Bid_submit_date = formatDateString(land_sale_bid.Bid_submit_date)
	land_sale_bid.Effective_date = formatDateString(land_sale_bid.Effective_date)
	land_sale_bid.Expiry_date = formatDateString(land_sale_bid.Expiry_date)
	land_sale_bid.Response_date = formatDateString(land_sale_bid.Response_date)
	land_sale_bid.Row_effective_date = formatDateString(land_sale_bid.Row_effective_date)
	land_sale_bid.Row_expiry_date = formatDateString(land_sale_bid.Row_expiry_date)






	rows, err := stmt.Exec(land_sale_bid.Land_sale_number, land_sale_bid.Land_sale_offering_id, land_sale_bid.Land_offering_bid_id, land_sale_bid.Active_ind, land_sale_bid.Address_for_service, land_sale_bid.Address_obs_no, land_sale_bid.Address_source, land_sale_bid.Bidder, land_sale_bid.Bidder_type, land_sale_bid.Bid_status, land_sale_bid.Bid_status_date, land_sale_bid.Bid_submit_date, land_sale_bid.Bid_type, land_sale_bid.Cash_bid_type, land_sale_bid.Confidential_ind, land_sale_bid.Contingency_desc, land_sale_bid.Contingency_ind, land_sale_bid.Effective_date, land_sale_bid.Expiry_date, land_sale_bid.Payment_instruction_id, land_sale_bid.Ppdm_guid, land_sale_bid.Priority_order, land_sale_bid.Remark, land_sale_bid.Response, land_sale_bid.Response_date, land_sale_bid.Source, land_sale_bid.Successful_ind, land_sale_bid.Work_bid_ind, land_sale_bid.Work_bid_remark, land_sale_bid.Row_changed_by, land_sale_bid.Row_changed_date, land_sale_bid.Row_created_by, land_sale_bid.Row_effective_date, land_sale_bid.Row_expiry_date, land_sale_bid.Row_quality, land_sale_bid.Jurisdiction)
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

func PatchLandSaleBid(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update land_sale_bid set "
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
		} else if key == "bid_status_date" || key == "bid_submit_date" || key == "effective_date" || key == "expiry_date" || key == "response_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteLandSaleBid(c *fiber.Ctx) error {
	id := c.Params("id")
	var land_sale_bid dto.Land_sale_bid
	land_sale_bid.Jurisdiction = id

	stmt, err := db.Prepare("delete from land_sale_bid where jurisdiction = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(land_sale_bid.Jurisdiction)
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



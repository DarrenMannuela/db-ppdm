package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLandSaleRequest(c *fiber.Ctx) error {
	rows, err := db.Query("select * from land_sale_request")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Land_sale_request

	for rows.Next() {
		var land_sale_request dto.Land_sale_request
		if err := rows.Scan(&land_sale_request.Land_request_id, &land_sale_request.Active_ind, &land_sale_request.Advance_booking_ind, &land_sale_request.Ami_ind, &land_sale_request.Area_id, &land_sale_request.Area_type, &land_sale_request.Bid_ind, &land_sale_request.Broker_ba_id, &land_sale_request.Contact_ba_id, &land_sale_request.Critical_date, &land_sale_request.Drilling_priority_ind, &land_sale_request.Effective_date, &land_sale_request.Expiry_date, &land_sale_request.Granted_right_type, &land_sale_request.Gross_size, &land_sale_request.Gross_size_ouom, &land_sale_request.Investigator, &land_sale_request.Jurisdiction, &land_sale_request.Land_req_status, &land_sale_request.Land_right_id, &land_sale_request.Land_right_subtype, &land_sale_request.Land_sale_number, &land_sale_request.Land_sale_offering_id, &land_sale_request.Lessor_refused_date, &land_sale_request.Originator_ba_id, &land_sale_request.Ppdm_guid, &land_sale_request.Previous_request_id, &land_sale_request.Purpose, &land_sale_request.Reference_num, &land_sale_request.Remark, &land_sale_request.Requested_sale_date, &land_sale_request.Requested_sale_name, &land_sale_request.Requester, &land_sale_request.Request_cancel_date, &land_sale_request.Request_description, &land_sale_request.Request_received_date, &land_sale_request.Request_size, &land_sale_request.Request_size_ouom, &land_sale_request.Request_type, &land_sale_request.Response_date, &land_sale_request.Response_desc, &land_sale_request.Source, &land_sale_request.Submitted_date, &land_sale_request.Term_duration, &land_sale_request.Term_duration_ouom, &land_sale_request.Uwi, &land_sale_request.Well_license_id, &land_sale_request.Well_source, &land_sale_request.Withdrawn_date, &land_sale_request.Row_changed_by, &land_sale_request.Row_changed_date, &land_sale_request.Row_created_by, &land_sale_request.Row_created_date, &land_sale_request.Row_effective_date, &land_sale_request.Row_expiry_date, &land_sale_request.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append land_sale_request to result
		result = append(result, land_sale_request)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLandSaleRequest(c *fiber.Ctx) error {
	var land_sale_request dto.Land_sale_request

	setDefaults(&land_sale_request)

	if err := json.Unmarshal(c.Body(), &land_sale_request); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into land_sale_request values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56)")
	if err != nil {
		return err
	}
	land_sale_request.Row_created_date = formatDate(land_sale_request.Row_created_date)
	land_sale_request.Row_changed_date = formatDate(land_sale_request.Row_changed_date)
	land_sale_request.Critical_date = formatDateString(land_sale_request.Critical_date)
	land_sale_request.Effective_date = formatDateString(land_sale_request.Effective_date)
	land_sale_request.Expiry_date = formatDateString(land_sale_request.Expiry_date)
	land_sale_request.Lessor_refused_date = formatDateString(land_sale_request.Lessor_refused_date)
	land_sale_request.Requested_sale_date = formatDateString(land_sale_request.Requested_sale_date)
	land_sale_request.Request_cancel_date = formatDateString(land_sale_request.Request_cancel_date)
	land_sale_request.Request_received_date = formatDateString(land_sale_request.Request_received_date)
	land_sale_request.Response_date = formatDateString(land_sale_request.Response_date)
	land_sale_request.Submitted_date = formatDateString(land_sale_request.Submitted_date)
	land_sale_request.Withdrawn_date = formatDateString(land_sale_request.Withdrawn_date)
	land_sale_request.Row_effective_date = formatDateString(land_sale_request.Row_effective_date)
	land_sale_request.Row_expiry_date = formatDateString(land_sale_request.Row_expiry_date)






	rows, err := stmt.Exec(land_sale_request.Land_request_id, land_sale_request.Active_ind, land_sale_request.Advance_booking_ind, land_sale_request.Ami_ind, land_sale_request.Area_id, land_sale_request.Area_type, land_sale_request.Bid_ind, land_sale_request.Broker_ba_id, land_sale_request.Contact_ba_id, land_sale_request.Critical_date, land_sale_request.Drilling_priority_ind, land_sale_request.Effective_date, land_sale_request.Expiry_date, land_sale_request.Granted_right_type, land_sale_request.Gross_size, land_sale_request.Gross_size_ouom, land_sale_request.Investigator, land_sale_request.Jurisdiction, land_sale_request.Land_req_status, land_sale_request.Land_right_id, land_sale_request.Land_right_subtype, land_sale_request.Land_sale_number, land_sale_request.Land_sale_offering_id, land_sale_request.Lessor_refused_date, land_sale_request.Originator_ba_id, land_sale_request.Ppdm_guid, land_sale_request.Previous_request_id, land_sale_request.Purpose, land_sale_request.Reference_num, land_sale_request.Remark, land_sale_request.Requested_sale_date, land_sale_request.Requested_sale_name, land_sale_request.Requester, land_sale_request.Request_cancel_date, land_sale_request.Request_description, land_sale_request.Request_received_date, land_sale_request.Request_size, land_sale_request.Request_size_ouom, land_sale_request.Request_type, land_sale_request.Response_date, land_sale_request.Response_desc, land_sale_request.Source, land_sale_request.Submitted_date, land_sale_request.Term_duration, land_sale_request.Term_duration_ouom, land_sale_request.Uwi, land_sale_request.Well_license_id, land_sale_request.Well_source, land_sale_request.Withdrawn_date, land_sale_request.Row_changed_by, land_sale_request.Row_changed_date, land_sale_request.Row_created_by, land_sale_request.Row_created_date, land_sale_request.Row_effective_date, land_sale_request.Row_expiry_date, land_sale_request.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLandSaleRequest(c *fiber.Ctx) error {
	var land_sale_request dto.Land_sale_request

	setDefaults(&land_sale_request)

	if err := json.Unmarshal(c.Body(), &land_sale_request); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	land_sale_request.Land_request_id = id

    if land_sale_request.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update land_sale_request set active_ind = :1, advance_booking_ind = :2, ami_ind = :3, area_id = :4, area_type = :5, bid_ind = :6, broker_ba_id = :7, contact_ba_id = :8, critical_date = :9, drilling_priority_ind = :10, effective_date = :11, expiry_date = :12, granted_right_type = :13, gross_size = :14, gross_size_ouom = :15, investigator = :16, jurisdiction = :17, land_req_status = :18, land_right_id = :19, land_right_subtype = :20, land_sale_number = :21, land_sale_offering_id = :22, lessor_refused_date = :23, originator_ba_id = :24, ppdm_guid = :25, previous_request_id = :26, purpose = :27, reference_num = :28, remark = :29, requested_sale_date = :30, requested_sale_name = :31, requester = :32, request_cancel_date = :33, request_description = :34, request_received_date = :35, request_size = :36, request_size_ouom = :37, request_type = :38, response_date = :39, response_desc = :40, source = :41, submitted_date = :42, term_duration = :43, term_duration_ouom = :44, uwi = :45, well_license_id = :46, well_source = :47, withdrawn_date = :48, row_changed_by = :49, row_changed_date = :50, row_created_by = :51, row_effective_date = :52, row_expiry_date = :53, row_quality = :54 where land_request_id = :56")
	if err != nil {
		return err
	}

	land_sale_request.Row_changed_date = formatDate(land_sale_request.Row_changed_date)
	land_sale_request.Critical_date = formatDateString(land_sale_request.Critical_date)
	land_sale_request.Effective_date = formatDateString(land_sale_request.Effective_date)
	land_sale_request.Expiry_date = formatDateString(land_sale_request.Expiry_date)
	land_sale_request.Lessor_refused_date = formatDateString(land_sale_request.Lessor_refused_date)
	land_sale_request.Requested_sale_date = formatDateString(land_sale_request.Requested_sale_date)
	land_sale_request.Request_cancel_date = formatDateString(land_sale_request.Request_cancel_date)
	land_sale_request.Request_received_date = formatDateString(land_sale_request.Request_received_date)
	land_sale_request.Response_date = formatDateString(land_sale_request.Response_date)
	land_sale_request.Submitted_date = formatDateString(land_sale_request.Submitted_date)
	land_sale_request.Withdrawn_date = formatDateString(land_sale_request.Withdrawn_date)
	land_sale_request.Row_effective_date = formatDateString(land_sale_request.Row_effective_date)
	land_sale_request.Row_expiry_date = formatDateString(land_sale_request.Row_expiry_date)






	rows, err := stmt.Exec(land_sale_request.Active_ind, land_sale_request.Advance_booking_ind, land_sale_request.Ami_ind, land_sale_request.Area_id, land_sale_request.Area_type, land_sale_request.Bid_ind, land_sale_request.Broker_ba_id, land_sale_request.Contact_ba_id, land_sale_request.Critical_date, land_sale_request.Drilling_priority_ind, land_sale_request.Effective_date, land_sale_request.Expiry_date, land_sale_request.Granted_right_type, land_sale_request.Gross_size, land_sale_request.Gross_size_ouom, land_sale_request.Investigator, land_sale_request.Jurisdiction, land_sale_request.Land_req_status, land_sale_request.Land_right_id, land_sale_request.Land_right_subtype, land_sale_request.Land_sale_number, land_sale_request.Land_sale_offering_id, land_sale_request.Lessor_refused_date, land_sale_request.Originator_ba_id, land_sale_request.Ppdm_guid, land_sale_request.Previous_request_id, land_sale_request.Purpose, land_sale_request.Reference_num, land_sale_request.Remark, land_sale_request.Requested_sale_date, land_sale_request.Requested_sale_name, land_sale_request.Requester, land_sale_request.Request_cancel_date, land_sale_request.Request_description, land_sale_request.Request_received_date, land_sale_request.Request_size, land_sale_request.Request_size_ouom, land_sale_request.Request_type, land_sale_request.Response_date, land_sale_request.Response_desc, land_sale_request.Source, land_sale_request.Submitted_date, land_sale_request.Term_duration, land_sale_request.Term_duration_ouom, land_sale_request.Uwi, land_sale_request.Well_license_id, land_sale_request.Well_source, land_sale_request.Withdrawn_date, land_sale_request.Row_changed_by, land_sale_request.Row_changed_date, land_sale_request.Row_created_by, land_sale_request.Row_effective_date, land_sale_request.Row_expiry_date, land_sale_request.Row_quality, land_sale_request.Land_request_id)
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

func PatchLandSaleRequest(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update land_sale_request set "
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
		} else if key == "critical_date" || key == "effective_date" || key == "expiry_date" || key == "lessor_refused_date" || key == "requested_sale_date" || key == "request_cancel_date" || key == "request_received_date" || key == "response_date" || key == "submitted_date" || key == "withdrawn_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where land_request_id = :id"

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

func DeleteLandSaleRequest(c *fiber.Ctx) error {
	id := c.Params("id")
	var land_sale_request dto.Land_sale_request
	land_sale_request.Land_request_id = id

	stmt, err := db.Prepare("delete from land_sale_request where land_request_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(land_sale_request.Land_request_id)
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



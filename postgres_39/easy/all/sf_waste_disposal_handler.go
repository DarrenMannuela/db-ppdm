package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSfWasteDisposal(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sf_waste_disposal")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sf_waste_disposal

	for rows.Next() {
		var sf_waste_disposal dto.Sf_waste_disposal
		if err := rows.Scan(&sf_waste_disposal.Sf_subtype, &sf_waste_disposal.Waste_facility_id, &sf_waste_disposal.Waste_id, &sf_waste_disposal.Active_ind, &sf_waste_disposal.Adjust_reason, &sf_waste_disposal.Amount_adjustment, &sf_waste_disposal.Amount_adjustment_ouom, &sf_waste_disposal.Amount_adjustment_uom, &sf_waste_disposal.Amount_received, &sf_waste_disposal.Amount_received_ouom, &sf_waste_disposal.Amount_received_uom, &sf_waste_disposal.Amount_shipped, &sf_waste_disposal.Amount_shipped_ouom, &sf_waste_disposal.Amount_shipped_uom, &sf_waste_disposal.Contract_id, &sf_waste_disposal.Description, &sf_waste_disposal.Effective_date, &sf_waste_disposal.End_date, &sf_waste_disposal.Expiry_date, &sf_waste_disposal.Operator_ba_id, &sf_waste_disposal.Origin_facility_id, &sf_waste_disposal.Origin_facility_type, &sf_waste_disposal.Origin_hse_incident_id, &sf_waste_disposal.Origin_uwi, &sf_waste_disposal.Ppdm_guid, &sf_waste_disposal.Provision_id, &sf_waste_disposal.Rate_schedule_id, &sf_waste_disposal.Receiver_ba_id, &sf_waste_disposal.Remark, &sf_waste_disposal.Reporting_uom, &sf_waste_disposal.Service_quality, &sf_waste_disposal.Service_type, &sf_waste_disposal.Shipped_date, &sf_waste_disposal.Source, &sf_waste_disposal.Transport_ba_id, &sf_waste_disposal.Waste_handling_method, &sf_waste_disposal.Waste_hazard, &sf_waste_disposal.Waste_material, &sf_waste_disposal.Waste_origin, &sf_waste_disposal.Row_changed_by, &sf_waste_disposal.Row_changed_date, &sf_waste_disposal.Row_created_by, &sf_waste_disposal.Row_created_date, &sf_waste_disposal.Row_effective_date, &sf_waste_disposal.Row_expiry_date, &sf_waste_disposal.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sf_waste_disposal to result
		result = append(result, sf_waste_disposal)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSfWasteDisposal(c *fiber.Ctx) error {
	var sf_waste_disposal dto.Sf_waste_disposal

	setDefaults(&sf_waste_disposal)

	if err := json.Unmarshal(c.Body(), &sf_waste_disposal); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sf_waste_disposal values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46)")
	if err != nil {
		return err
	}
	sf_waste_disposal.Row_created_date = formatDate(sf_waste_disposal.Row_created_date)
	sf_waste_disposal.Row_changed_date = formatDate(sf_waste_disposal.Row_changed_date)
	sf_waste_disposal.Effective_date = formatDateString(sf_waste_disposal.Effective_date)
	sf_waste_disposal.End_date = formatDateString(sf_waste_disposal.End_date)
	sf_waste_disposal.Expiry_date = formatDateString(sf_waste_disposal.Expiry_date)
	sf_waste_disposal.Shipped_date = formatDateString(sf_waste_disposal.Shipped_date)
	sf_waste_disposal.Row_effective_date = formatDateString(sf_waste_disposal.Row_effective_date)
	sf_waste_disposal.Row_expiry_date = formatDateString(sf_waste_disposal.Row_expiry_date)






	rows, err := stmt.Exec(sf_waste_disposal.Sf_subtype, sf_waste_disposal.Waste_facility_id, sf_waste_disposal.Waste_id, sf_waste_disposal.Active_ind, sf_waste_disposal.Adjust_reason, sf_waste_disposal.Amount_adjustment, sf_waste_disposal.Amount_adjustment_ouom, sf_waste_disposal.Amount_adjustment_uom, sf_waste_disposal.Amount_received, sf_waste_disposal.Amount_received_ouom, sf_waste_disposal.Amount_received_uom, sf_waste_disposal.Amount_shipped, sf_waste_disposal.Amount_shipped_ouom, sf_waste_disposal.Amount_shipped_uom, sf_waste_disposal.Contract_id, sf_waste_disposal.Description, sf_waste_disposal.Effective_date, sf_waste_disposal.End_date, sf_waste_disposal.Expiry_date, sf_waste_disposal.Operator_ba_id, sf_waste_disposal.Origin_facility_id, sf_waste_disposal.Origin_facility_type, sf_waste_disposal.Origin_hse_incident_id, sf_waste_disposal.Origin_uwi, sf_waste_disposal.Ppdm_guid, sf_waste_disposal.Provision_id, sf_waste_disposal.Rate_schedule_id, sf_waste_disposal.Receiver_ba_id, sf_waste_disposal.Remark, sf_waste_disposal.Reporting_uom, sf_waste_disposal.Service_quality, sf_waste_disposal.Service_type, sf_waste_disposal.Shipped_date, sf_waste_disposal.Source, sf_waste_disposal.Transport_ba_id, sf_waste_disposal.Waste_handling_method, sf_waste_disposal.Waste_hazard, sf_waste_disposal.Waste_material, sf_waste_disposal.Waste_origin, sf_waste_disposal.Row_changed_by, sf_waste_disposal.Row_changed_date, sf_waste_disposal.Row_created_by, sf_waste_disposal.Row_created_date, sf_waste_disposal.Row_effective_date, sf_waste_disposal.Row_expiry_date, sf_waste_disposal.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSfWasteDisposal(c *fiber.Ctx) error {
	var sf_waste_disposal dto.Sf_waste_disposal

	setDefaults(&sf_waste_disposal)

	if err := json.Unmarshal(c.Body(), &sf_waste_disposal); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sf_waste_disposal.Sf_subtype = id

    if sf_waste_disposal.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sf_waste_disposal set waste_facility_id = :1, waste_id = :2, active_ind = :3, adjust_reason = :4, amount_adjustment = :5, amount_adjustment_ouom = :6, amount_adjustment_uom = :7, amount_received = :8, amount_received_ouom = :9, amount_received_uom = :10, amount_shipped = :11, amount_shipped_ouom = :12, amount_shipped_uom = :13, contract_id = :14, description = :15, effective_date = :16, end_date = :17, expiry_date = :18, operator_ba_id = :19, origin_facility_id = :20, origin_facility_type = :21, origin_hse_incident_id = :22, origin_uwi = :23, ppdm_guid = :24, provision_id = :25, rate_schedule_id = :26, receiver_ba_id = :27, remark = :28, reporting_uom = :29, service_quality = :30, service_type = :31, shipped_date = :32, source = :33, transport_ba_id = :34, waste_handling_method = :35, waste_hazard = :36, waste_material = :37, waste_origin = :38, row_changed_by = :39, row_changed_date = :40, row_created_by = :41, row_effective_date = :42, row_expiry_date = :43, row_quality = :44 where sf_subtype = :46")
	if err != nil {
		return err
	}

	sf_waste_disposal.Row_changed_date = formatDate(sf_waste_disposal.Row_changed_date)
	sf_waste_disposal.Effective_date = formatDateString(sf_waste_disposal.Effective_date)
	sf_waste_disposal.End_date = formatDateString(sf_waste_disposal.End_date)
	sf_waste_disposal.Expiry_date = formatDateString(sf_waste_disposal.Expiry_date)
	sf_waste_disposal.Shipped_date = formatDateString(sf_waste_disposal.Shipped_date)
	sf_waste_disposal.Row_effective_date = formatDateString(sf_waste_disposal.Row_effective_date)
	sf_waste_disposal.Row_expiry_date = formatDateString(sf_waste_disposal.Row_expiry_date)






	rows, err := stmt.Exec(sf_waste_disposal.Waste_facility_id, sf_waste_disposal.Waste_id, sf_waste_disposal.Active_ind, sf_waste_disposal.Adjust_reason, sf_waste_disposal.Amount_adjustment, sf_waste_disposal.Amount_adjustment_ouom, sf_waste_disposal.Amount_adjustment_uom, sf_waste_disposal.Amount_received, sf_waste_disposal.Amount_received_ouom, sf_waste_disposal.Amount_received_uom, sf_waste_disposal.Amount_shipped, sf_waste_disposal.Amount_shipped_ouom, sf_waste_disposal.Amount_shipped_uom, sf_waste_disposal.Contract_id, sf_waste_disposal.Description, sf_waste_disposal.Effective_date, sf_waste_disposal.End_date, sf_waste_disposal.Expiry_date, sf_waste_disposal.Operator_ba_id, sf_waste_disposal.Origin_facility_id, sf_waste_disposal.Origin_facility_type, sf_waste_disposal.Origin_hse_incident_id, sf_waste_disposal.Origin_uwi, sf_waste_disposal.Ppdm_guid, sf_waste_disposal.Provision_id, sf_waste_disposal.Rate_schedule_id, sf_waste_disposal.Receiver_ba_id, sf_waste_disposal.Remark, sf_waste_disposal.Reporting_uom, sf_waste_disposal.Service_quality, sf_waste_disposal.Service_type, sf_waste_disposal.Shipped_date, sf_waste_disposal.Source, sf_waste_disposal.Transport_ba_id, sf_waste_disposal.Waste_handling_method, sf_waste_disposal.Waste_hazard, sf_waste_disposal.Waste_material, sf_waste_disposal.Waste_origin, sf_waste_disposal.Row_changed_by, sf_waste_disposal.Row_changed_date, sf_waste_disposal.Row_created_by, sf_waste_disposal.Row_effective_date, sf_waste_disposal.Row_expiry_date, sf_waste_disposal.Row_quality, sf_waste_disposal.Sf_subtype)
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

func PatchSfWasteDisposal(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sf_waste_disposal set "
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
		} else if key == "effective_date" || key == "end_date" || key == "expiry_date" || key == "shipped_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where sf_subtype = :id"

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

func DeleteSfWasteDisposal(c *fiber.Ctx) error {
	id := c.Params("id")
	var sf_waste_disposal dto.Sf_waste_disposal
	sf_waste_disposal.Sf_subtype = id

	stmt, err := db.Prepare("delete from sf_waste_disposal where sf_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sf_waste_disposal.Sf_subtype)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFacilityBaService(c *fiber.Ctx) error {
	rows, err := db.Query("select * from facility_ba_service")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Facility_ba_service

	for rows.Next() {
		var facility_ba_service dto.Facility_ba_service
		if err := rows.Scan(&facility_ba_service.Facility_id, &facility_ba_service.Facility_type, &facility_ba_service.Provided_by, &facility_ba_service.Service_seq_no, &facility_ba_service.Active_ind, &facility_ba_service.Ba_service_type, &facility_ba_service.Contact_ba_id, &facility_ba_service.Contract_id, &facility_ba_service.Description, &facility_ba_service.Effective_date, &facility_ba_service.End_date, &facility_ba_service.Expiry_date, &facility_ba_service.Ppdm_guid, &facility_ba_service.Provision_id, &facility_ba_service.Rate_schedule_id, &facility_ba_service.Remark, &facility_ba_service.Represented_ba_id, &facility_ba_service.Service_quality, &facility_ba_service.Source, &facility_ba_service.Start_date, &facility_ba_service.Row_changed_by, &facility_ba_service.Row_changed_date, &facility_ba_service.Row_created_by, &facility_ba_service.Row_created_date, &facility_ba_service.Row_effective_date, &facility_ba_service.Row_expiry_date, &facility_ba_service.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append facility_ba_service to result
		result = append(result, facility_ba_service)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFacilityBaService(c *fiber.Ctx) error {
	var facility_ba_service dto.Facility_ba_service

	setDefaults(&facility_ba_service)

	if err := json.Unmarshal(c.Body(), &facility_ba_service); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into facility_ba_service values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27)")
	if err != nil {
		return err
	}
	facility_ba_service.Row_created_date = formatDate(facility_ba_service.Row_created_date)
	facility_ba_service.Row_changed_date = formatDate(facility_ba_service.Row_changed_date)
	facility_ba_service.Effective_date = formatDateString(facility_ba_service.Effective_date)
	facility_ba_service.End_date = formatDateString(facility_ba_service.End_date)
	facility_ba_service.Expiry_date = formatDateString(facility_ba_service.Expiry_date)
	facility_ba_service.Start_date = formatDateString(facility_ba_service.Start_date)
	facility_ba_service.Row_effective_date = formatDateString(facility_ba_service.Row_effective_date)
	facility_ba_service.Row_expiry_date = formatDateString(facility_ba_service.Row_expiry_date)






	rows, err := stmt.Exec(facility_ba_service.Facility_id, facility_ba_service.Facility_type, facility_ba_service.Provided_by, facility_ba_service.Service_seq_no, facility_ba_service.Active_ind, facility_ba_service.Ba_service_type, facility_ba_service.Contact_ba_id, facility_ba_service.Contract_id, facility_ba_service.Description, facility_ba_service.Effective_date, facility_ba_service.End_date, facility_ba_service.Expiry_date, facility_ba_service.Ppdm_guid, facility_ba_service.Provision_id, facility_ba_service.Rate_schedule_id, facility_ba_service.Remark, facility_ba_service.Represented_ba_id, facility_ba_service.Service_quality, facility_ba_service.Source, facility_ba_service.Start_date, facility_ba_service.Row_changed_by, facility_ba_service.Row_changed_date, facility_ba_service.Row_created_by, facility_ba_service.Row_created_date, facility_ba_service.Row_effective_date, facility_ba_service.Row_expiry_date, facility_ba_service.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFacilityBaService(c *fiber.Ctx) error {
	var facility_ba_service dto.Facility_ba_service

	setDefaults(&facility_ba_service)

	if err := json.Unmarshal(c.Body(), &facility_ba_service); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	facility_ba_service.Facility_id = id

    if facility_ba_service.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update facility_ba_service set facility_type = :1, provided_by = :2, service_seq_no = :3, active_ind = :4, ba_service_type = :5, contact_ba_id = :6, contract_id = :7, description = :8, effective_date = :9, end_date = :10, expiry_date = :11, ppdm_guid = :12, provision_id = :13, rate_schedule_id = :14, remark = :15, represented_ba_id = :16, service_quality = :17, source = :18, start_date = :19, row_changed_by = :20, row_changed_date = :21, row_created_by = :22, row_effective_date = :23, row_expiry_date = :24, row_quality = :25 where facility_id = :27")
	if err != nil {
		return err
	}

	facility_ba_service.Row_changed_date = formatDate(facility_ba_service.Row_changed_date)
	facility_ba_service.Effective_date = formatDateString(facility_ba_service.Effective_date)
	facility_ba_service.End_date = formatDateString(facility_ba_service.End_date)
	facility_ba_service.Expiry_date = formatDateString(facility_ba_service.Expiry_date)
	facility_ba_service.Start_date = formatDateString(facility_ba_service.Start_date)
	facility_ba_service.Row_effective_date = formatDateString(facility_ba_service.Row_effective_date)
	facility_ba_service.Row_expiry_date = formatDateString(facility_ba_service.Row_expiry_date)






	rows, err := stmt.Exec(facility_ba_service.Facility_type, facility_ba_service.Provided_by, facility_ba_service.Service_seq_no, facility_ba_service.Active_ind, facility_ba_service.Ba_service_type, facility_ba_service.Contact_ba_id, facility_ba_service.Contract_id, facility_ba_service.Description, facility_ba_service.Effective_date, facility_ba_service.End_date, facility_ba_service.Expiry_date, facility_ba_service.Ppdm_guid, facility_ba_service.Provision_id, facility_ba_service.Rate_schedule_id, facility_ba_service.Remark, facility_ba_service.Represented_ba_id, facility_ba_service.Service_quality, facility_ba_service.Source, facility_ba_service.Start_date, facility_ba_service.Row_changed_by, facility_ba_service.Row_changed_date, facility_ba_service.Row_created_by, facility_ba_service.Row_effective_date, facility_ba_service.Row_expiry_date, facility_ba_service.Row_quality, facility_ba_service.Facility_id)
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

func PatchFacilityBaService(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update facility_ba_service set "
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
		} else if key == "effective_date" || key == "end_date" || key == "expiry_date" || key == "start_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where facility_id = :id"

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

func DeleteFacilityBaService(c *fiber.Ctx) error {
	id := c.Params("id")
	var facility_ba_service dto.Facility_ba_service
	facility_ba_service.Facility_id = id

	stmt, err := db.Prepare("delete from facility_ba_service where facility_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(facility_ba_service.Facility_id)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetBaConsortiumService(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ba_consortium_service")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ba_consortium_service

	for rows.Next() {
		var ba_consortium_service dto.Ba_consortium_service
		if err := rows.Scan(&ba_consortium_service.Consortium_ba_id, &ba_consortium_service.Provided_by, &ba_consortium_service.Service_seq_no, &ba_consortium_service.Active_ind, &ba_consortium_service.Contact_ba_id, &ba_consortium_service.Contract_id, &ba_consortium_service.Description, &ba_consortium_service.Effective_date, &ba_consortium_service.End_date, &ba_consortium_service.Expiry_date, &ba_consortium_service.Ppdm_guid, &ba_consortium_service.Provision_id, &ba_consortium_service.Rate_schedule_id, &ba_consortium_service.Remark, &ba_consortium_service.Represented_ba_id, &ba_consortium_service.Service_quality, &ba_consortium_service.Service_type, &ba_consortium_service.Source, &ba_consortium_service.Start_date, &ba_consortium_service.Row_changed_by, &ba_consortium_service.Row_changed_date, &ba_consortium_service.Row_created_by, &ba_consortium_service.Row_created_date, &ba_consortium_service.Row_effective_date, &ba_consortium_service.Row_expiry_date, &ba_consortium_service.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ba_consortium_service to result
		result = append(result, ba_consortium_service)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetBaConsortiumService(c *fiber.Ctx) error {
	var ba_consortium_service dto.Ba_consortium_service

	setDefaults(&ba_consortium_service)

	if err := json.Unmarshal(c.Body(), &ba_consortium_service); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ba_consortium_service values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26)")
	if err != nil {
		return err
	}
	ba_consortium_service.Row_created_date = formatDate(ba_consortium_service.Row_created_date)
	ba_consortium_service.Row_changed_date = formatDate(ba_consortium_service.Row_changed_date)
	ba_consortium_service.Effective_date = formatDateString(ba_consortium_service.Effective_date)
	ba_consortium_service.End_date = formatDateString(ba_consortium_service.End_date)
	ba_consortium_service.Expiry_date = formatDateString(ba_consortium_service.Expiry_date)
	ba_consortium_service.Start_date = formatDateString(ba_consortium_service.Start_date)
	ba_consortium_service.Row_effective_date = formatDateString(ba_consortium_service.Row_effective_date)
	ba_consortium_service.Row_expiry_date = formatDateString(ba_consortium_service.Row_expiry_date)






	rows, err := stmt.Exec(ba_consortium_service.Consortium_ba_id, ba_consortium_service.Provided_by, ba_consortium_service.Service_seq_no, ba_consortium_service.Active_ind, ba_consortium_service.Contact_ba_id, ba_consortium_service.Contract_id, ba_consortium_service.Description, ba_consortium_service.Effective_date, ba_consortium_service.End_date, ba_consortium_service.Expiry_date, ba_consortium_service.Ppdm_guid, ba_consortium_service.Provision_id, ba_consortium_service.Rate_schedule_id, ba_consortium_service.Remark, ba_consortium_service.Represented_ba_id, ba_consortium_service.Service_quality, ba_consortium_service.Service_type, ba_consortium_service.Source, ba_consortium_service.Start_date, ba_consortium_service.Row_changed_by, ba_consortium_service.Row_changed_date, ba_consortium_service.Row_created_by, ba_consortium_service.Row_created_date, ba_consortium_service.Row_effective_date, ba_consortium_service.Row_expiry_date, ba_consortium_service.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateBaConsortiumService(c *fiber.Ctx) error {
	var ba_consortium_service dto.Ba_consortium_service

	setDefaults(&ba_consortium_service)

	if err := json.Unmarshal(c.Body(), &ba_consortium_service); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ba_consortium_service.Consortium_ba_id = id

    if ba_consortium_service.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ba_consortium_service set provided_by = :1, service_seq_no = :2, active_ind = :3, contact_ba_id = :4, contract_id = :5, description = :6, effective_date = :7, end_date = :8, expiry_date = :9, ppdm_guid = :10, provision_id = :11, rate_schedule_id = :12, remark = :13, represented_ba_id = :14, service_quality = :15, service_type = :16, source = :17, start_date = :18, row_changed_by = :19, row_changed_date = :20, row_created_by = :21, row_effective_date = :22, row_expiry_date = :23, row_quality = :24 where consortium_ba_id = :26")
	if err != nil {
		return err
	}

	ba_consortium_service.Row_changed_date = formatDate(ba_consortium_service.Row_changed_date)
	ba_consortium_service.Effective_date = formatDateString(ba_consortium_service.Effective_date)
	ba_consortium_service.End_date = formatDateString(ba_consortium_service.End_date)
	ba_consortium_service.Expiry_date = formatDateString(ba_consortium_service.Expiry_date)
	ba_consortium_service.Start_date = formatDateString(ba_consortium_service.Start_date)
	ba_consortium_service.Row_effective_date = formatDateString(ba_consortium_service.Row_effective_date)
	ba_consortium_service.Row_expiry_date = formatDateString(ba_consortium_service.Row_expiry_date)






	rows, err := stmt.Exec(ba_consortium_service.Provided_by, ba_consortium_service.Service_seq_no, ba_consortium_service.Active_ind, ba_consortium_service.Contact_ba_id, ba_consortium_service.Contract_id, ba_consortium_service.Description, ba_consortium_service.Effective_date, ba_consortium_service.End_date, ba_consortium_service.Expiry_date, ba_consortium_service.Ppdm_guid, ba_consortium_service.Provision_id, ba_consortium_service.Rate_schedule_id, ba_consortium_service.Remark, ba_consortium_service.Represented_ba_id, ba_consortium_service.Service_quality, ba_consortium_service.Service_type, ba_consortium_service.Source, ba_consortium_service.Start_date, ba_consortium_service.Row_changed_by, ba_consortium_service.Row_changed_date, ba_consortium_service.Row_created_by, ba_consortium_service.Row_effective_date, ba_consortium_service.Row_expiry_date, ba_consortium_service.Row_quality, ba_consortium_service.Consortium_ba_id)
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

func PatchBaConsortiumService(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ba_consortium_service set "
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
	query += " where consortium_ba_id = :id"

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

func DeleteBaConsortiumService(c *fiber.Ctx) error {
	id := c.Params("id")
	var ba_consortium_service dto.Ba_consortium_service
	ba_consortium_service.Consortium_ba_id = id

	stmt, err := db.Prepare("delete from ba_consortium_service where consortium_ba_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ba_consortium_service.Consortium_ba_id)
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



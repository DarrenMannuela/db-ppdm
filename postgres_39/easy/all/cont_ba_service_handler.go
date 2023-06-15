package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetContBaService(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cont_ba_service")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cont_ba_service

	for rows.Next() {
		var cont_ba_service dto.Cont_ba_service
		if err := rows.Scan(&cont_ba_service.Contract_id, &cont_ba_service.Provided_by, &cont_ba_service.Service_seq_no, &cont_ba_service.Active_ind, &cont_ba_service.Ba_service_type, &cont_ba_service.Contact_ba_id, &cont_ba_service.Description, &cont_ba_service.Effective_date, &cont_ba_service.End_date, &cont_ba_service.Expiry_date, &cont_ba_service.Original_party_name, &cont_ba_service.Ppdm_guid, &cont_ba_service.Provision_id, &cont_ba_service.Rate_schedule_id, &cont_ba_service.Remark, &cont_ba_service.Represented_ba_id, &cont_ba_service.Service_quality, &cont_ba_service.Source, &cont_ba_service.Start_date, &cont_ba_service.Row_changed_by, &cont_ba_service.Row_changed_date, &cont_ba_service.Row_created_by, &cont_ba_service.Row_created_date, &cont_ba_service.Row_effective_date, &cont_ba_service.Row_expiry_date, &cont_ba_service.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cont_ba_service to result
		result = append(result, cont_ba_service)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetContBaService(c *fiber.Ctx) error {
	var cont_ba_service dto.Cont_ba_service

	setDefaults(&cont_ba_service)

	if err := json.Unmarshal(c.Body(), &cont_ba_service); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cont_ba_service values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26)")
	if err != nil {
		return err
	}
	cont_ba_service.Row_created_date = formatDate(cont_ba_service.Row_created_date)
	cont_ba_service.Row_changed_date = formatDate(cont_ba_service.Row_changed_date)
	cont_ba_service.Effective_date = formatDateString(cont_ba_service.Effective_date)
	cont_ba_service.End_date = formatDateString(cont_ba_service.End_date)
	cont_ba_service.Expiry_date = formatDateString(cont_ba_service.Expiry_date)
	cont_ba_service.Start_date = formatDateString(cont_ba_service.Start_date)
	cont_ba_service.Row_effective_date = formatDateString(cont_ba_service.Row_effective_date)
	cont_ba_service.Row_expiry_date = formatDateString(cont_ba_service.Row_expiry_date)






	rows, err := stmt.Exec(cont_ba_service.Contract_id, cont_ba_service.Provided_by, cont_ba_service.Service_seq_no, cont_ba_service.Active_ind, cont_ba_service.Ba_service_type, cont_ba_service.Contact_ba_id, cont_ba_service.Description, cont_ba_service.Effective_date, cont_ba_service.End_date, cont_ba_service.Expiry_date, cont_ba_service.Original_party_name, cont_ba_service.Ppdm_guid, cont_ba_service.Provision_id, cont_ba_service.Rate_schedule_id, cont_ba_service.Remark, cont_ba_service.Represented_ba_id, cont_ba_service.Service_quality, cont_ba_service.Source, cont_ba_service.Start_date, cont_ba_service.Row_changed_by, cont_ba_service.Row_changed_date, cont_ba_service.Row_created_by, cont_ba_service.Row_created_date, cont_ba_service.Row_effective_date, cont_ba_service.Row_expiry_date, cont_ba_service.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateContBaService(c *fiber.Ctx) error {
	var cont_ba_service dto.Cont_ba_service

	setDefaults(&cont_ba_service)

	if err := json.Unmarshal(c.Body(), &cont_ba_service); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cont_ba_service.Contract_id = id

    if cont_ba_service.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cont_ba_service set provided_by = :1, service_seq_no = :2, active_ind = :3, ba_service_type = :4, contact_ba_id = :5, description = :6, effective_date = :7, end_date = :8, expiry_date = :9, original_party_name = :10, ppdm_guid = :11, provision_id = :12, rate_schedule_id = :13, remark = :14, represented_ba_id = :15, service_quality = :16, source = :17, start_date = :18, row_changed_by = :19, row_changed_date = :20, row_created_by = :21, row_effective_date = :22, row_expiry_date = :23, row_quality = :24 where contract_id = :26")
	if err != nil {
		return err
	}

	cont_ba_service.Row_changed_date = formatDate(cont_ba_service.Row_changed_date)
	cont_ba_service.Effective_date = formatDateString(cont_ba_service.Effective_date)
	cont_ba_service.End_date = formatDateString(cont_ba_service.End_date)
	cont_ba_service.Expiry_date = formatDateString(cont_ba_service.Expiry_date)
	cont_ba_service.Start_date = formatDateString(cont_ba_service.Start_date)
	cont_ba_service.Row_effective_date = formatDateString(cont_ba_service.Row_effective_date)
	cont_ba_service.Row_expiry_date = formatDateString(cont_ba_service.Row_expiry_date)






	rows, err := stmt.Exec(cont_ba_service.Provided_by, cont_ba_service.Service_seq_no, cont_ba_service.Active_ind, cont_ba_service.Ba_service_type, cont_ba_service.Contact_ba_id, cont_ba_service.Description, cont_ba_service.Effective_date, cont_ba_service.End_date, cont_ba_service.Expiry_date, cont_ba_service.Original_party_name, cont_ba_service.Ppdm_guid, cont_ba_service.Provision_id, cont_ba_service.Rate_schedule_id, cont_ba_service.Remark, cont_ba_service.Represented_ba_id, cont_ba_service.Service_quality, cont_ba_service.Source, cont_ba_service.Start_date, cont_ba_service.Row_changed_by, cont_ba_service.Row_changed_date, cont_ba_service.Row_created_by, cont_ba_service.Row_effective_date, cont_ba_service.Row_expiry_date, cont_ba_service.Row_quality, cont_ba_service.Contract_id)
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

func PatchContBaService(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cont_ba_service set "
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
	query += " where contract_id = :id"

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

func DeleteContBaService(c *fiber.Ctx) error {
	id := c.Params("id")
	var cont_ba_service dto.Cont_ba_service
	cont_ba_service.Contract_id = id

	stmt, err := db.Prepare("delete from cont_ba_service where contract_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cont_ba_service.Contract_id)
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



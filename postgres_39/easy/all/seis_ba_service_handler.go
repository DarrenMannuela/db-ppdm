package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisBaService(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_ba_service")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_ba_service

	for rows.Next() {
		var seis_ba_service dto.Seis_ba_service
		if err := rows.Scan(&seis_ba_service.Seis_set_subtype, &seis_ba_service.Seis_set_id, &seis_ba_service.Provided_by, &seis_ba_service.Service_seq_no, &seis_ba_service.Active_ind, &seis_ba_service.Ba_service_type, &seis_ba_service.Contact_ba_id, &seis_ba_service.Contract_id, &seis_ba_service.Description, &seis_ba_service.Effective_date, &seis_ba_service.End_date, &seis_ba_service.Expiry_date, &seis_ba_service.Ppdm_guid, &seis_ba_service.Provision_id, &seis_ba_service.Rate_schedule_id, &seis_ba_service.Remark, &seis_ba_service.Represented_ba_id, &seis_ba_service.Service_date, &seis_ba_service.Service_quality, &seis_ba_service.Service_remark, &seis_ba_service.Source, &seis_ba_service.Start_date, &seis_ba_service.Row_changed_by, &seis_ba_service.Row_changed_date, &seis_ba_service.Row_created_by, &seis_ba_service.Row_created_date, &seis_ba_service.Row_effective_date, &seis_ba_service.Row_expiry_date, &seis_ba_service.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_ba_service to result
		result = append(result, seis_ba_service)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisBaService(c *fiber.Ctx) error {
	var seis_ba_service dto.Seis_ba_service

	setDefaults(&seis_ba_service)

	if err := json.Unmarshal(c.Body(), &seis_ba_service); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_ba_service values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29)")
	if err != nil {
		return err
	}
	seis_ba_service.Row_created_date = formatDate(seis_ba_service.Row_created_date)
	seis_ba_service.Row_changed_date = formatDate(seis_ba_service.Row_changed_date)
	seis_ba_service.Effective_date = formatDateString(seis_ba_service.Effective_date)
	seis_ba_service.End_date = formatDateString(seis_ba_service.End_date)
	seis_ba_service.Expiry_date = formatDateString(seis_ba_service.Expiry_date)
	seis_ba_service.Service_date = formatDateString(seis_ba_service.Service_date)
	seis_ba_service.Start_date = formatDateString(seis_ba_service.Start_date)
	seis_ba_service.Row_effective_date = formatDateString(seis_ba_service.Row_effective_date)
	seis_ba_service.Row_expiry_date = formatDateString(seis_ba_service.Row_expiry_date)






	rows, err := stmt.Exec(seis_ba_service.Seis_set_subtype, seis_ba_service.Seis_set_id, seis_ba_service.Provided_by, seis_ba_service.Service_seq_no, seis_ba_service.Active_ind, seis_ba_service.Ba_service_type, seis_ba_service.Contact_ba_id, seis_ba_service.Contract_id, seis_ba_service.Description, seis_ba_service.Effective_date, seis_ba_service.End_date, seis_ba_service.Expiry_date, seis_ba_service.Ppdm_guid, seis_ba_service.Provision_id, seis_ba_service.Rate_schedule_id, seis_ba_service.Remark, seis_ba_service.Represented_ba_id, seis_ba_service.Service_date, seis_ba_service.Service_quality, seis_ba_service.Service_remark, seis_ba_service.Source, seis_ba_service.Start_date, seis_ba_service.Row_changed_by, seis_ba_service.Row_changed_date, seis_ba_service.Row_created_by, seis_ba_service.Row_created_date, seis_ba_service.Row_effective_date, seis_ba_service.Row_expiry_date, seis_ba_service.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisBaService(c *fiber.Ctx) error {
	var seis_ba_service dto.Seis_ba_service

	setDefaults(&seis_ba_service)

	if err := json.Unmarshal(c.Body(), &seis_ba_service); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_ba_service.Seis_set_subtype = id

    if seis_ba_service.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_ba_service set seis_set_id = :1, provided_by = :2, service_seq_no = :3, active_ind = :4, ba_service_type = :5, contact_ba_id = :6, contract_id = :7, description = :8, effective_date = :9, end_date = :10, expiry_date = :11, ppdm_guid = :12, provision_id = :13, rate_schedule_id = :14, remark = :15, represented_ba_id = :16, service_date = :17, service_quality = :18, service_remark = :19, source = :20, start_date = :21, row_changed_by = :22, row_changed_date = :23, row_created_by = :24, row_effective_date = :25, row_expiry_date = :26, row_quality = :27 where seis_set_subtype = :29")
	if err != nil {
		return err
	}

	seis_ba_service.Row_changed_date = formatDate(seis_ba_service.Row_changed_date)
	seis_ba_service.Effective_date = formatDateString(seis_ba_service.Effective_date)
	seis_ba_service.End_date = formatDateString(seis_ba_service.End_date)
	seis_ba_service.Expiry_date = formatDateString(seis_ba_service.Expiry_date)
	seis_ba_service.Service_date = formatDateString(seis_ba_service.Service_date)
	seis_ba_service.Start_date = formatDateString(seis_ba_service.Start_date)
	seis_ba_service.Row_effective_date = formatDateString(seis_ba_service.Row_effective_date)
	seis_ba_service.Row_expiry_date = formatDateString(seis_ba_service.Row_expiry_date)






	rows, err := stmt.Exec(seis_ba_service.Seis_set_id, seis_ba_service.Provided_by, seis_ba_service.Service_seq_no, seis_ba_service.Active_ind, seis_ba_service.Ba_service_type, seis_ba_service.Contact_ba_id, seis_ba_service.Contract_id, seis_ba_service.Description, seis_ba_service.Effective_date, seis_ba_service.End_date, seis_ba_service.Expiry_date, seis_ba_service.Ppdm_guid, seis_ba_service.Provision_id, seis_ba_service.Rate_schedule_id, seis_ba_service.Remark, seis_ba_service.Represented_ba_id, seis_ba_service.Service_date, seis_ba_service.Service_quality, seis_ba_service.Service_remark, seis_ba_service.Source, seis_ba_service.Start_date, seis_ba_service.Row_changed_by, seis_ba_service.Row_changed_date, seis_ba_service.Row_created_by, seis_ba_service.Row_effective_date, seis_ba_service.Row_expiry_date, seis_ba_service.Row_quality, seis_ba_service.Seis_set_subtype)
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

func PatchSeisBaService(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_ba_service set "
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
		} else if key == "effective_date" || key == "end_date" || key == "expiry_date" || key == "service_date" || key == "start_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where seis_set_subtype = :id"

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

func DeleteSeisBaService(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_ba_service dto.Seis_ba_service
	seis_ba_service.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_ba_service where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_ba_service.Seis_set_subtype)
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



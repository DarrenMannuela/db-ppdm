package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetObligBaService(c *fiber.Ctx) error {
	rows, err := db.Query("select * from oblig_ba_service")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Oblig_ba_service

	for rows.Next() {
		var oblig_ba_service dto.Oblig_ba_service
		if err := rows.Scan(&oblig_ba_service.Obligation_id, &oblig_ba_service.Obligation_seq_no, &oblig_ba_service.Provided_by, &oblig_ba_service.Oblig_service_seq_no, &oblig_ba_service.Active_ind, &oblig_ba_service.Contact_ba_id, &oblig_ba_service.Contract_id, &oblig_ba_service.Description, &oblig_ba_service.Effective_date, &oblig_ba_service.End_date, &oblig_ba_service.Expiry_date, &oblig_ba_service.Ppdm_guid, &oblig_ba_service.Provided_for, &oblig_ba_service.Provision_id, &oblig_ba_service.Rate_schedule_id, &oblig_ba_service.Remark, &oblig_ba_service.Service_quality, &oblig_ba_service.Service_type, &oblig_ba_service.Source, &oblig_ba_service.Start_date, &oblig_ba_service.Row_changed_by, &oblig_ba_service.Row_changed_date, &oblig_ba_service.Row_created_by, &oblig_ba_service.Row_created_date, &oblig_ba_service.Row_effective_date, &oblig_ba_service.Row_expiry_date, &oblig_ba_service.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append oblig_ba_service to result
		result = append(result, oblig_ba_service)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetObligBaService(c *fiber.Ctx) error {
	var oblig_ba_service dto.Oblig_ba_service

	setDefaults(&oblig_ba_service)

	if err := json.Unmarshal(c.Body(), &oblig_ba_service); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into oblig_ba_service values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27)")
	if err != nil {
		return err
	}
	oblig_ba_service.Row_created_date = formatDate(oblig_ba_service.Row_created_date)
	oblig_ba_service.Row_changed_date = formatDate(oblig_ba_service.Row_changed_date)
	oblig_ba_service.Effective_date = formatDateString(oblig_ba_service.Effective_date)
	oblig_ba_service.End_date = formatDateString(oblig_ba_service.End_date)
	oblig_ba_service.Expiry_date = formatDateString(oblig_ba_service.Expiry_date)
	oblig_ba_service.Start_date = formatDateString(oblig_ba_service.Start_date)
	oblig_ba_service.Row_effective_date = formatDateString(oblig_ba_service.Row_effective_date)
	oblig_ba_service.Row_expiry_date = formatDateString(oblig_ba_service.Row_expiry_date)






	rows, err := stmt.Exec(oblig_ba_service.Obligation_id, oblig_ba_service.Obligation_seq_no, oblig_ba_service.Provided_by, oblig_ba_service.Oblig_service_seq_no, oblig_ba_service.Active_ind, oblig_ba_service.Contact_ba_id, oblig_ba_service.Contract_id, oblig_ba_service.Description, oblig_ba_service.Effective_date, oblig_ba_service.End_date, oblig_ba_service.Expiry_date, oblig_ba_service.Ppdm_guid, oblig_ba_service.Provided_for, oblig_ba_service.Provision_id, oblig_ba_service.Rate_schedule_id, oblig_ba_service.Remark, oblig_ba_service.Service_quality, oblig_ba_service.Service_type, oblig_ba_service.Source, oblig_ba_service.Start_date, oblig_ba_service.Row_changed_by, oblig_ba_service.Row_changed_date, oblig_ba_service.Row_created_by, oblig_ba_service.Row_created_date, oblig_ba_service.Row_effective_date, oblig_ba_service.Row_expiry_date, oblig_ba_service.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateObligBaService(c *fiber.Ctx) error {
	var oblig_ba_service dto.Oblig_ba_service

	setDefaults(&oblig_ba_service)

	if err := json.Unmarshal(c.Body(), &oblig_ba_service); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	oblig_ba_service.Obligation_id = id

    if oblig_ba_service.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update oblig_ba_service set obligation_seq_no = :1, provided_by = :2, oblig_service_seq_no = :3, active_ind = :4, contact_ba_id = :5, contract_id = :6, description = :7, effective_date = :8, end_date = :9, expiry_date = :10, ppdm_guid = :11, provided_for = :12, provision_id = :13, rate_schedule_id = :14, remark = :15, service_quality = :16, service_type = :17, source = :18, start_date = :19, row_changed_by = :20, row_changed_date = :21, row_created_by = :22, row_effective_date = :23, row_expiry_date = :24, row_quality = :25 where obligation_id = :27")
	if err != nil {
		return err
	}

	oblig_ba_service.Row_changed_date = formatDate(oblig_ba_service.Row_changed_date)
	oblig_ba_service.Effective_date = formatDateString(oblig_ba_service.Effective_date)
	oblig_ba_service.End_date = formatDateString(oblig_ba_service.End_date)
	oblig_ba_service.Expiry_date = formatDateString(oblig_ba_service.Expiry_date)
	oblig_ba_service.Start_date = formatDateString(oblig_ba_service.Start_date)
	oblig_ba_service.Row_effective_date = formatDateString(oblig_ba_service.Row_effective_date)
	oblig_ba_service.Row_expiry_date = formatDateString(oblig_ba_service.Row_expiry_date)






	rows, err := stmt.Exec(oblig_ba_service.Obligation_seq_no, oblig_ba_service.Provided_by, oblig_ba_service.Oblig_service_seq_no, oblig_ba_service.Active_ind, oblig_ba_service.Contact_ba_id, oblig_ba_service.Contract_id, oblig_ba_service.Description, oblig_ba_service.Effective_date, oblig_ba_service.End_date, oblig_ba_service.Expiry_date, oblig_ba_service.Ppdm_guid, oblig_ba_service.Provided_for, oblig_ba_service.Provision_id, oblig_ba_service.Rate_schedule_id, oblig_ba_service.Remark, oblig_ba_service.Service_quality, oblig_ba_service.Service_type, oblig_ba_service.Source, oblig_ba_service.Start_date, oblig_ba_service.Row_changed_by, oblig_ba_service.Row_changed_date, oblig_ba_service.Row_created_by, oblig_ba_service.Row_effective_date, oblig_ba_service.Row_expiry_date, oblig_ba_service.Row_quality, oblig_ba_service.Obligation_id)
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

func PatchObligBaService(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update oblig_ba_service set "
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
	query += " where obligation_id = :id"

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

func DeleteObligBaService(c *fiber.Ctx) error {
	id := c.Params("id")
	var oblig_ba_service dto.Oblig_ba_service
	oblig_ba_service.Obligation_id = id

	stmt, err := db.Prepare("delete from oblig_ba_service where obligation_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(oblig_ba_service.Obligation_id)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellBaService(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_ba_service")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_ba_service

	for rows.Next() {
		var well_ba_service dto.Well_ba_service
		if err := rows.Scan(&well_ba_service.Uwi, &well_ba_service.Provided_by, &well_ba_service.Service_seq_no, &well_ba_service.Active_ind, &well_ba_service.Contact_ba_id, &well_ba_service.Contract_id, &well_ba_service.Description, &well_ba_service.Effective_date, &well_ba_service.End_date, &well_ba_service.Expiry_date, &well_ba_service.Ppdm_guid, &well_ba_service.Provided_for, &well_ba_service.Provision_id, &well_ba_service.Rate_schedule_id, &well_ba_service.Remark, &well_ba_service.Represented_ba_id, &well_ba_service.Service_date, &well_ba_service.Service_period, &well_ba_service.Service_period_uom, &well_ba_service.Service_quality, &well_ba_service.Service_time, &well_ba_service.Service_timezone, &well_ba_service.Service_time_desc, &well_ba_service.Service_type, &well_ba_service.Source, &well_ba_service.Start_date, &well_ba_service.Well_activity_obs_no, &well_ba_service.Well_activity_source, &well_ba_service.Well_drill_period_obs_no, &well_ba_service.Row_changed_by, &well_ba_service.Row_changed_date, &well_ba_service.Row_created_by, &well_ba_service.Row_created_date, &well_ba_service.Row_effective_date, &well_ba_service.Row_expiry_date, &well_ba_service.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_ba_service to result
		result = append(result, well_ba_service)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellBaService(c *fiber.Ctx) error {
	var well_ba_service dto.Well_ba_service

	setDefaults(&well_ba_service)

	if err := json.Unmarshal(c.Body(), &well_ba_service); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_ba_service values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36)")
	if err != nil {
		return err
	}
	well_ba_service.Row_created_date = formatDate(well_ba_service.Row_created_date)
	well_ba_service.Row_changed_date = formatDate(well_ba_service.Row_changed_date)
	well_ba_service.Effective_date = formatDateString(well_ba_service.Effective_date)
	well_ba_service.End_date = formatDateString(well_ba_service.End_date)
	well_ba_service.Expiry_date = formatDateString(well_ba_service.Expiry_date)
	well_ba_service.Service_date = formatDateString(well_ba_service.Service_date)
	well_ba_service.Service_time = formatDateString(well_ba_service.Service_time)
	well_ba_service.Start_date = formatDateString(well_ba_service.Start_date)
	well_ba_service.Row_effective_date = formatDateString(well_ba_service.Row_effective_date)
	well_ba_service.Row_expiry_date = formatDateString(well_ba_service.Row_expiry_date)






	rows, err := stmt.Exec(well_ba_service.Uwi, well_ba_service.Provided_by, well_ba_service.Service_seq_no, well_ba_service.Active_ind, well_ba_service.Contact_ba_id, well_ba_service.Contract_id, well_ba_service.Description, well_ba_service.Effective_date, well_ba_service.End_date, well_ba_service.Expiry_date, well_ba_service.Ppdm_guid, well_ba_service.Provided_for, well_ba_service.Provision_id, well_ba_service.Rate_schedule_id, well_ba_service.Remark, well_ba_service.Represented_ba_id, well_ba_service.Service_date, well_ba_service.Service_period, well_ba_service.Service_period_uom, well_ba_service.Service_quality, well_ba_service.Service_time, well_ba_service.Service_timezone, well_ba_service.Service_time_desc, well_ba_service.Service_type, well_ba_service.Source, well_ba_service.Start_date, well_ba_service.Well_activity_obs_no, well_ba_service.Well_activity_source, well_ba_service.Well_drill_period_obs_no, well_ba_service.Row_changed_by, well_ba_service.Row_changed_date, well_ba_service.Row_created_by, well_ba_service.Row_created_date, well_ba_service.Row_effective_date, well_ba_service.Row_expiry_date, well_ba_service.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellBaService(c *fiber.Ctx) error {
	var well_ba_service dto.Well_ba_service

	setDefaults(&well_ba_service)

	if err := json.Unmarshal(c.Body(), &well_ba_service); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_ba_service.Uwi = id

    if well_ba_service.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_ba_service set provided_by = :1, service_seq_no = :2, active_ind = :3, contact_ba_id = :4, contract_id = :5, description = :6, effective_date = :7, end_date = :8, expiry_date = :9, ppdm_guid = :10, provided_for = :11, provision_id = :12, rate_schedule_id = :13, remark = :14, represented_ba_id = :15, service_date = :16, service_period = :17, service_period_uom = :18, service_quality = :19, service_time = :20, service_timezone = :21, service_time_desc = :22, service_type = :23, source = :24, start_date = :25, well_activity_obs_no = :26, well_activity_source = :27, well_drill_period_obs_no = :28, row_changed_by = :29, row_changed_date = :30, row_created_by = :31, row_effective_date = :32, row_expiry_date = :33, row_quality = :34 where uwi = :36")
	if err != nil {
		return err
	}

	well_ba_service.Row_changed_date = formatDate(well_ba_service.Row_changed_date)
	well_ba_service.Effective_date = formatDateString(well_ba_service.Effective_date)
	well_ba_service.End_date = formatDateString(well_ba_service.End_date)
	well_ba_service.Expiry_date = formatDateString(well_ba_service.Expiry_date)
	well_ba_service.Service_date = formatDateString(well_ba_service.Service_date)
	well_ba_service.Service_time = formatDateString(well_ba_service.Service_time)
	well_ba_service.Start_date = formatDateString(well_ba_service.Start_date)
	well_ba_service.Row_effective_date = formatDateString(well_ba_service.Row_effective_date)
	well_ba_service.Row_expiry_date = formatDateString(well_ba_service.Row_expiry_date)






	rows, err := stmt.Exec(well_ba_service.Provided_by, well_ba_service.Service_seq_no, well_ba_service.Active_ind, well_ba_service.Contact_ba_id, well_ba_service.Contract_id, well_ba_service.Description, well_ba_service.Effective_date, well_ba_service.End_date, well_ba_service.Expiry_date, well_ba_service.Ppdm_guid, well_ba_service.Provided_for, well_ba_service.Provision_id, well_ba_service.Rate_schedule_id, well_ba_service.Remark, well_ba_service.Represented_ba_id, well_ba_service.Service_date, well_ba_service.Service_period, well_ba_service.Service_period_uom, well_ba_service.Service_quality, well_ba_service.Service_time, well_ba_service.Service_timezone, well_ba_service.Service_time_desc, well_ba_service.Service_type, well_ba_service.Source, well_ba_service.Start_date, well_ba_service.Well_activity_obs_no, well_ba_service.Well_activity_source, well_ba_service.Well_drill_period_obs_no, well_ba_service.Row_changed_by, well_ba_service.Row_changed_date, well_ba_service.Row_created_by, well_ba_service.Row_effective_date, well_ba_service.Row_expiry_date, well_ba_service.Row_quality, well_ba_service.Uwi)
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

func PatchWellBaService(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_ba_service set "
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
		} else if key == "effective_date" || key == "end_date" || key == "expiry_date" || key == "service_date" || key == "service_time" || key == "start_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where uwi = :id"

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

func DeleteWellBaService(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_ba_service dto.Well_ba_service
	well_ba_service.Uwi = id

	stmt, err := db.Prepare("delete from well_ba_service where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_ba_service.Uwi)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLandBaService(c *fiber.Ctx) error {
	rows, err := db.Query("select * from land_ba_service")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Land_ba_service

	for rows.Next() {
		var land_ba_service dto.Land_ba_service
		if err := rows.Scan(&land_ba_service.Land_right_subtype, &land_ba_service.Land_right_id, &land_ba_service.Provided_by, &land_ba_service.Service_seq_no, &land_ba_service.Active_ind, &land_ba_service.Ba_service_type, &land_ba_service.Contact_ba_id, &land_ba_service.Contract_id, &land_ba_service.Description, &land_ba_service.Effective_date, &land_ba_service.End_date, &land_ba_service.Expiry_date, &land_ba_service.Organization_id, &land_ba_service.Organization_seq_no, &land_ba_service.Ppdm_guid, &land_ba_service.Provided_for_ba_id, &land_ba_service.Provision_id, &land_ba_service.Rate_schedule_id, &land_ba_service.Remark, &land_ba_service.Represented_ba_id, &land_ba_service.Service_quality, &land_ba_service.Source, &land_ba_service.Start_date, &land_ba_service.Row_changed_by, &land_ba_service.Row_changed_date, &land_ba_service.Row_created_by, &land_ba_service.Row_created_date, &land_ba_service.Row_effective_date, &land_ba_service.Row_expiry_date, &land_ba_service.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append land_ba_service to result
		result = append(result, land_ba_service)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLandBaService(c *fiber.Ctx) error {
	var land_ba_service dto.Land_ba_service

	setDefaults(&land_ba_service)

	if err := json.Unmarshal(c.Body(), &land_ba_service); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into land_ba_service values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30)")
	if err != nil {
		return err
	}
	land_ba_service.Row_created_date = formatDate(land_ba_service.Row_created_date)
	land_ba_service.Row_changed_date = formatDate(land_ba_service.Row_changed_date)
	land_ba_service.Effective_date = formatDateString(land_ba_service.Effective_date)
	land_ba_service.End_date = formatDateString(land_ba_service.End_date)
	land_ba_service.Expiry_date = formatDateString(land_ba_service.Expiry_date)
	land_ba_service.Start_date = formatDateString(land_ba_service.Start_date)
	land_ba_service.Row_effective_date = formatDateString(land_ba_service.Row_effective_date)
	land_ba_service.Row_expiry_date = formatDateString(land_ba_service.Row_expiry_date)






	rows, err := stmt.Exec(land_ba_service.Land_right_subtype, land_ba_service.Land_right_id, land_ba_service.Provided_by, land_ba_service.Service_seq_no, land_ba_service.Active_ind, land_ba_service.Ba_service_type, land_ba_service.Contact_ba_id, land_ba_service.Contract_id, land_ba_service.Description, land_ba_service.Effective_date, land_ba_service.End_date, land_ba_service.Expiry_date, land_ba_service.Organization_id, land_ba_service.Organization_seq_no, land_ba_service.Ppdm_guid, land_ba_service.Provided_for_ba_id, land_ba_service.Provision_id, land_ba_service.Rate_schedule_id, land_ba_service.Remark, land_ba_service.Represented_ba_id, land_ba_service.Service_quality, land_ba_service.Source, land_ba_service.Start_date, land_ba_service.Row_changed_by, land_ba_service.Row_changed_date, land_ba_service.Row_created_by, land_ba_service.Row_created_date, land_ba_service.Row_effective_date, land_ba_service.Row_expiry_date, land_ba_service.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLandBaService(c *fiber.Ctx) error {
	var land_ba_service dto.Land_ba_service

	setDefaults(&land_ba_service)

	if err := json.Unmarshal(c.Body(), &land_ba_service); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	land_ba_service.Land_right_subtype = id

    if land_ba_service.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update land_ba_service set land_right_id = :1, provided_by = :2, service_seq_no = :3, active_ind = :4, ba_service_type = :5, contact_ba_id = :6, contract_id = :7, description = :8, effective_date = :9, end_date = :10, expiry_date = :11, organization_id = :12, organization_seq_no = :13, ppdm_guid = :14, provided_for_ba_id = :15, provision_id = :16, rate_schedule_id = :17, remark = :18, represented_ba_id = :19, service_quality = :20, source = :21, start_date = :22, row_changed_by = :23, row_changed_date = :24, row_created_by = :25, row_effective_date = :26, row_expiry_date = :27, row_quality = :28 where land_right_subtype = :30")
	if err != nil {
		return err
	}

	land_ba_service.Row_changed_date = formatDate(land_ba_service.Row_changed_date)
	land_ba_service.Effective_date = formatDateString(land_ba_service.Effective_date)
	land_ba_service.End_date = formatDateString(land_ba_service.End_date)
	land_ba_service.Expiry_date = formatDateString(land_ba_service.Expiry_date)
	land_ba_service.Start_date = formatDateString(land_ba_service.Start_date)
	land_ba_service.Row_effective_date = formatDateString(land_ba_service.Row_effective_date)
	land_ba_service.Row_expiry_date = formatDateString(land_ba_service.Row_expiry_date)






	rows, err := stmt.Exec(land_ba_service.Land_right_id, land_ba_service.Provided_by, land_ba_service.Service_seq_no, land_ba_service.Active_ind, land_ba_service.Ba_service_type, land_ba_service.Contact_ba_id, land_ba_service.Contract_id, land_ba_service.Description, land_ba_service.Effective_date, land_ba_service.End_date, land_ba_service.Expiry_date, land_ba_service.Organization_id, land_ba_service.Organization_seq_no, land_ba_service.Ppdm_guid, land_ba_service.Provided_for_ba_id, land_ba_service.Provision_id, land_ba_service.Rate_schedule_id, land_ba_service.Remark, land_ba_service.Represented_ba_id, land_ba_service.Service_quality, land_ba_service.Source, land_ba_service.Start_date, land_ba_service.Row_changed_by, land_ba_service.Row_changed_date, land_ba_service.Row_created_by, land_ba_service.Row_effective_date, land_ba_service.Row_expiry_date, land_ba_service.Row_quality, land_ba_service.Land_right_subtype)
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

func PatchLandBaService(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update land_ba_service set "
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
	query += " where land_right_subtype = :id"

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

func DeleteLandBaService(c *fiber.Ctx) error {
	id := c.Params("id")
	var land_ba_service dto.Land_ba_service
	land_ba_service.Land_right_subtype = id

	stmt, err := db.Prepare("delete from land_ba_service where land_right_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(land_ba_service.Land_right_subtype)
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



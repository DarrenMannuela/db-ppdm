package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLandSaleBaService(c *fiber.Ctx) error {
	rows, err := db.Query("select * from land_sale_ba_service")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Land_sale_ba_service

	for rows.Next() {
		var land_sale_ba_service dto.Land_sale_ba_service
		if err := rows.Scan(&land_sale_ba_service.Jurisdiction, &land_sale_ba_service.Land_sale_number, &land_sale_ba_service.Land_sale_offering_id, &land_sale_ba_service.Provided_by, &land_sale_ba_service.Service_seq_no, &land_sale_ba_service.Active_ind, &land_sale_ba_service.Ba_service_type, &land_sale_ba_service.Contact_ba_id, &land_sale_ba_service.Description, &land_sale_ba_service.Effective_date, &land_sale_ba_service.End_date, &land_sale_ba_service.Expiry_date, &land_sale_ba_service.Organization_id, &land_sale_ba_service.Organization_seq_no, &land_sale_ba_service.Ppdm_guid, &land_sale_ba_service.Provided_for_ba_id, &land_sale_ba_service.Rate_schedule_id, &land_sale_ba_service.Remark, &land_sale_ba_service.Represented_ba_id, &land_sale_ba_service.Service_quality, &land_sale_ba_service.Source, &land_sale_ba_service.Start_date, &land_sale_ba_service.Row_changed_by, &land_sale_ba_service.Row_changed_date, &land_sale_ba_service.Row_created_by, &land_sale_ba_service.Row_created_date, &land_sale_ba_service.Row_effective_date, &land_sale_ba_service.Row_expiry_date, &land_sale_ba_service.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append land_sale_ba_service to result
		result = append(result, land_sale_ba_service)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLandSaleBaService(c *fiber.Ctx) error {
	var land_sale_ba_service dto.Land_sale_ba_service

	setDefaults(&land_sale_ba_service)

	if err := json.Unmarshal(c.Body(), &land_sale_ba_service); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into land_sale_ba_service values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29)")
	if err != nil {
		return err
	}
	land_sale_ba_service.Row_created_date = formatDate(land_sale_ba_service.Row_created_date)
	land_sale_ba_service.Row_changed_date = formatDate(land_sale_ba_service.Row_changed_date)
	land_sale_ba_service.Effective_date = formatDateString(land_sale_ba_service.Effective_date)
	land_sale_ba_service.End_date = formatDateString(land_sale_ba_service.End_date)
	land_sale_ba_service.Expiry_date = formatDateString(land_sale_ba_service.Expiry_date)
	land_sale_ba_service.Start_date = formatDateString(land_sale_ba_service.Start_date)
	land_sale_ba_service.Row_effective_date = formatDateString(land_sale_ba_service.Row_effective_date)
	land_sale_ba_service.Row_expiry_date = formatDateString(land_sale_ba_service.Row_expiry_date)






	rows, err := stmt.Exec(land_sale_ba_service.Jurisdiction, land_sale_ba_service.Land_sale_number, land_sale_ba_service.Land_sale_offering_id, land_sale_ba_service.Provided_by, land_sale_ba_service.Service_seq_no, land_sale_ba_service.Active_ind, land_sale_ba_service.Ba_service_type, land_sale_ba_service.Contact_ba_id, land_sale_ba_service.Description, land_sale_ba_service.Effective_date, land_sale_ba_service.End_date, land_sale_ba_service.Expiry_date, land_sale_ba_service.Organization_id, land_sale_ba_service.Organization_seq_no, land_sale_ba_service.Ppdm_guid, land_sale_ba_service.Provided_for_ba_id, land_sale_ba_service.Rate_schedule_id, land_sale_ba_service.Remark, land_sale_ba_service.Represented_ba_id, land_sale_ba_service.Service_quality, land_sale_ba_service.Source, land_sale_ba_service.Start_date, land_sale_ba_service.Row_changed_by, land_sale_ba_service.Row_changed_date, land_sale_ba_service.Row_created_by, land_sale_ba_service.Row_created_date, land_sale_ba_service.Row_effective_date, land_sale_ba_service.Row_expiry_date, land_sale_ba_service.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLandSaleBaService(c *fiber.Ctx) error {
	var land_sale_ba_service dto.Land_sale_ba_service

	setDefaults(&land_sale_ba_service)

	if err := json.Unmarshal(c.Body(), &land_sale_ba_service); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	land_sale_ba_service.Jurisdiction = id

    if land_sale_ba_service.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update land_sale_ba_service set land_sale_number = :1, land_sale_offering_id = :2, provided_by = :3, service_seq_no = :4, active_ind = :5, ba_service_type = :6, contact_ba_id = :7, description = :8, effective_date = :9, end_date = :10, expiry_date = :11, organization_id = :12, organization_seq_no = :13, ppdm_guid = :14, provided_for_ba_id = :15, rate_schedule_id = :16, remark = :17, represented_ba_id = :18, service_quality = :19, source = :20, start_date = :21, row_changed_by = :22, row_changed_date = :23, row_created_by = :24, row_effective_date = :25, row_expiry_date = :26, row_quality = :27 where jurisdiction = :29")
	if err != nil {
		return err
	}

	land_sale_ba_service.Row_changed_date = formatDate(land_sale_ba_service.Row_changed_date)
	land_sale_ba_service.Effective_date = formatDateString(land_sale_ba_service.Effective_date)
	land_sale_ba_service.End_date = formatDateString(land_sale_ba_service.End_date)
	land_sale_ba_service.Expiry_date = formatDateString(land_sale_ba_service.Expiry_date)
	land_sale_ba_service.Start_date = formatDateString(land_sale_ba_service.Start_date)
	land_sale_ba_service.Row_effective_date = formatDateString(land_sale_ba_service.Row_effective_date)
	land_sale_ba_service.Row_expiry_date = formatDateString(land_sale_ba_service.Row_expiry_date)






	rows, err := stmt.Exec(land_sale_ba_service.Land_sale_number, land_sale_ba_service.Land_sale_offering_id, land_sale_ba_service.Provided_by, land_sale_ba_service.Service_seq_no, land_sale_ba_service.Active_ind, land_sale_ba_service.Ba_service_type, land_sale_ba_service.Contact_ba_id, land_sale_ba_service.Description, land_sale_ba_service.Effective_date, land_sale_ba_service.End_date, land_sale_ba_service.Expiry_date, land_sale_ba_service.Organization_id, land_sale_ba_service.Organization_seq_no, land_sale_ba_service.Ppdm_guid, land_sale_ba_service.Provided_for_ba_id, land_sale_ba_service.Rate_schedule_id, land_sale_ba_service.Remark, land_sale_ba_service.Represented_ba_id, land_sale_ba_service.Service_quality, land_sale_ba_service.Source, land_sale_ba_service.Start_date, land_sale_ba_service.Row_changed_by, land_sale_ba_service.Row_changed_date, land_sale_ba_service.Row_created_by, land_sale_ba_service.Row_effective_date, land_sale_ba_service.Row_expiry_date, land_sale_ba_service.Row_quality, land_sale_ba_service.Jurisdiction)
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

func PatchLandSaleBaService(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update land_sale_ba_service set "
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

func DeleteLandSaleBaService(c *fiber.Ctx) error {
	id := c.Params("id")
	var land_sale_ba_service dto.Land_sale_ba_service
	land_sale_ba_service.Jurisdiction = id

	stmt, err := db.Prepare("delete from land_sale_ba_service where jurisdiction = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(land_sale_ba_service.Jurisdiction)
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



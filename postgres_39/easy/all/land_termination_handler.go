package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLandTermination(c *fiber.Ctx) error {
	rows, err := db.Query("select * from land_termination")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Land_termination

	for rows.Next() {
		var land_termination dto.Land_termination
		if err := rows.Scan(&land_termination.Lr_termination_id, &land_termination.Lr_termination_seq_no, &land_termination.Active_ind, &land_termination.Business_associate_id, &land_termination.Description, &land_termination.Effective_date, &land_termination.Expiry_date, &land_termination.Fulfilled_date, &land_termination.Fulfilled_user, &land_termination.Jurisdiction, &land_termination.Land_request_id, &land_termination.Land_right_id, &land_termination.Land_right_subtype, &land_termination.Land_sale_number, &land_termination.Land_sale_offering_id, &land_termination.Mineral_zone_id, &land_termination.Notification_id, &land_termination.Ppdm_guid, &land_termination.Remark, &land_termination.Source, &land_termination.Spatial_description_id, &land_termination.Spatial_obs_no, &land_termination.Termination_reqmt, &land_termination.Termination_type, &land_termination.Row_changed_by, &land_termination.Row_changed_date, &land_termination.Row_created_by, &land_termination.Row_created_date, &land_termination.Row_effective_date, &land_termination.Row_expiry_date, &land_termination.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append land_termination to result
		result = append(result, land_termination)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLandTermination(c *fiber.Ctx) error {
	var land_termination dto.Land_termination

	setDefaults(&land_termination)

	if err := json.Unmarshal(c.Body(), &land_termination); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into land_termination values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31)")
	if err != nil {
		return err
	}
	land_termination.Row_created_date = formatDate(land_termination.Row_created_date)
	land_termination.Row_changed_date = formatDate(land_termination.Row_changed_date)
	land_termination.Effective_date = formatDateString(land_termination.Effective_date)
	land_termination.Expiry_date = formatDateString(land_termination.Expiry_date)
	land_termination.Fulfilled_date = formatDateString(land_termination.Fulfilled_date)
	land_termination.Row_effective_date = formatDateString(land_termination.Row_effective_date)
	land_termination.Row_expiry_date = formatDateString(land_termination.Row_expiry_date)






	rows, err := stmt.Exec(land_termination.Lr_termination_id, land_termination.Lr_termination_seq_no, land_termination.Active_ind, land_termination.Business_associate_id, land_termination.Description, land_termination.Effective_date, land_termination.Expiry_date, land_termination.Fulfilled_date, land_termination.Fulfilled_user, land_termination.Jurisdiction, land_termination.Land_request_id, land_termination.Land_right_id, land_termination.Land_right_subtype, land_termination.Land_sale_number, land_termination.Land_sale_offering_id, land_termination.Mineral_zone_id, land_termination.Notification_id, land_termination.Ppdm_guid, land_termination.Remark, land_termination.Source, land_termination.Spatial_description_id, land_termination.Spatial_obs_no, land_termination.Termination_reqmt, land_termination.Termination_type, land_termination.Row_changed_by, land_termination.Row_changed_date, land_termination.Row_created_by, land_termination.Row_created_date, land_termination.Row_effective_date, land_termination.Row_expiry_date, land_termination.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLandTermination(c *fiber.Ctx) error {
	var land_termination dto.Land_termination

	setDefaults(&land_termination)

	if err := json.Unmarshal(c.Body(), &land_termination); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	land_termination.Lr_termination_id = id

    if land_termination.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update land_termination set lr_termination_seq_no = :1, active_ind = :2, business_associate_id = :3, description = :4, effective_date = :5, expiry_date = :6, fulfilled_date = :7, fulfilled_user = :8, jurisdiction = :9, land_request_id = :10, land_right_id = :11, land_right_subtype = :12, land_sale_number = :13, land_sale_offering_id = :14, mineral_zone_id = :15, notification_id = :16, ppdm_guid = :17, remark = :18, source = :19, spatial_description_id = :20, spatial_obs_no = :21, termination_reqmt = :22, termination_type = :23, row_changed_by = :24, row_changed_date = :25, row_created_by = :26, row_effective_date = :27, row_expiry_date = :28, row_quality = :29 where lr_termination_id = :31")
	if err != nil {
		return err
	}

	land_termination.Row_changed_date = formatDate(land_termination.Row_changed_date)
	land_termination.Effective_date = formatDateString(land_termination.Effective_date)
	land_termination.Expiry_date = formatDateString(land_termination.Expiry_date)
	land_termination.Fulfilled_date = formatDateString(land_termination.Fulfilled_date)
	land_termination.Row_effective_date = formatDateString(land_termination.Row_effective_date)
	land_termination.Row_expiry_date = formatDateString(land_termination.Row_expiry_date)






	rows, err := stmt.Exec(land_termination.Lr_termination_seq_no, land_termination.Active_ind, land_termination.Business_associate_id, land_termination.Description, land_termination.Effective_date, land_termination.Expiry_date, land_termination.Fulfilled_date, land_termination.Fulfilled_user, land_termination.Jurisdiction, land_termination.Land_request_id, land_termination.Land_right_id, land_termination.Land_right_subtype, land_termination.Land_sale_number, land_termination.Land_sale_offering_id, land_termination.Mineral_zone_id, land_termination.Notification_id, land_termination.Ppdm_guid, land_termination.Remark, land_termination.Source, land_termination.Spatial_description_id, land_termination.Spatial_obs_no, land_termination.Termination_reqmt, land_termination.Termination_type, land_termination.Row_changed_by, land_termination.Row_changed_date, land_termination.Row_created_by, land_termination.Row_effective_date, land_termination.Row_expiry_date, land_termination.Row_quality, land_termination.Lr_termination_id)
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

func PatchLandTermination(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update land_termination set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "fulfilled_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where lr_termination_id = :id"

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

func DeleteLandTermination(c *fiber.Ctx) error {
	id := c.Params("id")
	var land_termination dto.Land_termination
	land_termination.Lr_termination_id = id

	stmt, err := db.Prepare("delete from land_termination where lr_termination_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(land_termination.Lr_termination_id)
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



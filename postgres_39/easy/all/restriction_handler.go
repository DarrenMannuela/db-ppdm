package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRestriction(c *fiber.Ctx) error {
	rows, err := db.Query("select * from restriction")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Restriction

	for rows.Next() {
		var restriction dto.Restriction
		if err := rows.Scan(&restriction.Restriction_id, &restriction.Restriction_version, &restriction.Active_ind, &restriction.Description, &restriction.Effective_date, &restriction.End_date, &restriction.End_date_event, &restriction.Expiry_date, &restriction.Ppdm_guid, &restriction.Regulatory_act, &restriction.Remark, &restriction.Restriction_class, &restriction.Restriction_name, &restriction.Restriction_type, &restriction.Rest_duration_type, &restriction.Source, &restriction.Source_document_id, &restriction.Start_date, &restriction.Start_date_event, &restriction.Row_changed_by, &restriction.Row_changed_date, &restriction.Row_created_by, &restriction.Row_created_date, &restriction.Row_effective_date, &restriction.Row_expiry_date, &restriction.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append restriction to result
		result = append(result, restriction)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRestriction(c *fiber.Ctx) error {
	var restriction dto.Restriction

	setDefaults(&restriction)

	if err := json.Unmarshal(c.Body(), &restriction); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into restriction values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26)")
	if err != nil {
		return err
	}
	restriction.Row_created_date = formatDate(restriction.Row_created_date)
	restriction.Row_changed_date = formatDate(restriction.Row_changed_date)
	restriction.Effective_date = formatDateString(restriction.Effective_date)
	restriction.End_date = formatDateString(restriction.End_date)
	restriction.End_date_event = formatDateString(restriction.End_date_event)
	restriction.Expiry_date = formatDateString(restriction.Expiry_date)
	restriction.Start_date = formatDateString(restriction.Start_date)
	restriction.Start_date_event = formatDateString(restriction.Start_date_event)
	restriction.Row_effective_date = formatDateString(restriction.Row_effective_date)
	restriction.Row_expiry_date = formatDateString(restriction.Row_expiry_date)






	rows, err := stmt.Exec(restriction.Restriction_id, restriction.Restriction_version, restriction.Active_ind, restriction.Description, restriction.Effective_date, restriction.End_date, restriction.End_date_event, restriction.Expiry_date, restriction.Ppdm_guid, restriction.Regulatory_act, restriction.Remark, restriction.Restriction_class, restriction.Restriction_name, restriction.Restriction_type, restriction.Rest_duration_type, restriction.Source, restriction.Source_document_id, restriction.Start_date, restriction.Start_date_event, restriction.Row_changed_by, restriction.Row_changed_date, restriction.Row_created_by, restriction.Row_created_date, restriction.Row_effective_date, restriction.Row_expiry_date, restriction.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRestriction(c *fiber.Ctx) error {
	var restriction dto.Restriction

	setDefaults(&restriction)

	if err := json.Unmarshal(c.Body(), &restriction); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	restriction.Restriction_id = id

    if restriction.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update restriction set restriction_version = :1, active_ind = :2, description = :3, effective_date = :4, end_date = :5, end_date_event = :6, expiry_date = :7, ppdm_guid = :8, regulatory_act = :9, remark = :10, restriction_class = :11, restriction_name = :12, restriction_type = :13, rest_duration_type = :14, source = :15, source_document_id = :16, start_date = :17, start_date_event = :18, row_changed_by = :19, row_changed_date = :20, row_created_by = :21, row_effective_date = :22, row_expiry_date = :23, row_quality = :24 where restriction_id = :26")
	if err != nil {
		return err
	}

	restriction.Row_changed_date = formatDate(restriction.Row_changed_date)
	restriction.Effective_date = formatDateString(restriction.Effective_date)
	restriction.End_date = formatDateString(restriction.End_date)
	restriction.End_date_event = formatDateString(restriction.End_date_event)
	restriction.Expiry_date = formatDateString(restriction.Expiry_date)
	restriction.Start_date = formatDateString(restriction.Start_date)
	restriction.Start_date_event = formatDateString(restriction.Start_date_event)
	restriction.Row_effective_date = formatDateString(restriction.Row_effective_date)
	restriction.Row_expiry_date = formatDateString(restriction.Row_expiry_date)






	rows, err := stmt.Exec(restriction.Restriction_version, restriction.Active_ind, restriction.Description, restriction.Effective_date, restriction.End_date, restriction.End_date_event, restriction.Expiry_date, restriction.Ppdm_guid, restriction.Regulatory_act, restriction.Remark, restriction.Restriction_class, restriction.Restriction_name, restriction.Restriction_type, restriction.Rest_duration_type, restriction.Source, restriction.Source_document_id, restriction.Start_date, restriction.Start_date_event, restriction.Row_changed_by, restriction.Row_changed_date, restriction.Row_created_by, restriction.Row_effective_date, restriction.Row_expiry_date, restriction.Row_quality, restriction.Restriction_id)
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

func PatchRestriction(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update restriction set "
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
		} else if key == "effective_date" || key == "end_date" || key == "end_date_event" || key == "expiry_date" || key == "start_date" || key == "start_date_event" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where restriction_id = :id"

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

func DeleteRestriction(c *fiber.Ctx) error {
	id := c.Params("id")
	var restriction dto.Restriction
	restriction.Restriction_id = id

	stmt, err := db.Prepare("delete from restriction where restriction_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(restriction.Restriction_id)
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



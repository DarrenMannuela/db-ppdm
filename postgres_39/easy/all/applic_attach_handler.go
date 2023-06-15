package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetApplicAttach(c *fiber.Ctx) error {
	rows, err := db.Query("select * from applic_attach")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Applic_attach

	for rows.Next() {
		var applic_attach dto.Applic_attach
		if err := rows.Scan(&applic_attach.Application_id, &applic_attach.Attachment_id, &applic_attach.Active_ind, &applic_attach.Attachment_description, &applic_attach.Attachment_type, &applic_attach.Effective_date, &applic_attach.Expiry_date, &applic_attach.Physical_item_id, &applic_attach.Ppdm_guid, &applic_attach.Received_date, &applic_attach.Received_ind, &applic_attach.Remark, &applic_attach.Sent_by, &applic_attach.Sent_date, &applic_attach.Sent_ind, &applic_attach.Source, &applic_attach.Row_changed_by, &applic_attach.Row_changed_date, &applic_attach.Row_created_by, &applic_attach.Row_created_date, &applic_attach.Row_effective_date, &applic_attach.Row_expiry_date, &applic_attach.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append applic_attach to result
		result = append(result, applic_attach)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetApplicAttach(c *fiber.Ctx) error {
	var applic_attach dto.Applic_attach

	setDefaults(&applic_attach)

	if err := json.Unmarshal(c.Body(), &applic_attach); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into applic_attach values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	applic_attach.Row_created_date = formatDate(applic_attach.Row_created_date)
	applic_attach.Row_changed_date = formatDate(applic_attach.Row_changed_date)
	applic_attach.Effective_date = formatDateString(applic_attach.Effective_date)
	applic_attach.Expiry_date = formatDateString(applic_attach.Expiry_date)
	applic_attach.Received_date = formatDateString(applic_attach.Received_date)
	applic_attach.Sent_date = formatDateString(applic_attach.Sent_date)
	applic_attach.Row_effective_date = formatDateString(applic_attach.Row_effective_date)
	applic_attach.Row_expiry_date = formatDateString(applic_attach.Row_expiry_date)






	rows, err := stmt.Exec(applic_attach.Application_id, applic_attach.Attachment_id, applic_attach.Active_ind, applic_attach.Attachment_description, applic_attach.Attachment_type, applic_attach.Effective_date, applic_attach.Expiry_date, applic_attach.Physical_item_id, applic_attach.Ppdm_guid, applic_attach.Received_date, applic_attach.Received_ind, applic_attach.Remark, applic_attach.Sent_by, applic_attach.Sent_date, applic_attach.Sent_ind, applic_attach.Source, applic_attach.Row_changed_by, applic_attach.Row_changed_date, applic_attach.Row_created_by, applic_attach.Row_created_date, applic_attach.Row_effective_date, applic_attach.Row_expiry_date, applic_attach.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateApplicAttach(c *fiber.Ctx) error {
	var applic_attach dto.Applic_attach

	setDefaults(&applic_attach)

	if err := json.Unmarshal(c.Body(), &applic_attach); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	applic_attach.Application_id = id

    if applic_attach.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update applic_attach set attachment_id = :1, active_ind = :2, attachment_description = :3, attachment_type = :4, effective_date = :5, expiry_date = :6, physical_item_id = :7, ppdm_guid = :8, received_date = :9, received_ind = :10, remark = :11, sent_by = :12, sent_date = :13, sent_ind = :14, source = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where application_id = :23")
	if err != nil {
		return err
	}

	applic_attach.Row_changed_date = formatDate(applic_attach.Row_changed_date)
	applic_attach.Effective_date = formatDateString(applic_attach.Effective_date)
	applic_attach.Expiry_date = formatDateString(applic_attach.Expiry_date)
	applic_attach.Received_date = formatDateString(applic_attach.Received_date)
	applic_attach.Sent_date = formatDateString(applic_attach.Sent_date)
	applic_attach.Row_effective_date = formatDateString(applic_attach.Row_effective_date)
	applic_attach.Row_expiry_date = formatDateString(applic_attach.Row_expiry_date)






	rows, err := stmt.Exec(applic_attach.Attachment_id, applic_attach.Active_ind, applic_attach.Attachment_description, applic_attach.Attachment_type, applic_attach.Effective_date, applic_attach.Expiry_date, applic_attach.Physical_item_id, applic_attach.Ppdm_guid, applic_attach.Received_date, applic_attach.Received_ind, applic_attach.Remark, applic_attach.Sent_by, applic_attach.Sent_date, applic_attach.Sent_ind, applic_attach.Source, applic_attach.Row_changed_by, applic_attach.Row_changed_date, applic_attach.Row_created_by, applic_attach.Row_effective_date, applic_attach.Row_expiry_date, applic_attach.Row_quality, applic_attach.Application_id)
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

func PatchApplicAttach(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update applic_attach set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "received_date" || key == "sent_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where application_id = :id"

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

func DeleteApplicAttach(c *fiber.Ctx) error {
	id := c.Params("id")
	var applic_attach dto.Applic_attach
	applic_attach.Application_id = id

	stmt, err := db.Prepare("delete from applic_attach where application_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(applic_attach.Application_id)
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



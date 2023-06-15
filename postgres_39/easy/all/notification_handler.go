package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetNotification(c *fiber.Ctx) error {
	rows, err := db.Query("select * from notification")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Notification

	for rows.Next() {
		var notification dto.Notification
		if err := rows.Scan(&notification.Notification_id, &notification.Active_ind, &notification.Business_associate_id, &notification.Contract_ind, &notification.Critical_date, &notification.Effective_date, &notification.Expiry_date, &notification.Land_right_id, &notification.Land_right_ind, &notification.Land_right_subtype, &notification.Notification_type, &notification.Obligation_id, &notification.Obligation_ind, &notification.Obligation_seq_no, &notification.Ppdm_guid, &notification.Remark, &notification.Served_ind, &notification.Source, &notification.Row_changed_by, &notification.Row_changed_date, &notification.Row_created_by, &notification.Row_created_date, &notification.Row_effective_date, &notification.Row_expiry_date, &notification.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append notification to result
		result = append(result, notification)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetNotification(c *fiber.Ctx) error {
	var notification dto.Notification

	setDefaults(&notification)

	if err := json.Unmarshal(c.Body(), &notification); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into notification values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25)")
	if err != nil {
		return err
	}
	notification.Row_created_date = formatDate(notification.Row_created_date)
	notification.Row_changed_date = formatDate(notification.Row_changed_date)
	notification.Critical_date = formatDateString(notification.Critical_date)
	notification.Effective_date = formatDateString(notification.Effective_date)
	notification.Expiry_date = formatDateString(notification.Expiry_date)
	notification.Row_effective_date = formatDateString(notification.Row_effective_date)
	notification.Row_expiry_date = formatDateString(notification.Row_expiry_date)






	rows, err := stmt.Exec(notification.Notification_id, notification.Active_ind, notification.Business_associate_id, notification.Contract_ind, notification.Critical_date, notification.Effective_date, notification.Expiry_date, notification.Land_right_id, notification.Land_right_ind, notification.Land_right_subtype, notification.Notification_type, notification.Obligation_id, notification.Obligation_ind, notification.Obligation_seq_no, notification.Ppdm_guid, notification.Remark, notification.Served_ind, notification.Source, notification.Row_changed_by, notification.Row_changed_date, notification.Row_created_by, notification.Row_created_date, notification.Row_effective_date, notification.Row_expiry_date, notification.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateNotification(c *fiber.Ctx) error {
	var notification dto.Notification

	setDefaults(&notification)

	if err := json.Unmarshal(c.Body(), &notification); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	notification.Notification_id = id

    if notification.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update notification set active_ind = :1, business_associate_id = :2, contract_ind = :3, critical_date = :4, effective_date = :5, expiry_date = :6, land_right_id = :7, land_right_ind = :8, land_right_subtype = :9, notification_type = :10, obligation_id = :11, obligation_ind = :12, obligation_seq_no = :13, ppdm_guid = :14, remark = :15, served_ind = :16, source = :17, row_changed_by = :18, row_changed_date = :19, row_created_by = :20, row_effective_date = :21, row_expiry_date = :22, row_quality = :23 where notification_id = :25")
	if err != nil {
		return err
	}

	notification.Row_changed_date = formatDate(notification.Row_changed_date)
	notification.Critical_date = formatDateString(notification.Critical_date)
	notification.Effective_date = formatDateString(notification.Effective_date)
	notification.Expiry_date = formatDateString(notification.Expiry_date)
	notification.Row_effective_date = formatDateString(notification.Row_effective_date)
	notification.Row_expiry_date = formatDateString(notification.Row_expiry_date)






	rows, err := stmt.Exec(notification.Active_ind, notification.Business_associate_id, notification.Contract_ind, notification.Critical_date, notification.Effective_date, notification.Expiry_date, notification.Land_right_id, notification.Land_right_ind, notification.Land_right_subtype, notification.Notification_type, notification.Obligation_id, notification.Obligation_ind, notification.Obligation_seq_no, notification.Ppdm_guid, notification.Remark, notification.Served_ind, notification.Source, notification.Row_changed_by, notification.Row_changed_date, notification.Row_created_by, notification.Row_effective_date, notification.Row_expiry_date, notification.Row_quality, notification.Notification_id)
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

func PatchNotification(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update notification set "
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
		} else if key == "critical_date" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where notification_id = :id"

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

func DeleteNotification(c *fiber.Ctx) error {
	id := c.Params("id")
	var notification dto.Notification
	notification.Notification_id = id

	stmt, err := db.Prepare("delete from notification where notification_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(notification.Notification_id)
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



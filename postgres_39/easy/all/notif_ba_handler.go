package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetNotifBa(c *fiber.Ctx) error {
	rows, err := db.Query("select * from notif_ba")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Notif_ba

	for rows.Next() {
		var notif_ba dto.Notif_ba
		if err := rows.Scan(&notif_ba.Notification_id, &notif_ba.Business_associate_id, &notif_ba.Active_ind, &notif_ba.Effective_date, &notif_ba.Expiry_date, &notif_ba.Ppdm_guid, &notif_ba.Received_date, &notif_ba.Received_ind, &notif_ba.Remark, &notif_ba.Response, &notif_ba.Served_date, &notif_ba.Served_ind, &notif_ba.Source, &notif_ba.Row_changed_by, &notif_ba.Row_changed_date, &notif_ba.Row_created_by, &notif_ba.Row_created_date, &notif_ba.Row_effective_date, &notif_ba.Row_expiry_date, &notif_ba.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append notif_ba to result
		result = append(result, notif_ba)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetNotifBa(c *fiber.Ctx) error {
	var notif_ba dto.Notif_ba

	setDefaults(&notif_ba)

	if err := json.Unmarshal(c.Body(), &notif_ba); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into notif_ba values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	notif_ba.Row_created_date = formatDate(notif_ba.Row_created_date)
	notif_ba.Row_changed_date = formatDate(notif_ba.Row_changed_date)
	notif_ba.Effective_date = formatDateString(notif_ba.Effective_date)
	notif_ba.Expiry_date = formatDateString(notif_ba.Expiry_date)
	notif_ba.Received_date = formatDateString(notif_ba.Received_date)
	notif_ba.Served_date = formatDateString(notif_ba.Served_date)
	notif_ba.Row_effective_date = formatDateString(notif_ba.Row_effective_date)
	notif_ba.Row_expiry_date = formatDateString(notif_ba.Row_expiry_date)






	rows, err := stmt.Exec(notif_ba.Notification_id, notif_ba.Business_associate_id, notif_ba.Active_ind, notif_ba.Effective_date, notif_ba.Expiry_date, notif_ba.Ppdm_guid, notif_ba.Received_date, notif_ba.Received_ind, notif_ba.Remark, notif_ba.Response, notif_ba.Served_date, notif_ba.Served_ind, notif_ba.Source, notif_ba.Row_changed_by, notif_ba.Row_changed_date, notif_ba.Row_created_by, notif_ba.Row_created_date, notif_ba.Row_effective_date, notif_ba.Row_expiry_date, notif_ba.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateNotifBa(c *fiber.Ctx) error {
	var notif_ba dto.Notif_ba

	setDefaults(&notif_ba)

	if err := json.Unmarshal(c.Body(), &notif_ba); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	notif_ba.Notification_id = id

    if notif_ba.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update notif_ba set business_associate_id = :1, active_ind = :2, effective_date = :3, expiry_date = :4, ppdm_guid = :5, received_date = :6, received_ind = :7, remark = :8, response = :9, served_date = :10, served_ind = :11, source = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where notification_id = :20")
	if err != nil {
		return err
	}

	notif_ba.Row_changed_date = formatDate(notif_ba.Row_changed_date)
	notif_ba.Effective_date = formatDateString(notif_ba.Effective_date)
	notif_ba.Expiry_date = formatDateString(notif_ba.Expiry_date)
	notif_ba.Received_date = formatDateString(notif_ba.Received_date)
	notif_ba.Served_date = formatDateString(notif_ba.Served_date)
	notif_ba.Row_effective_date = formatDateString(notif_ba.Row_effective_date)
	notif_ba.Row_expiry_date = formatDateString(notif_ba.Row_expiry_date)






	rows, err := stmt.Exec(notif_ba.Business_associate_id, notif_ba.Active_ind, notif_ba.Effective_date, notif_ba.Expiry_date, notif_ba.Ppdm_guid, notif_ba.Received_date, notif_ba.Received_ind, notif_ba.Remark, notif_ba.Response, notif_ba.Served_date, notif_ba.Served_ind, notif_ba.Source, notif_ba.Row_changed_by, notif_ba.Row_changed_date, notif_ba.Row_created_by, notif_ba.Row_effective_date, notif_ba.Row_expiry_date, notif_ba.Row_quality, notif_ba.Notification_id)
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

func PatchNotifBa(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update notif_ba set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "received_date" || key == "served_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteNotifBa(c *fiber.Ctx) error {
	id := c.Params("id")
	var notif_ba dto.Notif_ba
	notif_ba.Notification_id = id

	stmt, err := db.Prepare("delete from notif_ba where notification_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(notif_ba.Notification_id)
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



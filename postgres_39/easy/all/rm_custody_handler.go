package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmCustody(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_custody")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_custody

	for rows.Next() {
		var rm_custody dto.Rm_custody
		if err := rows.Scan(&rm_custody.Physical_item_id, &rm_custody.Custody_id, &rm_custody.Active_ind, &rm_custody.Effective_date, &rm_custody.Expiry_date, &rm_custody.Ppdm_guid, &rm_custody.Receive_by, &rm_custody.Receive_date, &rm_custody.Registration_num, &rm_custody.Remark, &rm_custody.Send_by, &rm_custody.Send_date, &rm_custody.Send_method, &rm_custody.Source, &rm_custody.Row_changed_by, &rm_custody.Row_changed_date, &rm_custody.Row_created_by, &rm_custody.Row_created_date, &rm_custody.Row_effective_date, &rm_custody.Row_expiry_date, &rm_custody.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_custody to result
		result = append(result, rm_custody)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmCustody(c *fiber.Ctx) error {
	var rm_custody dto.Rm_custody

	setDefaults(&rm_custody)

	if err := json.Unmarshal(c.Body(), &rm_custody); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_custody values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	rm_custody.Row_created_date = formatDate(rm_custody.Row_created_date)
	rm_custody.Row_changed_date = formatDate(rm_custody.Row_changed_date)
	rm_custody.Effective_date = formatDateString(rm_custody.Effective_date)
	rm_custody.Expiry_date = formatDateString(rm_custody.Expiry_date)
	rm_custody.Receive_date = formatDateString(rm_custody.Receive_date)
	rm_custody.Send_date = formatDateString(rm_custody.Send_date)
	rm_custody.Row_effective_date = formatDateString(rm_custody.Row_effective_date)
	rm_custody.Row_expiry_date = formatDateString(rm_custody.Row_expiry_date)






	rows, err := stmt.Exec(rm_custody.Physical_item_id, rm_custody.Custody_id, rm_custody.Active_ind, rm_custody.Effective_date, rm_custody.Expiry_date, rm_custody.Ppdm_guid, rm_custody.Receive_by, rm_custody.Receive_date, rm_custody.Registration_num, rm_custody.Remark, rm_custody.Send_by, rm_custody.Send_date, rm_custody.Send_method, rm_custody.Source, rm_custody.Row_changed_by, rm_custody.Row_changed_date, rm_custody.Row_created_by, rm_custody.Row_created_date, rm_custody.Row_effective_date, rm_custody.Row_expiry_date, rm_custody.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmCustody(c *fiber.Ctx) error {
	var rm_custody dto.Rm_custody

	setDefaults(&rm_custody)

	if err := json.Unmarshal(c.Body(), &rm_custody); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_custody.Physical_item_id = id

    if rm_custody.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_custody set custody_id = :1, active_ind = :2, effective_date = :3, expiry_date = :4, ppdm_guid = :5, receive_by = :6, receive_date = :7, registration_num = :8, remark = :9, send_by = :10, send_date = :11, send_method = :12, source = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where physical_item_id = :21")
	if err != nil {
		return err
	}

	rm_custody.Row_changed_date = formatDate(rm_custody.Row_changed_date)
	rm_custody.Effective_date = formatDateString(rm_custody.Effective_date)
	rm_custody.Expiry_date = formatDateString(rm_custody.Expiry_date)
	rm_custody.Receive_date = formatDateString(rm_custody.Receive_date)
	rm_custody.Send_date = formatDateString(rm_custody.Send_date)
	rm_custody.Row_effective_date = formatDateString(rm_custody.Row_effective_date)
	rm_custody.Row_expiry_date = formatDateString(rm_custody.Row_expiry_date)






	rows, err := stmt.Exec(rm_custody.Custody_id, rm_custody.Active_ind, rm_custody.Effective_date, rm_custody.Expiry_date, rm_custody.Ppdm_guid, rm_custody.Receive_by, rm_custody.Receive_date, rm_custody.Registration_num, rm_custody.Remark, rm_custody.Send_by, rm_custody.Send_date, rm_custody.Send_method, rm_custody.Source, rm_custody.Row_changed_by, rm_custody.Row_changed_date, rm_custody.Row_created_by, rm_custody.Row_effective_date, rm_custody.Row_expiry_date, rm_custody.Row_quality, rm_custody.Physical_item_id)
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

func PatchRmCustody(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_custody set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "receive_date" || key == "send_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where physical_item_id = :id"

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

func DeleteRmCustody(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_custody dto.Rm_custody
	rm_custody.Physical_item_id = id

	stmt, err := db.Prepare("delete from rm_custody where physical_item_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_custody.Physical_item_id)
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



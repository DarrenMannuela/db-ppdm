package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmInfoItemStatus(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_info_item_status")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_info_item_status

	for rows.Next() {
		var rm_info_item_status dto.Rm_info_item_status
		if err := rows.Scan(&rm_info_item_status.Info_item_subtype, &rm_info_item_status.Information_item_id, &rm_info_item_status.Status_id, &rm_info_item_status.Active_ind, &rm_info_item_status.Effective_date, &rm_info_item_status.Expiry_date, &rm_info_item_status.Ppdm_guid, &rm_info_item_status.Remark, &rm_info_item_status.Source, &rm_info_item_status.Status, &rm_info_item_status.Status_date, &rm_info_item_status.Status_type, &rm_info_item_status.Row_changed_by, &rm_info_item_status.Row_changed_date, &rm_info_item_status.Row_created_by, &rm_info_item_status.Row_created_date, &rm_info_item_status.Row_effective_date, &rm_info_item_status.Row_expiry_date, &rm_info_item_status.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_info_item_status to result
		result = append(result, rm_info_item_status)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmInfoItemStatus(c *fiber.Ctx) error {
	var rm_info_item_status dto.Rm_info_item_status

	setDefaults(&rm_info_item_status)

	if err := json.Unmarshal(c.Body(), &rm_info_item_status); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_info_item_status values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	rm_info_item_status.Row_created_date = formatDate(rm_info_item_status.Row_created_date)
	rm_info_item_status.Row_changed_date = formatDate(rm_info_item_status.Row_changed_date)
	rm_info_item_status.Effective_date = formatDateString(rm_info_item_status.Effective_date)
	rm_info_item_status.Expiry_date = formatDateString(rm_info_item_status.Expiry_date)
	rm_info_item_status.Status_date = formatDateString(rm_info_item_status.Status_date)
	rm_info_item_status.Row_effective_date = formatDateString(rm_info_item_status.Row_effective_date)
	rm_info_item_status.Row_expiry_date = formatDateString(rm_info_item_status.Row_expiry_date)






	rows, err := stmt.Exec(rm_info_item_status.Info_item_subtype, rm_info_item_status.Information_item_id, rm_info_item_status.Status_id, rm_info_item_status.Active_ind, rm_info_item_status.Effective_date, rm_info_item_status.Expiry_date, rm_info_item_status.Ppdm_guid, rm_info_item_status.Remark, rm_info_item_status.Source, rm_info_item_status.Status, rm_info_item_status.Status_date, rm_info_item_status.Status_type, rm_info_item_status.Row_changed_by, rm_info_item_status.Row_changed_date, rm_info_item_status.Row_created_by, rm_info_item_status.Row_created_date, rm_info_item_status.Row_effective_date, rm_info_item_status.Row_expiry_date, rm_info_item_status.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmInfoItemStatus(c *fiber.Ctx) error {
	var rm_info_item_status dto.Rm_info_item_status

	setDefaults(&rm_info_item_status)

	if err := json.Unmarshal(c.Body(), &rm_info_item_status); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_info_item_status.Info_item_subtype = id

    if rm_info_item_status.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_info_item_status set information_item_id = :1, status_id = :2, active_ind = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, source = :8, status = :9, status_date = :10, status_type = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where info_item_subtype = :19")
	if err != nil {
		return err
	}

	rm_info_item_status.Row_changed_date = formatDate(rm_info_item_status.Row_changed_date)
	rm_info_item_status.Effective_date = formatDateString(rm_info_item_status.Effective_date)
	rm_info_item_status.Expiry_date = formatDateString(rm_info_item_status.Expiry_date)
	rm_info_item_status.Status_date = formatDateString(rm_info_item_status.Status_date)
	rm_info_item_status.Row_effective_date = formatDateString(rm_info_item_status.Row_effective_date)
	rm_info_item_status.Row_expiry_date = formatDateString(rm_info_item_status.Row_expiry_date)






	rows, err := stmt.Exec(rm_info_item_status.Information_item_id, rm_info_item_status.Status_id, rm_info_item_status.Active_ind, rm_info_item_status.Effective_date, rm_info_item_status.Expiry_date, rm_info_item_status.Ppdm_guid, rm_info_item_status.Remark, rm_info_item_status.Source, rm_info_item_status.Status, rm_info_item_status.Status_date, rm_info_item_status.Status_type, rm_info_item_status.Row_changed_by, rm_info_item_status.Row_changed_date, rm_info_item_status.Row_created_by, rm_info_item_status.Row_effective_date, rm_info_item_status.Row_expiry_date, rm_info_item_status.Row_quality, rm_info_item_status.Info_item_subtype)
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

func PatchRmInfoItemStatus(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_info_item_status set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "status_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where info_item_subtype = :id"

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

func DeleteRmInfoItemStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_info_item_status dto.Rm_info_item_status
	rm_info_item_status.Info_item_subtype = id

	stmt, err := db.Prepare("delete from rm_info_item_status where info_item_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_info_item_status.Info_item_subtype)
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



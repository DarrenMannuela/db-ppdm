package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmInfoItemMaint(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_info_item_maint")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_info_item_maint

	for rows.Next() {
		var rm_info_item_maint dto.Rm_info_item_maint
		if err := rows.Scan(&rm_info_item_maint.Info_item_subtype, &rm_info_item_maint.Information_item_id, &rm_info_item_maint.Maintain_id, &rm_info_item_maint.Active_ind, &rm_info_item_maint.Effective_date, &rm_info_item_maint.Expiry_date, &rm_info_item_maint.Instruction, &rm_info_item_maint.Maint_ba_id, &rm_info_item_maint.Maint_complete_date, &rm_info_item_maint.Maint_due_date, &rm_info_item_maint.Maint_period, &rm_info_item_maint.Maint_period_type, &rm_info_item_maint.Ppdm_guid, &rm_info_item_maint.Remark, &rm_info_item_maint.Scheduled_ind, &rm_info_item_maint.Source, &rm_info_item_maint.Row_changed_by, &rm_info_item_maint.Row_changed_date, &rm_info_item_maint.Row_created_by, &rm_info_item_maint.Row_created_date, &rm_info_item_maint.Row_effective_date, &rm_info_item_maint.Row_expiry_date, &rm_info_item_maint.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_info_item_maint to result
		result = append(result, rm_info_item_maint)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmInfoItemMaint(c *fiber.Ctx) error {
	var rm_info_item_maint dto.Rm_info_item_maint

	setDefaults(&rm_info_item_maint)

	if err := json.Unmarshal(c.Body(), &rm_info_item_maint); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_info_item_maint values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	rm_info_item_maint.Row_created_date = formatDate(rm_info_item_maint.Row_created_date)
	rm_info_item_maint.Row_changed_date = formatDate(rm_info_item_maint.Row_changed_date)
	rm_info_item_maint.Effective_date = formatDateString(rm_info_item_maint.Effective_date)
	rm_info_item_maint.Expiry_date = formatDateString(rm_info_item_maint.Expiry_date)
	rm_info_item_maint.Maint_complete_date = formatDateString(rm_info_item_maint.Maint_complete_date)
	rm_info_item_maint.Maint_due_date = formatDateString(rm_info_item_maint.Maint_due_date)
	rm_info_item_maint.Row_effective_date = formatDateString(rm_info_item_maint.Row_effective_date)
	rm_info_item_maint.Row_expiry_date = formatDateString(rm_info_item_maint.Row_expiry_date)






	rows, err := stmt.Exec(rm_info_item_maint.Info_item_subtype, rm_info_item_maint.Information_item_id, rm_info_item_maint.Maintain_id, rm_info_item_maint.Active_ind, rm_info_item_maint.Effective_date, rm_info_item_maint.Expiry_date, rm_info_item_maint.Instruction, rm_info_item_maint.Maint_ba_id, rm_info_item_maint.Maint_complete_date, rm_info_item_maint.Maint_due_date, rm_info_item_maint.Maint_period, rm_info_item_maint.Maint_period_type, rm_info_item_maint.Ppdm_guid, rm_info_item_maint.Remark, rm_info_item_maint.Scheduled_ind, rm_info_item_maint.Source, rm_info_item_maint.Row_changed_by, rm_info_item_maint.Row_changed_date, rm_info_item_maint.Row_created_by, rm_info_item_maint.Row_created_date, rm_info_item_maint.Row_effective_date, rm_info_item_maint.Row_expiry_date, rm_info_item_maint.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmInfoItemMaint(c *fiber.Ctx) error {
	var rm_info_item_maint dto.Rm_info_item_maint

	setDefaults(&rm_info_item_maint)

	if err := json.Unmarshal(c.Body(), &rm_info_item_maint); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_info_item_maint.Info_item_subtype = id

    if rm_info_item_maint.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_info_item_maint set information_item_id = :1, maintain_id = :2, active_ind = :3, effective_date = :4, expiry_date = :5, instruction = :6, maint_ba_id = :7, maint_complete_date = :8, maint_due_date = :9, maint_period = :10, maint_period_type = :11, ppdm_guid = :12, remark = :13, scheduled_ind = :14, source = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where info_item_subtype = :23")
	if err != nil {
		return err
	}

	rm_info_item_maint.Row_changed_date = formatDate(rm_info_item_maint.Row_changed_date)
	rm_info_item_maint.Effective_date = formatDateString(rm_info_item_maint.Effective_date)
	rm_info_item_maint.Expiry_date = formatDateString(rm_info_item_maint.Expiry_date)
	rm_info_item_maint.Maint_complete_date = formatDateString(rm_info_item_maint.Maint_complete_date)
	rm_info_item_maint.Maint_due_date = formatDateString(rm_info_item_maint.Maint_due_date)
	rm_info_item_maint.Row_effective_date = formatDateString(rm_info_item_maint.Row_effective_date)
	rm_info_item_maint.Row_expiry_date = formatDateString(rm_info_item_maint.Row_expiry_date)






	rows, err := stmt.Exec(rm_info_item_maint.Information_item_id, rm_info_item_maint.Maintain_id, rm_info_item_maint.Active_ind, rm_info_item_maint.Effective_date, rm_info_item_maint.Expiry_date, rm_info_item_maint.Instruction, rm_info_item_maint.Maint_ba_id, rm_info_item_maint.Maint_complete_date, rm_info_item_maint.Maint_due_date, rm_info_item_maint.Maint_period, rm_info_item_maint.Maint_period_type, rm_info_item_maint.Ppdm_guid, rm_info_item_maint.Remark, rm_info_item_maint.Scheduled_ind, rm_info_item_maint.Source, rm_info_item_maint.Row_changed_by, rm_info_item_maint.Row_changed_date, rm_info_item_maint.Row_created_by, rm_info_item_maint.Row_effective_date, rm_info_item_maint.Row_expiry_date, rm_info_item_maint.Row_quality, rm_info_item_maint.Info_item_subtype)
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

func PatchRmInfoItemMaint(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_info_item_maint set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "maint_complete_date" || key == "maint_due_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteRmInfoItemMaint(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_info_item_maint dto.Rm_info_item_maint
	rm_info_item_maint.Info_item_subtype = id

	stmt, err := db.Prepare("delete from rm_info_item_maint where info_item_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_info_item_maint.Info_item_subtype)
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



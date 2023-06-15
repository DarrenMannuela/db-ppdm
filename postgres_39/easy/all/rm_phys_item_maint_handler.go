package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmPhysItemMaint(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_phys_item_maint")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_phys_item_maint

	for rows.Next() {
		var rm_phys_item_maint dto.Rm_phys_item_maint
		if err := rows.Scan(&rm_phys_item_maint.Physical_item_id, &rm_phys_item_maint.Maint_id, &rm_phys_item_maint.Active_ind, &rm_phys_item_maint.Actual_date, &rm_phys_item_maint.Effective_date, &rm_phys_item_maint.Expiry_date, &rm_phys_item_maint.Maintain_process, &rm_phys_item_maint.Maint_ba_id, &rm_phys_item_maint.Maint_by_name, &rm_phys_item_maint.Ppdm_guid, &rm_phys_item_maint.Remark, &rm_phys_item_maint.Scheduled_date, &rm_phys_item_maint.Source, &rm_phys_item_maint.Row_changed_by, &rm_phys_item_maint.Row_changed_date, &rm_phys_item_maint.Row_created_by, &rm_phys_item_maint.Row_created_date, &rm_phys_item_maint.Row_effective_date, &rm_phys_item_maint.Row_expiry_date, &rm_phys_item_maint.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_phys_item_maint to result
		result = append(result, rm_phys_item_maint)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmPhysItemMaint(c *fiber.Ctx) error {
	var rm_phys_item_maint dto.Rm_phys_item_maint

	setDefaults(&rm_phys_item_maint)

	if err := json.Unmarshal(c.Body(), &rm_phys_item_maint); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_phys_item_maint values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	rm_phys_item_maint.Row_created_date = formatDate(rm_phys_item_maint.Row_created_date)
	rm_phys_item_maint.Row_changed_date = formatDate(rm_phys_item_maint.Row_changed_date)
	rm_phys_item_maint.Actual_date = formatDateString(rm_phys_item_maint.Actual_date)
	rm_phys_item_maint.Effective_date = formatDateString(rm_phys_item_maint.Effective_date)
	rm_phys_item_maint.Expiry_date = formatDateString(rm_phys_item_maint.Expiry_date)
	rm_phys_item_maint.Scheduled_date = formatDateString(rm_phys_item_maint.Scheduled_date)
	rm_phys_item_maint.Row_effective_date = formatDateString(rm_phys_item_maint.Row_effective_date)
	rm_phys_item_maint.Row_expiry_date = formatDateString(rm_phys_item_maint.Row_expiry_date)






	rows, err := stmt.Exec(rm_phys_item_maint.Physical_item_id, rm_phys_item_maint.Maint_id, rm_phys_item_maint.Active_ind, rm_phys_item_maint.Actual_date, rm_phys_item_maint.Effective_date, rm_phys_item_maint.Expiry_date, rm_phys_item_maint.Maintain_process, rm_phys_item_maint.Maint_ba_id, rm_phys_item_maint.Maint_by_name, rm_phys_item_maint.Ppdm_guid, rm_phys_item_maint.Remark, rm_phys_item_maint.Scheduled_date, rm_phys_item_maint.Source, rm_phys_item_maint.Row_changed_by, rm_phys_item_maint.Row_changed_date, rm_phys_item_maint.Row_created_by, rm_phys_item_maint.Row_created_date, rm_phys_item_maint.Row_effective_date, rm_phys_item_maint.Row_expiry_date, rm_phys_item_maint.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmPhysItemMaint(c *fiber.Ctx) error {
	var rm_phys_item_maint dto.Rm_phys_item_maint

	setDefaults(&rm_phys_item_maint)

	if err := json.Unmarshal(c.Body(), &rm_phys_item_maint); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_phys_item_maint.Physical_item_id = id

    if rm_phys_item_maint.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_phys_item_maint set maint_id = :1, active_ind = :2, actual_date = :3, effective_date = :4, expiry_date = :5, maintain_process = :6, maint_ba_id = :7, maint_by_name = :8, ppdm_guid = :9, remark = :10, scheduled_date = :11, source = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where physical_item_id = :20")
	if err != nil {
		return err
	}

	rm_phys_item_maint.Row_changed_date = formatDate(rm_phys_item_maint.Row_changed_date)
	rm_phys_item_maint.Actual_date = formatDateString(rm_phys_item_maint.Actual_date)
	rm_phys_item_maint.Effective_date = formatDateString(rm_phys_item_maint.Effective_date)
	rm_phys_item_maint.Expiry_date = formatDateString(rm_phys_item_maint.Expiry_date)
	rm_phys_item_maint.Scheduled_date = formatDateString(rm_phys_item_maint.Scheduled_date)
	rm_phys_item_maint.Row_effective_date = formatDateString(rm_phys_item_maint.Row_effective_date)
	rm_phys_item_maint.Row_expiry_date = formatDateString(rm_phys_item_maint.Row_expiry_date)






	rows, err := stmt.Exec(rm_phys_item_maint.Maint_id, rm_phys_item_maint.Active_ind, rm_phys_item_maint.Actual_date, rm_phys_item_maint.Effective_date, rm_phys_item_maint.Expiry_date, rm_phys_item_maint.Maintain_process, rm_phys_item_maint.Maint_ba_id, rm_phys_item_maint.Maint_by_name, rm_phys_item_maint.Ppdm_guid, rm_phys_item_maint.Remark, rm_phys_item_maint.Scheduled_date, rm_phys_item_maint.Source, rm_phys_item_maint.Row_changed_by, rm_phys_item_maint.Row_changed_date, rm_phys_item_maint.Row_created_by, rm_phys_item_maint.Row_effective_date, rm_phys_item_maint.Row_expiry_date, rm_phys_item_maint.Row_quality, rm_phys_item_maint.Physical_item_id)
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

func PatchRmPhysItemMaint(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_phys_item_maint set "
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
		} else if key == "actual_date" || key == "effective_date" || key == "expiry_date" || key == "scheduled_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteRmPhysItemMaint(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_phys_item_maint dto.Rm_phys_item_maint
	rm_phys_item_maint.Physical_item_id = id

	stmt, err := db.Prepare("delete from rm_phys_item_maint where physical_item_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_phys_item_maint.Physical_item_id)
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



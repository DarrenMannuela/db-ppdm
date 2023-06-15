package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetEquipmentMaintStatus(c *fiber.Ctx) error {
	rows, err := db.Query("select * from equipment_maint_status")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Equipment_maint_status

	for rows.Next() {
		var equipment_maint_status dto.Equipment_maint_status
		if err := rows.Scan(&equipment_maint_status.Equipment_id, &equipment_maint_status.Equip_maint_id, &equipment_maint_status.Status_id, &equipment_maint_status.Active_ind, &equipment_maint_status.Effective_date, &equipment_maint_status.Expiry_date, &equipment_maint_status.Maintain_status, &equipment_maint_status.Maintain_status_type, &equipment_maint_status.Ppdm_guid, &equipment_maint_status.Remark, &equipment_maint_status.Source, &equipment_maint_status.Row_changed_by, &equipment_maint_status.Row_changed_date, &equipment_maint_status.Row_created_by, &equipment_maint_status.Row_created_date, &equipment_maint_status.Row_effective_date, &equipment_maint_status.Row_expiry_date, &equipment_maint_status.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append equipment_maint_status to result
		result = append(result, equipment_maint_status)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetEquipmentMaintStatus(c *fiber.Ctx) error {
	var equipment_maint_status dto.Equipment_maint_status

	setDefaults(&equipment_maint_status)

	if err := json.Unmarshal(c.Body(), &equipment_maint_status); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into equipment_maint_status values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	equipment_maint_status.Row_created_date = formatDate(equipment_maint_status.Row_created_date)
	equipment_maint_status.Row_changed_date = formatDate(equipment_maint_status.Row_changed_date)
	equipment_maint_status.Effective_date = formatDateString(equipment_maint_status.Effective_date)
	equipment_maint_status.Expiry_date = formatDateString(equipment_maint_status.Expiry_date)
	equipment_maint_status.Row_effective_date = formatDateString(equipment_maint_status.Row_effective_date)
	equipment_maint_status.Row_expiry_date = formatDateString(equipment_maint_status.Row_expiry_date)






	rows, err := stmt.Exec(equipment_maint_status.Equipment_id, equipment_maint_status.Equip_maint_id, equipment_maint_status.Status_id, equipment_maint_status.Active_ind, equipment_maint_status.Effective_date, equipment_maint_status.Expiry_date, equipment_maint_status.Maintain_status, equipment_maint_status.Maintain_status_type, equipment_maint_status.Ppdm_guid, equipment_maint_status.Remark, equipment_maint_status.Source, equipment_maint_status.Row_changed_by, equipment_maint_status.Row_changed_date, equipment_maint_status.Row_created_by, equipment_maint_status.Row_created_date, equipment_maint_status.Row_effective_date, equipment_maint_status.Row_expiry_date, equipment_maint_status.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateEquipmentMaintStatus(c *fiber.Ctx) error {
	var equipment_maint_status dto.Equipment_maint_status

	setDefaults(&equipment_maint_status)

	if err := json.Unmarshal(c.Body(), &equipment_maint_status); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	equipment_maint_status.Equipment_id = id

    if equipment_maint_status.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update equipment_maint_status set equip_maint_id = :1, status_id = :2, active_ind = :3, effective_date = :4, expiry_date = :5, maintain_status = :6, maintain_status_type = :7, ppdm_guid = :8, remark = :9, source = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where equipment_id = :18")
	if err != nil {
		return err
	}

	equipment_maint_status.Row_changed_date = formatDate(equipment_maint_status.Row_changed_date)
	equipment_maint_status.Effective_date = formatDateString(equipment_maint_status.Effective_date)
	equipment_maint_status.Expiry_date = formatDateString(equipment_maint_status.Expiry_date)
	equipment_maint_status.Row_effective_date = formatDateString(equipment_maint_status.Row_effective_date)
	equipment_maint_status.Row_expiry_date = formatDateString(equipment_maint_status.Row_expiry_date)






	rows, err := stmt.Exec(equipment_maint_status.Equip_maint_id, equipment_maint_status.Status_id, equipment_maint_status.Active_ind, equipment_maint_status.Effective_date, equipment_maint_status.Expiry_date, equipment_maint_status.Maintain_status, equipment_maint_status.Maintain_status_type, equipment_maint_status.Ppdm_guid, equipment_maint_status.Remark, equipment_maint_status.Source, equipment_maint_status.Row_changed_by, equipment_maint_status.Row_changed_date, equipment_maint_status.Row_created_by, equipment_maint_status.Row_effective_date, equipment_maint_status.Row_expiry_date, equipment_maint_status.Row_quality, equipment_maint_status.Equipment_id)
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

func PatchEquipmentMaintStatus(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update equipment_maint_status set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where equipment_id = :id"

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

func DeleteEquipmentMaintStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var equipment_maint_status dto.Equipment_maint_status
	equipment_maint_status.Equipment_id = id

	stmt, err := db.Prepare("delete from equipment_maint_status where equipment_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(equipment_maint_status.Equipment_id)
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



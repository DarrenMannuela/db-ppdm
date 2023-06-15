package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetEquipmentStatus(c *fiber.Ctx) error {
	rows, err := db.Query("select * from equipment_status")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Equipment_status

	for rows.Next() {
		var equipment_status dto.Equipment_status
		if err := rows.Scan(&equipment_status.Equipment_id, &equipment_status.Equip_status_id, &equipment_status.Active_ind, &equipment_status.Effective_date, &equipment_status.End_time, &equipment_status.Equip_status, &equipment_status.Equip_status_type, &equipment_status.Expiry_date, &equipment_status.Percent_capability, &equipment_status.Ppdm_guid, &equipment_status.Preferred_ind, &equipment_status.Remark, &equipment_status.Source, &equipment_status.Start_time, &equipment_status.Timezone, &equipment_status.Row_changed_by, &equipment_status.Row_changed_date, &equipment_status.Row_created_by, &equipment_status.Row_created_date, &equipment_status.Row_effective_date, &equipment_status.Row_expiry_date, &equipment_status.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append equipment_status to result
		result = append(result, equipment_status)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetEquipmentStatus(c *fiber.Ctx) error {
	var equipment_status dto.Equipment_status

	setDefaults(&equipment_status)

	if err := json.Unmarshal(c.Body(), &equipment_status); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into equipment_status values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	equipment_status.Row_created_date = formatDate(equipment_status.Row_created_date)
	equipment_status.Row_changed_date = formatDate(equipment_status.Row_changed_date)
	equipment_status.Effective_date = formatDateString(equipment_status.Effective_date)
	equipment_status.End_time = formatDateString(equipment_status.End_time)
	equipment_status.Expiry_date = formatDateString(equipment_status.Expiry_date)
	equipment_status.Start_time = formatDateString(equipment_status.Start_time)
	equipment_status.Row_effective_date = formatDateString(equipment_status.Row_effective_date)
	equipment_status.Row_expiry_date = formatDateString(equipment_status.Row_expiry_date)






	rows, err := stmt.Exec(equipment_status.Equipment_id, equipment_status.Equip_status_id, equipment_status.Active_ind, equipment_status.Effective_date, equipment_status.End_time, equipment_status.Equip_status, equipment_status.Equip_status_type, equipment_status.Expiry_date, equipment_status.Percent_capability, equipment_status.Ppdm_guid, equipment_status.Preferred_ind, equipment_status.Remark, equipment_status.Source, equipment_status.Start_time, equipment_status.Timezone, equipment_status.Row_changed_by, equipment_status.Row_changed_date, equipment_status.Row_created_by, equipment_status.Row_created_date, equipment_status.Row_effective_date, equipment_status.Row_expiry_date, equipment_status.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateEquipmentStatus(c *fiber.Ctx) error {
	var equipment_status dto.Equipment_status

	setDefaults(&equipment_status)

	if err := json.Unmarshal(c.Body(), &equipment_status); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	equipment_status.Equipment_id = id

    if equipment_status.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update equipment_status set equip_status_id = :1, active_ind = :2, effective_date = :3, end_time = :4, equip_status = :5, equip_status_type = :6, expiry_date = :7, percent_capability = :8, ppdm_guid = :9, preferred_ind = :10, remark = :11, source = :12, start_time = :13, timezone = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where equipment_id = :22")
	if err != nil {
		return err
	}

	equipment_status.Row_changed_date = formatDate(equipment_status.Row_changed_date)
	equipment_status.Effective_date = formatDateString(equipment_status.Effective_date)
	equipment_status.End_time = formatDateString(equipment_status.End_time)
	equipment_status.Expiry_date = formatDateString(equipment_status.Expiry_date)
	equipment_status.Start_time = formatDateString(equipment_status.Start_time)
	equipment_status.Row_effective_date = formatDateString(equipment_status.Row_effective_date)
	equipment_status.Row_expiry_date = formatDateString(equipment_status.Row_expiry_date)






	rows, err := stmt.Exec(equipment_status.Equip_status_id, equipment_status.Active_ind, equipment_status.Effective_date, equipment_status.End_time, equipment_status.Equip_status, equipment_status.Equip_status_type, equipment_status.Expiry_date, equipment_status.Percent_capability, equipment_status.Ppdm_guid, equipment_status.Preferred_ind, equipment_status.Remark, equipment_status.Source, equipment_status.Start_time, equipment_status.Timezone, equipment_status.Row_changed_by, equipment_status.Row_changed_date, equipment_status.Row_created_by, equipment_status.Row_effective_date, equipment_status.Row_expiry_date, equipment_status.Row_quality, equipment_status.Equipment_id)
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

func PatchEquipmentStatus(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update equipment_status set "
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
		} else if key == "effective_date" || key == "end_time" || key == "expiry_date" || key == "start_time" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteEquipmentStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var equipment_status dto.Equipment_status
	equipment_status.Equipment_id = id

	stmt, err := db.Prepare("delete from equipment_status where equipment_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(equipment_status.Equipment_id)
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



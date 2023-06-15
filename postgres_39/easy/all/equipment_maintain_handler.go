package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetEquipmentMaintain(c *fiber.Ctx) error {
	rows, err := db.Query("select * from equipment_maintain")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Equipment_maintain

	for rows.Next() {
		var equipment_maintain dto.Equipment_maintain
		if err := rows.Scan(&equipment_maintain.Equipment_id, &equipment_maintain.Equip_maint_id, &equipment_maintain.Active_ind, &equipment_maintain.Actual_end_date, &equipment_maintain.Actual_start_date, &equipment_maintain.Catalogue_equip_id, &equipment_maintain.Completed_by_ba_id, &equipment_maintain.Effective_date, &equipment_maintain.End_date, &equipment_maintain.Expiry_date, &equipment_maintain.Failure_ind, &equipment_maintain.Location_ba_address_obs_no, &equipment_maintain.Location_ba_id, &equipment_maintain.Location_ba_source, &equipment_maintain.Maint_location_type, &equipment_maintain.Maint_reason, &equipment_maintain.Maint_type, &equipment_maintain.Ppdm_guid, &equipment_maintain.Project_id, &equipment_maintain.Remark, &equipment_maintain.Scheduled_date, &equipment_maintain.Scheduled_ind, &equipment_maintain.Source, &equipment_maintain.Start_date, &equipment_maintain.System_condition, &equipment_maintain.Row_changed_by, &equipment_maintain.Row_changed_date, &equipment_maintain.Row_created_by, &equipment_maintain.Row_created_date, &equipment_maintain.Row_effective_date, &equipment_maintain.Row_expiry_date, &equipment_maintain.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append equipment_maintain to result
		result = append(result, equipment_maintain)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetEquipmentMaintain(c *fiber.Ctx) error {
	var equipment_maintain dto.Equipment_maintain

	setDefaults(&equipment_maintain)

	if err := json.Unmarshal(c.Body(), &equipment_maintain); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into equipment_maintain values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32)")
	if err != nil {
		return err
	}
	equipment_maintain.Row_created_date = formatDate(equipment_maintain.Row_created_date)
	equipment_maintain.Row_changed_date = formatDate(equipment_maintain.Row_changed_date)
	equipment_maintain.Actual_end_date = formatDateString(equipment_maintain.Actual_end_date)
	equipment_maintain.Actual_start_date = formatDateString(equipment_maintain.Actual_start_date)
	equipment_maintain.Effective_date = formatDateString(equipment_maintain.Effective_date)
	equipment_maintain.End_date = formatDateString(equipment_maintain.End_date)
	equipment_maintain.Expiry_date = formatDateString(equipment_maintain.Expiry_date)
	equipment_maintain.Scheduled_date = formatDateString(equipment_maintain.Scheduled_date)
	equipment_maintain.Start_date = formatDateString(equipment_maintain.Start_date)
	equipment_maintain.Row_effective_date = formatDateString(equipment_maintain.Row_effective_date)
	equipment_maintain.Row_expiry_date = formatDateString(equipment_maintain.Row_expiry_date)






	rows, err := stmt.Exec(equipment_maintain.Equipment_id, equipment_maintain.Equip_maint_id, equipment_maintain.Active_ind, equipment_maintain.Actual_end_date, equipment_maintain.Actual_start_date, equipment_maintain.Catalogue_equip_id, equipment_maintain.Completed_by_ba_id, equipment_maintain.Effective_date, equipment_maintain.End_date, equipment_maintain.Expiry_date, equipment_maintain.Failure_ind, equipment_maintain.Location_ba_address_obs_no, equipment_maintain.Location_ba_id, equipment_maintain.Location_ba_source, equipment_maintain.Maint_location_type, equipment_maintain.Maint_reason, equipment_maintain.Maint_type, equipment_maintain.Ppdm_guid, equipment_maintain.Project_id, equipment_maintain.Remark, equipment_maintain.Scheduled_date, equipment_maintain.Scheduled_ind, equipment_maintain.Source, equipment_maintain.Start_date, equipment_maintain.System_condition, equipment_maintain.Row_changed_by, equipment_maintain.Row_changed_date, equipment_maintain.Row_created_by, equipment_maintain.Row_created_date, equipment_maintain.Row_effective_date, equipment_maintain.Row_expiry_date, equipment_maintain.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateEquipmentMaintain(c *fiber.Ctx) error {
	var equipment_maintain dto.Equipment_maintain

	setDefaults(&equipment_maintain)

	if err := json.Unmarshal(c.Body(), &equipment_maintain); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	equipment_maintain.Equipment_id = id

    if equipment_maintain.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update equipment_maintain set equip_maint_id = :1, active_ind = :2, actual_end_date = :3, actual_start_date = :4, catalogue_equip_id = :5, completed_by_ba_id = :6, effective_date = :7, end_date = :8, expiry_date = :9, failure_ind = :10, location_ba_address_obs_no = :11, location_ba_id = :12, location_ba_source = :13, maint_location_type = :14, maint_reason = :15, maint_type = :16, ppdm_guid = :17, project_id = :18, remark = :19, scheduled_date = :20, scheduled_ind = :21, source = :22, start_date = :23, system_condition = :24, row_changed_by = :25, row_changed_date = :26, row_created_by = :27, row_effective_date = :28, row_expiry_date = :29, row_quality = :30 where equipment_id = :32")
	if err != nil {
		return err
	}

	equipment_maintain.Row_changed_date = formatDate(equipment_maintain.Row_changed_date)
	equipment_maintain.Actual_end_date = formatDateString(equipment_maintain.Actual_end_date)
	equipment_maintain.Actual_start_date = formatDateString(equipment_maintain.Actual_start_date)
	equipment_maintain.Effective_date = formatDateString(equipment_maintain.Effective_date)
	equipment_maintain.End_date = formatDateString(equipment_maintain.End_date)
	equipment_maintain.Expiry_date = formatDateString(equipment_maintain.Expiry_date)
	equipment_maintain.Scheduled_date = formatDateString(equipment_maintain.Scheduled_date)
	equipment_maintain.Start_date = formatDateString(equipment_maintain.Start_date)
	equipment_maintain.Row_effective_date = formatDateString(equipment_maintain.Row_effective_date)
	equipment_maintain.Row_expiry_date = formatDateString(equipment_maintain.Row_expiry_date)






	rows, err := stmt.Exec(equipment_maintain.Equip_maint_id, equipment_maintain.Active_ind, equipment_maintain.Actual_end_date, equipment_maintain.Actual_start_date, equipment_maintain.Catalogue_equip_id, equipment_maintain.Completed_by_ba_id, equipment_maintain.Effective_date, equipment_maintain.End_date, equipment_maintain.Expiry_date, equipment_maintain.Failure_ind, equipment_maintain.Location_ba_address_obs_no, equipment_maintain.Location_ba_id, equipment_maintain.Location_ba_source, equipment_maintain.Maint_location_type, equipment_maintain.Maint_reason, equipment_maintain.Maint_type, equipment_maintain.Ppdm_guid, equipment_maintain.Project_id, equipment_maintain.Remark, equipment_maintain.Scheduled_date, equipment_maintain.Scheduled_ind, equipment_maintain.Source, equipment_maintain.Start_date, equipment_maintain.System_condition, equipment_maintain.Row_changed_by, equipment_maintain.Row_changed_date, equipment_maintain.Row_created_by, equipment_maintain.Row_effective_date, equipment_maintain.Row_expiry_date, equipment_maintain.Row_quality, equipment_maintain.Equipment_id)
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

func PatchEquipmentMaintain(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update equipment_maintain set "
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
		} else if key == "actual_end_date" || key == "actual_start_date" || key == "effective_date" || key == "end_date" || key == "expiry_date" || key == "scheduled_date" || key == "start_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteEquipmentMaintain(c *fiber.Ctx) error {
	id := c.Params("id")
	var equipment_maintain dto.Equipment_maintain
	equipment_maintain.Equipment_id = id

	stmt, err := db.Prepare("delete from equipment_maintain where equipment_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(equipment_maintain.Equipment_id)
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



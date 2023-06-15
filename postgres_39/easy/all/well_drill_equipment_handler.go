package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellDrillEquipment(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_drill_equipment")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_drill_equipment

	for rows.Next() {
		var well_drill_equipment dto.Well_drill_equipment
		if err := rows.Scan(&well_drill_equipment.Uwi, &well_drill_equipment.Equipment_id, &well_drill_equipment.Equipment_obs_no, &well_drill_equipment.Active_ind, &well_drill_equipment.Effective_date, &well_drill_equipment.Expiry_date, &well_drill_equipment.Offsite_date, &well_drill_equipment.Offsite_time, &well_drill_equipment.Onsite_date, &well_drill_equipment.Onsite_time, &well_drill_equipment.Operated_by_ba_id, &well_drill_equipment.Parent_equipment_id, &well_drill_equipment.Period_on_well, &well_drill_equipment.Period_on_well_ouom, &well_drill_equipment.Ppdm_guid, &well_drill_equipment.Reference_num, &well_drill_equipment.Remark, &well_drill_equipment.Rented_ind, &well_drill_equipment.Source, &well_drill_equipment.Timezone, &well_drill_equipment.Row_changed_by, &well_drill_equipment.Row_changed_date, &well_drill_equipment.Row_created_by, &well_drill_equipment.Row_created_date, &well_drill_equipment.Row_effective_date, &well_drill_equipment.Row_expiry_date, &well_drill_equipment.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_drill_equipment to result
		result = append(result, well_drill_equipment)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellDrillEquipment(c *fiber.Ctx) error {
	var well_drill_equipment dto.Well_drill_equipment

	setDefaults(&well_drill_equipment)

	if err := json.Unmarshal(c.Body(), &well_drill_equipment); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_drill_equipment values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27)")
	if err != nil {
		return err
	}
	well_drill_equipment.Row_created_date = formatDate(well_drill_equipment.Row_created_date)
	well_drill_equipment.Row_changed_date = formatDate(well_drill_equipment.Row_changed_date)
	well_drill_equipment.Effective_date = formatDateString(well_drill_equipment.Effective_date)
	well_drill_equipment.Expiry_date = formatDateString(well_drill_equipment.Expiry_date)
	well_drill_equipment.Offsite_date = formatDateString(well_drill_equipment.Offsite_date)
	well_drill_equipment.Offsite_time = formatDateString(well_drill_equipment.Offsite_time)
	well_drill_equipment.Onsite_date = formatDateString(well_drill_equipment.Onsite_date)
	well_drill_equipment.Onsite_time = formatDateString(well_drill_equipment.Onsite_time)
	well_drill_equipment.Row_effective_date = formatDateString(well_drill_equipment.Row_effective_date)
	well_drill_equipment.Row_expiry_date = formatDateString(well_drill_equipment.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_equipment.Uwi, well_drill_equipment.Equipment_id, well_drill_equipment.Equipment_obs_no, well_drill_equipment.Active_ind, well_drill_equipment.Effective_date, well_drill_equipment.Expiry_date, well_drill_equipment.Offsite_date, well_drill_equipment.Offsite_time, well_drill_equipment.Onsite_date, well_drill_equipment.Onsite_time, well_drill_equipment.Operated_by_ba_id, well_drill_equipment.Parent_equipment_id, well_drill_equipment.Period_on_well, well_drill_equipment.Period_on_well_ouom, well_drill_equipment.Ppdm_guid, well_drill_equipment.Reference_num, well_drill_equipment.Remark, well_drill_equipment.Rented_ind, well_drill_equipment.Source, well_drill_equipment.Timezone, well_drill_equipment.Row_changed_by, well_drill_equipment.Row_changed_date, well_drill_equipment.Row_created_by, well_drill_equipment.Row_created_date, well_drill_equipment.Row_effective_date, well_drill_equipment.Row_expiry_date, well_drill_equipment.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellDrillEquipment(c *fiber.Ctx) error {
	var well_drill_equipment dto.Well_drill_equipment

	setDefaults(&well_drill_equipment)

	if err := json.Unmarshal(c.Body(), &well_drill_equipment); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_drill_equipment.Uwi = id

    if well_drill_equipment.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_drill_equipment set equipment_id = :1, equipment_obs_no = :2, active_ind = :3, effective_date = :4, expiry_date = :5, offsite_date = :6, offsite_time = :7, onsite_date = :8, onsite_time = :9, operated_by_ba_id = :10, parent_equipment_id = :11, period_on_well = :12, period_on_well_ouom = :13, ppdm_guid = :14, reference_num = :15, remark = :16, rented_ind = :17, source = :18, timezone = :19, row_changed_by = :20, row_changed_date = :21, row_created_by = :22, row_effective_date = :23, row_expiry_date = :24, row_quality = :25 where uwi = :27")
	if err != nil {
		return err
	}

	well_drill_equipment.Row_changed_date = formatDate(well_drill_equipment.Row_changed_date)
	well_drill_equipment.Effective_date = formatDateString(well_drill_equipment.Effective_date)
	well_drill_equipment.Expiry_date = formatDateString(well_drill_equipment.Expiry_date)
	well_drill_equipment.Offsite_date = formatDateString(well_drill_equipment.Offsite_date)
	well_drill_equipment.Offsite_time = formatDateString(well_drill_equipment.Offsite_time)
	well_drill_equipment.Onsite_date = formatDateString(well_drill_equipment.Onsite_date)
	well_drill_equipment.Onsite_time = formatDateString(well_drill_equipment.Onsite_time)
	well_drill_equipment.Row_effective_date = formatDateString(well_drill_equipment.Row_effective_date)
	well_drill_equipment.Row_expiry_date = formatDateString(well_drill_equipment.Row_expiry_date)






	rows, err := stmt.Exec(well_drill_equipment.Equipment_id, well_drill_equipment.Equipment_obs_no, well_drill_equipment.Active_ind, well_drill_equipment.Effective_date, well_drill_equipment.Expiry_date, well_drill_equipment.Offsite_date, well_drill_equipment.Offsite_time, well_drill_equipment.Onsite_date, well_drill_equipment.Onsite_time, well_drill_equipment.Operated_by_ba_id, well_drill_equipment.Parent_equipment_id, well_drill_equipment.Period_on_well, well_drill_equipment.Period_on_well_ouom, well_drill_equipment.Ppdm_guid, well_drill_equipment.Reference_num, well_drill_equipment.Remark, well_drill_equipment.Rented_ind, well_drill_equipment.Source, well_drill_equipment.Timezone, well_drill_equipment.Row_changed_by, well_drill_equipment.Row_changed_date, well_drill_equipment.Row_created_by, well_drill_equipment.Row_effective_date, well_drill_equipment.Row_expiry_date, well_drill_equipment.Row_quality, well_drill_equipment.Uwi)
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

func PatchWellDrillEquipment(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_drill_equipment set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "offsite_date" || key == "offsite_time" || key == "onsite_date" || key == "onsite_time" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where uwi = :id"

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

func DeleteWellDrillEquipment(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_drill_equipment dto.Well_drill_equipment
	well_drill_equipment.Uwi = id

	stmt, err := db.Prepare("delete from well_drill_equipment where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_drill_equipment.Uwi)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellEquipment(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_equipment")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_equipment

	for rows.Next() {
		var well_equipment dto.Well_equipment
		if err := rows.Scan(&well_equipment.Uwi, &well_equipment.Source, &well_equipment.Equipment_id, &well_equipment.Equip_obs_no, &well_equipment.Active_ind, &well_equipment.Effective_date, &well_equipment.Expiry_date, &well_equipment.Install_base_depth, &well_equipment.Install_base_depth_ouom, &well_equipment.Install_date, &well_equipment.Install_seq_no, &well_equipment.Install_top_depth, &well_equipment.Install_top_depth_ouom, &well_equipment.Parent_equipment_id, &well_equipment.Ppdm_guid, &well_equipment.Remark, &well_equipment.Removal_date, &well_equipment.String_id, &well_equipment.String_source, &well_equipment.Row_changed_by, &well_equipment.Row_changed_date, &well_equipment.Row_created_by, &well_equipment.Row_created_date, &well_equipment.Row_effective_date, &well_equipment.Row_expiry_date, &well_equipment.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_equipment to result
		result = append(result, well_equipment)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellEquipment(c *fiber.Ctx) error {
	var well_equipment dto.Well_equipment

	setDefaults(&well_equipment)

	if err := json.Unmarshal(c.Body(), &well_equipment); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_equipment values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26)")
	if err != nil {
		return err
	}
	well_equipment.Row_created_date = formatDate(well_equipment.Row_created_date)
	well_equipment.Row_changed_date = formatDate(well_equipment.Row_changed_date)
	well_equipment.Effective_date = formatDateString(well_equipment.Effective_date)
	well_equipment.Expiry_date = formatDateString(well_equipment.Expiry_date)
	well_equipment.Install_date = formatDateString(well_equipment.Install_date)
	well_equipment.Removal_date = formatDateString(well_equipment.Removal_date)
	well_equipment.Row_effective_date = formatDateString(well_equipment.Row_effective_date)
	well_equipment.Row_expiry_date = formatDateString(well_equipment.Row_expiry_date)






	rows, err := stmt.Exec(well_equipment.Uwi, well_equipment.Source, well_equipment.Equipment_id, well_equipment.Equip_obs_no, well_equipment.Active_ind, well_equipment.Effective_date, well_equipment.Expiry_date, well_equipment.Install_base_depth, well_equipment.Install_base_depth_ouom, well_equipment.Install_date, well_equipment.Install_seq_no, well_equipment.Install_top_depth, well_equipment.Install_top_depth_ouom, well_equipment.Parent_equipment_id, well_equipment.Ppdm_guid, well_equipment.Remark, well_equipment.Removal_date, well_equipment.String_id, well_equipment.String_source, well_equipment.Row_changed_by, well_equipment.Row_changed_date, well_equipment.Row_created_by, well_equipment.Row_created_date, well_equipment.Row_effective_date, well_equipment.Row_expiry_date, well_equipment.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellEquipment(c *fiber.Ctx) error {
	var well_equipment dto.Well_equipment

	setDefaults(&well_equipment)

	if err := json.Unmarshal(c.Body(), &well_equipment); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_equipment.Uwi = id

    if well_equipment.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_equipment set source = :1, equipment_id = :2, equip_obs_no = :3, active_ind = :4, effective_date = :5, expiry_date = :6, install_base_depth = :7, install_base_depth_ouom = :8, install_date = :9, install_seq_no = :10, install_top_depth = :11, install_top_depth_ouom = :12, parent_equipment_id = :13, ppdm_guid = :14, remark = :15, removal_date = :16, string_id = :17, string_source = :18, row_changed_by = :19, row_changed_date = :20, row_created_by = :21, row_effective_date = :22, row_expiry_date = :23, row_quality = :24 where uwi = :26")
	if err != nil {
		return err
	}

	well_equipment.Row_changed_date = formatDate(well_equipment.Row_changed_date)
	well_equipment.Effective_date = formatDateString(well_equipment.Effective_date)
	well_equipment.Expiry_date = formatDateString(well_equipment.Expiry_date)
	well_equipment.Install_date = formatDateString(well_equipment.Install_date)
	well_equipment.Removal_date = formatDateString(well_equipment.Removal_date)
	well_equipment.Row_effective_date = formatDateString(well_equipment.Row_effective_date)
	well_equipment.Row_expiry_date = formatDateString(well_equipment.Row_expiry_date)






	rows, err := stmt.Exec(well_equipment.Source, well_equipment.Equipment_id, well_equipment.Equip_obs_no, well_equipment.Active_ind, well_equipment.Effective_date, well_equipment.Expiry_date, well_equipment.Install_base_depth, well_equipment.Install_base_depth_ouom, well_equipment.Install_date, well_equipment.Install_seq_no, well_equipment.Install_top_depth, well_equipment.Install_top_depth_ouom, well_equipment.Parent_equipment_id, well_equipment.Ppdm_guid, well_equipment.Remark, well_equipment.Removal_date, well_equipment.String_id, well_equipment.String_source, well_equipment.Row_changed_by, well_equipment.Row_changed_date, well_equipment.Row_created_by, well_equipment.Row_effective_date, well_equipment.Row_expiry_date, well_equipment.Row_quality, well_equipment.Uwi)
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

func PatchWellEquipment(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_equipment set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "install_date" || key == "removal_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellEquipment(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_equipment dto.Well_equipment
	well_equipment.Uwi = id

	stmt, err := db.Prepare("delete from well_equipment where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_equipment.Uwi)
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



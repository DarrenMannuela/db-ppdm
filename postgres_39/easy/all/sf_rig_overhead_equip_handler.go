package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSfRigOverheadEquip(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sf_rig_overhead_equip")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sf_rig_overhead_equip

	for rows.Next() {
		var sf_rig_overhead_equip dto.Sf_rig_overhead_equip
		if err := rows.Scan(&sf_rig_overhead_equip.Sf_subtype, &sf_rig_overhead_equip.Rig_id, &sf_rig_overhead_equip.Equip_id, &sf_rig_overhead_equip.Active_ind, &sf_rig_overhead_equip.Capacity, &sf_rig_overhead_equip.Capacity_ouom, &sf_rig_overhead_equip.Capacity_type, &sf_rig_overhead_equip.Capacity_uom, &sf_rig_overhead_equip.Catalogue_equip_id, &sf_rig_overhead_equip.Effective_date, &sf_rig_overhead_equip.Equipment_id, &sf_rig_overhead_equip.Expiry_date, &sf_rig_overhead_equip.Input_type, &sf_rig_overhead_equip.Inside_diameter, &sf_rig_overhead_equip.Inside_diameter_ouom, &sf_rig_overhead_equip.Install_date, &sf_rig_overhead_equip.Overhead_count, &sf_rig_overhead_equip.Overhead_equip_type, &sf_rig_overhead_equip.Ppdm_guid, &sf_rig_overhead_equip.Remark, &sf_rig_overhead_equip.Remove_date, &sf_rig_overhead_equip.Source, &sf_rig_overhead_equip.Row_changed_by, &sf_rig_overhead_equip.Row_changed_date, &sf_rig_overhead_equip.Row_created_by, &sf_rig_overhead_equip.Row_created_date, &sf_rig_overhead_equip.Row_effective_date, &sf_rig_overhead_equip.Row_expiry_date, &sf_rig_overhead_equip.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sf_rig_overhead_equip to result
		result = append(result, sf_rig_overhead_equip)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSfRigOverheadEquip(c *fiber.Ctx) error {
	var sf_rig_overhead_equip dto.Sf_rig_overhead_equip

	setDefaults(&sf_rig_overhead_equip)

	if err := json.Unmarshal(c.Body(), &sf_rig_overhead_equip); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sf_rig_overhead_equip values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29)")
	if err != nil {
		return err
	}
	sf_rig_overhead_equip.Row_created_date = formatDate(sf_rig_overhead_equip.Row_created_date)
	sf_rig_overhead_equip.Row_changed_date = formatDate(sf_rig_overhead_equip.Row_changed_date)
	sf_rig_overhead_equip.Effective_date = formatDateString(sf_rig_overhead_equip.Effective_date)
	sf_rig_overhead_equip.Expiry_date = formatDateString(sf_rig_overhead_equip.Expiry_date)
	sf_rig_overhead_equip.Install_date = formatDateString(sf_rig_overhead_equip.Install_date)
	sf_rig_overhead_equip.Remove_date = formatDateString(sf_rig_overhead_equip.Remove_date)
	sf_rig_overhead_equip.Row_effective_date = formatDateString(sf_rig_overhead_equip.Row_effective_date)
	sf_rig_overhead_equip.Row_expiry_date = formatDateString(sf_rig_overhead_equip.Row_expiry_date)






	rows, err := stmt.Exec(sf_rig_overhead_equip.Sf_subtype, sf_rig_overhead_equip.Rig_id, sf_rig_overhead_equip.Equip_id, sf_rig_overhead_equip.Active_ind, sf_rig_overhead_equip.Capacity, sf_rig_overhead_equip.Capacity_ouom, sf_rig_overhead_equip.Capacity_type, sf_rig_overhead_equip.Capacity_uom, sf_rig_overhead_equip.Catalogue_equip_id, sf_rig_overhead_equip.Effective_date, sf_rig_overhead_equip.Equipment_id, sf_rig_overhead_equip.Expiry_date, sf_rig_overhead_equip.Input_type, sf_rig_overhead_equip.Inside_diameter, sf_rig_overhead_equip.Inside_diameter_ouom, sf_rig_overhead_equip.Install_date, sf_rig_overhead_equip.Overhead_count, sf_rig_overhead_equip.Overhead_equip_type, sf_rig_overhead_equip.Ppdm_guid, sf_rig_overhead_equip.Remark, sf_rig_overhead_equip.Remove_date, sf_rig_overhead_equip.Source, sf_rig_overhead_equip.Row_changed_by, sf_rig_overhead_equip.Row_changed_date, sf_rig_overhead_equip.Row_created_by, sf_rig_overhead_equip.Row_created_date, sf_rig_overhead_equip.Row_effective_date, sf_rig_overhead_equip.Row_expiry_date, sf_rig_overhead_equip.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSfRigOverheadEquip(c *fiber.Ctx) error {
	var sf_rig_overhead_equip dto.Sf_rig_overhead_equip

	setDefaults(&sf_rig_overhead_equip)

	if err := json.Unmarshal(c.Body(), &sf_rig_overhead_equip); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sf_rig_overhead_equip.Sf_subtype = id

    if sf_rig_overhead_equip.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sf_rig_overhead_equip set rig_id = :1, equip_id = :2, active_ind = :3, capacity = :4, capacity_ouom = :5, capacity_type = :6, capacity_uom = :7, catalogue_equip_id = :8, effective_date = :9, equipment_id = :10, expiry_date = :11, input_type = :12, inside_diameter = :13, inside_diameter_ouom = :14, install_date = :15, overhead_count = :16, overhead_equip_type = :17, ppdm_guid = :18, remark = :19, remove_date = :20, source = :21, row_changed_by = :22, row_changed_date = :23, row_created_by = :24, row_effective_date = :25, row_expiry_date = :26, row_quality = :27 where sf_subtype = :29")
	if err != nil {
		return err
	}

	sf_rig_overhead_equip.Row_changed_date = formatDate(sf_rig_overhead_equip.Row_changed_date)
	sf_rig_overhead_equip.Effective_date = formatDateString(sf_rig_overhead_equip.Effective_date)
	sf_rig_overhead_equip.Expiry_date = formatDateString(sf_rig_overhead_equip.Expiry_date)
	sf_rig_overhead_equip.Install_date = formatDateString(sf_rig_overhead_equip.Install_date)
	sf_rig_overhead_equip.Remove_date = formatDateString(sf_rig_overhead_equip.Remove_date)
	sf_rig_overhead_equip.Row_effective_date = formatDateString(sf_rig_overhead_equip.Row_effective_date)
	sf_rig_overhead_equip.Row_expiry_date = formatDateString(sf_rig_overhead_equip.Row_expiry_date)






	rows, err := stmt.Exec(sf_rig_overhead_equip.Rig_id, sf_rig_overhead_equip.Equip_id, sf_rig_overhead_equip.Active_ind, sf_rig_overhead_equip.Capacity, sf_rig_overhead_equip.Capacity_ouom, sf_rig_overhead_equip.Capacity_type, sf_rig_overhead_equip.Capacity_uom, sf_rig_overhead_equip.Catalogue_equip_id, sf_rig_overhead_equip.Effective_date, sf_rig_overhead_equip.Equipment_id, sf_rig_overhead_equip.Expiry_date, sf_rig_overhead_equip.Input_type, sf_rig_overhead_equip.Inside_diameter, sf_rig_overhead_equip.Inside_diameter_ouom, sf_rig_overhead_equip.Install_date, sf_rig_overhead_equip.Overhead_count, sf_rig_overhead_equip.Overhead_equip_type, sf_rig_overhead_equip.Ppdm_guid, sf_rig_overhead_equip.Remark, sf_rig_overhead_equip.Remove_date, sf_rig_overhead_equip.Source, sf_rig_overhead_equip.Row_changed_by, sf_rig_overhead_equip.Row_changed_date, sf_rig_overhead_equip.Row_created_by, sf_rig_overhead_equip.Row_effective_date, sf_rig_overhead_equip.Row_expiry_date, sf_rig_overhead_equip.Row_quality, sf_rig_overhead_equip.Sf_subtype)
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

func PatchSfRigOverheadEquip(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sf_rig_overhead_equip set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "install_date" || key == "remove_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where sf_subtype = :id"

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

func DeleteSfRigOverheadEquip(c *fiber.Ctx) error {
	id := c.Params("id")
	var sf_rig_overhead_equip dto.Sf_rig_overhead_equip
	sf_rig_overhead_equip.Sf_subtype = id

	stmt, err := db.Prepare("delete from sf_rig_overhead_equip where sf_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sf_rig_overhead_equip.Sf_subtype)
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



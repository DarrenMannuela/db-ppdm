package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSfRigShaker(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sf_rig_shaker")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sf_rig_shaker

	for rows.Next() {
		var sf_rig_shaker dto.Sf_rig_shaker
		if err := rows.Scan(&sf_rig_shaker.Sf_subtype, &sf_rig_shaker.Rig_id, &sf_rig_shaker.Shaker_id, &sf_rig_shaker.Active_ind, &sf_rig_shaker.Capacity, &sf_rig_shaker.Capacity_ouom, &sf_rig_shaker.Catalogue_equip_id, &sf_rig_shaker.Effective_date, &sf_rig_shaker.Equipment_id, &sf_rig_shaker.Expiry_date, &sf_rig_shaker.Input_type, &sf_rig_shaker.Install_date, &sf_rig_shaker.Position_desc, &sf_rig_shaker.Ppdm_guid, &sf_rig_shaker.Reference_num, &sf_rig_shaker.Remark, &sf_rig_shaker.Remove_date, &sf_rig_shaker.Shaker_count, &sf_rig_shaker.Shaker_size, &sf_rig_shaker.Shaker_size_desc, &sf_rig_shaker.Shaker_size_ouom, &sf_rig_shaker.Source, &sf_rig_shaker.Row_changed_by, &sf_rig_shaker.Row_changed_date, &sf_rig_shaker.Row_created_by, &sf_rig_shaker.Row_created_date, &sf_rig_shaker.Row_effective_date, &sf_rig_shaker.Row_expiry_date, &sf_rig_shaker.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sf_rig_shaker to result
		result = append(result, sf_rig_shaker)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSfRigShaker(c *fiber.Ctx) error {
	var sf_rig_shaker dto.Sf_rig_shaker

	setDefaults(&sf_rig_shaker)

	if err := json.Unmarshal(c.Body(), &sf_rig_shaker); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sf_rig_shaker values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29)")
	if err != nil {
		return err
	}
	sf_rig_shaker.Row_created_date = formatDate(sf_rig_shaker.Row_created_date)
	sf_rig_shaker.Row_changed_date = formatDate(sf_rig_shaker.Row_changed_date)
	sf_rig_shaker.Effective_date = formatDateString(sf_rig_shaker.Effective_date)
	sf_rig_shaker.Expiry_date = formatDateString(sf_rig_shaker.Expiry_date)
	sf_rig_shaker.Install_date = formatDateString(sf_rig_shaker.Install_date)
	sf_rig_shaker.Remove_date = formatDateString(sf_rig_shaker.Remove_date)
	sf_rig_shaker.Row_effective_date = formatDateString(sf_rig_shaker.Row_effective_date)
	sf_rig_shaker.Row_expiry_date = formatDateString(sf_rig_shaker.Row_expiry_date)






	rows, err := stmt.Exec(sf_rig_shaker.Sf_subtype, sf_rig_shaker.Rig_id, sf_rig_shaker.Shaker_id, sf_rig_shaker.Active_ind, sf_rig_shaker.Capacity, sf_rig_shaker.Capacity_ouom, sf_rig_shaker.Catalogue_equip_id, sf_rig_shaker.Effective_date, sf_rig_shaker.Equipment_id, sf_rig_shaker.Expiry_date, sf_rig_shaker.Input_type, sf_rig_shaker.Install_date, sf_rig_shaker.Position_desc, sf_rig_shaker.Ppdm_guid, sf_rig_shaker.Reference_num, sf_rig_shaker.Remark, sf_rig_shaker.Remove_date, sf_rig_shaker.Shaker_count, sf_rig_shaker.Shaker_size, sf_rig_shaker.Shaker_size_desc, sf_rig_shaker.Shaker_size_ouom, sf_rig_shaker.Source, sf_rig_shaker.Row_changed_by, sf_rig_shaker.Row_changed_date, sf_rig_shaker.Row_created_by, sf_rig_shaker.Row_created_date, sf_rig_shaker.Row_effective_date, sf_rig_shaker.Row_expiry_date, sf_rig_shaker.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSfRigShaker(c *fiber.Ctx) error {
	var sf_rig_shaker dto.Sf_rig_shaker

	setDefaults(&sf_rig_shaker)

	if err := json.Unmarshal(c.Body(), &sf_rig_shaker); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sf_rig_shaker.Sf_subtype = id

    if sf_rig_shaker.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sf_rig_shaker set rig_id = :1, shaker_id = :2, active_ind = :3, capacity = :4, capacity_ouom = :5, catalogue_equip_id = :6, effective_date = :7, equipment_id = :8, expiry_date = :9, input_type = :10, install_date = :11, position_desc = :12, ppdm_guid = :13, reference_num = :14, remark = :15, remove_date = :16, shaker_count = :17, shaker_size = :18, shaker_size_desc = :19, shaker_size_ouom = :20, source = :21, row_changed_by = :22, row_changed_date = :23, row_created_by = :24, row_effective_date = :25, row_expiry_date = :26, row_quality = :27 where sf_subtype = :29")
	if err != nil {
		return err
	}

	sf_rig_shaker.Row_changed_date = formatDate(sf_rig_shaker.Row_changed_date)
	sf_rig_shaker.Effective_date = formatDateString(sf_rig_shaker.Effective_date)
	sf_rig_shaker.Expiry_date = formatDateString(sf_rig_shaker.Expiry_date)
	sf_rig_shaker.Install_date = formatDateString(sf_rig_shaker.Install_date)
	sf_rig_shaker.Remove_date = formatDateString(sf_rig_shaker.Remove_date)
	sf_rig_shaker.Row_effective_date = formatDateString(sf_rig_shaker.Row_effective_date)
	sf_rig_shaker.Row_expiry_date = formatDateString(sf_rig_shaker.Row_expiry_date)






	rows, err := stmt.Exec(sf_rig_shaker.Rig_id, sf_rig_shaker.Shaker_id, sf_rig_shaker.Active_ind, sf_rig_shaker.Capacity, sf_rig_shaker.Capacity_ouom, sf_rig_shaker.Catalogue_equip_id, sf_rig_shaker.Effective_date, sf_rig_shaker.Equipment_id, sf_rig_shaker.Expiry_date, sf_rig_shaker.Input_type, sf_rig_shaker.Install_date, sf_rig_shaker.Position_desc, sf_rig_shaker.Ppdm_guid, sf_rig_shaker.Reference_num, sf_rig_shaker.Remark, sf_rig_shaker.Remove_date, sf_rig_shaker.Shaker_count, sf_rig_shaker.Shaker_size, sf_rig_shaker.Shaker_size_desc, sf_rig_shaker.Shaker_size_ouom, sf_rig_shaker.Source, sf_rig_shaker.Row_changed_by, sf_rig_shaker.Row_changed_date, sf_rig_shaker.Row_created_by, sf_rig_shaker.Row_effective_date, sf_rig_shaker.Row_expiry_date, sf_rig_shaker.Row_quality, sf_rig_shaker.Sf_subtype)
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

func PatchSfRigShaker(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sf_rig_shaker set "
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

func DeleteSfRigShaker(c *fiber.Ctx) error {
	id := c.Params("id")
	var sf_rig_shaker dto.Sf_rig_shaker
	sf_rig_shaker.Sf_subtype = id

	stmt, err := db.Prepare("delete from sf_rig_shaker where sf_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sf_rig_shaker.Sf_subtype)
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



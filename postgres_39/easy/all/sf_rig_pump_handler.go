package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSfRigPump(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sf_rig_pump")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sf_rig_pump

	for rows.Next() {
		var sf_rig_pump dto.Sf_rig_pump
		if err := rows.Scan(&sf_rig_pump.Sf_subtype, &sf_rig_pump.Rig_id, &sf_rig_pump.Pump_id, &sf_rig_pump.Active_ind, &sf_rig_pump.Catalogue_equip_id, &sf_rig_pump.Effective_date, &sf_rig_pump.Equipment_id, &sf_rig_pump.Expiry_date, &sf_rig_pump.Input_type, &sf_rig_pump.Install_date, &sf_rig_pump.Ppdm_guid, &sf_rig_pump.Pump_count, &sf_rig_pump.Pump_function, &sf_rig_pump.Pump_hp, &sf_rig_pump.Pump_hp_ouom, &sf_rig_pump.Pump_rating, &sf_rig_pump.Pump_rating_ouom, &sf_rig_pump.Pump_type, &sf_rig_pump.Reference_num, &sf_rig_pump.Remark, &sf_rig_pump.Remove_date, &sf_rig_pump.Source, &sf_rig_pump.Row_changed_by, &sf_rig_pump.Row_changed_date, &sf_rig_pump.Row_created_by, &sf_rig_pump.Row_created_date, &sf_rig_pump.Row_effective_date, &sf_rig_pump.Row_expiry_date, &sf_rig_pump.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sf_rig_pump to result
		result = append(result, sf_rig_pump)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSfRigPump(c *fiber.Ctx) error {
	var sf_rig_pump dto.Sf_rig_pump

	setDefaults(&sf_rig_pump)

	if err := json.Unmarshal(c.Body(), &sf_rig_pump); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sf_rig_pump values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29)")
	if err != nil {
		return err
	}
	sf_rig_pump.Row_created_date = formatDate(sf_rig_pump.Row_created_date)
	sf_rig_pump.Row_changed_date = formatDate(sf_rig_pump.Row_changed_date)
	sf_rig_pump.Effective_date = formatDateString(sf_rig_pump.Effective_date)
	sf_rig_pump.Expiry_date = formatDateString(sf_rig_pump.Expiry_date)
	sf_rig_pump.Install_date = formatDateString(sf_rig_pump.Install_date)
	sf_rig_pump.Remove_date = formatDateString(sf_rig_pump.Remove_date)
	sf_rig_pump.Row_effective_date = formatDateString(sf_rig_pump.Row_effective_date)
	sf_rig_pump.Row_expiry_date = formatDateString(sf_rig_pump.Row_expiry_date)






	rows, err := stmt.Exec(sf_rig_pump.Sf_subtype, sf_rig_pump.Rig_id, sf_rig_pump.Pump_id, sf_rig_pump.Active_ind, sf_rig_pump.Catalogue_equip_id, sf_rig_pump.Effective_date, sf_rig_pump.Equipment_id, sf_rig_pump.Expiry_date, sf_rig_pump.Input_type, sf_rig_pump.Install_date, sf_rig_pump.Ppdm_guid, sf_rig_pump.Pump_count, sf_rig_pump.Pump_function, sf_rig_pump.Pump_hp, sf_rig_pump.Pump_hp_ouom, sf_rig_pump.Pump_rating, sf_rig_pump.Pump_rating_ouom, sf_rig_pump.Pump_type, sf_rig_pump.Reference_num, sf_rig_pump.Remark, sf_rig_pump.Remove_date, sf_rig_pump.Source, sf_rig_pump.Row_changed_by, sf_rig_pump.Row_changed_date, sf_rig_pump.Row_created_by, sf_rig_pump.Row_created_date, sf_rig_pump.Row_effective_date, sf_rig_pump.Row_expiry_date, sf_rig_pump.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSfRigPump(c *fiber.Ctx) error {
	var sf_rig_pump dto.Sf_rig_pump

	setDefaults(&sf_rig_pump)

	if err := json.Unmarshal(c.Body(), &sf_rig_pump); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sf_rig_pump.Sf_subtype = id

    if sf_rig_pump.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sf_rig_pump set rig_id = :1, pump_id = :2, active_ind = :3, catalogue_equip_id = :4, effective_date = :5, equipment_id = :6, expiry_date = :7, input_type = :8, install_date = :9, ppdm_guid = :10, pump_count = :11, pump_function = :12, pump_hp = :13, pump_hp_ouom = :14, pump_rating = :15, pump_rating_ouom = :16, pump_type = :17, reference_num = :18, remark = :19, remove_date = :20, source = :21, row_changed_by = :22, row_changed_date = :23, row_created_by = :24, row_effective_date = :25, row_expiry_date = :26, row_quality = :27 where sf_subtype = :29")
	if err != nil {
		return err
	}

	sf_rig_pump.Row_changed_date = formatDate(sf_rig_pump.Row_changed_date)
	sf_rig_pump.Effective_date = formatDateString(sf_rig_pump.Effective_date)
	sf_rig_pump.Expiry_date = formatDateString(sf_rig_pump.Expiry_date)
	sf_rig_pump.Install_date = formatDateString(sf_rig_pump.Install_date)
	sf_rig_pump.Remove_date = formatDateString(sf_rig_pump.Remove_date)
	sf_rig_pump.Row_effective_date = formatDateString(sf_rig_pump.Row_effective_date)
	sf_rig_pump.Row_expiry_date = formatDateString(sf_rig_pump.Row_expiry_date)






	rows, err := stmt.Exec(sf_rig_pump.Rig_id, sf_rig_pump.Pump_id, sf_rig_pump.Active_ind, sf_rig_pump.Catalogue_equip_id, sf_rig_pump.Effective_date, sf_rig_pump.Equipment_id, sf_rig_pump.Expiry_date, sf_rig_pump.Input_type, sf_rig_pump.Install_date, sf_rig_pump.Ppdm_guid, sf_rig_pump.Pump_count, sf_rig_pump.Pump_function, sf_rig_pump.Pump_hp, sf_rig_pump.Pump_hp_ouom, sf_rig_pump.Pump_rating, sf_rig_pump.Pump_rating_ouom, sf_rig_pump.Pump_type, sf_rig_pump.Reference_num, sf_rig_pump.Remark, sf_rig_pump.Remove_date, sf_rig_pump.Source, sf_rig_pump.Row_changed_by, sf_rig_pump.Row_changed_date, sf_rig_pump.Row_created_by, sf_rig_pump.Row_effective_date, sf_rig_pump.Row_expiry_date, sf_rig_pump.Row_quality, sf_rig_pump.Sf_subtype)
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

func PatchSfRigPump(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sf_rig_pump set "
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

func DeleteSfRigPump(c *fiber.Ctx) error {
	id := c.Params("id")
	var sf_rig_pump dto.Sf_rig_pump
	sf_rig_pump.Sf_subtype = id

	stmt, err := db.Prepare("delete from sf_rig_pump where sf_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sf_rig_pump.Sf_subtype)
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



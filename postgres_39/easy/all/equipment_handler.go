package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetEquipment(c *fiber.Ctx) error {
	rows, err := db.Query("select * from equipment")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Equipment

	for rows.Next() {
		var equipment dto.Equipment
		if err := rows.Scan(&equipment.Equipment_id, &equipment.Acquire_date_new, &equipment.Active_ind, &equipment.Catalogue_equip_id, &equipment.Commission_date, &equipment.Current_owner_ba_id, &equipment.Decommission_date, &equipment.Description, &equipment.Effective_date, &equipment.Equipment_group, &equipment.Equipment_name, &equipment.Equipment_type, &equipment.Expiry_date, &equipment.Manufacturer_ba_id, &equipment.Ppdm_guid, &equipment.Purchase_date, &equipment.Reference_num, &equipment.Remark, &equipment.Serial_num, &equipment.Source, &equipment.Row_changed_by, &equipment.Row_changed_date, &equipment.Row_created_by, &equipment.Row_created_date, &equipment.Row_effective_date, &equipment.Row_expiry_date, &equipment.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append equipment to result
		result = append(result, equipment)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetEquipment(c *fiber.Ctx) error {
	var equipment dto.Equipment

	setDefaults(&equipment)

	if err := json.Unmarshal(c.Body(), &equipment); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into equipment values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27)")
	if err != nil {
		return err
	}
	equipment.Row_created_date = formatDate(equipment.Row_created_date)
	equipment.Row_changed_date = formatDate(equipment.Row_changed_date)
	equipment.Acquire_date_new = formatDateString(equipment.Acquire_date_new)
	equipment.Commission_date = formatDateString(equipment.Commission_date)
	equipment.Decommission_date = formatDateString(equipment.Decommission_date)
	equipment.Effective_date = formatDateString(equipment.Effective_date)
	equipment.Expiry_date = formatDateString(equipment.Expiry_date)
	equipment.Purchase_date = formatDateString(equipment.Purchase_date)
	equipment.Row_effective_date = formatDateString(equipment.Row_effective_date)
	equipment.Row_expiry_date = formatDateString(equipment.Row_expiry_date)






	rows, err := stmt.Exec(equipment.Equipment_id, equipment.Acquire_date_new, equipment.Active_ind, equipment.Catalogue_equip_id, equipment.Commission_date, equipment.Current_owner_ba_id, equipment.Decommission_date, equipment.Description, equipment.Effective_date, equipment.Equipment_group, equipment.Equipment_name, equipment.Equipment_type, equipment.Expiry_date, equipment.Manufacturer_ba_id, equipment.Ppdm_guid, equipment.Purchase_date, equipment.Reference_num, equipment.Remark, equipment.Serial_num, equipment.Source, equipment.Row_changed_by, equipment.Row_changed_date, equipment.Row_created_by, equipment.Row_created_date, equipment.Row_effective_date, equipment.Row_expiry_date, equipment.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateEquipment(c *fiber.Ctx) error {
	var equipment dto.Equipment

	setDefaults(&equipment)

	if err := json.Unmarshal(c.Body(), &equipment); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	equipment.Equipment_id = id

    if equipment.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update equipment set acquire_date_new = :1, active_ind = :2, catalogue_equip_id = :3, commission_date = :4, current_owner_ba_id = :5, decommission_date = :6, description = :7, effective_date = :8, equipment_group = :9, equipment_name = :10, equipment_type = :11, expiry_date = :12, manufacturer_ba_id = :13, ppdm_guid = :14, purchase_date = :15, reference_num = :16, remark = :17, serial_num = :18, source = :19, row_changed_by = :20, row_changed_date = :21, row_created_by = :22, row_effective_date = :23, row_expiry_date = :24, row_quality = :25 where equipment_id = :27")
	if err != nil {
		return err
	}

	equipment.Row_changed_date = formatDate(equipment.Row_changed_date)
	equipment.Acquire_date_new = formatDateString(equipment.Acquire_date_new)
	equipment.Commission_date = formatDateString(equipment.Commission_date)
	equipment.Decommission_date = formatDateString(equipment.Decommission_date)
	equipment.Effective_date = formatDateString(equipment.Effective_date)
	equipment.Expiry_date = formatDateString(equipment.Expiry_date)
	equipment.Purchase_date = formatDateString(equipment.Purchase_date)
	equipment.Row_effective_date = formatDateString(equipment.Row_effective_date)
	equipment.Row_expiry_date = formatDateString(equipment.Row_expiry_date)






	rows, err := stmt.Exec(equipment.Acquire_date_new, equipment.Active_ind, equipment.Catalogue_equip_id, equipment.Commission_date, equipment.Current_owner_ba_id, equipment.Decommission_date, equipment.Description, equipment.Effective_date, equipment.Equipment_group, equipment.Equipment_name, equipment.Equipment_type, equipment.Expiry_date, equipment.Manufacturer_ba_id, equipment.Ppdm_guid, equipment.Purchase_date, equipment.Reference_num, equipment.Remark, equipment.Serial_num, equipment.Source, equipment.Row_changed_by, equipment.Row_changed_date, equipment.Row_created_by, equipment.Row_effective_date, equipment.Row_expiry_date, equipment.Row_quality, equipment.Equipment_id)
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

func PatchEquipment(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update equipment set "
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
		} else if key == "acquire_date_new" || key == "commission_date" || key == "decommission_date" || key == "effective_date" || key == "expiry_date" || key == "purchase_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteEquipment(c *fiber.Ctx) error {
	id := c.Params("id")
	var equipment dto.Equipment
	equipment.Equipment_id = id

	stmt, err := db.Prepare("delete from equipment where equipment_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(equipment.Equipment_id)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetEquipmentBa(c *fiber.Ctx) error {
	rows, err := db.Query("select * from equipment_ba")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Equipment_ba

	for rows.Next() {
		var equipment_ba dto.Equipment_ba
		if err := rows.Scan(&equipment_ba.Equip_id, &equipment_ba.Ba_obs_no, &equipment_ba.Acquire_date, &equipment_ba.Active_ind, &equipment_ba.Effective_date, &equipment_ba.Equip_ba_id, &equipment_ba.Equip_lease_ind, &equipment_ba.Equip_owned_ind, &equipment_ba.Equip_rent_ind, &equipment_ba.Expiry_date, &equipment_ba.Finance_id, &equipment_ba.Ppdm_guid, &equipment_ba.Preferred_name, &equipment_ba.Purchase_date, &equipment_ba.Release_date, &equipment_ba.Remark, &equipment_ba.Role_type, &equipment_ba.Source, &equipment_ba.Row_changed_by, &equipment_ba.Row_changed_date, &equipment_ba.Row_created_by, &equipment_ba.Row_created_date, &equipment_ba.Row_effective_date, &equipment_ba.Row_expiry_date, &equipment_ba.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append equipment_ba to result
		result = append(result, equipment_ba)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetEquipmentBa(c *fiber.Ctx) error {
	var equipment_ba dto.Equipment_ba

	setDefaults(&equipment_ba)

	if err := json.Unmarshal(c.Body(), &equipment_ba); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into equipment_ba values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25)")
	if err != nil {
		return err
	}
	equipment_ba.Row_created_date = formatDate(equipment_ba.Row_created_date)
	equipment_ba.Row_changed_date = formatDate(equipment_ba.Row_changed_date)
	equipment_ba.Acquire_date = formatDateString(equipment_ba.Acquire_date)
	equipment_ba.Effective_date = formatDateString(equipment_ba.Effective_date)
	equipment_ba.Expiry_date = formatDateString(equipment_ba.Expiry_date)
	equipment_ba.Purchase_date = formatDateString(equipment_ba.Purchase_date)
	equipment_ba.Release_date = formatDateString(equipment_ba.Release_date)
	equipment_ba.Row_effective_date = formatDateString(equipment_ba.Row_effective_date)
	equipment_ba.Row_expiry_date = formatDateString(equipment_ba.Row_expiry_date)






	rows, err := stmt.Exec(equipment_ba.Equip_id, equipment_ba.Ba_obs_no, equipment_ba.Acquire_date, equipment_ba.Active_ind, equipment_ba.Effective_date, equipment_ba.Equip_ba_id, equipment_ba.Equip_lease_ind, equipment_ba.Equip_owned_ind, equipment_ba.Equip_rent_ind, equipment_ba.Expiry_date, equipment_ba.Finance_id, equipment_ba.Ppdm_guid, equipment_ba.Preferred_name, equipment_ba.Purchase_date, equipment_ba.Release_date, equipment_ba.Remark, equipment_ba.Role_type, equipment_ba.Source, equipment_ba.Row_changed_by, equipment_ba.Row_changed_date, equipment_ba.Row_created_by, equipment_ba.Row_created_date, equipment_ba.Row_effective_date, equipment_ba.Row_expiry_date, equipment_ba.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateEquipmentBa(c *fiber.Ctx) error {
	var equipment_ba dto.Equipment_ba

	setDefaults(&equipment_ba)

	if err := json.Unmarshal(c.Body(), &equipment_ba); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	equipment_ba.Equip_id = id

    if equipment_ba.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update equipment_ba set ba_obs_no = :1, acquire_date = :2, active_ind = :3, effective_date = :4, equip_ba_id = :5, equip_lease_ind = :6, equip_owned_ind = :7, equip_rent_ind = :8, expiry_date = :9, finance_id = :10, ppdm_guid = :11, preferred_name = :12, purchase_date = :13, release_date = :14, remark = :15, role_type = :16, source = :17, row_changed_by = :18, row_changed_date = :19, row_created_by = :20, row_effective_date = :21, row_expiry_date = :22, row_quality = :23 where equip_id = :25")
	if err != nil {
		return err
	}

	equipment_ba.Row_changed_date = formatDate(equipment_ba.Row_changed_date)
	equipment_ba.Acquire_date = formatDateString(equipment_ba.Acquire_date)
	equipment_ba.Effective_date = formatDateString(equipment_ba.Effective_date)
	equipment_ba.Expiry_date = formatDateString(equipment_ba.Expiry_date)
	equipment_ba.Purchase_date = formatDateString(equipment_ba.Purchase_date)
	equipment_ba.Release_date = formatDateString(equipment_ba.Release_date)
	equipment_ba.Row_effective_date = formatDateString(equipment_ba.Row_effective_date)
	equipment_ba.Row_expiry_date = formatDateString(equipment_ba.Row_expiry_date)






	rows, err := stmt.Exec(equipment_ba.Ba_obs_no, equipment_ba.Acquire_date, equipment_ba.Active_ind, equipment_ba.Effective_date, equipment_ba.Equip_ba_id, equipment_ba.Equip_lease_ind, equipment_ba.Equip_owned_ind, equipment_ba.Equip_rent_ind, equipment_ba.Expiry_date, equipment_ba.Finance_id, equipment_ba.Ppdm_guid, equipment_ba.Preferred_name, equipment_ba.Purchase_date, equipment_ba.Release_date, equipment_ba.Remark, equipment_ba.Role_type, equipment_ba.Source, equipment_ba.Row_changed_by, equipment_ba.Row_changed_date, equipment_ba.Row_created_by, equipment_ba.Row_effective_date, equipment_ba.Row_expiry_date, equipment_ba.Row_quality, equipment_ba.Equip_id)
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

func PatchEquipmentBa(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update equipment_ba set "
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
		} else if key == "acquire_date" || key == "effective_date" || key == "expiry_date" || key == "purchase_date" || key == "release_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where equip_id = :id"

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

func DeleteEquipmentBa(c *fiber.Ctx) error {
	id := c.Params("id")
	var equipment_ba dto.Equipment_ba
	equipment_ba.Equip_id = id

	stmt, err := db.Prepare("delete from equipment_ba where equip_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(equipment_ba.Equip_id)
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



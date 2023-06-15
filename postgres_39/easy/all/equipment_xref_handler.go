package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetEquipmentXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from equipment_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Equipment_xref

	for rows.Next() {
		var equipment_xref dto.Equipment_xref
		if err := rows.Scan(&equipment_xref.Equipment_id, &equipment_xref.Equipment_part_id, &equipment_xref.Equipment_xref_obs_no, &equipment_xref.Active_ind, &equipment_xref.Commission_date, &equipment_xref.Decommission_date, &equipment_xref.Description, &equipment_xref.Effective_date, &equipment_xref.Equip_xref_type, &equipment_xref.Expiry_date, &equipment_xref.Interchangable_ind, &equipment_xref.Ppdm_guid, &equipment_xref.Remark, &equipment_xref.Source, &equipment_xref.Row_changed_by, &equipment_xref.Row_changed_date, &equipment_xref.Row_created_by, &equipment_xref.Row_created_date, &equipment_xref.Row_effective_date, &equipment_xref.Row_expiry_date, &equipment_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append equipment_xref to result
		result = append(result, equipment_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetEquipmentXref(c *fiber.Ctx) error {
	var equipment_xref dto.Equipment_xref

	setDefaults(&equipment_xref)

	if err := json.Unmarshal(c.Body(), &equipment_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into equipment_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	equipment_xref.Row_created_date = formatDate(equipment_xref.Row_created_date)
	equipment_xref.Row_changed_date = formatDate(equipment_xref.Row_changed_date)
	equipment_xref.Commission_date = formatDateString(equipment_xref.Commission_date)
	equipment_xref.Decommission_date = formatDateString(equipment_xref.Decommission_date)
	equipment_xref.Effective_date = formatDateString(equipment_xref.Effective_date)
	equipment_xref.Expiry_date = formatDateString(equipment_xref.Expiry_date)
	equipment_xref.Row_effective_date = formatDateString(equipment_xref.Row_effective_date)
	equipment_xref.Row_expiry_date = formatDateString(equipment_xref.Row_expiry_date)






	rows, err := stmt.Exec(equipment_xref.Equipment_id, equipment_xref.Equipment_part_id, equipment_xref.Equipment_xref_obs_no, equipment_xref.Active_ind, equipment_xref.Commission_date, equipment_xref.Decommission_date, equipment_xref.Description, equipment_xref.Effective_date, equipment_xref.Equip_xref_type, equipment_xref.Expiry_date, equipment_xref.Interchangable_ind, equipment_xref.Ppdm_guid, equipment_xref.Remark, equipment_xref.Source, equipment_xref.Row_changed_by, equipment_xref.Row_changed_date, equipment_xref.Row_created_by, equipment_xref.Row_created_date, equipment_xref.Row_effective_date, equipment_xref.Row_expiry_date, equipment_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateEquipmentXref(c *fiber.Ctx) error {
	var equipment_xref dto.Equipment_xref

	setDefaults(&equipment_xref)

	if err := json.Unmarshal(c.Body(), &equipment_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	equipment_xref.Equipment_id = id

    if equipment_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update equipment_xref set equipment_part_id = :1, equipment_xref_obs_no = :2, active_ind = :3, commission_date = :4, decommission_date = :5, description = :6, effective_date = :7, equip_xref_type = :8, expiry_date = :9, interchangable_ind = :10, ppdm_guid = :11, remark = :12, source = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where equipment_id = :21")
	if err != nil {
		return err
	}

	equipment_xref.Row_changed_date = formatDate(equipment_xref.Row_changed_date)
	equipment_xref.Commission_date = formatDateString(equipment_xref.Commission_date)
	equipment_xref.Decommission_date = formatDateString(equipment_xref.Decommission_date)
	equipment_xref.Effective_date = formatDateString(equipment_xref.Effective_date)
	equipment_xref.Expiry_date = formatDateString(equipment_xref.Expiry_date)
	equipment_xref.Row_effective_date = formatDateString(equipment_xref.Row_effective_date)
	equipment_xref.Row_expiry_date = formatDateString(equipment_xref.Row_expiry_date)






	rows, err := stmt.Exec(equipment_xref.Equipment_part_id, equipment_xref.Equipment_xref_obs_no, equipment_xref.Active_ind, equipment_xref.Commission_date, equipment_xref.Decommission_date, equipment_xref.Description, equipment_xref.Effective_date, equipment_xref.Equip_xref_type, equipment_xref.Expiry_date, equipment_xref.Interchangable_ind, equipment_xref.Ppdm_guid, equipment_xref.Remark, equipment_xref.Source, equipment_xref.Row_changed_by, equipment_xref.Row_changed_date, equipment_xref.Row_created_by, equipment_xref.Row_effective_date, equipment_xref.Row_expiry_date, equipment_xref.Row_quality, equipment_xref.Equipment_id)
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

func PatchEquipmentXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update equipment_xref set "
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
		} else if key == "commission_date" || key == "decommission_date" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteEquipmentXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var equipment_xref dto.Equipment_xref
	equipment_xref.Equipment_id = id

	stmt, err := db.Prepare("delete from equipment_xref where equipment_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(equipment_xref.Equipment_id)
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



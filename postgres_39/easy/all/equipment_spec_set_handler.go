package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetEquipmentSpecSet(c *fiber.Ctx) error {
	rows, err := db.Query("select * from equipment_spec_set")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Equipment_spec_set

	for rows.Next() {
		var equipment_spec_set dto.Equipment_spec_set
		if err := rows.Scan(&equipment_spec_set.Spec_set_id, &equipment_spec_set.Active_ind, &equipment_spec_set.Effective_date, &equipment_spec_set.Equipment_type, &equipment_spec_set.Expiry_date, &equipment_spec_set.Owner_ba_id, &equipment_spec_set.Ppdm_guid, &equipment_spec_set.Preferred_name, &equipment_spec_set.Remark, &equipment_spec_set.Source, &equipment_spec_set.Spec_set_desc, &equipment_spec_set.Spec_set_type, &equipment_spec_set.Row_changed_by, &equipment_spec_set.Row_changed_date, &equipment_spec_set.Row_created_by, &equipment_spec_set.Row_created_date, &equipment_spec_set.Row_effective_date, &equipment_spec_set.Row_expiry_date, &equipment_spec_set.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append equipment_spec_set to result
		result = append(result, equipment_spec_set)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetEquipmentSpecSet(c *fiber.Ctx) error {
	var equipment_spec_set dto.Equipment_spec_set

	setDefaults(&equipment_spec_set)

	if err := json.Unmarshal(c.Body(), &equipment_spec_set); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into equipment_spec_set values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	equipment_spec_set.Row_created_date = formatDate(equipment_spec_set.Row_created_date)
	equipment_spec_set.Row_changed_date = formatDate(equipment_spec_set.Row_changed_date)
	equipment_spec_set.Effective_date = formatDateString(equipment_spec_set.Effective_date)
	equipment_spec_set.Expiry_date = formatDateString(equipment_spec_set.Expiry_date)
	equipment_spec_set.Row_effective_date = formatDateString(equipment_spec_set.Row_effective_date)
	equipment_spec_set.Row_expiry_date = formatDateString(equipment_spec_set.Row_expiry_date)






	rows, err := stmt.Exec(equipment_spec_set.Spec_set_id, equipment_spec_set.Active_ind, equipment_spec_set.Effective_date, equipment_spec_set.Equipment_type, equipment_spec_set.Expiry_date, equipment_spec_set.Owner_ba_id, equipment_spec_set.Ppdm_guid, equipment_spec_set.Preferred_name, equipment_spec_set.Remark, equipment_spec_set.Source, equipment_spec_set.Spec_set_desc, equipment_spec_set.Spec_set_type, equipment_spec_set.Row_changed_by, equipment_spec_set.Row_changed_date, equipment_spec_set.Row_created_by, equipment_spec_set.Row_created_date, equipment_spec_set.Row_effective_date, equipment_spec_set.Row_expiry_date, equipment_spec_set.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateEquipmentSpecSet(c *fiber.Ctx) error {
	var equipment_spec_set dto.Equipment_spec_set

	setDefaults(&equipment_spec_set)

	if err := json.Unmarshal(c.Body(), &equipment_spec_set); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	equipment_spec_set.Spec_set_id = id

    if equipment_spec_set.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update equipment_spec_set set active_ind = :1, effective_date = :2, equipment_type = :3, expiry_date = :4, owner_ba_id = :5, ppdm_guid = :6, preferred_name = :7, remark = :8, source = :9, spec_set_desc = :10, spec_set_type = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where spec_set_id = :19")
	if err != nil {
		return err
	}

	equipment_spec_set.Row_changed_date = formatDate(equipment_spec_set.Row_changed_date)
	equipment_spec_set.Effective_date = formatDateString(equipment_spec_set.Effective_date)
	equipment_spec_set.Expiry_date = formatDateString(equipment_spec_set.Expiry_date)
	equipment_spec_set.Row_effective_date = formatDateString(equipment_spec_set.Row_effective_date)
	equipment_spec_set.Row_expiry_date = formatDateString(equipment_spec_set.Row_expiry_date)






	rows, err := stmt.Exec(equipment_spec_set.Active_ind, equipment_spec_set.Effective_date, equipment_spec_set.Equipment_type, equipment_spec_set.Expiry_date, equipment_spec_set.Owner_ba_id, equipment_spec_set.Ppdm_guid, equipment_spec_set.Preferred_name, equipment_spec_set.Remark, equipment_spec_set.Source, equipment_spec_set.Spec_set_desc, equipment_spec_set.Spec_set_type, equipment_spec_set.Row_changed_by, equipment_spec_set.Row_changed_date, equipment_spec_set.Row_created_by, equipment_spec_set.Row_effective_date, equipment_spec_set.Row_expiry_date, equipment_spec_set.Row_quality, equipment_spec_set.Spec_set_id)
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

func PatchEquipmentSpecSet(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update equipment_spec_set set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where spec_set_id = :id"

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

func DeleteEquipmentSpecSet(c *fiber.Ctx) error {
	id := c.Params("id")
	var equipment_spec_set dto.Equipment_spec_set
	equipment_spec_set.Spec_set_id = id

	stmt, err := db.Prepare("delete from equipment_spec_set where spec_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(equipment_spec_set.Spec_set_id)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetEquipmentSpecSetSpec(c *fiber.Ctx) error {
	rows, err := db.Query("select * from equipment_spec_set_spec")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Equipment_spec_set_spec

	for rows.Next() {
		var equipment_spec_set_spec dto.Equipment_spec_set_spec
		if err := rows.Scan(&equipment_spec_set_spec.Spec_set_id, &equipment_spec_set_spec.Spec_type, &equipment_spec_set_spec.Active_ind, &equipment_spec_set_spec.Effective_date, &equipment_spec_set_spec.Expiry_date, &equipment_spec_set_spec.Ppdm_guid, &equipment_spec_set_spec.Preferred_name, &equipment_spec_set_spec.Reference_value_type, &equipment_spec_set_spec.Remark, &equipment_spec_set_spec.Source, &equipment_spec_set_spec.Row_changed_by, &equipment_spec_set_spec.Row_changed_date, &equipment_spec_set_spec.Row_created_by, &equipment_spec_set_spec.Row_created_date, &equipment_spec_set_spec.Row_effective_date, &equipment_spec_set_spec.Row_expiry_date, &equipment_spec_set_spec.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append equipment_spec_set_spec to result
		result = append(result, equipment_spec_set_spec)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetEquipmentSpecSetSpec(c *fiber.Ctx) error {
	var equipment_spec_set_spec dto.Equipment_spec_set_spec

	setDefaults(&equipment_spec_set_spec)

	if err := json.Unmarshal(c.Body(), &equipment_spec_set_spec); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into equipment_spec_set_spec values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	equipment_spec_set_spec.Row_created_date = formatDate(equipment_spec_set_spec.Row_created_date)
	equipment_spec_set_spec.Row_changed_date = formatDate(equipment_spec_set_spec.Row_changed_date)
	equipment_spec_set_spec.Effective_date = formatDateString(equipment_spec_set_spec.Effective_date)
	equipment_spec_set_spec.Expiry_date = formatDateString(equipment_spec_set_spec.Expiry_date)
	equipment_spec_set_spec.Row_effective_date = formatDateString(equipment_spec_set_spec.Row_effective_date)
	equipment_spec_set_spec.Row_expiry_date = formatDateString(equipment_spec_set_spec.Row_expiry_date)






	rows, err := stmt.Exec(equipment_spec_set_spec.Spec_set_id, equipment_spec_set_spec.Spec_type, equipment_spec_set_spec.Active_ind, equipment_spec_set_spec.Effective_date, equipment_spec_set_spec.Expiry_date, equipment_spec_set_spec.Ppdm_guid, equipment_spec_set_spec.Preferred_name, equipment_spec_set_spec.Reference_value_type, equipment_spec_set_spec.Remark, equipment_spec_set_spec.Source, equipment_spec_set_spec.Row_changed_by, equipment_spec_set_spec.Row_changed_date, equipment_spec_set_spec.Row_created_by, equipment_spec_set_spec.Row_created_date, equipment_spec_set_spec.Row_effective_date, equipment_spec_set_spec.Row_expiry_date, equipment_spec_set_spec.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateEquipmentSpecSetSpec(c *fiber.Ctx) error {
	var equipment_spec_set_spec dto.Equipment_spec_set_spec

	setDefaults(&equipment_spec_set_spec)

	if err := json.Unmarshal(c.Body(), &equipment_spec_set_spec); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	equipment_spec_set_spec.Spec_set_id = id

    if equipment_spec_set_spec.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update equipment_spec_set_spec set spec_type = :1, active_ind = :2, effective_date = :3, expiry_date = :4, ppdm_guid = :5, preferred_name = :6, reference_value_type = :7, remark = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where spec_set_id = :17")
	if err != nil {
		return err
	}

	equipment_spec_set_spec.Row_changed_date = formatDate(equipment_spec_set_spec.Row_changed_date)
	equipment_spec_set_spec.Effective_date = formatDateString(equipment_spec_set_spec.Effective_date)
	equipment_spec_set_spec.Expiry_date = formatDateString(equipment_spec_set_spec.Expiry_date)
	equipment_spec_set_spec.Row_effective_date = formatDateString(equipment_spec_set_spec.Row_effective_date)
	equipment_spec_set_spec.Row_expiry_date = formatDateString(equipment_spec_set_spec.Row_expiry_date)






	rows, err := stmt.Exec(equipment_spec_set_spec.Spec_type, equipment_spec_set_spec.Active_ind, equipment_spec_set_spec.Effective_date, equipment_spec_set_spec.Expiry_date, equipment_spec_set_spec.Ppdm_guid, equipment_spec_set_spec.Preferred_name, equipment_spec_set_spec.Reference_value_type, equipment_spec_set_spec.Remark, equipment_spec_set_spec.Source, equipment_spec_set_spec.Row_changed_by, equipment_spec_set_spec.Row_changed_date, equipment_spec_set_spec.Row_created_by, equipment_spec_set_spec.Row_effective_date, equipment_spec_set_spec.Row_expiry_date, equipment_spec_set_spec.Row_quality, equipment_spec_set_spec.Spec_set_id)
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

func PatchEquipmentSpecSetSpec(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update equipment_spec_set_spec set "
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

func DeleteEquipmentSpecSetSpec(c *fiber.Ctx) error {
	id := c.Params("id")
	var equipment_spec_set_spec dto.Equipment_spec_set_spec
	equipment_spec_set_spec.Spec_set_id = id

	stmt, err := db.Prepare("delete from equipment_spec_set_spec where spec_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(equipment_spec_set_spec.Spec_set_id)
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



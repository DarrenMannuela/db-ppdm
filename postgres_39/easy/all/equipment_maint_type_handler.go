package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetEquipmentMaintType(c *fiber.Ctx) error {
	rows, err := db.Query("select * from equipment_maint_type")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Equipment_maint_type

	for rows.Next() {
		var equipment_maint_type dto.Equipment_maint_type
		if err := rows.Scan(&equipment_maint_type.Catalogue_equip_id, &equipment_maint_type.Maint_type, &equipment_maint_type.Abbreviation, &equipment_maint_type.Active_ind, &equipment_maint_type.Effective_date, &equipment_maint_type.Expiry_date, &equipment_maint_type.Long_name, &equipment_maint_type.Ppdm_guid, &equipment_maint_type.Remark, &equipment_maint_type.Short_name, &equipment_maint_type.Source, &equipment_maint_type.Row_changed_by, &equipment_maint_type.Row_changed_date, &equipment_maint_type.Row_created_by, &equipment_maint_type.Row_created_date, &equipment_maint_type.Row_effective_date, &equipment_maint_type.Row_expiry_date, &equipment_maint_type.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append equipment_maint_type to result
		result = append(result, equipment_maint_type)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetEquipmentMaintType(c *fiber.Ctx) error {
	var equipment_maint_type dto.Equipment_maint_type

	setDefaults(&equipment_maint_type)

	if err := json.Unmarshal(c.Body(), &equipment_maint_type); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into equipment_maint_type values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	equipment_maint_type.Row_created_date = formatDate(equipment_maint_type.Row_created_date)
	equipment_maint_type.Row_changed_date = formatDate(equipment_maint_type.Row_changed_date)
	equipment_maint_type.Effective_date = formatDateString(equipment_maint_type.Effective_date)
	equipment_maint_type.Expiry_date = formatDateString(equipment_maint_type.Expiry_date)
	equipment_maint_type.Row_effective_date = formatDateString(equipment_maint_type.Row_effective_date)
	equipment_maint_type.Row_expiry_date = formatDateString(equipment_maint_type.Row_expiry_date)






	rows, err := stmt.Exec(equipment_maint_type.Catalogue_equip_id, equipment_maint_type.Maint_type, equipment_maint_type.Abbreviation, equipment_maint_type.Active_ind, equipment_maint_type.Effective_date, equipment_maint_type.Expiry_date, equipment_maint_type.Long_name, equipment_maint_type.Ppdm_guid, equipment_maint_type.Remark, equipment_maint_type.Short_name, equipment_maint_type.Source, equipment_maint_type.Row_changed_by, equipment_maint_type.Row_changed_date, equipment_maint_type.Row_created_by, equipment_maint_type.Row_created_date, equipment_maint_type.Row_effective_date, equipment_maint_type.Row_expiry_date, equipment_maint_type.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateEquipmentMaintType(c *fiber.Ctx) error {
	var equipment_maint_type dto.Equipment_maint_type

	setDefaults(&equipment_maint_type)

	if err := json.Unmarshal(c.Body(), &equipment_maint_type); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	equipment_maint_type.Catalogue_equip_id = id

    if equipment_maint_type.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update equipment_maint_type set maint_type = :1, abbreviation = :2, active_ind = :3, effective_date = :4, expiry_date = :5, long_name = :6, ppdm_guid = :7, remark = :8, short_name = :9, source = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where catalogue_equip_id = :18")
	if err != nil {
		return err
	}

	equipment_maint_type.Row_changed_date = formatDate(equipment_maint_type.Row_changed_date)
	equipment_maint_type.Effective_date = formatDateString(equipment_maint_type.Effective_date)
	equipment_maint_type.Expiry_date = formatDateString(equipment_maint_type.Expiry_date)
	equipment_maint_type.Row_effective_date = formatDateString(equipment_maint_type.Row_effective_date)
	equipment_maint_type.Row_expiry_date = formatDateString(equipment_maint_type.Row_expiry_date)






	rows, err := stmt.Exec(equipment_maint_type.Maint_type, equipment_maint_type.Abbreviation, equipment_maint_type.Active_ind, equipment_maint_type.Effective_date, equipment_maint_type.Expiry_date, equipment_maint_type.Long_name, equipment_maint_type.Ppdm_guid, equipment_maint_type.Remark, equipment_maint_type.Short_name, equipment_maint_type.Source, equipment_maint_type.Row_changed_by, equipment_maint_type.Row_changed_date, equipment_maint_type.Row_created_by, equipment_maint_type.Row_effective_date, equipment_maint_type.Row_expiry_date, equipment_maint_type.Row_quality, equipment_maint_type.Catalogue_equip_id)
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

func PatchEquipmentMaintType(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update equipment_maint_type set "
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
	query += " where catalogue_equip_id = :id"

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

func DeleteEquipmentMaintType(c *fiber.Ctx) error {
	id := c.Params("id")
	var equipment_maint_type dto.Equipment_maint_type
	equipment_maint_type.Catalogue_equip_id = id

	stmt, err := db.Prepare("delete from equipment_maint_type where catalogue_equip_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(equipment_maint_type.Catalogue_equip_id)
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



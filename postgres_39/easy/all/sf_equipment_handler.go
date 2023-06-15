package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSfEquipment(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sf_equipment")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sf_equipment

	for rows.Next() {
		var sf_equipment dto.Sf_equipment
		if err := rows.Scan(&sf_equipment.Sf_subtype, &sf_equipment.Support_facility_id, &sf_equipment.Component_id, &sf_equipment.Active_ind, &sf_equipment.Catalogue_equip_id, &sf_equipment.Component_count, &sf_equipment.Description, &sf_equipment.Effective_date, &sf_equipment.Equipment_id, &sf_equipment.Equipment_name, &sf_equipment.Expiry_date, &sf_equipment.Install_date, &sf_equipment.Ppdm_guid, &sf_equipment.Purchase_date, &sf_equipment.Reference_num, &sf_equipment.Remark, &sf_equipment.Remove_date, &sf_equipment.Source, &sf_equipment.Row_changed_by, &sf_equipment.Row_changed_date, &sf_equipment.Row_created_by, &sf_equipment.Row_created_date, &sf_equipment.Row_effective_date, &sf_equipment.Row_expiry_date, &sf_equipment.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sf_equipment to result
		result = append(result, sf_equipment)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSfEquipment(c *fiber.Ctx) error {
	var sf_equipment dto.Sf_equipment

	setDefaults(&sf_equipment)

	if err := json.Unmarshal(c.Body(), &sf_equipment); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sf_equipment values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25)")
	if err != nil {
		return err
	}
	sf_equipment.Row_created_date = formatDate(sf_equipment.Row_created_date)
	sf_equipment.Row_changed_date = formatDate(sf_equipment.Row_changed_date)
	sf_equipment.Effective_date = formatDateString(sf_equipment.Effective_date)
	sf_equipment.Expiry_date = formatDateString(sf_equipment.Expiry_date)
	sf_equipment.Install_date = formatDateString(sf_equipment.Install_date)
	sf_equipment.Purchase_date = formatDateString(sf_equipment.Purchase_date)
	sf_equipment.Remove_date = formatDateString(sf_equipment.Remove_date)
	sf_equipment.Row_effective_date = formatDateString(sf_equipment.Row_effective_date)
	sf_equipment.Row_expiry_date = formatDateString(sf_equipment.Row_expiry_date)






	rows, err := stmt.Exec(sf_equipment.Sf_subtype, sf_equipment.Support_facility_id, sf_equipment.Component_id, sf_equipment.Active_ind, sf_equipment.Catalogue_equip_id, sf_equipment.Component_count, sf_equipment.Description, sf_equipment.Effective_date, sf_equipment.Equipment_id, sf_equipment.Equipment_name, sf_equipment.Expiry_date, sf_equipment.Install_date, sf_equipment.Ppdm_guid, sf_equipment.Purchase_date, sf_equipment.Reference_num, sf_equipment.Remark, sf_equipment.Remove_date, sf_equipment.Source, sf_equipment.Row_changed_by, sf_equipment.Row_changed_date, sf_equipment.Row_created_by, sf_equipment.Row_created_date, sf_equipment.Row_effective_date, sf_equipment.Row_expiry_date, sf_equipment.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSfEquipment(c *fiber.Ctx) error {
	var sf_equipment dto.Sf_equipment

	setDefaults(&sf_equipment)

	if err := json.Unmarshal(c.Body(), &sf_equipment); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sf_equipment.Sf_subtype = id

    if sf_equipment.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sf_equipment set support_facility_id = :1, component_id = :2, active_ind = :3, catalogue_equip_id = :4, component_count = :5, description = :6, effective_date = :7, equipment_id = :8, equipment_name = :9, expiry_date = :10, install_date = :11, ppdm_guid = :12, purchase_date = :13, reference_num = :14, remark = :15, remove_date = :16, source = :17, row_changed_by = :18, row_changed_date = :19, row_created_by = :20, row_effective_date = :21, row_expiry_date = :22, row_quality = :23 where sf_subtype = :25")
	if err != nil {
		return err
	}

	sf_equipment.Row_changed_date = formatDate(sf_equipment.Row_changed_date)
	sf_equipment.Effective_date = formatDateString(sf_equipment.Effective_date)
	sf_equipment.Expiry_date = formatDateString(sf_equipment.Expiry_date)
	sf_equipment.Install_date = formatDateString(sf_equipment.Install_date)
	sf_equipment.Purchase_date = formatDateString(sf_equipment.Purchase_date)
	sf_equipment.Remove_date = formatDateString(sf_equipment.Remove_date)
	sf_equipment.Row_effective_date = formatDateString(sf_equipment.Row_effective_date)
	sf_equipment.Row_expiry_date = formatDateString(sf_equipment.Row_expiry_date)






	rows, err := stmt.Exec(sf_equipment.Support_facility_id, sf_equipment.Component_id, sf_equipment.Active_ind, sf_equipment.Catalogue_equip_id, sf_equipment.Component_count, sf_equipment.Description, sf_equipment.Effective_date, sf_equipment.Equipment_id, sf_equipment.Equipment_name, sf_equipment.Expiry_date, sf_equipment.Install_date, sf_equipment.Ppdm_guid, sf_equipment.Purchase_date, sf_equipment.Reference_num, sf_equipment.Remark, sf_equipment.Remove_date, sf_equipment.Source, sf_equipment.Row_changed_by, sf_equipment.Row_changed_date, sf_equipment.Row_created_by, sf_equipment.Row_effective_date, sf_equipment.Row_expiry_date, sf_equipment.Row_quality, sf_equipment.Sf_subtype)
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

func PatchSfEquipment(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sf_equipment set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "install_date" || key == "purchase_date" || key == "remove_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteSfEquipment(c *fiber.Ctx) error {
	id := c.Params("id")
	var sf_equipment dto.Sf_equipment
	sf_equipment.Sf_subtype = id

	stmt, err := db.Prepare("delete from sf_equipment where sf_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sf_equipment.Sf_subtype)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetCatEquipment(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cat_equipment")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cat_equipment

	for rows.Next() {
		var cat_equipment dto.Cat_equipment
		if err := rows.Scan(&cat_equipment.Catalogue_equip_id, &cat_equipment.Active_ind, &cat_equipment.Cat_equip_group, &cat_equipment.Cat_equip_type, &cat_equipment.Effective_date, &cat_equipment.Equipment_name, &cat_equipment.Expiry_date, &cat_equipment.Install_loc_type, &cat_equipment.Manufacturer_ba_id, &cat_equipment.Model_num, &cat_equipment.Ppdm_guid, &cat_equipment.Remark, &cat_equipment.Source, &cat_equipment.Row_changed_by, &cat_equipment.Row_changed_date, &cat_equipment.Row_created_by, &cat_equipment.Row_created_date, &cat_equipment.Row_effective_date, &cat_equipment.Row_expiry_date, &cat_equipment.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cat_equipment to result
		result = append(result, cat_equipment)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetCatEquipment(c *fiber.Ctx) error {
	var cat_equipment dto.Cat_equipment

	setDefaults(&cat_equipment)

	if err := json.Unmarshal(c.Body(), &cat_equipment); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cat_equipment values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	cat_equipment.Row_created_date = formatDate(cat_equipment.Row_created_date)
	cat_equipment.Row_changed_date = formatDate(cat_equipment.Row_changed_date)
	cat_equipment.Effective_date = formatDateString(cat_equipment.Effective_date)
	cat_equipment.Expiry_date = formatDateString(cat_equipment.Expiry_date)
	cat_equipment.Row_effective_date = formatDateString(cat_equipment.Row_effective_date)
	cat_equipment.Row_expiry_date = formatDateString(cat_equipment.Row_expiry_date)






	rows, err := stmt.Exec(cat_equipment.Catalogue_equip_id, cat_equipment.Active_ind, cat_equipment.Cat_equip_group, cat_equipment.Cat_equip_type, cat_equipment.Effective_date, cat_equipment.Equipment_name, cat_equipment.Expiry_date, cat_equipment.Install_loc_type, cat_equipment.Manufacturer_ba_id, cat_equipment.Model_num, cat_equipment.Ppdm_guid, cat_equipment.Remark, cat_equipment.Source, cat_equipment.Row_changed_by, cat_equipment.Row_changed_date, cat_equipment.Row_created_by, cat_equipment.Row_created_date, cat_equipment.Row_effective_date, cat_equipment.Row_expiry_date, cat_equipment.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateCatEquipment(c *fiber.Ctx) error {
	var cat_equipment dto.Cat_equipment

	setDefaults(&cat_equipment)

	if err := json.Unmarshal(c.Body(), &cat_equipment); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cat_equipment.Catalogue_equip_id = id

    if cat_equipment.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cat_equipment set active_ind = :1, cat_equip_group = :2, cat_equip_type = :3, effective_date = :4, equipment_name = :5, expiry_date = :6, install_loc_type = :7, manufacturer_ba_id = :8, model_num = :9, ppdm_guid = :10, remark = :11, source = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where catalogue_equip_id = :20")
	if err != nil {
		return err
	}

	cat_equipment.Row_changed_date = formatDate(cat_equipment.Row_changed_date)
	cat_equipment.Effective_date = formatDateString(cat_equipment.Effective_date)
	cat_equipment.Expiry_date = formatDateString(cat_equipment.Expiry_date)
	cat_equipment.Row_effective_date = formatDateString(cat_equipment.Row_effective_date)
	cat_equipment.Row_expiry_date = formatDateString(cat_equipment.Row_expiry_date)






	rows, err := stmt.Exec(cat_equipment.Active_ind, cat_equipment.Cat_equip_group, cat_equipment.Cat_equip_type, cat_equipment.Effective_date, cat_equipment.Equipment_name, cat_equipment.Expiry_date, cat_equipment.Install_loc_type, cat_equipment.Manufacturer_ba_id, cat_equipment.Model_num, cat_equipment.Ppdm_guid, cat_equipment.Remark, cat_equipment.Source, cat_equipment.Row_changed_by, cat_equipment.Row_changed_date, cat_equipment.Row_created_by, cat_equipment.Row_effective_date, cat_equipment.Row_expiry_date, cat_equipment.Row_quality, cat_equipment.Catalogue_equip_id)
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

func PatchCatEquipment(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cat_equipment set "
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

func DeleteCatEquipment(c *fiber.Ctx) error {
	id := c.Params("id")
	var cat_equipment dto.Cat_equipment
	cat_equipment.Catalogue_equip_id = id

	stmt, err := db.Prepare("delete from cat_equipment where catalogue_equip_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cat_equipment.Catalogue_equip_id)
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



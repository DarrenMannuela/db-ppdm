package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSfRigGenerator(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sf_rig_generator")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sf_rig_generator

	for rows.Next() {
		var sf_rig_generator dto.Sf_rig_generator
		if err := rows.Scan(&sf_rig_generator.Sf_subtype, &sf_rig_generator.Rig_id, &sf_rig_generator.Generator_id, &sf_rig_generator.Active_ind, &sf_rig_generator.Catalogue_equip_id, &sf_rig_generator.Effective_date, &sf_rig_generator.Equipment_id, &sf_rig_generator.Expiry_date, &sf_rig_generator.Generator_count, &sf_rig_generator.Generator_type, &sf_rig_generator.Input_type, &sf_rig_generator.Install_date, &sf_rig_generator.Power, &sf_rig_generator.Power_ouom, &sf_rig_generator.Power_rating, &sf_rig_generator.Power_rating_ouom, &sf_rig_generator.Ppdm_guid, &sf_rig_generator.Reference_num, &sf_rig_generator.Remark, &sf_rig_generator.Remove_date, &sf_rig_generator.Source, &sf_rig_generator.Row_changed_by, &sf_rig_generator.Row_changed_date, &sf_rig_generator.Row_created_by, &sf_rig_generator.Row_created_date, &sf_rig_generator.Row_effective_date, &sf_rig_generator.Row_expiry_date, &sf_rig_generator.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sf_rig_generator to result
		result = append(result, sf_rig_generator)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSfRigGenerator(c *fiber.Ctx) error {
	var sf_rig_generator dto.Sf_rig_generator

	setDefaults(&sf_rig_generator)

	if err := json.Unmarshal(c.Body(), &sf_rig_generator); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sf_rig_generator values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28)")
	if err != nil {
		return err
	}
	sf_rig_generator.Row_created_date = formatDate(sf_rig_generator.Row_created_date)
	sf_rig_generator.Row_changed_date = formatDate(sf_rig_generator.Row_changed_date)
	sf_rig_generator.Effective_date = formatDateString(sf_rig_generator.Effective_date)
	sf_rig_generator.Expiry_date = formatDateString(sf_rig_generator.Expiry_date)
	sf_rig_generator.Install_date = formatDateString(sf_rig_generator.Install_date)
	sf_rig_generator.Remove_date = formatDateString(sf_rig_generator.Remove_date)
	sf_rig_generator.Row_effective_date = formatDateString(sf_rig_generator.Row_effective_date)
	sf_rig_generator.Row_expiry_date = formatDateString(sf_rig_generator.Row_expiry_date)






	rows, err := stmt.Exec(sf_rig_generator.Sf_subtype, sf_rig_generator.Rig_id, sf_rig_generator.Generator_id, sf_rig_generator.Active_ind, sf_rig_generator.Catalogue_equip_id, sf_rig_generator.Effective_date, sf_rig_generator.Equipment_id, sf_rig_generator.Expiry_date, sf_rig_generator.Generator_count, sf_rig_generator.Generator_type, sf_rig_generator.Input_type, sf_rig_generator.Install_date, sf_rig_generator.Power, sf_rig_generator.Power_ouom, sf_rig_generator.Power_rating, sf_rig_generator.Power_rating_ouom, sf_rig_generator.Ppdm_guid, sf_rig_generator.Reference_num, sf_rig_generator.Remark, sf_rig_generator.Remove_date, sf_rig_generator.Source, sf_rig_generator.Row_changed_by, sf_rig_generator.Row_changed_date, sf_rig_generator.Row_created_by, sf_rig_generator.Row_created_date, sf_rig_generator.Row_effective_date, sf_rig_generator.Row_expiry_date, sf_rig_generator.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSfRigGenerator(c *fiber.Ctx) error {
	var sf_rig_generator dto.Sf_rig_generator

	setDefaults(&sf_rig_generator)

	if err := json.Unmarshal(c.Body(), &sf_rig_generator); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sf_rig_generator.Sf_subtype = id

    if sf_rig_generator.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sf_rig_generator set rig_id = :1, generator_id = :2, active_ind = :3, catalogue_equip_id = :4, effective_date = :5, equipment_id = :6, expiry_date = :7, generator_count = :8, generator_type = :9, input_type = :10, install_date = :11, power = :12, power_ouom = :13, power_rating = :14, power_rating_ouom = :15, ppdm_guid = :16, reference_num = :17, remark = :18, remove_date = :19, source = :20, row_changed_by = :21, row_changed_date = :22, row_created_by = :23, row_effective_date = :24, row_expiry_date = :25, row_quality = :26 where sf_subtype = :28")
	if err != nil {
		return err
	}

	sf_rig_generator.Row_changed_date = formatDate(sf_rig_generator.Row_changed_date)
	sf_rig_generator.Effective_date = formatDateString(sf_rig_generator.Effective_date)
	sf_rig_generator.Expiry_date = formatDateString(sf_rig_generator.Expiry_date)
	sf_rig_generator.Install_date = formatDateString(sf_rig_generator.Install_date)
	sf_rig_generator.Remove_date = formatDateString(sf_rig_generator.Remove_date)
	sf_rig_generator.Row_effective_date = formatDateString(sf_rig_generator.Row_effective_date)
	sf_rig_generator.Row_expiry_date = formatDateString(sf_rig_generator.Row_expiry_date)






	rows, err := stmt.Exec(sf_rig_generator.Rig_id, sf_rig_generator.Generator_id, sf_rig_generator.Active_ind, sf_rig_generator.Catalogue_equip_id, sf_rig_generator.Effective_date, sf_rig_generator.Equipment_id, sf_rig_generator.Expiry_date, sf_rig_generator.Generator_count, sf_rig_generator.Generator_type, sf_rig_generator.Input_type, sf_rig_generator.Install_date, sf_rig_generator.Power, sf_rig_generator.Power_ouom, sf_rig_generator.Power_rating, sf_rig_generator.Power_rating_ouom, sf_rig_generator.Ppdm_guid, sf_rig_generator.Reference_num, sf_rig_generator.Remark, sf_rig_generator.Remove_date, sf_rig_generator.Source, sf_rig_generator.Row_changed_by, sf_rig_generator.Row_changed_date, sf_rig_generator.Row_created_by, sf_rig_generator.Row_effective_date, sf_rig_generator.Row_expiry_date, sf_rig_generator.Row_quality, sf_rig_generator.Sf_subtype)
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

func PatchSfRigGenerator(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sf_rig_generator set "
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

func DeleteSfRigGenerator(c *fiber.Ctx) error {
	id := c.Params("id")
	var sf_rig_generator dto.Sf_rig_generator
	sf_rig_generator.Sf_subtype = id

	stmt, err := db.Prepare("delete from sf_rig_generator where sf_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sf_rig_generator.Sf_subtype)
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



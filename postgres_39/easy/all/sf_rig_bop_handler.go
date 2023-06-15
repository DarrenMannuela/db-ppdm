package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSfRigBop(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sf_rig_bop")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sf_rig_bop

	for rows.Next() {
		var sf_rig_bop dto.Sf_rig_bop
		if err := rows.Scan(&sf_rig_bop.Sf_subtype, &sf_rig_bop.Rig_id, &sf_rig_bop.Bop_id, &sf_rig_bop.Active_ind, &sf_rig_bop.Annular_count, &sf_rig_bop.Bop_count, &sf_rig_bop.Bop_diameter, &sf_rig_bop.Bop_diameter_ouom, &sf_rig_bop.Bop_nace_certified_ind, &sf_rig_bop.Bop_position_desc, &sf_rig_bop.Bop_pressure_rating, &sf_rig_bop.Bop_pressure_rating_ouom, &sf_rig_bop.Bop_type, &sf_rig_bop.Capacity, &sf_rig_bop.Capacity_ouom, &sf_rig_bop.Catalogue_equip_id, &sf_rig_bop.Double_count, &sf_rig_bop.Effective_date, &sf_rig_bop.Equipment_id, &sf_rig_bop.Expiry_date, &sf_rig_bop.Input_type, &sf_rig_bop.Install_date, &sf_rig_bop.Ppdm_guid, &sf_rig_bop.Reference_num, &sf_rig_bop.Remark, &sf_rig_bop.Remove_date, &sf_rig_bop.Single_count, &sf_rig_bop.Source, &sf_rig_bop.Row_changed_by, &sf_rig_bop.Row_changed_date, &sf_rig_bop.Row_created_by, &sf_rig_bop.Row_created_date, &sf_rig_bop.Row_effective_date, &sf_rig_bop.Row_expiry_date, &sf_rig_bop.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sf_rig_bop to result
		result = append(result, sf_rig_bop)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSfRigBop(c *fiber.Ctx) error {
	var sf_rig_bop dto.Sf_rig_bop

	setDefaults(&sf_rig_bop)

	if err := json.Unmarshal(c.Body(), &sf_rig_bop); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sf_rig_bop values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35)")
	if err != nil {
		return err
	}
	sf_rig_bop.Row_created_date = formatDate(sf_rig_bop.Row_created_date)
	sf_rig_bop.Row_changed_date = formatDate(sf_rig_bop.Row_changed_date)
	sf_rig_bop.Effective_date = formatDateString(sf_rig_bop.Effective_date)
	sf_rig_bop.Expiry_date = formatDateString(sf_rig_bop.Expiry_date)
	sf_rig_bop.Install_date = formatDateString(sf_rig_bop.Install_date)
	sf_rig_bop.Remove_date = formatDateString(sf_rig_bop.Remove_date)
	sf_rig_bop.Row_effective_date = formatDateString(sf_rig_bop.Row_effective_date)
	sf_rig_bop.Row_expiry_date = formatDateString(sf_rig_bop.Row_expiry_date)






	rows, err := stmt.Exec(sf_rig_bop.Sf_subtype, sf_rig_bop.Rig_id, sf_rig_bop.Bop_id, sf_rig_bop.Active_ind, sf_rig_bop.Annular_count, sf_rig_bop.Bop_count, sf_rig_bop.Bop_diameter, sf_rig_bop.Bop_diameter_ouom, sf_rig_bop.Bop_nace_certified_ind, sf_rig_bop.Bop_position_desc, sf_rig_bop.Bop_pressure_rating, sf_rig_bop.Bop_pressure_rating_ouom, sf_rig_bop.Bop_type, sf_rig_bop.Capacity, sf_rig_bop.Capacity_ouom, sf_rig_bop.Catalogue_equip_id, sf_rig_bop.Double_count, sf_rig_bop.Effective_date, sf_rig_bop.Equipment_id, sf_rig_bop.Expiry_date, sf_rig_bop.Input_type, sf_rig_bop.Install_date, sf_rig_bop.Ppdm_guid, sf_rig_bop.Reference_num, sf_rig_bop.Remark, sf_rig_bop.Remove_date, sf_rig_bop.Single_count, sf_rig_bop.Source, sf_rig_bop.Row_changed_by, sf_rig_bop.Row_changed_date, sf_rig_bop.Row_created_by, sf_rig_bop.Row_created_date, sf_rig_bop.Row_effective_date, sf_rig_bop.Row_expiry_date, sf_rig_bop.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSfRigBop(c *fiber.Ctx) error {
	var sf_rig_bop dto.Sf_rig_bop

	setDefaults(&sf_rig_bop)

	if err := json.Unmarshal(c.Body(), &sf_rig_bop); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sf_rig_bop.Sf_subtype = id

    if sf_rig_bop.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sf_rig_bop set rig_id = :1, bop_id = :2, active_ind = :3, annular_count = :4, bop_count = :5, bop_diameter = :6, bop_diameter_ouom = :7, bop_nace_certified_ind = :8, bop_position_desc = :9, bop_pressure_rating = :10, bop_pressure_rating_ouom = :11, bop_type = :12, capacity = :13, capacity_ouom = :14, catalogue_equip_id = :15, double_count = :16, effective_date = :17, equipment_id = :18, expiry_date = :19, input_type = :20, install_date = :21, ppdm_guid = :22, reference_num = :23, remark = :24, remove_date = :25, single_count = :26, source = :27, row_changed_by = :28, row_changed_date = :29, row_created_by = :30, row_effective_date = :31, row_expiry_date = :32, row_quality = :33 where sf_subtype = :35")
	if err != nil {
		return err
	}

	sf_rig_bop.Row_changed_date = formatDate(sf_rig_bop.Row_changed_date)
	sf_rig_bop.Effective_date = formatDateString(sf_rig_bop.Effective_date)
	sf_rig_bop.Expiry_date = formatDateString(sf_rig_bop.Expiry_date)
	sf_rig_bop.Install_date = formatDateString(sf_rig_bop.Install_date)
	sf_rig_bop.Remove_date = formatDateString(sf_rig_bop.Remove_date)
	sf_rig_bop.Row_effective_date = formatDateString(sf_rig_bop.Row_effective_date)
	sf_rig_bop.Row_expiry_date = formatDateString(sf_rig_bop.Row_expiry_date)






	rows, err := stmt.Exec(sf_rig_bop.Rig_id, sf_rig_bop.Bop_id, sf_rig_bop.Active_ind, sf_rig_bop.Annular_count, sf_rig_bop.Bop_count, sf_rig_bop.Bop_diameter, sf_rig_bop.Bop_diameter_ouom, sf_rig_bop.Bop_nace_certified_ind, sf_rig_bop.Bop_position_desc, sf_rig_bop.Bop_pressure_rating, sf_rig_bop.Bop_pressure_rating_ouom, sf_rig_bop.Bop_type, sf_rig_bop.Capacity, sf_rig_bop.Capacity_ouom, sf_rig_bop.Catalogue_equip_id, sf_rig_bop.Double_count, sf_rig_bop.Effective_date, sf_rig_bop.Equipment_id, sf_rig_bop.Expiry_date, sf_rig_bop.Input_type, sf_rig_bop.Install_date, sf_rig_bop.Ppdm_guid, sf_rig_bop.Reference_num, sf_rig_bop.Remark, sf_rig_bop.Remove_date, sf_rig_bop.Single_count, sf_rig_bop.Source, sf_rig_bop.Row_changed_by, sf_rig_bop.Row_changed_date, sf_rig_bop.Row_created_by, sf_rig_bop.Row_effective_date, sf_rig_bop.Row_expiry_date, sf_rig_bop.Row_quality, sf_rig_bop.Sf_subtype)
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

func PatchSfRigBop(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sf_rig_bop set "
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

func DeleteSfRigBop(c *fiber.Ctx) error {
	id := c.Params("id")
	var sf_rig_bop dto.Sf_rig_bop
	sf_rig_bop.Sf_subtype = id

	stmt, err := db.Prepare("delete from sf_rig_bop where sf_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sf_rig_bop.Sf_subtype)
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



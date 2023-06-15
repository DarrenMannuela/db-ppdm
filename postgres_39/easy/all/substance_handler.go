package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSubstance(c *fiber.Ctx) error {
	rows, err := db.Query("select * from substance")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Substance

	for rows.Next() {
		var substance dto.Substance
		if err := rows.Scan(&substance.Substance_id, &substance.Active_ind, &substance.Api_gravity_max, &substance.Api_gravity_min, &substance.Atomic_mass, &substance.Atomic_mass_number, &substance.Atomic_weight, &substance.Carbon_count, &substance.Carbon_count_max, &substance.Carbon_count_min, &substance.Composition_formula, &substance.Conversion_quantity, &substance.Effective_date, &substance.Expiry_date, &substance.Ionic_charge, &substance.Molecular_mass, &substance.M_z_ion, &substance.Phase_type, &substance.Ph_value, &substance.Ppdm_guid, &substance.Preferred_ind, &substance.Preferred_long_name, &substance.Preferred_short_name, &substance.Preferred_uom_id, &substance.Property_set_id, &substance.Remark, &substance.Source, &substance.Source_document_id, &substance.Substance_definition, &substance.Valence_value, &substance.Row_changed_by, &substance.Row_changed_date, &substance.Row_created_by, &substance.Row_created_date, &substance.Row_effective_date, &substance.Row_expiry_date, &substance.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append substance to result
		result = append(result, substance)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSubstance(c *fiber.Ctx) error {
	var substance dto.Substance

	setDefaults(&substance)

	if err := json.Unmarshal(c.Body(), &substance); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into substance values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37)")
	if err != nil {
		return err
	}
	substance.Row_created_date = formatDate(substance.Row_created_date)
	substance.Row_changed_date = formatDate(substance.Row_changed_date)
	substance.Effective_date = formatDateString(substance.Effective_date)
	substance.Expiry_date = formatDateString(substance.Expiry_date)
	substance.Row_effective_date = formatDateString(substance.Row_effective_date)
	substance.Row_expiry_date = formatDateString(substance.Row_expiry_date)






	rows, err := stmt.Exec(substance.Substance_id, substance.Active_ind, substance.Api_gravity_max, substance.Api_gravity_min, substance.Atomic_mass, substance.Atomic_mass_number, substance.Atomic_weight, substance.Carbon_count, substance.Carbon_count_max, substance.Carbon_count_min, substance.Composition_formula, substance.Conversion_quantity, substance.Effective_date, substance.Expiry_date, substance.Ionic_charge, substance.Molecular_mass, substance.M_z_ion, substance.Phase_type, substance.Ph_value, substance.Ppdm_guid, substance.Preferred_ind, substance.Preferred_long_name, substance.Preferred_short_name, substance.Preferred_uom_id, substance.Property_set_id, substance.Remark, substance.Source, substance.Source_document_id, substance.Substance_definition, substance.Valence_value, substance.Row_changed_by, substance.Row_changed_date, substance.Row_created_by, substance.Row_created_date, substance.Row_effective_date, substance.Row_expiry_date, substance.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSubstance(c *fiber.Ctx) error {
	var substance dto.Substance

	setDefaults(&substance)

	if err := json.Unmarshal(c.Body(), &substance); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	substance.Substance_id = id

    if substance.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update substance set active_ind = :1, api_gravity_max = :2, api_gravity_min = :3, atomic_mass = :4, atomic_mass_number = :5, atomic_weight = :6, carbon_count = :7, carbon_count_max = :8, carbon_count_min = :9, composition_formula = :10, conversion_quantity = :11, effective_date = :12, expiry_date = :13, ionic_charge = :14, molecular_mass = :15, m_z_ion = :16, phase_type = :17, ph_value = :18, ppdm_guid = :19, preferred_ind = :20, preferred_long_name = :21, preferred_short_name = :22, preferred_uom_id = :23, property_set_id = :24, remark = :25, source = :26, source_document_id = :27, substance_definition = :28, valence_value = :29, row_changed_by = :30, row_changed_date = :31, row_created_by = :32, row_effective_date = :33, row_expiry_date = :34, row_quality = :35 where substance_id = :37")
	if err != nil {
		return err
	}

	substance.Row_changed_date = formatDate(substance.Row_changed_date)
	substance.Effective_date = formatDateString(substance.Effective_date)
	substance.Expiry_date = formatDateString(substance.Expiry_date)
	substance.Row_effective_date = formatDateString(substance.Row_effective_date)
	substance.Row_expiry_date = formatDateString(substance.Row_expiry_date)






	rows, err := stmt.Exec(substance.Active_ind, substance.Api_gravity_max, substance.Api_gravity_min, substance.Atomic_mass, substance.Atomic_mass_number, substance.Atomic_weight, substance.Carbon_count, substance.Carbon_count_max, substance.Carbon_count_min, substance.Composition_formula, substance.Conversion_quantity, substance.Effective_date, substance.Expiry_date, substance.Ionic_charge, substance.Molecular_mass, substance.M_z_ion, substance.Phase_type, substance.Ph_value, substance.Ppdm_guid, substance.Preferred_ind, substance.Preferred_long_name, substance.Preferred_short_name, substance.Preferred_uom_id, substance.Property_set_id, substance.Remark, substance.Source, substance.Source_document_id, substance.Substance_definition, substance.Valence_value, substance.Row_changed_by, substance.Row_changed_date, substance.Row_created_by, substance.Row_effective_date, substance.Row_expiry_date, substance.Row_quality, substance.Substance_id)
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

func PatchSubstance(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update substance set "
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
	query += " where substance_id = :id"

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

func DeleteSubstance(c *fiber.Ctx) error {
	id := c.Params("id")
	var substance dto.Substance
	substance.Substance_id = id

	stmt, err := db.Prepare("delete from substance where substance_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(substance.Substance_id)
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



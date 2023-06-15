package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSubstanceBehavior(c *fiber.Ctx) error {
	rows, err := db.Query("select * from substance_behavior")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Substance_behavior

	for rows.Next() {
		var substance_behavior dto.Substance_behavior
		if err := rows.Scan(&substance_behavior.Substance_id, &substance_behavior.Active_ind, &substance_behavior.Composite_ind, &substance_behavior.Drilling_mud_ind, &substance_behavior.Effective_date, &substance_behavior.Element_ind, &substance_behavior.Expiry_date, &substance_behavior.Hydrocarbon_ind, &substance_behavior.Ion_ind, &substance_behavior.Isomer_ind, &substance_behavior.Isotope_ind, &substance_behavior.Mineral_ind, &substance_behavior.Organic_matter_ind, &substance_behavior.Ppdm_guid, &substance_behavior.Preferred_ind, &substance_behavior.Preferred_long_name, &substance_behavior.Preferred_short_name, &substance_behavior.Production_stuff_ind, &substance_behavior.Property_set_id, &substance_behavior.Remark, &substance_behavior.Solvent_ind, &substance_behavior.Source, &substance_behavior.Source_document_id, &substance_behavior.Row_changed_by, &substance_behavior.Row_changed_date, &substance_behavior.Row_created_by, &substance_behavior.Row_created_date, &substance_behavior.Row_effective_date, &substance_behavior.Row_expiry_date, &substance_behavior.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append substance_behavior to result
		result = append(result, substance_behavior)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSubstanceBehavior(c *fiber.Ctx) error {
	var substance_behavior dto.Substance_behavior

	setDefaults(&substance_behavior)

	if err := json.Unmarshal(c.Body(), &substance_behavior); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into substance_behavior values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30)")
	if err != nil {
		return err
	}
	substance_behavior.Row_created_date = formatDate(substance_behavior.Row_created_date)
	substance_behavior.Row_changed_date = formatDate(substance_behavior.Row_changed_date)
	substance_behavior.Effective_date = formatDateString(substance_behavior.Effective_date)
	substance_behavior.Expiry_date = formatDateString(substance_behavior.Expiry_date)
	substance_behavior.Row_effective_date = formatDateString(substance_behavior.Row_effective_date)
	substance_behavior.Row_expiry_date = formatDateString(substance_behavior.Row_expiry_date)






	rows, err := stmt.Exec(substance_behavior.Substance_id, substance_behavior.Active_ind, substance_behavior.Composite_ind, substance_behavior.Drilling_mud_ind, substance_behavior.Effective_date, substance_behavior.Element_ind, substance_behavior.Expiry_date, substance_behavior.Hydrocarbon_ind, substance_behavior.Ion_ind, substance_behavior.Isomer_ind, substance_behavior.Isotope_ind, substance_behavior.Mineral_ind, substance_behavior.Organic_matter_ind, substance_behavior.Ppdm_guid, substance_behavior.Preferred_ind, substance_behavior.Preferred_long_name, substance_behavior.Preferred_short_name, substance_behavior.Production_stuff_ind, substance_behavior.Property_set_id, substance_behavior.Remark, substance_behavior.Solvent_ind, substance_behavior.Source, substance_behavior.Source_document_id, substance_behavior.Row_changed_by, substance_behavior.Row_changed_date, substance_behavior.Row_created_by, substance_behavior.Row_created_date, substance_behavior.Row_effective_date, substance_behavior.Row_expiry_date, substance_behavior.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSubstanceBehavior(c *fiber.Ctx) error {
	var substance_behavior dto.Substance_behavior

	setDefaults(&substance_behavior)

	if err := json.Unmarshal(c.Body(), &substance_behavior); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	substance_behavior.Substance_id = id

    if substance_behavior.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update substance_behavior set active_ind = :1, composite_ind = :2, drilling_mud_ind = :3, effective_date = :4, element_ind = :5, expiry_date = :6, hydrocarbon_ind = :7, ion_ind = :8, isomer_ind = :9, isotope_ind = :10, mineral_ind = :11, organic_matter_ind = :12, ppdm_guid = :13, preferred_ind = :14, preferred_long_name = :15, preferred_short_name = :16, production_stuff_ind = :17, property_set_id = :18, remark = :19, solvent_ind = :20, source = :21, source_document_id = :22, row_changed_by = :23, row_changed_date = :24, row_created_by = :25, row_effective_date = :26, row_expiry_date = :27, row_quality = :28 where substance_id = :30")
	if err != nil {
		return err
	}

	substance_behavior.Row_changed_date = formatDate(substance_behavior.Row_changed_date)
	substance_behavior.Effective_date = formatDateString(substance_behavior.Effective_date)
	substance_behavior.Expiry_date = formatDateString(substance_behavior.Expiry_date)
	substance_behavior.Row_effective_date = formatDateString(substance_behavior.Row_effective_date)
	substance_behavior.Row_expiry_date = formatDateString(substance_behavior.Row_expiry_date)






	rows, err := stmt.Exec(substance_behavior.Active_ind, substance_behavior.Composite_ind, substance_behavior.Drilling_mud_ind, substance_behavior.Effective_date, substance_behavior.Element_ind, substance_behavior.Expiry_date, substance_behavior.Hydrocarbon_ind, substance_behavior.Ion_ind, substance_behavior.Isomer_ind, substance_behavior.Isotope_ind, substance_behavior.Mineral_ind, substance_behavior.Organic_matter_ind, substance_behavior.Ppdm_guid, substance_behavior.Preferred_ind, substance_behavior.Preferred_long_name, substance_behavior.Preferred_short_name, substance_behavior.Production_stuff_ind, substance_behavior.Property_set_id, substance_behavior.Remark, substance_behavior.Solvent_ind, substance_behavior.Source, substance_behavior.Source_document_id, substance_behavior.Row_changed_by, substance_behavior.Row_changed_date, substance_behavior.Row_created_by, substance_behavior.Row_effective_date, substance_behavior.Row_expiry_date, substance_behavior.Row_quality, substance_behavior.Substance_id)
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

func PatchSubstanceBehavior(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update substance_behavior set "
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

func DeleteSubstanceBehavior(c *fiber.Ctx) error {
	id := c.Params("id")
	var substance_behavior dto.Substance_behavior
	substance_behavior.Substance_id = id

	stmt, err := db.Prepare("delete from substance_behavior where substance_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(substance_behavior.Substance_id)
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



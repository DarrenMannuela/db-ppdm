package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSubstanceComposition(c *fiber.Ctx) error {
	rows, err := db.Query("select * from substance_composition")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Substance_composition

	for rows.Next() {
		var substance_composition dto.Substance_composition
		if err := rows.Scan(&substance_composition.Substance_id, &substance_composition.Sub_substance_id, &substance_composition.Composition_obs_no, &substance_composition.Active_ind, &substance_composition.Defined_by_ba_id, &substance_composition.Effective_date, &substance_composition.Expiry_date, &substance_composition.Formula, &substance_composition.Ppdm_guid, &substance_composition.Remark, &substance_composition.Source, &substance_composition.Source_document_id, &substance_composition.Substance_component_type, &substance_composition.Row_changed_by, &substance_composition.Row_changed_date, &substance_composition.Row_created_by, &substance_composition.Row_created_date, &substance_composition.Row_effective_date, &substance_composition.Row_expiry_date, &substance_composition.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append substance_composition to result
		result = append(result, substance_composition)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSubstanceComposition(c *fiber.Ctx) error {
	var substance_composition dto.Substance_composition

	setDefaults(&substance_composition)

	if err := json.Unmarshal(c.Body(), &substance_composition); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into substance_composition values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	substance_composition.Row_created_date = formatDate(substance_composition.Row_created_date)
	substance_composition.Row_changed_date = formatDate(substance_composition.Row_changed_date)
	substance_composition.Effective_date = formatDateString(substance_composition.Effective_date)
	substance_composition.Expiry_date = formatDateString(substance_composition.Expiry_date)
	substance_composition.Row_effective_date = formatDateString(substance_composition.Row_effective_date)
	substance_composition.Row_expiry_date = formatDateString(substance_composition.Row_expiry_date)






	rows, err := stmt.Exec(substance_composition.Substance_id, substance_composition.Sub_substance_id, substance_composition.Composition_obs_no, substance_composition.Active_ind, substance_composition.Defined_by_ba_id, substance_composition.Effective_date, substance_composition.Expiry_date, substance_composition.Formula, substance_composition.Ppdm_guid, substance_composition.Remark, substance_composition.Source, substance_composition.Source_document_id, substance_composition.Substance_component_type, substance_composition.Row_changed_by, substance_composition.Row_changed_date, substance_composition.Row_created_by, substance_composition.Row_created_date, substance_composition.Row_effective_date, substance_composition.Row_expiry_date, substance_composition.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSubstanceComposition(c *fiber.Ctx) error {
	var substance_composition dto.Substance_composition

	setDefaults(&substance_composition)

	if err := json.Unmarshal(c.Body(), &substance_composition); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	substance_composition.Substance_id = id

    if substance_composition.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update substance_composition set sub_substance_id = :1, composition_obs_no = :2, active_ind = :3, defined_by_ba_id = :4, effective_date = :5, expiry_date = :6, formula = :7, ppdm_guid = :8, remark = :9, source = :10, source_document_id = :11, substance_component_type = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where substance_id = :20")
	if err != nil {
		return err
	}

	substance_composition.Row_changed_date = formatDate(substance_composition.Row_changed_date)
	substance_composition.Effective_date = formatDateString(substance_composition.Effective_date)
	substance_composition.Expiry_date = formatDateString(substance_composition.Expiry_date)
	substance_composition.Row_effective_date = formatDateString(substance_composition.Row_effective_date)
	substance_composition.Row_expiry_date = formatDateString(substance_composition.Row_expiry_date)






	rows, err := stmt.Exec(substance_composition.Sub_substance_id, substance_composition.Composition_obs_no, substance_composition.Active_ind, substance_composition.Defined_by_ba_id, substance_composition.Effective_date, substance_composition.Expiry_date, substance_composition.Formula, substance_composition.Ppdm_guid, substance_composition.Remark, substance_composition.Source, substance_composition.Source_document_id, substance_composition.Substance_component_type, substance_composition.Row_changed_by, substance_composition.Row_changed_date, substance_composition.Row_created_by, substance_composition.Row_effective_date, substance_composition.Row_expiry_date, substance_composition.Row_quality, substance_composition.Substance_id)
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

func PatchSubstanceComposition(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update substance_composition set "
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

func DeleteSubstanceComposition(c *fiber.Ctx) error {
	id := c.Params("id")
	var substance_composition dto.Substance_composition
	substance_composition.Substance_id = id

	stmt, err := db.Prepare("delete from substance_composition where substance_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(substance_composition.Substance_id)
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



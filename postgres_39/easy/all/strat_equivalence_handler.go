package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetStratEquivalence(c *fiber.Ctx) error {
	rows, err := db.Query("select * from strat_equivalence")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Strat_equivalence

	for rows.Next() {
		var strat_equivalence dto.Strat_equivalence
		if err := rows.Scan(&strat_equivalence.Strat_name_set_id, &strat_equivalence.Strat_unit_id, &strat_equivalence.Equiv_strat_name_set_id, &strat_equivalence.Equiv_strat_unit_id, &strat_equivalence.Equiv_id, &strat_equivalence.Source, &strat_equivalence.Active_ind, &strat_equivalence.Area_id, &strat_equivalence.Area_id2, &strat_equivalence.Area_type, &strat_equivalence.Area_type2, &strat_equivalence.Effective_date, &strat_equivalence.Equiv_direction, &strat_equivalence.Equiv_interp_id, &strat_equivalence.Equiv_strat_column_id, &strat_equivalence.Equiv_strat_column_source, &strat_equivalence.Expiry_date, &strat_equivalence.Ppdm_guid, &strat_equivalence.Remark, &strat_equivalence.Source_document_id, &strat_equivalence.Strat_column_id, &strat_equivalence.Strat_column_source, &strat_equivalence.Strat_equivalence_type, &strat_equivalence.Strat_interp_id, &strat_equivalence.Row_changed_by, &strat_equivalence.Row_changed_date, &strat_equivalence.Row_created_by, &strat_equivalence.Row_created_date, &strat_equivalence.Row_effective_date, &strat_equivalence.Row_expiry_date, &strat_equivalence.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append strat_equivalence to result
		result = append(result, strat_equivalence)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetStratEquivalence(c *fiber.Ctx) error {
	var strat_equivalence dto.Strat_equivalence

	setDefaults(&strat_equivalence)

	if err := json.Unmarshal(c.Body(), &strat_equivalence); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into strat_equivalence values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31)")
	if err != nil {
		return err
	}
	strat_equivalence.Row_created_date = formatDate(strat_equivalence.Row_created_date)
	strat_equivalence.Row_changed_date = formatDate(strat_equivalence.Row_changed_date)
	strat_equivalence.Effective_date = formatDateString(strat_equivalence.Effective_date)
	strat_equivalence.Expiry_date = formatDateString(strat_equivalence.Expiry_date)
	strat_equivalence.Row_effective_date = formatDateString(strat_equivalence.Row_effective_date)
	strat_equivalence.Row_expiry_date = formatDateString(strat_equivalence.Row_expiry_date)






	rows, err := stmt.Exec(strat_equivalence.Strat_name_set_id, strat_equivalence.Strat_unit_id, strat_equivalence.Equiv_strat_name_set_id, strat_equivalence.Equiv_strat_unit_id, strat_equivalence.Equiv_id, strat_equivalence.Source, strat_equivalence.Active_ind, strat_equivalence.Area_id, strat_equivalence.Area_id2, strat_equivalence.Area_type, strat_equivalence.Area_type2, strat_equivalence.Effective_date, strat_equivalence.Equiv_direction, strat_equivalence.Equiv_interp_id, strat_equivalence.Equiv_strat_column_id, strat_equivalence.Equiv_strat_column_source, strat_equivalence.Expiry_date, strat_equivalence.Ppdm_guid, strat_equivalence.Remark, strat_equivalence.Source_document_id, strat_equivalence.Strat_column_id, strat_equivalence.Strat_column_source, strat_equivalence.Strat_equivalence_type, strat_equivalence.Strat_interp_id, strat_equivalence.Row_changed_by, strat_equivalence.Row_changed_date, strat_equivalence.Row_created_by, strat_equivalence.Row_created_date, strat_equivalence.Row_effective_date, strat_equivalence.Row_expiry_date, strat_equivalence.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateStratEquivalence(c *fiber.Ctx) error {
	var strat_equivalence dto.Strat_equivalence

	setDefaults(&strat_equivalence)

	if err := json.Unmarshal(c.Body(), &strat_equivalence); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	strat_equivalence.Strat_name_set_id = id

    if strat_equivalence.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update strat_equivalence set strat_unit_id = :1, equiv_strat_name_set_id = :2, equiv_strat_unit_id = :3, equiv_id = :4, source = :5, active_ind = :6, area_id = :7, area_id2 = :8, area_type = :9, area_type2 = :10, effective_date = :11, equiv_direction = :12, equiv_interp_id = :13, equiv_strat_column_id = :14, equiv_strat_column_source = :15, expiry_date = :16, ppdm_guid = :17, remark = :18, source_document_id = :19, strat_column_id = :20, strat_column_source = :21, strat_equivalence_type = :22, strat_interp_id = :23, row_changed_by = :24, row_changed_date = :25, row_created_by = :26, row_effective_date = :27, row_expiry_date = :28, row_quality = :29 where strat_name_set_id = :31")
	if err != nil {
		return err
	}

	strat_equivalence.Row_changed_date = formatDate(strat_equivalence.Row_changed_date)
	strat_equivalence.Effective_date = formatDateString(strat_equivalence.Effective_date)
	strat_equivalence.Expiry_date = formatDateString(strat_equivalence.Expiry_date)
	strat_equivalence.Row_effective_date = formatDateString(strat_equivalence.Row_effective_date)
	strat_equivalence.Row_expiry_date = formatDateString(strat_equivalence.Row_expiry_date)






	rows, err := stmt.Exec(strat_equivalence.Strat_unit_id, strat_equivalence.Equiv_strat_name_set_id, strat_equivalence.Equiv_strat_unit_id, strat_equivalence.Equiv_id, strat_equivalence.Source, strat_equivalence.Active_ind, strat_equivalence.Area_id, strat_equivalence.Area_id2, strat_equivalence.Area_type, strat_equivalence.Area_type2, strat_equivalence.Effective_date, strat_equivalence.Equiv_direction, strat_equivalence.Equiv_interp_id, strat_equivalence.Equiv_strat_column_id, strat_equivalence.Equiv_strat_column_source, strat_equivalence.Expiry_date, strat_equivalence.Ppdm_guid, strat_equivalence.Remark, strat_equivalence.Source_document_id, strat_equivalence.Strat_column_id, strat_equivalence.Strat_column_source, strat_equivalence.Strat_equivalence_type, strat_equivalence.Strat_interp_id, strat_equivalence.Row_changed_by, strat_equivalence.Row_changed_date, strat_equivalence.Row_created_by, strat_equivalence.Row_effective_date, strat_equivalence.Row_expiry_date, strat_equivalence.Row_quality, strat_equivalence.Strat_name_set_id)
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

func PatchStratEquivalence(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update strat_equivalence set "
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
	query += " where strat_name_set_id = :id"

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

func DeleteStratEquivalence(c *fiber.Ctx) error {
	id := c.Params("id")
	var strat_equivalence dto.Strat_equivalence
	strat_equivalence.Strat_name_set_id = id

	stmt, err := db.Prepare("delete from strat_equivalence where strat_name_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(strat_equivalence.Strat_name_set_id)
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



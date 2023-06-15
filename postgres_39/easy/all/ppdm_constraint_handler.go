package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmConstraint(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_constraint")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_constraint

	for rows.Next() {
		var ppdm_constraint dto.Ppdm_constraint
		if err := rows.Scan(&ppdm_constraint.System_id, &ppdm_constraint.Table_name, &ppdm_constraint.Constraint_name, &ppdm_constraint.Active_ind, &ppdm_constraint.Constraint_full_name, &ppdm_constraint.Constraint_type, &ppdm_constraint.Effective_date, &ppdm_constraint.Expiry_date, &ppdm_constraint.Extension_ind, &ppdm_constraint.Ppdm_guid, &ppdm_constraint.Referenced_constraint_name, &ppdm_constraint.Referenced_system_id, &ppdm_constraint.Referenced_table_name, &ppdm_constraint.Remark, &ppdm_constraint.Source, &ppdm_constraint.Row_changed_by, &ppdm_constraint.Row_changed_date, &ppdm_constraint.Row_created_by, &ppdm_constraint.Row_created_date, &ppdm_constraint.Row_effective_date, &ppdm_constraint.Row_expiry_date, &ppdm_constraint.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_constraint to result
		result = append(result, ppdm_constraint)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmConstraint(c *fiber.Ctx) error {
	var ppdm_constraint dto.Ppdm_constraint

	setDefaults(&ppdm_constraint)

	if err := json.Unmarshal(c.Body(), &ppdm_constraint); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_constraint values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	ppdm_constraint.Row_created_date = formatDate(ppdm_constraint.Row_created_date)
	ppdm_constraint.Row_changed_date = formatDate(ppdm_constraint.Row_changed_date)
	ppdm_constraint.Effective_date = formatDateString(ppdm_constraint.Effective_date)
	ppdm_constraint.Expiry_date = formatDateString(ppdm_constraint.Expiry_date)
	ppdm_constraint.Row_effective_date = formatDateString(ppdm_constraint.Row_effective_date)
	ppdm_constraint.Row_expiry_date = formatDateString(ppdm_constraint.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_constraint.System_id, ppdm_constraint.Table_name, ppdm_constraint.Constraint_name, ppdm_constraint.Active_ind, ppdm_constraint.Constraint_full_name, ppdm_constraint.Constraint_type, ppdm_constraint.Effective_date, ppdm_constraint.Expiry_date, ppdm_constraint.Extension_ind, ppdm_constraint.Ppdm_guid, ppdm_constraint.Referenced_constraint_name, ppdm_constraint.Referenced_system_id, ppdm_constraint.Referenced_table_name, ppdm_constraint.Remark, ppdm_constraint.Source, ppdm_constraint.Row_changed_by, ppdm_constraint.Row_changed_date, ppdm_constraint.Row_created_by, ppdm_constraint.Row_created_date, ppdm_constraint.Row_effective_date, ppdm_constraint.Row_expiry_date, ppdm_constraint.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmConstraint(c *fiber.Ctx) error {
	var ppdm_constraint dto.Ppdm_constraint

	setDefaults(&ppdm_constraint)

	if err := json.Unmarshal(c.Body(), &ppdm_constraint); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_constraint.System_id = id

    if ppdm_constraint.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_constraint set table_name = :1, constraint_name = :2, active_ind = :3, constraint_full_name = :4, constraint_type = :5, effective_date = :6, expiry_date = :7, extension_ind = :8, ppdm_guid = :9, referenced_constraint_name = :10, referenced_system_id = :11, referenced_table_name = :12, remark = :13, source = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where system_id = :22")
	if err != nil {
		return err
	}

	ppdm_constraint.Row_changed_date = formatDate(ppdm_constraint.Row_changed_date)
	ppdm_constraint.Effective_date = formatDateString(ppdm_constraint.Effective_date)
	ppdm_constraint.Expiry_date = formatDateString(ppdm_constraint.Expiry_date)
	ppdm_constraint.Row_effective_date = formatDateString(ppdm_constraint.Row_effective_date)
	ppdm_constraint.Row_expiry_date = formatDateString(ppdm_constraint.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_constraint.Table_name, ppdm_constraint.Constraint_name, ppdm_constraint.Active_ind, ppdm_constraint.Constraint_full_name, ppdm_constraint.Constraint_type, ppdm_constraint.Effective_date, ppdm_constraint.Expiry_date, ppdm_constraint.Extension_ind, ppdm_constraint.Ppdm_guid, ppdm_constraint.Referenced_constraint_name, ppdm_constraint.Referenced_system_id, ppdm_constraint.Referenced_table_name, ppdm_constraint.Remark, ppdm_constraint.Source, ppdm_constraint.Row_changed_by, ppdm_constraint.Row_changed_date, ppdm_constraint.Row_created_by, ppdm_constraint.Row_effective_date, ppdm_constraint.Row_expiry_date, ppdm_constraint.Row_quality, ppdm_constraint.System_id)
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

func PatchPpdmConstraint(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_constraint set "
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
	query += " where system_id = :id"

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

func DeletePpdmConstraint(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_constraint dto.Ppdm_constraint
	ppdm_constraint.System_id = id

	stmt, err := db.Prepare("delete from ppdm_constraint where system_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_constraint.System_id)
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



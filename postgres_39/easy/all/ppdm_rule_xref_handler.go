package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmRuleXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_rule_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_rule_xref

	for rows.Next() {
		var ppdm_rule_xref dto.Ppdm_rule_xref
		if err := rows.Scan(&ppdm_rule_xref.Dependency_id, &ppdm_rule_xref.Rule_id, &ppdm_rule_xref.Rule_id2, &ppdm_rule_xref.Xref_obs_no, &ppdm_rule_xref.Active_ind, &ppdm_rule_xref.Description, &ppdm_rule_xref.Effective_date, &ppdm_rule_xref.Expiry_date, &ppdm_rule_xref.Ppdm_guid, &ppdm_rule_xref.Remark, &ppdm_rule_xref.Rule_xref_condition, &ppdm_rule_xref.Rule_xref_type, &ppdm_rule_xref.Source, &ppdm_rule_xref.Row_changed_by, &ppdm_rule_xref.Row_changed_date, &ppdm_rule_xref.Row_created_by, &ppdm_rule_xref.Row_created_date, &ppdm_rule_xref.Row_effective_date, &ppdm_rule_xref.Row_expiry_date, &ppdm_rule_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_rule_xref to result
		result = append(result, ppdm_rule_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmRuleXref(c *fiber.Ctx) error {
	var ppdm_rule_xref dto.Ppdm_rule_xref

	setDefaults(&ppdm_rule_xref)

	if err := json.Unmarshal(c.Body(), &ppdm_rule_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_rule_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	ppdm_rule_xref.Row_created_date = formatDate(ppdm_rule_xref.Row_created_date)
	ppdm_rule_xref.Row_changed_date = formatDate(ppdm_rule_xref.Row_changed_date)
	ppdm_rule_xref.Effective_date = formatDateString(ppdm_rule_xref.Effective_date)
	ppdm_rule_xref.Expiry_date = formatDateString(ppdm_rule_xref.Expiry_date)
	ppdm_rule_xref.Row_effective_date = formatDateString(ppdm_rule_xref.Row_effective_date)
	ppdm_rule_xref.Row_expiry_date = formatDateString(ppdm_rule_xref.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_rule_xref.Dependency_id, ppdm_rule_xref.Rule_id, ppdm_rule_xref.Rule_id2, ppdm_rule_xref.Xref_obs_no, ppdm_rule_xref.Active_ind, ppdm_rule_xref.Description, ppdm_rule_xref.Effective_date, ppdm_rule_xref.Expiry_date, ppdm_rule_xref.Ppdm_guid, ppdm_rule_xref.Remark, ppdm_rule_xref.Rule_xref_condition, ppdm_rule_xref.Rule_xref_type, ppdm_rule_xref.Source, ppdm_rule_xref.Row_changed_by, ppdm_rule_xref.Row_changed_date, ppdm_rule_xref.Row_created_by, ppdm_rule_xref.Row_created_date, ppdm_rule_xref.Row_effective_date, ppdm_rule_xref.Row_expiry_date, ppdm_rule_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmRuleXref(c *fiber.Ctx) error {
	var ppdm_rule_xref dto.Ppdm_rule_xref

	setDefaults(&ppdm_rule_xref)

	if err := json.Unmarshal(c.Body(), &ppdm_rule_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_rule_xref.Dependency_id = id

    if ppdm_rule_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_rule_xref set rule_id = :1, rule_id2 = :2, xref_obs_no = :3, active_ind = :4, description = :5, effective_date = :6, expiry_date = :7, ppdm_guid = :8, remark = :9, rule_xref_condition = :10, rule_xref_type = :11, source = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where dependency_id = :20")
	if err != nil {
		return err
	}

	ppdm_rule_xref.Row_changed_date = formatDate(ppdm_rule_xref.Row_changed_date)
	ppdm_rule_xref.Effective_date = formatDateString(ppdm_rule_xref.Effective_date)
	ppdm_rule_xref.Expiry_date = formatDateString(ppdm_rule_xref.Expiry_date)
	ppdm_rule_xref.Row_effective_date = formatDateString(ppdm_rule_xref.Row_effective_date)
	ppdm_rule_xref.Row_expiry_date = formatDateString(ppdm_rule_xref.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_rule_xref.Rule_id, ppdm_rule_xref.Rule_id2, ppdm_rule_xref.Xref_obs_no, ppdm_rule_xref.Active_ind, ppdm_rule_xref.Description, ppdm_rule_xref.Effective_date, ppdm_rule_xref.Expiry_date, ppdm_rule_xref.Ppdm_guid, ppdm_rule_xref.Remark, ppdm_rule_xref.Rule_xref_condition, ppdm_rule_xref.Rule_xref_type, ppdm_rule_xref.Source, ppdm_rule_xref.Row_changed_by, ppdm_rule_xref.Row_changed_date, ppdm_rule_xref.Row_created_by, ppdm_rule_xref.Row_effective_date, ppdm_rule_xref.Row_expiry_date, ppdm_rule_xref.Row_quality, ppdm_rule_xref.Dependency_id)
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

func PatchPpdmRuleXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_rule_xref set "
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
	query += " where dependency_id = :id"

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

func DeletePpdmRuleXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_rule_xref dto.Ppdm_rule_xref
	ppdm_rule_xref.Dependency_id = id

	stmt, err := db.Prepare("delete from ppdm_rule_xref where dependency_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_rule_xref.Dependency_id)
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



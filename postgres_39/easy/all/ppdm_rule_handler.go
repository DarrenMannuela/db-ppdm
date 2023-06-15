package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmRule(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_rule")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_rule

	for rows.Next() {
		var ppdm_rule dto.Ppdm_rule
		if err := rows.Scan(&ppdm_rule.Rule_id, &ppdm_rule.Active_ind, &ppdm_rule.Current_status, &ppdm_rule.Effective_date, &ppdm_rule.Expiry_date, &ppdm_rule.Ppdm_guid, &ppdm_rule.Remark, &ppdm_rule.Rule_class, &ppdm_rule.Rule_desc, &ppdm_rule.Rule_purpose, &ppdm_rule.Rule_query, &ppdm_rule.Source, &ppdm_rule.Use_condition_desc, &ppdm_rule.Use_condition_type, &ppdm_rule.Row_changed_by, &ppdm_rule.Row_changed_date, &ppdm_rule.Row_created_by, &ppdm_rule.Row_created_date, &ppdm_rule.Row_effective_date, &ppdm_rule.Row_expiry_date, &ppdm_rule.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_rule to result
		result = append(result, ppdm_rule)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmRule(c *fiber.Ctx) error {
	var ppdm_rule dto.Ppdm_rule

	setDefaults(&ppdm_rule)

	if err := json.Unmarshal(c.Body(), &ppdm_rule); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_rule values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	ppdm_rule.Row_created_date = formatDate(ppdm_rule.Row_created_date)
	ppdm_rule.Row_changed_date = formatDate(ppdm_rule.Row_changed_date)
	ppdm_rule.Effective_date = formatDateString(ppdm_rule.Effective_date)
	ppdm_rule.Expiry_date = formatDateString(ppdm_rule.Expiry_date)
	ppdm_rule.Row_effective_date = formatDateString(ppdm_rule.Row_effective_date)
	ppdm_rule.Row_expiry_date = formatDateString(ppdm_rule.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_rule.Rule_id, ppdm_rule.Active_ind, ppdm_rule.Current_status, ppdm_rule.Effective_date, ppdm_rule.Expiry_date, ppdm_rule.Ppdm_guid, ppdm_rule.Remark, ppdm_rule.Rule_class, ppdm_rule.Rule_desc, ppdm_rule.Rule_purpose, ppdm_rule.Rule_query, ppdm_rule.Source, ppdm_rule.Use_condition_desc, ppdm_rule.Use_condition_type, ppdm_rule.Row_changed_by, ppdm_rule.Row_changed_date, ppdm_rule.Row_created_by, ppdm_rule.Row_created_date, ppdm_rule.Row_effective_date, ppdm_rule.Row_expiry_date, ppdm_rule.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmRule(c *fiber.Ctx) error {
	var ppdm_rule dto.Ppdm_rule

	setDefaults(&ppdm_rule)

	if err := json.Unmarshal(c.Body(), &ppdm_rule); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_rule.Rule_id = id

    if ppdm_rule.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_rule set active_ind = :1, current_status = :2, effective_date = :3, expiry_date = :4, ppdm_guid = :5, remark = :6, rule_class = :7, rule_desc = :8, rule_purpose = :9, rule_query = :10, source = :11, use_condition_desc = :12, use_condition_type = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where rule_id = :21")
	if err != nil {
		return err
	}

	ppdm_rule.Row_changed_date = formatDate(ppdm_rule.Row_changed_date)
	ppdm_rule.Effective_date = formatDateString(ppdm_rule.Effective_date)
	ppdm_rule.Expiry_date = formatDateString(ppdm_rule.Expiry_date)
	ppdm_rule.Row_effective_date = formatDateString(ppdm_rule.Row_effective_date)
	ppdm_rule.Row_expiry_date = formatDateString(ppdm_rule.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_rule.Active_ind, ppdm_rule.Current_status, ppdm_rule.Effective_date, ppdm_rule.Expiry_date, ppdm_rule.Ppdm_guid, ppdm_rule.Remark, ppdm_rule.Rule_class, ppdm_rule.Rule_desc, ppdm_rule.Rule_purpose, ppdm_rule.Rule_query, ppdm_rule.Source, ppdm_rule.Use_condition_desc, ppdm_rule.Use_condition_type, ppdm_rule.Row_changed_by, ppdm_rule.Row_changed_date, ppdm_rule.Row_created_by, ppdm_rule.Row_effective_date, ppdm_rule.Row_expiry_date, ppdm_rule.Row_quality, ppdm_rule.Rule_id)
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

func PatchPpdmRule(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_rule set "
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
	query += " where rule_id = :id"

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

func DeletePpdmRule(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_rule dto.Ppdm_rule
	ppdm_rule.Rule_id = id

	stmt, err := db.Prepare("delete from ppdm_rule where rule_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_rule.Rule_id)
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



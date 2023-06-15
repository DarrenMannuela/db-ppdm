package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmRuleEnforcement(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_rule_enforcement")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_rule_enforcement

	for rows.Next() {
		var ppdm_rule_enforcement dto.Ppdm_rule_enforcement
		if err := rows.Scan(&ppdm_rule_enforcement.Rule_id, &ppdm_rule_enforcement.Enforcement_id, &ppdm_rule_enforcement.Abbreviation, &ppdm_rule_enforcement.Active_ind, &ppdm_rule_enforcement.Effective_date, &ppdm_rule_enforcement.Enforce_method, &ppdm_rule_enforcement.Expiry_date, &ppdm_rule_enforcement.Long_name, &ppdm_rule_enforcement.Owner_ba_id, &ppdm_rule_enforcement.Ppdm_guid, &ppdm_rule_enforcement.Procedure_id, &ppdm_rule_enforcement.Remark, &ppdm_rule_enforcement.Rule_fail_result, &ppdm_rule_enforcement.Short_name, &ppdm_rule_enforcement.Source, &ppdm_rule_enforcement.Sw_application_id, &ppdm_rule_enforcement.Row_changed_by, &ppdm_rule_enforcement.Row_changed_date, &ppdm_rule_enforcement.Row_created_by, &ppdm_rule_enforcement.Row_created_date, &ppdm_rule_enforcement.Row_effective_date, &ppdm_rule_enforcement.Row_expiry_date, &ppdm_rule_enforcement.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_rule_enforcement to result
		result = append(result, ppdm_rule_enforcement)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmRuleEnforcement(c *fiber.Ctx) error {
	var ppdm_rule_enforcement dto.Ppdm_rule_enforcement

	setDefaults(&ppdm_rule_enforcement)

	if err := json.Unmarshal(c.Body(), &ppdm_rule_enforcement); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_rule_enforcement values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	ppdm_rule_enforcement.Row_created_date = formatDate(ppdm_rule_enforcement.Row_created_date)
	ppdm_rule_enforcement.Row_changed_date = formatDate(ppdm_rule_enforcement.Row_changed_date)
	ppdm_rule_enforcement.Effective_date = formatDateString(ppdm_rule_enforcement.Effective_date)
	ppdm_rule_enforcement.Expiry_date = formatDateString(ppdm_rule_enforcement.Expiry_date)
	ppdm_rule_enforcement.Row_effective_date = formatDateString(ppdm_rule_enforcement.Row_effective_date)
	ppdm_rule_enforcement.Row_expiry_date = formatDateString(ppdm_rule_enforcement.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_rule_enforcement.Rule_id, ppdm_rule_enforcement.Enforcement_id, ppdm_rule_enforcement.Abbreviation, ppdm_rule_enforcement.Active_ind, ppdm_rule_enforcement.Effective_date, ppdm_rule_enforcement.Enforce_method, ppdm_rule_enforcement.Expiry_date, ppdm_rule_enforcement.Long_name, ppdm_rule_enforcement.Owner_ba_id, ppdm_rule_enforcement.Ppdm_guid, ppdm_rule_enforcement.Procedure_id, ppdm_rule_enforcement.Remark, ppdm_rule_enforcement.Rule_fail_result, ppdm_rule_enforcement.Short_name, ppdm_rule_enforcement.Source, ppdm_rule_enforcement.Sw_application_id, ppdm_rule_enforcement.Row_changed_by, ppdm_rule_enforcement.Row_changed_date, ppdm_rule_enforcement.Row_created_by, ppdm_rule_enforcement.Row_created_date, ppdm_rule_enforcement.Row_effective_date, ppdm_rule_enforcement.Row_expiry_date, ppdm_rule_enforcement.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmRuleEnforcement(c *fiber.Ctx) error {
	var ppdm_rule_enforcement dto.Ppdm_rule_enforcement

	setDefaults(&ppdm_rule_enforcement)

	if err := json.Unmarshal(c.Body(), &ppdm_rule_enforcement); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_rule_enforcement.Rule_id = id

    if ppdm_rule_enforcement.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_rule_enforcement set enforcement_id = :1, abbreviation = :2, active_ind = :3, effective_date = :4, enforce_method = :5, expiry_date = :6, long_name = :7, owner_ba_id = :8, ppdm_guid = :9, procedure_id = :10, remark = :11, rule_fail_result = :12, short_name = :13, source = :14, sw_application_id = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where rule_id = :23")
	if err != nil {
		return err
	}

	ppdm_rule_enforcement.Row_changed_date = formatDate(ppdm_rule_enforcement.Row_changed_date)
	ppdm_rule_enforcement.Effective_date = formatDateString(ppdm_rule_enforcement.Effective_date)
	ppdm_rule_enforcement.Expiry_date = formatDateString(ppdm_rule_enforcement.Expiry_date)
	ppdm_rule_enforcement.Row_effective_date = formatDateString(ppdm_rule_enforcement.Row_effective_date)
	ppdm_rule_enforcement.Row_expiry_date = formatDateString(ppdm_rule_enforcement.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_rule_enforcement.Enforcement_id, ppdm_rule_enforcement.Abbreviation, ppdm_rule_enforcement.Active_ind, ppdm_rule_enforcement.Effective_date, ppdm_rule_enforcement.Enforce_method, ppdm_rule_enforcement.Expiry_date, ppdm_rule_enforcement.Long_name, ppdm_rule_enforcement.Owner_ba_id, ppdm_rule_enforcement.Ppdm_guid, ppdm_rule_enforcement.Procedure_id, ppdm_rule_enforcement.Remark, ppdm_rule_enforcement.Rule_fail_result, ppdm_rule_enforcement.Short_name, ppdm_rule_enforcement.Source, ppdm_rule_enforcement.Sw_application_id, ppdm_rule_enforcement.Row_changed_by, ppdm_rule_enforcement.Row_changed_date, ppdm_rule_enforcement.Row_created_by, ppdm_rule_enforcement.Row_effective_date, ppdm_rule_enforcement.Row_expiry_date, ppdm_rule_enforcement.Row_quality, ppdm_rule_enforcement.Rule_id)
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

func PatchPpdmRuleEnforcement(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_rule_enforcement set "
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

func DeletePpdmRuleEnforcement(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_rule_enforcement dto.Ppdm_rule_enforcement
	ppdm_rule_enforcement.Rule_id = id

	stmt, err := db.Prepare("delete from ppdm_rule_enforcement where rule_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_rule_enforcement.Rule_id)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRPpdmRuleUseCond(c *fiber.Ctx) error {
	rows, err := db.Query("select * from r_ppdm_rule_use_cond")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.R_ppdm_rule_use_cond

	for rows.Next() {
		var r_ppdm_rule_use_cond dto.R_ppdm_rule_use_cond
		if err := rows.Scan(&r_ppdm_rule_use_cond.Use_condition_type, &r_ppdm_rule_use_cond.Abbreviation, &r_ppdm_rule_use_cond.Active_ind, &r_ppdm_rule_use_cond.Effective_date, &r_ppdm_rule_use_cond.Expiry_date, &r_ppdm_rule_use_cond.Long_name, &r_ppdm_rule_use_cond.Ppdm_guid, &r_ppdm_rule_use_cond.Remark, &r_ppdm_rule_use_cond.Short_name, &r_ppdm_rule_use_cond.Source, &r_ppdm_rule_use_cond.Row_changed_by, &r_ppdm_rule_use_cond.Row_changed_date, &r_ppdm_rule_use_cond.Row_created_by, &r_ppdm_rule_use_cond.Row_created_date, &r_ppdm_rule_use_cond.Row_effective_date, &r_ppdm_rule_use_cond.Row_expiry_date, &r_ppdm_rule_use_cond.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append r_ppdm_rule_use_cond to result
		result = append(result, r_ppdm_rule_use_cond)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRPpdmRuleUseCond(c *fiber.Ctx) error {
	var r_ppdm_rule_use_cond dto.R_ppdm_rule_use_cond

	setDefaults(&r_ppdm_rule_use_cond)

	if err := json.Unmarshal(c.Body(), &r_ppdm_rule_use_cond); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into r_ppdm_rule_use_cond values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	r_ppdm_rule_use_cond.Row_created_date = formatDate(r_ppdm_rule_use_cond.Row_created_date)
	r_ppdm_rule_use_cond.Row_changed_date = formatDate(r_ppdm_rule_use_cond.Row_changed_date)
	r_ppdm_rule_use_cond.Effective_date = formatDateString(r_ppdm_rule_use_cond.Effective_date)
	r_ppdm_rule_use_cond.Expiry_date = formatDateString(r_ppdm_rule_use_cond.Expiry_date)
	r_ppdm_rule_use_cond.Row_effective_date = formatDateString(r_ppdm_rule_use_cond.Row_effective_date)
	r_ppdm_rule_use_cond.Row_expiry_date = formatDateString(r_ppdm_rule_use_cond.Row_expiry_date)






	rows, err := stmt.Exec(r_ppdm_rule_use_cond.Use_condition_type, r_ppdm_rule_use_cond.Abbreviation, r_ppdm_rule_use_cond.Active_ind, r_ppdm_rule_use_cond.Effective_date, r_ppdm_rule_use_cond.Expiry_date, r_ppdm_rule_use_cond.Long_name, r_ppdm_rule_use_cond.Ppdm_guid, r_ppdm_rule_use_cond.Remark, r_ppdm_rule_use_cond.Short_name, r_ppdm_rule_use_cond.Source, r_ppdm_rule_use_cond.Row_changed_by, r_ppdm_rule_use_cond.Row_changed_date, r_ppdm_rule_use_cond.Row_created_by, r_ppdm_rule_use_cond.Row_created_date, r_ppdm_rule_use_cond.Row_effective_date, r_ppdm_rule_use_cond.Row_expiry_date, r_ppdm_rule_use_cond.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRPpdmRuleUseCond(c *fiber.Ctx) error {
	var r_ppdm_rule_use_cond dto.R_ppdm_rule_use_cond

	setDefaults(&r_ppdm_rule_use_cond)

	if err := json.Unmarshal(c.Body(), &r_ppdm_rule_use_cond); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	r_ppdm_rule_use_cond.Use_condition_type = id

    if r_ppdm_rule_use_cond.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update r_ppdm_rule_use_cond set abbreviation = :1, active_ind = :2, effective_date = :3, expiry_date = :4, long_name = :5, ppdm_guid = :6, remark = :7, short_name = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where use_condition_type = :17")
	if err != nil {
		return err
	}

	r_ppdm_rule_use_cond.Row_changed_date = formatDate(r_ppdm_rule_use_cond.Row_changed_date)
	r_ppdm_rule_use_cond.Effective_date = formatDateString(r_ppdm_rule_use_cond.Effective_date)
	r_ppdm_rule_use_cond.Expiry_date = formatDateString(r_ppdm_rule_use_cond.Expiry_date)
	r_ppdm_rule_use_cond.Row_effective_date = formatDateString(r_ppdm_rule_use_cond.Row_effective_date)
	r_ppdm_rule_use_cond.Row_expiry_date = formatDateString(r_ppdm_rule_use_cond.Row_expiry_date)






	rows, err := stmt.Exec(r_ppdm_rule_use_cond.Abbreviation, r_ppdm_rule_use_cond.Active_ind, r_ppdm_rule_use_cond.Effective_date, r_ppdm_rule_use_cond.Expiry_date, r_ppdm_rule_use_cond.Long_name, r_ppdm_rule_use_cond.Ppdm_guid, r_ppdm_rule_use_cond.Remark, r_ppdm_rule_use_cond.Short_name, r_ppdm_rule_use_cond.Source, r_ppdm_rule_use_cond.Row_changed_by, r_ppdm_rule_use_cond.Row_changed_date, r_ppdm_rule_use_cond.Row_created_by, r_ppdm_rule_use_cond.Row_effective_date, r_ppdm_rule_use_cond.Row_expiry_date, r_ppdm_rule_use_cond.Row_quality, r_ppdm_rule_use_cond.Use_condition_type)
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

func PatchRPpdmRuleUseCond(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update r_ppdm_rule_use_cond set "
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
	query += " where use_condition_type = :id"

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

func DeleteRPpdmRuleUseCond(c *fiber.Ctx) error {
	id := c.Params("id")
	var r_ppdm_rule_use_cond dto.R_ppdm_rule_use_cond
	r_ppdm_rule_use_cond.Use_condition_type = id

	stmt, err := db.Prepare("delete from r_ppdm_rule_use_cond where use_condition_type = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(r_ppdm_rule_use_cond.Use_condition_type)
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



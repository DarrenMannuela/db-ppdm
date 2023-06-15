package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPdenDeclineCondition(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pden_decline_condition")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pden_decline_condition

	for rows.Next() {
		var pden_decline_condition dto.Pden_decline_condition
		if err := rows.Scan(&pden_decline_condition.Pden_id, &pden_decline_condition.Pden_subtype, &pden_decline_condition.Pden_source, &pden_decline_condition.Product_type, &pden_decline_condition.Case_id, &pden_decline_condition.Condition_id, &pden_decline_condition.Period_type, &pden_decline_condition.Period_id, &pden_decline_condition.Active_ind, &pden_decline_condition.Condition_code, &pden_decline_condition.Condition_desc, &pden_decline_condition.Condition_type, &pden_decline_condition.Condition_value, &pden_decline_condition.Effective_date, &pden_decline_condition.Expiry_date, &pden_decline_condition.Ppdm_guid, &pden_decline_condition.Remark, &pden_decline_condition.Source, &pden_decline_condition.Volume_date, &pden_decline_condition.Volume_date_desc, &pden_decline_condition.Row_changed_by, &pden_decline_condition.Row_changed_date, &pden_decline_condition.Row_created_by, &pden_decline_condition.Row_created_date, &pden_decline_condition.Row_effective_date, &pden_decline_condition.Row_expiry_date, &pden_decline_condition.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pden_decline_condition to result
		result = append(result, pden_decline_condition)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPdenDeclineCondition(c *fiber.Ctx) error {
	var pden_decline_condition dto.Pden_decline_condition

	setDefaults(&pden_decline_condition)

	if err := json.Unmarshal(c.Body(), &pden_decline_condition); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pden_decline_condition values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27)")
	if err != nil {
		return err
	}
	pden_decline_condition.Row_created_date = formatDate(pden_decline_condition.Row_created_date)
	pden_decline_condition.Row_changed_date = formatDate(pden_decline_condition.Row_changed_date)
	pden_decline_condition.Effective_date = formatDateString(pden_decline_condition.Effective_date)
	pden_decline_condition.Expiry_date = formatDateString(pden_decline_condition.Expiry_date)
	pden_decline_condition.Volume_date = formatDateString(pden_decline_condition.Volume_date)
	pden_decline_condition.Row_effective_date = formatDateString(pden_decline_condition.Row_effective_date)
	pden_decline_condition.Row_expiry_date = formatDateString(pden_decline_condition.Row_expiry_date)






	rows, err := stmt.Exec(pden_decline_condition.Pden_id, pden_decline_condition.Pden_subtype, pden_decline_condition.Pden_source, pden_decline_condition.Product_type, pden_decline_condition.Case_id, pden_decline_condition.Condition_id, pden_decline_condition.Period_type, pden_decline_condition.Period_id, pden_decline_condition.Active_ind, pden_decline_condition.Condition_code, pden_decline_condition.Condition_desc, pden_decline_condition.Condition_type, pden_decline_condition.Condition_value, pden_decline_condition.Effective_date, pden_decline_condition.Expiry_date, pden_decline_condition.Ppdm_guid, pden_decline_condition.Remark, pden_decline_condition.Source, pden_decline_condition.Volume_date, pden_decline_condition.Volume_date_desc, pden_decline_condition.Row_changed_by, pden_decline_condition.Row_changed_date, pden_decline_condition.Row_created_by, pden_decline_condition.Row_created_date, pden_decline_condition.Row_effective_date, pden_decline_condition.Row_expiry_date, pden_decline_condition.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePdenDeclineCondition(c *fiber.Ctx) error {
	var pden_decline_condition dto.Pden_decline_condition

	setDefaults(&pden_decline_condition)

	if err := json.Unmarshal(c.Body(), &pden_decline_condition); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pden_decline_condition.Pden_id = id

    if pden_decline_condition.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pden_decline_condition set pden_subtype = :1, pden_source = :2, product_type = :3, case_id = :4, condition_id = :5, period_type = :6, period_id = :7, active_ind = :8, condition_code = :9, condition_desc = :10, condition_type = :11, condition_value = :12, effective_date = :13, expiry_date = :14, ppdm_guid = :15, remark = :16, source = :17, volume_date = :18, volume_date_desc = :19, row_changed_by = :20, row_changed_date = :21, row_created_by = :22, row_effective_date = :23, row_expiry_date = :24, row_quality = :25 where pden_id = :27")
	if err != nil {
		return err
	}

	pden_decline_condition.Row_changed_date = formatDate(pden_decline_condition.Row_changed_date)
	pden_decline_condition.Effective_date = formatDateString(pden_decline_condition.Effective_date)
	pden_decline_condition.Expiry_date = formatDateString(pden_decline_condition.Expiry_date)
	pden_decline_condition.Volume_date = formatDateString(pden_decline_condition.Volume_date)
	pden_decline_condition.Row_effective_date = formatDateString(pden_decline_condition.Row_effective_date)
	pden_decline_condition.Row_expiry_date = formatDateString(pden_decline_condition.Row_expiry_date)






	rows, err := stmt.Exec(pden_decline_condition.Pden_subtype, pden_decline_condition.Pden_source, pden_decline_condition.Product_type, pden_decline_condition.Case_id, pden_decline_condition.Condition_id, pden_decline_condition.Period_type, pden_decline_condition.Period_id, pden_decline_condition.Active_ind, pden_decline_condition.Condition_code, pden_decline_condition.Condition_desc, pden_decline_condition.Condition_type, pden_decline_condition.Condition_value, pden_decline_condition.Effective_date, pden_decline_condition.Expiry_date, pden_decline_condition.Ppdm_guid, pden_decline_condition.Remark, pden_decline_condition.Source, pden_decline_condition.Volume_date, pden_decline_condition.Volume_date_desc, pden_decline_condition.Row_changed_by, pden_decline_condition.Row_changed_date, pden_decline_condition.Row_created_by, pden_decline_condition.Row_effective_date, pden_decline_condition.Row_expiry_date, pden_decline_condition.Row_quality, pden_decline_condition.Pden_id)
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

func PatchPdenDeclineCondition(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pden_decline_condition set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "volume_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where pden_id = :id"

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

func DeletePdenDeclineCondition(c *fiber.Ctx) error {
	id := c.Params("id")
	var pden_decline_condition dto.Pden_decline_condition
	pden_decline_condition.Pden_id = id

	stmt, err := db.Prepare("delete from pden_decline_condition where pden_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pden_decline_condition.Pden_id)
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



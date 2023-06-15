package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmRuleRemark(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_rule_remark")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_rule_remark

	for rows.Next() {
		var ppdm_rule_remark dto.Ppdm_rule_remark
		if err := rows.Scan(&ppdm_rule_remark.Rule_id, &ppdm_rule_remark.Remark_type, &ppdm_rule_remark.Remark_seq_no, &ppdm_rule_remark.Active_ind, &ppdm_rule_remark.Effective_date, &ppdm_rule_remark.Expiry_date, &ppdm_rule_remark.Ppdm_guid, &ppdm_rule_remark.Remark, &ppdm_rule_remark.Remark_by_ba_id, &ppdm_rule_remark.Remark_date, &ppdm_rule_remark.Source, &ppdm_rule_remark.Row_changed_by, &ppdm_rule_remark.Row_changed_date, &ppdm_rule_remark.Row_created_by, &ppdm_rule_remark.Row_created_date, &ppdm_rule_remark.Row_effective_date, &ppdm_rule_remark.Row_expiry_date, &ppdm_rule_remark.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_rule_remark to result
		result = append(result, ppdm_rule_remark)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmRuleRemark(c *fiber.Ctx) error {
	var ppdm_rule_remark dto.Ppdm_rule_remark

	setDefaults(&ppdm_rule_remark)

	if err := json.Unmarshal(c.Body(), &ppdm_rule_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_rule_remark values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	ppdm_rule_remark.Row_created_date = formatDate(ppdm_rule_remark.Row_created_date)
	ppdm_rule_remark.Row_changed_date = formatDate(ppdm_rule_remark.Row_changed_date)
	ppdm_rule_remark.Effective_date = formatDateString(ppdm_rule_remark.Effective_date)
	ppdm_rule_remark.Expiry_date = formatDateString(ppdm_rule_remark.Expiry_date)
	ppdm_rule_remark.Remark_date = formatDateString(ppdm_rule_remark.Remark_date)
	ppdm_rule_remark.Row_effective_date = formatDateString(ppdm_rule_remark.Row_effective_date)
	ppdm_rule_remark.Row_expiry_date = formatDateString(ppdm_rule_remark.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_rule_remark.Rule_id, ppdm_rule_remark.Remark_type, ppdm_rule_remark.Remark_seq_no, ppdm_rule_remark.Active_ind, ppdm_rule_remark.Effective_date, ppdm_rule_remark.Expiry_date, ppdm_rule_remark.Ppdm_guid, ppdm_rule_remark.Remark, ppdm_rule_remark.Remark_by_ba_id, ppdm_rule_remark.Remark_date, ppdm_rule_remark.Source, ppdm_rule_remark.Row_changed_by, ppdm_rule_remark.Row_changed_date, ppdm_rule_remark.Row_created_by, ppdm_rule_remark.Row_created_date, ppdm_rule_remark.Row_effective_date, ppdm_rule_remark.Row_expiry_date, ppdm_rule_remark.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmRuleRemark(c *fiber.Ctx) error {
	var ppdm_rule_remark dto.Ppdm_rule_remark

	setDefaults(&ppdm_rule_remark)

	if err := json.Unmarshal(c.Body(), &ppdm_rule_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_rule_remark.Rule_id = id

    if ppdm_rule_remark.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_rule_remark set remark_type = :1, remark_seq_no = :2, active_ind = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, remark_by_ba_id = :8, remark_date = :9, source = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where rule_id = :18")
	if err != nil {
		return err
	}

	ppdm_rule_remark.Row_changed_date = formatDate(ppdm_rule_remark.Row_changed_date)
	ppdm_rule_remark.Effective_date = formatDateString(ppdm_rule_remark.Effective_date)
	ppdm_rule_remark.Expiry_date = formatDateString(ppdm_rule_remark.Expiry_date)
	ppdm_rule_remark.Remark_date = formatDateString(ppdm_rule_remark.Remark_date)
	ppdm_rule_remark.Row_effective_date = formatDateString(ppdm_rule_remark.Row_effective_date)
	ppdm_rule_remark.Row_expiry_date = formatDateString(ppdm_rule_remark.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_rule_remark.Remark_type, ppdm_rule_remark.Remark_seq_no, ppdm_rule_remark.Active_ind, ppdm_rule_remark.Effective_date, ppdm_rule_remark.Expiry_date, ppdm_rule_remark.Ppdm_guid, ppdm_rule_remark.Remark, ppdm_rule_remark.Remark_by_ba_id, ppdm_rule_remark.Remark_date, ppdm_rule_remark.Source, ppdm_rule_remark.Row_changed_by, ppdm_rule_remark.Row_changed_date, ppdm_rule_remark.Row_created_by, ppdm_rule_remark.Row_effective_date, ppdm_rule_remark.Row_expiry_date, ppdm_rule_remark.Row_quality, ppdm_rule_remark.Rule_id)
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

func PatchPpdmRuleRemark(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_rule_remark set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "remark_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeletePpdmRuleRemark(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_rule_remark dto.Ppdm_rule_remark
	ppdm_rule_remark.Rule_id = id

	stmt, err := db.Prepare("delete from ppdm_rule_remark where rule_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_rule_remark.Rule_id)
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



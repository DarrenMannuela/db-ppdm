package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmRuleDetail(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_rule_detail")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_rule_detail

	for rows.Next() {
		var ppdm_rule_detail dto.Ppdm_rule_detail
		if err := rows.Scan(&ppdm_rule_detail.Rule_id, &ppdm_rule_detail.Detail_seq_no, &ppdm_rule_detail.Active_ind, &ppdm_rule_detail.Average_value, &ppdm_rule_detail.Average_value_ouom, &ppdm_rule_detail.Average_value_uom, &ppdm_rule_detail.Boolean_rule, &ppdm_rule_detail.Business_rule, &ppdm_rule_detail.Date_format_desc, &ppdm_rule_detail.Effective_date, &ppdm_rule_detail.Expiry_date, &ppdm_rule_detail.Max_date, &ppdm_rule_detail.Max_value, &ppdm_rule_detail.Max_value_ouom, &ppdm_rule_detail.Max_value_uom, &ppdm_rule_detail.Min_date, &ppdm_rule_detail.Min_value, &ppdm_rule_detail.Min_value_ouom, &ppdm_rule_detail.Min_value_uom, &ppdm_rule_detail.Ppdm_guid, &ppdm_rule_detail.Referenced_column_name, &ppdm_rule_detail.Reference_column_name2, &ppdm_rule_detail.Reference_system_id, &ppdm_rule_detail.Reference_table_name, &ppdm_rule_detail.Reference_table_name2, &ppdm_rule_detail.Reference_value, &ppdm_rule_detail.Reference_value_ouom, &ppdm_rule_detail.Reference_value_type, &ppdm_rule_detail.Reference_value_uom, &ppdm_rule_detail.Remark, &ppdm_rule_detail.Rule_desc, &ppdm_rule_detail.Rule_detail_type, &ppdm_rule_detail.Source, &ppdm_rule_detail.Row_changed_by, &ppdm_rule_detail.Row_changed_date, &ppdm_rule_detail.Row_created_by, &ppdm_rule_detail.Row_created_date, &ppdm_rule_detail.Row_effective_date, &ppdm_rule_detail.Row_expiry_date, &ppdm_rule_detail.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_rule_detail to result
		result = append(result, ppdm_rule_detail)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmRuleDetail(c *fiber.Ctx) error {
	var ppdm_rule_detail dto.Ppdm_rule_detail

	setDefaults(&ppdm_rule_detail)

	if err := json.Unmarshal(c.Body(), &ppdm_rule_detail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_rule_detail values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40)")
	if err != nil {
		return err
	}
	ppdm_rule_detail.Row_created_date = formatDate(ppdm_rule_detail.Row_created_date)
	ppdm_rule_detail.Row_changed_date = formatDate(ppdm_rule_detail.Row_changed_date)
	ppdm_rule_detail.Date_format_desc = formatDateString(ppdm_rule_detail.Date_format_desc)
	ppdm_rule_detail.Effective_date = formatDateString(ppdm_rule_detail.Effective_date)
	ppdm_rule_detail.Expiry_date = formatDateString(ppdm_rule_detail.Expiry_date)
	ppdm_rule_detail.Max_date = formatDateString(ppdm_rule_detail.Max_date)
	ppdm_rule_detail.Min_date = formatDateString(ppdm_rule_detail.Min_date)
	ppdm_rule_detail.Row_effective_date = formatDateString(ppdm_rule_detail.Row_effective_date)
	ppdm_rule_detail.Row_expiry_date = formatDateString(ppdm_rule_detail.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_rule_detail.Rule_id, ppdm_rule_detail.Detail_seq_no, ppdm_rule_detail.Active_ind, ppdm_rule_detail.Average_value, ppdm_rule_detail.Average_value_ouom, ppdm_rule_detail.Average_value_uom, ppdm_rule_detail.Boolean_rule, ppdm_rule_detail.Business_rule, ppdm_rule_detail.Date_format_desc, ppdm_rule_detail.Effective_date, ppdm_rule_detail.Expiry_date, ppdm_rule_detail.Max_date, ppdm_rule_detail.Max_value, ppdm_rule_detail.Max_value_ouom, ppdm_rule_detail.Max_value_uom, ppdm_rule_detail.Min_date, ppdm_rule_detail.Min_value, ppdm_rule_detail.Min_value_ouom, ppdm_rule_detail.Min_value_uom, ppdm_rule_detail.Ppdm_guid, ppdm_rule_detail.Referenced_column_name, ppdm_rule_detail.Reference_column_name2, ppdm_rule_detail.Reference_system_id, ppdm_rule_detail.Reference_table_name, ppdm_rule_detail.Reference_table_name2, ppdm_rule_detail.Reference_value, ppdm_rule_detail.Reference_value_ouom, ppdm_rule_detail.Reference_value_type, ppdm_rule_detail.Reference_value_uom, ppdm_rule_detail.Remark, ppdm_rule_detail.Rule_desc, ppdm_rule_detail.Rule_detail_type, ppdm_rule_detail.Source, ppdm_rule_detail.Row_changed_by, ppdm_rule_detail.Row_changed_date, ppdm_rule_detail.Row_created_by, ppdm_rule_detail.Row_created_date, ppdm_rule_detail.Row_effective_date, ppdm_rule_detail.Row_expiry_date, ppdm_rule_detail.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmRuleDetail(c *fiber.Ctx) error {
	var ppdm_rule_detail dto.Ppdm_rule_detail

	setDefaults(&ppdm_rule_detail)

	if err := json.Unmarshal(c.Body(), &ppdm_rule_detail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_rule_detail.Rule_id = id

    if ppdm_rule_detail.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_rule_detail set detail_seq_no = :1, active_ind = :2, average_value = :3, average_value_ouom = :4, average_value_uom = :5, boolean_rule = :6, business_rule = :7, date_format_desc = :8, effective_date = :9, expiry_date = :10, max_date = :11, max_value = :12, max_value_ouom = :13, max_value_uom = :14, min_date = :15, min_value = :16, min_value_ouom = :17, min_value_uom = :18, ppdm_guid = :19, referenced_column_name = :20, reference_column_name2 = :21, reference_system_id = :22, reference_table_name = :23, reference_table_name2 = :24, reference_value = :25, reference_value_ouom = :26, reference_value_type = :27, reference_value_uom = :28, remark = :29, rule_desc = :30, rule_detail_type = :31, source = :32, row_changed_by = :33, row_changed_date = :34, row_created_by = :35, row_effective_date = :36, row_expiry_date = :37, row_quality = :38 where rule_id = :40")
	if err != nil {
		return err
	}

	ppdm_rule_detail.Row_changed_date = formatDate(ppdm_rule_detail.Row_changed_date)
	ppdm_rule_detail.Date_format_desc = formatDateString(ppdm_rule_detail.Date_format_desc)
	ppdm_rule_detail.Effective_date = formatDateString(ppdm_rule_detail.Effective_date)
	ppdm_rule_detail.Expiry_date = formatDateString(ppdm_rule_detail.Expiry_date)
	ppdm_rule_detail.Max_date = formatDateString(ppdm_rule_detail.Max_date)
	ppdm_rule_detail.Min_date = formatDateString(ppdm_rule_detail.Min_date)
	ppdm_rule_detail.Row_effective_date = formatDateString(ppdm_rule_detail.Row_effective_date)
	ppdm_rule_detail.Row_expiry_date = formatDateString(ppdm_rule_detail.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_rule_detail.Detail_seq_no, ppdm_rule_detail.Active_ind, ppdm_rule_detail.Average_value, ppdm_rule_detail.Average_value_ouom, ppdm_rule_detail.Average_value_uom, ppdm_rule_detail.Boolean_rule, ppdm_rule_detail.Business_rule, ppdm_rule_detail.Date_format_desc, ppdm_rule_detail.Effective_date, ppdm_rule_detail.Expiry_date, ppdm_rule_detail.Max_date, ppdm_rule_detail.Max_value, ppdm_rule_detail.Max_value_ouom, ppdm_rule_detail.Max_value_uom, ppdm_rule_detail.Min_date, ppdm_rule_detail.Min_value, ppdm_rule_detail.Min_value_ouom, ppdm_rule_detail.Min_value_uom, ppdm_rule_detail.Ppdm_guid, ppdm_rule_detail.Referenced_column_name, ppdm_rule_detail.Reference_column_name2, ppdm_rule_detail.Reference_system_id, ppdm_rule_detail.Reference_table_name, ppdm_rule_detail.Reference_table_name2, ppdm_rule_detail.Reference_value, ppdm_rule_detail.Reference_value_ouom, ppdm_rule_detail.Reference_value_type, ppdm_rule_detail.Reference_value_uom, ppdm_rule_detail.Remark, ppdm_rule_detail.Rule_desc, ppdm_rule_detail.Rule_detail_type, ppdm_rule_detail.Source, ppdm_rule_detail.Row_changed_by, ppdm_rule_detail.Row_changed_date, ppdm_rule_detail.Row_created_by, ppdm_rule_detail.Row_effective_date, ppdm_rule_detail.Row_expiry_date, ppdm_rule_detail.Row_quality, ppdm_rule_detail.Rule_id)
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

func PatchPpdmRuleDetail(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_rule_detail set "
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
		} else if key == "date_format_desc" || key == "effective_date" || key == "expiry_date" || key == "max_date" || key == "min_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeletePpdmRuleDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_rule_detail dto.Ppdm_rule_detail
	ppdm_rule_detail.Rule_id = id

	stmt, err := db.Prepare("delete from ppdm_rule_detail where rule_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_rule_detail.Rule_id)
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



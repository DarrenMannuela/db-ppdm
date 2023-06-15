package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmMapRule(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_map_rule")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_map_rule

	for rows.Next() {
		var ppdm_map_rule dto.Ppdm_map_rule
		if err := rows.Scan(&ppdm_map_rule.Map_id, &ppdm_map_rule.Map_detail_obs_no, &ppdm_map_rule.Rule_seq_no, &ppdm_map_rule.Active_ind, &ppdm_map_rule.Create_method, &ppdm_map_rule.Date_format_desc, &ppdm_map_rule.Dep_column_name, &ppdm_map_rule.Dep_schema_entity_id, &ppdm_map_rule.Dep_system_id, &ppdm_map_rule.Dep_table_name, &ppdm_map_rule.Effective_date, &ppdm_map_rule.Expiry_date, &ppdm_map_rule.Map_rule_type, &ppdm_map_rule.Max_value, &ppdm_map_rule.Max_value_ouom, &ppdm_map_rule.Max_value_uom, &ppdm_map_rule.Min_value, &ppdm_map_rule.Min_value_ouom, &ppdm_map_rule.Min_value_uom, &ppdm_map_rule.Ppdm_guid, &ppdm_map_rule.Preferred_ind, &ppdm_map_rule.Procedure_id, &ppdm_map_rule.Procedure_system_id, &ppdm_map_rule.Remark, &ppdm_map_rule.Ring_seq_no, &ppdm_map_rule.Rule_desc, &ppdm_map_rule.Rule_owner_ba_id, &ppdm_map_rule.Rule_version_num, &ppdm_map_rule.Source, &ppdm_map_rule.Sw_application_id, &ppdm_map_rule.Row_changed_by, &ppdm_map_rule.Row_changed_date, &ppdm_map_rule.Row_created_by, &ppdm_map_rule.Row_created_date, &ppdm_map_rule.Row_effective_date, &ppdm_map_rule.Row_expiry_date, &ppdm_map_rule.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_map_rule to result
		result = append(result, ppdm_map_rule)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmMapRule(c *fiber.Ctx) error {
	var ppdm_map_rule dto.Ppdm_map_rule

	setDefaults(&ppdm_map_rule)

	if err := json.Unmarshal(c.Body(), &ppdm_map_rule); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_map_rule values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37)")
	if err != nil {
		return err
	}
	ppdm_map_rule.Row_created_date = formatDate(ppdm_map_rule.Row_created_date)
	ppdm_map_rule.Row_changed_date = formatDate(ppdm_map_rule.Row_changed_date)
	ppdm_map_rule.Date_format_desc = formatDateString(ppdm_map_rule.Date_format_desc)
	ppdm_map_rule.Effective_date = formatDateString(ppdm_map_rule.Effective_date)
	ppdm_map_rule.Expiry_date = formatDateString(ppdm_map_rule.Expiry_date)
	ppdm_map_rule.Row_effective_date = formatDateString(ppdm_map_rule.Row_effective_date)
	ppdm_map_rule.Row_expiry_date = formatDateString(ppdm_map_rule.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_map_rule.Map_id, ppdm_map_rule.Map_detail_obs_no, ppdm_map_rule.Rule_seq_no, ppdm_map_rule.Active_ind, ppdm_map_rule.Create_method, ppdm_map_rule.Date_format_desc, ppdm_map_rule.Dep_column_name, ppdm_map_rule.Dep_schema_entity_id, ppdm_map_rule.Dep_system_id, ppdm_map_rule.Dep_table_name, ppdm_map_rule.Effective_date, ppdm_map_rule.Expiry_date, ppdm_map_rule.Map_rule_type, ppdm_map_rule.Max_value, ppdm_map_rule.Max_value_ouom, ppdm_map_rule.Max_value_uom, ppdm_map_rule.Min_value, ppdm_map_rule.Min_value_ouom, ppdm_map_rule.Min_value_uom, ppdm_map_rule.Ppdm_guid, ppdm_map_rule.Preferred_ind, ppdm_map_rule.Procedure_id, ppdm_map_rule.Procedure_system_id, ppdm_map_rule.Remark, ppdm_map_rule.Ring_seq_no, ppdm_map_rule.Rule_desc, ppdm_map_rule.Rule_owner_ba_id, ppdm_map_rule.Rule_version_num, ppdm_map_rule.Source, ppdm_map_rule.Sw_application_id, ppdm_map_rule.Row_changed_by, ppdm_map_rule.Row_changed_date, ppdm_map_rule.Row_created_by, ppdm_map_rule.Row_created_date, ppdm_map_rule.Row_effective_date, ppdm_map_rule.Row_expiry_date, ppdm_map_rule.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmMapRule(c *fiber.Ctx) error {
	var ppdm_map_rule dto.Ppdm_map_rule

	setDefaults(&ppdm_map_rule)

	if err := json.Unmarshal(c.Body(), &ppdm_map_rule); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_map_rule.Map_id = id

    if ppdm_map_rule.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_map_rule set map_detail_obs_no = :1, rule_seq_no = :2, active_ind = :3, create_method = :4, date_format_desc = :5, dep_column_name = :6, dep_schema_entity_id = :7, dep_system_id = :8, dep_table_name = :9, effective_date = :10, expiry_date = :11, map_rule_type = :12, max_value = :13, max_value_ouom = :14, max_value_uom = :15, min_value = :16, min_value_ouom = :17, min_value_uom = :18, ppdm_guid = :19, preferred_ind = :20, procedure_id = :21, procedure_system_id = :22, remark = :23, ring_seq_no = :24, rule_desc = :25, rule_owner_ba_id = :26, rule_version_num = :27, source = :28, sw_application_id = :29, row_changed_by = :30, row_changed_date = :31, row_created_by = :32, row_effective_date = :33, row_expiry_date = :34, row_quality = :35 where map_id = :37")
	if err != nil {
		return err
	}

	ppdm_map_rule.Row_changed_date = formatDate(ppdm_map_rule.Row_changed_date)
	ppdm_map_rule.Date_format_desc = formatDateString(ppdm_map_rule.Date_format_desc)
	ppdm_map_rule.Effective_date = formatDateString(ppdm_map_rule.Effective_date)
	ppdm_map_rule.Expiry_date = formatDateString(ppdm_map_rule.Expiry_date)
	ppdm_map_rule.Row_effective_date = formatDateString(ppdm_map_rule.Row_effective_date)
	ppdm_map_rule.Row_expiry_date = formatDateString(ppdm_map_rule.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_map_rule.Map_detail_obs_no, ppdm_map_rule.Rule_seq_no, ppdm_map_rule.Active_ind, ppdm_map_rule.Create_method, ppdm_map_rule.Date_format_desc, ppdm_map_rule.Dep_column_name, ppdm_map_rule.Dep_schema_entity_id, ppdm_map_rule.Dep_system_id, ppdm_map_rule.Dep_table_name, ppdm_map_rule.Effective_date, ppdm_map_rule.Expiry_date, ppdm_map_rule.Map_rule_type, ppdm_map_rule.Max_value, ppdm_map_rule.Max_value_ouom, ppdm_map_rule.Max_value_uom, ppdm_map_rule.Min_value, ppdm_map_rule.Min_value_ouom, ppdm_map_rule.Min_value_uom, ppdm_map_rule.Ppdm_guid, ppdm_map_rule.Preferred_ind, ppdm_map_rule.Procedure_id, ppdm_map_rule.Procedure_system_id, ppdm_map_rule.Remark, ppdm_map_rule.Ring_seq_no, ppdm_map_rule.Rule_desc, ppdm_map_rule.Rule_owner_ba_id, ppdm_map_rule.Rule_version_num, ppdm_map_rule.Source, ppdm_map_rule.Sw_application_id, ppdm_map_rule.Row_changed_by, ppdm_map_rule.Row_changed_date, ppdm_map_rule.Row_created_by, ppdm_map_rule.Row_effective_date, ppdm_map_rule.Row_expiry_date, ppdm_map_rule.Row_quality, ppdm_map_rule.Map_id)
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

func PatchPpdmMapRule(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_map_rule set "
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
		} else if key == "date_format_desc" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where map_id = :id"

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

func DeletePpdmMapRule(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_map_rule dto.Ppdm_map_rule
	ppdm_map_rule.Map_id = id

	stmt, err := db.Prepare("delete from ppdm_map_rule where map_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_map_rule.Map_id)
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



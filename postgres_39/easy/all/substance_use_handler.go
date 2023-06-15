package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSubstanceUse(c *fiber.Ctx) error {
	rows, err := db.Query("select * from substance_use")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Substance_use

	for rows.Next() {
		var substance_use dto.Substance_use
		if err := rows.Scan(&substance_use.Substance_id, &substance_use.Substance_use_obs_no, &substance_use.Active_ind, &substance_use.Area_id, &substance_use.Area_type, &substance_use.Business_associate_id, &substance_use.Contract_id, &substance_use.Effective_date, &substance_use.Expiry_date, &substance_use.Land_right_id, &substance_use.Land_right_subtype, &substance_use.Ppdm_column_name, &substance_use.Ppdm_guid, &substance_use.Ppdm_system_id, &substance_use.Ppdm_table_name, &substance_use.Preferred_ind, &substance_use.Project_id, &substance_use.Remark, &substance_use.Source, &substance_use.Substance_alias_id, &substance_use.Substance_use_rule_desc, &substance_use.Substance_use_rule_type, &substance_use.Row_changed_by, &substance_use.Row_changed_date, &substance_use.Row_created_by, &substance_use.Row_created_date, &substance_use.Row_effective_date, &substance_use.Row_expiry_date, &substance_use.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append substance_use to result
		result = append(result, substance_use)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSubstanceUse(c *fiber.Ctx) error {
	var substance_use dto.Substance_use

	setDefaults(&substance_use)

	if err := json.Unmarshal(c.Body(), &substance_use); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into substance_use values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29)")
	if err != nil {
		return err
	}
	substance_use.Row_created_date = formatDate(substance_use.Row_created_date)
	substance_use.Row_changed_date = formatDate(substance_use.Row_changed_date)
	substance_use.Effective_date = formatDateString(substance_use.Effective_date)
	substance_use.Expiry_date = formatDateString(substance_use.Expiry_date)
	substance_use.Row_effective_date = formatDateString(substance_use.Row_effective_date)
	substance_use.Row_expiry_date = formatDateString(substance_use.Row_expiry_date)






	rows, err := stmt.Exec(substance_use.Substance_id, substance_use.Substance_use_obs_no, substance_use.Active_ind, substance_use.Area_id, substance_use.Area_type, substance_use.Business_associate_id, substance_use.Contract_id, substance_use.Effective_date, substance_use.Expiry_date, substance_use.Land_right_id, substance_use.Land_right_subtype, substance_use.Ppdm_column_name, substance_use.Ppdm_guid, substance_use.Ppdm_system_id, substance_use.Ppdm_table_name, substance_use.Preferred_ind, substance_use.Project_id, substance_use.Remark, substance_use.Source, substance_use.Substance_alias_id, substance_use.Substance_use_rule_desc, substance_use.Substance_use_rule_type, substance_use.Row_changed_by, substance_use.Row_changed_date, substance_use.Row_created_by, substance_use.Row_created_date, substance_use.Row_effective_date, substance_use.Row_expiry_date, substance_use.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSubstanceUse(c *fiber.Ctx) error {
	var substance_use dto.Substance_use

	setDefaults(&substance_use)

	if err := json.Unmarshal(c.Body(), &substance_use); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	substance_use.Substance_id = id

    if substance_use.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update substance_use set substance_use_obs_no = :1, active_ind = :2, area_id = :3, area_type = :4, business_associate_id = :5, contract_id = :6, effective_date = :7, expiry_date = :8, land_right_id = :9, land_right_subtype = :10, ppdm_column_name = :11, ppdm_guid = :12, ppdm_system_id = :13, ppdm_table_name = :14, preferred_ind = :15, project_id = :16, remark = :17, source = :18, substance_alias_id = :19, substance_use_rule_desc = :20, substance_use_rule_type = :21, row_changed_by = :22, row_changed_date = :23, row_created_by = :24, row_effective_date = :25, row_expiry_date = :26, row_quality = :27 where substance_id = :29")
	if err != nil {
		return err
	}

	substance_use.Row_changed_date = formatDate(substance_use.Row_changed_date)
	substance_use.Effective_date = formatDateString(substance_use.Effective_date)
	substance_use.Expiry_date = formatDateString(substance_use.Expiry_date)
	substance_use.Row_effective_date = formatDateString(substance_use.Row_effective_date)
	substance_use.Row_expiry_date = formatDateString(substance_use.Row_expiry_date)






	rows, err := stmt.Exec(substance_use.Substance_use_obs_no, substance_use.Active_ind, substance_use.Area_id, substance_use.Area_type, substance_use.Business_associate_id, substance_use.Contract_id, substance_use.Effective_date, substance_use.Expiry_date, substance_use.Land_right_id, substance_use.Land_right_subtype, substance_use.Ppdm_column_name, substance_use.Ppdm_guid, substance_use.Ppdm_system_id, substance_use.Ppdm_table_name, substance_use.Preferred_ind, substance_use.Project_id, substance_use.Remark, substance_use.Source, substance_use.Substance_alias_id, substance_use.Substance_use_rule_desc, substance_use.Substance_use_rule_type, substance_use.Row_changed_by, substance_use.Row_changed_date, substance_use.Row_created_by, substance_use.Row_effective_date, substance_use.Row_expiry_date, substance_use.Row_quality, substance_use.Substance_id)
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

func PatchSubstanceUse(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update substance_use set "
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
	query += " where substance_id = :id"

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

func DeleteSubstanceUse(c *fiber.Ctx) error {
	id := c.Params("id")
	var substance_use dto.Substance_use
	substance_use.Substance_id = id

	stmt, err := db.Prepare("delete from substance_use where substance_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(substance_use.Substance_id)
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



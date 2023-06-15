package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmGroupObject(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_group_object")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_group_object

	for rows.Next() {
		var ppdm_group_object dto.Ppdm_group_object
		if err := rows.Scan(&ppdm_group_object.Group_id, &ppdm_group_object.Object_obs_no, &ppdm_group_object.Active_ind, &ppdm_group_object.Code_version_obs_no, &ppdm_group_object.Code_version_source, &ppdm_group_object.Column_alias_id, &ppdm_group_object.Column_name, &ppdm_group_object.Constraint_name, &ppdm_group_object.Effective_date, &ppdm_group_object.Expiry_date, &ppdm_group_object.Group_use, &ppdm_group_object.Index_id, &ppdm_group_object.Object_type, &ppdm_group_object.Output_font, &ppdm_group_object.Output_font_backgr_color, &ppdm_group_object.Output_font_color, &ppdm_group_object.Output_font_effect, &ppdm_group_object.Output_font_size, &ppdm_group_object.Output_font_size_uom, &ppdm_group_object.Output_heading, &ppdm_group_object.Output_length, &ppdm_group_object.Output_length_uom, &ppdm_group_object.Output_precision, &ppdm_group_object.Ppdm_guid, &ppdm_group_object.Preferred_uom, &ppdm_group_object.Procedure_id, &ppdm_group_object.Property_set_id, &ppdm_group_object.Remark, &ppdm_group_object.Rule_id, &ppdm_group_object.Source, &ppdm_group_object.Sw_application_id, &ppdm_group_object.System_id, &ppdm_group_object.Table_alias, &ppdm_group_object.Table_name, &ppdm_group_object.Use_rule_description, &ppdm_group_object.Row_changed_by, &ppdm_group_object.Row_changed_date, &ppdm_group_object.Row_created_by, &ppdm_group_object.Row_created_date, &ppdm_group_object.Row_effective_date, &ppdm_group_object.Row_expiry_date, &ppdm_group_object.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_group_object to result
		result = append(result, ppdm_group_object)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmGroupObject(c *fiber.Ctx) error {
	var ppdm_group_object dto.Ppdm_group_object

	setDefaults(&ppdm_group_object)

	if err := json.Unmarshal(c.Body(), &ppdm_group_object); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_group_object values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42)")
	if err != nil {
		return err
	}
	ppdm_group_object.Row_created_date = formatDate(ppdm_group_object.Row_created_date)
	ppdm_group_object.Row_changed_date = formatDate(ppdm_group_object.Row_changed_date)
	ppdm_group_object.Effective_date = formatDateString(ppdm_group_object.Effective_date)
	ppdm_group_object.Expiry_date = formatDateString(ppdm_group_object.Expiry_date)
	ppdm_group_object.Row_effective_date = formatDateString(ppdm_group_object.Row_effective_date)
	ppdm_group_object.Row_expiry_date = formatDateString(ppdm_group_object.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_group_object.Group_id, ppdm_group_object.Object_obs_no, ppdm_group_object.Active_ind, ppdm_group_object.Code_version_obs_no, ppdm_group_object.Code_version_source, ppdm_group_object.Column_alias_id, ppdm_group_object.Column_name, ppdm_group_object.Constraint_name, ppdm_group_object.Effective_date, ppdm_group_object.Expiry_date, ppdm_group_object.Group_use, ppdm_group_object.Index_id, ppdm_group_object.Object_type, ppdm_group_object.Output_font, ppdm_group_object.Output_font_backgr_color, ppdm_group_object.Output_font_color, ppdm_group_object.Output_font_effect, ppdm_group_object.Output_font_size, ppdm_group_object.Output_font_size_uom, ppdm_group_object.Output_heading, ppdm_group_object.Output_length, ppdm_group_object.Output_length_uom, ppdm_group_object.Output_precision, ppdm_group_object.Ppdm_guid, ppdm_group_object.Preferred_uom, ppdm_group_object.Procedure_id, ppdm_group_object.Property_set_id, ppdm_group_object.Remark, ppdm_group_object.Rule_id, ppdm_group_object.Source, ppdm_group_object.Sw_application_id, ppdm_group_object.System_id, ppdm_group_object.Table_alias, ppdm_group_object.Table_name, ppdm_group_object.Use_rule_description, ppdm_group_object.Row_changed_by, ppdm_group_object.Row_changed_date, ppdm_group_object.Row_created_by, ppdm_group_object.Row_created_date, ppdm_group_object.Row_effective_date, ppdm_group_object.Row_expiry_date, ppdm_group_object.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmGroupObject(c *fiber.Ctx) error {
	var ppdm_group_object dto.Ppdm_group_object

	setDefaults(&ppdm_group_object)

	if err := json.Unmarshal(c.Body(), &ppdm_group_object); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_group_object.Group_id = id

    if ppdm_group_object.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_group_object set object_obs_no = :1, active_ind = :2, code_version_obs_no = :3, code_version_source = :4, column_alias_id = :5, column_name = :6, constraint_name = :7, effective_date = :8, expiry_date = :9, group_use = :10, index_id = :11, object_type = :12, output_font = :13, output_font_backgr_color = :14, output_font_color = :15, output_font_effect = :16, output_font_size = :17, output_font_size_uom = :18, output_heading = :19, output_length = :20, output_length_uom = :21, output_precision = :22, ppdm_guid = :23, preferred_uom = :24, procedure_id = :25, property_set_id = :26, remark = :27, rule_id = :28, source = :29, sw_application_id = :30, system_id = :31, table_alias = :32, table_name = :33, use_rule_description = :34, row_changed_by = :35, row_changed_date = :36, row_created_by = :37, row_effective_date = :38, row_expiry_date = :39, row_quality = :40 where group_id = :42")
	if err != nil {
		return err
	}

	ppdm_group_object.Row_changed_date = formatDate(ppdm_group_object.Row_changed_date)
	ppdm_group_object.Effective_date = formatDateString(ppdm_group_object.Effective_date)
	ppdm_group_object.Expiry_date = formatDateString(ppdm_group_object.Expiry_date)
	ppdm_group_object.Row_effective_date = formatDateString(ppdm_group_object.Row_effective_date)
	ppdm_group_object.Row_expiry_date = formatDateString(ppdm_group_object.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_group_object.Object_obs_no, ppdm_group_object.Active_ind, ppdm_group_object.Code_version_obs_no, ppdm_group_object.Code_version_source, ppdm_group_object.Column_alias_id, ppdm_group_object.Column_name, ppdm_group_object.Constraint_name, ppdm_group_object.Effective_date, ppdm_group_object.Expiry_date, ppdm_group_object.Group_use, ppdm_group_object.Index_id, ppdm_group_object.Object_type, ppdm_group_object.Output_font, ppdm_group_object.Output_font_backgr_color, ppdm_group_object.Output_font_color, ppdm_group_object.Output_font_effect, ppdm_group_object.Output_font_size, ppdm_group_object.Output_font_size_uom, ppdm_group_object.Output_heading, ppdm_group_object.Output_length, ppdm_group_object.Output_length_uom, ppdm_group_object.Output_precision, ppdm_group_object.Ppdm_guid, ppdm_group_object.Preferred_uom, ppdm_group_object.Procedure_id, ppdm_group_object.Property_set_id, ppdm_group_object.Remark, ppdm_group_object.Rule_id, ppdm_group_object.Source, ppdm_group_object.Sw_application_id, ppdm_group_object.System_id, ppdm_group_object.Table_alias, ppdm_group_object.Table_name, ppdm_group_object.Use_rule_description, ppdm_group_object.Row_changed_by, ppdm_group_object.Row_changed_date, ppdm_group_object.Row_created_by, ppdm_group_object.Row_effective_date, ppdm_group_object.Row_expiry_date, ppdm_group_object.Row_quality, ppdm_group_object.Group_id)
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

func PatchPpdmGroupObject(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_group_object set "
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
	query += " where group_id = :id"

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

func DeletePpdmGroupObject(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_group_object dto.Ppdm_group_object
	ppdm_group_object.Group_id = id

	stmt, err := db.Prepare("delete from ppdm_group_object where group_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_group_object.Group_id)
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



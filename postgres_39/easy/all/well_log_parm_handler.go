package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellLogParm(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_log_parm")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_log_parm

	for rows.Next() {
		var well_log_parm dto.Well_log_parm
		if err := rows.Scan(&well_log_parm.Uwi, &well_log_parm.Well_log_id, &well_log_parm.Well_log_source, &well_log_parm.Parameter_seq_no, &well_log_parm.Active_ind, &well_log_parm.Array_ind, &well_log_parm.Description, &well_log_parm.Dictionary_id, &well_log_parm.Dict_parameter_id, &well_log_parm.Effective_date, &well_log_parm.Expiry_date, &well_log_parm.Parameter_output, &well_log_parm.Parameter_text_value, &well_log_parm.Parameter_value, &well_log_parm.Parameter_value_ouom, &well_log_parm.Parameter_value_uom, &well_log_parm.Ppdm_guid, &well_log_parm.Remark, &well_log_parm.Reported_desc, &well_log_parm.Reported_mnemonic, &well_log_parm.Source, &well_log_parm.Row_changed_by, &well_log_parm.Row_changed_date, &well_log_parm.Row_created_by, &well_log_parm.Row_created_date, &well_log_parm.Row_effective_date, &well_log_parm.Row_expiry_date, &well_log_parm.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_log_parm to result
		result = append(result, well_log_parm)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellLogParm(c *fiber.Ctx) error {
	var well_log_parm dto.Well_log_parm

	setDefaults(&well_log_parm)

	if err := json.Unmarshal(c.Body(), &well_log_parm); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_log_parm values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28)")
	if err != nil {
		return err
	}
	well_log_parm.Row_created_date = formatDate(well_log_parm.Row_created_date)
	well_log_parm.Row_changed_date = formatDate(well_log_parm.Row_changed_date)
	well_log_parm.Effective_date = formatDateString(well_log_parm.Effective_date)
	well_log_parm.Expiry_date = formatDateString(well_log_parm.Expiry_date)
	well_log_parm.Row_effective_date = formatDateString(well_log_parm.Row_effective_date)
	well_log_parm.Row_expiry_date = formatDateString(well_log_parm.Row_expiry_date)






	rows, err := stmt.Exec(well_log_parm.Uwi, well_log_parm.Well_log_id, well_log_parm.Well_log_source, well_log_parm.Parameter_seq_no, well_log_parm.Active_ind, well_log_parm.Array_ind, well_log_parm.Description, well_log_parm.Dictionary_id, well_log_parm.Dict_parameter_id, well_log_parm.Effective_date, well_log_parm.Expiry_date, well_log_parm.Parameter_output, well_log_parm.Parameter_text_value, well_log_parm.Parameter_value, well_log_parm.Parameter_value_ouom, well_log_parm.Parameter_value_uom, well_log_parm.Ppdm_guid, well_log_parm.Remark, well_log_parm.Reported_desc, well_log_parm.Reported_mnemonic, well_log_parm.Source, well_log_parm.Row_changed_by, well_log_parm.Row_changed_date, well_log_parm.Row_created_by, well_log_parm.Row_created_date, well_log_parm.Row_effective_date, well_log_parm.Row_expiry_date, well_log_parm.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellLogParm(c *fiber.Ctx) error {
	var well_log_parm dto.Well_log_parm

	setDefaults(&well_log_parm)

	if err := json.Unmarshal(c.Body(), &well_log_parm); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_log_parm.Uwi = id

    if well_log_parm.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_log_parm set well_log_id = :1, well_log_source = :2, parameter_seq_no = :3, active_ind = :4, array_ind = :5, description = :6, dictionary_id = :7, dict_parameter_id = :8, effective_date = :9, expiry_date = :10, parameter_output = :11, parameter_text_value = :12, parameter_value = :13, parameter_value_ouom = :14, parameter_value_uom = :15, ppdm_guid = :16, remark = :17, reported_desc = :18, reported_mnemonic = :19, source = :20, row_changed_by = :21, row_changed_date = :22, row_created_by = :23, row_effective_date = :24, row_expiry_date = :25, row_quality = :26 where uwi = :28")
	if err != nil {
		return err
	}

	well_log_parm.Row_changed_date = formatDate(well_log_parm.Row_changed_date)
	well_log_parm.Effective_date = formatDateString(well_log_parm.Effective_date)
	well_log_parm.Expiry_date = formatDateString(well_log_parm.Expiry_date)
	well_log_parm.Row_effective_date = formatDateString(well_log_parm.Row_effective_date)
	well_log_parm.Row_expiry_date = formatDateString(well_log_parm.Row_expiry_date)






	rows, err := stmt.Exec(well_log_parm.Well_log_id, well_log_parm.Well_log_source, well_log_parm.Parameter_seq_no, well_log_parm.Active_ind, well_log_parm.Array_ind, well_log_parm.Description, well_log_parm.Dictionary_id, well_log_parm.Dict_parameter_id, well_log_parm.Effective_date, well_log_parm.Expiry_date, well_log_parm.Parameter_output, well_log_parm.Parameter_text_value, well_log_parm.Parameter_value, well_log_parm.Parameter_value_ouom, well_log_parm.Parameter_value_uom, well_log_parm.Ppdm_guid, well_log_parm.Remark, well_log_parm.Reported_desc, well_log_parm.Reported_mnemonic, well_log_parm.Source, well_log_parm.Row_changed_by, well_log_parm.Row_changed_date, well_log_parm.Row_created_by, well_log_parm.Row_effective_date, well_log_parm.Row_expiry_date, well_log_parm.Row_quality, well_log_parm.Uwi)
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

func PatchWellLogParm(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_log_parm set "
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
	query += " where uwi = :id"

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

func DeleteWellLogParm(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_log_parm dto.Well_log_parm
	well_log_parm.Uwi = id

	stmt, err := db.Prepare("delete from well_log_parm where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_log_parm.Uwi)
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



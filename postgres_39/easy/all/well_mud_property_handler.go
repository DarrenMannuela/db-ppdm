package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellMudProperty(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_mud_property")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_mud_property

	for rows.Next() {
		var well_mud_property dto.Well_mud_property
		if err := rows.Scan(&well_mud_property.Sample_analysis_id, &well_mud_property.Sample_analysis_source, &well_mud_property.Property_obs_no, &well_mud_property.Active_ind, &well_mud_property.Average_value, &well_mud_property.Average_value_ouom, &well_mud_property.Average_value_uom, &well_mud_property.Date_format_desc, &well_mud_property.Effective_date, &well_mud_property.Expiry_date, &well_mud_property.Max_value, &well_mud_property.Max_value_ouom, &well_mud_property.Max_value_uom, &well_mud_property.Min_value, &well_mud_property.Min_value_ouom, &well_mud_property.Min_value_uom, &well_mud_property.Mud_anl_obs_no, &well_mud_property.Mud_property, &well_mud_property.Mud_property_code, &well_mud_property.Mud_property_desc, &well_mud_property.Mud_property_ref, &well_mud_property.Ppdm_guid, &well_mud_property.Reference_value, &well_mud_property.Reference_value_ouom, &well_mud_property.Reference_value_uom, &well_mud_property.Remark, &well_mud_property.Sample_type, &well_mud_property.Source, &well_mud_property.Step_seq_no, &well_mud_property.Uwi, &well_mud_property.Row_changed_by, &well_mud_property.Row_changed_date, &well_mud_property.Row_created_by, &well_mud_property.Row_created_date, &well_mud_property.Row_effective_date, &well_mud_property.Row_expiry_date, &well_mud_property.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_mud_property to result
		result = append(result, well_mud_property)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellMudProperty(c *fiber.Ctx) error {
	var well_mud_property dto.Well_mud_property

	setDefaults(&well_mud_property)

	if err := json.Unmarshal(c.Body(), &well_mud_property); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_mud_property values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37)")
	if err != nil {
		return err
	}
	well_mud_property.Row_created_date = formatDate(well_mud_property.Row_created_date)
	well_mud_property.Row_changed_date = formatDate(well_mud_property.Row_changed_date)
	well_mud_property.Date_format_desc = formatDateString(well_mud_property.Date_format_desc)
	well_mud_property.Effective_date = formatDateString(well_mud_property.Effective_date)
	well_mud_property.Expiry_date = formatDateString(well_mud_property.Expiry_date)
	well_mud_property.Row_effective_date = formatDateString(well_mud_property.Row_effective_date)
	well_mud_property.Row_expiry_date = formatDateString(well_mud_property.Row_expiry_date)






	rows, err := stmt.Exec(well_mud_property.Sample_analysis_id, well_mud_property.Sample_analysis_source, well_mud_property.Property_obs_no, well_mud_property.Active_ind, well_mud_property.Average_value, well_mud_property.Average_value_ouom, well_mud_property.Average_value_uom, well_mud_property.Date_format_desc, well_mud_property.Effective_date, well_mud_property.Expiry_date, well_mud_property.Max_value, well_mud_property.Max_value_ouom, well_mud_property.Max_value_uom, well_mud_property.Min_value, well_mud_property.Min_value_ouom, well_mud_property.Min_value_uom, well_mud_property.Mud_anl_obs_no, well_mud_property.Mud_property, well_mud_property.Mud_property_code, well_mud_property.Mud_property_desc, well_mud_property.Mud_property_ref, well_mud_property.Ppdm_guid, well_mud_property.Reference_value, well_mud_property.Reference_value_ouom, well_mud_property.Reference_value_uom, well_mud_property.Remark, well_mud_property.Sample_type, well_mud_property.Source, well_mud_property.Step_seq_no, well_mud_property.Uwi, well_mud_property.Row_changed_by, well_mud_property.Row_changed_date, well_mud_property.Row_created_by, well_mud_property.Row_created_date, well_mud_property.Row_effective_date, well_mud_property.Row_expiry_date, well_mud_property.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellMudProperty(c *fiber.Ctx) error {
	var well_mud_property dto.Well_mud_property

	setDefaults(&well_mud_property)

	if err := json.Unmarshal(c.Body(), &well_mud_property); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_mud_property.Sample_analysis_id = id

    if well_mud_property.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_mud_property set sample_analysis_source = :1, property_obs_no = :2, active_ind = :3, average_value = :4, average_value_ouom = :5, average_value_uom = :6, date_format_desc = :7, effective_date = :8, expiry_date = :9, max_value = :10, max_value_ouom = :11, max_value_uom = :12, min_value = :13, min_value_ouom = :14, min_value_uom = :15, mud_anl_obs_no = :16, mud_property = :17, mud_property_code = :18, mud_property_desc = :19, mud_property_ref = :20, ppdm_guid = :21, reference_value = :22, reference_value_ouom = :23, reference_value_uom = :24, remark = :25, sample_type = :26, source = :27, step_seq_no = :28, uwi = :29, row_changed_by = :30, row_changed_date = :31, row_created_by = :32, row_effective_date = :33, row_expiry_date = :34, row_quality = :35 where sample_analysis_id = :37")
	if err != nil {
		return err
	}

	well_mud_property.Row_changed_date = formatDate(well_mud_property.Row_changed_date)
	well_mud_property.Date_format_desc = formatDateString(well_mud_property.Date_format_desc)
	well_mud_property.Effective_date = formatDateString(well_mud_property.Effective_date)
	well_mud_property.Expiry_date = formatDateString(well_mud_property.Expiry_date)
	well_mud_property.Row_effective_date = formatDateString(well_mud_property.Row_effective_date)
	well_mud_property.Row_expiry_date = formatDateString(well_mud_property.Row_expiry_date)






	rows, err := stmt.Exec(well_mud_property.Sample_analysis_source, well_mud_property.Property_obs_no, well_mud_property.Active_ind, well_mud_property.Average_value, well_mud_property.Average_value_ouom, well_mud_property.Average_value_uom, well_mud_property.Date_format_desc, well_mud_property.Effective_date, well_mud_property.Expiry_date, well_mud_property.Max_value, well_mud_property.Max_value_ouom, well_mud_property.Max_value_uom, well_mud_property.Min_value, well_mud_property.Min_value_ouom, well_mud_property.Min_value_uom, well_mud_property.Mud_anl_obs_no, well_mud_property.Mud_property, well_mud_property.Mud_property_code, well_mud_property.Mud_property_desc, well_mud_property.Mud_property_ref, well_mud_property.Ppdm_guid, well_mud_property.Reference_value, well_mud_property.Reference_value_ouom, well_mud_property.Reference_value_uom, well_mud_property.Remark, well_mud_property.Sample_type, well_mud_property.Source, well_mud_property.Step_seq_no, well_mud_property.Uwi, well_mud_property.Row_changed_by, well_mud_property.Row_changed_date, well_mud_property.Row_created_by, well_mud_property.Row_effective_date, well_mud_property.Row_expiry_date, well_mud_property.Row_quality, well_mud_property.Sample_analysis_id)
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

func PatchWellMudProperty(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_mud_property set "
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
	query += " where sample_analysis_id = :id"

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

func DeleteWellMudProperty(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_mud_property dto.Well_mud_property
	well_mud_property.Sample_analysis_id = id

	stmt, err := db.Prepare("delete from well_mud_property where sample_analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_mud_property.Sample_analysis_id)
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



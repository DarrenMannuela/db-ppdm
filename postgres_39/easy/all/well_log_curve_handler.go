package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellLogCurve(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_log_curve")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_log_curve

	for rows.Next() {
		var well_log_curve dto.Well_log_curve
		if err := rows.Scan(&well_log_curve.Uwi, &well_log_curve.Curve_id, &well_log_curve.Acquired_for_ba_id, &well_log_curve.Active_ind, &well_log_curve.Api_code_system, &well_log_curve.Api_curve_class, &well_log_curve.Api_curve_code, &well_log_curve.Api_curve_modifier, &well_log_curve.Api_log_code, &well_log_curve.Base_curve_ind, &well_log_curve.Bypass_ind, &well_log_curve.Cased_hole_ind, &well_log_curve.Composite_ind, &well_log_curve.Curve_ouom, &well_log_curve.Curve_output_type, &well_log_curve.Curve_quality, &well_log_curve.Date_format_desc, &well_log_curve.Dictionary_id, &well_log_curve.Dict_curve_id, &well_log_curve.Effective_date, &well_log_curve.Expiry_date, &well_log_curve.Explicit_index_ind, &well_log_curve.First_good_value, &well_log_curve.First_good_value_index, &well_log_curve.Frame_id, &well_log_curve.Good_value_type, &well_log_curve.Index_curve_id, &well_log_curve.Index_ouom, &well_log_curve.Index_uom, &well_log_curve.Job_id, &well_log_curve.Last_good_value, &well_log_curve.Last_good_value_index, &well_log_curve.Log_tool_pass_no, &well_log_curve.Log_tool_type, &well_log_curve.Max_index, &well_log_curve.Max_value, &well_log_curve.Max_value_index, &well_log_curve.Mean_value, &well_log_curve.Mean_value_std_dev, &well_log_curve.Min_index, &well_log_curve.Min_value, &well_log_curve.Min_value_index, &well_log_curve.Multiple_index_ind, &well_log_curve.Mwd_ind, &well_log_curve.Null_count, &well_log_curve.Null_representation, &well_log_curve.Order_seq_no, &well_log_curve.Ppdm_guid, &well_log_curve.Preferred_ind, &well_log_curve.Primary_index_type, &well_log_curve.Remark, &well_log_curve.Reported_desc, &well_log_curve.Reported_mnemonic, &well_log_curve.Reported_unit_mnemonic, &well_log_curve.Source, &well_log_curve.Trip_obs_no, &well_log_curve.Value_count, &well_log_curve.Well_log_id, &well_log_curve.Well_log_job_source, &well_log_curve.Well_log_source, &well_log_curve.Row_changed_by, &well_log_curve.Row_changed_date, &well_log_curve.Row_created_by, &well_log_curve.Row_created_date, &well_log_curve.Row_effective_date, &well_log_curve.Row_expiry_date, &well_log_curve.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_log_curve to result
		result = append(result, well_log_curve)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellLogCurve(c *fiber.Ctx) error {
	var well_log_curve dto.Well_log_curve

	setDefaults(&well_log_curve)

	if err := json.Unmarshal(c.Body(), &well_log_curve); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_log_curve values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67)")
	if err != nil {
		return err
	}
	well_log_curve.Row_created_date = formatDate(well_log_curve.Row_created_date)
	well_log_curve.Row_changed_date = formatDate(well_log_curve.Row_changed_date)
	well_log_curve.Date_format_desc = formatDateString(well_log_curve.Date_format_desc)
	well_log_curve.Effective_date = formatDateString(well_log_curve.Effective_date)
	well_log_curve.Expiry_date = formatDateString(well_log_curve.Expiry_date)
	well_log_curve.Row_effective_date = formatDateString(well_log_curve.Row_effective_date)
	well_log_curve.Row_expiry_date = formatDateString(well_log_curve.Row_expiry_date)






	rows, err := stmt.Exec(well_log_curve.Uwi, well_log_curve.Curve_id, well_log_curve.Acquired_for_ba_id, well_log_curve.Active_ind, well_log_curve.Api_code_system, well_log_curve.Api_curve_class, well_log_curve.Api_curve_code, well_log_curve.Api_curve_modifier, well_log_curve.Api_log_code, well_log_curve.Base_curve_ind, well_log_curve.Bypass_ind, well_log_curve.Cased_hole_ind, well_log_curve.Composite_ind, well_log_curve.Curve_ouom, well_log_curve.Curve_output_type, well_log_curve.Curve_quality, well_log_curve.Date_format_desc, well_log_curve.Dictionary_id, well_log_curve.Dict_curve_id, well_log_curve.Effective_date, well_log_curve.Expiry_date, well_log_curve.Explicit_index_ind, well_log_curve.First_good_value, well_log_curve.First_good_value_index, well_log_curve.Frame_id, well_log_curve.Good_value_type, well_log_curve.Index_curve_id, well_log_curve.Index_ouom, well_log_curve.Index_uom, well_log_curve.Job_id, well_log_curve.Last_good_value, well_log_curve.Last_good_value_index, well_log_curve.Log_tool_pass_no, well_log_curve.Log_tool_type, well_log_curve.Max_index, well_log_curve.Max_value, well_log_curve.Max_value_index, well_log_curve.Mean_value, well_log_curve.Mean_value_std_dev, well_log_curve.Min_index, well_log_curve.Min_value, well_log_curve.Min_value_index, well_log_curve.Multiple_index_ind, well_log_curve.Mwd_ind, well_log_curve.Null_count, well_log_curve.Null_representation, well_log_curve.Order_seq_no, well_log_curve.Ppdm_guid, well_log_curve.Preferred_ind, well_log_curve.Primary_index_type, well_log_curve.Remark, well_log_curve.Reported_desc, well_log_curve.Reported_mnemonic, well_log_curve.Reported_unit_mnemonic, well_log_curve.Source, well_log_curve.Trip_obs_no, well_log_curve.Value_count, well_log_curve.Well_log_id, well_log_curve.Well_log_job_source, well_log_curve.Well_log_source, well_log_curve.Row_changed_by, well_log_curve.Row_changed_date, well_log_curve.Row_created_by, well_log_curve.Row_created_date, well_log_curve.Row_effective_date, well_log_curve.Row_expiry_date, well_log_curve.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellLogCurve(c *fiber.Ctx) error {
	var well_log_curve dto.Well_log_curve

	setDefaults(&well_log_curve)

	if err := json.Unmarshal(c.Body(), &well_log_curve); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_log_curve.Uwi = id

    if well_log_curve.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_log_curve set curve_id = :1, acquired_for_ba_id = :2, active_ind = :3, api_code_system = :4, api_curve_class = :5, api_curve_code = :6, api_curve_modifier = :7, api_log_code = :8, base_curve_ind = :9, bypass_ind = :10, cased_hole_ind = :11, composite_ind = :12, curve_ouom = :13, curve_output_type = :14, curve_quality = :15, date_format_desc = :16, dictionary_id = :17, dict_curve_id = :18, effective_date = :19, expiry_date = :20, explicit_index_ind = :21, first_good_value = :22, first_good_value_index = :23, frame_id = :24, good_value_type = :25, index_curve_id = :26, index_ouom = :27, index_uom = :28, job_id = :29, last_good_value = :30, last_good_value_index = :31, log_tool_pass_no = :32, log_tool_type = :33, max_index = :34, max_value = :35, max_value_index = :36, mean_value = :37, mean_value_std_dev = :38, min_index = :39, min_value = :40, min_value_index = :41, multiple_index_ind = :42, mwd_ind = :43, null_count = :44, null_representation = :45, order_seq_no = :46, ppdm_guid = :47, preferred_ind = :48, primary_index_type = :49, remark = :50, reported_desc = :51, reported_mnemonic = :52, reported_unit_mnemonic = :53, source = :54, trip_obs_no = :55, value_count = :56, well_log_id = :57, well_log_job_source = :58, well_log_source = :59, row_changed_by = :60, row_changed_date = :61, row_created_by = :62, row_effective_date = :63, row_expiry_date = :64, row_quality = :65 where uwi = :67")
	if err != nil {
		return err
	}

	well_log_curve.Row_changed_date = formatDate(well_log_curve.Row_changed_date)
	well_log_curve.Date_format_desc = formatDateString(well_log_curve.Date_format_desc)
	well_log_curve.Effective_date = formatDateString(well_log_curve.Effective_date)
	well_log_curve.Expiry_date = formatDateString(well_log_curve.Expiry_date)
	well_log_curve.Row_effective_date = formatDateString(well_log_curve.Row_effective_date)
	well_log_curve.Row_expiry_date = formatDateString(well_log_curve.Row_expiry_date)






	rows, err := stmt.Exec(well_log_curve.Curve_id, well_log_curve.Acquired_for_ba_id, well_log_curve.Active_ind, well_log_curve.Api_code_system, well_log_curve.Api_curve_class, well_log_curve.Api_curve_code, well_log_curve.Api_curve_modifier, well_log_curve.Api_log_code, well_log_curve.Base_curve_ind, well_log_curve.Bypass_ind, well_log_curve.Cased_hole_ind, well_log_curve.Composite_ind, well_log_curve.Curve_ouom, well_log_curve.Curve_output_type, well_log_curve.Curve_quality, well_log_curve.Date_format_desc, well_log_curve.Dictionary_id, well_log_curve.Dict_curve_id, well_log_curve.Effective_date, well_log_curve.Expiry_date, well_log_curve.Explicit_index_ind, well_log_curve.First_good_value, well_log_curve.First_good_value_index, well_log_curve.Frame_id, well_log_curve.Good_value_type, well_log_curve.Index_curve_id, well_log_curve.Index_ouom, well_log_curve.Index_uom, well_log_curve.Job_id, well_log_curve.Last_good_value, well_log_curve.Last_good_value_index, well_log_curve.Log_tool_pass_no, well_log_curve.Log_tool_type, well_log_curve.Max_index, well_log_curve.Max_value, well_log_curve.Max_value_index, well_log_curve.Mean_value, well_log_curve.Mean_value_std_dev, well_log_curve.Min_index, well_log_curve.Min_value, well_log_curve.Min_value_index, well_log_curve.Multiple_index_ind, well_log_curve.Mwd_ind, well_log_curve.Null_count, well_log_curve.Null_representation, well_log_curve.Order_seq_no, well_log_curve.Ppdm_guid, well_log_curve.Preferred_ind, well_log_curve.Primary_index_type, well_log_curve.Remark, well_log_curve.Reported_desc, well_log_curve.Reported_mnemonic, well_log_curve.Reported_unit_mnemonic, well_log_curve.Source, well_log_curve.Trip_obs_no, well_log_curve.Value_count, well_log_curve.Well_log_id, well_log_curve.Well_log_job_source, well_log_curve.Well_log_source, well_log_curve.Row_changed_by, well_log_curve.Row_changed_date, well_log_curve.Row_created_by, well_log_curve.Row_effective_date, well_log_curve.Row_expiry_date, well_log_curve.Row_quality, well_log_curve.Uwi)
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

func PatchWellLogCurve(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_log_curve set "
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

func DeleteWellLogCurve(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_log_curve dto.Well_log_curve
	well_log_curve.Uwi = id

	stmt, err := db.Prepare("delete from well_log_curve where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_log_curve.Uwi)
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



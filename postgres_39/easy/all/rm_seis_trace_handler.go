package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRmSeisTrace(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rm_seis_trace")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rm_seis_trace

	for rows.Next() {
		var rm_seis_trace dto.Rm_seis_trace
		if err := rows.Scan(&rm_seis_trace.Info_item_subtype, &rm_seis_trace.Information_item_id, &rm_seis_trace.Active_ind, &rm_seis_trace.Datum_vel_correction, &rm_seis_trace.Datum_vel_correction_ouom, &rm_seis_trace.Effective_date, &rm_seis_trace.Expiry_date, &rm_seis_trace.First_sample_time, &rm_seis_trace.First_sample_timezone, &rm_seis_trace.Horizontal_scale, &rm_seis_trace.Horizontal_scale_ouom, &rm_seis_trace.Phase, &rm_seis_trace.Polarity, &rm_seis_trace.Ppdm_guid, &rm_seis_trace.Record_length, &rm_seis_trace.Record_length_ouom, &rm_seis_trace.Reference_datum, &rm_seis_trace.Reference_datum_ouom, &rm_seis_trace.Remark, &rm_seis_trace.Replacement_vel, &rm_seis_trace.Replacement_vel_ouom, &rm_seis_trace.Reported_first_ref_num, &rm_seis_trace.Reported_last_ref_num, &rm_seis_trace.Reported_ref_num_type, &rm_seis_trace.Sample_rate, &rm_seis_trace.Sample_rate_ouom, &rm_seis_trace.Sample_type, &rm_seis_trace.Source, &rm_seis_trace.Static_correction, &rm_seis_trace.Static_correction_ouom, &rm_seis_trace.Variable_area_ind, &rm_seis_trace.Vertical_scale, &rm_seis_trace.Vertical_scale_ouom, &rm_seis_trace.Row_changed_by, &rm_seis_trace.Row_changed_date, &rm_seis_trace.Row_created_by, &rm_seis_trace.Row_created_date, &rm_seis_trace.Row_effective_date, &rm_seis_trace.Row_expiry_date, &rm_seis_trace.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rm_seis_trace to result
		result = append(result, rm_seis_trace)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRmSeisTrace(c *fiber.Ctx) error {
	var rm_seis_trace dto.Rm_seis_trace

	setDefaults(&rm_seis_trace)

	if err := json.Unmarshal(c.Body(), &rm_seis_trace); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rm_seis_trace values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40)")
	if err != nil {
		return err
	}
	rm_seis_trace.Row_created_date = formatDate(rm_seis_trace.Row_created_date)
	rm_seis_trace.Row_changed_date = formatDate(rm_seis_trace.Row_changed_date)
	rm_seis_trace.Effective_date = formatDateString(rm_seis_trace.Effective_date)
	rm_seis_trace.Expiry_date = formatDateString(rm_seis_trace.Expiry_date)
	rm_seis_trace.Row_effective_date = formatDateString(rm_seis_trace.Row_effective_date)
	rm_seis_trace.Row_expiry_date = formatDateString(rm_seis_trace.Row_expiry_date)






	rows, err := stmt.Exec(rm_seis_trace.Info_item_subtype, rm_seis_trace.Information_item_id, rm_seis_trace.Active_ind, rm_seis_trace.Datum_vel_correction, rm_seis_trace.Datum_vel_correction_ouom, rm_seis_trace.Effective_date, rm_seis_trace.Expiry_date, rm_seis_trace.First_sample_time, rm_seis_trace.First_sample_timezone, rm_seis_trace.Horizontal_scale, rm_seis_trace.Horizontal_scale_ouom, rm_seis_trace.Phase, rm_seis_trace.Polarity, rm_seis_trace.Ppdm_guid, rm_seis_trace.Record_length, rm_seis_trace.Record_length_ouom, rm_seis_trace.Reference_datum, rm_seis_trace.Reference_datum_ouom, rm_seis_trace.Remark, rm_seis_trace.Replacement_vel, rm_seis_trace.Replacement_vel_ouom, rm_seis_trace.Reported_first_ref_num, rm_seis_trace.Reported_last_ref_num, rm_seis_trace.Reported_ref_num_type, rm_seis_trace.Sample_rate, rm_seis_trace.Sample_rate_ouom, rm_seis_trace.Sample_type, rm_seis_trace.Source, rm_seis_trace.Static_correction, rm_seis_trace.Static_correction_ouom, rm_seis_trace.Variable_area_ind, rm_seis_trace.Vertical_scale, rm_seis_trace.Vertical_scale_ouom, rm_seis_trace.Row_changed_by, rm_seis_trace.Row_changed_date, rm_seis_trace.Row_created_by, rm_seis_trace.Row_created_date, rm_seis_trace.Row_effective_date, rm_seis_trace.Row_expiry_date, rm_seis_trace.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRmSeisTrace(c *fiber.Ctx) error {
	var rm_seis_trace dto.Rm_seis_trace

	setDefaults(&rm_seis_trace)

	if err := json.Unmarshal(c.Body(), &rm_seis_trace); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rm_seis_trace.Info_item_subtype = id

    if rm_seis_trace.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rm_seis_trace set information_item_id = :1, active_ind = :2, datum_vel_correction = :3, datum_vel_correction_ouom = :4, effective_date = :5, expiry_date = :6, first_sample_time = :7, first_sample_timezone = :8, horizontal_scale = :9, horizontal_scale_ouom = :10, phase = :11, polarity = :12, ppdm_guid = :13, record_length = :14, record_length_ouom = :15, reference_datum = :16, reference_datum_ouom = :17, remark = :18, replacement_vel = :19, replacement_vel_ouom = :20, reported_first_ref_num = :21, reported_last_ref_num = :22, reported_ref_num_type = :23, sample_rate = :24, sample_rate_ouom = :25, sample_type = :26, source = :27, static_correction = :28, static_correction_ouom = :29, variable_area_ind = :30, vertical_scale = :31, vertical_scale_ouom = :32, row_changed_by = :33, row_changed_date = :34, row_created_by = :35, row_effective_date = :36, row_expiry_date = :37, row_quality = :38 where info_item_subtype = :40")
	if err != nil {
		return err
	}

	rm_seis_trace.Row_changed_date = formatDate(rm_seis_trace.Row_changed_date)
	rm_seis_trace.Effective_date = formatDateString(rm_seis_trace.Effective_date)
	rm_seis_trace.Expiry_date = formatDateString(rm_seis_trace.Expiry_date)
	rm_seis_trace.Row_effective_date = formatDateString(rm_seis_trace.Row_effective_date)
	rm_seis_trace.Row_expiry_date = formatDateString(rm_seis_trace.Row_expiry_date)






	rows, err := stmt.Exec(rm_seis_trace.Information_item_id, rm_seis_trace.Active_ind, rm_seis_trace.Datum_vel_correction, rm_seis_trace.Datum_vel_correction_ouom, rm_seis_trace.Effective_date, rm_seis_trace.Expiry_date, rm_seis_trace.First_sample_time, rm_seis_trace.First_sample_timezone, rm_seis_trace.Horizontal_scale, rm_seis_trace.Horizontal_scale_ouom, rm_seis_trace.Phase, rm_seis_trace.Polarity, rm_seis_trace.Ppdm_guid, rm_seis_trace.Record_length, rm_seis_trace.Record_length_ouom, rm_seis_trace.Reference_datum, rm_seis_trace.Reference_datum_ouom, rm_seis_trace.Remark, rm_seis_trace.Replacement_vel, rm_seis_trace.Replacement_vel_ouom, rm_seis_trace.Reported_first_ref_num, rm_seis_trace.Reported_last_ref_num, rm_seis_trace.Reported_ref_num_type, rm_seis_trace.Sample_rate, rm_seis_trace.Sample_rate_ouom, rm_seis_trace.Sample_type, rm_seis_trace.Source, rm_seis_trace.Static_correction, rm_seis_trace.Static_correction_ouom, rm_seis_trace.Variable_area_ind, rm_seis_trace.Vertical_scale, rm_seis_trace.Vertical_scale_ouom, rm_seis_trace.Row_changed_by, rm_seis_trace.Row_changed_date, rm_seis_trace.Row_created_by, rm_seis_trace.Row_effective_date, rm_seis_trace.Row_expiry_date, rm_seis_trace.Row_quality, rm_seis_trace.Info_item_subtype)
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

func PatchRmSeisTrace(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rm_seis_trace set "
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
	query += " where info_item_subtype = :id"

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

func DeleteRmSeisTrace(c *fiber.Ctx) error {
	id := c.Params("id")
	var rm_seis_trace dto.Rm_seis_trace
	rm_seis_trace.Info_item_subtype = id

	stmt, err := db.Prepare("delete from rm_seis_trace where info_item_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rm_seis_trace.Info_item_subtype)
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



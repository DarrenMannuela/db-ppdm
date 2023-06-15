package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPdenFlowMeasurement(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pden_flow_measurement")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pden_flow_measurement

	for rows.Next() {
		var pden_flow_measurement dto.Pden_flow_measurement
		if err := rows.Scan(&pden_flow_measurement.Pden_id, &pden_flow_measurement.Pden_subtype, &pden_flow_measurement.Pden_source, &pden_flow_measurement.Product_type, &pden_flow_measurement.Amendment_seq_no, &pden_flow_measurement.Period_type, &pden_flow_measurement.Measurement_obs_no, &pden_flow_measurement.Active_ind, &pden_flow_measurement.Amend_reason, &pden_flow_measurement.Casing_pressure, &pden_flow_measurement.Casing_pressure_ouom, &pden_flow_measurement.Choke_position, &pden_flow_measurement.Choke_size, &pden_flow_measurement.Choke_size_ouom, &pden_flow_measurement.Differential_pressure, &pden_flow_measurement.Diff_pressure_ouom, &pden_flow_measurement.Effective_date, &pden_flow_measurement.Expiry_date, &pden_flow_measurement.Flow_rate, &pden_flow_measurement.Flow_rate_ouom, &pden_flow_measurement.Measurement_date, &pden_flow_measurement.Measurement_date_desc, &pden_flow_measurement.Measurement_time, &pden_flow_measurement.Measurement_timezone, &pden_flow_measurement.Measurement_type, &pden_flow_measurement.Meas_temperature, &pden_flow_measurement.Meas_temperature_ouom, &pden_flow_measurement.Posted_date, &pden_flow_measurement.Ppdm_guid, &pden_flow_measurement.Production_volume, &pden_flow_measurement.Production_volume_ouom, &pden_flow_measurement.Production_volume_uom, &pden_flow_measurement.Prod_time_elapsed, &pden_flow_measurement.Prod_time_elapsed_ouom, &pden_flow_measurement.Remark, &pden_flow_measurement.Source, &pden_flow_measurement.Static_pressure, &pden_flow_measurement.Static_pressure_ouom, &pden_flow_measurement.Tubing_pressure, &pden_flow_measurement.Tubing_pressure_ouom, &pden_flow_measurement.Row_changed_by, &pden_flow_measurement.Row_changed_date, &pden_flow_measurement.Row_created_by, &pden_flow_measurement.Row_created_date, &pden_flow_measurement.Row_effective_date, &pden_flow_measurement.Row_expiry_date, &pden_flow_measurement.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pden_flow_measurement to result
		result = append(result, pden_flow_measurement)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPdenFlowMeasurement(c *fiber.Ctx) error {
	var pden_flow_measurement dto.Pden_flow_measurement

	setDefaults(&pden_flow_measurement)

	if err := json.Unmarshal(c.Body(), &pden_flow_measurement); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pden_flow_measurement values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47)")
	if err != nil {
		return err
	}
	pden_flow_measurement.Row_created_date = formatDate(pden_flow_measurement.Row_created_date)
	pden_flow_measurement.Row_changed_date = formatDate(pden_flow_measurement.Row_changed_date)
	pden_flow_measurement.Effective_date = formatDateString(pden_flow_measurement.Effective_date)
	pden_flow_measurement.Expiry_date = formatDateString(pden_flow_measurement.Expiry_date)
	pden_flow_measurement.Measurement_date = formatDateString(pden_flow_measurement.Measurement_date)
	pden_flow_measurement.Posted_date = formatDateString(pden_flow_measurement.Posted_date)
	pden_flow_measurement.Row_effective_date = formatDateString(pden_flow_measurement.Row_effective_date)
	pden_flow_measurement.Row_expiry_date = formatDateString(pden_flow_measurement.Row_expiry_date)






	rows, err := stmt.Exec(pden_flow_measurement.Pden_id, pden_flow_measurement.Pden_subtype, pden_flow_measurement.Pden_source, pden_flow_measurement.Product_type, pden_flow_measurement.Amendment_seq_no, pden_flow_measurement.Period_type, pden_flow_measurement.Measurement_obs_no, pden_flow_measurement.Active_ind, pden_flow_measurement.Amend_reason, pden_flow_measurement.Casing_pressure, pden_flow_measurement.Casing_pressure_ouom, pden_flow_measurement.Choke_position, pden_flow_measurement.Choke_size, pden_flow_measurement.Choke_size_ouom, pden_flow_measurement.Differential_pressure, pden_flow_measurement.Diff_pressure_ouom, pden_flow_measurement.Effective_date, pden_flow_measurement.Expiry_date, pden_flow_measurement.Flow_rate, pden_flow_measurement.Flow_rate_ouom, pden_flow_measurement.Measurement_date, pden_flow_measurement.Measurement_date_desc, pden_flow_measurement.Measurement_time, pden_flow_measurement.Measurement_timezone, pden_flow_measurement.Measurement_type, pden_flow_measurement.Meas_temperature, pden_flow_measurement.Meas_temperature_ouom, pden_flow_measurement.Posted_date, pden_flow_measurement.Ppdm_guid, pden_flow_measurement.Production_volume, pden_flow_measurement.Production_volume_ouom, pden_flow_measurement.Production_volume_uom, pden_flow_measurement.Prod_time_elapsed, pden_flow_measurement.Prod_time_elapsed_ouom, pden_flow_measurement.Remark, pden_flow_measurement.Source, pden_flow_measurement.Static_pressure, pden_flow_measurement.Static_pressure_ouom, pden_flow_measurement.Tubing_pressure, pden_flow_measurement.Tubing_pressure_ouom, pden_flow_measurement.Row_changed_by, pden_flow_measurement.Row_changed_date, pden_flow_measurement.Row_created_by, pden_flow_measurement.Row_created_date, pden_flow_measurement.Row_effective_date, pden_flow_measurement.Row_expiry_date, pden_flow_measurement.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePdenFlowMeasurement(c *fiber.Ctx) error {
	var pden_flow_measurement dto.Pden_flow_measurement

	setDefaults(&pden_flow_measurement)

	if err := json.Unmarshal(c.Body(), &pden_flow_measurement); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pden_flow_measurement.Pden_id = id

    if pden_flow_measurement.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pden_flow_measurement set pden_subtype = :1, pden_source = :2, product_type = :3, amendment_seq_no = :4, period_type = :5, measurement_obs_no = :6, active_ind = :7, amend_reason = :8, casing_pressure = :9, casing_pressure_ouom = :10, choke_position = :11, choke_size = :12, choke_size_ouom = :13, differential_pressure = :14, diff_pressure_ouom = :15, effective_date = :16, expiry_date = :17, flow_rate = :18, flow_rate_ouom = :19, measurement_date = :20, measurement_date_desc = :21, measurement_time = :22, measurement_timezone = :23, measurement_type = :24, meas_temperature = :25, meas_temperature_ouom = :26, posted_date = :27, ppdm_guid = :28, production_volume = :29, production_volume_ouom = :30, production_volume_uom = :31, prod_time_elapsed = :32, prod_time_elapsed_ouom = :33, remark = :34, source = :35, static_pressure = :36, static_pressure_ouom = :37, tubing_pressure = :38, tubing_pressure_ouom = :39, row_changed_by = :40, row_changed_date = :41, row_created_by = :42, row_effective_date = :43, row_expiry_date = :44, row_quality = :45 where pden_id = :47")
	if err != nil {
		return err
	}

	pden_flow_measurement.Row_changed_date = formatDate(pden_flow_measurement.Row_changed_date)
	pden_flow_measurement.Effective_date = formatDateString(pden_flow_measurement.Effective_date)
	pden_flow_measurement.Expiry_date = formatDateString(pden_flow_measurement.Expiry_date)
	pden_flow_measurement.Measurement_date = formatDateString(pden_flow_measurement.Measurement_date)
	pden_flow_measurement.Posted_date = formatDateString(pden_flow_measurement.Posted_date)
	pden_flow_measurement.Row_effective_date = formatDateString(pden_flow_measurement.Row_effective_date)
	pden_flow_measurement.Row_expiry_date = formatDateString(pden_flow_measurement.Row_expiry_date)






	rows, err := stmt.Exec(pden_flow_measurement.Pden_subtype, pden_flow_measurement.Pden_source, pden_flow_measurement.Product_type, pden_flow_measurement.Amendment_seq_no, pden_flow_measurement.Period_type, pden_flow_measurement.Measurement_obs_no, pden_flow_measurement.Active_ind, pden_flow_measurement.Amend_reason, pden_flow_measurement.Casing_pressure, pden_flow_measurement.Casing_pressure_ouom, pden_flow_measurement.Choke_position, pden_flow_measurement.Choke_size, pden_flow_measurement.Choke_size_ouom, pden_flow_measurement.Differential_pressure, pden_flow_measurement.Diff_pressure_ouom, pden_flow_measurement.Effective_date, pden_flow_measurement.Expiry_date, pden_flow_measurement.Flow_rate, pden_flow_measurement.Flow_rate_ouom, pden_flow_measurement.Measurement_date, pden_flow_measurement.Measurement_date_desc, pden_flow_measurement.Measurement_time, pden_flow_measurement.Measurement_timezone, pden_flow_measurement.Measurement_type, pden_flow_measurement.Meas_temperature, pden_flow_measurement.Meas_temperature_ouom, pden_flow_measurement.Posted_date, pden_flow_measurement.Ppdm_guid, pden_flow_measurement.Production_volume, pden_flow_measurement.Production_volume_ouom, pden_flow_measurement.Production_volume_uom, pden_flow_measurement.Prod_time_elapsed, pden_flow_measurement.Prod_time_elapsed_ouom, pden_flow_measurement.Remark, pden_flow_measurement.Source, pden_flow_measurement.Static_pressure, pden_flow_measurement.Static_pressure_ouom, pden_flow_measurement.Tubing_pressure, pden_flow_measurement.Tubing_pressure_ouom, pden_flow_measurement.Row_changed_by, pden_flow_measurement.Row_changed_date, pden_flow_measurement.Row_created_by, pden_flow_measurement.Row_effective_date, pden_flow_measurement.Row_expiry_date, pden_flow_measurement.Row_quality, pden_flow_measurement.Pden_id)
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

func PatchPdenFlowMeasurement(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pden_flow_measurement set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "measurement_date" || key == "posted_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeletePdenFlowMeasurement(c *fiber.Ctx) error {
	id := c.Params("id")
	var pden_flow_measurement dto.Pden_flow_measurement
	pden_flow_measurement.Pden_id = id

	stmt, err := db.Prepare("delete from pden_flow_measurement where pden_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pden_flow_measurement.Pden_id)
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



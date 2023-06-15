package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellPressureBh(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_pressure_bh")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_pressure_bh

	for rows.Next() {
		var well_pressure_bh dto.Well_pressure_bh
		if err := rows.Scan(&well_pressure_bh.Uwi, &well_pressure_bh.Source, &well_pressure_bh.Pressure_obs_no, &well_pressure_bh.Bhp_obs_no, &well_pressure_bh.Active_ind, &well_pressure_bh.Bhfp, &well_pressure_bh.Bhfp_ouom, &well_pressure_bh.Bhp_method, &well_pressure_bh.Bh_test_code, &well_pressure_bh.Datum_pressure, &well_pressure_bh.Datum_pressure_ouom, &well_pressure_bh.Effective_date, &well_pressure_bh.Expiry_date, &well_pressure_bh.Packer_depth, &well_pressure_bh.Packer_depth_ouom, &well_pressure_bh.Ppdm_guid, &well_pressure_bh.Pressure_gradient, &well_pressure_bh.Pressure_gradient_ouom, &well_pressure_bh.Recorder_datum, &well_pressure_bh.Recorder_datum_ouom, &well_pressure_bh.Remark, &well_pressure_bh.Reported_run_tvd, &well_pressure_bh.Reported_run_tvd_ouom, &well_pressure_bh.Run_depth, &well_pressure_bh.Run_depth_ouom, &well_pressure_bh.Run_depth_temperature, &well_pressure_bh.Run_depth_temperature_ouom, &well_pressure_bh.Shutin_period, &well_pressure_bh.Shutin_period_ouom, &well_pressure_bh.Survey_date, &well_pressure_bh.Test_duration, &well_pressure_bh.Test_duration_ouom, &well_pressure_bh.Tubing_size_desc, &well_pressure_bh.Well_head_pressure, &well_pressure_bh.Well_head_pressure_ouom, &well_pressure_bh.Row_changed_by, &well_pressure_bh.Row_changed_date, &well_pressure_bh.Row_created_by, &well_pressure_bh.Row_created_date, &well_pressure_bh.Row_effective_date, &well_pressure_bh.Row_expiry_date, &well_pressure_bh.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_pressure_bh to result
		result = append(result, well_pressure_bh)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellPressureBh(c *fiber.Ctx) error {
	var well_pressure_bh dto.Well_pressure_bh

	setDefaults(&well_pressure_bh)

	if err := json.Unmarshal(c.Body(), &well_pressure_bh); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_pressure_bh values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42)")
	if err != nil {
		return err
	}
	well_pressure_bh.Row_created_date = formatDate(well_pressure_bh.Row_created_date)
	well_pressure_bh.Row_changed_date = formatDate(well_pressure_bh.Row_changed_date)
	well_pressure_bh.Effective_date = formatDateString(well_pressure_bh.Effective_date)
	well_pressure_bh.Expiry_date = formatDateString(well_pressure_bh.Expiry_date)
	well_pressure_bh.Survey_date = formatDateString(well_pressure_bh.Survey_date)
	well_pressure_bh.Row_effective_date = formatDateString(well_pressure_bh.Row_effective_date)
	well_pressure_bh.Row_expiry_date = formatDateString(well_pressure_bh.Row_expiry_date)






	rows, err := stmt.Exec(well_pressure_bh.Uwi, well_pressure_bh.Source, well_pressure_bh.Pressure_obs_no, well_pressure_bh.Bhp_obs_no, well_pressure_bh.Active_ind, well_pressure_bh.Bhfp, well_pressure_bh.Bhfp_ouom, well_pressure_bh.Bhp_method, well_pressure_bh.Bh_test_code, well_pressure_bh.Datum_pressure, well_pressure_bh.Datum_pressure_ouom, well_pressure_bh.Effective_date, well_pressure_bh.Expiry_date, well_pressure_bh.Packer_depth, well_pressure_bh.Packer_depth_ouom, well_pressure_bh.Ppdm_guid, well_pressure_bh.Pressure_gradient, well_pressure_bh.Pressure_gradient_ouom, well_pressure_bh.Recorder_datum, well_pressure_bh.Recorder_datum_ouom, well_pressure_bh.Remark, well_pressure_bh.Reported_run_tvd, well_pressure_bh.Reported_run_tvd_ouom, well_pressure_bh.Run_depth, well_pressure_bh.Run_depth_ouom, well_pressure_bh.Run_depth_temperature, well_pressure_bh.Run_depth_temperature_ouom, well_pressure_bh.Shutin_period, well_pressure_bh.Shutin_period_ouom, well_pressure_bh.Survey_date, well_pressure_bh.Test_duration, well_pressure_bh.Test_duration_ouom, well_pressure_bh.Tubing_size_desc, well_pressure_bh.Well_head_pressure, well_pressure_bh.Well_head_pressure_ouom, well_pressure_bh.Row_changed_by, well_pressure_bh.Row_changed_date, well_pressure_bh.Row_created_by, well_pressure_bh.Row_created_date, well_pressure_bh.Row_effective_date, well_pressure_bh.Row_expiry_date, well_pressure_bh.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellPressureBh(c *fiber.Ctx) error {
	var well_pressure_bh dto.Well_pressure_bh

	setDefaults(&well_pressure_bh)

	if err := json.Unmarshal(c.Body(), &well_pressure_bh); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_pressure_bh.Uwi = id

    if well_pressure_bh.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_pressure_bh set source = :1, pressure_obs_no = :2, bhp_obs_no = :3, active_ind = :4, bhfp = :5, bhfp_ouom = :6, bhp_method = :7, bh_test_code = :8, datum_pressure = :9, datum_pressure_ouom = :10, effective_date = :11, expiry_date = :12, packer_depth = :13, packer_depth_ouom = :14, ppdm_guid = :15, pressure_gradient = :16, pressure_gradient_ouom = :17, recorder_datum = :18, recorder_datum_ouom = :19, remark = :20, reported_run_tvd = :21, reported_run_tvd_ouom = :22, run_depth = :23, run_depth_ouom = :24, run_depth_temperature = :25, run_depth_temperature_ouom = :26, shutin_period = :27, shutin_period_ouom = :28, survey_date = :29, test_duration = :30, test_duration_ouom = :31, tubing_size_desc = :32, well_head_pressure = :33, well_head_pressure_ouom = :34, row_changed_by = :35, row_changed_date = :36, row_created_by = :37, row_effective_date = :38, row_expiry_date = :39, row_quality = :40 where uwi = :42")
	if err != nil {
		return err
	}

	well_pressure_bh.Row_changed_date = formatDate(well_pressure_bh.Row_changed_date)
	well_pressure_bh.Effective_date = formatDateString(well_pressure_bh.Effective_date)
	well_pressure_bh.Expiry_date = formatDateString(well_pressure_bh.Expiry_date)
	well_pressure_bh.Survey_date = formatDateString(well_pressure_bh.Survey_date)
	well_pressure_bh.Row_effective_date = formatDateString(well_pressure_bh.Row_effective_date)
	well_pressure_bh.Row_expiry_date = formatDateString(well_pressure_bh.Row_expiry_date)






	rows, err := stmt.Exec(well_pressure_bh.Source, well_pressure_bh.Pressure_obs_no, well_pressure_bh.Bhp_obs_no, well_pressure_bh.Active_ind, well_pressure_bh.Bhfp, well_pressure_bh.Bhfp_ouom, well_pressure_bh.Bhp_method, well_pressure_bh.Bh_test_code, well_pressure_bh.Datum_pressure, well_pressure_bh.Datum_pressure_ouom, well_pressure_bh.Effective_date, well_pressure_bh.Expiry_date, well_pressure_bh.Packer_depth, well_pressure_bh.Packer_depth_ouom, well_pressure_bh.Ppdm_guid, well_pressure_bh.Pressure_gradient, well_pressure_bh.Pressure_gradient_ouom, well_pressure_bh.Recorder_datum, well_pressure_bh.Recorder_datum_ouom, well_pressure_bh.Remark, well_pressure_bh.Reported_run_tvd, well_pressure_bh.Reported_run_tvd_ouom, well_pressure_bh.Run_depth, well_pressure_bh.Run_depth_ouom, well_pressure_bh.Run_depth_temperature, well_pressure_bh.Run_depth_temperature_ouom, well_pressure_bh.Shutin_period, well_pressure_bh.Shutin_period_ouom, well_pressure_bh.Survey_date, well_pressure_bh.Test_duration, well_pressure_bh.Test_duration_ouom, well_pressure_bh.Tubing_size_desc, well_pressure_bh.Well_head_pressure, well_pressure_bh.Well_head_pressure_ouom, well_pressure_bh.Row_changed_by, well_pressure_bh.Row_changed_date, well_pressure_bh.Row_created_by, well_pressure_bh.Row_effective_date, well_pressure_bh.Row_expiry_date, well_pressure_bh.Row_quality, well_pressure_bh.Uwi)
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

func PatchWellPressureBh(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_pressure_bh set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "survey_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellPressureBh(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_pressure_bh dto.Well_pressure_bh
	well_pressure_bh.Uwi = id

	stmt, err := db.Prepare("delete from well_pressure_bh where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_pressure_bh.Uwi)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellLogTrip(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_log_trip")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_log_trip

	for rows.Next() {
		var well_log_trip dto.Well_log_trip
		if err := rows.Scan(&well_log_trip.Uwi, &well_log_trip.Source, &well_log_trip.Job_id, &well_log_trip.Trip_obs_no, &well_log_trip.Active_ind, &well_log_trip.Base_depth, &well_log_trip.Base_depth_ouom, &well_log_trip.Base_strat_unit_id, &well_log_trip.Effective_date, &well_log_trip.Expiry_date, &well_log_trip.Logging_service_type, &well_log_trip.Max_depth, &well_log_trip.Max_depth_ouom, &well_log_trip.Max_temperature, &well_log_trip.Max_temperature_ouom, &well_log_trip.Mud_sample_id, &well_log_trip.Mud_sample_type, &well_log_trip.Mud_source, &well_log_trip.Observer, &well_log_trip.On_bottom_date, &well_log_trip.On_bottom_time, &well_log_trip.On_bottom_timezone, &well_log_trip.Ppdm_guid, &well_log_trip.Remark, &well_log_trip.Reported_tvd, &well_log_trip.Reported_tvd_ouom, &well_log_trip.Report_apd, &well_log_trip.Report_log_datum, &well_log_trip.Report_log_datum_elev, &well_log_trip.Report_log_datum_elev_ouom, &well_log_trip.Report_log_run, &well_log_trip.Report_perm_datum, &well_log_trip.Report_perm_datum_elev, &well_log_trip.Report_perm_datum_elev_ouom, &well_log_trip.Strat_name_set_id, &well_log_trip.Top_depth, &well_log_trip.Top_depth_ouom, &well_log_trip.Top_strat_unit_id, &well_log_trip.Trip_date, &well_log_trip.Tubing_bottom_depth, &well_log_trip.Tubing_bottom_depth_ouom, &well_log_trip.Row_changed_by, &well_log_trip.Row_changed_date, &well_log_trip.Row_created_by, &well_log_trip.Row_created_date, &well_log_trip.Row_effective_date, &well_log_trip.Row_expiry_date, &well_log_trip.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_log_trip to result
		result = append(result, well_log_trip)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellLogTrip(c *fiber.Ctx) error {
	var well_log_trip dto.Well_log_trip

	setDefaults(&well_log_trip)

	if err := json.Unmarshal(c.Body(), &well_log_trip); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_log_trip values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48)")
	if err != nil {
		return err
	}
	well_log_trip.Row_created_date = formatDate(well_log_trip.Row_created_date)
	well_log_trip.Row_changed_date = formatDate(well_log_trip.Row_changed_date)
	well_log_trip.Effective_date = formatDateString(well_log_trip.Effective_date)
	well_log_trip.Expiry_date = formatDateString(well_log_trip.Expiry_date)
	well_log_trip.On_bottom_date = formatDateString(well_log_trip.On_bottom_date)
	well_log_trip.On_bottom_time = formatDateString(well_log_trip.On_bottom_time)
	well_log_trip.Trip_date = formatDateString(well_log_trip.Trip_date)
	well_log_trip.Row_effective_date = formatDateString(well_log_trip.Row_effective_date)
	well_log_trip.Row_expiry_date = formatDateString(well_log_trip.Row_expiry_date)






	rows, err := stmt.Exec(well_log_trip.Uwi, well_log_trip.Source, well_log_trip.Job_id, well_log_trip.Trip_obs_no, well_log_trip.Active_ind, well_log_trip.Base_depth, well_log_trip.Base_depth_ouom, well_log_trip.Base_strat_unit_id, well_log_trip.Effective_date, well_log_trip.Expiry_date, well_log_trip.Logging_service_type, well_log_trip.Max_depth, well_log_trip.Max_depth_ouom, well_log_trip.Max_temperature, well_log_trip.Max_temperature_ouom, well_log_trip.Mud_sample_id, well_log_trip.Mud_sample_type, well_log_trip.Mud_source, well_log_trip.Observer, well_log_trip.On_bottom_date, well_log_trip.On_bottom_time, well_log_trip.On_bottom_timezone, well_log_trip.Ppdm_guid, well_log_trip.Remark, well_log_trip.Reported_tvd, well_log_trip.Reported_tvd_ouom, well_log_trip.Report_apd, well_log_trip.Report_log_datum, well_log_trip.Report_log_datum_elev, well_log_trip.Report_log_datum_elev_ouom, well_log_trip.Report_log_run, well_log_trip.Report_perm_datum, well_log_trip.Report_perm_datum_elev, well_log_trip.Report_perm_datum_elev_ouom, well_log_trip.Strat_name_set_id, well_log_trip.Top_depth, well_log_trip.Top_depth_ouom, well_log_trip.Top_strat_unit_id, well_log_trip.Trip_date, well_log_trip.Tubing_bottom_depth, well_log_trip.Tubing_bottom_depth_ouom, well_log_trip.Row_changed_by, well_log_trip.Row_changed_date, well_log_trip.Row_created_by, well_log_trip.Row_created_date, well_log_trip.Row_effective_date, well_log_trip.Row_expiry_date, well_log_trip.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellLogTrip(c *fiber.Ctx) error {
	var well_log_trip dto.Well_log_trip

	setDefaults(&well_log_trip)

	if err := json.Unmarshal(c.Body(), &well_log_trip); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_log_trip.Uwi = id

    if well_log_trip.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_log_trip set source = :1, job_id = :2, trip_obs_no = :3, active_ind = :4, base_depth = :5, base_depth_ouom = :6, base_strat_unit_id = :7, effective_date = :8, expiry_date = :9, logging_service_type = :10, max_depth = :11, max_depth_ouom = :12, max_temperature = :13, max_temperature_ouom = :14, mud_sample_id = :15, mud_sample_type = :16, mud_source = :17, observer = :18, on_bottom_date = :19, on_bottom_time = :20, on_bottom_timezone = :21, ppdm_guid = :22, remark = :23, reported_tvd = :24, reported_tvd_ouom = :25, report_apd = :26, report_log_datum = :27, report_log_datum_elev = :28, report_log_datum_elev_ouom = :29, report_log_run = :30, report_perm_datum = :31, report_perm_datum_elev = :32, report_perm_datum_elev_ouom = :33, strat_name_set_id = :34, top_depth = :35, top_depth_ouom = :36, top_strat_unit_id = :37, trip_date = :38, tubing_bottom_depth = :39, tubing_bottom_depth_ouom = :40, row_changed_by = :41, row_changed_date = :42, row_created_by = :43, row_effective_date = :44, row_expiry_date = :45, row_quality = :46 where uwi = :48")
	if err != nil {
		return err
	}

	well_log_trip.Row_changed_date = formatDate(well_log_trip.Row_changed_date)
	well_log_trip.Effective_date = formatDateString(well_log_trip.Effective_date)
	well_log_trip.Expiry_date = formatDateString(well_log_trip.Expiry_date)
	well_log_trip.On_bottom_date = formatDateString(well_log_trip.On_bottom_date)
	well_log_trip.On_bottom_time = formatDateString(well_log_trip.On_bottom_time)
	well_log_trip.Trip_date = formatDateString(well_log_trip.Trip_date)
	well_log_trip.Row_effective_date = formatDateString(well_log_trip.Row_effective_date)
	well_log_trip.Row_expiry_date = formatDateString(well_log_trip.Row_expiry_date)






	rows, err := stmt.Exec(well_log_trip.Source, well_log_trip.Job_id, well_log_trip.Trip_obs_no, well_log_trip.Active_ind, well_log_trip.Base_depth, well_log_trip.Base_depth_ouom, well_log_trip.Base_strat_unit_id, well_log_trip.Effective_date, well_log_trip.Expiry_date, well_log_trip.Logging_service_type, well_log_trip.Max_depth, well_log_trip.Max_depth_ouom, well_log_trip.Max_temperature, well_log_trip.Max_temperature_ouom, well_log_trip.Mud_sample_id, well_log_trip.Mud_sample_type, well_log_trip.Mud_source, well_log_trip.Observer, well_log_trip.On_bottom_date, well_log_trip.On_bottom_time, well_log_trip.On_bottom_timezone, well_log_trip.Ppdm_guid, well_log_trip.Remark, well_log_trip.Reported_tvd, well_log_trip.Reported_tvd_ouom, well_log_trip.Report_apd, well_log_trip.Report_log_datum, well_log_trip.Report_log_datum_elev, well_log_trip.Report_log_datum_elev_ouom, well_log_trip.Report_log_run, well_log_trip.Report_perm_datum, well_log_trip.Report_perm_datum_elev, well_log_trip.Report_perm_datum_elev_ouom, well_log_trip.Strat_name_set_id, well_log_trip.Top_depth, well_log_trip.Top_depth_ouom, well_log_trip.Top_strat_unit_id, well_log_trip.Trip_date, well_log_trip.Tubing_bottom_depth, well_log_trip.Tubing_bottom_depth_ouom, well_log_trip.Row_changed_by, well_log_trip.Row_changed_date, well_log_trip.Row_created_by, well_log_trip.Row_effective_date, well_log_trip.Row_expiry_date, well_log_trip.Row_quality, well_log_trip.Uwi)
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

func PatchWellLogTrip(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_log_trip set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "on_bottom_date" || key == "on_bottom_time" || key == "trip_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellLogTrip(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_log_trip dto.Well_log_trip
	well_log_trip.Uwi = id

	stmt, err := db.Prepare("delete from well_log_trip where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_log_trip.Uwi)
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



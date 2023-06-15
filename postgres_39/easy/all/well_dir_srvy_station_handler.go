package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellDirSrvyStation(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_dir_srvy_station")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_dir_srvy_station

	for rows.Next() {
		var well_dir_srvy_station dto.Well_dir_srvy_station
		if err := rows.Scan(&well_dir_srvy_station.Uwi, &well_dir_srvy_station.Survey_id, &well_dir_srvy_station.Source, &well_dir_srvy_station.Depth_obs_no, &well_dir_srvy_station.Accuracy_problem_ind, &well_dir_srvy_station.Accuracy_problem_reason, &well_dir_srvy_station.Active_ind, &well_dir_srvy_station.Azimuth, &well_dir_srvy_station.Azimuth_ouom, &well_dir_srvy_station.Closure_dist, &well_dir_srvy_station.Closure_dist_dir, &well_dir_srvy_station.Closure_dist_dir_ouom, &well_dir_srvy_station.Closure_dist_ouom, &well_dir_srvy_station.Delta_lat, &well_dir_srvy_station.Delta_long, &well_dir_srvy_station.Dog_leg_severity, &well_dir_srvy_station.Effective_date, &well_dir_srvy_station.Expiry_date, &well_dir_srvy_station.Inclination, &well_dir_srvy_station.Inclination_ouom, &well_dir_srvy_station.Latitude, &well_dir_srvy_station.Longitude, &well_dir_srvy_station.Period_obs_no, &well_dir_srvy_station.Point_type, &well_dir_srvy_station.Ppdm_guid, &well_dir_srvy_station.Rad_uncert_dist, &well_dir_srvy_station.Rad_uncert_dist_ouom, &well_dir_srvy_station.Rad_uncert_dist_reason, &well_dir_srvy_station.Remark, &well_dir_srvy_station.Rpt_azimuth, &well_dir_srvy_station.Rpt_azimuth_ouom, &well_dir_srvy_station.Station_id, &well_dir_srvy_station.Station_md, &well_dir_srvy_station.Station_md_ouom, &well_dir_srvy_station.Station_tvd, &well_dir_srvy_station.Station_tvdss, &well_dir_srvy_station.Station_tvdss_ouom, &well_dir_srvy_station.Station_tvd_ouom, &well_dir_srvy_station.Survey_run_num, &well_dir_srvy_station.Survey_type, &well_dir_srvy_station.Vertical_section, &well_dir_srvy_station.Vertical_section_ouom, &well_dir_srvy_station.X_offset, &well_dir_srvy_station.X_offset_ew_direction, &well_dir_srvy_station.X_offset_ouom, &well_dir_srvy_station.Y_offset, &well_dir_srvy_station.Y_offset_ns_direction, &well_dir_srvy_station.Y_offset_ouom, &well_dir_srvy_station.Row_changed_by, &well_dir_srvy_station.Row_changed_date, &well_dir_srvy_station.Row_created_by, &well_dir_srvy_station.Row_created_date, &well_dir_srvy_station.Row_effective_date, &well_dir_srvy_station.Row_expiry_date, &well_dir_srvy_station.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_dir_srvy_station to result
		result = append(result, well_dir_srvy_station)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellDirSrvyStation(c *fiber.Ctx) error {
	var well_dir_srvy_station dto.Well_dir_srvy_station

	setDefaults(&well_dir_srvy_station)

	if err := json.Unmarshal(c.Body(), &well_dir_srvy_station); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_dir_srvy_station values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55)")
	if err != nil {
		return err
	}
	well_dir_srvy_station.Row_created_date = formatDate(well_dir_srvy_station.Row_created_date)
	well_dir_srvy_station.Row_changed_date = formatDate(well_dir_srvy_station.Row_changed_date)
	well_dir_srvy_station.Effective_date = formatDateString(well_dir_srvy_station.Effective_date)
	well_dir_srvy_station.Expiry_date = formatDateString(well_dir_srvy_station.Expiry_date)
	well_dir_srvy_station.Row_effective_date = formatDateString(well_dir_srvy_station.Row_effective_date)
	well_dir_srvy_station.Row_expiry_date = formatDateString(well_dir_srvy_station.Row_expiry_date)






	rows, err := stmt.Exec(well_dir_srvy_station.Uwi, well_dir_srvy_station.Survey_id, well_dir_srvy_station.Source, well_dir_srvy_station.Depth_obs_no, well_dir_srvy_station.Accuracy_problem_ind, well_dir_srvy_station.Accuracy_problem_reason, well_dir_srvy_station.Active_ind, well_dir_srvy_station.Azimuth, well_dir_srvy_station.Azimuth_ouom, well_dir_srvy_station.Closure_dist, well_dir_srvy_station.Closure_dist_dir, well_dir_srvy_station.Closure_dist_dir_ouom, well_dir_srvy_station.Closure_dist_ouom, well_dir_srvy_station.Delta_lat, well_dir_srvy_station.Delta_long, well_dir_srvy_station.Dog_leg_severity, well_dir_srvy_station.Effective_date, well_dir_srvy_station.Expiry_date, well_dir_srvy_station.Inclination, well_dir_srvy_station.Inclination_ouom, well_dir_srvy_station.Latitude, well_dir_srvy_station.Longitude, well_dir_srvy_station.Period_obs_no, well_dir_srvy_station.Point_type, well_dir_srvy_station.Ppdm_guid, well_dir_srvy_station.Rad_uncert_dist, well_dir_srvy_station.Rad_uncert_dist_ouom, well_dir_srvy_station.Rad_uncert_dist_reason, well_dir_srvy_station.Remark, well_dir_srvy_station.Rpt_azimuth, well_dir_srvy_station.Rpt_azimuth_ouom, well_dir_srvy_station.Station_id, well_dir_srvy_station.Station_md, well_dir_srvy_station.Station_md_ouom, well_dir_srvy_station.Station_tvd, well_dir_srvy_station.Station_tvdss, well_dir_srvy_station.Station_tvdss_ouom, well_dir_srvy_station.Station_tvd_ouom, well_dir_srvy_station.Survey_run_num, well_dir_srvy_station.Survey_type, well_dir_srvy_station.Vertical_section, well_dir_srvy_station.Vertical_section_ouom, well_dir_srvy_station.X_offset, well_dir_srvy_station.X_offset_ew_direction, well_dir_srvy_station.X_offset_ouom, well_dir_srvy_station.Y_offset, well_dir_srvy_station.Y_offset_ns_direction, well_dir_srvy_station.Y_offset_ouom, well_dir_srvy_station.Row_changed_by, well_dir_srvy_station.Row_changed_date, well_dir_srvy_station.Row_created_by, well_dir_srvy_station.Row_created_date, well_dir_srvy_station.Row_effective_date, well_dir_srvy_station.Row_expiry_date, well_dir_srvy_station.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellDirSrvyStation(c *fiber.Ctx) error {
	var well_dir_srvy_station dto.Well_dir_srvy_station

	setDefaults(&well_dir_srvy_station)

	if err := json.Unmarshal(c.Body(), &well_dir_srvy_station); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_dir_srvy_station.Uwi = id

    if well_dir_srvy_station.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_dir_srvy_station set survey_id = :1, source = :2, depth_obs_no = :3, accuracy_problem_ind = :4, accuracy_problem_reason = :5, active_ind = :6, azimuth = :7, azimuth_ouom = :8, closure_dist = :9, closure_dist_dir = :10, closure_dist_dir_ouom = :11, closure_dist_ouom = :12, delta_lat = :13, delta_long = :14, dog_leg_severity = :15, effective_date = :16, expiry_date = :17, inclination = :18, inclination_ouom = :19, latitude = :20, longitude = :21, period_obs_no = :22, point_type = :23, ppdm_guid = :24, rad_uncert_dist = :25, rad_uncert_dist_ouom = :26, rad_uncert_dist_reason = :27, remark = :28, rpt_azimuth = :29, rpt_azimuth_ouom = :30, station_id = :31, station_md = :32, station_md_ouom = :33, station_tvd = :34, station_tvdss = :35, station_tvdss_ouom = :36, station_tvd_ouom = :37, survey_run_num = :38, survey_type = :39, vertical_section = :40, vertical_section_ouom = :41, x_offset = :42, x_offset_ew_direction = :43, x_offset_ouom = :44, y_offset = :45, y_offset_ns_direction = :46, y_offset_ouom = :47, row_changed_by = :48, row_changed_date = :49, row_created_by = :50, row_effective_date = :51, row_expiry_date = :52, row_quality = :53 where uwi = :55")
	if err != nil {
		return err
	}

	well_dir_srvy_station.Row_changed_date = formatDate(well_dir_srvy_station.Row_changed_date)
	well_dir_srvy_station.Effective_date = formatDateString(well_dir_srvy_station.Effective_date)
	well_dir_srvy_station.Expiry_date = formatDateString(well_dir_srvy_station.Expiry_date)
	well_dir_srvy_station.Row_effective_date = formatDateString(well_dir_srvy_station.Row_effective_date)
	well_dir_srvy_station.Row_expiry_date = formatDateString(well_dir_srvy_station.Row_expiry_date)






	rows, err := stmt.Exec(well_dir_srvy_station.Survey_id, well_dir_srvy_station.Source, well_dir_srvy_station.Depth_obs_no, well_dir_srvy_station.Accuracy_problem_ind, well_dir_srvy_station.Accuracy_problem_reason, well_dir_srvy_station.Active_ind, well_dir_srvy_station.Azimuth, well_dir_srvy_station.Azimuth_ouom, well_dir_srvy_station.Closure_dist, well_dir_srvy_station.Closure_dist_dir, well_dir_srvy_station.Closure_dist_dir_ouom, well_dir_srvy_station.Closure_dist_ouom, well_dir_srvy_station.Delta_lat, well_dir_srvy_station.Delta_long, well_dir_srvy_station.Dog_leg_severity, well_dir_srvy_station.Effective_date, well_dir_srvy_station.Expiry_date, well_dir_srvy_station.Inclination, well_dir_srvy_station.Inclination_ouom, well_dir_srvy_station.Latitude, well_dir_srvy_station.Longitude, well_dir_srvy_station.Period_obs_no, well_dir_srvy_station.Point_type, well_dir_srvy_station.Ppdm_guid, well_dir_srvy_station.Rad_uncert_dist, well_dir_srvy_station.Rad_uncert_dist_ouom, well_dir_srvy_station.Rad_uncert_dist_reason, well_dir_srvy_station.Remark, well_dir_srvy_station.Rpt_azimuth, well_dir_srvy_station.Rpt_azimuth_ouom, well_dir_srvy_station.Station_id, well_dir_srvy_station.Station_md, well_dir_srvy_station.Station_md_ouom, well_dir_srvy_station.Station_tvd, well_dir_srvy_station.Station_tvdss, well_dir_srvy_station.Station_tvdss_ouom, well_dir_srvy_station.Station_tvd_ouom, well_dir_srvy_station.Survey_run_num, well_dir_srvy_station.Survey_type, well_dir_srvy_station.Vertical_section, well_dir_srvy_station.Vertical_section_ouom, well_dir_srvy_station.X_offset, well_dir_srvy_station.X_offset_ew_direction, well_dir_srvy_station.X_offset_ouom, well_dir_srvy_station.Y_offset, well_dir_srvy_station.Y_offset_ns_direction, well_dir_srvy_station.Y_offset_ouom, well_dir_srvy_station.Row_changed_by, well_dir_srvy_station.Row_changed_date, well_dir_srvy_station.Row_created_by, well_dir_srvy_station.Row_effective_date, well_dir_srvy_station.Row_expiry_date, well_dir_srvy_station.Row_quality, well_dir_srvy_station.Uwi)
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

func PatchWellDirSrvyStation(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_dir_srvy_station set "
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

func DeleteWellDirSrvyStation(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_dir_srvy_station dto.Well_dir_srvy_station
	well_dir_srvy_station.Uwi = id

	stmt, err := db.Prepare("delete from well_dir_srvy_station where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_dir_srvy_station.Uwi)
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



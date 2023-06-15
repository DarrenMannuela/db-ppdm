package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellDirSrvyStVersion(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_dir_srvy_st_version")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_dir_srvy_st_version

	for rows.Next() {
		var well_dir_srvy_st_version dto.Well_dir_srvy_st_version
		if err := rows.Scan(&well_dir_srvy_st_version.Uwi, &well_dir_srvy_st_version.Survey_id, &well_dir_srvy_st_version.Source, &well_dir_srvy_st_version.Version_obs_no, &well_dir_srvy_st_version.Depth_obs_no, &well_dir_srvy_st_version.Accuracy_problem_ind, &well_dir_srvy_st_version.Accuracy_problem_reason, &well_dir_srvy_st_version.Active_ind, &well_dir_srvy_st_version.Azimuth, &well_dir_srvy_st_version.Azimuth_ouom, &well_dir_srvy_st_version.Closure_dist, &well_dir_srvy_st_version.Closure_dist_dir, &well_dir_srvy_st_version.Closure_dist_dir_ouom, &well_dir_srvy_st_version.Closure_dist_ouom, &well_dir_srvy_st_version.Delta_lat, &well_dir_srvy_st_version.Delta_long, &well_dir_srvy_st_version.Dog_leg_severity, &well_dir_srvy_st_version.Effective_date, &well_dir_srvy_st_version.Expiry_date, &well_dir_srvy_st_version.Inclination, &well_dir_srvy_st_version.Inclination_ouom, &well_dir_srvy_st_version.Latitude, &well_dir_srvy_st_version.Longitude, &well_dir_srvy_st_version.Period_obs_no, &well_dir_srvy_st_version.Point_type, &well_dir_srvy_st_version.Ppdm_guid, &well_dir_srvy_st_version.Rad_uncert_dist, &well_dir_srvy_st_version.Rad_uncert_dist_ouom, &well_dir_srvy_st_version.Rad_uncert_dist_reason, &well_dir_srvy_st_version.Remark, &well_dir_srvy_st_version.Rpt_azimuth, &well_dir_srvy_st_version.Rpt_azimuth_ouom, &well_dir_srvy_st_version.Station_id, &well_dir_srvy_st_version.Station_md, &well_dir_srvy_st_version.Station_md_ouom, &well_dir_srvy_st_version.Station_tvd, &well_dir_srvy_st_version.Station_tvdss, &well_dir_srvy_st_version.Station_tvdss_ouom, &well_dir_srvy_st_version.Station_tvd_ouom, &well_dir_srvy_st_version.Survey_run_num, &well_dir_srvy_st_version.Survey_type, &well_dir_srvy_st_version.Vertical_section, &well_dir_srvy_st_version.Vertical_section_ouom, &well_dir_srvy_st_version.X_offset, &well_dir_srvy_st_version.X_offset_ew_direction, &well_dir_srvy_st_version.X_offset_ouom, &well_dir_srvy_st_version.Y_offset, &well_dir_srvy_st_version.Y_offset_ns_direction, &well_dir_srvy_st_version.Y_offset_ouom, &well_dir_srvy_st_version.Row_changed_by, &well_dir_srvy_st_version.Row_changed_date, &well_dir_srvy_st_version.Row_created_by, &well_dir_srvy_st_version.Row_created_date, &well_dir_srvy_st_version.Row_effective_date, &well_dir_srvy_st_version.Row_expiry_date, &well_dir_srvy_st_version.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_dir_srvy_st_version to result
		result = append(result, well_dir_srvy_st_version)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellDirSrvyStVersion(c *fiber.Ctx) error {
	var well_dir_srvy_st_version dto.Well_dir_srvy_st_version

	setDefaults(&well_dir_srvy_st_version)

	if err := json.Unmarshal(c.Body(), &well_dir_srvy_st_version); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_dir_srvy_st_version values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56)")
	if err != nil {
		return err
	}
	well_dir_srvy_st_version.Row_created_date = formatDate(well_dir_srvy_st_version.Row_created_date)
	well_dir_srvy_st_version.Row_changed_date = formatDate(well_dir_srvy_st_version.Row_changed_date)
	well_dir_srvy_st_version.Effective_date = formatDateString(well_dir_srvy_st_version.Effective_date)
	well_dir_srvy_st_version.Expiry_date = formatDateString(well_dir_srvy_st_version.Expiry_date)
	well_dir_srvy_st_version.Row_effective_date = formatDateString(well_dir_srvy_st_version.Row_effective_date)
	well_dir_srvy_st_version.Row_expiry_date = formatDateString(well_dir_srvy_st_version.Row_expiry_date)






	rows, err := stmt.Exec(well_dir_srvy_st_version.Uwi, well_dir_srvy_st_version.Survey_id, well_dir_srvy_st_version.Source, well_dir_srvy_st_version.Version_obs_no, well_dir_srvy_st_version.Depth_obs_no, well_dir_srvy_st_version.Accuracy_problem_ind, well_dir_srvy_st_version.Accuracy_problem_reason, well_dir_srvy_st_version.Active_ind, well_dir_srvy_st_version.Azimuth, well_dir_srvy_st_version.Azimuth_ouom, well_dir_srvy_st_version.Closure_dist, well_dir_srvy_st_version.Closure_dist_dir, well_dir_srvy_st_version.Closure_dist_dir_ouom, well_dir_srvy_st_version.Closure_dist_ouom, well_dir_srvy_st_version.Delta_lat, well_dir_srvy_st_version.Delta_long, well_dir_srvy_st_version.Dog_leg_severity, well_dir_srvy_st_version.Effective_date, well_dir_srvy_st_version.Expiry_date, well_dir_srvy_st_version.Inclination, well_dir_srvy_st_version.Inclination_ouom, well_dir_srvy_st_version.Latitude, well_dir_srvy_st_version.Longitude, well_dir_srvy_st_version.Period_obs_no, well_dir_srvy_st_version.Point_type, well_dir_srvy_st_version.Ppdm_guid, well_dir_srvy_st_version.Rad_uncert_dist, well_dir_srvy_st_version.Rad_uncert_dist_ouom, well_dir_srvy_st_version.Rad_uncert_dist_reason, well_dir_srvy_st_version.Remark, well_dir_srvy_st_version.Rpt_azimuth, well_dir_srvy_st_version.Rpt_azimuth_ouom, well_dir_srvy_st_version.Station_id, well_dir_srvy_st_version.Station_md, well_dir_srvy_st_version.Station_md_ouom, well_dir_srvy_st_version.Station_tvd, well_dir_srvy_st_version.Station_tvdss, well_dir_srvy_st_version.Station_tvdss_ouom, well_dir_srvy_st_version.Station_tvd_ouom, well_dir_srvy_st_version.Survey_run_num, well_dir_srvy_st_version.Survey_type, well_dir_srvy_st_version.Vertical_section, well_dir_srvy_st_version.Vertical_section_ouom, well_dir_srvy_st_version.X_offset, well_dir_srvy_st_version.X_offset_ew_direction, well_dir_srvy_st_version.X_offset_ouom, well_dir_srvy_st_version.Y_offset, well_dir_srvy_st_version.Y_offset_ns_direction, well_dir_srvy_st_version.Y_offset_ouom, well_dir_srvy_st_version.Row_changed_by, well_dir_srvy_st_version.Row_changed_date, well_dir_srvy_st_version.Row_created_by, well_dir_srvy_st_version.Row_created_date, well_dir_srvy_st_version.Row_effective_date, well_dir_srvy_st_version.Row_expiry_date, well_dir_srvy_st_version.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellDirSrvyStVersion(c *fiber.Ctx) error {
	var well_dir_srvy_st_version dto.Well_dir_srvy_st_version

	setDefaults(&well_dir_srvy_st_version)

	if err := json.Unmarshal(c.Body(), &well_dir_srvy_st_version); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_dir_srvy_st_version.Uwi = id

    if well_dir_srvy_st_version.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_dir_srvy_st_version set survey_id = :1, source = :2, version_obs_no = :3, depth_obs_no = :4, accuracy_problem_ind = :5, accuracy_problem_reason = :6, active_ind = :7, azimuth = :8, azimuth_ouom = :9, closure_dist = :10, closure_dist_dir = :11, closure_dist_dir_ouom = :12, closure_dist_ouom = :13, delta_lat = :14, delta_long = :15, dog_leg_severity = :16, effective_date = :17, expiry_date = :18, inclination = :19, inclination_ouom = :20, latitude = :21, longitude = :22, period_obs_no = :23, point_type = :24, ppdm_guid = :25, rad_uncert_dist = :26, rad_uncert_dist_ouom = :27, rad_uncert_dist_reason = :28, remark = :29, rpt_azimuth = :30, rpt_azimuth_ouom = :31, station_id = :32, station_md = :33, station_md_ouom = :34, station_tvd = :35, station_tvdss = :36, station_tvdss_ouom = :37, station_tvd_ouom = :38, survey_run_num = :39, survey_type = :40, vertical_section = :41, vertical_section_ouom = :42, x_offset = :43, x_offset_ew_direction = :44, x_offset_ouom = :45, y_offset = :46, y_offset_ns_direction = :47, y_offset_ouom = :48, row_changed_by = :49, row_changed_date = :50, row_created_by = :51, row_effective_date = :52, row_expiry_date = :53, row_quality = :54 where uwi = :56")
	if err != nil {
		return err
	}

	well_dir_srvy_st_version.Row_changed_date = formatDate(well_dir_srvy_st_version.Row_changed_date)
	well_dir_srvy_st_version.Effective_date = formatDateString(well_dir_srvy_st_version.Effective_date)
	well_dir_srvy_st_version.Expiry_date = formatDateString(well_dir_srvy_st_version.Expiry_date)
	well_dir_srvy_st_version.Row_effective_date = formatDateString(well_dir_srvy_st_version.Row_effective_date)
	well_dir_srvy_st_version.Row_expiry_date = formatDateString(well_dir_srvy_st_version.Row_expiry_date)






	rows, err := stmt.Exec(well_dir_srvy_st_version.Survey_id, well_dir_srvy_st_version.Source, well_dir_srvy_st_version.Version_obs_no, well_dir_srvy_st_version.Depth_obs_no, well_dir_srvy_st_version.Accuracy_problem_ind, well_dir_srvy_st_version.Accuracy_problem_reason, well_dir_srvy_st_version.Active_ind, well_dir_srvy_st_version.Azimuth, well_dir_srvy_st_version.Azimuth_ouom, well_dir_srvy_st_version.Closure_dist, well_dir_srvy_st_version.Closure_dist_dir, well_dir_srvy_st_version.Closure_dist_dir_ouom, well_dir_srvy_st_version.Closure_dist_ouom, well_dir_srvy_st_version.Delta_lat, well_dir_srvy_st_version.Delta_long, well_dir_srvy_st_version.Dog_leg_severity, well_dir_srvy_st_version.Effective_date, well_dir_srvy_st_version.Expiry_date, well_dir_srvy_st_version.Inclination, well_dir_srvy_st_version.Inclination_ouom, well_dir_srvy_st_version.Latitude, well_dir_srvy_st_version.Longitude, well_dir_srvy_st_version.Period_obs_no, well_dir_srvy_st_version.Point_type, well_dir_srvy_st_version.Ppdm_guid, well_dir_srvy_st_version.Rad_uncert_dist, well_dir_srvy_st_version.Rad_uncert_dist_ouom, well_dir_srvy_st_version.Rad_uncert_dist_reason, well_dir_srvy_st_version.Remark, well_dir_srvy_st_version.Rpt_azimuth, well_dir_srvy_st_version.Rpt_azimuth_ouom, well_dir_srvy_st_version.Station_id, well_dir_srvy_st_version.Station_md, well_dir_srvy_st_version.Station_md_ouom, well_dir_srvy_st_version.Station_tvd, well_dir_srvy_st_version.Station_tvdss, well_dir_srvy_st_version.Station_tvdss_ouom, well_dir_srvy_st_version.Station_tvd_ouom, well_dir_srvy_st_version.Survey_run_num, well_dir_srvy_st_version.Survey_type, well_dir_srvy_st_version.Vertical_section, well_dir_srvy_st_version.Vertical_section_ouom, well_dir_srvy_st_version.X_offset, well_dir_srvy_st_version.X_offset_ew_direction, well_dir_srvy_st_version.X_offset_ouom, well_dir_srvy_st_version.Y_offset, well_dir_srvy_st_version.Y_offset_ns_direction, well_dir_srvy_st_version.Y_offset_ouom, well_dir_srvy_st_version.Row_changed_by, well_dir_srvy_st_version.Row_changed_date, well_dir_srvy_st_version.Row_created_by, well_dir_srvy_st_version.Row_effective_date, well_dir_srvy_st_version.Row_expiry_date, well_dir_srvy_st_version.Row_quality, well_dir_srvy_st_version.Uwi)
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

func PatchWellDirSrvyStVersion(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_dir_srvy_st_version set "
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

func DeleteWellDirSrvyStVersion(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_dir_srvy_st_version dto.Well_dir_srvy_st_version
	well_dir_srvy_st_version.Uwi = id

	stmt, err := db.Prepare("delete from well_dir_srvy_st_version where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_dir_srvy_st_version.Uwi)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellDirSrvyVersion(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_dir_srvy_version")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_dir_srvy_version

	for rows.Next() {
		var well_dir_srvy_version dto.Well_dir_srvy_version
		if err := rows.Scan(&well_dir_srvy_version.Uwi, &well_dir_srvy_version.Survey_id, &well_dir_srvy_version.Source, &well_dir_srvy_version.Version_obs_no, &well_dir_srvy_version.Active_ind, &well_dir_srvy_version.Azimuth_coord_sys_id, &well_dir_srvy_version.Azimuth_coord_sys_qualifier, &well_dir_srvy_version.Azimuth_coord_sys_remark, &well_dir_srvy_version.Azimuth_corr_angle, &well_dir_srvy_version.Azimuth_corr_angle_ouom, &well_dir_srvy_version.Azimuth_corr_angle_type, &well_dir_srvy_version.Azimuth_north_type, &well_dir_srvy_version.Base_depth, &well_dir_srvy_version.Base_depth_ouom, &well_dir_srvy_version.Compute_method, &well_dir_srvy_version.Confidential_type, &well_dir_srvy_version.Coord_acquisition_id, &well_dir_srvy_version.Dir_survey_class, &well_dir_srvy_version.Effective_date, &well_dir_srvy_version.Expiry_date, &well_dir_srvy_version.Extrapolate_depth, &well_dir_srvy_version.Extrapolate_depth_ouom, &well_dir_srvy_version.Extrapolate_ind, &well_dir_srvy_version.Geog_coord_sys_id, &well_dir_srvy_version.Local_coord_sys_id, &well_dir_srvy_version.Magnetic_declination, &well_dir_srvy_version.Magnetic_declination_ouom, &well_dir_srvy_version.Mag_decl_ew_direction, &well_dir_srvy_version.Offset_coord_sys_id, &well_dir_srvy_version.Offset_coord_sys_qualifier, &well_dir_srvy_version.Offset_coord_sys_remark, &well_dir_srvy_version.Offset_north_type, &well_dir_srvy_version.Plane_of_proposal, &well_dir_srvy_version.Plane_of_proposal_ouom, &well_dir_srvy_version.Ppdm_guid, &well_dir_srvy_version.Preferred_ind, &well_dir_srvy_version.Rad_uncert_dist, &well_dir_srvy_version.Rad_uncert_dist_ouom, &well_dir_srvy_version.Rad_uncert_dist_reason, &well_dir_srvy_version.Record_mode, &well_dir_srvy_version.Remark, &well_dir_srvy_version.Report_apd, &well_dir_srvy_version.Report_log_datum, &well_dir_srvy_version.Report_log_datum_elev, &well_dir_srvy_version.Report_log_datum_elev_ouom, &well_dir_srvy_version.Report_perm_datum, &well_dir_srvy_version.Report_perm_datum_elev, &well_dir_srvy_version.Report_perm_datum_elev_ouom, &well_dir_srvy_version.Resample_interval, &well_dir_srvy_version.Resample_interval_ouom, &well_dir_srvy_version.Rpt_coord_sys_id, &well_dir_srvy_version.Rpt_coord_sys_qualifier, &well_dir_srvy_version.Rpt_coord_sys_remark, &well_dir_srvy_version.Rpt_corr_angle, &well_dir_srvy_version.Rpt_corr_angle_ouom, &well_dir_srvy_version.Rpt_corr_angle_type, &well_dir_srvy_version.Rpt_north_type, &well_dir_srvy_version.Rpt_survey_type, &well_dir_srvy_version.Rpt_tie_azimuth, &well_dir_srvy_version.Rpt_tie_azimuth_ouom, &well_dir_srvy_version.Rpt_tie_inclination, &well_dir_srvy_version.Rpt_tie_inclination_ouom, &well_dir_srvy_version.Rpt_tie_md, &well_dir_srvy_version.Rpt_tie_md_ouom, &well_dir_srvy_version.Rpt_tie_point_ind, &well_dir_srvy_version.Rpt_tie_tvd, &well_dir_srvy_version.Rpt_tie_tvd_ouom, &well_dir_srvy_version.Rpt_tie_x_offset, &well_dir_srvy_version.Rpt_tie_x_offset_ew_dir, &well_dir_srvy_version.Rpt_tie_x_offset_ouom, &well_dir_srvy_version.Rpt_tie_y_offset, &well_dir_srvy_version.Rpt_tie_y_offset_ns_dir, &well_dir_srvy_version.Rpt_tie_y_offset_ouom, &well_dir_srvy_version.Source_document_id, &well_dir_srvy_version.Survey_company_ba_id, &well_dir_srvy_version.Survey_compiler_ba_id, &well_dir_srvy_version.Survey_digital_prvdr_ba_id, &well_dir_srvy_version.Survey_end_date, &well_dir_srvy_version.Survey_end_date_qualifier, &well_dir_srvy_version.Survey_process_method, &well_dir_srvy_version.Survey_quality, &well_dir_srvy_version.Survey_ref_num, &well_dir_srvy_version.Survey_report_type, &well_dir_srvy_version.Survey_run_num, &well_dir_srvy_version.Survey_start_date, &well_dir_srvy_version.Survey_start_date_qualifier, &well_dir_srvy_version.Survey_type, &well_dir_srvy_version.Top_depth, &well_dir_srvy_version.Top_depth_ouom, &well_dir_srvy_version.Row_changed_by, &well_dir_srvy_version.Row_changed_date, &well_dir_srvy_version.Row_created_by, &well_dir_srvy_version.Row_created_date, &well_dir_srvy_version.Row_effective_date, &well_dir_srvy_version.Row_expiry_date, &well_dir_srvy_version.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_dir_srvy_version to result
		result = append(result, well_dir_srvy_version)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellDirSrvyVersion(c *fiber.Ctx) error {
	var well_dir_srvy_version dto.Well_dir_srvy_version

	setDefaults(&well_dir_srvy_version)

	if err := json.Unmarshal(c.Body(), &well_dir_srvy_version); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_dir_srvy_version values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95, :96)")
	if err != nil {
		return err
	}
	well_dir_srvy_version.Row_created_date = formatDate(well_dir_srvy_version.Row_created_date)
	well_dir_srvy_version.Row_changed_date = formatDate(well_dir_srvy_version.Row_changed_date)
	well_dir_srvy_version.Effective_date = formatDateString(well_dir_srvy_version.Effective_date)
	well_dir_srvy_version.Expiry_date = formatDateString(well_dir_srvy_version.Expiry_date)
	well_dir_srvy_version.Survey_end_date = formatDateString(well_dir_srvy_version.Survey_end_date)
	well_dir_srvy_version.Survey_start_date = formatDateString(well_dir_srvy_version.Survey_start_date)
	well_dir_srvy_version.Row_effective_date = formatDateString(well_dir_srvy_version.Row_effective_date)
	well_dir_srvy_version.Row_expiry_date = formatDateString(well_dir_srvy_version.Row_expiry_date)






	rows, err := stmt.Exec(well_dir_srvy_version.Uwi, well_dir_srvy_version.Survey_id, well_dir_srvy_version.Source, well_dir_srvy_version.Version_obs_no, well_dir_srvy_version.Active_ind, well_dir_srvy_version.Azimuth_coord_sys_id, well_dir_srvy_version.Azimuth_coord_sys_qualifier, well_dir_srvy_version.Azimuth_coord_sys_remark, well_dir_srvy_version.Azimuth_corr_angle, well_dir_srvy_version.Azimuth_corr_angle_ouom, well_dir_srvy_version.Azimuth_corr_angle_type, well_dir_srvy_version.Azimuth_north_type, well_dir_srvy_version.Base_depth, well_dir_srvy_version.Base_depth_ouom, well_dir_srvy_version.Compute_method, well_dir_srvy_version.Confidential_type, well_dir_srvy_version.Coord_acquisition_id, well_dir_srvy_version.Dir_survey_class, well_dir_srvy_version.Effective_date, well_dir_srvy_version.Expiry_date, well_dir_srvy_version.Extrapolate_depth, well_dir_srvy_version.Extrapolate_depth_ouom, well_dir_srvy_version.Extrapolate_ind, well_dir_srvy_version.Geog_coord_sys_id, well_dir_srvy_version.Local_coord_sys_id, well_dir_srvy_version.Magnetic_declination, well_dir_srvy_version.Magnetic_declination_ouom, well_dir_srvy_version.Mag_decl_ew_direction, well_dir_srvy_version.Offset_coord_sys_id, well_dir_srvy_version.Offset_coord_sys_qualifier, well_dir_srvy_version.Offset_coord_sys_remark, well_dir_srvy_version.Offset_north_type, well_dir_srvy_version.Plane_of_proposal, well_dir_srvy_version.Plane_of_proposal_ouom, well_dir_srvy_version.Ppdm_guid, well_dir_srvy_version.Preferred_ind, well_dir_srvy_version.Rad_uncert_dist, well_dir_srvy_version.Rad_uncert_dist_ouom, well_dir_srvy_version.Rad_uncert_dist_reason, well_dir_srvy_version.Record_mode, well_dir_srvy_version.Remark, well_dir_srvy_version.Report_apd, well_dir_srvy_version.Report_log_datum, well_dir_srvy_version.Report_log_datum_elev, well_dir_srvy_version.Report_log_datum_elev_ouom, well_dir_srvy_version.Report_perm_datum, well_dir_srvy_version.Report_perm_datum_elev, well_dir_srvy_version.Report_perm_datum_elev_ouom, well_dir_srvy_version.Resample_interval, well_dir_srvy_version.Resample_interval_ouom, well_dir_srvy_version.Rpt_coord_sys_id, well_dir_srvy_version.Rpt_coord_sys_qualifier, well_dir_srvy_version.Rpt_coord_sys_remark, well_dir_srvy_version.Rpt_corr_angle, well_dir_srvy_version.Rpt_corr_angle_ouom, well_dir_srvy_version.Rpt_corr_angle_type, well_dir_srvy_version.Rpt_north_type, well_dir_srvy_version.Rpt_survey_type, well_dir_srvy_version.Rpt_tie_azimuth, well_dir_srvy_version.Rpt_tie_azimuth_ouom, well_dir_srvy_version.Rpt_tie_inclination, well_dir_srvy_version.Rpt_tie_inclination_ouom, well_dir_srvy_version.Rpt_tie_md, well_dir_srvy_version.Rpt_tie_md_ouom, well_dir_srvy_version.Rpt_tie_point_ind, well_dir_srvy_version.Rpt_tie_tvd, well_dir_srvy_version.Rpt_tie_tvd_ouom, well_dir_srvy_version.Rpt_tie_x_offset, well_dir_srvy_version.Rpt_tie_x_offset_ew_dir, well_dir_srvy_version.Rpt_tie_x_offset_ouom, well_dir_srvy_version.Rpt_tie_y_offset, well_dir_srvy_version.Rpt_tie_y_offset_ns_dir, well_dir_srvy_version.Rpt_tie_y_offset_ouom, well_dir_srvy_version.Source_document_id, well_dir_srvy_version.Survey_company_ba_id, well_dir_srvy_version.Survey_compiler_ba_id, well_dir_srvy_version.Survey_digital_prvdr_ba_id, well_dir_srvy_version.Survey_end_date, well_dir_srvy_version.Survey_end_date_qualifier, well_dir_srvy_version.Survey_process_method, well_dir_srvy_version.Survey_quality, well_dir_srvy_version.Survey_ref_num, well_dir_srvy_version.Survey_report_type, well_dir_srvy_version.Survey_run_num, well_dir_srvy_version.Survey_start_date, well_dir_srvy_version.Survey_start_date_qualifier, well_dir_srvy_version.Survey_type, well_dir_srvy_version.Top_depth, well_dir_srvy_version.Top_depth_ouom, well_dir_srvy_version.Row_changed_by, well_dir_srvy_version.Row_changed_date, well_dir_srvy_version.Row_created_by, well_dir_srvy_version.Row_created_date, well_dir_srvy_version.Row_effective_date, well_dir_srvy_version.Row_expiry_date, well_dir_srvy_version.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellDirSrvyVersion(c *fiber.Ctx) error {
	var well_dir_srvy_version dto.Well_dir_srvy_version

	setDefaults(&well_dir_srvy_version)

	if err := json.Unmarshal(c.Body(), &well_dir_srvy_version); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_dir_srvy_version.Uwi = id

    if well_dir_srvy_version.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_dir_srvy_version set survey_id = :1, source = :2, version_obs_no = :3, active_ind = :4, azimuth_coord_sys_id = :5, azimuth_coord_sys_qualifier = :6, azimuth_coord_sys_remark = :7, azimuth_corr_angle = :8, azimuth_corr_angle_ouom = :9, azimuth_corr_angle_type = :10, azimuth_north_type = :11, base_depth = :12, base_depth_ouom = :13, compute_method = :14, confidential_type = :15, coord_acquisition_id = :16, dir_survey_class = :17, effective_date = :18, expiry_date = :19, extrapolate_depth = :20, extrapolate_depth_ouom = :21, extrapolate_ind = :22, geog_coord_sys_id = :23, local_coord_sys_id = :24, magnetic_declination = :25, magnetic_declination_ouom = :26, mag_decl_ew_direction = :27, offset_coord_sys_id = :28, offset_coord_sys_qualifier = :29, offset_coord_sys_remark = :30, offset_north_type = :31, plane_of_proposal = :32, plane_of_proposal_ouom = :33, ppdm_guid = :34, preferred_ind = :35, rad_uncert_dist = :36, rad_uncert_dist_ouom = :37, rad_uncert_dist_reason = :38, record_mode = :39, remark = :40, report_apd = :41, report_log_datum = :42, report_log_datum_elev = :43, report_log_datum_elev_ouom = :44, report_perm_datum = :45, report_perm_datum_elev = :46, report_perm_datum_elev_ouom = :47, resample_interval = :48, resample_interval_ouom = :49, rpt_coord_sys_id = :50, rpt_coord_sys_qualifier = :51, rpt_coord_sys_remark = :52, rpt_corr_angle = :53, rpt_corr_angle_ouom = :54, rpt_corr_angle_type = :55, rpt_north_type = :56, rpt_survey_type = :57, rpt_tie_azimuth = :58, rpt_tie_azimuth_ouom = :59, rpt_tie_inclination = :60, rpt_tie_inclination_ouom = :61, rpt_tie_md = :62, rpt_tie_md_ouom = :63, rpt_tie_point_ind = :64, rpt_tie_tvd = :65, rpt_tie_tvd_ouom = :66, rpt_tie_x_offset = :67, rpt_tie_x_offset_ew_dir = :68, rpt_tie_x_offset_ouom = :69, rpt_tie_y_offset = :70, rpt_tie_y_offset_ns_dir = :71, rpt_tie_y_offset_ouom = :72, source_document_id = :73, survey_company_ba_id = :74, survey_compiler_ba_id = :75, survey_digital_prvdr_ba_id = :76, survey_end_date = :77, survey_end_date_qualifier = :78, survey_process_method = :79, survey_quality = :80, survey_ref_num = :81, survey_report_type = :82, survey_run_num = :83, survey_start_date = :84, survey_start_date_qualifier = :85, survey_type = :86, top_depth = :87, top_depth_ouom = :88, row_changed_by = :89, row_changed_date = :90, row_created_by = :91, row_effective_date = :92, row_expiry_date = :93, row_quality = :94 where uwi = :96")
	if err != nil {
		return err
	}

	well_dir_srvy_version.Row_changed_date = formatDate(well_dir_srvy_version.Row_changed_date)
	well_dir_srvy_version.Effective_date = formatDateString(well_dir_srvy_version.Effective_date)
	well_dir_srvy_version.Expiry_date = formatDateString(well_dir_srvy_version.Expiry_date)
	well_dir_srvy_version.Survey_end_date = formatDateString(well_dir_srvy_version.Survey_end_date)
	well_dir_srvy_version.Survey_start_date = formatDateString(well_dir_srvy_version.Survey_start_date)
	well_dir_srvy_version.Row_effective_date = formatDateString(well_dir_srvy_version.Row_effective_date)
	well_dir_srvy_version.Row_expiry_date = formatDateString(well_dir_srvy_version.Row_expiry_date)






	rows, err := stmt.Exec(well_dir_srvy_version.Survey_id, well_dir_srvy_version.Source, well_dir_srvy_version.Version_obs_no, well_dir_srvy_version.Active_ind, well_dir_srvy_version.Azimuth_coord_sys_id, well_dir_srvy_version.Azimuth_coord_sys_qualifier, well_dir_srvy_version.Azimuth_coord_sys_remark, well_dir_srvy_version.Azimuth_corr_angle, well_dir_srvy_version.Azimuth_corr_angle_ouom, well_dir_srvy_version.Azimuth_corr_angle_type, well_dir_srvy_version.Azimuth_north_type, well_dir_srvy_version.Base_depth, well_dir_srvy_version.Base_depth_ouom, well_dir_srvy_version.Compute_method, well_dir_srvy_version.Confidential_type, well_dir_srvy_version.Coord_acquisition_id, well_dir_srvy_version.Dir_survey_class, well_dir_srvy_version.Effective_date, well_dir_srvy_version.Expiry_date, well_dir_srvy_version.Extrapolate_depth, well_dir_srvy_version.Extrapolate_depth_ouom, well_dir_srvy_version.Extrapolate_ind, well_dir_srvy_version.Geog_coord_sys_id, well_dir_srvy_version.Local_coord_sys_id, well_dir_srvy_version.Magnetic_declination, well_dir_srvy_version.Magnetic_declination_ouom, well_dir_srvy_version.Mag_decl_ew_direction, well_dir_srvy_version.Offset_coord_sys_id, well_dir_srvy_version.Offset_coord_sys_qualifier, well_dir_srvy_version.Offset_coord_sys_remark, well_dir_srvy_version.Offset_north_type, well_dir_srvy_version.Plane_of_proposal, well_dir_srvy_version.Plane_of_proposal_ouom, well_dir_srvy_version.Ppdm_guid, well_dir_srvy_version.Preferred_ind, well_dir_srvy_version.Rad_uncert_dist, well_dir_srvy_version.Rad_uncert_dist_ouom, well_dir_srvy_version.Rad_uncert_dist_reason, well_dir_srvy_version.Record_mode, well_dir_srvy_version.Remark, well_dir_srvy_version.Report_apd, well_dir_srvy_version.Report_log_datum, well_dir_srvy_version.Report_log_datum_elev, well_dir_srvy_version.Report_log_datum_elev_ouom, well_dir_srvy_version.Report_perm_datum, well_dir_srvy_version.Report_perm_datum_elev, well_dir_srvy_version.Report_perm_datum_elev_ouom, well_dir_srvy_version.Resample_interval, well_dir_srvy_version.Resample_interval_ouom, well_dir_srvy_version.Rpt_coord_sys_id, well_dir_srvy_version.Rpt_coord_sys_qualifier, well_dir_srvy_version.Rpt_coord_sys_remark, well_dir_srvy_version.Rpt_corr_angle, well_dir_srvy_version.Rpt_corr_angle_ouom, well_dir_srvy_version.Rpt_corr_angle_type, well_dir_srvy_version.Rpt_north_type, well_dir_srvy_version.Rpt_survey_type, well_dir_srvy_version.Rpt_tie_azimuth, well_dir_srvy_version.Rpt_tie_azimuth_ouom, well_dir_srvy_version.Rpt_tie_inclination, well_dir_srvy_version.Rpt_tie_inclination_ouom, well_dir_srvy_version.Rpt_tie_md, well_dir_srvy_version.Rpt_tie_md_ouom, well_dir_srvy_version.Rpt_tie_point_ind, well_dir_srvy_version.Rpt_tie_tvd, well_dir_srvy_version.Rpt_tie_tvd_ouom, well_dir_srvy_version.Rpt_tie_x_offset, well_dir_srvy_version.Rpt_tie_x_offset_ew_dir, well_dir_srvy_version.Rpt_tie_x_offset_ouom, well_dir_srvy_version.Rpt_tie_y_offset, well_dir_srvy_version.Rpt_tie_y_offset_ns_dir, well_dir_srvy_version.Rpt_tie_y_offset_ouom, well_dir_srvy_version.Source_document_id, well_dir_srvy_version.Survey_company_ba_id, well_dir_srvy_version.Survey_compiler_ba_id, well_dir_srvy_version.Survey_digital_prvdr_ba_id, well_dir_srvy_version.Survey_end_date, well_dir_srvy_version.Survey_end_date_qualifier, well_dir_srvy_version.Survey_process_method, well_dir_srvy_version.Survey_quality, well_dir_srvy_version.Survey_ref_num, well_dir_srvy_version.Survey_report_type, well_dir_srvy_version.Survey_run_num, well_dir_srvy_version.Survey_start_date, well_dir_srvy_version.Survey_start_date_qualifier, well_dir_srvy_version.Survey_type, well_dir_srvy_version.Top_depth, well_dir_srvy_version.Top_depth_ouom, well_dir_srvy_version.Row_changed_by, well_dir_srvy_version.Row_changed_date, well_dir_srvy_version.Row_created_by, well_dir_srvy_version.Row_effective_date, well_dir_srvy_version.Row_expiry_date, well_dir_srvy_version.Row_quality, well_dir_srvy_version.Uwi)
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

func PatchWellDirSrvyVersion(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_dir_srvy_version set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "survey_end_date" || key == "survey_start_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellDirSrvyVersion(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_dir_srvy_version dto.Well_dir_srvy_version
	well_dir_srvy_version.Uwi = id

	stmt, err := db.Prepare("delete from well_dir_srvy_version where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_dir_srvy_version.Uwi)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellDirSrvy(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_dir_srvy")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_dir_srvy

	for rows.Next() {
		var well_dir_srvy dto.Well_dir_srvy
		if err := rows.Scan(&well_dir_srvy.Uwi, &well_dir_srvy.Survey_id, &well_dir_srvy.Source, &well_dir_srvy.Active_ind, &well_dir_srvy.Azimuth_coord_sys_id, &well_dir_srvy.Azimuth_coord_sys_qualifier, &well_dir_srvy.Azimuth_coord_sys_remark, &well_dir_srvy.Azimuth_corr_angle, &well_dir_srvy.Azimuth_corr_angle_ouom, &well_dir_srvy.Azimuth_corr_angle_type, &well_dir_srvy.Azimuth_north_type, &well_dir_srvy.Base_depth, &well_dir_srvy.Base_depth_ouom, &well_dir_srvy.Compute_method, &well_dir_srvy.Confidential_type, &well_dir_srvy.Coord_acquisition_id, &well_dir_srvy.Dir_survey_class, &well_dir_srvy.Effective_date, &well_dir_srvy.Expiry_date, &well_dir_srvy.Extrapolate_depth, &well_dir_srvy.Extrapolate_depth_ouom, &well_dir_srvy.Extrapolate_ind, &well_dir_srvy.Geog_coord_sys_id, &well_dir_srvy.Local_coord_sys_id, &well_dir_srvy.Magnetic_declination, &well_dir_srvy.Magnetic_declination_ouom, &well_dir_srvy.Mag_decl_ew_direction, &well_dir_srvy.Offset_coord_sys_id, &well_dir_srvy.Offset_coord_sys_qualifier, &well_dir_srvy.Offset_coord_sys_remark, &well_dir_srvy.Offset_north_type, &well_dir_srvy.Plane_of_proposal, &well_dir_srvy.Plane_of_proposal_ouom, &well_dir_srvy.Ppdm_guid, &well_dir_srvy.Preferred_ind, &well_dir_srvy.Rad_uncert_dist, &well_dir_srvy.Rad_uncert_dist_ouom, &well_dir_srvy.Rad_uncert_dist_reason, &well_dir_srvy.Record_mode, &well_dir_srvy.Remark, &well_dir_srvy.Report_apd, &well_dir_srvy.Report_log_datum, &well_dir_srvy.Report_log_datum_elev, &well_dir_srvy.Report_log_datum_elev_ouom, &well_dir_srvy.Report_perm_datum, &well_dir_srvy.Report_perm_datum_elev, &well_dir_srvy.Report_perm_datum_elev_ouom, &well_dir_srvy.Resample_interval, &well_dir_srvy.Resample_interval_ouom, &well_dir_srvy.Rpt_coord_sys_id, &well_dir_srvy.Rpt_coord_sys_qualifier, &well_dir_srvy.Rpt_coord_sys_remark, &well_dir_srvy.Rpt_corr_angle, &well_dir_srvy.Rpt_corr_angle_ouom, &well_dir_srvy.Rpt_corr_angle_type, &well_dir_srvy.Rpt_north_type, &well_dir_srvy.Rpt_survey_type, &well_dir_srvy.Rpt_tie_azimuth, &well_dir_srvy.Rpt_tie_azimuth_ouom, &well_dir_srvy.Rpt_tie_inclination, &well_dir_srvy.Rpt_tie_inclination_ouom, &well_dir_srvy.Rpt_tie_md, &well_dir_srvy.Rpt_tie_md_ouom, &well_dir_srvy.Rpt_tie_point_ind, &well_dir_srvy.Rpt_tie_tvd, &well_dir_srvy.Rpt_tie_tvd_ouom, &well_dir_srvy.Rpt_tie_x_offset, &well_dir_srvy.Rpt_tie_x_offset_ew_dir, &well_dir_srvy.Rpt_tie_x_offset_ouom, &well_dir_srvy.Rpt_tie_y_offset, &well_dir_srvy.Rpt_tie_y_offset_ns_dir, &well_dir_srvy.Rpt_tie_y_offset_ouom, &well_dir_srvy.Source_document_id, &well_dir_srvy.Survey_company_ba_id, &well_dir_srvy.Survey_compiler_ba_id, &well_dir_srvy.Survey_digital_prvdr_ba_id, &well_dir_srvy.Survey_end_date, &well_dir_srvy.Survey_end_date_qualifier, &well_dir_srvy.Survey_process_method, &well_dir_srvy.Survey_quality, &well_dir_srvy.Survey_ref_num, &well_dir_srvy.Survey_report_type, &well_dir_srvy.Survey_run_num, &well_dir_srvy.Survey_start_date, &well_dir_srvy.Survey_start_date_qualifier, &well_dir_srvy.Survey_type, &well_dir_srvy.Top_depth, &well_dir_srvy.Top_depth_ouom, &well_dir_srvy.Row_changed_by, &well_dir_srvy.Row_changed_date, &well_dir_srvy.Row_created_by, &well_dir_srvy.Row_created_date, &well_dir_srvy.Row_effective_date, &well_dir_srvy.Row_expiry_date, &well_dir_srvy.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_dir_srvy to result
		result = append(result, well_dir_srvy)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellDirSrvy(c *fiber.Ctx) error {
	var well_dir_srvy dto.Well_dir_srvy

	setDefaults(&well_dir_srvy)

	if err := json.Unmarshal(c.Body(), &well_dir_srvy); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_dir_srvy values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89, :90, :91, :92, :93, :94, :95)")
	if err != nil {
		return err
	}
	well_dir_srvy.Row_created_date = formatDate(well_dir_srvy.Row_created_date)
	well_dir_srvy.Row_changed_date = formatDate(well_dir_srvy.Row_changed_date)
	well_dir_srvy.Effective_date = formatDateString(well_dir_srvy.Effective_date)
	well_dir_srvy.Expiry_date = formatDateString(well_dir_srvy.Expiry_date)
	well_dir_srvy.Survey_end_date = formatDateString(well_dir_srvy.Survey_end_date)
	well_dir_srvy.Survey_start_date = formatDateString(well_dir_srvy.Survey_start_date)
	well_dir_srvy.Row_effective_date = formatDateString(well_dir_srvy.Row_effective_date)
	well_dir_srvy.Row_expiry_date = formatDateString(well_dir_srvy.Row_expiry_date)






	rows, err := stmt.Exec(well_dir_srvy.Uwi, well_dir_srvy.Survey_id, well_dir_srvy.Source, well_dir_srvy.Active_ind, well_dir_srvy.Azimuth_coord_sys_id, well_dir_srvy.Azimuth_coord_sys_qualifier, well_dir_srvy.Azimuth_coord_sys_remark, well_dir_srvy.Azimuth_corr_angle, well_dir_srvy.Azimuth_corr_angle_ouom, well_dir_srvy.Azimuth_corr_angle_type, well_dir_srvy.Azimuth_north_type, well_dir_srvy.Base_depth, well_dir_srvy.Base_depth_ouom, well_dir_srvy.Compute_method, well_dir_srvy.Confidential_type, well_dir_srvy.Coord_acquisition_id, well_dir_srvy.Dir_survey_class, well_dir_srvy.Effective_date, well_dir_srvy.Expiry_date, well_dir_srvy.Extrapolate_depth, well_dir_srvy.Extrapolate_depth_ouom, well_dir_srvy.Extrapolate_ind, well_dir_srvy.Geog_coord_sys_id, well_dir_srvy.Local_coord_sys_id, well_dir_srvy.Magnetic_declination, well_dir_srvy.Magnetic_declination_ouom, well_dir_srvy.Mag_decl_ew_direction, well_dir_srvy.Offset_coord_sys_id, well_dir_srvy.Offset_coord_sys_qualifier, well_dir_srvy.Offset_coord_sys_remark, well_dir_srvy.Offset_north_type, well_dir_srvy.Plane_of_proposal, well_dir_srvy.Plane_of_proposal_ouom, well_dir_srvy.Ppdm_guid, well_dir_srvy.Preferred_ind, well_dir_srvy.Rad_uncert_dist, well_dir_srvy.Rad_uncert_dist_ouom, well_dir_srvy.Rad_uncert_dist_reason, well_dir_srvy.Record_mode, well_dir_srvy.Remark, well_dir_srvy.Report_apd, well_dir_srvy.Report_log_datum, well_dir_srvy.Report_log_datum_elev, well_dir_srvy.Report_log_datum_elev_ouom, well_dir_srvy.Report_perm_datum, well_dir_srvy.Report_perm_datum_elev, well_dir_srvy.Report_perm_datum_elev_ouom, well_dir_srvy.Resample_interval, well_dir_srvy.Resample_interval_ouom, well_dir_srvy.Rpt_coord_sys_id, well_dir_srvy.Rpt_coord_sys_qualifier, well_dir_srvy.Rpt_coord_sys_remark, well_dir_srvy.Rpt_corr_angle, well_dir_srvy.Rpt_corr_angle_ouom, well_dir_srvy.Rpt_corr_angle_type, well_dir_srvy.Rpt_north_type, well_dir_srvy.Rpt_survey_type, well_dir_srvy.Rpt_tie_azimuth, well_dir_srvy.Rpt_tie_azimuth_ouom, well_dir_srvy.Rpt_tie_inclination, well_dir_srvy.Rpt_tie_inclination_ouom, well_dir_srvy.Rpt_tie_md, well_dir_srvy.Rpt_tie_md_ouom, well_dir_srvy.Rpt_tie_point_ind, well_dir_srvy.Rpt_tie_tvd, well_dir_srvy.Rpt_tie_tvd_ouom, well_dir_srvy.Rpt_tie_x_offset, well_dir_srvy.Rpt_tie_x_offset_ew_dir, well_dir_srvy.Rpt_tie_x_offset_ouom, well_dir_srvy.Rpt_tie_y_offset, well_dir_srvy.Rpt_tie_y_offset_ns_dir, well_dir_srvy.Rpt_tie_y_offset_ouom, well_dir_srvy.Source_document_id, well_dir_srvy.Survey_company_ba_id, well_dir_srvy.Survey_compiler_ba_id, well_dir_srvy.Survey_digital_prvdr_ba_id, well_dir_srvy.Survey_end_date, well_dir_srvy.Survey_end_date_qualifier, well_dir_srvy.Survey_process_method, well_dir_srvy.Survey_quality, well_dir_srvy.Survey_ref_num, well_dir_srvy.Survey_report_type, well_dir_srvy.Survey_run_num, well_dir_srvy.Survey_start_date, well_dir_srvy.Survey_start_date_qualifier, well_dir_srvy.Survey_type, well_dir_srvy.Top_depth, well_dir_srvy.Top_depth_ouom, well_dir_srvy.Row_changed_by, well_dir_srvy.Row_changed_date, well_dir_srvy.Row_created_by, well_dir_srvy.Row_created_date, well_dir_srvy.Row_effective_date, well_dir_srvy.Row_expiry_date, well_dir_srvy.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellDirSrvy(c *fiber.Ctx) error {
	var well_dir_srvy dto.Well_dir_srvy

	setDefaults(&well_dir_srvy)

	if err := json.Unmarshal(c.Body(), &well_dir_srvy); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_dir_srvy.Uwi = id

    if well_dir_srvy.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_dir_srvy set survey_id = :1, source = :2, active_ind = :3, azimuth_coord_sys_id = :4, azimuth_coord_sys_qualifier = :5, azimuth_coord_sys_remark = :6, azimuth_corr_angle = :7, azimuth_corr_angle_ouom = :8, azimuth_corr_angle_type = :9, azimuth_north_type = :10, base_depth = :11, base_depth_ouom = :12, compute_method = :13, confidential_type = :14, coord_acquisition_id = :15, dir_survey_class = :16, effective_date = :17, expiry_date = :18, extrapolate_depth = :19, extrapolate_depth_ouom = :20, extrapolate_ind = :21, geog_coord_sys_id = :22, local_coord_sys_id = :23, magnetic_declination = :24, magnetic_declination_ouom = :25, mag_decl_ew_direction = :26, offset_coord_sys_id = :27, offset_coord_sys_qualifier = :28, offset_coord_sys_remark = :29, offset_north_type = :30, plane_of_proposal = :31, plane_of_proposal_ouom = :32, ppdm_guid = :33, preferred_ind = :34, rad_uncert_dist = :35, rad_uncert_dist_ouom = :36, rad_uncert_dist_reason = :37, record_mode = :38, remark = :39, report_apd = :40, report_log_datum = :41, report_log_datum_elev = :42, report_log_datum_elev_ouom = :43, report_perm_datum = :44, report_perm_datum_elev = :45, report_perm_datum_elev_ouom = :46, resample_interval = :47, resample_interval_ouom = :48, rpt_coord_sys_id = :49, rpt_coord_sys_qualifier = :50, rpt_coord_sys_remark = :51, rpt_corr_angle = :52, rpt_corr_angle_ouom = :53, rpt_corr_angle_type = :54, rpt_north_type = :55, rpt_survey_type = :56, rpt_tie_azimuth = :57, rpt_tie_azimuth_ouom = :58, rpt_tie_inclination = :59, rpt_tie_inclination_ouom = :60, rpt_tie_md = :61, rpt_tie_md_ouom = :62, rpt_tie_point_ind = :63, rpt_tie_tvd = :64, rpt_tie_tvd_ouom = :65, rpt_tie_x_offset = :66, rpt_tie_x_offset_ew_dir = :67, rpt_tie_x_offset_ouom = :68, rpt_tie_y_offset = :69, rpt_tie_y_offset_ns_dir = :70, rpt_tie_y_offset_ouom = :71, source_document_id = :72, survey_company_ba_id = :73, survey_compiler_ba_id = :74, survey_digital_prvdr_ba_id = :75, survey_end_date = :76, survey_end_date_qualifier = :77, survey_process_method = :78, survey_quality = :79, survey_ref_num = :80, survey_report_type = :81, survey_run_num = :82, survey_start_date = :83, survey_start_date_qualifier = :84, survey_type = :85, top_depth = :86, top_depth_ouom = :87, row_changed_by = :88, row_changed_date = :89, row_created_by = :90, row_effective_date = :91, row_expiry_date = :92, row_quality = :93 where uwi = :95")
	if err != nil {
		return err
	}

	well_dir_srvy.Row_changed_date = formatDate(well_dir_srvy.Row_changed_date)
	well_dir_srvy.Effective_date = formatDateString(well_dir_srvy.Effective_date)
	well_dir_srvy.Expiry_date = formatDateString(well_dir_srvy.Expiry_date)
	well_dir_srvy.Survey_end_date = formatDateString(well_dir_srvy.Survey_end_date)
	well_dir_srvy.Survey_start_date = formatDateString(well_dir_srvy.Survey_start_date)
	well_dir_srvy.Row_effective_date = formatDateString(well_dir_srvy.Row_effective_date)
	well_dir_srvy.Row_expiry_date = formatDateString(well_dir_srvy.Row_expiry_date)






	rows, err := stmt.Exec(well_dir_srvy.Survey_id, well_dir_srvy.Source, well_dir_srvy.Active_ind, well_dir_srvy.Azimuth_coord_sys_id, well_dir_srvy.Azimuth_coord_sys_qualifier, well_dir_srvy.Azimuth_coord_sys_remark, well_dir_srvy.Azimuth_corr_angle, well_dir_srvy.Azimuth_corr_angle_ouom, well_dir_srvy.Azimuth_corr_angle_type, well_dir_srvy.Azimuth_north_type, well_dir_srvy.Base_depth, well_dir_srvy.Base_depth_ouom, well_dir_srvy.Compute_method, well_dir_srvy.Confidential_type, well_dir_srvy.Coord_acquisition_id, well_dir_srvy.Dir_survey_class, well_dir_srvy.Effective_date, well_dir_srvy.Expiry_date, well_dir_srvy.Extrapolate_depth, well_dir_srvy.Extrapolate_depth_ouom, well_dir_srvy.Extrapolate_ind, well_dir_srvy.Geog_coord_sys_id, well_dir_srvy.Local_coord_sys_id, well_dir_srvy.Magnetic_declination, well_dir_srvy.Magnetic_declination_ouom, well_dir_srvy.Mag_decl_ew_direction, well_dir_srvy.Offset_coord_sys_id, well_dir_srvy.Offset_coord_sys_qualifier, well_dir_srvy.Offset_coord_sys_remark, well_dir_srvy.Offset_north_type, well_dir_srvy.Plane_of_proposal, well_dir_srvy.Plane_of_proposal_ouom, well_dir_srvy.Ppdm_guid, well_dir_srvy.Preferred_ind, well_dir_srvy.Rad_uncert_dist, well_dir_srvy.Rad_uncert_dist_ouom, well_dir_srvy.Rad_uncert_dist_reason, well_dir_srvy.Record_mode, well_dir_srvy.Remark, well_dir_srvy.Report_apd, well_dir_srvy.Report_log_datum, well_dir_srvy.Report_log_datum_elev, well_dir_srvy.Report_log_datum_elev_ouom, well_dir_srvy.Report_perm_datum, well_dir_srvy.Report_perm_datum_elev, well_dir_srvy.Report_perm_datum_elev_ouom, well_dir_srvy.Resample_interval, well_dir_srvy.Resample_interval_ouom, well_dir_srvy.Rpt_coord_sys_id, well_dir_srvy.Rpt_coord_sys_qualifier, well_dir_srvy.Rpt_coord_sys_remark, well_dir_srvy.Rpt_corr_angle, well_dir_srvy.Rpt_corr_angle_ouom, well_dir_srvy.Rpt_corr_angle_type, well_dir_srvy.Rpt_north_type, well_dir_srvy.Rpt_survey_type, well_dir_srvy.Rpt_tie_azimuth, well_dir_srvy.Rpt_tie_azimuth_ouom, well_dir_srvy.Rpt_tie_inclination, well_dir_srvy.Rpt_tie_inclination_ouom, well_dir_srvy.Rpt_tie_md, well_dir_srvy.Rpt_tie_md_ouom, well_dir_srvy.Rpt_tie_point_ind, well_dir_srvy.Rpt_tie_tvd, well_dir_srvy.Rpt_tie_tvd_ouom, well_dir_srvy.Rpt_tie_x_offset, well_dir_srvy.Rpt_tie_x_offset_ew_dir, well_dir_srvy.Rpt_tie_x_offset_ouom, well_dir_srvy.Rpt_tie_y_offset, well_dir_srvy.Rpt_tie_y_offset_ns_dir, well_dir_srvy.Rpt_tie_y_offset_ouom, well_dir_srvy.Source_document_id, well_dir_srvy.Survey_company_ba_id, well_dir_srvy.Survey_compiler_ba_id, well_dir_srvy.Survey_digital_prvdr_ba_id, well_dir_srvy.Survey_end_date, well_dir_srvy.Survey_end_date_qualifier, well_dir_srvy.Survey_process_method, well_dir_srvy.Survey_quality, well_dir_srvy.Survey_ref_num, well_dir_srvy.Survey_report_type, well_dir_srvy.Survey_run_num, well_dir_srvy.Survey_start_date, well_dir_srvy.Survey_start_date_qualifier, well_dir_srvy.Survey_type, well_dir_srvy.Top_depth, well_dir_srvy.Top_depth_ouom, well_dir_srvy.Row_changed_by, well_dir_srvy.Row_changed_date, well_dir_srvy.Row_created_by, well_dir_srvy.Row_effective_date, well_dir_srvy.Row_expiry_date, well_dir_srvy.Row_quality, well_dir_srvy.Uwi)
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

func PatchWellDirSrvy(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_dir_srvy set "
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

func DeleteWellDirSrvy(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_dir_srvy dto.Well_dir_srvy
	well_dir_srvy.Uwi = id

	stmt, err := db.Prepare("delete from well_dir_srvy where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_dir_srvy.Uwi)
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



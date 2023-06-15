package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisAcqtnDesign(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_acqtn_design")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_acqtn_design

	for rows.Next() {
		var seis_acqtn_design dto.Seis_acqtn_design
		if err := rows.Scan(&seis_acqtn_design.Acqtn_design_id, &seis_acqtn_design.Acqtn_completed_date, &seis_acqtn_design.Acqtn_completed_date_desc, &seis_acqtn_design.Acqtn_dimension, &seis_acqtn_design.Acqtn_direction, &seis_acqtn_design.Acqtn_inline_bin_size, &seis_acqtn_design.Acqtn_inline_bin_size_ouom, &seis_acqtn_design.Acqtn_remark, &seis_acqtn_design.Acqtn_shotpt_interval, &seis_acqtn_design.Acqtn_shotpt_interval_ouom, &seis_acqtn_design.Acqtn_shot_line_spacing, &seis_acqtn_design.Acqtn_shot_line_spacing_ouom, &seis_acqtn_design.Acqtn_shot_time_intvl, &seis_acqtn_design.Acqtn_shot_time_intvl_ouom, &seis_acqtn_design.Acqtn_start_date, &seis_acqtn_design.Acqtn_start_date_desc, &seis_acqtn_design.Acqtn_xline_bin_size, &seis_acqtn_design.Acqtn_xline_bin_size_ouom, &seis_acqtn_design.Active_ind, &seis_acqtn_design.Actual_ind, &seis_acqtn_design.Cdp_coverage, &seis_acqtn_design.Effective_date, &seis_acqtn_design.Energy_charge_size, &seis_acqtn_design.Energy_charge_size_ouom, &seis_acqtn_design.Energy_oprg_psr, &seis_acqtn_design.Energy_oprg_psr_ouom, &seis_acqtn_design.Energy_oprg_volume, &seis_acqtn_design.Energy_oprg_volume_ouom, &seis_acqtn_design.Energy_shot_depth, &seis_acqtn_design.Energy_shot_depth_ouom, &seis_acqtn_design.Energy_src_array_spc, &seis_acqtn_design.Energy_src_array_spc_ouom, &seis_acqtn_design.Energy_src_array_type, &seis_acqtn_design.Energy_src_make, &seis_acqtn_design.Energy_src_per_shot, &seis_acqtn_design.Energy_sweep_duration, &seis_acqtn_design.Energy_sweep_duration_ouom, &seis_acqtn_design.Energy_sweep_end_freq, &seis_acqtn_design.Energy_sweep_freq_ouom, &seis_acqtn_design.Energy_sweep_mvup_dist, &seis_acqtn_design.Energy_sweep_mvup_dist_ouom, &seis_acqtn_design.Energy_sweep_no, &seis_acqtn_design.Energy_sweep_st_freq, &seis_acqtn_design.Energy_sweep_taper, &seis_acqtn_design.Energy_sweep_taper_ouom, &seis_acqtn_design.Energy_sweep_type, &seis_acqtn_design.Energy_type, &seis_acqtn_design.Environment, &seis_acqtn_design.Expiry_date, &seis_acqtn_design.Monitor_depth, &seis_acqtn_design.Monitor_depth_ouom, &seis_acqtn_design.Nominal_ind, &seis_acqtn_design.Ppdm_guid, &seis_acqtn_design.Rcrd_channel_count, &seis_acqtn_design.Rcrd_format_type, &seis_acqtn_design.Rcrd_gain_mode, &seis_acqtn_design.Rcrd_hf_freq, &seis_acqtn_design.Rcrd_hf_slope, &seis_acqtn_design.Rcrd_lf_freq, &seis_acqtn_design.Rcrd_lf_slope, &seis_acqtn_design.Rcrd_make, &seis_acqtn_design.Rcrd_near_surf_corr, &seis_acqtn_design.Rcrd_near_surf_corr_ouom, &seis_acqtn_design.Rcrd_nf_freq, &seis_acqtn_design.Rcrd_nf_ind, &seis_acqtn_design.Rcrd_polarity, &seis_acqtn_design.Rcrd_rec_length, &seis_acqtn_design.Rcrd_rec_length_ouom, &seis_acqtn_design.Rcrd_sample_rate, &seis_acqtn_design.Rcrd_sample_rate_ouom, &seis_acqtn_design.Rcvr_line_spacing, &seis_acqtn_design.Rcvr_line_spacing_ouom, &seis_acqtn_design.Remark, &seis_acqtn_design.Rep_water_acoustic_vel, &seis_acqtn_design.Rep_water_acoustic_vel_ouom, &seis_acqtn_design.Shot_by, &seis_acqtn_design.Shot_for, &seis_acqtn_design.Source, &seis_acqtn_design.Well_src_azimuth, &seis_acqtn_design.Well_src_azimuth_north, &seis_acqtn_design.Well_src_offset, &seis_acqtn_design.Well_src_offset_ouom, &seis_acqtn_design.Row_changed_by, &seis_acqtn_design.Row_changed_date, &seis_acqtn_design.Row_created_by, &seis_acqtn_design.Row_created_date, &seis_acqtn_design.Row_effective_date, &seis_acqtn_design.Row_expiry_date, &seis_acqtn_design.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_acqtn_design to result
		result = append(result, seis_acqtn_design)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisAcqtnDesign(c *fiber.Ctx) error {
	var seis_acqtn_design dto.Seis_acqtn_design

	setDefaults(&seis_acqtn_design)

	if err := json.Unmarshal(c.Body(), &seis_acqtn_design); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_acqtn_design values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77, :78, :79, :80, :81, :82, :83, :84, :85, :86, :87, :88, :89)")
	if err != nil {
		return err
	}
	seis_acqtn_design.Row_created_date = formatDate(seis_acqtn_design.Row_created_date)
	seis_acqtn_design.Row_changed_date = formatDate(seis_acqtn_design.Row_changed_date)
	seis_acqtn_design.Acqtn_completed_date = formatDateString(seis_acqtn_design.Acqtn_completed_date)
	seis_acqtn_design.Acqtn_start_date = formatDateString(seis_acqtn_design.Acqtn_start_date)
	seis_acqtn_design.Effective_date = formatDateString(seis_acqtn_design.Effective_date)
	seis_acqtn_design.Expiry_date = formatDateString(seis_acqtn_design.Expiry_date)
	seis_acqtn_design.Row_effective_date = formatDateString(seis_acqtn_design.Row_effective_date)
	seis_acqtn_design.Row_expiry_date = formatDateString(seis_acqtn_design.Row_expiry_date)






	rows, err := stmt.Exec(seis_acqtn_design.Acqtn_design_id, seis_acqtn_design.Acqtn_completed_date, seis_acqtn_design.Acqtn_completed_date_desc, seis_acqtn_design.Acqtn_dimension, seis_acqtn_design.Acqtn_direction, seis_acqtn_design.Acqtn_inline_bin_size, seis_acqtn_design.Acqtn_inline_bin_size_ouom, seis_acqtn_design.Acqtn_remark, seis_acqtn_design.Acqtn_shotpt_interval, seis_acqtn_design.Acqtn_shotpt_interval_ouom, seis_acqtn_design.Acqtn_shot_line_spacing, seis_acqtn_design.Acqtn_shot_line_spacing_ouom, seis_acqtn_design.Acqtn_shot_time_intvl, seis_acqtn_design.Acqtn_shot_time_intvl_ouom, seis_acqtn_design.Acqtn_start_date, seis_acqtn_design.Acqtn_start_date_desc, seis_acqtn_design.Acqtn_xline_bin_size, seis_acqtn_design.Acqtn_xline_bin_size_ouom, seis_acqtn_design.Active_ind, seis_acqtn_design.Actual_ind, seis_acqtn_design.Cdp_coverage, seis_acqtn_design.Effective_date, seis_acqtn_design.Energy_charge_size, seis_acqtn_design.Energy_charge_size_ouom, seis_acqtn_design.Energy_oprg_psr, seis_acqtn_design.Energy_oprg_psr_ouom, seis_acqtn_design.Energy_oprg_volume, seis_acqtn_design.Energy_oprg_volume_ouom, seis_acqtn_design.Energy_shot_depth, seis_acqtn_design.Energy_shot_depth_ouom, seis_acqtn_design.Energy_src_array_spc, seis_acqtn_design.Energy_src_array_spc_ouom, seis_acqtn_design.Energy_src_array_type, seis_acqtn_design.Energy_src_make, seis_acqtn_design.Energy_src_per_shot, seis_acqtn_design.Energy_sweep_duration, seis_acqtn_design.Energy_sweep_duration_ouom, seis_acqtn_design.Energy_sweep_end_freq, seis_acqtn_design.Energy_sweep_freq_ouom, seis_acqtn_design.Energy_sweep_mvup_dist, seis_acqtn_design.Energy_sweep_mvup_dist_ouom, seis_acqtn_design.Energy_sweep_no, seis_acqtn_design.Energy_sweep_st_freq, seis_acqtn_design.Energy_sweep_taper, seis_acqtn_design.Energy_sweep_taper_ouom, seis_acqtn_design.Energy_sweep_type, seis_acqtn_design.Energy_type, seis_acqtn_design.Environment, seis_acqtn_design.Expiry_date, seis_acqtn_design.Monitor_depth, seis_acqtn_design.Monitor_depth_ouom, seis_acqtn_design.Nominal_ind, seis_acqtn_design.Ppdm_guid, seis_acqtn_design.Rcrd_channel_count, seis_acqtn_design.Rcrd_format_type, seis_acqtn_design.Rcrd_gain_mode, seis_acqtn_design.Rcrd_hf_freq, seis_acqtn_design.Rcrd_hf_slope, seis_acqtn_design.Rcrd_lf_freq, seis_acqtn_design.Rcrd_lf_slope, seis_acqtn_design.Rcrd_make, seis_acqtn_design.Rcrd_near_surf_corr, seis_acqtn_design.Rcrd_near_surf_corr_ouom, seis_acqtn_design.Rcrd_nf_freq, seis_acqtn_design.Rcrd_nf_ind, seis_acqtn_design.Rcrd_polarity, seis_acqtn_design.Rcrd_rec_length, seis_acqtn_design.Rcrd_rec_length_ouom, seis_acqtn_design.Rcrd_sample_rate, seis_acqtn_design.Rcrd_sample_rate_ouom, seis_acqtn_design.Rcvr_line_spacing, seis_acqtn_design.Rcvr_line_spacing_ouom, seis_acqtn_design.Remark, seis_acqtn_design.Rep_water_acoustic_vel, seis_acqtn_design.Rep_water_acoustic_vel_ouom, seis_acqtn_design.Shot_by, seis_acqtn_design.Shot_for, seis_acqtn_design.Source, seis_acqtn_design.Well_src_azimuth, seis_acqtn_design.Well_src_azimuth_north, seis_acqtn_design.Well_src_offset, seis_acqtn_design.Well_src_offset_ouom, seis_acqtn_design.Row_changed_by, seis_acqtn_design.Row_changed_date, seis_acqtn_design.Row_created_by, seis_acqtn_design.Row_created_date, seis_acqtn_design.Row_effective_date, seis_acqtn_design.Row_expiry_date, seis_acqtn_design.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisAcqtnDesign(c *fiber.Ctx) error {
	var seis_acqtn_design dto.Seis_acqtn_design

	setDefaults(&seis_acqtn_design)

	if err := json.Unmarshal(c.Body(), &seis_acqtn_design); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_acqtn_design.Acqtn_design_id = id

    if seis_acqtn_design.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_acqtn_design set acqtn_completed_date = :1, acqtn_completed_date_desc = :2, acqtn_dimension = :3, acqtn_direction = :4, acqtn_inline_bin_size = :5, acqtn_inline_bin_size_ouom = :6, acqtn_remark = :7, acqtn_shotpt_interval = :8, acqtn_shotpt_interval_ouom = :9, acqtn_shot_line_spacing = :10, acqtn_shot_line_spacing_ouom = :11, acqtn_shot_time_intvl = :12, acqtn_shot_time_intvl_ouom = :13, acqtn_start_date = :14, acqtn_start_date_desc = :15, acqtn_xline_bin_size = :16, acqtn_xline_bin_size_ouom = :17, active_ind = :18, actual_ind = :19, cdp_coverage = :20, effective_date = :21, energy_charge_size = :22, energy_charge_size_ouom = :23, energy_oprg_psr = :24, energy_oprg_psr_ouom = :25, energy_oprg_volume = :26, energy_oprg_volume_ouom = :27, energy_shot_depth = :28, energy_shot_depth_ouom = :29, energy_src_array_spc = :30, energy_src_array_spc_ouom = :31, energy_src_array_type = :32, energy_src_make = :33, energy_src_per_shot = :34, energy_sweep_duration = :35, energy_sweep_duration_ouom = :36, energy_sweep_end_freq = :37, energy_sweep_freq_ouom = :38, energy_sweep_mvup_dist = :39, energy_sweep_mvup_dist_ouom = :40, energy_sweep_no = :41, energy_sweep_st_freq = :42, energy_sweep_taper = :43, energy_sweep_taper_ouom = :44, energy_sweep_type = :45, energy_type = :46, environment = :47, expiry_date = :48, monitor_depth = :49, monitor_depth_ouom = :50, nominal_ind = :51, ppdm_guid = :52, rcrd_channel_count = :53, rcrd_format_type = :54, rcrd_gain_mode = :55, rcrd_hf_freq = :56, rcrd_hf_slope = :57, rcrd_lf_freq = :58, rcrd_lf_slope = :59, rcrd_make = :60, rcrd_near_surf_corr = :61, rcrd_near_surf_corr_ouom = :62, rcrd_nf_freq = :63, rcrd_nf_ind = :64, rcrd_polarity = :65, rcrd_rec_length = :66, rcrd_rec_length_ouom = :67, rcrd_sample_rate = :68, rcrd_sample_rate_ouom = :69, rcvr_line_spacing = :70, rcvr_line_spacing_ouom = :71, remark = :72, rep_water_acoustic_vel = :73, rep_water_acoustic_vel_ouom = :74, shot_by = :75, shot_for = :76, source = :77, well_src_azimuth = :78, well_src_azimuth_north = :79, well_src_offset = :80, well_src_offset_ouom = :81, row_changed_by = :82, row_changed_date = :83, row_created_by = :84, row_effective_date = :85, row_expiry_date = :86, row_quality = :87 where acqtn_design_id = :89")
	if err != nil {
		return err
	}

	seis_acqtn_design.Row_changed_date = formatDate(seis_acqtn_design.Row_changed_date)
	seis_acqtn_design.Acqtn_completed_date = formatDateString(seis_acqtn_design.Acqtn_completed_date)
	seis_acqtn_design.Acqtn_start_date = formatDateString(seis_acqtn_design.Acqtn_start_date)
	seis_acqtn_design.Effective_date = formatDateString(seis_acqtn_design.Effective_date)
	seis_acqtn_design.Expiry_date = formatDateString(seis_acqtn_design.Expiry_date)
	seis_acqtn_design.Row_effective_date = formatDateString(seis_acqtn_design.Row_effective_date)
	seis_acqtn_design.Row_expiry_date = formatDateString(seis_acqtn_design.Row_expiry_date)






	rows, err := stmt.Exec(seis_acqtn_design.Acqtn_completed_date, seis_acqtn_design.Acqtn_completed_date_desc, seis_acqtn_design.Acqtn_dimension, seis_acqtn_design.Acqtn_direction, seis_acqtn_design.Acqtn_inline_bin_size, seis_acqtn_design.Acqtn_inline_bin_size_ouom, seis_acqtn_design.Acqtn_remark, seis_acqtn_design.Acqtn_shotpt_interval, seis_acqtn_design.Acqtn_shotpt_interval_ouom, seis_acqtn_design.Acqtn_shot_line_spacing, seis_acqtn_design.Acqtn_shot_line_spacing_ouom, seis_acqtn_design.Acqtn_shot_time_intvl, seis_acqtn_design.Acqtn_shot_time_intvl_ouom, seis_acqtn_design.Acqtn_start_date, seis_acqtn_design.Acqtn_start_date_desc, seis_acqtn_design.Acqtn_xline_bin_size, seis_acqtn_design.Acqtn_xline_bin_size_ouom, seis_acqtn_design.Active_ind, seis_acqtn_design.Actual_ind, seis_acqtn_design.Cdp_coverage, seis_acqtn_design.Effective_date, seis_acqtn_design.Energy_charge_size, seis_acqtn_design.Energy_charge_size_ouom, seis_acqtn_design.Energy_oprg_psr, seis_acqtn_design.Energy_oprg_psr_ouom, seis_acqtn_design.Energy_oprg_volume, seis_acqtn_design.Energy_oprg_volume_ouom, seis_acqtn_design.Energy_shot_depth, seis_acqtn_design.Energy_shot_depth_ouom, seis_acqtn_design.Energy_src_array_spc, seis_acqtn_design.Energy_src_array_spc_ouom, seis_acqtn_design.Energy_src_array_type, seis_acqtn_design.Energy_src_make, seis_acqtn_design.Energy_src_per_shot, seis_acqtn_design.Energy_sweep_duration, seis_acqtn_design.Energy_sweep_duration_ouom, seis_acqtn_design.Energy_sweep_end_freq, seis_acqtn_design.Energy_sweep_freq_ouom, seis_acqtn_design.Energy_sweep_mvup_dist, seis_acqtn_design.Energy_sweep_mvup_dist_ouom, seis_acqtn_design.Energy_sweep_no, seis_acqtn_design.Energy_sweep_st_freq, seis_acqtn_design.Energy_sweep_taper, seis_acqtn_design.Energy_sweep_taper_ouom, seis_acqtn_design.Energy_sweep_type, seis_acqtn_design.Energy_type, seis_acqtn_design.Environment, seis_acqtn_design.Expiry_date, seis_acqtn_design.Monitor_depth, seis_acqtn_design.Monitor_depth_ouom, seis_acqtn_design.Nominal_ind, seis_acqtn_design.Ppdm_guid, seis_acqtn_design.Rcrd_channel_count, seis_acqtn_design.Rcrd_format_type, seis_acqtn_design.Rcrd_gain_mode, seis_acqtn_design.Rcrd_hf_freq, seis_acqtn_design.Rcrd_hf_slope, seis_acqtn_design.Rcrd_lf_freq, seis_acqtn_design.Rcrd_lf_slope, seis_acqtn_design.Rcrd_make, seis_acqtn_design.Rcrd_near_surf_corr, seis_acqtn_design.Rcrd_near_surf_corr_ouom, seis_acqtn_design.Rcrd_nf_freq, seis_acqtn_design.Rcrd_nf_ind, seis_acqtn_design.Rcrd_polarity, seis_acqtn_design.Rcrd_rec_length, seis_acqtn_design.Rcrd_rec_length_ouom, seis_acqtn_design.Rcrd_sample_rate, seis_acqtn_design.Rcrd_sample_rate_ouom, seis_acqtn_design.Rcvr_line_spacing, seis_acqtn_design.Rcvr_line_spacing_ouom, seis_acqtn_design.Remark, seis_acqtn_design.Rep_water_acoustic_vel, seis_acqtn_design.Rep_water_acoustic_vel_ouom, seis_acqtn_design.Shot_by, seis_acqtn_design.Shot_for, seis_acqtn_design.Source, seis_acqtn_design.Well_src_azimuth, seis_acqtn_design.Well_src_azimuth_north, seis_acqtn_design.Well_src_offset, seis_acqtn_design.Well_src_offset_ouom, seis_acqtn_design.Row_changed_by, seis_acqtn_design.Row_changed_date, seis_acqtn_design.Row_created_by, seis_acqtn_design.Row_effective_date, seis_acqtn_design.Row_expiry_date, seis_acqtn_design.Row_quality, seis_acqtn_design.Acqtn_design_id)
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

func PatchSeisAcqtnDesign(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_acqtn_design set "
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
		} else if key == "acqtn_completed_date" || key == "acqtn_start_date" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where acqtn_design_id = :id"

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

func DeleteSeisAcqtnDesign(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_acqtn_design dto.Seis_acqtn_design
	seis_acqtn_design.Acqtn_design_id = id

	stmt, err := db.Prepare("delete from seis_acqtn_design where acqtn_design_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_acqtn_design.Acqtn_design_id)
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



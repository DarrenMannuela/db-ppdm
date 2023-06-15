package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellLicense(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_license")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_license

	for rows.Next() {
		var well_license dto.Well_license
		if err := rows.Scan(&well_license.Uwi, &well_license.Source, &well_license.License_id, &well_license.Active_ind, &well_license.Agent, &well_license.Application_id, &well_license.Authorized_strat_unit_id, &well_license.Bidding_round_num, &well_license.Contractor, &well_license.Direction_to_loc, &well_license.Direction_to_loc_ouom, &well_license.Distance_ref_point, &well_license.Distance_to_loc, &well_license.Distance_to_loc_ouom, &well_license.Drill_rig_num, &well_license.Drill_slot_no, &well_license.Drill_tool, &well_license.Effective_date, &well_license.Exception_granted, &well_license.Exception_requested, &well_license.Expired_ind, &well_license.Expiry_date, &well_license.Fees_paid_ind, &well_license.Licensee, &well_license.Licensee_contact_id, &well_license.License_date, &well_license.License_num, &well_license.No_of_wells, &well_license.Offshore_completion_type, &well_license.Permit_reference_num, &well_license.Permit_reissue_date, &well_license.Permit_type, &well_license.Platform_name, &well_license.Ppdm_guid, &well_license.Projected_depth, &well_license.Projected_depth_ouom, &well_license.Projected_strat_unit_id, &well_license.Projected_tvd, &well_license.Projected_tvd_ouom, &well_license.Proposed_spud_date, &well_license.Purpose, &well_license.Rate_schedule_id, &well_license.Regulation_text, &well_license.Regulation_text_source_doc, &well_license.Regulatory_agency, &well_license.Regulatory_contact_id, &well_license.Remark, &well_license.Rig_code, &well_license.Rig_sf_subtype, &well_license.Rig_substr_height, &well_license.Rig_substr_height_ouom, &well_license.Rig_support_facility_id, &well_license.Rig_type, &well_license.Section_of_regulation, &well_license.Strat_name_set_id, &well_license.Surveyor, &well_license.Target_objective_fluid, &well_license.Violation_ind, &well_license.Row_changed_by, &well_license.Row_changed_date, &well_license.Row_created_by, &well_license.Row_created_date, &well_license.Row_effective_date, &well_license.Row_expiry_date, &well_license.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_license to result
		result = append(result, well_license)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellLicense(c *fiber.Ctx) error {
	var well_license dto.Well_license

	setDefaults(&well_license)

	if err := json.Unmarshal(c.Body(), &well_license); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_license values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65)")
	if err != nil {
		return err
	}
	well_license.Row_created_date = formatDate(well_license.Row_created_date)
	well_license.Row_changed_date = formatDate(well_license.Row_changed_date)
	well_license.Effective_date = formatDateString(well_license.Effective_date)
	well_license.Expiry_date = formatDateString(well_license.Expiry_date)
	well_license.License_date = formatDateString(well_license.License_date)
	well_license.Permit_reissue_date = formatDateString(well_license.Permit_reissue_date)
	well_license.Proposed_spud_date = formatDateString(well_license.Proposed_spud_date)
	well_license.Row_effective_date = formatDateString(well_license.Row_effective_date)
	well_license.Row_expiry_date = formatDateString(well_license.Row_expiry_date)






	rows, err := stmt.Exec(well_license.Uwi, well_license.Source, well_license.License_id, well_license.Active_ind, well_license.Agent, well_license.Application_id, well_license.Authorized_strat_unit_id, well_license.Bidding_round_num, well_license.Contractor, well_license.Direction_to_loc, well_license.Direction_to_loc_ouom, well_license.Distance_ref_point, well_license.Distance_to_loc, well_license.Distance_to_loc_ouom, well_license.Drill_rig_num, well_license.Drill_slot_no, well_license.Drill_tool, well_license.Effective_date, well_license.Exception_granted, well_license.Exception_requested, well_license.Expired_ind, well_license.Expiry_date, well_license.Fees_paid_ind, well_license.Licensee, well_license.Licensee_contact_id, well_license.License_date, well_license.License_num, well_license.No_of_wells, well_license.Offshore_completion_type, well_license.Permit_reference_num, well_license.Permit_reissue_date, well_license.Permit_type, well_license.Platform_name, well_license.Ppdm_guid, well_license.Projected_depth, well_license.Projected_depth_ouom, well_license.Projected_strat_unit_id, well_license.Projected_tvd, well_license.Projected_tvd_ouom, well_license.Proposed_spud_date, well_license.Purpose, well_license.Rate_schedule_id, well_license.Regulation_text, well_license.Regulation_text_source_doc, well_license.Regulatory_agency, well_license.Regulatory_contact_id, well_license.Remark, well_license.Rig_code, well_license.Rig_sf_subtype, well_license.Rig_substr_height, well_license.Rig_substr_height_ouom, well_license.Rig_support_facility_id, well_license.Rig_type, well_license.Section_of_regulation, well_license.Strat_name_set_id, well_license.Surveyor, well_license.Target_objective_fluid, well_license.Violation_ind, well_license.Row_changed_by, well_license.Row_changed_date, well_license.Row_created_by, well_license.Row_created_date, well_license.Row_effective_date, well_license.Row_expiry_date, well_license.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellLicense(c *fiber.Ctx) error {
	var well_license dto.Well_license

	setDefaults(&well_license)

	if err := json.Unmarshal(c.Body(), &well_license); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_license.Uwi = id

    if well_license.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_license set source = :1, license_id = :2, active_ind = :3, agent = :4, application_id = :5, authorized_strat_unit_id = :6, bidding_round_num = :7, contractor = :8, direction_to_loc = :9, direction_to_loc_ouom = :10, distance_ref_point = :11, distance_to_loc = :12, distance_to_loc_ouom = :13, drill_rig_num = :14, drill_slot_no = :15, drill_tool = :16, effective_date = :17, exception_granted = :18, exception_requested = :19, expired_ind = :20, expiry_date = :21, fees_paid_ind = :22, licensee = :23, licensee_contact_id = :24, license_date = :25, license_num = :26, no_of_wells = :27, offshore_completion_type = :28, permit_reference_num = :29, permit_reissue_date = :30, permit_type = :31, platform_name = :32, ppdm_guid = :33, projected_depth = :34, projected_depth_ouom = :35, projected_strat_unit_id = :36, projected_tvd = :37, projected_tvd_ouom = :38, proposed_spud_date = :39, purpose = :40, rate_schedule_id = :41, regulation_text = :42, regulation_text_source_doc = :43, regulatory_agency = :44, regulatory_contact_id = :45, remark = :46, rig_code = :47, rig_sf_subtype = :48, rig_substr_height = :49, rig_substr_height_ouom = :50, rig_support_facility_id = :51, rig_type = :52, section_of_regulation = :53, strat_name_set_id = :54, surveyor = :55, target_objective_fluid = :56, violation_ind = :57, row_changed_by = :58, row_changed_date = :59, row_created_by = :60, row_effective_date = :61, row_expiry_date = :62, row_quality = :63 where uwi = :65")
	if err != nil {
		return err
	}

	well_license.Row_changed_date = formatDate(well_license.Row_changed_date)
	well_license.Effective_date = formatDateString(well_license.Effective_date)
	well_license.Expiry_date = formatDateString(well_license.Expiry_date)
	well_license.License_date = formatDateString(well_license.License_date)
	well_license.Permit_reissue_date = formatDateString(well_license.Permit_reissue_date)
	well_license.Proposed_spud_date = formatDateString(well_license.Proposed_spud_date)
	well_license.Row_effective_date = formatDateString(well_license.Row_effective_date)
	well_license.Row_expiry_date = formatDateString(well_license.Row_expiry_date)






	rows, err := stmt.Exec(well_license.Source, well_license.License_id, well_license.Active_ind, well_license.Agent, well_license.Application_id, well_license.Authorized_strat_unit_id, well_license.Bidding_round_num, well_license.Contractor, well_license.Direction_to_loc, well_license.Direction_to_loc_ouom, well_license.Distance_ref_point, well_license.Distance_to_loc, well_license.Distance_to_loc_ouom, well_license.Drill_rig_num, well_license.Drill_slot_no, well_license.Drill_tool, well_license.Effective_date, well_license.Exception_granted, well_license.Exception_requested, well_license.Expired_ind, well_license.Expiry_date, well_license.Fees_paid_ind, well_license.Licensee, well_license.Licensee_contact_id, well_license.License_date, well_license.License_num, well_license.No_of_wells, well_license.Offshore_completion_type, well_license.Permit_reference_num, well_license.Permit_reissue_date, well_license.Permit_type, well_license.Platform_name, well_license.Ppdm_guid, well_license.Projected_depth, well_license.Projected_depth_ouom, well_license.Projected_strat_unit_id, well_license.Projected_tvd, well_license.Projected_tvd_ouom, well_license.Proposed_spud_date, well_license.Purpose, well_license.Rate_schedule_id, well_license.Regulation_text, well_license.Regulation_text_source_doc, well_license.Regulatory_agency, well_license.Regulatory_contact_id, well_license.Remark, well_license.Rig_code, well_license.Rig_sf_subtype, well_license.Rig_substr_height, well_license.Rig_substr_height_ouom, well_license.Rig_support_facility_id, well_license.Rig_type, well_license.Section_of_regulation, well_license.Strat_name_set_id, well_license.Surveyor, well_license.Target_objective_fluid, well_license.Violation_ind, well_license.Row_changed_by, well_license.Row_changed_date, well_license.Row_created_by, well_license.Row_effective_date, well_license.Row_expiry_date, well_license.Row_quality, well_license.Uwi)
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

func PatchWellLicense(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_license set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "license_date" || key == "permit_reissue_date" || key == "proposed_spud_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellLicense(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_license dto.Well_license
	well_license.Uwi = id

	stmt, err := db.Prepare("delete from well_license where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_license.Uwi)
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



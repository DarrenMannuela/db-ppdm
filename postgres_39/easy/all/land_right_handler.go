package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLandRight(c *fiber.Ctx) error {
	rows, err := db.Query("select * from land_right")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Land_right

	for rows.Next() {
		var land_right dto.Land_right
		if err := rows.Scan(&land_right.Land_right_subtype, &land_right.Land_right_id, &land_right.Acqtn_date, &land_right.Active_ind, &land_right.Area_id, &land_right.Area_type, &land_right.Calculated_interest, &land_right.Case_serial_num, &land_right.Confidential_ind, &land_right.Cost_center_num, &land_right.Effective_date, &land_right.Expiry_date, &land_right.First_platform_approve_date, &land_right.First_platform_install_date, &land_right.First_production_date, &land_right.Fractional_interest, &land_right.Granted_right_type, &land_right.Gross_size, &land_right.Gross_size_ouom, &land_right.High_water_depth, &land_right.High_water_depth_ouom, &land_right.Inactivation_date, &land_right.Incent_cert_no, &land_right.Issue_date, &land_right.Jurisdiction, &land_right.Land_acqtn_method, &land_right.Land_case_action, &land_right.Land_case_type, &land_right.Land_property_type, &land_right.Land_right_category, &land_right.Land_sale_number, &land_right.Land_sale_offering_id, &land_right.Last_action_date, &land_right.Lessor_name, &land_right.Lessor_num, &land_right.Lessor_type, &land_right.Low_water_depth, &land_right.Low_water_depth_ouom, &land_right.Municipality, &land_right.Occupant_name, &land_right.Offshore_distance, &land_right.Offshore_distance_ouom, &land_right.Oil_sand_deposit, &land_right.Overlap_ind, &land_right.Platform_count, &land_right.Ppdm_guid, &land_right.Primary_term, &land_right.Primary_term_ouom, &land_right.Producing_ind, &land_right.Proprietary_num, &land_right.Qualification_date, &land_right.Reclamation_cert_num, &land_right.Reclamation_end_date, &land_right.Reclamation_start_date, &land_right.Relinquishable_ind, &land_right.Remark, &land_right.Rental_allocation_ind, &land_right.Report_acreage_ind, &land_right.Scheme_approval_num, &land_right.Scheme_expiry_date, &land_right.Secondary_term, &land_right.Secondary_term_ouom, &land_right.Source, &land_right.Subsurface_ind, &land_right.Surface_ind, &land_right.Termin_notice_period, &land_right.Termin_notice_period_ouom, &land_right.Term_exiry_date, &land_right.Unit_operated_ind, &land_right.Well_qualific_type, &land_right.Row_changed_by, &land_right.Row_changed_date, &land_right.Row_created_by, &land_right.Row_created_date, &land_right.Row_effective_date, &land_right.Row_expiry_date, &land_right.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append land_right to result
		result = append(result, land_right)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLandRight(c *fiber.Ctx) error {
	var land_right dto.Land_right

	setDefaults(&land_right)

	if err := json.Unmarshal(c.Body(), &land_right); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into land_right values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53, :54, :55, :56, :57, :58, :59, :60, :61, :62, :63, :64, :65, :66, :67, :68, :69, :70, :71, :72, :73, :74, :75, :76, :77)")
	if err != nil {
		return err
	}
	land_right.Row_created_date = formatDate(land_right.Row_created_date)
	land_right.Row_changed_date = formatDate(land_right.Row_changed_date)
	land_right.Acqtn_date = formatDateString(land_right.Acqtn_date)
	land_right.Effective_date = formatDateString(land_right.Effective_date)
	land_right.Expiry_date = formatDateString(land_right.Expiry_date)
	land_right.First_platform_approve_date = formatDateString(land_right.First_platform_approve_date)
	land_right.First_platform_install_date = formatDateString(land_right.First_platform_install_date)
	land_right.First_production_date = formatDateString(land_right.First_production_date)
	land_right.Inactivation_date = formatDateString(land_right.Inactivation_date)
	land_right.Issue_date = formatDateString(land_right.Issue_date)
	land_right.Last_action_date = formatDateString(land_right.Last_action_date)
	land_right.Qualification_date = formatDateString(land_right.Qualification_date)
	land_right.Reclamation_end_date = formatDateString(land_right.Reclamation_end_date)
	land_right.Reclamation_start_date = formatDateString(land_right.Reclamation_start_date)
	land_right.Scheme_expiry_date = formatDateString(land_right.Scheme_expiry_date)
	land_right.Term_exiry_date = formatDateString(land_right.Term_exiry_date)
	land_right.Row_effective_date = formatDateString(land_right.Row_effective_date)
	land_right.Row_expiry_date = formatDateString(land_right.Row_expiry_date)






	rows, err := stmt.Exec(land_right.Land_right_subtype, land_right.Land_right_id, land_right.Acqtn_date, land_right.Active_ind, land_right.Area_id, land_right.Area_type, land_right.Calculated_interest, land_right.Case_serial_num, land_right.Confidential_ind, land_right.Cost_center_num, land_right.Effective_date, land_right.Expiry_date, land_right.First_platform_approve_date, land_right.First_platform_install_date, land_right.First_production_date, land_right.Fractional_interest, land_right.Granted_right_type, land_right.Gross_size, land_right.Gross_size_ouom, land_right.High_water_depth, land_right.High_water_depth_ouom, land_right.Inactivation_date, land_right.Incent_cert_no, land_right.Issue_date, land_right.Jurisdiction, land_right.Land_acqtn_method, land_right.Land_case_action, land_right.Land_case_type, land_right.Land_property_type, land_right.Land_right_category, land_right.Land_sale_number, land_right.Land_sale_offering_id, land_right.Last_action_date, land_right.Lessor_name, land_right.Lessor_num, land_right.Lessor_type, land_right.Low_water_depth, land_right.Low_water_depth_ouom, land_right.Municipality, land_right.Occupant_name, land_right.Offshore_distance, land_right.Offshore_distance_ouom, land_right.Oil_sand_deposit, land_right.Overlap_ind, land_right.Platform_count, land_right.Ppdm_guid, land_right.Primary_term, land_right.Primary_term_ouom, land_right.Producing_ind, land_right.Proprietary_num, land_right.Qualification_date, land_right.Reclamation_cert_num, land_right.Reclamation_end_date, land_right.Reclamation_start_date, land_right.Relinquishable_ind, land_right.Remark, land_right.Rental_allocation_ind, land_right.Report_acreage_ind, land_right.Scheme_approval_num, land_right.Scheme_expiry_date, land_right.Secondary_term, land_right.Secondary_term_ouom, land_right.Source, land_right.Subsurface_ind, land_right.Surface_ind, land_right.Termin_notice_period, land_right.Termin_notice_period_ouom, land_right.Term_exiry_date, land_right.Unit_operated_ind, land_right.Well_qualific_type, land_right.Row_changed_by, land_right.Row_changed_date, land_right.Row_created_by, land_right.Row_created_date, land_right.Row_effective_date, land_right.Row_expiry_date, land_right.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLandRight(c *fiber.Ctx) error {
	var land_right dto.Land_right

	setDefaults(&land_right)

	if err := json.Unmarshal(c.Body(), &land_right); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	land_right.Land_right_subtype = id

    if land_right.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update land_right set land_right_id = :1, acqtn_date = :2, active_ind = :3, area_id = :4, area_type = :5, calculated_interest = :6, case_serial_num = :7, confidential_ind = :8, cost_center_num = :9, effective_date = :10, expiry_date = :11, first_platform_approve_date = :12, first_platform_install_date = :13, first_production_date = :14, fractional_interest = :15, granted_right_type = :16, gross_size = :17, gross_size_ouom = :18, high_water_depth = :19, high_water_depth_ouom = :20, inactivation_date = :21, incent_cert_no = :22, issue_date = :23, jurisdiction = :24, land_acqtn_method = :25, land_case_action = :26, land_case_type = :27, land_property_type = :28, land_right_category = :29, land_sale_number = :30, land_sale_offering_id = :31, last_action_date = :32, lessor_name = :33, lessor_num = :34, lessor_type = :35, low_water_depth = :36, low_water_depth_ouom = :37, municipality = :38, occupant_name = :39, offshore_distance = :40, offshore_distance_ouom = :41, oil_sand_deposit = :42, overlap_ind = :43, platform_count = :44, ppdm_guid = :45, primary_term = :46, primary_term_ouom = :47, producing_ind = :48, proprietary_num = :49, qualification_date = :50, reclamation_cert_num = :51, reclamation_end_date = :52, reclamation_start_date = :53, relinquishable_ind = :54, remark = :55, rental_allocation_ind = :56, report_acreage_ind = :57, scheme_approval_num = :58, scheme_expiry_date = :59, secondary_term = :60, secondary_term_ouom = :61, source = :62, subsurface_ind = :63, surface_ind = :64, termin_notice_period = :65, termin_notice_period_ouom = :66, term_exiry_date = :67, unit_operated_ind = :68, well_qualific_type = :69, row_changed_by = :70, row_changed_date = :71, row_created_by = :72, row_effective_date = :73, row_expiry_date = :74, row_quality = :75 where land_right_subtype = :77")
	if err != nil {
		return err
	}

	land_right.Row_changed_date = formatDate(land_right.Row_changed_date)
	land_right.Acqtn_date = formatDateString(land_right.Acqtn_date)
	land_right.Effective_date = formatDateString(land_right.Effective_date)
	land_right.Expiry_date = formatDateString(land_right.Expiry_date)
	land_right.First_platform_approve_date = formatDateString(land_right.First_platform_approve_date)
	land_right.First_platform_install_date = formatDateString(land_right.First_platform_install_date)
	land_right.First_production_date = formatDateString(land_right.First_production_date)
	land_right.Inactivation_date = formatDateString(land_right.Inactivation_date)
	land_right.Issue_date = formatDateString(land_right.Issue_date)
	land_right.Last_action_date = formatDateString(land_right.Last_action_date)
	land_right.Qualification_date = formatDateString(land_right.Qualification_date)
	land_right.Reclamation_end_date = formatDateString(land_right.Reclamation_end_date)
	land_right.Reclamation_start_date = formatDateString(land_right.Reclamation_start_date)
	land_right.Scheme_expiry_date = formatDateString(land_right.Scheme_expiry_date)
	land_right.Term_exiry_date = formatDateString(land_right.Term_exiry_date)
	land_right.Row_effective_date = formatDateString(land_right.Row_effective_date)
	land_right.Row_expiry_date = formatDateString(land_right.Row_expiry_date)






	rows, err := stmt.Exec(land_right.Land_right_id, land_right.Acqtn_date, land_right.Active_ind, land_right.Area_id, land_right.Area_type, land_right.Calculated_interest, land_right.Case_serial_num, land_right.Confidential_ind, land_right.Cost_center_num, land_right.Effective_date, land_right.Expiry_date, land_right.First_platform_approve_date, land_right.First_platform_install_date, land_right.First_production_date, land_right.Fractional_interest, land_right.Granted_right_type, land_right.Gross_size, land_right.Gross_size_ouom, land_right.High_water_depth, land_right.High_water_depth_ouom, land_right.Inactivation_date, land_right.Incent_cert_no, land_right.Issue_date, land_right.Jurisdiction, land_right.Land_acqtn_method, land_right.Land_case_action, land_right.Land_case_type, land_right.Land_property_type, land_right.Land_right_category, land_right.Land_sale_number, land_right.Land_sale_offering_id, land_right.Last_action_date, land_right.Lessor_name, land_right.Lessor_num, land_right.Lessor_type, land_right.Low_water_depth, land_right.Low_water_depth_ouom, land_right.Municipality, land_right.Occupant_name, land_right.Offshore_distance, land_right.Offshore_distance_ouom, land_right.Oil_sand_deposit, land_right.Overlap_ind, land_right.Platform_count, land_right.Ppdm_guid, land_right.Primary_term, land_right.Primary_term_ouom, land_right.Producing_ind, land_right.Proprietary_num, land_right.Qualification_date, land_right.Reclamation_cert_num, land_right.Reclamation_end_date, land_right.Reclamation_start_date, land_right.Relinquishable_ind, land_right.Remark, land_right.Rental_allocation_ind, land_right.Report_acreage_ind, land_right.Scheme_approval_num, land_right.Scheme_expiry_date, land_right.Secondary_term, land_right.Secondary_term_ouom, land_right.Source, land_right.Subsurface_ind, land_right.Surface_ind, land_right.Termin_notice_period, land_right.Termin_notice_period_ouom, land_right.Term_exiry_date, land_right.Unit_operated_ind, land_right.Well_qualific_type, land_right.Row_changed_by, land_right.Row_changed_date, land_right.Row_created_by, land_right.Row_effective_date, land_right.Row_expiry_date, land_right.Row_quality, land_right.Land_right_subtype)
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

func PatchLandRight(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update land_right set "
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
		} else if key == "acqtn_date" || key == "effective_date" || key == "expiry_date" || key == "first_platform_approve_date" || key == "first_platform_install_date" || key == "first_production_date" || key == "inactivation_date" || key == "issue_date" || key == "last_action_date" || key == "qualification_date" || key == "reclamation_end_date" || key == "reclamation_start_date" || key == "scheme_expiry_date" || key == "term_exiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where land_right_subtype = :id"

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

func DeleteLandRight(c *fiber.Ctx) error {
	id := c.Params("id")
	var land_right dto.Land_right
	land_right.Land_right_subtype = id

	stmt, err := db.Prepare("delete from land_right where land_right_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(land_right.Land_right_subtype)
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



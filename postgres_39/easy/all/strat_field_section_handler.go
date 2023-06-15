package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetStratFieldSection(c *fiber.Ctx) error {
	rows, err := db.Query("select * from strat_field_section")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Strat_field_section

	for rows.Next() {
		var strat_field_section dto.Strat_field_section
		if err := rows.Scan(&strat_field_section.Field_station_id, &strat_field_section.Strat_name_set_id, &strat_field_section.Strat_unit_id, &strat_field_section.Interp_id, &strat_field_section.Active_ind, &strat_field_section.Area_id, &strat_field_section.Area_type, &strat_field_section.Azimuth_north_type, &strat_field_section.Certified_ind, &strat_field_section.Conformity_relationship, &strat_field_section.Dip_angle, &strat_field_section.Dip_direction, &strat_field_section.Effective_date, &strat_field_section.Expiry_date, &strat_field_section.Fault_heave, &strat_field_section.Fault_throw, &strat_field_section.Interpretation_version, &strat_field_section.Interpreter, &strat_field_section.Lithology, &strat_field_section.Missing_section, &strat_field_section.Missing_strat_type, &strat_field_section.Ordinal_seq_no, &strat_field_section.Overturned_ind, &strat_field_section.Pick_date, &strat_field_section.Pick_location, &strat_field_section.Pick_qualifier, &strat_field_section.Pick_qualif_reason, &strat_field_section.Pick_quality, &strat_field_section.Pick_version_type, &strat_field_section.Ppdm_guid, &strat_field_section.Preferred_pick_ind, &strat_field_section.Remark, &strat_field_section.Repeat_section, &strat_field_section.Repeat_strat_occur_no, &strat_field_section.Repeat_strat_type, &strat_field_section.Source, &strat_field_section.Source_document_id, &strat_field_section.Strat_interpret_method, &strat_field_section.Strat_unit_md, &strat_field_section.Strat_unit_md_ouom, &strat_field_section.Strat_unit_tvd, &strat_field_section.Strike, &strat_field_section.Sw_application_id, &strat_field_section.True_thickness_ind, &strat_field_section.Tvd_method, &strat_field_section.Version_obs_no, &strat_field_section.Row_changed_by, &strat_field_section.Row_changed_date, &strat_field_section.Row_created_by, &strat_field_section.Row_created_date, &strat_field_section.Row_effective_date, &strat_field_section.Row_expiry_date, &strat_field_section.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append strat_field_section to result
		result = append(result, strat_field_section)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetStratFieldSection(c *fiber.Ctx) error {
	var strat_field_section dto.Strat_field_section

	setDefaults(&strat_field_section)

	if err := json.Unmarshal(c.Body(), &strat_field_section); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into strat_field_section values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51, :52, :53)")
	if err != nil {
		return err
	}
	strat_field_section.Row_created_date = formatDate(strat_field_section.Row_created_date)
	strat_field_section.Row_changed_date = formatDate(strat_field_section.Row_changed_date)
	strat_field_section.Effective_date = formatDateString(strat_field_section.Effective_date)
	strat_field_section.Expiry_date = formatDateString(strat_field_section.Expiry_date)
	strat_field_section.Pick_date = formatDateString(strat_field_section.Pick_date)
	strat_field_section.Row_effective_date = formatDateString(strat_field_section.Row_effective_date)
	strat_field_section.Row_expiry_date = formatDateString(strat_field_section.Row_expiry_date)






	rows, err := stmt.Exec(strat_field_section.Field_station_id, strat_field_section.Strat_name_set_id, strat_field_section.Strat_unit_id, strat_field_section.Interp_id, strat_field_section.Active_ind, strat_field_section.Area_id, strat_field_section.Area_type, strat_field_section.Azimuth_north_type, strat_field_section.Certified_ind, strat_field_section.Conformity_relationship, strat_field_section.Dip_angle, strat_field_section.Dip_direction, strat_field_section.Effective_date, strat_field_section.Expiry_date, strat_field_section.Fault_heave, strat_field_section.Fault_throw, strat_field_section.Interpretation_version, strat_field_section.Interpreter, strat_field_section.Lithology, strat_field_section.Missing_section, strat_field_section.Missing_strat_type, strat_field_section.Ordinal_seq_no, strat_field_section.Overturned_ind, strat_field_section.Pick_date, strat_field_section.Pick_location, strat_field_section.Pick_qualifier, strat_field_section.Pick_qualif_reason, strat_field_section.Pick_quality, strat_field_section.Pick_version_type, strat_field_section.Ppdm_guid, strat_field_section.Preferred_pick_ind, strat_field_section.Remark, strat_field_section.Repeat_section, strat_field_section.Repeat_strat_occur_no, strat_field_section.Repeat_strat_type, strat_field_section.Source, strat_field_section.Source_document_id, strat_field_section.Strat_interpret_method, strat_field_section.Strat_unit_md, strat_field_section.Strat_unit_md_ouom, strat_field_section.Strat_unit_tvd, strat_field_section.Strike, strat_field_section.Sw_application_id, strat_field_section.True_thickness_ind, strat_field_section.Tvd_method, strat_field_section.Version_obs_no, strat_field_section.Row_changed_by, strat_field_section.Row_changed_date, strat_field_section.Row_created_by, strat_field_section.Row_created_date, strat_field_section.Row_effective_date, strat_field_section.Row_expiry_date, strat_field_section.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateStratFieldSection(c *fiber.Ctx) error {
	var strat_field_section dto.Strat_field_section

	setDefaults(&strat_field_section)

	if err := json.Unmarshal(c.Body(), &strat_field_section); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	strat_field_section.Field_station_id = id

    if strat_field_section.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update strat_field_section set strat_name_set_id = :1, strat_unit_id = :2, interp_id = :3, active_ind = :4, area_id = :5, area_type = :6, azimuth_north_type = :7, certified_ind = :8, conformity_relationship = :9, dip_angle = :10, dip_direction = :11, effective_date = :12, expiry_date = :13, fault_heave = :14, fault_throw = :15, interpretation_version = :16, interpreter = :17, lithology = :18, missing_section = :19, missing_strat_type = :20, ordinal_seq_no = :21, overturned_ind = :22, pick_date = :23, pick_location = :24, pick_qualifier = :25, pick_qualif_reason = :26, pick_quality = :27, pick_version_type = :28, ppdm_guid = :29, preferred_pick_ind = :30, remark = :31, repeat_section = :32, repeat_strat_occur_no = :33, repeat_strat_type = :34, source = :35, source_document_id = :36, strat_interpret_method = :37, strat_unit_md = :38, strat_unit_md_ouom = :39, strat_unit_tvd = :40, strike = :41, sw_application_id = :42, true_thickness_ind = :43, tvd_method = :44, version_obs_no = :45, row_changed_by = :46, row_changed_date = :47, row_created_by = :48, row_effective_date = :49, row_expiry_date = :50, row_quality = :51 where field_station_id = :53")
	if err != nil {
		return err
	}

	strat_field_section.Row_changed_date = formatDate(strat_field_section.Row_changed_date)
	strat_field_section.Effective_date = formatDateString(strat_field_section.Effective_date)
	strat_field_section.Expiry_date = formatDateString(strat_field_section.Expiry_date)
	strat_field_section.Pick_date = formatDateString(strat_field_section.Pick_date)
	strat_field_section.Row_effective_date = formatDateString(strat_field_section.Row_effective_date)
	strat_field_section.Row_expiry_date = formatDateString(strat_field_section.Row_expiry_date)






	rows, err := stmt.Exec(strat_field_section.Strat_name_set_id, strat_field_section.Strat_unit_id, strat_field_section.Interp_id, strat_field_section.Active_ind, strat_field_section.Area_id, strat_field_section.Area_type, strat_field_section.Azimuth_north_type, strat_field_section.Certified_ind, strat_field_section.Conformity_relationship, strat_field_section.Dip_angle, strat_field_section.Dip_direction, strat_field_section.Effective_date, strat_field_section.Expiry_date, strat_field_section.Fault_heave, strat_field_section.Fault_throw, strat_field_section.Interpretation_version, strat_field_section.Interpreter, strat_field_section.Lithology, strat_field_section.Missing_section, strat_field_section.Missing_strat_type, strat_field_section.Ordinal_seq_no, strat_field_section.Overturned_ind, strat_field_section.Pick_date, strat_field_section.Pick_location, strat_field_section.Pick_qualifier, strat_field_section.Pick_qualif_reason, strat_field_section.Pick_quality, strat_field_section.Pick_version_type, strat_field_section.Ppdm_guid, strat_field_section.Preferred_pick_ind, strat_field_section.Remark, strat_field_section.Repeat_section, strat_field_section.Repeat_strat_occur_no, strat_field_section.Repeat_strat_type, strat_field_section.Source, strat_field_section.Source_document_id, strat_field_section.Strat_interpret_method, strat_field_section.Strat_unit_md, strat_field_section.Strat_unit_md_ouom, strat_field_section.Strat_unit_tvd, strat_field_section.Strike, strat_field_section.Sw_application_id, strat_field_section.True_thickness_ind, strat_field_section.Tvd_method, strat_field_section.Version_obs_no, strat_field_section.Row_changed_by, strat_field_section.Row_changed_date, strat_field_section.Row_created_by, strat_field_section.Row_effective_date, strat_field_section.Row_expiry_date, strat_field_section.Row_quality, strat_field_section.Field_station_id)
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

func PatchStratFieldSection(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update strat_field_section set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "pick_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where field_station_id = :id"

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

func DeleteStratFieldSection(c *fiber.Ctx) error {
	id := c.Params("id")
	var strat_field_section dto.Strat_field_section
	strat_field_section.Field_station_id = id

	stmt, err := db.Prepare("delete from strat_field_section where field_station_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(strat_field_section.Field_station_id)
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



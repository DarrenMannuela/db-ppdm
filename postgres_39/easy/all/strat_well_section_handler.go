package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetStratWellSection(c *fiber.Ctx) error {
	rows, err := db.Query("select * from strat_well_section")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Strat_well_section

	for rows.Next() {
		var strat_well_section dto.Strat_well_section
		if err := rows.Scan(&strat_well_section.Uwi, &strat_well_section.Strat_name_set_id, &strat_well_section.Strat_unit_id, &strat_well_section.Interp_id, &strat_well_section.Active_ind, &strat_well_section.Area_id, &strat_well_section.Area_type, &strat_well_section.Azimuth_north_type, &strat_well_section.Certified_ind, &strat_well_section.Conformity_relationship, &strat_well_section.Dip_angle, &strat_well_section.Dip_direction, &strat_well_section.Dominant_lithology, &strat_well_section.Effective_date, &strat_well_section.Expiry_date, &strat_well_section.Fault_heave, &strat_well_section.Fault_throw, &strat_well_section.Interpreter, &strat_well_section.Missing_section, &strat_well_section.Missing_strat_type, &strat_well_section.Ordinal_seq_no, &strat_well_section.Overturned_ind, &strat_well_section.Pick_date, &strat_well_section.Pick_depth, &strat_well_section.Pick_depth_ouom, &strat_well_section.Pick_location, &strat_well_section.Pick_qualifier, &strat_well_section.Pick_qualif_reason, &strat_well_section.Pick_quality, &strat_well_section.Pick_tvd, &strat_well_section.Pick_version_type, &strat_well_section.Ppdm_guid, &strat_well_section.Preferred_pick_ind, &strat_well_section.Remark, &strat_well_section.Repeat_section, &strat_well_section.Repeat_strat_occur_no, &strat_well_section.Repeat_strat_type, &strat_well_section.Source, &strat_well_section.Source_document_id, &strat_well_section.Strat_interpret_method, &strat_well_section.Strike, &strat_well_section.Sw_application_id, &strat_well_section.Tvd_method, &strat_well_section.Version_obs_no, &strat_well_section.Row_changed_by, &strat_well_section.Row_changed_date, &strat_well_section.Row_created_by, &strat_well_section.Row_created_date, &strat_well_section.Row_effective_date, &strat_well_section.Row_expiry_date, &strat_well_section.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append strat_well_section to result
		result = append(result, strat_well_section)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetStratWellSection(c *fiber.Ctx) error {
	var strat_well_section dto.Strat_well_section

	setDefaults(&strat_well_section)

	if err := json.Unmarshal(c.Body(), &strat_well_section); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into strat_well_section values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49, :50, :51)")
	if err != nil {
		return err
	}
	strat_well_section.Row_created_date = formatDate(strat_well_section.Row_created_date)
	strat_well_section.Row_changed_date = formatDate(strat_well_section.Row_changed_date)
	strat_well_section.Effective_date = formatDateString(strat_well_section.Effective_date)
	strat_well_section.Expiry_date = formatDateString(strat_well_section.Expiry_date)
	strat_well_section.Pick_date = formatDateString(strat_well_section.Pick_date)
	strat_well_section.Row_effective_date = formatDateString(strat_well_section.Row_effective_date)
	strat_well_section.Row_expiry_date = formatDateString(strat_well_section.Row_expiry_date)






	rows, err := stmt.Exec(strat_well_section.Uwi, strat_well_section.Strat_name_set_id, strat_well_section.Strat_unit_id, strat_well_section.Interp_id, strat_well_section.Active_ind, strat_well_section.Area_id, strat_well_section.Area_type, strat_well_section.Azimuth_north_type, strat_well_section.Certified_ind, strat_well_section.Conformity_relationship, strat_well_section.Dip_angle, strat_well_section.Dip_direction, strat_well_section.Dominant_lithology, strat_well_section.Effective_date, strat_well_section.Expiry_date, strat_well_section.Fault_heave, strat_well_section.Fault_throw, strat_well_section.Interpreter, strat_well_section.Missing_section, strat_well_section.Missing_strat_type, strat_well_section.Ordinal_seq_no, strat_well_section.Overturned_ind, strat_well_section.Pick_date, strat_well_section.Pick_depth, strat_well_section.Pick_depth_ouom, strat_well_section.Pick_location, strat_well_section.Pick_qualifier, strat_well_section.Pick_qualif_reason, strat_well_section.Pick_quality, strat_well_section.Pick_tvd, strat_well_section.Pick_version_type, strat_well_section.Ppdm_guid, strat_well_section.Preferred_pick_ind, strat_well_section.Remark, strat_well_section.Repeat_section, strat_well_section.Repeat_strat_occur_no, strat_well_section.Repeat_strat_type, strat_well_section.Source, strat_well_section.Source_document_id, strat_well_section.Strat_interpret_method, strat_well_section.Strike, strat_well_section.Sw_application_id, strat_well_section.Tvd_method, strat_well_section.Version_obs_no, strat_well_section.Row_changed_by, strat_well_section.Row_changed_date, strat_well_section.Row_created_by, strat_well_section.Row_created_date, strat_well_section.Row_effective_date, strat_well_section.Row_expiry_date, strat_well_section.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateStratWellSection(c *fiber.Ctx) error {
	var strat_well_section dto.Strat_well_section

	setDefaults(&strat_well_section)

	if err := json.Unmarshal(c.Body(), &strat_well_section); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	strat_well_section.Uwi = id

    if strat_well_section.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update strat_well_section set strat_name_set_id = :1, strat_unit_id = :2, interp_id = :3, active_ind = :4, area_id = :5, area_type = :6, azimuth_north_type = :7, certified_ind = :8, conformity_relationship = :9, dip_angle = :10, dip_direction = :11, dominant_lithology = :12, effective_date = :13, expiry_date = :14, fault_heave = :15, fault_throw = :16, interpreter = :17, missing_section = :18, missing_strat_type = :19, ordinal_seq_no = :20, overturned_ind = :21, pick_date = :22, pick_depth = :23, pick_depth_ouom = :24, pick_location = :25, pick_qualifier = :26, pick_qualif_reason = :27, pick_quality = :28, pick_tvd = :29, pick_version_type = :30, ppdm_guid = :31, preferred_pick_ind = :32, remark = :33, repeat_section = :34, repeat_strat_occur_no = :35, repeat_strat_type = :36, source = :37, source_document_id = :38, strat_interpret_method = :39, strike = :40, sw_application_id = :41, tvd_method = :42, version_obs_no = :43, row_changed_by = :44, row_changed_date = :45, row_created_by = :46, row_effective_date = :47, row_expiry_date = :48, row_quality = :49 where uwi = :51")
	if err != nil {
		return err
	}

	strat_well_section.Row_changed_date = formatDate(strat_well_section.Row_changed_date)
	strat_well_section.Effective_date = formatDateString(strat_well_section.Effective_date)
	strat_well_section.Expiry_date = formatDateString(strat_well_section.Expiry_date)
	strat_well_section.Pick_date = formatDateString(strat_well_section.Pick_date)
	strat_well_section.Row_effective_date = formatDateString(strat_well_section.Row_effective_date)
	strat_well_section.Row_expiry_date = formatDateString(strat_well_section.Row_expiry_date)






	rows, err := stmt.Exec(strat_well_section.Strat_name_set_id, strat_well_section.Strat_unit_id, strat_well_section.Interp_id, strat_well_section.Active_ind, strat_well_section.Area_id, strat_well_section.Area_type, strat_well_section.Azimuth_north_type, strat_well_section.Certified_ind, strat_well_section.Conformity_relationship, strat_well_section.Dip_angle, strat_well_section.Dip_direction, strat_well_section.Dominant_lithology, strat_well_section.Effective_date, strat_well_section.Expiry_date, strat_well_section.Fault_heave, strat_well_section.Fault_throw, strat_well_section.Interpreter, strat_well_section.Missing_section, strat_well_section.Missing_strat_type, strat_well_section.Ordinal_seq_no, strat_well_section.Overturned_ind, strat_well_section.Pick_date, strat_well_section.Pick_depth, strat_well_section.Pick_depth_ouom, strat_well_section.Pick_location, strat_well_section.Pick_qualifier, strat_well_section.Pick_qualif_reason, strat_well_section.Pick_quality, strat_well_section.Pick_tvd, strat_well_section.Pick_version_type, strat_well_section.Ppdm_guid, strat_well_section.Preferred_pick_ind, strat_well_section.Remark, strat_well_section.Repeat_section, strat_well_section.Repeat_strat_occur_no, strat_well_section.Repeat_strat_type, strat_well_section.Source, strat_well_section.Source_document_id, strat_well_section.Strat_interpret_method, strat_well_section.Strike, strat_well_section.Sw_application_id, strat_well_section.Tvd_method, strat_well_section.Version_obs_no, strat_well_section.Row_changed_by, strat_well_section.Row_changed_date, strat_well_section.Row_created_by, strat_well_section.Row_effective_date, strat_well_section.Row_expiry_date, strat_well_section.Row_quality, strat_well_section.Uwi)
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

func PatchStratWellSection(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update strat_well_section set "
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

func DeleteStratWellSection(c *fiber.Ctx) error {
	id := c.Params("id")
	var strat_well_section dto.Strat_well_section
	strat_well_section.Uwi = id

	stmt, err := db.Prepare("delete from strat_well_section where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(strat_well_section.Uwi)
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



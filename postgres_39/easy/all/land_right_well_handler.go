package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLandRightWell(c *fiber.Ctx) error {
	rows, err := db.Query("select * from land_right_well")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Land_right_well

	for rows.Next() {
		var land_right_well dto.Land_right_well
		if err := rows.Scan(&land_right_well.Uwi, &land_right_well.Land_right_subtype, &land_right_well.Land_right_id, &land_right_well.Lr_well_seq_no, &land_right_well.Active_ind, &land_right_well.Completion_obs_no, &land_right_well.Effective_date, &land_right_well.Expiry_date, &land_right_well.Gas_percent_psu, &land_right_well.Key_well_ind, &land_right_well.Oil_percent_psu, &land_right_well.Ppdm_guid, &land_right_well.Pr_str_form_obs_no, &land_right_well.Pr_str_source, &land_right_well.Remark, &land_right_well.Source, &land_right_well.Spacing_complete_ind, &land_right_well.Spacing_unit_id, &land_right_well.String_id, &land_right_well.Well_in_tract_ind, &land_right_well.Well_relationship_type, &land_right_well.Row_changed_by, &land_right_well.Row_changed_date, &land_right_well.Row_created_by, &land_right_well.Row_created_date, &land_right_well.Row_effective_date, &land_right_well.Row_expiry_date, &land_right_well.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append land_right_well to result
		result = append(result, land_right_well)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLandRightWell(c *fiber.Ctx) error {
	var land_right_well dto.Land_right_well

	setDefaults(&land_right_well)

	if err := json.Unmarshal(c.Body(), &land_right_well); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into land_right_well values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28)")
	if err != nil {
		return err
	}
	land_right_well.Row_created_date = formatDate(land_right_well.Row_created_date)
	land_right_well.Row_changed_date = formatDate(land_right_well.Row_changed_date)
	land_right_well.Effective_date = formatDateString(land_right_well.Effective_date)
	land_right_well.Expiry_date = formatDateString(land_right_well.Expiry_date)
	land_right_well.Row_effective_date = formatDateString(land_right_well.Row_effective_date)
	land_right_well.Row_expiry_date = formatDateString(land_right_well.Row_expiry_date)






	rows, err := stmt.Exec(land_right_well.Uwi, land_right_well.Land_right_subtype, land_right_well.Land_right_id, land_right_well.Lr_well_seq_no, land_right_well.Active_ind, land_right_well.Completion_obs_no, land_right_well.Effective_date, land_right_well.Expiry_date, land_right_well.Gas_percent_psu, land_right_well.Key_well_ind, land_right_well.Oil_percent_psu, land_right_well.Ppdm_guid, land_right_well.Pr_str_form_obs_no, land_right_well.Pr_str_source, land_right_well.Remark, land_right_well.Source, land_right_well.Spacing_complete_ind, land_right_well.Spacing_unit_id, land_right_well.String_id, land_right_well.Well_in_tract_ind, land_right_well.Well_relationship_type, land_right_well.Row_changed_by, land_right_well.Row_changed_date, land_right_well.Row_created_by, land_right_well.Row_created_date, land_right_well.Row_effective_date, land_right_well.Row_expiry_date, land_right_well.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLandRightWell(c *fiber.Ctx) error {
	var land_right_well dto.Land_right_well

	setDefaults(&land_right_well)

	if err := json.Unmarshal(c.Body(), &land_right_well); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	land_right_well.Uwi = id

    if land_right_well.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update land_right_well set land_right_subtype = :1, land_right_id = :2, lr_well_seq_no = :3, active_ind = :4, completion_obs_no = :5, effective_date = :6, expiry_date = :7, gas_percent_psu = :8, key_well_ind = :9, oil_percent_psu = :10, ppdm_guid = :11, pr_str_form_obs_no = :12, pr_str_source = :13, remark = :14, source = :15, spacing_complete_ind = :16, spacing_unit_id = :17, string_id = :18, well_in_tract_ind = :19, well_relationship_type = :20, row_changed_by = :21, row_changed_date = :22, row_created_by = :23, row_effective_date = :24, row_expiry_date = :25, row_quality = :26 where uwi = :28")
	if err != nil {
		return err
	}

	land_right_well.Row_changed_date = formatDate(land_right_well.Row_changed_date)
	land_right_well.Effective_date = formatDateString(land_right_well.Effective_date)
	land_right_well.Expiry_date = formatDateString(land_right_well.Expiry_date)
	land_right_well.Row_effective_date = formatDateString(land_right_well.Row_effective_date)
	land_right_well.Row_expiry_date = formatDateString(land_right_well.Row_expiry_date)






	rows, err := stmt.Exec(land_right_well.Land_right_subtype, land_right_well.Land_right_id, land_right_well.Lr_well_seq_no, land_right_well.Active_ind, land_right_well.Completion_obs_no, land_right_well.Effective_date, land_right_well.Expiry_date, land_right_well.Gas_percent_psu, land_right_well.Key_well_ind, land_right_well.Oil_percent_psu, land_right_well.Ppdm_guid, land_right_well.Pr_str_form_obs_no, land_right_well.Pr_str_source, land_right_well.Remark, land_right_well.Source, land_right_well.Spacing_complete_ind, land_right_well.Spacing_unit_id, land_right_well.String_id, land_right_well.Well_in_tract_ind, land_right_well.Well_relationship_type, land_right_well.Row_changed_by, land_right_well.Row_changed_date, land_right_well.Row_created_by, land_right_well.Row_effective_date, land_right_well.Row_expiry_date, land_right_well.Row_quality, land_right_well.Uwi)
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

func PatchLandRightWell(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update land_right_well set "
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

func DeleteLandRightWell(c *fiber.Ctx) error {
	id := c.Params("id")
	var land_right_well dto.Land_right_well
	land_right_well.Uwi = id

	stmt, err := db.Prepare("delete from land_right_well where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(land_right_well.Uwi)
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



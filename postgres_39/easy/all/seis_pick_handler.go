package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisPick(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_pick")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_pick

	for rows.Next() {
		var seis_pick dto.Seis_pick
		if err := rows.Scan(&seis_pick.Seis_set_subtype, &seis_pick.Interp_set_id, &seis_pick.Surface_id, &seis_pick.Seis_pick_id, &seis_pick.Active_ind, &seis_pick.Bin_grid_id, &seis_pick.Bin_point_id, &seis_pick.Bin_seis_set_id, &seis_pick.Bin_seis_set_subtype, &seis_pick.Bin_source, &seis_pick.Date_format_desc, &seis_pick.Effective_date, &seis_pick.Expiry_date, &seis_pick.Interp_type, &seis_pick.Pick_depth, &seis_pick.Pick_depth_ouom, &seis_pick.Pick_description, &seis_pick.Pick_method, &seis_pick.Pick_qualifier, &seis_pick.Pick_qualif_reason, &seis_pick.Pick_quality, &seis_pick.Pick_version_type, &seis_pick.Ppdm_guid, &seis_pick.Preferred_pick_ind, &seis_pick.Remark, &seis_pick.Seis_pick_value, &seis_pick.Seis_pick_value_ouom, &seis_pick.Seis_pick_value_uom, &seis_pick.Seis_point_id, &seis_pick.Seis_point_set_id, &seis_pick.Seis_point_set_subtype, &seis_pick.Source, &seis_pick.Trace_id, &seis_pick.Row_changed_by, &seis_pick.Row_changed_date, &seis_pick.Row_created_by, &seis_pick.Row_created_date, &seis_pick.Row_effective_date, &seis_pick.Row_expiry_date, &seis_pick.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_pick to result
		result = append(result, seis_pick)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisPick(c *fiber.Ctx) error {
	var seis_pick dto.Seis_pick

	setDefaults(&seis_pick)

	if err := json.Unmarshal(c.Body(), &seis_pick); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_pick values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40)")
	if err != nil {
		return err
	}
	seis_pick.Row_created_date = formatDate(seis_pick.Row_created_date)
	seis_pick.Row_changed_date = formatDate(seis_pick.Row_changed_date)
	seis_pick.Date_format_desc = formatDateString(seis_pick.Date_format_desc)
	seis_pick.Effective_date = formatDateString(seis_pick.Effective_date)
	seis_pick.Expiry_date = formatDateString(seis_pick.Expiry_date)
	seis_pick.Row_effective_date = formatDateString(seis_pick.Row_effective_date)
	seis_pick.Row_expiry_date = formatDateString(seis_pick.Row_expiry_date)






	rows, err := stmt.Exec(seis_pick.Seis_set_subtype, seis_pick.Interp_set_id, seis_pick.Surface_id, seis_pick.Seis_pick_id, seis_pick.Active_ind, seis_pick.Bin_grid_id, seis_pick.Bin_point_id, seis_pick.Bin_seis_set_id, seis_pick.Bin_seis_set_subtype, seis_pick.Bin_source, seis_pick.Date_format_desc, seis_pick.Effective_date, seis_pick.Expiry_date, seis_pick.Interp_type, seis_pick.Pick_depth, seis_pick.Pick_depth_ouom, seis_pick.Pick_description, seis_pick.Pick_method, seis_pick.Pick_qualifier, seis_pick.Pick_qualif_reason, seis_pick.Pick_quality, seis_pick.Pick_version_type, seis_pick.Ppdm_guid, seis_pick.Preferred_pick_ind, seis_pick.Remark, seis_pick.Seis_pick_value, seis_pick.Seis_pick_value_ouom, seis_pick.Seis_pick_value_uom, seis_pick.Seis_point_id, seis_pick.Seis_point_set_id, seis_pick.Seis_point_set_subtype, seis_pick.Source, seis_pick.Trace_id, seis_pick.Row_changed_by, seis_pick.Row_changed_date, seis_pick.Row_created_by, seis_pick.Row_created_date, seis_pick.Row_effective_date, seis_pick.Row_expiry_date, seis_pick.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisPick(c *fiber.Ctx) error {
	var seis_pick dto.Seis_pick

	setDefaults(&seis_pick)

	if err := json.Unmarshal(c.Body(), &seis_pick); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_pick.Seis_set_subtype = id

    if seis_pick.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_pick set interp_set_id = :1, surface_id = :2, seis_pick_id = :3, active_ind = :4, bin_grid_id = :5, bin_point_id = :6, bin_seis_set_id = :7, bin_seis_set_subtype = :8, bin_source = :9, date_format_desc = :10, effective_date = :11, expiry_date = :12, interp_type = :13, pick_depth = :14, pick_depth_ouom = :15, pick_description = :16, pick_method = :17, pick_qualifier = :18, pick_qualif_reason = :19, pick_quality = :20, pick_version_type = :21, ppdm_guid = :22, preferred_pick_ind = :23, remark = :24, seis_pick_value = :25, seis_pick_value_ouom = :26, seis_pick_value_uom = :27, seis_point_id = :28, seis_point_set_id = :29, seis_point_set_subtype = :30, source = :31, trace_id = :32, row_changed_by = :33, row_changed_date = :34, row_created_by = :35, row_effective_date = :36, row_expiry_date = :37, row_quality = :38 where seis_set_subtype = :40")
	if err != nil {
		return err
	}

	seis_pick.Row_changed_date = formatDate(seis_pick.Row_changed_date)
	seis_pick.Date_format_desc = formatDateString(seis_pick.Date_format_desc)
	seis_pick.Effective_date = formatDateString(seis_pick.Effective_date)
	seis_pick.Expiry_date = formatDateString(seis_pick.Expiry_date)
	seis_pick.Row_effective_date = formatDateString(seis_pick.Row_effective_date)
	seis_pick.Row_expiry_date = formatDateString(seis_pick.Row_expiry_date)






	rows, err := stmt.Exec(seis_pick.Interp_set_id, seis_pick.Surface_id, seis_pick.Seis_pick_id, seis_pick.Active_ind, seis_pick.Bin_grid_id, seis_pick.Bin_point_id, seis_pick.Bin_seis_set_id, seis_pick.Bin_seis_set_subtype, seis_pick.Bin_source, seis_pick.Date_format_desc, seis_pick.Effective_date, seis_pick.Expiry_date, seis_pick.Interp_type, seis_pick.Pick_depth, seis_pick.Pick_depth_ouom, seis_pick.Pick_description, seis_pick.Pick_method, seis_pick.Pick_qualifier, seis_pick.Pick_qualif_reason, seis_pick.Pick_quality, seis_pick.Pick_version_type, seis_pick.Ppdm_guid, seis_pick.Preferred_pick_ind, seis_pick.Remark, seis_pick.Seis_pick_value, seis_pick.Seis_pick_value_ouom, seis_pick.Seis_pick_value_uom, seis_pick.Seis_point_id, seis_pick.Seis_point_set_id, seis_pick.Seis_point_set_subtype, seis_pick.Source, seis_pick.Trace_id, seis_pick.Row_changed_by, seis_pick.Row_changed_date, seis_pick.Row_created_by, seis_pick.Row_effective_date, seis_pick.Row_expiry_date, seis_pick.Row_quality, seis_pick.Seis_set_subtype)
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

func PatchSeisPick(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_pick set "
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
		} else if key == "date_format_desc" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where seis_set_subtype = :id"

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

func DeleteSeisPick(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_pick dto.Seis_pick
	seis_pick.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_pick where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_pick.Seis_set_subtype)
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



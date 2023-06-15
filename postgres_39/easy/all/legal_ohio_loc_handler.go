package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLegalOhioLoc(c *fiber.Ctx) error {
	rows, err := db.Query("select * from legal_ohio_loc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Legal_ohio_loc

	for rows.Next() {
		var legal_ohio_loc dto.Legal_ohio_loc
		if err := rows.Scan(&legal_ohio_loc.Legal_ohio_loc_id, &legal_ohio_loc.Location_type, &legal_ohio_loc.Source, &legal_ohio_loc.Active_ind, &legal_ohio_loc.Area_id, &legal_ohio_loc.Area_type, &legal_ohio_loc.Coord_system_id, &legal_ohio_loc.Effective_date, &legal_ohio_loc.Ew_footage, &legal_ohio_loc.Ew_footage_ouom, &legal_ohio_loc.Ew_start_line, &legal_ohio_loc.Expiry_date, &legal_ohio_loc.Footage_origin, &legal_ohio_loc.Map_quad_min, &legal_ohio_loc.Map_quad_name, &legal_ohio_loc.Ns_footage, &legal_ohio_loc.Ns_footage_ouom, &legal_ohio_loc.Ns_start_line, &legal_ohio_loc.Ohio_allotment, &legal_ohio_loc.Ohio_division, &legal_ohio_loc.Ohio_fraction, &legal_ohio_loc.Ohio_land_subdivision_name, &legal_ohio_loc.Ohio_other_subdivision, &legal_ohio_loc.Ohio_quarter_township, &legal_ohio_loc.Ohio_range, &legal_ohio_loc.Ohio_range_dir, &legal_ohio_loc.Ohio_township, &legal_ohio_loc.Ohio_township_dir, &legal_ohio_loc.Ohio_township_name, &legal_ohio_loc.Ohio_tract, &legal_ohio_loc.Ohio_twp_lot_code, &legal_ohio_loc.Ohio_twp_section_code, &legal_ohio_loc.Ppdm_guid, &legal_ohio_loc.Principal_meridian, &legal_ohio_loc.Remark, &legal_ohio_loc.Scaled_footage_ind, &legal_ohio_loc.Spot_code, &legal_ohio_loc.Uwi, &legal_ohio_loc.Well_node_id, &legal_ohio_loc.Row_changed_by, &legal_ohio_loc.Row_changed_date, &legal_ohio_loc.Row_created_by, &legal_ohio_loc.Row_created_date, &legal_ohio_loc.Row_effective_date, &legal_ohio_loc.Row_expiry_date, &legal_ohio_loc.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append legal_ohio_loc to result
		result = append(result, legal_ohio_loc)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLegalOhioLoc(c *fiber.Ctx) error {
	var legal_ohio_loc dto.Legal_ohio_loc

	setDefaults(&legal_ohio_loc)

	if err := json.Unmarshal(c.Body(), &legal_ohio_loc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into legal_ohio_loc values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46)")
	if err != nil {
		return err
	}
	legal_ohio_loc.Row_created_date = formatDate(legal_ohio_loc.Row_created_date)
	legal_ohio_loc.Row_changed_date = formatDate(legal_ohio_loc.Row_changed_date)
	legal_ohio_loc.Effective_date = formatDateString(legal_ohio_loc.Effective_date)
	legal_ohio_loc.Expiry_date = formatDateString(legal_ohio_loc.Expiry_date)
	legal_ohio_loc.Row_effective_date = formatDateString(legal_ohio_loc.Row_effective_date)
	legal_ohio_loc.Row_expiry_date = formatDateString(legal_ohio_loc.Row_expiry_date)






	rows, err := stmt.Exec(legal_ohio_loc.Legal_ohio_loc_id, legal_ohio_loc.Location_type, legal_ohio_loc.Source, legal_ohio_loc.Active_ind, legal_ohio_loc.Area_id, legal_ohio_loc.Area_type, legal_ohio_loc.Coord_system_id, legal_ohio_loc.Effective_date, legal_ohio_loc.Ew_footage, legal_ohio_loc.Ew_footage_ouom, legal_ohio_loc.Ew_start_line, legal_ohio_loc.Expiry_date, legal_ohio_loc.Footage_origin, legal_ohio_loc.Map_quad_min, legal_ohio_loc.Map_quad_name, legal_ohio_loc.Ns_footage, legal_ohio_loc.Ns_footage_ouom, legal_ohio_loc.Ns_start_line, legal_ohio_loc.Ohio_allotment, legal_ohio_loc.Ohio_division, legal_ohio_loc.Ohio_fraction, legal_ohio_loc.Ohio_land_subdivision_name, legal_ohio_loc.Ohio_other_subdivision, legal_ohio_loc.Ohio_quarter_township, legal_ohio_loc.Ohio_range, legal_ohio_loc.Ohio_range_dir, legal_ohio_loc.Ohio_township, legal_ohio_loc.Ohio_township_dir, legal_ohio_loc.Ohio_township_name, legal_ohio_loc.Ohio_tract, legal_ohio_loc.Ohio_twp_lot_code, legal_ohio_loc.Ohio_twp_section_code, legal_ohio_loc.Ppdm_guid, legal_ohio_loc.Principal_meridian, legal_ohio_loc.Remark, legal_ohio_loc.Scaled_footage_ind, legal_ohio_loc.Spot_code, legal_ohio_loc.Uwi, legal_ohio_loc.Well_node_id, legal_ohio_loc.Row_changed_by, legal_ohio_loc.Row_changed_date, legal_ohio_loc.Row_created_by, legal_ohio_loc.Row_created_date, legal_ohio_loc.Row_effective_date, legal_ohio_loc.Row_expiry_date, legal_ohio_loc.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLegalOhioLoc(c *fiber.Ctx) error {
	var legal_ohio_loc dto.Legal_ohio_loc

	setDefaults(&legal_ohio_loc)

	if err := json.Unmarshal(c.Body(), &legal_ohio_loc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	legal_ohio_loc.Legal_ohio_loc_id = id

    if legal_ohio_loc.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update legal_ohio_loc set location_type = :1, source = :2, active_ind = :3, area_id = :4, area_type = :5, coord_system_id = :6, effective_date = :7, ew_footage = :8, ew_footage_ouom = :9, ew_start_line = :10, expiry_date = :11, footage_origin = :12, map_quad_min = :13, map_quad_name = :14, ns_footage = :15, ns_footage_ouom = :16, ns_start_line = :17, ohio_allotment = :18, ohio_division = :19, ohio_fraction = :20, ohio_land_subdivision_name = :21, ohio_other_subdivision = :22, ohio_quarter_township = :23, ohio_range = :24, ohio_range_dir = :25, ohio_township = :26, ohio_township_dir = :27, ohio_township_name = :28, ohio_tract = :29, ohio_twp_lot_code = :30, ohio_twp_section_code = :31, ppdm_guid = :32, principal_meridian = :33, remark = :34, scaled_footage_ind = :35, spot_code = :36, uwi = :37, well_node_id = :38, row_changed_by = :39, row_changed_date = :40, row_created_by = :41, row_effective_date = :42, row_expiry_date = :43, row_quality = :44 where legal_ohio_loc_id = :46")
	if err != nil {
		return err
	}

	legal_ohio_loc.Row_changed_date = formatDate(legal_ohio_loc.Row_changed_date)
	legal_ohio_loc.Effective_date = formatDateString(legal_ohio_loc.Effective_date)
	legal_ohio_loc.Expiry_date = formatDateString(legal_ohio_loc.Expiry_date)
	legal_ohio_loc.Row_effective_date = formatDateString(legal_ohio_loc.Row_effective_date)
	legal_ohio_loc.Row_expiry_date = formatDateString(legal_ohio_loc.Row_expiry_date)






	rows, err := stmt.Exec(legal_ohio_loc.Location_type, legal_ohio_loc.Source, legal_ohio_loc.Active_ind, legal_ohio_loc.Area_id, legal_ohio_loc.Area_type, legal_ohio_loc.Coord_system_id, legal_ohio_loc.Effective_date, legal_ohio_loc.Ew_footage, legal_ohio_loc.Ew_footage_ouom, legal_ohio_loc.Ew_start_line, legal_ohio_loc.Expiry_date, legal_ohio_loc.Footage_origin, legal_ohio_loc.Map_quad_min, legal_ohio_loc.Map_quad_name, legal_ohio_loc.Ns_footage, legal_ohio_loc.Ns_footage_ouom, legal_ohio_loc.Ns_start_line, legal_ohio_loc.Ohio_allotment, legal_ohio_loc.Ohio_division, legal_ohio_loc.Ohio_fraction, legal_ohio_loc.Ohio_land_subdivision_name, legal_ohio_loc.Ohio_other_subdivision, legal_ohio_loc.Ohio_quarter_township, legal_ohio_loc.Ohio_range, legal_ohio_loc.Ohio_range_dir, legal_ohio_loc.Ohio_township, legal_ohio_loc.Ohio_township_dir, legal_ohio_loc.Ohio_township_name, legal_ohio_loc.Ohio_tract, legal_ohio_loc.Ohio_twp_lot_code, legal_ohio_loc.Ohio_twp_section_code, legal_ohio_loc.Ppdm_guid, legal_ohio_loc.Principal_meridian, legal_ohio_loc.Remark, legal_ohio_loc.Scaled_footage_ind, legal_ohio_loc.Spot_code, legal_ohio_loc.Uwi, legal_ohio_loc.Well_node_id, legal_ohio_loc.Row_changed_by, legal_ohio_loc.Row_changed_date, legal_ohio_loc.Row_created_by, legal_ohio_loc.Row_effective_date, legal_ohio_loc.Row_expiry_date, legal_ohio_loc.Row_quality, legal_ohio_loc.Legal_ohio_loc_id)
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

func PatchLegalOhioLoc(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update legal_ohio_loc set "
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
	query += " where legal_ohio_loc_id = :id"

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

func DeleteLegalOhioLoc(c *fiber.Ctx) error {
	id := c.Params("id")
	var legal_ohio_loc dto.Legal_ohio_loc
	legal_ohio_loc.Legal_ohio_loc_id = id

	stmt, err := db.Prepare("delete from legal_ohio_loc where legal_ohio_loc_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(legal_ohio_loc.Legal_ohio_loc_id)
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



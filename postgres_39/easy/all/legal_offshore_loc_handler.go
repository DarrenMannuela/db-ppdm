package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLegalOffshoreLoc(c *fiber.Ctx) error {
	rows, err := db.Query("select * from legal_offshore_loc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Legal_offshore_loc

	for rows.Next() {
		var legal_offshore_loc dto.Legal_offshore_loc
		if err := rows.Scan(&legal_offshore_loc.Legal_offshore_loc_id, &legal_offshore_loc.Location_type, &legal_offshore_loc.Source, &legal_offshore_loc.Active_ind, &legal_offshore_loc.Area_id, &legal_offshore_loc.Area_type, &legal_offshore_loc.Block_addition, &legal_offshore_loc.Block_modifier, &legal_offshore_loc.Coord_system_id, &legal_offshore_loc.Effective_date, &legal_offshore_loc.Ew_footage, &legal_offshore_loc.Ew_footage_ouom, &legal_offshore_loc.Ew_start_line, &legal_offshore_loc.Expiry_date, &legal_offshore_loc.Footage_origin, &legal_offshore_loc.Governing_agency_ba_id, &legal_offshore_loc.Grid_block_ew, &legal_offshore_loc.Grid_block_ns, &legal_offshore_loc.Grid_block_range, &legal_offshore_loc.Grid_block_tier, &legal_offshore_loc.Ns_footage, &legal_offshore_loc.Ns_footage_ouom, &legal_offshore_loc.Ns_start_line, &legal_offshore_loc.Ocs_num, &legal_offshore_loc.Offshore_area_code, &legal_offshore_loc.Offshore_block, &legal_offshore_loc.Ppdm_guid, &legal_offshore_loc.Projected_ew_direction, &legal_offshore_loc.Projected_meridian, &legal_offshore_loc.Projected_ns_direction, &legal_offshore_loc.Projected_range, &legal_offshore_loc.Projected_section, &legal_offshore_loc.Projected_township, &legal_offshore_loc.Remark, &legal_offshore_loc.Scaled_footage_ind, &legal_offshore_loc.Utm_quadrant, &legal_offshore_loc.Uwi, &legal_offshore_loc.Water_bottom_zone, &legal_offshore_loc.Well_node_id, &legal_offshore_loc.Row_changed_by, &legal_offshore_loc.Row_changed_date, &legal_offshore_loc.Row_created_by, &legal_offshore_loc.Row_created_date, &legal_offshore_loc.Row_effective_date, &legal_offshore_loc.Row_expiry_date, &legal_offshore_loc.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append legal_offshore_loc to result
		result = append(result, legal_offshore_loc)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLegalOffshoreLoc(c *fiber.Ctx) error {
	var legal_offshore_loc dto.Legal_offshore_loc

	setDefaults(&legal_offshore_loc)

	if err := json.Unmarshal(c.Body(), &legal_offshore_loc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into legal_offshore_loc values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46)")
	if err != nil {
		return err
	}
	legal_offshore_loc.Row_created_date = formatDate(legal_offshore_loc.Row_created_date)
	legal_offshore_loc.Row_changed_date = formatDate(legal_offshore_loc.Row_changed_date)
	legal_offshore_loc.Effective_date = formatDateString(legal_offshore_loc.Effective_date)
	legal_offshore_loc.Expiry_date = formatDateString(legal_offshore_loc.Expiry_date)
	legal_offshore_loc.Row_effective_date = formatDateString(legal_offshore_loc.Row_effective_date)
	legal_offshore_loc.Row_expiry_date = formatDateString(legal_offshore_loc.Row_expiry_date)






	rows, err := stmt.Exec(legal_offshore_loc.Legal_offshore_loc_id, legal_offshore_loc.Location_type, legal_offshore_loc.Source, legal_offshore_loc.Active_ind, legal_offshore_loc.Area_id, legal_offshore_loc.Area_type, legal_offshore_loc.Block_addition, legal_offshore_loc.Block_modifier, legal_offshore_loc.Coord_system_id, legal_offshore_loc.Effective_date, legal_offshore_loc.Ew_footage, legal_offshore_loc.Ew_footage_ouom, legal_offshore_loc.Ew_start_line, legal_offshore_loc.Expiry_date, legal_offshore_loc.Footage_origin, legal_offshore_loc.Governing_agency_ba_id, legal_offshore_loc.Grid_block_ew, legal_offshore_loc.Grid_block_ns, legal_offshore_loc.Grid_block_range, legal_offshore_loc.Grid_block_tier, legal_offshore_loc.Ns_footage, legal_offshore_loc.Ns_footage_ouom, legal_offshore_loc.Ns_start_line, legal_offshore_loc.Ocs_num, legal_offshore_loc.Offshore_area_code, legal_offshore_loc.Offshore_block, legal_offshore_loc.Ppdm_guid, legal_offshore_loc.Projected_ew_direction, legal_offshore_loc.Projected_meridian, legal_offshore_loc.Projected_ns_direction, legal_offshore_loc.Projected_range, legal_offshore_loc.Projected_section, legal_offshore_loc.Projected_township, legal_offshore_loc.Remark, legal_offshore_loc.Scaled_footage_ind, legal_offshore_loc.Utm_quadrant, legal_offshore_loc.Uwi, legal_offshore_loc.Water_bottom_zone, legal_offshore_loc.Well_node_id, legal_offshore_loc.Row_changed_by, legal_offshore_loc.Row_changed_date, legal_offshore_loc.Row_created_by, legal_offshore_loc.Row_created_date, legal_offshore_loc.Row_effective_date, legal_offshore_loc.Row_expiry_date, legal_offshore_loc.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLegalOffshoreLoc(c *fiber.Ctx) error {
	var legal_offshore_loc dto.Legal_offshore_loc

	setDefaults(&legal_offshore_loc)

	if err := json.Unmarshal(c.Body(), &legal_offshore_loc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	legal_offshore_loc.Legal_offshore_loc_id = id

    if legal_offshore_loc.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update legal_offshore_loc set location_type = :1, source = :2, active_ind = :3, area_id = :4, area_type = :5, block_addition = :6, block_modifier = :7, coord_system_id = :8, effective_date = :9, ew_footage = :10, ew_footage_ouom = :11, ew_start_line = :12, expiry_date = :13, footage_origin = :14, governing_agency_ba_id = :15, grid_block_ew = :16, grid_block_ns = :17, grid_block_range = :18, grid_block_tier = :19, ns_footage = :20, ns_footage_ouom = :21, ns_start_line = :22, ocs_num = :23, offshore_area_code = :24, offshore_block = :25, ppdm_guid = :26, projected_ew_direction = :27, projected_meridian = :28, projected_ns_direction = :29, projected_range = :30, projected_section = :31, projected_township = :32, remark = :33, scaled_footage_ind = :34, utm_quadrant = :35, uwi = :36, water_bottom_zone = :37, well_node_id = :38, row_changed_by = :39, row_changed_date = :40, row_created_by = :41, row_effective_date = :42, row_expiry_date = :43, row_quality = :44 where legal_offshore_loc_id = :46")
	if err != nil {
		return err
	}

	legal_offshore_loc.Row_changed_date = formatDate(legal_offshore_loc.Row_changed_date)
	legal_offshore_loc.Effective_date = formatDateString(legal_offshore_loc.Effective_date)
	legal_offshore_loc.Expiry_date = formatDateString(legal_offshore_loc.Expiry_date)
	legal_offshore_loc.Row_effective_date = formatDateString(legal_offshore_loc.Row_effective_date)
	legal_offshore_loc.Row_expiry_date = formatDateString(legal_offshore_loc.Row_expiry_date)






	rows, err := stmt.Exec(legal_offshore_loc.Location_type, legal_offshore_loc.Source, legal_offshore_loc.Active_ind, legal_offshore_loc.Area_id, legal_offshore_loc.Area_type, legal_offshore_loc.Block_addition, legal_offshore_loc.Block_modifier, legal_offshore_loc.Coord_system_id, legal_offshore_loc.Effective_date, legal_offshore_loc.Ew_footage, legal_offshore_loc.Ew_footage_ouom, legal_offshore_loc.Ew_start_line, legal_offshore_loc.Expiry_date, legal_offshore_loc.Footage_origin, legal_offshore_loc.Governing_agency_ba_id, legal_offshore_loc.Grid_block_ew, legal_offshore_loc.Grid_block_ns, legal_offshore_loc.Grid_block_range, legal_offshore_loc.Grid_block_tier, legal_offshore_loc.Ns_footage, legal_offshore_loc.Ns_footage_ouom, legal_offshore_loc.Ns_start_line, legal_offshore_loc.Ocs_num, legal_offshore_loc.Offshore_area_code, legal_offshore_loc.Offshore_block, legal_offshore_loc.Ppdm_guid, legal_offshore_loc.Projected_ew_direction, legal_offshore_loc.Projected_meridian, legal_offshore_loc.Projected_ns_direction, legal_offshore_loc.Projected_range, legal_offshore_loc.Projected_section, legal_offshore_loc.Projected_township, legal_offshore_loc.Remark, legal_offshore_loc.Scaled_footage_ind, legal_offshore_loc.Utm_quadrant, legal_offshore_loc.Uwi, legal_offshore_loc.Water_bottom_zone, legal_offshore_loc.Well_node_id, legal_offshore_loc.Row_changed_by, legal_offshore_loc.Row_changed_date, legal_offshore_loc.Row_created_by, legal_offshore_loc.Row_effective_date, legal_offshore_loc.Row_expiry_date, legal_offshore_loc.Row_quality, legal_offshore_loc.Legal_offshore_loc_id)
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

func PatchLegalOffshoreLoc(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update legal_offshore_loc set "
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
	query += " where legal_offshore_loc_id = :id"

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

func DeleteLegalOffshoreLoc(c *fiber.Ctx) error {
	id := c.Params("id")
	var legal_offshore_loc dto.Legal_offshore_loc
	legal_offshore_loc.Legal_offshore_loc_id = id

	stmt, err := db.Prepare("delete from legal_offshore_loc where legal_offshore_loc_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(legal_offshore_loc.Legal_offshore_loc_id)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLegalTexasLoc(c *fiber.Ctx) error {
	rows, err := db.Query("select * from legal_texas_loc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Legal_texas_loc

	for rows.Next() {
		var legal_texas_loc dto.Legal_texas_loc
		if err := rows.Scan(&legal_texas_loc.Legal_texas_loc_id, &legal_texas_loc.Location_type, &legal_texas_loc.Source, &legal_texas_loc.Abstract_num, &legal_texas_loc.Active_ind, &legal_texas_loc.Area_id, &legal_texas_loc.Area_type, &legal_texas_loc.Block_fraction, &legal_texas_loc.Coord_system_id, &legal_texas_loc.Effective_date, &legal_texas_loc.Ew_direction, &legal_texas_loc.Ew_footage, &legal_texas_loc.Ew_footage_ouom, &legal_texas_loc.Ew_start_line, &legal_texas_loc.Expiry_date, &legal_texas_loc.Footage_origin, &legal_texas_loc.Labor, &legal_texas_loc.League, &legal_texas_loc.Ns_direction, &legal_texas_loc.Ns_footage, &legal_texas_loc.Ns_footage_ouom, &legal_texas_loc.Ns_start_line, &legal_texas_loc.Porcion_num, &legal_texas_loc.Porcion_survey_name, &legal_texas_loc.Ppdm_guid, &legal_texas_loc.Remark, &legal_texas_loc.Scaled_footage_ind, &legal_texas_loc.Section_fraction, &legal_texas_loc.Spot_code, &legal_texas_loc.Texas_block, &legal_texas_loc.Texas_lot, &legal_texas_loc.Texas_section, &legal_texas_loc.Texas_share, &legal_texas_loc.Texas_subdivision, &legal_texas_loc.Texas_survey, &legal_texas_loc.Texas_township, &legal_texas_loc.Uwi, &legal_texas_loc.Well_node_id, &legal_texas_loc.Row_changed_by, &legal_texas_loc.Row_changed_date, &legal_texas_loc.Row_created_by, &legal_texas_loc.Row_created_date, &legal_texas_loc.Row_effective_date, &legal_texas_loc.Row_expiry_date, &legal_texas_loc.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append legal_texas_loc to result
		result = append(result, legal_texas_loc)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLegalTexasLoc(c *fiber.Ctx) error {
	var legal_texas_loc dto.Legal_texas_loc

	setDefaults(&legal_texas_loc)

	if err := json.Unmarshal(c.Body(), &legal_texas_loc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into legal_texas_loc values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45)")
	if err != nil {
		return err
	}
	legal_texas_loc.Row_created_date = formatDate(legal_texas_loc.Row_created_date)
	legal_texas_loc.Row_changed_date = formatDate(legal_texas_loc.Row_changed_date)
	legal_texas_loc.Effective_date = formatDateString(legal_texas_loc.Effective_date)
	legal_texas_loc.Expiry_date = formatDateString(legal_texas_loc.Expiry_date)
	legal_texas_loc.Row_effective_date = formatDateString(legal_texas_loc.Row_effective_date)
	legal_texas_loc.Row_expiry_date = formatDateString(legal_texas_loc.Row_expiry_date)






	rows, err := stmt.Exec(legal_texas_loc.Legal_texas_loc_id, legal_texas_loc.Location_type, legal_texas_loc.Source, legal_texas_loc.Abstract_num, legal_texas_loc.Active_ind, legal_texas_loc.Area_id, legal_texas_loc.Area_type, legal_texas_loc.Block_fraction, legal_texas_loc.Coord_system_id, legal_texas_loc.Effective_date, legal_texas_loc.Ew_direction, legal_texas_loc.Ew_footage, legal_texas_loc.Ew_footage_ouom, legal_texas_loc.Ew_start_line, legal_texas_loc.Expiry_date, legal_texas_loc.Footage_origin, legal_texas_loc.Labor, legal_texas_loc.League, legal_texas_loc.Ns_direction, legal_texas_loc.Ns_footage, legal_texas_loc.Ns_footage_ouom, legal_texas_loc.Ns_start_line, legal_texas_loc.Porcion_num, legal_texas_loc.Porcion_survey_name, legal_texas_loc.Ppdm_guid, legal_texas_loc.Remark, legal_texas_loc.Scaled_footage_ind, legal_texas_loc.Section_fraction, legal_texas_loc.Spot_code, legal_texas_loc.Texas_block, legal_texas_loc.Texas_lot, legal_texas_loc.Texas_section, legal_texas_loc.Texas_share, legal_texas_loc.Texas_subdivision, legal_texas_loc.Texas_survey, legal_texas_loc.Texas_township, legal_texas_loc.Uwi, legal_texas_loc.Well_node_id, legal_texas_loc.Row_changed_by, legal_texas_loc.Row_changed_date, legal_texas_loc.Row_created_by, legal_texas_loc.Row_created_date, legal_texas_loc.Row_effective_date, legal_texas_loc.Row_expiry_date, legal_texas_loc.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLegalTexasLoc(c *fiber.Ctx) error {
	var legal_texas_loc dto.Legal_texas_loc

	setDefaults(&legal_texas_loc)

	if err := json.Unmarshal(c.Body(), &legal_texas_loc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	legal_texas_loc.Legal_texas_loc_id = id

    if legal_texas_loc.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update legal_texas_loc set location_type = :1, source = :2, abstract_num = :3, active_ind = :4, area_id = :5, area_type = :6, block_fraction = :7, coord_system_id = :8, effective_date = :9, ew_direction = :10, ew_footage = :11, ew_footage_ouom = :12, ew_start_line = :13, expiry_date = :14, footage_origin = :15, labor = :16, league = :17, ns_direction = :18, ns_footage = :19, ns_footage_ouom = :20, ns_start_line = :21, porcion_num = :22, porcion_survey_name = :23, ppdm_guid = :24, remark = :25, scaled_footage_ind = :26, section_fraction = :27, spot_code = :28, texas_block = :29, texas_lot = :30, texas_section = :31, texas_share = :32, texas_subdivision = :33, texas_survey = :34, texas_township = :35, uwi = :36, well_node_id = :37, row_changed_by = :38, row_changed_date = :39, row_created_by = :40, row_effective_date = :41, row_expiry_date = :42, row_quality = :43 where legal_texas_loc_id = :45")
	if err != nil {
		return err
	}

	legal_texas_loc.Row_changed_date = formatDate(legal_texas_loc.Row_changed_date)
	legal_texas_loc.Effective_date = formatDateString(legal_texas_loc.Effective_date)
	legal_texas_loc.Expiry_date = formatDateString(legal_texas_loc.Expiry_date)
	legal_texas_loc.Row_effective_date = formatDateString(legal_texas_loc.Row_effective_date)
	legal_texas_loc.Row_expiry_date = formatDateString(legal_texas_loc.Row_expiry_date)






	rows, err := stmt.Exec(legal_texas_loc.Location_type, legal_texas_loc.Source, legal_texas_loc.Abstract_num, legal_texas_loc.Active_ind, legal_texas_loc.Area_id, legal_texas_loc.Area_type, legal_texas_loc.Block_fraction, legal_texas_loc.Coord_system_id, legal_texas_loc.Effective_date, legal_texas_loc.Ew_direction, legal_texas_loc.Ew_footage, legal_texas_loc.Ew_footage_ouom, legal_texas_loc.Ew_start_line, legal_texas_loc.Expiry_date, legal_texas_loc.Footage_origin, legal_texas_loc.Labor, legal_texas_loc.League, legal_texas_loc.Ns_direction, legal_texas_loc.Ns_footage, legal_texas_loc.Ns_footage_ouom, legal_texas_loc.Ns_start_line, legal_texas_loc.Porcion_num, legal_texas_loc.Porcion_survey_name, legal_texas_loc.Ppdm_guid, legal_texas_loc.Remark, legal_texas_loc.Scaled_footage_ind, legal_texas_loc.Section_fraction, legal_texas_loc.Spot_code, legal_texas_loc.Texas_block, legal_texas_loc.Texas_lot, legal_texas_loc.Texas_section, legal_texas_loc.Texas_share, legal_texas_loc.Texas_subdivision, legal_texas_loc.Texas_survey, legal_texas_loc.Texas_township, legal_texas_loc.Uwi, legal_texas_loc.Well_node_id, legal_texas_loc.Row_changed_by, legal_texas_loc.Row_changed_date, legal_texas_loc.Row_created_by, legal_texas_loc.Row_effective_date, legal_texas_loc.Row_expiry_date, legal_texas_loc.Row_quality, legal_texas_loc.Legal_texas_loc_id)
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

func PatchLegalTexasLoc(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update legal_texas_loc set "
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
	query += " where legal_texas_loc_id = :id"

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

func DeleteLegalTexasLoc(c *fiber.Ctx) error {
	id := c.Params("id")
	var legal_texas_loc dto.Legal_texas_loc
	legal_texas_loc.Legal_texas_loc_id = id

	stmt, err := db.Prepare("delete from legal_texas_loc where legal_texas_loc_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(legal_texas_loc.Legal_texas_loc_id)
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



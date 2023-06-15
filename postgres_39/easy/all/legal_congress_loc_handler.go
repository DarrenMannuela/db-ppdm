package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLegalCongressLoc(c *fiber.Ctx) error {
	rows, err := db.Query("select * from legal_congress_loc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Legal_congress_loc

	for rows.Next() {
		var legal_congress_loc dto.Legal_congress_loc
		if err := rows.Scan(&legal_congress_loc.Legal_congress_loc_id, &legal_congress_loc.Location_type, &legal_congress_loc.Source, &legal_congress_loc.Active_ind, &legal_congress_loc.Area_id, &legal_congress_loc.Area_type, &legal_congress_loc.Congress_lot_num, &legal_congress_loc.Congress_range, &legal_congress_loc.Congress_section, &legal_congress_loc.Congress_section_suffix, &legal_congress_loc.Congress_township, &legal_congress_loc.Congress_twp_name, &legal_congress_loc.Coord_system_id, &legal_congress_loc.Effective_date, &legal_congress_loc.Ew_direction, &legal_congress_loc.Ew_footage, &legal_congress_loc.Ew_footage_ouom, &legal_congress_loc.Ew_start_line, &legal_congress_loc.Expiry_date, &legal_congress_loc.Footage_origin, &legal_congress_loc.Ns_direction, &legal_congress_loc.Ns_footage, &legal_congress_loc.Ns_footage_ouom, &legal_congress_loc.Ns_start_line, &legal_congress_loc.Ppdm_guid, &legal_congress_loc.Principal_meridian, &legal_congress_loc.Remark, &legal_congress_loc.Scaled_footage_ind, &legal_congress_loc.Section_type, &legal_congress_loc.Spot_code, &legal_congress_loc.Uwi, &legal_congress_loc.Well_node_id, &legal_congress_loc.Row_changed_by, &legal_congress_loc.Row_changed_date, &legal_congress_loc.Row_created_by, &legal_congress_loc.Row_created_date, &legal_congress_loc.Row_effective_date, &legal_congress_loc.Row_expiry_date, &legal_congress_loc.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append legal_congress_loc to result
		result = append(result, legal_congress_loc)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLegalCongressLoc(c *fiber.Ctx) error {
	var legal_congress_loc dto.Legal_congress_loc

	setDefaults(&legal_congress_loc)

	if err := json.Unmarshal(c.Body(), &legal_congress_loc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into legal_congress_loc values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39)")
	if err != nil {
		return err
	}
	legal_congress_loc.Row_created_date = formatDate(legal_congress_loc.Row_created_date)
	legal_congress_loc.Row_changed_date = formatDate(legal_congress_loc.Row_changed_date)
	legal_congress_loc.Effective_date = formatDateString(legal_congress_loc.Effective_date)
	legal_congress_loc.Expiry_date = formatDateString(legal_congress_loc.Expiry_date)
	legal_congress_loc.Row_effective_date = formatDateString(legal_congress_loc.Row_effective_date)
	legal_congress_loc.Row_expiry_date = formatDateString(legal_congress_loc.Row_expiry_date)






	rows, err := stmt.Exec(legal_congress_loc.Legal_congress_loc_id, legal_congress_loc.Location_type, legal_congress_loc.Source, legal_congress_loc.Active_ind, legal_congress_loc.Area_id, legal_congress_loc.Area_type, legal_congress_loc.Congress_lot_num, legal_congress_loc.Congress_range, legal_congress_loc.Congress_section, legal_congress_loc.Congress_section_suffix, legal_congress_loc.Congress_township, legal_congress_loc.Congress_twp_name, legal_congress_loc.Coord_system_id, legal_congress_loc.Effective_date, legal_congress_loc.Ew_direction, legal_congress_loc.Ew_footage, legal_congress_loc.Ew_footage_ouom, legal_congress_loc.Ew_start_line, legal_congress_loc.Expiry_date, legal_congress_loc.Footage_origin, legal_congress_loc.Ns_direction, legal_congress_loc.Ns_footage, legal_congress_loc.Ns_footage_ouom, legal_congress_loc.Ns_start_line, legal_congress_loc.Ppdm_guid, legal_congress_loc.Principal_meridian, legal_congress_loc.Remark, legal_congress_loc.Scaled_footage_ind, legal_congress_loc.Section_type, legal_congress_loc.Spot_code, legal_congress_loc.Uwi, legal_congress_loc.Well_node_id, legal_congress_loc.Row_changed_by, legal_congress_loc.Row_changed_date, legal_congress_loc.Row_created_by, legal_congress_loc.Row_created_date, legal_congress_loc.Row_effective_date, legal_congress_loc.Row_expiry_date, legal_congress_loc.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLegalCongressLoc(c *fiber.Ctx) error {
	var legal_congress_loc dto.Legal_congress_loc

	setDefaults(&legal_congress_loc)

	if err := json.Unmarshal(c.Body(), &legal_congress_loc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	legal_congress_loc.Legal_congress_loc_id = id

    if legal_congress_loc.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update legal_congress_loc set location_type = :1, source = :2, active_ind = :3, area_id = :4, area_type = :5, congress_lot_num = :6, congress_range = :7, congress_section = :8, congress_section_suffix = :9, congress_township = :10, congress_twp_name = :11, coord_system_id = :12, effective_date = :13, ew_direction = :14, ew_footage = :15, ew_footage_ouom = :16, ew_start_line = :17, expiry_date = :18, footage_origin = :19, ns_direction = :20, ns_footage = :21, ns_footage_ouom = :22, ns_start_line = :23, ppdm_guid = :24, principal_meridian = :25, remark = :26, scaled_footage_ind = :27, section_type = :28, spot_code = :29, uwi = :30, well_node_id = :31, row_changed_by = :32, row_changed_date = :33, row_created_by = :34, row_effective_date = :35, row_expiry_date = :36, row_quality = :37 where legal_congress_loc_id = :39")
	if err != nil {
		return err
	}

	legal_congress_loc.Row_changed_date = formatDate(legal_congress_loc.Row_changed_date)
	legal_congress_loc.Effective_date = formatDateString(legal_congress_loc.Effective_date)
	legal_congress_loc.Expiry_date = formatDateString(legal_congress_loc.Expiry_date)
	legal_congress_loc.Row_effective_date = formatDateString(legal_congress_loc.Row_effective_date)
	legal_congress_loc.Row_expiry_date = formatDateString(legal_congress_loc.Row_expiry_date)






	rows, err := stmt.Exec(legal_congress_loc.Location_type, legal_congress_loc.Source, legal_congress_loc.Active_ind, legal_congress_loc.Area_id, legal_congress_loc.Area_type, legal_congress_loc.Congress_lot_num, legal_congress_loc.Congress_range, legal_congress_loc.Congress_section, legal_congress_loc.Congress_section_suffix, legal_congress_loc.Congress_township, legal_congress_loc.Congress_twp_name, legal_congress_loc.Coord_system_id, legal_congress_loc.Effective_date, legal_congress_loc.Ew_direction, legal_congress_loc.Ew_footage, legal_congress_loc.Ew_footage_ouom, legal_congress_loc.Ew_start_line, legal_congress_loc.Expiry_date, legal_congress_loc.Footage_origin, legal_congress_loc.Ns_direction, legal_congress_loc.Ns_footage, legal_congress_loc.Ns_footage_ouom, legal_congress_loc.Ns_start_line, legal_congress_loc.Ppdm_guid, legal_congress_loc.Principal_meridian, legal_congress_loc.Remark, legal_congress_loc.Scaled_footage_ind, legal_congress_loc.Section_type, legal_congress_loc.Spot_code, legal_congress_loc.Uwi, legal_congress_loc.Well_node_id, legal_congress_loc.Row_changed_by, legal_congress_loc.Row_changed_date, legal_congress_loc.Row_created_by, legal_congress_loc.Row_effective_date, legal_congress_loc.Row_expiry_date, legal_congress_loc.Row_quality, legal_congress_loc.Legal_congress_loc_id)
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

func PatchLegalCongressLoc(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update legal_congress_loc set "
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
	query += " where legal_congress_loc_id = :id"

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

func DeleteLegalCongressLoc(c *fiber.Ctx) error {
	id := c.Params("id")
	var legal_congress_loc dto.Legal_congress_loc
	legal_congress_loc.Legal_congress_loc_id = id

	stmt, err := db.Prepare("delete from legal_congress_loc where legal_congress_loc_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(legal_congress_loc.Legal_congress_loc_id)
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



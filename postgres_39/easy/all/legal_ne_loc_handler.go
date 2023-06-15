package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLegalNeLoc(c *fiber.Ctx) error {
	rows, err := db.Query("select * from legal_ne_loc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Legal_ne_loc

	for rows.Next() {
		var legal_ne_loc dto.Legal_ne_loc
		if err := rows.Scan(&legal_ne_loc.Legal_ne_loc_id, &legal_ne_loc.Location_type, &legal_ne_loc.Source, &legal_ne_loc.Active_ind, &legal_ne_loc.Area_id, &legal_ne_loc.Area_type, &legal_ne_loc.Coord_system_id, &legal_ne_loc.Effective_date, &legal_ne_loc.Ew_footage, &legal_ne_loc.Ew_footage_ouom, &legal_ne_loc.Ew_start_line, &legal_ne_loc.Expiry_date, &legal_ne_loc.Footage_origin, &legal_ne_loc.Ne_district, &legal_ne_loc.Ne_lot_code, &legal_ne_loc.Ne_map_quad_min, &legal_ne_loc.Ne_map_quad_name, &legal_ne_loc.Ne_map_quad_section, &legal_ne_loc.Ne_township_name, &legal_ne_loc.Ns_footage, &legal_ne_loc.Ns_footage_ouom, &legal_ne_loc.Ns_start_line, &legal_ne_loc.Ppdm_guid, &legal_ne_loc.Reference_latitude, &legal_ne_loc.Reference_longitude, &legal_ne_loc.Remark, &legal_ne_loc.Scaled_footage_ind, &legal_ne_loc.Uwi, &legal_ne_loc.Well_node_id, &legal_ne_loc.Row_changed_by, &legal_ne_loc.Row_changed_date, &legal_ne_loc.Row_created_by, &legal_ne_loc.Row_created_date, &legal_ne_loc.Row_effective_date, &legal_ne_loc.Row_expiry_date, &legal_ne_loc.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append legal_ne_loc to result
		result = append(result, legal_ne_loc)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLegalNeLoc(c *fiber.Ctx) error {
	var legal_ne_loc dto.Legal_ne_loc

	setDefaults(&legal_ne_loc)

	if err := json.Unmarshal(c.Body(), &legal_ne_loc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into legal_ne_loc values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36)")
	if err != nil {
		return err
	}
	legal_ne_loc.Row_created_date = formatDate(legal_ne_loc.Row_created_date)
	legal_ne_loc.Row_changed_date = formatDate(legal_ne_loc.Row_changed_date)
	legal_ne_loc.Effective_date = formatDateString(legal_ne_loc.Effective_date)
	legal_ne_loc.Expiry_date = formatDateString(legal_ne_loc.Expiry_date)
	legal_ne_loc.Row_effective_date = formatDateString(legal_ne_loc.Row_effective_date)
	legal_ne_loc.Row_expiry_date = formatDateString(legal_ne_loc.Row_expiry_date)






	rows, err := stmt.Exec(legal_ne_loc.Legal_ne_loc_id, legal_ne_loc.Location_type, legal_ne_loc.Source, legal_ne_loc.Active_ind, legal_ne_loc.Area_id, legal_ne_loc.Area_type, legal_ne_loc.Coord_system_id, legal_ne_loc.Effective_date, legal_ne_loc.Ew_footage, legal_ne_loc.Ew_footage_ouom, legal_ne_loc.Ew_start_line, legal_ne_loc.Expiry_date, legal_ne_loc.Footage_origin, legal_ne_loc.Ne_district, legal_ne_loc.Ne_lot_code, legal_ne_loc.Ne_map_quad_min, legal_ne_loc.Ne_map_quad_name, legal_ne_loc.Ne_map_quad_section, legal_ne_loc.Ne_township_name, legal_ne_loc.Ns_footage, legal_ne_loc.Ns_footage_ouom, legal_ne_loc.Ns_start_line, legal_ne_loc.Ppdm_guid, legal_ne_loc.Reference_latitude, legal_ne_loc.Reference_longitude, legal_ne_loc.Remark, legal_ne_loc.Scaled_footage_ind, legal_ne_loc.Uwi, legal_ne_loc.Well_node_id, legal_ne_loc.Row_changed_by, legal_ne_loc.Row_changed_date, legal_ne_loc.Row_created_by, legal_ne_loc.Row_created_date, legal_ne_loc.Row_effective_date, legal_ne_loc.Row_expiry_date, legal_ne_loc.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLegalNeLoc(c *fiber.Ctx) error {
	var legal_ne_loc dto.Legal_ne_loc

	setDefaults(&legal_ne_loc)

	if err := json.Unmarshal(c.Body(), &legal_ne_loc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	legal_ne_loc.Legal_ne_loc_id = id

    if legal_ne_loc.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update legal_ne_loc set location_type = :1, source = :2, active_ind = :3, area_id = :4, area_type = :5, coord_system_id = :6, effective_date = :7, ew_footage = :8, ew_footage_ouom = :9, ew_start_line = :10, expiry_date = :11, footage_origin = :12, ne_district = :13, ne_lot_code = :14, ne_map_quad_min = :15, ne_map_quad_name = :16, ne_map_quad_section = :17, ne_township_name = :18, ns_footage = :19, ns_footage_ouom = :20, ns_start_line = :21, ppdm_guid = :22, reference_latitude = :23, reference_longitude = :24, remark = :25, scaled_footage_ind = :26, uwi = :27, well_node_id = :28, row_changed_by = :29, row_changed_date = :30, row_created_by = :31, row_effective_date = :32, row_expiry_date = :33, row_quality = :34 where legal_ne_loc_id = :36")
	if err != nil {
		return err
	}

	legal_ne_loc.Row_changed_date = formatDate(legal_ne_loc.Row_changed_date)
	legal_ne_loc.Effective_date = formatDateString(legal_ne_loc.Effective_date)
	legal_ne_loc.Expiry_date = formatDateString(legal_ne_loc.Expiry_date)
	legal_ne_loc.Row_effective_date = formatDateString(legal_ne_loc.Row_effective_date)
	legal_ne_loc.Row_expiry_date = formatDateString(legal_ne_loc.Row_expiry_date)






	rows, err := stmt.Exec(legal_ne_loc.Location_type, legal_ne_loc.Source, legal_ne_loc.Active_ind, legal_ne_loc.Area_id, legal_ne_loc.Area_type, legal_ne_loc.Coord_system_id, legal_ne_loc.Effective_date, legal_ne_loc.Ew_footage, legal_ne_loc.Ew_footage_ouom, legal_ne_loc.Ew_start_line, legal_ne_loc.Expiry_date, legal_ne_loc.Footage_origin, legal_ne_loc.Ne_district, legal_ne_loc.Ne_lot_code, legal_ne_loc.Ne_map_quad_min, legal_ne_loc.Ne_map_quad_name, legal_ne_loc.Ne_map_quad_section, legal_ne_loc.Ne_township_name, legal_ne_loc.Ns_footage, legal_ne_loc.Ns_footage_ouom, legal_ne_loc.Ns_start_line, legal_ne_loc.Ppdm_guid, legal_ne_loc.Reference_latitude, legal_ne_loc.Reference_longitude, legal_ne_loc.Remark, legal_ne_loc.Scaled_footage_ind, legal_ne_loc.Uwi, legal_ne_loc.Well_node_id, legal_ne_loc.Row_changed_by, legal_ne_loc.Row_changed_date, legal_ne_loc.Row_created_by, legal_ne_loc.Row_effective_date, legal_ne_loc.Row_expiry_date, legal_ne_loc.Row_quality, legal_ne_loc.Legal_ne_loc_id)
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

func PatchLegalNeLoc(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update legal_ne_loc set "
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
	query += " where legal_ne_loc_id = :id"

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

func DeleteLegalNeLoc(c *fiber.Ctx) error {
	id := c.Params("id")
	var legal_ne_loc dto.Legal_ne_loc
	legal_ne_loc.Legal_ne_loc_id = id

	stmt, err := db.Prepare("delete from legal_ne_loc where legal_ne_loc_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(legal_ne_loc.Legal_ne_loc_id)
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



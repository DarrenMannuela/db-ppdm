package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLegalCarterLoc(c *fiber.Ctx) error {
	rows, err := db.Query("select * from legal_carter_loc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Legal_carter_loc

	for rows.Next() {
		var legal_carter_loc dto.Legal_carter_loc
		if err := rows.Scan(&legal_carter_loc.Legal_carter_loc_id, &legal_carter_loc.Location_type, &legal_carter_loc.Source, &legal_carter_loc.Active_ind, &legal_carter_loc.Area_id, &legal_carter_loc.Area_type, &legal_carter_loc.Carter_range, &legal_carter_loc.Carter_section, &legal_carter_loc.Carter_township, &legal_carter_loc.Coord_system_id, &legal_carter_loc.Effective_date, &legal_carter_loc.Ew_direction, &legal_carter_loc.Ew_footage, &legal_carter_loc.Ew_footage_ouom, &legal_carter_loc.Ew_start_line, &legal_carter_loc.Expiry_date, &legal_carter_loc.Footage_origin, &legal_carter_loc.Map_quad_min, &legal_carter_loc.Map_quad_name, &legal_carter_loc.Ns_direction, &legal_carter_loc.Ns_footage, &legal_carter_loc.Ns_footage_ouom, &legal_carter_loc.Ns_start_line, &legal_carter_loc.Ppdm_guid, &legal_carter_loc.Remark, &legal_carter_loc.Scaled_footage_ind, &legal_carter_loc.Spot_code, &legal_carter_loc.Uwi, &legal_carter_loc.Well_node_id, &legal_carter_loc.Row_changed_by, &legal_carter_loc.Row_changed_date, &legal_carter_loc.Row_created_by, &legal_carter_loc.Row_created_date, &legal_carter_loc.Row_effective_date, &legal_carter_loc.Row_expiry_date, &legal_carter_loc.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append legal_carter_loc to result
		result = append(result, legal_carter_loc)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLegalCarterLoc(c *fiber.Ctx) error {
	var legal_carter_loc dto.Legal_carter_loc

	setDefaults(&legal_carter_loc)

	if err := json.Unmarshal(c.Body(), &legal_carter_loc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into legal_carter_loc values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36)")
	if err != nil {
		return err
	}
	legal_carter_loc.Row_created_date = formatDate(legal_carter_loc.Row_created_date)
	legal_carter_loc.Row_changed_date = formatDate(legal_carter_loc.Row_changed_date)
	legal_carter_loc.Effective_date = formatDateString(legal_carter_loc.Effective_date)
	legal_carter_loc.Expiry_date = formatDateString(legal_carter_loc.Expiry_date)
	legal_carter_loc.Row_effective_date = formatDateString(legal_carter_loc.Row_effective_date)
	legal_carter_loc.Row_expiry_date = formatDateString(legal_carter_loc.Row_expiry_date)






	rows, err := stmt.Exec(legal_carter_loc.Legal_carter_loc_id, legal_carter_loc.Location_type, legal_carter_loc.Source, legal_carter_loc.Active_ind, legal_carter_loc.Area_id, legal_carter_loc.Area_type, legal_carter_loc.Carter_range, legal_carter_loc.Carter_section, legal_carter_loc.Carter_township, legal_carter_loc.Coord_system_id, legal_carter_loc.Effective_date, legal_carter_loc.Ew_direction, legal_carter_loc.Ew_footage, legal_carter_loc.Ew_footage_ouom, legal_carter_loc.Ew_start_line, legal_carter_loc.Expiry_date, legal_carter_loc.Footage_origin, legal_carter_loc.Map_quad_min, legal_carter_loc.Map_quad_name, legal_carter_loc.Ns_direction, legal_carter_loc.Ns_footage, legal_carter_loc.Ns_footage_ouom, legal_carter_loc.Ns_start_line, legal_carter_loc.Ppdm_guid, legal_carter_loc.Remark, legal_carter_loc.Scaled_footage_ind, legal_carter_loc.Spot_code, legal_carter_loc.Uwi, legal_carter_loc.Well_node_id, legal_carter_loc.Row_changed_by, legal_carter_loc.Row_changed_date, legal_carter_loc.Row_created_by, legal_carter_loc.Row_created_date, legal_carter_loc.Row_effective_date, legal_carter_loc.Row_expiry_date, legal_carter_loc.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLegalCarterLoc(c *fiber.Ctx) error {
	var legal_carter_loc dto.Legal_carter_loc

	setDefaults(&legal_carter_loc)

	if err := json.Unmarshal(c.Body(), &legal_carter_loc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	legal_carter_loc.Legal_carter_loc_id = id

    if legal_carter_loc.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update legal_carter_loc set location_type = :1, source = :2, active_ind = :3, area_id = :4, area_type = :5, carter_range = :6, carter_section = :7, carter_township = :8, coord_system_id = :9, effective_date = :10, ew_direction = :11, ew_footage = :12, ew_footage_ouom = :13, ew_start_line = :14, expiry_date = :15, footage_origin = :16, map_quad_min = :17, map_quad_name = :18, ns_direction = :19, ns_footage = :20, ns_footage_ouom = :21, ns_start_line = :22, ppdm_guid = :23, remark = :24, scaled_footage_ind = :25, spot_code = :26, uwi = :27, well_node_id = :28, row_changed_by = :29, row_changed_date = :30, row_created_by = :31, row_effective_date = :32, row_expiry_date = :33, row_quality = :34 where legal_carter_loc_id = :36")
	if err != nil {
		return err
	}

	legal_carter_loc.Row_changed_date = formatDate(legal_carter_loc.Row_changed_date)
	legal_carter_loc.Effective_date = formatDateString(legal_carter_loc.Effective_date)
	legal_carter_loc.Expiry_date = formatDateString(legal_carter_loc.Expiry_date)
	legal_carter_loc.Row_effective_date = formatDateString(legal_carter_loc.Row_effective_date)
	legal_carter_loc.Row_expiry_date = formatDateString(legal_carter_loc.Row_expiry_date)






	rows, err := stmt.Exec(legal_carter_loc.Location_type, legal_carter_loc.Source, legal_carter_loc.Active_ind, legal_carter_loc.Area_id, legal_carter_loc.Area_type, legal_carter_loc.Carter_range, legal_carter_loc.Carter_section, legal_carter_loc.Carter_township, legal_carter_loc.Coord_system_id, legal_carter_loc.Effective_date, legal_carter_loc.Ew_direction, legal_carter_loc.Ew_footage, legal_carter_loc.Ew_footage_ouom, legal_carter_loc.Ew_start_line, legal_carter_loc.Expiry_date, legal_carter_loc.Footage_origin, legal_carter_loc.Map_quad_min, legal_carter_loc.Map_quad_name, legal_carter_loc.Ns_direction, legal_carter_loc.Ns_footage, legal_carter_loc.Ns_footage_ouom, legal_carter_loc.Ns_start_line, legal_carter_loc.Ppdm_guid, legal_carter_loc.Remark, legal_carter_loc.Scaled_footage_ind, legal_carter_loc.Spot_code, legal_carter_loc.Uwi, legal_carter_loc.Well_node_id, legal_carter_loc.Row_changed_by, legal_carter_loc.Row_changed_date, legal_carter_loc.Row_created_by, legal_carter_loc.Row_effective_date, legal_carter_loc.Row_expiry_date, legal_carter_loc.Row_quality, legal_carter_loc.Legal_carter_loc_id)
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

func PatchLegalCarterLoc(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update legal_carter_loc set "
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
	query += " where legal_carter_loc_id = :id"

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

func DeleteLegalCarterLoc(c *fiber.Ctx) error {
	id := c.Params("id")
	var legal_carter_loc dto.Legal_carter_loc
	legal_carter_loc.Legal_carter_loc_id = id

	stmt, err := db.Prepare("delete from legal_carter_loc where legal_carter_loc_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(legal_carter_loc.Legal_carter_loc_id)
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



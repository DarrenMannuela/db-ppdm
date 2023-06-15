package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLegalDlsLoc(c *fiber.Ctx) error {
	rows, err := db.Query("select * from legal_dls_loc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Legal_dls_loc

	for rows.Next() {
		var legal_dls_loc dto.Legal_dls_loc
		if err := rows.Scan(&legal_dls_loc.Legal_dls_loc_id, &legal_dls_loc.Location_type, &legal_dls_loc.Source, &legal_dls_loc.Active_ind, &legal_dls_loc.Coord_system_id, &legal_dls_loc.Dls_event_sequence, &legal_dls_loc.Dls_legal_subdivision, &legal_dls_loc.Dls_loc_exception, &legal_dls_loc.Dls_meridian, &legal_dls_loc.Dls_meridian_direction, &legal_dls_loc.Dls_range, &legal_dls_loc.Dls_range_modifier, &legal_dls_loc.Dls_section, &legal_dls_loc.Dls_township, &legal_dls_loc.Dls_township_modifier, &legal_dls_loc.Effective_date, &legal_dls_loc.Expiry_date, &legal_dls_loc.Ppdm_guid, &legal_dls_loc.Remark, &legal_dls_loc.Uwi, &legal_dls_loc.Well_node_id, &legal_dls_loc.Row_changed_by, &legal_dls_loc.Row_changed_date, &legal_dls_loc.Row_created_by, &legal_dls_loc.Row_created_date, &legal_dls_loc.Row_effective_date, &legal_dls_loc.Row_expiry_date, &legal_dls_loc.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append legal_dls_loc to result
		result = append(result, legal_dls_loc)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLegalDlsLoc(c *fiber.Ctx) error {
	var legal_dls_loc dto.Legal_dls_loc

	setDefaults(&legal_dls_loc)

	if err := json.Unmarshal(c.Body(), &legal_dls_loc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into legal_dls_loc values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28)")
	if err != nil {
		return err
	}
	legal_dls_loc.Row_created_date = formatDate(legal_dls_loc.Row_created_date)
	legal_dls_loc.Row_changed_date = formatDate(legal_dls_loc.Row_changed_date)
	legal_dls_loc.Effective_date = formatDateString(legal_dls_loc.Effective_date)
	legal_dls_loc.Expiry_date = formatDateString(legal_dls_loc.Expiry_date)
	legal_dls_loc.Row_effective_date = formatDateString(legal_dls_loc.Row_effective_date)
	legal_dls_loc.Row_expiry_date = formatDateString(legal_dls_loc.Row_expiry_date)






	rows, err := stmt.Exec(legal_dls_loc.Legal_dls_loc_id, legal_dls_loc.Location_type, legal_dls_loc.Source, legal_dls_loc.Active_ind, legal_dls_loc.Coord_system_id, legal_dls_loc.Dls_event_sequence, legal_dls_loc.Dls_legal_subdivision, legal_dls_loc.Dls_loc_exception, legal_dls_loc.Dls_meridian, legal_dls_loc.Dls_meridian_direction, legal_dls_loc.Dls_range, legal_dls_loc.Dls_range_modifier, legal_dls_loc.Dls_section, legal_dls_loc.Dls_township, legal_dls_loc.Dls_township_modifier, legal_dls_loc.Effective_date, legal_dls_loc.Expiry_date, legal_dls_loc.Ppdm_guid, legal_dls_loc.Remark, legal_dls_loc.Uwi, legal_dls_loc.Well_node_id, legal_dls_loc.Row_changed_by, legal_dls_loc.Row_changed_date, legal_dls_loc.Row_created_by, legal_dls_loc.Row_created_date, legal_dls_loc.Row_effective_date, legal_dls_loc.Row_expiry_date, legal_dls_loc.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLegalDlsLoc(c *fiber.Ctx) error {
	var legal_dls_loc dto.Legal_dls_loc

	setDefaults(&legal_dls_loc)

	if err := json.Unmarshal(c.Body(), &legal_dls_loc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	legal_dls_loc.Legal_dls_loc_id = id

    if legal_dls_loc.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update legal_dls_loc set location_type = :1, source = :2, active_ind = :3, coord_system_id = :4, dls_event_sequence = :5, dls_legal_subdivision = :6, dls_loc_exception = :7, dls_meridian = :8, dls_meridian_direction = :9, dls_range = :10, dls_range_modifier = :11, dls_section = :12, dls_township = :13, dls_township_modifier = :14, effective_date = :15, expiry_date = :16, ppdm_guid = :17, remark = :18, uwi = :19, well_node_id = :20, row_changed_by = :21, row_changed_date = :22, row_created_by = :23, row_effective_date = :24, row_expiry_date = :25, row_quality = :26 where legal_dls_loc_id = :28")
	if err != nil {
		return err
	}

	legal_dls_loc.Row_changed_date = formatDate(legal_dls_loc.Row_changed_date)
	legal_dls_loc.Effective_date = formatDateString(legal_dls_loc.Effective_date)
	legal_dls_loc.Expiry_date = formatDateString(legal_dls_loc.Expiry_date)
	legal_dls_loc.Row_effective_date = formatDateString(legal_dls_loc.Row_effective_date)
	legal_dls_loc.Row_expiry_date = formatDateString(legal_dls_loc.Row_expiry_date)






	rows, err := stmt.Exec(legal_dls_loc.Location_type, legal_dls_loc.Source, legal_dls_loc.Active_ind, legal_dls_loc.Coord_system_id, legal_dls_loc.Dls_event_sequence, legal_dls_loc.Dls_legal_subdivision, legal_dls_loc.Dls_loc_exception, legal_dls_loc.Dls_meridian, legal_dls_loc.Dls_meridian_direction, legal_dls_loc.Dls_range, legal_dls_loc.Dls_range_modifier, legal_dls_loc.Dls_section, legal_dls_loc.Dls_township, legal_dls_loc.Dls_township_modifier, legal_dls_loc.Effective_date, legal_dls_loc.Expiry_date, legal_dls_loc.Ppdm_guid, legal_dls_loc.Remark, legal_dls_loc.Uwi, legal_dls_loc.Well_node_id, legal_dls_loc.Row_changed_by, legal_dls_loc.Row_changed_date, legal_dls_loc.Row_created_by, legal_dls_loc.Row_effective_date, legal_dls_loc.Row_expiry_date, legal_dls_loc.Row_quality, legal_dls_loc.Legal_dls_loc_id)
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

func PatchLegalDlsLoc(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update legal_dls_loc set "
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
	query += " where legal_dls_loc_id = :id"

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

func DeleteLegalDlsLoc(c *fiber.Ctx) error {
	id := c.Params("id")
	var legal_dls_loc dto.Legal_dls_loc
	legal_dls_loc.Legal_dls_loc_id = id

	stmt, err := db.Prepare("delete from legal_dls_loc where legal_dls_loc_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(legal_dls_loc.Legal_dls_loc_id)
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



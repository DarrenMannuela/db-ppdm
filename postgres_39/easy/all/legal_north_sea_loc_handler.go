package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLegalNorthSeaLoc(c *fiber.Ctx) error {
	rows, err := db.Query("select * from legal_north_sea_loc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Legal_north_sea_loc

	for rows.Next() {
		var legal_north_sea_loc dto.Legal_north_sea_loc
		if err := rows.Scan(&legal_north_sea_loc.Legal_north_sea_loc_id, &legal_north_sea_loc.Location_type, &legal_north_sea_loc.Source, &legal_north_sea_loc.Active_ind, &legal_north_sea_loc.Block_no, &legal_north_sea_loc.Block_suffix, &legal_north_sea_loc.Coord_system_id, &legal_north_sea_loc.Effective_date, &legal_north_sea_loc.Expiry_date, &legal_north_sea_loc.Land_well_ind, &legal_north_sea_loc.Ppdm_guid, &legal_north_sea_loc.Principal_meridian, &legal_north_sea_loc.Quadrant, &legal_north_sea_loc.Quadrant_prefix, &legal_north_sea_loc.Remark, &legal_north_sea_loc.Uwi, &legal_north_sea_loc.Well_node_id, &legal_north_sea_loc.Well_no_prefix, &legal_north_sea_loc.Well_no_suffix, &legal_north_sea_loc.Well_seq_no, &legal_north_sea_loc.Row_changed_by, &legal_north_sea_loc.Row_changed_date, &legal_north_sea_loc.Row_created_by, &legal_north_sea_loc.Row_created_date, &legal_north_sea_loc.Row_effective_date, &legal_north_sea_loc.Row_expiry_date, &legal_north_sea_loc.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append legal_north_sea_loc to result
		result = append(result, legal_north_sea_loc)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLegalNorthSeaLoc(c *fiber.Ctx) error {
	var legal_north_sea_loc dto.Legal_north_sea_loc

	setDefaults(&legal_north_sea_loc)

	if err := json.Unmarshal(c.Body(), &legal_north_sea_loc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into legal_north_sea_loc values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27)")
	if err != nil {
		return err
	}
	legal_north_sea_loc.Row_created_date = formatDate(legal_north_sea_loc.Row_created_date)
	legal_north_sea_loc.Row_changed_date = formatDate(legal_north_sea_loc.Row_changed_date)
	legal_north_sea_loc.Effective_date = formatDateString(legal_north_sea_loc.Effective_date)
	legal_north_sea_loc.Expiry_date = formatDateString(legal_north_sea_loc.Expiry_date)
	legal_north_sea_loc.Row_effective_date = formatDateString(legal_north_sea_loc.Row_effective_date)
	legal_north_sea_loc.Row_expiry_date = formatDateString(legal_north_sea_loc.Row_expiry_date)






	rows, err := stmt.Exec(legal_north_sea_loc.Legal_north_sea_loc_id, legal_north_sea_loc.Location_type, legal_north_sea_loc.Source, legal_north_sea_loc.Active_ind, legal_north_sea_loc.Block_no, legal_north_sea_loc.Block_suffix, legal_north_sea_loc.Coord_system_id, legal_north_sea_loc.Effective_date, legal_north_sea_loc.Expiry_date, legal_north_sea_loc.Land_well_ind, legal_north_sea_loc.Ppdm_guid, legal_north_sea_loc.Principal_meridian, legal_north_sea_loc.Quadrant, legal_north_sea_loc.Quadrant_prefix, legal_north_sea_loc.Remark, legal_north_sea_loc.Uwi, legal_north_sea_loc.Well_node_id, legal_north_sea_loc.Well_no_prefix, legal_north_sea_loc.Well_no_suffix, legal_north_sea_loc.Well_seq_no, legal_north_sea_loc.Row_changed_by, legal_north_sea_loc.Row_changed_date, legal_north_sea_loc.Row_created_by, legal_north_sea_loc.Row_created_date, legal_north_sea_loc.Row_effective_date, legal_north_sea_loc.Row_expiry_date, legal_north_sea_loc.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLegalNorthSeaLoc(c *fiber.Ctx) error {
	var legal_north_sea_loc dto.Legal_north_sea_loc

	setDefaults(&legal_north_sea_loc)

	if err := json.Unmarshal(c.Body(), &legal_north_sea_loc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	legal_north_sea_loc.Legal_north_sea_loc_id = id

    if legal_north_sea_loc.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update legal_north_sea_loc set location_type = :1, source = :2, active_ind = :3, block_no = :4, block_suffix = :5, coord_system_id = :6, effective_date = :7, expiry_date = :8, land_well_ind = :9, ppdm_guid = :10, principal_meridian = :11, quadrant = :12, quadrant_prefix = :13, remark = :14, uwi = :15, well_node_id = :16, well_no_prefix = :17, well_no_suffix = :18, well_seq_no = :19, row_changed_by = :20, row_changed_date = :21, row_created_by = :22, row_effective_date = :23, row_expiry_date = :24, row_quality = :25 where legal_north_sea_loc_id = :27")
	if err != nil {
		return err
	}

	legal_north_sea_loc.Row_changed_date = formatDate(legal_north_sea_loc.Row_changed_date)
	legal_north_sea_loc.Effective_date = formatDateString(legal_north_sea_loc.Effective_date)
	legal_north_sea_loc.Expiry_date = formatDateString(legal_north_sea_loc.Expiry_date)
	legal_north_sea_loc.Row_effective_date = formatDateString(legal_north_sea_loc.Row_effective_date)
	legal_north_sea_loc.Row_expiry_date = formatDateString(legal_north_sea_loc.Row_expiry_date)






	rows, err := stmt.Exec(legal_north_sea_loc.Location_type, legal_north_sea_loc.Source, legal_north_sea_loc.Active_ind, legal_north_sea_loc.Block_no, legal_north_sea_loc.Block_suffix, legal_north_sea_loc.Coord_system_id, legal_north_sea_loc.Effective_date, legal_north_sea_loc.Expiry_date, legal_north_sea_loc.Land_well_ind, legal_north_sea_loc.Ppdm_guid, legal_north_sea_loc.Principal_meridian, legal_north_sea_loc.Quadrant, legal_north_sea_loc.Quadrant_prefix, legal_north_sea_loc.Remark, legal_north_sea_loc.Uwi, legal_north_sea_loc.Well_node_id, legal_north_sea_loc.Well_no_prefix, legal_north_sea_loc.Well_no_suffix, legal_north_sea_loc.Well_seq_no, legal_north_sea_loc.Row_changed_by, legal_north_sea_loc.Row_changed_date, legal_north_sea_loc.Row_created_by, legal_north_sea_loc.Row_effective_date, legal_north_sea_loc.Row_expiry_date, legal_north_sea_loc.Row_quality, legal_north_sea_loc.Legal_north_sea_loc_id)
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

func PatchLegalNorthSeaLoc(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update legal_north_sea_loc set "
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
	query += " where legal_north_sea_loc_id = :id"

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

func DeleteLegalNorthSeaLoc(c *fiber.Ctx) error {
	id := c.Params("id")
	var legal_north_sea_loc dto.Legal_north_sea_loc
	legal_north_sea_loc.Legal_north_sea_loc_id = id

	stmt, err := db.Prepare("delete from legal_north_sea_loc where legal_north_sea_loc_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(legal_north_sea_loc.Legal_north_sea_loc_id)
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



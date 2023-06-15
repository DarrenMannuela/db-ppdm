package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLegalNtsLoc(c *fiber.Ctx) error {
	rows, err := db.Query("select * from legal_nts_loc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Legal_nts_loc

	for rows.Next() {
		var legal_nts_loc dto.Legal_nts_loc
		if err := rows.Scan(&legal_nts_loc.Legal_nts_loc_id, &legal_nts_loc.Location_type, &legal_nts_loc.Source, &legal_nts_loc.Active_ind, &legal_nts_loc.Block, &legal_nts_loc.Coord_system_id, &legal_nts_loc.Effective_date, &legal_nts_loc.Event_sequence, &legal_nts_loc.Expiry_date, &legal_nts_loc.Letter_quadrangle, &legal_nts_loc.Nts_loc_exception, &legal_nts_loc.Ppdm_guid, &legal_nts_loc.Primary_quadrangle, &legal_nts_loc.Quarter_unit, &legal_nts_loc.Remark, &legal_nts_loc.Sixteenth, &legal_nts_loc.Unit, &legal_nts_loc.Uwi, &legal_nts_loc.Well_node_id, &legal_nts_loc.Row_changed_by, &legal_nts_loc.Row_changed_date, &legal_nts_loc.Row_created_by, &legal_nts_loc.Row_created_date, &legal_nts_loc.Row_effective_date, &legal_nts_loc.Row_expiry_date, &legal_nts_loc.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append legal_nts_loc to result
		result = append(result, legal_nts_loc)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLegalNtsLoc(c *fiber.Ctx) error {
	var legal_nts_loc dto.Legal_nts_loc

	setDefaults(&legal_nts_loc)

	if err := json.Unmarshal(c.Body(), &legal_nts_loc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into legal_nts_loc values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26)")
	if err != nil {
		return err
	}
	legal_nts_loc.Row_created_date = formatDate(legal_nts_loc.Row_created_date)
	legal_nts_loc.Row_changed_date = formatDate(legal_nts_loc.Row_changed_date)
	legal_nts_loc.Effective_date = formatDateString(legal_nts_loc.Effective_date)
	legal_nts_loc.Expiry_date = formatDateString(legal_nts_loc.Expiry_date)
	legal_nts_loc.Row_effective_date = formatDateString(legal_nts_loc.Row_effective_date)
	legal_nts_loc.Row_expiry_date = formatDateString(legal_nts_loc.Row_expiry_date)






	rows, err := stmt.Exec(legal_nts_loc.Legal_nts_loc_id, legal_nts_loc.Location_type, legal_nts_loc.Source, legal_nts_loc.Active_ind, legal_nts_loc.Block, legal_nts_loc.Coord_system_id, legal_nts_loc.Effective_date, legal_nts_loc.Event_sequence, legal_nts_loc.Expiry_date, legal_nts_loc.Letter_quadrangle, legal_nts_loc.Nts_loc_exception, legal_nts_loc.Ppdm_guid, legal_nts_loc.Primary_quadrangle, legal_nts_loc.Quarter_unit, legal_nts_loc.Remark, legal_nts_loc.Sixteenth, legal_nts_loc.Unit, legal_nts_loc.Uwi, legal_nts_loc.Well_node_id, legal_nts_loc.Row_changed_by, legal_nts_loc.Row_changed_date, legal_nts_loc.Row_created_by, legal_nts_loc.Row_created_date, legal_nts_loc.Row_effective_date, legal_nts_loc.Row_expiry_date, legal_nts_loc.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLegalNtsLoc(c *fiber.Ctx) error {
	var legal_nts_loc dto.Legal_nts_loc

	setDefaults(&legal_nts_loc)

	if err := json.Unmarshal(c.Body(), &legal_nts_loc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	legal_nts_loc.Legal_nts_loc_id = id

    if legal_nts_loc.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update legal_nts_loc set location_type = :1, source = :2, active_ind = :3, block = :4, coord_system_id = :5, effective_date = :6, event_sequence = :7, expiry_date = :8, letter_quadrangle = :9, nts_loc_exception = :10, ppdm_guid = :11, primary_quadrangle = :12, quarter_unit = :13, remark = :14, sixteenth = :15, unit = :16, uwi = :17, well_node_id = :18, row_changed_by = :19, row_changed_date = :20, row_created_by = :21, row_effective_date = :22, row_expiry_date = :23, row_quality = :24 where legal_nts_loc_id = :26")
	if err != nil {
		return err
	}

	legal_nts_loc.Row_changed_date = formatDate(legal_nts_loc.Row_changed_date)
	legal_nts_loc.Effective_date = formatDateString(legal_nts_loc.Effective_date)
	legal_nts_loc.Expiry_date = formatDateString(legal_nts_loc.Expiry_date)
	legal_nts_loc.Row_effective_date = formatDateString(legal_nts_loc.Row_effective_date)
	legal_nts_loc.Row_expiry_date = formatDateString(legal_nts_loc.Row_expiry_date)






	rows, err := stmt.Exec(legal_nts_loc.Location_type, legal_nts_loc.Source, legal_nts_loc.Active_ind, legal_nts_loc.Block, legal_nts_loc.Coord_system_id, legal_nts_loc.Effective_date, legal_nts_loc.Event_sequence, legal_nts_loc.Expiry_date, legal_nts_loc.Letter_quadrangle, legal_nts_loc.Nts_loc_exception, legal_nts_loc.Ppdm_guid, legal_nts_loc.Primary_quadrangle, legal_nts_loc.Quarter_unit, legal_nts_loc.Remark, legal_nts_loc.Sixteenth, legal_nts_loc.Unit, legal_nts_loc.Uwi, legal_nts_loc.Well_node_id, legal_nts_loc.Row_changed_by, legal_nts_loc.Row_changed_date, legal_nts_loc.Row_created_by, legal_nts_loc.Row_effective_date, legal_nts_loc.Row_expiry_date, legal_nts_loc.Row_quality, legal_nts_loc.Legal_nts_loc_id)
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

func PatchLegalNtsLoc(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update legal_nts_loc set "
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
	query += " where legal_nts_loc_id = :id"

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

func DeleteLegalNtsLoc(c *fiber.Ctx) error {
	id := c.Params("id")
	var legal_nts_loc dto.Legal_nts_loc
	legal_nts_loc.Legal_nts_loc_id = id

	stmt, err := db.Prepare("delete from legal_nts_loc where legal_nts_loc_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(legal_nts_loc.Legal_nts_loc_id)
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



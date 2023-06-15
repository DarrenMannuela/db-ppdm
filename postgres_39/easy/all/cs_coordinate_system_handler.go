package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetCsCoordinateSystem(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cs_coordinate_system")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cs_coordinate_system

	for rows.Next() {
		var cs_coordinate_system dto.Cs_coordinate_system
		if err := rows.Scan(&cs_coordinate_system.Coord_system_id, &cs_coordinate_system.Active_ind, &cs_coordinate_system.Coordinate_system_type, &cs_coordinate_system.Coord_system_abbreviation, &cs_coordinate_system.Coord_system_area, &cs_coordinate_system.Coord_system_long_name, &cs_coordinate_system.Coord_system_short_name, &cs_coordinate_system.Coord_system_uom, &cs_coordinate_system.Datum_ouom, &cs_coordinate_system.Effective_date, &cs_coordinate_system.Expiry_date, &cs_coordinate_system.Geodetic_datum, &cs_coordinate_system.N_value, &cs_coordinate_system.N_value_ouom, &cs_coordinate_system.Owner_ba_id, &cs_coordinate_system.Parent_coord_system_id, &cs_coordinate_system.Perspective_height, &cs_coordinate_system.Ppdm_guid, &cs_coordinate_system.Prime_meridian, &cs_coordinate_system.Principal_meridian, &cs_coordinate_system.Projection_type, &cs_coordinate_system.Reference_elev, &cs_coordinate_system.Reference_elev_ouom, &cs_coordinate_system.Remark, &cs_coordinate_system.Rotation_ind, &cs_coordinate_system.Source, &cs_coordinate_system.Source_document_id, &cs_coordinate_system.Vertical_datum_type, &cs_coordinate_system.Row_changed_by, &cs_coordinate_system.Row_changed_date, &cs_coordinate_system.Row_created_by, &cs_coordinate_system.Row_created_date, &cs_coordinate_system.Row_effective_date, &cs_coordinate_system.Row_expiry_date, &cs_coordinate_system.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cs_coordinate_system to result
		result = append(result, cs_coordinate_system)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetCsCoordinateSystem(c *fiber.Ctx) error {
	var cs_coordinate_system dto.Cs_coordinate_system

	setDefaults(&cs_coordinate_system)

	if err := json.Unmarshal(c.Body(), &cs_coordinate_system); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cs_coordinate_system values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35)")
	if err != nil {
		return err
	}
	cs_coordinate_system.Row_created_date = formatDate(cs_coordinate_system.Row_created_date)
	cs_coordinate_system.Row_changed_date = formatDate(cs_coordinate_system.Row_changed_date)
	cs_coordinate_system.Effective_date = formatDateString(cs_coordinate_system.Effective_date)
	cs_coordinate_system.Expiry_date = formatDateString(cs_coordinate_system.Expiry_date)
	cs_coordinate_system.Row_effective_date = formatDateString(cs_coordinate_system.Row_effective_date)
	cs_coordinate_system.Row_expiry_date = formatDateString(cs_coordinate_system.Row_expiry_date)






	rows, err := stmt.Exec(cs_coordinate_system.Coord_system_id, cs_coordinate_system.Active_ind, cs_coordinate_system.Coordinate_system_type, cs_coordinate_system.Coord_system_abbreviation, cs_coordinate_system.Coord_system_area, cs_coordinate_system.Coord_system_long_name, cs_coordinate_system.Coord_system_short_name, cs_coordinate_system.Coord_system_uom, cs_coordinate_system.Datum_ouom, cs_coordinate_system.Effective_date, cs_coordinate_system.Expiry_date, cs_coordinate_system.Geodetic_datum, cs_coordinate_system.N_value, cs_coordinate_system.N_value_ouom, cs_coordinate_system.Owner_ba_id, cs_coordinate_system.Parent_coord_system_id, cs_coordinate_system.Perspective_height, cs_coordinate_system.Ppdm_guid, cs_coordinate_system.Prime_meridian, cs_coordinate_system.Principal_meridian, cs_coordinate_system.Projection_type, cs_coordinate_system.Reference_elev, cs_coordinate_system.Reference_elev_ouom, cs_coordinate_system.Remark, cs_coordinate_system.Rotation_ind, cs_coordinate_system.Source, cs_coordinate_system.Source_document_id, cs_coordinate_system.Vertical_datum_type, cs_coordinate_system.Row_changed_by, cs_coordinate_system.Row_changed_date, cs_coordinate_system.Row_created_by, cs_coordinate_system.Row_created_date, cs_coordinate_system.Row_effective_date, cs_coordinate_system.Row_expiry_date, cs_coordinate_system.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateCsCoordinateSystem(c *fiber.Ctx) error {
	var cs_coordinate_system dto.Cs_coordinate_system

	setDefaults(&cs_coordinate_system)

	if err := json.Unmarshal(c.Body(), &cs_coordinate_system); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cs_coordinate_system.Coord_system_id = id

    if cs_coordinate_system.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cs_coordinate_system set active_ind = :1, coordinate_system_type = :2, coord_system_abbreviation = :3, coord_system_area = :4, coord_system_long_name = :5, coord_system_short_name = :6, coord_system_uom = :7, datum_ouom = :8, effective_date = :9, expiry_date = :10, geodetic_datum = :11, n_value = :12, n_value_ouom = :13, owner_ba_id = :14, parent_coord_system_id = :15, perspective_height = :16, ppdm_guid = :17, prime_meridian = :18, principal_meridian = :19, projection_type = :20, reference_elev = :21, reference_elev_ouom = :22, remark = :23, rotation_ind = :24, source = :25, source_document_id = :26, vertical_datum_type = :27, row_changed_by = :28, row_changed_date = :29, row_created_by = :30, row_effective_date = :31, row_expiry_date = :32, row_quality = :33 where coord_system_id = :35")
	if err != nil {
		return err
	}

	cs_coordinate_system.Row_changed_date = formatDate(cs_coordinate_system.Row_changed_date)
	cs_coordinate_system.Effective_date = formatDateString(cs_coordinate_system.Effective_date)
	cs_coordinate_system.Expiry_date = formatDateString(cs_coordinate_system.Expiry_date)
	cs_coordinate_system.Row_effective_date = formatDateString(cs_coordinate_system.Row_effective_date)
	cs_coordinate_system.Row_expiry_date = formatDateString(cs_coordinate_system.Row_expiry_date)






	rows, err := stmt.Exec(cs_coordinate_system.Active_ind, cs_coordinate_system.Coordinate_system_type, cs_coordinate_system.Coord_system_abbreviation, cs_coordinate_system.Coord_system_area, cs_coordinate_system.Coord_system_long_name, cs_coordinate_system.Coord_system_short_name, cs_coordinate_system.Coord_system_uom, cs_coordinate_system.Datum_ouom, cs_coordinate_system.Effective_date, cs_coordinate_system.Expiry_date, cs_coordinate_system.Geodetic_datum, cs_coordinate_system.N_value, cs_coordinate_system.N_value_ouom, cs_coordinate_system.Owner_ba_id, cs_coordinate_system.Parent_coord_system_id, cs_coordinate_system.Perspective_height, cs_coordinate_system.Ppdm_guid, cs_coordinate_system.Prime_meridian, cs_coordinate_system.Principal_meridian, cs_coordinate_system.Projection_type, cs_coordinate_system.Reference_elev, cs_coordinate_system.Reference_elev_ouom, cs_coordinate_system.Remark, cs_coordinate_system.Rotation_ind, cs_coordinate_system.Source, cs_coordinate_system.Source_document_id, cs_coordinate_system.Vertical_datum_type, cs_coordinate_system.Row_changed_by, cs_coordinate_system.Row_changed_date, cs_coordinate_system.Row_created_by, cs_coordinate_system.Row_effective_date, cs_coordinate_system.Row_expiry_date, cs_coordinate_system.Row_quality, cs_coordinate_system.Coord_system_id)
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

func PatchCsCoordinateSystem(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cs_coordinate_system set "
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
	query += " where coord_system_id = :id"

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

func DeleteCsCoordinateSystem(c *fiber.Ctx) error {
	id := c.Params("id")
	var cs_coordinate_system dto.Cs_coordinate_system
	cs_coordinate_system.Coord_system_id = id

	stmt, err := db.Prepare("delete from cs_coordinate_system where coord_system_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cs_coordinate_system.Coord_system_id)
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



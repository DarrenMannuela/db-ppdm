package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSfMonument(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sf_monument")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sf_monument

	for rows.Next() {
		var sf_monument dto.Sf_monument
		if err := rows.Scan(&sf_monument.Sf_subtype, &sf_monument.Monument_id, &sf_monument.Active_ind, &sf_monument.Coord_acquisition_id, &sf_monument.Effective_date, &sf_monument.Expiry_date, &sf_monument.Horiz_coord_system, &sf_monument.Local_coord_system_id, &sf_monument.Location_type, &sf_monument.Monument_elev, &sf_monument.Monument_elev_ouom, &sf_monument.Monument_latitude, &sf_monument.Monument_longitude, &sf_monument.Monument_name, &sf_monument.Ppdm_guid, &sf_monument.Remark, &sf_monument.Source, &sf_monument.Source_document_id, &sf_monument.Vert_coord_system, &sf_monument.Row_changed_by, &sf_monument.Row_changed_date, &sf_monument.Row_created_by, &sf_monument.Row_created_date, &sf_monument.Row_effective_date, &sf_monument.Row_expiry_date, &sf_monument.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sf_monument to result
		result = append(result, sf_monument)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSfMonument(c *fiber.Ctx) error {
	var sf_monument dto.Sf_monument

	setDefaults(&sf_monument)

	if err := json.Unmarshal(c.Body(), &sf_monument); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sf_monument values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26)")
	if err != nil {
		return err
	}
	sf_monument.Row_created_date = formatDate(sf_monument.Row_created_date)
	sf_monument.Row_changed_date = formatDate(sf_monument.Row_changed_date)
	sf_monument.Effective_date = formatDateString(sf_monument.Effective_date)
	sf_monument.Expiry_date = formatDateString(sf_monument.Expiry_date)
	sf_monument.Row_effective_date = formatDateString(sf_monument.Row_effective_date)
	sf_monument.Row_expiry_date = formatDateString(sf_monument.Row_expiry_date)






	rows, err := stmt.Exec(sf_monument.Sf_subtype, sf_monument.Monument_id, sf_monument.Active_ind, sf_monument.Coord_acquisition_id, sf_monument.Effective_date, sf_monument.Expiry_date, sf_monument.Horiz_coord_system, sf_monument.Local_coord_system_id, sf_monument.Location_type, sf_monument.Monument_elev, sf_monument.Monument_elev_ouom, sf_monument.Monument_latitude, sf_monument.Monument_longitude, sf_monument.Monument_name, sf_monument.Ppdm_guid, sf_monument.Remark, sf_monument.Source, sf_monument.Source_document_id, sf_monument.Vert_coord_system, sf_monument.Row_changed_by, sf_monument.Row_changed_date, sf_monument.Row_created_by, sf_monument.Row_created_date, sf_monument.Row_effective_date, sf_monument.Row_expiry_date, sf_monument.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSfMonument(c *fiber.Ctx) error {
	var sf_monument dto.Sf_monument

	setDefaults(&sf_monument)

	if err := json.Unmarshal(c.Body(), &sf_monument); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sf_monument.Sf_subtype = id

    if sf_monument.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sf_monument set monument_id = :1, active_ind = :2, coord_acquisition_id = :3, effective_date = :4, expiry_date = :5, horiz_coord_system = :6, local_coord_system_id = :7, location_type = :8, monument_elev = :9, monument_elev_ouom = :10, monument_latitude = :11, monument_longitude = :12, monument_name = :13, ppdm_guid = :14, remark = :15, source = :16, source_document_id = :17, vert_coord_system = :18, row_changed_by = :19, row_changed_date = :20, row_created_by = :21, row_effective_date = :22, row_expiry_date = :23, row_quality = :24 where sf_subtype = :26")
	if err != nil {
		return err
	}

	sf_monument.Row_changed_date = formatDate(sf_monument.Row_changed_date)
	sf_monument.Effective_date = formatDateString(sf_monument.Effective_date)
	sf_monument.Expiry_date = formatDateString(sf_monument.Expiry_date)
	sf_monument.Row_effective_date = formatDateString(sf_monument.Row_effective_date)
	sf_monument.Row_expiry_date = formatDateString(sf_monument.Row_expiry_date)






	rows, err := stmt.Exec(sf_monument.Monument_id, sf_monument.Active_ind, sf_monument.Coord_acquisition_id, sf_monument.Effective_date, sf_monument.Expiry_date, sf_monument.Horiz_coord_system, sf_monument.Local_coord_system_id, sf_monument.Location_type, sf_monument.Monument_elev, sf_monument.Monument_elev_ouom, sf_monument.Monument_latitude, sf_monument.Monument_longitude, sf_monument.Monument_name, sf_monument.Ppdm_guid, sf_monument.Remark, sf_monument.Source, sf_monument.Source_document_id, sf_monument.Vert_coord_system, sf_monument.Row_changed_by, sf_monument.Row_changed_date, sf_monument.Row_created_by, sf_monument.Row_effective_date, sf_monument.Row_expiry_date, sf_monument.Row_quality, sf_monument.Sf_subtype)
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

func PatchSfMonument(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sf_monument set "
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
	query += " where sf_subtype = :id"

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

func DeleteSfMonument(c *fiber.Ctx) error {
	id := c.Params("id")
	var sf_monument dto.Sf_monument
	sf_monument.Sf_subtype = id

	stmt, err := db.Prepare("delete from sf_monument where sf_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sf_monument.Sf_subtype)
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



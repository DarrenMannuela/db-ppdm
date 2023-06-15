package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetCsGeodeticDatum(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cs_geodetic_datum")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cs_geodetic_datum

	for rows.Next() {
		var cs_geodetic_datum dto.Cs_geodetic_datum
		if err := rows.Scan(&cs_geodetic_datum.Geodetic_datum, &cs_geodetic_datum.Active_ind, &cs_geodetic_datum.Datum_origin, &cs_geodetic_datum.Effective_date, &cs_geodetic_datum.Ellipsoid_id, &cs_geodetic_datum.Expiry_date, &cs_geodetic_datum.Geodetic_datum_area_id, &cs_geodetic_datum.Geodetic_datum_area_type, &cs_geodetic_datum.Geodetic_datum_name, &cs_geodetic_datum.Origin_angular_ouom, &cs_geodetic_datum.Origin_latitude, &cs_geodetic_datum.Origin_longitude, &cs_geodetic_datum.Origin_name, &cs_geodetic_datum.Ppdm_guid, &cs_geodetic_datum.Remark, &cs_geodetic_datum.Source, &cs_geodetic_datum.Source_document_id, &cs_geodetic_datum.Row_changed_by, &cs_geodetic_datum.Row_changed_date, &cs_geodetic_datum.Row_created_by, &cs_geodetic_datum.Row_created_date, &cs_geodetic_datum.Row_effective_date, &cs_geodetic_datum.Row_expiry_date, &cs_geodetic_datum.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cs_geodetic_datum to result
		result = append(result, cs_geodetic_datum)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetCsGeodeticDatum(c *fiber.Ctx) error {
	var cs_geodetic_datum dto.Cs_geodetic_datum

	setDefaults(&cs_geodetic_datum)

	if err := json.Unmarshal(c.Body(), &cs_geodetic_datum); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cs_geodetic_datum values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24)")
	if err != nil {
		return err
	}
	cs_geodetic_datum.Row_created_date = formatDate(cs_geodetic_datum.Row_created_date)
	cs_geodetic_datum.Row_changed_date = formatDate(cs_geodetic_datum.Row_changed_date)
	cs_geodetic_datum.Effective_date = formatDateString(cs_geodetic_datum.Effective_date)
	cs_geodetic_datum.Expiry_date = formatDateString(cs_geodetic_datum.Expiry_date)
	cs_geodetic_datum.Row_effective_date = formatDateString(cs_geodetic_datum.Row_effective_date)
	cs_geodetic_datum.Row_expiry_date = formatDateString(cs_geodetic_datum.Row_expiry_date)






	rows, err := stmt.Exec(cs_geodetic_datum.Geodetic_datum, cs_geodetic_datum.Active_ind, cs_geodetic_datum.Datum_origin, cs_geodetic_datum.Effective_date, cs_geodetic_datum.Ellipsoid_id, cs_geodetic_datum.Expiry_date, cs_geodetic_datum.Geodetic_datum_area_id, cs_geodetic_datum.Geodetic_datum_area_type, cs_geodetic_datum.Geodetic_datum_name, cs_geodetic_datum.Origin_angular_ouom, cs_geodetic_datum.Origin_latitude, cs_geodetic_datum.Origin_longitude, cs_geodetic_datum.Origin_name, cs_geodetic_datum.Ppdm_guid, cs_geodetic_datum.Remark, cs_geodetic_datum.Source, cs_geodetic_datum.Source_document_id, cs_geodetic_datum.Row_changed_by, cs_geodetic_datum.Row_changed_date, cs_geodetic_datum.Row_created_by, cs_geodetic_datum.Row_created_date, cs_geodetic_datum.Row_effective_date, cs_geodetic_datum.Row_expiry_date, cs_geodetic_datum.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateCsGeodeticDatum(c *fiber.Ctx) error {
	var cs_geodetic_datum dto.Cs_geodetic_datum

	setDefaults(&cs_geodetic_datum)

	if err := json.Unmarshal(c.Body(), &cs_geodetic_datum); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cs_geodetic_datum.Geodetic_datum = id

    if cs_geodetic_datum.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cs_geodetic_datum set active_ind = :1, datum_origin = :2, effective_date = :3, ellipsoid_id = :4, expiry_date = :5, geodetic_datum_area_id = :6, geodetic_datum_area_type = :7, geodetic_datum_name = :8, origin_angular_ouom = :9, origin_latitude = :10, origin_longitude = :11, origin_name = :12, ppdm_guid = :13, remark = :14, source = :15, source_document_id = :16, row_changed_by = :17, row_changed_date = :18, row_created_by = :19, row_effective_date = :20, row_expiry_date = :21, row_quality = :22 where geodetic_datum = :24")
	if err != nil {
		return err
	}

	cs_geodetic_datum.Row_changed_date = formatDate(cs_geodetic_datum.Row_changed_date)
	cs_geodetic_datum.Effective_date = formatDateString(cs_geodetic_datum.Effective_date)
	cs_geodetic_datum.Expiry_date = formatDateString(cs_geodetic_datum.Expiry_date)
	cs_geodetic_datum.Row_effective_date = formatDateString(cs_geodetic_datum.Row_effective_date)
	cs_geodetic_datum.Row_expiry_date = formatDateString(cs_geodetic_datum.Row_expiry_date)






	rows, err := stmt.Exec(cs_geodetic_datum.Active_ind, cs_geodetic_datum.Datum_origin, cs_geodetic_datum.Effective_date, cs_geodetic_datum.Ellipsoid_id, cs_geodetic_datum.Expiry_date, cs_geodetic_datum.Geodetic_datum_area_id, cs_geodetic_datum.Geodetic_datum_area_type, cs_geodetic_datum.Geodetic_datum_name, cs_geodetic_datum.Origin_angular_ouom, cs_geodetic_datum.Origin_latitude, cs_geodetic_datum.Origin_longitude, cs_geodetic_datum.Origin_name, cs_geodetic_datum.Ppdm_guid, cs_geodetic_datum.Remark, cs_geodetic_datum.Source, cs_geodetic_datum.Source_document_id, cs_geodetic_datum.Row_changed_by, cs_geodetic_datum.Row_changed_date, cs_geodetic_datum.Row_created_by, cs_geodetic_datum.Row_effective_date, cs_geodetic_datum.Row_expiry_date, cs_geodetic_datum.Row_quality, cs_geodetic_datum.Geodetic_datum)
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

func PatchCsGeodeticDatum(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cs_geodetic_datum set "
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
	query += " where geodetic_datum = :id"

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

func DeleteCsGeodeticDatum(c *fiber.Ctx) error {
	id := c.Params("id")
	var cs_geodetic_datum dto.Cs_geodetic_datum
	cs_geodetic_datum.Geodetic_datum = id

	stmt, err := db.Prepare("delete from cs_geodetic_datum where geodetic_datum = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cs_geodetic_datum.Geodetic_datum)
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



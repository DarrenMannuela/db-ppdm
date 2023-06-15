package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetCsPrimeMeridian(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cs_prime_meridian")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cs_prime_meridian

	for rows.Next() {
		var cs_prime_meridian dto.Cs_prime_meridian
		if err := rows.Scan(&cs_prime_meridian.Prime_meridian_id, &cs_prime_meridian.Active_ind, &cs_prime_meridian.Effective_date, &cs_prime_meridian.Expiry_date, &cs_prime_meridian.Greenwich_longitude, &cs_prime_meridian.Longitude_ouom, &cs_prime_meridian.Ppdm_guid, &cs_prime_meridian.Prime_meridian_abbreviation, &cs_prime_meridian.Prime_meridian_name, &cs_prime_meridian.Remark, &cs_prime_meridian.Source, &cs_prime_meridian.Source_document_id, &cs_prime_meridian.Row_changed_by, &cs_prime_meridian.Row_changed_date, &cs_prime_meridian.Row_created_by, &cs_prime_meridian.Row_created_date, &cs_prime_meridian.Row_effective_date, &cs_prime_meridian.Row_expiry_date, &cs_prime_meridian.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cs_prime_meridian to result
		result = append(result, cs_prime_meridian)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetCsPrimeMeridian(c *fiber.Ctx) error {
	var cs_prime_meridian dto.Cs_prime_meridian

	setDefaults(&cs_prime_meridian)

	if err := json.Unmarshal(c.Body(), &cs_prime_meridian); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cs_prime_meridian values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	cs_prime_meridian.Row_created_date = formatDate(cs_prime_meridian.Row_created_date)
	cs_prime_meridian.Row_changed_date = formatDate(cs_prime_meridian.Row_changed_date)
	cs_prime_meridian.Effective_date = formatDateString(cs_prime_meridian.Effective_date)
	cs_prime_meridian.Expiry_date = formatDateString(cs_prime_meridian.Expiry_date)
	cs_prime_meridian.Row_effective_date = formatDateString(cs_prime_meridian.Row_effective_date)
	cs_prime_meridian.Row_expiry_date = formatDateString(cs_prime_meridian.Row_expiry_date)






	rows, err := stmt.Exec(cs_prime_meridian.Prime_meridian_id, cs_prime_meridian.Active_ind, cs_prime_meridian.Effective_date, cs_prime_meridian.Expiry_date, cs_prime_meridian.Greenwich_longitude, cs_prime_meridian.Longitude_ouom, cs_prime_meridian.Ppdm_guid, cs_prime_meridian.Prime_meridian_abbreviation, cs_prime_meridian.Prime_meridian_name, cs_prime_meridian.Remark, cs_prime_meridian.Source, cs_prime_meridian.Source_document_id, cs_prime_meridian.Row_changed_by, cs_prime_meridian.Row_changed_date, cs_prime_meridian.Row_created_by, cs_prime_meridian.Row_created_date, cs_prime_meridian.Row_effective_date, cs_prime_meridian.Row_expiry_date, cs_prime_meridian.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateCsPrimeMeridian(c *fiber.Ctx) error {
	var cs_prime_meridian dto.Cs_prime_meridian

	setDefaults(&cs_prime_meridian)

	if err := json.Unmarshal(c.Body(), &cs_prime_meridian); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cs_prime_meridian.Prime_meridian_id = id

    if cs_prime_meridian.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cs_prime_meridian set active_ind = :1, effective_date = :2, expiry_date = :3, greenwich_longitude = :4, longitude_ouom = :5, ppdm_guid = :6, prime_meridian_abbreviation = :7, prime_meridian_name = :8, remark = :9, source = :10, source_document_id = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where prime_meridian_id = :19")
	if err != nil {
		return err
	}

	cs_prime_meridian.Row_changed_date = formatDate(cs_prime_meridian.Row_changed_date)
	cs_prime_meridian.Effective_date = formatDateString(cs_prime_meridian.Effective_date)
	cs_prime_meridian.Expiry_date = formatDateString(cs_prime_meridian.Expiry_date)
	cs_prime_meridian.Row_effective_date = formatDateString(cs_prime_meridian.Row_effective_date)
	cs_prime_meridian.Row_expiry_date = formatDateString(cs_prime_meridian.Row_expiry_date)






	rows, err := stmt.Exec(cs_prime_meridian.Active_ind, cs_prime_meridian.Effective_date, cs_prime_meridian.Expiry_date, cs_prime_meridian.Greenwich_longitude, cs_prime_meridian.Longitude_ouom, cs_prime_meridian.Ppdm_guid, cs_prime_meridian.Prime_meridian_abbreviation, cs_prime_meridian.Prime_meridian_name, cs_prime_meridian.Remark, cs_prime_meridian.Source, cs_prime_meridian.Source_document_id, cs_prime_meridian.Row_changed_by, cs_prime_meridian.Row_changed_date, cs_prime_meridian.Row_created_by, cs_prime_meridian.Row_effective_date, cs_prime_meridian.Row_expiry_date, cs_prime_meridian.Row_quality, cs_prime_meridian.Prime_meridian_id)
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

func PatchCsPrimeMeridian(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cs_prime_meridian set "
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
	query += " where prime_meridian_id = :id"

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

func DeleteCsPrimeMeridian(c *fiber.Ctx) error {
	id := c.Params("id")
	var cs_prime_meridian dto.Cs_prime_meridian
	cs_prime_meridian.Prime_meridian_id = id

	stmt, err := db.Prepare("delete from cs_prime_meridian where prime_meridian_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cs_prime_meridian.Prime_meridian_id)
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



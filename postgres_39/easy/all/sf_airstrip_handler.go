package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSfAirstrip(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sf_airstrip")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sf_airstrip

	for rows.Next() {
		var sf_airstrip dto.Sf_airstrip
		if err := rows.Scan(&sf_airstrip.Sf_subtype, &sf_airstrip.Support_facility_id, &sf_airstrip.Active_ind, &sf_airstrip.Airstrip_type, &sf_airstrip.Effective_date, &sf_airstrip.Expiry_date, &sf_airstrip.Length, &sf_airstrip.Length_ouom, &sf_airstrip.Ppdm_guid, &sf_airstrip.Remark, &sf_airstrip.Source, &sf_airstrip.Surface_type, &sf_airstrip.Row_changed_by, &sf_airstrip.Row_changed_date, &sf_airstrip.Row_created_by, &sf_airstrip.Row_created_date, &sf_airstrip.Row_effective_date, &sf_airstrip.Row_expiry_date, &sf_airstrip.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sf_airstrip to result
		result = append(result, sf_airstrip)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSfAirstrip(c *fiber.Ctx) error {
	var sf_airstrip dto.Sf_airstrip

	setDefaults(&sf_airstrip)

	if err := json.Unmarshal(c.Body(), &sf_airstrip); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sf_airstrip values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	sf_airstrip.Row_created_date = formatDate(sf_airstrip.Row_created_date)
	sf_airstrip.Row_changed_date = formatDate(sf_airstrip.Row_changed_date)
	sf_airstrip.Effective_date = formatDateString(sf_airstrip.Effective_date)
	sf_airstrip.Expiry_date = formatDateString(sf_airstrip.Expiry_date)
	sf_airstrip.Row_effective_date = formatDateString(sf_airstrip.Row_effective_date)
	sf_airstrip.Row_expiry_date = formatDateString(sf_airstrip.Row_expiry_date)






	rows, err := stmt.Exec(sf_airstrip.Sf_subtype, sf_airstrip.Support_facility_id, sf_airstrip.Active_ind, sf_airstrip.Airstrip_type, sf_airstrip.Effective_date, sf_airstrip.Expiry_date, sf_airstrip.Length, sf_airstrip.Length_ouom, sf_airstrip.Ppdm_guid, sf_airstrip.Remark, sf_airstrip.Source, sf_airstrip.Surface_type, sf_airstrip.Row_changed_by, sf_airstrip.Row_changed_date, sf_airstrip.Row_created_by, sf_airstrip.Row_created_date, sf_airstrip.Row_effective_date, sf_airstrip.Row_expiry_date, sf_airstrip.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSfAirstrip(c *fiber.Ctx) error {
	var sf_airstrip dto.Sf_airstrip

	setDefaults(&sf_airstrip)

	if err := json.Unmarshal(c.Body(), &sf_airstrip); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sf_airstrip.Sf_subtype = id

    if sf_airstrip.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sf_airstrip set support_facility_id = :1, active_ind = :2, airstrip_type = :3, effective_date = :4, expiry_date = :5, length = :6, length_ouom = :7, ppdm_guid = :8, remark = :9, source = :10, surface_type = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where sf_subtype = :19")
	if err != nil {
		return err
	}

	sf_airstrip.Row_changed_date = formatDate(sf_airstrip.Row_changed_date)
	sf_airstrip.Effective_date = formatDateString(sf_airstrip.Effective_date)
	sf_airstrip.Expiry_date = formatDateString(sf_airstrip.Expiry_date)
	sf_airstrip.Row_effective_date = formatDateString(sf_airstrip.Row_effective_date)
	sf_airstrip.Row_expiry_date = formatDateString(sf_airstrip.Row_expiry_date)






	rows, err := stmt.Exec(sf_airstrip.Support_facility_id, sf_airstrip.Active_ind, sf_airstrip.Airstrip_type, sf_airstrip.Effective_date, sf_airstrip.Expiry_date, sf_airstrip.Length, sf_airstrip.Length_ouom, sf_airstrip.Ppdm_guid, sf_airstrip.Remark, sf_airstrip.Source, sf_airstrip.Surface_type, sf_airstrip.Row_changed_by, sf_airstrip.Row_changed_date, sf_airstrip.Row_created_by, sf_airstrip.Row_effective_date, sf_airstrip.Row_expiry_date, sf_airstrip.Row_quality, sf_airstrip.Sf_subtype)
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

func PatchSfAirstrip(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sf_airstrip set "
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

func DeleteSfAirstrip(c *fiber.Ctx) error {
	id := c.Params("id")
	var sf_airstrip dto.Sf_airstrip
	sf_airstrip.Sf_subtype = id

	stmt, err := db.Prepare("delete from sf_airstrip where sf_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sf_airstrip.Sf_subtype)
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



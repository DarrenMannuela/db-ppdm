package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSfRestriction(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sf_restriction")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sf_restriction

	for rows.Next() {
		var sf_restriction dto.Sf_restriction
		if err := rows.Scan(&sf_restriction.Sf_subtype, &sf_restriction.Support_facility_id, &sf_restriction.Restriction_id, &sf_restriction.Restriction_version, &sf_restriction.Active_ind, &sf_restriction.Effective_date, &sf_restriction.Expiry_date, &sf_restriction.Ppdm_guid, &sf_restriction.Remark, &sf_restriction.Source, &sf_restriction.Row_changed_by, &sf_restriction.Row_changed_date, &sf_restriction.Row_created_by, &sf_restriction.Row_created_date, &sf_restriction.Row_effective_date, &sf_restriction.Row_expiry_date, &sf_restriction.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sf_restriction to result
		result = append(result, sf_restriction)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSfRestriction(c *fiber.Ctx) error {
	var sf_restriction dto.Sf_restriction

	setDefaults(&sf_restriction)

	if err := json.Unmarshal(c.Body(), &sf_restriction); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sf_restriction values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	sf_restriction.Row_created_date = formatDate(sf_restriction.Row_created_date)
	sf_restriction.Row_changed_date = formatDate(sf_restriction.Row_changed_date)
	sf_restriction.Effective_date = formatDateString(sf_restriction.Effective_date)
	sf_restriction.Expiry_date = formatDateString(sf_restriction.Expiry_date)
	sf_restriction.Row_effective_date = formatDateString(sf_restriction.Row_effective_date)
	sf_restriction.Row_expiry_date = formatDateString(sf_restriction.Row_expiry_date)






	rows, err := stmt.Exec(sf_restriction.Sf_subtype, sf_restriction.Support_facility_id, sf_restriction.Restriction_id, sf_restriction.Restriction_version, sf_restriction.Active_ind, sf_restriction.Effective_date, sf_restriction.Expiry_date, sf_restriction.Ppdm_guid, sf_restriction.Remark, sf_restriction.Source, sf_restriction.Row_changed_by, sf_restriction.Row_changed_date, sf_restriction.Row_created_by, sf_restriction.Row_created_date, sf_restriction.Row_effective_date, sf_restriction.Row_expiry_date, sf_restriction.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSfRestriction(c *fiber.Ctx) error {
	var sf_restriction dto.Sf_restriction

	setDefaults(&sf_restriction)

	if err := json.Unmarshal(c.Body(), &sf_restriction); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sf_restriction.Sf_subtype = id

    if sf_restriction.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sf_restriction set support_facility_id = :1, restriction_id = :2, restriction_version = :3, active_ind = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where sf_subtype = :17")
	if err != nil {
		return err
	}

	sf_restriction.Row_changed_date = formatDate(sf_restriction.Row_changed_date)
	sf_restriction.Effective_date = formatDateString(sf_restriction.Effective_date)
	sf_restriction.Expiry_date = formatDateString(sf_restriction.Expiry_date)
	sf_restriction.Row_effective_date = formatDateString(sf_restriction.Row_effective_date)
	sf_restriction.Row_expiry_date = formatDateString(sf_restriction.Row_expiry_date)






	rows, err := stmt.Exec(sf_restriction.Support_facility_id, sf_restriction.Restriction_id, sf_restriction.Restriction_version, sf_restriction.Active_ind, sf_restriction.Effective_date, sf_restriction.Expiry_date, sf_restriction.Ppdm_guid, sf_restriction.Remark, sf_restriction.Source, sf_restriction.Row_changed_by, sf_restriction.Row_changed_date, sf_restriction.Row_created_by, sf_restriction.Row_effective_date, sf_restriction.Row_expiry_date, sf_restriction.Row_quality, sf_restriction.Sf_subtype)
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

func PatchSfRestriction(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sf_restriction set "
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

func DeleteSfRestriction(c *fiber.Ctx) error {
	id := c.Params("id")
	var sf_restriction dto.Sf_restriction
	sf_restriction.Sf_subtype = id

	stmt, err := db.Prepare("delete from sf_restriction where sf_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sf_restriction.Sf_subtype)
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



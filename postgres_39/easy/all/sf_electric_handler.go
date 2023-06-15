package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSfElectric(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sf_electric")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sf_electric

	for rows.Next() {
		var sf_electric dto.Sf_electric
		if err := rows.Scan(&sf_electric.Sf_subtype, &sf_electric.Support_facility_id, &sf_electric.Active_ind, &sf_electric.Effective_date, &sf_electric.Electric_facility_type, &sf_electric.Expiry_date, &sf_electric.Ppdm_guid, &sf_electric.Remark, &sf_electric.Source, &sf_electric.Row_changed_by, &sf_electric.Row_changed_date, &sf_electric.Row_created_by, &sf_electric.Row_created_date, &sf_electric.Row_effective_date, &sf_electric.Row_expiry_date, &sf_electric.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sf_electric to result
		result = append(result, sf_electric)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSfElectric(c *fiber.Ctx) error {
	var sf_electric dto.Sf_electric

	setDefaults(&sf_electric)

	if err := json.Unmarshal(c.Body(), &sf_electric); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sf_electric values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16)")
	if err != nil {
		return err
	}
	sf_electric.Row_created_date = formatDate(sf_electric.Row_created_date)
	sf_electric.Row_changed_date = formatDate(sf_electric.Row_changed_date)
	sf_electric.Effective_date = formatDateString(sf_electric.Effective_date)
	sf_electric.Expiry_date = formatDateString(sf_electric.Expiry_date)
	sf_electric.Row_effective_date = formatDateString(sf_electric.Row_effective_date)
	sf_electric.Row_expiry_date = formatDateString(sf_electric.Row_expiry_date)






	rows, err := stmt.Exec(sf_electric.Sf_subtype, sf_electric.Support_facility_id, sf_electric.Active_ind, sf_electric.Effective_date, sf_electric.Electric_facility_type, sf_electric.Expiry_date, sf_electric.Ppdm_guid, sf_electric.Remark, sf_electric.Source, sf_electric.Row_changed_by, sf_electric.Row_changed_date, sf_electric.Row_created_by, sf_electric.Row_created_date, sf_electric.Row_effective_date, sf_electric.Row_expiry_date, sf_electric.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSfElectric(c *fiber.Ctx) error {
	var sf_electric dto.Sf_electric

	setDefaults(&sf_electric)

	if err := json.Unmarshal(c.Body(), &sf_electric); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sf_electric.Sf_subtype = id

    if sf_electric.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sf_electric set support_facility_id = :1, active_ind = :2, effective_date = :3, electric_facility_type = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, source = :8, row_changed_by = :9, row_changed_date = :10, row_created_by = :11, row_effective_date = :12, row_expiry_date = :13, row_quality = :14 where sf_subtype = :16")
	if err != nil {
		return err
	}

	sf_electric.Row_changed_date = formatDate(sf_electric.Row_changed_date)
	sf_electric.Effective_date = formatDateString(sf_electric.Effective_date)
	sf_electric.Expiry_date = formatDateString(sf_electric.Expiry_date)
	sf_electric.Row_effective_date = formatDateString(sf_electric.Row_effective_date)
	sf_electric.Row_expiry_date = formatDateString(sf_electric.Row_expiry_date)






	rows, err := stmt.Exec(sf_electric.Support_facility_id, sf_electric.Active_ind, sf_electric.Effective_date, sf_electric.Electric_facility_type, sf_electric.Expiry_date, sf_electric.Ppdm_guid, sf_electric.Remark, sf_electric.Source, sf_electric.Row_changed_by, sf_electric.Row_changed_date, sf_electric.Row_created_by, sf_electric.Row_effective_date, sf_electric.Row_expiry_date, sf_electric.Row_quality, sf_electric.Sf_subtype)
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

func PatchSfElectric(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sf_electric set "
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

func DeleteSfElectric(c *fiber.Ctx) error {
	id := c.Params("id")
	var sf_electric dto.Sf_electric
	sf_electric.Sf_subtype = id

	stmt, err := db.Prepare("delete from sf_electric where sf_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sf_electric.Sf_subtype)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSfXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from sf_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Sf_xref

	for rows.Next() {
		var sf_xref dto.Sf_xref
		if err := rows.Scan(&sf_xref.Sf_subtype, &sf_xref.Support_facility_id, &sf_xref.Sf_subtype2, &sf_xref.Support_facility_id2, &sf_xref.Active_ind, &sf_xref.Effective_date, &sf_xref.Expiry_date, &sf_xref.From_distance, &sf_xref.From_distance_ouom, &sf_xref.Portion_percent, &sf_xref.Ppdm_guid, &sf_xref.Remark, &sf_xref.Source, &sf_xref.To_distance, &sf_xref.To_distance_ouom, &sf_xref.Xref_type, &sf_xref.Row_changed_by, &sf_xref.Row_changed_date, &sf_xref.Row_created_by, &sf_xref.Row_created_date, &sf_xref.Row_effective_date, &sf_xref.Row_expiry_date, &sf_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append sf_xref to result
		result = append(result, sf_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSfXref(c *fiber.Ctx) error {
	var sf_xref dto.Sf_xref

	setDefaults(&sf_xref)

	if err := json.Unmarshal(c.Body(), &sf_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into sf_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	sf_xref.Row_created_date = formatDate(sf_xref.Row_created_date)
	sf_xref.Row_changed_date = formatDate(sf_xref.Row_changed_date)
	sf_xref.Effective_date = formatDateString(sf_xref.Effective_date)
	sf_xref.Expiry_date = formatDateString(sf_xref.Expiry_date)
	sf_xref.Row_effective_date = formatDateString(sf_xref.Row_effective_date)
	sf_xref.Row_expiry_date = formatDateString(sf_xref.Row_expiry_date)






	rows, err := stmt.Exec(sf_xref.Sf_subtype, sf_xref.Support_facility_id, sf_xref.Sf_subtype2, sf_xref.Support_facility_id2, sf_xref.Active_ind, sf_xref.Effective_date, sf_xref.Expiry_date, sf_xref.From_distance, sf_xref.From_distance_ouom, sf_xref.Portion_percent, sf_xref.Ppdm_guid, sf_xref.Remark, sf_xref.Source, sf_xref.To_distance, sf_xref.To_distance_ouom, sf_xref.Xref_type, sf_xref.Row_changed_by, sf_xref.Row_changed_date, sf_xref.Row_created_by, sf_xref.Row_created_date, sf_xref.Row_effective_date, sf_xref.Row_expiry_date, sf_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSfXref(c *fiber.Ctx) error {
	var sf_xref dto.Sf_xref

	setDefaults(&sf_xref)

	if err := json.Unmarshal(c.Body(), &sf_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	sf_xref.Sf_subtype = id

    if sf_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update sf_xref set support_facility_id = :1, sf_subtype2 = :2, support_facility_id2 = :3, active_ind = :4, effective_date = :5, expiry_date = :6, from_distance = :7, from_distance_ouom = :8, portion_percent = :9, ppdm_guid = :10, remark = :11, source = :12, to_distance = :13, to_distance_ouom = :14, xref_type = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where sf_subtype = :23")
	if err != nil {
		return err
	}

	sf_xref.Row_changed_date = formatDate(sf_xref.Row_changed_date)
	sf_xref.Effective_date = formatDateString(sf_xref.Effective_date)
	sf_xref.Expiry_date = formatDateString(sf_xref.Expiry_date)
	sf_xref.Row_effective_date = formatDateString(sf_xref.Row_effective_date)
	sf_xref.Row_expiry_date = formatDateString(sf_xref.Row_expiry_date)






	rows, err := stmt.Exec(sf_xref.Support_facility_id, sf_xref.Sf_subtype2, sf_xref.Support_facility_id2, sf_xref.Active_ind, sf_xref.Effective_date, sf_xref.Expiry_date, sf_xref.From_distance, sf_xref.From_distance_ouom, sf_xref.Portion_percent, sf_xref.Ppdm_guid, sf_xref.Remark, sf_xref.Source, sf_xref.To_distance, sf_xref.To_distance_ouom, sf_xref.Xref_type, sf_xref.Row_changed_by, sf_xref.Row_changed_date, sf_xref.Row_created_by, sf_xref.Row_effective_date, sf_xref.Row_expiry_date, sf_xref.Row_quality, sf_xref.Sf_subtype)
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

func PatchSfXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update sf_xref set "
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

func DeleteSfXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var sf_xref dto.Sf_xref
	sf_xref.Sf_subtype = id

	stmt, err := db.Prepare("delete from sf_xref where sf_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(sf_xref.Sf_subtype)
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



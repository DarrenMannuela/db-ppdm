package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRaLithology(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ra_lithology")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ra_lithology

	for rows.Next() {
		var ra_lithology dto.Ra_lithology
		if err := rows.Scan(&ra_lithology.Lithology, &ra_lithology.Alias_id, &ra_lithology.Abbreviation, &ra_lithology.Active_ind, &ra_lithology.Alias_long_name, &ra_lithology.Alias_reason_type, &ra_lithology.Alias_short_name, &ra_lithology.Alias_type, &ra_lithology.Amended_date, &ra_lithology.Created_date, &ra_lithology.Effective_date, &ra_lithology.Expiry_date, &ra_lithology.Lithology, &ra_lithology.Original_ind, &ra_lithology.Owner_ba_id, &ra_lithology.Ppdm_guid, &ra_lithology.Preferred_ind, &ra_lithology.Reason_desc, &ra_lithology.Remark, &ra_lithology.Source, &ra_lithology.Source_document, &ra_lithology.Struckoff_date, &ra_lithology.Sw_application_id, &ra_lithology.Use_rule, &ra_lithology.Row_changed_by, &ra_lithology.Row_changed_date, &ra_lithology.Row_created_by, &ra_lithology.Row_created_date, &ra_lithology.Row_effective_date, &ra_lithology.Row_expiry_date, &ra_lithology.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ra_lithology to result
		result = append(result, ra_lithology)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRaLithology(c *fiber.Ctx) error {
	var ra_lithology dto.Ra_lithology

	setDefaults(&ra_lithology)

	if err := json.Unmarshal(c.Body(), &ra_lithology); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ra_lithology values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31)")
	if err != nil {
		return err
	}
	ra_lithology.Row_created_date = formatDate(ra_lithology.Row_created_date)
	ra_lithology.Row_changed_date = formatDate(ra_lithology.Row_changed_date)
	ra_lithology.Amended_date = formatDateString(ra_lithology.Amended_date)
	ra_lithology.Created_date = formatDateString(ra_lithology.Created_date)
	ra_lithology.Effective_date = formatDateString(ra_lithology.Effective_date)
	ra_lithology.Expiry_date = formatDateString(ra_lithology.Expiry_date)
	ra_lithology.Struckoff_date = formatDateString(ra_lithology.Struckoff_date)
	ra_lithology.Row_effective_date = formatDateString(ra_lithology.Row_effective_date)
	ra_lithology.Row_expiry_date = formatDateString(ra_lithology.Row_expiry_date)






	rows, err := stmt.Exec(ra_lithology.Lithology, ra_lithology.Alias_id, ra_lithology.Abbreviation, ra_lithology.Active_ind, ra_lithology.Alias_long_name, ra_lithology.Alias_reason_type, ra_lithology.Alias_short_name, ra_lithology.Alias_type, ra_lithology.Amended_date, ra_lithology.Created_date, ra_lithology.Effective_date, ra_lithology.Expiry_date, ra_lithology.Lithology, ra_lithology.Original_ind, ra_lithology.Owner_ba_id, ra_lithology.Ppdm_guid, ra_lithology.Preferred_ind, ra_lithology.Reason_desc, ra_lithology.Remark, ra_lithology.Source, ra_lithology.Source_document, ra_lithology.Struckoff_date, ra_lithology.Sw_application_id, ra_lithology.Use_rule, ra_lithology.Row_changed_by, ra_lithology.Row_changed_date, ra_lithology.Row_created_by, ra_lithology.Row_created_date, ra_lithology.Row_effective_date, ra_lithology.Row_expiry_date, ra_lithology.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRaLithology(c *fiber.Ctx) error {
	var ra_lithology dto.Ra_lithology

	setDefaults(&ra_lithology)

	if err := json.Unmarshal(c.Body(), &ra_lithology); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ra_lithology.Lithology = id

    if ra_lithology.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ra_lithology set alias_id = :1, abbreviation = :2, active_ind = :3, alias_long_name = :4, alias_reason_type = :5, alias_short_name = :6, alias_type = :7, amended_date = :8, created_date = :9, effective_date = :10, expiry_date = :11, lithology = :12, original_ind = :13, owner_ba_id = :14, ppdm_guid = :15, preferred_ind = :16, reason_desc = :17, remark = :18, source = :19, source_document = :20, struckoff_date = :21, sw_application_id = :22, use_rule = :23, row_changed_by = :24, row_changed_date = :25, row_created_by = :26, row_effective_date = :27, row_expiry_date = :28, row_quality = :29 where lithology = :31")
	if err != nil {
		return err
	}

	ra_lithology.Row_changed_date = formatDate(ra_lithology.Row_changed_date)
	ra_lithology.Amended_date = formatDateString(ra_lithology.Amended_date)
	ra_lithology.Created_date = formatDateString(ra_lithology.Created_date)
	ra_lithology.Effective_date = formatDateString(ra_lithology.Effective_date)
	ra_lithology.Expiry_date = formatDateString(ra_lithology.Expiry_date)
	ra_lithology.Struckoff_date = formatDateString(ra_lithology.Struckoff_date)
	ra_lithology.Row_effective_date = formatDateString(ra_lithology.Row_effective_date)
	ra_lithology.Row_expiry_date = formatDateString(ra_lithology.Row_expiry_date)






	rows, err := stmt.Exec(ra_lithology.Alias_id, ra_lithology.Abbreviation, ra_lithology.Active_ind, ra_lithology.Alias_long_name, ra_lithology.Alias_reason_type, ra_lithology.Alias_short_name, ra_lithology.Alias_type, ra_lithology.Amended_date, ra_lithology.Created_date, ra_lithology.Effective_date, ra_lithology.Expiry_date, ra_lithology.Lithology, ra_lithology.Original_ind, ra_lithology.Owner_ba_id, ra_lithology.Ppdm_guid, ra_lithology.Preferred_ind, ra_lithology.Reason_desc, ra_lithology.Remark, ra_lithology.Source, ra_lithology.Source_document, ra_lithology.Struckoff_date, ra_lithology.Sw_application_id, ra_lithology.Use_rule, ra_lithology.Row_changed_by, ra_lithology.Row_changed_date, ra_lithology.Row_created_by, ra_lithology.Row_effective_date, ra_lithology.Row_expiry_date, ra_lithology.Row_quality, ra_lithology.Lithology)
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

func PatchRaLithology(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ra_lithology set "
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
		} else if key == "amended_date" || key == "created_date" || key == "effective_date" || key == "expiry_date" || key == "struckoff_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where lithology = :id"

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

func DeleteRaLithology(c *fiber.Ctx) error {
	id := c.Params("id")
	var ra_lithology dto.Ra_lithology
	ra_lithology.Lithology = id

	stmt, err := db.Prepare("delete from ra_lithology where lithology = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ra_lithology.Lithology)
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



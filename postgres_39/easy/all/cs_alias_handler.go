package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetCsAlias(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cs_alias")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cs_alias

	for rows.Next() {
		var cs_alias dto.Cs_alias
		if err := rows.Scan(&cs_alias.Coord_system_id, &cs_alias.Coord_system_alias_id, &cs_alias.Abbreviation, &cs_alias.Active_ind, &cs_alias.Alias_long_name, &cs_alias.Alias_reason_type, &cs_alias.Alias_short_name, &cs_alias.Alias_type, &cs_alias.Amended_date, &cs_alias.Created_date, &cs_alias.Effective_date, &cs_alias.Expiry_date, &cs_alias.Coord_system_id, &cs_alias.Original_ind, &cs_alias.Owner_ba_id, &cs_alias.Ppdm_guid, &cs_alias.Preferred_ind, &cs_alias.Reason_desc, &cs_alias.Remark, &cs_alias.Source, &cs_alias.Source_document_id, &cs_alias.Struckoff_date, &cs_alias.Sw_application_id, &cs_alias.Use_rule, &cs_alias.Row_changed_by, &cs_alias.Row_changed_date, &cs_alias.Row_created_by, &cs_alias.Row_created_date, &cs_alias.Row_effective_date, &cs_alias.Row_expiry_date, &cs_alias.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cs_alias to result
		result = append(result, cs_alias)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetCsAlias(c *fiber.Ctx) error {
	var cs_alias dto.Cs_alias

	setDefaults(&cs_alias)

	if err := json.Unmarshal(c.Body(), &cs_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cs_alias values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31)")
	if err != nil {
		return err
	}
	cs_alias.Row_created_date = formatDate(cs_alias.Row_created_date)
	cs_alias.Row_changed_date = formatDate(cs_alias.Row_changed_date)
	cs_alias.Amended_date = formatDateString(cs_alias.Amended_date)
	cs_alias.Created_date = formatDateString(cs_alias.Created_date)
	cs_alias.Effective_date = formatDateString(cs_alias.Effective_date)
	cs_alias.Expiry_date = formatDateString(cs_alias.Expiry_date)
	cs_alias.Struckoff_date = formatDateString(cs_alias.Struckoff_date)
	cs_alias.Row_effective_date = formatDateString(cs_alias.Row_effective_date)
	cs_alias.Row_expiry_date = formatDateString(cs_alias.Row_expiry_date)






	rows, err := stmt.Exec(cs_alias.Coord_system_id, cs_alias.Coord_system_alias_id, cs_alias.Abbreviation, cs_alias.Active_ind, cs_alias.Alias_long_name, cs_alias.Alias_reason_type, cs_alias.Alias_short_name, cs_alias.Alias_type, cs_alias.Amended_date, cs_alias.Created_date, cs_alias.Effective_date, cs_alias.Expiry_date, cs_alias.Coord_system_id, cs_alias.Original_ind, cs_alias.Owner_ba_id, cs_alias.Ppdm_guid, cs_alias.Preferred_ind, cs_alias.Reason_desc, cs_alias.Remark, cs_alias.Source, cs_alias.Source_document_id, cs_alias.Struckoff_date, cs_alias.Sw_application_id, cs_alias.Use_rule, cs_alias.Row_changed_by, cs_alias.Row_changed_date, cs_alias.Row_created_by, cs_alias.Row_created_date, cs_alias.Row_effective_date, cs_alias.Row_expiry_date, cs_alias.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateCsAlias(c *fiber.Ctx) error {
	var cs_alias dto.Cs_alias

	setDefaults(&cs_alias)

	if err := json.Unmarshal(c.Body(), &cs_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cs_alias.Coord_system_id = id

    if cs_alias.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cs_alias set coord_system_alias_id = :1, abbreviation = :2, active_ind = :3, alias_long_name = :4, alias_reason_type = :5, alias_short_name = :6, alias_type = :7, amended_date = :8, created_date = :9, effective_date = :10, expiry_date = :11, coord_system_id = :12, original_ind = :13, owner_ba_id = :14, ppdm_guid = :15, preferred_ind = :16, reason_desc = :17, remark = :18, source = :19, source_document_id = :20, struckoff_date = :21, sw_application_id = :22, use_rule = :23, row_changed_by = :24, row_changed_date = :25, row_created_by = :26, row_effective_date = :27, row_expiry_date = :28, row_quality = :29 where coord_system_id = :31")
	if err != nil {
		return err
	}

	cs_alias.Row_changed_date = formatDate(cs_alias.Row_changed_date)
	cs_alias.Amended_date = formatDateString(cs_alias.Amended_date)
	cs_alias.Created_date = formatDateString(cs_alias.Created_date)
	cs_alias.Effective_date = formatDateString(cs_alias.Effective_date)
	cs_alias.Expiry_date = formatDateString(cs_alias.Expiry_date)
	cs_alias.Struckoff_date = formatDateString(cs_alias.Struckoff_date)
	cs_alias.Row_effective_date = formatDateString(cs_alias.Row_effective_date)
	cs_alias.Row_expiry_date = formatDateString(cs_alias.Row_expiry_date)






	rows, err := stmt.Exec(cs_alias.Coord_system_alias_id, cs_alias.Abbreviation, cs_alias.Active_ind, cs_alias.Alias_long_name, cs_alias.Alias_reason_type, cs_alias.Alias_short_name, cs_alias.Alias_type, cs_alias.Amended_date, cs_alias.Created_date, cs_alias.Effective_date, cs_alias.Expiry_date, cs_alias.Coord_system_id, cs_alias.Original_ind, cs_alias.Owner_ba_id, cs_alias.Ppdm_guid, cs_alias.Preferred_ind, cs_alias.Reason_desc, cs_alias.Remark, cs_alias.Source, cs_alias.Source_document_id, cs_alias.Struckoff_date, cs_alias.Sw_application_id, cs_alias.Use_rule, cs_alias.Row_changed_by, cs_alias.Row_changed_date, cs_alias.Row_created_by, cs_alias.Row_effective_date, cs_alias.Row_expiry_date, cs_alias.Row_quality, cs_alias.Coord_system_id)
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

func PatchCsAlias(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cs_alias set "
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

func DeleteCsAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var cs_alias dto.Cs_alias
	cs_alias.Coord_system_id = id

	stmt, err := db.Prepare("delete from cs_alias where coord_system_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cs_alias.Coord_system_id)
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



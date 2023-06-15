package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRaLithScaleScheme(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ra_lith_scale_scheme")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ra_lith_scale_scheme

	for rows.Next() {
		var ra_lith_scale_scheme dto.Ra_lith_scale_scheme
		if err := rows.Scan(&ra_lith_scale_scheme.Scale_scheme, &ra_lith_scale_scheme.Alias_id, &ra_lith_scale_scheme.Abbreviation, &ra_lith_scale_scheme.Active_ind, &ra_lith_scale_scheme.Alias_long_name, &ra_lith_scale_scheme.Alias_reason_type, &ra_lith_scale_scheme.Alias_short_name, &ra_lith_scale_scheme.Alias_type, &ra_lith_scale_scheme.Amended_date, &ra_lith_scale_scheme.Created_date, &ra_lith_scale_scheme.Effective_date, &ra_lith_scale_scheme.Expiry_date, &ra_lith_scale_scheme.Scale_scheme, &ra_lith_scale_scheme.Original_ind, &ra_lith_scale_scheme.Owner_ba_id, &ra_lith_scale_scheme.Ppdm_guid, &ra_lith_scale_scheme.Preferred_ind, &ra_lith_scale_scheme.Reason_desc, &ra_lith_scale_scheme.Remark, &ra_lith_scale_scheme.Source, &ra_lith_scale_scheme.Source_document, &ra_lith_scale_scheme.Struckoff_date, &ra_lith_scale_scheme.Sw_application_id, &ra_lith_scale_scheme.Use_rule, &ra_lith_scale_scheme.Row_changed_by, &ra_lith_scale_scheme.Row_changed_date, &ra_lith_scale_scheme.Row_created_by, &ra_lith_scale_scheme.Row_created_date, &ra_lith_scale_scheme.Row_effective_date, &ra_lith_scale_scheme.Row_expiry_date, &ra_lith_scale_scheme.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ra_lith_scale_scheme to result
		result = append(result, ra_lith_scale_scheme)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRaLithScaleScheme(c *fiber.Ctx) error {
	var ra_lith_scale_scheme dto.Ra_lith_scale_scheme

	setDefaults(&ra_lith_scale_scheme)

	if err := json.Unmarshal(c.Body(), &ra_lith_scale_scheme); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ra_lith_scale_scheme values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31)")
	if err != nil {
		return err
	}
	ra_lith_scale_scheme.Row_created_date = formatDate(ra_lith_scale_scheme.Row_created_date)
	ra_lith_scale_scheme.Row_changed_date = formatDate(ra_lith_scale_scheme.Row_changed_date)
	ra_lith_scale_scheme.Amended_date = formatDateString(ra_lith_scale_scheme.Amended_date)
	ra_lith_scale_scheme.Created_date = formatDateString(ra_lith_scale_scheme.Created_date)
	ra_lith_scale_scheme.Effective_date = formatDateString(ra_lith_scale_scheme.Effective_date)
	ra_lith_scale_scheme.Expiry_date = formatDateString(ra_lith_scale_scheme.Expiry_date)
	ra_lith_scale_scheme.Struckoff_date = formatDateString(ra_lith_scale_scheme.Struckoff_date)
	ra_lith_scale_scheme.Row_effective_date = formatDateString(ra_lith_scale_scheme.Row_effective_date)
	ra_lith_scale_scheme.Row_expiry_date = formatDateString(ra_lith_scale_scheme.Row_expiry_date)






	rows, err := stmt.Exec(ra_lith_scale_scheme.Scale_scheme, ra_lith_scale_scheme.Alias_id, ra_lith_scale_scheme.Abbreviation, ra_lith_scale_scheme.Active_ind, ra_lith_scale_scheme.Alias_long_name, ra_lith_scale_scheme.Alias_reason_type, ra_lith_scale_scheme.Alias_short_name, ra_lith_scale_scheme.Alias_type, ra_lith_scale_scheme.Amended_date, ra_lith_scale_scheme.Created_date, ra_lith_scale_scheme.Effective_date, ra_lith_scale_scheme.Expiry_date, ra_lith_scale_scheme.Scale_scheme, ra_lith_scale_scheme.Original_ind, ra_lith_scale_scheme.Owner_ba_id, ra_lith_scale_scheme.Ppdm_guid, ra_lith_scale_scheme.Preferred_ind, ra_lith_scale_scheme.Reason_desc, ra_lith_scale_scheme.Remark, ra_lith_scale_scheme.Source, ra_lith_scale_scheme.Source_document, ra_lith_scale_scheme.Struckoff_date, ra_lith_scale_scheme.Sw_application_id, ra_lith_scale_scheme.Use_rule, ra_lith_scale_scheme.Row_changed_by, ra_lith_scale_scheme.Row_changed_date, ra_lith_scale_scheme.Row_created_by, ra_lith_scale_scheme.Row_created_date, ra_lith_scale_scheme.Row_effective_date, ra_lith_scale_scheme.Row_expiry_date, ra_lith_scale_scheme.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRaLithScaleScheme(c *fiber.Ctx) error {
	var ra_lith_scale_scheme dto.Ra_lith_scale_scheme

	setDefaults(&ra_lith_scale_scheme)

	if err := json.Unmarshal(c.Body(), &ra_lith_scale_scheme); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ra_lith_scale_scheme.Scale_scheme = id

    if ra_lith_scale_scheme.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ra_lith_scale_scheme set alias_id = :1, abbreviation = :2, active_ind = :3, alias_long_name = :4, alias_reason_type = :5, alias_short_name = :6, alias_type = :7, amended_date = :8, created_date = :9, effective_date = :10, expiry_date = :11, scale_scheme = :12, original_ind = :13, owner_ba_id = :14, ppdm_guid = :15, preferred_ind = :16, reason_desc = :17, remark = :18, source = :19, source_document = :20, struckoff_date = :21, sw_application_id = :22, use_rule = :23, row_changed_by = :24, row_changed_date = :25, row_created_by = :26, row_effective_date = :27, row_expiry_date = :28, row_quality = :29 where scale_scheme = :31")
	if err != nil {
		return err
	}

	ra_lith_scale_scheme.Row_changed_date = formatDate(ra_lith_scale_scheme.Row_changed_date)
	ra_lith_scale_scheme.Amended_date = formatDateString(ra_lith_scale_scheme.Amended_date)
	ra_lith_scale_scheme.Created_date = formatDateString(ra_lith_scale_scheme.Created_date)
	ra_lith_scale_scheme.Effective_date = formatDateString(ra_lith_scale_scheme.Effective_date)
	ra_lith_scale_scheme.Expiry_date = formatDateString(ra_lith_scale_scheme.Expiry_date)
	ra_lith_scale_scheme.Struckoff_date = formatDateString(ra_lith_scale_scheme.Struckoff_date)
	ra_lith_scale_scheme.Row_effective_date = formatDateString(ra_lith_scale_scheme.Row_effective_date)
	ra_lith_scale_scheme.Row_expiry_date = formatDateString(ra_lith_scale_scheme.Row_expiry_date)






	rows, err := stmt.Exec(ra_lith_scale_scheme.Alias_id, ra_lith_scale_scheme.Abbreviation, ra_lith_scale_scheme.Active_ind, ra_lith_scale_scheme.Alias_long_name, ra_lith_scale_scheme.Alias_reason_type, ra_lith_scale_scheme.Alias_short_name, ra_lith_scale_scheme.Alias_type, ra_lith_scale_scheme.Amended_date, ra_lith_scale_scheme.Created_date, ra_lith_scale_scheme.Effective_date, ra_lith_scale_scheme.Expiry_date, ra_lith_scale_scheme.Scale_scheme, ra_lith_scale_scheme.Original_ind, ra_lith_scale_scheme.Owner_ba_id, ra_lith_scale_scheme.Ppdm_guid, ra_lith_scale_scheme.Preferred_ind, ra_lith_scale_scheme.Reason_desc, ra_lith_scale_scheme.Remark, ra_lith_scale_scheme.Source, ra_lith_scale_scheme.Source_document, ra_lith_scale_scheme.Struckoff_date, ra_lith_scale_scheme.Sw_application_id, ra_lith_scale_scheme.Use_rule, ra_lith_scale_scheme.Row_changed_by, ra_lith_scale_scheme.Row_changed_date, ra_lith_scale_scheme.Row_created_by, ra_lith_scale_scheme.Row_effective_date, ra_lith_scale_scheme.Row_expiry_date, ra_lith_scale_scheme.Row_quality, ra_lith_scale_scheme.Scale_scheme)
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

func PatchRaLithScaleScheme(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ra_lith_scale_scheme set "
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
	query += " where scale_scheme = :id"

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

func DeleteRaLithScaleScheme(c *fiber.Ctx) error {
	id := c.Params("id")
	var ra_lith_scale_scheme dto.Ra_lith_scale_scheme
	ra_lith_scale_scheme.Scale_scheme = id

	stmt, err := db.Prepare("delete from ra_lith_scale_scheme where scale_scheme = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ra_lith_scale_scheme.Scale_scheme)
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



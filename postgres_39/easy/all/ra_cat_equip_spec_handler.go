package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRaCatEquipSpec(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ra_cat_equip_spec")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ra_cat_equip_spec

	for rows.Next() {
		var ra_cat_equip_spec dto.Ra_cat_equip_spec
		if err := rows.Scan(&ra_cat_equip_spec.Spec_type, &ra_cat_equip_spec.Alias_id, &ra_cat_equip_spec.Abbreviation, &ra_cat_equip_spec.Active_ind, &ra_cat_equip_spec.Alias_long_name, &ra_cat_equip_spec.Alias_reason_type, &ra_cat_equip_spec.Alias_short_name, &ra_cat_equip_spec.Alias_type, &ra_cat_equip_spec.Amended_date, &ra_cat_equip_spec.Created_date, &ra_cat_equip_spec.Effective_date, &ra_cat_equip_spec.Expiry_date, &ra_cat_equip_spec.Spec_type, &ra_cat_equip_spec.Original_ind, &ra_cat_equip_spec.Owner_ba_id, &ra_cat_equip_spec.Ppdm_guid, &ra_cat_equip_spec.Preferred_ind, &ra_cat_equip_spec.Reason_desc, &ra_cat_equip_spec.Remark, &ra_cat_equip_spec.Source, &ra_cat_equip_spec.Source_document, &ra_cat_equip_spec.Struckoff_date, &ra_cat_equip_spec.Sw_application_id, &ra_cat_equip_spec.Use_rule, &ra_cat_equip_spec.Row_changed_by, &ra_cat_equip_spec.Row_changed_date, &ra_cat_equip_spec.Row_created_by, &ra_cat_equip_spec.Row_created_date, &ra_cat_equip_spec.Row_effective_date, &ra_cat_equip_spec.Row_expiry_date, &ra_cat_equip_spec.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ra_cat_equip_spec to result
		result = append(result, ra_cat_equip_spec)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRaCatEquipSpec(c *fiber.Ctx) error {
	var ra_cat_equip_spec dto.Ra_cat_equip_spec

	setDefaults(&ra_cat_equip_spec)

	if err := json.Unmarshal(c.Body(), &ra_cat_equip_spec); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ra_cat_equip_spec values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31)")
	if err != nil {
		return err
	}
	ra_cat_equip_spec.Row_created_date = formatDate(ra_cat_equip_spec.Row_created_date)
	ra_cat_equip_spec.Row_changed_date = formatDate(ra_cat_equip_spec.Row_changed_date)
	ra_cat_equip_spec.Amended_date = formatDateString(ra_cat_equip_spec.Amended_date)
	ra_cat_equip_spec.Created_date = formatDateString(ra_cat_equip_spec.Created_date)
	ra_cat_equip_spec.Effective_date = formatDateString(ra_cat_equip_spec.Effective_date)
	ra_cat_equip_spec.Expiry_date = formatDateString(ra_cat_equip_spec.Expiry_date)
	ra_cat_equip_spec.Struckoff_date = formatDateString(ra_cat_equip_spec.Struckoff_date)
	ra_cat_equip_spec.Row_effective_date = formatDateString(ra_cat_equip_spec.Row_effective_date)
	ra_cat_equip_spec.Row_expiry_date = formatDateString(ra_cat_equip_spec.Row_expiry_date)






	rows, err := stmt.Exec(ra_cat_equip_spec.Spec_type, ra_cat_equip_spec.Alias_id, ra_cat_equip_spec.Abbreviation, ra_cat_equip_spec.Active_ind, ra_cat_equip_spec.Alias_long_name, ra_cat_equip_spec.Alias_reason_type, ra_cat_equip_spec.Alias_short_name, ra_cat_equip_spec.Alias_type, ra_cat_equip_spec.Amended_date, ra_cat_equip_spec.Created_date, ra_cat_equip_spec.Effective_date, ra_cat_equip_spec.Expiry_date, ra_cat_equip_spec.Spec_type, ra_cat_equip_spec.Original_ind, ra_cat_equip_spec.Owner_ba_id, ra_cat_equip_spec.Ppdm_guid, ra_cat_equip_spec.Preferred_ind, ra_cat_equip_spec.Reason_desc, ra_cat_equip_spec.Remark, ra_cat_equip_spec.Source, ra_cat_equip_spec.Source_document, ra_cat_equip_spec.Struckoff_date, ra_cat_equip_spec.Sw_application_id, ra_cat_equip_spec.Use_rule, ra_cat_equip_spec.Row_changed_by, ra_cat_equip_spec.Row_changed_date, ra_cat_equip_spec.Row_created_by, ra_cat_equip_spec.Row_created_date, ra_cat_equip_spec.Row_effective_date, ra_cat_equip_spec.Row_expiry_date, ra_cat_equip_spec.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRaCatEquipSpec(c *fiber.Ctx) error {
	var ra_cat_equip_spec dto.Ra_cat_equip_spec

	setDefaults(&ra_cat_equip_spec)

	if err := json.Unmarshal(c.Body(), &ra_cat_equip_spec); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ra_cat_equip_spec.Spec_type = id

    if ra_cat_equip_spec.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ra_cat_equip_spec set alias_id = :1, abbreviation = :2, active_ind = :3, alias_long_name = :4, alias_reason_type = :5, alias_short_name = :6, alias_type = :7, amended_date = :8, created_date = :9, effective_date = :10, expiry_date = :11, spec_type = :12, original_ind = :13, owner_ba_id = :14, ppdm_guid = :15, preferred_ind = :16, reason_desc = :17, remark = :18, source = :19, source_document = :20, struckoff_date = :21, sw_application_id = :22, use_rule = :23, row_changed_by = :24, row_changed_date = :25, row_created_by = :26, row_effective_date = :27, row_expiry_date = :28, row_quality = :29 where spec_type = :31")
	if err != nil {
		return err
	}

	ra_cat_equip_spec.Row_changed_date = formatDate(ra_cat_equip_spec.Row_changed_date)
	ra_cat_equip_spec.Amended_date = formatDateString(ra_cat_equip_spec.Amended_date)
	ra_cat_equip_spec.Created_date = formatDateString(ra_cat_equip_spec.Created_date)
	ra_cat_equip_spec.Effective_date = formatDateString(ra_cat_equip_spec.Effective_date)
	ra_cat_equip_spec.Expiry_date = formatDateString(ra_cat_equip_spec.Expiry_date)
	ra_cat_equip_spec.Struckoff_date = formatDateString(ra_cat_equip_spec.Struckoff_date)
	ra_cat_equip_spec.Row_effective_date = formatDateString(ra_cat_equip_spec.Row_effective_date)
	ra_cat_equip_spec.Row_expiry_date = formatDateString(ra_cat_equip_spec.Row_expiry_date)






	rows, err := stmt.Exec(ra_cat_equip_spec.Alias_id, ra_cat_equip_spec.Abbreviation, ra_cat_equip_spec.Active_ind, ra_cat_equip_spec.Alias_long_name, ra_cat_equip_spec.Alias_reason_type, ra_cat_equip_spec.Alias_short_name, ra_cat_equip_spec.Alias_type, ra_cat_equip_spec.Amended_date, ra_cat_equip_spec.Created_date, ra_cat_equip_spec.Effective_date, ra_cat_equip_spec.Expiry_date, ra_cat_equip_spec.Spec_type, ra_cat_equip_spec.Original_ind, ra_cat_equip_spec.Owner_ba_id, ra_cat_equip_spec.Ppdm_guid, ra_cat_equip_spec.Preferred_ind, ra_cat_equip_spec.Reason_desc, ra_cat_equip_spec.Remark, ra_cat_equip_spec.Source, ra_cat_equip_spec.Source_document, ra_cat_equip_spec.Struckoff_date, ra_cat_equip_spec.Sw_application_id, ra_cat_equip_spec.Use_rule, ra_cat_equip_spec.Row_changed_by, ra_cat_equip_spec.Row_changed_date, ra_cat_equip_spec.Row_created_by, ra_cat_equip_spec.Row_effective_date, ra_cat_equip_spec.Row_expiry_date, ra_cat_equip_spec.Row_quality, ra_cat_equip_spec.Spec_type)
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

func PatchRaCatEquipSpec(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ra_cat_equip_spec set "
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
	query += " where spec_type = :id"

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

func DeleteRaCatEquipSpec(c *fiber.Ctx) error {
	id := c.Params("id")
	var ra_cat_equip_spec dto.Ra_cat_equip_spec
	ra_cat_equip_spec.Spec_type = id

	stmt, err := db.Prepare("delete from ra_cat_equip_spec where spec_type = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ra_cat_equip_spec.Spec_type)
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



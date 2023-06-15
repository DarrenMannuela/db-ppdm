package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLandAlias(c *fiber.Ctx) error {
	rows, err := db.Query("select * from land_alias")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Land_alias

	for rows.Next() {
		var land_alias dto.Land_alias
		if err := rows.Scan(&land_alias.Land_right_subtype, &land_alias.Land_right_id, &land_alias.Lr_alias_id, &land_alias.Abbreviation, &land_alias.Active_ind, &land_alias.Alias_long_name, &land_alias.Alias_number, &land_alias.Alias_reason_type, &land_alias.Alias_short_name, &land_alias.Alias_type, &land_alias.Amended_date, &land_alias.Created_date, &land_alias.Effective_date, &land_alias.Expiry_date, &land_alias.Land_ref_num_type, &land_alias.Land_right_subtype, &land_alias.Original_ind, &land_alias.Owner_ba_id, &land_alias.Ppdm_guid, &land_alias.Preferred_ind, &land_alias.Reason_desc, &land_alias.Reference_num, &land_alias.Remark, &land_alias.Source, &land_alias.Source_document_id, &land_alias.Struckoff_date, &land_alias.Sw_application_id, &land_alias.Use_rule, &land_alias.Row_changed_by, &land_alias.Row_changed_date, &land_alias.Row_created_by, &land_alias.Row_created_date, &land_alias.Row_effective_date, &land_alias.Row_expiry_date, &land_alias.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append land_alias to result
		result = append(result, land_alias)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLandAlias(c *fiber.Ctx) error {
	var land_alias dto.Land_alias

	setDefaults(&land_alias)

	if err := json.Unmarshal(c.Body(), &land_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into land_alias values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35)")
	if err != nil {
		return err
	}
	land_alias.Row_created_date = formatDate(land_alias.Row_created_date)
	land_alias.Row_changed_date = formatDate(land_alias.Row_changed_date)
	land_alias.Amended_date = formatDateString(land_alias.Amended_date)
	land_alias.Created_date = formatDateString(land_alias.Created_date)
	land_alias.Effective_date = formatDateString(land_alias.Effective_date)
	land_alias.Expiry_date = formatDateString(land_alias.Expiry_date)
	land_alias.Struckoff_date = formatDateString(land_alias.Struckoff_date)
	land_alias.Row_effective_date = formatDateString(land_alias.Row_effective_date)
	land_alias.Row_expiry_date = formatDateString(land_alias.Row_expiry_date)






	rows, err := stmt.Exec(land_alias.Land_right_subtype, land_alias.Land_right_id, land_alias.Lr_alias_id, land_alias.Abbreviation, land_alias.Active_ind, land_alias.Alias_long_name, land_alias.Alias_number, land_alias.Alias_reason_type, land_alias.Alias_short_name, land_alias.Alias_type, land_alias.Amended_date, land_alias.Created_date, land_alias.Effective_date, land_alias.Expiry_date, land_alias.Land_ref_num_type, land_alias.Land_right_subtype, land_alias.Original_ind, land_alias.Owner_ba_id, land_alias.Ppdm_guid, land_alias.Preferred_ind, land_alias.Reason_desc, land_alias.Reference_num, land_alias.Remark, land_alias.Source, land_alias.Source_document_id, land_alias.Struckoff_date, land_alias.Sw_application_id, land_alias.Use_rule, land_alias.Row_changed_by, land_alias.Row_changed_date, land_alias.Row_created_by, land_alias.Row_created_date, land_alias.Row_effective_date, land_alias.Row_expiry_date, land_alias.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLandAlias(c *fiber.Ctx) error {
	var land_alias dto.Land_alias

	setDefaults(&land_alias)

	if err := json.Unmarshal(c.Body(), &land_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	land_alias.Land_right_subtype = id

    if land_alias.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update land_alias set land_right_id = :1, lr_alias_id = :2, abbreviation = :3, active_ind = :4, alias_long_name = :5, alias_number = :6, alias_reason_type = :7, alias_short_name = :8, alias_type = :9, amended_date = :10, created_date = :11, effective_date = :12, expiry_date = :13, land_ref_num_type = :14, land_right_subtype = :15, original_ind = :16, owner_ba_id = :17, ppdm_guid = :18, preferred_ind = :19, reason_desc = :20, reference_num = :21, remark = :22, source = :23, source_document_id = :24, struckoff_date = :25, sw_application_id = :26, use_rule = :27, row_changed_by = :28, row_changed_date = :29, row_created_by = :30, row_effective_date = :31, row_expiry_date = :32, row_quality = :33 where land_right_subtype = :35")
	if err != nil {
		return err
	}

	land_alias.Row_changed_date = formatDate(land_alias.Row_changed_date)
	land_alias.Amended_date = formatDateString(land_alias.Amended_date)
	land_alias.Created_date = formatDateString(land_alias.Created_date)
	land_alias.Effective_date = formatDateString(land_alias.Effective_date)
	land_alias.Expiry_date = formatDateString(land_alias.Expiry_date)
	land_alias.Struckoff_date = formatDateString(land_alias.Struckoff_date)
	land_alias.Row_effective_date = formatDateString(land_alias.Row_effective_date)
	land_alias.Row_expiry_date = formatDateString(land_alias.Row_expiry_date)






	rows, err := stmt.Exec(land_alias.Land_right_id, land_alias.Lr_alias_id, land_alias.Abbreviation, land_alias.Active_ind, land_alias.Alias_long_name, land_alias.Alias_number, land_alias.Alias_reason_type, land_alias.Alias_short_name, land_alias.Alias_type, land_alias.Amended_date, land_alias.Created_date, land_alias.Effective_date, land_alias.Expiry_date, land_alias.Land_ref_num_type, land_alias.Land_right_subtype, land_alias.Original_ind, land_alias.Owner_ba_id, land_alias.Ppdm_guid, land_alias.Preferred_ind, land_alias.Reason_desc, land_alias.Reference_num, land_alias.Remark, land_alias.Source, land_alias.Source_document_id, land_alias.Struckoff_date, land_alias.Sw_application_id, land_alias.Use_rule, land_alias.Row_changed_by, land_alias.Row_changed_date, land_alias.Row_created_by, land_alias.Row_effective_date, land_alias.Row_expiry_date, land_alias.Row_quality, land_alias.Land_right_subtype)
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

func PatchLandAlias(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update land_alias set "
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
	query += " where land_right_subtype = :id"

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

func DeleteLandAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var land_alias dto.Land_alias
	land_alias.Land_right_subtype = id

	stmt, err := db.Prepare("delete from land_alias where land_right_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(land_alias.Land_right_subtype)
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



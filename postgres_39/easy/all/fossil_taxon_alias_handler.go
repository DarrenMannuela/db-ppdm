package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFossilTaxonAlias(c *fiber.Ctx) error {
	rows, err := db.Query("select * from fossil_taxon_alias")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Fossil_taxon_alias

	for rows.Next() {
		var fossil_taxon_alias dto.Fossil_taxon_alias
		if err := rows.Scan(&fossil_taxon_alias.Taxon_leaf_id, &fossil_taxon_alias.Fossil_alias_id, &fossil_taxon_alias.Abbreviation, &fossil_taxon_alias.Active_ind, &fossil_taxon_alias.Alias_fossil_id, &fossil_taxon_alias.Alias_long_name, &fossil_taxon_alias.Alias_reason_type, &fossil_taxon_alias.Alias_short_name, &fossil_taxon_alias.Alias_type, &fossil_taxon_alias.Amended_date, &fossil_taxon_alias.Created_date, &fossil_taxon_alias.Effective_date, &fossil_taxon_alias.Expiry_date, &fossil_taxon_alias.Taxon_leaf_id, &fossil_taxon_alias.Original_ind, &fossil_taxon_alias.Owner_ba_id, &fossil_taxon_alias.Ppdm_guid, &fossil_taxon_alias.Preferred_ind, &fossil_taxon_alias.Reason_desc, &fossil_taxon_alias.Remark, &fossil_taxon_alias.Source, &fossil_taxon_alias.Source_document_id, &fossil_taxon_alias.Struckoff_date, &fossil_taxon_alias.Sw_application_id, &fossil_taxon_alias.Use_rule, &fossil_taxon_alias.Row_changed_by, &fossil_taxon_alias.Row_changed_date, &fossil_taxon_alias.Row_created_by, &fossil_taxon_alias.Row_created_date, &fossil_taxon_alias.Row_effective_date, &fossil_taxon_alias.Row_expiry_date, &fossil_taxon_alias.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append fossil_taxon_alias to result
		result = append(result, fossil_taxon_alias)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFossilTaxonAlias(c *fiber.Ctx) error {
	var fossil_taxon_alias dto.Fossil_taxon_alias

	setDefaults(&fossil_taxon_alias)

	if err := json.Unmarshal(c.Body(), &fossil_taxon_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into fossil_taxon_alias values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32)")
	if err != nil {
		return err
	}
	fossil_taxon_alias.Row_created_date = formatDate(fossil_taxon_alias.Row_created_date)
	fossil_taxon_alias.Row_changed_date = formatDate(fossil_taxon_alias.Row_changed_date)
	fossil_taxon_alias.Amended_date = formatDateString(fossil_taxon_alias.Amended_date)
	fossil_taxon_alias.Created_date = formatDateString(fossil_taxon_alias.Created_date)
	fossil_taxon_alias.Effective_date = formatDateString(fossil_taxon_alias.Effective_date)
	fossil_taxon_alias.Expiry_date = formatDateString(fossil_taxon_alias.Expiry_date)
	fossil_taxon_alias.Struckoff_date = formatDateString(fossil_taxon_alias.Struckoff_date)
	fossil_taxon_alias.Row_effective_date = formatDateString(fossil_taxon_alias.Row_effective_date)
	fossil_taxon_alias.Row_expiry_date = formatDateString(fossil_taxon_alias.Row_expiry_date)






	rows, err := stmt.Exec(fossil_taxon_alias.Taxon_leaf_id, fossil_taxon_alias.Fossil_alias_id, fossil_taxon_alias.Abbreviation, fossil_taxon_alias.Active_ind, fossil_taxon_alias.Alias_fossil_id, fossil_taxon_alias.Alias_long_name, fossil_taxon_alias.Alias_reason_type, fossil_taxon_alias.Alias_short_name, fossil_taxon_alias.Alias_type, fossil_taxon_alias.Amended_date, fossil_taxon_alias.Created_date, fossil_taxon_alias.Effective_date, fossil_taxon_alias.Expiry_date, fossil_taxon_alias.Taxon_leaf_id, fossil_taxon_alias.Original_ind, fossil_taxon_alias.Owner_ba_id, fossil_taxon_alias.Ppdm_guid, fossil_taxon_alias.Preferred_ind, fossil_taxon_alias.Reason_desc, fossil_taxon_alias.Remark, fossil_taxon_alias.Source, fossil_taxon_alias.Source_document_id, fossil_taxon_alias.Struckoff_date, fossil_taxon_alias.Sw_application_id, fossil_taxon_alias.Use_rule, fossil_taxon_alias.Row_changed_by, fossil_taxon_alias.Row_changed_date, fossil_taxon_alias.Row_created_by, fossil_taxon_alias.Row_created_date, fossil_taxon_alias.Row_effective_date, fossil_taxon_alias.Row_expiry_date, fossil_taxon_alias.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFossilTaxonAlias(c *fiber.Ctx) error {
	var fossil_taxon_alias dto.Fossil_taxon_alias

	setDefaults(&fossil_taxon_alias)

	if err := json.Unmarshal(c.Body(), &fossil_taxon_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	fossil_taxon_alias.Taxon_leaf_id = id

    if fossil_taxon_alias.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update fossil_taxon_alias set fossil_alias_id = :1, abbreviation = :2, active_ind = :3, alias_fossil_id = :4, alias_long_name = :5, alias_reason_type = :6, alias_short_name = :7, alias_type = :8, amended_date = :9, created_date = :10, effective_date = :11, expiry_date = :12, taxon_leaf_id = :13, original_ind = :14, owner_ba_id = :15, ppdm_guid = :16, preferred_ind = :17, reason_desc = :18, remark = :19, source = :20, source_document_id = :21, struckoff_date = :22, sw_application_id = :23, use_rule = :24, row_changed_by = :25, row_changed_date = :26, row_created_by = :27, row_effective_date = :28, row_expiry_date = :29, row_quality = :30 where taxon_leaf_id = :32")
	if err != nil {
		return err
	}

	fossil_taxon_alias.Row_changed_date = formatDate(fossil_taxon_alias.Row_changed_date)
	fossil_taxon_alias.Amended_date = formatDateString(fossil_taxon_alias.Amended_date)
	fossil_taxon_alias.Created_date = formatDateString(fossil_taxon_alias.Created_date)
	fossil_taxon_alias.Effective_date = formatDateString(fossil_taxon_alias.Effective_date)
	fossil_taxon_alias.Expiry_date = formatDateString(fossil_taxon_alias.Expiry_date)
	fossil_taxon_alias.Struckoff_date = formatDateString(fossil_taxon_alias.Struckoff_date)
	fossil_taxon_alias.Row_effective_date = formatDateString(fossil_taxon_alias.Row_effective_date)
	fossil_taxon_alias.Row_expiry_date = formatDateString(fossil_taxon_alias.Row_expiry_date)






	rows, err := stmt.Exec(fossil_taxon_alias.Fossil_alias_id, fossil_taxon_alias.Abbreviation, fossil_taxon_alias.Active_ind, fossil_taxon_alias.Alias_fossil_id, fossil_taxon_alias.Alias_long_name, fossil_taxon_alias.Alias_reason_type, fossil_taxon_alias.Alias_short_name, fossil_taxon_alias.Alias_type, fossil_taxon_alias.Amended_date, fossil_taxon_alias.Created_date, fossil_taxon_alias.Effective_date, fossil_taxon_alias.Expiry_date, fossil_taxon_alias.Taxon_leaf_id, fossil_taxon_alias.Original_ind, fossil_taxon_alias.Owner_ba_id, fossil_taxon_alias.Ppdm_guid, fossil_taxon_alias.Preferred_ind, fossil_taxon_alias.Reason_desc, fossil_taxon_alias.Remark, fossil_taxon_alias.Source, fossil_taxon_alias.Source_document_id, fossil_taxon_alias.Struckoff_date, fossil_taxon_alias.Sw_application_id, fossil_taxon_alias.Use_rule, fossil_taxon_alias.Row_changed_by, fossil_taxon_alias.Row_changed_date, fossil_taxon_alias.Row_created_by, fossil_taxon_alias.Row_effective_date, fossil_taxon_alias.Row_expiry_date, fossil_taxon_alias.Row_quality, fossil_taxon_alias.Taxon_leaf_id)
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

func PatchFossilTaxonAlias(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update fossil_taxon_alias set "
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
	query += " where taxon_leaf_id = :id"

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

func DeleteFossilTaxonAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var fossil_taxon_alias dto.Fossil_taxon_alias
	fossil_taxon_alias.Taxon_leaf_id = id

	stmt, err := db.Prepare("delete from fossil_taxon_alias where taxon_leaf_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(fossil_taxon_alias.Taxon_leaf_id)
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



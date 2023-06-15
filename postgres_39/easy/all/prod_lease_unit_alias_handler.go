package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetProdLeaseUnitAlias(c *fiber.Ctx) error {
	rows, err := db.Query("select * from prod_lease_unit_alias")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Prod_lease_unit_alias

	for rows.Next() {
		var prod_lease_unit_alias dto.Prod_lease_unit_alias
		if err := rows.Scan(&prod_lease_unit_alias.Lease_unit_id, &prod_lease_unit_alias.Lease_unit_source, &prod_lease_unit_alias.Prod_lease_unit_alias_id, &prod_lease_unit_alias.Alias_source, &prod_lease_unit_alias.Abbreviation, &prod_lease_unit_alias.Active_ind, &prod_lease_unit_alias.Alias_long_name, &prod_lease_unit_alias.Alias_reason_type, &prod_lease_unit_alias.Alias_short_name, &prod_lease_unit_alias.Alias_type, &prod_lease_unit_alias.Amended_date, &prod_lease_unit_alias.Created_date, &prod_lease_unit_alias.Effective_date, &prod_lease_unit_alias.Expiry_date, &prod_lease_unit_alias.Lease_unit_id, &prod_lease_unit_alias.Original_ind, &prod_lease_unit_alias.Owner_ba_id, &prod_lease_unit_alias.Ppdm_guid, &prod_lease_unit_alias.Preferred_ind, &prod_lease_unit_alias.Reason_desc, &prod_lease_unit_alias.Remark, &prod_lease_unit_alias.Source_document_id, &prod_lease_unit_alias.Struckoff_date, &prod_lease_unit_alias.Sw_application_id, &prod_lease_unit_alias.Use_rule, &prod_lease_unit_alias.Row_changed_by, &prod_lease_unit_alias.Row_changed_date, &prod_lease_unit_alias.Row_created_by, &prod_lease_unit_alias.Row_created_date, &prod_lease_unit_alias.Row_effective_date, &prod_lease_unit_alias.Row_expiry_date, &prod_lease_unit_alias.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append prod_lease_unit_alias to result
		result = append(result, prod_lease_unit_alias)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetProdLeaseUnitAlias(c *fiber.Ctx) error {
	var prod_lease_unit_alias dto.Prod_lease_unit_alias

	setDefaults(&prod_lease_unit_alias)

	if err := json.Unmarshal(c.Body(), &prod_lease_unit_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into prod_lease_unit_alias values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32)")
	if err != nil {
		return err
	}
	prod_lease_unit_alias.Row_created_date = formatDate(prod_lease_unit_alias.Row_created_date)
	prod_lease_unit_alias.Row_changed_date = formatDate(prod_lease_unit_alias.Row_changed_date)
	prod_lease_unit_alias.Amended_date = formatDateString(prod_lease_unit_alias.Amended_date)
	prod_lease_unit_alias.Created_date = formatDateString(prod_lease_unit_alias.Created_date)
	prod_lease_unit_alias.Effective_date = formatDateString(prod_lease_unit_alias.Effective_date)
	prod_lease_unit_alias.Expiry_date = formatDateString(prod_lease_unit_alias.Expiry_date)
	prod_lease_unit_alias.Struckoff_date = formatDateString(prod_lease_unit_alias.Struckoff_date)
	prod_lease_unit_alias.Row_effective_date = formatDateString(prod_lease_unit_alias.Row_effective_date)
	prod_lease_unit_alias.Row_expiry_date = formatDateString(prod_lease_unit_alias.Row_expiry_date)






	rows, err := stmt.Exec(prod_lease_unit_alias.Lease_unit_id, prod_lease_unit_alias.Lease_unit_source, prod_lease_unit_alias.Prod_lease_unit_alias_id, prod_lease_unit_alias.Alias_source, prod_lease_unit_alias.Abbreviation, prod_lease_unit_alias.Active_ind, prod_lease_unit_alias.Alias_long_name, prod_lease_unit_alias.Alias_reason_type, prod_lease_unit_alias.Alias_short_name, prod_lease_unit_alias.Alias_type, prod_lease_unit_alias.Amended_date, prod_lease_unit_alias.Created_date, prod_lease_unit_alias.Effective_date, prod_lease_unit_alias.Expiry_date, prod_lease_unit_alias.Lease_unit_id, prod_lease_unit_alias.Original_ind, prod_lease_unit_alias.Owner_ba_id, prod_lease_unit_alias.Ppdm_guid, prod_lease_unit_alias.Preferred_ind, prod_lease_unit_alias.Reason_desc, prod_lease_unit_alias.Remark, prod_lease_unit_alias.Source_document_id, prod_lease_unit_alias.Struckoff_date, prod_lease_unit_alias.Sw_application_id, prod_lease_unit_alias.Use_rule, prod_lease_unit_alias.Row_changed_by, prod_lease_unit_alias.Row_changed_date, prod_lease_unit_alias.Row_created_by, prod_lease_unit_alias.Row_created_date, prod_lease_unit_alias.Row_effective_date, prod_lease_unit_alias.Row_expiry_date, prod_lease_unit_alias.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateProdLeaseUnitAlias(c *fiber.Ctx) error {
	var prod_lease_unit_alias dto.Prod_lease_unit_alias

	setDefaults(&prod_lease_unit_alias)

	if err := json.Unmarshal(c.Body(), &prod_lease_unit_alias); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	prod_lease_unit_alias.Lease_unit_id = id

    if prod_lease_unit_alias.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update prod_lease_unit_alias set lease_unit_source = :1, prod_lease_unit_alias_id = :2, alias_source = :3, abbreviation = :4, active_ind = :5, alias_long_name = :6, alias_reason_type = :7, alias_short_name = :8, alias_type = :9, amended_date = :10, created_date = :11, effective_date = :12, expiry_date = :13, lease_unit_id = :14, original_ind = :15, owner_ba_id = :16, ppdm_guid = :17, preferred_ind = :18, reason_desc = :19, remark = :20, source_document_id = :21, struckoff_date = :22, sw_application_id = :23, use_rule = :24, row_changed_by = :25, row_changed_date = :26, row_created_by = :27, row_effective_date = :28, row_expiry_date = :29, row_quality = :30 where lease_unit_id = :32")
	if err != nil {
		return err
	}

	prod_lease_unit_alias.Row_changed_date = formatDate(prod_lease_unit_alias.Row_changed_date)
	prod_lease_unit_alias.Amended_date = formatDateString(prod_lease_unit_alias.Amended_date)
	prod_lease_unit_alias.Created_date = formatDateString(prod_lease_unit_alias.Created_date)
	prod_lease_unit_alias.Effective_date = formatDateString(prod_lease_unit_alias.Effective_date)
	prod_lease_unit_alias.Expiry_date = formatDateString(prod_lease_unit_alias.Expiry_date)
	prod_lease_unit_alias.Struckoff_date = formatDateString(prod_lease_unit_alias.Struckoff_date)
	prod_lease_unit_alias.Row_effective_date = formatDateString(prod_lease_unit_alias.Row_effective_date)
	prod_lease_unit_alias.Row_expiry_date = formatDateString(prod_lease_unit_alias.Row_expiry_date)






	rows, err := stmt.Exec(prod_lease_unit_alias.Lease_unit_source, prod_lease_unit_alias.Prod_lease_unit_alias_id, prod_lease_unit_alias.Alias_source, prod_lease_unit_alias.Abbreviation, prod_lease_unit_alias.Active_ind, prod_lease_unit_alias.Alias_long_name, prod_lease_unit_alias.Alias_reason_type, prod_lease_unit_alias.Alias_short_name, prod_lease_unit_alias.Alias_type, prod_lease_unit_alias.Amended_date, prod_lease_unit_alias.Created_date, prod_lease_unit_alias.Effective_date, prod_lease_unit_alias.Expiry_date, prod_lease_unit_alias.Lease_unit_id, prod_lease_unit_alias.Original_ind, prod_lease_unit_alias.Owner_ba_id, prod_lease_unit_alias.Ppdm_guid, prod_lease_unit_alias.Preferred_ind, prod_lease_unit_alias.Reason_desc, prod_lease_unit_alias.Remark, prod_lease_unit_alias.Source_document_id, prod_lease_unit_alias.Struckoff_date, prod_lease_unit_alias.Sw_application_id, prod_lease_unit_alias.Use_rule, prod_lease_unit_alias.Row_changed_by, prod_lease_unit_alias.Row_changed_date, prod_lease_unit_alias.Row_created_by, prod_lease_unit_alias.Row_effective_date, prod_lease_unit_alias.Row_expiry_date, prod_lease_unit_alias.Row_quality, prod_lease_unit_alias.Lease_unit_id)
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

func PatchProdLeaseUnitAlias(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update prod_lease_unit_alias set "
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
	query += " where lease_unit_id = :id"

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

func DeleteProdLeaseUnitAlias(c *fiber.Ctx) error {
	id := c.Params("id")
	var prod_lease_unit_alias dto.Prod_lease_unit_alias
	prod_lease_unit_alias.Lease_unit_id = id

	stmt, err := db.Prepare("delete from prod_lease_unit_alias where lease_unit_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(prod_lease_unit_alias.Lease_unit_id)
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



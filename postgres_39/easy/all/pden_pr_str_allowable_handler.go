package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPdenPrStrAllowable(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pden_pr_str_allowable")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pden_pr_str_allowable

	for rows.Next() {
		var pden_pr_str_allowable dto.Pden_pr_str_allowable
		if err := rows.Scan(&pden_pr_str_allowable.Pden_subtype, &pden_pr_str_allowable.Pden_id, &pden_pr_str_allowable.Pden_source, &pden_pr_str_allowable.Uwi, &pden_pr_str_allowable.String_source, &pden_pr_str_allowable.String_id, &pden_pr_str_allowable.Pden_prs_xref_seq_no, &pden_pr_str_allowable.Allowable_obs_no, &pden_pr_str_allowable.Active_ind, &pden_pr_str_allowable.Allowable, &pden_pr_str_allowable.Allowable_date, &pden_pr_str_allowable.Allowable_date_desc, &pden_pr_str_allowable.Allowable_ouom, &pden_pr_str_allowable.Allowable_uom, &pden_pr_str_allowable.Effective_date, &pden_pr_str_allowable.Expiry_date, &pden_pr_str_allowable.Period_type, &pden_pr_str_allowable.Ppdm_guid, &pden_pr_str_allowable.Product_type, &pden_pr_str_allowable.Remark, &pden_pr_str_allowable.Source, &pden_pr_str_allowable.Row_changed_by, &pden_pr_str_allowable.Row_changed_date, &pden_pr_str_allowable.Row_created_by, &pden_pr_str_allowable.Row_created_date, &pden_pr_str_allowable.Row_effective_date, &pden_pr_str_allowable.Row_expiry_date, &pden_pr_str_allowable.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pden_pr_str_allowable to result
		result = append(result, pden_pr_str_allowable)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPdenPrStrAllowable(c *fiber.Ctx) error {
	var pden_pr_str_allowable dto.Pden_pr_str_allowable

	setDefaults(&pden_pr_str_allowable)

	if err := json.Unmarshal(c.Body(), &pden_pr_str_allowable); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pden_pr_str_allowable values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28)")
	if err != nil {
		return err
	}
	pden_pr_str_allowable.Row_created_date = formatDate(pden_pr_str_allowable.Row_created_date)
	pden_pr_str_allowable.Row_changed_date = formatDate(pden_pr_str_allowable.Row_changed_date)
	pden_pr_str_allowable.Allowable_date = formatDateString(pden_pr_str_allowable.Allowable_date)
	pden_pr_str_allowable.Effective_date = formatDateString(pden_pr_str_allowable.Effective_date)
	pden_pr_str_allowable.Expiry_date = formatDateString(pden_pr_str_allowable.Expiry_date)
	pden_pr_str_allowable.Row_effective_date = formatDateString(pden_pr_str_allowable.Row_effective_date)
	pden_pr_str_allowable.Row_expiry_date = formatDateString(pden_pr_str_allowable.Row_expiry_date)






	rows, err := stmt.Exec(pden_pr_str_allowable.Pden_subtype, pden_pr_str_allowable.Pden_id, pden_pr_str_allowable.Pden_source, pden_pr_str_allowable.Uwi, pden_pr_str_allowable.String_source, pden_pr_str_allowable.String_id, pden_pr_str_allowable.Pden_prs_xref_seq_no, pden_pr_str_allowable.Allowable_obs_no, pden_pr_str_allowable.Active_ind, pden_pr_str_allowable.Allowable, pden_pr_str_allowable.Allowable_date, pden_pr_str_allowable.Allowable_date_desc, pden_pr_str_allowable.Allowable_ouom, pden_pr_str_allowable.Allowable_uom, pden_pr_str_allowable.Effective_date, pden_pr_str_allowable.Expiry_date, pden_pr_str_allowable.Period_type, pden_pr_str_allowable.Ppdm_guid, pden_pr_str_allowable.Product_type, pden_pr_str_allowable.Remark, pden_pr_str_allowable.Source, pden_pr_str_allowable.Row_changed_by, pden_pr_str_allowable.Row_changed_date, pden_pr_str_allowable.Row_created_by, pden_pr_str_allowable.Row_created_date, pden_pr_str_allowable.Row_effective_date, pden_pr_str_allowable.Row_expiry_date, pden_pr_str_allowable.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePdenPrStrAllowable(c *fiber.Ctx) error {
	var pden_pr_str_allowable dto.Pden_pr_str_allowable

	setDefaults(&pden_pr_str_allowable)

	if err := json.Unmarshal(c.Body(), &pden_pr_str_allowable); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pden_pr_str_allowable.Pden_subtype = id

    if pden_pr_str_allowable.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pden_pr_str_allowable set pden_id = :1, pden_source = :2, uwi = :3, string_source = :4, string_id = :5, pden_prs_xref_seq_no = :6, allowable_obs_no = :7, active_ind = :8, allowable = :9, allowable_date = :10, allowable_date_desc = :11, allowable_ouom = :12, allowable_uom = :13, effective_date = :14, expiry_date = :15, period_type = :16, ppdm_guid = :17, product_type = :18, remark = :19, source = :20, row_changed_by = :21, row_changed_date = :22, row_created_by = :23, row_effective_date = :24, row_expiry_date = :25, row_quality = :26 where pden_subtype = :28")
	if err != nil {
		return err
	}

	pden_pr_str_allowable.Row_changed_date = formatDate(pden_pr_str_allowable.Row_changed_date)
	pden_pr_str_allowable.Allowable_date = formatDateString(pden_pr_str_allowable.Allowable_date)
	pden_pr_str_allowable.Effective_date = formatDateString(pden_pr_str_allowable.Effective_date)
	pden_pr_str_allowable.Expiry_date = formatDateString(pden_pr_str_allowable.Expiry_date)
	pden_pr_str_allowable.Row_effective_date = formatDateString(pden_pr_str_allowable.Row_effective_date)
	pden_pr_str_allowable.Row_expiry_date = formatDateString(pden_pr_str_allowable.Row_expiry_date)






	rows, err := stmt.Exec(pden_pr_str_allowable.Pden_id, pden_pr_str_allowable.Pden_source, pden_pr_str_allowable.Uwi, pden_pr_str_allowable.String_source, pden_pr_str_allowable.String_id, pden_pr_str_allowable.Pden_prs_xref_seq_no, pden_pr_str_allowable.Allowable_obs_no, pden_pr_str_allowable.Active_ind, pden_pr_str_allowable.Allowable, pden_pr_str_allowable.Allowable_date, pden_pr_str_allowable.Allowable_date_desc, pden_pr_str_allowable.Allowable_ouom, pden_pr_str_allowable.Allowable_uom, pden_pr_str_allowable.Effective_date, pden_pr_str_allowable.Expiry_date, pden_pr_str_allowable.Period_type, pden_pr_str_allowable.Ppdm_guid, pden_pr_str_allowable.Product_type, pden_pr_str_allowable.Remark, pden_pr_str_allowable.Source, pden_pr_str_allowable.Row_changed_by, pden_pr_str_allowable.Row_changed_date, pden_pr_str_allowable.Row_created_by, pden_pr_str_allowable.Row_effective_date, pden_pr_str_allowable.Row_expiry_date, pden_pr_str_allowable.Row_quality, pden_pr_str_allowable.Pden_subtype)
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

func PatchPdenPrStrAllowable(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pden_pr_str_allowable set "
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
		} else if key == "allowable_date" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where pden_subtype = :id"

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

func DeletePdenPrStrAllowable(c *fiber.Ctx) error {
	id := c.Params("id")
	var pden_pr_str_allowable dto.Pden_pr_str_allowable
	pden_pr_str_allowable.Pden_subtype = id

	stmt, err := db.Prepare("delete from pden_pr_str_allowable where pden_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pden_pr_str_allowable.Pden_subtype)
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



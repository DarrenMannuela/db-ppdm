package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPdenPrStrForm(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pden_pr_str_form")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pden_pr_str_form

	for rows.Next() {
		var pden_pr_str_form dto.Pden_pr_str_form
		if err := rows.Scan(&pden_pr_str_form.Pden_subtype, &pden_pr_str_form.Pden_id, &pden_pr_str_form.Pden_source, &pden_pr_str_form.Active_ind, &pden_pr_str_form.Effective_date, &pden_pr_str_form.Expiry_date, &pden_pr_str_form.Ppdm_guid, &pden_pr_str_form.Pr_str_form_obs_no, &pden_pr_str_form.Remark, &pden_pr_str_form.Source, &pden_pr_str_form.Strat_name_set_id, &pden_pr_str_form.Strat_unit_id, &pden_pr_str_form.String_id, &pden_pr_str_form.String_source, &pden_pr_str_form.Uwi, &pden_pr_str_form.Row_changed_by, &pden_pr_str_form.Row_changed_date, &pden_pr_str_form.Row_created_by, &pden_pr_str_form.Row_created_date, &pden_pr_str_form.Row_effective_date, &pden_pr_str_form.Row_expiry_date, &pden_pr_str_form.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pden_pr_str_form to result
		result = append(result, pden_pr_str_form)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPdenPrStrForm(c *fiber.Ctx) error {
	var pden_pr_str_form dto.Pden_pr_str_form

	setDefaults(&pden_pr_str_form)

	if err := json.Unmarshal(c.Body(), &pden_pr_str_form); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pden_pr_str_form values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	pden_pr_str_form.Row_created_date = formatDate(pden_pr_str_form.Row_created_date)
	pden_pr_str_form.Row_changed_date = formatDate(pden_pr_str_form.Row_changed_date)
	pden_pr_str_form.Effective_date = formatDateString(pden_pr_str_form.Effective_date)
	pden_pr_str_form.Expiry_date = formatDateString(pden_pr_str_form.Expiry_date)
	pden_pr_str_form.Row_effective_date = formatDateString(pden_pr_str_form.Row_effective_date)
	pden_pr_str_form.Row_expiry_date = formatDateString(pden_pr_str_form.Row_expiry_date)






	rows, err := stmt.Exec(pden_pr_str_form.Pden_subtype, pden_pr_str_form.Pden_id, pden_pr_str_form.Pden_source, pden_pr_str_form.Active_ind, pden_pr_str_form.Effective_date, pden_pr_str_form.Expiry_date, pden_pr_str_form.Ppdm_guid, pden_pr_str_form.Pr_str_form_obs_no, pden_pr_str_form.Remark, pden_pr_str_form.Source, pden_pr_str_form.Strat_name_set_id, pden_pr_str_form.Strat_unit_id, pden_pr_str_form.String_id, pden_pr_str_form.String_source, pden_pr_str_form.Uwi, pden_pr_str_form.Row_changed_by, pden_pr_str_form.Row_changed_date, pden_pr_str_form.Row_created_by, pden_pr_str_form.Row_created_date, pden_pr_str_form.Row_effective_date, pden_pr_str_form.Row_expiry_date, pden_pr_str_form.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePdenPrStrForm(c *fiber.Ctx) error {
	var pden_pr_str_form dto.Pden_pr_str_form

	setDefaults(&pden_pr_str_form)

	if err := json.Unmarshal(c.Body(), &pden_pr_str_form); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pden_pr_str_form.Pden_subtype = id

    if pden_pr_str_form.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pden_pr_str_form set pden_id = :1, pden_source = :2, active_ind = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, pr_str_form_obs_no = :7, remark = :8, source = :9, strat_name_set_id = :10, strat_unit_id = :11, string_id = :12, string_source = :13, uwi = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where pden_subtype = :22")
	if err != nil {
		return err
	}

	pden_pr_str_form.Row_changed_date = formatDate(pden_pr_str_form.Row_changed_date)
	pden_pr_str_form.Effective_date = formatDateString(pden_pr_str_form.Effective_date)
	pden_pr_str_form.Expiry_date = formatDateString(pden_pr_str_form.Expiry_date)
	pden_pr_str_form.Row_effective_date = formatDateString(pden_pr_str_form.Row_effective_date)
	pden_pr_str_form.Row_expiry_date = formatDateString(pden_pr_str_form.Row_expiry_date)






	rows, err := stmt.Exec(pden_pr_str_form.Pden_id, pden_pr_str_form.Pden_source, pden_pr_str_form.Active_ind, pden_pr_str_form.Effective_date, pden_pr_str_form.Expiry_date, pden_pr_str_form.Ppdm_guid, pden_pr_str_form.Pr_str_form_obs_no, pden_pr_str_form.Remark, pden_pr_str_form.Source, pden_pr_str_form.Strat_name_set_id, pden_pr_str_form.Strat_unit_id, pden_pr_str_form.String_id, pden_pr_str_form.String_source, pden_pr_str_form.Uwi, pden_pr_str_form.Row_changed_by, pden_pr_str_form.Row_changed_date, pden_pr_str_form.Row_created_by, pden_pr_str_form.Row_effective_date, pden_pr_str_form.Row_expiry_date, pden_pr_str_form.Row_quality, pden_pr_str_form.Pden_subtype)
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

func PatchPdenPrStrForm(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pden_pr_str_form set "
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

func DeletePdenPrStrForm(c *fiber.Ctx) error {
	id := c.Params("id")
	var pden_pr_str_form dto.Pden_pr_str_form
	pden_pr_str_form.Pden_subtype = id

	stmt, err := db.Prepare("delete from pden_pr_str_form where pden_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pden_pr_str_form.Pden_subtype)
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



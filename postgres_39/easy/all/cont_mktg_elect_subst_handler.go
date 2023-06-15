package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetContMktgElectSubst(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cont_mktg_elect_subst")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cont_mktg_elect_subst

	for rows.Next() {
		var cont_mktg_elect_subst dto.Cont_mktg_elect_subst
		if err := rows.Scan(&cont_mktg_elect_subst.Contract_id, &cont_mktg_elect_subst.Operating_procedure_id, &cont_mktg_elect_subst.Substance_id, &cont_mktg_elect_subst.Substance_obs_no, &cont_mktg_elect_subst.Active_ind, &cont_mktg_elect_subst.Both_ind, &cont_mktg_elect_subst.Effective_date, &cont_mktg_elect_subst.Election_cost, &cont_mktg_elect_subst.Election_cost_qualifier, &cont_mktg_elect_subst.Election_percent, &cont_mktg_elect_subst.Election_percent_qualifier, &cont_mktg_elect_subst.Expiry_date, &cont_mktg_elect_subst.Ppdm_guid, &cont_mktg_elect_subst.Remark, &cont_mktg_elect_subst.Source, &cont_mktg_elect_subst.Substance_ouom, &cont_mktg_elect_subst.Substance_uom, &cont_mktg_elect_subst.Row_changed_by, &cont_mktg_elect_subst.Row_changed_date, &cont_mktg_elect_subst.Row_created_by, &cont_mktg_elect_subst.Row_created_date, &cont_mktg_elect_subst.Row_effective_date, &cont_mktg_elect_subst.Row_expiry_date, &cont_mktg_elect_subst.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cont_mktg_elect_subst to result
		result = append(result, cont_mktg_elect_subst)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetContMktgElectSubst(c *fiber.Ctx) error {
	var cont_mktg_elect_subst dto.Cont_mktg_elect_subst

	setDefaults(&cont_mktg_elect_subst)

	if err := json.Unmarshal(c.Body(), &cont_mktg_elect_subst); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cont_mktg_elect_subst values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24)")
	if err != nil {
		return err
	}
	cont_mktg_elect_subst.Row_created_date = formatDate(cont_mktg_elect_subst.Row_created_date)
	cont_mktg_elect_subst.Row_changed_date = formatDate(cont_mktg_elect_subst.Row_changed_date)
	cont_mktg_elect_subst.Effective_date = formatDateString(cont_mktg_elect_subst.Effective_date)
	cont_mktg_elect_subst.Expiry_date = formatDateString(cont_mktg_elect_subst.Expiry_date)
	cont_mktg_elect_subst.Row_effective_date = formatDateString(cont_mktg_elect_subst.Row_effective_date)
	cont_mktg_elect_subst.Row_expiry_date = formatDateString(cont_mktg_elect_subst.Row_expiry_date)






	rows, err := stmt.Exec(cont_mktg_elect_subst.Contract_id, cont_mktg_elect_subst.Operating_procedure_id, cont_mktg_elect_subst.Substance_id, cont_mktg_elect_subst.Substance_obs_no, cont_mktg_elect_subst.Active_ind, cont_mktg_elect_subst.Both_ind, cont_mktg_elect_subst.Effective_date, cont_mktg_elect_subst.Election_cost, cont_mktg_elect_subst.Election_cost_qualifier, cont_mktg_elect_subst.Election_percent, cont_mktg_elect_subst.Election_percent_qualifier, cont_mktg_elect_subst.Expiry_date, cont_mktg_elect_subst.Ppdm_guid, cont_mktg_elect_subst.Remark, cont_mktg_elect_subst.Source, cont_mktg_elect_subst.Substance_ouom, cont_mktg_elect_subst.Substance_uom, cont_mktg_elect_subst.Row_changed_by, cont_mktg_elect_subst.Row_changed_date, cont_mktg_elect_subst.Row_created_by, cont_mktg_elect_subst.Row_created_date, cont_mktg_elect_subst.Row_effective_date, cont_mktg_elect_subst.Row_expiry_date, cont_mktg_elect_subst.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateContMktgElectSubst(c *fiber.Ctx) error {
	var cont_mktg_elect_subst dto.Cont_mktg_elect_subst

	setDefaults(&cont_mktg_elect_subst)

	if err := json.Unmarshal(c.Body(), &cont_mktg_elect_subst); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cont_mktg_elect_subst.Contract_id = id

    if cont_mktg_elect_subst.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cont_mktg_elect_subst set operating_procedure_id = :1, substance_id = :2, substance_obs_no = :3, active_ind = :4, both_ind = :5, effective_date = :6, election_cost = :7, election_cost_qualifier = :8, election_percent = :9, election_percent_qualifier = :10, expiry_date = :11, ppdm_guid = :12, remark = :13, source = :14, substance_ouom = :15, substance_uom = :16, row_changed_by = :17, row_changed_date = :18, row_created_by = :19, row_effective_date = :20, row_expiry_date = :21, row_quality = :22 where contract_id = :24")
	if err != nil {
		return err
	}

	cont_mktg_elect_subst.Row_changed_date = formatDate(cont_mktg_elect_subst.Row_changed_date)
	cont_mktg_elect_subst.Effective_date = formatDateString(cont_mktg_elect_subst.Effective_date)
	cont_mktg_elect_subst.Expiry_date = formatDateString(cont_mktg_elect_subst.Expiry_date)
	cont_mktg_elect_subst.Row_effective_date = formatDateString(cont_mktg_elect_subst.Row_effective_date)
	cont_mktg_elect_subst.Row_expiry_date = formatDateString(cont_mktg_elect_subst.Row_expiry_date)






	rows, err := stmt.Exec(cont_mktg_elect_subst.Operating_procedure_id, cont_mktg_elect_subst.Substance_id, cont_mktg_elect_subst.Substance_obs_no, cont_mktg_elect_subst.Active_ind, cont_mktg_elect_subst.Both_ind, cont_mktg_elect_subst.Effective_date, cont_mktg_elect_subst.Election_cost, cont_mktg_elect_subst.Election_cost_qualifier, cont_mktg_elect_subst.Election_percent, cont_mktg_elect_subst.Election_percent_qualifier, cont_mktg_elect_subst.Expiry_date, cont_mktg_elect_subst.Ppdm_guid, cont_mktg_elect_subst.Remark, cont_mktg_elect_subst.Source, cont_mktg_elect_subst.Substance_ouom, cont_mktg_elect_subst.Substance_uom, cont_mktg_elect_subst.Row_changed_by, cont_mktg_elect_subst.Row_changed_date, cont_mktg_elect_subst.Row_created_by, cont_mktg_elect_subst.Row_effective_date, cont_mktg_elect_subst.Row_expiry_date, cont_mktg_elect_subst.Row_quality, cont_mktg_elect_subst.Contract_id)
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

func PatchContMktgElectSubst(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cont_mktg_elect_subst set "
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
	query += " where contract_id = :id"

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

func DeleteContMktgElectSubst(c *fiber.Ctx) error {
	id := c.Params("id")
	var cont_mktg_elect_subst dto.Cont_mktg_elect_subst
	cont_mktg_elect_subst.Contract_id = id

	stmt, err := db.Prepare("delete from cont_mktg_elect_subst where contract_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cont_mktg_elect_subst.Contract_id)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFinance(c *fiber.Ctx) error {
	rows, err := db.Query("select * from finance")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Finance

	for rows.Next() {
		var finance dto.Finance
		if err := rows.Scan(&finance.Finance_id, &finance.Active_ind, &finance.Actual_cost, &finance.Authorized_by_ba_id, &finance.Budget_cost, &finance.Currency_conversion, &finance.Currency_ouom, &finance.Effective_date, &finance.Expiry_date, &finance.Finance_type, &finance.Fin_status, &finance.Issue_date, &finance.Limit_amount, &finance.Ppdm_guid, &finance.Reference_number, &finance.Remark, &finance.Source, &finance.Tax_credit_code, &finance.Row_changed_by, &finance.Row_changed_date, &finance.Row_created_by, &finance.Row_created_date, &finance.Row_effective_date, &finance.Row_expiry_date, &finance.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append finance to result
		result = append(result, finance)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFinance(c *fiber.Ctx) error {
	var finance dto.Finance

	setDefaults(&finance)

	if err := json.Unmarshal(c.Body(), &finance); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into finance values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25)")
	if err != nil {
		return err
	}
	finance.Row_created_date = formatDate(finance.Row_created_date)
	finance.Row_changed_date = formatDate(finance.Row_changed_date)
	finance.Effective_date = formatDateString(finance.Effective_date)
	finance.Expiry_date = formatDateString(finance.Expiry_date)
	finance.Issue_date = formatDateString(finance.Issue_date)
	finance.Row_effective_date = formatDateString(finance.Row_effective_date)
	finance.Row_expiry_date = formatDateString(finance.Row_expiry_date)






	rows, err := stmt.Exec(finance.Finance_id, finance.Active_ind, finance.Actual_cost, finance.Authorized_by_ba_id, finance.Budget_cost, finance.Currency_conversion, finance.Currency_ouom, finance.Effective_date, finance.Expiry_date, finance.Finance_type, finance.Fin_status, finance.Issue_date, finance.Limit_amount, finance.Ppdm_guid, finance.Reference_number, finance.Remark, finance.Source, finance.Tax_credit_code, finance.Row_changed_by, finance.Row_changed_date, finance.Row_created_by, finance.Row_created_date, finance.Row_effective_date, finance.Row_expiry_date, finance.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFinance(c *fiber.Ctx) error {
	var finance dto.Finance

	setDefaults(&finance)

	if err := json.Unmarshal(c.Body(), &finance); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	finance.Finance_id = id

    if finance.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update finance set active_ind = :1, actual_cost = :2, authorized_by_ba_id = :3, budget_cost = :4, currency_conversion = :5, currency_ouom = :6, effective_date = :7, expiry_date = :8, finance_type = :9, fin_status = :10, issue_date = :11, limit_amount = :12, ppdm_guid = :13, reference_number = :14, remark = :15, source = :16, tax_credit_code = :17, row_changed_by = :18, row_changed_date = :19, row_created_by = :20, row_effective_date = :21, row_expiry_date = :22, row_quality = :23 where finance_id = :25")
	if err != nil {
		return err
	}

	finance.Row_changed_date = formatDate(finance.Row_changed_date)
	finance.Effective_date = formatDateString(finance.Effective_date)
	finance.Expiry_date = formatDateString(finance.Expiry_date)
	finance.Issue_date = formatDateString(finance.Issue_date)
	finance.Row_effective_date = formatDateString(finance.Row_effective_date)
	finance.Row_expiry_date = formatDateString(finance.Row_expiry_date)






	rows, err := stmt.Exec(finance.Active_ind, finance.Actual_cost, finance.Authorized_by_ba_id, finance.Budget_cost, finance.Currency_conversion, finance.Currency_ouom, finance.Effective_date, finance.Expiry_date, finance.Finance_type, finance.Fin_status, finance.Issue_date, finance.Limit_amount, finance.Ppdm_guid, finance.Reference_number, finance.Remark, finance.Source, finance.Tax_credit_code, finance.Row_changed_by, finance.Row_changed_date, finance.Row_created_by, finance.Row_effective_date, finance.Row_expiry_date, finance.Row_quality, finance.Finance_id)
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

func PatchFinance(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update finance set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "issue_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where finance_id = :id"

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

func DeleteFinance(c *fiber.Ctx) error {
	id := c.Params("id")
	var finance dto.Finance
	finance.Finance_id = id

	stmt, err := db.Prepare("delete from finance where finance_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(finance.Finance_id)
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



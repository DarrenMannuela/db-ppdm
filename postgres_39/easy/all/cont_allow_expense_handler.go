package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetContAllowExpense(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cont_allow_expense")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cont_allow_expense

	for rows.Next() {
		var cont_allow_expense dto.Cont_allow_expense
		if err := rows.Scan(&cont_allow_expense.Contract_id, &cont_allow_expense.Allow_expense_type, &cont_allow_expense.Allow_expense_obs_no, &cont_allow_expense.Account_proc_obs_no, &cont_allow_expense.Account_proc_type, &cont_allow_expense.Active_ind, &cont_allow_expense.Allow_percent, &cont_allow_expense.Allow_percent_qualifier, &cont_allow_expense.Currency_conversion, &cont_allow_expense.Currency_ouom, &cont_allow_expense.Effective_date, &cont_allow_expense.Expense_frequency, &cont_allow_expense.Expense_frequency_uom, &cont_allow_expense.Expense_rate, &cont_allow_expense.Expense_rate_qualifier, &cont_allow_expense.Expiry_date, &cont_allow_expense.Ineligible_ind, &cont_allow_expense.Ppdm_guid, &cont_allow_expense.Remark, &cont_allow_expense.Source, &cont_allow_expense.Row_changed_by, &cont_allow_expense.Row_changed_date, &cont_allow_expense.Row_created_by, &cont_allow_expense.Row_created_date, &cont_allow_expense.Row_effective_date, &cont_allow_expense.Row_expiry_date, &cont_allow_expense.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cont_allow_expense to result
		result = append(result, cont_allow_expense)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetContAllowExpense(c *fiber.Ctx) error {
	var cont_allow_expense dto.Cont_allow_expense

	setDefaults(&cont_allow_expense)

	if err := json.Unmarshal(c.Body(), &cont_allow_expense); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cont_allow_expense values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27)")
	if err != nil {
		return err
	}
	cont_allow_expense.Row_created_date = formatDate(cont_allow_expense.Row_created_date)
	cont_allow_expense.Row_changed_date = formatDate(cont_allow_expense.Row_changed_date)
	cont_allow_expense.Effective_date = formatDateString(cont_allow_expense.Effective_date)
	cont_allow_expense.Expiry_date = formatDateString(cont_allow_expense.Expiry_date)
	cont_allow_expense.Row_effective_date = formatDateString(cont_allow_expense.Row_effective_date)
	cont_allow_expense.Row_expiry_date = formatDateString(cont_allow_expense.Row_expiry_date)






	rows, err := stmt.Exec(cont_allow_expense.Contract_id, cont_allow_expense.Allow_expense_type, cont_allow_expense.Allow_expense_obs_no, cont_allow_expense.Account_proc_obs_no, cont_allow_expense.Account_proc_type, cont_allow_expense.Active_ind, cont_allow_expense.Allow_percent, cont_allow_expense.Allow_percent_qualifier, cont_allow_expense.Currency_conversion, cont_allow_expense.Currency_ouom, cont_allow_expense.Effective_date, cont_allow_expense.Expense_frequency, cont_allow_expense.Expense_frequency_uom, cont_allow_expense.Expense_rate, cont_allow_expense.Expense_rate_qualifier, cont_allow_expense.Expiry_date, cont_allow_expense.Ineligible_ind, cont_allow_expense.Ppdm_guid, cont_allow_expense.Remark, cont_allow_expense.Source, cont_allow_expense.Row_changed_by, cont_allow_expense.Row_changed_date, cont_allow_expense.Row_created_by, cont_allow_expense.Row_created_date, cont_allow_expense.Row_effective_date, cont_allow_expense.Row_expiry_date, cont_allow_expense.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateContAllowExpense(c *fiber.Ctx) error {
	var cont_allow_expense dto.Cont_allow_expense

	setDefaults(&cont_allow_expense)

	if err := json.Unmarshal(c.Body(), &cont_allow_expense); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cont_allow_expense.Contract_id = id

    if cont_allow_expense.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cont_allow_expense set allow_expense_type = :1, allow_expense_obs_no = :2, account_proc_obs_no = :3, account_proc_type = :4, active_ind = :5, allow_percent = :6, allow_percent_qualifier = :7, currency_conversion = :8, currency_ouom = :9, effective_date = :10, expense_frequency = :11, expense_frequency_uom = :12, expense_rate = :13, expense_rate_qualifier = :14, expiry_date = :15, ineligible_ind = :16, ppdm_guid = :17, remark = :18, source = :19, row_changed_by = :20, row_changed_date = :21, row_created_by = :22, row_effective_date = :23, row_expiry_date = :24, row_quality = :25 where contract_id = :27")
	if err != nil {
		return err
	}

	cont_allow_expense.Row_changed_date = formatDate(cont_allow_expense.Row_changed_date)
	cont_allow_expense.Effective_date = formatDateString(cont_allow_expense.Effective_date)
	cont_allow_expense.Expiry_date = formatDateString(cont_allow_expense.Expiry_date)
	cont_allow_expense.Row_effective_date = formatDateString(cont_allow_expense.Row_effective_date)
	cont_allow_expense.Row_expiry_date = formatDateString(cont_allow_expense.Row_expiry_date)






	rows, err := stmt.Exec(cont_allow_expense.Allow_expense_type, cont_allow_expense.Allow_expense_obs_no, cont_allow_expense.Account_proc_obs_no, cont_allow_expense.Account_proc_type, cont_allow_expense.Active_ind, cont_allow_expense.Allow_percent, cont_allow_expense.Allow_percent_qualifier, cont_allow_expense.Currency_conversion, cont_allow_expense.Currency_ouom, cont_allow_expense.Effective_date, cont_allow_expense.Expense_frequency, cont_allow_expense.Expense_frequency_uom, cont_allow_expense.Expense_rate, cont_allow_expense.Expense_rate_qualifier, cont_allow_expense.Expiry_date, cont_allow_expense.Ineligible_ind, cont_allow_expense.Ppdm_guid, cont_allow_expense.Remark, cont_allow_expense.Source, cont_allow_expense.Row_changed_by, cont_allow_expense.Row_changed_date, cont_allow_expense.Row_created_by, cont_allow_expense.Row_effective_date, cont_allow_expense.Row_expiry_date, cont_allow_expense.Row_quality, cont_allow_expense.Contract_id)
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

func PatchContAllowExpense(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cont_allow_expense set "
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

func DeleteContAllowExpense(c *fiber.Ctx) error {
	id := c.Params("id")
	var cont_allow_expense dto.Cont_allow_expense
	cont_allow_expense.Contract_id = id

	stmt, err := db.Prepare("delete from cont_allow_expense where contract_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cont_allow_expense.Contract_id)
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



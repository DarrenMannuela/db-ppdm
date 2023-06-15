package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetObligPaymentInstr(c *fiber.Ctx) error {
	rows, err := db.Query("select * from oblig_payment_instr")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Oblig_payment_instr

	for rows.Next() {
		var oblig_payment_instr dto.Oblig_payment_instr
		if err := rows.Scan(&oblig_payment_instr.Payee_ba_id, &oblig_payment_instr.Payment_instruction_id, &oblig_payment_instr.Aba_number, &oblig_payment_instr.Active_ind, &oblig_payment_instr.Bank_address_obs_no, &oblig_payment_instr.Bank_address_source, &oblig_payment_instr.Bank_ba_id, &oblig_payment_instr.Bank_fee, &oblig_payment_instr.Currency_conversion, &oblig_payment_instr.Currency_ouom, &oblig_payment_instr.Depository_num, &oblig_payment_instr.Effective_date, &oblig_payment_instr.Expiry_date, &oblig_payment_instr.Invalid_date, &oblig_payment_instr.Pay_method, &oblig_payment_instr.Ppdm_guid, &oblig_payment_instr.Remark, &oblig_payment_instr.Source, &oblig_payment_instr.Suspend_payment_ind, &oblig_payment_instr.Suspend_payment_reason, &oblig_payment_instr.Row_changed_by, &oblig_payment_instr.Row_changed_date, &oblig_payment_instr.Row_created_by, &oblig_payment_instr.Row_created_date, &oblig_payment_instr.Row_effective_date, &oblig_payment_instr.Row_expiry_date, &oblig_payment_instr.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append oblig_payment_instr to result
		result = append(result, oblig_payment_instr)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetObligPaymentInstr(c *fiber.Ctx) error {
	var oblig_payment_instr dto.Oblig_payment_instr

	setDefaults(&oblig_payment_instr)

	if err := json.Unmarshal(c.Body(), &oblig_payment_instr); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into oblig_payment_instr values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27)")
	if err != nil {
		return err
	}
	oblig_payment_instr.Row_created_date = formatDate(oblig_payment_instr.Row_created_date)
	oblig_payment_instr.Row_changed_date = formatDate(oblig_payment_instr.Row_changed_date)
	oblig_payment_instr.Effective_date = formatDateString(oblig_payment_instr.Effective_date)
	oblig_payment_instr.Expiry_date = formatDateString(oblig_payment_instr.Expiry_date)
	oblig_payment_instr.Invalid_date = formatDateString(oblig_payment_instr.Invalid_date)
	oblig_payment_instr.Row_effective_date = formatDateString(oblig_payment_instr.Row_effective_date)
	oblig_payment_instr.Row_expiry_date = formatDateString(oblig_payment_instr.Row_expiry_date)






	rows, err := stmt.Exec(oblig_payment_instr.Payee_ba_id, oblig_payment_instr.Payment_instruction_id, oblig_payment_instr.Aba_number, oblig_payment_instr.Active_ind, oblig_payment_instr.Bank_address_obs_no, oblig_payment_instr.Bank_address_source, oblig_payment_instr.Bank_ba_id, oblig_payment_instr.Bank_fee, oblig_payment_instr.Currency_conversion, oblig_payment_instr.Currency_ouom, oblig_payment_instr.Depository_num, oblig_payment_instr.Effective_date, oblig_payment_instr.Expiry_date, oblig_payment_instr.Invalid_date, oblig_payment_instr.Pay_method, oblig_payment_instr.Ppdm_guid, oblig_payment_instr.Remark, oblig_payment_instr.Source, oblig_payment_instr.Suspend_payment_ind, oblig_payment_instr.Suspend_payment_reason, oblig_payment_instr.Row_changed_by, oblig_payment_instr.Row_changed_date, oblig_payment_instr.Row_created_by, oblig_payment_instr.Row_created_date, oblig_payment_instr.Row_effective_date, oblig_payment_instr.Row_expiry_date, oblig_payment_instr.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateObligPaymentInstr(c *fiber.Ctx) error {
	var oblig_payment_instr dto.Oblig_payment_instr

	setDefaults(&oblig_payment_instr)

	if err := json.Unmarshal(c.Body(), &oblig_payment_instr); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	oblig_payment_instr.Payee_ba_id = id

    if oblig_payment_instr.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update oblig_payment_instr set payment_instruction_id = :1, aba_number = :2, active_ind = :3, bank_address_obs_no = :4, bank_address_source = :5, bank_ba_id = :6, bank_fee = :7, currency_conversion = :8, currency_ouom = :9, depository_num = :10, effective_date = :11, expiry_date = :12, invalid_date = :13, pay_method = :14, ppdm_guid = :15, remark = :16, source = :17, suspend_payment_ind = :18, suspend_payment_reason = :19, row_changed_by = :20, row_changed_date = :21, row_created_by = :22, row_effective_date = :23, row_expiry_date = :24, row_quality = :25 where payee_ba_id = :27")
	if err != nil {
		return err
	}

	oblig_payment_instr.Row_changed_date = formatDate(oblig_payment_instr.Row_changed_date)
	oblig_payment_instr.Effective_date = formatDateString(oblig_payment_instr.Effective_date)
	oblig_payment_instr.Expiry_date = formatDateString(oblig_payment_instr.Expiry_date)
	oblig_payment_instr.Invalid_date = formatDateString(oblig_payment_instr.Invalid_date)
	oblig_payment_instr.Row_effective_date = formatDateString(oblig_payment_instr.Row_effective_date)
	oblig_payment_instr.Row_expiry_date = formatDateString(oblig_payment_instr.Row_expiry_date)






	rows, err := stmt.Exec(oblig_payment_instr.Payment_instruction_id, oblig_payment_instr.Aba_number, oblig_payment_instr.Active_ind, oblig_payment_instr.Bank_address_obs_no, oblig_payment_instr.Bank_address_source, oblig_payment_instr.Bank_ba_id, oblig_payment_instr.Bank_fee, oblig_payment_instr.Currency_conversion, oblig_payment_instr.Currency_ouom, oblig_payment_instr.Depository_num, oblig_payment_instr.Effective_date, oblig_payment_instr.Expiry_date, oblig_payment_instr.Invalid_date, oblig_payment_instr.Pay_method, oblig_payment_instr.Ppdm_guid, oblig_payment_instr.Remark, oblig_payment_instr.Source, oblig_payment_instr.Suspend_payment_ind, oblig_payment_instr.Suspend_payment_reason, oblig_payment_instr.Row_changed_by, oblig_payment_instr.Row_changed_date, oblig_payment_instr.Row_created_by, oblig_payment_instr.Row_effective_date, oblig_payment_instr.Row_expiry_date, oblig_payment_instr.Row_quality, oblig_payment_instr.Payee_ba_id)
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

func PatchObligPaymentInstr(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update oblig_payment_instr set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "invalid_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where payee_ba_id = :id"

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

func DeleteObligPaymentInstr(c *fiber.Ctx) error {
	id := c.Params("id")
	var oblig_payment_instr dto.Oblig_payment_instr
	oblig_payment_instr.Payee_ba_id = id

	stmt, err := db.Prepare("delete from oblig_payment_instr where payee_ba_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(oblig_payment_instr.Payee_ba_id)
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



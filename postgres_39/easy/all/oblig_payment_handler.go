package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetObligPayment(c *fiber.Ctx) error {
	rows, err := db.Query("select * from oblig_payment")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Oblig_payment

	for rows.Next() {
		var oblig_payment dto.Oblig_payment
		if err := rows.Scan(&oblig_payment.Obligation_id, &oblig_payment.Obligation_seq_no, &oblig_payment.Payee_payor_ba_id, &oblig_payment.Payee_payor, &oblig_payment.Active_ind, &oblig_payment.Cheque_remark, &oblig_payment.Contract_id, &oblig_payment.Currency_conversion, &oblig_payment.Currency_ouom, &oblig_payment.Currency_uom, &oblig_payment.Effective_date, &oblig_payment.Expiry_date, &oblig_payment.Gross_cost, &oblig_payment.Invoice_amount, &oblig_payment.Invoice_num, &oblig_payment.Land_rental_type, &oblig_payment.Net_cost, &oblig_payment.Payment_due_date, &oblig_payment.Payment_ind, &oblig_payment.Payment_instruction_id, &oblig_payment.Payment_type, &oblig_payment.Pay_method, &oblig_payment.Percent_of_gross, &oblig_payment.Ppdm_guid, &oblig_payment.Provision_id, &oblig_payment.Rate_schedule_id, &oblig_payment.Rate_type, &oblig_payment.Receipt_ind, &oblig_payment.Related_obligation_id, &oblig_payment.Related_obligation_seq_no, &oblig_payment.Related_oblig_ba_id, &oblig_payment.Related_oblig_payee_payor, &oblig_payment.Remark, &oblig_payment.Royalty_type, &oblig_payment.Source, &oblig_payment.Stop_payment_ind, &oblig_payment.Substance_id, &oblig_payment.Substance_seq_no, &oblig_payment.Row_changed_by, &oblig_payment.Row_changed_date, &oblig_payment.Row_created_by, &oblig_payment.Row_created_date, &oblig_payment.Row_effective_date, &oblig_payment.Row_expiry_date, &oblig_payment.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append oblig_payment to result
		result = append(result, oblig_payment)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetObligPayment(c *fiber.Ctx) error {
	var oblig_payment dto.Oblig_payment

	setDefaults(&oblig_payment)

	if err := json.Unmarshal(c.Body(), &oblig_payment); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into oblig_payment values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45)")
	if err != nil {
		return err
	}
	oblig_payment.Row_created_date = formatDate(oblig_payment.Row_created_date)
	oblig_payment.Row_changed_date = formatDate(oblig_payment.Row_changed_date)
	oblig_payment.Effective_date = formatDateString(oblig_payment.Effective_date)
	oblig_payment.Expiry_date = formatDateString(oblig_payment.Expiry_date)
	oblig_payment.Payment_due_date = formatDateString(oblig_payment.Payment_due_date)
	oblig_payment.Row_effective_date = formatDateString(oblig_payment.Row_effective_date)
	oblig_payment.Row_expiry_date = formatDateString(oblig_payment.Row_expiry_date)






	rows, err := stmt.Exec(oblig_payment.Obligation_id, oblig_payment.Obligation_seq_no, oblig_payment.Payee_payor_ba_id, oblig_payment.Payee_payor, oblig_payment.Active_ind, oblig_payment.Cheque_remark, oblig_payment.Contract_id, oblig_payment.Currency_conversion, oblig_payment.Currency_ouom, oblig_payment.Currency_uom, oblig_payment.Effective_date, oblig_payment.Expiry_date, oblig_payment.Gross_cost, oblig_payment.Invoice_amount, oblig_payment.Invoice_num, oblig_payment.Land_rental_type, oblig_payment.Net_cost, oblig_payment.Payment_due_date, oblig_payment.Payment_ind, oblig_payment.Payment_instruction_id, oblig_payment.Payment_type, oblig_payment.Pay_method, oblig_payment.Percent_of_gross, oblig_payment.Ppdm_guid, oblig_payment.Provision_id, oblig_payment.Rate_schedule_id, oblig_payment.Rate_type, oblig_payment.Receipt_ind, oblig_payment.Related_obligation_id, oblig_payment.Related_obligation_seq_no, oblig_payment.Related_oblig_ba_id, oblig_payment.Related_oblig_payee_payor, oblig_payment.Remark, oblig_payment.Royalty_type, oblig_payment.Source, oblig_payment.Stop_payment_ind, oblig_payment.Substance_id, oblig_payment.Substance_seq_no, oblig_payment.Row_changed_by, oblig_payment.Row_changed_date, oblig_payment.Row_created_by, oblig_payment.Row_created_date, oblig_payment.Row_effective_date, oblig_payment.Row_expiry_date, oblig_payment.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateObligPayment(c *fiber.Ctx) error {
	var oblig_payment dto.Oblig_payment

	setDefaults(&oblig_payment)

	if err := json.Unmarshal(c.Body(), &oblig_payment); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	oblig_payment.Obligation_id = id

    if oblig_payment.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update oblig_payment set obligation_seq_no = :1, payee_payor_ba_id = :2, payee_payor = :3, active_ind = :4, cheque_remark = :5, contract_id = :6, currency_conversion = :7, currency_ouom = :8, currency_uom = :9, effective_date = :10, expiry_date = :11, gross_cost = :12, invoice_amount = :13, invoice_num = :14, land_rental_type = :15, net_cost = :16, payment_due_date = :17, payment_ind = :18, payment_instruction_id = :19, payment_type = :20, pay_method = :21, percent_of_gross = :22, ppdm_guid = :23, provision_id = :24, rate_schedule_id = :25, rate_type = :26, receipt_ind = :27, related_obligation_id = :28, related_obligation_seq_no = :29, related_oblig_ba_id = :30, related_oblig_payee_payor = :31, remark = :32, royalty_type = :33, source = :34, stop_payment_ind = :35, substance_id = :36, substance_seq_no = :37, row_changed_by = :38, row_changed_date = :39, row_created_by = :40, row_effective_date = :41, row_expiry_date = :42, row_quality = :43 where obligation_id = :45")
	if err != nil {
		return err
	}

	oblig_payment.Row_changed_date = formatDate(oblig_payment.Row_changed_date)
	oblig_payment.Effective_date = formatDateString(oblig_payment.Effective_date)
	oblig_payment.Expiry_date = formatDateString(oblig_payment.Expiry_date)
	oblig_payment.Payment_due_date = formatDateString(oblig_payment.Payment_due_date)
	oblig_payment.Row_effective_date = formatDateString(oblig_payment.Row_effective_date)
	oblig_payment.Row_expiry_date = formatDateString(oblig_payment.Row_expiry_date)






	rows, err := stmt.Exec(oblig_payment.Obligation_seq_no, oblig_payment.Payee_payor_ba_id, oblig_payment.Payee_payor, oblig_payment.Active_ind, oblig_payment.Cheque_remark, oblig_payment.Contract_id, oblig_payment.Currency_conversion, oblig_payment.Currency_ouom, oblig_payment.Currency_uom, oblig_payment.Effective_date, oblig_payment.Expiry_date, oblig_payment.Gross_cost, oblig_payment.Invoice_amount, oblig_payment.Invoice_num, oblig_payment.Land_rental_type, oblig_payment.Net_cost, oblig_payment.Payment_due_date, oblig_payment.Payment_ind, oblig_payment.Payment_instruction_id, oblig_payment.Payment_type, oblig_payment.Pay_method, oblig_payment.Percent_of_gross, oblig_payment.Ppdm_guid, oblig_payment.Provision_id, oblig_payment.Rate_schedule_id, oblig_payment.Rate_type, oblig_payment.Receipt_ind, oblig_payment.Related_obligation_id, oblig_payment.Related_obligation_seq_no, oblig_payment.Related_oblig_ba_id, oblig_payment.Related_oblig_payee_payor, oblig_payment.Remark, oblig_payment.Royalty_type, oblig_payment.Source, oblig_payment.Stop_payment_ind, oblig_payment.Substance_id, oblig_payment.Substance_seq_no, oblig_payment.Row_changed_by, oblig_payment.Row_changed_date, oblig_payment.Row_created_by, oblig_payment.Row_effective_date, oblig_payment.Row_expiry_date, oblig_payment.Row_quality, oblig_payment.Obligation_id)
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

func PatchObligPayment(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update oblig_payment set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "payment_due_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where obligation_id = :id"

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

func DeleteObligPayment(c *fiber.Ctx) error {
	id := c.Params("id")
	var oblig_payment dto.Oblig_payment
	oblig_payment.Obligation_id = id

	stmt, err := db.Prepare("delete from oblig_payment where obligation_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(oblig_payment.Obligation_id)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetObligPayDetail(c *fiber.Ctx) error {
	rows, err := db.Query("select * from oblig_pay_detail")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Oblig_pay_detail

	for rows.Next() {
		var oblig_pay_detail dto.Oblig_pay_detail
		if err := rows.Scan(&oblig_pay_detail.Obligation_id, &oblig_pay_detail.Obligation_seq_no, &oblig_pay_detail.Payee_payor_ba_id, &oblig_pay_detail.Payee_payor, &oblig_pay_detail.Detail_id, &oblig_pay_detail.Active_ind, &oblig_pay_detail.Cheque_number, &oblig_pay_detail.Currency_conversion, &oblig_pay_detail.Currency_ouom, &oblig_pay_detail.Deduction_ind, &oblig_pay_detail.Effective_date, &oblig_pay_detail.Expiry_date, &oblig_pay_detail.Payment_amount, &oblig_pay_detail.Payment_date, &oblig_pay_detail.Pay_detail_type, &oblig_pay_detail.Percent_of_payment, &oblig_pay_detail.Ppdm_guid, &oblig_pay_detail.Rate, &oblig_pay_detail.Rate_ouom, &oblig_pay_detail.Remark, &oblig_pay_detail.Source, &oblig_pay_detail.Row_changed_by, &oblig_pay_detail.Row_changed_date, &oblig_pay_detail.Row_created_by, &oblig_pay_detail.Row_created_date, &oblig_pay_detail.Row_effective_date, &oblig_pay_detail.Row_expiry_date, &oblig_pay_detail.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append oblig_pay_detail to result
		result = append(result, oblig_pay_detail)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetObligPayDetail(c *fiber.Ctx) error {
	var oblig_pay_detail dto.Oblig_pay_detail

	setDefaults(&oblig_pay_detail)

	if err := json.Unmarshal(c.Body(), &oblig_pay_detail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into oblig_pay_detail values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28)")
	if err != nil {
		return err
	}
	oblig_pay_detail.Row_created_date = formatDate(oblig_pay_detail.Row_created_date)
	oblig_pay_detail.Row_changed_date = formatDate(oblig_pay_detail.Row_changed_date)
	oblig_pay_detail.Effective_date = formatDateString(oblig_pay_detail.Effective_date)
	oblig_pay_detail.Expiry_date = formatDateString(oblig_pay_detail.Expiry_date)
	oblig_pay_detail.Payment_date = formatDateString(oblig_pay_detail.Payment_date)
	oblig_pay_detail.Row_effective_date = formatDateString(oblig_pay_detail.Row_effective_date)
	oblig_pay_detail.Row_expiry_date = formatDateString(oblig_pay_detail.Row_expiry_date)






	rows, err := stmt.Exec(oblig_pay_detail.Obligation_id, oblig_pay_detail.Obligation_seq_no, oblig_pay_detail.Payee_payor_ba_id, oblig_pay_detail.Payee_payor, oblig_pay_detail.Detail_id, oblig_pay_detail.Active_ind, oblig_pay_detail.Cheque_number, oblig_pay_detail.Currency_conversion, oblig_pay_detail.Currency_ouom, oblig_pay_detail.Deduction_ind, oblig_pay_detail.Effective_date, oblig_pay_detail.Expiry_date, oblig_pay_detail.Payment_amount, oblig_pay_detail.Payment_date, oblig_pay_detail.Pay_detail_type, oblig_pay_detail.Percent_of_payment, oblig_pay_detail.Ppdm_guid, oblig_pay_detail.Rate, oblig_pay_detail.Rate_ouom, oblig_pay_detail.Remark, oblig_pay_detail.Source, oblig_pay_detail.Row_changed_by, oblig_pay_detail.Row_changed_date, oblig_pay_detail.Row_created_by, oblig_pay_detail.Row_created_date, oblig_pay_detail.Row_effective_date, oblig_pay_detail.Row_expiry_date, oblig_pay_detail.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateObligPayDetail(c *fiber.Ctx) error {
	var oblig_pay_detail dto.Oblig_pay_detail

	setDefaults(&oblig_pay_detail)

	if err := json.Unmarshal(c.Body(), &oblig_pay_detail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	oblig_pay_detail.Obligation_id = id

    if oblig_pay_detail.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update oblig_pay_detail set obligation_seq_no = :1, payee_payor_ba_id = :2, payee_payor = :3, detail_id = :4, active_ind = :5, cheque_number = :6, currency_conversion = :7, currency_ouom = :8, deduction_ind = :9, effective_date = :10, expiry_date = :11, payment_amount = :12, payment_date = :13, pay_detail_type = :14, percent_of_payment = :15, ppdm_guid = :16, rate = :17, rate_ouom = :18, remark = :19, source = :20, row_changed_by = :21, row_changed_date = :22, row_created_by = :23, row_effective_date = :24, row_expiry_date = :25, row_quality = :26 where obligation_id = :28")
	if err != nil {
		return err
	}

	oblig_pay_detail.Row_changed_date = formatDate(oblig_pay_detail.Row_changed_date)
	oblig_pay_detail.Effective_date = formatDateString(oblig_pay_detail.Effective_date)
	oblig_pay_detail.Expiry_date = formatDateString(oblig_pay_detail.Expiry_date)
	oblig_pay_detail.Payment_date = formatDateString(oblig_pay_detail.Payment_date)
	oblig_pay_detail.Row_effective_date = formatDateString(oblig_pay_detail.Row_effective_date)
	oblig_pay_detail.Row_expiry_date = formatDateString(oblig_pay_detail.Row_expiry_date)






	rows, err := stmt.Exec(oblig_pay_detail.Obligation_seq_no, oblig_pay_detail.Payee_payor_ba_id, oblig_pay_detail.Payee_payor, oblig_pay_detail.Detail_id, oblig_pay_detail.Active_ind, oblig_pay_detail.Cheque_number, oblig_pay_detail.Currency_conversion, oblig_pay_detail.Currency_ouom, oblig_pay_detail.Deduction_ind, oblig_pay_detail.Effective_date, oblig_pay_detail.Expiry_date, oblig_pay_detail.Payment_amount, oblig_pay_detail.Payment_date, oblig_pay_detail.Pay_detail_type, oblig_pay_detail.Percent_of_payment, oblig_pay_detail.Ppdm_guid, oblig_pay_detail.Rate, oblig_pay_detail.Rate_ouom, oblig_pay_detail.Remark, oblig_pay_detail.Source, oblig_pay_detail.Row_changed_by, oblig_pay_detail.Row_changed_date, oblig_pay_detail.Row_created_by, oblig_pay_detail.Row_effective_date, oblig_pay_detail.Row_expiry_date, oblig_pay_detail.Row_quality, oblig_pay_detail.Obligation_id)
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

func PatchObligPayDetail(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update oblig_pay_detail set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "payment_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteObligPayDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	var oblig_pay_detail dto.Oblig_pay_detail
	oblig_pay_detail.Obligation_id = id

	stmt, err := db.Prepare("delete from oblig_pay_detail where obligation_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(oblig_pay_detail.Obligation_id)
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



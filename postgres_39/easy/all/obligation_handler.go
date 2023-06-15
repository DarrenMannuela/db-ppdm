package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetObligation(c *fiber.Ctx) error {
	rows, err := db.Query("select * from obligation")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Obligation

	for rows.Next() {
		var obligation dto.Obligation
		if err := rows.Scan(&obligation.Obligation_id, &obligation.Obligation_seq_no, &obligation.Active_ind, &obligation.Calculation_method, &obligation.Convertible_ind, &obligation.Critical_date, &obligation.Currency_conversion, &obligation.Currency_ouom, &obligation.Description, &obligation.Effective_date, &obligation.Expiry_date, &obligation.Fulfilled_date, &obligation.Fulfilled_ind, &obligation.Gross_obligation_cost, &obligation.Instrument_id, &obligation.Liability_release_date, &obligation.Net_obligation_cost, &obligation.Notice_period_length, &obligation.Notice_period_ouom, &obligation.Obligation_category, &obligation.Obligation_duration, &obligation.Obligation_duration_ouom, &obligation.Obligation_frequency, &obligation.Obligation_type, &obligation.Payment_ind, &obligation.Payment_responsibility, &obligation.Percentage, &obligation.Potential_obligation_desc, &obligation.Potential_obligation_ind, &obligation.Ppdm_guid, &obligation.Prepaid_ind, &obligation.Remark, &obligation.Resp_party_ba_id, &obligation.Review_frequency, &obligation.Royalty_owner_ba_id, &obligation.Royalty_payor_ba_id, &obligation.Royalty_type, &obligation.Source, &obligation.Start_date, &obligation.Substance_id, &obligation.Trigger_method, &obligation.Work_obligation_desc, &obligation.Row_changed_by, &obligation.Row_changed_date, &obligation.Row_created_by, &obligation.Row_created_date, &obligation.Row_effective_date, &obligation.Row_expiry_date, &obligation.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append obligation to result
		result = append(result, obligation)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetObligation(c *fiber.Ctx) error {
	var obligation dto.Obligation

	setDefaults(&obligation)

	if err := json.Unmarshal(c.Body(), &obligation); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into obligation values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49)")
	if err != nil {
		return err
	}
	obligation.Row_created_date = formatDate(obligation.Row_created_date)
	obligation.Row_changed_date = formatDate(obligation.Row_changed_date)
	obligation.Critical_date = formatDateString(obligation.Critical_date)
	obligation.Effective_date = formatDateString(obligation.Effective_date)
	obligation.Expiry_date = formatDateString(obligation.Expiry_date)
	obligation.Fulfilled_date = formatDateString(obligation.Fulfilled_date)
	obligation.Liability_release_date = formatDateString(obligation.Liability_release_date)
	obligation.Start_date = formatDateString(obligation.Start_date)
	obligation.Row_effective_date = formatDateString(obligation.Row_effective_date)
	obligation.Row_expiry_date = formatDateString(obligation.Row_expiry_date)






	rows, err := stmt.Exec(obligation.Obligation_id, obligation.Obligation_seq_no, obligation.Active_ind, obligation.Calculation_method, obligation.Convertible_ind, obligation.Critical_date, obligation.Currency_conversion, obligation.Currency_ouom, obligation.Description, obligation.Effective_date, obligation.Expiry_date, obligation.Fulfilled_date, obligation.Fulfilled_ind, obligation.Gross_obligation_cost, obligation.Instrument_id, obligation.Liability_release_date, obligation.Net_obligation_cost, obligation.Notice_period_length, obligation.Notice_period_ouom, obligation.Obligation_category, obligation.Obligation_duration, obligation.Obligation_duration_ouom, obligation.Obligation_frequency, obligation.Obligation_type, obligation.Payment_ind, obligation.Payment_responsibility, obligation.Percentage, obligation.Potential_obligation_desc, obligation.Potential_obligation_ind, obligation.Ppdm_guid, obligation.Prepaid_ind, obligation.Remark, obligation.Resp_party_ba_id, obligation.Review_frequency, obligation.Royalty_owner_ba_id, obligation.Royalty_payor_ba_id, obligation.Royalty_type, obligation.Source, obligation.Start_date, obligation.Substance_id, obligation.Trigger_method, obligation.Work_obligation_desc, obligation.Row_changed_by, obligation.Row_changed_date, obligation.Row_created_by, obligation.Row_created_date, obligation.Row_effective_date, obligation.Row_expiry_date, obligation.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateObligation(c *fiber.Ctx) error {
	var obligation dto.Obligation

	setDefaults(&obligation)

	if err := json.Unmarshal(c.Body(), &obligation); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	obligation.Obligation_id = id

    if obligation.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update obligation set obligation_seq_no = :1, active_ind = :2, calculation_method = :3, convertible_ind = :4, critical_date = :5, currency_conversion = :6, currency_ouom = :7, description = :8, effective_date = :9, expiry_date = :10, fulfilled_date = :11, fulfilled_ind = :12, gross_obligation_cost = :13, instrument_id = :14, liability_release_date = :15, net_obligation_cost = :16, notice_period_length = :17, notice_period_ouom = :18, obligation_category = :19, obligation_duration = :20, obligation_duration_ouom = :21, obligation_frequency = :22, obligation_type = :23, payment_ind = :24, payment_responsibility = :25, percentage = :26, potential_obligation_desc = :27, potential_obligation_ind = :28, ppdm_guid = :29, prepaid_ind = :30, remark = :31, resp_party_ba_id = :32, review_frequency = :33, royalty_owner_ba_id = :34, royalty_payor_ba_id = :35, royalty_type = :36, source = :37, start_date = :38, substance_id = :39, trigger_method = :40, work_obligation_desc = :41, row_changed_by = :42, row_changed_date = :43, row_created_by = :44, row_effective_date = :45, row_expiry_date = :46, row_quality = :47 where obligation_id = :49")
	if err != nil {
		return err
	}

	obligation.Row_changed_date = formatDate(obligation.Row_changed_date)
	obligation.Critical_date = formatDateString(obligation.Critical_date)
	obligation.Effective_date = formatDateString(obligation.Effective_date)
	obligation.Expiry_date = formatDateString(obligation.Expiry_date)
	obligation.Fulfilled_date = formatDateString(obligation.Fulfilled_date)
	obligation.Liability_release_date = formatDateString(obligation.Liability_release_date)
	obligation.Start_date = formatDateString(obligation.Start_date)
	obligation.Row_effective_date = formatDateString(obligation.Row_effective_date)
	obligation.Row_expiry_date = formatDateString(obligation.Row_expiry_date)






	rows, err := stmt.Exec(obligation.Obligation_seq_no, obligation.Active_ind, obligation.Calculation_method, obligation.Convertible_ind, obligation.Critical_date, obligation.Currency_conversion, obligation.Currency_ouom, obligation.Description, obligation.Effective_date, obligation.Expiry_date, obligation.Fulfilled_date, obligation.Fulfilled_ind, obligation.Gross_obligation_cost, obligation.Instrument_id, obligation.Liability_release_date, obligation.Net_obligation_cost, obligation.Notice_period_length, obligation.Notice_period_ouom, obligation.Obligation_category, obligation.Obligation_duration, obligation.Obligation_duration_ouom, obligation.Obligation_frequency, obligation.Obligation_type, obligation.Payment_ind, obligation.Payment_responsibility, obligation.Percentage, obligation.Potential_obligation_desc, obligation.Potential_obligation_ind, obligation.Ppdm_guid, obligation.Prepaid_ind, obligation.Remark, obligation.Resp_party_ba_id, obligation.Review_frequency, obligation.Royalty_owner_ba_id, obligation.Royalty_payor_ba_id, obligation.Royalty_type, obligation.Source, obligation.Start_date, obligation.Substance_id, obligation.Trigger_method, obligation.Work_obligation_desc, obligation.Row_changed_by, obligation.Row_changed_date, obligation.Row_created_by, obligation.Row_effective_date, obligation.Row_expiry_date, obligation.Row_quality, obligation.Obligation_id)
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

func PatchObligation(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update obligation set "
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
		} else if key == "critical_date" || key == "effective_date" || key == "expiry_date" || key == "fulfilled_date" || key == "liability_release_date" || key == "start_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteObligation(c *fiber.Ctx) error {
	id := c.Params("id")
	var obligation dto.Obligation
	obligation.Obligation_id = id

	stmt, err := db.Prepare("delete from obligation where obligation_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(obligation.Obligation_id)
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



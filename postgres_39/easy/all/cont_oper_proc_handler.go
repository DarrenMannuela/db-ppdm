package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetContOperProc(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cont_oper_proc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cont_oper_proc

	for rows.Next() {
		var cont_oper_proc dto.Cont_oper_proc
		if err := rows.Scan(&cont_oper_proc.Contract_id, &cont_oper_proc.Operating_procedure_id, &cont_oper_proc.Active_ind, &cont_oper_proc.Adjust_13_month_desc, &cont_oper_proc.Adjust_13_month_ind, &cont_oper_proc.Casing_point_election, &cont_oper_proc.Claim_amount_limit, &cont_oper_proc.Claim_amount_limit_ouom, &cont_oper_proc.Dev_penalty_cost, &cont_oper_proc.Dev_penalty_cost_ouom, &cont_oper_proc.Dev_penalty_percent, &cont_oper_proc.Effective_date, &cont_oper_proc.Equip_penalty_cost, &cont_oper_proc.Equip_penalty_cost_ouom, &cont_oper_proc.Equip_penalty_percent, &cont_oper_proc.Expiry_date, &cont_oper_proc.Expl_penalty_cost, &cont_oper_proc.Expl_penalty_cost_ouom, &cont_oper_proc.Expl_penalty_percent, &cont_oper_proc.Inflation_adjustment_ind, &cont_oper_proc.Insurance_election, &cont_oper_proc.Mktg_fee_elect_cost, &cont_oper_proc.Mktg_fee_elect_cost_ouom, &cont_oper_proc.Mktg_fee_elect_percent, &cont_oper_proc.Nat_currency_conversion, &cont_oper_proc.Nat_currency_uom, &cont_oper_proc.Operating_procedure_type, &cont_oper_proc.Ppdm_guid, &cont_oper_proc.Recog_on_assignment_ind, &cont_oper_proc.Remark, &cont_oper_proc.Rofr_ind, &cont_oper_proc.Rofr_reply_term, &cont_oper_proc.Rofr_reply_term_uom, &cont_oper_proc.Source, &cont_oper_proc.Source_document_id, &cont_oper_proc.Statute_limit_period, &cont_oper_proc.Statute_limit_period_uom, &cont_oper_proc.Surrender_notice, &cont_oper_proc.Surrender_notice_uom, &cont_oper_proc.Title_preserve_desc, &cont_oper_proc.Unauthorized_amount, &cont_oper_proc.Unauthorized_amount_ouom, &cont_oper_proc.Row_changed_by, &cont_oper_proc.Row_changed_date, &cont_oper_proc.Row_created_by, &cont_oper_proc.Row_created_date, &cont_oper_proc.Row_effective_date, &cont_oper_proc.Row_expiry_date, &cont_oper_proc.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cont_oper_proc to result
		result = append(result, cont_oper_proc)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetContOperProc(c *fiber.Ctx) error {
	var cont_oper_proc dto.Cont_oper_proc

	setDefaults(&cont_oper_proc)

	if err := json.Unmarshal(c.Body(), &cont_oper_proc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cont_oper_proc values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48, :49)")
	if err != nil {
		return err
	}
	cont_oper_proc.Row_created_date = formatDate(cont_oper_proc.Row_created_date)
	cont_oper_proc.Row_changed_date = formatDate(cont_oper_proc.Row_changed_date)
	cont_oper_proc.Effective_date = formatDateString(cont_oper_proc.Effective_date)
	cont_oper_proc.Expiry_date = formatDateString(cont_oper_proc.Expiry_date)
	cont_oper_proc.Row_effective_date = formatDateString(cont_oper_proc.Row_effective_date)
	cont_oper_proc.Row_expiry_date = formatDateString(cont_oper_proc.Row_expiry_date)






	rows, err := stmt.Exec(cont_oper_proc.Contract_id, cont_oper_proc.Operating_procedure_id, cont_oper_proc.Active_ind, cont_oper_proc.Adjust_13_month_desc, cont_oper_proc.Adjust_13_month_ind, cont_oper_proc.Casing_point_election, cont_oper_proc.Claim_amount_limit, cont_oper_proc.Claim_amount_limit_ouom, cont_oper_proc.Dev_penalty_cost, cont_oper_proc.Dev_penalty_cost_ouom, cont_oper_proc.Dev_penalty_percent, cont_oper_proc.Effective_date, cont_oper_proc.Equip_penalty_cost, cont_oper_proc.Equip_penalty_cost_ouom, cont_oper_proc.Equip_penalty_percent, cont_oper_proc.Expiry_date, cont_oper_proc.Expl_penalty_cost, cont_oper_proc.Expl_penalty_cost_ouom, cont_oper_proc.Expl_penalty_percent, cont_oper_proc.Inflation_adjustment_ind, cont_oper_proc.Insurance_election, cont_oper_proc.Mktg_fee_elect_cost, cont_oper_proc.Mktg_fee_elect_cost_ouom, cont_oper_proc.Mktg_fee_elect_percent, cont_oper_proc.Nat_currency_conversion, cont_oper_proc.Nat_currency_uom, cont_oper_proc.Operating_procedure_type, cont_oper_proc.Ppdm_guid, cont_oper_proc.Recog_on_assignment_ind, cont_oper_proc.Remark, cont_oper_proc.Rofr_ind, cont_oper_proc.Rofr_reply_term, cont_oper_proc.Rofr_reply_term_uom, cont_oper_proc.Source, cont_oper_proc.Source_document_id, cont_oper_proc.Statute_limit_period, cont_oper_proc.Statute_limit_period_uom, cont_oper_proc.Surrender_notice, cont_oper_proc.Surrender_notice_uom, cont_oper_proc.Title_preserve_desc, cont_oper_proc.Unauthorized_amount, cont_oper_proc.Unauthorized_amount_ouom, cont_oper_proc.Row_changed_by, cont_oper_proc.Row_changed_date, cont_oper_proc.Row_created_by, cont_oper_proc.Row_created_date, cont_oper_proc.Row_effective_date, cont_oper_proc.Row_expiry_date, cont_oper_proc.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateContOperProc(c *fiber.Ctx) error {
	var cont_oper_proc dto.Cont_oper_proc

	setDefaults(&cont_oper_proc)

	if err := json.Unmarshal(c.Body(), &cont_oper_proc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cont_oper_proc.Contract_id = id

    if cont_oper_proc.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cont_oper_proc set operating_procedure_id = :1, active_ind = :2, adjust_13_month_desc = :3, adjust_13_month_ind = :4, casing_point_election = :5, claim_amount_limit = :6, claim_amount_limit_ouom = :7, dev_penalty_cost = :8, dev_penalty_cost_ouom = :9, dev_penalty_percent = :10, effective_date = :11, equip_penalty_cost = :12, equip_penalty_cost_ouom = :13, equip_penalty_percent = :14, expiry_date = :15, expl_penalty_cost = :16, expl_penalty_cost_ouom = :17, expl_penalty_percent = :18, inflation_adjustment_ind = :19, insurance_election = :20, mktg_fee_elect_cost = :21, mktg_fee_elect_cost_ouom = :22, mktg_fee_elect_percent = :23, nat_currency_conversion = :24, nat_currency_uom = :25, operating_procedure_type = :26, ppdm_guid = :27, recog_on_assignment_ind = :28, remark = :29, rofr_ind = :30, rofr_reply_term = :31, rofr_reply_term_uom = :32, source = :33, source_document_id = :34, statute_limit_period = :35, statute_limit_period_uom = :36, surrender_notice = :37, surrender_notice_uom = :38, title_preserve_desc = :39, unauthorized_amount = :40, unauthorized_amount_ouom = :41, row_changed_by = :42, row_changed_date = :43, row_created_by = :44, row_effective_date = :45, row_expiry_date = :46, row_quality = :47 where contract_id = :49")
	if err != nil {
		return err
	}

	cont_oper_proc.Row_changed_date = formatDate(cont_oper_proc.Row_changed_date)
	cont_oper_proc.Effective_date = formatDateString(cont_oper_proc.Effective_date)
	cont_oper_proc.Expiry_date = formatDateString(cont_oper_proc.Expiry_date)
	cont_oper_proc.Row_effective_date = formatDateString(cont_oper_proc.Row_effective_date)
	cont_oper_proc.Row_expiry_date = formatDateString(cont_oper_proc.Row_expiry_date)






	rows, err := stmt.Exec(cont_oper_proc.Operating_procedure_id, cont_oper_proc.Active_ind, cont_oper_proc.Adjust_13_month_desc, cont_oper_proc.Adjust_13_month_ind, cont_oper_proc.Casing_point_election, cont_oper_proc.Claim_amount_limit, cont_oper_proc.Claim_amount_limit_ouom, cont_oper_proc.Dev_penalty_cost, cont_oper_proc.Dev_penalty_cost_ouom, cont_oper_proc.Dev_penalty_percent, cont_oper_proc.Effective_date, cont_oper_proc.Equip_penalty_cost, cont_oper_proc.Equip_penalty_cost_ouom, cont_oper_proc.Equip_penalty_percent, cont_oper_proc.Expiry_date, cont_oper_proc.Expl_penalty_cost, cont_oper_proc.Expl_penalty_cost_ouom, cont_oper_proc.Expl_penalty_percent, cont_oper_proc.Inflation_adjustment_ind, cont_oper_proc.Insurance_election, cont_oper_proc.Mktg_fee_elect_cost, cont_oper_proc.Mktg_fee_elect_cost_ouom, cont_oper_proc.Mktg_fee_elect_percent, cont_oper_proc.Nat_currency_conversion, cont_oper_proc.Nat_currency_uom, cont_oper_proc.Operating_procedure_type, cont_oper_proc.Ppdm_guid, cont_oper_proc.Recog_on_assignment_ind, cont_oper_proc.Remark, cont_oper_proc.Rofr_ind, cont_oper_proc.Rofr_reply_term, cont_oper_proc.Rofr_reply_term_uom, cont_oper_proc.Source, cont_oper_proc.Source_document_id, cont_oper_proc.Statute_limit_period, cont_oper_proc.Statute_limit_period_uom, cont_oper_proc.Surrender_notice, cont_oper_proc.Surrender_notice_uom, cont_oper_proc.Title_preserve_desc, cont_oper_proc.Unauthorized_amount, cont_oper_proc.Unauthorized_amount_ouom, cont_oper_proc.Row_changed_by, cont_oper_proc.Row_changed_date, cont_oper_proc.Row_created_by, cont_oper_proc.Row_effective_date, cont_oper_proc.Row_expiry_date, cont_oper_proc.Row_quality, cont_oper_proc.Contract_id)
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

func PatchContOperProc(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cont_oper_proc set "
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

func DeleteContOperProc(c *fiber.Ctx) error {
	id := c.Params("id")
	var cont_oper_proc dto.Cont_oper_proc
	cont_oper_proc.Contract_id = id

	stmt, err := db.Prepare("delete from cont_oper_proc where contract_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cont_oper_proc.Contract_id)
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



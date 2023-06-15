package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetContract(c *fiber.Ctx) error {
	rows, err := db.Query("select * from contract")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Contract

	for rows.Next() {
		var contract dto.Contract
		if err := rows.Scan(&contract.Contract_id, &contract.Active_ind, &contract.Ami_aoe_ind, &contract.Area_id, &contract.Area_type, &contract.Assignment_proc_ind, &contract.Casing_point_election, &contract.Closing_date, &contract.Confidential_ind, &contract.Contract_name, &contract.Contract_num, &contract.Currency_conversion, &contract.Currency_ouom, &contract.Effective_date, &contract.Expiry_date, &contract.Extension_condition, &contract.Governing_contract_ind, &contract.Liability_period, &contract.Liability_period_ouom, &contract.Liability_release_date, &contract.Nat_currency_conversion, &contract.Nat_currency_uom, &contract.Operating_proc_ind, &contract.Ppdm_guid, &contract.Primary_term, &contract.Primary_term_ouom, &contract.Remark, &contract.Rofr_ind, &contract.Secondary_term, &contract.Secondary_term_ouom, &contract.Source, &contract.Source_document, &contract.Surrender_notice_term, &contract.Surrender_notice_term_ouom, &contract.Termination_date, &contract.Voting_proc_ind, &contract.Row_changed_by, &contract.Row_changed_date, &contract.Row_created_by, &contract.Row_created_date, &contract.Row_effective_date, &contract.Row_expiry_date, &contract.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append contract to result
		result = append(result, contract)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetContract(c *fiber.Ctx) error {
	var contract dto.Contract

	setDefaults(&contract)

	if err := json.Unmarshal(c.Body(), &contract); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into contract values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43)")
	if err != nil {
		return err
	}
	contract.Row_created_date = formatDate(contract.Row_created_date)
	contract.Row_changed_date = formatDate(contract.Row_changed_date)
	contract.Closing_date = formatDateString(contract.Closing_date)
	contract.Effective_date = formatDateString(contract.Effective_date)
	contract.Expiry_date = formatDateString(contract.Expiry_date)
	contract.Liability_release_date = formatDateString(contract.Liability_release_date)
	contract.Termination_date = formatDateString(contract.Termination_date)
	contract.Row_effective_date = formatDateString(contract.Row_effective_date)
	contract.Row_expiry_date = formatDateString(contract.Row_expiry_date)






	rows, err := stmt.Exec(contract.Contract_id, contract.Active_ind, contract.Ami_aoe_ind, contract.Area_id, contract.Area_type, contract.Assignment_proc_ind, contract.Casing_point_election, contract.Closing_date, contract.Confidential_ind, contract.Contract_name, contract.Contract_num, contract.Currency_conversion, contract.Currency_ouom, contract.Effective_date, contract.Expiry_date, contract.Extension_condition, contract.Governing_contract_ind, contract.Liability_period, contract.Liability_period_ouom, contract.Liability_release_date, contract.Nat_currency_conversion, contract.Nat_currency_uom, contract.Operating_proc_ind, contract.Ppdm_guid, contract.Primary_term, contract.Primary_term_ouom, contract.Remark, contract.Rofr_ind, contract.Secondary_term, contract.Secondary_term_ouom, contract.Source, contract.Source_document, contract.Surrender_notice_term, contract.Surrender_notice_term_ouom, contract.Termination_date, contract.Voting_proc_ind, contract.Row_changed_by, contract.Row_changed_date, contract.Row_created_by, contract.Row_created_date, contract.Row_effective_date, contract.Row_expiry_date, contract.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateContract(c *fiber.Ctx) error {
	var contract dto.Contract

	setDefaults(&contract)

	if err := json.Unmarshal(c.Body(), &contract); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	contract.Contract_id = id

    if contract.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update contract set active_ind = :1, ami_aoe_ind = :2, area_id = :3, area_type = :4, assignment_proc_ind = :5, casing_point_election = :6, closing_date = :7, confidential_ind = :8, contract_name = :9, contract_num = :10, currency_conversion = :11, currency_ouom = :12, effective_date = :13, expiry_date = :14, extension_condition = :15, governing_contract_ind = :16, liability_period = :17, liability_period_ouom = :18, liability_release_date = :19, nat_currency_conversion = :20, nat_currency_uom = :21, operating_proc_ind = :22, ppdm_guid = :23, primary_term = :24, primary_term_ouom = :25, remark = :26, rofr_ind = :27, secondary_term = :28, secondary_term_ouom = :29, source = :30, source_document = :31, surrender_notice_term = :32, surrender_notice_term_ouom = :33, termination_date = :34, voting_proc_ind = :35, row_changed_by = :36, row_changed_date = :37, row_created_by = :38, row_effective_date = :39, row_expiry_date = :40, row_quality = :41 where contract_id = :43")
	if err != nil {
		return err
	}

	contract.Row_changed_date = formatDate(contract.Row_changed_date)
	contract.Closing_date = formatDateString(contract.Closing_date)
	contract.Effective_date = formatDateString(contract.Effective_date)
	contract.Expiry_date = formatDateString(contract.Expiry_date)
	contract.Liability_release_date = formatDateString(contract.Liability_release_date)
	contract.Termination_date = formatDateString(contract.Termination_date)
	contract.Row_effective_date = formatDateString(contract.Row_effective_date)
	contract.Row_expiry_date = formatDateString(contract.Row_expiry_date)






	rows, err := stmt.Exec(contract.Active_ind, contract.Ami_aoe_ind, contract.Area_id, contract.Area_type, contract.Assignment_proc_ind, contract.Casing_point_election, contract.Closing_date, contract.Confidential_ind, contract.Contract_name, contract.Contract_num, contract.Currency_conversion, contract.Currency_ouom, contract.Effective_date, contract.Expiry_date, contract.Extension_condition, contract.Governing_contract_ind, contract.Liability_period, contract.Liability_period_ouom, contract.Liability_release_date, contract.Nat_currency_conversion, contract.Nat_currency_uom, contract.Operating_proc_ind, contract.Ppdm_guid, contract.Primary_term, contract.Primary_term_ouom, contract.Remark, contract.Rofr_ind, contract.Secondary_term, contract.Secondary_term_ouom, contract.Source, contract.Source_document, contract.Surrender_notice_term, contract.Surrender_notice_term_ouom, contract.Termination_date, contract.Voting_proc_ind, contract.Row_changed_by, contract.Row_changed_date, contract.Row_created_by, contract.Row_effective_date, contract.Row_expiry_date, contract.Row_quality, contract.Contract_id)
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

func PatchContract(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update contract set "
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
		} else if key == "closing_date" || key == "effective_date" || key == "expiry_date" || key == "liability_release_date" || key == "termination_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteContract(c *fiber.Ctx) error {
	id := c.Params("id")
	var contract dto.Contract
	contract.Contract_id = id

	stmt, err := db.Prepare("delete from contract where contract_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(contract.Contract_id)
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



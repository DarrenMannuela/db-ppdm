package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetIntSetPartner(c *fiber.Ctx) error {
	rows, err := db.Query("select * from int_set_partner")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Int_set_partner

	for rows.Next() {
		var int_set_partner dto.Int_set_partner
		if err := rows.Scan(&int_set_partner.Interest_set_id, &int_set_partner.Interest_set_seq_no, &int_set_partner.Partner_ba_id, &int_set_partner.Partner_obs_no, &int_set_partner.Active_ind, &int_set_partner.Address_obs_no, &int_set_partner.Address_source, &int_set_partner.Breach_desc, &int_set_partner.Breach_ind, &int_set_partner.Change_reason, &int_set_partner.Confidential_ind, &int_set_partner.Contract_id, &int_set_partner.Effective_date, &int_set_partner.Escrow_desc, &int_set_partner.Escrow_ind, &int_set_partner.Expiry_date, &int_set_partner.Gross_fract_interest, &int_set_partner.Gross_percent_interest, &int_set_partner.Inactive_date, &int_set_partner.Instrument_id, &int_set_partner.Interest_set_role, &int_set_partner.Net_fract_interest, &int_set_partner.Net_percent_interest, &int_set_partner.Penalty_ind, &int_set_partner.Ppdm_guid, &int_set_partner.Provision_id, &int_set_partner.Remark, &int_set_partner.Right_to_earn_desc, &int_set_partner.Right_to_earn_ind, &int_set_partner.Source, &int_set_partner.Title_own_type, &int_set_partner.Unit_operated_ind, &int_set_partner.Row_changed_by, &int_set_partner.Row_changed_date, &int_set_partner.Row_created_by, &int_set_partner.Row_created_date, &int_set_partner.Row_effective_date, &int_set_partner.Row_expiry_date, &int_set_partner.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append int_set_partner to result
		result = append(result, int_set_partner)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetIntSetPartner(c *fiber.Ctx) error {
	var int_set_partner dto.Int_set_partner

	setDefaults(&int_set_partner)

	if err := json.Unmarshal(c.Body(), &int_set_partner); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into int_set_partner values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39)")
	if err != nil {
		return err
	}
	int_set_partner.Row_created_date = formatDate(int_set_partner.Row_created_date)
	int_set_partner.Row_changed_date = formatDate(int_set_partner.Row_changed_date)
	int_set_partner.Effective_date = formatDateString(int_set_partner.Effective_date)
	int_set_partner.Expiry_date = formatDateString(int_set_partner.Expiry_date)
	int_set_partner.Inactive_date = formatDateString(int_set_partner.Inactive_date)
	int_set_partner.Row_effective_date = formatDateString(int_set_partner.Row_effective_date)
	int_set_partner.Row_expiry_date = formatDateString(int_set_partner.Row_expiry_date)






	rows, err := stmt.Exec(int_set_partner.Interest_set_id, int_set_partner.Interest_set_seq_no, int_set_partner.Partner_ba_id, int_set_partner.Partner_obs_no, int_set_partner.Active_ind, int_set_partner.Address_obs_no, int_set_partner.Address_source, int_set_partner.Breach_desc, int_set_partner.Breach_ind, int_set_partner.Change_reason, int_set_partner.Confidential_ind, int_set_partner.Contract_id, int_set_partner.Effective_date, int_set_partner.Escrow_desc, int_set_partner.Escrow_ind, int_set_partner.Expiry_date, int_set_partner.Gross_fract_interest, int_set_partner.Gross_percent_interest, int_set_partner.Inactive_date, int_set_partner.Instrument_id, int_set_partner.Interest_set_role, int_set_partner.Net_fract_interest, int_set_partner.Net_percent_interest, int_set_partner.Penalty_ind, int_set_partner.Ppdm_guid, int_set_partner.Provision_id, int_set_partner.Remark, int_set_partner.Right_to_earn_desc, int_set_partner.Right_to_earn_ind, int_set_partner.Source, int_set_partner.Title_own_type, int_set_partner.Unit_operated_ind, int_set_partner.Row_changed_by, int_set_partner.Row_changed_date, int_set_partner.Row_created_by, int_set_partner.Row_created_date, int_set_partner.Row_effective_date, int_set_partner.Row_expiry_date, int_set_partner.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateIntSetPartner(c *fiber.Ctx) error {
	var int_set_partner dto.Int_set_partner

	setDefaults(&int_set_partner)

	if err := json.Unmarshal(c.Body(), &int_set_partner); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	int_set_partner.Interest_set_id = id

    if int_set_partner.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update int_set_partner set interest_set_seq_no = :1, partner_ba_id = :2, partner_obs_no = :3, active_ind = :4, address_obs_no = :5, address_source = :6, breach_desc = :7, breach_ind = :8, change_reason = :9, confidential_ind = :10, contract_id = :11, effective_date = :12, escrow_desc = :13, escrow_ind = :14, expiry_date = :15, gross_fract_interest = :16, gross_percent_interest = :17, inactive_date = :18, instrument_id = :19, interest_set_role = :20, net_fract_interest = :21, net_percent_interest = :22, penalty_ind = :23, ppdm_guid = :24, provision_id = :25, remark = :26, right_to_earn_desc = :27, right_to_earn_ind = :28, source = :29, title_own_type = :30, unit_operated_ind = :31, row_changed_by = :32, row_changed_date = :33, row_created_by = :34, row_effective_date = :35, row_expiry_date = :36, row_quality = :37 where interest_set_id = :39")
	if err != nil {
		return err
	}

	int_set_partner.Row_changed_date = formatDate(int_set_partner.Row_changed_date)
	int_set_partner.Effective_date = formatDateString(int_set_partner.Effective_date)
	int_set_partner.Expiry_date = formatDateString(int_set_partner.Expiry_date)
	int_set_partner.Inactive_date = formatDateString(int_set_partner.Inactive_date)
	int_set_partner.Row_effective_date = formatDateString(int_set_partner.Row_effective_date)
	int_set_partner.Row_expiry_date = formatDateString(int_set_partner.Row_expiry_date)






	rows, err := stmt.Exec(int_set_partner.Interest_set_seq_no, int_set_partner.Partner_ba_id, int_set_partner.Partner_obs_no, int_set_partner.Active_ind, int_set_partner.Address_obs_no, int_set_partner.Address_source, int_set_partner.Breach_desc, int_set_partner.Breach_ind, int_set_partner.Change_reason, int_set_partner.Confidential_ind, int_set_partner.Contract_id, int_set_partner.Effective_date, int_set_partner.Escrow_desc, int_set_partner.Escrow_ind, int_set_partner.Expiry_date, int_set_partner.Gross_fract_interest, int_set_partner.Gross_percent_interest, int_set_partner.Inactive_date, int_set_partner.Instrument_id, int_set_partner.Interest_set_role, int_set_partner.Net_fract_interest, int_set_partner.Net_percent_interest, int_set_partner.Penalty_ind, int_set_partner.Ppdm_guid, int_set_partner.Provision_id, int_set_partner.Remark, int_set_partner.Right_to_earn_desc, int_set_partner.Right_to_earn_ind, int_set_partner.Source, int_set_partner.Title_own_type, int_set_partner.Unit_operated_ind, int_set_partner.Row_changed_by, int_set_partner.Row_changed_date, int_set_partner.Row_created_by, int_set_partner.Row_effective_date, int_set_partner.Row_expiry_date, int_set_partner.Row_quality, int_set_partner.Interest_set_id)
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

func PatchIntSetPartner(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update int_set_partner set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "inactive_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where interest_set_id = :id"

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

func DeleteIntSetPartner(c *fiber.Ctx) error {
	id := c.Params("id")
	var int_set_partner dto.Int_set_partner
	int_set_partner.Interest_set_id = id

	stmt, err := db.Prepare("delete from int_set_partner where interest_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(int_set_partner.Interest_set_id)
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



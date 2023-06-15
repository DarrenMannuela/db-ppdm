package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetResentVolSummary(c *fiber.Ctx) error {
	rows, err := db.Query("select * from resent_vol_summary")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Resent_vol_summary

	for rows.Next() {
		var resent_vol_summary dto.Resent_vol_summary
		if err := rows.Scan(&resent_vol_summary.Resent_id, &resent_vol_summary.Reserve_class_id, &resent_vol_summary.Product_type, &resent_vol_summary.Summary_id, &resent_vol_summary.Summary_obs_no, &resent_vol_summary.Active_ind, &resent_vol_summary.Analyst_ba_id, &resent_vol_summary.Approved_by_ba_id, &resent_vol_summary.Approved_date, &resent_vol_summary.Created_date, &resent_vol_summary.Created_date_desc, &resent_vol_summary.Current_balance, &resent_vol_summary.Current_balance_ouom, &resent_vol_summary.Current_balance_uom, &resent_vol_summary.Decline_case_id, &resent_vol_summary.Effective_date, &resent_vol_summary.Expiry_date, &resent_vol_summary.Gross_summary_ind, &resent_vol_summary.Interest_set_id, &resent_vol_summary.Interest_set_partner, &resent_vol_summary.Interest_set_seq_no, &resent_vol_summary.Material_balance_case_id, &resent_vol_summary.Net_summary_ind, &resent_vol_summary.Open_balance, &resent_vol_summary.Open_balance_ouom, &resent_vol_summary.Open_balance_uom, &resent_vol_summary.Partner_obs_no, &resent_vol_summary.Pden_id, &resent_vol_summary.Pden_product_type, &resent_vol_summary.Pden_source, &resent_vol_summary.Pden_subtype, &resent_vol_summary.Ppdm_guid, &resent_vol_summary.Remark, &resent_vol_summary.Report_ind, &resent_vol_summary.Source, &resent_vol_summary.Volume_method, &resent_vol_summary.Vol_anal_case_id, &resent_vol_summary.Yield_parent_product, &resent_vol_summary.Yield_rate, &resent_vol_summary.Yield_rate_ouom, &resent_vol_summary.Yield_rate_uom, &resent_vol_summary.Row_changed_by, &resent_vol_summary.Row_changed_date, &resent_vol_summary.Row_created_by, &resent_vol_summary.Row_created_date, &resent_vol_summary.Row_effective_date, &resent_vol_summary.Row_expiry_date, &resent_vol_summary.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append resent_vol_summary to result
		result = append(result, resent_vol_summary)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetResentVolSummary(c *fiber.Ctx) error {
	var resent_vol_summary dto.Resent_vol_summary

	setDefaults(&resent_vol_summary)

	if err := json.Unmarshal(c.Body(), &resent_vol_summary); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into resent_vol_summary values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39, :40, :41, :42, :43, :44, :45, :46, :47, :48)")
	if err != nil {
		return err
	}
	resent_vol_summary.Row_created_date = formatDate(resent_vol_summary.Row_created_date)
	resent_vol_summary.Row_changed_date = formatDate(resent_vol_summary.Row_changed_date)
	resent_vol_summary.Approved_date = formatDateString(resent_vol_summary.Approved_date)
	resent_vol_summary.Created_date = formatDateString(resent_vol_summary.Created_date)
	resent_vol_summary.Effective_date = formatDateString(resent_vol_summary.Effective_date)
	resent_vol_summary.Expiry_date = formatDateString(resent_vol_summary.Expiry_date)
	resent_vol_summary.Row_effective_date = formatDateString(resent_vol_summary.Row_effective_date)
	resent_vol_summary.Row_expiry_date = formatDateString(resent_vol_summary.Row_expiry_date)






	rows, err := stmt.Exec(resent_vol_summary.Resent_id, resent_vol_summary.Reserve_class_id, resent_vol_summary.Product_type, resent_vol_summary.Summary_id, resent_vol_summary.Summary_obs_no, resent_vol_summary.Active_ind, resent_vol_summary.Analyst_ba_id, resent_vol_summary.Approved_by_ba_id, resent_vol_summary.Approved_date, resent_vol_summary.Created_date, resent_vol_summary.Created_date_desc, resent_vol_summary.Current_balance, resent_vol_summary.Current_balance_ouom, resent_vol_summary.Current_balance_uom, resent_vol_summary.Decline_case_id, resent_vol_summary.Effective_date, resent_vol_summary.Expiry_date, resent_vol_summary.Gross_summary_ind, resent_vol_summary.Interest_set_id, resent_vol_summary.Interest_set_partner, resent_vol_summary.Interest_set_seq_no, resent_vol_summary.Material_balance_case_id, resent_vol_summary.Net_summary_ind, resent_vol_summary.Open_balance, resent_vol_summary.Open_balance_ouom, resent_vol_summary.Open_balance_uom, resent_vol_summary.Partner_obs_no, resent_vol_summary.Pden_id, resent_vol_summary.Pden_product_type, resent_vol_summary.Pden_source, resent_vol_summary.Pden_subtype, resent_vol_summary.Ppdm_guid, resent_vol_summary.Remark, resent_vol_summary.Report_ind, resent_vol_summary.Source, resent_vol_summary.Volume_method, resent_vol_summary.Vol_anal_case_id, resent_vol_summary.Yield_parent_product, resent_vol_summary.Yield_rate, resent_vol_summary.Yield_rate_ouom, resent_vol_summary.Yield_rate_uom, resent_vol_summary.Row_changed_by, resent_vol_summary.Row_changed_date, resent_vol_summary.Row_created_by, resent_vol_summary.Row_created_date, resent_vol_summary.Row_effective_date, resent_vol_summary.Row_expiry_date, resent_vol_summary.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateResentVolSummary(c *fiber.Ctx) error {
	var resent_vol_summary dto.Resent_vol_summary

	setDefaults(&resent_vol_summary)

	if err := json.Unmarshal(c.Body(), &resent_vol_summary); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	resent_vol_summary.Resent_id = id

    if resent_vol_summary.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update resent_vol_summary set reserve_class_id = :1, product_type = :2, summary_id = :3, summary_obs_no = :4, active_ind = :5, analyst_ba_id = :6, approved_by_ba_id = :7, approved_date = :8, created_date = :9, created_date_desc = :10, current_balance = :11, current_balance_ouom = :12, current_balance_uom = :13, decline_case_id = :14, effective_date = :15, expiry_date = :16, gross_summary_ind = :17, interest_set_id = :18, interest_set_partner = :19, interest_set_seq_no = :20, material_balance_case_id = :21, net_summary_ind = :22, open_balance = :23, open_balance_ouom = :24, open_balance_uom = :25, partner_obs_no = :26, pden_id = :27, pden_product_type = :28, pden_source = :29, pden_subtype = :30, ppdm_guid = :31, remark = :32, report_ind = :33, source = :34, volume_method = :35, vol_anal_case_id = :36, yield_parent_product = :37, yield_rate = :38, yield_rate_ouom = :39, yield_rate_uom = :40, row_changed_by = :41, row_changed_date = :42, row_created_by = :43, row_effective_date = :44, row_expiry_date = :45, row_quality = :46 where resent_id = :48")
	if err != nil {
		return err
	}

	resent_vol_summary.Row_changed_date = formatDate(resent_vol_summary.Row_changed_date)
	resent_vol_summary.Approved_date = formatDateString(resent_vol_summary.Approved_date)
	resent_vol_summary.Created_date = formatDateString(resent_vol_summary.Created_date)
	resent_vol_summary.Effective_date = formatDateString(resent_vol_summary.Effective_date)
	resent_vol_summary.Expiry_date = formatDateString(resent_vol_summary.Expiry_date)
	resent_vol_summary.Row_effective_date = formatDateString(resent_vol_summary.Row_effective_date)
	resent_vol_summary.Row_expiry_date = formatDateString(resent_vol_summary.Row_expiry_date)






	rows, err := stmt.Exec(resent_vol_summary.Reserve_class_id, resent_vol_summary.Product_type, resent_vol_summary.Summary_id, resent_vol_summary.Summary_obs_no, resent_vol_summary.Active_ind, resent_vol_summary.Analyst_ba_id, resent_vol_summary.Approved_by_ba_id, resent_vol_summary.Approved_date, resent_vol_summary.Created_date, resent_vol_summary.Created_date_desc, resent_vol_summary.Current_balance, resent_vol_summary.Current_balance_ouom, resent_vol_summary.Current_balance_uom, resent_vol_summary.Decline_case_id, resent_vol_summary.Effective_date, resent_vol_summary.Expiry_date, resent_vol_summary.Gross_summary_ind, resent_vol_summary.Interest_set_id, resent_vol_summary.Interest_set_partner, resent_vol_summary.Interest_set_seq_no, resent_vol_summary.Material_balance_case_id, resent_vol_summary.Net_summary_ind, resent_vol_summary.Open_balance, resent_vol_summary.Open_balance_ouom, resent_vol_summary.Open_balance_uom, resent_vol_summary.Partner_obs_no, resent_vol_summary.Pden_id, resent_vol_summary.Pden_product_type, resent_vol_summary.Pden_source, resent_vol_summary.Pden_subtype, resent_vol_summary.Ppdm_guid, resent_vol_summary.Remark, resent_vol_summary.Report_ind, resent_vol_summary.Source, resent_vol_summary.Volume_method, resent_vol_summary.Vol_anal_case_id, resent_vol_summary.Yield_parent_product, resent_vol_summary.Yield_rate, resent_vol_summary.Yield_rate_ouom, resent_vol_summary.Yield_rate_uom, resent_vol_summary.Row_changed_by, resent_vol_summary.Row_changed_date, resent_vol_summary.Row_created_by, resent_vol_summary.Row_effective_date, resent_vol_summary.Row_expiry_date, resent_vol_summary.Row_quality, resent_vol_summary.Resent_id)
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

func PatchResentVolSummary(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update resent_vol_summary set "
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
		} else if key == "approved_date" || key == "created_date" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where resent_id = :id"

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

func DeleteResentVolSummary(c *fiber.Ctx) error {
	id := c.Params("id")
	var resent_vol_summary dto.Resent_vol_summary
	resent_vol_summary.Resent_id = id

	stmt, err := db.Prepare("delete from resent_vol_summary where resent_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(resent_vol_summary.Resent_id)
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



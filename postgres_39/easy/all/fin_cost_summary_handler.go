package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetFinCostSummary(c *fiber.Ctx) error {
	rows, err := db.Query("select * from fin_cost_summary")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Fin_cost_summary

	for rows.Next() {
		var fin_cost_summary dto.Fin_cost_summary
		if err := rows.Scan(&fin_cost_summary.Cost_id, &fin_cost_summary.Active_ind, &fin_cost_summary.Ami_ind, &fin_cost_summary.Confidential_ind, &fin_cost_summary.Cost_amount, &fin_cost_summary.Cost_per_size, &fin_cost_summary.Cost_per_size_ouom, &fin_cost_summary.Cost_type, &fin_cost_summary.Currency_conversion, &fin_cost_summary.Currency_ouom, &fin_cost_summary.Effective_date, &fin_cost_summary.Expiry_date, &fin_cost_summary.Finance_component_id, &fin_cost_summary.Finance_id, &fin_cost_summary.Paid_date, &fin_cost_summary.Parent_cost_id, &fin_cost_summary.Percent_of_parent, &fin_cost_summary.Ppdm_guid, &fin_cost_summary.Remark, &fin_cost_summary.Report_cost_ind, &fin_cost_summary.Source, &fin_cost_summary.Submit_date, &fin_cost_summary.Row_changed_by, &fin_cost_summary.Row_changed_date, &fin_cost_summary.Row_created_by, &fin_cost_summary.Row_created_date, &fin_cost_summary.Row_effective_date, &fin_cost_summary.Row_expiry_date, &fin_cost_summary.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append fin_cost_summary to result
		result = append(result, fin_cost_summary)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetFinCostSummary(c *fiber.Ctx) error {
	var fin_cost_summary dto.Fin_cost_summary

	setDefaults(&fin_cost_summary)

	if err := json.Unmarshal(c.Body(), &fin_cost_summary); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into fin_cost_summary values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29)")
	if err != nil {
		return err
	}
	fin_cost_summary.Row_created_date = formatDate(fin_cost_summary.Row_created_date)
	fin_cost_summary.Row_changed_date = formatDate(fin_cost_summary.Row_changed_date)
	fin_cost_summary.Effective_date = formatDateString(fin_cost_summary.Effective_date)
	fin_cost_summary.Expiry_date = formatDateString(fin_cost_summary.Expiry_date)
	fin_cost_summary.Paid_date = formatDateString(fin_cost_summary.Paid_date)
	fin_cost_summary.Submit_date = formatDateString(fin_cost_summary.Submit_date)
	fin_cost_summary.Row_effective_date = formatDateString(fin_cost_summary.Row_effective_date)
	fin_cost_summary.Row_expiry_date = formatDateString(fin_cost_summary.Row_expiry_date)






	rows, err := stmt.Exec(fin_cost_summary.Cost_id, fin_cost_summary.Active_ind, fin_cost_summary.Ami_ind, fin_cost_summary.Confidential_ind, fin_cost_summary.Cost_amount, fin_cost_summary.Cost_per_size, fin_cost_summary.Cost_per_size_ouom, fin_cost_summary.Cost_type, fin_cost_summary.Currency_conversion, fin_cost_summary.Currency_ouom, fin_cost_summary.Effective_date, fin_cost_summary.Expiry_date, fin_cost_summary.Finance_component_id, fin_cost_summary.Finance_id, fin_cost_summary.Paid_date, fin_cost_summary.Parent_cost_id, fin_cost_summary.Percent_of_parent, fin_cost_summary.Ppdm_guid, fin_cost_summary.Remark, fin_cost_summary.Report_cost_ind, fin_cost_summary.Source, fin_cost_summary.Submit_date, fin_cost_summary.Row_changed_by, fin_cost_summary.Row_changed_date, fin_cost_summary.Row_created_by, fin_cost_summary.Row_created_date, fin_cost_summary.Row_effective_date, fin_cost_summary.Row_expiry_date, fin_cost_summary.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateFinCostSummary(c *fiber.Ctx) error {
	var fin_cost_summary dto.Fin_cost_summary

	setDefaults(&fin_cost_summary)

	if err := json.Unmarshal(c.Body(), &fin_cost_summary); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	fin_cost_summary.Cost_id = id

    if fin_cost_summary.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update fin_cost_summary set active_ind = :1, ami_ind = :2, confidential_ind = :3, cost_amount = :4, cost_per_size = :5, cost_per_size_ouom = :6, cost_type = :7, currency_conversion = :8, currency_ouom = :9, effective_date = :10, expiry_date = :11, finance_component_id = :12, finance_id = :13, paid_date = :14, parent_cost_id = :15, percent_of_parent = :16, ppdm_guid = :17, remark = :18, report_cost_ind = :19, source = :20, submit_date = :21, row_changed_by = :22, row_changed_date = :23, row_created_by = :24, row_effective_date = :25, row_expiry_date = :26, row_quality = :27 where cost_id = :29")
	if err != nil {
		return err
	}

	fin_cost_summary.Row_changed_date = formatDate(fin_cost_summary.Row_changed_date)
	fin_cost_summary.Effective_date = formatDateString(fin_cost_summary.Effective_date)
	fin_cost_summary.Expiry_date = formatDateString(fin_cost_summary.Expiry_date)
	fin_cost_summary.Paid_date = formatDateString(fin_cost_summary.Paid_date)
	fin_cost_summary.Submit_date = formatDateString(fin_cost_summary.Submit_date)
	fin_cost_summary.Row_effective_date = formatDateString(fin_cost_summary.Row_effective_date)
	fin_cost_summary.Row_expiry_date = formatDateString(fin_cost_summary.Row_expiry_date)






	rows, err := stmt.Exec(fin_cost_summary.Active_ind, fin_cost_summary.Ami_ind, fin_cost_summary.Confidential_ind, fin_cost_summary.Cost_amount, fin_cost_summary.Cost_per_size, fin_cost_summary.Cost_per_size_ouom, fin_cost_summary.Cost_type, fin_cost_summary.Currency_conversion, fin_cost_summary.Currency_ouom, fin_cost_summary.Effective_date, fin_cost_summary.Expiry_date, fin_cost_summary.Finance_component_id, fin_cost_summary.Finance_id, fin_cost_summary.Paid_date, fin_cost_summary.Parent_cost_id, fin_cost_summary.Percent_of_parent, fin_cost_summary.Ppdm_guid, fin_cost_summary.Remark, fin_cost_summary.Report_cost_ind, fin_cost_summary.Source, fin_cost_summary.Submit_date, fin_cost_summary.Row_changed_by, fin_cost_summary.Row_changed_date, fin_cost_summary.Row_created_by, fin_cost_summary.Row_effective_date, fin_cost_summary.Row_expiry_date, fin_cost_summary.Row_quality, fin_cost_summary.Cost_id)
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

func PatchFinCostSummary(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update fin_cost_summary set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "paid_date" || key == "submit_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where cost_id = :id"

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

func DeleteFinCostSummary(c *fiber.Ctx) error {
	id := c.Params("id")
	var fin_cost_summary dto.Fin_cost_summary
	fin_cost_summary.Cost_id = id

	stmt, err := db.Prepare("delete from fin_cost_summary where cost_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(fin_cost_summary.Cost_id)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPaleoSummary(c *fiber.Ctx) error {
	rows, err := db.Query("select * from paleo_summary")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Paleo_summary

	for rows.Next() {
		var paleo_summary dto.Paleo_summary
		if err := rows.Scan(&paleo_summary.Paleo_summary_id, &paleo_summary.Access_condition, &paleo_summary.Active_ind, &paleo_summary.Analysis_end_date, &paleo_summary.Analysis_start_date, &paleo_summary.Confidential_ind, &paleo_summary.Confidential_reason, &paleo_summary.Confidential_release_date, &paleo_summary.Confidential_term, &paleo_summary.Confidential_type, &paleo_summary.Diversity_count, &paleo_summary.Effective_date, &paleo_summary.Expiry_date, &paleo_summary.Ppdm_guid, &paleo_summary.Preservation_quality, &paleo_summary.Reference_num, &paleo_summary.Remark, &paleo_summary.Report_date, &paleo_summary.Report_title, &paleo_summary.Source, &paleo_summary.Total_sample_interval, &paleo_summary.Total_sample_interval_ouom, &paleo_summary.Row_changed_by, &paleo_summary.Row_changed_date, &paleo_summary.Row_created_by, &paleo_summary.Row_created_date, &paleo_summary.Row_effective_date, &paleo_summary.Row_expiry_date, &paleo_summary.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append paleo_summary to result
		result = append(result, paleo_summary)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPaleoSummary(c *fiber.Ctx) error {
	var paleo_summary dto.Paleo_summary

	setDefaults(&paleo_summary)

	if err := json.Unmarshal(c.Body(), &paleo_summary); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into paleo_summary values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29)")
	if err != nil {
		return err
	}
	paleo_summary.Row_created_date = formatDate(paleo_summary.Row_created_date)
	paleo_summary.Row_changed_date = formatDate(paleo_summary.Row_changed_date)
	paleo_summary.Analysis_end_date = formatDateString(paleo_summary.Analysis_end_date)
	paleo_summary.Analysis_start_date = formatDateString(paleo_summary.Analysis_start_date)
	paleo_summary.Confidential_release_date = formatDateString(paleo_summary.Confidential_release_date)
	paleo_summary.Effective_date = formatDateString(paleo_summary.Effective_date)
	paleo_summary.Expiry_date = formatDateString(paleo_summary.Expiry_date)
	paleo_summary.Report_date = formatDateString(paleo_summary.Report_date)
	paleo_summary.Row_effective_date = formatDateString(paleo_summary.Row_effective_date)
	paleo_summary.Row_expiry_date = formatDateString(paleo_summary.Row_expiry_date)






	rows, err := stmt.Exec(paleo_summary.Paleo_summary_id, paleo_summary.Access_condition, paleo_summary.Active_ind, paleo_summary.Analysis_end_date, paleo_summary.Analysis_start_date, paleo_summary.Confidential_ind, paleo_summary.Confidential_reason, paleo_summary.Confidential_release_date, paleo_summary.Confidential_term, paleo_summary.Confidential_type, paleo_summary.Diversity_count, paleo_summary.Effective_date, paleo_summary.Expiry_date, paleo_summary.Ppdm_guid, paleo_summary.Preservation_quality, paleo_summary.Reference_num, paleo_summary.Remark, paleo_summary.Report_date, paleo_summary.Report_title, paleo_summary.Source, paleo_summary.Total_sample_interval, paleo_summary.Total_sample_interval_ouom, paleo_summary.Row_changed_by, paleo_summary.Row_changed_date, paleo_summary.Row_created_by, paleo_summary.Row_created_date, paleo_summary.Row_effective_date, paleo_summary.Row_expiry_date, paleo_summary.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePaleoSummary(c *fiber.Ctx) error {
	var paleo_summary dto.Paleo_summary

	setDefaults(&paleo_summary)

	if err := json.Unmarshal(c.Body(), &paleo_summary); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	paleo_summary.Paleo_summary_id = id

    if paleo_summary.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update paleo_summary set access_condition = :1, active_ind = :2, analysis_end_date = :3, analysis_start_date = :4, confidential_ind = :5, confidential_reason = :6, confidential_release_date = :7, confidential_term = :8, confidential_type = :9, diversity_count = :10, effective_date = :11, expiry_date = :12, ppdm_guid = :13, preservation_quality = :14, reference_num = :15, remark = :16, report_date = :17, report_title = :18, source = :19, total_sample_interval = :20, total_sample_interval_ouom = :21, row_changed_by = :22, row_changed_date = :23, row_created_by = :24, row_effective_date = :25, row_expiry_date = :26, row_quality = :27 where paleo_summary_id = :29")
	if err != nil {
		return err
	}

	paleo_summary.Row_changed_date = formatDate(paleo_summary.Row_changed_date)
	paleo_summary.Analysis_end_date = formatDateString(paleo_summary.Analysis_end_date)
	paleo_summary.Analysis_start_date = formatDateString(paleo_summary.Analysis_start_date)
	paleo_summary.Confidential_release_date = formatDateString(paleo_summary.Confidential_release_date)
	paleo_summary.Effective_date = formatDateString(paleo_summary.Effective_date)
	paleo_summary.Expiry_date = formatDateString(paleo_summary.Expiry_date)
	paleo_summary.Report_date = formatDateString(paleo_summary.Report_date)
	paleo_summary.Row_effective_date = formatDateString(paleo_summary.Row_effective_date)
	paleo_summary.Row_expiry_date = formatDateString(paleo_summary.Row_expiry_date)






	rows, err := stmt.Exec(paleo_summary.Access_condition, paleo_summary.Active_ind, paleo_summary.Analysis_end_date, paleo_summary.Analysis_start_date, paleo_summary.Confidential_ind, paleo_summary.Confidential_reason, paleo_summary.Confidential_release_date, paleo_summary.Confidential_term, paleo_summary.Confidential_type, paleo_summary.Diversity_count, paleo_summary.Effective_date, paleo_summary.Expiry_date, paleo_summary.Ppdm_guid, paleo_summary.Preservation_quality, paleo_summary.Reference_num, paleo_summary.Remark, paleo_summary.Report_date, paleo_summary.Report_title, paleo_summary.Source, paleo_summary.Total_sample_interval, paleo_summary.Total_sample_interval_ouom, paleo_summary.Row_changed_by, paleo_summary.Row_changed_date, paleo_summary.Row_created_by, paleo_summary.Row_effective_date, paleo_summary.Row_expiry_date, paleo_summary.Row_quality, paleo_summary.Paleo_summary_id)
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

func PatchPaleoSummary(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update paleo_summary set "
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
		} else if key == "analysis_end_date" || key == "analysis_start_date" || key == "confidential_release_date" || key == "effective_date" || key == "expiry_date" || key == "report_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where paleo_summary_id = :id"

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

func DeletePaleoSummary(c *fiber.Ctx) error {
	id := c.Params("id")
	var paleo_summary dto.Paleo_summary
	paleo_summary.Paleo_summary_id = id

	stmt, err := db.Prepare("delete from paleo_summary where paleo_summary_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(paleo_summary.Paleo_summary_id)
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



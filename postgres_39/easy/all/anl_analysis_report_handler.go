package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlAnalysisReport(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_analysis_report")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_analysis_report

	for rows.Next() {
		var anl_analysis_report dto.Anl_analysis_report
		if err := rows.Scan(&anl_analysis_report.Analysis_id, &anl_analysis_report.Anl_source, &anl_analysis_report.Active_ind, &anl_analysis_report.Analysis_date, &anl_analysis_report.Analysis_purpose, &anl_analysis_report.Analysis_quality, &anl_analysis_report.Base_depth, &anl_analysis_report.Base_depth_ouom, &anl_analysis_report.Base_strat_unit_id, &anl_analysis_report.Effective_date, &anl_analysis_report.End_date, &anl_analysis_report.Expiry_date, &anl_analysis_report.Ppdm_guid, &anl_analysis_report.Received_date, &anl_analysis_report.Remark, &anl_analysis_report.Reported_date, &anl_analysis_report.Reported_tvd, &anl_analysis_report.Reported_tvd_ouom, &anl_analysis_report.Sample_date, &anl_analysis_report.Sample_location, &anl_analysis_report.Start_date, &anl_analysis_report.Strat_name_set_id, &anl_analysis_report.Study_type, &anl_analysis_report.Top_depth, &anl_analysis_report.Top_depth_ouom, &anl_analysis_report.Top_strat_unit_id, &anl_analysis_report.Row_changed_by, &anl_analysis_report.Row_changed_date, &anl_analysis_report.Row_created_by, &anl_analysis_report.Row_created_date, &anl_analysis_report.Row_effective_date, &anl_analysis_report.Row_expiry_date, &anl_analysis_report.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_analysis_report to result
		result = append(result, anl_analysis_report)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlAnalysisReport(c *fiber.Ctx) error {
	var anl_analysis_report dto.Anl_analysis_report

	setDefaults(&anl_analysis_report)

	if err := json.Unmarshal(c.Body(), &anl_analysis_report); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_analysis_report values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33)")
	if err != nil {
		return err
	}
	anl_analysis_report.Row_created_date = formatDate(anl_analysis_report.Row_created_date)
	anl_analysis_report.Row_changed_date = formatDate(anl_analysis_report.Row_changed_date)
	anl_analysis_report.Analysis_date = formatDateString(anl_analysis_report.Analysis_date)
	anl_analysis_report.Effective_date = formatDateString(anl_analysis_report.Effective_date)
	anl_analysis_report.End_date = formatDateString(anl_analysis_report.End_date)
	anl_analysis_report.Expiry_date = formatDateString(anl_analysis_report.Expiry_date)
	anl_analysis_report.Received_date = formatDateString(anl_analysis_report.Received_date)
	anl_analysis_report.Reported_date = formatDateString(anl_analysis_report.Reported_date)
	anl_analysis_report.Sample_date = formatDateString(anl_analysis_report.Sample_date)
	anl_analysis_report.Start_date = formatDateString(anl_analysis_report.Start_date)
	anl_analysis_report.Row_effective_date = formatDateString(anl_analysis_report.Row_effective_date)
	anl_analysis_report.Row_expiry_date = formatDateString(anl_analysis_report.Row_expiry_date)






	rows, err := stmt.Exec(anl_analysis_report.Analysis_id, anl_analysis_report.Anl_source, anl_analysis_report.Active_ind, anl_analysis_report.Analysis_date, anl_analysis_report.Analysis_purpose, anl_analysis_report.Analysis_quality, anl_analysis_report.Base_depth, anl_analysis_report.Base_depth_ouom, anl_analysis_report.Base_strat_unit_id, anl_analysis_report.Effective_date, anl_analysis_report.End_date, anl_analysis_report.Expiry_date, anl_analysis_report.Ppdm_guid, anl_analysis_report.Received_date, anl_analysis_report.Remark, anl_analysis_report.Reported_date, anl_analysis_report.Reported_tvd, anl_analysis_report.Reported_tvd_ouom, anl_analysis_report.Sample_date, anl_analysis_report.Sample_location, anl_analysis_report.Start_date, anl_analysis_report.Strat_name_set_id, anl_analysis_report.Study_type, anl_analysis_report.Top_depth, anl_analysis_report.Top_depth_ouom, anl_analysis_report.Top_strat_unit_id, anl_analysis_report.Row_changed_by, anl_analysis_report.Row_changed_date, anl_analysis_report.Row_created_by, anl_analysis_report.Row_created_date, anl_analysis_report.Row_effective_date, anl_analysis_report.Row_expiry_date, anl_analysis_report.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlAnalysisReport(c *fiber.Ctx) error {
	var anl_analysis_report dto.Anl_analysis_report

	setDefaults(&anl_analysis_report)

	if err := json.Unmarshal(c.Body(), &anl_analysis_report); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_analysis_report.Analysis_id = id

    if anl_analysis_report.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_analysis_report set anl_source = :1, active_ind = :2, analysis_date = :3, analysis_purpose = :4, analysis_quality = :5, base_depth = :6, base_depth_ouom = :7, base_strat_unit_id = :8, effective_date = :9, end_date = :10, expiry_date = :11, ppdm_guid = :12, received_date = :13, remark = :14, reported_date = :15, reported_tvd = :16, reported_tvd_ouom = :17, sample_date = :18, sample_location = :19, start_date = :20, strat_name_set_id = :21, study_type = :22, top_depth = :23, top_depth_ouom = :24, top_strat_unit_id = :25, row_changed_by = :26, row_changed_date = :27, row_created_by = :28, row_effective_date = :29, row_expiry_date = :30, row_quality = :31 where analysis_id = :33")
	if err != nil {
		return err
	}

	anl_analysis_report.Row_changed_date = formatDate(anl_analysis_report.Row_changed_date)
	anl_analysis_report.Analysis_date = formatDateString(anl_analysis_report.Analysis_date)
	anl_analysis_report.Effective_date = formatDateString(anl_analysis_report.Effective_date)
	anl_analysis_report.End_date = formatDateString(anl_analysis_report.End_date)
	anl_analysis_report.Expiry_date = formatDateString(anl_analysis_report.Expiry_date)
	anl_analysis_report.Received_date = formatDateString(anl_analysis_report.Received_date)
	anl_analysis_report.Reported_date = formatDateString(anl_analysis_report.Reported_date)
	anl_analysis_report.Sample_date = formatDateString(anl_analysis_report.Sample_date)
	anl_analysis_report.Start_date = formatDateString(anl_analysis_report.Start_date)
	anl_analysis_report.Row_effective_date = formatDateString(anl_analysis_report.Row_effective_date)
	anl_analysis_report.Row_expiry_date = formatDateString(anl_analysis_report.Row_expiry_date)






	rows, err := stmt.Exec(anl_analysis_report.Anl_source, anl_analysis_report.Active_ind, anl_analysis_report.Analysis_date, anl_analysis_report.Analysis_purpose, anl_analysis_report.Analysis_quality, anl_analysis_report.Base_depth, anl_analysis_report.Base_depth_ouom, anl_analysis_report.Base_strat_unit_id, anl_analysis_report.Effective_date, anl_analysis_report.End_date, anl_analysis_report.Expiry_date, anl_analysis_report.Ppdm_guid, anl_analysis_report.Received_date, anl_analysis_report.Remark, anl_analysis_report.Reported_date, anl_analysis_report.Reported_tvd, anl_analysis_report.Reported_tvd_ouom, anl_analysis_report.Sample_date, anl_analysis_report.Sample_location, anl_analysis_report.Start_date, anl_analysis_report.Strat_name_set_id, anl_analysis_report.Study_type, anl_analysis_report.Top_depth, anl_analysis_report.Top_depth_ouom, anl_analysis_report.Top_strat_unit_id, anl_analysis_report.Row_changed_by, anl_analysis_report.Row_changed_date, anl_analysis_report.Row_created_by, anl_analysis_report.Row_effective_date, anl_analysis_report.Row_expiry_date, anl_analysis_report.Row_quality, anl_analysis_report.Analysis_id)
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

func PatchAnlAnalysisReport(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_analysis_report set "
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
		} else if key == "analysis_date" || key == "effective_date" || key == "end_date" || key == "expiry_date" || key == "received_date" || key == "reported_date" || key == "sample_date" || key == "start_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where analysis_id = :id"

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

func DeleteAnlAnalysisReport(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_analysis_report dto.Anl_analysis_report
	anl_analysis_report.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_analysis_report where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_analysis_report.Analysis_id)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellCoreAnalysis(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_core_analysis")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_core_analysis

	for rows.Next() {
		var well_core_analysis dto.Well_core_analysis
		if err := rows.Scan(&well_core_analysis.Uwi, &well_core_analysis.Source, &well_core_analysis.Core_id, &well_core_analysis.Analysis_obs_no, &well_core_analysis.Active_ind, &well_core_analysis.Analysis_date, &well_core_analysis.Analyzing_company, &well_core_analysis.Analyzing_company_loc, &well_core_analysis.Analyzing_file_num, &well_core_analysis.Core_analyst_name, &well_core_analysis.Effective_date, &well_core_analysis.Expiry_date, &well_core_analysis.Ppdm_guid, &well_core_analysis.Primary_sample_type, &well_core_analysis.Remark, &well_core_analysis.Sample_diameter, &well_core_analysis.Sample_diameter_ouom, &well_core_analysis.Sample_length, &well_core_analysis.Sample_length_ouom, &well_core_analysis.Sample_shape, &well_core_analysis.Second_sample_type, &well_core_analysis.Row_changed_by, &well_core_analysis.Row_changed_date, &well_core_analysis.Row_created_by, &well_core_analysis.Row_created_date, &well_core_analysis.Row_effective_date, &well_core_analysis.Row_expiry_date, &well_core_analysis.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_core_analysis to result
		result = append(result, well_core_analysis)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellCoreAnalysis(c *fiber.Ctx) error {
	var well_core_analysis dto.Well_core_analysis

	setDefaults(&well_core_analysis)

	if err := json.Unmarshal(c.Body(), &well_core_analysis); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_core_analysis values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28)")
	if err != nil {
		return err
	}
	well_core_analysis.Row_created_date = formatDate(well_core_analysis.Row_created_date)
	well_core_analysis.Row_changed_date = formatDate(well_core_analysis.Row_changed_date)
	well_core_analysis.Analysis_date = formatDateString(well_core_analysis.Analysis_date)
	well_core_analysis.Effective_date = formatDateString(well_core_analysis.Effective_date)
	well_core_analysis.Expiry_date = formatDateString(well_core_analysis.Expiry_date)
	well_core_analysis.Row_effective_date = formatDateString(well_core_analysis.Row_effective_date)
	well_core_analysis.Row_expiry_date = formatDateString(well_core_analysis.Row_expiry_date)






	rows, err := stmt.Exec(well_core_analysis.Uwi, well_core_analysis.Source, well_core_analysis.Core_id, well_core_analysis.Analysis_obs_no, well_core_analysis.Active_ind, well_core_analysis.Analysis_date, well_core_analysis.Analyzing_company, well_core_analysis.Analyzing_company_loc, well_core_analysis.Analyzing_file_num, well_core_analysis.Core_analyst_name, well_core_analysis.Effective_date, well_core_analysis.Expiry_date, well_core_analysis.Ppdm_guid, well_core_analysis.Primary_sample_type, well_core_analysis.Remark, well_core_analysis.Sample_diameter, well_core_analysis.Sample_diameter_ouom, well_core_analysis.Sample_length, well_core_analysis.Sample_length_ouom, well_core_analysis.Sample_shape, well_core_analysis.Second_sample_type, well_core_analysis.Row_changed_by, well_core_analysis.Row_changed_date, well_core_analysis.Row_created_by, well_core_analysis.Row_created_date, well_core_analysis.Row_effective_date, well_core_analysis.Row_expiry_date, well_core_analysis.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellCoreAnalysis(c *fiber.Ctx) error {
	var well_core_analysis dto.Well_core_analysis

	setDefaults(&well_core_analysis)

	if err := json.Unmarshal(c.Body(), &well_core_analysis); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_core_analysis.Uwi = id

    if well_core_analysis.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_core_analysis set source = :1, core_id = :2, analysis_obs_no = :3, active_ind = :4, analysis_date = :5, analyzing_company = :6, analyzing_company_loc = :7, analyzing_file_num = :8, core_analyst_name = :9, effective_date = :10, expiry_date = :11, ppdm_guid = :12, primary_sample_type = :13, remark = :14, sample_diameter = :15, sample_diameter_ouom = :16, sample_length = :17, sample_length_ouom = :18, sample_shape = :19, second_sample_type = :20, row_changed_by = :21, row_changed_date = :22, row_created_by = :23, row_effective_date = :24, row_expiry_date = :25, row_quality = :26 where uwi = :28")
	if err != nil {
		return err
	}

	well_core_analysis.Row_changed_date = formatDate(well_core_analysis.Row_changed_date)
	well_core_analysis.Analysis_date = formatDateString(well_core_analysis.Analysis_date)
	well_core_analysis.Effective_date = formatDateString(well_core_analysis.Effective_date)
	well_core_analysis.Expiry_date = formatDateString(well_core_analysis.Expiry_date)
	well_core_analysis.Row_effective_date = formatDateString(well_core_analysis.Row_effective_date)
	well_core_analysis.Row_expiry_date = formatDateString(well_core_analysis.Row_expiry_date)






	rows, err := stmt.Exec(well_core_analysis.Source, well_core_analysis.Core_id, well_core_analysis.Analysis_obs_no, well_core_analysis.Active_ind, well_core_analysis.Analysis_date, well_core_analysis.Analyzing_company, well_core_analysis.Analyzing_company_loc, well_core_analysis.Analyzing_file_num, well_core_analysis.Core_analyst_name, well_core_analysis.Effective_date, well_core_analysis.Expiry_date, well_core_analysis.Ppdm_guid, well_core_analysis.Primary_sample_type, well_core_analysis.Remark, well_core_analysis.Sample_diameter, well_core_analysis.Sample_diameter_ouom, well_core_analysis.Sample_length, well_core_analysis.Sample_length_ouom, well_core_analysis.Sample_shape, well_core_analysis.Second_sample_type, well_core_analysis.Row_changed_by, well_core_analysis.Row_changed_date, well_core_analysis.Row_created_by, well_core_analysis.Row_effective_date, well_core_analysis.Row_expiry_date, well_core_analysis.Row_quality, well_core_analysis.Uwi)
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

func PatchWellCoreAnalysis(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_core_analysis set "
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
		} else if key == "analysis_date" || key == "effective_date" || key == "expiry_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where uwi = :id"

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

func DeleteWellCoreAnalysis(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_core_analysis dto.Well_core_analysis
	well_core_analysis.Uwi = id

	stmt, err := db.Prepare("delete from well_core_analysis where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_core_analysis.Uwi)
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



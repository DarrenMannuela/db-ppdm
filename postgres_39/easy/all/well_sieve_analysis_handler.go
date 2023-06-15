package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellSieveAnalysis(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_sieve_analysis")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_sieve_analysis

	for rows.Next() {
		var well_sieve_analysis dto.Well_sieve_analysis
		if err := rows.Scan(&well_sieve_analysis.Uwi, &well_sieve_analysis.Source, &well_sieve_analysis.Analysis_obs_no, &well_sieve_analysis.Active_ind, &well_sieve_analysis.Analysis_date, &well_sieve_analysis.Base_depth, &well_sieve_analysis.Base_depth_ouom, &well_sieve_analysis.Effective_date, &well_sieve_analysis.Expiry_date, &well_sieve_analysis.Interval_depth, &well_sieve_analysis.Interval_depth_ouom, &well_sieve_analysis.Interval_length, &well_sieve_analysis.Interval_length_ouom, &well_sieve_analysis.Laboratory, &well_sieve_analysis.Lab_file_num, &well_sieve_analysis.Ppdm_guid, &well_sieve_analysis.Remark, &well_sieve_analysis.Top_depth, &well_sieve_analysis.Top_depth_ouom, &well_sieve_analysis.Row_changed_by, &well_sieve_analysis.Row_changed_date, &well_sieve_analysis.Row_created_by, &well_sieve_analysis.Row_created_date, &well_sieve_analysis.Row_effective_date, &well_sieve_analysis.Row_expiry_date, &well_sieve_analysis.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_sieve_analysis to result
		result = append(result, well_sieve_analysis)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellSieveAnalysis(c *fiber.Ctx) error {
	var well_sieve_analysis dto.Well_sieve_analysis

	setDefaults(&well_sieve_analysis)

	if err := json.Unmarshal(c.Body(), &well_sieve_analysis); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_sieve_analysis values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26)")
	if err != nil {
		return err
	}
	well_sieve_analysis.Row_created_date = formatDate(well_sieve_analysis.Row_created_date)
	well_sieve_analysis.Row_changed_date = formatDate(well_sieve_analysis.Row_changed_date)
	well_sieve_analysis.Analysis_date = formatDateString(well_sieve_analysis.Analysis_date)
	well_sieve_analysis.Effective_date = formatDateString(well_sieve_analysis.Effective_date)
	well_sieve_analysis.Expiry_date = formatDateString(well_sieve_analysis.Expiry_date)
	well_sieve_analysis.Row_effective_date = formatDateString(well_sieve_analysis.Row_effective_date)
	well_sieve_analysis.Row_expiry_date = formatDateString(well_sieve_analysis.Row_expiry_date)






	rows, err := stmt.Exec(well_sieve_analysis.Uwi, well_sieve_analysis.Source, well_sieve_analysis.Analysis_obs_no, well_sieve_analysis.Active_ind, well_sieve_analysis.Analysis_date, well_sieve_analysis.Base_depth, well_sieve_analysis.Base_depth_ouom, well_sieve_analysis.Effective_date, well_sieve_analysis.Expiry_date, well_sieve_analysis.Interval_depth, well_sieve_analysis.Interval_depth_ouom, well_sieve_analysis.Interval_length, well_sieve_analysis.Interval_length_ouom, well_sieve_analysis.Laboratory, well_sieve_analysis.Lab_file_num, well_sieve_analysis.Ppdm_guid, well_sieve_analysis.Remark, well_sieve_analysis.Top_depth, well_sieve_analysis.Top_depth_ouom, well_sieve_analysis.Row_changed_by, well_sieve_analysis.Row_changed_date, well_sieve_analysis.Row_created_by, well_sieve_analysis.Row_created_date, well_sieve_analysis.Row_effective_date, well_sieve_analysis.Row_expiry_date, well_sieve_analysis.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellSieveAnalysis(c *fiber.Ctx) error {
	var well_sieve_analysis dto.Well_sieve_analysis

	setDefaults(&well_sieve_analysis)

	if err := json.Unmarshal(c.Body(), &well_sieve_analysis); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_sieve_analysis.Uwi = id

    if well_sieve_analysis.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_sieve_analysis set source = :1, analysis_obs_no = :2, active_ind = :3, analysis_date = :4, base_depth = :5, base_depth_ouom = :6, effective_date = :7, expiry_date = :8, interval_depth = :9, interval_depth_ouom = :10, interval_length = :11, interval_length_ouom = :12, laboratory = :13, lab_file_num = :14, ppdm_guid = :15, remark = :16, top_depth = :17, top_depth_ouom = :18, row_changed_by = :19, row_changed_date = :20, row_created_by = :21, row_effective_date = :22, row_expiry_date = :23, row_quality = :24 where uwi = :26")
	if err != nil {
		return err
	}

	well_sieve_analysis.Row_changed_date = formatDate(well_sieve_analysis.Row_changed_date)
	well_sieve_analysis.Analysis_date = formatDateString(well_sieve_analysis.Analysis_date)
	well_sieve_analysis.Effective_date = formatDateString(well_sieve_analysis.Effective_date)
	well_sieve_analysis.Expiry_date = formatDateString(well_sieve_analysis.Expiry_date)
	well_sieve_analysis.Row_effective_date = formatDateString(well_sieve_analysis.Row_effective_date)
	well_sieve_analysis.Row_expiry_date = formatDateString(well_sieve_analysis.Row_expiry_date)






	rows, err := stmt.Exec(well_sieve_analysis.Source, well_sieve_analysis.Analysis_obs_no, well_sieve_analysis.Active_ind, well_sieve_analysis.Analysis_date, well_sieve_analysis.Base_depth, well_sieve_analysis.Base_depth_ouom, well_sieve_analysis.Effective_date, well_sieve_analysis.Expiry_date, well_sieve_analysis.Interval_depth, well_sieve_analysis.Interval_depth_ouom, well_sieve_analysis.Interval_length, well_sieve_analysis.Interval_length_ouom, well_sieve_analysis.Laboratory, well_sieve_analysis.Lab_file_num, well_sieve_analysis.Ppdm_guid, well_sieve_analysis.Remark, well_sieve_analysis.Top_depth, well_sieve_analysis.Top_depth_ouom, well_sieve_analysis.Row_changed_by, well_sieve_analysis.Row_changed_date, well_sieve_analysis.Row_created_by, well_sieve_analysis.Row_effective_date, well_sieve_analysis.Row_expiry_date, well_sieve_analysis.Row_quality, well_sieve_analysis.Uwi)
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

func PatchWellSieveAnalysis(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_sieve_analysis set "
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

func DeleteWellSieveAnalysis(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_sieve_analysis dto.Well_sieve_analysis
	well_sieve_analysis.Uwi = id

	stmt, err := db.Prepare("delete from well_sieve_analysis where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_sieve_analysis.Uwi)
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



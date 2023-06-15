package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellCoreSampleRmk(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_core_sample_rmk")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_core_sample_rmk

	for rows.Next() {
		var well_core_sample_rmk dto.Well_core_sample_rmk
		if err := rows.Scan(&well_core_sample_rmk.Uwi, &well_core_sample_rmk.Source, &well_core_sample_rmk.Core_id, &well_core_sample_rmk.Analysis_obs_no, &well_core_sample_rmk.Sample_num, &well_core_sample_rmk.Sample_analysis_obs_no, &well_core_sample_rmk.Remark_obs_no, &well_core_sample_rmk.Active_ind, &well_core_sample_rmk.Effective_date, &well_core_sample_rmk.Expiry_date, &well_core_sample_rmk.Ppdm_guid, &well_core_sample_rmk.Remark, &well_core_sample_rmk.Remark_source, &well_core_sample_rmk.Row_changed_by, &well_core_sample_rmk.Row_changed_date, &well_core_sample_rmk.Row_created_by, &well_core_sample_rmk.Row_created_date, &well_core_sample_rmk.Row_effective_date, &well_core_sample_rmk.Row_expiry_date, &well_core_sample_rmk.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_core_sample_rmk to result
		result = append(result, well_core_sample_rmk)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellCoreSampleRmk(c *fiber.Ctx) error {
	var well_core_sample_rmk dto.Well_core_sample_rmk

	setDefaults(&well_core_sample_rmk)

	if err := json.Unmarshal(c.Body(), &well_core_sample_rmk); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_core_sample_rmk values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	well_core_sample_rmk.Row_created_date = formatDate(well_core_sample_rmk.Row_created_date)
	well_core_sample_rmk.Row_changed_date = formatDate(well_core_sample_rmk.Row_changed_date)
	well_core_sample_rmk.Effective_date = formatDateString(well_core_sample_rmk.Effective_date)
	well_core_sample_rmk.Expiry_date = formatDateString(well_core_sample_rmk.Expiry_date)
	well_core_sample_rmk.Row_effective_date = formatDateString(well_core_sample_rmk.Row_effective_date)
	well_core_sample_rmk.Row_expiry_date = formatDateString(well_core_sample_rmk.Row_expiry_date)






	rows, err := stmt.Exec(well_core_sample_rmk.Uwi, well_core_sample_rmk.Source, well_core_sample_rmk.Core_id, well_core_sample_rmk.Analysis_obs_no, well_core_sample_rmk.Sample_num, well_core_sample_rmk.Sample_analysis_obs_no, well_core_sample_rmk.Remark_obs_no, well_core_sample_rmk.Active_ind, well_core_sample_rmk.Effective_date, well_core_sample_rmk.Expiry_date, well_core_sample_rmk.Ppdm_guid, well_core_sample_rmk.Remark, well_core_sample_rmk.Remark_source, well_core_sample_rmk.Row_changed_by, well_core_sample_rmk.Row_changed_date, well_core_sample_rmk.Row_created_by, well_core_sample_rmk.Row_created_date, well_core_sample_rmk.Row_effective_date, well_core_sample_rmk.Row_expiry_date, well_core_sample_rmk.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellCoreSampleRmk(c *fiber.Ctx) error {
	var well_core_sample_rmk dto.Well_core_sample_rmk

	setDefaults(&well_core_sample_rmk)

	if err := json.Unmarshal(c.Body(), &well_core_sample_rmk); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_core_sample_rmk.Uwi = id

    if well_core_sample_rmk.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_core_sample_rmk set source = :1, core_id = :2, analysis_obs_no = :3, sample_num = :4, sample_analysis_obs_no = :5, remark_obs_no = :6, active_ind = :7, effective_date = :8, expiry_date = :9, ppdm_guid = :10, remark = :11, remark_source = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where uwi = :20")
	if err != nil {
		return err
	}

	well_core_sample_rmk.Row_changed_date = formatDate(well_core_sample_rmk.Row_changed_date)
	well_core_sample_rmk.Effective_date = formatDateString(well_core_sample_rmk.Effective_date)
	well_core_sample_rmk.Expiry_date = formatDateString(well_core_sample_rmk.Expiry_date)
	well_core_sample_rmk.Row_effective_date = formatDateString(well_core_sample_rmk.Row_effective_date)
	well_core_sample_rmk.Row_expiry_date = formatDateString(well_core_sample_rmk.Row_expiry_date)






	rows, err := stmt.Exec(well_core_sample_rmk.Source, well_core_sample_rmk.Core_id, well_core_sample_rmk.Analysis_obs_no, well_core_sample_rmk.Sample_num, well_core_sample_rmk.Sample_analysis_obs_no, well_core_sample_rmk.Remark_obs_no, well_core_sample_rmk.Active_ind, well_core_sample_rmk.Effective_date, well_core_sample_rmk.Expiry_date, well_core_sample_rmk.Ppdm_guid, well_core_sample_rmk.Remark, well_core_sample_rmk.Remark_source, well_core_sample_rmk.Row_changed_by, well_core_sample_rmk.Row_changed_date, well_core_sample_rmk.Row_created_by, well_core_sample_rmk.Row_effective_date, well_core_sample_rmk.Row_expiry_date, well_core_sample_rmk.Row_quality, well_core_sample_rmk.Uwi)
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

func PatchWellCoreSampleRmk(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_core_sample_rmk set "
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

func DeleteWellCoreSampleRmk(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_core_sample_rmk dto.Well_core_sample_rmk
	well_core_sample_rmk.Uwi = id

	stmt, err := db.Prepare("delete from well_core_sample_rmk where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_core_sample_rmk.Uwi)
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



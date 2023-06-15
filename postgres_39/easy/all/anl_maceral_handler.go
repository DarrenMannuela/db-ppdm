package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlMaceral(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_maceral")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_maceral

	for rows.Next() {
		var anl_maceral dto.Anl_maceral
		if err := rows.Scan(&anl_maceral.Analysis_id, &anl_maceral.Anl_source, &anl_maceral.Maceral_anl_obs_no, &anl_maceral.Active_ind, &anl_maceral.Effective_date, &anl_maceral.Expiry_date, &anl_maceral.Lithology_desc, &anl_maceral.Maceral_amount_type, &anl_maceral.Ppdm_guid, &anl_maceral.Preferred_ind, &anl_maceral.Problem_ind, &anl_maceral.Remark, &anl_maceral.Sample_tot_maceral_val, &anl_maceral.Sample_tot_maceral_val_ouom, &anl_maceral.Source, &anl_maceral.Step_seq_no, &anl_maceral.Substance_id, &anl_maceral.Total_maceral_val, &anl_maceral.Total_maceral_val_ouom, &anl_maceral.Row_changed_by, &anl_maceral.Row_changed_date, &anl_maceral.Row_created_by, &anl_maceral.Row_created_date, &anl_maceral.Row_effective_date, &anl_maceral.Row_expiry_date, &anl_maceral.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_maceral to result
		result = append(result, anl_maceral)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlMaceral(c *fiber.Ctx) error {
	var anl_maceral dto.Anl_maceral

	setDefaults(&anl_maceral)

	if err := json.Unmarshal(c.Body(), &anl_maceral); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_maceral values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26)")
	if err != nil {
		return err
	}
	anl_maceral.Row_created_date = formatDate(anl_maceral.Row_created_date)
	anl_maceral.Row_changed_date = formatDate(anl_maceral.Row_changed_date)
	anl_maceral.Effective_date = formatDateString(anl_maceral.Effective_date)
	anl_maceral.Expiry_date = formatDateString(anl_maceral.Expiry_date)
	anl_maceral.Row_effective_date = formatDateString(anl_maceral.Row_effective_date)
	anl_maceral.Row_expiry_date = formatDateString(anl_maceral.Row_expiry_date)






	rows, err := stmt.Exec(anl_maceral.Analysis_id, anl_maceral.Anl_source, anl_maceral.Maceral_anl_obs_no, anl_maceral.Active_ind, anl_maceral.Effective_date, anl_maceral.Expiry_date, anl_maceral.Lithology_desc, anl_maceral.Maceral_amount_type, anl_maceral.Ppdm_guid, anl_maceral.Preferred_ind, anl_maceral.Problem_ind, anl_maceral.Remark, anl_maceral.Sample_tot_maceral_val, anl_maceral.Sample_tot_maceral_val_ouom, anl_maceral.Source, anl_maceral.Step_seq_no, anl_maceral.Substance_id, anl_maceral.Total_maceral_val, anl_maceral.Total_maceral_val_ouom, anl_maceral.Row_changed_by, anl_maceral.Row_changed_date, anl_maceral.Row_created_by, anl_maceral.Row_created_date, anl_maceral.Row_effective_date, anl_maceral.Row_expiry_date, anl_maceral.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlMaceral(c *fiber.Ctx) error {
	var anl_maceral dto.Anl_maceral

	setDefaults(&anl_maceral)

	if err := json.Unmarshal(c.Body(), &anl_maceral); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_maceral.Analysis_id = id

    if anl_maceral.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_maceral set anl_source = :1, maceral_anl_obs_no = :2, active_ind = :3, effective_date = :4, expiry_date = :5, lithology_desc = :6, maceral_amount_type = :7, ppdm_guid = :8, preferred_ind = :9, problem_ind = :10, remark = :11, sample_tot_maceral_val = :12, sample_tot_maceral_val_ouom = :13, source = :14, step_seq_no = :15, substance_id = :16, total_maceral_val = :17, total_maceral_val_ouom = :18, row_changed_by = :19, row_changed_date = :20, row_created_by = :21, row_effective_date = :22, row_expiry_date = :23, row_quality = :24 where analysis_id = :26")
	if err != nil {
		return err
	}

	anl_maceral.Row_changed_date = formatDate(anl_maceral.Row_changed_date)
	anl_maceral.Effective_date = formatDateString(anl_maceral.Effective_date)
	anl_maceral.Expiry_date = formatDateString(anl_maceral.Expiry_date)
	anl_maceral.Row_effective_date = formatDateString(anl_maceral.Row_effective_date)
	anl_maceral.Row_expiry_date = formatDateString(anl_maceral.Row_expiry_date)






	rows, err := stmt.Exec(anl_maceral.Anl_source, anl_maceral.Maceral_anl_obs_no, anl_maceral.Active_ind, anl_maceral.Effective_date, anl_maceral.Expiry_date, anl_maceral.Lithology_desc, anl_maceral.Maceral_amount_type, anl_maceral.Ppdm_guid, anl_maceral.Preferred_ind, anl_maceral.Problem_ind, anl_maceral.Remark, anl_maceral.Sample_tot_maceral_val, anl_maceral.Sample_tot_maceral_val_ouom, anl_maceral.Source, anl_maceral.Step_seq_no, anl_maceral.Substance_id, anl_maceral.Total_maceral_val, anl_maceral.Total_maceral_val_ouom, anl_maceral.Row_changed_by, anl_maceral.Row_changed_date, anl_maceral.Row_created_by, anl_maceral.Row_effective_date, anl_maceral.Row_expiry_date, anl_maceral.Row_quality, anl_maceral.Analysis_id)
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

func PatchAnlMaceral(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_maceral set "
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

func DeleteAnlMaceral(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_maceral dto.Anl_maceral
	anl_maceral.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_maceral where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_maceral.Analysis_id)
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



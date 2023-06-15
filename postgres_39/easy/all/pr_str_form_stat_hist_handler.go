package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPrStrFormStatHist(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pr_str_form_stat_hist")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pr_str_form_stat_hist

	for rows.Next() {
		var pr_str_form_stat_hist dto.Pr_str_form_stat_hist
		if err := rows.Scan(&pr_str_form_stat_hist.Uwi, &pr_str_form_stat_hist.Prod_string_source, &pr_str_form_stat_hist.String_id, &pr_str_form_stat_hist.Pr_str_form_obs_no, &pr_str_form_stat_hist.Status_id, &pr_str_form_stat_hist.Active_ind, &pr_str_form_stat_hist.Effective_date, &pr_str_form_stat_hist.End_time, &pr_str_form_stat_hist.Expiry_date, &pr_str_form_stat_hist.Percent_capability, &pr_str_form_stat_hist.Ppdm_guid, &pr_str_form_stat_hist.Remark, &pr_str_form_stat_hist.Source, &pr_str_form_stat_hist.Start_time, &pr_str_form_stat_hist.Status, &pr_str_form_stat_hist.Status_date, &pr_str_form_stat_hist.Status_type, &pr_str_form_stat_hist.Timezone, &pr_str_form_stat_hist.Row_changed_by, &pr_str_form_stat_hist.Row_changed_date, &pr_str_form_stat_hist.Row_created_by, &pr_str_form_stat_hist.Row_created_date, &pr_str_form_stat_hist.Row_effective_date, &pr_str_form_stat_hist.Row_expiry_date, &pr_str_form_stat_hist.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pr_str_form_stat_hist to result
		result = append(result, pr_str_form_stat_hist)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPrStrFormStatHist(c *fiber.Ctx) error {
	var pr_str_form_stat_hist dto.Pr_str_form_stat_hist

	setDefaults(&pr_str_form_stat_hist)

	if err := json.Unmarshal(c.Body(), &pr_str_form_stat_hist); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pr_str_form_stat_hist values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25)")
	if err != nil {
		return err
	}
	pr_str_form_stat_hist.Row_created_date = formatDate(pr_str_form_stat_hist.Row_created_date)
	pr_str_form_stat_hist.Row_changed_date = formatDate(pr_str_form_stat_hist.Row_changed_date)
	pr_str_form_stat_hist.Effective_date = formatDateString(pr_str_form_stat_hist.Effective_date)
	pr_str_form_stat_hist.End_time = formatDateString(pr_str_form_stat_hist.End_time)
	pr_str_form_stat_hist.Expiry_date = formatDateString(pr_str_form_stat_hist.Expiry_date)
	pr_str_form_stat_hist.Start_time = formatDateString(pr_str_form_stat_hist.Start_time)
	pr_str_form_stat_hist.Status_date = formatDateString(pr_str_form_stat_hist.Status_date)
	pr_str_form_stat_hist.Row_effective_date = formatDateString(pr_str_form_stat_hist.Row_effective_date)
	pr_str_form_stat_hist.Row_expiry_date = formatDateString(pr_str_form_stat_hist.Row_expiry_date)






	rows, err := stmt.Exec(pr_str_form_stat_hist.Uwi, pr_str_form_stat_hist.Prod_string_source, pr_str_form_stat_hist.String_id, pr_str_form_stat_hist.Pr_str_form_obs_no, pr_str_form_stat_hist.Status_id, pr_str_form_stat_hist.Active_ind, pr_str_form_stat_hist.Effective_date, pr_str_form_stat_hist.End_time, pr_str_form_stat_hist.Expiry_date, pr_str_form_stat_hist.Percent_capability, pr_str_form_stat_hist.Ppdm_guid, pr_str_form_stat_hist.Remark, pr_str_form_stat_hist.Source, pr_str_form_stat_hist.Start_time, pr_str_form_stat_hist.Status, pr_str_form_stat_hist.Status_date, pr_str_form_stat_hist.Status_type, pr_str_form_stat_hist.Timezone, pr_str_form_stat_hist.Row_changed_by, pr_str_form_stat_hist.Row_changed_date, pr_str_form_stat_hist.Row_created_by, pr_str_form_stat_hist.Row_created_date, pr_str_form_stat_hist.Row_effective_date, pr_str_form_stat_hist.Row_expiry_date, pr_str_form_stat_hist.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePrStrFormStatHist(c *fiber.Ctx) error {
	var pr_str_form_stat_hist dto.Pr_str_form_stat_hist

	setDefaults(&pr_str_form_stat_hist)

	if err := json.Unmarshal(c.Body(), &pr_str_form_stat_hist); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pr_str_form_stat_hist.Uwi = id

    if pr_str_form_stat_hist.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pr_str_form_stat_hist set prod_string_source = :1, string_id = :2, pr_str_form_obs_no = :3, status_id = :4, active_ind = :5, effective_date = :6, end_time = :7, expiry_date = :8, percent_capability = :9, ppdm_guid = :10, remark = :11, source = :12, start_time = :13, status = :14, status_date = :15, status_type = :16, timezone = :17, row_changed_by = :18, row_changed_date = :19, row_created_by = :20, row_effective_date = :21, row_expiry_date = :22, row_quality = :23 where uwi = :25")
	if err != nil {
		return err
	}

	pr_str_form_stat_hist.Row_changed_date = formatDate(pr_str_form_stat_hist.Row_changed_date)
	pr_str_form_stat_hist.Effective_date = formatDateString(pr_str_form_stat_hist.Effective_date)
	pr_str_form_stat_hist.End_time = formatDateString(pr_str_form_stat_hist.End_time)
	pr_str_form_stat_hist.Expiry_date = formatDateString(pr_str_form_stat_hist.Expiry_date)
	pr_str_form_stat_hist.Start_time = formatDateString(pr_str_form_stat_hist.Start_time)
	pr_str_form_stat_hist.Status_date = formatDateString(pr_str_form_stat_hist.Status_date)
	pr_str_form_stat_hist.Row_effective_date = formatDateString(pr_str_form_stat_hist.Row_effective_date)
	pr_str_form_stat_hist.Row_expiry_date = formatDateString(pr_str_form_stat_hist.Row_expiry_date)






	rows, err := stmt.Exec(pr_str_form_stat_hist.Prod_string_source, pr_str_form_stat_hist.String_id, pr_str_form_stat_hist.Pr_str_form_obs_no, pr_str_form_stat_hist.Status_id, pr_str_form_stat_hist.Active_ind, pr_str_form_stat_hist.Effective_date, pr_str_form_stat_hist.End_time, pr_str_form_stat_hist.Expiry_date, pr_str_form_stat_hist.Percent_capability, pr_str_form_stat_hist.Ppdm_guid, pr_str_form_stat_hist.Remark, pr_str_form_stat_hist.Source, pr_str_form_stat_hist.Start_time, pr_str_form_stat_hist.Status, pr_str_form_stat_hist.Status_date, pr_str_form_stat_hist.Status_type, pr_str_form_stat_hist.Timezone, pr_str_form_stat_hist.Row_changed_by, pr_str_form_stat_hist.Row_changed_date, pr_str_form_stat_hist.Row_created_by, pr_str_form_stat_hist.Row_effective_date, pr_str_form_stat_hist.Row_expiry_date, pr_str_form_stat_hist.Row_quality, pr_str_form_stat_hist.Uwi)
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

func PatchPrStrFormStatHist(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pr_str_form_stat_hist set "
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
		} else if key == "effective_date" || key == "end_time" || key == "expiry_date" || key == "start_time" || key == "status_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeletePrStrFormStatHist(c *fiber.Ctx) error {
	id := c.Params("id")
	var pr_str_form_stat_hist dto.Pr_str_form_stat_hist
	pr_str_form_stat_hist.Uwi = id

	stmt, err := db.Prepare("delete from pr_str_form_stat_hist where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pr_str_form_stat_hist.Uwi)
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



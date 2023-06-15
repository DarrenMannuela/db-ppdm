package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetResentEcoRun(c *fiber.Ctx) error {
	rows, err := db.Query("select * from resent_eco_run")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Resent_eco_run

	for rows.Next() {
		var resent_eco_run dto.Resent_eco_run
		if err := rows.Scan(&resent_eco_run.Resent_id, &resent_eco_run.Reserve_class_id, &resent_eco_run.Economics_run_id, &resent_eco_run.Active_ind, &resent_eco_run.Currency_conversion, &resent_eco_run.Currency_ouom, &resent_eco_run.Currency_uom, &resent_eco_run.Economic_forecast, &resent_eco_run.Economic_scenario, &resent_eco_run.Effective_date, &resent_eco_run.Expiry_date, &resent_eco_run.Gross_ind, &resent_eco_run.Interest_set_id, &resent_eco_run.Interest_set_seq_no, &resent_eco_run.Net_ind, &resent_eco_run.Net_present_value, &resent_eco_run.Partner_ba_id, &resent_eco_run.Partner_obs_no, &resent_eco_run.Ppdm_guid, &resent_eco_run.Remark, &resent_eco_run.Reserve_life_index, &resent_eco_run.Source, &resent_eco_run.Tech_forecast, &resent_eco_run.Row_changed_by, &resent_eco_run.Row_changed_date, &resent_eco_run.Row_created_by, &resent_eco_run.Row_created_date, &resent_eco_run.Row_effective_date, &resent_eco_run.Row_expiry_date, &resent_eco_run.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append resent_eco_run to result
		result = append(result, resent_eco_run)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetResentEcoRun(c *fiber.Ctx) error {
	var resent_eco_run dto.Resent_eco_run

	setDefaults(&resent_eco_run)

	if err := json.Unmarshal(c.Body(), &resent_eco_run); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into resent_eco_run values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30)")
	if err != nil {
		return err
	}
	resent_eco_run.Row_created_date = formatDate(resent_eco_run.Row_created_date)
	resent_eco_run.Row_changed_date = formatDate(resent_eco_run.Row_changed_date)
	resent_eco_run.Effective_date = formatDateString(resent_eco_run.Effective_date)
	resent_eco_run.Expiry_date = formatDateString(resent_eco_run.Expiry_date)
	resent_eco_run.Row_effective_date = formatDateString(resent_eco_run.Row_effective_date)
	resent_eco_run.Row_expiry_date = formatDateString(resent_eco_run.Row_expiry_date)






	rows, err := stmt.Exec(resent_eco_run.Resent_id, resent_eco_run.Reserve_class_id, resent_eco_run.Economics_run_id, resent_eco_run.Active_ind, resent_eco_run.Currency_conversion, resent_eco_run.Currency_ouom, resent_eco_run.Currency_uom, resent_eco_run.Economic_forecast, resent_eco_run.Economic_scenario, resent_eco_run.Effective_date, resent_eco_run.Expiry_date, resent_eco_run.Gross_ind, resent_eco_run.Interest_set_id, resent_eco_run.Interest_set_seq_no, resent_eco_run.Net_ind, resent_eco_run.Net_present_value, resent_eco_run.Partner_ba_id, resent_eco_run.Partner_obs_no, resent_eco_run.Ppdm_guid, resent_eco_run.Remark, resent_eco_run.Reserve_life_index, resent_eco_run.Source, resent_eco_run.Tech_forecast, resent_eco_run.Row_changed_by, resent_eco_run.Row_changed_date, resent_eco_run.Row_created_by, resent_eco_run.Row_created_date, resent_eco_run.Row_effective_date, resent_eco_run.Row_expiry_date, resent_eco_run.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateResentEcoRun(c *fiber.Ctx) error {
	var resent_eco_run dto.Resent_eco_run

	setDefaults(&resent_eco_run)

	if err := json.Unmarshal(c.Body(), &resent_eco_run); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	resent_eco_run.Resent_id = id

    if resent_eco_run.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update resent_eco_run set reserve_class_id = :1, economics_run_id = :2, active_ind = :3, currency_conversion = :4, currency_ouom = :5, currency_uom = :6, economic_forecast = :7, economic_scenario = :8, effective_date = :9, expiry_date = :10, gross_ind = :11, interest_set_id = :12, interest_set_seq_no = :13, net_ind = :14, net_present_value = :15, partner_ba_id = :16, partner_obs_no = :17, ppdm_guid = :18, remark = :19, reserve_life_index = :20, source = :21, tech_forecast = :22, row_changed_by = :23, row_changed_date = :24, row_created_by = :25, row_effective_date = :26, row_expiry_date = :27, row_quality = :28 where resent_id = :30")
	if err != nil {
		return err
	}

	resent_eco_run.Row_changed_date = formatDate(resent_eco_run.Row_changed_date)
	resent_eco_run.Effective_date = formatDateString(resent_eco_run.Effective_date)
	resent_eco_run.Expiry_date = formatDateString(resent_eco_run.Expiry_date)
	resent_eco_run.Row_effective_date = formatDateString(resent_eco_run.Row_effective_date)
	resent_eco_run.Row_expiry_date = formatDateString(resent_eco_run.Row_expiry_date)






	rows, err := stmt.Exec(resent_eco_run.Reserve_class_id, resent_eco_run.Economics_run_id, resent_eco_run.Active_ind, resent_eco_run.Currency_conversion, resent_eco_run.Currency_ouom, resent_eco_run.Currency_uom, resent_eco_run.Economic_forecast, resent_eco_run.Economic_scenario, resent_eco_run.Effective_date, resent_eco_run.Expiry_date, resent_eco_run.Gross_ind, resent_eco_run.Interest_set_id, resent_eco_run.Interest_set_seq_no, resent_eco_run.Net_ind, resent_eco_run.Net_present_value, resent_eco_run.Partner_ba_id, resent_eco_run.Partner_obs_no, resent_eco_run.Ppdm_guid, resent_eco_run.Remark, resent_eco_run.Reserve_life_index, resent_eco_run.Source, resent_eco_run.Tech_forecast, resent_eco_run.Row_changed_by, resent_eco_run.Row_changed_date, resent_eco_run.Row_created_by, resent_eco_run.Row_effective_date, resent_eco_run.Row_expiry_date, resent_eco_run.Row_quality, resent_eco_run.Resent_id)
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

func PatchResentEcoRun(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update resent_eco_run set "
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

func DeleteResentEcoRun(c *fiber.Ctx) error {
	id := c.Params("id")
	var resent_eco_run dto.Resent_eco_run
	resent_eco_run.Resent_id = id

	stmt, err := db.Prepare("delete from resent_eco_run where resent_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(resent_eco_run.Resent_id)
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



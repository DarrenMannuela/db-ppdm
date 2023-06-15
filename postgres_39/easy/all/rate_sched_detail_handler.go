package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRateSchedDetail(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rate_sched_detail")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rate_sched_detail

	for rows.Next() {
		var rate_sched_detail dto.Rate_sched_detail
		if err := rows.Scan(&rate_sched_detail.Rate_schedule_id, &rate_sched_detail.Rate_type, &rate_sched_detail.Active_ind, &rate_sched_detail.Currency_conversion, &rate_sched_detail.Currency_ouom, &rate_sched_detail.Effective_date, &rate_sched_detail.Expiry_date, &rate_sched_detail.Nat_currency_conversion, &rate_sched_detail.Nat_currency_uom, &rate_sched_detail.Period_type, &rate_sched_detail.Ppdm_guid, &rate_sched_detail.Rate_condition, &rate_sched_detail.Rate_cost, &rate_sched_detail.Rate_cost_uom, &rate_sched_detail.Rate_formula, &rate_sched_detail.Rate_percent, &rate_sched_detail.Remark, &rate_sched_detail.Source, &rate_sched_detail.Taxable_ind, &rate_sched_detail.Row_changed_by, &rate_sched_detail.Row_changed_date, &rate_sched_detail.Row_created_by, &rate_sched_detail.Row_created_date, &rate_sched_detail.Row_effective_date, &rate_sched_detail.Row_expiry_date, &rate_sched_detail.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rate_sched_detail to result
		result = append(result, rate_sched_detail)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRateSchedDetail(c *fiber.Ctx) error {
	var rate_sched_detail dto.Rate_sched_detail

	setDefaults(&rate_sched_detail)

	if err := json.Unmarshal(c.Body(), &rate_sched_detail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rate_sched_detail values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26)")
	if err != nil {
		return err
	}
	rate_sched_detail.Row_created_date = formatDate(rate_sched_detail.Row_created_date)
	rate_sched_detail.Row_changed_date = formatDate(rate_sched_detail.Row_changed_date)
	rate_sched_detail.Effective_date = formatDateString(rate_sched_detail.Effective_date)
	rate_sched_detail.Expiry_date = formatDateString(rate_sched_detail.Expiry_date)
	rate_sched_detail.Row_effective_date = formatDateString(rate_sched_detail.Row_effective_date)
	rate_sched_detail.Row_expiry_date = formatDateString(rate_sched_detail.Row_expiry_date)






	rows, err := stmt.Exec(rate_sched_detail.Rate_schedule_id, rate_sched_detail.Rate_type, rate_sched_detail.Active_ind, rate_sched_detail.Currency_conversion, rate_sched_detail.Currency_ouom, rate_sched_detail.Effective_date, rate_sched_detail.Expiry_date, rate_sched_detail.Nat_currency_conversion, rate_sched_detail.Nat_currency_uom, rate_sched_detail.Period_type, rate_sched_detail.Ppdm_guid, rate_sched_detail.Rate_condition, rate_sched_detail.Rate_cost, rate_sched_detail.Rate_cost_uom, rate_sched_detail.Rate_formula, rate_sched_detail.Rate_percent, rate_sched_detail.Remark, rate_sched_detail.Source, rate_sched_detail.Taxable_ind, rate_sched_detail.Row_changed_by, rate_sched_detail.Row_changed_date, rate_sched_detail.Row_created_by, rate_sched_detail.Row_created_date, rate_sched_detail.Row_effective_date, rate_sched_detail.Row_expiry_date, rate_sched_detail.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRateSchedDetail(c *fiber.Ctx) error {
	var rate_sched_detail dto.Rate_sched_detail

	setDefaults(&rate_sched_detail)

	if err := json.Unmarshal(c.Body(), &rate_sched_detail); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rate_sched_detail.Rate_schedule_id = id

    if rate_sched_detail.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rate_sched_detail set rate_type = :1, active_ind = :2, currency_conversion = :3, currency_ouom = :4, effective_date = :5, expiry_date = :6, nat_currency_conversion = :7, nat_currency_uom = :8, period_type = :9, ppdm_guid = :10, rate_condition = :11, rate_cost = :12, rate_cost_uom = :13, rate_formula = :14, rate_percent = :15, remark = :16, source = :17, taxable_ind = :18, row_changed_by = :19, row_changed_date = :20, row_created_by = :21, row_effective_date = :22, row_expiry_date = :23, row_quality = :24 where rate_schedule_id = :26")
	if err != nil {
		return err
	}

	rate_sched_detail.Row_changed_date = formatDate(rate_sched_detail.Row_changed_date)
	rate_sched_detail.Effective_date = formatDateString(rate_sched_detail.Effective_date)
	rate_sched_detail.Expiry_date = formatDateString(rate_sched_detail.Expiry_date)
	rate_sched_detail.Row_effective_date = formatDateString(rate_sched_detail.Row_effective_date)
	rate_sched_detail.Row_expiry_date = formatDateString(rate_sched_detail.Row_expiry_date)






	rows, err := stmt.Exec(rate_sched_detail.Rate_type, rate_sched_detail.Active_ind, rate_sched_detail.Currency_conversion, rate_sched_detail.Currency_ouom, rate_sched_detail.Effective_date, rate_sched_detail.Expiry_date, rate_sched_detail.Nat_currency_conversion, rate_sched_detail.Nat_currency_uom, rate_sched_detail.Period_type, rate_sched_detail.Ppdm_guid, rate_sched_detail.Rate_condition, rate_sched_detail.Rate_cost, rate_sched_detail.Rate_cost_uom, rate_sched_detail.Rate_formula, rate_sched_detail.Rate_percent, rate_sched_detail.Remark, rate_sched_detail.Source, rate_sched_detail.Taxable_ind, rate_sched_detail.Row_changed_by, rate_sched_detail.Row_changed_date, rate_sched_detail.Row_created_by, rate_sched_detail.Row_effective_date, rate_sched_detail.Row_expiry_date, rate_sched_detail.Row_quality, rate_sched_detail.Rate_schedule_id)
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

func PatchRateSchedDetail(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rate_sched_detail set "
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
	query += " where rate_schedule_id = :id"

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

func DeleteRateSchedDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	var rate_sched_detail dto.Rate_sched_detail
	rate_sched_detail.Rate_schedule_id = id

	stmt, err := db.Prepare("delete from rate_sched_detail where rate_schedule_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rate_sched_detail.Rate_schedule_id)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetRateSchedule(c *fiber.Ctx) error {
	rows, err := db.Query("select * from rate_schedule")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Rate_schedule

	for rows.Next() {
		var rate_schedule dto.Rate_schedule
		if err := rows.Scan(&rate_schedule.Rate_schedule_id, &rate_schedule.Active_ind, &rate_schedule.Area_id, &rate_schedule.Area_type, &rate_schedule.Change_notice, &rate_schedule.Contract_id, &rate_schedule.Effective_date, &rate_schedule.Expiry_date, &rate_schedule.Owner_ba_add_obs_no, &rate_schedule.Owner_ba_add_source, &rate_schedule.Owner_ba_id, &rate_schedule.Pay_method, &rate_schedule.Ppdm_guid, &rate_schedule.Preferred_currency_uom, &rate_schedule.Provision_id, &rate_schedule.Rate_schedule_type, &rate_schedule.Remark, &rate_schedule.Review_frequency, &rate_schedule.Source, &rate_schedule.Source_document_id, &rate_schedule.Spatial_description_id, &rate_schedule.Spatial_obs_no, &rate_schedule.Row_changed_by, &rate_schedule.Row_changed_date, &rate_schedule.Row_created_by, &rate_schedule.Row_created_date, &rate_schedule.Row_effective_date, &rate_schedule.Row_expiry_date, &rate_schedule.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append rate_schedule to result
		result = append(result, rate_schedule)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetRateSchedule(c *fiber.Ctx) error {
	var rate_schedule dto.Rate_schedule

	setDefaults(&rate_schedule)

	if err := json.Unmarshal(c.Body(), &rate_schedule); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into rate_schedule values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29)")
	if err != nil {
		return err
	}
	rate_schedule.Row_created_date = formatDate(rate_schedule.Row_created_date)
	rate_schedule.Row_changed_date = formatDate(rate_schedule.Row_changed_date)
	rate_schedule.Effective_date = formatDateString(rate_schedule.Effective_date)
	rate_schedule.Expiry_date = formatDateString(rate_schedule.Expiry_date)
	rate_schedule.Row_effective_date = formatDateString(rate_schedule.Row_effective_date)
	rate_schedule.Row_expiry_date = formatDateString(rate_schedule.Row_expiry_date)






	rows, err := stmt.Exec(rate_schedule.Rate_schedule_id, rate_schedule.Active_ind, rate_schedule.Area_id, rate_schedule.Area_type, rate_schedule.Change_notice, rate_schedule.Contract_id, rate_schedule.Effective_date, rate_schedule.Expiry_date, rate_schedule.Owner_ba_add_obs_no, rate_schedule.Owner_ba_add_source, rate_schedule.Owner_ba_id, rate_schedule.Pay_method, rate_schedule.Ppdm_guid, rate_schedule.Preferred_currency_uom, rate_schedule.Provision_id, rate_schedule.Rate_schedule_type, rate_schedule.Remark, rate_schedule.Review_frequency, rate_schedule.Source, rate_schedule.Source_document_id, rate_schedule.Spatial_description_id, rate_schedule.Spatial_obs_no, rate_schedule.Row_changed_by, rate_schedule.Row_changed_date, rate_schedule.Row_created_by, rate_schedule.Row_created_date, rate_schedule.Row_effective_date, rate_schedule.Row_expiry_date, rate_schedule.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateRateSchedule(c *fiber.Ctx) error {
	var rate_schedule dto.Rate_schedule

	setDefaults(&rate_schedule)

	if err := json.Unmarshal(c.Body(), &rate_schedule); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	rate_schedule.Rate_schedule_id = id

    if rate_schedule.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update rate_schedule set active_ind = :1, area_id = :2, area_type = :3, change_notice = :4, contract_id = :5, effective_date = :6, expiry_date = :7, owner_ba_add_obs_no = :8, owner_ba_add_source = :9, owner_ba_id = :10, pay_method = :11, ppdm_guid = :12, preferred_currency_uom = :13, provision_id = :14, rate_schedule_type = :15, remark = :16, review_frequency = :17, source = :18, source_document_id = :19, spatial_description_id = :20, spatial_obs_no = :21, row_changed_by = :22, row_changed_date = :23, row_created_by = :24, row_effective_date = :25, row_expiry_date = :26, row_quality = :27 where rate_schedule_id = :29")
	if err != nil {
		return err
	}

	rate_schedule.Row_changed_date = formatDate(rate_schedule.Row_changed_date)
	rate_schedule.Effective_date = formatDateString(rate_schedule.Effective_date)
	rate_schedule.Expiry_date = formatDateString(rate_schedule.Expiry_date)
	rate_schedule.Row_effective_date = formatDateString(rate_schedule.Row_effective_date)
	rate_schedule.Row_expiry_date = formatDateString(rate_schedule.Row_expiry_date)






	rows, err := stmt.Exec(rate_schedule.Active_ind, rate_schedule.Area_id, rate_schedule.Area_type, rate_schedule.Change_notice, rate_schedule.Contract_id, rate_schedule.Effective_date, rate_schedule.Expiry_date, rate_schedule.Owner_ba_add_obs_no, rate_schedule.Owner_ba_add_source, rate_schedule.Owner_ba_id, rate_schedule.Pay_method, rate_schedule.Ppdm_guid, rate_schedule.Preferred_currency_uom, rate_schedule.Provision_id, rate_schedule.Rate_schedule_type, rate_schedule.Remark, rate_schedule.Review_frequency, rate_schedule.Source, rate_schedule.Source_document_id, rate_schedule.Spatial_description_id, rate_schedule.Spatial_obs_no, rate_schedule.Row_changed_by, rate_schedule.Row_changed_date, rate_schedule.Row_created_by, rate_schedule.Row_effective_date, rate_schedule.Row_expiry_date, rate_schedule.Row_quality, rate_schedule.Rate_schedule_id)
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

func PatchRateSchedule(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update rate_schedule set "
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

func DeleteRateSchedule(c *fiber.Ctx) error {
	id := c.Params("id")
	var rate_schedule dto.Rate_schedule
	rate_schedule.Rate_schedule_id = id

	stmt, err := db.Prepare("delete from rate_schedule where rate_schedule_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(rate_schedule.Rate_schedule_id)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPpdmMetricValue(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ppdm_metric_value")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ppdm_metric_value

	for rows.Next() {
		var ppdm_metric_value dto.Ppdm_metric_value
		if err := rows.Scan(&ppdm_metric_value.Metric_id, &ppdm_metric_value.Metric_obs_no, &ppdm_metric_value.Active_ind, &ppdm_metric_value.Average_value, &ppdm_metric_value.Average_value_ouom, &ppdm_metric_value.Average_value_uom, &ppdm_metric_value.Cost, &ppdm_metric_value.Currency_conversion, &ppdm_metric_value.Currency_ouom, &ppdm_metric_value.Currency_uom, &ppdm_metric_value.Date_format_desc, &ppdm_metric_value.Effective_date, &ppdm_metric_value.Expiry_date, &ppdm_metric_value.Max_date, &ppdm_metric_value.Max_value, &ppdm_metric_value.Max_value_ouom, &ppdm_metric_value.Max_value_uom, &ppdm_metric_value.Metric_code, &ppdm_metric_value.Metric_desc, &ppdm_metric_value.Metric_ind, &ppdm_metric_value.Metric_type, &ppdm_metric_value.Min_date, &ppdm_metric_value.Min_value, &ppdm_metric_value.Min_value_ouom, &ppdm_metric_value.Min_value_uom, &ppdm_metric_value.Ppdm_guid, &ppdm_metric_value.Reference_value, &ppdm_metric_value.Reference_value_ouom, &ppdm_metric_value.Reference_value_type, &ppdm_metric_value.Reference_value_uom, &ppdm_metric_value.Remark, &ppdm_metric_value.Source, &ppdm_metric_value.Row_changed_by, &ppdm_metric_value.Row_changed_date, &ppdm_metric_value.Row_created_by, &ppdm_metric_value.Row_created_date, &ppdm_metric_value.Row_effective_date, &ppdm_metric_value.Row_expiry_date, &ppdm_metric_value.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ppdm_metric_value to result
		result = append(result, ppdm_metric_value)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPpdmMetricValue(c *fiber.Ctx) error {
	var ppdm_metric_value dto.Ppdm_metric_value

	setDefaults(&ppdm_metric_value)

	if err := json.Unmarshal(c.Body(), &ppdm_metric_value); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ppdm_metric_value values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32, :33, :34, :35, :36, :37, :38, :39)")
	if err != nil {
		return err
	}
	ppdm_metric_value.Row_created_date = formatDate(ppdm_metric_value.Row_created_date)
	ppdm_metric_value.Row_changed_date = formatDate(ppdm_metric_value.Row_changed_date)
	ppdm_metric_value.Date_format_desc = formatDateString(ppdm_metric_value.Date_format_desc)
	ppdm_metric_value.Effective_date = formatDateString(ppdm_metric_value.Effective_date)
	ppdm_metric_value.Expiry_date = formatDateString(ppdm_metric_value.Expiry_date)
	ppdm_metric_value.Max_date = formatDateString(ppdm_metric_value.Max_date)
	ppdm_metric_value.Min_date = formatDateString(ppdm_metric_value.Min_date)
	ppdm_metric_value.Row_effective_date = formatDateString(ppdm_metric_value.Row_effective_date)
	ppdm_metric_value.Row_expiry_date = formatDateString(ppdm_metric_value.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_metric_value.Metric_id, ppdm_metric_value.Metric_obs_no, ppdm_metric_value.Active_ind, ppdm_metric_value.Average_value, ppdm_metric_value.Average_value_ouom, ppdm_metric_value.Average_value_uom, ppdm_metric_value.Cost, ppdm_metric_value.Currency_conversion, ppdm_metric_value.Currency_ouom, ppdm_metric_value.Currency_uom, ppdm_metric_value.Date_format_desc, ppdm_metric_value.Effective_date, ppdm_metric_value.Expiry_date, ppdm_metric_value.Max_date, ppdm_metric_value.Max_value, ppdm_metric_value.Max_value_ouom, ppdm_metric_value.Max_value_uom, ppdm_metric_value.Metric_code, ppdm_metric_value.Metric_desc, ppdm_metric_value.Metric_ind, ppdm_metric_value.Metric_type, ppdm_metric_value.Min_date, ppdm_metric_value.Min_value, ppdm_metric_value.Min_value_ouom, ppdm_metric_value.Min_value_uom, ppdm_metric_value.Ppdm_guid, ppdm_metric_value.Reference_value, ppdm_metric_value.Reference_value_ouom, ppdm_metric_value.Reference_value_type, ppdm_metric_value.Reference_value_uom, ppdm_metric_value.Remark, ppdm_metric_value.Source, ppdm_metric_value.Row_changed_by, ppdm_metric_value.Row_changed_date, ppdm_metric_value.Row_created_by, ppdm_metric_value.Row_created_date, ppdm_metric_value.Row_effective_date, ppdm_metric_value.Row_expiry_date, ppdm_metric_value.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePpdmMetricValue(c *fiber.Ctx) error {
	var ppdm_metric_value dto.Ppdm_metric_value

	setDefaults(&ppdm_metric_value)

	if err := json.Unmarshal(c.Body(), &ppdm_metric_value); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ppdm_metric_value.Metric_id = id

    if ppdm_metric_value.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ppdm_metric_value set metric_obs_no = :1, active_ind = :2, average_value = :3, average_value_ouom = :4, average_value_uom = :5, cost = :6, currency_conversion = :7, currency_ouom = :8, currency_uom = :9, date_format_desc = :10, effective_date = :11, expiry_date = :12, max_date = :13, max_value = :14, max_value_ouom = :15, max_value_uom = :16, metric_code = :17, metric_desc = :18, metric_ind = :19, metric_type = :20, min_date = :21, min_value = :22, min_value_ouom = :23, min_value_uom = :24, ppdm_guid = :25, reference_value = :26, reference_value_ouom = :27, reference_value_type = :28, reference_value_uom = :29, remark = :30, source = :31, row_changed_by = :32, row_changed_date = :33, row_created_by = :34, row_effective_date = :35, row_expiry_date = :36, row_quality = :37 where metric_id = :39")
	if err != nil {
		return err
	}

	ppdm_metric_value.Row_changed_date = formatDate(ppdm_metric_value.Row_changed_date)
	ppdm_metric_value.Date_format_desc = formatDateString(ppdm_metric_value.Date_format_desc)
	ppdm_metric_value.Effective_date = formatDateString(ppdm_metric_value.Effective_date)
	ppdm_metric_value.Expiry_date = formatDateString(ppdm_metric_value.Expiry_date)
	ppdm_metric_value.Max_date = formatDateString(ppdm_metric_value.Max_date)
	ppdm_metric_value.Min_date = formatDateString(ppdm_metric_value.Min_date)
	ppdm_metric_value.Row_effective_date = formatDateString(ppdm_metric_value.Row_effective_date)
	ppdm_metric_value.Row_expiry_date = formatDateString(ppdm_metric_value.Row_expiry_date)






	rows, err := stmt.Exec(ppdm_metric_value.Metric_obs_no, ppdm_metric_value.Active_ind, ppdm_metric_value.Average_value, ppdm_metric_value.Average_value_ouom, ppdm_metric_value.Average_value_uom, ppdm_metric_value.Cost, ppdm_metric_value.Currency_conversion, ppdm_metric_value.Currency_ouom, ppdm_metric_value.Currency_uom, ppdm_metric_value.Date_format_desc, ppdm_metric_value.Effective_date, ppdm_metric_value.Expiry_date, ppdm_metric_value.Max_date, ppdm_metric_value.Max_value, ppdm_metric_value.Max_value_ouom, ppdm_metric_value.Max_value_uom, ppdm_metric_value.Metric_code, ppdm_metric_value.Metric_desc, ppdm_metric_value.Metric_ind, ppdm_metric_value.Metric_type, ppdm_metric_value.Min_date, ppdm_metric_value.Min_value, ppdm_metric_value.Min_value_ouom, ppdm_metric_value.Min_value_uom, ppdm_metric_value.Ppdm_guid, ppdm_metric_value.Reference_value, ppdm_metric_value.Reference_value_ouom, ppdm_metric_value.Reference_value_type, ppdm_metric_value.Reference_value_uom, ppdm_metric_value.Remark, ppdm_metric_value.Source, ppdm_metric_value.Row_changed_by, ppdm_metric_value.Row_changed_date, ppdm_metric_value.Row_created_by, ppdm_metric_value.Row_effective_date, ppdm_metric_value.Row_expiry_date, ppdm_metric_value.Row_quality, ppdm_metric_value.Metric_id)
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

func PatchPpdmMetricValue(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ppdm_metric_value set "
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
		} else if key == "date_format_desc" || key == "effective_date" || key == "expiry_date" || key == "max_date" || key == "min_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where metric_id = :id"

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

func DeletePpdmMetricValue(c *fiber.Ctx) error {
	id := c.Params("id")
	var ppdm_metric_value dto.Ppdm_metric_value
	ppdm_metric_value.Metric_id = id

	stmt, err := db.Prepare("delete from ppdm_metric_value where metric_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ppdm_metric_value.Metric_id)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetProdStrStatHist(c *fiber.Ctx) error {
	rows, err := db.Query("select * from prod_str_stat_hist")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Prod_str_stat_hist

	for rows.Next() {
		var prod_str_stat_hist dto.Prod_str_stat_hist
		if err := rows.Scan(&prod_str_stat_hist.Uwi, &prod_str_stat_hist.Prod_string_source, &prod_str_stat_hist.String_id, &prod_str_stat_hist.Status_id, &prod_str_stat_hist.Active_ind, &prod_str_stat_hist.Effective_date, &prod_str_stat_hist.End_time, &prod_str_stat_hist.Expiry_date, &prod_str_stat_hist.Percent_capability, &prod_str_stat_hist.Ppdm_guid, &prod_str_stat_hist.Remark, &prod_str_stat_hist.Source, &prod_str_stat_hist.Start_time, &prod_str_stat_hist.Status, &prod_str_stat_hist.Status_date, &prod_str_stat_hist.Status_type, &prod_str_stat_hist.Timezone, &prod_str_stat_hist.Row_changed_by, &prod_str_stat_hist.Row_changed_date, &prod_str_stat_hist.Row_created_by, &prod_str_stat_hist.Row_created_date, &prod_str_stat_hist.Row_effective_date, &prod_str_stat_hist.Row_expiry_date, &prod_str_stat_hist.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append prod_str_stat_hist to result
		result = append(result, prod_str_stat_hist)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetProdStrStatHist(c *fiber.Ctx) error {
	var prod_str_stat_hist dto.Prod_str_stat_hist

	setDefaults(&prod_str_stat_hist)

	if err := json.Unmarshal(c.Body(), &prod_str_stat_hist); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into prod_str_stat_hist values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24)")
	if err != nil {
		return err
	}
	prod_str_stat_hist.Row_created_date = formatDate(prod_str_stat_hist.Row_created_date)
	prod_str_stat_hist.Row_changed_date = formatDate(prod_str_stat_hist.Row_changed_date)
	prod_str_stat_hist.Effective_date = formatDateString(prod_str_stat_hist.Effective_date)
	prod_str_stat_hist.End_time = formatDateString(prod_str_stat_hist.End_time)
	prod_str_stat_hist.Expiry_date = formatDateString(prod_str_stat_hist.Expiry_date)
	prod_str_stat_hist.Start_time = formatDateString(prod_str_stat_hist.Start_time)
	prod_str_stat_hist.Status_date = formatDateString(prod_str_stat_hist.Status_date)
	prod_str_stat_hist.Row_effective_date = formatDateString(prod_str_stat_hist.Row_effective_date)
	prod_str_stat_hist.Row_expiry_date = formatDateString(prod_str_stat_hist.Row_expiry_date)






	rows, err := stmt.Exec(prod_str_stat_hist.Uwi, prod_str_stat_hist.Prod_string_source, prod_str_stat_hist.String_id, prod_str_stat_hist.Status_id, prod_str_stat_hist.Active_ind, prod_str_stat_hist.Effective_date, prod_str_stat_hist.End_time, prod_str_stat_hist.Expiry_date, prod_str_stat_hist.Percent_capability, prod_str_stat_hist.Ppdm_guid, prod_str_stat_hist.Remark, prod_str_stat_hist.Source, prod_str_stat_hist.Start_time, prod_str_stat_hist.Status, prod_str_stat_hist.Status_date, prod_str_stat_hist.Status_type, prod_str_stat_hist.Timezone, prod_str_stat_hist.Row_changed_by, prod_str_stat_hist.Row_changed_date, prod_str_stat_hist.Row_created_by, prod_str_stat_hist.Row_created_date, prod_str_stat_hist.Row_effective_date, prod_str_stat_hist.Row_expiry_date, prod_str_stat_hist.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateProdStrStatHist(c *fiber.Ctx) error {
	var prod_str_stat_hist dto.Prod_str_stat_hist

	setDefaults(&prod_str_stat_hist)

	if err := json.Unmarshal(c.Body(), &prod_str_stat_hist); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	prod_str_stat_hist.Uwi = id

    if prod_str_stat_hist.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update prod_str_stat_hist set prod_string_source = :1, string_id = :2, status_id = :3, active_ind = :4, effective_date = :5, end_time = :6, expiry_date = :7, percent_capability = :8, ppdm_guid = :9, remark = :10, source = :11, start_time = :12, status = :13, status_date = :14, status_type = :15, timezone = :16, row_changed_by = :17, row_changed_date = :18, row_created_by = :19, row_effective_date = :20, row_expiry_date = :21, row_quality = :22 where uwi = :24")
	if err != nil {
		return err
	}

	prod_str_stat_hist.Row_changed_date = formatDate(prod_str_stat_hist.Row_changed_date)
	prod_str_stat_hist.Effective_date = formatDateString(prod_str_stat_hist.Effective_date)
	prod_str_stat_hist.End_time = formatDateString(prod_str_stat_hist.End_time)
	prod_str_stat_hist.Expiry_date = formatDateString(prod_str_stat_hist.Expiry_date)
	prod_str_stat_hist.Start_time = formatDateString(prod_str_stat_hist.Start_time)
	prod_str_stat_hist.Status_date = formatDateString(prod_str_stat_hist.Status_date)
	prod_str_stat_hist.Row_effective_date = formatDateString(prod_str_stat_hist.Row_effective_date)
	prod_str_stat_hist.Row_expiry_date = formatDateString(prod_str_stat_hist.Row_expiry_date)






	rows, err := stmt.Exec(prod_str_stat_hist.Prod_string_source, prod_str_stat_hist.String_id, prod_str_stat_hist.Status_id, prod_str_stat_hist.Active_ind, prod_str_stat_hist.Effective_date, prod_str_stat_hist.End_time, prod_str_stat_hist.Expiry_date, prod_str_stat_hist.Percent_capability, prod_str_stat_hist.Ppdm_guid, prod_str_stat_hist.Remark, prod_str_stat_hist.Source, prod_str_stat_hist.Start_time, prod_str_stat_hist.Status, prod_str_stat_hist.Status_date, prod_str_stat_hist.Status_type, prod_str_stat_hist.Timezone, prod_str_stat_hist.Row_changed_by, prod_str_stat_hist.Row_changed_date, prod_str_stat_hist.Row_created_by, prod_str_stat_hist.Row_effective_date, prod_str_stat_hist.Row_expiry_date, prod_str_stat_hist.Row_quality, prod_str_stat_hist.Uwi)
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

func PatchProdStrStatHist(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update prod_str_stat_hist set "
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

func DeleteProdStrStatHist(c *fiber.Ctx) error {
	id := c.Params("id")
	var prod_str_stat_hist dto.Prod_str_stat_hist
	prod_str_stat_hist.Uwi = id

	stmt, err := db.Prepare("delete from prod_str_stat_hist where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(prod_str_stat_hist.Uwi)
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



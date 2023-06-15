package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPdenStatusHist(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pden_status_hist")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pden_status_hist

	for rows.Next() {
		var pden_status_hist dto.Pden_status_hist
		if err := rows.Scan(&pden_status_hist.Pden_subtype, &pden_status_hist.Pden_id, &pden_status_hist.Pden_source, &pden_status_hist.Status_id, &pden_status_hist.Active_ind, &pden_status_hist.Effective_date, &pden_status_hist.End_time, &pden_status_hist.Expiry_date, &pden_status_hist.Percent_capability, &pden_status_hist.Ppdm_guid, &pden_status_hist.Remark, &pden_status_hist.Source, &pden_status_hist.Start_time, &pden_status_hist.Status, &pden_status_hist.Status_date, &pden_status_hist.Status_type, &pden_status_hist.Timezone, &pden_status_hist.Row_changed_by, &pden_status_hist.Row_changed_date, &pden_status_hist.Row_created_by, &pden_status_hist.Row_created_date, &pden_status_hist.Row_effective_date, &pden_status_hist.Row_expiry_date, &pden_status_hist.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pden_status_hist to result
		result = append(result, pden_status_hist)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPdenStatusHist(c *fiber.Ctx) error {
	var pden_status_hist dto.Pden_status_hist

	setDefaults(&pden_status_hist)

	if err := json.Unmarshal(c.Body(), &pden_status_hist); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pden_status_hist values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24)")
	if err != nil {
		return err
	}
	pden_status_hist.Row_created_date = formatDate(pden_status_hist.Row_created_date)
	pden_status_hist.Row_changed_date = formatDate(pden_status_hist.Row_changed_date)
	pden_status_hist.Effective_date = formatDateString(pden_status_hist.Effective_date)
	pden_status_hist.End_time = formatDateString(pden_status_hist.End_time)
	pden_status_hist.Expiry_date = formatDateString(pden_status_hist.Expiry_date)
	pden_status_hist.Start_time = formatDateString(pden_status_hist.Start_time)
	pden_status_hist.Status_date = formatDateString(pden_status_hist.Status_date)
	pden_status_hist.Row_effective_date = formatDateString(pden_status_hist.Row_effective_date)
	pden_status_hist.Row_expiry_date = formatDateString(pden_status_hist.Row_expiry_date)






	rows, err := stmt.Exec(pden_status_hist.Pden_subtype, pden_status_hist.Pden_id, pden_status_hist.Pden_source, pden_status_hist.Status_id, pden_status_hist.Active_ind, pden_status_hist.Effective_date, pden_status_hist.End_time, pden_status_hist.Expiry_date, pden_status_hist.Percent_capability, pden_status_hist.Ppdm_guid, pden_status_hist.Remark, pden_status_hist.Source, pden_status_hist.Start_time, pden_status_hist.Status, pden_status_hist.Status_date, pden_status_hist.Status_type, pden_status_hist.Timezone, pden_status_hist.Row_changed_by, pden_status_hist.Row_changed_date, pden_status_hist.Row_created_by, pden_status_hist.Row_created_date, pden_status_hist.Row_effective_date, pden_status_hist.Row_expiry_date, pden_status_hist.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePdenStatusHist(c *fiber.Ctx) error {
	var pden_status_hist dto.Pden_status_hist

	setDefaults(&pden_status_hist)

	if err := json.Unmarshal(c.Body(), &pden_status_hist); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pden_status_hist.Pden_subtype = id

    if pden_status_hist.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pden_status_hist set pden_id = :1, pden_source = :2, status_id = :3, active_ind = :4, effective_date = :5, end_time = :6, expiry_date = :7, percent_capability = :8, ppdm_guid = :9, remark = :10, source = :11, start_time = :12, status = :13, status_date = :14, status_type = :15, timezone = :16, row_changed_by = :17, row_changed_date = :18, row_created_by = :19, row_effective_date = :20, row_expiry_date = :21, row_quality = :22 where pden_subtype = :24")
	if err != nil {
		return err
	}

	pden_status_hist.Row_changed_date = formatDate(pden_status_hist.Row_changed_date)
	pden_status_hist.Effective_date = formatDateString(pden_status_hist.Effective_date)
	pden_status_hist.End_time = formatDateString(pden_status_hist.End_time)
	pden_status_hist.Expiry_date = formatDateString(pden_status_hist.Expiry_date)
	pden_status_hist.Start_time = formatDateString(pden_status_hist.Start_time)
	pden_status_hist.Status_date = formatDateString(pden_status_hist.Status_date)
	pden_status_hist.Row_effective_date = formatDateString(pden_status_hist.Row_effective_date)
	pden_status_hist.Row_expiry_date = formatDateString(pden_status_hist.Row_expiry_date)






	rows, err := stmt.Exec(pden_status_hist.Pden_id, pden_status_hist.Pden_source, pden_status_hist.Status_id, pden_status_hist.Active_ind, pden_status_hist.Effective_date, pden_status_hist.End_time, pden_status_hist.Expiry_date, pden_status_hist.Percent_capability, pden_status_hist.Ppdm_guid, pden_status_hist.Remark, pden_status_hist.Source, pden_status_hist.Start_time, pden_status_hist.Status, pden_status_hist.Status_date, pden_status_hist.Status_type, pden_status_hist.Timezone, pden_status_hist.Row_changed_by, pden_status_hist.Row_changed_date, pden_status_hist.Row_created_by, pden_status_hist.Row_effective_date, pden_status_hist.Row_expiry_date, pden_status_hist.Row_quality, pden_status_hist.Pden_subtype)
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

func PatchPdenStatusHist(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pden_status_hist set "
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
	query += " where pden_subtype = :id"

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

func DeletePdenStatusHist(c *fiber.Ctx) error {
	id := c.Params("id")
	var pden_status_hist dto.Pden_status_hist
	pden_status_hist.Pden_subtype = id

	stmt, err := db.Prepare("delete from pden_status_hist where pden_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pden_status_hist.Pden_subtype)
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



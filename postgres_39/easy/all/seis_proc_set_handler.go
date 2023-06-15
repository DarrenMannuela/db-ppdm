package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisProcSet(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_proc_set")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_proc_set

	for rows.Next() {
		var seis_proc_set dto.Seis_proc_set
		if err := rows.Scan(&seis_proc_set.Seis_set_subtype, &seis_proc_set.Seis_proc_set_id, &seis_proc_set.Active_ind, &seis_proc_set.Area_id, &seis_proc_set.Area_type, &seis_proc_set.Description, &seis_proc_set.Effective_date, &seis_proc_set.End_date, &seis_proc_set.Expiry_date, &seis_proc_set.Max_latitude, &seis_proc_set.Max_longitude, &seis_proc_set.Min_latitude, &seis_proc_set.Min_longitude, &seis_proc_set.Objective, &seis_proc_set.Original_proc_ind, &seis_proc_set.Ppdm_guid, &seis_proc_set.Processed_for, &seis_proc_set.Processing_company, &seis_proc_set.Processing_name, &seis_proc_set.Process_status, &seis_proc_set.Proc_set_type, &seis_proc_set.Remark, &seis_proc_set.Reprocessed_ind, &seis_proc_set.Source, &seis_proc_set.Start_date, &seis_proc_set.Row_changed_by, &seis_proc_set.Row_changed_date, &seis_proc_set.Row_created_by, &seis_proc_set.Row_created_date, &seis_proc_set.Row_effective_date, &seis_proc_set.Row_expiry_date, &seis_proc_set.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_proc_set to result
		result = append(result, seis_proc_set)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisProcSet(c *fiber.Ctx) error {
	var seis_proc_set dto.Seis_proc_set

	setDefaults(&seis_proc_set)

	if err := json.Unmarshal(c.Body(), &seis_proc_set); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_proc_set values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30, :31, :32)")
	if err != nil {
		return err
	}
	seis_proc_set.Row_created_date = formatDate(seis_proc_set.Row_created_date)
	seis_proc_set.Row_changed_date = formatDate(seis_proc_set.Row_changed_date)
	seis_proc_set.Effective_date = formatDateString(seis_proc_set.Effective_date)
	seis_proc_set.End_date = formatDateString(seis_proc_set.End_date)
	seis_proc_set.Expiry_date = formatDateString(seis_proc_set.Expiry_date)
	seis_proc_set.Start_date = formatDateString(seis_proc_set.Start_date)
	seis_proc_set.Row_effective_date = formatDateString(seis_proc_set.Row_effective_date)
	seis_proc_set.Row_expiry_date = formatDateString(seis_proc_set.Row_expiry_date)






	rows, err := stmt.Exec(seis_proc_set.Seis_set_subtype, seis_proc_set.Seis_proc_set_id, seis_proc_set.Active_ind, seis_proc_set.Area_id, seis_proc_set.Area_type, seis_proc_set.Description, seis_proc_set.Effective_date, seis_proc_set.End_date, seis_proc_set.Expiry_date, seis_proc_set.Max_latitude, seis_proc_set.Max_longitude, seis_proc_set.Min_latitude, seis_proc_set.Min_longitude, seis_proc_set.Objective, seis_proc_set.Original_proc_ind, seis_proc_set.Ppdm_guid, seis_proc_set.Processed_for, seis_proc_set.Processing_company, seis_proc_set.Processing_name, seis_proc_set.Process_status, seis_proc_set.Proc_set_type, seis_proc_set.Remark, seis_proc_set.Reprocessed_ind, seis_proc_set.Source, seis_proc_set.Start_date, seis_proc_set.Row_changed_by, seis_proc_set.Row_changed_date, seis_proc_set.Row_created_by, seis_proc_set.Row_created_date, seis_proc_set.Row_effective_date, seis_proc_set.Row_expiry_date, seis_proc_set.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisProcSet(c *fiber.Ctx) error {
	var seis_proc_set dto.Seis_proc_set

	setDefaults(&seis_proc_set)

	if err := json.Unmarshal(c.Body(), &seis_proc_set); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_proc_set.Seis_set_subtype = id

    if seis_proc_set.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_proc_set set seis_proc_set_id = :1, active_ind = :2, area_id = :3, area_type = :4, description = :5, effective_date = :6, end_date = :7, expiry_date = :8, max_latitude = :9, max_longitude = :10, min_latitude = :11, min_longitude = :12, objective = :13, original_proc_ind = :14, ppdm_guid = :15, processed_for = :16, processing_company = :17, processing_name = :18, process_status = :19, proc_set_type = :20, remark = :21, reprocessed_ind = :22, source = :23, start_date = :24, row_changed_by = :25, row_changed_date = :26, row_created_by = :27, row_effective_date = :28, row_expiry_date = :29, row_quality = :30 where seis_set_subtype = :32")
	if err != nil {
		return err
	}

	seis_proc_set.Row_changed_date = formatDate(seis_proc_set.Row_changed_date)
	seis_proc_set.Effective_date = formatDateString(seis_proc_set.Effective_date)
	seis_proc_set.End_date = formatDateString(seis_proc_set.End_date)
	seis_proc_set.Expiry_date = formatDateString(seis_proc_set.Expiry_date)
	seis_proc_set.Start_date = formatDateString(seis_proc_set.Start_date)
	seis_proc_set.Row_effective_date = formatDateString(seis_proc_set.Row_effective_date)
	seis_proc_set.Row_expiry_date = formatDateString(seis_proc_set.Row_expiry_date)






	rows, err := stmt.Exec(seis_proc_set.Seis_proc_set_id, seis_proc_set.Active_ind, seis_proc_set.Area_id, seis_proc_set.Area_type, seis_proc_set.Description, seis_proc_set.Effective_date, seis_proc_set.End_date, seis_proc_set.Expiry_date, seis_proc_set.Max_latitude, seis_proc_set.Max_longitude, seis_proc_set.Min_latitude, seis_proc_set.Min_longitude, seis_proc_set.Objective, seis_proc_set.Original_proc_ind, seis_proc_set.Ppdm_guid, seis_proc_set.Processed_for, seis_proc_set.Processing_company, seis_proc_set.Processing_name, seis_proc_set.Process_status, seis_proc_set.Proc_set_type, seis_proc_set.Remark, seis_proc_set.Reprocessed_ind, seis_proc_set.Source, seis_proc_set.Start_date, seis_proc_set.Row_changed_by, seis_proc_set.Row_changed_date, seis_proc_set.Row_created_by, seis_proc_set.Row_effective_date, seis_proc_set.Row_expiry_date, seis_proc_set.Row_quality, seis_proc_set.Seis_set_subtype)
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

func PatchSeisProcSet(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_proc_set set "
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
		} else if key == "effective_date" || key == "end_date" || key == "expiry_date" || key == "start_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where seis_set_subtype = :id"

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

func DeleteSeisProcSet(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_proc_set dto.Seis_proc_set
	seis_proc_set.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_proc_set where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_proc_set.Seis_set_subtype)
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



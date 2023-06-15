package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetSeisBinPointTrace(c *fiber.Ctx) error {
	rows, err := db.Query("select * from seis_bin_point_trace")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Seis_bin_point_trace

	for rows.Next() {
		var seis_bin_point_trace dto.Seis_bin_point_trace
		if err := rows.Scan(&seis_bin_point_trace.Seis_set_subtype, &seis_bin_point_trace.Seis_set_id, &seis_bin_point_trace.Bin_grid_id, &seis_bin_point_trace.Bin_grid_source, &seis_bin_point_trace.Bin_point_id, &seis_bin_point_trace.Trace_id, &seis_bin_point_trace.Active_ind, &seis_bin_point_trace.Effective_date, &seis_bin_point_trace.Expiry_date, &seis_bin_point_trace.Information_item_id, &seis_bin_point_trace.Info_item_subtype, &seis_bin_point_trace.Ppdm_guid, &seis_bin_point_trace.Remark, &seis_bin_point_trace.Source, &seis_bin_point_trace.Row_changed_by, &seis_bin_point_trace.Row_changed_date, &seis_bin_point_trace.Row_created_by, &seis_bin_point_trace.Row_created_date, &seis_bin_point_trace.Row_effective_date, &seis_bin_point_trace.Row_expiry_date, &seis_bin_point_trace.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append seis_bin_point_trace to result
		result = append(result, seis_bin_point_trace)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetSeisBinPointTrace(c *fiber.Ctx) error {
	var seis_bin_point_trace dto.Seis_bin_point_trace

	setDefaults(&seis_bin_point_trace)

	if err := json.Unmarshal(c.Body(), &seis_bin_point_trace); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into seis_bin_point_trace values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	seis_bin_point_trace.Row_created_date = formatDate(seis_bin_point_trace.Row_created_date)
	seis_bin_point_trace.Row_changed_date = formatDate(seis_bin_point_trace.Row_changed_date)
	seis_bin_point_trace.Effective_date = formatDateString(seis_bin_point_trace.Effective_date)
	seis_bin_point_trace.Expiry_date = formatDateString(seis_bin_point_trace.Expiry_date)
	seis_bin_point_trace.Row_effective_date = formatDateString(seis_bin_point_trace.Row_effective_date)
	seis_bin_point_trace.Row_expiry_date = formatDateString(seis_bin_point_trace.Row_expiry_date)






	rows, err := stmt.Exec(seis_bin_point_trace.Seis_set_subtype, seis_bin_point_trace.Seis_set_id, seis_bin_point_trace.Bin_grid_id, seis_bin_point_trace.Bin_grid_source, seis_bin_point_trace.Bin_point_id, seis_bin_point_trace.Trace_id, seis_bin_point_trace.Active_ind, seis_bin_point_trace.Effective_date, seis_bin_point_trace.Expiry_date, seis_bin_point_trace.Information_item_id, seis_bin_point_trace.Info_item_subtype, seis_bin_point_trace.Ppdm_guid, seis_bin_point_trace.Remark, seis_bin_point_trace.Source, seis_bin_point_trace.Row_changed_by, seis_bin_point_trace.Row_changed_date, seis_bin_point_trace.Row_created_by, seis_bin_point_trace.Row_created_date, seis_bin_point_trace.Row_effective_date, seis_bin_point_trace.Row_expiry_date, seis_bin_point_trace.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateSeisBinPointTrace(c *fiber.Ctx) error {
	var seis_bin_point_trace dto.Seis_bin_point_trace

	setDefaults(&seis_bin_point_trace)

	if err := json.Unmarshal(c.Body(), &seis_bin_point_trace); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	seis_bin_point_trace.Seis_set_subtype = id

    if seis_bin_point_trace.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update seis_bin_point_trace set seis_set_id = :1, bin_grid_id = :2, bin_grid_source = :3, bin_point_id = :4, trace_id = :5, active_ind = :6, effective_date = :7, expiry_date = :8, information_item_id = :9, info_item_subtype = :10, ppdm_guid = :11, remark = :12, source = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where seis_set_subtype = :21")
	if err != nil {
		return err
	}

	seis_bin_point_trace.Row_changed_date = formatDate(seis_bin_point_trace.Row_changed_date)
	seis_bin_point_trace.Effective_date = formatDateString(seis_bin_point_trace.Effective_date)
	seis_bin_point_trace.Expiry_date = formatDateString(seis_bin_point_trace.Expiry_date)
	seis_bin_point_trace.Row_effective_date = formatDateString(seis_bin_point_trace.Row_effective_date)
	seis_bin_point_trace.Row_expiry_date = formatDateString(seis_bin_point_trace.Row_expiry_date)






	rows, err := stmt.Exec(seis_bin_point_trace.Seis_set_id, seis_bin_point_trace.Bin_grid_id, seis_bin_point_trace.Bin_grid_source, seis_bin_point_trace.Bin_point_id, seis_bin_point_trace.Trace_id, seis_bin_point_trace.Active_ind, seis_bin_point_trace.Effective_date, seis_bin_point_trace.Expiry_date, seis_bin_point_trace.Information_item_id, seis_bin_point_trace.Info_item_subtype, seis_bin_point_trace.Ppdm_guid, seis_bin_point_trace.Remark, seis_bin_point_trace.Source, seis_bin_point_trace.Row_changed_by, seis_bin_point_trace.Row_changed_date, seis_bin_point_trace.Row_created_by, seis_bin_point_trace.Row_effective_date, seis_bin_point_trace.Row_expiry_date, seis_bin_point_trace.Row_quality, seis_bin_point_trace.Seis_set_subtype)
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

func PatchSeisBinPointTrace(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update seis_bin_point_trace set "
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

func DeleteSeisBinPointTrace(c *fiber.Ctx) error {
	id := c.Params("id")
	var seis_bin_point_trace dto.Seis_bin_point_trace
	seis_bin_point_trace.Seis_set_subtype = id

	stmt, err := db.Prepare("delete from seis_bin_point_trace where seis_set_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(seis_bin_point_trace.Seis_set_subtype)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPdenWellReportStream(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pden_well_report_stream")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pden_well_report_stream

	for rows.Next() {
		var pden_well_report_stream dto.Pden_well_report_stream
		if err := rows.Scan(&pden_well_report_stream.Pden_subtype, &pden_well_report_stream.Pden_id, &pden_well_report_stream.Pden_source, &pden_well_report_stream.Uwi, &pden_well_report_stream.Wrs_obs_no, &pden_well_report_stream.Active_ind, &pden_well_report_stream.Effective_date, &pden_well_report_stream.Expiry_date, &pden_well_report_stream.Ppdm_guid, &pden_well_report_stream.Remark, &pden_well_report_stream.Substance_id, &pden_well_report_stream.Substance_percent, &pden_well_report_stream.Row_changed_by, &pden_well_report_stream.Row_changed_date, &pden_well_report_stream.Row_created_by, &pden_well_report_stream.Row_created_date, &pden_well_report_stream.Row_effective_date, &pden_well_report_stream.Row_expiry_date, &pden_well_report_stream.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pden_well_report_stream to result
		result = append(result, pden_well_report_stream)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPdenWellReportStream(c *fiber.Ctx) error {
	var pden_well_report_stream dto.Pden_well_report_stream

	setDefaults(&pden_well_report_stream)

	if err := json.Unmarshal(c.Body(), &pden_well_report_stream); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pden_well_report_stream values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	pden_well_report_stream.Row_created_date = formatDate(pden_well_report_stream.Row_created_date)
	pden_well_report_stream.Row_changed_date = formatDate(pden_well_report_stream.Row_changed_date)
	pden_well_report_stream.Effective_date = formatDateString(pden_well_report_stream.Effective_date)
	pden_well_report_stream.Expiry_date = formatDateString(pden_well_report_stream.Expiry_date)
	pden_well_report_stream.Row_effective_date = formatDateString(pden_well_report_stream.Row_effective_date)
	pden_well_report_stream.Row_expiry_date = formatDateString(pden_well_report_stream.Row_expiry_date)






	rows, err := stmt.Exec(pden_well_report_stream.Pden_subtype, pden_well_report_stream.Pden_id, pden_well_report_stream.Pden_source, pden_well_report_stream.Uwi, pden_well_report_stream.Wrs_obs_no, pden_well_report_stream.Active_ind, pden_well_report_stream.Effective_date, pden_well_report_stream.Expiry_date, pden_well_report_stream.Ppdm_guid, pden_well_report_stream.Remark, pden_well_report_stream.Substance_id, pden_well_report_stream.Substance_percent, pden_well_report_stream.Row_changed_by, pden_well_report_stream.Row_changed_date, pden_well_report_stream.Row_created_by, pden_well_report_stream.Row_created_date, pden_well_report_stream.Row_effective_date, pden_well_report_stream.Row_expiry_date, pden_well_report_stream.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePdenWellReportStream(c *fiber.Ctx) error {
	var pden_well_report_stream dto.Pden_well_report_stream

	setDefaults(&pden_well_report_stream)

	if err := json.Unmarshal(c.Body(), &pden_well_report_stream); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pden_well_report_stream.Pden_subtype = id

    if pden_well_report_stream.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pden_well_report_stream set pden_id = :1, pden_source = :2, uwi = :3, wrs_obs_no = :4, active_ind = :5, effective_date = :6, expiry_date = :7, ppdm_guid = :8, remark = :9, substance_id = :10, substance_percent = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where pden_subtype = :19")
	if err != nil {
		return err
	}

	pden_well_report_stream.Row_changed_date = formatDate(pden_well_report_stream.Row_changed_date)
	pden_well_report_stream.Effective_date = formatDateString(pden_well_report_stream.Effective_date)
	pden_well_report_stream.Expiry_date = formatDateString(pden_well_report_stream.Expiry_date)
	pden_well_report_stream.Row_effective_date = formatDateString(pden_well_report_stream.Row_effective_date)
	pden_well_report_stream.Row_expiry_date = formatDateString(pden_well_report_stream.Row_expiry_date)






	rows, err := stmt.Exec(pden_well_report_stream.Pden_id, pden_well_report_stream.Pden_source, pden_well_report_stream.Uwi, pden_well_report_stream.Wrs_obs_no, pden_well_report_stream.Active_ind, pden_well_report_stream.Effective_date, pden_well_report_stream.Expiry_date, pden_well_report_stream.Ppdm_guid, pden_well_report_stream.Remark, pden_well_report_stream.Substance_id, pden_well_report_stream.Substance_percent, pden_well_report_stream.Row_changed_by, pden_well_report_stream.Row_changed_date, pden_well_report_stream.Row_created_by, pden_well_report_stream.Row_effective_date, pden_well_report_stream.Row_expiry_date, pden_well_report_stream.Row_quality, pden_well_report_stream.Pden_subtype)
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

func PatchPdenWellReportStream(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pden_well_report_stream set "
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

func DeletePdenWellReportStream(c *fiber.Ctx) error {
	id := c.Params("id")
	var pden_well_report_stream dto.Pden_well_report_stream
	pden_well_report_stream.Pden_subtype = id

	stmt, err := db.Prepare("delete from pden_well_report_stream where pden_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pden_well_report_stream.Pden_subtype)
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



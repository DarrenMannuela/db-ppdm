package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLithMeasuredSec(c *fiber.Ctx) error {
	rows, err := db.Query("select * from lith_measured_sec")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Lith_measured_sec

	for rows.Next() {
		var lith_measured_sec dto.Lith_measured_sec
		if err := rows.Scan(&lith_measured_sec.Meas_section_id, &lith_measured_sec.Source, &lith_measured_sec.Active_ind, &lith_measured_sec.Author, &lith_measured_sec.Description, &lith_measured_sec.Effective_date, &lith_measured_sec.Expiry_date, &lith_measured_sec.Location_desc, &lith_measured_sec.Location_qualifier, &lith_measured_sec.Node_id, &lith_measured_sec.Ppdm_guid, &lith_measured_sec.Publication_reference_text, &lith_measured_sec.Reference_date, &lith_measured_sec.Remark, &lith_measured_sec.Source_document_id, &lith_measured_sec.Strat_name_set_id, &lith_measured_sec.Strat_unit_id, &lith_measured_sec.Row_changed_by, &lith_measured_sec.Row_changed_date, &lith_measured_sec.Row_created_by, &lith_measured_sec.Row_created_date, &lith_measured_sec.Row_effective_date, &lith_measured_sec.Row_expiry_date, &lith_measured_sec.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append lith_measured_sec to result
		result = append(result, lith_measured_sec)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLithMeasuredSec(c *fiber.Ctx) error {
	var lith_measured_sec dto.Lith_measured_sec

	setDefaults(&lith_measured_sec)

	if err := json.Unmarshal(c.Body(), &lith_measured_sec); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into lith_measured_sec values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24)")
	if err != nil {
		return err
	}
	lith_measured_sec.Row_created_date = formatDate(lith_measured_sec.Row_created_date)
	lith_measured_sec.Row_changed_date = formatDate(lith_measured_sec.Row_changed_date)
	lith_measured_sec.Effective_date = formatDateString(lith_measured_sec.Effective_date)
	lith_measured_sec.Expiry_date = formatDateString(lith_measured_sec.Expiry_date)
	lith_measured_sec.Reference_date = formatDateString(lith_measured_sec.Reference_date)
	lith_measured_sec.Row_effective_date = formatDateString(lith_measured_sec.Row_effective_date)
	lith_measured_sec.Row_expiry_date = formatDateString(lith_measured_sec.Row_expiry_date)






	rows, err := stmt.Exec(lith_measured_sec.Meas_section_id, lith_measured_sec.Source, lith_measured_sec.Active_ind, lith_measured_sec.Author, lith_measured_sec.Description, lith_measured_sec.Effective_date, lith_measured_sec.Expiry_date, lith_measured_sec.Location_desc, lith_measured_sec.Location_qualifier, lith_measured_sec.Node_id, lith_measured_sec.Ppdm_guid, lith_measured_sec.Publication_reference_text, lith_measured_sec.Reference_date, lith_measured_sec.Remark, lith_measured_sec.Source_document_id, lith_measured_sec.Strat_name_set_id, lith_measured_sec.Strat_unit_id, lith_measured_sec.Row_changed_by, lith_measured_sec.Row_changed_date, lith_measured_sec.Row_created_by, lith_measured_sec.Row_created_date, lith_measured_sec.Row_effective_date, lith_measured_sec.Row_expiry_date, lith_measured_sec.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLithMeasuredSec(c *fiber.Ctx) error {
	var lith_measured_sec dto.Lith_measured_sec

	setDefaults(&lith_measured_sec)

	if err := json.Unmarshal(c.Body(), &lith_measured_sec); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	lith_measured_sec.Meas_section_id = id

    if lith_measured_sec.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update lith_measured_sec set source = :1, active_ind = :2, author = :3, description = :4, effective_date = :5, expiry_date = :6, location_desc = :7, location_qualifier = :8, node_id = :9, ppdm_guid = :10, publication_reference_text = :11, reference_date = :12, remark = :13, source_document_id = :14, strat_name_set_id = :15, strat_unit_id = :16, row_changed_by = :17, row_changed_date = :18, row_created_by = :19, row_effective_date = :20, row_expiry_date = :21, row_quality = :22 where meas_section_id = :24")
	if err != nil {
		return err
	}

	lith_measured_sec.Row_changed_date = formatDate(lith_measured_sec.Row_changed_date)
	lith_measured_sec.Effective_date = formatDateString(lith_measured_sec.Effective_date)
	lith_measured_sec.Expiry_date = formatDateString(lith_measured_sec.Expiry_date)
	lith_measured_sec.Reference_date = formatDateString(lith_measured_sec.Reference_date)
	lith_measured_sec.Row_effective_date = formatDateString(lith_measured_sec.Row_effective_date)
	lith_measured_sec.Row_expiry_date = formatDateString(lith_measured_sec.Row_expiry_date)






	rows, err := stmt.Exec(lith_measured_sec.Source, lith_measured_sec.Active_ind, lith_measured_sec.Author, lith_measured_sec.Description, lith_measured_sec.Effective_date, lith_measured_sec.Expiry_date, lith_measured_sec.Location_desc, lith_measured_sec.Location_qualifier, lith_measured_sec.Node_id, lith_measured_sec.Ppdm_guid, lith_measured_sec.Publication_reference_text, lith_measured_sec.Reference_date, lith_measured_sec.Remark, lith_measured_sec.Source_document_id, lith_measured_sec.Strat_name_set_id, lith_measured_sec.Strat_unit_id, lith_measured_sec.Row_changed_by, lith_measured_sec.Row_changed_date, lith_measured_sec.Row_created_by, lith_measured_sec.Row_effective_date, lith_measured_sec.Row_expiry_date, lith_measured_sec.Row_quality, lith_measured_sec.Meas_section_id)
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

func PatchLithMeasuredSec(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update lith_measured_sec set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "reference_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where meas_section_id = :id"

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

func DeleteLithMeasuredSec(c *fiber.Ctx) error {
	id := c.Params("id")
	var lith_measured_sec dto.Lith_measured_sec
	lith_measured_sec.Meas_section_id = id

	stmt, err := db.Prepare("delete from lith_measured_sec where meas_section_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(lith_measured_sec.Meas_section_id)
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



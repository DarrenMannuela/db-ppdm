package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlRemark(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_remark")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_remark

	for rows.Next() {
		var anl_remark dto.Anl_remark
		if err := rows.Scan(&anl_remark.Analysis_id, &anl_remark.Anl_source, &anl_remark.Remark_seq_no, &anl_remark.Active_ind, &anl_remark.Effective_date, &anl_remark.Expiry_date, &anl_remark.Ppdm_guid, &anl_remark.Referenced_column_name, &anl_remark.Referenced_ppdm_guid, &anl_remark.Referenced_system_id, &anl_remark.Referenced_table_name, &anl_remark.Remark, &anl_remark.Remark_ba_id, &anl_remark.Remark_type, &anl_remark.Source, &anl_remark.Step_seq_no, &anl_remark.Row_changed_by, &anl_remark.Row_changed_date, &anl_remark.Row_created_by, &anl_remark.Row_created_date, &anl_remark.Row_effective_date, &anl_remark.Row_expiry_date, &anl_remark.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_remark to result
		result = append(result, anl_remark)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlRemark(c *fiber.Ctx) error {
	var anl_remark dto.Anl_remark

	setDefaults(&anl_remark)

	if err := json.Unmarshal(c.Body(), &anl_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_remark values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23)")
	if err != nil {
		return err
	}
	anl_remark.Row_created_date = formatDate(anl_remark.Row_created_date)
	anl_remark.Row_changed_date = formatDate(anl_remark.Row_changed_date)
	anl_remark.Effective_date = formatDateString(anl_remark.Effective_date)
	anl_remark.Expiry_date = formatDateString(anl_remark.Expiry_date)
	anl_remark.Row_effective_date = formatDateString(anl_remark.Row_effective_date)
	anl_remark.Row_expiry_date = formatDateString(anl_remark.Row_expiry_date)






	rows, err := stmt.Exec(anl_remark.Analysis_id, anl_remark.Anl_source, anl_remark.Remark_seq_no, anl_remark.Active_ind, anl_remark.Effective_date, anl_remark.Expiry_date, anl_remark.Ppdm_guid, anl_remark.Referenced_column_name, anl_remark.Referenced_ppdm_guid, anl_remark.Referenced_system_id, anl_remark.Referenced_table_name, anl_remark.Remark, anl_remark.Remark_ba_id, anl_remark.Remark_type, anl_remark.Source, anl_remark.Step_seq_no, anl_remark.Row_changed_by, anl_remark.Row_changed_date, anl_remark.Row_created_by, anl_remark.Row_created_date, anl_remark.Row_effective_date, anl_remark.Row_expiry_date, anl_remark.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlRemark(c *fiber.Ctx) error {
	var anl_remark dto.Anl_remark

	setDefaults(&anl_remark)

	if err := json.Unmarshal(c.Body(), &anl_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_remark.Analysis_id = id

    if anl_remark.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_remark set anl_source = :1, remark_seq_no = :2, active_ind = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, referenced_column_name = :7, referenced_ppdm_guid = :8, referenced_system_id = :9, referenced_table_name = :10, remark = :11, remark_ba_id = :12, remark_type = :13, source = :14, step_seq_no = :15, row_changed_by = :16, row_changed_date = :17, row_created_by = :18, row_effective_date = :19, row_expiry_date = :20, row_quality = :21 where analysis_id = :23")
	if err != nil {
		return err
	}

	anl_remark.Row_changed_date = formatDate(anl_remark.Row_changed_date)
	anl_remark.Effective_date = formatDateString(anl_remark.Effective_date)
	anl_remark.Expiry_date = formatDateString(anl_remark.Expiry_date)
	anl_remark.Row_effective_date = formatDateString(anl_remark.Row_effective_date)
	anl_remark.Row_expiry_date = formatDateString(anl_remark.Row_expiry_date)






	rows, err := stmt.Exec(anl_remark.Anl_source, anl_remark.Remark_seq_no, anl_remark.Active_ind, anl_remark.Effective_date, anl_remark.Expiry_date, anl_remark.Ppdm_guid, anl_remark.Referenced_column_name, anl_remark.Referenced_ppdm_guid, anl_remark.Referenced_system_id, anl_remark.Referenced_table_name, anl_remark.Remark, anl_remark.Remark_ba_id, anl_remark.Remark_type, anl_remark.Source, anl_remark.Step_seq_no, anl_remark.Row_changed_by, anl_remark.Row_changed_date, anl_remark.Row_created_by, anl_remark.Row_effective_date, anl_remark.Row_expiry_date, anl_remark.Row_quality, anl_remark.Analysis_id)
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

func PatchAnlRemark(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_remark set "
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
	query += " where analysis_id = :id"

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

func DeleteAnlRemark(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_remark dto.Anl_remark
	anl_remark.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_remark where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_remark.Analysis_id)
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



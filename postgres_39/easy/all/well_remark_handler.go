package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellRemark(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_remark")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_remark

	for rows.Next() {
		var well_remark dto.Well_remark
		if err := rows.Scan(&well_remark.Uwi, &well_remark.Source, &well_remark.Remark_type, &well_remark.Remark_seq_no, &well_remark.Active_ind, &well_remark.Effective_date, &well_remark.Expiry_date, &well_remark.Ppdm_guid, &well_remark.Remark, &well_remark.Remark_date, &well_remark.Remark_strat_unit_id, &well_remark.Strat_name_set_id, &well_remark.Row_changed_by, &well_remark.Row_changed_date, &well_remark.Row_created_by, &well_remark.Row_created_date, &well_remark.Row_effective_date, &well_remark.Row_expiry_date, &well_remark.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_remark to result
		result = append(result, well_remark)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellRemark(c *fiber.Ctx) error {
	var well_remark dto.Well_remark

	setDefaults(&well_remark)

	if err := json.Unmarshal(c.Body(), &well_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_remark values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	well_remark.Row_created_date = formatDate(well_remark.Row_created_date)
	well_remark.Row_changed_date = formatDate(well_remark.Row_changed_date)
	well_remark.Effective_date = formatDateString(well_remark.Effective_date)
	well_remark.Expiry_date = formatDateString(well_remark.Expiry_date)
	well_remark.Remark_date = formatDateString(well_remark.Remark_date)
	well_remark.Row_effective_date = formatDateString(well_remark.Row_effective_date)
	well_remark.Row_expiry_date = formatDateString(well_remark.Row_expiry_date)






	rows, err := stmt.Exec(well_remark.Uwi, well_remark.Source, well_remark.Remark_type, well_remark.Remark_seq_no, well_remark.Active_ind, well_remark.Effective_date, well_remark.Expiry_date, well_remark.Ppdm_guid, well_remark.Remark, well_remark.Remark_date, well_remark.Remark_strat_unit_id, well_remark.Strat_name_set_id, well_remark.Row_changed_by, well_remark.Row_changed_date, well_remark.Row_created_by, well_remark.Row_created_date, well_remark.Row_effective_date, well_remark.Row_expiry_date, well_remark.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellRemark(c *fiber.Ctx) error {
	var well_remark dto.Well_remark

	setDefaults(&well_remark)

	if err := json.Unmarshal(c.Body(), &well_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_remark.Uwi = id

    if well_remark.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_remark set source = :1, remark_type = :2, remark_seq_no = :3, active_ind = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, remark_date = :9, remark_strat_unit_id = :10, strat_name_set_id = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where uwi = :19")
	if err != nil {
		return err
	}

	well_remark.Row_changed_date = formatDate(well_remark.Row_changed_date)
	well_remark.Effective_date = formatDateString(well_remark.Effective_date)
	well_remark.Expiry_date = formatDateString(well_remark.Expiry_date)
	well_remark.Remark_date = formatDateString(well_remark.Remark_date)
	well_remark.Row_effective_date = formatDateString(well_remark.Row_effective_date)
	well_remark.Row_expiry_date = formatDateString(well_remark.Row_expiry_date)






	rows, err := stmt.Exec(well_remark.Source, well_remark.Remark_type, well_remark.Remark_seq_no, well_remark.Active_ind, well_remark.Effective_date, well_remark.Expiry_date, well_remark.Ppdm_guid, well_remark.Remark, well_remark.Remark_date, well_remark.Remark_strat_unit_id, well_remark.Strat_name_set_id, well_remark.Row_changed_by, well_remark.Row_changed_date, well_remark.Row_created_by, well_remark.Row_effective_date, well_remark.Row_expiry_date, well_remark.Row_quality, well_remark.Uwi)
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

func PatchWellRemark(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_remark set "
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
		} else if key == "effective_date" || key == "expiry_date" || key == "remark_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteWellRemark(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_remark dto.Well_remark
	well_remark.Uwi = id

	stmt, err := db.Prepare("delete from well_remark where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_remark.Uwi)
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



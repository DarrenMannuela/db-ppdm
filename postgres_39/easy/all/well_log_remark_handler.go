package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetWellLogRemark(c *fiber.Ctx) error {
	rows, err := db.Query("select * from well_log_remark")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Well_log_remark

	for rows.Next() {
		var well_log_remark dto.Well_log_remark
		if err := rows.Scan(&well_log_remark.Uwi, &well_log_remark.Well_log_id, &well_log_remark.Well_log_source, &well_log_remark.Remark_obs_no, &well_log_remark.Active_ind, &well_log_remark.Effective_date, &well_log_remark.Expiry_date, &well_log_remark.Made_by_ba_id, &well_log_remark.Ppdm_guid, &well_log_remark.Remark, &well_log_remark.Remark_type, &well_log_remark.Remark_use_type, &well_log_remark.Source, &well_log_remark.Row_changed_by, &well_log_remark.Row_changed_date, &well_log_remark.Row_created_by, &well_log_remark.Row_created_date, &well_log_remark.Row_effective_date, &well_log_remark.Row_expiry_date, &well_log_remark.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append well_log_remark to result
		result = append(result, well_log_remark)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetWellLogRemark(c *fiber.Ctx) error {
	var well_log_remark dto.Well_log_remark

	setDefaults(&well_log_remark)

	if err := json.Unmarshal(c.Body(), &well_log_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into well_log_remark values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	well_log_remark.Row_created_date = formatDate(well_log_remark.Row_created_date)
	well_log_remark.Row_changed_date = formatDate(well_log_remark.Row_changed_date)
	well_log_remark.Effective_date = formatDateString(well_log_remark.Effective_date)
	well_log_remark.Expiry_date = formatDateString(well_log_remark.Expiry_date)
	well_log_remark.Row_effective_date = formatDateString(well_log_remark.Row_effective_date)
	well_log_remark.Row_expiry_date = formatDateString(well_log_remark.Row_expiry_date)






	rows, err := stmt.Exec(well_log_remark.Uwi, well_log_remark.Well_log_id, well_log_remark.Well_log_source, well_log_remark.Remark_obs_no, well_log_remark.Active_ind, well_log_remark.Effective_date, well_log_remark.Expiry_date, well_log_remark.Made_by_ba_id, well_log_remark.Ppdm_guid, well_log_remark.Remark, well_log_remark.Remark_type, well_log_remark.Remark_use_type, well_log_remark.Source, well_log_remark.Row_changed_by, well_log_remark.Row_changed_date, well_log_remark.Row_created_by, well_log_remark.Row_created_date, well_log_remark.Row_effective_date, well_log_remark.Row_expiry_date, well_log_remark.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateWellLogRemark(c *fiber.Ctx) error {
	var well_log_remark dto.Well_log_remark

	setDefaults(&well_log_remark)

	if err := json.Unmarshal(c.Body(), &well_log_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	well_log_remark.Uwi = id

    if well_log_remark.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update well_log_remark set well_log_id = :1, well_log_source = :2, remark_obs_no = :3, active_ind = :4, effective_date = :5, expiry_date = :6, made_by_ba_id = :7, ppdm_guid = :8, remark = :9, remark_type = :10, remark_use_type = :11, source = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where uwi = :20")
	if err != nil {
		return err
	}

	well_log_remark.Row_changed_date = formatDate(well_log_remark.Row_changed_date)
	well_log_remark.Effective_date = formatDateString(well_log_remark.Effective_date)
	well_log_remark.Expiry_date = formatDateString(well_log_remark.Expiry_date)
	well_log_remark.Row_effective_date = formatDateString(well_log_remark.Row_effective_date)
	well_log_remark.Row_expiry_date = formatDateString(well_log_remark.Row_expiry_date)






	rows, err := stmt.Exec(well_log_remark.Well_log_id, well_log_remark.Well_log_source, well_log_remark.Remark_obs_no, well_log_remark.Active_ind, well_log_remark.Effective_date, well_log_remark.Expiry_date, well_log_remark.Made_by_ba_id, well_log_remark.Ppdm_guid, well_log_remark.Remark, well_log_remark.Remark_type, well_log_remark.Remark_use_type, well_log_remark.Source, well_log_remark.Row_changed_by, well_log_remark.Row_changed_date, well_log_remark.Row_created_by, well_log_remark.Row_effective_date, well_log_remark.Row_expiry_date, well_log_remark.Row_quality, well_log_remark.Uwi)
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

func PatchWellLogRemark(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update well_log_remark set "
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

func DeleteWellLogRemark(c *fiber.Ctx) error {
	id := c.Params("id")
	var well_log_remark dto.Well_log_remark
	well_log_remark.Uwi = id

	stmt, err := db.Prepare("delete from well_log_remark where uwi = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(well_log_remark.Uwi)
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



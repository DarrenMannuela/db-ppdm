package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetLithLogRemark(c *fiber.Ctx) error {
	rows, err := db.Query("select * from lith_log_remark")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Lith_log_remark

	for rows.Next() {
		var lith_log_remark dto.Lith_log_remark
		if err := rows.Scan(&lith_log_remark.Lithology_log_id, &lith_log_remark.Lith_log_source, &lith_log_remark.Remark_seq_no, &lith_log_remark.Active_ind, &lith_log_remark.Effective_date, &lith_log_remark.Expiry_date, &lith_log_remark.Ppdm_guid, &lith_log_remark.Remark, &lith_log_remark.Remark_ba_id, &lith_log_remark.Row_changed_by, &lith_log_remark.Row_changed_date, &lith_log_remark.Row_created_by, &lith_log_remark.Row_created_date, &lith_log_remark.Row_effective_date, &lith_log_remark.Row_expiry_date, &lith_log_remark.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append lith_log_remark to result
		result = append(result, lith_log_remark)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetLithLogRemark(c *fiber.Ctx) error {
	var lith_log_remark dto.Lith_log_remark

	setDefaults(&lith_log_remark)

	if err := json.Unmarshal(c.Body(), &lith_log_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into lith_log_remark values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16)")
	if err != nil {
		return err
	}
	lith_log_remark.Row_created_date = formatDate(lith_log_remark.Row_created_date)
	lith_log_remark.Row_changed_date = formatDate(lith_log_remark.Row_changed_date)
	lith_log_remark.Effective_date = formatDateString(lith_log_remark.Effective_date)
	lith_log_remark.Expiry_date = formatDateString(lith_log_remark.Expiry_date)
	lith_log_remark.Row_effective_date = formatDateString(lith_log_remark.Row_effective_date)
	lith_log_remark.Row_expiry_date = formatDateString(lith_log_remark.Row_expiry_date)






	rows, err := stmt.Exec(lith_log_remark.Lithology_log_id, lith_log_remark.Lith_log_source, lith_log_remark.Remark_seq_no, lith_log_remark.Active_ind, lith_log_remark.Effective_date, lith_log_remark.Expiry_date, lith_log_remark.Ppdm_guid, lith_log_remark.Remark, lith_log_remark.Remark_ba_id, lith_log_remark.Row_changed_by, lith_log_remark.Row_changed_date, lith_log_remark.Row_created_by, lith_log_remark.Row_created_date, lith_log_remark.Row_effective_date, lith_log_remark.Row_expiry_date, lith_log_remark.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateLithLogRemark(c *fiber.Ctx) error {
	var lith_log_remark dto.Lith_log_remark

	setDefaults(&lith_log_remark)

	if err := json.Unmarshal(c.Body(), &lith_log_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	lith_log_remark.Lithology_log_id = id

    if lith_log_remark.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update lith_log_remark set lith_log_source = :1, remark_seq_no = :2, active_ind = :3, effective_date = :4, expiry_date = :5, ppdm_guid = :6, remark = :7, remark_ba_id = :8, row_changed_by = :9, row_changed_date = :10, row_created_by = :11, row_effective_date = :12, row_expiry_date = :13, row_quality = :14 where lithology_log_id = :16")
	if err != nil {
		return err
	}

	lith_log_remark.Row_changed_date = formatDate(lith_log_remark.Row_changed_date)
	lith_log_remark.Effective_date = formatDateString(lith_log_remark.Effective_date)
	lith_log_remark.Expiry_date = formatDateString(lith_log_remark.Expiry_date)
	lith_log_remark.Row_effective_date = formatDateString(lith_log_remark.Row_effective_date)
	lith_log_remark.Row_expiry_date = formatDateString(lith_log_remark.Row_expiry_date)






	rows, err := stmt.Exec(lith_log_remark.Lith_log_source, lith_log_remark.Remark_seq_no, lith_log_remark.Active_ind, lith_log_remark.Effective_date, lith_log_remark.Expiry_date, lith_log_remark.Ppdm_guid, lith_log_remark.Remark, lith_log_remark.Remark_ba_id, lith_log_remark.Row_changed_by, lith_log_remark.Row_changed_date, lith_log_remark.Row_created_by, lith_log_remark.Row_effective_date, lith_log_remark.Row_expiry_date, lith_log_remark.Row_quality, lith_log_remark.Lithology_log_id)
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

func PatchLithLogRemark(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update lith_log_remark set "
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
	query += " where lithology_log_id = :id"

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

func DeleteLithLogRemark(c *fiber.Ctx) error {
	id := c.Params("id")
	var lith_log_remark dto.Lith_log_remark
	lith_log_remark.Lithology_log_id = id

	stmt, err := db.Prepare("delete from lith_log_remark where lithology_log_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(lith_log_remark.Lithology_log_id)
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



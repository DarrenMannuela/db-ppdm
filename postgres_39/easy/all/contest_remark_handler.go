package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetContestRemark(c *fiber.Ctx) error {
	rows, err := db.Query("select * from contest_remark")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Contest_remark

	for rows.Next() {
		var contest_remark dto.Contest_remark
		if err := rows.Scan(&contest_remark.Contest_id, &contest_remark.Remark_id, &contest_remark.Active_ind, &contest_remark.Effective_date, &contest_remark.Expiry_date, &contest_remark.Made_by, &contest_remark.Ppdm_guid, &contest_remark.Remark, &contest_remark.Remark_date, &contest_remark.Remark_seq_no, &contest_remark.Remark_type, &contest_remark.Source, &contest_remark.Row_changed_by, &contest_remark.Row_changed_date, &contest_remark.Row_created_by, &contest_remark.Row_created_date, &contest_remark.Row_effective_date, &contest_remark.Row_expiry_date, &contest_remark.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append contest_remark to result
		result = append(result, contest_remark)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetContestRemark(c *fiber.Ctx) error {
	var contest_remark dto.Contest_remark

	setDefaults(&contest_remark)

	if err := json.Unmarshal(c.Body(), &contest_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into contest_remark values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	contest_remark.Row_created_date = formatDate(contest_remark.Row_created_date)
	contest_remark.Row_changed_date = formatDate(contest_remark.Row_changed_date)
	contest_remark.Effective_date = formatDateString(contest_remark.Effective_date)
	contest_remark.Expiry_date = formatDateString(contest_remark.Expiry_date)
	contest_remark.Remark_date = formatDateString(contest_remark.Remark_date)
	contest_remark.Row_effective_date = formatDateString(contest_remark.Row_effective_date)
	contest_remark.Row_expiry_date = formatDateString(contest_remark.Row_expiry_date)






	rows, err := stmt.Exec(contest_remark.Contest_id, contest_remark.Remark_id, contest_remark.Active_ind, contest_remark.Effective_date, contest_remark.Expiry_date, contest_remark.Made_by, contest_remark.Ppdm_guid, contest_remark.Remark, contest_remark.Remark_date, contest_remark.Remark_seq_no, contest_remark.Remark_type, contest_remark.Source, contest_remark.Row_changed_by, contest_remark.Row_changed_date, contest_remark.Row_created_by, contest_remark.Row_created_date, contest_remark.Row_effective_date, contest_remark.Row_expiry_date, contest_remark.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateContestRemark(c *fiber.Ctx) error {
	var contest_remark dto.Contest_remark

	setDefaults(&contest_remark)

	if err := json.Unmarshal(c.Body(), &contest_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	contest_remark.Contest_id = id

    if contest_remark.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update contest_remark set remark_id = :1, active_ind = :2, effective_date = :3, expiry_date = :4, made_by = :5, ppdm_guid = :6, remark = :7, remark_date = :8, remark_seq_no = :9, remark_type = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where contest_id = :19")
	if err != nil {
		return err
	}

	contest_remark.Row_changed_date = formatDate(contest_remark.Row_changed_date)
	contest_remark.Effective_date = formatDateString(contest_remark.Effective_date)
	contest_remark.Expiry_date = formatDateString(contest_remark.Expiry_date)
	contest_remark.Remark_date = formatDateString(contest_remark.Remark_date)
	contest_remark.Row_effective_date = formatDateString(contest_remark.Row_effective_date)
	contest_remark.Row_expiry_date = formatDateString(contest_remark.Row_expiry_date)






	rows, err := stmt.Exec(contest_remark.Remark_id, contest_remark.Active_ind, contest_remark.Effective_date, contest_remark.Expiry_date, contest_remark.Made_by, contest_remark.Ppdm_guid, contest_remark.Remark, contest_remark.Remark_date, contest_remark.Remark_seq_no, contest_remark.Remark_type, contest_remark.Source, contest_remark.Row_changed_by, contest_remark.Row_changed_date, contest_remark.Row_created_by, contest_remark.Row_effective_date, contest_remark.Row_expiry_date, contest_remark.Row_quality, contest_remark.Contest_id)
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

func PatchContestRemark(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update contest_remark set "
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
	query += " where contest_id = :id"

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

func DeleteContestRemark(c *fiber.Ctx) error {
	id := c.Params("id")
	var contest_remark dto.Contest_remark
	contest_remark.Contest_id = id

	stmt, err := db.Prepare("delete from contest_remark where contest_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(contest_remark.Contest_id)
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



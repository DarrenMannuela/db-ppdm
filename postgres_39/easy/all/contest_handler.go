package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetContest(c *fiber.Ctx) error {
	rows, err := db.Query("select * from contest")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Contest

	for rows.Next() {
		var contest dto.Contest
		if err := rows.Scan(&contest.Contest_id, &contest.Active_ind, &contest.Effective_date, &contest.End_date, &contest.Expiry_date, &contest.Lr_contest_type, &contest.Ppdm_guid, &contest.Reason, &contest.Remark, &contest.Resolution_date, &contest.Resolution_method, &contest.Resolution_remark, &contest.Source, &contest.Start_date, &contest.Row_changed_by, &contest.Row_changed_date, &contest.Row_created_by, &contest.Row_created_date, &contest.Row_effective_date, &contest.Row_expiry_date, &contest.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append contest to result
		result = append(result, contest)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetContest(c *fiber.Ctx) error {
	var contest dto.Contest

	setDefaults(&contest)

	if err := json.Unmarshal(c.Body(), &contest); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into contest values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21)")
	if err != nil {
		return err
	}
	contest.Row_created_date = formatDate(contest.Row_created_date)
	contest.Row_changed_date = formatDate(contest.Row_changed_date)
	contest.Effective_date = formatDateString(contest.Effective_date)
	contest.End_date = formatDateString(contest.End_date)
	contest.Expiry_date = formatDateString(contest.Expiry_date)
	contest.Resolution_date = formatDateString(contest.Resolution_date)
	contest.Start_date = formatDateString(contest.Start_date)
	contest.Row_effective_date = formatDateString(contest.Row_effective_date)
	contest.Row_expiry_date = formatDateString(contest.Row_expiry_date)






	rows, err := stmt.Exec(contest.Contest_id, contest.Active_ind, contest.Effective_date, contest.End_date, contest.Expiry_date, contest.Lr_contest_type, contest.Ppdm_guid, contest.Reason, contest.Remark, contest.Resolution_date, contest.Resolution_method, contest.Resolution_remark, contest.Source, contest.Start_date, contest.Row_changed_by, contest.Row_changed_date, contest.Row_created_by, contest.Row_created_date, contest.Row_effective_date, contest.Row_expiry_date, contest.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateContest(c *fiber.Ctx) error {
	var contest dto.Contest

	setDefaults(&contest)

	if err := json.Unmarshal(c.Body(), &contest); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	contest.Contest_id = id

    if contest.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update contest set active_ind = :1, effective_date = :2, end_date = :3, expiry_date = :4, lr_contest_type = :5, ppdm_guid = :6, reason = :7, remark = :8, resolution_date = :9, resolution_method = :10, resolution_remark = :11, source = :12, start_date = :13, row_changed_by = :14, row_changed_date = :15, row_created_by = :16, row_effective_date = :17, row_expiry_date = :18, row_quality = :19 where contest_id = :21")
	if err != nil {
		return err
	}

	contest.Row_changed_date = formatDate(contest.Row_changed_date)
	contest.Effective_date = formatDateString(contest.Effective_date)
	contest.End_date = formatDateString(contest.End_date)
	contest.Expiry_date = formatDateString(contest.Expiry_date)
	contest.Resolution_date = formatDateString(contest.Resolution_date)
	contest.Start_date = formatDateString(contest.Start_date)
	contest.Row_effective_date = formatDateString(contest.Row_effective_date)
	contest.Row_expiry_date = formatDateString(contest.Row_expiry_date)






	rows, err := stmt.Exec(contest.Active_ind, contest.Effective_date, contest.End_date, contest.Expiry_date, contest.Lr_contest_type, contest.Ppdm_guid, contest.Reason, contest.Remark, contest.Resolution_date, contest.Resolution_method, contest.Resolution_remark, contest.Source, contest.Start_date, contest.Row_changed_by, contest.Row_changed_date, contest.Row_created_by, contest.Row_effective_date, contest.Row_expiry_date, contest.Row_quality, contest.Contest_id)
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

func PatchContest(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update contest set "
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
		} else if key == "effective_date" || key == "end_date" || key == "expiry_date" || key == "resolution_date" || key == "start_date" || key == "row_effective_date" || key == "row_expiry_date"      {
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

func DeleteContest(c *fiber.Ctx) error {
	id := c.Params("id")
	var contest dto.Contest
	contest.Contest_id = id

	stmt, err := db.Prepare("delete from contest where contest_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(contest.Contest_id)
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



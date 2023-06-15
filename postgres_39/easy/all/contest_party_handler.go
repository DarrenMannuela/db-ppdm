package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetContestParty(c *fiber.Ctx) error {
	rows, err := db.Query("select * from contest_party")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Contest_party

	for rows.Next() {
		var contest_party dto.Contest_party
		if err := rows.Scan(&contest_party.Contest_id, &contest_party.Business_associate_id, &contest_party.Ba_obs_no, &contest_party.Active_ind, &contest_party.Effective_date, &contest_party.Expiry_date, &contest_party.Party_role, &contest_party.Ppdm_guid, &contest_party.Remark, &contest_party.Source, &contest_party.Row_changed_by, &contest_party.Row_changed_date, &contest_party.Row_created_by, &contest_party.Row_created_date, &contest_party.Row_effective_date, &contest_party.Row_expiry_date, &contest_party.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append contest_party to result
		result = append(result, contest_party)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetContestParty(c *fiber.Ctx) error {
	var contest_party dto.Contest_party

	setDefaults(&contest_party)

	if err := json.Unmarshal(c.Body(), &contest_party); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into contest_party values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	contest_party.Row_created_date = formatDate(contest_party.Row_created_date)
	contest_party.Row_changed_date = formatDate(contest_party.Row_changed_date)
	contest_party.Effective_date = formatDateString(contest_party.Effective_date)
	contest_party.Expiry_date = formatDateString(contest_party.Expiry_date)
	contest_party.Row_effective_date = formatDateString(contest_party.Row_effective_date)
	contest_party.Row_expiry_date = formatDateString(contest_party.Row_expiry_date)






	rows, err := stmt.Exec(contest_party.Contest_id, contest_party.Business_associate_id, contest_party.Ba_obs_no, contest_party.Active_ind, contest_party.Effective_date, contest_party.Expiry_date, contest_party.Party_role, contest_party.Ppdm_guid, contest_party.Remark, contest_party.Source, contest_party.Row_changed_by, contest_party.Row_changed_date, contest_party.Row_created_by, contest_party.Row_created_date, contest_party.Row_effective_date, contest_party.Row_expiry_date, contest_party.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateContestParty(c *fiber.Ctx) error {
	var contest_party dto.Contest_party

	setDefaults(&contest_party)

	if err := json.Unmarshal(c.Body(), &contest_party); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	contest_party.Contest_id = id

    if contest_party.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update contest_party set business_associate_id = :1, ba_obs_no = :2, active_ind = :3, effective_date = :4, expiry_date = :5, party_role = :6, ppdm_guid = :7, remark = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where contest_id = :17")
	if err != nil {
		return err
	}

	contest_party.Row_changed_date = formatDate(contest_party.Row_changed_date)
	contest_party.Effective_date = formatDateString(contest_party.Effective_date)
	contest_party.Expiry_date = formatDateString(contest_party.Expiry_date)
	contest_party.Row_effective_date = formatDateString(contest_party.Row_effective_date)
	contest_party.Row_expiry_date = formatDateString(contest_party.Row_expiry_date)






	rows, err := stmt.Exec(contest_party.Business_associate_id, contest_party.Ba_obs_no, contest_party.Active_ind, contest_party.Effective_date, contest_party.Expiry_date, contest_party.Party_role, contest_party.Ppdm_guid, contest_party.Remark, contest_party.Source, contest_party.Row_changed_by, contest_party.Row_changed_date, contest_party.Row_created_by, contest_party.Row_effective_date, contest_party.Row_expiry_date, contest_party.Row_quality, contest_party.Contest_id)
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

func PatchContestParty(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update contest_party set "
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

func DeleteContestParty(c *fiber.Ctx) error {
	id := c.Params("id")
	var contest_party dto.Contest_party
	contest_party.Contest_id = id

	stmt, err := db.Prepare("delete from contest_party where contest_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(contest_party.Contest_id)
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



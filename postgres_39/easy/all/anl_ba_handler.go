package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetAnlBa(c *fiber.Ctx) error {
	rows, err := db.Query("select * from anl_ba")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Anl_ba

	for rows.Next() {
		var anl_ba dto.Anl_ba
		if err := rows.Scan(&anl_ba.Analysis_id, &anl_ba.Anl_source, &anl_ba.Ba_obs_no, &anl_ba.Active_ind, &anl_ba.Analysis_ba_id, &anl_ba.Ba_role_type, &anl_ba.Effective_date, &anl_ba.Expiry_date, &anl_ba.Ppdm_guid, &anl_ba.Problem_ind, &anl_ba.Remark, &anl_ba.Source, &anl_ba.Step_seq_no, &anl_ba.Row_changed_by, &anl_ba.Row_changed_date, &anl_ba.Row_created_by, &anl_ba.Row_created_date, &anl_ba.Row_effective_date, &anl_ba.Row_expiry_date, &anl_ba.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append anl_ba to result
		result = append(result, anl_ba)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetAnlBa(c *fiber.Ctx) error {
	var anl_ba dto.Anl_ba

	setDefaults(&anl_ba)

	if err := json.Unmarshal(c.Body(), &anl_ba); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into anl_ba values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	anl_ba.Row_created_date = formatDate(anl_ba.Row_created_date)
	anl_ba.Row_changed_date = formatDate(anl_ba.Row_changed_date)
	anl_ba.Effective_date = formatDateString(anl_ba.Effective_date)
	anl_ba.Expiry_date = formatDateString(anl_ba.Expiry_date)
	anl_ba.Row_effective_date = formatDateString(anl_ba.Row_effective_date)
	anl_ba.Row_expiry_date = formatDateString(anl_ba.Row_expiry_date)






	rows, err := stmt.Exec(anl_ba.Analysis_id, anl_ba.Anl_source, anl_ba.Ba_obs_no, anl_ba.Active_ind, anl_ba.Analysis_ba_id, anl_ba.Ba_role_type, anl_ba.Effective_date, anl_ba.Expiry_date, anl_ba.Ppdm_guid, anl_ba.Problem_ind, anl_ba.Remark, anl_ba.Source, anl_ba.Step_seq_no, anl_ba.Row_changed_by, anl_ba.Row_changed_date, anl_ba.Row_created_by, anl_ba.Row_created_date, anl_ba.Row_effective_date, anl_ba.Row_expiry_date, anl_ba.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateAnlBa(c *fiber.Ctx) error {
	var anl_ba dto.Anl_ba

	setDefaults(&anl_ba)

	if err := json.Unmarshal(c.Body(), &anl_ba); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	anl_ba.Analysis_id = id

    if anl_ba.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update anl_ba set anl_source = :1, ba_obs_no = :2, active_ind = :3, analysis_ba_id = :4, ba_role_type = :5, effective_date = :6, expiry_date = :7, ppdm_guid = :8, problem_ind = :9, remark = :10, source = :11, step_seq_no = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where analysis_id = :20")
	if err != nil {
		return err
	}

	anl_ba.Row_changed_date = formatDate(anl_ba.Row_changed_date)
	anl_ba.Effective_date = formatDateString(anl_ba.Effective_date)
	anl_ba.Expiry_date = formatDateString(anl_ba.Expiry_date)
	anl_ba.Row_effective_date = formatDateString(anl_ba.Row_effective_date)
	anl_ba.Row_expiry_date = formatDateString(anl_ba.Row_expiry_date)






	rows, err := stmt.Exec(anl_ba.Anl_source, anl_ba.Ba_obs_no, anl_ba.Active_ind, anl_ba.Analysis_ba_id, anl_ba.Ba_role_type, anl_ba.Effective_date, anl_ba.Expiry_date, anl_ba.Ppdm_guid, anl_ba.Problem_ind, anl_ba.Remark, anl_ba.Source, anl_ba.Step_seq_no, anl_ba.Row_changed_by, anl_ba.Row_changed_date, anl_ba.Row_created_by, anl_ba.Row_effective_date, anl_ba.Row_expiry_date, anl_ba.Row_quality, anl_ba.Analysis_id)
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

func PatchAnlBa(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update anl_ba set "
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

func DeleteAnlBa(c *fiber.Ctx) error {
	id := c.Params("id")
	var anl_ba dto.Anl_ba
	anl_ba.Analysis_id = id

	stmt, err := db.Prepare("delete from anl_ba where analysis_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(anl_ba.Analysis_id)
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



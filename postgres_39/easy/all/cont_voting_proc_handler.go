package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetContVotingProc(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cont_voting_proc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cont_voting_proc

	for rows.Next() {
		var cont_voting_proc dto.Cont_voting_proc
		if err := rows.Scan(&cont_voting_proc.Contract_id, &cont_voting_proc.Voting_procedure_id, &cont_voting_proc.Active_ind, &cont_voting_proc.Defeat_count, &cont_voting_proc.Defeat_percent, &cont_voting_proc.Effective_date, &cont_voting_proc.Expiry_date, &cont_voting_proc.Interest_percent, &cont_voting_proc.No_vote_response, &cont_voting_proc.Ppdm_guid, &cont_voting_proc.Quorum_count, &cont_voting_proc.Remark, &cont_voting_proc.Response_period, &cont_voting_proc.Response_period_uom, &cont_voting_proc.Source, &cont_voting_proc.Source_document_id, &cont_voting_proc.Vote_procedure_type, &cont_voting_proc.Row_changed_by, &cont_voting_proc.Row_changed_date, &cont_voting_proc.Row_created_by, &cont_voting_proc.Row_created_date, &cont_voting_proc.Row_effective_date, &cont_voting_proc.Row_expiry_date, &cont_voting_proc.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cont_voting_proc to result
		result = append(result, cont_voting_proc)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetContVotingProc(c *fiber.Ctx) error {
	var cont_voting_proc dto.Cont_voting_proc

	setDefaults(&cont_voting_proc)

	if err := json.Unmarshal(c.Body(), &cont_voting_proc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cont_voting_proc values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24)")
	if err != nil {
		return err
	}
	cont_voting_proc.Row_created_date = formatDate(cont_voting_proc.Row_created_date)
	cont_voting_proc.Row_changed_date = formatDate(cont_voting_proc.Row_changed_date)
	cont_voting_proc.Effective_date = formatDateString(cont_voting_proc.Effective_date)
	cont_voting_proc.Expiry_date = formatDateString(cont_voting_proc.Expiry_date)
	cont_voting_proc.Row_effective_date = formatDateString(cont_voting_proc.Row_effective_date)
	cont_voting_proc.Row_expiry_date = formatDateString(cont_voting_proc.Row_expiry_date)






	rows, err := stmt.Exec(cont_voting_proc.Contract_id, cont_voting_proc.Voting_procedure_id, cont_voting_proc.Active_ind, cont_voting_proc.Defeat_count, cont_voting_proc.Defeat_percent, cont_voting_proc.Effective_date, cont_voting_proc.Expiry_date, cont_voting_proc.Interest_percent, cont_voting_proc.No_vote_response, cont_voting_proc.Ppdm_guid, cont_voting_proc.Quorum_count, cont_voting_proc.Remark, cont_voting_proc.Response_period, cont_voting_proc.Response_period_uom, cont_voting_proc.Source, cont_voting_proc.Source_document_id, cont_voting_proc.Vote_procedure_type, cont_voting_proc.Row_changed_by, cont_voting_proc.Row_changed_date, cont_voting_proc.Row_created_by, cont_voting_proc.Row_created_date, cont_voting_proc.Row_effective_date, cont_voting_proc.Row_expiry_date, cont_voting_proc.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateContVotingProc(c *fiber.Ctx) error {
	var cont_voting_proc dto.Cont_voting_proc

	setDefaults(&cont_voting_proc)

	if err := json.Unmarshal(c.Body(), &cont_voting_proc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cont_voting_proc.Contract_id = id

    if cont_voting_proc.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cont_voting_proc set voting_procedure_id = :1, active_ind = :2, defeat_count = :3, defeat_percent = :4, effective_date = :5, expiry_date = :6, interest_percent = :7, no_vote_response = :8, ppdm_guid = :9, quorum_count = :10, remark = :11, response_period = :12, response_period_uom = :13, source = :14, source_document_id = :15, vote_procedure_type = :16, row_changed_by = :17, row_changed_date = :18, row_created_by = :19, row_effective_date = :20, row_expiry_date = :21, row_quality = :22 where contract_id = :24")
	if err != nil {
		return err
	}

	cont_voting_proc.Row_changed_date = formatDate(cont_voting_proc.Row_changed_date)
	cont_voting_proc.Effective_date = formatDateString(cont_voting_proc.Effective_date)
	cont_voting_proc.Expiry_date = formatDateString(cont_voting_proc.Expiry_date)
	cont_voting_proc.Row_effective_date = formatDateString(cont_voting_proc.Row_effective_date)
	cont_voting_proc.Row_expiry_date = formatDateString(cont_voting_proc.Row_expiry_date)






	rows, err := stmt.Exec(cont_voting_proc.Voting_procedure_id, cont_voting_proc.Active_ind, cont_voting_proc.Defeat_count, cont_voting_proc.Defeat_percent, cont_voting_proc.Effective_date, cont_voting_proc.Expiry_date, cont_voting_proc.Interest_percent, cont_voting_proc.No_vote_response, cont_voting_proc.Ppdm_guid, cont_voting_proc.Quorum_count, cont_voting_proc.Remark, cont_voting_proc.Response_period, cont_voting_proc.Response_period_uom, cont_voting_proc.Source, cont_voting_proc.Source_document_id, cont_voting_proc.Vote_procedure_type, cont_voting_proc.Row_changed_by, cont_voting_proc.Row_changed_date, cont_voting_proc.Row_created_by, cont_voting_proc.Row_effective_date, cont_voting_proc.Row_expiry_date, cont_voting_proc.Row_quality, cont_voting_proc.Contract_id)
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

func PatchContVotingProc(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cont_voting_proc set "
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
	query += " where contract_id = :id"

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

func DeleteContVotingProc(c *fiber.Ctx) error {
	id := c.Params("id")
	var cont_voting_proc dto.Cont_voting_proc
	cont_voting_proc.Contract_id = id

	stmt, err := db.Prepare("delete from cont_voting_proc where contract_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cont_voting_proc.Contract_id)
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



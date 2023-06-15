package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetIntSetStatus(c *fiber.Ctx) error {
	rows, err := db.Query("select * from int_set_status")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Int_set_status

	for rows.Next() {
		var int_set_status dto.Int_set_status
		if err := rows.Scan(&int_set_status.Interest_set_id, &int_set_status.Interest_set_seq_no, &int_set_status.Status_obs_no, &int_set_status.Active_ind, &int_set_status.Effective_date, &int_set_status.Effective_term, &int_set_status.Effective_term_ouom, &int_set_status.Expiry_date, &int_set_status.Ppdm_guid, &int_set_status.Reason, &int_set_status.Remark, &int_set_status.Source, &int_set_status.Status, &int_set_status.Status_type, &int_set_status.Undetermined_term_ind, &int_set_status.Row_changed_by, &int_set_status.Row_changed_date, &int_set_status.Row_created_by, &int_set_status.Row_created_date, &int_set_status.Row_effective_date, &int_set_status.Row_expiry_date, &int_set_status.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append int_set_status to result
		result = append(result, int_set_status)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetIntSetStatus(c *fiber.Ctx) error {
	var int_set_status dto.Int_set_status

	setDefaults(&int_set_status)

	if err := json.Unmarshal(c.Body(), &int_set_status); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into int_set_status values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	int_set_status.Row_created_date = formatDate(int_set_status.Row_created_date)
	int_set_status.Row_changed_date = formatDate(int_set_status.Row_changed_date)
	int_set_status.Effective_date = formatDateString(int_set_status.Effective_date)
	int_set_status.Expiry_date = formatDateString(int_set_status.Expiry_date)
	int_set_status.Row_effective_date = formatDateString(int_set_status.Row_effective_date)
	int_set_status.Row_expiry_date = formatDateString(int_set_status.Row_expiry_date)






	rows, err := stmt.Exec(int_set_status.Interest_set_id, int_set_status.Interest_set_seq_no, int_set_status.Status_obs_no, int_set_status.Active_ind, int_set_status.Effective_date, int_set_status.Effective_term, int_set_status.Effective_term_ouom, int_set_status.Expiry_date, int_set_status.Ppdm_guid, int_set_status.Reason, int_set_status.Remark, int_set_status.Source, int_set_status.Status, int_set_status.Status_type, int_set_status.Undetermined_term_ind, int_set_status.Row_changed_by, int_set_status.Row_changed_date, int_set_status.Row_created_by, int_set_status.Row_created_date, int_set_status.Row_effective_date, int_set_status.Row_expiry_date, int_set_status.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateIntSetStatus(c *fiber.Ctx) error {
	var int_set_status dto.Int_set_status

	setDefaults(&int_set_status)

	if err := json.Unmarshal(c.Body(), &int_set_status); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	int_set_status.Interest_set_id = id

    if int_set_status.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update int_set_status set interest_set_seq_no = :1, status_obs_no = :2, active_ind = :3, effective_date = :4, effective_term = :5, effective_term_ouom = :6, expiry_date = :7, ppdm_guid = :8, reason = :9, remark = :10, source = :11, status = :12, status_type = :13, undetermined_term_ind = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where interest_set_id = :22")
	if err != nil {
		return err
	}

	int_set_status.Row_changed_date = formatDate(int_set_status.Row_changed_date)
	int_set_status.Effective_date = formatDateString(int_set_status.Effective_date)
	int_set_status.Expiry_date = formatDateString(int_set_status.Expiry_date)
	int_set_status.Row_effective_date = formatDateString(int_set_status.Row_effective_date)
	int_set_status.Row_expiry_date = formatDateString(int_set_status.Row_expiry_date)






	rows, err := stmt.Exec(int_set_status.Interest_set_seq_no, int_set_status.Status_obs_no, int_set_status.Active_ind, int_set_status.Effective_date, int_set_status.Effective_term, int_set_status.Effective_term_ouom, int_set_status.Expiry_date, int_set_status.Ppdm_guid, int_set_status.Reason, int_set_status.Remark, int_set_status.Source, int_set_status.Status, int_set_status.Status_type, int_set_status.Undetermined_term_ind, int_set_status.Row_changed_by, int_set_status.Row_changed_date, int_set_status.Row_created_by, int_set_status.Row_effective_date, int_set_status.Row_expiry_date, int_set_status.Row_quality, int_set_status.Interest_set_id)
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

func PatchIntSetStatus(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update int_set_status set "
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
	query += " where interest_set_id = :id"

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

func DeleteIntSetStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var int_set_status dto.Int_set_status
	int_set_status.Interest_set_id = id

	stmt, err := db.Prepare("delete from int_set_status where interest_set_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(int_set_status.Interest_set_id)
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



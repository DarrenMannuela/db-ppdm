package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetContRemark(c *fiber.Ctx) error {
	rows, err := db.Query("select * from cont_remark")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Cont_remark

	for rows.Next() {
		var cont_remark dto.Cont_remark
		if err := rows.Scan(&cont_remark.Contract_id, &cont_remark.Remark_id, &cont_remark.Remark_seq_no, &cont_remark.Active_ind, &cont_remark.Author, &cont_remark.Effective_date, &cont_remark.Expiry_date, &cont_remark.Ppdm_guid, &cont_remark.Remark, &cont_remark.Remark_date, &cont_remark.Remark_type, &cont_remark.Source, &cont_remark.Row_changed_by, &cont_remark.Row_changed_date, &cont_remark.Row_created_by, &cont_remark.Row_created_date, &cont_remark.Row_effective_date, &cont_remark.Row_expiry_date, &cont_remark.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append cont_remark to result
		result = append(result, cont_remark)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetContRemark(c *fiber.Ctx) error {
	var cont_remark dto.Cont_remark

	setDefaults(&cont_remark)

	if err := json.Unmarshal(c.Body(), &cont_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into cont_remark values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	cont_remark.Row_created_date = formatDate(cont_remark.Row_created_date)
	cont_remark.Row_changed_date = formatDate(cont_remark.Row_changed_date)
	cont_remark.Effective_date = formatDateString(cont_remark.Effective_date)
	cont_remark.Expiry_date = formatDateString(cont_remark.Expiry_date)
	cont_remark.Remark_date = formatDateString(cont_remark.Remark_date)
	cont_remark.Row_effective_date = formatDateString(cont_remark.Row_effective_date)
	cont_remark.Row_expiry_date = formatDateString(cont_remark.Row_expiry_date)






	rows, err := stmt.Exec(cont_remark.Contract_id, cont_remark.Remark_id, cont_remark.Remark_seq_no, cont_remark.Active_ind, cont_remark.Author, cont_remark.Effective_date, cont_remark.Expiry_date, cont_remark.Ppdm_guid, cont_remark.Remark, cont_remark.Remark_date, cont_remark.Remark_type, cont_remark.Source, cont_remark.Row_changed_by, cont_remark.Row_changed_date, cont_remark.Row_created_by, cont_remark.Row_created_date, cont_remark.Row_effective_date, cont_remark.Row_expiry_date, cont_remark.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateContRemark(c *fiber.Ctx) error {
	var cont_remark dto.Cont_remark

	setDefaults(&cont_remark)

	if err := json.Unmarshal(c.Body(), &cont_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	cont_remark.Contract_id = id

    if cont_remark.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update cont_remark set remark_id = :1, remark_seq_no = :2, active_ind = :3, author = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, remark_date = :9, remark_type = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where contract_id = :19")
	if err != nil {
		return err
	}

	cont_remark.Row_changed_date = formatDate(cont_remark.Row_changed_date)
	cont_remark.Effective_date = formatDateString(cont_remark.Effective_date)
	cont_remark.Expiry_date = formatDateString(cont_remark.Expiry_date)
	cont_remark.Remark_date = formatDateString(cont_remark.Remark_date)
	cont_remark.Row_effective_date = formatDateString(cont_remark.Row_effective_date)
	cont_remark.Row_expiry_date = formatDateString(cont_remark.Row_expiry_date)






	rows, err := stmt.Exec(cont_remark.Remark_id, cont_remark.Remark_seq_no, cont_remark.Active_ind, cont_remark.Author, cont_remark.Effective_date, cont_remark.Expiry_date, cont_remark.Ppdm_guid, cont_remark.Remark, cont_remark.Remark_date, cont_remark.Remark_type, cont_remark.Source, cont_remark.Row_changed_by, cont_remark.Row_changed_date, cont_remark.Row_created_by, cont_remark.Row_effective_date, cont_remark.Row_expiry_date, cont_remark.Row_quality, cont_remark.Contract_id)
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

func PatchContRemark(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update cont_remark set "
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

func DeleteContRemark(c *fiber.Ctx) error {
	id := c.Params("id")
	var cont_remark dto.Cont_remark
	cont_remark.Contract_id = id

	stmt, err := db.Prepare("delete from cont_remark where contract_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(cont_remark.Contract_id)
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



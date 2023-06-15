package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetObligXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from oblig_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Oblig_xref

	for rows.Next() {
		var oblig_xref dto.Oblig_xref
		if err := rows.Scan(&oblig_xref.Obligation_id, &oblig_xref.Obligation_seq_no, &oblig_xref.Obligation_id_2, &oblig_xref.Obligation_seq_no_2, &oblig_xref.Xref_seq_no, &oblig_xref.Active_ind, &oblig_xref.Effective_date, &oblig_xref.Expiry_date, &oblig_xref.Oblig_xref_type, &oblig_xref.Ppdm_guid, &oblig_xref.Remark, &oblig_xref.Source, &oblig_xref.Row_changed_by, &oblig_xref.Row_changed_date, &oblig_xref.Row_created_by, &oblig_xref.Row_created_date, &oblig_xref.Row_effective_date, &oblig_xref.Row_expiry_date, &oblig_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append oblig_xref to result
		result = append(result, oblig_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetObligXref(c *fiber.Ctx) error {
	var oblig_xref dto.Oblig_xref

	setDefaults(&oblig_xref)

	if err := json.Unmarshal(c.Body(), &oblig_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into oblig_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	oblig_xref.Row_created_date = formatDate(oblig_xref.Row_created_date)
	oblig_xref.Row_changed_date = formatDate(oblig_xref.Row_changed_date)
	oblig_xref.Effective_date = formatDateString(oblig_xref.Effective_date)
	oblig_xref.Expiry_date = formatDateString(oblig_xref.Expiry_date)
	oblig_xref.Row_effective_date = formatDateString(oblig_xref.Row_effective_date)
	oblig_xref.Row_expiry_date = formatDateString(oblig_xref.Row_expiry_date)






	rows, err := stmt.Exec(oblig_xref.Obligation_id, oblig_xref.Obligation_seq_no, oblig_xref.Obligation_id_2, oblig_xref.Obligation_seq_no_2, oblig_xref.Xref_seq_no, oblig_xref.Active_ind, oblig_xref.Effective_date, oblig_xref.Expiry_date, oblig_xref.Oblig_xref_type, oblig_xref.Ppdm_guid, oblig_xref.Remark, oblig_xref.Source, oblig_xref.Row_changed_by, oblig_xref.Row_changed_date, oblig_xref.Row_created_by, oblig_xref.Row_created_date, oblig_xref.Row_effective_date, oblig_xref.Row_expiry_date, oblig_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateObligXref(c *fiber.Ctx) error {
	var oblig_xref dto.Oblig_xref

	setDefaults(&oblig_xref)

	if err := json.Unmarshal(c.Body(), &oblig_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	oblig_xref.Obligation_id = id

    if oblig_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update oblig_xref set obligation_seq_no = :1, obligation_id_2 = :2, obligation_seq_no_2 = :3, xref_seq_no = :4, active_ind = :5, effective_date = :6, expiry_date = :7, oblig_xref_type = :8, ppdm_guid = :9, remark = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where obligation_id = :19")
	if err != nil {
		return err
	}

	oblig_xref.Row_changed_date = formatDate(oblig_xref.Row_changed_date)
	oblig_xref.Effective_date = formatDateString(oblig_xref.Effective_date)
	oblig_xref.Expiry_date = formatDateString(oblig_xref.Expiry_date)
	oblig_xref.Row_effective_date = formatDateString(oblig_xref.Row_effective_date)
	oblig_xref.Row_expiry_date = formatDateString(oblig_xref.Row_expiry_date)






	rows, err := stmt.Exec(oblig_xref.Obligation_seq_no, oblig_xref.Obligation_id_2, oblig_xref.Obligation_seq_no_2, oblig_xref.Xref_seq_no, oblig_xref.Active_ind, oblig_xref.Effective_date, oblig_xref.Expiry_date, oblig_xref.Oblig_xref_type, oblig_xref.Ppdm_guid, oblig_xref.Remark, oblig_xref.Source, oblig_xref.Row_changed_by, oblig_xref.Row_changed_date, oblig_xref.Row_created_by, oblig_xref.Row_effective_date, oblig_xref.Row_expiry_date, oblig_xref.Row_quality, oblig_xref.Obligation_id)
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

func PatchObligXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update oblig_xref set "
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
	query += " where obligation_id = :id"

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

func DeleteObligXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var oblig_xref dto.Oblig_xref
	oblig_xref.Obligation_id = id

	stmt, err := db.Prepare("delete from oblig_xref where obligation_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(oblig_xref.Obligation_id)
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



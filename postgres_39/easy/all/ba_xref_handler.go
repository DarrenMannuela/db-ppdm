package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetBaXref(c *fiber.Ctx) error {
	rows, err := db.Query("select * from ba_xref")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Ba_xref

	for rows.Next() {
		var ba_xref dto.Ba_xref
		if err := rows.Scan(&ba_xref.Business_associate_id, &ba_xref.New_ba_id, &ba_xref.Ba_xref_seq_no, &ba_xref.Active_ind, &ba_xref.Ba_xref_type, &ba_xref.Effective_date, &ba_xref.Expiry_date, &ba_xref.Ppdm_guid, &ba_xref.Remark, &ba_xref.Source, &ba_xref.Row_changed_by, &ba_xref.Row_changed_date, &ba_xref.Row_created_by, &ba_xref.Row_created_date, &ba_xref.Row_effective_date, &ba_xref.Row_expiry_date, &ba_xref.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append ba_xref to result
		result = append(result, ba_xref)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetBaXref(c *fiber.Ctx) error {
	var ba_xref dto.Ba_xref

	setDefaults(&ba_xref)

	if err := json.Unmarshal(c.Body(), &ba_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into ba_xref values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17)")
	if err != nil {
		return err
	}
	ba_xref.Row_created_date = formatDate(ba_xref.Row_created_date)
	ba_xref.Row_changed_date = formatDate(ba_xref.Row_changed_date)
	ba_xref.Effective_date = formatDateString(ba_xref.Effective_date)
	ba_xref.Expiry_date = formatDateString(ba_xref.Expiry_date)
	ba_xref.Row_effective_date = formatDateString(ba_xref.Row_effective_date)
	ba_xref.Row_expiry_date = formatDateString(ba_xref.Row_expiry_date)






	rows, err := stmt.Exec(ba_xref.Business_associate_id, ba_xref.New_ba_id, ba_xref.Ba_xref_seq_no, ba_xref.Active_ind, ba_xref.Ba_xref_type, ba_xref.Effective_date, ba_xref.Expiry_date, ba_xref.Ppdm_guid, ba_xref.Remark, ba_xref.Source, ba_xref.Row_changed_by, ba_xref.Row_changed_date, ba_xref.Row_created_by, ba_xref.Row_created_date, ba_xref.Row_effective_date, ba_xref.Row_expiry_date, ba_xref.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateBaXref(c *fiber.Ctx) error {
	var ba_xref dto.Ba_xref

	setDefaults(&ba_xref)

	if err := json.Unmarshal(c.Body(), &ba_xref); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	ba_xref.Business_associate_id = id

    if ba_xref.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update ba_xref set new_ba_id = :1, ba_xref_seq_no = :2, active_ind = :3, ba_xref_type = :4, effective_date = :5, expiry_date = :6, ppdm_guid = :7, remark = :8, source = :9, row_changed_by = :10, row_changed_date = :11, row_created_by = :12, row_effective_date = :13, row_expiry_date = :14, row_quality = :15 where business_associate_id = :17")
	if err != nil {
		return err
	}

	ba_xref.Row_changed_date = formatDate(ba_xref.Row_changed_date)
	ba_xref.Effective_date = formatDateString(ba_xref.Effective_date)
	ba_xref.Expiry_date = formatDateString(ba_xref.Expiry_date)
	ba_xref.Row_effective_date = formatDateString(ba_xref.Row_effective_date)
	ba_xref.Row_expiry_date = formatDateString(ba_xref.Row_expiry_date)






	rows, err := stmt.Exec(ba_xref.New_ba_id, ba_xref.Ba_xref_seq_no, ba_xref.Active_ind, ba_xref.Ba_xref_type, ba_xref.Effective_date, ba_xref.Expiry_date, ba_xref.Ppdm_guid, ba_xref.Remark, ba_xref.Source, ba_xref.Row_changed_by, ba_xref.Row_changed_date, ba_xref.Row_created_by, ba_xref.Row_effective_date, ba_xref.Row_expiry_date, ba_xref.Row_quality, ba_xref.Business_associate_id)
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

func PatchBaXref(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update ba_xref set "
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
	query += " where business_associate_id = :id"

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

func DeleteBaXref(c *fiber.Ctx) error {
	id := c.Params("id")
	var ba_xref dto.Ba_xref
	ba_xref.Business_associate_id = id

	stmt, err := db.Prepare("delete from ba_xref where business_associate_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(ba_xref.Business_associate_id)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetObligRemark(c *fiber.Ctx) error {
	rows, err := db.Query("select * from oblig_remark")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Oblig_remark

	for rows.Next() {
		var oblig_remark dto.Oblig_remark
		if err := rows.Scan(&oblig_remark.Obligation_id, &oblig_remark.Obligation_seq_no, &oblig_remark.Remark_seq_no, &oblig_remark.Active_ind, &oblig_remark.Effective_date, &oblig_remark.Expiry_date, &oblig_remark.Made_by_ba_id, &oblig_remark.Ppdm_guid, &oblig_remark.Remark, &oblig_remark.Remark_type, &oblig_remark.Source, &oblig_remark.Row_changed_by, &oblig_remark.Row_changed_date, &oblig_remark.Row_created_by, &oblig_remark.Row_created_date, &oblig_remark.Row_effective_date, &oblig_remark.Row_expiry_date, &oblig_remark.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append oblig_remark to result
		result = append(result, oblig_remark)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetObligRemark(c *fiber.Ctx) error {
	var oblig_remark dto.Oblig_remark

	setDefaults(&oblig_remark)

	if err := json.Unmarshal(c.Body(), &oblig_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into oblig_remark values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18)")
	if err != nil {
		return err
	}
	oblig_remark.Row_created_date = formatDate(oblig_remark.Row_created_date)
	oblig_remark.Row_changed_date = formatDate(oblig_remark.Row_changed_date)
	oblig_remark.Effective_date = formatDateString(oblig_remark.Effective_date)
	oblig_remark.Expiry_date = formatDateString(oblig_remark.Expiry_date)
	oblig_remark.Row_effective_date = formatDateString(oblig_remark.Row_effective_date)
	oblig_remark.Row_expiry_date = formatDateString(oblig_remark.Row_expiry_date)






	rows, err := stmt.Exec(oblig_remark.Obligation_id, oblig_remark.Obligation_seq_no, oblig_remark.Remark_seq_no, oblig_remark.Active_ind, oblig_remark.Effective_date, oblig_remark.Expiry_date, oblig_remark.Made_by_ba_id, oblig_remark.Ppdm_guid, oblig_remark.Remark, oblig_remark.Remark_type, oblig_remark.Source, oblig_remark.Row_changed_by, oblig_remark.Row_changed_date, oblig_remark.Row_created_by, oblig_remark.Row_created_date, oblig_remark.Row_effective_date, oblig_remark.Row_expiry_date, oblig_remark.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateObligRemark(c *fiber.Ctx) error {
	var oblig_remark dto.Oblig_remark

	setDefaults(&oblig_remark)

	if err := json.Unmarshal(c.Body(), &oblig_remark); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	oblig_remark.Obligation_id = id

    if oblig_remark.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update oblig_remark set obligation_seq_no = :1, remark_seq_no = :2, active_ind = :3, effective_date = :4, expiry_date = :5, made_by_ba_id = :6, ppdm_guid = :7, remark = :8, remark_type = :9, source = :10, row_changed_by = :11, row_changed_date = :12, row_created_by = :13, row_effective_date = :14, row_expiry_date = :15, row_quality = :16 where obligation_id = :18")
	if err != nil {
		return err
	}

	oblig_remark.Row_changed_date = formatDate(oblig_remark.Row_changed_date)
	oblig_remark.Effective_date = formatDateString(oblig_remark.Effective_date)
	oblig_remark.Expiry_date = formatDateString(oblig_remark.Expiry_date)
	oblig_remark.Row_effective_date = formatDateString(oblig_remark.Row_effective_date)
	oblig_remark.Row_expiry_date = formatDateString(oblig_remark.Row_expiry_date)






	rows, err := stmt.Exec(oblig_remark.Obligation_seq_no, oblig_remark.Remark_seq_no, oblig_remark.Active_ind, oblig_remark.Effective_date, oblig_remark.Expiry_date, oblig_remark.Made_by_ba_id, oblig_remark.Ppdm_guid, oblig_remark.Remark, oblig_remark.Remark_type, oblig_remark.Source, oblig_remark.Row_changed_by, oblig_remark.Row_changed_date, oblig_remark.Row_created_by, oblig_remark.Row_effective_date, oblig_remark.Row_expiry_date, oblig_remark.Row_quality, oblig_remark.Obligation_id)
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

func PatchObligRemark(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update oblig_remark set "
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

func DeleteObligRemark(c *fiber.Ctx) error {
	id := c.Params("id")
	var oblig_remark dto.Oblig_remark
	oblig_remark.Obligation_id = id

	stmt, err := db.Prepare("delete from oblig_remark where obligation_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(oblig_remark.Obligation_id)
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



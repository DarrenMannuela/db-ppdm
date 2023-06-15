package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetObligSubstance(c *fiber.Ctx) error {
	rows, err := db.Query("select * from oblig_substance")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Oblig_substance

	for rows.Next() {
		var oblig_substance dto.Oblig_substance
		if err := rows.Scan(&oblig_substance.Obligation_id, &oblig_substance.Obligation_seq_no, &oblig_substance.Substance_id, &oblig_substance.Substance_seq_no, &oblig_substance.Active_ind, &oblig_substance.Contract_id, &oblig_substance.Effective_date, &oblig_substance.Excluded_ind, &oblig_substance.Expiry_date, &oblig_substance.Included_ind, &oblig_substance.Ppdm_guid, &oblig_substance.Remark, &oblig_substance.Source, &oblig_substance.Row_changed_by, &oblig_substance.Row_changed_date, &oblig_substance.Row_created_by, &oblig_substance.Row_created_date, &oblig_substance.Row_effective_date, &oblig_substance.Row_expiry_date, &oblig_substance.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append oblig_substance to result
		result = append(result, oblig_substance)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetObligSubstance(c *fiber.Ctx) error {
	var oblig_substance dto.Oblig_substance

	setDefaults(&oblig_substance)

	if err := json.Unmarshal(c.Body(), &oblig_substance); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into oblig_substance values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20)")
	if err != nil {
		return err
	}
	oblig_substance.Row_created_date = formatDate(oblig_substance.Row_created_date)
	oblig_substance.Row_changed_date = formatDate(oblig_substance.Row_changed_date)
	oblig_substance.Effective_date = formatDateString(oblig_substance.Effective_date)
	oblig_substance.Expiry_date = formatDateString(oblig_substance.Expiry_date)
	oblig_substance.Row_effective_date = formatDateString(oblig_substance.Row_effective_date)
	oblig_substance.Row_expiry_date = formatDateString(oblig_substance.Row_expiry_date)






	rows, err := stmt.Exec(oblig_substance.Obligation_id, oblig_substance.Obligation_seq_no, oblig_substance.Substance_id, oblig_substance.Substance_seq_no, oblig_substance.Active_ind, oblig_substance.Contract_id, oblig_substance.Effective_date, oblig_substance.Excluded_ind, oblig_substance.Expiry_date, oblig_substance.Included_ind, oblig_substance.Ppdm_guid, oblig_substance.Remark, oblig_substance.Source, oblig_substance.Row_changed_by, oblig_substance.Row_changed_date, oblig_substance.Row_created_by, oblig_substance.Row_created_date, oblig_substance.Row_effective_date, oblig_substance.Row_expiry_date, oblig_substance.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateObligSubstance(c *fiber.Ctx) error {
	var oblig_substance dto.Oblig_substance

	setDefaults(&oblig_substance)

	if err := json.Unmarshal(c.Body(), &oblig_substance); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	oblig_substance.Obligation_id = id

    if oblig_substance.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update oblig_substance set obligation_seq_no = :1, substance_id = :2, substance_seq_no = :3, active_ind = :4, contract_id = :5, effective_date = :6, excluded_ind = :7, expiry_date = :8, included_ind = :9, ppdm_guid = :10, remark = :11, source = :12, row_changed_by = :13, row_changed_date = :14, row_created_by = :15, row_effective_date = :16, row_expiry_date = :17, row_quality = :18 where obligation_id = :20")
	if err != nil {
		return err
	}

	oblig_substance.Row_changed_date = formatDate(oblig_substance.Row_changed_date)
	oblig_substance.Effective_date = formatDateString(oblig_substance.Effective_date)
	oblig_substance.Expiry_date = formatDateString(oblig_substance.Expiry_date)
	oblig_substance.Row_effective_date = formatDateString(oblig_substance.Row_effective_date)
	oblig_substance.Row_expiry_date = formatDateString(oblig_substance.Row_expiry_date)






	rows, err := stmt.Exec(oblig_substance.Obligation_seq_no, oblig_substance.Substance_id, oblig_substance.Substance_seq_no, oblig_substance.Active_ind, oblig_substance.Contract_id, oblig_substance.Effective_date, oblig_substance.Excluded_ind, oblig_substance.Expiry_date, oblig_substance.Included_ind, oblig_substance.Ppdm_guid, oblig_substance.Remark, oblig_substance.Source, oblig_substance.Row_changed_by, oblig_substance.Row_changed_date, oblig_substance.Row_created_by, oblig_substance.Row_effective_date, oblig_substance.Row_expiry_date, oblig_substance.Row_quality, oblig_substance.Obligation_id)
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

func PatchObligSubstance(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update oblig_substance set "
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

func DeleteObligSubstance(c *fiber.Ctx) error {
	id := c.Params("id")
	var oblig_substance dto.Oblig_substance
	oblig_substance.Obligation_id = id

	stmt, err := db.Prepare("delete from oblig_substance where obligation_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(oblig_substance.Obligation_id)
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



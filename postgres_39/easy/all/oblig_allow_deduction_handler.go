package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetObligAllowDeduction(c *fiber.Ctx) error {
	rows, err := db.Query("select * from oblig_allow_deduction")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Oblig_allow_deduction

	for rows.Next() {
		var oblig_allow_deduction dto.Oblig_allow_deduction
		if err := rows.Scan(&oblig_allow_deduction.Obligation_id, &oblig_allow_deduction.Obligation_seq_no, &oblig_allow_deduction.Allow_deduction_id, &oblig_allow_deduction.Active_ind, &oblig_allow_deduction.Allow_deduction_type, &oblig_allow_deduction.Allow_deduct_remark, &oblig_allow_deduction.Contract_id, &oblig_allow_deduction.Effective_date, &oblig_allow_deduction.Expiry_date, &oblig_allow_deduction.Ppdm_guid, &oblig_allow_deduction.Remark, &oblig_allow_deduction.Source, &oblig_allow_deduction.Row_changed_by, &oblig_allow_deduction.Row_changed_date, &oblig_allow_deduction.Row_created_by, &oblig_allow_deduction.Row_created_date, &oblig_allow_deduction.Row_effective_date, &oblig_allow_deduction.Row_expiry_date, &oblig_allow_deduction.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append oblig_allow_deduction to result
		result = append(result, oblig_allow_deduction)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetObligAllowDeduction(c *fiber.Ctx) error {
	var oblig_allow_deduction dto.Oblig_allow_deduction

	setDefaults(&oblig_allow_deduction)

	if err := json.Unmarshal(c.Body(), &oblig_allow_deduction); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into oblig_allow_deduction values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	oblig_allow_deduction.Row_created_date = formatDate(oblig_allow_deduction.Row_created_date)
	oblig_allow_deduction.Row_changed_date = formatDate(oblig_allow_deduction.Row_changed_date)
	oblig_allow_deduction.Effective_date = formatDateString(oblig_allow_deduction.Effective_date)
	oblig_allow_deduction.Expiry_date = formatDateString(oblig_allow_deduction.Expiry_date)
	oblig_allow_deduction.Row_effective_date = formatDateString(oblig_allow_deduction.Row_effective_date)
	oblig_allow_deduction.Row_expiry_date = formatDateString(oblig_allow_deduction.Row_expiry_date)






	rows, err := stmt.Exec(oblig_allow_deduction.Obligation_id, oblig_allow_deduction.Obligation_seq_no, oblig_allow_deduction.Allow_deduction_id, oblig_allow_deduction.Active_ind, oblig_allow_deduction.Allow_deduction_type, oblig_allow_deduction.Allow_deduct_remark, oblig_allow_deduction.Contract_id, oblig_allow_deduction.Effective_date, oblig_allow_deduction.Expiry_date, oblig_allow_deduction.Ppdm_guid, oblig_allow_deduction.Remark, oblig_allow_deduction.Source, oblig_allow_deduction.Row_changed_by, oblig_allow_deduction.Row_changed_date, oblig_allow_deduction.Row_created_by, oblig_allow_deduction.Row_created_date, oblig_allow_deduction.Row_effective_date, oblig_allow_deduction.Row_expiry_date, oblig_allow_deduction.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateObligAllowDeduction(c *fiber.Ctx) error {
	var oblig_allow_deduction dto.Oblig_allow_deduction

	setDefaults(&oblig_allow_deduction)

	if err := json.Unmarshal(c.Body(), &oblig_allow_deduction); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	oblig_allow_deduction.Obligation_id = id

    if oblig_allow_deduction.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update oblig_allow_deduction set obligation_seq_no = :1, allow_deduction_id = :2, active_ind = :3, allow_deduction_type = :4, allow_deduct_remark = :5, contract_id = :6, effective_date = :7, expiry_date = :8, ppdm_guid = :9, remark = :10, source = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where obligation_id = :19")
	if err != nil {
		return err
	}

	oblig_allow_deduction.Row_changed_date = formatDate(oblig_allow_deduction.Row_changed_date)
	oblig_allow_deduction.Effective_date = formatDateString(oblig_allow_deduction.Effective_date)
	oblig_allow_deduction.Expiry_date = formatDateString(oblig_allow_deduction.Expiry_date)
	oblig_allow_deduction.Row_effective_date = formatDateString(oblig_allow_deduction.Row_effective_date)
	oblig_allow_deduction.Row_expiry_date = formatDateString(oblig_allow_deduction.Row_expiry_date)






	rows, err := stmt.Exec(oblig_allow_deduction.Obligation_seq_no, oblig_allow_deduction.Allow_deduction_id, oblig_allow_deduction.Active_ind, oblig_allow_deduction.Allow_deduction_type, oblig_allow_deduction.Allow_deduct_remark, oblig_allow_deduction.Contract_id, oblig_allow_deduction.Effective_date, oblig_allow_deduction.Expiry_date, oblig_allow_deduction.Ppdm_guid, oblig_allow_deduction.Remark, oblig_allow_deduction.Source, oblig_allow_deduction.Row_changed_by, oblig_allow_deduction.Row_changed_date, oblig_allow_deduction.Row_created_by, oblig_allow_deduction.Row_effective_date, oblig_allow_deduction.Row_expiry_date, oblig_allow_deduction.Row_quality, oblig_allow_deduction.Obligation_id)
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

func PatchObligAllowDeduction(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update oblig_allow_deduction set "
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

func DeleteObligAllowDeduction(c *fiber.Ctx) error {
	id := c.Params("id")
	var oblig_allow_deduction dto.Oblig_allow_deduction
	oblig_allow_deduction.Obligation_id = id

	stmt, err := db.Prepare("delete from oblig_allow_deduction where obligation_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(oblig_allow_deduction.Obligation_id)
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



package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetReserveClassFormula(c *fiber.Ctx) error {
	rows, err := db.Query("select * from reserve_class_formula")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Reserve_class_formula

	for rows.Next() {
		var reserve_class_formula dto.Reserve_class_formula
		if err := rows.Scan(&reserve_class_formula.Reserve_class_id, &reserve_class_formula.Formula_id, &reserve_class_formula.Active_ind, &reserve_class_formula.Effective_date, &reserve_class_formula.Expiry_date, &reserve_class_formula.Formula, &reserve_class_formula.Ppdm_guid, &reserve_class_formula.Remark, &reserve_class_formula.Source, &reserve_class_formula.Use_rule, &reserve_class_formula.Use_seq_no, &reserve_class_formula.Use_type, &reserve_class_formula.Row_changed_by, &reserve_class_formula.Row_changed_date, &reserve_class_formula.Row_created_by, &reserve_class_formula.Row_created_date, &reserve_class_formula.Row_effective_date, &reserve_class_formula.Row_expiry_date, &reserve_class_formula.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append reserve_class_formula to result
		result = append(result, reserve_class_formula)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetReserveClassFormula(c *fiber.Ctx) error {
	var reserve_class_formula dto.Reserve_class_formula

	setDefaults(&reserve_class_formula)

	if err := json.Unmarshal(c.Body(), &reserve_class_formula); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into reserve_class_formula values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19)")
	if err != nil {
		return err
	}
	reserve_class_formula.Row_created_date = formatDate(reserve_class_formula.Row_created_date)
	reserve_class_formula.Row_changed_date = formatDate(reserve_class_formula.Row_changed_date)
	reserve_class_formula.Effective_date = formatDateString(reserve_class_formula.Effective_date)
	reserve_class_formula.Expiry_date = formatDateString(reserve_class_formula.Expiry_date)
	reserve_class_formula.Row_effective_date = formatDateString(reserve_class_formula.Row_effective_date)
	reserve_class_formula.Row_expiry_date = formatDateString(reserve_class_formula.Row_expiry_date)






	rows, err := stmt.Exec(reserve_class_formula.Reserve_class_id, reserve_class_formula.Formula_id, reserve_class_formula.Active_ind, reserve_class_formula.Effective_date, reserve_class_formula.Expiry_date, reserve_class_formula.Formula, reserve_class_formula.Ppdm_guid, reserve_class_formula.Remark, reserve_class_formula.Source, reserve_class_formula.Use_rule, reserve_class_formula.Use_seq_no, reserve_class_formula.Use_type, reserve_class_formula.Row_changed_by, reserve_class_formula.Row_changed_date, reserve_class_formula.Row_created_by, reserve_class_formula.Row_created_date, reserve_class_formula.Row_effective_date, reserve_class_formula.Row_expiry_date, reserve_class_formula.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateReserveClassFormula(c *fiber.Ctx) error {
	var reserve_class_formula dto.Reserve_class_formula

	setDefaults(&reserve_class_formula)

	if err := json.Unmarshal(c.Body(), &reserve_class_formula); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	reserve_class_formula.Reserve_class_id = id

    if reserve_class_formula.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update reserve_class_formula set formula_id = :1, active_ind = :2, effective_date = :3, expiry_date = :4, formula = :5, ppdm_guid = :6, remark = :7, source = :8, use_rule = :9, use_seq_no = :10, use_type = :11, row_changed_by = :12, row_changed_date = :13, row_created_by = :14, row_effective_date = :15, row_expiry_date = :16, row_quality = :17 where reserve_class_id = :19")
	if err != nil {
		return err
	}

	reserve_class_formula.Row_changed_date = formatDate(reserve_class_formula.Row_changed_date)
	reserve_class_formula.Effective_date = formatDateString(reserve_class_formula.Effective_date)
	reserve_class_formula.Expiry_date = formatDateString(reserve_class_formula.Expiry_date)
	reserve_class_formula.Row_effective_date = formatDateString(reserve_class_formula.Row_effective_date)
	reserve_class_formula.Row_expiry_date = formatDateString(reserve_class_formula.Row_expiry_date)






	rows, err := stmt.Exec(reserve_class_formula.Formula_id, reserve_class_formula.Active_ind, reserve_class_formula.Effective_date, reserve_class_formula.Expiry_date, reserve_class_formula.Formula, reserve_class_formula.Ppdm_guid, reserve_class_formula.Remark, reserve_class_formula.Source, reserve_class_formula.Use_rule, reserve_class_formula.Use_seq_no, reserve_class_formula.Use_type, reserve_class_formula.Row_changed_by, reserve_class_formula.Row_changed_date, reserve_class_formula.Row_created_by, reserve_class_formula.Row_effective_date, reserve_class_formula.Row_expiry_date, reserve_class_formula.Row_quality, reserve_class_formula.Reserve_class_id)
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

func PatchReserveClassFormula(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update reserve_class_formula set "
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
	query += " where reserve_class_id = :id"

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

func DeleteReserveClassFormula(c *fiber.Ctx) error {
	id := c.Params("id")
	var reserve_class_formula dto.Reserve_class_formula
	reserve_class_formula.Reserve_class_id = id

	stmt, err := db.Prepare("delete from reserve_class_formula where reserve_class_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(reserve_class_formula.Reserve_class_id)
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



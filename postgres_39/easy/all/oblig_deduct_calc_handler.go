package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetObligDeductCalc(c *fiber.Ctx) error {
	rows, err := db.Query("select * from oblig_deduct_calc")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Oblig_deduct_calc

	for rows.Next() {
		var oblig_deduct_calc dto.Oblig_deduct_calc
		if err := rows.Scan(&oblig_deduct_calc.Obligation_id, &oblig_deduct_calc.Obligation_seq_no, &oblig_deduct_calc.Deduction_id, &oblig_deduct_calc.Calculation_id, &oblig_deduct_calc.Active_ind, &oblig_deduct_calc.Calculation_formula, &oblig_deduct_calc.Calculation_method, &oblig_deduct_calc.Calculation_point, &oblig_deduct_calc.Contract_id, &oblig_deduct_calc.Effective_date, &oblig_deduct_calc.Expiry_date, &oblig_deduct_calc.Ppdm_guid, &oblig_deduct_calc.Remark, &oblig_deduct_calc.Source, &oblig_deduct_calc.Substance_id, &oblig_deduct_calc.Row_changed_by, &oblig_deduct_calc.Row_changed_date, &oblig_deduct_calc.Row_created_by, &oblig_deduct_calc.Row_created_date, &oblig_deduct_calc.Row_effective_date, &oblig_deduct_calc.Row_expiry_date, &oblig_deduct_calc.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append oblig_deduct_calc to result
		result = append(result, oblig_deduct_calc)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetObligDeductCalc(c *fiber.Ctx) error {
	var oblig_deduct_calc dto.Oblig_deduct_calc

	setDefaults(&oblig_deduct_calc)

	if err := json.Unmarshal(c.Body(), &oblig_deduct_calc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into oblig_deduct_calc values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22)")
	if err != nil {
		return err
	}
	oblig_deduct_calc.Row_created_date = formatDate(oblig_deduct_calc.Row_created_date)
	oblig_deduct_calc.Row_changed_date = formatDate(oblig_deduct_calc.Row_changed_date)
	oblig_deduct_calc.Effective_date = formatDateString(oblig_deduct_calc.Effective_date)
	oblig_deduct_calc.Expiry_date = formatDateString(oblig_deduct_calc.Expiry_date)
	oblig_deduct_calc.Row_effective_date = formatDateString(oblig_deduct_calc.Row_effective_date)
	oblig_deduct_calc.Row_expiry_date = formatDateString(oblig_deduct_calc.Row_expiry_date)






	rows, err := stmt.Exec(oblig_deduct_calc.Obligation_id, oblig_deduct_calc.Obligation_seq_no, oblig_deduct_calc.Deduction_id, oblig_deduct_calc.Calculation_id, oblig_deduct_calc.Active_ind, oblig_deduct_calc.Calculation_formula, oblig_deduct_calc.Calculation_method, oblig_deduct_calc.Calculation_point, oblig_deduct_calc.Contract_id, oblig_deduct_calc.Effective_date, oblig_deduct_calc.Expiry_date, oblig_deduct_calc.Ppdm_guid, oblig_deduct_calc.Remark, oblig_deduct_calc.Source, oblig_deduct_calc.Substance_id, oblig_deduct_calc.Row_changed_by, oblig_deduct_calc.Row_changed_date, oblig_deduct_calc.Row_created_by, oblig_deduct_calc.Row_created_date, oblig_deduct_calc.Row_effective_date, oblig_deduct_calc.Row_expiry_date, oblig_deduct_calc.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdateObligDeductCalc(c *fiber.Ctx) error {
	var oblig_deduct_calc dto.Oblig_deduct_calc

	setDefaults(&oblig_deduct_calc)

	if err := json.Unmarshal(c.Body(), &oblig_deduct_calc); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	oblig_deduct_calc.Obligation_id = id

    if oblig_deduct_calc.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update oblig_deduct_calc set obligation_seq_no = :1, deduction_id = :2, calculation_id = :3, active_ind = :4, calculation_formula = :5, calculation_method = :6, calculation_point = :7, contract_id = :8, effective_date = :9, expiry_date = :10, ppdm_guid = :11, remark = :12, source = :13, substance_id = :14, row_changed_by = :15, row_changed_date = :16, row_created_by = :17, row_effective_date = :18, row_expiry_date = :19, row_quality = :20 where obligation_id = :22")
	if err != nil {
		return err
	}

	oblig_deduct_calc.Row_changed_date = formatDate(oblig_deduct_calc.Row_changed_date)
	oblig_deduct_calc.Effective_date = formatDateString(oblig_deduct_calc.Effective_date)
	oblig_deduct_calc.Expiry_date = formatDateString(oblig_deduct_calc.Expiry_date)
	oblig_deduct_calc.Row_effective_date = formatDateString(oblig_deduct_calc.Row_effective_date)
	oblig_deduct_calc.Row_expiry_date = formatDateString(oblig_deduct_calc.Row_expiry_date)






	rows, err := stmt.Exec(oblig_deduct_calc.Obligation_seq_no, oblig_deduct_calc.Deduction_id, oblig_deduct_calc.Calculation_id, oblig_deduct_calc.Active_ind, oblig_deduct_calc.Calculation_formula, oblig_deduct_calc.Calculation_method, oblig_deduct_calc.Calculation_point, oblig_deduct_calc.Contract_id, oblig_deduct_calc.Effective_date, oblig_deduct_calc.Expiry_date, oblig_deduct_calc.Ppdm_guid, oblig_deduct_calc.Remark, oblig_deduct_calc.Source, oblig_deduct_calc.Substance_id, oblig_deduct_calc.Row_changed_by, oblig_deduct_calc.Row_changed_date, oblig_deduct_calc.Row_created_by, oblig_deduct_calc.Row_effective_date, oblig_deduct_calc.Row_expiry_date, oblig_deduct_calc.Row_quality, oblig_deduct_calc.Obligation_id)
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

func PatchObligDeductCalc(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update oblig_deduct_calc set "
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

func DeleteObligDeductCalc(c *fiber.Ctx) error {
	id := c.Params("id")
	var oblig_deduct_calc dto.Oblig_deduct_calc
	oblig_deduct_calc.Obligation_id = id

	stmt, err := db.Prepare("delete from oblig_deduct_calc where obligation_id = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(oblig_deduct_calc.Obligation_id)
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



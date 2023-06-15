package apiv1

import (
	dto "github.com/DarrenMannuela/pt_gtn_bibliography/dto"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"encoding/json"
)


func GetPdenDeclineCase(c *fiber.Ctx) error {
	rows, err := db.Query("select * from pden_decline_case")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []dto.Pden_decline_case

	for rows.Next() {
		var pden_decline_case dto.Pden_decline_case
		if err := rows.Scan(&pden_decline_case.Pden_subtype, &pden_decline_case.Pden_id, &pden_decline_case.Pden_source, &pden_decline_case.Product_type, &pden_decline_case.Case_id, &pden_decline_case.Active_ind, &pden_decline_case.Case_name, &pden_decline_case.Decline_curve_type, &pden_decline_case.Decline_type, &pden_decline_case.Effective_date, &pden_decline_case.End_date, &pden_decline_case.Expiry_date, &pden_decline_case.Final_decline, &pden_decline_case.Final_rate, &pden_decline_case.Initial_decline, &pden_decline_case.Initial_rate, &pden_decline_case.Ppdm_guid, &pden_decline_case.Project_id, &pden_decline_case.Remark, &pden_decline_case.Source, &pden_decline_case.Start_date, &pden_decline_case.Volume, &pden_decline_case.Volume_ouom, &pden_decline_case.Row_changed_by, &pden_decline_case.Row_changed_date, &pden_decline_case.Row_created_by, &pden_decline_case.Row_created_date, &pden_decline_case.Row_effective_date, &pden_decline_case.Row_expiry_date, &pden_decline_case.Row_quality); err != nil {
			return err // Exit if we get an error
		}

		// Append pden_decline_case to result
		result = append(result, pden_decline_case)
	}
	// Return result in JSON format
	return c.JSON(result)
	 
}

func SetPdenDeclineCase(c *fiber.Ctx) error {
	var pden_decline_case dto.Pden_decline_case

	setDefaults(&pden_decline_case)

	if err := json.Unmarshal(c.Body(), &pden_decline_case); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	stmt, err := db.Prepare("insert into pden_decline_case values(:1, :2, :3, :4, :5, :6, :7, :8, :9, :10, :11, :12, :13, :14, :15, :16, :17, :18, :19, :20, :21, :22, :23, :24, :25, :26, :27, :28, :29, :30)")
	if err != nil {
		return err
	}
	pden_decline_case.Row_created_date = formatDate(pden_decline_case.Row_created_date)
	pden_decline_case.Row_changed_date = formatDate(pden_decline_case.Row_changed_date)
	pden_decline_case.Effective_date = formatDateString(pden_decline_case.Effective_date)
	pden_decline_case.End_date = formatDateString(pden_decline_case.End_date)
	pden_decline_case.Expiry_date = formatDateString(pden_decline_case.Expiry_date)
	pden_decline_case.Start_date = formatDateString(pden_decline_case.Start_date)
	pden_decline_case.Row_effective_date = formatDateString(pden_decline_case.Row_effective_date)
	pden_decline_case.Row_expiry_date = formatDateString(pden_decline_case.Row_expiry_date)






	rows, err := stmt.Exec(pden_decline_case.Pden_subtype, pden_decline_case.Pden_id, pden_decline_case.Pden_source, pden_decline_case.Product_type, pden_decline_case.Case_id, pden_decline_case.Active_ind, pden_decline_case.Case_name, pden_decline_case.Decline_curve_type, pden_decline_case.Decline_type, pden_decline_case.Effective_date, pden_decline_case.End_date, pden_decline_case.Expiry_date, pden_decline_case.Final_decline, pden_decline_case.Final_rate, pden_decline_case.Initial_decline, pden_decline_case.Initial_rate, pden_decline_case.Ppdm_guid, pden_decline_case.Project_id, pden_decline_case.Remark, pden_decline_case.Source, pden_decline_case.Start_date, pden_decline_case.Volume, pden_decline_case.Volume_ouom, pden_decline_case.Row_changed_by, pden_decline_case.Row_changed_date, pden_decline_case.Row_created_by, pden_decline_case.Row_created_date, pden_decline_case.Row_effective_date, pden_decline_case.Row_expiry_date, pden_decline_case.Row_quality)
	
	if err != nil {
		return err
	}

	return c.JSON(rows)
}

func UpdatePdenDeclineCase(c *fiber.Ctx) error {
	var pden_decline_case dto.Pden_decline_case

	setDefaults(&pden_decline_case)

	if err := json.Unmarshal(c.Body(), &pden_decline_case); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	id := c.Params("id")
	pden_decline_case.Pden_subtype = id

    if pden_decline_case.Row_created_date != nil {
        return c.Status(400).SendString("Cannot update row_created_date")
    }

	stmt, err := db.Prepare("update pden_decline_case set pden_id = :1, pden_source = :2, product_type = :3, case_id = :4, active_ind = :5, case_name = :6, decline_curve_type = :7, decline_type = :8, effective_date = :9, end_date = :10, expiry_date = :11, final_decline = :12, final_rate = :13, initial_decline = :14, initial_rate = :15, ppdm_guid = :16, project_id = :17, remark = :18, source = :19, start_date = :20, volume = :21, volume_ouom = :22, row_changed_by = :23, row_changed_date = :24, row_created_by = :25, row_effective_date = :26, row_expiry_date = :27, row_quality = :28 where pden_subtype = :30")
	if err != nil {
		return err
	}

	pden_decline_case.Row_changed_date = formatDate(pden_decline_case.Row_changed_date)
	pden_decline_case.Effective_date = formatDateString(pden_decline_case.Effective_date)
	pden_decline_case.End_date = formatDateString(pden_decline_case.End_date)
	pden_decline_case.Expiry_date = formatDateString(pden_decline_case.Expiry_date)
	pden_decline_case.Start_date = formatDateString(pden_decline_case.Start_date)
	pden_decline_case.Row_effective_date = formatDateString(pden_decline_case.Row_effective_date)
	pden_decline_case.Row_expiry_date = formatDateString(pden_decline_case.Row_expiry_date)






	rows, err := stmt.Exec(pden_decline_case.Pden_id, pden_decline_case.Pden_source, pden_decline_case.Product_type, pden_decline_case.Case_id, pden_decline_case.Active_ind, pden_decline_case.Case_name, pden_decline_case.Decline_curve_type, pden_decline_case.Decline_type, pden_decline_case.Effective_date, pden_decline_case.End_date, pden_decline_case.Expiry_date, pden_decline_case.Final_decline, pden_decline_case.Final_rate, pden_decline_case.Initial_decline, pden_decline_case.Initial_rate, pden_decline_case.Ppdm_guid, pden_decline_case.Project_id, pden_decline_case.Remark, pden_decline_case.Source, pden_decline_case.Start_date, pden_decline_case.Volume, pden_decline_case.Volume_ouom, pden_decline_case.Row_changed_by, pden_decline_case.Row_changed_date, pden_decline_case.Row_created_by, pden_decline_case.Row_effective_date, pden_decline_case.Row_expiry_date, pden_decline_case.Row_quality, pden_decline_case.Pden_subtype)
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

func PatchPdenDeclineCase(c *fiber.Ctx) error {
	id := c.Params("id")

	updatedFields := make(map[string]interface{})

	if err := c.BodyParser(&updatedFields); err != nil {
		return err
	}
	if _, ok := updatedFields["row_created_date"]; ok {
        return c.Status(400).SendString("Cannot update row_created_date field")
    }

	query := "update pden_decline_case set "
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
		} else if key == "effective_date" || key == "end_date" || key == "expiry_date" || key == "start_date" || key == "row_effective_date" || key == "row_expiry_date"      {
			if date, ok := value.(string); ok && len(date) > 0 {
				formattedDate := formatDateString(&date)
				value = formattedDate
			}
		}
		values = append(values, value)
	}
	query += " where pden_subtype = :id"

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

func DeletePdenDeclineCase(c *fiber.Ctx) error {
	id := c.Params("id")
	var pden_decline_case dto.Pden_decline_case
	pden_decline_case.Pden_subtype = id

	stmt, err := db.Prepare("delete from pden_decline_case where pden_subtype = :1")
	if err != nil {
		return err
	}

	rows, err := stmt.Exec(pden_decline_case.Pden_subtype)
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


